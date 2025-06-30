package apis

import (
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
)

const EndpointName = "ebs"

type Apis struct {
	EbsAttachEbsApi              *EbsAttachEbsApi
	EbsDetachEbsApi              *EbsDetachEbsApi
	EbsRenewEbsApi               *EbsRenewEbsApi
	EbsRefundEbsApi              *EbsRefundEbsApi
	EbsNewEbsApi                 *EbsNewEbsApi
	EbsQueryEbsListApi           *EbsQueryEbsListApi
	EbsQueryEbsByIDApi           *EbsQueryEbsByIDApi
	EbsQueryEbsByNameApi         *EbsQueryEbsByNameApi
	EbsResizeEbsApi              *EbsResizeEbsApi
	EbsNewFromSnapshotEbsSnapApi *EbsNewFromSnapshotEbsSnapApi
}

func NewApis(endpointUrl string, client *core.CtyunClient) *Apis {
	client.RegisterEndpoint(core.Endpoint{
		Name: EndpointName,
		Url:  endpointUrl,
	})
	return &Apis{
		EbsAttachEbsApi:              NewEbsAttachEbsApi(client),
		EbsDetachEbsApi:              NewEbsDetachEbsApi(client),
		EbsRenewEbsApi:               NewEbsRenewEbsApi(client),
		EbsRefundEbsApi:              NewEbsRefundEbsApi(client),
		EbsNewEbsApi:                 NewEbsNewEbsApi(client),
		EbsQueryEbsListApi:           NewEbsQueryEbsListApi(client),
		EbsQueryEbsByIDApi:           NewEbsQueryEbsByIDApi(client),
		EbsQueryEbsByNameApi:         NewEbsQueryEbsByNameApi(client),
		EbsResizeEbsApi:              NewEbsResizeEbsApi(client),
		EbsNewFromSnapshotEbsSnapApi: NewEbsNewFromSnapshotEbsSnapApi(client),
	}
}
