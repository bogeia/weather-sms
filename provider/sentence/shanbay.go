package sentence

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bogeia/weather-sms/provider/config"
	"github.com/bogeia/weather-sms/provider/models"
	"github.com/pkg/errors"
)

func GetShanBayInfo(conf config.Sentence) (str string, err error) {
	resp, err := httpGetDoTimeout(conf.ShanBayURL, conf.Timeout)
	if err != nil {
		return "", errors.Wrap(err, "httpGetDoTimeout failed")
	}

	tem := new(models.ShanBay)
	if err = json.Unmarshal(resp, &tem); err != nil {
		return "", errors.Wrap(err, "Unmarshal failed")
	}

	tem.Content = strings.TrimSpace(tem.Content)
	tem.Translation = strings.TrimSpace(tem.Translation)
	if len(tem.Content) > 3 && len(tem.Content) > 3 {
		str = fmt.Sprintf("%s\n%s", tem.Content, tem.Translation)
	}

	return
}
