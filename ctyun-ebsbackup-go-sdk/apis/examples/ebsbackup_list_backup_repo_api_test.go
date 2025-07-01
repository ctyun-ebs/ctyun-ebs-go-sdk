package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/apis"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"testing"
)

func TestEbsbackupListBackupRepoApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	new_apis := apis.NewApis("<YOUR_ENDPOINT>", client)
	api := new_apis.EbsbackupListBackupRepoApi

	// 构造请求
	var hideExpire bool = true
	var asc bool = true
	request := &apis.EbsbackupListBackupRepoRequest{
		RegionID:       "81f7728662dd11ec810800155d307d5b",
		RepositoryName: "test-repo",
		RepositoryID:   "9915c3f4-8d78-445a-a1da-d8d9287d506b",
		Status:         "active",
		HideExpire:     &hideExpire,
		QueryContent:   "repo",
		PageNo:         1,
		PageSize:       10,
		Asc:            &asc,
		Sort:           "createdTime",
		ProjectID:      "0",
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
