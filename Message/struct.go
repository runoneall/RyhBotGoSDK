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

type UploadImageResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ImageKey string `json:"imageKey"`
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
