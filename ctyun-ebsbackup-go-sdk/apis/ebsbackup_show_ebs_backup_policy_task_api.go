package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupShowEbsBackupPolicyTaskApi
/* 查询备份策略创建的备份任务
 */type EbsbackupShowEbsBackupPolicyTaskApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupShowEbsBackupPolicyTaskApi(client *core.CtyunClient) *EbsbackupShowEbsBackupPolicyTaskApi {
	return &EbsbackupShowEbsBackupPolicyTaskApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodGet,
			UrlPath:      "/v4/ebs-backup/policy/show-task",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupShowEbsBackupPolicyTaskApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupShowEbsBackupPolicyTaskRequest) (*EbsbackupShowEbsBackupPolicyTaskResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	ctReq.AddParam("regionID", req.RegionID)
	ctReq.AddParam("policyID", req.PolicyID)
	ctReq.AddParam("taskID", req.TaskID)
	response, err := a.client.RequestToEndpoint(ctx, ctReq)
	if err != nil {
		return nil, err
	}
	var resp EbsbackupShowEbsBackupPolicyTaskResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupShowEbsBackupPolicyTaskRequest struct {
	RegionID string `json:"regionID,omitempty"` /*  资源池ID  */
	PolicyID string `json:"policyID,omitempty"` /*  备份策略ID  */
	TaskID   string `json:"taskID,omitempty"`   /*  备份任务ID  */
}

type EbsbackupShowEbsBackupPolicyTaskResponse struct {
	StatusCode  int32                                              `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                             `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                             `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupShowEbsBackupPolicyTaskReturnObjResponse `json:"returnObj"`             /*  返回对象  */
	ErrorCode   string                                             `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                             `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupShowEbsBackupPolicyTaskReturnObjResponse struct{}
