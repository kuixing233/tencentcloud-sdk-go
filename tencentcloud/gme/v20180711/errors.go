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

package v20180711

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 登录态过期。
	FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"

	// 欠费不可操作。
	FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 回调地址不正确
	INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"

	// 日期无效。
	INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"

	// 输入日期超出60天查询范围。
	INVALIDPARAMETER_DATEOUTOFSIXTYDAYS = "InvalidParameter.DateOutOfSixtyDays"

	// 标签不正确
	INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"

	// 查询时间范围错误。
	INVALIDPARAMETER_TIMERANGEERROR = "InvalidParameter.TimeRangeError"

	// BizId 参数错误
	INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"

	// 输入删除类型应为1或2。
	INVALIDPARAMETERVALUE_INVALIDDELETETYPE = "InvalidParameterValue.InvalidDeleteType"

	// RecordMode参数错误
	INVALIDPARAMETERVALUE_INVALIDRECORDMODE = "InvalidParameterValue.InvalidRecordMode"

	// RoomId 参数错误
	INVALIDPARAMETERVALUE_INVALIDROOMID = "InvalidParameterValue.InvalidRoomId"

	// 不需要填写StrUid，请填写Uid
	INVALIDPARAMETERVALUE_INVALIDSTRUIN = "InvalidParameterValue.InvalidStrUin"

	// 黑白名单格式错误
	INVALIDPARAMETERVALUE_INVALIDSUBSCRIBERECORDUSERIDS = "InvalidParameterValue.InvalidSubscribeRecordUserIds"

	// 白名单个数超过20个
	INVALIDPARAMETERVALUE_INVALIDSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidSubscribeUserIds"

	// taskid参数错误
	INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"

	// 黑名单个数超过20个
	INVALIDPARAMETERVALUE_INVALIDUNSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidUNSubscribeUserIds"

	// 不能同时输入uid和strUid
	INVALIDPARAMETERVALUE_INVALIDUINORSTRUIN = "InvalidParameterValue.InvalidUinOrStrUin"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 创建应用数已达上限。
	LIMITEXCEEDED_APPLICATION = "LimitExceeded.Application"

	// 缺少参数。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// task已存在
	RESOURCEINUSE_TASKINUSE = "ResourceInUse.TaskInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 应用ID不正确
	RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"

	// 房间不存在
	RESOURCENOTFOUND_ROOMNOTFOUND = "ResourceNotFound.RoomNotFound"

	// 任务ID不存在
	RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 创建应用不被授权。
	UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"

	// 该用户未进行实名认证。
	UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 超过最大踢人数量
	UNSUPPORTEDOPERATION_LIMITDELETEEXCEEDED = "UnsupportedOperation.LimitDeleteExceeded"

	// 语音转文本开关未开启。
	UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"

	// 录制服务未开通
	UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
)
