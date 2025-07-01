package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupDeleteEbsBackupApi
/* 删除云硬盘备份。
 */type EbsbackupDeleteEbsBackupApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupDeleteEbsBackupApi(client *core.CtyunClient) *EbsbackupDeleteEbsBackupApi {
	return &EbsbackupDeleteEbsBackupApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/delete",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupDeleteEbsBackupApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupDeleteEbsBackupRequest) (*EbsbackupDeleteEbsBackupResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupDeleteEbsBackupRequest
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
	var resp EbsbackupDeleteEbsBackupResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupDeleteEbsBackupRequest struct {
	RegionID string `json:"regionID,omitempty"` /*  资源池ID */
	BackupID string `json:"backupID,omitempty"` /*  云硬盘备份ID  */
}

type EbsbackupDeleteEbsBackupResponse struct {
	StatusCode  int32                                      `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）。  */
	Message     string                                     `json:"message,omitempty"`     /*  成功或失败时的描述，一般为英文描述。  */
	Description string                                     `json:"description,omitempty"` /*  成功或失败时的描述，一般为中文描述。  */
	ReturnObj   *EbsbackupDeleteEbsBackupReturnObjResponse `json:"returnObj"`             /*  无实际对象返回，值为{}  */
	ErrorCode   string                                     `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码。  */
	Error       string                                     `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码。  */
}

type EbsbackupDeleteEbsBackupReturnObjResponse struct {
	TaskID string `json:"taskID,omitempty"` /*  删除任务ID。  */
}
