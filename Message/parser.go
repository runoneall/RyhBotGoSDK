package message

import "encoding/json"

func ParseSendMessageResponse(body []byte) SendMessageResponse {
	var resp SendMessageResponse
	json.Unmarshal(body, &resp)
	return resp
}

func ParseImageUploadResponse(body []byte) UploadImageResponse {
	var resp UploadImageResponse
	json.Unmarshal(body, &resp)
	return resp
}

func ParseBatchSendMessageResponse(body []byte) BatchSendMessageResponse {
	var resp BatchSendMessageResponse
	json.Unmarshal(body, &resp)
	return resp
}

func ParseEditMessageResponse(body []byte) EditMessageResponse {
	var resp EditMessageResponse
	json.Unmarshal(body, &resp)
	return resp
}

func ParseDeleteMessageResponse(body []byte) DeleteMessageResponse {
	var resp DeleteMessageResponse
	json.Unmarshal(body, &resp)
	return resp
}

func ParseGetMessageResponse(body []byte) GetMessageResponse {
	var resp GetMessageResponse
	json.Unmarshal(body, &resp)
	return resp
}

func ParseSetBoardResponse(body []byte) SetBoardResponse {
	var resp SetBoardResponse
	json.Unmarshal(body, &resp)
	return resp
}

func ParseCancelBoardResponse(body []byte) CancelBoardResponse {
	var resp CancelBoardResponse
	json.Unmarshal(body, &resp)
	return resp
}
