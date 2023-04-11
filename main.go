package main

import (
	_ "embed"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bogeia/weather-sms/provider/config"
	"github.com/bogeia/weather-sms/provider/weather"
)

//go:embed script/imessage.scpt
var imessageScript []byte

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

	cmd := exec.Command("osascript", "-e", string(imessageScript), conf.Server.Phone, message)
	if _, err = cmd.CombinedOutput(); err != nil {
		log.Fatalf("send message failed, err:(%v)", err)
	}

	log.Println("📨 send successful")
}
