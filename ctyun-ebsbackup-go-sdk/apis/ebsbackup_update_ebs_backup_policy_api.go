package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupUpdateEbsBackupPolicyApi
/* 更新云硬盘备份策略。
 */type EbsbackupUpdateEbsBackupPolicyApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupUpdateEbsBackupPolicyApi(client *core.CtyunClient) *EbsbackupUpdateEbsBackupPolicyApi {
	return &EbsbackupUpdateEbsBackupPolicyApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs-backup/policy/update",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupUpdateEbsBackupPolicyApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupUpdateEbsBackupPolicyRequest) (*EbsbackupUpdateEbsBackupPolicyResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsbackupUpdateEbsBackupPolicyRequest
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
	var resp EbsbackupUpdateEbsBackupPolicyResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupUpdateEbsBackupPolicyRequest struct {
	RegionID              string `json:"regionID,omitempty"`           /*  资源池ID  */
	PolicyID              string `json:"policyID,omitempty"`           /*  备份策略ID，您可以通过<a href="https://www.ctyun.cn/document/10026752/10040084">查询备份策略</a>获取  */
	PolicyName            string `json:"policyName,omitempty"`         /*  策略名，唯一，不可重复  */
	CycleType             string `json:"cycleType,omitempty"`          /*  备份周期类型，day-按天备份，week-按星期备份  */
	CycleDay              int32  `json:"cycleDay,omitempty"`           /*  备份周期，只有cycleType为day时需设置  */
	CycleWeek             string `json:"cycleWeek,omitempty"`          /*  备份周期，只有cycleType为week时需设置，则取值范围0-6代表星期日-星期六，如果一周有多天备份，以逗号隔开  */
	Time                  string `json:"time,omitempty"`               /*  备份整点时间，取值范围0-23，如果一天内多个时间节点备份，以逗号隔开  */
	RetentionType         string `json:"retentionType,omitempty"`      /*  备份保留类型，num-按数量保留，date-按时间保留，all-全部保留  */
	RetentionNum          int32  `json:"retentionNum,omitempty"`       /*  保留数量，只有retentionType为num时需设置,取值范围1-99999  */
	RetentionDay          int32  `json:"retentionDay,omitempty"`       /*  保留天数，只有retentionType为date时需设置，取值1-99999  */
	RemainFirstOfCurMonth *bool  `json:"remainFirstOfCurMonth"`        /*  是否保留每个月第一个备份，在retentionType为num时可设置，默认false  */
	FullBackupInterval    int32  `json:"fullBackupInterval,omitempty"` /*  启用周期性全量备份。-1代表不开启，默认为-1；取值范围为[-1,100]，即每执行n次增量备份后，执行一次全量备份；若传入为0，代表每一次均为全量备份。  */
	AdvRetentionStatus    *bool  `json:"advRetentionStatus"`           /*  是否启用高级保留策略，取值范围：
	●true：启用
	●false：不启用
	默认为false。只有当保留类型为num时，才可启用高级保留策略。  */
	AdvRetention *EbsbackupUpdateEbsBackupPolicyAdvRetentionRequest `json:"advRetention"` /*  高级保留策略内容，配合advRetentionStatus使用。若启用了高级保留策略，可以通过该参数配置具体保留内容。  */
}

type EbsbackupUpdateEbsBackupPolicyAdvRetentionRequest struct {
	AdvDay   int32 `json:"advDay,omitempty"`   /*  ● 保留n天内，每天最新的一个备份。● 单位为天，取值范围：[0, 100]，默认值0  */
	AdvWeek  int32 `json:"advWeek,omitempty"`  /*  ● 保留n周内，每周最新的一个备份。● 单位为周，取值范围：[0, 100]，默认值0  */
	AdvMonth int32 `json:"advMonth,omitempty"` /*  ● 保留n月内，每月最新的一个备份。● 单位为月，取值范围：[0, 100]，默认值0  */
	AdvYear  int32 `json:"advYear,omitempty"`  /*  ● 保留n年内，每年最新的一个备份。● 单位为年，取值范围：[0, 100]，默认值0  */
}

type EbsbackupUpdateEbsBackupPolicyResponse struct {
	StatusCode  int32                                            `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）  */
	Message     string                                           `json:"message,omitempty"`     /*  错误信息的英文描述  */
	Description string                                           `json:"description,omitempty"` /*  错误信息的本地化描述（中文）  */
	ReturnObj   *EbsbackupUpdateEbsBackupPolicyReturnObjResponse `json:"returnObj"`             /*  返回对象  */
	ErrorCode   string                                           `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
	Error       string                                           `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码  */
}

type EbsbackupUpdateEbsBackupPolicyReturnObjResponse struct{}
