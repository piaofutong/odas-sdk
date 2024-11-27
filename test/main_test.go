package test

import (
	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/auth"
	"os"
	"testing"
)

var (
	accessId     = "abcdefg"
	accessKey    = "abcdefg"
	token        = "999d432ba819b623c14d3a728b8f55d1818672b4"
	sid          = 3385
	start        = "2024-09-01"
	end          = "2024-11-01"
	startCompare = "2024-07-01"
	endCompare   = "2024-09-01"
	lid          = "116157,116155"
	excludeLid   = "116157,116156"
	gid          = 2
)

func TestMain(m *testing.M) {
	odas.SetLocalMode()
	os.Exit(m.Run())
}

func TestService_Token(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := auth.NewTokenRequest(accessId, accessKey)
	var r auth.TokenResponse
	err := iam.Do(req, &r)
	if err != nil {
		t.Fatal(err)
	}
}
