package server

type ServerHandler struct {
	PreCall             func(data YunhuMessage) YunhuMessage
	NormalMessage       func(data YunhuMessageEventNormal)
	CommandMessage      func(data YunhuMessageEventNormal)
	BotFollowMessage    func(data YunhuMessageEventBotFollow)
	BotUnfollowMessage  func(data YunhuMessageEventBotUnfollow)
	GroupJoinMessage    func(data YunhuMessageEventGroupJoin)
	GroupLeaveMessage   func(data YunhuMessageEventGroupLeave)
	ButtonReportMessage func(data YunhuMessageEventButtonReport)
	BotSettingMessage   func(data YunhuMessageEventBotSetting)
	AllTypeMessage      func(data interface{})
}

type YunhuMessage struct {
	Version string `json:"version"`
	Header  struct {
		EventId   string `json:"eventId"`
		EventTime int64  `json:"eventTime"`
		EventType string `json:"eventType"`
	} `json:"header"`
	Event interface{} `json:"event"`
}

type YunhuMessageEventNormal struct {
	Chat struct {
		ChatId   string `json:"chatId"`
		ChatType string `json:"chatType"`
	} `json:"chat"`
	Sender struct {
		SenderId        string `json:"senderId"`
		SenderType      string `json:"senderType"`
		SenderUserLevel string `json:"senderUserLevel"`
		SenderNickname  string `json:"senderNickname"`
	} `json:"sender"`
	Message struct {
		MsgId           string                 `json:"msgId"`
		ParentId        string                 `json:"parentId"`
		ContentType     string                 `json:"contentType"`
		Content         map[string]interface{} `json:"content"`
		InstructionId   int64                  `json:"instructionId"`
		InstructionName string                 `json:"instructionName"`
		CommandId       int64                  `json:"commandId"`
		CommandName     string                 `json:"commandName"`
	} `json:"message"`
}

type YunhuMessageEventBotFollow struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

type YunhuMessageEventBotUnfollow struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

type YunhuMessageEventGroupJoin struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

type YunhuMessageEventGroupLeave struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

type YunhuMessageEventButtonReport struct {
	MsgId    string `json:"msgId"`
	RecvId   string `json:"recvId"`
	RecvType string `json:"recvType"`
	UserId   string `json:"userId"`
	Value    string `json:"value"`
	Time     int64  `json:"time"`
}

type YunhuMessageEventBotSetting struct {
	Time        int64                  `json:"time"`
	ChatId      string                 `json:"chatId"`
	ChatType    string                 `json:"chatType"`
	GroupId     string                 `json:"groupId"`
	GroupName   string                 `json:"groupName"`
	AvatarUrl   string                 `json:"avatarUrl"`
	SettingJson map[string]interface{} `json:"settingJson"`
}
