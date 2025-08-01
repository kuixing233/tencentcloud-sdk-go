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

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-12-05"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewApplyCertificateRequest() (request *ApplyCertificateRequest) {
    request = &ApplyCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ApplyCertificate")
    
    
    return
}

func NewApplyCertificateResponse() (response *ApplyCertificateResponse) {
    response = &ApplyCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyCertificate
// 本接口（ApplyCertificate）用于免费证书申请。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_GENCSRFAIL = "FailedOperation.GenCSRFail"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_MAINDOMAINCERTIFICATECOUNTLIMIT = "FailedOperation.MainDomainCertificateCountLimit"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_PACKAGECOUNTLIMIT = "FailedOperation.PackageCountLimit"
//  FAILEDOPERATION_PACKAGEEXPIRED = "FailedOperation.PackageExpired"
//  FAILEDOPERATION_PACKAGENOTFOUND = "FailedOperation.PackageNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"
//  INVALIDPARAMETERVALUE_CERTIFICATEEMAILPARSEINVALID = "InvalidParameterValue.CertificateEmailParseInvalid"
func (c *Client) ApplyCertificate(request *ApplyCertificateRequest) (response *ApplyCertificateResponse, err error) {
    return c.ApplyCertificateWithContext(context.Background(), request)
}

// ApplyCertificate
// 本接口（ApplyCertificate）用于免费证书申请。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_GENCSRFAIL = "FailedOperation.GenCSRFail"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_MAINDOMAINCERTIFICATECOUNTLIMIT = "FailedOperation.MainDomainCertificateCountLimit"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_PACKAGECOUNTLIMIT = "FailedOperation.PackageCountLimit"
//  FAILEDOPERATION_PACKAGEEXPIRED = "FailedOperation.PackageExpired"
//  FAILEDOPERATION_PACKAGENOTFOUND = "FailedOperation.PackageNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"
//  INVALIDPARAMETERVALUE_CERTIFICATEEMAILPARSEINVALID = "InvalidParameterValue.CertificateEmailParseInvalid"
func (c *Client) ApplyCertificateWithContext(ctx context.Context, request *ApplyCertificateRequest) (response *ApplyCertificateResponse, err error) {
    if request == nil {
        request = NewApplyCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "ApplyCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCancelAuditCertificateRequest() (request *CancelAuditCertificateRequest) {
    request = &CancelAuditCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CancelAuditCertificate")
    
    
    return
}

func NewCancelAuditCertificateResponse() (response *CancelAuditCertificateResponse) {
    response = &CancelAuditCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelAuditCertificate
// 取消证书审核
//
// 可能返回的错误码:
//  FAILEDOPERATION_CANCELAUDITCERTIFICATEFAILED = "FailedOperation.CancelAuditCertificateFailed"
//  FAILEDOPERATION_CERTIFICATECANCELFAILEDMULTI = "FailedOperation.CertificateCancelFailedMulti"
//  FAILEDOPERATION_CERTIFICATENOTFOUNDORCANTCANCEL = "FailedOperation.CertificateNotFoundOrCantCancel"
func (c *Client) CancelAuditCertificate(request *CancelAuditCertificateRequest) (response *CancelAuditCertificateResponse, err error) {
    return c.CancelAuditCertificateWithContext(context.Background(), request)
}

// CancelAuditCertificate
// 取消证书审核
//
// 可能返回的错误码:
//  FAILEDOPERATION_CANCELAUDITCERTIFICATEFAILED = "FailedOperation.CancelAuditCertificateFailed"
//  FAILEDOPERATION_CERTIFICATECANCELFAILEDMULTI = "FailedOperation.CertificateCancelFailedMulti"
//  FAILEDOPERATION_CERTIFICATENOTFOUNDORCANTCANCEL = "FailedOperation.CertificateNotFoundOrCantCancel"
func (c *Client) CancelAuditCertificateWithContext(ctx context.Context, request *CancelAuditCertificateRequest) (response *CancelAuditCertificateResponse, err error) {
    if request == nil {
        request = NewCancelAuditCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CancelAuditCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelAuditCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelAuditCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCancelCertificateOrderRequest() (request *CancelCertificateOrderRequest) {
    request = &CancelCertificateOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CancelCertificateOrder")
    
    
    return
}

func NewCancelCertificateOrderResponse() (response *CancelCertificateOrderResponse) {
    response = &CancelCertificateOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelCertificateOrder
// 取消证书订单。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) CancelCertificateOrder(request *CancelCertificateOrderRequest) (response *CancelCertificateOrderResponse, err error) {
    return c.CancelCertificateOrderWithContext(context.Background(), request)
}

// CancelCertificateOrder
// 取消证书订单。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) CancelCertificateOrderWithContext(ctx context.Context, request *CancelCertificateOrderRequest) (response *CancelCertificateOrderResponse, err error) {
    if request == nil {
        request = NewCancelCertificateOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CancelCertificateOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelCertificateOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelCertificateOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCertificateInfoSubmitRequest() (request *CertificateInfoSubmitRequest) {
    request = &CertificateInfoSubmitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CertificateInfoSubmit")
    
    
    return
}

func NewCertificateInfoSubmitResponse() (response *CertificateInfoSubmitResponse) {
    response = &CertificateInfoSubmitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CertificateInfoSubmit
// 付费提交证书资料
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINFOSUBMITCONTACTEMAILEMPTY = "FailedOperation.CertificateInfoSubmitContactEmailEmpty"
//  FAILEDOPERATION_CERTIFICATEINFOSUBMITDOMAININVALID = "FailedOperation.CertificateInfoSubmitDomainInvalid"
//  FAILEDOPERATION_CERTIFICATEINFOSUBMITSTATUSINVALID = "FailedOperation.CertificateInfoSubmitStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEORDERSUBMITROOTCAINVALID = "FailedOperation.CertificateOrderSubmitRootCAInvalid"
//  FAILEDOPERATION_CERTIFICATESUBMITHOSTINGCERTERROR = "FailedOperation.CertificateSubmitHostingCertError"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RENEWALGORITHMINVALID = "InvalidParameter.RenewAlgorithmInvalid"
func (c *Client) CertificateInfoSubmit(request *CertificateInfoSubmitRequest) (response *CertificateInfoSubmitResponse, err error) {
    return c.CertificateInfoSubmitWithContext(context.Background(), request)
}

// CertificateInfoSubmit
// 付费提交证书资料
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINFOSUBMITCONTACTEMAILEMPTY = "FailedOperation.CertificateInfoSubmitContactEmailEmpty"
//  FAILEDOPERATION_CERTIFICATEINFOSUBMITDOMAININVALID = "FailedOperation.CertificateInfoSubmitDomainInvalid"
//  FAILEDOPERATION_CERTIFICATEINFOSUBMITSTATUSINVALID = "FailedOperation.CertificateInfoSubmitStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEORDERSUBMITROOTCAINVALID = "FailedOperation.CertificateOrderSubmitRootCAInvalid"
//  FAILEDOPERATION_CERTIFICATESUBMITHOSTINGCERTERROR = "FailedOperation.CertificateSubmitHostingCertError"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RENEWALGORITHMINVALID = "InvalidParameter.RenewAlgorithmInvalid"
func (c *Client) CertificateInfoSubmitWithContext(ctx context.Context, request *CertificateInfoSubmitRequest) (response *CertificateInfoSubmitResponse, err error) {
    if request == nil {
        request = NewCertificateInfoSubmitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CertificateInfoSubmit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CertificateInfoSubmit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCertificateInfoSubmitResponse()
    err = c.Send(request, response)
    return
}

func NewCertificateOrderSubmitRequest() (request *CertificateOrderSubmitRequest) {
    request = &CertificateOrderSubmitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CertificateOrderSubmit")
    
    
    return
}

func NewCertificateOrderSubmitResponse() (response *CertificateOrderSubmitResponse) {
    response = &CertificateOrderSubmitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CertificateOrderSubmit
// 提交付费证书订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINFOSUBMITSTATUSINVALID = "FailedOperation.CertificateInfoSubmitStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEORDERSUBMITPROCESSING = "FailedOperation.CertificateOrderSubmitProcessing"
//  FAILEDOPERATION_CERTIFICATEPREVERIFYEMAILINVALID = "FailedOperation.CertificatePreVerifyEmailInvalid"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
func (c *Client) CertificateOrderSubmit(request *CertificateOrderSubmitRequest) (response *CertificateOrderSubmitResponse, err error) {
    return c.CertificateOrderSubmitWithContext(context.Background(), request)
}

// CertificateOrderSubmit
// 提交付费证书订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINFOSUBMITSTATUSINVALID = "FailedOperation.CertificateInfoSubmitStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEORDERSUBMITPROCESSING = "FailedOperation.CertificateOrderSubmitProcessing"
//  FAILEDOPERATION_CERTIFICATEPREVERIFYEMAILINVALID = "FailedOperation.CertificatePreVerifyEmailInvalid"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
func (c *Client) CertificateOrderSubmitWithContext(ctx context.Context, request *CertificateOrderSubmitRequest) (response *CertificateOrderSubmitResponse, err error) {
    if request == nil {
        request = NewCertificateOrderSubmitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CertificateOrderSubmit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CertificateOrderSubmit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCertificateOrderSubmitResponse()
    err = c.Send(request, response)
    return
}

func NewCheckCertificateChainRequest() (request *CheckCertificateChainRequest) {
    request = &CheckCertificateChainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CheckCertificateChain")
    
    
    return
}

func NewCheckCertificateChainResponse() (response *CheckCertificateChainResponse) {
    response = &CheckCertificateChainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckCertificateChain
// 本接口（CheckCertificateChain）用于检查证书链是否完整。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) CheckCertificateChain(request *CheckCertificateChainRequest) (response *CheckCertificateChainResponse, err error) {
    return c.CheckCertificateChainWithContext(context.Background(), request)
}

// CheckCertificateChain
// 本接口（CheckCertificateChain）用于检查证书链是否完整。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) CheckCertificateChainWithContext(ctx context.Context, request *CheckCertificateChainRequest) (response *CheckCertificateChainResponse, err error) {
    if request == nil {
        request = NewCheckCertificateChainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CheckCertificateChain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckCertificateChain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckCertificateChainResponse()
    err = c.Send(request, response)
    return
}

func NewCheckCertificateDomainVerificationRequest() (request *CheckCertificateDomainVerificationRequest) {
    request = &CheckCertificateDomainVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CheckCertificateDomainVerification")
    
    
    return
}

func NewCheckCertificateDomainVerificationResponse() (response *CheckCertificateDomainVerificationResponse) {
    response = &CheckCertificateDomainVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckCertificateDomainVerification
// 检查证书域名验证结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWVERIFY = "FailedOperation.CertificateStatusNotAllowVerify"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
func (c *Client) CheckCertificateDomainVerification(request *CheckCertificateDomainVerificationRequest) (response *CheckCertificateDomainVerificationResponse, err error) {
    return c.CheckCertificateDomainVerificationWithContext(context.Background(), request)
}

// CheckCertificateDomainVerification
// 检查证书域名验证结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWVERIFY = "FailedOperation.CertificateStatusNotAllowVerify"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
func (c *Client) CheckCertificateDomainVerificationWithContext(ctx context.Context, request *CheckCertificateDomainVerificationRequest) (response *CheckCertificateDomainVerificationResponse, err error) {
    if request == nil {
        request = NewCheckCertificateDomainVerificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CheckCertificateDomainVerification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckCertificateDomainVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckCertificateDomainVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckCertificateExistRequest() (request *CheckCertificateExistRequest) {
    request = &CheckCertificateExistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CheckCertificateExist")
    
    
    return
}

func NewCheckCertificateExistResponse() (response *CheckCertificateExistResponse) {
    response = &CheckCertificateExistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckCertificateExist
// 根据证书内容检测当前账号下是否存在一致的证书， 存在则返回证书ID， 不存在则返回空
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWVERIFY = "FailedOperation.CertificateStatusNotAllowVerify"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
func (c *Client) CheckCertificateExist(request *CheckCertificateExistRequest) (response *CheckCertificateExistResponse, err error) {
    return c.CheckCertificateExistWithContext(context.Background(), request)
}

// CheckCertificateExist
// 根据证书内容检测当前账号下是否存在一致的证书， 存在则返回证书ID， 不存在则返回空
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWVERIFY = "FailedOperation.CertificateStatusNotAllowVerify"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
func (c *Client) CheckCertificateExistWithContext(ctx context.Context, request *CheckCertificateExistRequest) (response *CheckCertificateExistResponse, err error) {
    if request == nil {
        request = NewCheckCertificateExistRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CheckCertificateExist")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckCertificateExist require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckCertificateExistResponse()
    err = c.Send(request, response)
    return
}

func NewCommitCertificateInformationRequest() (request *CommitCertificateInformationRequest) {
    request = &CommitCertificateInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CommitCertificateInformation")
    
    
    return
}

func NewCommitCertificateInformationResponse() (response *CommitCertificateInformationResponse) {
    response = &CommitCertificateInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CommitCertificateInformation
// 付费证书提交订单； 本接口不维护新功能， 可使用新接口进行提交， [CertificateOrderSubmit](https://cloud.tencent.com/document/product/400/116032)
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
func (c *Client) CommitCertificateInformation(request *CommitCertificateInformationRequest) (response *CommitCertificateInformationResponse, err error) {
    return c.CommitCertificateInformationWithContext(context.Background(), request)
}

// CommitCertificateInformation
// 付费证书提交订单； 本接口不维护新功能， 可使用新接口进行提交， [CertificateOrderSubmit](https://cloud.tencent.com/document/product/400/116032)
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
func (c *Client) CommitCertificateInformationWithContext(ctx context.Context, request *CommitCertificateInformationRequest) (response *CommitCertificateInformationResponse, err error) {
    if request == nil {
        request = NewCommitCertificateInformationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CommitCertificateInformation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CommitCertificateInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCommitCertificateInformationResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteCertificateRequest() (request *CompleteCertificateRequest) {
    request = &CompleteCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CompleteCertificate")
    
    
    return
}

func NewCompleteCertificateResponse() (response *CompleteCertificateResponse) {
    response = &CompleteCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CompleteCertificate
// 本接口（CompleteCertificate）用于主动触发证书验证。DNSPod和Wotrus品牌的证书不支持使用此接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUNDORAUDITED = "FailedOperation.CertificateNotFoundOrAudited"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
func (c *Client) CompleteCertificate(request *CompleteCertificateRequest) (response *CompleteCertificateResponse, err error) {
    return c.CompleteCertificateWithContext(context.Background(), request)
}

// CompleteCertificate
// 本接口（CompleteCertificate）用于主动触发证书验证。DNSPod和Wotrus品牌的证书不支持使用此接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUNDORAUDITED = "FailedOperation.CertificateNotFoundOrAudited"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
func (c *Client) CompleteCertificateWithContext(ctx context.Context, request *CompleteCertificateRequest) (response *CompleteCertificateResponse, err error) {
    if request == nil {
        request = NewCompleteCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CompleteCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompleteCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCompleteCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCertificateRequest() (request *CreateCertificateRequest) {
    request = &CreateCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CreateCertificate")
    
    
    return
}

func NewCreateCertificateResponse() (response *CreateCertificateResponse) {
    response = &CreateCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCertificate
// 本接口（CreateCertificate）用于创建付费证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateCertificate(request *CreateCertificateRequest) (response *CreateCertificateResponse, err error) {
    return c.CreateCertificateWithContext(context.Background(), request)
}

// CreateCertificate
// 本接口（CreateCertificate）用于创建付费证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateCertificateWithContext(ctx context.Context, request *CreateCertificateRequest) (response *CreateCertificateResponse, err error) {
    if request == nil {
        request = NewCreateCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CreateCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCertificateBindResourceSyncTaskRequest() (request *CreateCertificateBindResourceSyncTaskRequest) {
    request = &CreateCertificateBindResourceSyncTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CreateCertificateBindResourceSyncTask")
    
    
    return
}

func NewCreateCertificateBindResourceSyncTaskResponse() (response *CreateCertificateBindResourceSyncTaskResponse) {
    response = &CreateCertificateBindResourceSyncTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCertificateBindResourceSyncTask
// 创建证书绑定关联云资源异步任务， 该接口用于查询证书关联云资源。 若证书ID已存在查询云资源任务，则结果返回该任务ID。关联云资源类型，支持以下云资源：clb、cdn、waf、live、vod、ddos、tke、apigateway、tcb、teo（edgeOne）、cos。查询关联云资源结果使用DescribeCertificateBindResourceTaskResult接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
func (c *Client) CreateCertificateBindResourceSyncTask(request *CreateCertificateBindResourceSyncTaskRequest) (response *CreateCertificateBindResourceSyncTaskResponse, err error) {
    return c.CreateCertificateBindResourceSyncTaskWithContext(context.Background(), request)
}

// CreateCertificateBindResourceSyncTask
// 创建证书绑定关联云资源异步任务， 该接口用于查询证书关联云资源。 若证书ID已存在查询云资源任务，则结果返回该任务ID。关联云资源类型，支持以下云资源：clb、cdn、waf、live、vod、ddos、tke、apigateway、tcb、teo（edgeOne）、cos。查询关联云资源结果使用DescribeCertificateBindResourceTaskResult接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
func (c *Client) CreateCertificateBindResourceSyncTaskWithContext(ctx context.Context, request *CreateCertificateBindResourceSyncTaskRequest) (response *CreateCertificateBindResourceSyncTaskResponse, err error) {
    if request == nil {
        request = NewCreateCertificateBindResourceSyncTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CreateCertificateBindResourceSyncTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCertificateBindResourceSyncTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCertificateBindResourceSyncTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCertificateByPackageRequest() (request *CreateCertificateByPackageRequest) {
    request = &CreateCertificateByPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CreateCertificateByPackage")
    
    
    return
}

func NewCreateCertificateByPackageResponse() (response *CreateCertificateByPackageResponse) {
    response = &CreateCertificateByPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCertificateByPackage
// 使用权益点创建证书
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRADEERROR = "FailedOperation.TradeError"
//  INVALIDPARAMETER_DOMAINCOUNTINVALID = "InvalidParameter.DomainCountInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"
//  INVALIDPARAMETER_PERIODINVALID = "InvalidParameter.PeriodInvalid"
//  INVALIDPARAMETER_PRODUCTPIDINVALID = "InvalidParameter.ProductPidInvalid"
//  INVALIDPARAMETER_RENEWALGORITHMINVALID = "InvalidParameter.RenewAlgorithmInvalid"
//  INVALIDPARAMETER_RENEWCSRINVALID = "InvalidParameter.RenewCsrInvalid"
//  INVALIDPARAMETER_RENEWGENCSRMETHODINVALID = "InvalidParameter.RenewGenCsrMethodInvalid"
func (c *Client) CreateCertificateByPackage(request *CreateCertificateByPackageRequest) (response *CreateCertificateByPackageResponse, err error) {
    return c.CreateCertificateByPackageWithContext(context.Background(), request)
}

// CreateCertificateByPackage
// 使用权益点创建证书
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRADEERROR = "FailedOperation.TradeError"
//  INVALIDPARAMETER_DOMAINCOUNTINVALID = "InvalidParameter.DomainCountInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"
//  INVALIDPARAMETER_PERIODINVALID = "InvalidParameter.PeriodInvalid"
//  INVALIDPARAMETER_PRODUCTPIDINVALID = "InvalidParameter.ProductPidInvalid"
//  INVALIDPARAMETER_RENEWALGORITHMINVALID = "InvalidParameter.RenewAlgorithmInvalid"
//  INVALIDPARAMETER_RENEWCSRINVALID = "InvalidParameter.RenewCsrInvalid"
//  INVALIDPARAMETER_RENEWGENCSRMETHODINVALID = "InvalidParameter.RenewGenCsrMethodInvalid"
func (c *Client) CreateCertificateByPackageWithContext(ctx context.Context, request *CreateCertificateByPackageRequest) (response *CreateCertificateByPackageResponse, err error) {
    if request == nil {
        request = NewCreateCertificateByPackageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "CreateCertificateByPackage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCertificateByPackage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCertificateByPackageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCertificateRequest() (request *DeleteCertificateRequest) {
    request = &DeleteCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DeleteCertificate")
    
    
    return
}

func NewDeleteCertificateResponse() (response *DeleteCertificateResponse) {
    response = &DeleteCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCertificate
// 本接口（DeleteCertificate）用于删除证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_BOUNDRESOURCES = "FailedOperation.BoundResources"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_DELETEFAILEDTIMENOTUP = "FailedOperation.DeleteFailedTimeNotUp"
//  FAILEDOPERATION_DELETERESOURCEFAILED = "FailedOperation.DeleteResourceFailed"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_RENEWNOTISSUED = "FailedOperation.RenewNotIssued"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeleteCertificate(request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
    return c.DeleteCertificateWithContext(context.Background(), request)
}

// DeleteCertificate
// 本接口（DeleteCertificate）用于删除证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_BOUNDRESOURCES = "FailedOperation.BoundResources"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_DELETEFAILEDTIMENOTUP = "FailedOperation.DeleteFailedTimeNotUp"
//  FAILEDOPERATION_DELETERESOURCEFAILED = "FailedOperation.DeleteResourceFailed"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_RENEWNOTISSUED = "FailedOperation.RenewNotIssued"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeleteCertificateWithContext(ctx context.Context, request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
    if request == nil {
        request = NewDeleteCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DeleteCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCertificatesRequest() (request *DeleteCertificatesRequest) {
    request = &DeleteCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DeleteCertificates")
    
    
    return
}

func NewDeleteCertificatesResponse() (response *DeleteCertificatesResponse) {
    response = &DeleteCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCertificates
// 批量删除证书，删除证书前支持查询证书是否关联了腾讯云云资源 （需自定义配置参数，参数名称：IsSync）
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CERTIFICATESYNCTASKCERTIFICATEIDCOUNTINVALID = "FailedOperation.CertificateSyncTaskCertificateIdCountInvalid"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
func (c *Client) DeleteCertificates(request *DeleteCertificatesRequest) (response *DeleteCertificatesResponse, err error) {
    return c.DeleteCertificatesWithContext(context.Background(), request)
}

// DeleteCertificates
// 批量删除证书，删除证书前支持查询证书是否关联了腾讯云云资源 （需自定义配置参数，参数名称：IsSync）
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CERTIFICATESYNCTASKCERTIFICATEIDCOUNTINVALID = "FailedOperation.CertificateSyncTaskCertificateIdCountInvalid"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
func (c *Client) DeleteCertificatesWithContext(ctx context.Context, request *DeleteCertificatesRequest) (response *DeleteCertificatesResponse, err error) {
    if request == nil {
        request = NewDeleteCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DeleteCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteManagerRequest() (request *DeleteManagerRequest) {
    request = &DeleteManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DeleteManager")
    
    
    return
}

func NewDeleteManagerResponse() (response *DeleteManagerResponse) {
    response = &DeleteManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteManager
// 删除管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_MANAGERCANNOTDELETECA = "FailedOperation.ManagerCanNotDeleteCa"
//  FAILEDOPERATION_MANAGERCANNOTDELETECERT = "FailedOperation.ManagerCanNotDeleteCert"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DeleteManager(request *DeleteManagerRequest) (response *DeleteManagerResponse, err error) {
    return c.DeleteManagerWithContext(context.Background(), request)
}

// DeleteManager
// 删除管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_MANAGERCANNOTDELETECA = "FailedOperation.ManagerCanNotDeleteCa"
//  FAILEDOPERATION_MANAGERCANNOTDELETECERT = "FailedOperation.ManagerCanNotDeleteCert"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DeleteManagerWithContext(ctx context.Context, request *DeleteManagerRequest) (response *DeleteManagerResponse, err error) {
    if request == nil {
        request = NewDeleteManagerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DeleteManager")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteManagerResponse()
    err = c.Send(request, response)
    return
}

func NewDeployCertificateInstanceRequest() (request *DeployCertificateInstanceRequest) {
    request = &DeployCertificateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DeployCertificateInstance")
    
    
    return
}

func NewDeployCertificateInstanceResponse() (response *DeployCertificateInstanceResponse) {
    response = &DeployCertificateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployCertificateInstance
// 证书部署到云资源实例列表，本接口只会创建部署任务， 部署任务结果可通过DescribeHostDeployRecordDetail查询。本接口创建部署任务时，会校验证书和部署实例的匹配关系，存在不匹配的则会创建部署任务失败。以下为匹配关系校验规则：
//
// - 若待部署的证书和传入实例域名的当前绑定的证书一致， 则不会创建成功
//
// - 若待部署的证书和传入域名不匹配， 则不会创建成功
//
// - 若部署clb实例时， 7层监听器下无规则，则不会创建成功
//
// - 若部署clb实例时， 7层监听器未开启SNI，该监听器下存在任一域名和证书不匹配， 则不会创建成功
//
// - 若部署clb实例时，监听器规则为正则表示式， 则不会创建成功
//
// 
//
// <dx-alert infotype="explain" title="">一个证书ID，相同的资源类型，只能创建一个部署任务，必须等部署任务执行完成，才能创建新的部署任务</dx-alert>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEFORMATERROR = "FailedOperation.CertificateDeployInstanceFormatError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_REGIONNOTEMPTY = "InvalidParameter.RegionNotEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeployCertificateInstance(request *DeployCertificateInstanceRequest) (response *DeployCertificateInstanceResponse, err error) {
    return c.DeployCertificateInstanceWithContext(context.Background(), request)
}

// DeployCertificateInstance
// 证书部署到云资源实例列表，本接口只会创建部署任务， 部署任务结果可通过DescribeHostDeployRecordDetail查询。本接口创建部署任务时，会校验证书和部署实例的匹配关系，存在不匹配的则会创建部署任务失败。以下为匹配关系校验规则：
//
// - 若待部署的证书和传入实例域名的当前绑定的证书一致， 则不会创建成功
//
// - 若待部署的证书和传入域名不匹配， 则不会创建成功
//
// - 若部署clb实例时， 7层监听器下无规则，则不会创建成功
//
// - 若部署clb实例时， 7层监听器未开启SNI，该监听器下存在任一域名和证书不匹配， 则不会创建成功
//
// - 若部署clb实例时，监听器规则为正则表示式， 则不会创建成功
//
// 
//
// <dx-alert infotype="explain" title="">一个证书ID，相同的资源类型，只能创建一个部署任务，必须等部署任务执行完成，才能创建新的部署任务</dx-alert>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEFORMATERROR = "FailedOperation.CertificateDeployInstanceFormatError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_REGIONNOTEMPTY = "InvalidParameter.RegionNotEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeployCertificateInstanceWithContext(ctx context.Context, request *DeployCertificateInstanceRequest) (response *DeployCertificateInstanceResponse, err error) {
    if request == nil {
        request = NewDeployCertificateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DeployCertificateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployCertificateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployCertificateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeployCertificateRecordRetryRequest() (request *DeployCertificateRecordRetryRequest) {
    request = &DeployCertificateRecordRetryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DeployCertificateRecordRetry")
    
    
    return
}

func NewDeployCertificateRecordRetryResponse() (response *DeployCertificateRecordRetryResponse) {
    response = &DeployCertificateRecordRetryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployCertificateRecordRetry
// 云资源部署重试部署记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeployCertificateRecordRetry(request *DeployCertificateRecordRetryRequest) (response *DeployCertificateRecordRetryResponse, err error) {
    return c.DeployCertificateRecordRetryWithContext(context.Background(), request)
}

// DeployCertificateRecordRetry
// 云资源部署重试部署记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeployCertificateRecordRetryWithContext(ctx context.Context, request *DeployCertificateRecordRetryRequest) (response *DeployCertificateRecordRetryResponse, err error) {
    if request == nil {
        request = NewDeployCertificateRecordRetryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DeployCertificateRecordRetry")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployCertificateRecordRetry require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployCertificateRecordRetryResponse()
    err = c.Send(request, response)
    return
}

func NewDeployCertificateRecordRollbackRequest() (request *DeployCertificateRecordRollbackRequest) {
    request = &DeployCertificateRecordRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DeployCertificateRecordRollback")
    
    
    return
}

func NewDeployCertificateRecordRollbackResponse() (response *DeployCertificateRecordRollbackResponse) {
    response = &DeployCertificateRecordRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployCertificateRecordRollback
// 云资源部署成功记录回滚， 部署失败的记录不会回滚； 接口调用成功后， 会创建一个用于回滚的部署任务， 并返回该任务的ID
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYNOTEXIST = "FailedOperation.CertificateDeployNotExist"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeployCertificateRecordRollback(request *DeployCertificateRecordRollbackRequest) (response *DeployCertificateRecordRollbackResponse, err error) {
    return c.DeployCertificateRecordRollbackWithContext(context.Background(), request)
}

// DeployCertificateRecordRollback
// 云资源部署成功记录回滚， 部署失败的记录不会回滚； 接口调用成功后， 会创建一个用于回滚的部署任务， 并返回该任务的ID
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYNOTEXIST = "FailedOperation.CertificateDeployNotExist"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeployCertificateRecordRollbackWithContext(ctx context.Context, request *DeployCertificateRecordRollbackRequest) (response *DeployCertificateRecordRollbackResponse, err error) {
    if request == nil {
        request = NewDeployCertificateRecordRollbackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DeployCertificateRecordRollback")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployCertificateRecordRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployCertificateRecordRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateRequest() (request *DescribeCertificateRequest) {
    request = &DescribeCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificate")
    
    
    return
}

func NewDescribeCertificateResponse() (response *DescribeCertificateResponse) {
    response = &DescribeCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificate
// 本接口（DescribeCertificate）用于获取证书信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificate(request *DescribeCertificateRequest) (response *DescribeCertificateResponse, err error) {
    return c.DescribeCertificateWithContext(context.Background(), request)
}

// DescribeCertificate
// 本接口（DescribeCertificate）用于获取证书信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificateWithContext(ctx context.Context, request *DescribeCertificateRequest) (response *DescribeCertificateResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateBindResourceTaskDetailRequest() (request *DescribeCertificateBindResourceTaskDetailRequest) {
    request = &DescribeCertificateBindResourceTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateBindResourceTaskDetail")
    
    
    return
}

func NewDescribeCertificateBindResourceTaskDetailResponse() (response *DescribeCertificateBindResourceTaskDetailResponse) {
    response = &DescribeCertificateBindResourceTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificateBindResourceTaskDetail
// 查询CreateCertificateBindResourceSyncTask任务结果， 返回证书关联云资源异步任务结果， 支持以下云资源：clb、cdn、waf、live、vod、ddos、tke、apigateway、tcb、teo（edgeOne）、cos
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATESYNCTASKIDINVALID = "FailedOperation.CertificateSyncTaskIdInvalid"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
func (c *Client) DescribeCertificateBindResourceTaskDetail(request *DescribeCertificateBindResourceTaskDetailRequest) (response *DescribeCertificateBindResourceTaskDetailResponse, err error) {
    return c.DescribeCertificateBindResourceTaskDetailWithContext(context.Background(), request)
}

// DescribeCertificateBindResourceTaskDetail
// 查询CreateCertificateBindResourceSyncTask任务结果， 返回证书关联云资源异步任务结果， 支持以下云资源：clb、cdn、waf、live、vod、ddos、tke、apigateway、tcb、teo（edgeOne）、cos
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATESYNCTASKIDINVALID = "FailedOperation.CertificateSyncTaskIdInvalid"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
func (c *Client) DescribeCertificateBindResourceTaskDetailWithContext(ctx context.Context, request *DescribeCertificateBindResourceTaskDetailRequest) (response *DescribeCertificateBindResourceTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateBindResourceTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeCertificateBindResourceTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateBindResourceTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateBindResourceTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateBindResourceTaskResultRequest() (request *DescribeCertificateBindResourceTaskResultRequest) {
    request = &DescribeCertificateBindResourceTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateBindResourceTaskResult")
    
    
    return
}

func NewDescribeCertificateBindResourceTaskResultResponse() (response *DescribeCertificateBindResourceTaskResultResponse) {
    response = &DescribeCertificateBindResourceTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificateBindResourceTaskResult
// 查询CreateCertificateBindResourceSyncTask任务结果， 返回证书关联云资源异步任务结果， 支持以下云资源：clb、cdn、waf、live、vod、ddos、tke、apigateway、tcb、teo（edgeOne）、cos
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATESYNCTASKIDINVALID = "FailedOperation.CertificateSyncTaskIdInvalid"
func (c *Client) DescribeCertificateBindResourceTaskResult(request *DescribeCertificateBindResourceTaskResultRequest) (response *DescribeCertificateBindResourceTaskResultResponse, err error) {
    return c.DescribeCertificateBindResourceTaskResultWithContext(context.Background(), request)
}

// DescribeCertificateBindResourceTaskResult
// 查询CreateCertificateBindResourceSyncTask任务结果， 返回证书关联云资源异步任务结果， 支持以下云资源：clb、cdn、waf、live、vod、ddos、tke、apigateway、tcb、teo（edgeOne）、cos
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATESYNCTASKIDINVALID = "FailedOperation.CertificateSyncTaskIdInvalid"
func (c *Client) DescribeCertificateBindResourceTaskResultWithContext(ctx context.Context, request *DescribeCertificateBindResourceTaskResultRequest) (response *DescribeCertificateBindResourceTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateBindResourceTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeCertificateBindResourceTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateBindResourceTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateBindResourceTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateDetailRequest() (request *DescribeCertificateDetailRequest) {
    request = &DescribeCertificateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateDetail")
    
    
    return
}

func NewDescribeCertificateDetailResponse() (response *DescribeCertificateDetailResponse) {
    response = &DescribeCertificateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificateDetail
// 获取证书详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DescribeCertificateDetail(request *DescribeCertificateDetailRequest) (response *DescribeCertificateDetailResponse, err error) {
    return c.DescribeCertificateDetailWithContext(context.Background(), request)
}

// DescribeCertificateDetail
// 获取证书详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DescribeCertificateDetailWithContext(ctx context.Context, request *DescribeCertificateDetailRequest) (response *DescribeCertificateDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeCertificateDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateOperateLogsRequest() (request *DescribeCertificateOperateLogsRequest) {
    request = &DescribeCertificateOperateLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateOperateLogs")
    
    
    return
}

func NewDescribeCertificateOperateLogsResponse() (response *DescribeCertificateOperateLogsResponse) {
    response = &DescribeCertificateOperateLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificateOperateLogs
// 获取用户账号下有关证书的操作日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCertificateOperateLogs(request *DescribeCertificateOperateLogsRequest) (response *DescribeCertificateOperateLogsResponse, err error) {
    return c.DescribeCertificateOperateLogsWithContext(context.Background(), request)
}

// DescribeCertificateOperateLogs
// 获取用户账号下有关证书的操作日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCertificateOperateLogsWithContext(ctx context.Context, request *DescribeCertificateOperateLogsRequest) (response *DescribeCertificateOperateLogsResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateOperateLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeCertificateOperateLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateOperateLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateOperateLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificatesRequest() (request *DescribeCertificatesRequest) {
    request = &DescribeCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificates")
    
    
    return
}

func NewDescribeCertificatesResponse() (response *DescribeCertificatesResponse) {
    response = &DescribeCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificates
// 本接口（DescribeCertificates）用于获取证书列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificates(request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
    return c.DescribeCertificatesWithContext(context.Background(), request)
}

// DescribeCertificates
// 本接口（DescribeCertificates）用于获取证书列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificatesWithContext(ctx context.Context, request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompaniesRequest() (request *DescribeCompaniesRequest) {
    request = &DescribeCompaniesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCompanies")
    
    
    return
}

func NewDescribeCompaniesResponse() (response *DescribeCompaniesResponse) {
    response = &DescribeCompaniesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCompanies
// 查询公司列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeCompanies(request *DescribeCompaniesRequest) (response *DescribeCompaniesResponse, err error) {
    return c.DescribeCompaniesWithContext(context.Background(), request)
}

// DescribeCompanies
// 查询公司列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeCompaniesWithContext(ctx context.Context, request *DescribeCompaniesRequest) (response *DescribeCompaniesResponse, err error) {
    if request == nil {
        request = NewDescribeCompaniesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeCompanies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompanies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompaniesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeleteCertificatesTaskResultRequest() (request *DescribeDeleteCertificatesTaskResultRequest) {
    request = &DescribeDeleteCertificatesTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeDeleteCertificatesTaskResult")
    
    
    return
}

func NewDescribeDeleteCertificatesTaskResultResponse() (response *DescribeDeleteCertificatesTaskResultResponse) {
    response = &DescribeDeleteCertificatesTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeleteCertificatesTaskResult
// 查询批量删除任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDeleteCertificatesTaskResult(request *DescribeDeleteCertificatesTaskResultRequest) (response *DescribeDeleteCertificatesTaskResultResponse, err error) {
    return c.DescribeDeleteCertificatesTaskResultWithContext(context.Background(), request)
}

// DescribeDeleteCertificatesTaskResult
// 查询批量删除任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDeleteCertificatesTaskResultWithContext(ctx context.Context, request *DescribeDeleteCertificatesTaskResultRequest) (response *DescribeDeleteCertificatesTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeDeleteCertificatesTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeDeleteCertificatesTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeleteCertificatesTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeleteCertificatesTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeployedResourcesRequest() (request *DescribeDeployedResourcesRequest) {
    request = &DescribeDeployedResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeDeployedResources")
    
    
    return
}

func NewDescribeDeployedResourcesResponse() (response *DescribeDeployedResourcesResponse) {
    response = &DescribeDeployedResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeployedResources
// 证书查询关联资源， 最新查询接口请使用CreateCertificateBindResourceSyncTask， 可以查询更多支持的云资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATEIDNUMBERLIMIT = "InvalidParameter.CertificateIdNumberLimit"
//  INVALIDPARAMETER_CONTAINSINVALIDCERTIFICATEID = "InvalidParameter.ContainsInvalidCertificateId"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDeployedResources(request *DescribeDeployedResourcesRequest) (response *DescribeDeployedResourcesResponse, err error) {
    return c.DescribeDeployedResourcesWithContext(context.Background(), request)
}

// DescribeDeployedResources
// 证书查询关联资源， 最新查询接口请使用CreateCertificateBindResourceSyncTask， 可以查询更多支持的云资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATEIDNUMBERLIMIT = "InvalidParameter.CertificateIdNumberLimit"
//  INVALIDPARAMETER_CONTAINSINVALIDCERTIFICATEID = "InvalidParameter.ContainsInvalidCertificateId"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDeployedResourcesWithContext(ctx context.Context, request *DescribeDeployedResourcesRequest) (response *DescribeDeployedResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeDeployedResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeDeployedResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeployedResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeployedResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDownloadCertificateUrlRequest() (request *DescribeDownloadCertificateUrlRequest) {
    request = &DescribeDownloadCertificateUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeDownloadCertificateUrl")
    
    
    return
}

func NewDescribeDownloadCertificateUrlResponse() (response *DescribeDownloadCertificateUrlResponse) {
    response = &DescribeDownloadCertificateUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDownloadCertificateUrl
// 获取下载证书链接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEDOWNLOADSERVICETYPENOTSUPPORT = "FailedOperation.CertificateDownloadServiceTypeNotSupport"
//  FAILEDOPERATION_CERTIFICATENOTALLOWDOWNLOAD = "FailedOperation.CertificateNotAllowDownload"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWDOWNLOAD = "FailedOperation.CertificateStatusNotAllowDownload"
//  FAILEDOPERATION_CERTIFICATETYPEINFONOTFOUND = "FailedOperation.CertificateTypeInfoNotFound"
func (c *Client) DescribeDownloadCertificateUrl(request *DescribeDownloadCertificateUrlRequest) (response *DescribeDownloadCertificateUrlResponse, err error) {
    return c.DescribeDownloadCertificateUrlWithContext(context.Background(), request)
}

// DescribeDownloadCertificateUrl
// 获取下载证书链接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEDOWNLOADSERVICETYPENOTSUPPORT = "FailedOperation.CertificateDownloadServiceTypeNotSupport"
//  FAILEDOPERATION_CERTIFICATENOTALLOWDOWNLOAD = "FailedOperation.CertificateNotAllowDownload"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWDOWNLOAD = "FailedOperation.CertificateStatusNotAllowDownload"
//  FAILEDOPERATION_CERTIFICATETYPEINFONOTFOUND = "FailedOperation.CertificateTypeInfoNotFound"
func (c *Client) DescribeDownloadCertificateUrlWithContext(ctx context.Context, request *DescribeDownloadCertificateUrlRequest) (response *DescribeDownloadCertificateUrlResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadCertificateUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeDownloadCertificateUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDownloadCertificateUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDownloadCertificateUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostApiGatewayInstanceListRequest() (request *DescribeHostApiGatewayInstanceListRequest) {
    request = &DescribeHostApiGatewayInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostApiGatewayInstanceList")
    
    
    return
}

func NewDescribeHostApiGatewayInstanceListResponse() (response *DescribeHostApiGatewayInstanceListResponse) {
    response = &DescribeHostApiGatewayInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostApiGatewayInstanceList
// 查询证书apiGateway云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostApiGatewayInstanceList(request *DescribeHostApiGatewayInstanceListRequest) (response *DescribeHostApiGatewayInstanceListResponse, err error) {
    return c.DescribeHostApiGatewayInstanceListWithContext(context.Background(), request)
}

// DescribeHostApiGatewayInstanceList
// 查询证书apiGateway云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostApiGatewayInstanceListWithContext(ctx context.Context, request *DescribeHostApiGatewayInstanceListRequest) (response *DescribeHostApiGatewayInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostApiGatewayInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostApiGatewayInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostApiGatewayInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostApiGatewayInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostCdnInstanceListRequest() (request *DescribeHostCdnInstanceListRequest) {
    request = &DescribeHostCdnInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostCdnInstanceList")
    
    
    return
}

func NewDescribeHostCdnInstanceListResponse() (response *DescribeHostCdnInstanceListResponse) {
    response = &DescribeHostCdnInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostCdnInstanceList
// 查询证书cdn云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostCdnInstanceList(request *DescribeHostCdnInstanceListRequest) (response *DescribeHostCdnInstanceListResponse, err error) {
    return c.DescribeHostCdnInstanceListWithContext(context.Background(), request)
}

// DescribeHostCdnInstanceList
// 查询证书cdn云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostCdnInstanceListWithContext(ctx context.Context, request *DescribeHostCdnInstanceListRequest) (response *DescribeHostCdnInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostCdnInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostCdnInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostCdnInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostCdnInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostClbInstanceListRequest() (request *DescribeHostClbInstanceListRequest) {
    request = &DescribeHostClbInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostClbInstanceList")
    
    
    return
}

func NewDescribeHostClbInstanceListResponse() (response *DescribeHostClbInstanceListResponse) {
    response = &DescribeHostClbInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostClbInstanceList
// 查询证书clb云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostClbInstanceList(request *DescribeHostClbInstanceListRequest) (response *DescribeHostClbInstanceListResponse, err error) {
    return c.DescribeHostClbInstanceListWithContext(context.Background(), request)
}

// DescribeHostClbInstanceList
// 查询证书clb云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostClbInstanceListWithContext(ctx context.Context, request *DescribeHostClbInstanceListRequest) (response *DescribeHostClbInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostClbInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostClbInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostClbInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostClbInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostCosInstanceListRequest() (request *DescribeHostCosInstanceListRequest) {
    request = &DescribeHostCosInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostCosInstanceList")
    
    
    return
}

func NewDescribeHostCosInstanceListResponse() (response *DescribeHostCosInstanceListResponse) {
    response = &DescribeHostCosInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostCosInstanceList
// 查询证书cos云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostCosInstanceList(request *DescribeHostCosInstanceListRequest) (response *DescribeHostCosInstanceListResponse, err error) {
    return c.DescribeHostCosInstanceListWithContext(context.Background(), request)
}

// DescribeHostCosInstanceList
// 查询证书cos云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostCosInstanceListWithContext(ctx context.Context, request *DescribeHostCosInstanceListRequest) (response *DescribeHostCosInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostCosInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostCosInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostCosInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostCosInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostDdosInstanceListRequest() (request *DescribeHostDdosInstanceListRequest) {
    request = &DescribeHostDdosInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostDdosInstanceList")
    
    
    return
}

func NewDescribeHostDdosInstanceListResponse() (response *DescribeHostDdosInstanceListResponse) {
    response = &DescribeHostDdosInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostDdosInstanceList
// 查询证书ddos云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostDdosInstanceList(request *DescribeHostDdosInstanceListRequest) (response *DescribeHostDdosInstanceListResponse, err error) {
    return c.DescribeHostDdosInstanceListWithContext(context.Background(), request)
}

// DescribeHostDdosInstanceList
// 查询证书ddos云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostDdosInstanceListWithContext(ctx context.Context, request *DescribeHostDdosInstanceListRequest) (response *DescribeHostDdosInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostDdosInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostDdosInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostDdosInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostDdosInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostDeployRecordRequest() (request *DescribeHostDeployRecordRequest) {
    request = &DescribeHostDeployRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostDeployRecord")
    
    
    return
}

func NewDescribeHostDeployRecordResponse() (response *DescribeHostDeployRecordResponse) {
    response = &DescribeHostDeployRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostDeployRecord
// 查询证书云资源部署记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostDeployRecord(request *DescribeHostDeployRecordRequest) (response *DescribeHostDeployRecordResponse, err error) {
    return c.DescribeHostDeployRecordWithContext(context.Background(), request)
}

// DescribeHostDeployRecord
// 查询证书云资源部署记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostDeployRecordWithContext(ctx context.Context, request *DescribeHostDeployRecordRequest) (response *DescribeHostDeployRecordResponse, err error) {
    if request == nil {
        request = NewDescribeHostDeployRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostDeployRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostDeployRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostDeployRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostDeployRecordDetailRequest() (request *DescribeHostDeployRecordDetailRequest) {
    request = &DescribeHostDeployRecordDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostDeployRecordDetail")
    
    
    return
}

func NewDescribeHostDeployRecordDetailResponse() (response *DescribeHostDeployRecordDetailResponse) {
    response = &DescribeHostDeployRecordDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostDeployRecordDetail
// 查询证书云资源部署记录详情列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYNOTEXIST = "FailedOperation.CertificateDeployNotExist"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostDeployRecordDetail(request *DescribeHostDeployRecordDetailRequest) (response *DescribeHostDeployRecordDetailResponse, err error) {
    return c.DescribeHostDeployRecordDetailWithContext(context.Background(), request)
}

// DescribeHostDeployRecordDetail
// 查询证书云资源部署记录详情列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYNOTEXIST = "FailedOperation.CertificateDeployNotExist"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostDeployRecordDetailWithContext(ctx context.Context, request *DescribeHostDeployRecordDetailRequest) (response *DescribeHostDeployRecordDetailResponse, err error) {
    if request == nil {
        request = NewDescribeHostDeployRecordDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostDeployRecordDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostDeployRecordDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostDeployRecordDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostLighthouseInstanceListRequest() (request *DescribeHostLighthouseInstanceListRequest) {
    request = &DescribeHostLighthouseInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostLighthouseInstanceList")
    
    
    return
}

func NewDescribeHostLighthouseInstanceListResponse() (response *DescribeHostLighthouseInstanceListResponse) {
    response = &DescribeHostLighthouseInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostLighthouseInstanceList
// 查询证书Lighthouse云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostLighthouseInstanceList(request *DescribeHostLighthouseInstanceListRequest) (response *DescribeHostLighthouseInstanceListResponse, err error) {
    return c.DescribeHostLighthouseInstanceListWithContext(context.Background(), request)
}

// DescribeHostLighthouseInstanceList
// 查询证书Lighthouse云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostLighthouseInstanceListWithContext(ctx context.Context, request *DescribeHostLighthouseInstanceListRequest) (response *DescribeHostLighthouseInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostLighthouseInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostLighthouseInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostLighthouseInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostLighthouseInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostLiveInstanceListRequest() (request *DescribeHostLiveInstanceListRequest) {
    request = &DescribeHostLiveInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostLiveInstanceList")
    
    
    return
}

func NewDescribeHostLiveInstanceListResponse() (response *DescribeHostLiveInstanceListResponse) {
    response = &DescribeHostLiveInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostLiveInstanceList
// 查询证书live云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostLiveInstanceList(request *DescribeHostLiveInstanceListRequest) (response *DescribeHostLiveInstanceListResponse, err error) {
    return c.DescribeHostLiveInstanceListWithContext(context.Background(), request)
}

// DescribeHostLiveInstanceList
// 查询证书live云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostLiveInstanceListWithContext(ctx context.Context, request *DescribeHostLiveInstanceListRequest) (response *DescribeHostLiveInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostLiveInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostLiveInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostLiveInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostLiveInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostTeoInstanceListRequest() (request *DescribeHostTeoInstanceListRequest) {
    request = &DescribeHostTeoInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostTeoInstanceList")
    
    
    return
}

func NewDescribeHostTeoInstanceListResponse() (response *DescribeHostTeoInstanceListResponse) {
    response = &DescribeHostTeoInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostTeoInstanceList
// 查询证书EdgeOne云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostTeoInstanceList(request *DescribeHostTeoInstanceListRequest) (response *DescribeHostTeoInstanceListResponse, err error) {
    return c.DescribeHostTeoInstanceListWithContext(context.Background(), request)
}

// DescribeHostTeoInstanceList
// 查询证书EdgeOne云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostTeoInstanceListWithContext(ctx context.Context, request *DescribeHostTeoInstanceListRequest) (response *DescribeHostTeoInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostTeoInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostTeoInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostTeoInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostTeoInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostTkeInstanceListRequest() (request *DescribeHostTkeInstanceListRequest) {
    request = &DescribeHostTkeInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostTkeInstanceList")
    
    
    return
}

func NewDescribeHostTkeInstanceListResponse() (response *DescribeHostTkeInstanceListResponse) {
    response = &DescribeHostTkeInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostTkeInstanceList
// 查询证书tke云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostTkeInstanceList(request *DescribeHostTkeInstanceListRequest) (response *DescribeHostTkeInstanceListResponse, err error) {
    return c.DescribeHostTkeInstanceListWithContext(context.Background(), request)
}

// DescribeHostTkeInstanceList
// 查询证书tke云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostTkeInstanceListWithContext(ctx context.Context, request *DescribeHostTkeInstanceListRequest) (response *DescribeHostTkeInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostTkeInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostTkeInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostTkeInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostTkeInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostUpdateRecordRequest() (request *DescribeHostUpdateRecordRequest) {
    request = &DescribeHostUpdateRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostUpdateRecord")
    
    
    return
}

func NewDescribeHostUpdateRecordResponse() (response *DescribeHostUpdateRecordResponse) {
    response = &DescribeHostUpdateRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostUpdateRecord
// 查询证书云资源更新记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUpdateRecord(request *DescribeHostUpdateRecordRequest) (response *DescribeHostUpdateRecordResponse, err error) {
    return c.DescribeHostUpdateRecordWithContext(context.Background(), request)
}

// DescribeHostUpdateRecord
// 查询证书云资源更新记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUpdateRecordWithContext(ctx context.Context, request *DescribeHostUpdateRecordRequest) (response *DescribeHostUpdateRecordResponse, err error) {
    if request == nil {
        request = NewDescribeHostUpdateRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostUpdateRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostUpdateRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostUpdateRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostUpdateRecordDetailRequest() (request *DescribeHostUpdateRecordDetailRequest) {
    request = &DescribeHostUpdateRecordDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostUpdateRecordDetail")
    
    
    return
}

func NewDescribeHostUpdateRecordDetailResponse() (response *DescribeHostUpdateRecordDetailResponse) {
    response = &DescribeHostUpdateRecordDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostUpdateRecordDetail
// 查询证书云资源更新记录详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYNOTEXIST = "FailedOperation.CertificateDeployNotExist"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUpdateRecordDetail(request *DescribeHostUpdateRecordDetailRequest) (response *DescribeHostUpdateRecordDetailResponse, err error) {
    return c.DescribeHostUpdateRecordDetailWithContext(context.Background(), request)
}

// DescribeHostUpdateRecordDetail
// 查询证书云资源更新记录详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYNOTEXIST = "FailedOperation.CertificateDeployNotExist"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUpdateRecordDetailWithContext(ctx context.Context, request *DescribeHostUpdateRecordDetailRequest) (response *DescribeHostUpdateRecordDetailResponse, err error) {
    if request == nil {
        request = NewDescribeHostUpdateRecordDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostUpdateRecordDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostUpdateRecordDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostUpdateRecordDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostUploadUpdateRecordRequest() (request *DescribeHostUploadUpdateRecordRequest) {
    request = &DescribeHostUploadUpdateRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostUploadUpdateRecord")
    
    
    return
}

func NewDescribeHostUploadUpdateRecordResponse() (response *DescribeHostUploadUpdateRecordResponse) {
    response = &DescribeHostUploadUpdateRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostUploadUpdateRecord
// 查询证书云资源更新（证书ID不变）记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUploadUpdateRecord(request *DescribeHostUploadUpdateRecordRequest) (response *DescribeHostUploadUpdateRecordResponse, err error) {
    return c.DescribeHostUploadUpdateRecordWithContext(context.Background(), request)
}

// DescribeHostUploadUpdateRecord
// 查询证书云资源更新（证书ID不变）记录列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUploadUpdateRecordWithContext(ctx context.Context, request *DescribeHostUploadUpdateRecordRequest) (response *DescribeHostUploadUpdateRecordResponse, err error) {
    if request == nil {
        request = NewDescribeHostUploadUpdateRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostUploadUpdateRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostUploadUpdateRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostUploadUpdateRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostUploadUpdateRecordDetailRequest() (request *DescribeHostUploadUpdateRecordDetailRequest) {
    request = &DescribeHostUploadUpdateRecordDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostUploadUpdateRecordDetail")
    
    
    return
}

func NewDescribeHostUploadUpdateRecordDetailResponse() (response *DescribeHostUploadUpdateRecordDetailResponse) {
    response = &DescribeHostUploadUpdateRecordDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostUploadUpdateRecordDetail
// 查询证书更新（证书ID不变）部署记录详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeHostUploadUpdateRecordDetail(request *DescribeHostUploadUpdateRecordDetailRequest) (response *DescribeHostUploadUpdateRecordDetailResponse, err error) {
    return c.DescribeHostUploadUpdateRecordDetailWithContext(context.Background(), request)
}

// DescribeHostUploadUpdateRecordDetail
// 查询证书更新（证书ID不变）部署记录详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeHostUploadUpdateRecordDetailWithContext(ctx context.Context, request *DescribeHostUploadUpdateRecordDetailRequest) (response *DescribeHostUploadUpdateRecordDetailResponse, err error) {
    if request == nil {
        request = NewDescribeHostUploadUpdateRecordDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostUploadUpdateRecordDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostUploadUpdateRecordDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostUploadUpdateRecordDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostVodInstanceListRequest() (request *DescribeHostVodInstanceListRequest) {
    request = &DescribeHostVodInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostVodInstanceList")
    
    
    return
}

func NewDescribeHostVodInstanceListResponse() (response *DescribeHostVodInstanceListResponse) {
    response = &DescribeHostVodInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostVodInstanceList
// 查询证书Vod云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostVodInstanceList(request *DescribeHostVodInstanceListRequest) (response *DescribeHostVodInstanceListResponse, err error) {
    return c.DescribeHostVodInstanceListWithContext(context.Background(), request)
}

// DescribeHostVodInstanceList
// 查询证书Vod云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostVodInstanceListWithContext(ctx context.Context, request *DescribeHostVodInstanceListRequest) (response *DescribeHostVodInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostVodInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostVodInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostVodInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostVodInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostWafInstanceListRequest() (request *DescribeHostWafInstanceListRequest) {
    request = &DescribeHostWafInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostWafInstanceList")
    
    
    return
}

func NewDescribeHostWafInstanceListResponse() (response *DescribeHostWafInstanceListResponse) {
    response = &DescribeHostWafInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostWafInstanceList
// 查询证书waf云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostWafInstanceList(request *DescribeHostWafInstanceListRequest) (response *DescribeHostWafInstanceListResponse, err error) {
    return c.DescribeHostWafInstanceListWithContext(context.Background(), request)
}

// DescribeHostWafInstanceList
// 查询证书waf云资源部署实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostWafInstanceListWithContext(ctx context.Context, request *DescribeHostWafInstanceListRequest) (response *DescribeHostWafInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostWafInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeHostWafInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostWafInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostWafInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeManagerDetailRequest() (request *DescribeManagerDetailRequest) {
    request = &DescribeManagerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeManagerDetail")
    
    
    return
}

func NewDescribeManagerDetailResponse() (response *DescribeManagerDetailResponse) {
    response = &DescribeManagerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeManagerDetail
// 查询管理人详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DescribeManagerDetail(request *DescribeManagerDetailRequest) (response *DescribeManagerDetailResponse, err error) {
    return c.DescribeManagerDetailWithContext(context.Background(), request)
}

// DescribeManagerDetail
// 查询管理人详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DescribeManagerDetailWithContext(ctx context.Context, request *DescribeManagerDetailRequest) (response *DescribeManagerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeManagerDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeManagerDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeManagerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeManagerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeManagersRequest() (request *DescribeManagersRequest) {
    request = &DescribeManagersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeManagers")
    
    
    return
}

func NewDescribeManagersResponse() (response *DescribeManagersResponse) {
    response = &DescribeManagersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeManagers
// 查询管理人列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeManagers(request *DescribeManagersRequest) (response *DescribeManagersResponse, err error) {
    return c.DescribeManagersWithContext(context.Background(), request)
}

// DescribeManagers
// 查询管理人列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeManagersWithContext(ctx context.Context, request *DescribeManagersRequest) (response *DescribeManagersResponse, err error) {
    if request == nil {
        request = NewDescribeManagersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribeManagers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeManagers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeManagersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackagesRequest() (request *DescribePackagesRequest) {
    request = &DescribePackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribePackages")
    
    
    return
}

func NewDescribePackagesResponse() (response *DescribePackagesResponse) {
    response = &DescribePackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePackages
// 获得权益包列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribePackages(request *DescribePackagesRequest) (response *DescribePackagesResponse, err error) {
    return c.DescribePackagesWithContext(context.Background(), request)
}

// DescribePackages
// 获得权益包列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribePackagesWithContext(ctx context.Context, request *DescribePackagesRequest) (response *DescribePackagesResponse, err error) {
    if request == nil {
        request = NewDescribePackagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DescribePackages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadCertificateRequest() (request *DownloadCertificateRequest) {
    request = &DownloadCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DownloadCertificate")
    
    
    return
}

func NewDownloadCertificateResponse() (response *DownloadCertificateResponse) {
    response = &DownloadCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadCertificate
// 本接口（DownloadCertificate）用于下载证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
func (c *Client) DownloadCertificate(request *DownloadCertificateRequest) (response *DownloadCertificateResponse, err error) {
    return c.DownloadCertificateWithContext(context.Background(), request)
}

// DownloadCertificate
// 本接口（DownloadCertificate）用于下载证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
func (c *Client) DownloadCertificateWithContext(ctx context.Context, request *DownloadCertificateRequest) (response *DownloadCertificateResponse, err error) {
    if request == nil {
        request = NewDownloadCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "DownloadCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateAliasRequest() (request *ModifyCertificateAliasRequest) {
    request = &ModifyCertificateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificateAlias")
    
    
    return
}

func NewModifyCertificateAliasResponse() (response *ModifyCertificateAliasResponse) {
    response = &ModifyCertificateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCertificateAlias
// 用户传入证书id和备注来修改证书备注。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCertificateAlias(request *ModifyCertificateAliasRequest) (response *ModifyCertificateAliasResponse, err error) {
    return c.ModifyCertificateAliasWithContext(context.Background(), request)
}

// ModifyCertificateAlias
// 用户传入证书id和备注来修改证书备注。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCertificateAliasWithContext(ctx context.Context, request *ModifyCertificateAliasRequest) (response *ModifyCertificateAliasResponse, err error) {
    if request == nil {
        request = NewModifyCertificateAliasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "ModifyCertificateAlias")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificateAlias require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateProjectRequest() (request *ModifyCertificateProjectRequest) {
    request = &ModifyCertificateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificateProject")
    
    
    return
}

func NewModifyCertificateProjectResponse() (response *ModifyCertificateProjectResponse) {
    response = &ModifyCertificateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCertificateProject
// 批量修改证书所属项目。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyCertificateProject(request *ModifyCertificateProjectRequest) (response *ModifyCertificateProjectResponse, err error) {
    return c.ModifyCertificateProjectWithContext(context.Background(), request)
}

// ModifyCertificateProject
// 批量修改证书所属项目。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyCertificateProjectWithContext(ctx context.Context, request *ModifyCertificateProjectRequest) (response *ModifyCertificateProjectResponse, err error) {
    if request == nil {
        request = NewModifyCertificateProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "ModifyCertificateProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateResubmitRequest() (request *ModifyCertificateResubmitRequest) {
    request = &ModifyCertificateResubmitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificateResubmit")
    
    
    return
}

func NewModifyCertificateResubmitResponse() (response *ModifyCertificateResubmitResponse) {
    response = &ModifyCertificateResubmitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCertificateResubmit
// 针对审核失败或审核取消的付费证书，重新发起审核
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWRESUBMIT = "FailedOperation.CertificateStatusNotAllowResubmit"
//  INVALIDPARAMETER_CERTIFICATESTATUSNOTALLOWRESUBMIT = "InvalidParameter.CertificateStatusNotAllowResubmit"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
func (c *Client) ModifyCertificateResubmit(request *ModifyCertificateResubmitRequest) (response *ModifyCertificateResubmitResponse, err error) {
    return c.ModifyCertificateResubmitWithContext(context.Background(), request)
}

// ModifyCertificateResubmit
// 针对审核失败或审核取消的付费证书，重新发起审核
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWRESUBMIT = "FailedOperation.CertificateStatusNotAllowResubmit"
//  INVALIDPARAMETER_CERTIFICATESTATUSNOTALLOWRESUBMIT = "InvalidParameter.CertificateStatusNotAllowResubmit"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
func (c *Client) ModifyCertificateResubmitWithContext(ctx context.Context, request *ModifyCertificateResubmitRequest) (response *ModifyCertificateResubmitResponse, err error) {
    if request == nil {
        request = NewModifyCertificateResubmitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "ModifyCertificateResubmit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificateResubmit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificateResubmitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificatesExpiringNotificationSwitchRequest() (request *ModifyCertificatesExpiringNotificationSwitchRequest) {
    request = &ModifyCertificatesExpiringNotificationSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificatesExpiringNotificationSwitch")
    
    
    return
}

func NewModifyCertificatesExpiringNotificationSwitchResponse() (response *ModifyCertificatesExpiringNotificationSwitchResponse) {
    response = &ModifyCertificatesExpiringNotificationSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCertificatesExpiringNotificationSwitch
// 修改忽略证书到期通知。打开或关闭证书到期通知。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INVALIDPARAMETER_CERTIFICATESNUMBEREXCEEDED = "InvalidParameter.CertificatesNumberExceeded"
func (c *Client) ModifyCertificatesExpiringNotificationSwitch(request *ModifyCertificatesExpiringNotificationSwitchRequest) (response *ModifyCertificatesExpiringNotificationSwitchResponse, err error) {
    return c.ModifyCertificatesExpiringNotificationSwitchWithContext(context.Background(), request)
}

// ModifyCertificatesExpiringNotificationSwitch
// 修改忽略证书到期通知。打开或关闭证书到期通知。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INVALIDPARAMETER_CERTIFICATESNUMBEREXCEEDED = "InvalidParameter.CertificatesNumberExceeded"
func (c *Client) ModifyCertificatesExpiringNotificationSwitchWithContext(ctx context.Context, request *ModifyCertificatesExpiringNotificationSwitchRequest) (response *ModifyCertificatesExpiringNotificationSwitchResponse, err error) {
    if request == nil {
        request = NewModifyCertificatesExpiringNotificationSwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "ModifyCertificatesExpiringNotificationSwitch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificatesExpiringNotificationSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificatesExpiringNotificationSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceCertificateRequest() (request *ReplaceCertificateRequest) {
    request = &ReplaceCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ReplaceCertificate")
    
    
    return
}

func NewReplaceCertificateResponse() (response *ReplaceCertificateResponse) {
    response = &ReplaceCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceCertificate
// 本接口（ReplaceCertificate）用于重颁发证书。已申请的免费证书仅支持 RSA 算法、密钥对参数为2048的证书重颁发，并且目前仅支持1次重颁发。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RENEWALGORITHMINVALID = "InvalidParameter.RenewAlgorithmInvalid"
func (c *Client) ReplaceCertificate(request *ReplaceCertificateRequest) (response *ReplaceCertificateResponse, err error) {
    return c.ReplaceCertificateWithContext(context.Background(), request)
}

// ReplaceCertificate
// 本接口（ReplaceCertificate）用于重颁发证书。已申请的免费证书仅支持 RSA 算法、密钥对参数为2048的证书重颁发，并且目前仅支持1次重颁发。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RENEWALGORITHMINVALID = "InvalidParameter.RenewAlgorithmInvalid"
func (c *Client) ReplaceCertificateWithContext(ctx context.Context, request *ReplaceCertificateRequest) (response *ReplaceCertificateResponse, err error) {
    if request == nil {
        request = NewReplaceCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "ReplaceCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeCertificateRequest() (request *RevokeCertificateRequest) {
    request = &RevokeCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "RevokeCertificate")
    
    
    return
}

func NewRevokeCertificateResponse() (response *RevokeCertificateResponse) {
    response = &RevokeCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RevokeCertificate
// 本接口（RevokeCertificate）用于吊销证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_REVOKEFAILED = "FailedOperation.RevokeFailed"
//  FAILEDOPERATION_REVOKERESOURCEFAILED = "FailedOperation.RevokeResourceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) RevokeCertificate(request *RevokeCertificateRequest) (response *RevokeCertificateResponse, err error) {
    return c.RevokeCertificateWithContext(context.Background(), request)
}

// RevokeCertificate
// 本接口（RevokeCertificate）用于吊销证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_REVOKEFAILED = "FailedOperation.RevokeFailed"
//  FAILEDOPERATION_REVOKERESOURCEFAILED = "FailedOperation.RevokeResourceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) RevokeCertificateWithContext(ctx context.Context, request *RevokeCertificateRequest) (response *RevokeCertificateResponse, err error) {
    if request == nil {
        request = NewRevokeCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "RevokeCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevokeCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevokeCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitAuditManagerRequest() (request *SubmitAuditManagerRequest) {
    request = &SubmitAuditManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "SubmitAuditManager")
    
    
    return
}

func NewSubmitAuditManagerResponse() (response *SubmitAuditManagerResponse) {
    response = &SubmitAuditManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitAuditManager
// 重新提交审核管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ILLEGALMANAGERSTATUS = "FailedOperation.IllegalManagerStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) SubmitAuditManager(request *SubmitAuditManagerRequest) (response *SubmitAuditManagerResponse, err error) {
    return c.SubmitAuditManagerWithContext(context.Background(), request)
}

// SubmitAuditManager
// 重新提交审核管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ILLEGALMANAGERSTATUS = "FailedOperation.IllegalManagerStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) SubmitAuditManagerWithContext(ctx context.Context, request *SubmitAuditManagerRequest) (response *SubmitAuditManagerResponse, err error) {
    if request == nil {
        request = NewSubmitAuditManagerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "SubmitAuditManager")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitAuditManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitAuditManagerResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitCertificateInformationRequest() (request *SubmitCertificateInformationRequest) {
    request = &SubmitCertificateInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "SubmitCertificateInformation")
    
    
    return
}

func NewSubmitCertificateInformationResponse() (response *SubmitCertificateInformationResponse) {
    response = &SubmitCertificateInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitCertificateInformation
// 付费证书提交资料； 本接口不维护新功能， 可使用新接口进行资料提交， [CertificateInfoSubmit](https://cloud.tencent.com/document/product/400/116033)
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) SubmitCertificateInformation(request *SubmitCertificateInformationRequest) (response *SubmitCertificateInformationResponse, err error) {
    return c.SubmitCertificateInformationWithContext(context.Background(), request)
}

// SubmitCertificateInformation
// 付费证书提交资料； 本接口不维护新功能， 可使用新接口进行资料提交， [CertificateInfoSubmit](https://cloud.tencent.com/document/product/400/116033)
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) SubmitCertificateInformationWithContext(ctx context.Context, request *SubmitCertificateInformationRequest) (response *SubmitCertificateInformationResponse, err error) {
    if request == nil {
        request = NewSubmitCertificateInformationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "SubmitCertificateInformation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitCertificateInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitCertificateInformationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCertificateInstanceRequest() (request *UpdateCertificateInstanceRequest) {
    request = &UpdateCertificateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UpdateCertificateInstance")
    
    
    return
}

func NewUpdateCertificateInstanceResponse() (response *UpdateCertificateInstanceResponse) {
    response = &UpdateCertificateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCertificateInstance
// 一键更新旧证书资源，本接口为异步接口， 调用之后DeployRecordId为0表示任务进行中， 重复请求这个接口， 当返回DeployRecordId大于0则表示任务创建成功。 未创建成功则会抛出异常
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEEMPTY = "FailedOperation.CertificateDeployInstanceEmpty"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  INVALIDPARAMETERVALUE_INTERMEDIATECERTNOTSAME = "InvalidParameterValue.IntermediateCertNotSame"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateInstance(request *UpdateCertificateInstanceRequest) (response *UpdateCertificateInstanceResponse, err error) {
    return c.UpdateCertificateInstanceWithContext(context.Background(), request)
}

// UpdateCertificateInstance
// 一键更新旧证书资源，本接口为异步接口， 调用之后DeployRecordId为0表示任务进行中， 重复请求这个接口， 当返回DeployRecordId大于0则表示任务创建成功。 未创建成功则会抛出异常
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEEMPTY = "FailedOperation.CertificateDeployInstanceEmpty"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  INVALIDPARAMETERVALUE_INTERMEDIATECERTNOTSAME = "InvalidParameterValue.IntermediateCertNotSame"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateInstanceWithContext(ctx context.Context, request *UpdateCertificateInstanceRequest) (response *UpdateCertificateInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateCertificateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "UpdateCertificateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCertificateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCertificateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCertificateRecordRetryRequest() (request *UpdateCertificateRecordRetryRequest) {
    request = &UpdateCertificateRecordRetryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UpdateCertificateRecordRetry")
    
    
    return
}

func NewUpdateCertificateRecordRetryResponse() (response *UpdateCertificateRecordRetryResponse) {
    response = &UpdateCertificateRecordRetryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCertificateRecordRetry
// 云资源更新重试部署记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateRecordRetry(request *UpdateCertificateRecordRetryRequest) (response *UpdateCertificateRecordRetryResponse, err error) {
    return c.UpdateCertificateRecordRetryWithContext(context.Background(), request)
}

// UpdateCertificateRecordRetry
// 云资源更新重试部署记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateRecordRetryWithContext(ctx context.Context, request *UpdateCertificateRecordRetryRequest) (response *UpdateCertificateRecordRetryResponse, err error) {
    if request == nil {
        request = NewUpdateCertificateRecordRetryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "UpdateCertificateRecordRetry")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCertificateRecordRetry require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCertificateRecordRetryResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCertificateRecordRollbackRequest() (request *UpdateCertificateRecordRollbackRequest) {
    request = &UpdateCertificateRecordRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UpdateCertificateRecordRollback")
    
    
    return
}

func NewUpdateCertificateRecordRollbackResponse() (response *UpdateCertificateRecordRollbackResponse) {
    response = &UpdateCertificateRecordRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCertificateRecordRollback
// 云资源更新成功记录回滚， 只对更新已成功的记录回滚
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateRecordRollback(request *UpdateCertificateRecordRollbackRequest) (response *UpdateCertificateRecordRollbackResponse, err error) {
    return c.UpdateCertificateRecordRollbackWithContext(context.Background(), request)
}

// UpdateCertificateRecordRollback
// 云资源更新成功记录回滚， 只对更新已成功的记录回滚
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateRecordRollbackWithContext(ctx context.Context, request *UpdateCertificateRecordRollbackRequest) (response *UpdateCertificateRecordRollbackResponse, err error) {
    if request == nil {
        request = NewUpdateCertificateRecordRollbackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "UpdateCertificateRecordRollback")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCertificateRecordRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCertificateRecordRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewUploadCertificateRequest() (request *UploadCertificateRequest) {
    request = &UploadCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadCertificate")
    
    
    return
}

func NewUploadCertificateResponse() (response *UploadCertificateResponse) {
    response = &UploadCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadCertificate
// 本接口（UploadCertificate）用于上传证书。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATECAERROR = "FailedOperation.CertificateCaError"
//  FAILEDOPERATION_CERTIFICATECHAINERROR = "FailedOperation.CertificateChainError"
//  FAILEDOPERATION_CERTIFICATEENCRYPTINVALID = "FailedOperation.CertificateEncryptInvalid"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSMATCHPRIVATEKEY = "InvalidParameter.MissMatchPrivateKey"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadCertificate(request *UploadCertificateRequest) (response *UploadCertificateResponse, err error) {
    return c.UploadCertificateWithContext(context.Background(), request)
}

// UploadCertificate
// 本接口（UploadCertificate）用于上传证书。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATECAERROR = "FailedOperation.CertificateCaError"
//  FAILEDOPERATION_CERTIFICATECHAINERROR = "FailedOperation.CertificateChainError"
//  FAILEDOPERATION_CERTIFICATEENCRYPTINVALID = "FailedOperation.CertificateEncryptInvalid"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSMATCHPRIVATEKEY = "InvalidParameter.MissMatchPrivateKey"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadCertificateWithContext(ctx context.Context, request *UploadCertificateRequest) (response *UploadCertificateResponse, err error) {
    if request == nil {
        request = NewUploadCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "UploadCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewUploadConfirmLetterRequest() (request *UploadConfirmLetterRequest) {
    request = &UploadConfirmLetterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadConfirmLetter")
    
    
    return
}

func NewUploadConfirmLetterResponse() (response *UploadConfirmLetterResponse) {
    response = &UploadConfirmLetterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadConfirmLetter
// 本接口（UploadConfirmLetter）上传证书确认函，不再维护其功能，请用户前往腾讯云证书控制台上传证书确认函
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"
//  FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMAT = "FailedOperation.InvalidConfirmLetterFormat"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMATWOSIGN = "FailedOperation.InvalidConfirmLetterFormatWosign"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_UPLOADCONFIRMCERTIFICATENOTSUPPORT = "FailedOperation.UploadConfirmCertificateNotSupport"
//  FAILEDOPERATION_UPLOADCONFIRMCERTIFICATENOTSUPPORTDV = "FailedOperation.UploadConfirmCertificateNotSupportDV"
//  INTERNALERROR = "InternalError"
func (c *Client) UploadConfirmLetter(request *UploadConfirmLetterRequest) (response *UploadConfirmLetterResponse, err error) {
    return c.UploadConfirmLetterWithContext(context.Background(), request)
}

// UploadConfirmLetter
// 本接口（UploadConfirmLetter）上传证书确认函，不再维护其功能，请用户前往腾讯云证书控制台上传证书确认函
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"
//  FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMAT = "FailedOperation.InvalidConfirmLetterFormat"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMATWOSIGN = "FailedOperation.InvalidConfirmLetterFormatWosign"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_UPLOADCONFIRMCERTIFICATENOTSUPPORT = "FailedOperation.UploadConfirmCertificateNotSupport"
//  FAILEDOPERATION_UPLOADCONFIRMCERTIFICATENOTSUPPORTDV = "FailedOperation.UploadConfirmCertificateNotSupportDV"
//  INTERNALERROR = "InternalError"
func (c *Client) UploadConfirmLetterWithContext(ctx context.Context, request *UploadConfirmLetterRequest) (response *UploadConfirmLetterResponse, err error) {
    if request == nil {
        request = NewUploadConfirmLetterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "UploadConfirmLetter")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadConfirmLetter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadConfirmLetterResponse()
    err = c.Send(request, response)
    return
}

func NewUploadRevokeLetterRequest() (request *UploadRevokeLetterRequest) {
    request = &UploadRevokeLetterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadRevokeLetter")
    
    
    return
}

func NewUploadRevokeLetterResponse() (response *UploadRevokeLetterResponse) {
    response = &UploadRevokeLetterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadRevokeLetter
// 本接口（UploadRevokeLetter）上传证书吊销确认函，不再维护其功能，请用户前往腾讯云证书控制台上传证书吊销确认函
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"
//  FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"
//  FAILEDOPERATION_FILETOOLARGE = "FailedOperation.FileTooLarge"
//  FAILEDOPERATION_FILETOOSMALL = "FailedOperation.FileTooSmall"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDFILETYPE = "FailedOperation.InvalidFileType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_UPLOADCONFIRMCERTIFICATENOTSUPPORT = "FailedOperation.UploadConfirmCertificateNotSupport"
//  FAILEDOPERATION_UPLOADCONFIRMCERTIFICATENOTSUPPORTDV = "FailedOperation.UploadConfirmCertificateNotSupportDV"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadRevokeLetter(request *UploadRevokeLetterRequest) (response *UploadRevokeLetterResponse, err error) {
    return c.UploadRevokeLetterWithContext(context.Background(), request)
}

// UploadRevokeLetter
// 本接口（UploadRevokeLetter）上传证书吊销确认函，不再维护其功能，请用户前往腾讯云证书控制台上传证书吊销确认函
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"
//  FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"
//  FAILEDOPERATION_FILETOOLARGE = "FailedOperation.FileTooLarge"
//  FAILEDOPERATION_FILETOOSMALL = "FailedOperation.FileTooSmall"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDFILETYPE = "FailedOperation.InvalidFileType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_UPLOADCONFIRMCERTIFICATENOTSUPPORT = "FailedOperation.UploadConfirmCertificateNotSupport"
//  FAILEDOPERATION_UPLOADCONFIRMCERTIFICATENOTSUPPORTDV = "FailedOperation.UploadConfirmCertificateNotSupportDV"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadRevokeLetterWithContext(ctx context.Context, request *UploadRevokeLetterRequest) (response *UploadRevokeLetterResponse, err error) {
    if request == nil {
        request = NewUploadRevokeLetterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "UploadRevokeLetter")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadRevokeLetter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadRevokeLetterResponse()
    err = c.Send(request, response)
    return
}

func NewUploadUpdateCertificateInstanceRequest() (request *UploadUpdateCertificateInstanceRequest) {
    request = &UploadUpdateCertificateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadUpdateCertificateInstance")
    
    
    return
}

func NewUploadUpdateCertificateInstanceResponse() (response *UploadUpdateCertificateInstanceResponse) {
    response = &UploadUpdateCertificateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadUpdateCertificateInstance
// 更新证书内容（证书ID不变）并更新关联的云资源，本接口为异步接口， 调用之后DeployRecordId为0表示任务进行中， 重复请求这个接口， 当返回DeployRecordId大于0则表示任务创建成功。 未创建成功则会抛出异常
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEEMPTY = "FailedOperation.CertificateDeployInstanceEmpty"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_CERTIFICATEWHITEFUNCERROR = "FailedOperation.CertificateWhiteFuncError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  INVALIDPARAMETERVALUE_INTERMEDIATECERTNOTSAME = "InvalidParameterValue.IntermediateCertNotSame"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateInstance(request *UploadUpdateCertificateInstanceRequest) (response *UploadUpdateCertificateInstanceResponse, err error) {
    return c.UploadUpdateCertificateInstanceWithContext(context.Background(), request)
}

// UploadUpdateCertificateInstance
// 更新证书内容（证书ID不变）并更新关联的云资源，本接口为异步接口， 调用之后DeployRecordId为0表示任务进行中， 重复请求这个接口， 当返回DeployRecordId大于0则表示任务创建成功。 未创建成功则会抛出异常
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEEMPTY = "FailedOperation.CertificateDeployInstanceEmpty"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_CERTIFICATEWHITEFUNCERROR = "FailedOperation.CertificateWhiteFuncError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  INVALIDPARAMETERVALUE_INTERMEDIATECERTNOTSAME = "InvalidParameterValue.IntermediateCertNotSame"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateInstanceWithContext(ctx context.Context, request *UploadUpdateCertificateInstanceRequest) (response *UploadUpdateCertificateInstanceResponse, err error) {
    if request == nil {
        request = NewUploadUpdateCertificateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "UploadUpdateCertificateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadUpdateCertificateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadUpdateCertificateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUploadUpdateCertificateRecordRetryRequest() (request *UploadUpdateCertificateRecordRetryRequest) {
    request = &UploadUpdateCertificateRecordRetryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadUpdateCertificateRecordRetry")
    
    
    return
}

func NewUploadUpdateCertificateRecordRetryResponse() (response *UploadUpdateCertificateRecordRetryResponse) {
    response = &UploadUpdateCertificateRecordRetryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadUpdateCertificateRecordRetry
// 云资源更新（证书ID不变）重试部署记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateRecordRetry(request *UploadUpdateCertificateRecordRetryRequest) (response *UploadUpdateCertificateRecordRetryResponse, err error) {
    return c.UploadUpdateCertificateRecordRetryWithContext(context.Background(), request)
}

// UploadUpdateCertificateRecordRetry
// 云资源更新（证书ID不变）重试部署记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateRecordRetryWithContext(ctx context.Context, request *UploadUpdateCertificateRecordRetryRequest) (response *UploadUpdateCertificateRecordRetryResponse, err error) {
    if request == nil {
        request = NewUploadUpdateCertificateRecordRetryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "UploadUpdateCertificateRecordRetry")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadUpdateCertificateRecordRetry require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadUpdateCertificateRecordRetryResponse()
    err = c.Send(request, response)
    return
}

func NewUploadUpdateCertificateRecordRollbackRequest() (request *UploadUpdateCertificateRecordRollbackRequest) {
    request = &UploadUpdateCertificateRecordRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadUpdateCertificateRecordRollback")
    
    
    return
}

func NewUploadUpdateCertificateRecordRollbackResponse() (response *UploadUpdateCertificateRecordRollbackResponse) {
    response = &UploadUpdateCertificateRecordRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadUpdateCertificateRecordRollback
// 云资源更新成功（证书ID不变）记录回滚， 会对全量任务进行回滚
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateRecordRollback(request *UploadUpdateCertificateRecordRollbackRequest) (response *UploadUpdateCertificateRecordRollbackResponse, err error) {
    return c.UploadUpdateCertificateRecordRollbackWithContext(context.Background(), request)
}

// UploadUpdateCertificateRecordRollback
// 云资源更新成功（证书ID不变）记录回滚， 会对全量任务进行回滚
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateRecordRollbackWithContext(ctx context.Context, request *UploadUpdateCertificateRecordRollbackRequest) (response *UploadUpdateCertificateRecordRollbackResponse, err error) {
    if request == nil {
        request = NewUploadUpdateCertificateRecordRollbackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "UploadUpdateCertificateRecordRollback")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadUpdateCertificateRecordRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadUpdateCertificateRecordRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyManagerRequest() (request *VerifyManagerRequest) {
    request = &VerifyManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "VerifyManager")
    
    
    return
}

func NewVerifyManagerResponse() (response *VerifyManagerResponse) {
    response = &VerifyManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyManager
// 重新核验管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) VerifyManager(request *VerifyManagerRequest) (response *VerifyManagerResponse, err error) {
    return c.VerifyManagerWithContext(context.Background(), request)
}

// VerifyManager
// 重新核验管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) VerifyManagerWithContext(ctx context.Context, request *VerifyManagerRequest) (response *VerifyManagerResponse, err error) {
    if request == nil {
        request = NewVerifyManagerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ssl", APIVersion, "VerifyManager")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyManagerResponse()
    err = c.Send(request, response)
    return
}
