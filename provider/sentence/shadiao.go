package sentence

import (
	"encoding/json"
	"strings"

	"github.com/bogeia/weather-sms/provider/config"
	"github.com/bogeia/weather-sms/provider/models"
	"github.com/pkg/errors"
)

func GetShaDiaoInfo(conf config.Sentence) (str string, err error) {
	resp, err := httpGetDoTimeout(conf.ShaDiaoHost, conf.Timeout)
	if err != nil {
		return "", errors.Wrap(err, "httpGetDoTimeout failed")
	}

	tem := new(models.ShaDiao)
	if err = json.Unmarshal(resp, &tem); err != nil {
		return "", errors.Wrap(err, "Unmarshal failed")
	}

	if tem.Data.Text != "" {
		str = strings.TrimSpace(tem.Data.Text)
	}
	return
}
