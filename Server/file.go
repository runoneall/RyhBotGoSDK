package server

import (
	"encoding/json"
	"os"
)

func WriteRequestToFile(requestBody []byte) {
	var data map[string]interface{}
	json.Unmarshal(requestBody, &data)
	req_json, _ := json.MarshalIndent(data, "", "  ")
	file, err := os.Create("RecvMsg.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(req_json)
	if err != nil {
		panic(err)
	}
}
