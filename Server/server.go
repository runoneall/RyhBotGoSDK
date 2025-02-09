package server

func PreCallDefaultHandler(data interface{}) interface{} {
	return data
}
func MessageDefaultHandler(data interface{}) {}

func (sh *ServerHandler) Init() {
	sh.PreCall = PreCallDefaultHandler
	sh.NormalMessage = MessageDefaultHandler
	sh.CommandMessage = MessageDefaultHandler
	sh.BotFollowMessage = MessageDefaultHandler
	sh.BotUnfollowMessage = MessageDefaultHandler
	sh.GroupJoinMessage = MessageDefaultHandler
	sh.GroupLeaveMessage = MessageDefaultHandler
	sh.ButtonReportMessage = MessageDefaultHandler
	sh.AllTypeMessage = MessageDefaultHandler
}

func (sh *ServerHandler) SetPreCallHandler(handler func(data interface{}) interface{}) {
	sh.PreCall = handler
}

func (sh *ServerHandler) SetNormalMessageHandler(handler func(data interface{})) {
	sh.NormalMessage = handler
}

func (sh *ServerHandler) SetCommandMessageHandler(handler func(data interface{})) {
	sh.CommandMessage = handler
}

func (sh *ServerHandler) SetBotFollowMessageHandler(handler func(data interface{})) {
	sh.BotFollowMessage = handler
}

func (sh *ServerHandler) SetBotUnfollowMessageHandler(handler func(data interface{})) {
	sh.BotUnfollowMessage = handler
}

func (sh *ServerHandler) SetGroupJoinMessageHandler(handler func(data interface{})) {
	sh.GroupJoinMessage = handler
}

func (sh *ServerHandler) SetGroupLeaveMessageHandler(handler func(data interface{})) {
	sh.GroupLeaveMessage = handler
}

func (sh *ServerHandler) SetButtonReportMessageHandler(handler func(data interface{})) {
	sh.ButtonReportMessage = handler
}

func (sh *ServerHandler) SetAllTypeMessageHandler(handler func(data interface{})) {
	sh.AllTypeMessage = handler
}
