package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bogeia/weather-sms/provider/config"
	"github.com/bogeia/weather-sms/provider/weather"
)

func main() {
	conf, err := config.Load(filepath.Join("./", "conf", "config.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	conf.Server.Title = strings.ReplaceAll(conf.Server.Title, " ", "\n")
	conf.Server.End = strings.ReplaceAll(conf.Server.End, " ", "\n")
	conf.Server.Form = strings.ReplaceAll(conf.Server.Form, " ", "\n")

	info, err := weather.GetWeatherInfo(conf.Weather)
	if err != nil {
		log.Fatal(err)
	}

	message := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s",
		conf.Server.Title, info, conf.Server.End, conf.Server.Form)

	cmd := fmt.Sprintf(`osascript %s "%s" "%s"`,
		filepath.Join("script", "message.scpt"),
		conf.Server.Phone,
		message)

	if _, err = exec.Command("/bin/sh", "-c", cmd).CombinedOutput(); err != nil {
		log.Panicf("send message failed, err:(%v)", err)
	}
}
