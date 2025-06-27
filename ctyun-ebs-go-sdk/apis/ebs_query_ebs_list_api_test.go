package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsQueryEbsListApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	apis := NewApis("<YOUR_ENDPOINT>", client)
	api := apis.EbsQueryEbsListApi

	// 构造请求
	request := &EbsQueryEbsListRequest{
		RegionID:       "41f64827f25f468595ffa3a5deb5d15d",
		PageNo:         1,
		PageSize:       10,
		DiskStatus:     "attached",
		AzName:         "az1",
		ProjectID:      "0",
		DiskType:       "SAS",
		DiskMode:       "VBD",
		MultiAttach:    "true",
		IsSystemVolume: "false",
		IsEncrypt:      "false",
		QueryContent:   "test",
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
