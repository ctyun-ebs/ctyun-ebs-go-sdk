package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/apis"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsRefundEbsApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	new_apis := apis.NewApis("<YOUR_ENDPOINT>", client)
	api := new_apis.EbsRefundEbsApi

	// 构造请求
	var deleteSnapWithEbs bool = true
	request := &apis.EbsRefundEbsRequest{
		ClientToken:       "refund0211v1",
		DiskID:            "0ae97ef5-6ee2-44af-9d05-1a509b0a1bxx",
		RegionID:          "81f7728662dd11ec810800155d307d5b",
		DeleteSnapWithEbs: &deleteSnapWithEbs,
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
