package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupCancelBackupTaskApi
/* 取消云硬盘备份执行中的备份任务。
 */type EbsbackupCancelBackupTaskApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupCancelBackupTaskApi(client *core.CtyunClient) *EbsbackupCancelBackupTaskApi {
	return &EbsbackupCancelBackupTaskApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/task/cancel-task",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupCancelBackupTaskApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupCancelBackupTaskRequest) (*EbsbackupCancelBackupTaskResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupCancelBackupTaskRequest
	}{
		req,
	}, a.template.ContentType)
	if err != nil {
		return nil, err
	}
	response, err := a.client.RequestToEndpoint(ctx, ctReq)
	if err != nil {
		return nil, err
	}
	var resp EbsbackupCancelBackupTaskResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupCancelBackupTaskRequest struct {
	RegionID string `json:"regionID,omitempty"` /*  资源池ID */
	TaskID   string `json:"taskID,omitempty"`   /*  云硬盘备份任务ID。  */
}

type EbsbackupCancelBackupTaskResponse struct {
	StatusCode  int32  `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）。  */
	Message     string `json:"message,omitempty"`     /*  成功或失败时的描述，一般为英文描述。  */
	Description string `json:"description,omitempty"` /*  成功或失败时的描述，一般为中文描述。  */
	ErrorCode   string `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码。  */
	Error       string `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码。  */
}
