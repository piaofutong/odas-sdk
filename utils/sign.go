package utils

import (
	"crypto/md5"
	"fmt"
	"log/slog"
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
	slog.With(
		slog.String("encoded", encoded),
		slog.String("uri", s.Uri),
		slog.String("method", s.Method),
		slog.String("token", s.Token),
		slog.String("timestamp", s.Timestamp),
		slog.String("secret", s.AccessKey),
		slog.String("sign", sign),
	).Info("请求sign的所有参数")
	return sign
}
