package odas

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	AccessId  string
	AccessKey string
}

func (o *Client) Do(token string, r IRequest, v any) error {
	u := fmt.Sprintf("%s%s", BaseURL, r.Api())
	request, err := http.NewRequest(r.Method(), u, bytes.NewBuffer(r.Body()))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", r.ContentType())
	if r.Api() != "/token" {
		timestamp := strconv.Itoa(int(time.Now().UnixMilli()))
		request.Header.Set("X-TOKEN", token)
		request.Header.Set("X-TIMESTAMP", timestamp)
		signature := o.sign(request.Method, request.URL.RequestURI(), token, timestamp)
		request.Header.Set("X-SIGNATURE", signature)
	}

	cli := http.Client{}
	response, err := cli.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %d", response.StatusCode)
	}
	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var reply ODASResponse
	err = json.Unmarshal(respBytes, &reply)
	if err != nil {
		return err
	}
	if reply.Code != Ok {
		return fmt.Errorf("code %d, message: %s", reply.Code, reply.Msg)
	}
	var result json.RawMessage
	if reply.Result != nil {
		result = reply.Result
	} else {
		result = reply.Data
	}
	err = json.Unmarshal(result, &v)
	if err != nil {
		return err
	}
	return nil
}

func (o *Client) sign(method, uri, token, timestamp string) string {
	values := url.Values{}
	values.Add("method", method)
	values.Add("api", uri)
	values.Add("token", token)
	values.Add("timestamp", timestamp)
	values.Add("secret", o.AccessKey)
	encoded := values.Encode()
	sum := md5.Sum([]byte(encoded))
	sign := fmt.Sprintf("%x", sum[:])
	return sign
}

func NewClient(accessId, accessKey string) *Client {
	client := &Client{AccessId: accessId, AccessKey: accessKey}
	return client
}
