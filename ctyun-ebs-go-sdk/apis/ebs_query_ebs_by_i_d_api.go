package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"net/http"
)

// EbsQueryEbsByIDApi
/* 基于磁盘ID查询云硬盘详情。
 */type EbsQueryEbsByIDApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsQueryEbsByIDApi(client *core.CtyunClient) *EbsQueryEbsByIDApi {
	return &EbsQueryEbsByIDApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodGet,
			UrlPath:      "/v4/ebs/info-ebs",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsQueryEbsByIDApi) Do(ctx context.Context, credential core.Credential, req *EbsQueryEbsByIDRequest) (*EbsQueryEbsByIDResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	ctReq.AddParam("diskID", req.DiskID)
	if req.RegionID != "" {
		ctReq.AddParam("regionID", req.RegionID)
	}
	response, err := a.client.RequestToEndpoint(ctx, ctReq)
	if err != nil {
		return nil, err
	}
	var resp EbsQueryEbsByIDResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsQueryEbsByIDRequest struct {
	DiskID   string `json:"diskID,omitempty"`   /*  云硬盘ID。  */
	RegionID string `json:"regionID,omitempty"` /*  资源池ID。如本地语境支持保存regionID，那么建议传递。  */
}

type EbsQueryEbsByIDResponse struct {
	StatusCode  int32                             `json:"statusCode,omitempty"`  /*  返回状态码(800为成功，900为失败)。  */
	Message     string                            `json:"message,omitempty"`     /*  成功或失败时的描述，一般为英文描述。  */
	Description string                            `json:"description,omitempty"` /*  成功或失败时的描述，一般为中文描述。  */
	ReturnObj   *EbsQueryEbsByIDReturnObjResponse `json:"returnObj"`             /*  云硬盘详情对象。  */
	ErrorCode   string                            `json:"errorCode,omitempty"`   /*  业务细分码，为product.module.code三段式码，请参考错误码。  */
	Error       string                            `json:"error,omitempty"`       /*  业务细分码，为product.module.code三段式大驼峰码，请参考错误码。  */
}

type EbsQueryEbsByIDReturnObjResponse struct {
	DiskName string `json:"diskName,omitempty"` /*  云硬盘名称。  */
	DiskID   string `json:"diskID,omitempty"`   /*  云硬盘ID。  */
	DiskSize int32  `json:"diskSize,omitempty"` /*  磁盘大小，单位为GB。  */
	DiskType string `json:"diskType,omitempty"` /*  磁盘规格类型，取值为：
	●SATA：普通IO。
	●SAS：高IO。
	●SSD：超高IO。
	●FAST-SSD：极速型SSD。
	●XSSD-0、XSSD-1、XSSD-2：X系列云硬盘。  */
	DiskMode string `json:"diskMode,omitempty"` /*  云硬盘磁盘模式，取值为：
	●VBD（Virtual Block Device）：虚拟块存储设备。
	●ISCSI （Internet Small Computer System Interface）：小型计算机系统接口。
	●FCSAN（Fibre Channel SAN）：光纤通道协议的SAN网络。  */
	DiskStatus       string                                         `json:"diskStatus,omitempty"`       /*  参考 https://www.ctyun.cn/document/10027696/10168629  */
	CreateTime       int64                                          `json:"createTime,omitempty"`       /*  创建时刻，epoch时戳，精度毫秒。  */
	UpdateTime       int64                                          `json:"updateTime,omitempty"`       /*  更新时刻，epoch时戳，精度毫秒。  */
	ExpireTime       int64                                          `json:"expireTime,omitempty"`       /*  过期时刻，epoch时戳，精度毫秒。  */
	IsSystemVolume   *bool                                          `json:"isSystemVolume"`             /*  只有为系统盘时才返回该字段。  */
	IsPackaged       *bool                                          `json:"isPackaged"`                 /*  是否随云主机一起订购。  */
	InstanceName     string                                         `json:"instanceName,omitempty"`     /*  绑定的云主机名称，有挂载时才返回。  */
	InstanceID       string                                         `json:"instanceID,omitempty"`       /*  绑定的云主机ID，有挂载时才返回。  */
	InstanceStatus   string                                         `json:"instanceStatus,omitempty"`   /*  云主机状态，参考https://www.ctyun.cn/document/10027696/10168629  */
	MultiAttach      *bool                                          `json:"multiAttach"`                /*  是否是共享云硬盘。  */
	Attachments      []*EbsQueryEbsByIDReturnObjAttachmentsResponse `json:"attachments"`                /*  挂载信息。如果是共享挂载云硬盘，则返回多项；无挂载时不返回该字段。  */
	ProjectID        string                                         `json:"projectID,omitempty"`        /*  资源所属的企业项目ID。  */
	IsEncrypt        *bool                                          `json:"isEncrypt"`                  /*  是否是加密盘。  */
	KmsUUID          string                                         `json:"kmsUUID,omitempty"`          /*  加密盘密钥UUID，是加密盘时才返回。  */
	OnDemand         *bool                                          `json:"onDemand"`                   /*  是否按需订购，按需时才返回该字段。  */
	CycleType        string                                         `json:"cycleType,omitempty"`        /*  包周期类型，year：年，month：月。非按需时才返回。  */
	CycleCount       int32                                          `json:"cycleCount,omitempty"`       /*  包周期数，非按需时才返回。  */
	RegionID         string                                         `json:"regionID,omitempty"`         /*  资源池ID。  */
	AzName           string                                         `json:"azName,omitempty"`           /*  多可用区下的可用区名称。  */
	DiskFreeze       *bool                                          `json:"diskFreeze,omitempty"`       /*  云硬盘是否已冻结。  */
	ProvisionedIops  int32                                          `json:"provisionedIops,omitempty"`  /*  XSSD类型盘的预配置iops，未配置返回0，其他类型盘不返回。  */
	VolumeSource     string                                         `json:"volumeSource,omitempty"`     /*  云硬盘源快照ID，若不是从快照创建的则返回null。  */
	SnapshotPolicyID string                                         `json:"snapshotPolicyID,omitempty"` /*  云硬盘绑定的快照策略ID，若没有绑定则返回null。  */
	Labels           []*EbsQueryEbsByIDReturnObjLabelsResponse      `json:"labels"`                     /*  标签信息。资源池支持标签功能才返回。  */
}

type EbsQueryEbsByIDReturnObjAttachmentsResponse struct {
	InstanceID   string `json:"instanceID,omitempty"`   /*  绑定的云主机ID。  */
	AttachmentID string `json:"attachmentID,omitempty"` /*  挂载ID。  */
	Device       string `json:"device,omitempty"`       /*  挂载设备名，例如/dev/sda。  */
}

type EbsQueryEbsByIDReturnObjLabelsResponse struct {
	Key     string `json:"key,omitempty"`     /*  标签键。  */
	Value   string `json:"value,omitempty"`   /*  标签值。  */
	LabelId string `json:"labelId,omitempty"` /*  标签ID。  */
}
