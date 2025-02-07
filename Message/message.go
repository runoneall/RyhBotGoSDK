package message

var msgBaseApi string = "https://chat-go.jwzhd.com/open-apis/v1/bot/"
var msgSendApi string = msgBaseApi + "send"
var msgBatchSendApi string = msgBaseApi + "batch_send"
var msgEditApi string = msgBaseApi + "edit"
var msgRecallApi string = msgBaseApi + "recall"
var msgGetApi string = msgBaseApi + "messages"
var imgUploadApi string = "https://chat-go.jwzhd.com/open-apis/v1/image/upload"
var NoButton = make([][]ButtonItem, 0)
var NoParent = ""

func SetToken(t string) {
	msgSendApi += "?token=" + t
	msgBatchSendApi += "?token=" + t
	msgEditApi += "?token=" + t
	msgRecallApi += "?token=" + t
	msgGetApi += "?token=" + t
	imgUploadApi += "?token=" + t
}

func SendText(recvId string, recvType string, text string, buttons [][]ButtonItem, parentId string) SendMessageResponse {
	m := SendMessage{
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: "text",
		Content: SendMessageContent{
			Text:    text,
			Buttons: buttons,
		},
		ParentId: parentId,
	}
	body, _ := http_post(msgSendApi, m)
	return ParseSendMessageResponse(body)
}

func SendMarkdown(recvId string, recvType string, text string, buttons [][]ButtonItem, parentId string) SendMessageResponse {
	m := SendMessage{
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: "markdown",
		Content: SendMessageContent{
			Text:    text,
			Buttons: buttons,
		},
		ParentId: parentId,
	}
	body, _ := http_post(msgSendApi, m)
	return ParseSendMessageResponse(body)
}

func SendImage(recvId string, recvType string, image []byte, buttons [][]ButtonItem, parentId string) SendMessageResponse {
	body, _ := http_imageUpload(imgUploadApi, image)
	imgUploadResp := ParseImageUploadResponse(body)
	m := SendMessage{
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: "image",
		Content: SendImageContent{
			ImageKey: imgUploadResp.Data.ImageKey,
			Buttons:  buttons,
		},
		ParentId: parentId,
	}
	body, _ = http_post(msgSendApi, m)
	return ParseSendMessageResponse(body)
}

func SendHtml(recvId string, recvType string, text string, buttons [][]ButtonItem, parentId string) SendMessageResponse {
	m := SendMessage{
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: "html",
		Content: SendMessageContent{
			Text:    text,
			Buttons: buttons,
		},
		ParentId: parentId,
	}
	body, _ := http_post(msgSendApi, m)
	return ParseSendMessageResponse(body)
}

func BatchSendText(recvIds []string, recvType string, text string, buttons [][]ButtonItem, parentId string) BatchSendMessageResponse {
	m := BatchSendMessage{
		RecvIds:     recvIds,
		RecvType:    recvType,
		ContentType: "text",
		Content: SendMessageContent{
			Text:    text,
			Buttons: buttons,
		},
	}
	body, _ := http_post(msgBatchSendApi, m)
	return ParseBatchSendMessageResponse(body)
}

func BatchSendMarkdown(recvIds []string, recvType string, text string, buttons [][]ButtonItem, parentId string) BatchSendMessageResponse {
	m := BatchSendMessage{
		RecvIds:     recvIds,
		RecvType:    recvType,
		ContentType: "markdown",
		Content: SendMessageContent{
			Text:    text,
			Buttons: buttons,
		},
	}
	body, _ := http_post(msgBatchSendApi, m)
	return ParseBatchSendMessageResponse(body)
}
