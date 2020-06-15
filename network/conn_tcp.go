// /////////////////////////////////////////////////////////////////////////////
// tcp 客户端

package network

import (
	"encoding/json"
	"net"

	"github.com/zpab123/sco/protocol"
	"github.com/zpab123/sco/state"
	"github.com/zpab123/zaplog"
)

// /////////////////////////////////////////////////////////////////////////////
// TcpConn

// tcp 客户端
type TcpConn struct {
	addr     string              // 远端地址
	socket   *Socket             // socket
	stateMgr *state.StateManager // 状态管理
	handler  IClientHandler      // 消息处理
}

// 新建1个 tcp 连接
func NewTcpConn(addr string) *TcpConn {
	// 状态管理
	st := state.NewStateManager()

	c := TcpConn{
		addr:     addr,
		stateMgr: st,
	}
	// 设置为初始化状态
	c.stateMgr.SetState(C_CLI_ST_INIT)

	return &c
}

// 启动
func (this *TcpConn) Run() error {
	conn, err := net.Dial("tcp", this.addr)
	if nil != err {
		return err
	}

	s, err := NewSocket(conn)
	if nil != err {
		return err
	}

	this.socket = s

	// 发送线程
	go this.sendLoop()
	// 心跳？
	// 接收线程
	go this.recvLoop() // 接收循环

	return nil
}

// 停止
func (this *TcpConn) Stop() {
	this.socket.Close()
}

// 发送1个 packet 消息
func (this *TcpConn) SendPacket(pkt *Packet) error {
	// 状态效验
	if this.stateMgr.GetState() != C_CLI_ST_WORKING {
		return errState
	}

	this.socket.SendPacket(pkt)

	return nil
}

// 发送 []byte
func (this *TcpConn) SendBytes(bytes []byte) error {
	// 状态效验
	if this.stateMgr.GetState() != C_CLI_ST_WORKING {
		return errState
	}

	this.socket.SendBytes(bytes)

	return nil
}

// 打印信息
func (this *TcpConn) String() string {
	return this.socket.String()
}

// 设置处理器
func (this *TcpConn) SetHandler(h IClientHandler) {
	if nil != h {
		this.handler = h
	}
}

// 接收线程
func (this *TcpConn) recvLoop() {
	defer func() {
		zaplog.Debugf("[TcpConn] recvLoop 结束")
		this.socket.SendPacket(nil) // 用于结束 sendLoop
	}()

	for {
		pkt, err := this.socket.RecvPacket()
		if nil != err {
			return
		}

		if nil != pkt {
			this.onPacket(pkt)
			continue
		}
	}
}

// 发送线程
func (this *TcpConn) sendLoop() {
	defer func() {
		zaplog.Debugf("[TcpConn] sendLoop 结束")
		this.Stop()
	}()

	for {
		err := this.socket.Flush()
		if nil != err {
			break
		}
	}
}

// 发送握手请求
func (this *TcpConn) reqHandShake() {
	req := protocol.HandshakeReq{
		Key:      "sco",
		Acceptor: 1,
	}

	data, err := json.Marshal(&req)
	if nil != err {
		zaplog.Debugf("[TcpConn] 编码握手消息失败")
		this.Stop()
		return
	}

	pkt := NewPacket(protocol.C_MID_HANDSHAKE)
	pkt.AppendBytes(data)

	this.socket.SendBytes(pkt.Data())
}

// 发送握手ack
func (this *TcpConn) sendAck() {
	pkt := NewPacket(protocol.C_MID_HANDSHAKE_ACK)
	this.socket.SendBytes(pkt.Data())

	this.stateMgr.SetState(C_CLI_ST_WORKING)
}

// 收到1个 pakcet
func (this *TcpConn) onPacket(pkt *Packet) {
	switch pkt.mid {
	case protocol.C_MID_HANDSHAKE: // 远端握手结果
		this.onHandshake(pkt.GetBody())
	case protocol.C_MID_HEARTBEAT: // 心跳
		//
	default:
		this.handle(pkt) // 处理
	}
}

// 握手结果
func (this *TcpConn) onHandshake(data []byte) {
	res := protocol.HandshakeRes{}
	err := json.Unmarshal(data, &res)
	if nil != err {
		zaplog.Debugf("[TcpConn] 解码握手结果失败")
		this.Stop()
		return
	}

	if res.Code == protocol.C_CODE_OK {
		this.sendAck()
	}
}

// 需要处理的消息
func (this *TcpConn) handle(pkt *Packet) {
	if this.stateMgr.GetState() != C_CLI_ST_WORKING {
		this.Stop()
		return
	}

	if nil != this.handler {
		this.handler.OnPacket(this, pkt)
	}
}
