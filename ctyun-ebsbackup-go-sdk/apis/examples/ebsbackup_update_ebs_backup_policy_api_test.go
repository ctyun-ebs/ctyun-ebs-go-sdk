package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/apis"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"testing"
)

func TestEbsbackupUpdateEbsBackupPolicyApi_Do(t *testing.T) {
	// 初始化
	client := core.DefaultClient()
	credential := core.NewCredential("<YOUR_AK>", "<YOUR_SK>")
	// credential := core.CredentialFromEnv()
	new_apis := apis.NewApis("<YOUR_ENDPOINT>", client)
	api := new_apis.EbsbackupUpdateEbsBackupPolicyApi

	// 构造请求
	var remainFirstOfCurMonth bool = false
	var advRetentionStatus bool = true
	request := &apis.EbsbackupUpdateEbsBackupPolicyRequest{
		RegionID:              "81f7728662dd11ec810800155d307d5b",
		PolicyID:              "d15e7d402f8f11ed81370242ac110006",
		PolicyName:            "test-policy",
		CycleType:             "day",
		CycleDay:              1,
		CycleWeek:             "0,2,6",
		Time:                  "1,20",
		RetentionType:         "num",
		RetentionNum:          1,
		RetentionDay:          1,
		RemainFirstOfCurMonth: &remainFirstOfCurMonth,
		FullBackupInterval:    -1,
		AdvRetentionStatus:    &advRetentionStatus,
		AdvRetention: &apis.EbsbackupUpdateEbsBackupPolicyAdvRetentionRequest{
			AdvDay:   1,
			AdvWeek:  1,
			AdvMonth: 1,
			AdvYear:  0,
		},
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
