package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupDisableEbsBackupPolicyApi
/* 停用云硬盘备份策略
 */type EbsbackupDisableEbsBackupPolicyApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupDisableEbsBackupPolicyApi(client *core.CtyunClient) *EbsbackupDisableEbsBackupPolicyApi {
	return &EbsbackupDisableEbsBackupPolicyApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/policy/disable",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupDisableEbsBackupPolicyApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupDisableEbsBackupPolicyRequest) (*EbsbackupDisableEbsBackupPolicyResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupDisableEbsBackupPolicyRequest
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
	var resp EbsbackupDisableEbsBackupPolicyResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupDisableEbsBackupPolicyRequest struct {
	RegionID string `json:"regionID,omitempty"` /*  资源池ID  */
	PolicyID string `json:"policyID,omitempty"` /*  备份策略ID  */
}

type EbsbackupDisableEbsBackupPolicyResponse struct {
	StatusCode int32 `json:"statusCode,omitempty"` /*  返回状态码（800为成功，900为失败）
	 */
	Message     string                                            `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                            `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupDisableEbsBackupPolicyReturnObjResponse `json:"returnObj"`             /*  无实际对象返回，值为{}  */
	ErrorCode   string                                            `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                            `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupDisableEbsBackupPolicyReturnObjResponse struct{}
