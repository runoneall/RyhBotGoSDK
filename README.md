# RyhBotGoSDK

## Install
- Init project(if not): `go mod init project_name`
- Clone this repository: `git clone https://github.com/runoneall/RyhBotGoSDK.git`
- Add `replace ryhbotgosdk => ./RyhBotGoSDK` to `go.mod` file
- Run `go mod tidy`

## Import
```go
import (
	"ryhbotgosdk/message"
	"ryhbotgosdk/server"
)
```