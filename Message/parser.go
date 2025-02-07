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
