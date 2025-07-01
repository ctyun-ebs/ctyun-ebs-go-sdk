package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupEbsBackupPolicyUnbindRepoApi
/* 备份策略解绑云硬盘备份存储库
 */type EbsbackupEbsBackupPolicyUnbindRepoApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupEbsBackupPolicyUnbindRepoApi(client *core.CtyunClient) *EbsbackupEbsBackupPolicyUnbindRepoApi {
	return &EbsbackupEbsBackupPolicyUnbindRepoApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/policy/unbind-repo",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupEbsBackupPolicyUnbindRepoApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupEbsBackupPolicyUnbindRepoRequest) (*EbsbackupEbsBackupPolicyUnbindRepoResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupEbsBackupPolicyUnbindRepoRequest
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
	var resp EbsbackupEbsBackupPolicyUnbindRepoResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupEbsBackupPolicyUnbindRepoRequest struct {
	RegionID     string `json:"regionID,omitempty"`     /*  资源池ID   */
	PolicyIDs    string `json:"policyIDs,omitempty"`    /*  备份策略ID,如果解绑多个,请使用逗号隔开，  */
	RepositoryID string `json:"repositoryID,omitempty"` /*  云硬盘备份存储库ID，云硬盘备份存储库ID，  */
}

type EbsbackupEbsBackupPolicyUnbindRepoResponse struct {
	StatusCode  int32                                                `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                               `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                               `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupEbsBackupPolicyUnbindRepoReturnObjResponse `json:"returnObj"`             /*  无实际对象返回，值为{}  */
	ErrorCode   string                                               `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                               `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupEbsBackupPolicyUnbindRepoReturnObjResponse struct{}
