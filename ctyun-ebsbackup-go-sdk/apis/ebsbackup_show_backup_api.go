package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
)

// EbsbackupShowBackupApi
/* 查询云硬盘备份信息。
 */type EbsbackupShowBackupApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupShowBackupApi(client *core.CtyunClient) *EbsbackupShowBackupApi {
	return &EbsbackupShowBackupApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodGet,
			UrlPath:      "/v4/ebs-backup/show-backup",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupShowBackupApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupShowBackupRequest) (*EbsbackupShowBackupResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	ctReq.AddParam("regionID", req.RegionID)
	ctReq.AddParam("backupID", req.BackupID)
	response, err := a.client.RequestToEndpoint(ctx, ctReq)
	if err != nil {
		return nil, err
	}
	var resp EbsbackupShowBackupResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupShowBackupRequest struct {
	RegionID string `json:"regionID,omitempty"` /*  资源池ID，您可以查看<a href="https://www.ctyun.cn/document/10026730/10028695">地域和可用区</a>来了解资源池 <br />获取：<br /><span style="background-color: rgb(73, 204, 144);color: rgb(255,255,255);padding: 2px; margin:2px">查</span> <a  href="https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=5851&data=87">资源池列表查询</a>  */
	BackupID string `json:"backupID,omitempty"` /*  云硬盘备份ID。  */
}

type EbsbackupShowBackupResponse struct {
	StatusCode  int32                                 `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）。  */
	Message     string                                `json:"message,omitempty"`     /*  成功或失败时的描述，一般为英文描述。  */
	Description string                                `json:"description,omitempty"` /*  成功或失败时的描述，一般为中文描述。  */
	ReturnObj   *EbsbackupShowBackupReturnObjResponse `json:"returnObj"`             /*  返回对象  */
	ErrorCode   string                                `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码。  */
	Error       string                                `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码。  */
}

type EbsbackupShowBackupReturnObjResponse struct {
	RegionID     string `json:"regionID,omitempty"`     /*  资源池ID。  */
	BackupID     string `json:"backupID,omitempty"`     /*  云硬盘备份ID。  */
	BackupName   string `json:"backupName,omitempty"`   /*  云硬盘备份名称。  */
	BackupStatus string `json:"backupStatus,omitempty"` /*  云硬盘备份状态，取值范围：
	●available：可用。
	●error：错误。
	●restoring：恢复中。
	●creating：创建中。
	●deleting：删除中。
	●merging_backup：合并中。
	●frozen：已冻结。  */
	DiskSize            int32  `json:"diskSize,omitempty"`            /*  云硬盘大小，单位GB。  */
	BackupSize          int32  `json:"backupSize,omitempty"`          /*  云硬盘备份大小，单位Byte。  */
	Description         string `json:"description,omitempty"`         /*  云硬盘备份描述。  */
	RepositoryID        string `json:"repositoryID,omitempty"`        /*  备份存储库ID。  */
	RepositoryName      string `json:"repositoryName,omitempty"`      /*  备份存储库名称。  */
	CreatedTime         int32  `json:"createdTime,omitempty"`         /*  备份创建时间。  */
	UpdatedTime         int32  `json:"updatedTime,omitempty"`         /*  备份更新时间。  */
	FinishedTime        int32  `json:"finishedTime,omitempty"`        /*  备份完成时间。  */
	RestoredTime        int32  `json:"restoredTime,omitempty"`        /*  使用该云硬盘备份恢复数据时间。  */
	RestoreFinishedTime int32  `json:"restoreFinishedTime,omitempty"` /*  使用该云硬盘备份恢复完成时间。  */
	Freeze              *bool  `json:"freeze"`                        /*  备份是否冻结。  */
	DiskID              string `json:"diskID,omitempty"`              /*  云硬盘ID。  */
	DiskName            string `json:"diskName,omitempty"`            /*  云硬盘名称。  */
	Encrypted           *bool  `json:"encrypted"`                     /*  云硬盘是否加密。  */
	DiskType            string `json:"diskType,omitempty"`            /*  云硬盘类型，取值范围：
	●SATA：普通IO。
	●SAS：高IO。
	●SSD：超高IO。
	●FAST-SSD：极速型SSD。
	●XSSD-0、XSSD-1、XSSD-2：X系列云硬盘。  */
	Paas         *bool  `json:"paas"`                   /*  是否支持PAAS。  */
	InstanceID   string `json:"instanceID,omitempty"`   /*  云硬盘挂载的云主机ID。  */
	InstanceName string `json:"instanceName,omitempty"` /*  云硬盘挂载的云主机名称。  */
	ProjectID    string `json:"projectID,omitempty"`    /*  企业项目ID。  */
	BackupType   string `json:"backupType,omitempty"`   /*  备份类型，取值范围：
	●full-backup：全量备份
	●incremental-backup：增量备份  */
}
