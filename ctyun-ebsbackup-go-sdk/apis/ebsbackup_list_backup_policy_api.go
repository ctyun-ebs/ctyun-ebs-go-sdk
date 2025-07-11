package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
	"strconv"
)

// EbsbackupListBackupPolicyApi
/* 查询云硬盘备份策略列表。
 */type EbsbackupListBackupPolicyApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupListBackupPolicyApi(client *core.CtyunClient) *EbsbackupListBackupPolicyApi {
	return &EbsbackupListBackupPolicyApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodGet,
			UrlPath:      "/v4/ebs-backup/policy/list-policies",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupListBackupPolicyApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupListBackupPolicyRequest) (*EbsbackupListBackupPolicyResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	ctReq.AddParam("regionID", req.RegionID)
	if req.PageNo != 0 {
		ctReq.AddParam("pageNo", strconv.FormatInt(int64(req.PageNo), 10))
	}
	if req.PageSize != 0 {
		ctReq.AddParam("pageSize", strconv.FormatInt(int64(req.PageSize), 10))
	}
	if req.PolicyID != "" {
		ctReq.AddParam("policyID", req.PolicyID)
	}
	if req.PolicyName != "" {
		ctReq.AddParam("policyName", req.PolicyName)
	}
	if req.ProjectID != "" {
		ctReq.AddParam("projectID", req.ProjectID)
	}
	response, err := a.client.RequestToEndpoint(ctx, ctReq)
	if err != nil {
		return nil, err
	}
	var resp EbsbackupListBackupPolicyResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupListBackupPolicyRequest struct {
	RegionID   string `json:"regionID,omitempty"`   /*  资源池ID */
	PageNo     int32  `json:"pageNo,omitempty"`     /*  页码，默认值1  */
	PageSize   int32  `json:"pageSize,omitempty"`   /*  每页记录数目 ,默认10  */
	PolicyID   string `json:"policyID,omitempty"`   /*  备份策略ID  */
	PolicyName string `json:"policyName,omitempty"` /*  备份策略名，指定了policyID时，该参数会被忽略  */
	ProjectID  string `json:"projectID,omitempty"`  /*  企业项目ID，企业项目管理服务提供统一的云资源按企业项目管理，以及企业项目内的资源管理，成员管理。  */
}

type EbsbackupListBackupPolicyResponse struct {
	StatusCode  int32                                       `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                      `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                      `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupListBackupPolicyReturnObjResponse `json:"returnObj"`             /*  返回对象  */
	ErrorCode   string                                      `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                      `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupListBackupPolicyReturnObjResponse struct{}
