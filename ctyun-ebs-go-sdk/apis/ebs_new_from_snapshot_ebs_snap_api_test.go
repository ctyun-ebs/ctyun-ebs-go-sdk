package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsNewFromSnapshotEbsSnapApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	apis := NewApis("<YOUR_ENDPOINT>", client)
	api := apis.EbsNewFromSnapshotEbsSnapApi

	// 构造请求
	var multiAttach bool = false
	var onDemand bool = true
	request := &EbsNewFromSnapshotEbsSnapRequest{
		SnapshotID:  "3f868846-f47f-4619-a5b4-a02e9714f744",
		ClientToken: "cbe3840c-bda4-4102-b68f-98c9d7190d69",
		RegionID:    "41f64827f25f468595ffa3a5deb5d15d",
		MultiAttach: &multiAttach,
		ProjectID:   "0",
		DiskMode:    "VBD",
		DiskName:    "mydisk-0001",
		DiskSize:    1024,
		OnDemand:    &onDemand,
		CycleType:   "month",
		CycleCount:  2,
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
