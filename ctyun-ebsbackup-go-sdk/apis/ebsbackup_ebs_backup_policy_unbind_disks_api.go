package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupEbsBackupPolicyUnbindDisksApi
/* 备份策略解绑云硬盘
 */type EbsbackupEbsBackupPolicyUnbindDisksApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupEbsBackupPolicyUnbindDisksApi(client *core.CtyunClient) *EbsbackupEbsBackupPolicyUnbindDisksApi {
	return &EbsbackupEbsBackupPolicyUnbindDisksApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/policy/unbind-disks",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupEbsBackupPolicyUnbindDisksApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupEbsBackupPolicyUnbindDisksRequest) (*EbsbackupEbsBackupPolicyUnbindDisksResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupEbsBackupPolicyUnbindDisksRequest
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
	var resp EbsbackupEbsBackupPolicyUnbindDisksResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupEbsBackupPolicyUnbindDisksRequest struct {
	RegionID string `json:"regionID,omitempty"` /*  资源池ID  */
	PolicyID string `json:"policyID,omitempty"` /*  备份策略ID  */
	DiskIDs  string `json:"diskIDs,omitempty"`  /*  云硬盘ID  */
}

type EbsbackupEbsBackupPolicyUnbindDisksResponse struct {
	StatusCode  int32                                                 `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                                `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                                `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupEbsBackupPolicyUnbindDisksReturnObjResponse `json:"returnObj"`             /*  无实际对象返回，值为{}  */
	ErrorCode   string                                                `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                                `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupEbsBackupPolicyUnbindDisksReturnObjResponse struct{}
