package tourist

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/tourist"
)

func flow() {
	iam := odas.NewIAM("your_access_id", "your_access_key")
	token := "your_token"
	req := tourist.NewInoutByGroupId(12)
	var r tourist.InoutByGroupIdResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		panic(err)
	}
	for _, v := range r.List {
		fmt.Println(v.Time, v.In, v.Out)
	}
	fmt.Println(r.Total)
}
