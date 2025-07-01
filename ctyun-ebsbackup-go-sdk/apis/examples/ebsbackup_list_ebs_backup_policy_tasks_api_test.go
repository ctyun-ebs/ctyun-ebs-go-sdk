package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/apis"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"testing"
)

func TestEbsbackupListEbsBackupPolicyTasksApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	new_apis := apis.NewApis("<YOUR_ENDPOINT>", client)
	api := new_apis.EbsbackupListEbsBackupPolicyTasksApi

	// 构造请求
	var asc bool = true
	request := &apis.EbsbackupListEbsBackupPolicyTasksRequest{
		RegionID:   "81f7728662dd11ec810800155d307d5b",
		PolicyID:   "d15e7d402f8f11ed81370242ac110006",
		PageNo:     1,
		PageSize:   10,
		Asc:        &asc,
		Sort:       "createdTime",
		TaskStatus: 1,
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
