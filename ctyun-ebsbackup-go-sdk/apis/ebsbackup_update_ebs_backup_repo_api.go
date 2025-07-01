package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupUpdateEbsBackupRepoApi
/* 更新云硬盘备份存储库信息
 */type EbsbackupUpdateEbsBackupRepoApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupUpdateEbsBackupRepoApi(client *core.CtyunClient) *EbsbackupUpdateEbsBackupRepoApi {
	return &EbsbackupUpdateEbsBackupRepoApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/repo/update",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupUpdateEbsBackupRepoApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupUpdateEbsBackupRepoRequest) (*EbsbackupUpdateEbsBackupRepoResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupUpdateEbsBackupRepoRequest
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
	var resp EbsbackupUpdateEbsBackupRepoResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupUpdateEbsBackupRepoRequest struct {
	RegionID       string `json:"regionID,omitempty"`       /*  资源池ID */
	RepositoryID   string `json:"repositoryID,omitempty"`   /*  云硬盘备份存储库ID  */
	RepositoryName string `json:"repositoryName,omitempty"` /*  云硬盘备份存储库名称  */
}

type EbsbackupUpdateEbsBackupRepoResponse struct {
	StatusCode  int32                                          `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                         `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                         `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupUpdateEbsBackupRepoReturnObjResponse `json:"returnObj"`             /*  无实际对象返回，值为{}  */
	Error       string                                         `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupUpdateEbsBackupRepoReturnObjResponse struct{}
