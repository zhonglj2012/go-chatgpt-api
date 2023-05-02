package webdriver

import (
	"os"
	"time"

	"github.com/linweiyuan/go-chatgpt-api/api"
	"github.com/linweiyuan/go-chatgpt-api/util/logger"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var WebDriver selenium.WebDriver

//goland:noinspection GoUnhandledErrorResult,SpellCheckingInspection
func init() {
	chatgptProxyServer := os.Getenv("CHATGPT_PROXY_SERVER")
	if chatgptProxyServer == "" {
		logger.Error("CHATGPT_PROXY_SERVER is empty")
		return
	}
	logger.Info("CHATGPT_PROXY_SERVER is: " + chatgptProxyServer)

	chromeArgs := []string{
		"--no-sandbox",
		"--disable-gpu",
		"--disable-dev-shm-usage",
		"--disable-blink-features=AutomationControlled",
		"--incognito",
		"--headless=new",
	}

	networkProxyServer := os.Getenv("NETWORK_PROXY_SERVER")
	if networkProxyServer != "" {
		logger.Info("NETWORK_PROXY_SERVER is: " + networkProxyServer)
		chromeArgs = append(chromeArgs, "--proxy-server="+networkProxyServer)
	}

	WebDriver, _ = selenium.NewRemote(selenium.Capabilities{
		"chromeOptions": chrome.Capabilities{
			Args:            chromeArgs,
			ExcludeSwitches: []string{"enable-automation"},
		},
	}, chatgptProxyServer)

	if WebDriver == nil {
		logger.Error("Please make sure chatgpt proxy service is running")
		return
	}

	WebDriver.Get(api.ChatGPTUrl)

	if isReady(WebDriver) {
		logger.Info(api.ChatGPTWelcomeText)
	} else {
		if !isAccessDenied(WebDriver) {
			if HandleCaptcha(WebDriver) {
				logger.Info(api.ChatGPTWelcomeText)
			}
		}
	}

	go func() {
		ticker := time.NewTicker(api.RefreshEveryMinutes * time.Minute)
		for {
			select {
			case <-ticker.C:
				refresh()
			}
		}
	}()
}
