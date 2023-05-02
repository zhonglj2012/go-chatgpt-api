package api

import (
	"github.com/gin-gonic/gin"
)

const defaultErrorMessageKey = "errorMessage"

const (
	ChatGPTUrl         = "https://chat.openai.com/chat"
	ChatGPTWelcomeText = "API is ready to provider cookies."
	ChatGPTTitleText   = "ChatGPT"

	RefreshEveryMinutes = 1
)

func ReturnMessage(msg string) gin.H {
	return gin.H{
		defaultErrorMessageKey: msg,
	}
}
