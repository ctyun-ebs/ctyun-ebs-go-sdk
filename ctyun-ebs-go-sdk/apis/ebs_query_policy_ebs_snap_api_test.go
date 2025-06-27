package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsQueryPolicyEbsSnapApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	apis := NewApis("<YOUR_ENDPOINT>", client)
	api := apis.EbsQueryPolicyEbsSnapApi

	// 构造请求
	request := &EbsQueryPolicyEbsSnapRequest{
		RegionID:           "41f64827f25f468595ffa3a5deb5d15d",
		SnapshotPolicyID:   "73f321ea-62ff-11ec-a8bc-005056898fe0",
		SnapshotPolicyName: "ecs001",
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
