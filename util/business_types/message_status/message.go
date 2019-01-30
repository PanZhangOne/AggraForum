package message_status

// 全局消息业务类型
const (
	SYSTEM = iota
	ACTIVITY
)

// 消息状态
const (
	MessageUnread = iota
	MessageRead   = iota
)
