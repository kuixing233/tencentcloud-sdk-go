// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20191016

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddUserContactRequestParams struct {
	// 联系人姓名，大小写字母+数字+下划线，最小 2 位最大 60 位的长度， 不能以"_"开头，且联系人名保持唯一。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 邮箱地址，大小写字母、数字及下划线组成， 不能以"_"开头。
	ContactInfo *string `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// 服务产品类型，固定值："mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type AddUserContactRequest struct {
	*tchttp.BaseRequest
	
	// 联系人姓名，大小写字母+数字+下划线，最小 2 位最大 60 位的长度， 不能以"_"开头，且联系人名保持唯一。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 邮箱地址，大小写字母、数字及下划线组成， 不能以"_"开头。
	ContactInfo *string `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// 服务产品类型，固定值："mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *AddUserContactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserContactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ContactInfo")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserContactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserContactResponseParams struct {
	// 添加成功的联系人id。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddUserContactResponse struct {
	*tchttp.BaseResponse
	Response *AddUserContactResponseParams `json:"Response"`
}

func (r *AddUserContactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContactItem struct {
	// 联系人id。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 联系人姓名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 联系人绑定的邮箱。
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`
}

// Predefined struct for user
type CreateDBDiagReportTaskRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2020-11-08T14:00:00+08:00”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2020-11-09T14:00:00+08:00”。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否发送邮件: 0 - 否，1 - 是。
	SendMailFlag *int64 `json:"SendMailFlag,omitnil,omitempty" name:"SendMailFlag"`

	// 接收邮件的联系人ID数组。
	ContactPerson []*int64 `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// 接收邮件的联系组ID数组。
	ContactGroup []*int64 `json:"ContactGroup,omitnil,omitempty" name:"ContactGroup"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，"redis" - 云数据库 Redis，默认值为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateDBDiagReportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2020-11-08T14:00:00+08:00”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2020-11-09T14:00:00+08:00”。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否发送邮件: 0 - 否，1 - 是。
	SendMailFlag *int64 `json:"SendMailFlag,omitnil,omitempty" name:"SendMailFlag"`

	// 接收邮件的联系人ID数组。
	ContactPerson []*int64 `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// 接收邮件的联系组ID数组。
	ContactGroup []*int64 `json:"ContactGroup,omitnil,omitempty" name:"ContactGroup"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，"redis" - 云数据库 Redis，默认值为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CreateDBDiagReportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SendMailFlag")
	delete(f, "ContactPerson")
	delete(f, "ContactGroup")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBDiagReportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportTaskResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBDiagReportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBDiagReportTaskResponseParams `json:"Response"`
}

func (r *CreateDBDiagReportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportUrlRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 健康报告相应的任务ID，可通过DescribeDBDiagReportTasks查询。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateDBDiagReportUrlRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 健康报告相应的任务ID，可通过DescribeDBDiagReportTasks查询。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CreateDBDiagReportUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBDiagReportUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportUrlResponseParams struct {
	// 健康报告浏览地址。
	ReportUrl *string `json:"ReportUrl,omitnil,omitempty" name:"ReportUrl"`

	// 健康报告浏览地址到期时间戳（秒）。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBDiagReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBDiagReportUrlResponseParams `json:"Response"`
}

func (r *CreateDBDiagReportUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMailProfileRequestParams struct {
	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// 配置级别，支持值包括："User" - 用户级别，"Instance" - 实例级别，其中数据库巡检邮件配置为用户级别，定期生成邮件配置为实例级别。
	ProfileLevel *string `json:"ProfileLevel,omitnil,omitempty" name:"ProfileLevel"`

	// 配置名称，需要保持唯一性，数据库巡检邮件配置名称自拟；定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"schduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 配置绑定的实例ID，当配置级别为"Instance"时需要传入且只能为一个实例；当配置级别为“User”时，此参数不填。
	BindInstanceIds []*string `json:"BindInstanceIds,omitnil,omitempty" name:"BindInstanceIds"`
}

type CreateMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// 配置级别，支持值包括："User" - 用户级别，"Instance" - 实例级别，其中数据库巡检邮件配置为用户级别，定期生成邮件配置为实例级别。
	ProfileLevel *string `json:"ProfileLevel,omitnil,omitempty" name:"ProfileLevel"`

	// 配置名称，需要保持唯一性，数据库巡检邮件配置名称自拟；定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"schduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 配置绑定的实例ID，当配置级别为"Instance"时需要传入且只能为一个实例；当配置级别为“User”时，此参数不填。
	BindInstanceIds []*string `json:"BindInstanceIds,omitnil,omitempty" name:"BindInstanceIds"`
}

func (r *CreateMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileInfo")
	delete(f, "ProfileLevel")
	delete(f, "ProfileName")
	delete(f, "ProfileType")
	delete(f, "Product")
	delete(f, "BindInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMailProfileResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateMailProfileResponseParams `json:"Response"`
}

func (r *CreateMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulerMailProfileRequestParams struct {
	// 取值范围1-7，分别代表周一至周日。
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitnil,omitempty" name:"WeekConfiguration"`

	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// 配置名称，需要保持唯一性，定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"schduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 配置订阅的实例ID。
	BindInstanceId *string `json:"BindInstanceId,omitnil,omitempty" name:"BindInstanceId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateSchedulerMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// 取值范围1-7，分别代表周一至周日。
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitnil,omitempty" name:"WeekConfiguration"`

	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// 配置名称，需要保持唯一性，定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"schduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 配置订阅的实例ID。
	BindInstanceId *string `json:"BindInstanceId,omitnil,omitempty" name:"BindInstanceId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CreateSchedulerMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulerMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WeekConfiguration")
	delete(f, "ProfileInfo")
	delete(f, "ProfileName")
	delete(f, "BindInstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchedulerMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulerMailProfileResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSchedulerMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateSchedulerMailProfileResponseParams `json:"Response"`
}

func (r *CreateSchedulerMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulerMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityAuditLogExportTaskRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 导出日志开始时间，例如2020-12-28 00:00:00。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 导出日志结束时间，例如2020-12-28 01:00:00。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 日志风险等级列表，支持值包括：0 无风险；1 低风险；2 中风险；3 高风险。
	DangerLevels []*int64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`
}

type CreateSecurityAuditLogExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 导出日志开始时间，例如2020-12-28 00:00:00。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 导出日志结束时间，例如2020-12-28 01:00:00。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 日志风险等级列表，支持值包括：0 无风险；1 低风险；2 中风险；3 高风险。
	DangerLevels []*int64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`
}

func (r *CreateSecurityAuditLogExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAuditLogExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "DangerLevels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityAuditLogExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityAuditLogExportTaskResponseParams struct {
	// 日志导出任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityAuditLogExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityAuditLogExportTaskResponseParams `json:"Response"`
}

func (r *CreateSecurityAuditLogExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAuditLogExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityAuditLogExportTasksRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 日志导出任务Id列表，接口会忽略不存在或已删除的任务Id。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值： "mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DeleteSecurityAuditLogExportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 日志导出任务Id列表，接口会忽略不存在或已删除的任务Id。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值： "mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DeleteSecurityAuditLogExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAuditLogExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "AsyncRequestIds")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityAuditLogExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityAuditLogExportTasksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityAuditLogExportTasksResponseParams `json:"Response"`
}

func (r *DeleteSecurityAuditLogExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAuditLogExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserContactRequestParams struct {
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 联系人名数组，支持模糊搜索。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DescribeAllUserContactRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 联系人名数组，支持模糊搜索。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

func (r *DescribeAllUserContactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserContactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllUserContactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserContactResponseParams struct {
	// 联系人的总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 联系人的信息。
	Contacts []*ContactItem `json:"Contacts,omitnil,omitempty" name:"Contacts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllUserContactResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllUserContactResponseParams `json:"Response"`
}

