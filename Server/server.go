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
func MessageNormalDefaultHandler(data YunhuMessageEventNormal)             {}
func MessageBotFollowDefaultHandler(data YunhuMessageEventBotFollow)       {}
func MessageBotUnfollowDefaultHandler(data YunhuMessageEventBotUnfollow)   {}
func MessageGroupJoinDefaultHandler(data YunhuMessageEventGroupJoin)       {}
func MessageGroupLeaveDefaultHandler(data YunhuMessageEventGroupLeave)     {}
func MessageButtonReportDefaultHandler(data YunhuMessageEventButtonReport) {}
func MessageBotSettingDefaultHandler(data YunhuMessageEventBotSetting)     {}
func MessageAllTypeDefaultHandler(data interface{})                        {}

func (sh *ServerHandler) Init() {
	sh.PreCall = PreCallDefaultHandler
	sh.NormalMessage = MessageNormalDefaultHandler
	sh.CommandMessage = MessageNormalDefaultHandler
	sh.BotFollowMessage = MessageBotFollowDefaultHandler
	sh.BotUnfollowMessage = MessageBotUnfollowDefaultHandler
	sh.GroupJoinMessage = MessageGroupJoinDefaultHandler
	sh.GroupLeaveMessage = MessageGroupLeaveDefaultHandler
	sh.ButtonReportMessage = MessageButtonReportDefaultHandler
	sh.BotSettingMessage = MessageBotSettingDefaultHandler
	sh.AllTypeMessage = MessageAllTypeDefaultHandler
}

func (sh *ServerHandler) Start(port int) {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {

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
			yh_event := ParseYunhuMessageEventBotUnfollow(yh.Event)
			sh.BotUnfollowMessage(yh_event)
		case "group.join":
			yh_event := ParseYunhuMessageEventGroupJoin(yh.Event)
			sh.GroupJoinMessage(yh_event)
		case "group.leave":
			yh_event := ParseYunhuMessageEventGroupLeave(yh.Event)
			sh.GroupLeaveMessage(yh_event)
		case "button.report.inline":
			yh_event := ParseYunhuMessageEventButtonReport(yh.Event)
			sh.ButtonReportMessage(yh_event)
		case "bot.setting":
			yh_event := ParseYunhuMessageEventBotSetting(yh.Event)
			sh.BotSettingMessage(yh_event)
		}
		sh.AllTypeMessage(yh.Event)

	})
	portstring := fmt.Sprintf(":%d", port)
	fmt.Printf("* Starting server on %s\n", portstring)
	http.ListenAndServe(portstring, nil)
}
