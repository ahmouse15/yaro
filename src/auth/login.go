package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	webview "github.com/webview/webview_go"
)

// Returns code and idToken
func RunLoginWindow() string {
	wd, _ := os.Getwd()
	os.Setenv("WEBVIEW2_USER_DATA_FOLDER", wd+"/.webview")

	fmt.Println("pc sign: ", pcSign)

	const LOGIN_URL = "https://accounts.ea.com/connect/auth"

	url := LOGIN_URL + params

	window := webview.New(false)
	window.SetSize(600, 900, webview.HintNone)
	window.Navigate(url)

	var authUri string

	checkUri := func() {
		for {
			uri := window.GetAuthUri()
			if strings.HasPrefix(uri, "qrc") {
				fmt.Println("Found auth URI.")

				authUri = strings.Clone(uri)
				window.Dispatch(window.Terminate)
				window.Dispatch(window.Destroy)

				return
			}
			time.Sleep(time.Second / 2)
		}
	}
	go checkUri()

	window.Run()

	fmt.Println(authUri)
	loginCode, _ := parseAuthUri(authUri)

	return loginCode
}

func parseAuthUri(uri string) (code string, idToken string) {
	_, temp, _ := strings.Cut(uri, "#code=")
	code, _, _ = strings.Cut(temp, "&")

	_, idToken, _ = strings.Cut(uri, "id_token=")

	return
}
