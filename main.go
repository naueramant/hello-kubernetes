package main

import (
	"fmt"
	"os"

	"github.com/aymerick/raymond"
	"github.com/gobuffalo/packr/v2"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

var (
	box        = packr.New("assets", "./assets")
	html, logo string
)

type Config struct {
	Port int `default:"8080"`
}

func main() {
	template, err := box.FindString("index.hb")
	if err != nil {
		logrus.Fatal(err)
	}

	tpl, err := raymond.Parse(template)
	if err != nil {
		logrus.Fatal(err)
	}

	var conf Config
	if err := envconfig.Process("", &conf); err != nil {
		logrus.Fatal(err)
	}

	logo, err := box.FindString("logo.svg")
	if err != nil {
		logrus.Fatal(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		logrus.Fatal(err)
	}

	templateContext := map[string]string{
		"logo":     logo,
		"hostname": hostname,
	}

	html, err = tpl.Exec(templateContext)
	if err != nil {
		logrus.Fatal(err)
	}

	h := fasthttp.CompressHandler(requestHandler)

	logrus.Infof("Listening on port %d", conf.Port)
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", conf.Port), h); err != nil {
		logrus.Fatal(err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=UTF-8")
	fmt.Fprintf(ctx, html)
}
