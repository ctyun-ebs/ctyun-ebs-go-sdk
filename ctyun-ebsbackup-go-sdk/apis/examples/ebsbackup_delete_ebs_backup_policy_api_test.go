package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/apis"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"testing"
)

func TestEbsbackupDeleteEbsBackupPolicyApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	new_apis := apis.NewApis("<YOUR_ENDPOINT>", client)
	api := new_apis.EbsbackupDeleteEbsBackupPolicyApi

	// 构造请求
	request := &apis.EbsbackupDeleteEbsBackupPolicyRequest{
		RegionID:  "81f7728662dd11ec810800155d307d5b",
		PolicyIDs: "d15e7d402f8f11ed81370242ac110006，6ff161244e8611ed85170242ac110002",
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
