package odas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type DoOption struct {
	Token string
}

func NewDoOption() *DoOption {
	return &DoOption{}
}

type Option func(options *DoOption)

func WithToken(token string) Option {
	return func(options *DoOption) {
		options.Token = token
	}
}

type IAM struct {
	AccessId  string
	AccessKey string

	mutex   sync.Mutex
	Builder IBuilder
	Client  IClient
}

func (o *IAM) SetBuilder(builder IBuilder) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.Builder = builder
}

func (o *IAM) Do(req IRequest, v any, opts ...Option) error {
	if o.Builder == nil {
		o.mutex.Lock()
		if o.Builder == nil {
			o.Builder = NewBuilder(o.AccessKey)
		}
		o.mutex.Unlock()
	}
	var options = NewDoOption()
	for _, opt := range opts {
		opt(options)
	}
	if options.Token != "" {
		o.Builder.WithToken(options.Token)
	}
	request, err := o.Builder.Build(req)
	if err != nil {
		return err
	}
	return o.Client.Do(request, v)
}

func NewIAM(accessId, accessKey string) *IAM {
	return &IAM{
		AccessId:  accessId,
		AccessKey: accessKey,
		Builder:   NewBuilder(accessKey),
		Client:    NewClient(),
	}
}

type IClient interface {
	Do(req *http.Request, v any) error
}

type Client struct {
}

func (o *Client) Do(request *http.Request, v any) error {
	cli := http.Client{}
	response, err := cli.Do(request)
	if err != nil {
		return err
	}
	defer func() { _ = response.Body.Close() }()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %d", response.StatusCode)
	}
	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var reply Response
	err = json.Unmarshal(respBytes, &reply)
	if err != nil {
		return err
	}
	if !reply.IsOk() {
		return fmt.Errorf("code %d, message: %s", reply.Code, reply.Msg)
	}
	err = json.Unmarshal(reply.GetResult(), &v)
	if err != nil {
		return err
	}
	return nil
}

func NewClient() *Client {
	return &Client{}
}
