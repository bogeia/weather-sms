package models

// ResultWeather struct {}
type ResultWeather struct {
	Result struct {
		Alert            Alert    `json:"alert"`    // 预警信息
		Realtime         Realtime `json:"realtime"` // 实况信息
		Minutely         Minutely `json:"minutely"` // 分钟级别信息
		Hourly           Hourly   `json:"hourly"`   // 小时级别信息
		Daily            Daily    `json:"daily"`    // 天级别信息
		ForecastKeypoint string   `json:"forecast_keypoint"`
	} `json:"result"`
}

// Alert struct {}
type Alert struct {
	Status  string `json:"status"`
	Content []struct {
		Province      string    `json:"province"`
		Status        string    `json:"status"`
		Code          string    `json:"code"`
		Description   string    `json:"description"`
		RegionID      string    `json:"regionId"`
		County        string    `json:"county"`
		Pubtimestamp  int       `json:"pubtimestamp"`
		Latlon        []float64 `json:"latlon"`
		City          string    `json:"city"`
		AlertID       string    `json:"alertId"`
		Title         string    `json:"title"`
		Adcode        string    `json:"adcode"`
		Source        string    `json:"source"`
		Location      string    `json:"location"`
		RequestStatus string    `json:"request_status"`
	} `json:"content"`
}

// Realtime struct {}
type Realtime struct {
	Temperature float64 `json:"temperature"` // 地表 2 米气温
	Humidity    float64 `json:"humidity"`    // 地表 2 米湿度相对湿度(%)
	Skycon      string  `json:"skycon"`      // 天气现象
	Wind        struct {
		Speed     float64 `json:"speed"`     // 地表 10 米风速
		Direction float64 `json:"direction"` // 地表 10 米风向
	} `json:"wind"`
	ApparentTemperature float64 `json:"apparent_temperature"` // 体感温度
	AirQuality          struct {
		Pm25 int     `json:"pm25"` // PM25 浓度(μg/m3)
		Pm10 int     `json:"pm10"` // PM10 浓度(μg/m3)
		O3   int     `json:"o3"`   // 臭氧浓度(μg/m3)
		So2  int     `json:"so2"`  // 二氧化氮浓度(μg/m3)
		No2  int     `json:"no2"`  // 二氧化硫浓度(μg/m3)
		Co   float64 `json:"co"`   // 一氧化碳浓度(mg/m3)
		Aqi  struct {
			Chn int `json:"chn"` // 国标 AQI
			Usa int `json:"usa"`
		} `json:"aqi"`
		Description struct {
			Chn string `json:"chn"` // 国标 AQI
			Usa string `json:"usa"`
		} `json:"description"`
	} `json:"air_quality"`
	LifeIndex struct {
		Ultraviolet struct { // 紫外线
			Index float64 `json:"index"` // 等级
			Desc  string  `json:"desc"`  // 等级描述
		} `json:"ultraviolet"`
		Comfort struct { // 舒适度指数
			Index int    `json:"index"` // 等级
			Desc  string `json:"desc"`  // 等级描述
		} `json:"comfort"`
	} `json:"life_index"`
}

// Minutely struct {}
type Minutely struct {
	Status      string `json:"status"`
	Datasource  string `json:"datasource"`
	Description string `json:"description"`
}

// Hourly struct {}
type Hourly struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

// Daily struct {}
type Daily struct {
	Status      string     `json:"status"`
	Description string     `json:"description"`
	Astro       []struct { // 日出日落时间
		Sunrise struct {
			Time string `json:"time"` // 日出时间
		} `json:"sunrise"`
		Sunset struct {
			Time string `json:"time"` // 日落时间
		} `json:"sunset"`
	} `json:"astro"`
	Temperature []struct { // 全天气温
		common
	} `json:"temperature"`
	Skycon []struct { // 全天主要天气现象
		Date  string `json:"date"`
		Value string `json:"value"`
	} `json:"skycon"`
	Skycon08H20H []struct { // 白天主要天气现象
		Date  string `json:"date"`
		Value string `json:"value"`
	} `json:"skycon_08h_20h"`
	LifeIndex struct {
		Ultraviolet []struct { // 紫外线（天级别）
			commonLifeIndex
		} `json:"ultraviolet"`
		CarWashing []struct { // 洗车指数
			commonLifeIndex
		} `json:"carWashing"`
		Dressing []struct { // 穿衣指数
			commonLifeIndex
		} `json:"dressing"`
		Comfort []struct { // 舒适度指数
			commonLifeIndex
		} `json:"comfort"`
		ColdRisk []struct { // 感冒指数
			commonLifeIndex
		} `json:"coldRisk"`
	} `json:"life_index"`
}

// common struct {}
type common struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}

// commonLifeIndex struct {}
type commonLifeIndex struct {
	Date  string `json:"date"`
	Index string `json:"index"`
	Desc  string `json:"desc"`
}
