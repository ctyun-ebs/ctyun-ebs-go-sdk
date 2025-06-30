package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/apis"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsQueryEbsListApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("1bd1730c730945a99132011bd48ee085", "ad2b72c33c2f4d3cb38ae5b4a413bae6")
	// credential := core.CredentialFromEnv()
	new_apis := apis.NewApis("https://ebs-global.ctapi.ctyun.cn", client)
	api := new_apis.EbsQueryEbsListApi

	// 构造请求
	request := &apis.EbsQueryEbsListRequest{
		RegionID: "bb9fdb42056f11eda1610242ac110002",
		PageNo:   1,
		PageSize: 10,
		//DiskStatus:     "attached",
		//AzName:         "az1",
		//ProjectID:      "0",
		//DiskType:       "SAS",
		//DiskMode:       "VBD",
		//MultiAttach:    "true",
		//IsSystemVolume: "false",
		//IsEncrypt:      "false",
		//QueryContent:   "test",
	}

	// 发起调用
	response, err := api.Do(context.Background(), *credential, request)
	if err != nil {
		t.Log("request error:", err)
		t.Fail()
		return
	}
	t.Logf("%+v\n", *response)
}
