package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/go-chatgpt-api/api"
	"github.com/linweiyuan/go-chatgpt-api/webdriver"
)

//goland:noinspection GoUnhandledErrorResult
func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		cookies, err := webdriver.WebDriver.GetCookies()
		if err != nil {
			c.JSON(http.StatusInternalServerError, api.ReturnMessage(err.Error()))
			return
		}

		responseMap := make(map[string]string)
		for _, cookie := range cookies {
			responseMap[cookie.Name] = cookie.Value
		}
		c.JSON(http.StatusOK, responseMap)
	})

	router.Run(":8080")
}
