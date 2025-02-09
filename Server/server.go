package server

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func PreCallDefaultHandler(data YunhuMessage) YunhuMessage {
	return data
}
func MessageDefaultHandler(data interface{}) {}

func (sh *ServerHandler) Init() {
	sh.PreCall = PreCallDefaultHandler
	sh.NormalMessage = MessageDefaultHandler
	sh.CommandMessage = MessageDefaultHandler
	sh.BotFollowMessage = MessageDefaultHandler
	sh.BotUnfollowMessage = MessageDefaultHandler
	sh.GroupJoinMessage = MessageDefaultHandler
	sh.GroupLeaveMessage = MessageDefaultHandler
	sh.ButtonReportMessage = MessageDefaultHandler
	sh.BotSettingMessage = MessageDefaultHandler
	sh.AllTypeMessage = MessageDefaultHandler
}

func (sh *ServerHandler) SetPreCallHandler(handler func(data YunhuMessage) YunhuMessage) {
	sh.PreCall = handler
}

func (sh *ServerHandler) SetNormalMessageHandler(handler func(data interface{})) {
	sh.NormalMessage = handler
}

func (sh *ServerHandler) SetCommandMessageHandler(handler func(data interface{})) {
	sh.CommandMessage = handler
}

func (sh *ServerHandler) SetBotFollowMessageHandler(handler func(data interface{})) {
	sh.BotFollowMessage = handler
}

func (sh *ServerHandler) SetBotUnfollowMessageHandler(handler func(data interface{})) {
	sh.BotUnfollowMessage = handler
}

func (sh *ServerHandler) SetGroupJoinMessageHandler(handler func(data interface{})) {
	sh.GroupJoinMessage = handler
}

func (sh *ServerHandler) SetGroupLeaveMessageHandler(handler func(data interface{})) {
	sh.GroupLeaveMessage = handler
}

func (sh *ServerHandler) SetButtonReportMessageHandler(handler func(data interface{})) {
	sh.ButtonReportMessage = handler
}

func (sh *ServerHandler) SetBotSettingMessageHandler(handler func(data interface{})) {
	sh.BotSettingMessage = handler
}

func (sh *ServerHandler) SetAllTypeMessageHandler(handler func(data interface{})) {
	sh.AllTypeMessage = handler
}

func (sh *ServerHandler) Start(port int) {
	servYunhuHandler := func(resp http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodPost {
			http.Error(resp, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		req_body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer req.Body.Close()

		fmt.Printf("* %s POST %s ON %s\n", req.RemoteAddr, req.URL.Path, time.Now().Format("2006-01-02 15:04:05"))
		WriteRequestToFile(req_body)

		yh := sh.PreCall(ParseYunhuMessage(req_body))
		yh_header := yh.Header
		switch yh_header.EventType {
		case "message.receive.normal":
			yh_event := ParseYunhuMessageEventNormal(yh.Event)
			sh.NormalMessage(yh_event)
		case "message.receive.instruction":
			yh_event := ParseYunhuMessageEventNormal(yh.Event)
			sh.CommandMessage(yh_event)
		case "bot.followed":
			yh_event := ParseYunhuMessageEventBotFollow(yh.Event)
			sh.BotFollowMessage(yh_event)
		case "bot.unfollowed":
			yh_event := ParseYunhuMessageEventBotFollow(yh.Event)
			sh.BotUnfollowMessage(yh_event)
		case "bot.setting":
			yh_event := ParseYunhuMessageEventBotSetting(yh.Event)
			sh.BotSettingMessage(yh_event)
		case "group.join":
			yh_event := ParseYunhuMessageEventGroupJoin(yh.Event)
			sh.GroupJoinMessage(yh_event)
		case "group.leave":
			yh_event := ParseYunhuMessageEventGroupLeave(yh.Event)
			sh.GroupLeaveMessage(yh_event)
		case "button.report.inline":
			yh_event := ParseYunhuMessageEventButtonReport(yh.Event)
			sh.ButtonReportMessage(yh_event)
		}
		sh.AllTypeMessage(yh.Event)

	}
	http.HandleFunc("/", servYunhuHandler)
	portstring := fmt.Sprintf(":%d", port)
	fmt.Printf("* Starting server on %s\n", portstring)
	http.ListenAndServe(portstring, nil)
}
