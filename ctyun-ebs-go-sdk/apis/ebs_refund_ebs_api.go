package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
	"net/http"
)

// EbsRefundEbsApi
/* 您可以退订/删除一块包周期/按需的云硬盘，以释放存储空间资源。退订/删除云硬盘后，将停止对云硬盘收费。
 */ /* 包周期云硬盘退订时，按照原始订单实付价格折算退订金额并进行返还。
 */ /* 当云硬盘被退订/删除后，云硬盘的数据将无法被访问。
 */type EbsRefundEbsApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsRefundEbsApi(client *core.CtyunClient) *EbsRefundEbsApi {
	return &EbsRefundEbsApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodPost,
			UrlPath:      "/v4/ebs/refund-ebs",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsRefundEbsApi) Do(ctx context.Context, credential core.Credential, req *EbsRefundEbsRequest) (*EbsRefundEbsResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	_, err := ctReq.WriteJson(struct {
		*EbsRefundEbsRequest
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
	var resp EbsRefundEbsResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsRefundEbsRequest struct {
	ClientToken       string `json:"clientToken,omitempty"` /*  客户端存根，用于保证订单幂等性。要求单个云平台账户内唯一。  */
	DiskID            string `json:"diskID,omitempty"`      /*  需要退订的云硬盘ID。  */
	RegionID          string `json:"regionID,omitempty"`    /*  资源池ID。如本地语境支持保存regionID，那么建议传递。  */
	DeleteSnapWithEbs *bool  `json:"deleteSnapWithEbs"`     /*  设置快照是否随盘删除，只能设置为true。  */
}

type EbsRefundEbsResponse struct {
	StatusCode  int32                            `json:"statusCode,omitempty"`  /*  返回状态码(800为成功，900为处理中/失败)。  */
	Message     string                           `json:"message,omitempty"`     /*  成功或失败时的描述，一般为英文描述。  */
	Description string                           `json:"description,omitempty"` /*  成功或失败时的描述，一般为中文描述。  */
	ReturnObj   *EbsRefundEbsReturnObjResponse   `json:"returnObj"`             /*  返回结构体。  */
	ErrorCode   string                           `json:"errorCode,omitempty"`   /*  业务细分码，为product.module.code三段式码，请参考错误码。  */
	Error       string                           `json:"error,omitempty"`       /*  业务细分码，为product.module.code三段式大驼峰码，请参考错误码。  */
	ErrorDetail *EbsRefundEbsErrorDetailResponse `json:"errorDetail"`           /*  错误明细。一般情况下，会对订单侧(bss)的云硬盘订单业务相关的错误做明确的错误映射和提升，有唯一对应的errorCode。<br> 其他订单侧(bss)的错误，以Ebs.Order.ProcFailed的errorCode统一映射返回，并在errorDetail中返回订单侧的详细错误信息。  */
}

type EbsRefundEbsReturnObjResponse struct {
	MasterOrderID string `json:"masterOrderID,omitempty"` /*  退订订单号，可以使用该订单号确认资源的最终退订状态。  */
	MasterOrderNO string `json:"masterOrderNO,omitempty"` /*  退订订单号。  */
	RegionID      string `json:"regionID,omitempty"`      /*  资源池ID。  */
}

type EbsRefundEbsErrorDetailResponse struct {
	BssErrCode       string `json:"bssErrCode,omitempty"`       /*  bss错误明细码，包含于bss格式化JSON错误信息中。  */
	BssErrMsg        string `json:"bssErrMsg,omitempty"`        /*  bss错误信息，包含于bss格式化JSON错误信息中。  */
	BssOrigErr       string `json:"bssOrigErr,omitempty"`       /*  无法明确解码bss错误信息时，原样透出的bss错误信息。  */
	BssErrPrefixHint string `json:"bssErrPrefixHint,omitempty"` /*  bss格式化JSON错误信息的前置提示信息。  */
}
