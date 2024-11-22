package test

import (
	"github.com/piaofutong/odas-sdk/odas"
	"os"
	"testing"
)

var (
	accessId     = "abcdefg"
	accessKey    = "abcdefg"
	token        = "6021d7d7f3e1ed1b93afc6bffab5074b46f5ab21"
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
