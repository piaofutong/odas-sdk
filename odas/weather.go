package odas

import (
	"fmt"
)

type Weather struct {
	Code string
}

func (o *Weather) SetCode(code string) {
	o.Code = code
}

func (o *Weather) Api() string {
	return fmt.Sprintf("/tools/weather/%s", o.Code)
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

func NewWeather(code string) *Weather {
	return &Weather{Code: code}
}

type WeatherResponse struct {
	Now struct {
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
	} `json:"now"`
}
