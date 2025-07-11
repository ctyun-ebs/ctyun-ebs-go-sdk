package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/apis"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"testing"
)

func TestEbsbackupRestoreBackupApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	new_apis := apis.NewApis("<YOUR_ENDPOINT>", client)
	api := new_apis.EbsbackupRestoreBackupApi

	// 构造请求
	request := &apis.EbsbackupRestoreBackupRequest{
		RegionID: "81f7728662dd11ec810800155d307d5b",
		BackupID: "59093d15-8a3c-53b9-b61b-484af10a3e97",
		DiskID:   "0c582801-6b20-4e3a-956a-f3afbb5e9725",
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
