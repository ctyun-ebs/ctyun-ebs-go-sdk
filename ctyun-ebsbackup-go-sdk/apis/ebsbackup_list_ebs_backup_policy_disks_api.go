package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
	"strconv"
)

// EbsbackupListEbsBackupPolicyDisksApi
/* 查询云硬盘备份策略绑定的云硬盘列表
 */type EbsbackupListEbsBackupPolicyDisksApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupListEbsBackupPolicyDisksApi(client *core.CtyunClient) *EbsbackupListEbsBackupPolicyDisksApi {
	return &EbsbackupListEbsBackupPolicyDisksApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodGet,
			UrlPath:      "/v4/ebs-backup/policy/list-disks",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupListEbsBackupPolicyDisksApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupListEbsBackupPolicyDisksRequest) (*EbsbackupListEbsBackupPolicyDisksResponse, error) {
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
	if req.DiskID != "" {
		ctReq.AddParam("diskID", req.DiskID)
	}
	if req.DiskName != "" {
		ctReq.AddParam("diskName", req.DiskName)
	}
	response, err := a.client.RequestToEndpoint(ctx, ctReq)
	if err != nil {
		return nil, err
	}
	var resp EbsbackupListEbsBackupPolicyDisksResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupListEbsBackupPolicyDisksRequest struct {
	RegionID string `json:"regionID,omitempty"` /*  资源池ID  */
	PolicyID string `json:"policyID,omitempty"` /*  备份策略ID  */
	PageNo   int32  `json:"pageNo,omitempty"`   /*  页码，默认1  */
	PageSize int32  `json:"pageSize,omitempty"` /*  每页显示条目，默认10  */
	DiskID   string `json:"diskID,omitempty"`   /*  云硬盘ID  */
	DiskName string `json:"diskName,omitempty"` /*  云硬盘名称，模糊过滤，指定diskID时，该参数无效  */
}

type EbsbackupListEbsBackupPolicyDisksResponse struct {
	StatusCode  int32                                               `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                              `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                              `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupListEbsBackupPolicyDisksReturnObjResponse `json:"returnObj"`             /*  返回对象  */
	ErrorCode   string                                              `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                              `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupListEbsBackupPolicyDisksReturnObjResponse struct{}
