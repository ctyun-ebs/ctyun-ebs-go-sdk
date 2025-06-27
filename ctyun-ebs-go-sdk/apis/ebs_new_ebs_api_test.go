package apis

import (
	"context"
	"ctyun-go-sdk/core"
	"testing"
)

func TestEbsNewEbsApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	apis := NewApis("<YOUR_ENDPOINT>", client)
	api := apis.EbsNewEbsApi

	// 构造请求
	var multiAttach bool = false
	var isEncrypt bool = false
	var onDemand bool = false
	var deleteSnapWithEbs bool = false
	request := &EbsNewEbsRequest{
		ClientToken:       "20230211ebsspec7",
		RegionID:          "81f7728662dd11ec810800155d307d5b",
		MultiAttach:       &multiAttach,
		IsEncrypt:         &isEncrypt,
		KmsUUID:           "111d979e-5f30-4dd6-a167-c8c8cdd8aa7c",
		ProjectID:         "0",
		DiskMode:          "VBD",
		DiskType:          "SATA",
		DiskName:          "ebs-newspec-test0211v7",
		DiskSize:          10,
		OnDemand:          &onDemand,
		CycleType:         "month",
		CycleCount:        1,
		ImageID:           "sjsidfnsdfsf",
		AzName:            "az2",
		ProvisionedIops:   1,
		DeleteSnapWithEbs: &deleteSnapWithEbs,
		Labels: []*EbsNewEbsLabelsRequest{
			{
				Key:   "32ff",
				Value: "fe33",
			},
		},
		BackupID: "0ae97ef5-6ee2-44af-9d05-1a509b0a1be6",
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
