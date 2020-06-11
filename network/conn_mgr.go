// /////////////////////////////////////////////////////////////////////////////
// 连接管理

package network

import (
	"net"

	"github.com/zpab123/syncutil"
	"github.com/zpab123/zaplog"
	"golang.org/x/net/websocket"
)

// /////////////////////////////////////////////////////////////////////////////
// ConnMgr

// 连接管理
type ConnMgr struct {
	maxConn uint32                // 最大连接数量，超过此数值后，不再接收新连接
	connNum syncutil.AtomicUint32 // 当前连接数
	handler IHandler              // 消息处理器
}

// 新建1个 ConnMgr
func NewConnMgr(maxConn uint32) *ConnMgr {
	if 0 == maxConn {
		maxConn = C_CONN_MAX
	}

	mgr := ConnMgr{
		maxConn: maxConn,
	}

	return &mgr
}

// 停止连接管理
func (this *ConnMgr) Stop() {

}

// 收到1个新的 websocket 连接对象 [IWsConnManager]
func (this *ConnMgr) OnWsConn(wsconn *websocket.Conn) {
	// 超过最大连接数
	if this.connNum.Load() >= this.maxConn {
		wsconn.Close()
		zaplog.Warnf("[ConnMgr] 达到最大连接数，关闭新连接。当前连接数=%d", this.connNum.Load())
	}

	// 参数设置
	wsconn.PayloadType = websocket.BinaryFrame // 以二进制方式接受数据

	// 创建代理
	zaplog.Debugf("[ConnMgr] 新 ws 连接，ip=%s", wsconn.RemoteAddr())
	this.newAgent(wsconn, true)
}

// 收到1个新的 tcp 连接对象 [ITcpConnManager]
func (this *ConnMgr) OnTcpConn(conn net.Conn) {
	// 超过最大连接数
	if this.connNum.Load() >= this.maxConn {
		conn.Close()
		zaplog.Warnf("[ConnMgr] 达到最大连接数，关闭新连接。当前连接数=%d", this.connNum.Load())
	}

	// 创建代理
	zaplog.Debugf("[ConnMgr] 新 tcp 连接，ip=%s", conn.RemoteAddr())
	this.newAgent(conn, true)
}

// 设置 handler
func (this *ConnMgr) SetHandler(h IHandler) {
	if nil != h {
		this.handler = h
	}
}

// 创建代理
func (this *ConnMgr) newAgent(conn net.Conn, isWebSocket bool) {
	ao := NewAgentOpt()
	s := NewSocket(conn)

	a := NewAgent(s, ao)
	a.SetHandler(this.handler)
	a.Run()
}
