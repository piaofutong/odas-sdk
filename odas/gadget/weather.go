package gadget

import (
	"fmt"
	"net/url"
)

type WeatherRequestOption struct {
	enableForecast bool
	enableAQI      bool
	enableWarnings bool
	enableIndex    bool
}

func NewWeatherRequestOption() *WeatherRequestOption {
	return &WeatherRequestOption{
		enableForecast: false,
		enableAQI:      false,
		enableWarnings: false,
		enableIndex:    false,
	}
}

type WeatherOption func(options *WeatherRequestOption)

func WithEnableForecast() WeatherOption {
	return func(options *WeatherRequestOption) {
		options.enableForecast = true
	}
}

func WithEnableAQI() WeatherOption {
	return func(options *WeatherRequestOption) {
		options.enableAQI = true
	}
}
func WithEnableWarnings() WeatherOption {
	return func(options *WeatherRequestOption) {
		options.enableWarnings = true
	}

}
func WithEnableIndex() WeatherOption {
	return func(options *WeatherRequestOption) {
		options.enableIndex = true
	}
}

type Weather struct {
	Code    string
	Options *WeatherRequestOption
}

func (o *Weather) SetCode(code string) {
	o.Code = code
}

func (o *Weather) Api() string {
	u := fmt.Sprintf("/tools/weather/%s", o.Code)
	v := url.Values{}
	if o.Options.enableForecast {
		v.Add("forecast", "1")
	} else if o.Options.enableAQI {
		v.Add("aqi", "1")
	} else if o.Options.enableWarnings {
		v.Add("warnings", "1")
	} else if o.Options.enableIndex {
		v.Add("index", "1")
	}
	if len(v) > 0 {
		u = fmt.Sprintf("%s?%s", u, v.Encode())
	}
	return u
}

func (o *Weather) Body() []byte {
	return nil
}

func (o *Weather) Method() string {
	return "GET"
}

func (o *Weather) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (o *Weather) AuthRequired() bool {
	return false
}

func NewWeather(code string, options ...WeatherOption) *Weather {
	requestOption := NewWeatherRequestOption()
	for _, opt := range options {
		opt(requestOption)
	}
	return &Weather{Code: code, Options: requestOption}
}

type WeatherResponse struct {
	Now      WeatherNow        `json:"now"`
	Forecast []WeatherForecast `json:"forecast,omitempty"`
	Index    []WeatherIndex    `json:"index,omitempty"`
	AQI      WeatherAQI        `json:"aqi,omitempty"`
}

type WeatherNow struct {
	ObsTime   string `json:"obsTime"`
	Temp      string `json:"temp"`
	FeelsLike string `json:"feelsLike"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	Precip    string `json:"precip"`
	Pressure  string `json:"pressure"`
	Vis       string `json:"vis"`
	Cloud     string `json:"cloud"`
	Dew       string `json:"dew"`
}

type WeatherForecast struct {
	FxDate         string `json:"fxDate"`
	Sunrise        string `json:"sunrise"`
	Sunset         string `json:"sunset"`
	Moonrise       string `json:"moonrise"`
	Moonset        string `json:"moonset"`
	MoonPhase      string `json:"moonPhase"`
	MoonPhaseIcon  string `json:"moonPhaseIcon"`
	TempMax        string `json:"tempMax"`
	TempMin        string `json:"tempMin"`
	IconDay        string `json:"iconDay"`
	TextDay        string `json:"textDay"`
	IconNight      string `json:"iconNight"`
	TextNight      string `json:"textNight"`
	Wind360Day     string `json:"wind360Day"`
	WindDirDay     string `json:"windDirDay"`
	WindScaleDay   string `json:"windScaleDay"`
	WindSpeedDay   string `json:"windSpeedDay"`
	Wind360Night   string `json:"wind360Night"`
	WindDirNight   string `json:"windDirNight"`
	WindScaleNight string `json:"windScaleNight"`
	WindSpeedNight string `json:"windSpeedNight"`
	Humidity       string `json:"humidity"`
	Precip         string `json:"precip"`
	Pressure       string `json:"pressure"`
	Vis            string `json:"vis"`
	Cloud          string `json:"cloud"`
	UvIndex        string `json:"uvIndex"`
}

type WeatherIndex struct {
	Date     string `json:"date"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	Category string `json:"category"`
	Text     string `json:"text"`
}

type WeatherAQI struct {
	PubTime  string `json:"pubTime"`
	Aqi      string `json:"aqi"`
	Level    string `json:"level"`
	Category string `json:"category"`
	Primary  string `json:"primary"`
	Pm10     string `json:"pm10"`
	Pm2P5    string `json:"pm2p5"`
	No2      string `json:"no2"`
	So2      string `json:"so2"`
	Co       string `json:"co"`
	O3       string `json:"o3"`
}
