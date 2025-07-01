package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
	"strconv"
)

// EbsbackupListEbsBackupPolicyTasksApi
/* 查询备份策略创建的备份任务列表
 */type EbsbackupListEbsBackupPolicyTasksApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupListEbsBackupPolicyTasksApi(client *core.CtyunClient) *EbsbackupListEbsBackupPolicyTasksApi {
	return &EbsbackupListEbsBackupPolicyTasksApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodGet,
			UrlPath:      "/v4/ebs-backup/policy/list-tasks",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupListEbsBackupPolicyTasksApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupListEbsBackupPolicyTasksRequest) (*EbsbackupListEbsBackupPolicyTasksResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	ctReq.AddParam("regionID", req.RegionID)
	ctReq.AddParam("policyID", req.PolicyID)
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
	if req.TaskStatus != 0 {
		ctReq.AddParam("taskStatus", strconv.FormatInt(int64(req.TaskStatus), 10))
	}
	response, err := a.client.RequestToEndpoint(ctx, ctReq)
	if err != nil {
		return nil, err
	}
	var resp EbsbackupListEbsBackupPolicyTasksResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupListEbsBackupPolicyTasksRequest struct {
	RegionID   string `json:"regionID,omitempty"`   /*  资源池ID */
	PolicyID   string `json:"policyID,omitempty"`   /*  备份策略ID */
	PageNo     int32  `json:"pageNo,omitempty"`     /*  页码，默认值1  */
	PageSize   int32  `json:"pageSize,omitempty"`   /*  每页记录数目 ,默认10  */
	Asc        *bool  `json:"asc"`                  /*  和sort配合使用，是否升序排列，默认降序  */
	Sort       string `json:"sort,omitempty"`       /*  和asc配合使用，指定用于排序的字段。可选字段：createdTime/completedTime，默认createdTime  */
	TaskStatus int32  `json:"taskStatus,omitempty"` /*  备份任务状态，-1-失败，0-执行中，1-成功  */
}

type EbsbackupListEbsBackupPolicyTasksResponse struct {
	StatusCode  int32                                               `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                              `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                              `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupListEbsBackupPolicyTasksReturnObjResponse `json:"returnObj"`             /*  返回对象  */
	ErrorCode   string                                              `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                              `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupListEbsBackupPolicyTasksReturnObjResponse struct{}
