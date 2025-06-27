package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsDetachEbsApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	apis := NewApis("<YOUR_ENDPOINT>", client)
	api := apis.EbsDetachEbsApi

	// 构造请求
	request := &EbsDetachEbsRequest{
		DiskID:     "eff436e3d44040f1b306ab3a14530f02",
		RegionID:   "41f64827f25f468595ffa3a5deb5d15d",
		InstanceID: "sdff234d4dfs48950f1b45sd02132554",
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
