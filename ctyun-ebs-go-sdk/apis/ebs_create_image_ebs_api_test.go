package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsCreateImageEbsApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	apis := NewApis("<YOUR_ENDPOINT>", client)
	api := apis.EbsCreateImageEbsApi

	// 构造请求
	request := &EbsCreateImageEbsRequest{
		RegionID:    "81f7728662dd11ec810800155d307d5b",
		DiskID:      "83b27404-97fe-4407-a930-25432b7e754e",
		ImageName:   "img-ebs-83b2",
		Description: "private image for ebs-83b2",
		ProjectID:   "0",
		Key:         "group",
		Value:       "1-1",
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
