package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupDeleteRepoApi
/* 退订云硬盘备份存储库，该接口会涉及计费<br />

 */type EbsbackupDeleteRepoApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupDeleteRepoApi(client *core.CtyunClient) *EbsbackupDeleteRepoApi {
	return &EbsbackupDeleteRepoApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/repo/delete",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupDeleteRepoApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupDeleteRepoRequest) (*EbsbackupDeleteRepoResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupDeleteRepoRequest
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
	var resp EbsbackupDeleteRepoResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupDeleteRepoRequest struct {
	ClientToken  string `json:"clientToken,omitempty"`  /*  用于保证订单幂等性。要求单个云平台账户内唯一。使用同一个ClientToken值，其他请求参数相同时，则代表为同一个请求  */
	RegionID     string `json:"regionID,omitempty"`     /*  资源池ID  */
	RepositoryID string `json:"repositoryID,omitempty"` /*  云硬盘备份存储库ID  */
}

type EbsbackupDeleteRepoResponse struct {
	StatusCode  int32                                 `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ErrorCode   string                                `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回  */
	ReturnObj   *EbsbackupDeleteRepoReturnObjResponse `json:"returnObj"`             /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupDeleteRepoReturnObjResponse struct{}
