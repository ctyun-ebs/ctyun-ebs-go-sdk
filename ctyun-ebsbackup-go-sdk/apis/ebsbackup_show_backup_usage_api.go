package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupShowBackupUsageApi
/* 查询云硬盘备份实际占用存储大小
 */type EbsbackupShowBackupUsageApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupShowBackupUsageApi(client *core.CtyunClient) *EbsbackupShowBackupUsageApi {
	return &EbsbackupShowBackupUsageApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodGet,
			UrlPath:      "/v4/ebs-backup/show-usage",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupShowBackupUsageApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupShowBackupUsageRequest) (*EbsbackupShowBackupUsageResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	ctReq.AddParam("regionID", req.RegionID)
	ctReq.AddParam("backupID", req.BackupID)
	response, err := a.client.RequestToEndpoint(ctx, ctReq)
	if err != nil {
		return nil, err
	}
	var resp EbsbackupShowBackupUsageResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupShowBackupUsageRequest struct {
	RegionID string `json:"regionID,omitempty"` /*  资源池ID  */
	BackupID string `json:"backupID,omitempty"` /*  云硬盘备份ID  */
}

type EbsbackupShowBackupUsageResponse struct {
	StatusCode  int32                                      `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                     `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                     `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupShowBackupUsageReturnObjResponse `json:"returnObj"`             /*  是  */
	ErrorCode   string                                     `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回  */
	Error       string                                     `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupShowBackupUsageReturnObjResponse struct{}
