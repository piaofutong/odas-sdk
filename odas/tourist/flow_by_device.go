package tourist

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"net/http"
	"net/url"
	"strconv"
)

// FlowByDeviceReq 根据设备号查询出入园数据
type FlowByDeviceReq struct {
	Devices string `json:"devices"`
	Hour    int    `json:"hour"`
}

func (f FlowByDeviceReq) Api() string {
	params := url.Values{}
	if f.Devices != "" {
		params.Add("devices", f.Devices)
	}
	if f.Hour > 0 {
		params.Add("hour", strconv.Itoa(f.Hour))
	}
	return fmt.Sprintf("/v2/tourist/inout/flowByDevice?%s", params.Encode())
}

func (f FlowByDeviceReq) Body() []byte {
	return nil
}

func (f FlowByDeviceReq) Method() string {
	return http.MethodGet
}

func (f FlowByDeviceReq) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (f FlowByDeviceReq) AuthRequired() bool {
	return true
}

func NewFlowByDeviceReq(devices string, hour int) FlowByDeviceReq {
	return FlowByDeviceReq{
		Devices: devices,
		Hour:    hour,
	}
}

type FlowByDeviceResponse map[string]*odas.InoutStatVO
