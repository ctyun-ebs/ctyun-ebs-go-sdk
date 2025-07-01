package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupDeleteEbsBackupPolicyApi
/* 删除云硬盘备份策略
 */type EbsbackupDeleteEbsBackupPolicyApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupDeleteEbsBackupPolicyApi(client *core.CtyunClient) *EbsbackupDeleteEbsBackupPolicyApi {
	return &EbsbackupDeleteEbsBackupPolicyApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/policy/delete",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupDeleteEbsBackupPolicyApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupDeleteEbsBackupPolicyRequest) (*EbsbackupDeleteEbsBackupPolicyResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupDeleteEbsBackupPolicyRequest
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
	var resp EbsbackupDeleteEbsBackupPolicyResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupDeleteEbsBackupPolicyRequest struct {
	RegionID  string `json:"regionID,omitempty"`  /*  资源池ID  */
	PolicyIDs string `json:"policyIDs,omitempty"` /*  备份策略ID,如果删除多个,请使用逗号隔开。 */
}

type EbsbackupDeleteEbsBackupPolicyResponse struct {
	StatusCode  int32                                            `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                           `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                           `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupDeleteEbsBackupPolicyReturnObjResponse `json:"returnObj"`             /*  无实际对象返回，值为{}  */
	ErrorCode   string                                           `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                           `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupDeleteEbsBackupPolicyReturnObjResponse struct{}
