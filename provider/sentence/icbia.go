package sentence

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bogeia/weather-sms/provider/config"
	"github.com/bogeia/weather-sms/provider/models"
	"github.com/pkg/errors"
)

func GetIcBiaInfo(conf config.Sentence) (str string, err error) {
	resp, err := httpGetDoTimeout(conf.ICiBaURL, conf.Timeout)
	if err != nil {
		return "", errors.Wrap(err, "httpGetDoTimeout failed")
	}

	tem := new(models.Icbia)
	if err = json.Unmarshal(resp, &tem); err != nil {
		return "", errors.Wrap(err, "Unmarshal failed")
	}

	tem.Content = strings.TrimSpace(tem.Content)
	tem.Note = strings.TrimSpace(tem.Note)
	if tem.Content != "" && tem.Note != "" {
		str = fmt.Sprintf("%s\n%s", tem.Content, tem.Note)
	}
	return
}
