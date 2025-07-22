package tourist

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/inout"
)

func authorizationList() {
	iam := odas.NewIAM("your_access_id", "your_access_key")
	token := "your_token"
	req := &inout.AuthorizationListReq{
		Granter:  10086,
		Page:     1,
		PageSize: 10,
	}
	var r inout.AuthorizationListResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		panic(err)
	}
	for _, v := range r.List {
		fmt.Println(v.Grantee, v.Granter, v.GroupId)
	}
	fmt.Println(r.Pagination)
}
