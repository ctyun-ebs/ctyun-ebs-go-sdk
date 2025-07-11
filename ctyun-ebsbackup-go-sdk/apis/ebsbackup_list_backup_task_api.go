package apis

import (
	"context"
	"github.com/ctyun-ebs/ctyun-ebs-go-sdk/ctyun-ebsbackup-go-sdk/core"
	"net/http"
	"strconv"
)

// EbsbackupListBackupTaskApi
/* 查询云硬盘备份任务列表。
 */type EbsbackupListBackupTaskApi struct {
	template core.CtyunRequestTemplate
	client   *core.CtyunClient
}

func NewEbsbackupListBackupTaskApi(client *core.CtyunClient) *EbsbackupListBackupTaskApi {
	return &EbsbackupListBackupTaskApi{
		client: client,
		template: core.CtyunRequestTemplate{
			EndpointName: EndpointName,
			Method:       http.MethodGet,
			UrlPath:      "/v4/ebs-backup/task/list-task",
			ContentType:  "application/json",
		},
	}
}

func (a *EbsbackupListBackupTaskApi) Do(ctx context.Context, credential core.Credential, req *EbsbackupListBackupTaskRequest) (*EbsbackupListBackupTaskResponse, error) {
	builder := core.NewCtyunRequestBuilder(a.template)
	builder.WithCredential(credential)
	ctReq := builder.Build()
	ctReq.AddParam("regionID", req.RegionID)
	if req.TaskID != "" {
		ctReq.AddParam("taskID", req.TaskID)
	}
	if req.QueryContent != "" {
		ctReq.AddParam("queryContent", req.QueryContent)
	}
	if req.TaskStatus != "" {
		ctReq.AddParam("taskStatus", req.TaskStatus)
	}
	if req.TaskType != 0 {
		ctReq.AddParam("taskType", strconv.FormatInt(int64(req.TaskType), 10))
	}
	response, err := a.client.RequestToEndpoint(ctx, ctReq)
	if err != nil {
		return nil, err
	}
	var resp EbsbackupListBackupTaskResponse
	err = response.Parse(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EbsbackupListBackupTaskRequest struct {
	RegionID     string `json:"regionID,omitempty"`     /*  资源池ID  */
	TaskID       string `json:"taskID,omitempty"`       /*  云硬盘备份任务ID。  */
	QueryContent string `json:"queryContent,omitempty"` /*  该参数，可用于模糊过滤，任务ID/云硬盘ID/云硬盘名称/备份任务ID/备份名称/存储库名称，即上述6个字段如果包含该参数的值，则会被过滤出来。  */
	TaskStatus   string `json:"taskStatus,omitempty"`   /*  任务状态：<br />执行中:"running"<br />成功:"success"<br />失败:"failed"<br />已取消:"canceled"<br />取消中:"canceling"  */
	TaskType     int32  `json:"taskType,omitempty"`     /*  任务类型：<br />1:创建任务<br />2:删除任务<br />3:恢复任务  */
}

type EbsbackupListBackupTaskResponse struct {
	StatusCode  int32                                     `json:"statusCode,omitempty"`  /*  返回状态码（800为成功，900为失败）。  */
	Message     string                                    `json:"message,omitempty"`     /*  错误信息的英文描述。  */
	Description string                                    `json:"description,omitempty"` /*  错误信息的本地化描述（中文）。  */
	ReturnObj   *EbsbackupListBackupTaskReturnObjResponse `json:"returnObj"`             /*  成功时返回对象。  */
	ErrorCode   string                                    `json:"errorCode,omitempty"`   /*  业务错误细分码，发生错误时返回，为product.module.code三段式码。  */
	Error       string                                    `json:"error,omitempty"`       /*  业务错误细分码，发生错误时返回，为product.module.code三段式码。  */
}

type EbsbackupListBackupTaskReturnObjResponse struct {
	TaskList     []*EbsbackupListBackupTaskReturnObjTaskListResponse `json:"taskList"`               /*  云硬盘备份任务列表。  */
	TotalCount   int32                                               `json:"totalCount,omitempty"`   /*  云硬盘备份任务总数。  */
	CurrentCount int32                                               `json:"currentCount,omitempty"` /*  当前页云硬盘备份任务数。  */
}

type EbsbackupListBackupTaskReturnObjTaskListResponse struct {
	RegionID       string `json:"regionID,omitempty"`       /*  资源池ID。  */
	TaskID         string `json:"taskID,omitempty"`         /*  备份任务ID。  */
	TaskType       string `json:"taskType,omitempty"`       /*  任务类型:<br />创建任务:create<br />删除任务:delete<br />恢复任务:restore  */
	BackupID       string `json:"backupID,omitempty"`       /*  备份ID。  */
	BackupName     string `json:"backupName,omitempty"`     /*  备份名称。  */
	DiskID         string `json:"diskID,omitempty"`         /*  云硬盘ID。  */
	DiskName       string `json:"diskName,omitempty"`       /*  云硬盘名称。  */
	RepositoryID   string `json:"repositoryID,omitempty"`   /*  备份存储库ID。  */
	RepositoryName string `json:"repositoryName,omitempty"` /*  备份存储库名称。  */
	StatusCode     string `json:"statusCode,omitempty"`     /*  任务状态码。  */
	TaskProgress   int32  `json:"taskProgress,omitempty"`   /*  任务进度，取值为0-100。  */
	Status         string `json:"status,omitempty"`         /*  任务状态描述:<br />成功:success<br />执行中:running<br />失败:failed<br />已取消：canceled<br />取消中：canceling  */
	TaskStartTime  int32  `json:"taskStartTime,omitempty"`  /*  任务开始时间。  */
	TaskDoneTime   int32  `json:"taskDoneTime,omitempty"`   /*  任务完成时间。  */
	TaskErrMsg     string `json:"taskErrMsg,omitempty"`     /*  任务执行过程中的错误信息描述。  */
	TaskDetail     string `json:"taskDetail,omitempty"`     /*  任务执行过程中附加信息描述。  */
	Description    string `json:"description,omitempty"`    /*  任务描述。  */
	PolicyTaskID   string `json:"policyTaskID,omitempty"`   /*  关联的策略调度任务ID。  */
	BackupType     string `json:"backupType,omitempty"`     /*  备份类型:<br />全量备份:totalReplication<br />增量备份:incrementalReplication  */
	ProjectID      string `json:"projectID,omitempty"`      /*  企业项目ID  */
}