func (r *DescribeAllUserContactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserGroupRequestParams struct {
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 联系组名称数组，支持模糊搜索。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DescribeAllUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 联系组名称数组，支持模糊搜索。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

func (r *DescribeAllUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserGroupResponseParams struct {
	// 组总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 组信息。
	Groups []*GroupItem `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllUserGroupResponseParams `json:"Response"`
}

func (r *DescribeAllUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事件 ID 。通过“获取实例诊断历史DescribeDBDiagHistory”获取。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBDiagEventRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事件 ID 。通过“获取实例诊断历史DescribeDBDiagHistory”获取。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBDiagEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EventId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventResponseParams struct {
	// 诊断项。
	DiagItem *string `json:"DiagItem,omitnil,omitempty" name:"DiagItem"`

	// 诊断类型。
	DiagType *string `json:"DiagType,omitnil,omitempty" name:"DiagType"`

	// 事件 ID 。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 事件详情。
	Explanation *string `json:"Explanation,omitnil,omitempty" name:"Explanation"`

	// 概要。
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// 诊断出的问题。
	Problem *string `json:"Problem,omitnil,omitempty" name:"Problem"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 建议。
	Suggestions *string `json:"Suggestions,omitnil,omitempty" name:"Suggestions"`

	// 保留字段。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagEventResponseParams `json:"Response"`
}

func (r *DescribeDBDiagEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagHistoryRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-11 12:13:14”，结束时间与开始时间的间隔最大可为2天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBDiagHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-11 12:13:14”，结束时间与开始时间的间隔最大可为2天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBDiagHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagHistoryResponseParams struct {
	// 事件描述。
	Events []*DiagHistoryEventItem `json:"Events,omitnil,omitempty" name:"Events"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagHistoryResponseParams `json:"Response"`
}

func (r *DescribeDBDiagHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagReportTasksRequestParams struct {
	// 第一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 最后一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例ID数组，用于筛选指定实例的任务列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 任务的触发来源，支持的取值包括："DAILY_INSPECTION" - 实例巡检；"SCHEDULED" - 计划任务；"MANUAL" - 手动触发。
	Sources []*string `json:"Sources,omitnil,omitempty" name:"Sources"`

	// 报告的健康等级，支持的取值包括："HEALTH" - 健康；"SUB_HEALTH" - 亚健康；"RISK" - 危险；"HIGH_RISK" - 高危。
	HealthLevels *string `json:"HealthLevels,omitnil,omitempty" name:"HealthLevels"`

	// 任务的状态，支持的取值包括："created" - 新建；"chosen" - 待执行； "running" - 执行中；"failed" - 失败；"finished" - 已完成。
	TaskStatuses *string `json:"TaskStatuses,omitnil,omitempty" name:"TaskStatuses"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL；"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBDiagReportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 第一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 最后一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例ID数组，用于筛选指定实例的任务列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 任务的触发来源，支持的取值包括："DAILY_INSPECTION" - 实例巡检；"SCHEDULED" - 计划任务；"MANUAL" - 手动触发。
	Sources []*string `json:"Sources,omitnil,omitempty" name:"Sources"`

	// 报告的健康等级，支持的取值包括："HEALTH" - 健康；"SUB_HEALTH" - 亚健康；"RISK" - 危险；"HIGH_RISK" - 高危。
	HealthLevels *string `json:"HealthLevels,omitnil,omitempty" name:"HealthLevels"`

	// 任务的状态，支持的取值包括："created" - 新建；"chosen" - 待执行； "running" - 执行中；"failed" - 失败；"finished" - 已完成。
	TaskStatuses *string `json:"TaskStatuses,omitnil,omitempty" name:"TaskStatuses"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL；"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBDiagReportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceIds")
	delete(f, "Sources")
	delete(f, "HealthLevels")
	delete(f, "TaskStatuses")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagReportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagReportTasksResponseParams struct {
	// 任务总数目。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表。
	Tasks []*HealthReportTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagReportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagReportTasksResponseParams `json:"Response"`
}

func (r *DescribeDBDiagReportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSpaceStatusRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 时间段天数，截止日期为当日，默认为7天。
	RangeDays *int64 `json:"RangeDays,omitnil,omitempty" name:"RangeDays"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBSpaceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 时间段天数，截止日期为当日，默认为7天。
	RangeDays *int64 `json:"RangeDays,omitnil,omitempty" name:"RangeDays"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBSpaceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSpaceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RangeDays")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSpaceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSpaceStatusResponseParams struct {
	// 磁盘增长量(MB)。
	Growth *int64 `json:"Growth,omitnil,omitempty" name:"Growth"`

	// 磁盘剩余(MB)。
	Remain *int64 `json:"Remain,omitnil,omitempty" name:"Remain"`

	// 磁盘总量(MB)。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 预计可用天数。
	AvailableDays *int64 `json:"AvailableDays,omitnil,omitempty" name:"AvailableDays"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSpaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSpaceStatusResponseParams `json:"Response"`
}

func (r *DescribeDBSpaceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSpaceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagDBInstancesRequestParams struct {
	// 是否是DBbrain支持的实例，固定传 true。
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分页参数，偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页值。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据实例名称条件查询。
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 根据实例ID条件查询。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 根据地域条件查询。
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type DescribeDiagDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 是否是DBbrain支持的实例，固定传 true。
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分页参数，偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页值。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据实例名称条件查询。
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 根据实例ID条件查询。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 根据地域条件查询。
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

func (r *DescribeDiagDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsSupported")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceNames")
	delete(f, "InstanceIds")
	delete(f, "Regions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiagDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagDBInstancesResponseParams struct {
	// 实例总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 全实例巡检状态：0：开启全实例巡检；1：未开启全实例巡检。
	DbScanStatus *int64 `json:"DbScanStatus,omitnil,omitempty" name:"DbScanStatus"`

	// 实例相关信息。
	Items []*InstanceInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDiagDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiagDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDiagDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthScoreRequestParams struct {
	// 需要获取健康得分的实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取健康得分的时间。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeHealthScoreRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取健康得分的实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取健康得分的时间。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeHealthScoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Time")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHealthScoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthScoreResponseParams struct {
	// 健康得分以及异常扣分项。
	Data *HealthScoreInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHealthScoreResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHealthScoreResponseParams `json:"Response"`
}

func (r *DescribeHealthScoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMailProfileRequestParams struct {
	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单位，最大支持50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据邮件配置名称查询，定期发送的邮件配置名称遵循："scheduler_"+{instanceId}的规则。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`
}

type DescribeMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单位，最大支持50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据邮件配置名称查询，定期发送的邮件配置名称遵循："scheduler_"+{instanceId}的规则。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`
}

func (r *DescribeMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileType")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProfileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMailProfileResponseParams struct {
	// 邮件配置详情。
	ProfileList []*UserProfile `json:"ProfileList,omitnil,omitempty" name:"ProfileList"`

	// 邮件模板总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMailProfileResponseParams `json:"Response"`
}

func (r *DescribeMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySqlProcessListRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 线程的ID，用于筛选线程列表。
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 线程的操作账号名，用于筛选线程列表。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 线程的操作主机地址，用于筛选线程列表。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 线程的操作数据库，用于筛选线程列表。
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 线程的操作状态，用于筛选线程列表。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 线程的执行类型，用于筛选线程列表。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 线程的操作时长最小值，单位秒，用于筛选操作时长大于该值的线程列表。
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 线程的操作语句，用于筛选线程列表。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 返回数量，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeMySqlProcessListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 线程的ID，用于筛选线程列表。
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 线程的操作账号名，用于筛选线程列表。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 线程的操作主机地址，用于筛选线程列表。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 线程的操作数据库，用于筛选线程列表。
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 线程的操作状态，用于筛选线程列表。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 线程的执行类型，用于筛选线程列表。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 线程的操作时长最小值，单位秒，用于筛选操作时长大于该值的线程列表。
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 线程的操作语句，用于筛选线程列表。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 返回数量，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeMySqlProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySqlProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ID")
	delete(f, "User")
	delete(f, "Host")
	delete(f, "DB")
	delete(f, "State")
	delete(f, "Command")
	delete(f, "Time")
	delete(f, "Info")
	delete(f, "Limit")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMySqlProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySqlProcessListResponseParams struct {
	// 实时线程列表。
	ProcessList []*MySqlProcess `json:"ProcessList,omitnil,omitempty" name:"ProcessList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMySqlProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMySqlProcessListResponseParams `json:"Response"`
}

func (r *DescribeMySqlProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySqlProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogDownloadUrlsRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 异步任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSecurityAuditLogDownloadUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 异步任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeSecurityAuditLogDownloadUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogDownloadUrlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "AsyncRequestId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAuditLogDownloadUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogDownloadUrlsResponseParams struct {
	// 导出结果的COS链接列表。当结果集很大时，可能会切分为多个url下载。
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityAuditLogDownloadUrlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAuditLogDownloadUrlsResponseParams `json:"Response"`
}

func (r *DescribeSecurityAuditLogDownloadUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogDownloadUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogExportTasksRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 日志导出任务Id列表。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSecurityAuditLogExportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 日志导出任务Id列表。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSecurityAuditLogExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "Product")
	delete(f, "AsyncRequestIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAuditLogExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogExportTasksResponseParams struct {
	// 安全审计日志导出任务列表。
	Tasks []*SecLogExportTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 安全审计日志导出任务总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAuditLogExportTasksResponseParams `json:"Response"`
}

func (r *DescribeSecurityAuditLogExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTimeSeriesStatsRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-10 12:13:14”，结束时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSlowLogTimeSeriesStatsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-10 12:13:14”，结束时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeSlowLogTimeSeriesStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTimeSeriesStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogTimeSeriesStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTimeSeriesStatsResponseParams struct {
	// 柱间单位时间间隔，单位为秒。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 单位时间间隔内慢日志数量统计。
	TimeSeries []*TimeSlice `json:"TimeSeries,omitnil,omitempty" name:"TimeSeries"`

	// 单位时间间隔内的实例 cpu 利用率监控数据。
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogTimeSeriesStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogTimeSeriesStatsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogTimeSeriesStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTimeSeriesStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTopSqlsRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-10 12:13:14”，截止时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序键，目前支持 QueryTime,ExecTimes,RowsSent,LockTime以及RowsExamined 等排序键。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，支持ASC（升序）以及DESC（降序）。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数据库名称数组。
	SchemaList []*SchemaItem `json:"SchemaList,omitnil,omitempty" name:"SchemaList"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSlowLogTopSqlsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-10 12:13:14”，截止时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序键，目前支持 QueryTime,ExecTimes,RowsSent,LockTime以及RowsExamined 等排序键。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，支持ASC（升序）以及DESC（降序）。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数据库名称数组。
	SchemaList []*SchemaItem `json:"SchemaList,omitnil,omitempty" name:"SchemaList"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeSlowLogTopSqlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTopSqlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SortBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SchemaList")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogTopSqlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTopSqlsResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 慢日志 top sql 列表
	Rows []*SlowLogTopSqlItem `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogTopSqlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogTopSqlsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogTopSqlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTopSqlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogUserHostStatsRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询范围的开始时间，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围的结束时间，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSlowLogUserHostStatsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询范围的开始时间，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围的结束时间，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeSlowLogUserHostStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogUserHostStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogUserHostStatsResponseParams struct {
	// 来源地址数目。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 各来源地址的慢日志占比详情列表。
	Items []*SlowLogHost `json:"Items,omitnil,omitempty" name:"Items"`

	// 各来源用户名的慢日志占比详情列表。
	UserNameItems []*SlowLogUser `json:"UserNameItems,omitnil,omitempty" name:"UserNameItems"`

	// 来源用户数目。
	UserTotalCount *int64 `json:"UserTotalCount,omitnil,omitempty" name:"UserTotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogUserHostStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogUserHostStatsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogUserHostStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemaTimeSeriesRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceSchemaTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceSchemaTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemaTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemaTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemaTimeSeriesResponseParams struct {
	// 返回的Top库空间统计信息的时序数据列表。
	TopSpaceSchemaTimeSeries []*SchemaSpaceTimeSeries `json:"TopSpaceSchemaTimeSeries,omitnil,omitempty" name:"TopSpaceSchemaTimeSeries"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceSchemaTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemaTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceSchemaTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemaTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemasRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceSchemasRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceSchemasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemasResponseParams struct {
	// 返回的Top库空间统计信息列表。
	TopSpaceSchemas []*SchemaSpaceData `json:"TopSpaceSchemas,omitnil,omitempty" name:"TopSpaceSchemas"`

	// 采集库空间数据的时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceSchemasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemasResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceSchemasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTableTimeSeriesRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize，默认为 PhysicalFileSize。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceTableTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize，默认为 PhysicalFileSize。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceTableTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTableTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceTableTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTableTimeSeriesResponseParams struct {
	// 返回的Top表空间统计信息的时序数据列表。
	TopSpaceTableTimeSeries []*TableSpaceTimeSeries `json:"TopSpaceTableTimeSeries,omitnil,omitempty" name:"TopSpaceTableTimeSeries"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceTableTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceTableTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceTableTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTableTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTablesRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceTablesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTablesResponseParams struct {
	// 返回的Top表空间统计信息列表。
	TopSpaceTables []*TableSpaceData `json:"TopSpaceTables,omitnil,omitempty" name:"TopSpaceTables"`

	// 采集表空间数据的时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceTablesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 库名。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`
}

type DescribeUserSqlAdviceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 库名。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`
}

func (r *DescribeUserSqlAdviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SqlText")
	delete(f, "Schema")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSqlAdviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceResponseParams struct {
	// SQL优化建议，可解析为JSON数组。
	Advices *string `json:"Advices,omitnil,omitempty" name:"Advices"`

	// SQL优化建议备注，可解析为String数组。
	Comments *string `json:"Comments,omitnil,omitempty" name:"Comments"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 库名。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 相关表的DDL信息，可解析为JSON数组。
	Tables *string `json:"Tables,omitnil,omitempty" name:"Tables"`

	// SQL执行计划，可解析为JSON。
	SqlPlan *string `json:"SqlPlan,omitnil,omitempty" name:"SqlPlan"`

	// SQL优化后的成本节约详情，可解析为JSON。
	Cost *string `json:"Cost,omitnil,omitempty" name:"Cost"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserSqlAdviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserSqlAdviceResponseParams `json:"Response"`
}

func (r *DescribeUserSqlAdviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiagHistoryEventItem struct {
	// 诊断类型。
	DiagType *string `json:"DiagType,omitnil,omitempty" name:"DiagType"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事件 ID 。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 概要。
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// 诊断项。
	DiagItem *string `json:"DiagItem,omitnil,omitempty" name:"DiagItem"`

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 保留字段
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type EventInfo struct {
	// 事件 ID 。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 诊断类型。
	DiagType *string `json:"DiagType,omitnil,omitempty" name:"DiagType"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 概要。
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 扣分。
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`

	// 保留字段。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 告警数目。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type GroupItem struct {
	// 组id。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 组名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 组成员数量。
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`
}

type HealthReportTask struct {
	// 异步任务请求 ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 任务的触发来源，支持的取值包括："DAILY_INSPECTION" - 实例巡检；"SCHEDULED" - 定时生成；"MANUAL" - 手动触发。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 任务完成进度，单位%。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务开始执行时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务完成执行时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务所属实例的基础信息。
	InstanceInfo *InstanceBasicInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// 健康报告中的健康信息。
	HealthStatus *HealthStatus `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`
}

type HealthScoreInfo struct {
	// 异常详情。
	IssueTypes []*IssueTypeInfo `json:"IssueTypes,omitnil,omitempty" name:"IssueTypes"`

	// 异常事件总数。
	EventsTotalCount *int64 `json:"EventsTotalCount,omitnil,omitempty" name:"EventsTotalCount"`

	// 健康得分。
	HealthScore *int64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// 健康等级, 如："HEALTH", "SUB_HEALTH", "RISK", "HIGH_RISK"。
	HealthLevel *string `json:"HealthLevel,omitnil,omitempty" name:"HealthLevel"`
}

type HealthStatus struct {
	// 健康分数，满分100。
	HealthScore *int64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// 健康等级，取值包括："HEALTH" - 健康；"SUB_HEALTH" - 亚健康；"RISK"- 危险；"HIGH_RISK" - 高危。
	HealthLevel *string `json:"HealthLevel,omitnil,omitempty" name:"HealthLevel"`

	// 总扣分分数。
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`

	// 扣分详情。
	ScoreDetails []*ScoreDetail `json:"ScoreDetails,omitnil,omitempty" name:"ScoreDetails"`

	// 健康等级版本，默认为V1
	HealthLevelVersion *string `json:"HealthLevelVersion,omitnil,omitempty" name:"HealthLevelVersion"`
}

type InstanceBasicInfo struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例内网IP。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 实例内网Port。
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 实例产品。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例引擎版本。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// CPU数量，对于Redis为0。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例部署模式。
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 实例内存配置。
	InstanceConf *RedisInstanceConf `json:"InstanceConf,omitnil,omitempty" name:"InstanceConf"`

	// DBbrain是否支持该实例。
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// 实例内存，单位MB。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例子网统一ID，对于redis为空字符串。	
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 实例私有网络统一ID，对于redis为空字符串。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 实例磁盘容量，对于Redis为0。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`
}

type InstanceConfs struct {
	// 数据库巡检开关, Yes/No。
	DailyInspection *string `json:"DailyInspection,omitnil,omitempty" name:"DailyInspection"`

	// 实例概览开关，Yes/No。
	OverviewDisplay *string `json:"OverviewDisplay,omitnil,omitempty" name:"OverviewDisplay"`

	// redis大key分析的自定义分割符，仅redis使用
	KeyDelimiters []*string `json:"KeyDelimiters,omitnil,omitempty" name:"KeyDelimiters"`
}

type InstanceInfo struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例所属地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 健康得分。
	HealthScore *int64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// 所属产品。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 异常事件数量。
	EventCount *int64 `json:"EventCount,omitnil,omitempty" name:"EventCount"`

	// 实例类型：1:MASTER；2:DR，3：RO，4:SDR。
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 核心数。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，单位MB。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 硬盘存储，单位GB。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 数据库版本。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 内网地址。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 内网端口。
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 接入来源。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 分组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组组名。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 实例状态：0：发货中；1：运行正常；4：销毁中；5：隔离中。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 子网统一ID。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// cdb类型。
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// cdb实例初始化标志：0：未初始化；1：已初始化。
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// 任务状态。
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 私有网络统一ID。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 实例巡检/概览的状态。
	InstanceConf *InstanceConfs `json:"InstanceConf,omitnil,omitempty" name:"InstanceConf"`

	// 资源到期时间。
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// 是否是DBbrain支持的实例。
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// 实例安全审计日志开启状态：ON： 安全审计开启；OFF： 未开启安全审计。
	SecAuditStatus *string `json:"SecAuditStatus,omitnil,omitempty" name:"SecAuditStatus"`

	// 实例审计日志开启状态，ALL_AUDIT： 开启全审计；RULE_AUDIT： 开启规则审计；UNBOUND： 未开启审计。
	AuditPolicyStatus *string `json:"AuditPolicyStatus,omitnil,omitempty" name:"AuditPolicyStatus"`

	// 实例审计日志运行状态：normal： 运行中； paused： 欠费暂停。
	AuditRunningStatus *string `json:"AuditRunningStatus,omitnil,omitempty" name:"AuditRunningStatus"`

	// 内网vip。
	InternalVip *string `json:"InternalVip,omitnil,omitempty" name:"InternalVip"`

	// 内网port。
	InternalVport *int64 `json:"InternalVport,omitnil,omitempty" name:"InternalVport"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 所属集群ID（仅对集群数据库产品该字段非空，如TDSQL-C）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 所属集群名称（仅对集群数据库产品该字段非空，如TDSQL-C）。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type IssueTypeInfo struct {
	// 指标分类：AVAILABILITY：可用性，MAINTAINABILITY：可维护性，PERFORMANCE，性能，RELIABILITY可靠性。
	IssueType *string `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// 异常事件。
	Events []*EventInfo `json:"Events,omitnil,omitempty" name:"Events"`

	// 异常事件总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type MailConfiguration struct {
	// 是否开启邮件发送: 0, 否; 1, 是。
	SendMail *int64 `json:"SendMail,omitnil,omitempty" name:"SendMail"`

	// 地域配置, 如["ap-guangzhou", "ap-shanghai"]。巡检的邮件发送模板，配置需要发送巡检邮件的地域；订阅的邮件发送模板，配置当前订阅实例的所属地域。
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`

	// 发送指定的健康等级的报告, 如["HEALTH", "SUB_HEALTH", "RISK", "HIGH_RISK"]。
	HealthStatus []*string `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// 联系人id, 联系人/联系组不能都为空。
	ContactPerson []*int64 `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// 联系组id, 联系人/联系组不能都为空。
	ContactGroup []*int64 `json:"ContactGroup,omitnil,omitempty" name:"ContactGroup"`
}

// Predefined struct for user
type ModifyDiagDBInstanceConfRequestParams struct {
	// 巡检开关。
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitnil,omitempty" name:"InstanceConfs"`

	// 生效实例地域，取值为"All"，代表全地域。
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 指定更改巡检状态的实例ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type ModifyDiagDBInstanceConfRequest struct {
	*tchttp.BaseRequest
	
	// 巡检开关。
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitnil,omitempty" name:"InstanceConfs"`

	// 生效实例地域，取值为"All"，代表全地域。
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 指定更改巡检状态的实例ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *ModifyDiagDBInstanceConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiagDBInstanceConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceConfs")
	delete(f, "Regions")
	delete(f, "Product")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiagDBInstanceConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiagDBInstanceConfResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDiagDBInstanceConfResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiagDBInstanceConfResponseParams `json:"Response"`
}

func (r *ModifyDiagDBInstanceConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiagDBInstanceConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorFloatMetric struct {
	// 指标名称。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 指标单位。
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 指标值。
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type MonitorFloatMetricSeriesData struct {
	// 监控指标。
	Series []*MonitorFloatMetric `json:"Series,omitnil,omitempty" name:"Series"`

	// 监控指标对应的时间戳。
	Timestamp []*int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type MonitorMetric struct {
	// 指标名称。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 指标单位。
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 指标值。
	Values []*int64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type MonitorMetricSeriesData struct {
	// 监控指标。
	Series []*MonitorMetric `json:"Series,omitnil,omitempty" name:"Series"`

	// 监控指标对应的时间戳。
	Timestamp []*int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type MySqlProcess struct {
	// 线程ID。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 线程的操作账号名。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 线程的操作主机地址。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 线程的操作数据库。
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 线程的操作状态。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 线程的执行类型。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 线程的操作时长，单位秒。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 线程的操作语句。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`
}

type ProfileInfo struct {
	// 语言, 如"zh"。
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// 邮件模板的内容。
	MailConfiguration *MailConfiguration `json:"MailConfiguration,omitnil,omitempty" name:"MailConfiguration"`
}

type RedisInstanceConf struct {
	// 副本数量
	ReplicasNum *string `json:"ReplicasNum,omitnil,omitempty" name:"ReplicasNum"`

	// 分片数量
	ShardNum *string `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 分片内存大小，单位MB
	ShardSize *string `json:"ShardSize,omitnil,omitempty" name:"ShardSize"`
}

type SchemaItem struct {
	// 数据库名称
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`
}

type SchemaSpaceData struct {
	// 库名。
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// 数据空间（MB）。
	DataLength *float64 `json:"DataLength,omitnil,omitempty" name:"DataLength"`

	// 索引空间（MB）。
	IndexLength *float64 `json:"IndexLength,omitnil,omitempty" name:"IndexLength"`

	// 碎片空间（MB）。
	DataFree *float64 `json:"DataFree,omitnil,omitempty" name:"DataFree"`

	// 总使用空间（MB）。
	TotalLength *float64 `json:"TotalLength,omitnil,omitempty" name:"TotalLength"`

	// 碎片率（%）。
	FragRatio *float64 `json:"FragRatio,omitnil,omitempty" name:"FragRatio"`

	// 行数。
	TableRows *int64 `json:"TableRows,omitnil,omitempty" name:"TableRows"`

	// 库中所有表对应的独立物理文件大小加和（MB）。
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitnil,omitempty" name:"PhysicalFileSize"`
}

type SchemaSpaceTimeSeries struct {
	// 库名
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// 单位时间间隔内的空间指标数据。
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`
}

type ScoreDetail struct {
	// 扣分项分类，取值包括：可用性、可维护性、性能及可靠性。
	IssueType *string `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// 扣分总分。
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`

	// 扣分总分上限。
	ScoreLostMax *int64 `json:"ScoreLostMax,omitnil,omitempty" name:"ScoreLostMax"`

	// 扣分项列表。
	Items []*ScoreItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type ScoreItem struct {
	// 异常诊断项名称。
	DiagItem *string `json:"DiagItem,omitnil,omitempty" name:"DiagItem"`

	// 诊断项分类，取值包括：可用性、可维护性、性能及可靠性。
	IssueType *string `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// 健康等级，取值包括：信息、提示、告警、严重、致命。
	TopSeverity *string `json:"TopSeverity,omitnil,omitempty" name:"TopSeverity"`

	// 该异常诊断项出现次数。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 扣分分数。
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`
}

type SecLogExportTaskInfo struct {
	// 异步任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 任务开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务执行进度。
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 导出日志开始时间。
	LogStartTime *string `json:"LogStartTime,omitnil,omitempty" name:"LogStartTime"`

	// 导出日志结束时间。
	LogEndTime *string `json:"LogEndTime,omitnil,omitempty" name:"LogEndTime"`

	// 日志文件总大小，单位KB。
	TotalSize *uint64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 风险等级列表。0 无风险；1 低风险；2 中风险；3 高风险。
	DangerLevels []*uint64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`
}

type SlowLogHost struct {
	// 来源地址。
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// 该来源地址的慢日志数目占总数目的比例，单位%。
	Ratio *float64 `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// 该来源地址的慢日志数目。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type SlowLogTopSqlItem struct {
	// sql总锁等待时间
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// 最大锁等待时间
	LockTimeMax *float64 `json:"LockTimeMax,omitnil,omitempty" name:"LockTimeMax"`

	// 最小锁等待时间
	LockTimeMin *float64 `json:"LockTimeMin,omitnil,omitempty" name:"LockTimeMin"`

	// 总扫描行数
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// 最大扫描行数
	RowsExaminedMax *int64 `json:"RowsExaminedMax,omitnil,omitempty" name:"RowsExaminedMax"`

	// 最小扫描行数
	RowsExaminedMin *int64 `json:"RowsExaminedMin,omitnil,omitempty" name:"RowsExaminedMin"`

	// 总耗时
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// 最大执行时间
	QueryTimeMax *float64 `json:"QueryTimeMax,omitnil,omitempty" name:"QueryTimeMax"`

	// 最小执行时间
	QueryTimeMin *float64 `json:"QueryTimeMin,omitnil,omitempty" name:"QueryTimeMin"`

	// 总返回行数
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`

	// 最大返回行数
	RowsSentMax *int64 `json:"RowsSentMax,omitnil,omitempty" name:"RowsSentMax"`

	// 最小返回行数
	RowsSentMin *int64 `json:"RowsSentMin,omitnil,omitempty" name:"RowsSentMin"`

	// 执行次数
	ExecTimes *int64 `json:"ExecTimes,omitnil,omitempty" name:"ExecTimes"`

	// sql模板
	SqlTemplate *string `json:"SqlTemplate,omitnil,omitempty" name:"SqlTemplate"`

	// 带参数SQL（随机）
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 数据库名
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 总耗时占比
	QueryTimeRatio *float64 `json:"QueryTimeRatio,omitnil,omitempty" name:"QueryTimeRatio"`

	// sql总锁等待时间占比
	LockTimeRatio *float64 `json:"LockTimeRatio,omitnil,omitempty" name:"LockTimeRatio"`

	// 总扫描行数占比
	RowsExaminedRatio *float64 `json:"RowsExaminedRatio,omitnil,omitempty" name:"RowsExaminedRatio"`

	// 总返回行数占比
	RowsSentRatio *float64 `json:"RowsSentRatio,omitnil,omitempty" name:"RowsSentRatio"`

	// 平均执行时间
	QueryTimeAvg *float64 `json:"QueryTimeAvg,omitnil,omitempty" name:"QueryTimeAvg"`

	// 平均返回行数
	RowsSentAvg *float64 `json:"RowsSentAvg,omitnil,omitempty" name:"RowsSentAvg"`

	// 平均锁等待时间
	LockTimeAvg *float64 `json:"LockTimeAvg,omitnil,omitempty" name:"LockTimeAvg"`

	// 平均扫描行数
	RowsExaminedAvg *float64 `json:"RowsExaminedAvg,omitnil,omitempty" name:"RowsExaminedAvg"`

	// SQL模板的MD5值
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`
}

type SlowLogUser struct {
	// 来源用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 该来源用户名的慢日志数目占总数目的比例，单位%。
	Ratio *float64 `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// 该来源用户名的慢日志数目。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type TableSpaceData struct {
	// 表名。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 库名。
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// 库表的存储引擎。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 数据空间（MB）。
	DataLength *float64 `json:"DataLength,omitnil,omitempty" name:"DataLength"`

	// 索引空间（MB）。
	IndexLength *float64 `json:"IndexLength,omitnil,omitempty" name:"IndexLength"`

	// 碎片空间（MB）。
	DataFree *float64 `json:"DataFree,omitnil,omitempty" name:"DataFree"`

	// 总使用空间（MB）。
	TotalLength *float64 `json:"TotalLength,omitnil,omitempty" name:"TotalLength"`

	// 碎片率（%）。
	FragRatio *float64 `json:"FragRatio,omitnil,omitempty" name:"FragRatio"`

	// 行数。
	TableRows *int64 `json:"TableRows,omitnil,omitempty" name:"TableRows"`

	// 表对应的独立物理文件大小（MB）。
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitnil,omitempty" name:"PhysicalFileSize"`
}

type TableSpaceTimeSeries struct {
	// 表名。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 库名。
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// 库表的存储引擎。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 单位时间间隔内的空间指标数据。
	SeriesData *MonitorFloatMetricSeriesData `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`
}

type TimeSlice struct {
	// 总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 统计开始时间
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type UserProfile struct {
	// 配置的id。
	ProfileId *string `json:"ProfileId,omitnil,omitempty" name:"ProfileId"`

	// 配置类型。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 配置级别，"User"或"Instance"。
	ProfileLevel *string `json:"ProfileLevel,omitnil,omitempty" name:"ProfileLevel"`

	// 配置名称。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 配置详情。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`
}