package weather

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bogeia/weather-sms/provider/config"
)

var (
	once sync.Once
	opt  string
)

// ParseWeatherConfigToString ...
func ParseWeatherConfigToString(conf config.Weather) string {
	once.Do(func() {
		var info strings.Builder

		info.WriteString(fmt.Sprintf("%s/", conf.Host))
		info.WriteString(fmt.Sprintf("%s/", conf.Version))
		info.WriteString(fmt.Sprintf("%s/", conf.Token))
		info.WriteString(fmt.Sprintf("%s", conf.Location))

		opt = info.String()
	})
	return opt
}
