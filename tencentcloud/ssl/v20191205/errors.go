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

package v20191205

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作未授权。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 请检查是否有权限。
	FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"

	// 有未解绑的云资源，不能被删除。
	FAILEDOPERATION_BOUNDRESOURCES = "FailedOperation.BoundResources"

	// CAM鉴权出现错误。
	FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"

	// 由于未知原因证书取消失败，请您刷新页面后重试
	FAILEDOPERATION_CANCELAUDITCERTIFICATEFAILED = "FailedOperation.CancelAuditCertificateFailed"

	// 取消订单失败。
	FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"

	// 该证书已颁发， 不能删除。
	FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"

	// 免费证书申请1小时内不允许删除。
	FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"

	// 获取订单信息失败，请稍后重试。
	FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"

	// 温馨提示：上传的证书内容不符合CA证书的标准格式，请您核对证书类型是否有误
	FAILEDOPERATION_CERTIFICATECAERROR = "FailedOperation.CertificateCaError"

	// 多年期证书非首年无法终止审核，如您暂时不用证书，无需验证域名即可
	FAILEDOPERATION_CERTIFICATECANCELFAILEDMULTI = "FailedOperation.CertificateCancelFailedMulti"

	// 温馨提示：证书链检测异常，请核对每段证书是否存在异常
	FAILEDOPERATION_CERTIFICATECHAINERROR = "FailedOperation.CertificateChainError"

	// 记录状态必须完结才可以执行该操作。
	FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"

	// 证书部署有正在进行中的任务，请部署完成后再重试。
	FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"

	// 已选择的云资源无实例，无法更新，请重新核对后重试。
	FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEEMPTY = "FailedOperation.CertificateDeployInstanceEmpty"

	// 部署格式错误
	FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEFORMATERROR = "FailedOperation.CertificateDeployInstanceFormatError"

	// 证书部署记录不存在。
	FAILEDOPERATION_CERTIFICATEDEPLOYNOTEXIST = "FailedOperation.CertificateDeployNotExist"

	// 记录状态必须失败才可以执行该操作。
	FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"

	// 必须有部署成功的记录才可以回滚。
	FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"

	// 证书下载服务类型不支持。
	FAILEDOPERATION_CERTIFICATEDOWNLOADSERVICETYPENOTSUPPORT = "FailedOperation.CertificateDownloadServiceTypeNotSupport"

	// 证书内容疑似国密证书，与所选证书标准不符，请您核对选择的证书标准和证书内容是否有误
	FAILEDOPERATION_CERTIFICATEENCRYPTINVALID = "FailedOperation.CertificateEncryptInvalid"

	// 证书已存在。
	FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"

	// 当前证书不允许使用一键更新的功能。
	FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"

	// 当前为内部账号，账号涉及实例资源较多，无法使用部署功能，请联系SSL证书特殊处理。
	FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"

	// 当前账号下实例数量过多，无法正常加载，请您切换加载方式。切换后点击“刷新列表”等待一段时间后即可全部加载。
	FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"

	// 云资源类型无效。
	FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"

	// 邮箱不能为空
	FAILEDOPERATION_CERTIFICATEINFOSUBMITCONTACTEMAILEMPTY = "FailedOperation.CertificateInfoSubmitContactEmailEmpty"

	// 您输入的域名%s格式有误，请您核对后重新提交
	FAILEDOPERATION_CERTIFICATEINFOSUBMITDOMAININVALID = "FailedOperation.CertificateInfoSubmitDomainInvalid"

	// 当前证书状态不可以提交资料
	FAILEDOPERATION_CERTIFICATEINFOSUBMITSTATUSINVALID = "FailedOperation.CertificateInfoSubmitStatusInvalid"

	// 证书不符合标准。
	FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"

	// 温馨提示：证书内容和私钥不匹配，请您核对（请留意是否有多余的空格）
	FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"

	// 证书与私钥不对应。
	FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"

	// 该证书已设置成不允许下载。
	FAILEDOPERATION_CERTIFICATENOTALLOWDOWNLOAD = "FailedOperation.CertificateNotAllowDownload"

	// 证书不可用，请检查后再试。
	FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"

	// 证书不可以部署到实例列表下。
	FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"

	// 证书不存在。
	FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"

	// 该证书不存在或已签发。
	FAILEDOPERATION_CERTIFICATENOTFOUNDORAUDITED = "FailedOperation.CertificateNotFoundOrAudited"

	// 该证书不存在或不可取消。
	FAILEDOPERATION_CERTIFICATENOTFOUNDORCANTCANCEL = "FailedOperation.CertificateNotFoundOrCantCancel"

	// 订单正在处理中，请您耐心等待
	FAILEDOPERATION_CERTIFICATEORDERSUBMITPROCESSING = "FailedOperation.CertificateOrderSubmitProcessing"

	// 当前选择的根证书无效，请您重新选择
	FAILEDOPERATION_CERTIFICATEORDERSUBMITROOTCAINVALID = "FailedOperation.CertificateOrderSubmitRootCAInvalid"

	// 解析失败，请检查证书是否符合标准，请留意是否有多余的空格
	FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"

	// 您输入的邮箱格式有误，请您核对后重新提交
	FAILEDOPERATION_CERTIFICATEPREVERIFYEMAILINVALID = "FailedOperation.CertificatePreVerifyEmailInvalid"

	// 当前证书状态不允许下载。
	FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWDOWNLOAD = "FailedOperation.CertificateStatusNotAllowDownload"

	// 证书状态不允许重新申请。
	FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWRESUBMIT = "FailedOperation.CertificateStatusNotAllowResubmit"

	// 证书当前状态不允许进行域名验证。
	FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWVERIFY = "FailedOperation.CertificateStatusNotAllowVerify"

	// 上传证书的托管续费证书不能手动提交订单
	FAILEDOPERATION_CERTIFICATESUBMITHOSTINGCERTERROR = "FailedOperation.CertificateSubmitHostingCertError"

	// 证书操作参数无效，证书ID数量超过100。
	FAILEDOPERATION_CERTIFICATESYNCTASKCERTIFICATEIDCOUNTINVALID = "FailedOperation.CertificateSyncTaskCertificateIdCountInvalid"

	// 证书关联云资源查询任务ID无效
	FAILEDOPERATION_CERTIFICATESYNCTASKIDINVALID = "FailedOperation.CertificateSyncTaskIdInvalid"

	// 无法查到当前证书类型详情。
	FAILEDOPERATION_CERTIFICATETYPEINFONOTFOUND = "FailedOperation.CertificateTypeInfoNotFound"

	// 当前为白名单功能，非白名单用户无法使用该功能，请联系SSL证书特殊处理。
	FAILEDOPERATION_CERTIFICATEWHITEFUNCERROR = "FailedOperation.CertificateWhiteFuncError"

	// 证书确认函文件过大（需小于1.4M）。
	FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"

	// 证书确认函文件过小（需大于1KB）。
	FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"

	// 免费证书申请时间未超过1小时，不能被删除。
	FAILEDOPERATION_DELETEFAILEDTIMENOTUP = "FailedOperation.DeleteFailedTimeNotUp"

	// 证书已关联云资源，无法删除。
	FAILEDOPERATION_DELETERESOURCEFAILED = "FailedOperation.DeleteResourceFailed"

	// 免费证书数量超出限制。
	FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"

	// 文件尺寸太大（需小于1.4MB）。
	FAILEDOPERATION_FILETOOLARGE = "FailedOperation.FileTooLarge"

	// 文件尺寸太小，请上传清晰图片。
	FAILEDOPERATION_FILETOOSMALL = "FailedOperation.FileTooSmall"

	// 生产CSR失败
	FAILEDOPERATION_GENCSRFAIL = "FailedOperation.GenCSRFail"

	// 公司管理人状态错误。
	FAILEDOPERATION_ILLEGALMANAGERSTATUS = "FailedOperation.IllegalManagerStatus"

	// 证书来源错误。
	FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"

	// 证书状态不正确。
	FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"

	// 证书确认函格式错误（支持格式为jpg、jpeg、png、pdf）。
	FAILEDOPERATION_INVALIDCONFIRMLETTERFORMAT = "FailedOperation.InvalidConfirmLetterFormat"

	// 证书确认函格式错误（支持格式为jpg、pdf、gif）。
	FAILEDOPERATION_INVALIDCONFIRMLETTERFORMATWOSIGN = "FailedOperation.InvalidConfirmLetterFormatWosign"

	// 文件格式错误，请重新上传。
	FAILEDOPERATION_INVALIDFILETYPE = "FailedOperation.InvalidFileType"

	// 参数有误。
	FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"

	// 该主域（%s）下申请的免费证书数量已达到%s个上限，请购买付费证书。
	FAILEDOPERATION_MAINDOMAINCERTIFICATECOUNTLIMIT = "FailedOperation.MainDomainCertificateCountLimit"

	// 管理人信息已提交CA，不可以删除。
	FAILEDOPERATION_MANAGERCANNOTDELETECA = "FailedOperation.ManagerCanNotDeleteCa"

	// 管理人信息已关联证书，不可以删除。
	FAILEDOPERATION_MANAGERCANNOTDELETECERT = "FailedOperation.ManagerCanNotDeleteCert"

	// 当前 CA 机构访问繁忙，请稍后重试。
	FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"

	// 您没有该项目的操作权限。
	FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"

	// 尚未通过实名认证。
	FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"

	// 该订单已重签发。
	FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"

	// 重颁发失败。
	FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"

	// 剩余权益点不足。
	FAILEDOPERATION_PACKAGECOUNTLIMIT = "FailedOperation.PackageCountLimit"

	// 权益包已过期。
	FAILEDOPERATION_PACKAGEEXPIRED = "FailedOperation.PackageExpired"

	// 权益包不存在。
	FAILEDOPERATION_PACKAGENOTFOUND = "FailedOperation.PackageNotFound"

	// 续费证书未颁发，无法执行删除操作。
	FAILEDOPERATION_RENEWNOTISSUED = "FailedOperation.RenewNotIssued"

	// 证书吊销失败。
	FAILEDOPERATION_REVOKEFAILED = "FailedOperation.RevokeFailed"

	// 证书已关联云资源，无法吊销。
	FAILEDOPERATION_REVOKERESOURCEFAILED = "FailedOperation.RevokeResourceFailed"

	// 角色不存在，请前往授权。
	FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"

	// 系统错误。
	FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"

	// 计费中心错误。
	FAILEDOPERATION_TRADEERROR = "FailedOperation.TradeError"

	// 当前接口不支持上传确认函功能，请去腾讯云控制台进行操作
	FAILEDOPERATION_UPLOADCONFIRMCERTIFICATENOTSUPPORT = "FailedOperation.UploadConfirmCertificateNotSupport"

	// 当前证书是DV证书，不支持上传确认函
	FAILEDOPERATION_UPLOADCONFIRMCERTIFICATENOTSUPPORTDV = "FailedOperation.UploadConfirmCertificateNotSupportDV"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 后端服务响应为空。
	INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"

	// 后端服务响应错误。
	INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 查询的证书ID不能超过50个。
	INVALIDPARAMETER_CERTIFICATEIDNUMBERLIMIT = "InvalidParameter.CertificateIdNumberLimit"

	// 当前证书状态不允许重新提交申请。
	INVALIDPARAMETER_CERTIFICATESTATUSNOTALLOWRESUBMIT = "InvalidParameter.CertificateStatusNotAllowResubmit"

	// 证书数量超出限制。
	INVALIDPARAMETER_CERTIFICATESNUMBEREXCEEDED = "InvalidParameter.CertificatesNumberExceeded"

	// 包含无效的证书ID。
	INVALIDPARAMETER_CONTAINSINVALIDCERTIFICATEID = "InvalidParameter.ContainsInvalidCertificateId"

	// 域名数量无效。
	INVALIDPARAMETER_DOMAINCOUNTINVALID = "InvalidParameter.DomainCountInvalid"

	// 域名无效，请重新输入。
	INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"

	// 没有匹配的私钥。
	INVALIDPARAMETER_MISSMATCHPRIVATEKEY = "InvalidParameter.MissMatchPrivateKey"

	// 权益点ID列表无效。
	INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"

	// 证书期限无效。
	INVALIDPARAMETER_PERIODINVALID = "InvalidParameter.PeriodInvalid"

	// 产品PID无效。
	INVALIDPARAMETER_PRODUCTPIDINVALID = "InvalidParameter.ProductPidInvalid"

	// Region不允许为空
	INVALIDPARAMETER_REGIONNOTEMPTY = "InvalidParameter.RegionNotEmpty"

	// 算法无效。
	INVALIDPARAMETER_RENEWALGORITHMINVALID = "InvalidParameter.RenewAlgorithmInvalid"

	// CSR无效。
	INVALIDPARAMETER_RENEWCSRINVALID = "InvalidParameter.RenewCsrInvalid"

	// 生成CSR方式无效。
	INVALIDPARAMETER_RENEWGENCSRMETHODINVALID = "InvalidParameter.RenewGenCsrMethodInvalid"

	// 参数有误。
	INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 邮箱格式有误，请您重新输入
	INVALIDPARAMETERVALUE_CERTIFICATEEMAILPARSEINVALID = "InvalidParameterValue.CertificateEmailParseInvalid"

	// 中间证书不一致
	INVALIDPARAMETERVALUE_INTERMEDIATECERTNOTSAME = "InvalidParameterValue.IntermediateCertNotSame"

	// 单位时间内接口请求频率达到限制。
	LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 公司管理人不存在。
	RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"
)
