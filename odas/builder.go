package odas

import (
	"bytes"
	"fmt"
	"github.com/piaofutong/odas-sdk/utils"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type IBuilder interface {
	WithToken(token string)
	Build(req IRequest) (*http.Request, error)
}

type RequestBuilder struct {
	token     string
	accessKey string
}

func (r *RequestBuilder) WithToken(token string) {
	r.token = token
}

func (r *RequestBuilder) WithAccessKey(accessKey string) *RequestBuilder {
	r.accessKey = accessKey
	return r
}

func (r *RequestBuilder) Build(req IRequest) (*http.Request, error) {
	if req.AuthRequired() && r.token == "" {
		return nil, fmt.Errorf("token is required")
	}

	u := fmt.Sprintf("%s%s", baseURL, req.Api())
	request, err := http.NewRequest(req.Method(), u, bytes.NewBuffer(req.Body()))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", req.ContentType())

	if r.token != "" {
		timestamp := strconv.Itoa(int(time.Now().UnixMilli()))
		request.Header.Set("X-TOKEN", r.token)
		request.Header.Set("X-TIMESTAMP", timestamp)
		uri, _ := url.QueryUnescape(req.Api())
		signature := utils.Signature{
			AccessKey: r.accessKey,
			Method:    req.Method(),
			Uri:       uri,
			Token:     r.token,
			Timestamp: timestamp,
		}
		request.Header.Set("X-SIGNATURE", signature.Sign())
	}

	return request, nil
}

func NewBuilder(accessKey string) IBuilder {
	builder := &RequestBuilder{}
	builder.WithAccessKey(accessKey)
	return builder
}
