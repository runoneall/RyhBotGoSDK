package message

import "fmt"

var msgBaseApi string = "https://chat-go.jwzhd.com/open-apis/v1/bot/"
var msgSendApi string = msgBaseApi + "send"
var msgBatchSendApi string = msgBaseApi + "batch_send"
var msgEditApi string = msgBaseApi + "edit"
var msgRecallApi string = msgBaseApi + "recall"
var msgGetApi string = msgBaseApi + "messages"
var msgUserBoardApi string = msgBaseApi + "board"
var msgCancelUserBoardApi string = msgBaseApi + "board-dismiss"
var msgGlobalBoardApi string = msgBaseApi + "board-all"
var msgCancelGlobalBoardApi string = msgBaseApi + "board-all-dismiss"
var imgUploadApi string = "https://chat-go.jwzhd.com/open-apis/v1/image/upload"
var NoButton [][]ButtonItem = make([][]ButtonItem, 0)
var NoParent string = ""
var NoBoardExpire int64 = 0

func SetToken(t string) {
	msgSendApi += "?token=" + t
	msgBatchSendApi += "?token=" + t
	msgEditApi += "?token=" + t
	msgRecallApi += "?token=" + t
	msgGetApi += "?token=" + t
	msgUserBoardApi += "?token=" + t
	msgCancelUserBoardApi += "?token=" + t
	msgGlobalBoardApi += "?token=" + t
	msgCancelGlobalBoardApi += "?token=" + t
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

func BatchSendText(recvIds []string, recvType string, text string, buttons [][]ButtonItem) BatchSendMessageResponse {
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

func BatchSendMarkdown(recvIds []string, recvType string, text string, buttons [][]ButtonItem) BatchSendMessageResponse {
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

func BatchSendImage(recvIds []string, recvType string, image []byte, buttons [][]ButtonItem) BatchSendMessageResponse {
	body, _ := http_imageUpload(imgUploadApi, image)
	imgUploadResp := ParseImageUploadResponse(body)
	m := BatchSendMessage{
		RecvIds:     recvIds,
		RecvType:    recvType,
		ContentType: "image",
		Content: SendImageContent{
			ImageKey: imgUploadResp.Data.ImageKey,
			Buttons:  buttons,
		},
	}
	body, _ = http_post(msgBatchSendApi, m)
	return ParseBatchSendMessageResponse(body)
}

func BatchSendHtml(recvIds []string, recvType string, text string, buttons [][]ButtonItem) BatchSendMessageResponse {
	m := BatchSendMessage{
		RecvIds:     recvIds,
		RecvType:    recvType,
		ContentType: "html",
		Content: SendMessageContent{
			Text:    text,
			Buttons: buttons,
		},
	}
	body, _ := http_post(msgBatchSendApi, m)
	return ParseBatchSendMessageResponse(body)
}

func EditText(msgId string, recvId string, recvType string, text string, buttons [][]ButtonItem) EditMessageResponse {
	m := EditMessage{
		MsgId:       msgId,
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: "text",
		Content: SendMessageContent{
			Text:    text,
			Buttons: buttons,
		},
	}
	body, _ := http_post(msgEditApi, m)
	return ParseEditMessageResponse(body)
}

func EditMarkdown(msgId string, recvId string, recvType string, text string, buttons [][]ButtonItem) EditMessageResponse {
	m := EditMessage{
		MsgId:       msgId,
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: "markdown",
		Content: SendMessageContent{
			Text:    text,
			Buttons: buttons,
		},
	}
	body, _ := http_post(msgEditApi, m)
	return ParseEditMessageResponse(body)
}

func Delete(msgId string, recvId string, recvType string) DeleteMessageResponse {
	m := DeleteMessage{
		MsgId:    msgId,
		ChatId:   recvId,
		ChatType: recvType,
	}
	body, _ := http_post(msgRecallApi, m)
	return ParseDeleteMessageResponse(body)
}

func GetBeforeMessage(chatId string, chatType string, limit int64) GetMessageResponse {
	url := fmt.Sprintf("%s&chat-id=%s&chat-type=%s&before=%d", msgGetApi, chatId, chatType, limit)
	body, _ := http_get(url)
	return ParseGetMessageResponse(body)
}

func GetAfterMessage(chatId string, chatType string, msgId string, limit int64) GetMessageResponse {
	url := fmt.Sprintf("%s&chat-id=%s&chat-type=%s&message-id=%s&after=%d", msgGetApi, chatId, chatType, msgId, limit-1)
	body, _ := http_get(url)
	return ParseGetMessageResponse(body)
}

func UserTextBoard(recvId string, recvType string, text string, expireTime int64) SetBoardResponse {
	m := SetUserBoard{
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: "text",
		Content:     text,
		ExpireTime:  expireTime,
	}
	body, _ := http_post(msgUserBoardApi, m)
	return ParseSetBoardResponse(body)
}

func UserMarkdownBoard(recvId string, recvType string, text string, expireTime int64) SetBoardResponse {
	m := SetUserBoard{
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: "markdown",
		Content:     text,
		ExpireTime:  expireTime,
	}
	body, _ := http_post(msgUserBoardApi, m)
	return ParseSetBoardResponse(body)
}

func UserHtmlBoard(recvId string, recvType string, text string, expireTime int64) SetBoardResponse {
	m := SetUserBoard{
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: "html",
		Content:     text,
		ExpireTime:  expireTime,
	}
	body, _ := http_post(msgUserBoardApi, m)
	return ParseSetBoardResponse(body)
}

func GlobalTextBoard(content string, expireTime int64) SetBoardResponse {
	m := SetGlobalBoard{
		ContentType: "text",
		Content:     content,
		ExpireTime:  expireTime,
	}
	body, _ := http_post(msgGlobalBoardApi, m)
	return ParseSetBoardResponse(body)
}

func GlobalMarkdownBoard(content string, expireTime int64) SetBoardResponse {
	m := SetGlobalBoard{
		ContentType: "markdown",
		Content:     content,
		ExpireTime:  expireTime,
	}
	body, _ := http_post(msgGlobalBoardApi, m)
	return ParseSetBoardResponse(body)
}

func GlobalHtmlBoard(content string, expireTime int64) SetBoardResponse {
	m := SetGlobalBoard{
		ContentType: "html",
		Content:     content,
		ExpireTime:  expireTime,
	}
	body, _ := http_post(msgGlobalBoardApi, m)
	return ParseSetBoardResponse(body)
}

func CancelUserBoard(recvId string, recvType string) CancelBoardResponse {
	m := CancelBoard{
		RecvId:   recvId,
		RecvType: recvType,
	}
	body, _ := http_post(msgCancelUserBoardApi, m)
	return ParseCancelBoardResponse(body)
}

func CancelGlobalBoard() CancelBoardResponse {
	body, _ := http_post(msgCancelGlobalBoardApi, CancelBoard{})
	return ParseCancelBoardResponse(body)
}
