package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/apis"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsResizeEbsApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("1bd1730c730945a99132011bd48ee085", "ad2b72c33c2f4d3cb38ae5b4a413bae6")
	// credential := core.CredentialFromEnv()
	new_apis := apis.NewApis("https://ebs-global.ctapi.ctyun.cn", client)
	api := new_apis.EbsResizeEbsApi

	// 构造请求
	request := &apis.EbsResizeEbsRequest{
		DiskSize:    20,
		DiskID:      "af53c90b-0098-4df9-b790-688067fc67d6",
		RegionID:    "bb9fdb42056f11eda1610242ac110002",
		ClientToken: "resize0211v3",
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
