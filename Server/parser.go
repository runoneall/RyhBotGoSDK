package server

import (
	"encoding/json"
)

func ParseYunhuMessage(body []byte) YunhuMessage {
	var data YunhuMessage
	json.Unmarshal(body, &data)
	return data
}

func ParseYunhuMessageEventNormal(body interface{}) YunhuMessageEventNormal {
	data, _ := json.Marshal(body)
	var data2 YunhuMessageEventNormal
	json.Unmarshal(data, &data2)
	return data2
}

func ParseYunhuMessageEventBotFollow(body interface{}) YunhuMessageEventBotFollow {
	data, _ := json.Marshal(body)
	var data2 YunhuMessageEventBotFollow
	json.Unmarshal(data, &data2)
	return data2
}

func ParseYunhuMessageEventBotUnfollow(body interface{}) YunhuMessageEventBotUnfollow {
	data, _ := json.Marshal(body)
	var data2 YunhuMessageEventBotUnfollow
	json.Unmarshal(data, &data2)
	return data2
}

func ParseYunhuMessageEventGroupJoin(body interface{}) YunhuMessageEventGroupJoin {
	data, _ := json.Marshal(body)
	var data2 YunhuMessageEventGroupJoin
	json.Unmarshal(data, &data2)
	return data2
}

func ParseYunhuMessageEventGroupLeave(body interface{}) YunhuMessageEventGroupLeave {
	data, _ := json.Marshal(body)
	var data2 YunhuMessageEventGroupLeave
	json.Unmarshal(data, &data2)
	return data2
}

func ParseYunhuMessageEventButtonReport(body interface{}) YunhuMessageEventButtonReport {
	data, _ := json.Marshal(body)
	var data2 YunhuMessageEventButtonReport
	json.Unmarshal(data, &data2)
	return data2
}

func ParseYunhuMessageEventBotSetting(body interface{}) YunhuMessageEventBotSetting {
	data, _ := json.Marshal(body)
	var data2 YunhuMessageEventBotSetting
	json.Unmarshal(data, &data2)
	return data2
}
