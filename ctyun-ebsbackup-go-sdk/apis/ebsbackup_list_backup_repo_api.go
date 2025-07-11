package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
	"strconv"
)

// EbsbackupListBackupRepoApi
/* 查询云硬盘备份存储库列表。
 */type EbsbackupListBackupRepoApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupListBackupRepoApi(client *core.CtyunClient) *EbsbackupListBackupRepoApi {
	return &EbsbackupListBackupRepoApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodGet,
			UrlPath:      "/v4/ebs-backup/repo/list-repos",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupListBackupRepoApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupListBackupRepoRequest) (*EbsbackupListBackupRepoResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	ctReq.AddParam("regionID", req.RegionID)
	if req.RepositoryName != "" {
		ctReq.AddParam("repositoryName", req.RepositoryName)
	}
	if req.RepositoryID != "" {
		ctReq.AddParam("repositoryID", req.RepositoryID)
	}
	if req.Status != "" {
		ctReq.AddParam("status", req.Status)
	}
	if req.HideExpire != nil {
		ctReq.AddParam("hideExpire", strconv.FormatBool(*req.HideExpire))
	}
	if req.QueryContent != "" {
		ctReq.AddParam("queryContent", req.QueryContent)
	}
	if req.PageNo != 0 {
		ctReq.AddParam("pageNo", strconv.FormatInt(int64(req.PageNo), 10))
	}
	if req.PageSize != 0 {
		ctReq.AddParam("pageSize", strconv.FormatInt(int64(req.PageSize), 10))
	}
	if req.Asc != nil {
		ctReq.AddParam("asc", strconv.FormatBool(*req.Asc))
	}
	if req.Sort != "" {
		ctReq.AddParam("sort", req.Sort)
	}
	if req.ProjectID != "" {
		ctReq.AddParam("projectID", req.ProjectID)
	}
	response, err := a.client.RequestToEndpoint(ctx, ctReq)
	if err != nil {
		return nil, err
	}
	var resp EbsbackupListBackupRepoResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupListBackupRepoRequest struct {
	RegionID       string `json:"regionID,omitempty"`       /*  资源池ID  */
	RepositoryName string `json:"repositoryName,omitempty"` /*  云硬盘备份存储库名称。  */
	RepositoryID   string `json:"repositoryID,omitempty"`   /*  云硬盘备份存储库ID。  */
	Status         string `json:"status,omitempty"`         /*  云硬盘备份存储库状态，取值范围：
	●active：可用。
	●master_order_creating：主订单未完成。
	●freezing：已冻结。
	●expired：已过期。  */
	HideExpire   *bool  `json:"hideExpire"`             /*  是否隐藏过期的云硬盘备份存储库。  */
	QueryContent string `json:"queryContent,omitempty"` /*  目前，仅支持备份存储库名称的过滤。  */
	PageNo       int32  `json:"pageNo,omitempty"`       /*  页码，默认1。  */
	PageSize     int32  `json:"pageSize,omitempty"`     /*  每页记录数目 ,默认10。  */
	Asc          *bool  `json:"asc"`                    /*  和sort配合使用，是否升序排列。  */
	Sort         string `json:"sort,omitempty"`         /*  和asc配合使用，指定用于排序的字段。可选字段：createdTime/expiredTime/size/freeSize/usedSize/repositoryName  */
	ProjectID    string `json:"projectID,omitempty"`    /*  企业项目ID，企业项目管理服务提供统一的云资源按企业项目管理，以及企业项目内的资源管理，成员管理。注：默认值为"0"  */
}

type EbsbackupListBackupRepoResponse struct {
	StatusCode  int32                                     `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                    `json:"message,omitempty"`     /*  成功或失败时的描述，一般为英文描述。  */
	Description string                                    `json:"description,omitempty"` /*  成功或失败时的描述，一般为中文描述。  */
	ReturnObj   *EbsbackupListBackupRepoReturnObjResponse `json:"returnObj"`             /*  返回对象数组。  */
	ErrorCode   string                                    `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回。  */
	Error       string                                    `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码。  */
}

type EbsbackupListBackupRepoReturnObjResponse struct{}
