package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

func main() {
	httphandler()
}

func httphandler() {

	app := iris.New()
	party := app.Party("/")
	party.Get("login", func(c iris.Context) {
		host := "https://access.line.me/oauth2/v2.1/authorize"
		argstr := "?"
		args := map[string]string{
			"response_type": "code",
			"client_id":     "1656003757",
			"redirect_uri":  "https://linelogin-test-210519.herokuapp.com/auth",
			"state":         time.Now().Format("200601021504050000"),
			"scope":         "profile%20openid",
		}

		for i, v := range args {
			argstr += fmt.Sprintf("%s=%s&", i, v)
		}

		argstr = strings.TrimRight(argstr, "&")

		host += argstr

		c.JSON(map[string]string{
			"url": host,
		}, iris.JSON{UnescapeHTML: true})
	})

	party.Get("/auth", func(c iris.Context) {
		fmt.Println(c.Request().URL.RawQuery)
		body, _ := c.GetBody()
		fmt.Println(string(body))
	})

	app.Listen(":80")
}
