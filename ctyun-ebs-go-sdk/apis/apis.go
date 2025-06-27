package apis

import (
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebs-go-sdk/core"
)

const EndpointName = "ebs"

type Apis struct {
	EbsCreateEbsSnapApi             *EbsCreateEbsSnapApi
	EbsListEbsSnapApi               *EbsListEbsSnapApi
	EbsQuerySizeEbsSnapApi          *EbsQuerySizeEbsSnapApi
	EbsRollbackEbsSnapApi           *EbsRollbackEbsSnapApi
	EbsBatchRollbackEbsSnapApi      *EbsBatchRollbackEbsSnapApi
	EbsNewFromSnapshotEbsSnapApi    *EbsNewFromSnapshotEbsSnapApi
	EbsQueryPolicyEbsSnapApi        *EbsQueryPolicyEbsSnapApi
	EbsCreatePolicyEbsSnapApi       *EbsCreatePolicyEbsSnapApi
	EbsModifyPolicyEbsSnapApi       *EbsModifyPolicyEbsSnapApi
	EbsApplyPolicyEbsSnapApi        *EbsApplyPolicyEbsSnapApi
	EbsCancelPolicyEbsSnapApi       *EbsCancelPolicyEbsSnapApi
	EbsModifyPolicyStatusEbsSnapApi *EbsModifyPolicyStatusEbsSnapApi
	EbsDeleteEbsSnapApi             *EbsDeleteEbsSnapApi
	EbsDeletePolicyEbsSnapApi       *EbsDeletePolicyEbsSnapApi
	EbsAttachEbsApi                 *EbsAttachEbsApi
	EbsDetachEbsApi                 *EbsDetachEbsApi
	EbsRenewEbsApi                  *EbsRenewEbsApi
	EbsRefundEbsApi                 *EbsRefundEbsApi
	EbsNewEbsApi                    *EbsNewEbsApi
	EbsUpdateEbsNameApi             *EbsUpdateEbsNameApi
	EbsQueryEbsListApi              *EbsQueryEbsListApi
	EbsQueryEbsByIDApi              *EbsQueryEbsByIDApi
	EbsQueryEbsByNameApi            *EbsQueryEbsByNameApi
	EbsCreateOrderEbsSnapApi        *EbsCreateOrderEbsSnapApi
	EbsSetDeletePolicyEbsApi        *EbsSetDeletePolicyEbsApi
	EbsUpdateIopsEbsApi             *EbsUpdateIopsEbsApi
	EbsResizeEbsApi                 *EbsResizeEbsApi
	EbsCreateImageEbsApi            *EbsCreateImageEbsApi
}

func NewApis(endpointUrl string, client *core.CtyunClient) *Apis {
	client.RegisterEndpoint(core.Endpoint{
		Name: EndpointName,
		Url:  endpointUrl,
	})
	return &Apis{
		EbsCreateEbsSnapApi:             NewEbsCreateEbsSnapApi(client),
		EbsListEbsSnapApi:               NewEbsListEbsSnapApi(client),
		EbsQuerySizeEbsSnapApi:          NewEbsQuerySizeEbsSnapApi(client),
		EbsRollbackEbsSnapApi:           NewEbsRollbackEbsSnapApi(client),
		EbsBatchRollbackEbsSnapApi:      NewEbsBatchRollbackEbsSnapApi(client),
		EbsNewFromSnapshotEbsSnapApi:    NewEbsNewFromSnapshotEbsSnapApi(client),
		EbsQueryPolicyEbsSnapApi:        NewEbsQueryPolicyEbsSnapApi(client),
		EbsCreatePolicyEbsSnapApi:       NewEbsCreatePolicyEbsSnapApi(client),
		EbsModifyPolicyEbsSnapApi:       NewEbsModifyPolicyEbsSnapApi(client),
		EbsApplyPolicyEbsSnapApi:        NewEbsApplyPolicyEbsSnapApi(client),
		EbsCancelPolicyEbsSnapApi:       NewEbsCancelPolicyEbsSnapApi(client),
		EbsModifyPolicyStatusEbsSnapApi: NewEbsModifyPolicyStatusEbsSnapApi(client),
		EbsDeleteEbsSnapApi:             NewEbsDeleteEbsSnapApi(client),
		EbsDeletePolicyEbsSnapApi:       NewEbsDeletePolicyEbsSnapApi(client),
		EbsAttachEbsApi:                 NewEbsAttachEbsApi(client),
		EbsDetachEbsApi:                 NewEbsDetachEbsApi(client),
		EbsRenewEbsApi:                  NewEbsRenewEbsApi(client),
		EbsRefundEbsApi:                 NewEbsRefundEbsApi(client),
		EbsNewEbsApi:                    NewEbsNewEbsApi(client),
		EbsUpdateEbsNameApi:             NewEbsUpdateEbsNameApi(client),
		EbsQueryEbsListApi:              NewEbsQueryEbsListApi(client),
		EbsQueryEbsByIDApi:              NewEbsQueryEbsByIDApi(client),
		EbsQueryEbsByNameApi:            NewEbsQueryEbsByNameApi(client),
		EbsCreateOrderEbsSnapApi:        NewEbsCreateOrderEbsSnapApi(client),
		EbsSetDeletePolicyEbsApi:        NewEbsSetDeletePolicyEbsApi(client),
		EbsUpdateIopsEbsApi:             NewEbsUpdateIopsEbsApi(client),
		EbsResizeEbsApi:                 NewEbsResizeEbsApi(client),
		EbsCreateImageEbsApi:            NewEbsCreateImageEbsApi(client),
	}
}
