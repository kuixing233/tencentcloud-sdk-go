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

package v20190924

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 数据库错误。
	FAILEDOPERATION_DBERROR = "FailedOperation.DbError"

	// 依赖服务异常。
	FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"

	// Core 返回内容为空
	FAILEDOPERATION_EMPTYCOREBODY = "FailedOperation.EmptyCoreBody"

	// 获取数据库数据错误。
	FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"

	// Tcr实例请求无效的Hearder类型。
	FAILEDOPERATION_ERRORTCRINVALIDMEDIATYPE = "FailedOperation.ErrorTcrInvalidMediaType"

	// Tcr实例资源冲突。
	FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"

	// 没有Tcr操作权限。
	FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"

	// 获取数据库数据错误。
	FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"

	// 获取安全组策略错误。
	FAILEDOPERATION_GETSECURITYPOLICYFAIL = "FailedOperation.GetSecurityPolicyFail"

	// 获取TcrClient错误。
	FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"

	// 操作取消。
	FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"

	// 先决条件不满足。
	FAILEDOPERATION_PRECONDITIONFAILED = "FailedOperation.PreconditionFailed"

	// 交易失败。
	FAILEDOPERATION_TRADEFAILED = "FailedOperation.TradeFailed"

	// 校验仓库名称失败。
	FAILEDOPERATION_VALIDATEREGISTRYNAMEFAIL = "FailedOperation.ValidateRegistryNameFail"

	// 校验支持地域错误。
	FAILEDOPERATION_VALIDATESUPPORTEDREGIONFAIL = "FailedOperation.ValidateSupportedRegionFail"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 创建私有域失败。
	INTERNALERROR_CREATEPRIVATEZONE = "InternalError.CreatePrivateZone"

	// 创建私有域记录失败。
	INTERNALERROR_CREATEPRIVATEZONERECORD = "InternalError.CreatePrivateZoneRecord"

	// 数据库错误。
	INTERNALERROR_DBERROR = "InternalError.DbError"

	// 删除私有域记录失败。
	INTERNALERROR_DELETEPRIVATEZONERECORD = "InternalError.DeletePrivateZoneRecord"

	// 查询vpc私有解析状态失败。
	INTERNALERROR_DESCRIBEINTERNALENDPOINTDNSSTATUS = "InternalError.DescribeInternalEndpointDnsStatus"

	// 查询私有域列表失败。
	INTERNALERROR_DESCRIBEPRIVATEZONELIST = "InternalError.DescribePrivateZoneList"

	// 查询私有域记录列表失败。
	INTERNALERROR_DESCRIBEPRIVATEZONERECORDLIST = "InternalError.DescribePrivateZoneRecordList"

	// 查询开白vpc列表失败。
	INTERNALERROR_DESCRIBEPRIVATEZONESERVICELIST = "InternalError.DescribePrivateZoneServiceList"

	// 目标冲突。
	INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"

	// 目标不存在。
	INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"

	// 鉴权失败。
	INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"

	// 资源已存在。
	INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"

	// 资源超过配额。
	INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"

	// Tcr实例内部错误。
	INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"

	// Tcr实例请求无效的Hearder类型。
	INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"

	// Tcr实例资源冲突。
	INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"

	// 没有Tcr操作权限。
	INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"

	// 修改vpc与私有域关联关系失败。
	INTERNALERROR_MODIFYPRIVATEZONEVPC = "InternalError.ModifyPrivateZoneVpc"

	// 未知错误。
	INTERNALERROR_UNKNOWN = "InternalError.Unknown"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 用户请求中的信息与其namespace不匹配。
	INVALIDPARAMETER_ERRNSMISMATCH = "InvalidParameter.ErrNSMisMatch"

	// 命名空间名称已经存在。
	INVALIDPARAMETER_ERRNAMESPACEEXIST = "InvalidParameter.ErrNamespaceExist"

	// 命名空间已被占用。
	INVALIDPARAMETER_ERRNAMESPACERESERVED = "InvalidParameter.ErrNamespaceReserved"

	// 无效的参数，仓库已存在。
	INVALIDPARAMETER_ERRREPOEXIST = "InvalidParameter.ErrRepoExist"

	// 参数过多。
	INVALIDPARAMETER_ERRTOOLARGE = "InvalidParameter.ErrTooLarge"

	// 触发器名称已存在。
	INVALIDPARAMETER_ERRTRIGGEREXIST = "InvalidParameter.ErrTriggerExist"

	// 用户已经存在。
	INVALIDPARAMETER_ERRUSEREXIST = "InvalidParameter.ErrUserExist"

	// 实例名称已存在。
	INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"

	// 实例名称非法。
	INVALIDPARAMETER_ERRORNAMEILLEGAL = "InvalidParameter.ErrorNameIllegal"

	// 实例名称已保留。
	INVALIDPARAMETER_ERRORNAMERESERVED = "InvalidParameter.ErrorNameReserved"

	// 实例名称非法，格式不正确或者已保留。
	INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"

	// 云标签超过10个上线。
	INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"

	// 无效的TCR请求。
	INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"

	// 复制实例已存在。
	INVALIDPARAMETER_REPLICATIONEXISTS = "InvalidParameter.ReplicationExists"

	// 该地域不支持创建实例。
	INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"

	// 用户命名空间达到配额。
	LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"

	// 用户仓库已经达到最大配额。
	LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"

	// 触发器达到配额。
	LIMITEXCEEDED_ERRTRIGGERMAXLIMIT = "LimitExceeded.ErrTriggerMaxLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少参数。
	MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 配额超限。
	OPERATIONDENIED_QUOTAOVERLIMIT = "OperationDenied.QuotaOverLimit"

	// 实例状态异常。
	RESOURCEINSUFFICIENT_ERRORINSTANCENOTRUNNING = "ResourceInsufficient.ErrorInstanceNotRunning"

	// Vpc dsn解析状态异常或未删除。
	RESOURCEINSUFFICIENT_ERRORVPCDNSSTATUS = "ResourceInsufficient.ErrorVpcDnsStatus"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 用户没有创建命名空间。
	RESOURCENOTFOUND_ERRNONAMESPACE = "ResourceNotFound.ErrNoNamespace"

	// 仓库不存在。
	RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"

	// tag不存在。
	RESOURCENOTFOUND_ERRNOTAG = "ResourceNotFound.ErrNoTag"

	// 触发器不存在。
	RESOURCENOTFOUND_ERRNOTRIGGER = "ResourceNotFound.ErrNoTrigger"

	// 用户不存在（未注册）。
	RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"

	// Tcr实例中的资源未找到。
	RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未授权操作
	UNAUTHORIZEDOPERATION_ERRORTCRUNAUTHORIZED = "UnauthorizedOperation.ErrorTcrUnauthorized"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 创建私有域失败。
	UNSUPPORTEDOPERATION_CREATEPRIVATEZONE = "UnsupportedOperation.CreatePrivateZone"

	// 创建私有域记录失败。
	UNSUPPORTEDOPERATION_CREATEPRIVATEZONERECORD = "UnsupportedOperation.CreatePrivateZoneRecord"

	// 删除私有域记录失败。
	UNSUPPORTEDOPERATION_DELETEPRIVATEZONERECORD = "UnsupportedOperation.DeletePrivateZoneRecord"

	// 查询私有域列表失败。
	UNSUPPORTEDOPERATION_DESCRIBEPRIVATEZONELIST = "UnsupportedOperation.DescribePrivateZoneList"

	// 查询私有域记录列表失败。
	UNSUPPORTEDOPERATION_DESCRIBEPRIVATEZONERECORDLIST = "UnsupportedOperation.DescribePrivateZoneRecordList"

	// 个人版实例用户未初始化。
	UNSUPPORTEDOPERATION_ERRNOUSERINITIALIZED = "UnsupportedOperation.ErrNoUserInitialized"

	// 修改私有域记录失败。
	UNSUPPORTEDOPERATION_MODIFYPRIVATEZONERECORD = "UnsupportedOperation.ModifyPrivateZoneRecord"

	// 修改vpc与私有域关联关系失败。
	UNSUPPORTEDOPERATION_MODIFYPRIVATEZONEVPC = "UnsupportedOperation.ModifyPrivateZoneVpc"
)
