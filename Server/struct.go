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
	BotSettingMessage   func(data interface{})
	AllTypeMessage      func(data interface{})
}

// {
//     "event": {
//         "chat": {
//             "chatId": "748345168",
//             "chatType": "group"
//         },
//         "message": {
//             "chatId": "748345168",
//             "chatType": "group",
//             "commandId": 0,
//             "commandName": "",
//             "content": {
//                 "at": [
//                     "8826514"
//                 ],
//                 "parent": "Hehe Nya!: hhha",
//                 "text": "@Hehe Nya! ​test"
//             },
//             "contentType": "text",
//             "instructionId": 0,
//             "instructionName": "",
//             "msgId": "0744157a5c204a4aa56aac6ae86a5030",
//             "parentId": "9b38a340dfa440129b8d3869d647398c",
//             "sendTime": 1739089469918
//         },
//         "sender": {
//             "senderId": "9437329",
//             "senderNickname": "£″≈μ",
//             "senderType": "user",
//             "senderUserLevel": "owner"
//         }
//     },
//     "header": {
//         "eventId": "c1228d38994641efb6ba739ed71f2ae8",
//         "eventTime": 1739089469944,
//         "eventType": "message.receive.normal"
//     },
//     "version": "1.0"
// }

type YunhuMessage struct {
	Version string `json:"version"`
	Header  struct {
		EventId   string `json:"eventId"`
		EventTime int64  `json:"eventTime"`
		EventType string `json:"eventType"`
	} `json:"header"`
	Event interface{} `json:"event"`
}
