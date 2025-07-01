package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/apis"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsNewEbsApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("1bd1730c730945a99132011bd48ee085", "ad2b72c33c2f4d3cb38ae5b4a413bae6")
	// credential := core.CredentialFromEnv()
	new_apis := apis.NewApis("https://ebs-global.ctapi.ctyun.cn", client)
	api := new_apis.EbsNewEbsApi

	// 构造请求
	var multiAttach bool = false
	var isEncrypt bool = false
	var onDemand bool = false
	var deleteSnapWithEbs bool = false
	request := &apis.EbsNewEbsRequest{
		ClientToken: "20230211ebsspec7",
		RegionID:    "bb9fdb42056f11eda1610242ac110002",
		MultiAttach: &multiAttach,
		IsEncrypt:   &isEncrypt,
		//KmsUUID:           "111d979e-5f30-4dd6-a167-c8c8cdd8aa7c",
		ProjectID:  "0",
		DiskMode:   "VBD",
		DiskType:   "SATA",
		DiskName:   "ebs-newspec-test0211v7",
		DiskSize:   10,
		OnDemand:   &onDemand,
		CycleType:  "month",
		CycleCount: 1,
		//ImageID:           "sjsidfnsdfsf",
		AzName: "cn-huadong1-jsnj1A-public-ctcloud",
		//ProvisionedIops:   1,
		DeleteSnapWithEbs: &deleteSnapWithEbs,
		Labels: []*apis.EbsNewEbsLabelsRequest{
			{
				Key:   "32ff",
				Value: "fe33",
			},
		},
		//BackupID: "0ae97ef5-6ee2-44af-9d05-1a509b0a1be6",
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
