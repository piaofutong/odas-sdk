package main

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

func main() {
	iam := odas.NewIAM("", "")
	var r odas.WeatherResponse
	weatherRequest := odas.NewWeather("101280601", gadget.WithEnableAQI())
	err := iam.Do(weatherRequest, &r)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Now)
	fmt.Println(r.AQI)
}
