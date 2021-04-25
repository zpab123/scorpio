// /////////////////////////////////////////////////////////////////////////////
// 消息订阅

package module

// /////////////////////////////////////////////////////////////////////////////
// Subscriber

// 消息订阅者
type Subscriber struct {
	SuberId uint32      // 订阅者id
	MsgId   uint32      // 消息id
	MsgChan chan Messge // 接收消息的通道
}

// /////////////////////////////////////////////////////////////////////////////
// OffMsg

// 取消订阅
type OffMessge struct {
	SuberId uint32 // 订阅者id
	MsgId   uint32 // 消息id
}
