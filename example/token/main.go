package main

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas/auth"
	"github.com/piaofutong/odas-sdk/odas/core"
)

func main() {
	iam := core.NewIAM("your-api-key", "your-secret")
	tokenReq := auth.NewTokenRequest(iam.AccessId, iam.AccessKey)
	var r auth.TokenResponse
	err := iam.Do(tokenReq, &r)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.AccessToken)
}
