package utils

import (
	"crypto/md5"
	"fmt"
	"net/url"
)

type Signature struct {
	AccessKey string
	Method    string
	Uri       string
	Token     string
	Timestamp string
}

func (s *Signature) Sign() string {
	values := url.Values{}
	values.Add("method", s.Method)
	values.Add("api", s.Uri)
	values.Add("token", s.Token)
	values.Add("timestamp", s.Timestamp)
	values.Add("secret", s.AccessKey)
	encoded := values.Encode()
	sum := md5.Sum([]byte(encoded))
	sign := fmt.Sprintf("%x", sum[:])
	return sign
}
