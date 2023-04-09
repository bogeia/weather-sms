package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/avast/retry-go"
	"github.com/bogeia/weather-sms/provider/config"
	"github.com/bogeia/weather-sms/provider/models"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

// GetWeatherInfo get weather info
func GetWeatherInfo(conf config.Weather) (string, error) {
	realData, err := GetRealtime(conf)
	if err != nil {
		return "", errors.Wrap(err, "GetRealtime failed")
	}

	str, err := ConvertWeatherToString(realData)
	if err != nil {
		return "", errors.Wrap(err, "ConvertWeatherToString failed")
	}
	return str, nil
}

// GetRealtime  预警信息 + 实况（当时天气状况） + 未来的天气信息
func GetRealtime(conf config.Weather) (data []byte, err error) {
	err = retry.Do(
		func() error {
			url := ParseWeatherConfigToString(conf) + "/weather?alert=true&dailysteps=1&hourlysteps=24"

			req := &fasthttp.Request{}
			req.SetRequestURI(url)
			req.Header.SetMethod(http.MethodGet)
			req.Header.SetContentType("application/json; charset=utf-8")

			resp := &fasthttp.Response{}
			if err = fasthttp.DoTimeout(req, resp, conf.Timeout); err != nil {
				return errors.Wrap(err, "http do failed")
			}
			data = resp.Body()
			return nil
		})
	return
}

// ConvertWeatherToString ...
func ConvertWeatherToString(realDate []byte) (string, error) {
	r := new(models.ResultWeather)
	if err := json.Unmarshal(realDate, &r); err != nil {
		return "", err
	}

	var info strings.Builder

	if r.Result.Daily.Temperature != nil && len(r.Result.Daily.Temperature) > 0 {
		info.WriteString(fmt.Sprintf("今日温度为: %0.0f°C~%0.0f°\n", r.Result.Daily.Temperature[0].Min, r.Result.Daily.Temperature[0].Max))
	}
	info.WriteString(fmt.Sprintf("当前温度为: %0.0f°C\n", r.Result.Realtime.Temperature))
	info.WriteString(fmt.Sprintf("体感温度为: %0.1f°C\n", r.Result.Realtime.ApparentTemperature))
	info.WriteString(fmt.Sprintf("相对湿度为: %0.0f%%\n", r.Result.Realtime.Humidity*100))
	// info.WriteString(fmt.Sprintf("PM25浓度为: %d\n", r.Result.Realtime.AirQuality.Pm25))
	// info.WriteString(fmt.Sprintf("空气质量AQI为: %d,%s\n", r.Result.Realtime.AirQuality.Aqi.Chn, r.Result.Realtime.AirQuality.Description.Chn))
	if r.Result.Realtime.LifeIndex.Ultraviolet.Index > 0 {
		info.WriteString(fmt.Sprintf("紫外线指数为: %s\n", r.Result.Realtime.LifeIndex.Ultraviolet.Desc))
	}
	info.WriteString(fmt.Sprintf("舒适度指数为: %s\n", r.Result.Realtime.LifeIndex.Comfort.Desc))

	if r.Result.Alert.Content != nil && len(r.Result.Alert.Content) > 0 {
		info.WriteString(fmt.Sprintf("温馨提示: %s\n", r.Result.Alert.Content[0].Title))
		info.WriteString(fmt.Sprintf("%s\n", r.Result.Alert.Content[0].Description))
	}

	if r.Result.Hourly.Description == r.Result.ForecastKeypoint {
		info.WriteString(fmt.Sprintf("未來兩小時: %s\n", r.Result.ForecastKeypoint))
	} else {
		info.WriteString(fmt.Sprintf("当前天气状况: %s\n", r.Result.Hourly.Description))
		info.WriteString(fmt.Sprintf("未來兩小時: %s\n", r.Result.ForecastKeypoint))
	}

	if r.Result.Daily.Astro != nil && len(r.Result.Daily.Astro) > 0 {
		info.WriteString(fmt.Sprintf("日出: %s, 日落: %s", r.Result.Daily.Astro[0].Sunrise.Time, r.Result.Daily.Astro[0].Sunset.Time))
	}

	return info.String(), nil
}
