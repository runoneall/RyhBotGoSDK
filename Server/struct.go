package server

type ServerHandler struct {
	PreCall             func(data interface{}) interface{}
	NormalMessage       func(data interface{})
	CommandMessage      func(data interface{})
	BotFollowMessage    func(data interface{})
	BotUnfollowMessage  func(data interface{})
	GroupJoinMessage    func(data interface{})
	GroupLeaveMessage   func(data interface{})
	ButtonReportMessage func(data interface{})
	AllTypeMessage      func(data interface{})
}
