package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsModifyPolicyEbsSnapApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	apis := NewApis("<YOUR_ENDPOINT>", client)
	api := apis.EbsModifyPolicyEbsSnapApi

	// 构造请求
	request := &EbsModifyPolicyEbsSnapRequest{
		RegionID:           "41f64827f25f468595ffa3a5deb5d15d",
		SnapshotPolicyID:   "3641b283-0345-49a9-9c86-bebd963f1caa",
		SnapshotPolicyName: "policy-2",
		RepeatWeekdays:     "0,1,2",
		RepeatTimes:        "0,1,2",
		RetentionTime:      2,
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
