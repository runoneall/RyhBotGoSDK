package message

type ButtonItem struct {
	Text       string `json:"text"`
	ActionType int    `json:"actionType"`
	Url        string `json:"url"`
	Value      string `json:"value"`
}

type SendMessageContent struct {
	Text    string         `json:"text"`
	Buttons [][]ButtonItem `json:"buttons"`
}

type SendImageContent struct {
	ImageKey string         `json:"imageKey"`
	Buttons  [][]ButtonItem `json:"buttons"`
}

type UploadImageResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ImageKey string `json:"imageKey"`
	} `json:"data"`
}

type SendMessage struct {
	RecvId      string      `json:"recvId"`
	RecvType    string      `json:"recvType"`
	ContentType string      `json:"contentType"`
	Content     interface{} `json:"content"`
	ParentId    string      `json:"parentId"`
}

type SendMessageResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		MessageInfo struct {
			MsgId    string `json:"msgId"`
			RecvId   string `json:"recvId"`
			RecvType string `json:"recvType"`
		} `json:"messageInfo"`
	} `json:"data"`
}

type BatchSendMessage struct {
	RecvIds     []string    `json:"recvIds"`
	RecvType    string      `json:"recvType"`
	ContentType string      `json:"contentType"`
	Content     interface{} `json:"content"`
}

type BatchSendMessageResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		SuccessCount int64 `json:"successCount"`
		SuccessList  []struct {
			MsgId    string `json:"msgId"`
			RecvId   string `json:"recvId"`
			RecvType string `json:"recvType"`
		} `json:"successList"`
	} `json:"data"`
}

type EditMessage struct {
	MsgId       string      `json:"msgId"`
	RecvId      string      `json:"recvId"`
	RecvType    string      `json:"recvType"`
	ContentType string      `json:"contentType"`
	Content     interface{} `json:"content"`
}

type EditMessageResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		SuccessCount int64 `json:"successCount"`
	} `json:"data"`
}

type DeleteMessage struct {
	MsgId    string `json:"msgId"`
	ChatId   string `json:"chatId"`
	ChatType string `json:"chatType"`
}

type DeleteMessageResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type GetMessageResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []struct {
			MsgId          string                 `json:"msgId"`
			ParentId       string                 `json:"parentId"`
			SenderId       string                 `json:"senderId"`
			SenderType     string                 `json:"senderType"`
			SenderNickname string                 `json:"senderNickname"`
			ContentType    string                 `json:"contentType"`
			Content        map[string]interface{} `json:"content"`
			SendTime       int64                  `json:"sendTime"`
			CommandName    string                 `json:"commandName"`
			CommandId      int64                  `json:"commandId"`
		} `json:"list"`
		Total int64 `json:"total"`
	} `json:"data"`
}
