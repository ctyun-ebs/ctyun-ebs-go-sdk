package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupEbsBackupPolicyBindDisksApi
/* 备份策略绑定云硬盘
 */type EbsbackupEbsBackupPolicyBindDisksApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupEbsBackupPolicyBindDisksApi(client *core.CtyunClient) *EbsbackupEbsBackupPolicyBindDisksApi {
	return &EbsbackupEbsBackupPolicyBindDisksApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/policy/bind-disks",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupEbsBackupPolicyBindDisksApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupEbsBackupPolicyBindDisksRequest) (*EbsbackupEbsBackupPolicyBindDisksResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupEbsBackupPolicyBindDisksRequest
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
	var resp EbsbackupEbsBackupPolicyBindDisksResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupEbsBackupPolicyBindDisksRequest struct {
	RegionID string `json:"regionID,omitempty"` /*  资源池ID */
	PolicyID string `json:"policyID,omitempty"` /*  备份策略ID  */
	DiskIDs  string `json:"diskIDs,omitempty"`  /*  云硬盘ID,如果绑定多个,请使用逗号隔开，  */
}

type EbsbackupEbsBackupPolicyBindDisksResponse struct {
	StatusCode  int32  `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string `json:"description,omitempty"` /*  错误信息的本地化描述（中文）
	 */
	ReturnObj *EbsbackupEbsBackupPolicyBindDisksReturnObjResponse `json:"returnObj"`           /*  无实际对象返回，值为{}  */
	ErrorCode string                                              `json:"errorCode,omitempty"` /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error     string                                              `json:"error,omitempty"`     /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupEbsBackupPolicyBindDisksReturnObjResponse struct{}
