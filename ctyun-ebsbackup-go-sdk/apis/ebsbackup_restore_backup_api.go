package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupRestoreBackupApi
/* 恢复云硬盘备份。
 */type EbsbackupRestoreBackupApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupRestoreBackupApi(client *core.CtyunClient) *EbsbackupRestoreBackupApi {
	return &EbsbackupRestoreBackupApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/restore-backup",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupRestoreBackupApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupRestoreBackupRequest) (*EbsbackupRestoreBackupResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupRestoreBackupRequest
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
	var resp EbsbackupRestoreBackupResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupRestoreBackupRequest struct {
	RegionID string `json:"regionID,omitempty"` /*  资源池ID  */
	BackupID string `json:"backupID,omitempty"` /*  云硬盘备份ID  */
	DiskID   string `json:"diskID,omitempty"`   /*  云硬盘ID  */
}

type EbsbackupRestoreBackupResponse struct {
	StatusCode  int32                                    `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                   `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                   `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupRestoreBackupReturnObjResponse `json:"returnObj"`             /*  无实际对象返回，值为{}  */
	ErrorCode   string                                   `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                   `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupRestoreBackupReturnObjResponse struct {
	TaskID string `json:"taskID,omitempty"` /*  恢复任务ID  */
}
