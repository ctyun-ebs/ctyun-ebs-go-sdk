package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/apis"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"testing"
)

func TestEbsbackupListBackupTaskApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	new_apis := apis.NewApis("<YOUR_ENDPOINT>", client)
	api := new_apis.EbsbackupListBackupTaskApi

	// 构造请求
	request := &apis.EbsbackupListBackupTaskRequest{
		RegionID:     "81f7728662dd11ec810800155d307d5b",
		TaskID:       "59093d15-8a3c-53b9-b61b-484af10a3e97",
		QueryContent: "9915c3f4-8d78-445a-a1da-d8d9287d506b",
		TaskStatus:   "running",
		TaskType:     1,
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
