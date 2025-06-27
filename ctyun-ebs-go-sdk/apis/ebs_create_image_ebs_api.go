package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"net/http"
)

// EbsCreateImageEbsApi
/* 为未挂载的数据盘/系统盘创建私有镜像。
 */type EbsCreateImageEbsApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsCreateImageEbsApi(client *core.CtyunClient) *EbsCreateImageEbsApi {
	return &EbsCreateImageEbsApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs/create-image-ebs",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsCreateImageEbsApi) Do(ctx context.Context, credential core.Credential, req *EbsCreateImageEbsRequest) (*EbsCreateImageEbsResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsCreateImageEbsRequest
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
	var resp EbsCreateImageEbsResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsCreateImageEbsRequest struct {
	RegionID    string `json:"regionID,omitempty"`    /*  资源池ID。  */
	DiskID      string `json:"diskID,omitempty"`      /*  要创建镜像的云硬盘ID。  */
	ImageName   string `json:"imageName,omitempty"`   /*  镜像名称。注意：长度为2~32个字符，只能由数字、字母、“-”组成，不能以数字、“-”开头，且不能以“-”结尾。  */
	Description string `json:"description,omitempty"` /*  镜像描述信息。注意：长度为1~128个字符，不允许换行，不允许空白字符开头或结尾。  */
	ProjectID   string `json:"projectID,omitempty"`   /*  企业项目ID，默认值“0”。  */
	Key         string `json:"key,omitempty"`         /*  标签的key值，长度不能超过32个字符。  */
	Value       string `json:"value,omitempty"`       /*  标签的value值，长度不能超过32个字符。  */
}

type EbsCreateImageEbsResponse struct {
	StatusCode  int32                               `json:"statusCode,omitempty"`  /*  返回状态码(800为成功，900为失败)。  */
	Message     string                              `json:"message,omitempty"`     /*  成功或失败时的描述，一般为英文描述。  */
	Description string                              `json:"description,omitempty"` /*  成功或失败时的描述，一般为中文描述。  */
	ReturnObj   *EbsCreateImageEbsReturnObjResponse `json:"returnObj"`             /*  返回数据结构体。  */
	ErrorCode   string                              `json:"errorCode,omitempty"`   /*  业务细分码，为Product.Module.Code三段式码。请参考错误码。  */
	Error       string                              `json:"error,omitempty"`       /*  业务细分码，为Product.Module.Code三段式大驼峰码。请参考错误码。  */
}

type EbsCreateImageEbsReturnObjResponse struct {
	Images []*EbsCreateImageEbsReturnObjImagesResponse `json:"images"` /*  镜像列表。  */
}

type EbsCreateImageEbsReturnObjImagesResponse struct {
	ImageID         string `json:"imageID,omitempty"`         /*  镜像ID。  */
	ImageName       string `json:"imageName,omitempty"`       /*  镜像名称。  */
	ImageType       string `json:"imageType,omitempty"`       /*  镜像类型。取值范围（值：描述）：<br>（空或空字符串）：系统盘镜像；<br>data_disk_image：数据盘镜像。  */
	Size            int32  `json:"size,omitempty"`            /*  镜像大小，单位为byte，返回时正在队列中或者创建中，值为0。  */
	Status          string `json:"status,omitempty"`          /*  镜像状态，见枚举值。  */
	Visibility      string `json:"visibility,omitempty"`      /*  镜像可见类型，始终为private（私有镜像）。  */
	DiskID          string `json:"diskID,omitempty"`          /*  私有镜像来源的系统盘/数据盘ID。  */
	DiskSize        int32  `json:"diskSize,omitempty"`        /*  磁盘容量，单位为GB。  */
	DiskFormat      string `json:"diskFormat,omitempty"`      /*  磁盘格式，取值范围（值：描述）：<br>qcow2：QCOW2格式；<br>raw：RAW格式；<br>vhd：VHD格式；<br>vmdk：VMDK格式。  */
	AzName          string `json:"azName,omitempty"`          /*  可用区名称。  */
	ProjectID       string `json:"projectID,omitempty"`       /*  企业项目。  */
	Description     string `json:"description,omitempty"`     /*  镜像描述。  */
	Architecture    string `json:"architecture,omitempty"`    /*  镜像系统架构，取值范围（值：描述）：<br>aarch64：AArch64架构，仅支持UEFI启动方式；<br>x86_64：x86_64架构，支持BIOS和UEFI启动方式。  */
	ContainerFormat string `json:"containerFormat,omitempty"` /*  容器格式。  */
	CreatedTime     int32  `json:"createdTime,omitempty"`     /*  镜像创建时间，epoch秒数，即从1970-01-01 00:00:00 UTC到当前时间的秒数。  */
	UpdatedTime     int32  `json:"updatedTime,omitempty"`     /*  镜像更新时间，epoch秒数，即从1970-01-01 00:00:00 UTC到当前时间的秒数。  */
	MaximumRAM      int32  `json:"maximumRAM,omitempty"`      /*  最大内存。  */
	MinimumRAM      int32  `json:"minimumRAM,omitempty"`      /*  最小内存。  */
	OsDistro        string `json:"osDistro,omitempty"`        /*  操作系统的发行版名称。  */
	OsType          string `json:"osType,omitempty"`          /*  操作系统类型，取值范围（值：描述）：<br>linux：Linux系操作系统；<br>windows：Windows系操作系统。  */
	OsVersion       string `json:"osVersion,omitempty"`       /*  操作系统版本。  */
	SourceServerID  string `json:"sourceServerID,omitempty"`  /*  私有镜像来源的云主机/物理机ID。  */
}
