package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/core"
	"testing"
)

func TestEbsListEbsSnapApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	apis := NewApis("<YOUR_ENDPOINT>", client)
	api := apis.EbsListEbsSnapApi

	// 构造请求
	request := &EbsListEbsSnapRequest{
		RegionID:        "fc862f71-d629-4a0e-9fe0-b104963b3f05",
		DiskID:          "61f69121-4036-4f31-a148-457eb8a45fee",
		SnapshotID:      "cafa8c74-42e7-4e32-921c-d7b2c75dac27",
		SnapshotName:    "test-snap",
		SnapshotStatus:  "available",
		SnapshotType:    "manu",
		VolumeAttr:      "data",
		RetentionPolicy: "custom",
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
