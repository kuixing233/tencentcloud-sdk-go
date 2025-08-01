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

package v20180801

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-08-01"

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


func NewAddCasterInputInfoRequest() (request *AddCasterInputInfoRequest) {
    request = &AddCasterInputInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "AddCasterInputInfo")
    
    
    return
}

func NewAddCasterInputInfoResponse() (response *AddCasterInputInfoResponse) {
    response = &AddCasterInputInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCasterInputInfo
// 该接口用来向导播台中添加一个输入源，该输入源可以是拉流地址、或是一个文件链接
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_INPUTALREADYEXIST = "FailedOperation.InputAlreadyExist"
//  FAILEDOPERATION_INPUTUSEDINPGM = "FailedOperation.InputUsedInPgm"
//  FAILEDOPERATION_INPUTUSEDINPVW = "FailedOperation.InputUsedInPvw"
//  FAILEDOPERATION_STARTPULLFAILED = "FailedOperation.StartPullFailed"
//  FAILEDOPERATION_STOPPULLFAILED = "FailedOperation.StopPullFailed"
//  FAILEDOPERATION_TOOMUCHINPUT = "FailedOperation.TooMuchInput"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCasterInputInfo(request *AddCasterInputInfoRequest) (response *AddCasterInputInfoResponse, err error) {
    return c.AddCasterInputInfoWithContext(context.Background(), request)
}

// AddCasterInputInfo
// 该接口用来向导播台中添加一个输入源，该输入源可以是拉流地址、或是一个文件链接
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_INPUTALREADYEXIST = "FailedOperation.InputAlreadyExist"
//  FAILEDOPERATION_INPUTUSEDINPGM = "FailedOperation.InputUsedInPgm"
//  FAILEDOPERATION_INPUTUSEDINPVW = "FailedOperation.InputUsedInPvw"
//  FAILEDOPERATION_STARTPULLFAILED = "FailedOperation.StartPullFailed"
//  FAILEDOPERATION_STOPPULLFAILED = "FailedOperation.StopPullFailed"
//  FAILEDOPERATION_TOOMUCHINPUT = "FailedOperation.TooMuchInput"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCasterInputInfoWithContext(ctx context.Context, request *AddCasterInputInfoRequest) (response *AddCasterInputInfoResponse, err error) {
    if request == nil {
        request = NewAddCasterInputInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "AddCasterInputInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCasterInputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCasterInputInfoResponse()
    err = c.Send(request, response)
    return
}

func NewAddCasterLayoutInfoRequest() (request *AddCasterLayoutInfoRequest) {
    request = &AddCasterLayoutInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "AddCasterLayoutInfo")
    
    
    return
}

func NewAddCasterLayoutInfoResponse() (response *AddCasterLayoutInfoResponse) {
    response = &AddCasterLayoutInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCasterLayoutInfo
// 该接口用来增加导播台的布局参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTALREADYEXIST = "FailedOperation.LayoutAlreadyExist"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCasterLayoutInfo(request *AddCasterLayoutInfoRequest) (response *AddCasterLayoutInfoResponse, err error) {
    return c.AddCasterLayoutInfoWithContext(context.Background(), request)
}

// AddCasterLayoutInfo
// 该接口用来增加导播台的布局参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTALREADYEXIST = "FailedOperation.LayoutAlreadyExist"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCasterLayoutInfoWithContext(ctx context.Context, request *AddCasterLayoutInfoRequest) (response *AddCasterLayoutInfoResponse, err error) {
    if request == nil {
        request = NewAddCasterLayoutInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "AddCasterLayoutInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCasterLayoutInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCasterLayoutInfoResponse()
    err = c.Send(request, response)
    return
}

func NewAddCasterMarkPicInfoRequest() (request *AddCasterMarkPicInfoRequest) {
    request = &AddCasterMarkPicInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "AddCasterMarkPicInfo")
    
    
    return
}

func NewAddCasterMarkPicInfoResponse() (response *AddCasterMarkPicInfoResponse) {
    response = &AddCasterMarkPicInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCasterMarkPicInfo
// 该接口用来新增图片水印。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_MARKPICALREADYEXIST = "FailedOperation.MarkPicAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCasterMarkPicInfo(request *AddCasterMarkPicInfoRequest) (response *AddCasterMarkPicInfoResponse, err error) {
    return c.AddCasterMarkPicInfoWithContext(context.Background(), request)
}

// AddCasterMarkPicInfo
// 该接口用来新增图片水印。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_MARKPICALREADYEXIST = "FailedOperation.MarkPicAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCasterMarkPicInfoWithContext(ctx context.Context, request *AddCasterMarkPicInfoRequest) (response *AddCasterMarkPicInfoResponse, err error) {
    if request == nil {
        request = NewAddCasterMarkPicInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "AddCasterMarkPicInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCasterMarkPicInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCasterMarkPicInfoResponse()
    err = c.Send(request, response)
    return
}

func NewAddCasterMarkWordInfoRequest() (request *AddCasterMarkWordInfoRequest) {
    request = &AddCasterMarkWordInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "AddCasterMarkWordInfo")
    
    
    return
}

func NewAddCasterMarkWordInfoResponse() (response *AddCasterMarkWordInfoResponse) {
    response = &AddCasterMarkWordInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCasterMarkWordInfo
// 为导播台添加文本配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_MARKWORDALREADYEXIST = "FailedOperation.MarkWordAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCasterMarkWordInfo(request *AddCasterMarkWordInfoRequest) (response *AddCasterMarkWordInfoResponse, err error) {
    return c.AddCasterMarkWordInfoWithContext(context.Background(), request)
}

// AddCasterMarkWordInfo
// 为导播台添加文本配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_MARKWORDALREADYEXIST = "FailedOperation.MarkWordAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCasterMarkWordInfoWithContext(ctx context.Context, request *AddCasterMarkWordInfoRequest) (response *AddCasterMarkWordInfoResponse, err error) {
    if request == nil {
        request = NewAddCasterMarkWordInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "AddCasterMarkWordInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCasterMarkWordInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCasterMarkWordInfoResponse()
    err = c.Send(request, response)
    return
}

func NewAddCasterOutputInfoRequest() (request *AddCasterOutputInfoRequest) {
    request = &AddCasterOutputInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "AddCasterOutputInfo")
    
    
    return
}

func NewAddCasterOutputInfoResponse() (response *AddCasterOutputInfoResponse) {
    response = &AddCasterOutputInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCasterOutputInfo
// 该接口用来新增导播台推流信息。导播台主监启动后，将会将主监画面推向该接口设置的地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_OUTPUTALREADYEXIST = "FailedOperation.OutputAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCasterOutputInfo(request *AddCasterOutputInfoRequest) (response *AddCasterOutputInfoResponse, err error) {
    return c.AddCasterOutputInfoWithContext(context.Background(), request)
}

// AddCasterOutputInfo
// 该接口用来新增导播台推流信息。导播台主监启动后，将会将主监画面推向该接口设置的地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_OUTPUTALREADYEXIST = "FailedOperation.OutputAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddCasterOutputInfoWithContext(ctx context.Context, request *AddCasterOutputInfoRequest) (response *AddCasterOutputInfoResponse, err error) {
    if request == nil {
        request = NewAddCasterOutputInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "AddCasterOutputInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCasterOutputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCasterOutputInfoResponse()
    err = c.Send(request, response)
    return
}

func NewAddDelayLiveStreamRequest() (request *AddDelayLiveStreamRequest) {
    request = &AddDelayLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "AddDelayLiveStream")
    
    
    return
}

func NewAddDelayLiveStreamResponse() (response *AddDelayLiveStreamResponse) {
    response = &AddDelayLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddDelayLiveStream
// 针对大型活动直播，通过对直播流设置延时来控制现场与观众播放画面的时间间隔，避免突发状况造成影响。
//
// 
//
// 注意：如果在推流前设置延播，需要提前5分钟设置，目前该接口只支持流粒度。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddDelayLiveStream(request *AddDelayLiveStreamRequest) (response *AddDelayLiveStreamResponse, err error) {
    return c.AddDelayLiveStreamWithContext(context.Background(), request)
}

// AddDelayLiveStream
// 针对大型活动直播，通过对直播流设置延时来控制现场与观众播放画面的时间间隔，避免突发状况造成影响。
//
// 
//
// 注意：如果在推流前设置延播，需要提前5分钟设置，目前该接口只支持流粒度。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddDelayLiveStreamWithContext(ctx context.Context, request *AddDelayLiveStreamRequest) (response *AddDelayLiveStreamResponse, err error) {
    if request == nil {
        request = NewAddDelayLiveStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "AddDelayLiveStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddDelayLiveStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddDelayLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewAddLiveDomainRequest() (request *AddLiveDomainRequest) {
    request = &AddLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "AddLiveDomain")
    
    
    return
}

func NewAddLiveDomainResponse() (response *AddLiveDomainResponse) {
    response = &AddLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddLiveDomain
// 添加域名，一次只能提交一个域名。域名必须已备案。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  FAILEDOPERATION_DOMAINADDED = "FailedOperation.DomainAdded"
//  FAILEDOPERATION_DOMAINGSLBFAIL = "FailedOperation.DomainGslbFail"
//  FAILEDOPERATION_DOMAINNEEDREALNAME = "FailedOperation.DomainNeedRealName"
//  FAILEDOPERATION_DOMAINNEEDVERIFYOWNER = "FailedOperation.DomainNeedVerifyOwner"
//  FAILEDOPERATION_HOSTOUTLIMIT = "FailedOperation.HostOutLimit"
//  FAILEDOPERATION_PARENTDOMAINADDED = "FailedOperation.ParentDomainAdded"
//  FAILEDOPERATION_SUBDOMAINADDED = "FailedOperation.SubDomainAdded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHINESECHARACTERDETECTED = "InternalError.ChineseCharacterDetected"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINALREADYEXIST = "InternalError.DomainAlreadyExist"
//  INTERNALERROR_DOMAINFORMATERROR = "InternalError.DomainFormatError"
//  INTERNALERROR_DOMAINGSLBFAIL = "InternalError.DomainGslbFail"
//  INTERNALERROR_DOMAINISFAMOUS = "InternalError.DomainIsFamous"
//  INTERNALERROR_DOMAINISLIMITED = "InternalError.DomainIsLimited"
//  INTERNALERROR_DOMAINNORECORD = "InternalError.DomainNoRecord"
//  INTERNALERROR_DOMAINTOOLONG = "InternalError.DomainTooLong"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  INVALIDPARAMETER_DOMAINISFAMOUS = "InvalidParameter.DomainIsFamous"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_MPHOSTDELETE = "InvalidParameter.MpHostDelete"
//  INVALIDPARAMETER_MPPLUGINNOUSE = "InvalidParameter.MpPluginNoUse"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) AddLiveDomain(request *AddLiveDomainRequest) (response *AddLiveDomainResponse, err error) {
    return c.AddLiveDomainWithContext(context.Background(), request)
}

// AddLiveDomain
// 添加域名，一次只能提交一个域名。域名必须已备案。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  FAILEDOPERATION_DOMAINADDED = "FailedOperation.DomainAdded"
//  FAILEDOPERATION_DOMAINGSLBFAIL = "FailedOperation.DomainGslbFail"
//  FAILEDOPERATION_DOMAINNEEDREALNAME = "FailedOperation.DomainNeedRealName"
//  FAILEDOPERATION_DOMAINNEEDVERIFYOWNER = "FailedOperation.DomainNeedVerifyOwner"
//  FAILEDOPERATION_HOSTOUTLIMIT = "FailedOperation.HostOutLimit"
//  FAILEDOPERATION_PARENTDOMAINADDED = "FailedOperation.ParentDomainAdded"
//  FAILEDOPERATION_SUBDOMAINADDED = "FailedOperation.SubDomainAdded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHINESECHARACTERDETECTED = "InternalError.ChineseCharacterDetected"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINALREADYEXIST = "InternalError.DomainAlreadyExist"
//  INTERNALERROR_DOMAINFORMATERROR = "InternalError.DomainFormatError"
//  INTERNALERROR_DOMAINGSLBFAIL = "InternalError.DomainGslbFail"
//  INTERNALERROR_DOMAINISFAMOUS = "InternalError.DomainIsFamous"
//  INTERNALERROR_DOMAINISLIMITED = "InternalError.DomainIsLimited"
//  INTERNALERROR_DOMAINNORECORD = "InternalError.DomainNoRecord"
//  INTERNALERROR_DOMAINTOOLONG = "InternalError.DomainTooLong"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  INVALIDPARAMETER_DOMAINISFAMOUS = "InvalidParameter.DomainIsFamous"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_MPHOSTDELETE = "InvalidParameter.MpHostDelete"
//  INVALIDPARAMETER_MPPLUGINNOUSE = "InvalidParameter.MpPluginNoUse"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) AddLiveDomainWithContext(ctx context.Context, request *AddLiveDomainRequest) (response *AddLiveDomainResponse, err error) {
    if request == nil {
        request = NewAddLiveDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "AddLiveDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddLiveDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewAddLiveWatermarkRequest() (request *AddLiveWatermarkRequest) {
    request = &AddLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "AddLiveWatermark")
    
    
    return
}

func NewAddLiveWatermarkResponse() (response *AddLiveWatermarkResponse) {
    response = &AddLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddLiveWatermark
// 添加水印，成功返回水印 ID 后，需要调用[CreateLiveWatermarkRule](/document/product/267/32629)接口将水印 ID 绑定到流使用。 水印数量上限 100，超过后需要先删除，再添加。
//
// 可能返回的错误码:
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_WATERMARKADDERROR = "InternalError.WatermarkAddError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTFOUNT = "ResourceNotFound.UserNotFount"
func (c *Client) AddLiveWatermark(request *AddLiveWatermarkRequest) (response *AddLiveWatermarkResponse, err error) {
    return c.AddLiveWatermarkWithContext(context.Background(), request)
}

// AddLiveWatermark
// 添加水印，成功返回水印 ID 后，需要调用[CreateLiveWatermarkRule](/document/product/267/32629)接口将水印 ID 绑定到流使用。 水印数量上限 100，超过后需要先删除，再添加。
//
// 可能返回的错误码:
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_WATERMARKADDERROR = "InternalError.WatermarkAddError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTFOUNT = "ResourceNotFound.UserNotFount"
func (c *Client) AddLiveWatermarkWithContext(ctx context.Context, request *AddLiveWatermarkRequest) (response *AddLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewAddLiveWatermarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "AddLiveWatermark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddLiveWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewAuthenticateDomainOwnerRequest() (request *AuthenticateDomainOwnerRequest) {
    request = &AuthenticateDomainOwnerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "AuthenticateDomainOwner")
    
    
    return
}

func NewAuthenticateDomainOwnerResponse() (response *AuthenticateDomainOwnerResponse) {
    response = &AuthenticateDomainOwnerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AuthenticateDomainOwner
// 验证用户是否拥有特定直播域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
func (c *Client) AuthenticateDomainOwner(request *AuthenticateDomainOwnerRequest) (response *AuthenticateDomainOwnerResponse, err error) {
    return c.AuthenticateDomainOwnerWithContext(context.Background(), request)
}

// AuthenticateDomainOwner
// 验证用户是否拥有特定直播域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
func (c *Client) AuthenticateDomainOwnerWithContext(ctx context.Context, request *AuthenticateDomainOwnerRequest) (response *AuthenticateDomainOwnerResponse, err error) {
    if request == nil {
        request = NewAuthenticateDomainOwnerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "AuthenticateDomainOwner")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AuthenticateDomainOwner require credential")
    }

    request.SetContext(ctx)
    
    response = NewAuthenticateDomainOwnerResponse()
    err = c.Send(request, response)
    return
}

func NewCancelCommonMixStreamRequest() (request *CancelCommonMixStreamRequest) {
    request = &CancelCommonMixStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CancelCommonMixStream")
    
    
    return
}

func NewCancelCommonMixStreamResponse() (response *CancelCommonMixStreamResponse) {
    response = &CancelCommonMixStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelCommonMixStream
// 该接口用来取消混流。用法与 mix_streamv2.cancel_mix_stream 基本一致。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRERROR = "FailedOperation.CallOtherSvrError"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_CANCELSESSIONNOTEXIST = "FailedOperation.CancelSessionNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELSESSIONNOTEXIST = "InvalidParameter.CancelSessionNotExist"
//  INVALIDPARAMETER_OTHERERROR = "InvalidParameter.OtherError"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CancelCommonMixStream(request *CancelCommonMixStreamRequest) (response *CancelCommonMixStreamResponse, err error) {
    return c.CancelCommonMixStreamWithContext(context.Background(), request)
}

// CancelCommonMixStream
// 该接口用来取消混流。用法与 mix_streamv2.cancel_mix_stream 基本一致。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRERROR = "FailedOperation.CallOtherSvrError"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_CANCELSESSIONNOTEXIST = "FailedOperation.CancelSessionNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELSESSIONNOTEXIST = "InvalidParameter.CancelSessionNotExist"
//  INVALIDPARAMETER_OTHERERROR = "InvalidParameter.OtherError"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CancelCommonMixStreamWithContext(ctx context.Context, request *CancelCommonMixStreamRequest) (response *CancelCommonMixStreamResponse, err error) {
    if request == nil {
        request = NewCancelCommonMixStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CancelCommonMixStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelCommonMixStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelCommonMixStreamResponse()
    err = c.Send(request, response)
    return
}

func NewCopyCasterRequest() (request *CopyCasterRequest) {
    request = &CopyCasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CopyCaster")
    
    
    return
}

func NewCopyCasterResponse() (response *CopyCasterResponse) {
    response = &CopyCasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyCaster
// 该接口用来复制导播台配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CopyCaster(request *CopyCasterRequest) (response *CopyCasterResponse, err error) {
    return c.CopyCasterWithContext(context.Background(), request)
}

// CopyCaster
// 该接口用来复制导播台配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CopyCasterWithContext(ctx context.Context, request *CopyCasterRequest) (response *CopyCasterResponse, err error) {
    if request == nil {
        request = NewCopyCasterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CopyCaster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyCaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyCasterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditKeywordsRequest() (request *CreateAuditKeywordsRequest) {
    request = &CreateAuditKeywordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateAuditKeywords")
    
    
    return
}

func NewCreateAuditKeywordsResponse() (response *CreateAuditKeywordsResponse) {
    response = &CreateAuditKeywordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAuditKeywords
// 创建关键词，并关联到关键词库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateAuditKeywords(request *CreateAuditKeywordsRequest) (response *CreateAuditKeywordsResponse, err error) {
    return c.CreateAuditKeywordsWithContext(context.Background(), request)
}

// CreateAuditKeywords
// 创建关键词，并关联到关键词库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateAuditKeywordsWithContext(ctx context.Context, request *CreateAuditKeywordsRequest) (response *CreateAuditKeywordsResponse, err error) {
    if request == nil {
        request = NewCreateAuditKeywordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateAuditKeywords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditKeywords require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditKeywordsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCasterRequest() (request *CreateCasterRequest) {
    request = &CreateCasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateCaster")
    
    
    return
}

func NewCreateCasterResponse() (response *CreateCasterResponse) {
    response = &CreateCasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCaster
// 该接口用来创建新的导播台
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCaster(request *CreateCasterRequest) (response *CreateCasterResponse, err error) {
    return c.CreateCasterWithContext(context.Background(), request)
}

// CreateCaster
// 该接口用来创建新的导播台
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCasterWithContext(ctx context.Context, request *CreateCasterRequest) (response *CreateCasterResponse, err error) {
    if request == nil {
        request = NewCreateCasterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateCaster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCasterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCasterInputPushUrlRequest() (request *CreateCasterInputPushUrlRequest) {
    request = &CreateCasterInputPushUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateCasterInputPushUrl")
    
    
    return
}

func NewCreateCasterInputPushUrlResponse() (response *CreateCasterInputPushUrlResponse) {
    response = &CreateCasterInputPushUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCasterInputPushUrl
// 该接口用来生成导播台推流地址
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPROTOCOL = "InvalidParameter.InvalidProtocol"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCasterInputPushUrl(request *CreateCasterInputPushUrlRequest) (response *CreateCasterInputPushUrlResponse, err error) {
    return c.CreateCasterInputPushUrlWithContext(context.Background(), request)
}

// CreateCasterInputPushUrl
// 该接口用来生成导播台推流地址
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPROTOCOL = "InvalidParameter.InvalidProtocol"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCasterInputPushUrlWithContext(ctx context.Context, request *CreateCasterInputPushUrlRequest) (response *CreateCasterInputPushUrlResponse, err error) {
    if request == nil {
        request = NewCreateCasterInputPushUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateCasterInputPushUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCasterInputPushUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCasterInputPushUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCasterPgmRequest() (request *CreateCasterPgmRequest) {
    request = &CreateCasterPgmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateCasterPgm")
    
    
    return
}

func NewCreateCasterPgmResponse() (response *CreateCasterPgmResponse) {
    response = &CreateCasterPgmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCasterPgm
// 该接口用来启动主监任务，并将获取主监画面的播放地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_CREATECASTERTASKFAILED = "FailedOperation.CreateCasterTaskFailed"
//  FAILEDOPERATION_NOLVBPUSHORPLAYDOMAIN = "FailedOperation.NoLVBPushOrPlayDomain"
//  FAILEDOPERATION_RESOURCENOTENOUGH = "FailedOperation.ResourceNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCasterPgm(request *CreateCasterPgmRequest) (response *CreateCasterPgmResponse, err error) {
    return c.CreateCasterPgmWithContext(context.Background(), request)
}

// CreateCasterPgm
// 该接口用来启动主监任务，并将获取主监画面的播放地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_CREATECASTERTASKFAILED = "FailedOperation.CreateCasterTaskFailed"
//  FAILEDOPERATION_NOLVBPUSHORPLAYDOMAIN = "FailedOperation.NoLVBPushOrPlayDomain"
//  FAILEDOPERATION_RESOURCENOTENOUGH = "FailedOperation.ResourceNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCasterPgmWithContext(ctx context.Context, request *CreateCasterPgmRequest) (response *CreateCasterPgmResponse, err error) {
    if request == nil {
        request = NewCreateCasterPgmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateCasterPgm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCasterPgm require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCasterPgmResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCasterPgmFromPvwRequest() (request *CreateCasterPgmFromPvwRequest) {
    request = &CreateCasterPgmFromPvwRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateCasterPgmFromPvw")
    
    
    return
}

func NewCreateCasterPgmFromPvwResponse() (response *CreateCasterPgmFromPvwResponse) {
    response = &CreateCasterPgmFromPvwResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCasterPgmFromPvw
// 该接口用来将预监画面的布局、水印、字幕等配置，复制到主监画面中。
//
// 该接口使用时，预监任务需处于运行状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTEREXPIRED = "FailedOperation.CasterExpired"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_CREATECASTERTASKFAILED = "FailedOperation.CreateCasterTaskFailed"
//  FAILEDOPERATION_NOLVBPUSHORPLAYDOMAIN = "FailedOperation.NoLVBPushOrPlayDomain"
//  FAILEDOPERATION_RESOURCENOTENOUGH = "FailedOperation.ResourceNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCasterPgmFromPvw(request *CreateCasterPgmFromPvwRequest) (response *CreateCasterPgmFromPvwResponse, err error) {
    return c.CreateCasterPgmFromPvwWithContext(context.Background(), request)
}

// CreateCasterPgmFromPvw
// 该接口用来将预监画面的布局、水印、字幕等配置，复制到主监画面中。
//
// 该接口使用时，预监任务需处于运行状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTEREXPIRED = "FailedOperation.CasterExpired"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_CREATECASTERTASKFAILED = "FailedOperation.CreateCasterTaskFailed"
//  FAILEDOPERATION_NOLVBPUSHORPLAYDOMAIN = "FailedOperation.NoLVBPushOrPlayDomain"
//  FAILEDOPERATION_RESOURCENOTENOUGH = "FailedOperation.ResourceNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCasterPgmFromPvwWithContext(ctx context.Context, request *CreateCasterPgmFromPvwRequest) (response *CreateCasterPgmFromPvwResponse, err error) {
    if request == nil {
        request = NewCreateCasterPgmFromPvwRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateCasterPgmFromPvw")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCasterPgmFromPvw require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCasterPgmFromPvwResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCasterPvwRequest() (request *CreateCasterPvwRequest) {
    request = &CreateCasterPvwRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateCasterPvw")
    
    
    return
}

func NewCreateCasterPvwResponse() (response *CreateCasterPvwResponse) {
    response = &CreateCasterPvwResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCasterPvw
// 该接口用来启动预监任务，并将获取预监画面的播放地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTEREXPIRED = "FailedOperation.CasterExpired"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_CREATECASTERTASKFAILED = "FailedOperation.CreateCasterTaskFailed"
//  FAILEDOPERATION_RESOURCENOTENOUGH = "FailedOperation.ResourceNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCasterPvw(request *CreateCasterPvwRequest) (response *CreateCasterPvwResponse, err error) {
    return c.CreateCasterPvwWithContext(context.Background(), request)
}

// CreateCasterPvw
// 该接口用来启动预监任务，并将获取预监画面的播放地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTEREXPIRED = "FailedOperation.CasterExpired"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_CREATECASTERTASKFAILED = "FailedOperation.CreateCasterTaskFailed"
//  FAILEDOPERATION_RESOURCENOTENOUGH = "FailedOperation.ResourceNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCasterPvwWithContext(ctx context.Context, request *CreateCasterPvwRequest) (response *CreateCasterPvwResponse, err error) {
    if request == nil {
        request = NewCreateCasterPvwRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateCasterPvw")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCasterPvw require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCasterPvwResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCommonMixStreamRequest() (request *CreateCommonMixStreamRequest) {
    request = &CreateCommonMixStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateCommonMixStream")
    
    
    return
}

func NewCreateCommonMixStreamResponse() (response *CreateCommonMixStreamResponse) {
    response = &CreateCommonMixStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCommonMixStream
// 该接口用来创建通用混流。用法与旧接口 mix_streamv2.start_mix_stream_advanced 基本一致。
//
// 注意：当前最多支持16路混流。
//
// 最佳实践：https://cloud.tencent.com/document/product/267/45566
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRERROR = "FailedOperation.CallOtherSvrError"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_CANCELSESSIONNOTEXIST = "FailedOperation.CancelSessionNotExist"
//  FAILEDOPERATION_GETPICTUREURLERROR = "FailedOperation.GetPictureUrlError"
//  FAILEDOPERATION_GETSTREAMRESOLUTIONERROR = "FailedOperation.GetStreamResolutionError"
//  FAILEDOPERATION_PROCESSMIXERROR = "FailedOperation.ProcessMixError"
//  FAILEDOPERATION_STREAMNOTEXIST = "FailedOperation.StreamNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JIFEIOTHERERROR = "InternalError.JiFeiOtherError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELSESSIONNOTEXIST = "InvalidParameter.CancelSessionNotExist"
//  INVALIDPARAMETER_INPUTNUMLIMITEXCEEDED = "InvalidParameter.InputNumLimitExceeded"
//  INVALIDPARAMETER_INVALIDBACKGROUDRESOLUTION = "InvalidParameter.InvalidBackgroudResolution"
//  INVALIDPARAMETER_INVALIDBITRATE = "InvalidParameter.InvalidBitrate"
//  INVALIDPARAMETER_INVALIDCROPPARAM = "InvalidParameter.InvalidCropParam"
//  INVALIDPARAMETER_INVALIDLAYERPARAM = "InvalidParameter.InvalidLayerParam"
//  INVALIDPARAMETER_INVALIDOUTPUTSTREAMID = "InvalidParameter.InvalidOutputStreamID"
//  INVALIDPARAMETER_INVALIDOUTPUTTYPE = "InvalidParameter.InvalidOutputType"
//  INVALIDPARAMETER_INVALIDPICTUREID = "InvalidParameter.InvalidPictureID"
//  INVALIDPARAMETER_INVALIDROUNDRECTRADIUS = "InvalidParameter.InvalidRoundRectRadius"
//  INVALIDPARAMETER_OTHERERROR = "InvalidParameter.OtherError"
//  INVALIDPARAMETER_SESSIONOUTPUTSTREAMCHANGED = "InvalidParameter.SessionOutputStreamChanged"
//  INVALIDPARAMETER_TEMPLATENOTMATCHINPUTNUM = "InvalidParameter.TemplateNotMatchInputNum"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateCommonMixStream(request *CreateCommonMixStreamRequest) (response *CreateCommonMixStreamResponse, err error) {
    return c.CreateCommonMixStreamWithContext(context.Background(), request)
}

// CreateCommonMixStream
// 该接口用来创建通用混流。用法与旧接口 mix_streamv2.start_mix_stream_advanced 基本一致。
//
// 注意：当前最多支持16路混流。
//
// 最佳实践：https://cloud.tencent.com/document/product/267/45566
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRERROR = "FailedOperation.CallOtherSvrError"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_CANCELSESSIONNOTEXIST = "FailedOperation.CancelSessionNotExist"
//  FAILEDOPERATION_GETPICTUREURLERROR = "FailedOperation.GetPictureUrlError"
//  FAILEDOPERATION_GETSTREAMRESOLUTIONERROR = "FailedOperation.GetStreamResolutionError"
//  FAILEDOPERATION_PROCESSMIXERROR = "FailedOperation.ProcessMixError"
//  FAILEDOPERATION_STREAMNOTEXIST = "FailedOperation.StreamNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JIFEIOTHERERROR = "InternalError.JiFeiOtherError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELSESSIONNOTEXIST = "InvalidParameter.CancelSessionNotExist"
//  INVALIDPARAMETER_INPUTNUMLIMITEXCEEDED = "InvalidParameter.InputNumLimitExceeded"
//  INVALIDPARAMETER_INVALIDBACKGROUDRESOLUTION = "InvalidParameter.InvalidBackgroudResolution"
//  INVALIDPARAMETER_INVALIDBITRATE = "InvalidParameter.InvalidBitrate"
//  INVALIDPARAMETER_INVALIDCROPPARAM = "InvalidParameter.InvalidCropParam"
//  INVALIDPARAMETER_INVALIDLAYERPARAM = "InvalidParameter.InvalidLayerParam"
//  INVALIDPARAMETER_INVALIDOUTPUTSTREAMID = "InvalidParameter.InvalidOutputStreamID"
//  INVALIDPARAMETER_INVALIDOUTPUTTYPE = "InvalidParameter.InvalidOutputType"
//  INVALIDPARAMETER_INVALIDPICTUREID = "InvalidParameter.InvalidPictureID"
//  INVALIDPARAMETER_INVALIDROUNDRECTRADIUS = "InvalidParameter.InvalidRoundRectRadius"
//  INVALIDPARAMETER_OTHERERROR = "InvalidParameter.OtherError"
//  INVALIDPARAMETER_SESSIONOUTPUTSTREAMCHANGED = "InvalidParameter.SessionOutputStreamChanged"
//  INVALIDPARAMETER_TEMPLATENOTMATCHINPUTNUM = "InvalidParameter.TemplateNotMatchInputNum"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateCommonMixStreamWithContext(ctx context.Context, request *CreateCommonMixStreamRequest) (response *CreateCommonMixStreamResponse, err error) {
    if request == nil {
        request = NewCreateCommonMixStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateCommonMixStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCommonMixStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCommonMixStreamResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveCallbackRuleRequest() (request *CreateLiveCallbackRuleRequest) {
    request = &CreateLiveCallbackRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveCallbackRule")
    
    
    return
}

func NewCreateLiveCallbackRuleResponse() (response *CreateLiveCallbackRuleResponse) {
    response = &CreateLiveCallbackRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveCallbackRule
// 创建回调规则，需要先调用[CreateLiveCallbackTemplate](/document/product/267/32637)接口创建回调模板，将返回的模板id绑定到域名/路径进行使用。
//
// <br>回调协议相关文档：[事件消息通知](/document/product/267/32744)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveCallbackRule(request *CreateLiveCallbackRuleRequest) (response *CreateLiveCallbackRuleResponse, err error) {
    return c.CreateLiveCallbackRuleWithContext(context.Background(), request)
}

// CreateLiveCallbackRule
// 创建回调规则，需要先调用[CreateLiveCallbackTemplate](/document/product/267/32637)接口创建回调模板，将返回的模板id绑定到域名/路径进行使用。
//
// <br>回调协议相关文档：[事件消息通知](/document/product/267/32744)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveCallbackRuleWithContext(ctx context.Context, request *CreateLiveCallbackRuleRequest) (response *CreateLiveCallbackRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveCallbackRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveCallbackRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveCallbackRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveCallbackRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveCallbackTemplateRequest() (request *CreateLiveCallbackTemplateRequest) {
    request = &CreateLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveCallbackTemplate")
    
    
    return
}

func NewCreateLiveCallbackTemplateResponse() (response *CreateLiveCallbackTemplateResponse) {
    response = &CreateLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveCallbackTemplate
// 创建回调模板，数量上限：50，成功返回模板id后，需要调用[CreateLiveCallbackRule](/document/product/267/32638)接口将模板 ID 绑定到域名/路径使用。
//
// <br>回调协议相关文档：[事件消息通知](/document/product/267/32744)。
//
// 注意：至少填写一个回调 URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETER_URLNOTSAFE = "InvalidParameter.UrlNotSafe"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveCallbackTemplate(request *CreateLiveCallbackTemplateRequest) (response *CreateLiveCallbackTemplateResponse, err error) {
    return c.CreateLiveCallbackTemplateWithContext(context.Background(), request)
}

// CreateLiveCallbackTemplate
// 创建回调模板，数量上限：50，成功返回模板id后，需要调用[CreateLiveCallbackRule](/document/product/267/32638)接口将模板 ID 绑定到域名/路径使用。
//
// <br>回调协议相关文档：[事件消息通知](/document/product/267/32744)。
//
// 注意：至少填写一个回调 URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETER_URLNOTSAFE = "InvalidParameter.UrlNotSafe"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveCallbackTemplateWithContext(ctx context.Context, request *CreateLiveCallbackTemplateRequest) (response *CreateLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveCallbackTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveCallbackTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveCallbackTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLivePadRuleRequest() (request *CreateLivePadRuleRequest) {
    request = &CreateLivePadRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLivePadRule")
    
    
    return
}

func NewCreateLivePadRuleResponse() (response *CreateLivePadRuleResponse) {
    response = &CreateLivePadRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLivePadRule
// 创建直播垫片规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLivePadRule(request *CreateLivePadRuleRequest) (response *CreateLivePadRuleResponse, err error) {
    return c.CreateLivePadRuleWithContext(context.Background(), request)
}

// CreateLivePadRule
// 创建直播垫片规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLivePadRuleWithContext(ctx context.Context, request *CreateLivePadRuleRequest) (response *CreateLivePadRuleResponse, err error) {
    if request == nil {
        request = NewCreateLivePadRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLivePadRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLivePadRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLivePadRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLivePadTemplateRequest() (request *CreateLivePadTemplateRequest) {
    request = &CreateLivePadTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLivePadTemplate")
    
    
    return
}

func NewCreateLivePadTemplateResponse() (response *CreateLivePadTemplateResponse) {
    response = &CreateLivePadTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLivePadTemplate
// 创建直播垫片模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLivePadTemplate(request *CreateLivePadTemplateRequest) (response *CreateLivePadTemplateResponse, err error) {
    return c.CreateLivePadTemplateWithContext(context.Background(), request)
}

// CreateLivePadTemplate
// 创建直播垫片模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLivePadTemplateWithContext(ctx context.Context, request *CreateLivePadTemplateRequest) (response *CreateLivePadTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLivePadTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLivePadTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLivePadTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLivePadTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLivePullStreamTaskRequest() (request *CreateLivePullStreamTaskRequest) {
    request = &CreateLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLivePullStreamTask")
    
    
    return
}

func NewCreateLivePullStreamTaskResponse() (response *CreateLivePullStreamTaskResponse) {
    response = &CreateLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLivePullStreamTask
// 创建直播拉流任务。支持将外部已有的点播文件，或者直播源拉取过来转推到指定的目标地址。
//
// 注意：
//
// 1. 默认支持任务数上限200个，如有特殊需求，可通过提单到售后进行评估增加上限。
//
// 2. 源流视频编码目前只支持: H264, H265。其他编码格式建议先进行转码处理。
//
// 3. 源流音频编码目前只支持: AAC。其他编码格式建议先进行转码处理。
//
// 4. 可在控制台开启过期自动清理，避免过期任务占用任务数额度。
//
// 5. 拉流转推功能为计费增值服务，计费规则详情可参见[计费文档](https://cloud.tencent.com/document/product/267/53308)。
//
// 6. 拉流转推功能仅提供内容拉取与推送服务，请确保内容已获得授权并符合内容传播相关的法律法规。若内容有侵权或违规相关问题，云直播会停止相关的功能服务并保留追究法律责任的权利。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDBACKUPTOURL = "InvalidParameter.InvalidBackupToUrl"
//  INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"
//  INVALIDPARAMETER_INVALIDMIXINPUTPARAM = "InvalidParameter.InvalidMixInputParam"
//  INVALIDPARAMETER_INVALIDOUTPUTPARAM = "InvalidParameter.InvalidOutputParam"
//  INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"
//  INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"
//  INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"
//  INVALIDPARAMETER_INVALIDWATERMARK = "InvalidParameter.InvalidWatermark"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  INVALIDPARAMETER_TASKNUMMORETHANLIMIT = "InvalidParameter.TaskNumMoreThanLimit"
//  INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateLivePullStreamTask(request *CreateLivePullStreamTaskRequest) (response *CreateLivePullStreamTaskResponse, err error) {
    return c.CreateLivePullStreamTaskWithContext(context.Background(), request)
}

// CreateLivePullStreamTask
// 创建直播拉流任务。支持将外部已有的点播文件，或者直播源拉取过来转推到指定的目标地址。
//
// 注意：
//
// 1. 默认支持任务数上限200个，如有特殊需求，可通过提单到售后进行评估增加上限。
//
// 2. 源流视频编码目前只支持: H264, H265。其他编码格式建议先进行转码处理。
//
// 3. 源流音频编码目前只支持: AAC。其他编码格式建议先进行转码处理。
//
// 4. 可在控制台开启过期自动清理，避免过期任务占用任务数额度。
//
// 5. 拉流转推功能为计费增值服务，计费规则详情可参见[计费文档](https://cloud.tencent.com/document/product/267/53308)。
//
// 6. 拉流转推功能仅提供内容拉取与推送服务，请确保内容已获得授权并符合内容传播相关的法律法规。若内容有侵权或违规相关问题，云直播会停止相关的功能服务并保留追究法律责任的权利。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDBACKUPTOURL = "InvalidParameter.InvalidBackupToUrl"
//  INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"
//  INVALIDPARAMETER_INVALIDMIXINPUTPARAM = "InvalidParameter.InvalidMixInputParam"
//  INVALIDPARAMETER_INVALIDOUTPUTPARAM = "InvalidParameter.InvalidOutputParam"
//  INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"
//  INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"
//  INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"
//  INVALIDPARAMETER_INVALIDWATERMARK = "InvalidParameter.InvalidWatermark"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  INVALIDPARAMETER_TASKNUMMORETHANLIMIT = "InvalidParameter.TaskNumMoreThanLimit"
//  INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateLivePullStreamTaskWithContext(ctx context.Context, request *CreateLivePullStreamTaskRequest) (response *CreateLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewCreateLivePullStreamTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLivePullStreamTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLivePullStreamTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLivePullStreamTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordRequest() (request *CreateLiveRecordRequest) {
    request = &CreateLiveRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveRecord")
    
    
    return
}

func NewCreateLiveRecordResponse() (response *CreateLiveRecordResponse) {
    response = &CreateLiveRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveRecord
// - 使用前提
//
//   1. 录制文件存放于点播平台，所以用户如需使用录制功能，需首先自行开通点播服务。
//
//   2. 录制文件存放后相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，具体请参考 [对应文档](https://cloud.tencent.com/document/product/266/2838)。
//
// 
//
// - 模式说明
//
//   该接口支持两种录制模式：
//
//   1. 定时录制模式【默认模式】。
//
//     需要传入开始时间与结束时间，录制任务根据起止时间自动开始与结束。在所设置结束时间过期之前（且未调用StopLiveRecord提前终止任务），录制任务都是有效的，期间多次断流然后重推都会启动录制任务。
//
//   2. 实时视频录制模式。
//
//     忽略传入的开始时间，在录制任务创建后立即开始录制，录制时长支持最大为30分钟，如果传入的结束时间与当前时间差大于30分钟，则按30分钟计算，实时视频录制主要用于录制精彩视频场景，时长建议控制在5分钟以内。
//
// 
//
// - 注意事项
//
//   1. 调用接口超时设置应大于3秒，小于3秒重试以及按不同起止时间调用都有可能产生重复录制任务，进而导致额外录制费用。
//
//   2. 受限于音视频文件格式（FLV/MP4/HLS）对编码类型的支持，视频编码类型支持 H.264，音频编码类型支持 AAC。
//
//   3. 为避免恶意或非主观的频繁 API 请求，对定时录制模式最大创建任务数做了限制：其中，当天可以创建的最大任务数不超过4000（不含已删除的任务）；当前时刻并发运行的任务数不超过400。有超出此限制的需要提工单申请。
//
//   4. 此调用方式暂时不支持海外推流录制。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  LIMITEXCEEDED_MAXIMUMTASK = "LimitExceeded.MaximumTask"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  RESOURCEUNAVAILABLE_STREAMNOTEXIST = "ResourceUnavailable.StreamNotExist"
func (c *Client) CreateLiveRecord(request *CreateLiveRecordRequest) (response *CreateLiveRecordResponse, err error) {
    return c.CreateLiveRecordWithContext(context.Background(), request)
}

// CreateLiveRecord
// - 使用前提
//
//   1. 录制文件存放于点播平台，所以用户如需使用录制功能，需首先自行开通点播服务。
//
//   2. 录制文件存放后相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，具体请参考 [对应文档](https://cloud.tencent.com/document/product/266/2838)。
//
// 
//
// - 模式说明
//
//   该接口支持两种录制模式：
//
//   1. 定时录制模式【默认模式】。
//
//     需要传入开始时间与结束时间，录制任务根据起止时间自动开始与结束。在所设置结束时间过期之前（且未调用StopLiveRecord提前终止任务），录制任务都是有效的，期间多次断流然后重推都会启动录制任务。
//
//   2. 实时视频录制模式。
//
//     忽略传入的开始时间，在录制任务创建后立即开始录制，录制时长支持最大为30分钟，如果传入的结束时间与当前时间差大于30分钟，则按30分钟计算，实时视频录制主要用于录制精彩视频场景，时长建议控制在5分钟以内。
//
// 
//
// - 注意事项
//
//   1. 调用接口超时设置应大于3秒，小于3秒重试以及按不同起止时间调用都有可能产生重复录制任务，进而导致额外录制费用。
//
//   2. 受限于音视频文件格式（FLV/MP4/HLS）对编码类型的支持，视频编码类型支持 H.264，音频编码类型支持 AAC。
//
//   3. 为避免恶意或非主观的频繁 API 请求，对定时录制模式最大创建任务数做了限制：其中，当天可以创建的最大任务数不超过4000（不含已删除的任务）；当前时刻并发运行的任务数不超过400。有超出此限制的需要提工单申请。
//
//   4. 此调用方式暂时不支持海外推流录制。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  LIMITEXCEEDED_MAXIMUMTASK = "LimitExceeded.MaximumTask"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  RESOURCEUNAVAILABLE_STREAMNOTEXIST = "ResourceUnavailable.StreamNotExist"
func (c *Client) CreateLiveRecordWithContext(ctx context.Context, request *CreateLiveRecordRequest) (response *CreateLiveRecordResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordRuleRequest() (request *CreateLiveRecordRuleRequest) {
    request = &CreateLiveRecordRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveRecordRule")
    
    
    return
}

func NewCreateLiveRecordRuleResponse() (response *CreateLiveRecordRuleResponse) {
    response = &CreateLiveRecordRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveRecordRule
// 创建录制规则，需要先调用[CreateLiveRecordTemplate](/document/product/267/32614)接口创建录制模板，将返回的模板id绑定到流使用。
//
// <br>录制相关文档：[直播录制](/document/product/267/32739)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveRecordRule(request *CreateLiveRecordRuleRequest) (response *CreateLiveRecordRuleResponse, err error) {
    return c.CreateLiveRecordRuleWithContext(context.Background(), request)
}

// CreateLiveRecordRule
// 创建录制规则，需要先调用[CreateLiveRecordTemplate](/document/product/267/32614)接口创建录制模板，将返回的模板id绑定到流使用。
//
// <br>录制相关文档：[直播录制](/document/product/267/32739)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveRecordRuleWithContext(ctx context.Context, request *CreateLiveRecordRuleRequest) (response *CreateLiveRecordRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveRecordRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveRecordRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveRecordRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordTemplateRequest() (request *CreateLiveRecordTemplateRequest) {
    request = &CreateLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveRecordTemplate")
    
    
    return
}

func NewCreateLiveRecordTemplateResponse() (response *CreateLiveRecordTemplateResponse) {
    response = &CreateLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveRecordTemplate
// 创建录制模板，数量上限：50，成功返回模板id后，需要调用[CreateLiveRecordRule](/document/product/267/32615)接口，将模板id绑定到流进行使用。
//
// <br>录制相关文档：[直播录制](/document/product/267/32739)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveRecordTemplate(request *CreateLiveRecordTemplateRequest) (response *CreateLiveRecordTemplateResponse, err error) {
    return c.CreateLiveRecordTemplateWithContext(context.Background(), request)
}

// CreateLiveRecordTemplate
// 创建录制模板，数量上限：50，成功返回模板id后，需要调用[CreateLiveRecordRule](/document/product/267/32615)接口，将模板id绑定到流进行使用。
//
// <br>录制相关文档：[直播录制](/document/product/267/32739)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveRecordTemplateWithContext(ctx context.Context, request *CreateLiveRecordTemplateRequest) (response *CreateLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveSnapshotRuleRequest() (request *CreateLiveSnapshotRuleRequest) {
    request = &CreateLiveSnapshotRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveSnapshotRule")
    
    
    return
}

func NewCreateLiveSnapshotRuleResponse() (response *CreateLiveSnapshotRuleResponse) {
    response = &CreateLiveSnapshotRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveSnapshotRule
// 创建截图规则，需要先调用[CreateLiveSnapshotTemplate](/document/product/267/32624)接口创建截图模板，然后将返回的模板 ID 绑定到流进行使用。
//
// <br>截图相关文档：[直播截图](/document/product/267/32737)。
//
// 注意：单个域名仅支持关联一个截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveSnapshotRule(request *CreateLiveSnapshotRuleRequest) (response *CreateLiveSnapshotRuleResponse, err error) {
    return c.CreateLiveSnapshotRuleWithContext(context.Background(), request)
}

// CreateLiveSnapshotRule
// 创建截图规则，需要先调用[CreateLiveSnapshotTemplate](/document/product/267/32624)接口创建截图模板，然后将返回的模板 ID 绑定到流进行使用。
//
// <br>截图相关文档：[直播截图](/document/product/267/32737)。
//
// 注意：单个域名仅支持关联一个截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveSnapshotRuleWithContext(ctx context.Context, request *CreateLiveSnapshotRuleRequest) (response *CreateLiveSnapshotRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveSnapshotRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveSnapshotRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveSnapshotRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveSnapshotRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveSnapshotTemplateRequest() (request *CreateLiveSnapshotTemplateRequest) {
    request = &CreateLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveSnapshotTemplate")
    
    
    return
}

func NewCreateLiveSnapshotTemplateResponse() (response *CreateLiveSnapshotTemplateResponse) {
    response = &CreateLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveSnapshotTemplate
// 创建截图模板，数量上限：50，成功返回模板id后，需要调用[CreateLiveSnapshotRule](/document/product/267/32625)接口，将模板id绑定到流使用。
//
// <br>截图相关文档：[直播截图](/document/product/267/32737)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_COSBUCKETNOTEXIST = "FailedOperation.CosBucketNotExist"
//  FAILEDOPERATION_COSBUCKETNOTPERMISSION = "FailedOperation.CosBucketNotPermission"
//  FAILEDOPERATION_COSROLENOTEXISTS = "FailedOperation.CosRoleNotExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveSnapshotTemplate(request *CreateLiveSnapshotTemplateRequest) (response *CreateLiveSnapshotTemplateResponse, err error) {
    return c.CreateLiveSnapshotTemplateWithContext(context.Background(), request)
}

// CreateLiveSnapshotTemplate
// 创建截图模板，数量上限：50，成功返回模板id后，需要调用[CreateLiveSnapshotRule](/document/product/267/32625)接口，将模板id绑定到流使用。
//
// <br>截图相关文档：[直播截图](/document/product/267/32737)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_COSBUCKETNOTEXIST = "FailedOperation.CosBucketNotExist"
//  FAILEDOPERATION_COSBUCKETNOTPERMISSION = "FailedOperation.CosBucketNotPermission"
//  FAILEDOPERATION_COSROLENOTEXISTS = "FailedOperation.CosRoleNotExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveSnapshotTemplateWithContext(ctx context.Context, request *CreateLiveSnapshotTemplateRequest) (response *CreateLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveStreamMonitorRequest() (request *CreateLiveStreamMonitorRequest) {
    request = &CreateLiveStreamMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveStreamMonitor")
    
    
    return
}

func NewCreateLiveStreamMonitorResponse() (response *CreateLiveStreamMonitorResponse) {
    response = &CreateLiveStreamMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveStreamMonitor
// 该接口用来创建直播流监播任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_INPUTSTREAMMIXTYPENOTACCESSIBLE = "FailedOperation.InputStreamMixTypeNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_RELATESERVERNOTACCESSIBLE = "FailedOperation.RelateServerNotAccessible"
//  FAILEDOPERATION_RELATEDRUNNINGMONITORLIMITEXCEEDED = "FailedOperation.RelatedRunningMonitorLimitExceeded"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveStreamMonitor(request *CreateLiveStreamMonitorRequest) (response *CreateLiveStreamMonitorResponse, err error) {
    return c.CreateLiveStreamMonitorWithContext(context.Background(), request)
}

// CreateLiveStreamMonitor
// 该接口用来创建直播流监播任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_INPUTSTREAMMIXTYPENOTACCESSIBLE = "FailedOperation.InputStreamMixTypeNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_RELATESERVERNOTACCESSIBLE = "FailedOperation.RelateServerNotAccessible"
//  FAILEDOPERATION_RELATEDRUNNINGMONITORLIMITEXCEEDED = "FailedOperation.RelatedRunningMonitorLimitExceeded"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveStreamMonitorWithContext(ctx context.Context, request *CreateLiveStreamMonitorRequest) (response *CreateLiveStreamMonitorResponse, err error) {
    if request == nil {
        request = NewCreateLiveStreamMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveStreamMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveStreamMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveStreamMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveTimeShiftRuleRequest() (request *CreateLiveTimeShiftRuleRequest) {
    request = &CreateLiveTimeShiftRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveTimeShiftRule")
    
    
    return
}

func NewCreateLiveTimeShiftRuleResponse() (response *CreateLiveTimeShiftRuleResponse) {
    response = &CreateLiveTimeShiftRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveTimeShiftRule
// 创建直播时移规则，需要先调用[CreateLiveTimeShiftTemplate](/document/product/267/86169)接口创建直播时移模板，将返回的模板id绑定到流使用。
//
// <br>直播时移相关文档：[直播时移](/document/product/267/86134)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveTimeShiftRule(request *CreateLiveTimeShiftRuleRequest) (response *CreateLiveTimeShiftRuleResponse, err error) {
    return c.CreateLiveTimeShiftRuleWithContext(context.Background(), request)
}

// CreateLiveTimeShiftRule
// 创建直播时移规则，需要先调用[CreateLiveTimeShiftTemplate](/document/product/267/86169)接口创建直播时移模板，将返回的模板id绑定到流使用。
//
// <br>直播时移相关文档：[直播时移](/document/product/267/86134)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveTimeShiftRuleWithContext(ctx context.Context, request *CreateLiveTimeShiftRuleRequest) (response *CreateLiveTimeShiftRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveTimeShiftRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveTimeShiftRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveTimeShiftRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveTimeShiftRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveTimeShiftTemplateRequest() (request *CreateLiveTimeShiftTemplateRequest) {
    request = &CreateLiveTimeShiftTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveTimeShiftTemplate")
    
    
    return
}

func NewCreateLiveTimeShiftTemplateResponse() (response *CreateLiveTimeShiftTemplateResponse) {
    response = &CreateLiveTimeShiftTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveTimeShiftTemplate
// 创建直播时移模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTimeShiftTemplate(request *CreateLiveTimeShiftTemplateRequest) (response *CreateLiveTimeShiftTemplateResponse, err error) {
    return c.CreateLiveTimeShiftTemplateWithContext(context.Background(), request)
}

// CreateLiveTimeShiftTemplate
// 创建直播时移模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTimeShiftTemplateWithContext(ctx context.Context, request *CreateLiveTimeShiftTemplateRequest) (response *CreateLiveTimeShiftTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveTimeShiftTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveTimeShiftTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveTimeShiftTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveTimeShiftTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveTranscodeRuleRequest() (request *CreateLiveTranscodeRuleRequest) {
    request = &CreateLiveTranscodeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveTranscodeRule")
    
    
    return
}

func NewCreateLiveTranscodeRuleResponse() (response *CreateLiveTranscodeRuleResponse) {
    response = &CreateLiveTranscodeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveTranscodeRule
// 创建转码规则，数量上限：50，需要先调用[CreateLiveTranscodeTemplate](/document/product/267/32646)接口创建转码模板，将返回的模板id绑定到流使用。
//
// <br>转码相关文档：[直播转封装及转码](/document/product/267/32736)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTranscodeRule(request *CreateLiveTranscodeRuleRequest) (response *CreateLiveTranscodeRuleResponse, err error) {
    return c.CreateLiveTranscodeRuleWithContext(context.Background(), request)
}

// CreateLiveTranscodeRule
// 创建转码规则，数量上限：50，需要先调用[CreateLiveTranscodeTemplate](/document/product/267/32646)接口创建转码模板，将返回的模板id绑定到流使用。
//
// <br>转码相关文档：[直播转封装及转码](/document/product/267/32736)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTranscodeRuleWithContext(ctx context.Context, request *CreateLiveTranscodeRuleRequest) (response *CreateLiveTranscodeRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveTranscodeRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveTranscodeRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveTranscodeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveTranscodeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveTranscodeTemplateRequest() (request *CreateLiveTranscodeTemplateRequest) {
    request = &CreateLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveTranscodeTemplate")
    
    
    return
}

func NewCreateLiveTranscodeTemplateResponse() (response *CreateLiveTranscodeTemplateResponse) {
    response = &CreateLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveTranscodeTemplate
// 创建转码模板，数量上限：50，成功返回模板id后，需要调用[CreateLiveTranscodeRule](/document/product/267/32647)接口，将返回的模板id绑定到流使用。
//
// <br>转码相关文档：[直播转封装及转码](/document/product/267/32736)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_PROCESSORALREADYEXIST = "InternalError.ProcessorAlreadyExist"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_GOPMUSTEQUALANDEXISTS = "InvalidParameter.GopMustEqualAndExists"
//  INVALIDPARAMETER_PROCESSORALREADYEXIST = "InvalidParameter.ProcessorAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTranscodeTemplate(request *CreateLiveTranscodeTemplateRequest) (response *CreateLiveTranscodeTemplateResponse, err error) {
    return c.CreateLiveTranscodeTemplateWithContext(context.Background(), request)
}

// CreateLiveTranscodeTemplate
// 创建转码模板，数量上限：50，成功返回模板id后，需要调用[CreateLiveTranscodeRule](/document/product/267/32647)接口，将返回的模板id绑定到流使用。
//
// <br>转码相关文档：[直播转封装及转码](/document/product/267/32736)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_PROCESSORALREADYEXIST = "InternalError.ProcessorAlreadyExist"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_GOPMUSTEQUALANDEXISTS = "InvalidParameter.GopMustEqualAndExists"
//  INVALIDPARAMETER_PROCESSORALREADYEXIST = "InvalidParameter.ProcessorAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTranscodeTemplateWithContext(ctx context.Context, request *CreateLiveTranscodeTemplateRequest) (response *CreateLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveWatermarkRuleRequest() (request *CreateLiveWatermarkRuleRequest) {
    request = &CreateLiveWatermarkRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveWatermarkRule")
    
    
    return
}

func NewCreateLiveWatermarkRuleResponse() (response *CreateLiveWatermarkRuleResponse) {
    response = &CreateLiveWatermarkRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveWatermarkRule
// 创建水印规则，需要先调用[AddLiveWatermark](/document/product/267/30154)接口添加水印，将返回的水印id绑定到流使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveWatermarkRule(request *CreateLiveWatermarkRuleRequest) (response *CreateLiveWatermarkRuleResponse, err error) {
    return c.CreateLiveWatermarkRuleWithContext(context.Background(), request)
}

// CreateLiveWatermarkRule
// 创建水印规则，需要先调用[AddLiveWatermark](/document/product/267/30154)接口添加水印，将返回的水印id绑定到流使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveWatermarkRuleWithContext(ctx context.Context, request *CreateLiveWatermarkRuleRequest) (response *CreateLiveWatermarkRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveWatermarkRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateLiveWatermarkRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveWatermarkRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveWatermarkRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePullStreamConfigRequest() (request *CreatePullStreamConfigRequest) {
    request = &CreatePullStreamConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreatePullStreamConfig")
    
    
    return
}

func NewCreatePullStreamConfigResponse() (response *CreatePullStreamConfigResponse) {
    response = &CreatePullStreamConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePullStreamConfig
// 创建临时拉流转推任务，目前限制添加10条任务。
//
// 该接口已下线,请使用新接口 CreateLivePullStreamTask。
//
// 
//
// 注意：该接口用于创建临时拉流转推任务，
//
// 拉流源地址即 FromUrl 可以是腾讯或非腾讯数据源，
//
// 但转推目标地址即 ToUrl 目前限制为已注册的腾讯直播域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreatePullStreamConfig(request *CreatePullStreamConfigRequest) (response *CreatePullStreamConfigResponse, err error) {
    return c.CreatePullStreamConfigWithContext(context.Background(), request)
}

// CreatePullStreamConfig
// 创建临时拉流转推任务，目前限制添加10条任务。
//
// 该接口已下线,请使用新接口 CreateLivePullStreamTask。
//
// 
//
// 注意：该接口用于创建临时拉流转推任务，
//
// 拉流源地址即 FromUrl 可以是腾讯或非腾讯数据源，
//
// 但转推目标地址即 ToUrl 目前限制为已注册的腾讯直播域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreatePullStreamConfigWithContext(ctx context.Context, request *CreatePullStreamConfigRequest) (response *CreatePullStreamConfigResponse, err error) {
    if request == nil {
        request = NewCreatePullStreamConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreatePullStreamConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePullStreamConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePullStreamConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordTaskRequest() (request *CreateRecordTaskRequest) {
    request = &CreateRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateRecordTask")
    
    
    return
}

func NewCreateRecordTaskResponse() (response *CreateRecordTaskResponse) {
    response = &CreateRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRecordTask
// 创建一个在指定时间启动、结束的录制任务，并使用指定录制模板ID对应的配置进行录制。
//
// - 使用前提
//
// 1. 录制文件存放于点播平台或对象存储内，所以用户如需使用录制功能，需首先自行开通点播服务或对象存储服务。
//
// 2. 录制文件存放后相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，具体请参考[对应文档](https://cloud.tencent.com/document/product/266/2837)。
//
// - 注意事项
//
// 1. 断流会结束当前录制并生成录制文件。在结束时间到达之前任务仍然有效，期间只要正常推流都会正常录制，与是否多次推、断流无关。
//
// 2. 使用上避免创建时间段相互重叠的录制任务。若同一条流当前存在多个时段重叠的任务，为避免重复录制系统将启动最多3个录制任务。
//
// 3. 创建的录制任务记录在平台侧只保留3个月。
//
// 4. 当前录制任务管理API（[CreateRecordTask](https://cloud.tencent.com/document/product/267/45983)/[StopRecordTask](https://cloud.tencent.com/document/product/267/45981)/[DeleteRecordTask](https://cloud.tencent.com/document/product/267/45982)）与旧API（CreateLiveRecord/StopLiveRecord/DeleteLiveRecord）不兼容，两套接口不能混用。
//
// 5. 避免 创建录制任务 与 推流 操作同时进行，可能导致因录制任务未生效而引起任务延迟启动问题，两者操作间隔建议大于3秒。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRecordTask(request *CreateRecordTaskRequest) (response *CreateRecordTaskResponse, err error) {
    return c.CreateRecordTaskWithContext(context.Background(), request)
}

// CreateRecordTask
// 创建一个在指定时间启动、结束的录制任务，并使用指定录制模板ID对应的配置进行录制。
//
// - 使用前提
//
// 1. 录制文件存放于点播平台或对象存储内，所以用户如需使用录制功能，需首先自行开通点播服务或对象存储服务。
//
// 2. 录制文件存放后相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，具体请参考[对应文档](https://cloud.tencent.com/document/product/266/2837)。
//
// - 注意事项
//
// 1. 断流会结束当前录制并生成录制文件。在结束时间到达之前任务仍然有效，期间只要正常推流都会正常录制，与是否多次推、断流无关。
//
// 2. 使用上避免创建时间段相互重叠的录制任务。若同一条流当前存在多个时段重叠的任务，为避免重复录制系统将启动最多3个录制任务。
//
// 3. 创建的录制任务记录在平台侧只保留3个月。
//
// 4. 当前录制任务管理API（[CreateRecordTask](https://cloud.tencent.com/document/product/267/45983)/[StopRecordTask](https://cloud.tencent.com/document/product/267/45981)/[DeleteRecordTask](https://cloud.tencent.com/document/product/267/45982)）与旧API（CreateLiveRecord/StopLiveRecord/DeleteLiveRecord）不兼容，两套接口不能混用。
//
// 5. 避免 创建录制任务 与 推流 操作同时进行，可能导致因录制任务未生效而引起任务延迟启动问题，两者操作间隔建议大于3秒。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRecordTaskWithContext(ctx context.Context, request *CreateRecordTaskRequest) (response *CreateRecordTaskResponse, err error) {
    if request == nil {
        request = NewCreateRecordTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateRecordTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecordTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecordTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScreenshotTaskRequest() (request *CreateScreenshotTaskRequest) {
    request = &CreateScreenshotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateScreenshotTask")
    
    
    return
}

func NewCreateScreenshotTaskResponse() (response *CreateScreenshotTaskResponse) {
    response = &CreateScreenshotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateScreenshotTask
// 创建一个在指定时间启动、结束的截图任务，并使用指定截图模板ID对应的配置进行截图。
//
// - 注意事项
//
// 1. 断流会结束当前截图。在结束时间到达之前任务仍然有效，期间只要正常推流都会正常截图，与是否多次推、断流无关。
//
// 2. 使用上避免创建时间段相互重叠的截图任务。若同一条流当前存在多个时段重叠的任务，为避免重复系统将启动最多3个截图任务。
//
// 3. 创建的截图任务记录在平台侧只保留3个月。
//
// 4. 当前截图任务管理API（CreateScreenshotTask/StopScreenshotTask/DeleteScreenshotTask）与旧API（CreateLiveInstantSnapshot/StopLiveInstantSnapshot）不兼容，两套接口不能混用。
//
// 5. 避免 创建截图任务 与 推流 操作同时进行，可能导致因截图任务未生效而引起任务延迟启动问题，两者操作间隔建议大于3秒。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateScreenshotTask(request *CreateScreenshotTaskRequest) (response *CreateScreenshotTaskResponse, err error) {
    return c.CreateScreenshotTaskWithContext(context.Background(), request)
}

// CreateScreenshotTask
// 创建一个在指定时间启动、结束的截图任务，并使用指定截图模板ID对应的配置进行截图。
//
// - 注意事项
//
// 1. 断流会结束当前截图。在结束时间到达之前任务仍然有效，期间只要正常推流都会正常截图，与是否多次推、断流无关。
//
// 2. 使用上避免创建时间段相互重叠的截图任务。若同一条流当前存在多个时段重叠的任务，为避免重复系统将启动最多3个截图任务。
//
// 3. 创建的截图任务记录在平台侧只保留3个月。
//
// 4. 当前截图任务管理API（CreateScreenshotTask/StopScreenshotTask/DeleteScreenshotTask）与旧API（CreateLiveInstantSnapshot/StopLiveInstantSnapshot）不兼容，两套接口不能混用。
//
// 5. 避免 创建截图任务 与 推流 操作同时进行，可能导致因截图任务未生效而引起任务延迟启动问题，两者操作间隔建议大于3秒。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateScreenshotTaskWithContext(ctx context.Context, request *CreateScreenshotTaskRequest) (response *CreateScreenshotTaskResponse, err error) {
    if request == nil {
        request = NewCreateScreenshotTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "CreateScreenshotTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScreenshotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScreenshotTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditKeywordsRequest() (request *DeleteAuditKeywordsRequest) {
    request = &DeleteAuditKeywordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteAuditKeywords")
    
    
    return
}

func NewDeleteAuditKeywordsResponse() (response *DeleteAuditKeywordsResponse) {
    response = &DeleteAuditKeywordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuditKeywords
// 删除关键词信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAuditKeywords(request *DeleteAuditKeywordsRequest) (response *DeleteAuditKeywordsResponse, err error) {
    return c.DeleteAuditKeywordsWithContext(context.Background(), request)
}

// DeleteAuditKeywords
// 删除关键词信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAuditKeywordsWithContext(ctx context.Context, request *DeleteAuditKeywordsRequest) (response *DeleteAuditKeywordsResponse, err error) {
    if request == nil {
        request = NewDeleteAuditKeywordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteAuditKeywords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditKeywords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditKeywordsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCasterRequest() (request *DeleteCasterRequest) {
    request = &DeleteCasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteCaster")
    
    
    return
}

func NewDeleteCasterResponse() (response *DeleteCasterResponse) {
    response = &DeleteCasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCaster
// 该接口用来删除一个导播台的所有信息。
//
// 注意，调用该接口后，所有的导播台信息将被清除，包括正在直播的内容也将直接中断。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERBINDED = "FailedOperation.CasterBinded"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_STOPCASTERTASKFAILED = "FailedOperation.StopCasterTaskFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCaster(request *DeleteCasterRequest) (response *DeleteCasterResponse, err error) {
    return c.DeleteCasterWithContext(context.Background(), request)
}

// DeleteCaster
// 该接口用来删除一个导播台的所有信息。
//
// 注意，调用该接口后，所有的导播台信息将被清除，包括正在直播的内容也将直接中断。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERBINDED = "FailedOperation.CasterBinded"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_STOPCASTERTASKFAILED = "FailedOperation.StopCasterTaskFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterWithContext(ctx context.Context, request *DeleteCasterRequest) (response *DeleteCasterResponse, err error) {
    if request == nil {
        request = NewDeleteCasterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteCaster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCasterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCasterInputInfoRequest() (request *DeleteCasterInputInfoRequest) {
    request = &DeleteCasterInputInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteCasterInputInfo")
    
    
    return
}

func NewDeleteCasterInputInfoResponse() (response *DeleteCasterInputInfoResponse) {
    response = &DeleteCasterInputInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCasterInputInfo
// 该接口用来删除导播台中的输入源信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTEREXPIRED = "FailedOperation.CasterExpired"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_INPUTUSEDINAUTOCAST = "FailedOperation.InputUsedInAutoCast"
//  FAILEDOPERATION_INPUTUSEDINLAYOUT = "FailedOperation.InputUsedInLayout"
//  FAILEDOPERATION_INPUTUSEDINPGM = "FailedOperation.InputUsedInPgm"
//  FAILEDOPERATION_INPUTUSEDINPVW = "FailedOperation.InputUsedInPvw"
//  FAILEDOPERATION_STARTPULLFAILED = "FailedOperation.StartPullFailed"
//  FAILEDOPERATION_STOPPULLFAILED = "FailedOperation.StopPullFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterInputInfo(request *DeleteCasterInputInfoRequest) (response *DeleteCasterInputInfoResponse, err error) {
    return c.DeleteCasterInputInfoWithContext(context.Background(), request)
}

// DeleteCasterInputInfo
// 该接口用来删除导播台中的输入源信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTEREXPIRED = "FailedOperation.CasterExpired"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_INPUTUSEDINAUTOCAST = "FailedOperation.InputUsedInAutoCast"
//  FAILEDOPERATION_INPUTUSEDINLAYOUT = "FailedOperation.InputUsedInLayout"
//  FAILEDOPERATION_INPUTUSEDINPGM = "FailedOperation.InputUsedInPgm"
//  FAILEDOPERATION_INPUTUSEDINPVW = "FailedOperation.InputUsedInPvw"
//  FAILEDOPERATION_STARTPULLFAILED = "FailedOperation.StartPullFailed"
//  FAILEDOPERATION_STOPPULLFAILED = "FailedOperation.StopPullFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterInputInfoWithContext(ctx context.Context, request *DeleteCasterInputInfoRequest) (response *DeleteCasterInputInfoResponse, err error) {
    if request == nil {
        request = NewDeleteCasterInputInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteCasterInputInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCasterInputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCasterInputInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCasterLayoutInfoRequest() (request *DeleteCasterLayoutInfoRequest) {
    request = &DeleteCasterLayoutInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteCasterLayoutInfo")
    
    
    return
}

func NewDeleteCasterLayoutInfoResponse() (response *DeleteCasterLayoutInfoResponse) {
    response = &DeleteCasterLayoutInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCasterLayoutInfo
// 该接口用来将布局信息从导播台中删除
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINAUTOCAST = "FailedOperation.LayoutUsedInAutoCast"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterLayoutInfo(request *DeleteCasterLayoutInfoRequest) (response *DeleteCasterLayoutInfoResponse, err error) {
    return c.DeleteCasterLayoutInfoWithContext(context.Background(), request)
}

// DeleteCasterLayoutInfo
// 该接口用来将布局信息从导播台中删除
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINAUTOCAST = "FailedOperation.LayoutUsedInAutoCast"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterLayoutInfoWithContext(ctx context.Context, request *DeleteCasterLayoutInfoRequest) (response *DeleteCasterLayoutInfoResponse, err error) {
    if request == nil {
        request = NewDeleteCasterLayoutInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteCasterLayoutInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCasterLayoutInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCasterLayoutInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCasterMarkPicInfoRequest() (request *DeleteCasterMarkPicInfoRequest) {
    request = &DeleteCasterMarkPicInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteCasterMarkPicInfo")
    
    
    return
}

func NewDeleteCasterMarkPicInfoResponse() (response *DeleteCasterMarkPicInfoResponse) {
    response = &DeleteCasterMarkPicInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCasterMarkPicInfo
// 该接口用来删除导播台某个Index对应的水印。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  FAILEDOPERATION_MARKPICUSEDINAUTOCAST = "FailedOperation.MarkPicUsedInAutoCast"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterMarkPicInfo(request *DeleteCasterMarkPicInfoRequest) (response *DeleteCasterMarkPicInfoResponse, err error) {
    return c.DeleteCasterMarkPicInfoWithContext(context.Background(), request)
}

// DeleteCasterMarkPicInfo
// 该接口用来删除导播台某个Index对应的水印。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  FAILEDOPERATION_MARKPICUSEDINAUTOCAST = "FailedOperation.MarkPicUsedInAutoCast"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterMarkPicInfoWithContext(ctx context.Context, request *DeleteCasterMarkPicInfoRequest) (response *DeleteCasterMarkPicInfoResponse, err error) {
    if request == nil {
        request = NewDeleteCasterMarkPicInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteCasterMarkPicInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCasterMarkPicInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCasterMarkPicInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCasterMarkWordInfoRequest() (request *DeleteCasterMarkWordInfoRequest) {
    request = &DeleteCasterMarkWordInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteCasterMarkWordInfo")
    
    
    return
}

func NewDeleteCasterMarkWordInfoResponse() (response *DeleteCasterMarkWordInfoResponse) {
    response = &DeleteCasterMarkWordInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCasterMarkWordInfo
// 该接口用来删除导播台的文本配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  FAILEDOPERATION_MARKWORDUSEDINAUTOCAST = "FailedOperation.MarkWordUsedInAutoCast"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterMarkWordInfo(request *DeleteCasterMarkWordInfoRequest) (response *DeleteCasterMarkWordInfoResponse, err error) {
    return c.DeleteCasterMarkWordInfoWithContext(context.Background(), request)
}

// DeleteCasterMarkWordInfo
// 该接口用来删除导播台的文本配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  FAILEDOPERATION_MARKWORDUSEDINAUTOCAST = "FailedOperation.MarkWordUsedInAutoCast"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterMarkWordInfoWithContext(ctx context.Context, request *DeleteCasterMarkWordInfoRequest) (response *DeleteCasterMarkWordInfoResponse, err error) {
    if request == nil {
        request = NewDeleteCasterMarkWordInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteCasterMarkWordInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCasterMarkWordInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCasterMarkWordInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCasterOutputInfoRequest() (request *DeleteCasterOutputInfoRequest) {
    request = &DeleteCasterOutputInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteCasterOutputInfo")
    
    
    return
}

func NewDeleteCasterOutputInfoResponse() (response *DeleteCasterOutputInfoResponse) {
    response = &DeleteCasterOutputInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCasterOutputInfo
// 该接口用来删除导播台的推流信息。
//
// 注：若删除推流到腾讯云直播源站配置，即OutputIndex为0，OutputType为1的推流配置，在重新启动主监后，系统会自动重新生成一个推流到腾讯云直播源站配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterOutputInfo(request *DeleteCasterOutputInfoRequest) (response *DeleteCasterOutputInfoResponse, err error) {
    return c.DeleteCasterOutputInfoWithContext(context.Background(), request)
}

// DeleteCasterOutputInfo
// 该接口用来删除导播台的推流信息。
//
// 注：若删除推流到腾讯云直播源站配置，即OutputIndex为0，OutputType为1的推流配置，在重新启动主监后，系统会自动重新生成一个推流到腾讯云直播源站配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCasterOutputInfoWithContext(ctx context.Context, request *DeleteCasterOutputInfoRequest) (response *DeleteCasterOutputInfoResponse, err error) {
    if request == nil {
        request = NewDeleteCasterOutputInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteCasterOutputInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCasterOutputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCasterOutputInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveCallbackRuleRequest() (request *DeleteLiveCallbackRuleRequest) {
    request = &DeleteLiveCallbackRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveCallbackRule")
    
    
    return
}

func NewDeleteLiveCallbackRuleResponse() (response *DeleteLiveCallbackRuleResponse) {
    response = &DeleteLiveCallbackRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveCallbackRule
// 删除回调规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DeleteLiveCallbackRule(request *DeleteLiveCallbackRuleRequest) (response *DeleteLiveCallbackRuleResponse, err error) {
    return c.DeleteLiveCallbackRuleWithContext(context.Background(), request)
}

// DeleteLiveCallbackRule
// 删除回调规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DeleteLiveCallbackRuleWithContext(ctx context.Context, request *DeleteLiveCallbackRuleRequest) (response *DeleteLiveCallbackRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveCallbackRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveCallbackRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveCallbackRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveCallbackRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveCallbackTemplateRequest() (request *DeleteLiveCallbackTemplateRequest) {
    request = &DeleteLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveCallbackTemplate")
    
    
    return
}

func NewDeleteLiveCallbackTemplateResponse() (response *DeleteLiveCallbackTemplateResponse) {
    response = &DeleteLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveCallbackTemplate
// 删除回调模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveCallbackTemplate(request *DeleteLiveCallbackTemplateRequest) (response *DeleteLiveCallbackTemplateResponse, err error) {
    return c.DeleteLiveCallbackTemplateWithContext(context.Background(), request)
}

// DeleteLiveCallbackTemplate
// 删除回调模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveCallbackTemplateWithContext(ctx context.Context, request *DeleteLiveCallbackTemplateRequest) (response *DeleteLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveCallbackTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveCallbackTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveCallbackTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveDomainRequest() (request *DeleteLiveDomainRequest) {
    request = &DeleteLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveDomain")
    
    
    return
}

func NewDeleteLiveDomainResponse() (response *DeleteLiveDomainResponse) {
    response = &DeleteLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveDomain
// 删除已添加的直播域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  FAILEDOPERATION_JIFEINOENOUGHFUND = "FailedOperation.JiFeiNoEnoughFund"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_TAGUNBINDERROR = "FailedOperation.TagUnbindError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveDomain(request *DeleteLiveDomainRequest) (response *DeleteLiveDomainResponse, err error) {
    return c.DeleteLiveDomainWithContext(context.Background(), request)
}

// DeleteLiveDomain
// 删除已添加的直播域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  FAILEDOPERATION_JIFEINOENOUGHFUND = "FailedOperation.JiFeiNoEnoughFund"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_TAGUNBINDERROR = "FailedOperation.TagUnbindError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveDomainWithContext(ctx context.Context, request *DeleteLiveDomainRequest) (response *DeleteLiveDomainResponse, err error) {
    if request == nil {
        request = NewDeleteLiveDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLivePadRuleRequest() (request *DeleteLivePadRuleRequest) {
    request = &DeleteLivePadRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLivePadRule")
    
    
    return
}

func NewDeleteLivePadRuleResponse() (response *DeleteLivePadRuleResponse) {
    response = &DeleteLivePadRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLivePadRule
// 删除直播垫片规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLivePadRule(request *DeleteLivePadRuleRequest) (response *DeleteLivePadRuleResponse, err error) {
    return c.DeleteLivePadRuleWithContext(context.Background(), request)
}

// DeleteLivePadRule
// 删除直播垫片规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLivePadRuleWithContext(ctx context.Context, request *DeleteLivePadRuleRequest) (response *DeleteLivePadRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLivePadRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLivePadRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLivePadRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLivePadRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLivePadTemplateRequest() (request *DeleteLivePadTemplateRequest) {
    request = &DeleteLivePadTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLivePadTemplate")
    
    
    return
}

func NewDeleteLivePadTemplateResponse() (response *DeleteLivePadTemplateResponse) {
    response = &DeleteLivePadTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLivePadTemplate
// 删除直播垫片模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLivePadTemplate(request *DeleteLivePadTemplateRequest) (response *DeleteLivePadTemplateResponse, err error) {
    return c.DeleteLivePadTemplateWithContext(context.Background(), request)
}

// DeleteLivePadTemplate
// 删除直播垫片模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLivePadTemplateWithContext(ctx context.Context, request *DeleteLivePadTemplateRequest) (response *DeleteLivePadTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLivePadTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLivePadTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLivePadTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLivePadTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLivePullStreamTaskRequest() (request *DeleteLivePullStreamTaskRequest) {
    request = &DeleteLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLivePullStreamTask")
    
    
    return
}

func NewDeleteLivePullStreamTaskResponse() (response *DeleteLivePullStreamTaskResponse) {
    response = &DeleteLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLivePullStreamTask
// 删除接口 CreateLivePullStreamTask 创建的拉流任务。
//
// 注意：
//
// 1. 入参中的 TaskId 为 CreateLivePullStreamTask 接口创建时返回的TaskId。
//
// 2. 也可通过 DescribeLivePullStreamTasks 进行查询创建的任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteLivePullStreamTask(request *DeleteLivePullStreamTaskRequest) (response *DeleteLivePullStreamTaskResponse, err error) {
    return c.DeleteLivePullStreamTaskWithContext(context.Background(), request)
}

// DeleteLivePullStreamTask
// 删除接口 CreateLivePullStreamTask 创建的拉流任务。
//
// 注意：
//
// 1. 入参中的 TaskId 为 CreateLivePullStreamTask 接口创建时返回的TaskId。
//
// 2. 也可通过 DescribeLivePullStreamTasks 进行查询创建的任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteLivePullStreamTaskWithContext(ctx context.Context, request *DeleteLivePullStreamTaskRequest) (response *DeleteLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewDeleteLivePullStreamTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLivePullStreamTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLivePullStreamTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLivePullStreamTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordRequest() (request *DeleteLiveRecordRequest) {
    request = &DeleteLiveRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveRecord")
    
    
    return
}

func NewDeleteLiveRecordResponse() (response *DeleteLiveRecordResponse) {
    response = &DeleteLiveRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveRecord
// 注：DeleteLiveRecord 接口仅用于删除录制任务记录，不具备停止录制的功能，也不能删除正在进行中的录制。如果需要停止录制任务，请使用终止录制[StopLiveRecord](/document/product/267/30146) 接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALTERTASKSTATE = "FailedOperation.AlterTaskState"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecord(request *DeleteLiveRecordRequest) (response *DeleteLiveRecordResponse, err error) {
    return c.DeleteLiveRecordWithContext(context.Background(), request)
}

// DeleteLiveRecord
// 注：DeleteLiveRecord 接口仅用于删除录制任务记录，不具备停止录制的功能，也不能删除正在进行中的录制。如果需要停止录制任务，请使用终止录制[StopLiveRecord](/document/product/267/30146) 接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALTERTASKSTATE = "FailedOperation.AlterTaskState"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecordWithContext(ctx context.Context, request *DeleteLiveRecordRequest) (response *DeleteLiveRecordResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordRuleRequest() (request *DeleteLiveRecordRuleRequest) {
    request = &DeleteLiveRecordRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveRecordRule")
    
    
    return
}

func NewDeleteLiveRecordRuleResponse() (response *DeleteLiveRecordRuleResponse) {
    response = &DeleteLiveRecordRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveRecordRule
// 删除录制规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecordRule(request *DeleteLiveRecordRuleRequest) (response *DeleteLiveRecordRuleResponse, err error) {
    return c.DeleteLiveRecordRuleWithContext(context.Background(), request)
}

// DeleteLiveRecordRule
// 删除录制规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecordRuleWithContext(ctx context.Context, request *DeleteLiveRecordRuleRequest) (response *DeleteLiveRecordRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveRecordRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveRecordRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveRecordRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordTemplateRequest() (request *DeleteLiveRecordTemplateRequest) {
    request = &DeleteLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveRecordTemplate")
    
    
    return
}

func NewDeleteLiveRecordTemplateResponse() (response *DeleteLiveRecordTemplateResponse) {
    response = &DeleteLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveRecordTemplate
// 删除录制模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecordTemplate(request *DeleteLiveRecordTemplateRequest) (response *DeleteLiveRecordTemplateResponse, err error) {
    return c.DeleteLiveRecordTemplateWithContext(context.Background(), request)
}

// DeleteLiveRecordTemplate
// 删除录制模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecordTemplateWithContext(ctx context.Context, request *DeleteLiveRecordTemplateRequest) (response *DeleteLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveSnapshotRuleRequest() (request *DeleteLiveSnapshotRuleRequest) {
    request = &DeleteLiveSnapshotRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveSnapshotRule")
    
    
    return
}

func NewDeleteLiveSnapshotRuleResponse() (response *DeleteLiveSnapshotRuleResponse) {
    response = &DeleteLiveSnapshotRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveSnapshotRule
// 删除截图规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveSnapshotRule(request *DeleteLiveSnapshotRuleRequest) (response *DeleteLiveSnapshotRuleResponse, err error) {
    return c.DeleteLiveSnapshotRuleWithContext(context.Background(), request)
}

// DeleteLiveSnapshotRule
// 删除截图规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveSnapshotRuleWithContext(ctx context.Context, request *DeleteLiveSnapshotRuleRequest) (response *DeleteLiveSnapshotRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveSnapshotRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveSnapshotRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveSnapshotRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveSnapshotRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveSnapshotTemplateRequest() (request *DeleteLiveSnapshotTemplateRequest) {
    request = &DeleteLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveSnapshotTemplate")
    
    
    return
}

func NewDeleteLiveSnapshotTemplateResponse() (response *DeleteLiveSnapshotTemplateResponse) {
    response = &DeleteLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveSnapshotTemplate
// 删除截图模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveSnapshotTemplate(request *DeleteLiveSnapshotTemplateRequest) (response *DeleteLiveSnapshotTemplateResponse, err error) {
    return c.DeleteLiveSnapshotTemplateWithContext(context.Background(), request)
}

// DeleteLiveSnapshotTemplate
// 删除截图模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveSnapshotTemplateWithContext(ctx context.Context, request *DeleteLiveSnapshotTemplateRequest) (response *DeleteLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveStreamMonitorRequest() (request *DeleteLiveStreamMonitorRequest) {
    request = &DeleteLiveStreamMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveStreamMonitor")
    
    
    return
}

func NewDeleteLiveStreamMonitorResponse() (response *DeleteLiveStreamMonitorResponse) {
    response = &DeleteLiveStreamMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveStreamMonitor
// 该接口用来删除直播流监播任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveStreamMonitor(request *DeleteLiveStreamMonitorRequest) (response *DeleteLiveStreamMonitorResponse, err error) {
    return c.DeleteLiveStreamMonitorWithContext(context.Background(), request)
}

// DeleteLiveStreamMonitor
// 该接口用来删除直播流监播任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveStreamMonitorWithContext(ctx context.Context, request *DeleteLiveStreamMonitorRequest) (response *DeleteLiveStreamMonitorResponse, err error) {
    if request == nil {
        request = NewDeleteLiveStreamMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveStreamMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveStreamMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveStreamMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveTimeShiftRuleRequest() (request *DeleteLiveTimeShiftRuleRequest) {
    request = &DeleteLiveTimeShiftRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveTimeShiftRule")
    
    
    return
}

func NewDeleteLiveTimeShiftRuleResponse() (response *DeleteLiveTimeShiftRuleResponse) {
    response = &DeleteLiveTimeShiftRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveTimeShiftRule
// 删除直播时移规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTimeShiftRule(request *DeleteLiveTimeShiftRuleRequest) (response *DeleteLiveTimeShiftRuleResponse, err error) {
    return c.DeleteLiveTimeShiftRuleWithContext(context.Background(), request)
}

// DeleteLiveTimeShiftRule
// 删除直播时移规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTimeShiftRuleWithContext(ctx context.Context, request *DeleteLiveTimeShiftRuleRequest) (response *DeleteLiveTimeShiftRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTimeShiftRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveTimeShiftRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveTimeShiftRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveTimeShiftRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveTimeShiftTemplateRequest() (request *DeleteLiveTimeShiftTemplateRequest) {
    request = &DeleteLiveTimeShiftTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveTimeShiftTemplate")
    
    
    return
}

func NewDeleteLiveTimeShiftTemplateResponse() (response *DeleteLiveTimeShiftTemplateResponse) {
    response = &DeleteLiveTimeShiftTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveTimeShiftTemplate
// 删除直播时移模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTimeShiftTemplate(request *DeleteLiveTimeShiftTemplateRequest) (response *DeleteLiveTimeShiftTemplateResponse, err error) {
    return c.DeleteLiveTimeShiftTemplateWithContext(context.Background(), request)
}

// DeleteLiveTimeShiftTemplate
// 删除直播时移模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTimeShiftTemplateWithContext(ctx context.Context, request *DeleteLiveTimeShiftTemplateRequest) (response *DeleteLiveTimeShiftTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTimeShiftTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveTimeShiftTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveTimeShiftTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveTimeShiftTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveTranscodeRuleRequest() (request *DeleteLiveTranscodeRuleRequest) {
    request = &DeleteLiveTranscodeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveTranscodeRule")
    
    
    return
}

func NewDeleteLiveTranscodeRuleResponse() (response *DeleteLiveTranscodeRuleResponse) {
    response = &DeleteLiveTranscodeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveTranscodeRule
// 删除转码规则。
//
// DomainName+AppName+StreamName+TemplateId唯一标识单个转码规则，如需删除需要强匹配。其中TemplateId必填，其余参数为空时也需要传空字符串进行强匹配。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_RULENOTFOUND = "InvalidParameter.RuleNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTranscodeRule(request *DeleteLiveTranscodeRuleRequest) (response *DeleteLiveTranscodeRuleResponse, err error) {
    return c.DeleteLiveTranscodeRuleWithContext(context.Background(), request)
}

// DeleteLiveTranscodeRule
// 删除转码规则。
//
// DomainName+AppName+StreamName+TemplateId唯一标识单个转码规则，如需删除需要强匹配。其中TemplateId必填，其余参数为空时也需要传空字符串进行强匹配。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_RULENOTFOUND = "InvalidParameter.RuleNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTranscodeRuleWithContext(ctx context.Context, request *DeleteLiveTranscodeRuleRequest) (response *DeleteLiveTranscodeRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTranscodeRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveTranscodeRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveTranscodeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveTranscodeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveTranscodeTemplateRequest() (request *DeleteLiveTranscodeTemplateRequest) {
    request = &DeleteLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveTranscodeTemplate")
    
    
    return
}

func NewDeleteLiveTranscodeTemplateResponse() (response *DeleteLiveTranscodeTemplateResponse) {
    response = &DeleteLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveTranscodeTemplate
// 删除转码模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTranscodeTemplate(request *DeleteLiveTranscodeTemplateRequest) (response *DeleteLiveTranscodeTemplateResponse, err error) {
    return c.DeleteLiveTranscodeTemplateWithContext(context.Background(), request)
}

// DeleteLiveTranscodeTemplate
// 删除转码模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTranscodeTemplateWithContext(ctx context.Context, request *DeleteLiveTranscodeTemplateRequest) (response *DeleteLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveWatermarkRequest() (request *DeleteLiveWatermarkRequest) {
    request = &DeleteLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveWatermark")
    
    
    return
}

func NewDeleteLiveWatermarkResponse() (response *DeleteLiveWatermarkResponse) {
    response = &DeleteLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveWatermark
// 删除水印。
//
// 可能返回的错误码:
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETWATERMARKERROR = "InternalError.GetWatermarkError"
//  INTERNALERROR_WATERMARKNOTEXIST = "InternalError.WatermarkNotExist"
//  INVALIDPARAMETER_CONFINUSED = "InvalidParameter.ConfInUsed"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) DeleteLiveWatermark(request *DeleteLiveWatermarkRequest) (response *DeleteLiveWatermarkResponse, err error) {
    return c.DeleteLiveWatermarkWithContext(context.Background(), request)
}

// DeleteLiveWatermark
// 删除水印。
//
// 可能返回的错误码:
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETWATERMARKERROR = "InternalError.GetWatermarkError"
//  INTERNALERROR_WATERMARKNOTEXIST = "InternalError.WatermarkNotExist"
//  INVALIDPARAMETER_CONFINUSED = "InvalidParameter.ConfInUsed"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) DeleteLiveWatermarkWithContext(ctx context.Context, request *DeleteLiveWatermarkRequest) (response *DeleteLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewDeleteLiveWatermarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveWatermark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveWatermarkRuleRequest() (request *DeleteLiveWatermarkRuleRequest) {
    request = &DeleteLiveWatermarkRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveWatermarkRule")
    
    
    return
}

func NewDeleteLiveWatermarkRuleResponse() (response *DeleteLiveWatermarkRuleResponse) {
    response = &DeleteLiveWatermarkRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveWatermarkRule
// 删除水印规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveWatermarkRule(request *DeleteLiveWatermarkRuleRequest) (response *DeleteLiveWatermarkRuleResponse, err error) {
    return c.DeleteLiveWatermarkRuleWithContext(context.Background(), request)
}

// DeleteLiveWatermarkRule
// 删除水印规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveWatermarkRuleWithContext(ctx context.Context, request *DeleteLiveWatermarkRuleRequest) (response *DeleteLiveWatermarkRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveWatermarkRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteLiveWatermarkRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveWatermarkRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveWatermarkRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePullStreamConfigRequest() (request *DeletePullStreamConfigRequest) {
    request = &DeletePullStreamConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeletePullStreamConfig")
    
    
    return
}

func NewDeletePullStreamConfigResponse() (response *DeletePullStreamConfigResponse) {
    response = &DeletePullStreamConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePullStreamConfig
// 删除直播拉流配置。该接口已下线,请使用新接口 DeleteLivePullStreamTask。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeletePullStreamConfig(request *DeletePullStreamConfigRequest) (response *DeletePullStreamConfigResponse, err error) {
    return c.DeletePullStreamConfigWithContext(context.Background(), request)
}

// DeletePullStreamConfig
// 删除直播拉流配置。该接口已下线,请使用新接口 DeleteLivePullStreamTask。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeletePullStreamConfigWithContext(ctx context.Context, request *DeletePullStreamConfigRequest) (response *DeletePullStreamConfigResponse, err error) {
    if request == nil {
        request = NewDeletePullStreamConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeletePullStreamConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePullStreamConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePullStreamConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordTaskRequest() (request *DeleteRecordTaskRequest) {
    request = &DeleteRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteRecordTask")
    
    
    return
}

func NewDeleteRecordTaskResponse() (response *DeleteRecordTaskResponse) {
    response = &DeleteRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordTask
// 删除录制任务配置。删除操作不影响正在运行当中的任务，仅对删除之后新的推流有效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecordTask(request *DeleteRecordTaskRequest) (response *DeleteRecordTaskResponse, err error) {
    return c.DeleteRecordTaskWithContext(context.Background(), request)
}

// DeleteRecordTask
// 删除录制任务配置。删除操作不影响正在运行当中的任务，仅对删除之后新的推流有效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecordTaskWithContext(ctx context.Context, request *DeleteRecordTaskRequest) (response *DeleteRecordTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRecordTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteRecordTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScreenshotTaskRequest() (request *DeleteScreenshotTaskRequest) {
    request = &DeleteScreenshotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteScreenshotTask")
    
    
    return
}

func NewDeleteScreenshotTaskResponse() (response *DeleteScreenshotTaskResponse) {
    response = &DeleteScreenshotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteScreenshotTask
// 删除截图任务配置。删除操作不影响正在运行当中的任务，仅对删除之后新的推流有效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteScreenshotTask(request *DeleteScreenshotTaskRequest) (response *DeleteScreenshotTaskResponse, err error) {
    return c.DeleteScreenshotTaskWithContext(context.Background(), request)
}

// DeleteScreenshotTask
// 删除截图任务配置。删除操作不影响正在运行当中的任务，仅对删除之后新的推流有效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteScreenshotTaskWithContext(ctx context.Context, request *DeleteScreenshotTaskRequest) (response *DeleteScreenshotTaskResponse, err error) {
    if request == nil {
        request = NewDeleteScreenshotTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DeleteScreenshotTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScreenshotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScreenshotTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllStreamPlayInfoListRequest() (request *DescribeAllStreamPlayInfoListRequest) {
    request = &DescribeAllStreamPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeAllStreamPlayInfoList")
    
    
    return
}

func NewDescribeAllStreamPlayInfoListResponse() (response *DescribeAllStreamPlayInfoListResponse) {
    response = &DescribeAllStreamPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllStreamPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 输入某个时间点（1分钟维度），查询该时间点所有流的下行信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeAllStreamPlayInfoList(request *DescribeAllStreamPlayInfoListRequest) (response *DescribeAllStreamPlayInfoListResponse, err error) {
    return c.DescribeAllStreamPlayInfoListWithContext(context.Background(), request)
}

// DescribeAllStreamPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 输入某个时间点（1分钟维度），查询该时间点所有流的下行信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeAllStreamPlayInfoListWithContext(ctx context.Context, request *DescribeAllStreamPlayInfoListRequest) (response *DescribeAllStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeAllStreamPlayInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeAllStreamPlayInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllStreamPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllStreamPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAreaBillBandwidthAndFluxListRequest() (request *DescribeAreaBillBandwidthAndFluxListRequest) {
    request = &DescribeAreaBillBandwidthAndFluxListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeAreaBillBandwidthAndFluxList")
    
    
    return
}

func NewDescribeAreaBillBandwidthAndFluxListResponse() (response *DescribeAreaBillBandwidthAndFluxListResponse) {
    response = &DescribeAreaBillBandwidthAndFluxListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAreaBillBandwidthAndFluxList
// 海外分区直播播放带宽和流量数据查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeAreaBillBandwidthAndFluxList(request *DescribeAreaBillBandwidthAndFluxListRequest) (response *DescribeAreaBillBandwidthAndFluxListResponse, err error) {
    return c.DescribeAreaBillBandwidthAndFluxListWithContext(context.Background(), request)
}

// DescribeAreaBillBandwidthAndFluxList
// 海外分区直播播放带宽和流量数据查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeAreaBillBandwidthAndFluxListWithContext(ctx context.Context, request *DescribeAreaBillBandwidthAndFluxListRequest) (response *DescribeAreaBillBandwidthAndFluxListResponse, err error) {
    if request == nil {
        request = NewDescribeAreaBillBandwidthAndFluxListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeAreaBillBandwidthAndFluxList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAreaBillBandwidthAndFluxList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAreaBillBandwidthAndFluxListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditKeywordsRequest() (request *DescribeAuditKeywordsRequest) {
    request = &DescribeAuditKeywordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeAuditKeywords")
    
    
    return
}

func NewDescribeAuditKeywordsResponse() (response *DescribeAuditKeywordsResponse) {
    response = &DescribeAuditKeywordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditKeywords
// 获取关键词信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAuditKeywords(request *DescribeAuditKeywordsRequest) (response *DescribeAuditKeywordsResponse, err error) {
    return c.DescribeAuditKeywordsWithContext(context.Background(), request)
}

// DescribeAuditKeywords
// 获取关键词信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAuditKeywordsWithContext(ctx context.Context, request *DescribeAuditKeywordsRequest) (response *DescribeAuditKeywordsResponse, err error) {
    if request == nil {
        request = NewDescribeAuditKeywordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeAuditKeywords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditKeywords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditKeywordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupStreamListRequest() (request *DescribeBackupStreamListRequest) {
    request = &DescribeBackupStreamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeBackupStreamList")
    
    
    return
}

func NewDescribeBackupStreamListResponse() (response *DescribeBackupStreamListResponse) {
    response = &DescribeBackupStreamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupStreamList
// 返回正在直播中的流列表。适用于推流成功后查询在线流信息。
//
// 
//
// 注意：
//
// 1. 该接口仅提供辅助查询在线流列表功能，业务重要场景不可强依赖该接口。
//
// 2. 该接口仅适用于流数少于2万路的情况，对于流数较大用户请联系售后。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeBackupStreamList(request *DescribeBackupStreamListRequest) (response *DescribeBackupStreamListResponse, err error) {
    return c.DescribeBackupStreamListWithContext(context.Background(), request)
}

// DescribeBackupStreamList
// 返回正在直播中的流列表。适用于推流成功后查询在线流信息。
//
// 
//
// 注意：
//
// 1. 该接口仅提供辅助查询在线流列表功能，业务重要场景不可强依赖该接口。
//
// 2. 该接口仅适用于流数少于2万路的情况，对于流数较大用户请联系售后。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeBackupStreamListWithContext(ctx context.Context, request *DescribeBackupStreamListRequest) (response *DescribeBackupStreamListResponse, err error) {
    if request == nil {
        request = NewDescribeBackupStreamListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeBackupStreamList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupStreamList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupStreamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillBandwidthAndFluxListRequest() (request *DescribeBillBandwidthAndFluxListRequest) {
    request = &DescribeBillBandwidthAndFluxListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeBillBandwidthAndFluxList")
    
    
    return
}

func NewDescribeBillBandwidthAndFluxListResponse() (response *DescribeBillBandwidthAndFluxListResponse) {
    response = &DescribeBillBandwidthAndFluxListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillBandwidthAndFluxList
// 直播播放带宽和流量数据查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeBillBandwidthAndFluxList(request *DescribeBillBandwidthAndFluxListRequest) (response *DescribeBillBandwidthAndFluxListResponse, err error) {
    return c.DescribeBillBandwidthAndFluxListWithContext(context.Background(), request)
}

// DescribeBillBandwidthAndFluxList
// 直播播放带宽和流量数据查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeBillBandwidthAndFluxListWithContext(ctx context.Context, request *DescribeBillBandwidthAndFluxListRequest) (response *DescribeBillBandwidthAndFluxListResponse, err error) {
    if request == nil {
        request = NewDescribeBillBandwidthAndFluxListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeBillBandwidthAndFluxList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillBandwidthAndFluxList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillBandwidthAndFluxListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallbackRecordsListRequest() (request *DescribeCallbackRecordsListRequest) {
    request = &DescribeCallbackRecordsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCallbackRecordsList")
    
    
    return
}

func NewDescribeCallbackRecordsListResponse() (response *DescribeCallbackRecordsListResponse) {
    response = &DescribeCallbackRecordsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCallbackRecordsList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 用于查询回调事件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeCallbackRecordsList(request *DescribeCallbackRecordsListRequest) (response *DescribeCallbackRecordsListResponse, err error) {
    return c.DescribeCallbackRecordsListWithContext(context.Background(), request)
}

// DescribeCallbackRecordsList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 用于查询回调事件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeCallbackRecordsListWithContext(ctx context.Context, request *DescribeCallbackRecordsListRequest) (response *DescribeCallbackRecordsListResponse, err error) {
    if request == nil {
        request = NewDescribeCallbackRecordsListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCallbackRecordsList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallbackRecordsList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallbackRecordsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterRequest() (request *DescribeCasterRequest) {
    request = &DescribeCasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCaster")
    
    
    return
}

func NewDescribeCasterResponse() (response *DescribeCasterResponse) {
    response = &DescribeCasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCaster
// 查询导播台信息接口，用来查询导播台状态、描述、输出长、宽等信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_OUTOFSERVICE = "FailedOperation.OutOfService"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCaster(request *DescribeCasterRequest) (response *DescribeCasterResponse, err error) {
    return c.DescribeCasterWithContext(context.Background(), request)
}

// DescribeCaster
// 查询导播台信息接口，用来查询导播台状态、描述、输出长、宽等信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_OUTOFSERVICE = "FailedOperation.OutOfService"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterWithContext(ctx context.Context, request *DescribeCasterRequest) (response *DescribeCasterResponse, err error) {
    if request == nil {
        request = NewDescribeCasterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCaster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterDisplayInfoRequest() (request *DescribeCasterDisplayInfoRequest) {
    request = &DescribeCasterDisplayInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCasterDisplayInfo")
    
    
    return
}

func NewDescribeCasterDisplayInfoResponse() (response *DescribeCasterDisplayInfoResponse) {
    response = &DescribeCasterDisplayInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCasterDisplayInfo
// 查询导播台PVW任务和PGM任务的展示信息，包括使用的布局、水印、字幕等信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_CASTERNOTEXISTS = "ResourceNotFound.CasterNotExists"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterDisplayInfo(request *DescribeCasterDisplayInfoRequest) (response *DescribeCasterDisplayInfoResponse, err error) {
    return c.DescribeCasterDisplayInfoWithContext(context.Background(), request)
}

// DescribeCasterDisplayInfo
// 查询导播台PVW任务和PGM任务的展示信息，包括使用的布局、水印、字幕等信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_CASTERNOTEXISTS = "ResourceNotFound.CasterNotExists"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterDisplayInfoWithContext(ctx context.Context, request *DescribeCasterDisplayInfoRequest) (response *DescribeCasterDisplayInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCasterDisplayInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCasterDisplayInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCasterDisplayInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterDisplayInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterInputInfosRequest() (request *DescribeCasterInputInfosRequest) {
    request = &DescribeCasterInputInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCasterInputInfos")
    
    
    return
}

func NewDescribeCasterInputInfosResponse() (response *DescribeCasterInputInfosResponse) {
    response = &DescribeCasterInputInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCasterInputInfos
// 该接口用来查询导播台的输入源信息列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPROTOCOL = "InvalidParameter.InvalidProtocol"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterInputInfos(request *DescribeCasterInputInfosRequest) (response *DescribeCasterInputInfosResponse, err error) {
    return c.DescribeCasterInputInfosWithContext(context.Background(), request)
}

// DescribeCasterInputInfos
// 该接口用来查询导播台的输入源信息列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPROTOCOL = "InvalidParameter.InvalidProtocol"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterInputInfosWithContext(ctx context.Context, request *DescribeCasterInputInfosRequest) (response *DescribeCasterInputInfosResponse, err error) {
    if request == nil {
        request = NewDescribeCasterInputInfosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCasterInputInfos")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCasterInputInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterInputInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterLayoutInfosRequest() (request *DescribeCasterLayoutInfosRequest) {
    request = &DescribeCasterLayoutInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCasterLayoutInfos")
    
    
    return
}

func NewDescribeCasterLayoutInfosResponse() (response *DescribeCasterLayoutInfosResponse) {
    response = &DescribeCasterLayoutInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCasterLayoutInfos
// 该接口用来查询某个导播台的布局列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterLayoutInfos(request *DescribeCasterLayoutInfosRequest) (response *DescribeCasterLayoutInfosResponse, err error) {
    return c.DescribeCasterLayoutInfosWithContext(context.Background(), request)
}

// DescribeCasterLayoutInfos
// 该接口用来查询某个导播台的布局列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterLayoutInfosWithContext(ctx context.Context, request *DescribeCasterLayoutInfosRequest) (response *DescribeCasterLayoutInfosResponse, err error) {
    if request == nil {
        request = NewDescribeCasterLayoutInfosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCasterLayoutInfos")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCasterLayoutInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterLayoutInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterListRequest() (request *DescribeCasterListRequest) {
    request = &DescribeCasterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCasterList")
    
    
    return
}

func NewDescribeCasterListResponse() (response *DescribeCasterListResponse) {
    response = &DescribeCasterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCasterList
// 该接口用来查询账号下所有的导播台列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterList(request *DescribeCasterListRequest) (response *DescribeCasterListResponse, err error) {
    return c.DescribeCasterListWithContext(context.Background(), request)
}

// DescribeCasterList
// 该接口用来查询账号下所有的导播台列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterListWithContext(ctx context.Context, request *DescribeCasterListRequest) (response *DescribeCasterListResponse, err error) {
    if request == nil {
        request = NewDescribeCasterListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCasterList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCasterList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterMarkPicInfosRequest() (request *DescribeCasterMarkPicInfosRequest) {
    request = &DescribeCasterMarkPicInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCasterMarkPicInfos")
    
    
    return
}

func NewDescribeCasterMarkPicInfosResponse() (response *DescribeCasterMarkPicInfosResponse) {
    response = &DescribeCasterMarkPicInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCasterMarkPicInfos
// 该接口用来查询某个导播台的水印列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterMarkPicInfos(request *DescribeCasterMarkPicInfosRequest) (response *DescribeCasterMarkPicInfosResponse, err error) {
    return c.DescribeCasterMarkPicInfosWithContext(context.Background(), request)
}

// DescribeCasterMarkPicInfos
// 该接口用来查询某个导播台的水印列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterMarkPicInfosWithContext(ctx context.Context, request *DescribeCasterMarkPicInfosRequest) (response *DescribeCasterMarkPicInfosResponse, err error) {
    if request == nil {
        request = NewDescribeCasterMarkPicInfosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCasterMarkPicInfos")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCasterMarkPicInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterMarkPicInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterMarkWordInfosRequest() (request *DescribeCasterMarkWordInfosRequest) {
    request = &DescribeCasterMarkWordInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCasterMarkWordInfos")
    
    
    return
}

func NewDescribeCasterMarkWordInfosResponse() (response *DescribeCasterMarkWordInfosResponse) {
    response = &DescribeCasterMarkWordInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCasterMarkWordInfos
// 该接口用来查询某个导播台的文本列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterMarkWordInfos(request *DescribeCasterMarkWordInfosRequest) (response *DescribeCasterMarkWordInfosResponse, err error) {
    return c.DescribeCasterMarkWordInfosWithContext(context.Background(), request)
}

// DescribeCasterMarkWordInfos
// 该接口用来查询某个导播台的文本列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterMarkWordInfosWithContext(ctx context.Context, request *DescribeCasterMarkWordInfosRequest) (response *DescribeCasterMarkWordInfosResponse, err error) {
    if request == nil {
        request = NewDescribeCasterMarkWordInfosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCasterMarkWordInfos")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCasterMarkWordInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterMarkWordInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterOutputInfosRequest() (request *DescribeCasterOutputInfosRequest) {
    request = &DescribeCasterOutputInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCasterOutputInfos")
    
    
    return
}

func NewDescribeCasterOutputInfosResponse() (response *DescribeCasterOutputInfosResponse) {
    response = &DescribeCasterOutputInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCasterOutputInfos
// 该接口用来查询某个导播台的推流信息列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterOutputInfos(request *DescribeCasterOutputInfosRequest) (response *DescribeCasterOutputInfosResponse, err error) {
    return c.DescribeCasterOutputInfosWithContext(context.Background(), request)
}

// DescribeCasterOutputInfos
// 该接口用来查询某个导播台的推流信息列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterOutputInfosWithContext(ctx context.Context, request *DescribeCasterOutputInfosRequest) (response *DescribeCasterOutputInfosResponse, err error) {
    if request == nil {
        request = NewDescribeCasterOutputInfosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCasterOutputInfos")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCasterOutputInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterOutputInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterPlayUrlRequest() (request *DescribeCasterPlayUrlRequest) {
    request = &DescribeCasterPlayUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCasterPlayUrl")
    
    
    return
}

func NewDescribeCasterPlayUrlResponse() (response *DescribeCasterPlayUrlResponse) {
    response = &DescribeCasterPlayUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCasterPlayUrl
// 该接口用来获取导播台视频流的播放url，用来在页面上拉流展示。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_INPUTISNOTACTIVE = "FailedOperation.InputIsNotActive"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterPlayUrl(request *DescribeCasterPlayUrlRequest) (response *DescribeCasterPlayUrlResponse, err error) {
    return c.DescribeCasterPlayUrlWithContext(context.Background(), request)
}

// DescribeCasterPlayUrl
// 该接口用来获取导播台视频流的播放url，用来在页面上拉流展示。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_INPUTISNOTACTIVE = "FailedOperation.InputIsNotActive"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterPlayUrlWithContext(ctx context.Context, request *DescribeCasterPlayUrlRequest) (response *DescribeCasterPlayUrlResponse, err error) {
    if request == nil {
        request = NewDescribeCasterPlayUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCasterPlayUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCasterPlayUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterPlayUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterTransitionTypesRequest() (request *DescribeCasterTransitionTypesRequest) {
    request = &DescribeCasterTransitionTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCasterTransitionTypes")
    
    
    return
}

func NewDescribeCasterTransitionTypesResponse() (response *DescribeCasterTransitionTypesResponse) {
    response = &DescribeCasterTransitionTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCasterTransitionTypes
// 该接口用来获取所有的转场名称及其对应的素材url。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeCasterTransitionTypes(request *DescribeCasterTransitionTypesRequest) (response *DescribeCasterTransitionTypesResponse, err error) {
    return c.DescribeCasterTransitionTypesWithContext(context.Background(), request)
}

// DescribeCasterTransitionTypes
// 该接口用来获取所有的转场名称及其对应的素材url。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeCasterTransitionTypesWithContext(ctx context.Context, request *DescribeCasterTransitionTypesRequest) (response *DescribeCasterTransitionTypesResponse, err error) {
    if request == nil {
        request = NewDescribeCasterTransitionTypesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCasterTransitionTypes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCasterTransitionTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterTransitionTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCasterUserStatusRequest() (request *DescribeCasterUserStatusRequest) {
    request = &DescribeCasterUserStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeCasterUserStatus")
    
    
    return
}

func NewDescribeCasterUserStatusResponse() (response *DescribeCasterUserStatusResponse) {
    response = &DescribeCasterUserStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCasterUserStatus
// 本接口用来查询当前APPID导播台业务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterUserStatus(request *DescribeCasterUserStatusRequest) (response *DescribeCasterUserStatusResponse, err error) {
    return c.DescribeCasterUserStatusWithContext(context.Background(), request)
}

// DescribeCasterUserStatus
// 本接口用来查询当前APPID导播台业务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCasterUserStatusWithContext(ctx context.Context, request *DescribeCasterUserStatusRequest) (response *DescribeCasterUserStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCasterUserStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeCasterUserStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCasterUserStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCasterUserStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrentRecordStreamNumRequest() (request *DescribeConcurrentRecordStreamNumRequest) {
    request = &DescribeConcurrentRecordStreamNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeConcurrentRecordStreamNum")
    
    
    return
}

func NewDescribeConcurrentRecordStreamNumResponse() (response *DescribeConcurrentRecordStreamNumResponse) {
    response = &DescribeConcurrentRecordStreamNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrentRecordStreamNum
// 查询并发录制路数，对慢直播和普通直播适用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeConcurrentRecordStreamNum(request *DescribeConcurrentRecordStreamNumRequest) (response *DescribeConcurrentRecordStreamNumResponse, err error) {
    return c.DescribeConcurrentRecordStreamNumWithContext(context.Background(), request)
}

// DescribeConcurrentRecordStreamNum
// 查询并发录制路数，对慢直播和普通直播适用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeConcurrentRecordStreamNumWithContext(ctx context.Context, request *DescribeConcurrentRecordStreamNumRequest) (response *DescribeConcurrentRecordStreamNumResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrentRecordStreamNumRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeConcurrentRecordStreamNum")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrentRecordStreamNum require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrentRecordStreamNumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeliverBandwidthListRequest() (request *DescribeDeliverBandwidthListRequest) {
    request = &DescribeDeliverBandwidthListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeDeliverBandwidthList")
    
    
    return
}

func NewDescribeDeliverBandwidthListResponse() (response *DescribeDeliverBandwidthListResponse) {
    response = &DescribeDeliverBandwidthListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeliverBandwidthList
// 查询直播转推计费带宽，查询时间范围最大支持3个月内的数据，时间跨度最长31天。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeliverBandwidthList(request *DescribeDeliverBandwidthListRequest) (response *DescribeDeliverBandwidthListResponse, err error) {
    return c.DescribeDeliverBandwidthListWithContext(context.Background(), request)
}

// DescribeDeliverBandwidthList
// 查询直播转推计费带宽，查询时间范围最大支持3个月内的数据，时间跨度最长31天。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeliverBandwidthListWithContext(ctx context.Context, request *DescribeDeliverBandwidthListRequest) (response *DescribeDeliverBandwidthListResponse, err error) {
    if request == nil {
        request = NewDescribeDeliverBandwidthListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeDeliverBandwidthList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeliverBandwidthList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeliverBandwidthListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeliverLogDownListRequest() (request *DescribeDeliverLogDownListRequest) {
    request = &DescribeDeliverLogDownListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeDeliverLogDownList")
    
    
    return
}

func NewDescribeDeliverLogDownListResponse() (response *DescribeDeliverLogDownListResponse) {
    response = &DescribeDeliverLogDownListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeliverLogDownList
// 批量获取转推日志的URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVOKECDNAPIFAIL = "FailedOperation.InvokeCdnApiFail"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CDNLOGEMPTY = "ResourceNotFound.CdnLogEmpty"
//  RESOURCENOTFOUND_CDNTHEMEEMPTY = "ResourceNotFound.CdnThemeEmpty"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeDeliverLogDownList(request *DescribeDeliverLogDownListRequest) (response *DescribeDeliverLogDownListResponse, err error) {
    return c.DescribeDeliverLogDownListWithContext(context.Background(), request)
}

// DescribeDeliverLogDownList
// 批量获取转推日志的URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVOKECDNAPIFAIL = "FailedOperation.InvokeCdnApiFail"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CDNLOGEMPTY = "ResourceNotFound.CdnLogEmpty"
//  RESOURCENOTFOUND_CDNTHEMEEMPTY = "ResourceNotFound.CdnThemeEmpty"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeDeliverLogDownListWithContext(ctx context.Context, request *DescribeDeliverLogDownListRequest) (response *DescribeDeliverLogDownListResponse, err error) {
    if request == nil {
        request = NewDescribeDeliverLogDownListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeDeliverLogDownList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeliverLogDownList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeliverLogDownListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupProIspPlayInfoListRequest() (request *DescribeGroupProIspPlayInfoListRequest) {
    request = &DescribeGroupProIspPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeGroupProIspPlayInfoList")
    
    
    return
}

func NewDescribeGroupProIspPlayInfoListResponse() (response *DescribeGroupProIspPlayInfoListResponse) {
    response = &DescribeGroupProIspPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupProIspPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询按省份和运营商分组的下行播放数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupProIspPlayInfoList(request *DescribeGroupProIspPlayInfoListRequest) (response *DescribeGroupProIspPlayInfoListResponse, err error) {
    return c.DescribeGroupProIspPlayInfoListWithContext(context.Background(), request)
}

// DescribeGroupProIspPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询按省份和运营商分组的下行播放数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupProIspPlayInfoListWithContext(ctx context.Context, request *DescribeGroupProIspPlayInfoListRequest) (response *DescribeGroupProIspPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeGroupProIspPlayInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeGroupProIspPlayInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupProIspPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupProIspPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHttpStatusInfoListRequest() (request *DescribeHttpStatusInfoListRequest) {
    request = &DescribeHttpStatusInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeHttpStatusInfoList")
    
    
    return
}

func NewDescribeHttpStatusInfoListResponse() (response *DescribeHttpStatusInfoListResponse) {
    response = &DescribeHttpStatusInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHttpStatusInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询某段时间内5分钟粒度的各播放http状态码的个数。
//
// 备注：数据延迟1小时，如10:00-10:59点的数据12点才能查到。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeHttpStatusInfoList(request *DescribeHttpStatusInfoListRequest) (response *DescribeHttpStatusInfoListResponse, err error) {
    return c.DescribeHttpStatusInfoListWithContext(context.Background(), request)
}

// DescribeHttpStatusInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询某段时间内5分钟粒度的各播放http状态码的个数。
//
// 备注：数据延迟1小时，如10:00-10:59点的数据12点才能查到。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeHttpStatusInfoListWithContext(ctx context.Context, request *DescribeHttpStatusInfoListRequest) (response *DescribeHttpStatusInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeHttpStatusInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeHttpStatusInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHttpStatusInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHttpStatusInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCallbackRulesRequest() (request *DescribeLiveCallbackRulesRequest) {
    request = &DescribeLiveCallbackRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCallbackRules")
    
    
    return
}

func NewDescribeLiveCallbackRulesResponse() (response *DescribeLiveCallbackRulesResponse) {
    response = &DescribeLiveCallbackRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveCallbackRules
// 获取回调规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveCallbackRules(request *DescribeLiveCallbackRulesRequest) (response *DescribeLiveCallbackRulesResponse, err error) {
    return c.DescribeLiveCallbackRulesWithContext(context.Background(), request)
}

// DescribeLiveCallbackRules
// 获取回调规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveCallbackRulesWithContext(ctx context.Context, request *DescribeLiveCallbackRulesRequest) (response *DescribeLiveCallbackRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveCallbackRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCallbackRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveCallbackRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCallbackTemplateRequest() (request *DescribeLiveCallbackTemplateRequest) {
    request = &DescribeLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCallbackTemplate")
    
    
    return
}

func NewDescribeLiveCallbackTemplateResponse() (response *DescribeLiveCallbackTemplateResponse) {
    response = &DescribeLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveCallbackTemplate
// 获取单个回调模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCallbackTemplate(request *DescribeLiveCallbackTemplateRequest) (response *DescribeLiveCallbackTemplateResponse, err error) {
    return c.DescribeLiveCallbackTemplateWithContext(context.Background(), request)
}

// DescribeLiveCallbackTemplate
// 获取单个回调模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCallbackTemplateWithContext(ctx context.Context, request *DescribeLiveCallbackTemplateRequest) (response *DescribeLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveCallbackTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCallbackTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCallbackTemplatesRequest() (request *DescribeLiveCallbackTemplatesRequest) {
    request = &DescribeLiveCallbackTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCallbackTemplates")
    
    
    return
}

func NewDescribeLiveCallbackTemplatesResponse() (response *DescribeLiveCallbackTemplatesResponse) {
    response = &DescribeLiveCallbackTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveCallbackTemplates
// 获取回调模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveCallbackTemplates(request *DescribeLiveCallbackTemplatesRequest) (response *DescribeLiveCallbackTemplatesResponse, err error) {
    return c.DescribeLiveCallbackTemplatesWithContext(context.Background(), request)
}

// DescribeLiveCallbackTemplates
// 获取回调模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveCallbackTemplatesWithContext(ctx context.Context, request *DescribeLiveCallbackTemplatesRequest) (response *DescribeLiveCallbackTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveCallbackTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCallbackTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveCallbackTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCertRequest() (request *DescribeLiveCertRequest) {
    request = &DescribeLiveCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCert")
    
    
    return
}

func NewDescribeLiveCertResponse() (response *DescribeLiveCertResponse) {
    response = &DescribeLiveCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveCert
// 获取证书信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCert(request *DescribeLiveCertRequest) (response *DescribeLiveCertResponse, err error) {
    return c.DescribeLiveCertWithContext(context.Background(), request)
}

// DescribeLiveCert
// 获取证书信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCertWithContext(ctx context.Context, request *DescribeLiveCertRequest) (response *DescribeLiveCertResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCertRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveCert")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCertsRequest() (request *DescribeLiveCertsRequest) {
    request = &DescribeLiveCertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCerts")
    
    
    return
}

func NewDescribeLiveCertsResponse() (response *DescribeLiveCertsResponse) {
    response = &DescribeLiveCertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveCerts
// 获取证书信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCerts(request *DescribeLiveCertsRequest) (response *DescribeLiveCertsResponse, err error) {
    return c.DescribeLiveCertsWithContext(context.Background(), request)
}

// DescribeLiveCerts
// 获取证书信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCertsWithContext(ctx context.Context, request *DescribeLiveCertsRequest) (response *DescribeLiveCertsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCertsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveCerts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCerts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveCertsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCloudEffectListRequest() (request *DescribeLiveCloudEffectListRequest) {
    request = &DescribeLiveCloudEffectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCloudEffectList")
    
    
    return
}

func NewDescribeLiveCloudEffectListResponse() (response *DescribeLiveCloudEffectListResponse) {
    response = &DescribeLiveCloudEffectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveCloudEffectList
// 使用该接口查询云端特效列表，特效列表中包含一部分官方精品特效，同时包含用户自定义生成的特效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveCloudEffectList(request *DescribeLiveCloudEffectListRequest) (response *DescribeLiveCloudEffectListResponse, err error) {
    return c.DescribeLiveCloudEffectListWithContext(context.Background(), request)
}

// DescribeLiveCloudEffectList
// 使用该接口查询云端特效列表，特效列表中包含一部分官方精品特效，同时包含用户自定义生成的特效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveCloudEffectListWithContext(ctx context.Context, request *DescribeLiveCloudEffectListRequest) (response *DescribeLiveCloudEffectListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCloudEffectListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveCloudEffectList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCloudEffectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveCloudEffectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDelayInfoListRequest() (request *DescribeLiveDelayInfoListRequest) {
    request = &DescribeLiveDelayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDelayInfoList")
    
    
    return
}

func NewDescribeLiveDelayInfoListResponse() (response *DescribeLiveDelayInfoListResponse) {
    response = &DescribeLiveDelayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveDelayInfoList
// 获取直播延播列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveDelayInfoList(request *DescribeLiveDelayInfoListRequest) (response *DescribeLiveDelayInfoListResponse, err error) {
    return c.DescribeLiveDelayInfoListWithContext(context.Background(), request)
}

// DescribeLiveDelayInfoList
// 获取直播延播列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveDelayInfoListWithContext(ctx context.Context, request *DescribeLiveDelayInfoListRequest) (response *DescribeLiveDelayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDelayInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveDelayInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDelayInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveDelayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainRequest() (request *DescribeLiveDomainRequest) {
    request = &DescribeLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomain")
    
    
    return
}

func NewDescribeLiveDomainResponse() (response *DescribeLiveDomainResponse) {
    response = &DescribeLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveDomain
// 查询直播域名信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomain(request *DescribeLiveDomainRequest) (response *DescribeLiveDomainResponse, err error) {
    return c.DescribeLiveDomainWithContext(context.Background(), request)
}

// DescribeLiveDomain
// 查询直播域名信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainWithContext(ctx context.Context, request *DescribeLiveDomainRequest) (response *DescribeLiveDomainResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainCertRequest() (request *DescribeLiveDomainCertRequest) {
    request = &DescribeLiveDomainCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainCert")
    
    
    return
}

func NewDescribeLiveDomainCertResponse() (response *DescribeLiveDomainCertResponse) {
    response = &DescribeLiveDomainCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveDomainCert
// 获取域名证书信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainCert(request *DescribeLiveDomainCertRequest) (response *DescribeLiveDomainCertResponse, err error) {
    return c.DescribeLiveDomainCertWithContext(context.Background(), request)
}

// DescribeLiveDomainCert
// 获取域名证书信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainCertWithContext(ctx context.Context, request *DescribeLiveDomainCertRequest) (response *DescribeLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainCertRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveDomainCert")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomainCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainCertBindingsRequest() (request *DescribeLiveDomainCertBindingsRequest) {
    request = &DescribeLiveDomainCertBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainCertBindings")
    
    
    return
}

func NewDescribeLiveDomainCertBindingsResponse() (response *DescribeLiveDomainCertBindingsResponse) {
    response = &DescribeLiveDomainCertBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveDomainCertBindings
// 查询绑定证书的域名列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainCertBindings(request *DescribeLiveDomainCertBindingsRequest) (response *DescribeLiveDomainCertBindingsResponse, err error) {
    return c.DescribeLiveDomainCertBindingsWithContext(context.Background(), request)
}

// DescribeLiveDomainCertBindings
// 查询绑定证书的域名列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainCertBindingsWithContext(ctx context.Context, request *DescribeLiveDomainCertBindingsRequest) (response *DescribeLiveDomainCertBindingsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainCertBindingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveDomainCertBindings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomainCertBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveDomainCertBindingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainPlayInfoListRequest() (request *DescribeLiveDomainPlayInfoListRequest) {
    request = &DescribeLiveDomainPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainPlayInfoList")
    
    
    return
}

func NewDescribeLiveDomainPlayInfoListResponse() (response *DescribeLiveDomainPlayInfoListResponse) {
    response = &DescribeLiveDomainPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveDomainPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询实时的域名维度下行播放数据，由于数据处理有耗时，接口默认查询4分钟前的准实时数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLiveDomainPlayInfoList(request *DescribeLiveDomainPlayInfoListRequest) (response *DescribeLiveDomainPlayInfoListResponse, err error) {
    return c.DescribeLiveDomainPlayInfoListWithContext(context.Background(), request)
}

// DescribeLiveDomainPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询实时的域名维度下行播放数据，由于数据处理有耗时，接口默认查询4分钟前的准实时数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLiveDomainPlayInfoListWithContext(ctx context.Context, request *DescribeLiveDomainPlayInfoListRequest) (response *DescribeLiveDomainPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainPlayInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveDomainPlayInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomainPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveDomainPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainRefererRequest() (request *DescribeLiveDomainRefererRequest) {
    request = &DescribeLiveDomainRefererRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainReferer")
    
    
    return
}

func NewDescribeLiveDomainRefererResponse() (response *DescribeLiveDomainRefererResponse) {
    response = &DescribeLiveDomainRefererResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveDomainReferer
// 查询直播域名 Referer 黑白名单配置。
//
// 由于 Referer 信息包含在 http 协议中，在开启配置后，播放协议为 rtmp 或 WebRTC 不会校验 Referer 配置，仍可正常播放。如需配置 Referer 鉴权建议使用 http-flv 或 http-hls 协议播放。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainReferer(request *DescribeLiveDomainRefererRequest) (response *DescribeLiveDomainRefererResponse, err error) {
    return c.DescribeLiveDomainRefererWithContext(context.Background(), request)
}

// DescribeLiveDomainReferer
// 查询直播域名 Referer 黑白名单配置。
//
// 由于 Referer 信息包含在 http 协议中，在开启配置后，播放协议为 rtmp 或 WebRTC 不会校验 Referer 配置，仍可正常播放。如需配置 Referer 鉴权建议使用 http-flv 或 http-hls 协议播放。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainRefererWithContext(ctx context.Context, request *DescribeLiveDomainRefererRequest) (response *DescribeLiveDomainRefererResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainRefererRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveDomainReferer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomainReferer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveDomainRefererResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainsRequest() (request *DescribeLiveDomainsRequest) {
    request = &DescribeLiveDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomains")
    
    
    return
}

func NewDescribeLiveDomainsResponse() (response *DescribeLiveDomainsResponse) {
    response = &DescribeLiveDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveDomains
// 根据域名状态、类型等信息查询用户的域名信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomains(request *DescribeLiveDomainsRequest) (response *DescribeLiveDomainsResponse, err error) {
    return c.DescribeLiveDomainsWithContext(context.Background(), request)
}

// DescribeLiveDomains
// 根据域名状态、类型等信息查询用户的域名信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainsWithContext(ctx context.Context, request *DescribeLiveDomainsRequest) (response *DescribeLiveDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveEnhanceInfoListRequest() (request *DescribeLiveEnhanceInfoListRequest) {
    request = &DescribeLiveEnhanceInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveEnhanceInfoList")
    
    
    return
}

func NewDescribeLiveEnhanceInfoListResponse() (response *DescribeLiveEnhanceInfoListResponse) {
    response = &DescribeLiveEnhanceInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveEnhanceInfoList
// 查询直播增强用量明细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveEnhanceInfoList(request *DescribeLiveEnhanceInfoListRequest) (response *DescribeLiveEnhanceInfoListResponse, err error) {
    return c.DescribeLiveEnhanceInfoListWithContext(context.Background(), request)
}

// DescribeLiveEnhanceInfoList
// 查询直播增强用量明细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveEnhanceInfoListWithContext(ctx context.Context, request *DescribeLiveEnhanceInfoListRequest) (response *DescribeLiveEnhanceInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveEnhanceInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveEnhanceInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveEnhanceInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveEnhanceInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveForbidStreamListRequest() (request *DescribeLiveForbidStreamListRequest) {
    request = &DescribeLiveForbidStreamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveForbidStreamList")
    
    
    return
}

func NewDescribeLiveForbidStreamListResponse() (response *DescribeLiveForbidStreamListResponse) {
    response = &DescribeLiveForbidStreamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveForbidStreamList
// 获取禁推流列表。
//
// 
//
// 注意：该接口仅作为直播辅助查询接口，重要业务场景不可强依赖该接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveForbidStreamList(request *DescribeLiveForbidStreamListRequest) (response *DescribeLiveForbidStreamListResponse, err error) {
    return c.DescribeLiveForbidStreamListWithContext(context.Background(), request)
}

// DescribeLiveForbidStreamList
// 获取禁推流列表。
//
// 
//
// 注意：该接口仅作为直播辅助查询接口，重要业务场景不可强依赖该接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveForbidStreamListWithContext(ctx context.Context, request *DescribeLiveForbidStreamListRequest) (response *DescribeLiveForbidStreamListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveForbidStreamListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveForbidStreamList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveForbidStreamList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveForbidStreamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePackageInfoRequest() (request *DescribeLivePackageInfoRequest) {
    request = &DescribeLivePackageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePackageInfo")
    
    
    return
}

func NewDescribeLivePackageInfoResponse() (response *DescribeLivePackageInfoResponse) {
    response = &DescribeLivePackageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLivePackageInfo
// 查询用户套餐包总量、使用量、剩余量、包状态、购买时间和过期时间等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePackageInfo(request *DescribeLivePackageInfoRequest) (response *DescribeLivePackageInfoResponse, err error) {
    return c.DescribeLivePackageInfoWithContext(context.Background(), request)
}

// DescribeLivePackageInfo
// 查询用户套餐包总量、使用量、剩余量、包状态、购买时间和过期时间等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePackageInfoWithContext(ctx context.Context, request *DescribeLivePackageInfoRequest) (response *DescribeLivePackageInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLivePackageInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLivePackageInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePackageInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePackageInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePadRulesRequest() (request *DescribeLivePadRulesRequest) {
    request = &DescribeLivePadRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePadRules")
    
    
    return
}

func NewDescribeLivePadRulesResponse() (response *DescribeLivePadRulesResponse) {
    response = &DescribeLivePadRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLivePadRules
// 获取直播垫片规则列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLivePadRules(request *DescribeLivePadRulesRequest) (response *DescribeLivePadRulesResponse, err error) {
    return c.DescribeLivePadRulesWithContext(context.Background(), request)
}

// DescribeLivePadRules
// 获取直播垫片规则列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLivePadRulesWithContext(ctx context.Context, request *DescribeLivePadRulesRequest) (response *DescribeLivePadRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLivePadRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLivePadRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePadRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePadRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePadStreamListRequest() (request *DescribeLivePadStreamListRequest) {
    request = &DescribeLivePadStreamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePadStreamList")
    
    
    return
}

func NewDescribeLivePadStreamListResponse() (response *DescribeLivePadStreamListResponse) {
    response = &DescribeLivePadStreamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLivePadStreamList
// 使用该接口查询垫片流列表。垫片流状态更新存在一定延迟，可间隔30秒以上查询，避免频繁查询该接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLivePadStreamList(request *DescribeLivePadStreamListRequest) (response *DescribeLivePadStreamListResponse, err error) {
    return c.DescribeLivePadStreamListWithContext(context.Background(), request)
}

// DescribeLivePadStreamList
// 使用该接口查询垫片流列表。垫片流状态更新存在一定延迟，可间隔30秒以上查询，避免频繁查询该接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLivePadStreamListWithContext(ctx context.Context, request *DescribeLivePadStreamListRequest) (response *DescribeLivePadStreamListResponse, err error) {
    if request == nil {
        request = NewDescribeLivePadStreamListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLivePadStreamList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePadStreamList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePadStreamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePadTemplateRequest() (request *DescribeLivePadTemplateRequest) {
    request = &DescribeLivePadTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePadTemplate")
    
    
    return
}

func NewDescribeLivePadTemplateResponse() (response *DescribeLivePadTemplateResponse) {
    response = &DescribeLivePadTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLivePadTemplate
// 获取单个直播垫片模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePadTemplate(request *DescribeLivePadTemplateRequest) (response *DescribeLivePadTemplateResponse, err error) {
    return c.DescribeLivePadTemplateWithContext(context.Background(), request)
}

// DescribeLivePadTemplate
// 获取单个直播垫片模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePadTemplateWithContext(ctx context.Context, request *DescribeLivePadTemplateRequest) (response *DescribeLivePadTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLivePadTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLivePadTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePadTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePadTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePadTemplatesRequest() (request *DescribeLivePadTemplatesRequest) {
    request = &DescribeLivePadTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePadTemplates")
    
    
    return
}

func NewDescribeLivePadTemplatesResponse() (response *DescribeLivePadTemplatesResponse) {
    response = &DescribeLivePadTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLivePadTemplates
// 获取直播垫片模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLivePadTemplates(request *DescribeLivePadTemplatesRequest) (response *DescribeLivePadTemplatesResponse, err error) {
    return c.DescribeLivePadTemplatesWithContext(context.Background(), request)
}

// DescribeLivePadTemplates
// 获取直播垫片模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLivePadTemplatesWithContext(ctx context.Context, request *DescribeLivePadTemplatesRequest) (response *DescribeLivePadTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLivePadTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLivePadTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePadTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePadTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePlayAuthKeyRequest() (request *DescribeLivePlayAuthKeyRequest) {
    request = &DescribeLivePlayAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePlayAuthKey")
    
    
    return
}

func NewDescribeLivePlayAuthKeyResponse() (response *DescribeLivePlayAuthKeyResponse) {
    response = &DescribeLivePlayAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLivePlayAuthKey
// 查询播放鉴权key。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePlayAuthKey(request *DescribeLivePlayAuthKeyRequest) (response *DescribeLivePlayAuthKeyResponse, err error) {
    return c.DescribeLivePlayAuthKeyWithContext(context.Background(), request)
}

// DescribeLivePlayAuthKey
// 查询播放鉴权key。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePlayAuthKeyWithContext(ctx context.Context, request *DescribeLivePlayAuthKeyRequest) (response *DescribeLivePlayAuthKeyResponse, err error) {
    if request == nil {
        request = NewDescribeLivePlayAuthKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLivePlayAuthKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePlayAuthKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePlayAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePullStreamTaskStatusRequest() (request *DescribeLivePullStreamTaskStatusRequest) {
    request = &DescribeLivePullStreamTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePullStreamTaskStatus")
    
    
    return
}

func NewDescribeLivePullStreamTaskStatusResponse() (response *DescribeLivePullStreamTaskStatusResponse) {
    response = &DescribeLivePullStreamTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLivePullStreamTaskStatus
// 查询直播拉流任务状态信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLivePullStreamTaskStatus(request *DescribeLivePullStreamTaskStatusRequest) (response *DescribeLivePullStreamTaskStatusResponse, err error) {
    return c.DescribeLivePullStreamTaskStatusWithContext(context.Background(), request)
}

// DescribeLivePullStreamTaskStatus
// 查询直播拉流任务状态信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLivePullStreamTaskStatusWithContext(ctx context.Context, request *DescribeLivePullStreamTaskStatusRequest) (response *DescribeLivePullStreamTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeLivePullStreamTaskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLivePullStreamTaskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePullStreamTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePullStreamTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePullStreamTasksRequest() (request *DescribeLivePullStreamTasksRequest) {
    request = &DescribeLivePullStreamTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePullStreamTasks")
    
    
    return
}

func NewDescribeLivePullStreamTasksResponse() (response *DescribeLivePullStreamTasksResponse) {
    response = &DescribeLivePullStreamTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLivePullStreamTasks
// 查询使用 CreateLivePullStreamTask 接口创建的直播拉流任务。
//
// 排序方式：默认按更新时间 倒序排列。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLivePullStreamTasks(request *DescribeLivePullStreamTasksRequest) (response *DescribeLivePullStreamTasksResponse, err error) {
    return c.DescribeLivePullStreamTasksWithContext(context.Background(), request)
}

// DescribeLivePullStreamTasks
// 查询使用 CreateLivePullStreamTask 接口创建的直播拉流任务。
//
// 排序方式：默认按更新时间 倒序排列。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLivePullStreamTasksWithContext(ctx context.Context, request *DescribeLivePullStreamTasksRequest) (response *DescribeLivePullStreamTasksResponse, err error) {
    if request == nil {
        request = NewDescribeLivePullStreamTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLivePullStreamTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePullStreamTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePullStreamTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePushAuthKeyRequest() (request *DescribeLivePushAuthKeyRequest) {
    request = &DescribeLivePushAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePushAuthKey")
    
    
    return
}

func NewDescribeLivePushAuthKeyResponse() (response *DescribeLivePushAuthKeyResponse) {
    response = &DescribeLivePushAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLivePushAuthKey
// 查询直播推流鉴权key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PUSHDOMAINNORECORD = "InternalError.PushDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePushAuthKey(request *DescribeLivePushAuthKeyRequest) (response *DescribeLivePushAuthKeyResponse, err error) {
    return c.DescribeLivePushAuthKeyWithContext(context.Background(), request)
}

// DescribeLivePushAuthKey
// 查询直播推流鉴权key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PUSHDOMAINNORECORD = "InternalError.PushDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePushAuthKeyWithContext(ctx context.Context, request *DescribeLivePushAuthKeyRequest) (response *DescribeLivePushAuthKeyResponse, err error) {
    if request == nil {
        request = NewDescribeLivePushAuthKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLivePushAuthKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePushAuthKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePushAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordRulesRequest() (request *DescribeLiveRecordRulesRequest) {
    request = &DescribeLiveRecordRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveRecordRules")
    
    
    return
}

func NewDescribeLiveRecordRulesResponse() (response *DescribeLiveRecordRulesResponse) {
    response = &DescribeLiveRecordRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveRecordRules
// 获取录制规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordRules(request *DescribeLiveRecordRulesRequest) (response *DescribeLiveRecordRulesResponse, err error) {
    return c.DescribeLiveRecordRulesWithContext(context.Background(), request)
}

// DescribeLiveRecordRules
// 获取录制规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordRulesWithContext(ctx context.Context, request *DescribeLiveRecordRulesRequest) (response *DescribeLiveRecordRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveRecordRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveRecordRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveRecordRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordTemplateRequest() (request *DescribeLiveRecordTemplateRequest) {
    request = &DescribeLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveRecordTemplate")
    
    
    return
}

func NewDescribeLiveRecordTemplateResponse() (response *DescribeLiveRecordTemplateResponse) {
    response = &DescribeLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveRecordTemplate
// 获取单个录制模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordTemplate(request *DescribeLiveRecordTemplateRequest) (response *DescribeLiveRecordTemplateResponse, err error) {
    return c.DescribeLiveRecordTemplateWithContext(context.Background(), request)
}

// DescribeLiveRecordTemplate
// 获取单个录制模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordTemplateWithContext(ctx context.Context, request *DescribeLiveRecordTemplateRequest) (response *DescribeLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordTemplatesRequest() (request *DescribeLiveRecordTemplatesRequest) {
    request = &DescribeLiveRecordTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveRecordTemplates")
    
    
    return
}

func NewDescribeLiveRecordTemplatesResponse() (response *DescribeLiveRecordTemplatesResponse) {
    response = &DescribeLiveRecordTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveRecordTemplates
// 获取录制模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordTemplates(request *DescribeLiveRecordTemplatesRequest) (response *DescribeLiveRecordTemplatesResponse, err error) {
    return c.DescribeLiveRecordTemplatesWithContext(context.Background(), request)
}

// DescribeLiveRecordTemplates
// 获取录制模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordTemplatesWithContext(ctx context.Context, request *DescribeLiveRecordTemplatesRequest) (response *DescribeLiveRecordTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveRecordTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveRecordTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveRecordTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveSnapshotRulesRequest() (request *DescribeLiveSnapshotRulesRequest) {
    request = &DescribeLiveSnapshotRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveSnapshotRules")
    
    
    return
}

func NewDescribeLiveSnapshotRulesResponse() (response *DescribeLiveSnapshotRulesResponse) {
    response = &DescribeLiveSnapshotRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveSnapshotRules
// 获取截图规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveSnapshotRules(request *DescribeLiveSnapshotRulesRequest) (response *DescribeLiveSnapshotRulesResponse, err error) {
    return c.DescribeLiveSnapshotRulesWithContext(context.Background(), request)
}

// DescribeLiveSnapshotRules
// 获取截图规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveSnapshotRulesWithContext(ctx context.Context, request *DescribeLiveSnapshotRulesRequest) (response *DescribeLiveSnapshotRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveSnapshotRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveSnapshotRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveSnapshotRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveSnapshotTemplateRequest() (request *DescribeLiveSnapshotTemplateRequest) {
    request = &DescribeLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveSnapshotTemplate")
    
    
    return
}

func NewDescribeLiveSnapshotTemplateResponse() (response *DescribeLiveSnapshotTemplateResponse) {
    response = &DescribeLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveSnapshotTemplate
// 获取单个截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveSnapshotTemplate(request *DescribeLiveSnapshotTemplateRequest) (response *DescribeLiveSnapshotTemplateResponse, err error) {
    return c.DescribeLiveSnapshotTemplateWithContext(context.Background(), request)
}

// DescribeLiveSnapshotTemplate
// 获取单个截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveSnapshotTemplateWithContext(ctx context.Context, request *DescribeLiveSnapshotTemplateRequest) (response *DescribeLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveSnapshotTemplatesRequest() (request *DescribeLiveSnapshotTemplatesRequest) {
    request = &DescribeLiveSnapshotTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveSnapshotTemplates")
    
    
    return
}

func NewDescribeLiveSnapshotTemplatesResponse() (response *DescribeLiveSnapshotTemplatesResponse) {
    response = &DescribeLiveSnapshotTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveSnapshotTemplates
// 获取截图模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveSnapshotTemplates(request *DescribeLiveSnapshotTemplatesRequest) (response *DescribeLiveSnapshotTemplatesResponse, err error) {
    return c.DescribeLiveSnapshotTemplatesWithContext(context.Background(), request)
}

// DescribeLiveSnapshotTemplates
// 获取截图模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveSnapshotTemplatesWithContext(ctx context.Context, request *DescribeLiveSnapshotTemplatesRequest) (response *DescribeLiveSnapshotTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveSnapshotTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveSnapshotTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveSnapshotTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamEventListRequest() (request *DescribeLiveStreamEventListRequest) {
    request = &DescribeLiveStreamEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamEventList")
    
    
    return
}

func NewDescribeLiveStreamEventListResponse() (response *DescribeLiveStreamEventListResponse) {
    response = &DescribeLiveStreamEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveStreamEventList
// 用于查询推断流事件。<br>
//
// 
//
// 注意：
//
// 1. 该接口提供离线推断流记录查询功能，不可作为重要业务场景强依赖接口。
//
// 2. 该接口可通过使用IsFilter进行过滤，返回推流历史记录。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamEventList(request *DescribeLiveStreamEventListRequest) (response *DescribeLiveStreamEventListResponse, err error) {
    return c.DescribeLiveStreamEventListWithContext(context.Background(), request)
}

// DescribeLiveStreamEventList
// 用于查询推断流事件。<br>
//
// 
//
// 注意：
//
// 1. 该接口提供离线推断流记录查询功能，不可作为重要业务场景强依赖接口。
//
// 2. 该接口可通过使用IsFilter进行过滤，返回推流历史记录。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamEventListWithContext(ctx context.Context, request *DescribeLiveStreamEventListRequest) (response *DescribeLiveStreamEventListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamEventListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveStreamEventList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveStreamEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamMonitorRequest() (request *DescribeLiveStreamMonitorRequest) {
    request = &DescribeLiveStreamMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamMonitor")
    
    
    return
}

func NewDescribeLiveStreamMonitorResponse() (response *DescribeLiveStreamMonitorResponse) {
    response = &DescribeLiveStreamMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveStreamMonitor
// 该接口用来查询某个特定监播任务的配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveStreamMonitor(request *DescribeLiveStreamMonitorRequest) (response *DescribeLiveStreamMonitorResponse, err error) {
    return c.DescribeLiveStreamMonitorWithContext(context.Background(), request)
}

// DescribeLiveStreamMonitor
// 该接口用来查询某个特定监播任务的配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveStreamMonitorWithContext(ctx context.Context, request *DescribeLiveStreamMonitorRequest) (response *DescribeLiveStreamMonitorResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveStreamMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveStreamMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamMonitorListRequest() (request *DescribeLiveStreamMonitorListRequest) {
    request = &DescribeLiveStreamMonitorListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamMonitorList")
    
    
    return
}

func NewDescribeLiveStreamMonitorListResponse() (response *DescribeLiveStreamMonitorListResponse) {
    response = &DescribeLiveStreamMonitorListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveStreamMonitorList
// 该接口用来查询直播流监播任务配置的列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveStreamMonitorList(request *DescribeLiveStreamMonitorListRequest) (response *DescribeLiveStreamMonitorListResponse, err error) {
    return c.DescribeLiveStreamMonitorListWithContext(context.Background(), request)
}

// DescribeLiveStreamMonitorList
// 该接口用来查询直播流监播任务配置的列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveStreamMonitorListWithContext(ctx context.Context, request *DescribeLiveStreamMonitorListRequest) (response *DescribeLiveStreamMonitorListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamMonitorListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveStreamMonitorList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamMonitorList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveStreamMonitorListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamOnlineListRequest() (request *DescribeLiveStreamOnlineListRequest) {
    request = &DescribeLiveStreamOnlineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamOnlineList")
    
    
    return
}

func NewDescribeLiveStreamOnlineListResponse() (response *DescribeLiveStreamOnlineListResponse) {
    response = &DescribeLiveStreamOnlineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveStreamOnlineList
// 返回正在直播中的流列表。适用于推流成功后查询在线流信息。
//
// 
//
// 注意：
//
// 1. 该接口仅提供辅助查询在线流列表功能，业务重要场景不可强依赖该接口。
//
// 2. 该接口仅适用于流数少于2万路的情况，对于流数较大用户请联系售后。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamOnlineList(request *DescribeLiveStreamOnlineListRequest) (response *DescribeLiveStreamOnlineListResponse, err error) {
    return c.DescribeLiveStreamOnlineListWithContext(context.Background(), request)
}

// DescribeLiveStreamOnlineList
// 返回正在直播中的流列表。适用于推流成功后查询在线流信息。
//
// 
//
// 注意：
//
// 1. 该接口仅提供辅助查询在线流列表功能，业务重要场景不可强依赖该接口。
//
// 2. 该接口仅适用于流数少于2万路的情况，对于流数较大用户请联系售后。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamOnlineListWithContext(ctx context.Context, request *DescribeLiveStreamOnlineListRequest) (response *DescribeLiveStreamOnlineListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamOnlineListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveStreamOnlineList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamOnlineList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveStreamOnlineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamPublishedListRequest() (request *DescribeLiveStreamPublishedListRequest) {
    request = &DescribeLiveStreamPublishedListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamPublishedList")
    
    
    return
}

func NewDescribeLiveStreamPublishedListResponse() (response *DescribeLiveStreamPublishedListResponse) {
    response = &DescribeLiveStreamPublishedListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveStreamPublishedList
// 返回已经推过流的流列表。<br>
//
// 注意：分页最多支持查询1万条记录，可通过调整查询时间范围来获取更多数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamPublishedList(request *DescribeLiveStreamPublishedListRequest) (response *DescribeLiveStreamPublishedListResponse, err error) {
    return c.DescribeLiveStreamPublishedListWithContext(context.Background(), request)
}

// DescribeLiveStreamPublishedList
// 返回已经推过流的流列表。<br>
//
// 注意：分页最多支持查询1万条记录，可通过调整查询时间范围来获取更多数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamPublishedListWithContext(ctx context.Context, request *DescribeLiveStreamPublishedListRequest) (response *DescribeLiveStreamPublishedListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamPublishedListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveStreamPublishedList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamPublishedList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveStreamPublishedListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamPushInfoListRequest() (request *DescribeLiveStreamPushInfoListRequest) {
    request = &DescribeLiveStreamPushInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamPushInfoList")
    
    
    return
}

func NewDescribeLiveStreamPushInfoListResponse() (response *DescribeLiveStreamPushInfoListResponse) {
    response = &DescribeLiveStreamPushInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveStreamPushInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询所有实时流的推流信息，包括客户端IP，服务端IP，帧率，码率，域名，开始推流时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETDATAFAILED = "FailedOperation.GetDataFailed"
//  FAILEDOPERATION_HASNOTLIVINGSTREAM = "FailedOperation.HasNotLivingStream"
//  FAILEDOPERATION_QUERYUPLOADINFOFAILED = "FailedOperation.QueryUploadInfoFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDREQUEST = "InternalError.InvalidRequest"
//  INTERNALERROR_QUERYPROISPPLAYINFOERROR = "InternalError.QueryProIspPlayInfoError"
//  INTERNALERROR_QUERYUPLOADINFOFAILED = "InternalError.QueryUploadInfoFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveStreamPushInfoList(request *DescribeLiveStreamPushInfoListRequest) (response *DescribeLiveStreamPushInfoListResponse, err error) {
    return c.DescribeLiveStreamPushInfoListWithContext(context.Background(), request)
}

// DescribeLiveStreamPushInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询所有实时流的推流信息，包括客户端IP，服务端IP，帧率，码率，域名，开始推流时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETDATAFAILED = "FailedOperation.GetDataFailed"
//  FAILEDOPERATION_HASNOTLIVINGSTREAM = "FailedOperation.HasNotLivingStream"
//  FAILEDOPERATION_QUERYUPLOADINFOFAILED = "FailedOperation.QueryUploadInfoFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDREQUEST = "InternalError.InvalidRequest"
//  INTERNALERROR_QUERYPROISPPLAYINFOERROR = "InternalError.QueryProIspPlayInfoError"
//  INTERNALERROR_QUERYUPLOADINFOFAILED = "InternalError.QueryUploadInfoFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveStreamPushInfoListWithContext(ctx context.Context, request *DescribeLiveStreamPushInfoListRequest) (response *DescribeLiveStreamPushInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamPushInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveStreamPushInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamPushInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveStreamPushInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamStateRequest() (request *DescribeLiveStreamStateRequest) {
    request = &DescribeLiveStreamStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamState")
    
    
    return
}

func NewDescribeLiveStreamStateResponse() (response *DescribeLiveStreamStateResponse) {
    response = &DescribeLiveStreamStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveStreamState
// 返回直播中、无推流或者禁播等状态。
//
// 
//
// 使用建议：
//
// 该接口提供实时流状态查询功能，鉴于网络抖动等一些不可抗因素，使用该接口作为判断主播是否开播等重要业务场景时，请参考以下使用建议。
//
// 1. 优先使用业务自身的房间开关播逻辑，判断主播是否在线，譬如客户端开播信令和主播在线心跳等。
//
// 2. 对于没有房间管理的直播场景，可以结合以下方案综合判断。
//
// 2.1 根据[推断流事件通知](/document/product/267/20388) 判断主播在线状态。
//
// 2.2 通过定时（间隔>1min）查询[直播中的流接口](/document/api/267/20472)，判断主播是否在线。
//
// 2.3 通过 本接口 查询直播流状态，判断主播是否在线。
//
// 2.4 以上任一方式判断为在线，都认为主播开播中，并且接口查询超时或解析异常时，也默认为在线，减少对业务的影响。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamState(request *DescribeLiveStreamStateRequest) (response *DescribeLiveStreamStateResponse, err error) {
    return c.DescribeLiveStreamStateWithContext(context.Background(), request)
}

// DescribeLiveStreamState
// 返回直播中、无推流或者禁播等状态。
//
// 
//
// 使用建议：
//
// 该接口提供实时流状态查询功能，鉴于网络抖动等一些不可抗因素，使用该接口作为判断主播是否开播等重要业务场景时，请参考以下使用建议。
//
// 1. 优先使用业务自身的房间开关播逻辑，判断主播是否在线，譬如客户端开播信令和主播在线心跳等。
//
// 2. 对于没有房间管理的直播场景，可以结合以下方案综合判断。
//
// 2.1 根据[推断流事件通知](/document/product/267/20388) 判断主播在线状态。
//
// 2.2 通过定时（间隔>1min）查询[直播中的流接口](/document/api/267/20472)，判断主播是否在线。
//
// 2.3 通过 本接口 查询直播流状态，判断主播是否在线。
//
// 2.4 以上任一方式判断为在线，都认为主播开播中，并且接口查询超时或解析异常时，也默认为在线，减少对业务的影响。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamStateWithContext(ctx context.Context, request *DescribeLiveStreamStateRequest) (response *DescribeLiveStreamStateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamStateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveStreamState")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveStreamStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTimeShiftBillInfoListRequest() (request *DescribeLiveTimeShiftBillInfoListRequest) {
    request = &DescribeLiveTimeShiftBillInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTimeShiftBillInfoList")
    
    
    return
}

func NewDescribeLiveTimeShiftBillInfoListResponse() (response *DescribeLiveTimeShiftBillInfoListResponse) {
    response = &DescribeLiveTimeShiftBillInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTimeShiftBillInfoList
// 提供给客户对账，按天统计，维度：推流域名、时移文件时长（累加）、配置天数（不累加）、时移总时长（累加）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTimeShiftBillInfoList(request *DescribeLiveTimeShiftBillInfoListRequest) (response *DescribeLiveTimeShiftBillInfoListResponse, err error) {
    return c.DescribeLiveTimeShiftBillInfoListWithContext(context.Background(), request)
}

// DescribeLiveTimeShiftBillInfoList
// 提供给客户对账，按天统计，维度：推流域名、时移文件时长（累加）、配置天数（不累加）、时移总时长（累加）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTimeShiftBillInfoListWithContext(ctx context.Context, request *DescribeLiveTimeShiftBillInfoListRequest) (response *DescribeLiveTimeShiftBillInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTimeShiftBillInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveTimeShiftBillInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTimeShiftBillInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTimeShiftBillInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTimeShiftRulesRequest() (request *DescribeLiveTimeShiftRulesRequest) {
    request = &DescribeLiveTimeShiftRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTimeShiftRules")
    
    
    return
}

func NewDescribeLiveTimeShiftRulesResponse() (response *DescribeLiveTimeShiftRulesResponse) {
    response = &DescribeLiveTimeShiftRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTimeShiftRules
// 获取直播时移规则列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTimeShiftRules(request *DescribeLiveTimeShiftRulesRequest) (response *DescribeLiveTimeShiftRulesResponse, err error) {
    return c.DescribeLiveTimeShiftRulesWithContext(context.Background(), request)
}

// DescribeLiveTimeShiftRules
// 获取直播时移规则列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTimeShiftRulesWithContext(ctx context.Context, request *DescribeLiveTimeShiftRulesRequest) (response *DescribeLiveTimeShiftRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTimeShiftRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveTimeShiftRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTimeShiftRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTimeShiftRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTimeShiftTemplatesRequest() (request *DescribeLiveTimeShiftTemplatesRequest) {
    request = &DescribeLiveTimeShiftTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTimeShiftTemplates")
    
    
    return
}

func NewDescribeLiveTimeShiftTemplatesResponse() (response *DescribeLiveTimeShiftTemplatesResponse) {
    response = &DescribeLiveTimeShiftTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTimeShiftTemplates
// 获取直播时移模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTimeShiftTemplates(request *DescribeLiveTimeShiftTemplatesRequest) (response *DescribeLiveTimeShiftTemplatesResponse, err error) {
    return c.DescribeLiveTimeShiftTemplatesWithContext(context.Background(), request)
}

// DescribeLiveTimeShiftTemplates
// 获取直播时移模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTimeShiftTemplatesWithContext(ctx context.Context, request *DescribeLiveTimeShiftTemplatesRequest) (response *DescribeLiveTimeShiftTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTimeShiftTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveTimeShiftTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTimeShiftTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTimeShiftTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTimeShiftWriteSizeInfoListRequest() (request *DescribeLiveTimeShiftWriteSizeInfoListRequest) {
    request = &DescribeLiveTimeShiftWriteSizeInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTimeShiftWriteSizeInfoList")
    
    
    return
}

func NewDescribeLiveTimeShiftWriteSizeInfoListResponse() (response *DescribeLiveTimeShiftWriteSizeInfoListResponse) {
    response = &DescribeLiveTimeShiftWriteSizeInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTimeShiftWriteSizeInfoList
// 支持直播时移写入量数据查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTimeShiftWriteSizeInfoList(request *DescribeLiveTimeShiftWriteSizeInfoListRequest) (response *DescribeLiveTimeShiftWriteSizeInfoListResponse, err error) {
    return c.DescribeLiveTimeShiftWriteSizeInfoListWithContext(context.Background(), request)
}

// DescribeLiveTimeShiftWriteSizeInfoList
// 支持直播时移写入量数据查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTimeShiftWriteSizeInfoListWithContext(ctx context.Context, request *DescribeLiveTimeShiftWriteSizeInfoListRequest) (response *DescribeLiveTimeShiftWriteSizeInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTimeShiftWriteSizeInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveTimeShiftWriteSizeInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTimeShiftWriteSizeInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTimeShiftWriteSizeInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeDetailInfoRequest() (request *DescribeLiveTranscodeDetailInfoRequest) {
    request = &DescribeLiveTranscodeDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeDetailInfo")
    
    
    return
}

func NewDescribeLiveTranscodeDetailInfoResponse() (response *DescribeLiveTranscodeDetailInfoResponse) {
    response = &DescribeLiveTranscodeDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTranscodeDetailInfo
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 支持查询某天或某段时间的转码详细信息。由于转码数据量较大，如果查询时间跨度太长可能会拉不到数据，可以尝试将查询时间范围缩小些再重试。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTranscodeDetailInfo(request *DescribeLiveTranscodeDetailInfoRequest) (response *DescribeLiveTranscodeDetailInfoResponse, err error) {
    return c.DescribeLiveTranscodeDetailInfoWithContext(context.Background(), request)
}

// DescribeLiveTranscodeDetailInfo
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 支持查询某天或某段时间的转码详细信息。由于转码数据量较大，如果查询时间跨度太长可能会拉不到数据，可以尝试将查询时间范围缩小些再重试。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTranscodeDetailInfoWithContext(ctx context.Context, request *DescribeLiveTranscodeDetailInfoRequest) (response *DescribeLiveTranscodeDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeDetailInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveTranscodeDetailInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTranscodeDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTranscodeDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeRulesRequest() (request *DescribeLiveTranscodeRulesRequest) {
    request = &DescribeLiveTranscodeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeRules")
    
    
    return
}

func NewDescribeLiveTranscodeRulesResponse() (response *DescribeLiveTranscodeRulesResponse) {
    response = &DescribeLiveTranscodeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTranscodeRules
// 获取转码规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTranscodeRules(request *DescribeLiveTranscodeRulesRequest) (response *DescribeLiveTranscodeRulesResponse, err error) {
    return c.DescribeLiveTranscodeRulesWithContext(context.Background(), request)
}

// DescribeLiveTranscodeRules
// 获取转码规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTranscodeRulesWithContext(ctx context.Context, request *DescribeLiveTranscodeRulesRequest) (response *DescribeLiveTranscodeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveTranscodeRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTranscodeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTranscodeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeTemplateRequest() (request *DescribeLiveTranscodeTemplateRequest) {
    request = &DescribeLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeTemplate")
    
    
    return
}

func NewDescribeLiveTranscodeTemplateResponse() (response *DescribeLiveTranscodeTemplateResponse) {
    response = &DescribeLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTranscodeTemplate
// 获取单个转码模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTranscodeTemplate(request *DescribeLiveTranscodeTemplateRequest) (response *DescribeLiveTranscodeTemplateResponse, err error) {
    return c.DescribeLiveTranscodeTemplateWithContext(context.Background(), request)
}

// DescribeLiveTranscodeTemplate
// 获取单个转码模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTranscodeTemplateWithContext(ctx context.Context, request *DescribeLiveTranscodeTemplateRequest) (response *DescribeLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeTemplatesRequest() (request *DescribeLiveTranscodeTemplatesRequest) {
    request = &DescribeLiveTranscodeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeTemplates")
    
    
    return
}

func NewDescribeLiveTranscodeTemplatesResponse() (response *DescribeLiveTranscodeTemplatesResponse) {
    response = &DescribeLiveTranscodeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTranscodeTemplates
// 获取转码模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTranscodeTemplates(request *DescribeLiveTranscodeTemplatesRequest) (response *DescribeLiveTranscodeTemplatesResponse, err error) {
    return c.DescribeLiveTranscodeTemplatesWithContext(context.Background(), request)
}

// DescribeLiveTranscodeTemplates
// 获取转码模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTranscodeTemplatesWithContext(ctx context.Context, request *DescribeLiveTranscodeTemplatesRequest) (response *DescribeLiveTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveTranscodeTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTranscodeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeTotalInfoRequest() (request *DescribeLiveTranscodeTotalInfoRequest) {
    request = &DescribeLiveTranscodeTotalInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeTotalInfo")
    
    
    return
}

func NewDescribeLiveTranscodeTotalInfoResponse() (response *DescribeLiveTranscodeTotalInfoResponse) {
    response = &DescribeLiveTranscodeTotalInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTranscodeTotalInfo
// 查询转码总量数据，可查询近三个月内的数据。
//
// 注意：
//
// 如果是查询某一天内，则返回5分钟粒度数据；
//
// 如果是查询跨天或指定域名， 则返回1小时粒度数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLiveTranscodeTotalInfo(request *DescribeLiveTranscodeTotalInfoRequest) (response *DescribeLiveTranscodeTotalInfoResponse, err error) {
    return c.DescribeLiveTranscodeTotalInfoWithContext(context.Background(), request)
}

// DescribeLiveTranscodeTotalInfo
// 查询转码总量数据，可查询近三个月内的数据。
//
// 注意：
//
// 如果是查询某一天内，则返回5分钟粒度数据；
//
// 如果是查询跨天或指定域名， 则返回1小时粒度数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLiveTranscodeTotalInfoWithContext(ctx context.Context, request *DescribeLiveTranscodeTotalInfoRequest) (response *DescribeLiveTranscodeTotalInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeTotalInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveTranscodeTotalInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTranscodeTotalInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTranscodeTotalInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveWatermarkRequest() (request *DescribeLiveWatermarkRequest) {
    request = &DescribeLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveWatermark")
    
    
    return
}

func NewDescribeLiveWatermarkResponse() (response *DescribeLiveWatermarkResponse) {
    response = &DescribeLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveWatermark
// 获取单个水印信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermark(request *DescribeLiveWatermarkRequest) (response *DescribeLiveWatermarkResponse, err error) {
    return c.DescribeLiveWatermarkWithContext(context.Background(), request)
}

// DescribeLiveWatermark
// 获取单个水印信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermarkWithContext(ctx context.Context, request *DescribeLiveWatermarkRequest) (response *DescribeLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveWatermark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveWatermarkRulesRequest() (request *DescribeLiveWatermarkRulesRequest) {
    request = &DescribeLiveWatermarkRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveWatermarkRules")
    
    
    return
}

func NewDescribeLiveWatermarkRulesResponse() (response *DescribeLiveWatermarkRulesResponse) {
    response = &DescribeLiveWatermarkRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveWatermarkRules
// 获取水印规则列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermarkRules(request *DescribeLiveWatermarkRulesRequest) (response *DescribeLiveWatermarkRulesResponse, err error) {
    return c.DescribeLiveWatermarkRulesWithContext(context.Background(), request)
}

// DescribeLiveWatermarkRules
// 获取水印规则列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermarkRulesWithContext(ctx context.Context, request *DescribeLiveWatermarkRulesRequest) (response *DescribeLiveWatermarkRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarkRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveWatermarkRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveWatermarkRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveWatermarkRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveWatermarksRequest() (request *DescribeLiveWatermarksRequest) {
    request = &DescribeLiveWatermarksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveWatermarks")
    
    
    return
}

func NewDescribeLiveWatermarksResponse() (response *DescribeLiveWatermarksResponse) {
    response = &DescribeLiveWatermarksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveWatermarks
// 查询水印列表。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermarks(request *DescribeLiveWatermarksRequest) (response *DescribeLiveWatermarksResponse, err error) {
    return c.DescribeLiveWatermarksWithContext(context.Background(), request)
}

// DescribeLiveWatermarks
// 查询水印列表。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermarksWithContext(ctx context.Context, request *DescribeLiveWatermarksRequest) (response *DescribeLiveWatermarksResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveWatermarks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveWatermarks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveWatermarksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveXP2PDetailInfoListRequest() (request *DescribeLiveXP2PDetailInfoListRequest) {
    request = &DescribeLiveXP2PDetailInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveXP2PDetailInfoList")
    
    
    return
}

func NewDescribeLiveXP2PDetailInfoListResponse() (response *DescribeLiveXP2PDetailInfoListResponse) {
    response = &DescribeLiveXP2PDetailInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveXP2PDetailInfoList
// P2P流数据查询接口，用来获取流量、卡播和起播信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveXP2PDetailInfoList(request *DescribeLiveXP2PDetailInfoListRequest) (response *DescribeLiveXP2PDetailInfoListResponse, err error) {
    return c.DescribeLiveXP2PDetailInfoListWithContext(context.Background(), request)
}

// DescribeLiveXP2PDetailInfoList
// P2P流数据查询接口，用来获取流量、卡播和起播信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveXP2PDetailInfoListWithContext(ctx context.Context, request *DescribeLiveXP2PDetailInfoListRequest) (response *DescribeLiveXP2PDetailInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveXP2PDetailInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLiveXP2PDetailInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveXP2PDetailInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveXP2PDetailInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogDownloadListRequest() (request *DescribeLogDownloadListRequest) {
    request = &DescribeLogDownloadListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLogDownloadList")
    
    
    return
}

func NewDescribeLogDownloadListResponse() (response *DescribeLogDownloadListResponse) {
    response = &DescribeLogDownloadListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogDownloadList
// 批量获取日志URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVOKECDNAPIFAIL = "FailedOperation.InvokeCdnApiFail"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CDNLOGEMPTY = "ResourceNotFound.CdnLogEmpty"
//  RESOURCENOTFOUND_CDNTHEMEEMPTY = "ResourceNotFound.CdnThemeEmpty"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLogDownloadList(request *DescribeLogDownloadListRequest) (response *DescribeLogDownloadListResponse, err error) {
    return c.DescribeLogDownloadListWithContext(context.Background(), request)
}

// DescribeLogDownloadList
// 批量获取日志URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVOKECDNAPIFAIL = "FailedOperation.InvokeCdnApiFail"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CDNLOGEMPTY = "ResourceNotFound.CdnLogEmpty"
//  RESOURCENOTFOUND_CDNTHEMEEMPTY = "ResourceNotFound.CdnThemeEmpty"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLogDownloadListWithContext(ctx context.Context, request *DescribeLogDownloadListRequest) (response *DescribeLogDownloadListResponse, err error) {
    if request == nil {
        request = NewDescribeLogDownloadListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeLogDownloadList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogDownloadList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogDownloadListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorReportRequest() (request *DescribeMonitorReportRequest) {
    request = &DescribeMonitorReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeMonitorReport")
    
    
    return
}

func NewDescribeMonitorReportResponse() (response *DescribeMonitorReportResponse) {
    response = &DescribeMonitorReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMonitorReport
// 用来查询监播场次7天内的智能识别、断流、低帧率等信息的汇总报告。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeMonitorReport(request *DescribeMonitorReportRequest) (response *DescribeMonitorReportResponse, err error) {
    return c.DescribeMonitorReportWithContext(context.Background(), request)
}

// DescribeMonitorReport
// 用来查询监播场次7天内的智能识别、断流、低帧率等信息的汇总报告。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeMonitorReportWithContext(ctx context.Context, request *DescribeMonitorReportRequest) (response *DescribeMonitorReportResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorReportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeMonitorReport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMonitorReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMonitorReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlayErrorCodeDetailInfoListRequest() (request *DescribePlayErrorCodeDetailInfoListRequest) {
    request = &DescribePlayErrorCodeDetailInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribePlayErrorCodeDetailInfoList")
    
    
    return
}

func NewDescribePlayErrorCodeDetailInfoListResponse() (response *DescribePlayErrorCodeDetailInfoListResponse) {
    response = &DescribePlayErrorCodeDetailInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePlayErrorCodeDetailInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询下行播放错误码信息，某段时间内1分钟粒度的各http错误码出现的次数，包括4xx，5xx。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePlayErrorCodeDetailInfoList(request *DescribePlayErrorCodeDetailInfoListRequest) (response *DescribePlayErrorCodeDetailInfoListResponse, err error) {
    return c.DescribePlayErrorCodeDetailInfoListWithContext(context.Background(), request)
}

// DescribePlayErrorCodeDetailInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询下行播放错误码信息，某段时间内1分钟粒度的各http错误码出现的次数，包括4xx，5xx。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePlayErrorCodeDetailInfoListWithContext(ctx context.Context, request *DescribePlayErrorCodeDetailInfoListRequest) (response *DescribePlayErrorCodeDetailInfoListResponse, err error) {
    if request == nil {
        request = NewDescribePlayErrorCodeDetailInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribePlayErrorCodeDetailInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlayErrorCodeDetailInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePlayErrorCodeDetailInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlayErrorCodeSumInfoListRequest() (request *DescribePlayErrorCodeSumInfoListRequest) {
    request = &DescribePlayErrorCodeSumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribePlayErrorCodeSumInfoList")
    
    
    return
}

func NewDescribePlayErrorCodeSumInfoListResponse() (response *DescribePlayErrorCodeSumInfoListResponse) {
    response = &DescribePlayErrorCodeSumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePlayErrorCodeSumInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询下行播放错误码信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePlayErrorCodeSumInfoList(request *DescribePlayErrorCodeSumInfoListRequest) (response *DescribePlayErrorCodeSumInfoListResponse, err error) {
    return c.DescribePlayErrorCodeSumInfoListWithContext(context.Background(), request)
}

// DescribePlayErrorCodeSumInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询下行播放错误码信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePlayErrorCodeSumInfoListWithContext(ctx context.Context, request *DescribePlayErrorCodeSumInfoListRequest) (response *DescribePlayErrorCodeSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribePlayErrorCodeSumInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribePlayErrorCodeSumInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlayErrorCodeSumInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePlayErrorCodeSumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProIspPlaySumInfoListRequest() (request *DescribeProIspPlaySumInfoListRequest) {
    request = &DescribeProIspPlaySumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeProIspPlaySumInfoList")
    
    
    return
}

func NewDescribeProIspPlaySumInfoListResponse() (response *DescribeProIspPlaySumInfoListResponse) {
    response = &DescribeProIspPlaySumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProIspPlaySumInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询某段时间内每个国家地区每个省份每个运营商的平均每秒流量，总流量，总请求数信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeProIspPlaySumInfoList(request *DescribeProIspPlaySumInfoListRequest) (response *DescribeProIspPlaySumInfoListResponse, err error) {
    return c.DescribeProIspPlaySumInfoListWithContext(context.Background(), request)
}

// DescribeProIspPlaySumInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询某段时间内每个国家地区每个省份每个运营商的平均每秒流量，总流量，总请求数信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeProIspPlaySumInfoListWithContext(ctx context.Context, request *DescribeProIspPlaySumInfoListRequest) (response *DescribeProIspPlaySumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeProIspPlaySumInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeProIspPlaySumInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProIspPlaySumInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProIspPlaySumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProvinceIspPlayInfoListRequest() (request *DescribeProvinceIspPlayInfoListRequest) {
    request = &DescribeProvinceIspPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeProvinceIspPlayInfoList")
    
    
    return
}

func NewDescribeProvinceIspPlayInfoListResponse() (response *DescribeProvinceIspPlayInfoListResponse) {
    response = &DescribeProvinceIspPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProvinceIspPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询某省份某运营商下行播放数据，包括带宽，流量，请求数，并发连接数信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_HASNOTLIVINGSTREAM = "InternalError.HasNotLivingStream"
//  INTERNALERROR_INVALIDREQUEST = "InternalError.InvalidRequest"
//  INTERNALERROR_QUERYPROISPPLAYINFOERROR = "InternalError.QueryProIspPlayInfoError"
//  INTERNALERROR_QUERYUPLOADINFOFAILED = "InternalError.QueryUploadInfoFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeProvinceIspPlayInfoList(request *DescribeProvinceIspPlayInfoListRequest) (response *DescribeProvinceIspPlayInfoListResponse, err error) {
    return c.DescribeProvinceIspPlayInfoListWithContext(context.Background(), request)
}

// DescribeProvinceIspPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询某省份某运营商下行播放数据，包括带宽，流量，请求数，并发连接数信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_HASNOTLIVINGSTREAM = "InternalError.HasNotLivingStream"
//  INTERNALERROR_INVALIDREQUEST = "InternalError.InvalidRequest"
//  INTERNALERROR_QUERYPROISPPLAYINFOERROR = "InternalError.QueryProIspPlayInfoError"
//  INTERNALERROR_QUERYUPLOADINFOFAILED = "InternalError.QueryUploadInfoFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeProvinceIspPlayInfoListWithContext(ctx context.Context, request *DescribeProvinceIspPlayInfoListRequest) (response *DescribeProvinceIspPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeProvinceIspPlayInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeProvinceIspPlayInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProvinceIspPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProvinceIspPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePullStreamConfigsRequest() (request *DescribePullStreamConfigsRequest) {
    request = &DescribePullStreamConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribePullStreamConfigs")
    
    
    return
}

func NewDescribePullStreamConfigsResponse() (response *DescribePullStreamConfigsResponse) {
    response = &DescribePullStreamConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePullStreamConfigs
// 查询直播拉流配置。该接口已下线,请使用新接口 DescribeLivePullStreamTasks。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribePullStreamConfigs(request *DescribePullStreamConfigsRequest) (response *DescribePullStreamConfigsResponse, err error) {
    return c.DescribePullStreamConfigsWithContext(context.Background(), request)
}

// DescribePullStreamConfigs
// 查询直播拉流配置。该接口已下线,请使用新接口 DescribeLivePullStreamTasks。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribePullStreamConfigsWithContext(ctx context.Context, request *DescribePullStreamConfigsRequest) (response *DescribePullStreamConfigsResponse, err error) {
    if request == nil {
        request = NewDescribePullStreamConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribePullStreamConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePullStreamConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePullStreamConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePullTransformPushInfoRequest() (request *DescribePullTransformPushInfoRequest) {
    request = &DescribePullTransformPushInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribePullTransformPushInfo")
    
    
    return
}

func NewDescribePullTransformPushInfoResponse() (response *DescribePullTransformPushInfoResponse) {
    response = &DescribePullTransformPushInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePullTransformPushInfo
// 查询拉流转推任务的时长信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePullTransformPushInfo(request *DescribePullTransformPushInfoRequest) (response *DescribePullTransformPushInfoResponse, err error) {
    return c.DescribePullTransformPushInfoWithContext(context.Background(), request)
}

// DescribePullTransformPushInfo
// 查询拉流转推任务的时长信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePullTransformPushInfoWithContext(ctx context.Context, request *DescribePullTransformPushInfoRequest) (response *DescribePullTransformPushInfoResponse, err error) {
    if request == nil {
        request = NewDescribePullTransformPushInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribePullTransformPushInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePullTransformPushInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePullTransformPushInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePushBandwidthAndFluxListRequest() (request *DescribePushBandwidthAndFluxListRequest) {
    request = &DescribePushBandwidthAndFluxListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribePushBandwidthAndFluxList")
    
    
    return
}

func NewDescribePushBandwidthAndFluxListResponse() (response *DescribePushBandwidthAndFluxListResponse) {
    response = &DescribePushBandwidthAndFluxListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePushBandwidthAndFluxList
// 直播推流带宽和流量数据查询。
//
// 推流计费会先取全球推流用量和全球播放用量进行比较，满足计费条件后再按各地区用量出账。详情参见[计费文档](https://cloud.tencent.com/document/product/267/34175)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePushBandwidthAndFluxList(request *DescribePushBandwidthAndFluxListRequest) (response *DescribePushBandwidthAndFluxListResponse, err error) {
    return c.DescribePushBandwidthAndFluxListWithContext(context.Background(), request)
}

// DescribePushBandwidthAndFluxList
// 直播推流带宽和流量数据查询。
//
// 推流计费会先取全球推流用量和全球播放用量进行比较，满足计费条件后再按各地区用量出账。详情参见[计费文档](https://cloud.tencent.com/document/product/267/34175)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePushBandwidthAndFluxListWithContext(ctx context.Context, request *DescribePushBandwidthAndFluxListRequest) (response *DescribePushBandwidthAndFluxListResponse, err error) {
    if request == nil {
        request = NewDescribePushBandwidthAndFluxListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribePushBandwidthAndFluxList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePushBandwidthAndFluxList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePushBandwidthAndFluxListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordTaskRequest() (request *DescribeRecordTaskRequest) {
    request = &DescribeRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeRecordTask")
    
    
    return
}

func NewDescribeRecordTaskResponse() (response *DescribeRecordTaskResponse) {
    response = &DescribeRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordTask
// 查询指定时间段范围内启动和结束的录制任务列表。
//
// - 使用前提
//
// 1. 仅用于查询由 CreateRecordTask 接口创建的录制任务。
//
// 2. 不能查询被 DeleteRecordTask 接口删除以及已过期（平台侧保留3个月）的录制任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRecordTask(request *DescribeRecordTaskRequest) (response *DescribeRecordTaskResponse, err error) {
    return c.DescribeRecordTaskWithContext(context.Background(), request)
}

// DescribeRecordTask
// 查询指定时间段范围内启动和结束的录制任务列表。
//
// - 使用前提
//
// 1. 仅用于查询由 CreateRecordTask 接口创建的录制任务。
//
// 2. 不能查询被 DeleteRecordTask 接口删除以及已过期（平台侧保留3个月）的录制任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRecordTaskWithContext(ctx context.Context, request *DescribeRecordTaskRequest) (response *DescribeRecordTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRecordTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeRecordTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenShotSheetNumListRequest() (request *DescribeScreenShotSheetNumListRequest) {
    request = &DescribeScreenShotSheetNumListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeScreenShotSheetNumList")
    
    
    return
}

func NewDescribeScreenShotSheetNumListResponse() (response *DescribeScreenShotSheetNumListResponse) {
    response = &DescribeScreenShotSheetNumListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScreenShotSheetNumList
// 接口用来查询直播增值业务--截图的张数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeScreenShotSheetNumList(request *DescribeScreenShotSheetNumListRequest) (response *DescribeScreenShotSheetNumListResponse, err error) {
    return c.DescribeScreenShotSheetNumListWithContext(context.Background(), request)
}

// DescribeScreenShotSheetNumList
// 接口用来查询直播增值业务--截图的张数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeScreenShotSheetNumListWithContext(ctx context.Context, request *DescribeScreenShotSheetNumListRequest) (response *DescribeScreenShotSheetNumListResponse, err error) {
    if request == nil {
        request = NewDescribeScreenShotSheetNumListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeScreenShotSheetNumList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenShotSheetNumList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenShotSheetNumListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenshotTaskRequest() (request *DescribeScreenshotTaskRequest) {
    request = &DescribeScreenshotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeScreenshotTask")
    
    
    return
}

func NewDescribeScreenshotTaskResponse() (response *DescribeScreenshotTaskResponse) {
    response = &DescribeScreenshotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScreenshotTask
// 查询指定时间段范围内启动和结束的截图任务列表。
//
// - 使用前提
//
// 1. 仅用于查询由 CreateScreenshotTask接口创建的截图任务。
//
// 2. 不能查询被 DeleteScreenshotTask接口删除以及已过期（平台侧保留3个月）的截图任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScreenshotTask(request *DescribeScreenshotTaskRequest) (response *DescribeScreenshotTaskResponse, err error) {
    return c.DescribeScreenshotTaskWithContext(context.Background(), request)
}

// DescribeScreenshotTask
// 查询指定时间段范围内启动和结束的截图任务列表。
//
// - 使用前提
//
// 1. 仅用于查询由 CreateScreenshotTask接口创建的截图任务。
//
// 2. 不能查询被 DeleteScreenshotTask接口删除以及已过期（平台侧保留3个月）的截图任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScreenshotTaskWithContext(ctx context.Context, request *DescribeScreenshotTaskRequest) (response *DescribeScreenshotTaskResponse, err error) {
    if request == nil {
        request = NewDescribeScreenshotTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeScreenshotTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenshotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScreenshotTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamDayPlayInfoListRequest() (request *DescribeStreamDayPlayInfoListRequest) {
    request = &DescribeStreamDayPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeStreamDayPlayInfoList")
    
    
    return
}

func NewDescribeStreamDayPlayInfoListResponse() (response *DescribeStreamDayPlayInfoListResponse) {
    response = &DescribeStreamDayPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamDayPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询天维度每条流的播放数据，包括总流量等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamDayPlayInfoList(request *DescribeStreamDayPlayInfoListRequest) (response *DescribeStreamDayPlayInfoListResponse, err error) {
    return c.DescribeStreamDayPlayInfoListWithContext(context.Background(), request)
}

// DescribeStreamDayPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询天维度每条流的播放数据，包括总流量等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamDayPlayInfoListWithContext(ctx context.Context, request *DescribeStreamDayPlayInfoListRequest) (response *DescribeStreamDayPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamDayPlayInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeStreamDayPlayInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamDayPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamDayPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPlayInfoListRequest() (request *DescribeStreamPlayInfoListRequest) {
    request = &DescribeStreamPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeStreamPlayInfoList")
    
    
    return
}

func NewDescribeStreamPlayInfoListResponse() (response *DescribeStreamPlayInfoListResponse) {
    response = &DescribeStreamPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询播放数据，支持按流名称查询详细播放数据，也可按播放域名查询详细总数据，数据延迟4分钟左右。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamPlayInfoList(request *DescribeStreamPlayInfoListRequest) (response *DescribeStreamPlayInfoListResponse, err error) {
    return c.DescribeStreamPlayInfoListWithContext(context.Background(), request)
}

// DescribeStreamPlayInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询播放数据，支持按流名称查询详细播放数据，也可按播放域名查询详细总数据，数据延迟4分钟左右。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamPlayInfoListWithContext(ctx context.Context, request *DescribeStreamPlayInfoListRequest) (response *DescribeStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPlayInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeStreamPlayInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPushInfoListRequest() (request *DescribeStreamPushInfoListRequest) {
    request = &DescribeStreamPushInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeStreamPushInfoList")
    
    
    return
}

func NewDescribeStreamPushInfoListResponse() (response *DescribeStreamPushInfoListResponse) {
    response = &DescribeStreamPushInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPushInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询流id的上行推流质量数据，包括音视频的帧率，码率，流逝时间，编码格式等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamPushInfoList(request *DescribeStreamPushInfoListRequest) (response *DescribeStreamPushInfoListResponse, err error) {
    return c.DescribeStreamPushInfoListWithContext(context.Background(), request)
}

// DescribeStreamPushInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询流id的上行推流质量数据，包括音视频的帧率，码率，流逝时间，编码格式等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamPushInfoListWithContext(ctx context.Context, request *DescribeStreamPushInfoListRequest) (response *DescribeStreamPushInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPushInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeStreamPushInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPushInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPushInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimeShiftRecordDetailRequest() (request *DescribeTimeShiftRecordDetailRequest) {
    request = &DescribeTimeShiftRecordDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeTimeShiftRecordDetail")
    
    
    return
}

func NewDescribeTimeShiftRecordDetailResponse() (response *DescribeTimeShiftRecordDetailResponse) {
    response = &DescribeTimeShiftRecordDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimeShiftRecordDetail
// 前提调用 DescribeTimeShiftStreamList 获得请求必要参数。查询指定范围内的时移流录制详情，最大支持24小时范围查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTimeShiftRecordDetail(request *DescribeTimeShiftRecordDetailRequest) (response *DescribeTimeShiftRecordDetailResponse, err error) {
    return c.DescribeTimeShiftRecordDetailWithContext(context.Background(), request)
}

// DescribeTimeShiftRecordDetail
// 前提调用 DescribeTimeShiftStreamList 获得请求必要参数。查询指定范围内的时移流录制详情，最大支持24小时范围查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTimeShiftRecordDetailWithContext(ctx context.Context, request *DescribeTimeShiftRecordDetailRequest) (response *DescribeTimeShiftRecordDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTimeShiftRecordDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeTimeShiftRecordDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimeShiftRecordDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimeShiftRecordDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimeShiftStreamListRequest() (request *DescribeTimeShiftStreamListRequest) {
    request = &DescribeTimeShiftStreamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeTimeShiftStreamList")
    
    
    return
}

func NewDescribeTimeShiftStreamListResponse() (response *DescribeTimeShiftStreamListResponse) {
    response = &DescribeTimeShiftStreamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimeShiftStreamList
// 查询某个时间范围内所有时移流列表。最大支持查询24小时内的数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTimeShiftStreamList(request *DescribeTimeShiftStreamListRequest) (response *DescribeTimeShiftStreamListResponse, err error) {
    return c.DescribeTimeShiftStreamListWithContext(context.Background(), request)
}

// DescribeTimeShiftStreamList
// 查询某个时间范围内所有时移流列表。最大支持查询24小时内的数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTimeShiftStreamListWithContext(ctx context.Context, request *DescribeTimeShiftStreamListRequest) (response *DescribeTimeShiftStreamListResponse, err error) {
    if request == nil {
        request = NewDescribeTimeShiftStreamListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeTimeShiftStreamList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimeShiftStreamList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimeShiftStreamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopClientIpSumInfoListRequest() (request *DescribeTopClientIpSumInfoListRequest) {
    request = &DescribeTopClientIpSumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeTopClientIpSumInfoList")
    
    
    return
}

func NewDescribeTopClientIpSumInfoListResponse() (response *DescribeTopClientIpSumInfoListResponse) {
    response = &DescribeTopClientIpSumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopClientIpSumInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询某段时间top n客户端ip汇总信息（暂支持top 1000）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeTopClientIpSumInfoList(request *DescribeTopClientIpSumInfoListRequest) (response *DescribeTopClientIpSumInfoListResponse, err error) {
    return c.DescribeTopClientIpSumInfoListWithContext(context.Background(), request)
}

// DescribeTopClientIpSumInfoList
// 该接口为监控数据接口，数据采集及统计方式与计费数据不同，仅供运营分析使用，不能用于计费对账参考。
//
// 查询某段时间top n客户端ip汇总信息（暂支持top 1000）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeTopClientIpSumInfoListWithContext(ctx context.Context, request *DescribeTopClientIpSumInfoListRequest) (response *DescribeTopClientIpSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeTopClientIpSumInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeTopClientIpSumInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopClientIpSumInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopClientIpSumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeTaskNumRequest() (request *DescribeTranscodeTaskNumRequest) {
    request = &DescribeTranscodeTaskNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeTranscodeTaskNum")
    
    
    return
}

func NewDescribeTranscodeTaskNumResponse() (response *DescribeTranscodeTaskNumResponse) {
    response = &DescribeTranscodeTaskNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTranscodeTaskNum
// 查询转码任务数。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeTranscodeTaskNum(request *DescribeTranscodeTaskNumRequest) (response *DescribeTranscodeTaskNumResponse, err error) {
    return c.DescribeTranscodeTaskNumWithContext(context.Background(), request)
}

// DescribeTranscodeTaskNum
// 查询转码任务数。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeTranscodeTaskNumWithContext(ctx context.Context, request *DescribeTranscodeTaskNumRequest) (response *DescribeTranscodeTaskNumResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeTaskNumRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeTranscodeTaskNum")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeTaskNum require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeTaskNumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadStreamNumsRequest() (request *DescribeUploadStreamNumsRequest) {
    request = &DescribeUploadStreamNumsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeUploadStreamNums")
    
    
    return
}

func NewDescribeUploadStreamNumsResponse() (response *DescribeUploadStreamNumsResponse) {
    response = &DescribeUploadStreamNumsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUploadStreamNums
// 直播上行路数查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeUploadStreamNums(request *DescribeUploadStreamNumsRequest) (response *DescribeUploadStreamNumsResponse, err error) {
    return c.DescribeUploadStreamNumsWithContext(context.Background(), request)
}

// DescribeUploadStreamNums
// 直播上行路数查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeUploadStreamNumsWithContext(ctx context.Context, request *DescribeUploadStreamNumsRequest) (response *DescribeUploadStreamNumsResponse, err error) {
    if request == nil {
        request = NewDescribeUploadStreamNumsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeUploadStreamNums")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUploadStreamNums require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUploadStreamNumsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVisitTopSumInfoListRequest() (request *DescribeVisitTopSumInfoListRequest) {
    request = &DescribeVisitTopSumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeVisitTopSumInfoList")
    
    
    return
}

func NewDescribeVisitTopSumInfoListResponse() (response *DescribeVisitTopSumInfoListResponse) {
    response = &DescribeVisitTopSumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVisitTopSumInfoList
// 查询某时间段top n的域名或流id信息（暂支持top 1000）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeVisitTopSumInfoList(request *DescribeVisitTopSumInfoListRequest) (response *DescribeVisitTopSumInfoListResponse, err error) {
    return c.DescribeVisitTopSumInfoListWithContext(context.Background(), request)
}

// DescribeVisitTopSumInfoList
// 查询某时间段top n的域名或流id信息（暂支持top 1000）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeVisitTopSumInfoListWithContext(ctx context.Context, request *DescribeVisitTopSumInfoListRequest) (response *DescribeVisitTopSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeVisitTopSumInfoListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DescribeVisitTopSumInfoList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVisitTopSumInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVisitTopSumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDropLiveStreamRequest() (request *DropLiveStreamRequest) {
    request = &DropLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DropLiveStream")
    
    
    return
}

func NewDropLiveStreamResponse() (response *DropLiveStreamResponse) {
    response = &DropLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DropLiveStream
// 断开推流连接，但可以重新推流。
//
// 注：对已经不活跃的流，调用该断流接口时，接口返回成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DropLiveStream(request *DropLiveStreamRequest) (response *DropLiveStreamResponse, err error) {
    return c.DropLiveStreamWithContext(context.Background(), request)
}

// DropLiveStream
// 断开推流连接，但可以重新推流。
//
// 注：对已经不活跃的流，调用该断流接口时，接口返回成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DropLiveStreamWithContext(ctx context.Context, request *DropLiveStreamRequest) (response *DropLiveStreamResponse, err error) {
    if request == nil {
        request = NewDropLiveStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "DropLiveStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DropLiveStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewDropLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewEnableLiveDomainRequest() (request *EnableLiveDomainRequest) {
    request = &EnableLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "EnableLiveDomain")
    
    
    return
}

func NewEnableLiveDomainResponse() (response *EnableLiveDomainResponse) {
    response = &EnableLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableLiveDomain
// 启用状态为停用的直播域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SDKNOPACKAGE = "FailedOperation.SdkNoPackage"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CLOUDDOMAINISSTOP = "InvalidParameter.CloudDomainIsStop"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) EnableLiveDomain(request *EnableLiveDomainRequest) (response *EnableLiveDomainResponse, err error) {
    return c.EnableLiveDomainWithContext(context.Background(), request)
}

// EnableLiveDomain
// 启用状态为停用的直播域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SDKNOPACKAGE = "FailedOperation.SdkNoPackage"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CLOUDDOMAINISSTOP = "InvalidParameter.CloudDomainIsStop"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) EnableLiveDomainWithContext(ctx context.Context, request *EnableLiveDomainRequest) (response *EnableLiveDomainResponse, err error) {
    if request == nil {
        request = NewEnableLiveDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "EnableLiveDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableLiveDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewEnableOptimalSwitchingRequest() (request *EnableOptimalSwitchingRequest) {
    request = &EnableOptimalSwitchingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "EnableOptimalSwitching")
    
    
    return
}

func NewEnableOptimalSwitchingResponse() (response *EnableOptimalSwitchingResponse) {
    response = &EnableOptimalSwitchingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableOptimalSwitching
// 启用择优调度。
//
// 注意：流维度的择优调度，当主备流结束后自动失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SDKNOPACKAGE = "FailedOperation.SdkNoPackage"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CLOUDDOMAINISSTOP = "InvalidParameter.CloudDomainIsStop"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) EnableOptimalSwitching(request *EnableOptimalSwitchingRequest) (response *EnableOptimalSwitchingResponse, err error) {
    return c.EnableOptimalSwitchingWithContext(context.Background(), request)
}

// EnableOptimalSwitching
// 启用择优调度。
//
// 注意：流维度的择优调度，当主备流结束后自动失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SDKNOPACKAGE = "FailedOperation.SdkNoPackage"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CLOUDDOMAINISSTOP = "InvalidParameter.CloudDomainIsStop"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) EnableOptimalSwitchingWithContext(ctx context.Context, request *EnableOptimalSwitchingRequest) (response *EnableOptimalSwitchingResponse, err error) {
    if request == nil {
        request = NewEnableOptimalSwitchingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "EnableOptimalSwitching")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableOptimalSwitching require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableOptimalSwitchingResponse()
    err = c.Send(request, response)
    return
}

func NewForbidLiveDomainRequest() (request *ForbidLiveDomainRequest) {
    request = &ForbidLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ForbidLiveDomain")
    
    
    return
}

func NewForbidLiveDomainResponse() (response *ForbidLiveDomainResponse) {
    response = &ForbidLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ForbidLiveDomain
// 停止使用某个直播域名。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ForbidLiveDomain(request *ForbidLiveDomainRequest) (response *ForbidLiveDomainResponse, err error) {
    return c.ForbidLiveDomainWithContext(context.Background(), request)
}

// ForbidLiveDomain
// 停止使用某个直播域名。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ForbidLiveDomainWithContext(ctx context.Context, request *ForbidLiveDomainRequest) (response *ForbidLiveDomainResponse, err error) {
    if request == nil {
        request = NewForbidLiveDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ForbidLiveDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForbidLiveDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewForbidLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewForbidLiveStreamRequest() (request *ForbidLiveStreamRequest) {
    request = &ForbidLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ForbidLiveStream")
    
    
    return
}

func NewForbidLiveStreamResponse() (response *ForbidLiveStreamResponse) {
    response = &ForbidLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ForbidLiveStream
// 禁止某条流的推送，可以预设某个时刻将流恢复。
//
// 注意：
//
// 1. 默认只要流名称正确，禁推就会生效。
//
// 2. 如需要推流域名+推流路径+流名称 强匹配生效禁推，需提单联系售后开启配置。
//
// 3. 如果配置了域名分组，需填写准确推流域名，才可断掉当前推流。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ForbidLiveStream(request *ForbidLiveStreamRequest) (response *ForbidLiveStreamResponse, err error) {
    return c.ForbidLiveStreamWithContext(context.Background(), request)
}

// ForbidLiveStream
// 禁止某条流的推送，可以预设某个时刻将流恢复。
//
// 注意：
//
// 1. 默认只要流名称正确，禁推就会生效。
//
// 2. 如需要推流域名+推流路径+流名称 强匹配生效禁推，需提单联系售后开启配置。
//
// 3. 如果配置了域名分组，需填写准确推流域名，才可断掉当前推流。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ForbidLiveStreamWithContext(ctx context.Context, request *ForbidLiveStreamRequest) (response *ForbidLiveStreamResponse, err error) {
    if request == nil {
        request = NewForbidLiveStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ForbidLiveStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForbidLiveStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewForbidLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCasterRequest() (request *ModifyCasterRequest) {
    request = &ModifyCasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyCaster")
    
    
    return
}

func NewModifyCasterResponse() (response *ModifyCasterResponse) {
    response = &ModifyCasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCaster
// 该接口用来设置导播台的描述、名称、录制模板id等参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCaster(request *ModifyCasterRequest) (response *ModifyCasterResponse, err error) {
    return c.ModifyCasterWithContext(context.Background(), request)
}

// ModifyCaster
// 该接口用来设置导播台的描述、名称、录制模板id等参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterWithContext(ctx context.Context, request *ModifyCasterRequest) (response *ModifyCasterResponse, err error) {
    if request == nil {
        request = NewModifyCasterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyCaster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCasterResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCasterInputInfoRequest() (request *ModifyCasterInputInfoRequest) {
    request = &ModifyCasterInputInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyCasterInputInfo")
    
    
    return
}

func NewModifyCasterInputInfoResponse() (response *ModifyCasterInputInfoResponse) {
    response = &ModifyCasterInputInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCasterInputInfo
// 该接口用来修改已经设置过的输入源信息，如源地址，源类型等。
//
// 设置前，需保证待修改的输入源已经存在。若不存在，需使用AddCasterInputInfo接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_INPUTISNOTACTIVE = "FailedOperation.InputIsNotActive"
//  FAILEDOPERATION_INPUTNOTEXIST = "FailedOperation.InputNotExist"
//  FAILEDOPERATION_INPUTUSEDINAUTOCAST = "FailedOperation.InputUsedInAutoCast"
//  FAILEDOPERATION_INPUTUSEDINPGM = "FailedOperation.InputUsedInPgm"
//  FAILEDOPERATION_INPUTUSEDINPVW = "FailedOperation.InputUsedInPvw"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  FAILEDOPERATION_STARTPULLFAILED = "FailedOperation.StartPullFailed"
//  FAILEDOPERATION_STOPPULLFAILED = "FailedOperation.StopPullFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterInputInfo(request *ModifyCasterInputInfoRequest) (response *ModifyCasterInputInfoResponse, err error) {
    return c.ModifyCasterInputInfoWithContext(context.Background(), request)
}

// ModifyCasterInputInfo
// 该接口用来修改已经设置过的输入源信息，如源地址，源类型等。
//
// 设置前，需保证待修改的输入源已经存在。若不存在，需使用AddCasterInputInfo接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_INPUTISNOTACTIVE = "FailedOperation.InputIsNotActive"
//  FAILEDOPERATION_INPUTNOTEXIST = "FailedOperation.InputNotExist"
//  FAILEDOPERATION_INPUTUSEDINAUTOCAST = "FailedOperation.InputUsedInAutoCast"
//  FAILEDOPERATION_INPUTUSEDINPGM = "FailedOperation.InputUsedInPgm"
//  FAILEDOPERATION_INPUTUSEDINPVW = "FailedOperation.InputUsedInPvw"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  FAILEDOPERATION_STARTPULLFAILED = "FailedOperation.StartPullFailed"
//  FAILEDOPERATION_STOPPULLFAILED = "FailedOperation.StopPullFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterInputInfoWithContext(ctx context.Context, request *ModifyCasterInputInfoRequest) (response *ModifyCasterInputInfoResponse, err error) {
    if request == nil {
        request = NewModifyCasterInputInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyCasterInputInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCasterInputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCasterInputInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCasterLayoutInfoRequest() (request *ModifyCasterLayoutInfoRequest) {
    request = &ModifyCasterLayoutInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyCasterLayoutInfo")
    
    
    return
}

func NewModifyCasterLayoutInfoResponse() (response *ModifyCasterLayoutInfoResponse) {
    response = &ModifyCasterLayoutInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCasterLayoutInfo
// 该接口用来修改布局参数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTNOTEXIST = "FailedOperation.LayoutNotExist"
//  FAILEDOPERATION_LAYOUTUSEDINAUTOCAST = "FailedOperation.LayoutUsedInAutoCast"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterLayoutInfo(request *ModifyCasterLayoutInfoRequest) (response *ModifyCasterLayoutInfoResponse, err error) {
    return c.ModifyCasterLayoutInfoWithContext(context.Background(), request)
}

// ModifyCasterLayoutInfo
// 该接口用来修改布局参数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTNOTEXIST = "FailedOperation.LayoutNotExist"
//  FAILEDOPERATION_LAYOUTUSEDINAUTOCAST = "FailedOperation.LayoutUsedInAutoCast"
//  FAILEDOPERATION_LAYOUTUSEDINPGM = "FailedOperation.LayoutUsedInPgm"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterLayoutInfoWithContext(ctx context.Context, request *ModifyCasterLayoutInfoRequest) (response *ModifyCasterLayoutInfoResponse, err error) {
    if request == nil {
        request = NewModifyCasterLayoutInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyCasterLayoutInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCasterLayoutInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCasterLayoutInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCasterMarkPicInfoRequest() (request *ModifyCasterMarkPicInfoRequest) {
    request = &ModifyCasterMarkPicInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyCasterMarkPicInfo")
    
    
    return
}

func NewModifyCasterMarkPicInfoResponse() (response *ModifyCasterMarkPicInfoResponse) {
    response = &ModifyCasterMarkPicInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCasterMarkPicInfo
// 该接口用来修改导播台水印信息。
//
// 注意，修改的Index对应的水印需已存在
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_MARKPICNOTEXIST = "FailedOperation.MarkPicNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterMarkPicInfo(request *ModifyCasterMarkPicInfoRequest) (response *ModifyCasterMarkPicInfoResponse, err error) {
    return c.ModifyCasterMarkPicInfoWithContext(context.Background(), request)
}

// ModifyCasterMarkPicInfo
// 该接口用来修改导播台水印信息。
//
// 注意，修改的Index对应的水印需已存在
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_MARKPICNOTEXIST = "FailedOperation.MarkPicNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterMarkPicInfoWithContext(ctx context.Context, request *ModifyCasterMarkPicInfoRequest) (response *ModifyCasterMarkPicInfoResponse, err error) {
    if request == nil {
        request = NewModifyCasterMarkPicInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyCasterMarkPicInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCasterMarkPicInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCasterMarkPicInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCasterMarkWordInfoRequest() (request *ModifyCasterMarkWordInfoRequest) {
    request = &ModifyCasterMarkWordInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyCasterMarkWordInfo")
    
    
    return
}

func NewModifyCasterMarkWordInfoResponse() (response *ModifyCasterMarkWordInfoResponse) {
    response = &ModifyCasterMarkWordInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCasterMarkWordInfo
// 该接口用来修改导播台文本配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  FAILEDOPERATION_MARKWORDNOTEXIST = "FailedOperation.MarkWordNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterMarkWordInfo(request *ModifyCasterMarkWordInfoRequest) (response *ModifyCasterMarkWordInfoResponse, err error) {
    return c.ModifyCasterMarkWordInfoWithContext(context.Background(), request)
}

// ModifyCasterMarkWordInfo
// 该接口用来修改导播台文本配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_LAYOUTUSEDINPVW = "FailedOperation.LayoutUsedInPvw"
//  FAILEDOPERATION_MARKWORDNOTEXIST = "FailedOperation.MarkWordNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterMarkWordInfoWithContext(ctx context.Context, request *ModifyCasterMarkWordInfoRequest) (response *ModifyCasterMarkWordInfoResponse, err error) {
    if request == nil {
        request = NewModifyCasterMarkWordInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyCasterMarkWordInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCasterMarkWordInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCasterMarkWordInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCasterOutputInfoRequest() (request *ModifyCasterOutputInfoRequest) {
    request = &ModifyCasterOutputInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyCasterOutputInfo")
    
    
    return
}

func NewModifyCasterOutputInfoResponse() (response *ModifyCasterOutputInfoResponse) {
    response = &ModifyCasterOutputInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCasterOutputInfo
// 该接口用来修改导播台的推流信息。
//
// 注：只有在主监启动前设置才生效，主监启动后设置，下次推流生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_OUTPUTISNOTEXIST = "FailedOperation.OutputIsNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterOutputInfo(request *ModifyCasterOutputInfoRequest) (response *ModifyCasterOutputInfoResponse, err error) {
    return c.ModifyCasterOutputInfoWithContext(context.Background(), request)
}

// ModifyCasterOutputInfo
// 该接口用来修改导播台的推流信息。
//
// 注：只有在主监启动前设置才生效，主监启动后设置，下次推流生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_OUTPUTISNOTEXIST = "FailedOperation.OutputIsNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidUrl"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCasterOutputInfoWithContext(ctx context.Context, request *ModifyCasterOutputInfoRequest) (response *ModifyCasterOutputInfoResponse, err error) {
    if request == nil {
        request = NewModifyCasterOutputInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyCasterOutputInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCasterOutputInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCasterOutputInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveCallbackTemplateRequest() (request *ModifyLiveCallbackTemplateRequest) {
    request = &ModifyLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveCallbackTemplate")
    
    
    return
}

func NewModifyLiveCallbackTemplateResponse() (response *ModifyLiveCallbackTemplateResponse) {
    response = &ModifyLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveCallbackTemplate
// 修改回调模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETER_URLNOTSAFE = "InvalidParameter.UrlNotSafe"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveCallbackTemplate(request *ModifyLiveCallbackTemplateRequest) (response *ModifyLiveCallbackTemplateResponse, err error) {
    return c.ModifyLiveCallbackTemplateWithContext(context.Background(), request)
}

// ModifyLiveCallbackTemplate
// 修改回调模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETER_URLNOTSAFE = "InvalidParameter.UrlNotSafe"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveCallbackTemplateWithContext(ctx context.Context, request *ModifyLiveCallbackTemplateRequest) (response *ModifyLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveCallbackTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLiveCallbackTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveCallbackTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveDomainCertBindingsRequest() (request *ModifyLiveDomainCertBindingsRequest) {
    request = &ModifyLiveDomainCertBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveDomainCertBindings")
    
    
    return
}

func NewModifyLiveDomainCertBindingsResponse() (response *ModifyLiveDomainCertBindingsResponse) {
    response = &ModifyLiveDomainCertBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveDomainCertBindings
// 批量绑定证书对应的播放域名，并更新启用状态。
//
// 新建自有证书将自动上传至腾讯云ssl。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIGCDNFAILED = "FailedOperation.ConfigCDNFailed"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_CRTDATEINUSING = "InternalError.CrtDateInUsing"
//  INTERNALERROR_CRTDATENOTFOUND = "InternalError.CrtDateNotFound"
//  INTERNALERROR_CRTDATENOTLEGAL = "InternalError.CrtDateNotLegal"
//  INTERNALERROR_CRTDATEOVERDUE = "InternalError.CrtDateOverdue"
//  INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"
//  INTERNALERROR_CRTKEYNOTMATCH = "InternalError.CrtKeyNotMatch"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLOUDCRTIDERROR = "InvalidParameter.CloudCrtIdError"
//  INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"
//  INVALIDPARAMETER_CRTDATENOTFOUND = "InvalidParameter.CrtDateNotFound"
//  INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"
//  INVALIDPARAMETER_CRTDATEOVERDUE = "InvalidParameter.CrtDateOverdue"
//  INVALIDPARAMETER_CRTDOMAINNOTFOUND = "InvalidParameter.CrtDomainNotFound"
//  INVALIDPARAMETER_CRTKEYNOTMATCH = "InvalidParameter.CrtKeyNotMatch"
//  INVALIDPARAMETER_CRTORKEYNOTEXIST = "InvalidParameter.CrtOrKeyNotExist"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  RESOURCENOTFOUND_CRTDATENOTFOUND = "ResourceNotFound.CrtDateNotFound"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveDomainCertBindings(request *ModifyLiveDomainCertBindingsRequest) (response *ModifyLiveDomainCertBindingsResponse, err error) {
    return c.ModifyLiveDomainCertBindingsWithContext(context.Background(), request)
}

// ModifyLiveDomainCertBindings
// 批量绑定证书对应的播放域名，并更新启用状态。
//
// 新建自有证书将自动上传至腾讯云ssl。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIGCDNFAILED = "FailedOperation.ConfigCDNFailed"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_CRTDATEINUSING = "InternalError.CrtDateInUsing"
//  INTERNALERROR_CRTDATENOTFOUND = "InternalError.CrtDateNotFound"
//  INTERNALERROR_CRTDATENOTLEGAL = "InternalError.CrtDateNotLegal"
//  INTERNALERROR_CRTDATEOVERDUE = "InternalError.CrtDateOverdue"
//  INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"
//  INTERNALERROR_CRTKEYNOTMATCH = "InternalError.CrtKeyNotMatch"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLOUDCRTIDERROR = "InvalidParameter.CloudCrtIdError"
//  INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"
//  INVALIDPARAMETER_CRTDATENOTFOUND = "InvalidParameter.CrtDateNotFound"
//  INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"
//  INVALIDPARAMETER_CRTDATEOVERDUE = "InvalidParameter.CrtDateOverdue"
//  INVALIDPARAMETER_CRTDOMAINNOTFOUND = "InvalidParameter.CrtDomainNotFound"
//  INVALIDPARAMETER_CRTKEYNOTMATCH = "InvalidParameter.CrtKeyNotMatch"
//  INVALIDPARAMETER_CRTORKEYNOTEXIST = "InvalidParameter.CrtOrKeyNotExist"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  RESOURCENOTFOUND_CRTDATENOTFOUND = "ResourceNotFound.CrtDateNotFound"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveDomainCertBindingsWithContext(ctx context.Context, request *ModifyLiveDomainCertBindingsRequest) (response *ModifyLiveDomainCertBindingsResponse, err error) {
    if request == nil {
        request = NewModifyLiveDomainCertBindingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLiveDomainCertBindings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveDomainCertBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveDomainCertBindingsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveDomainRefererRequest() (request *ModifyLiveDomainRefererRequest) {
    request = &ModifyLiveDomainRefererRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveDomainReferer")
    
    
    return
}

func NewModifyLiveDomainRefererResponse() (response *ModifyLiveDomainRefererResponse) {
    response = &ModifyLiveDomainRefererResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveDomainReferer
// 设置直播域名 Referer 黑白名单。
//
// 由于 Referer 信息包含在 http 协议中，在开启配置后，播放协议为 rtmp 或 WebRTC 不会校验 Referer 配置，仍可正常播放。如需配置 Referer 鉴权建议使用 http-flv 或 http-hls 协议播放。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveDomainReferer(request *ModifyLiveDomainRefererRequest) (response *ModifyLiveDomainRefererResponse, err error) {
    return c.ModifyLiveDomainRefererWithContext(context.Background(), request)
}

// ModifyLiveDomainReferer
// 设置直播域名 Referer 黑白名单。
//
// 由于 Referer 信息包含在 http 协议中，在开启配置后，播放协议为 rtmp 或 WebRTC 不会校验 Referer 配置，仍可正常播放。如需配置 Referer 鉴权建议使用 http-flv 或 http-hls 协议播放。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveDomainRefererWithContext(ctx context.Context, request *ModifyLiveDomainRefererRequest) (response *ModifyLiveDomainRefererResponse, err error) {
    if request == nil {
        request = NewModifyLiveDomainRefererRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLiveDomainReferer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveDomainReferer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveDomainRefererResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePadTemplateRequest() (request *ModifyLivePadTemplateRequest) {
    request = &ModifyLivePadTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePadTemplate")
    
    
    return
}

func NewModifyLivePadTemplateResponse() (response *ModifyLivePadTemplateResponse) {
    response = &ModifyLivePadTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLivePadTemplate
// 修改直播垫片模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLivePadTemplate(request *ModifyLivePadTemplateRequest) (response *ModifyLivePadTemplateResponse, err error) {
    return c.ModifyLivePadTemplateWithContext(context.Background(), request)
}

// ModifyLivePadTemplate
// 修改直播垫片模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLivePadTemplateWithContext(ctx context.Context, request *ModifyLivePadTemplateRequest) (response *ModifyLivePadTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLivePadTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLivePadTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLivePadTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLivePadTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePlayAuthKeyRequest() (request *ModifyLivePlayAuthKeyRequest) {
    request = &ModifyLivePlayAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePlayAuthKey")
    
    
    return
}

func NewModifyLivePlayAuthKeyResponse() (response *ModifyLivePlayAuthKeyResponse) {
    response = &ModifyLivePlayAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLivePlayAuthKey
// 修改播放鉴权key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePlayAuthKey(request *ModifyLivePlayAuthKeyRequest) (response *ModifyLivePlayAuthKeyResponse, err error) {
    return c.ModifyLivePlayAuthKeyWithContext(context.Background(), request)
}

// ModifyLivePlayAuthKey
// 修改播放鉴权key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePlayAuthKeyWithContext(ctx context.Context, request *ModifyLivePlayAuthKeyRequest) (response *ModifyLivePlayAuthKeyResponse, err error) {
    if request == nil {
        request = NewModifyLivePlayAuthKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLivePlayAuthKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLivePlayAuthKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLivePlayAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePlayDomainRequest() (request *ModifyLivePlayDomainRequest) {
    request = &ModifyLivePlayDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePlayDomain")
    
    
    return
}

func NewModifyLivePlayDomainResponse() (response *ModifyLivePlayDomainResponse) {
    response = &ModifyLivePlayDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLivePlayDomain
// 修改播放域名信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINGSLBFAIL = "FailedOperation.DomainGslbFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePlayDomain(request *ModifyLivePlayDomainRequest) (response *ModifyLivePlayDomainResponse, err error) {
    return c.ModifyLivePlayDomainWithContext(context.Background(), request)
}

// ModifyLivePlayDomain
// 修改播放域名信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINGSLBFAIL = "FailedOperation.DomainGslbFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePlayDomainWithContext(ctx context.Context, request *ModifyLivePlayDomainRequest) (response *ModifyLivePlayDomainResponse, err error) {
    if request == nil {
        request = NewModifyLivePlayDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLivePlayDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLivePlayDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLivePlayDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePullStreamTaskRequest() (request *ModifyLivePullStreamTaskRequest) {
    request = &ModifyLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePullStreamTask")
    
    
    return
}

func NewModifyLivePullStreamTaskResponse() (response *ModifyLivePullStreamTaskResponse) {
    response = &ModifyLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLivePullStreamTask
// 更新直播拉流任务。 
//
// 1. 不支持修改拉流源类型，如需更换，请创建新任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDBACKUPTOURL = "InvalidParameter.InvalidBackupToUrl"
//  INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"
//  INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"
//  INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"
//  INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"
//  INVALIDPARAMETER_INVALIDWATERMARK = "InvalidParameter.InvalidWatermark"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyLivePullStreamTask(request *ModifyLivePullStreamTaskRequest) (response *ModifyLivePullStreamTaskResponse, err error) {
    return c.ModifyLivePullStreamTaskWithContext(context.Background(), request)
}

// ModifyLivePullStreamTask
// 更新直播拉流任务。 
//
// 1. 不支持修改拉流源类型，如需更换，请创建新任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDBACKUPTOURL = "InvalidParameter.InvalidBackupToUrl"
//  INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"
//  INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"
//  INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"
//  INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"
//  INVALIDPARAMETER_INVALIDWATERMARK = "InvalidParameter.InvalidWatermark"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyLivePullStreamTaskWithContext(ctx context.Context, request *ModifyLivePullStreamTaskRequest) (response *ModifyLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewModifyLivePullStreamTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLivePullStreamTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLivePullStreamTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLivePullStreamTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePushAuthKeyRequest() (request *ModifyLivePushAuthKeyRequest) {
    request = &ModifyLivePushAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePushAuthKey")
    
    
    return
}

func NewModifyLivePushAuthKeyResponse() (response *ModifyLivePushAuthKeyResponse) {
    response = &ModifyLivePushAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLivePushAuthKey
// 修改直播推流鉴权key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PUSHDOMAINNORECORD = "InternalError.PushDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePushAuthKey(request *ModifyLivePushAuthKeyRequest) (response *ModifyLivePushAuthKeyResponse, err error) {
    return c.ModifyLivePushAuthKeyWithContext(context.Background(), request)
}

// ModifyLivePushAuthKey
// 修改直播推流鉴权key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PUSHDOMAINNORECORD = "InternalError.PushDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePushAuthKeyWithContext(ctx context.Context, request *ModifyLivePushAuthKeyRequest) (response *ModifyLivePushAuthKeyResponse, err error) {
    if request == nil {
        request = NewModifyLivePushAuthKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLivePushAuthKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLivePushAuthKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLivePushAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveRecordTemplateRequest() (request *ModifyLiveRecordTemplateRequest) {
    request = &ModifyLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveRecordTemplate")
    
    
    return
}

func NewModifyLiveRecordTemplateResponse() (response *ModifyLiveRecordTemplateResponse) {
    response = &ModifyLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveRecordTemplate
// 修改录制模板配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveRecordTemplate(request *ModifyLiveRecordTemplateRequest) (response *ModifyLiveRecordTemplateResponse, err error) {
    return c.ModifyLiveRecordTemplateWithContext(context.Background(), request)
}

// ModifyLiveRecordTemplate
// 修改录制模板配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveRecordTemplateWithContext(ctx context.Context, request *ModifyLiveRecordTemplateRequest) (response *ModifyLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLiveRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveSnapshotTemplateRequest() (request *ModifyLiveSnapshotTemplateRequest) {
    request = &ModifyLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveSnapshotTemplate")
    
    
    return
}

func NewModifyLiveSnapshotTemplateResponse() (response *ModifyLiveSnapshotTemplateResponse) {
    response = &ModifyLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveSnapshotTemplate
// 修改截图模板配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_COSBUCKETNOTEXIST = "FailedOperation.CosBucketNotExist"
//  FAILEDOPERATION_COSBUCKETNOTPERMISSION = "FailedOperation.CosBucketNotPermission"
//  FAILEDOPERATION_COSROLENOTEXISTS = "FailedOperation.CosRoleNotExists"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveSnapshotTemplate(request *ModifyLiveSnapshotTemplateRequest) (response *ModifyLiveSnapshotTemplateResponse, err error) {
    return c.ModifyLiveSnapshotTemplateWithContext(context.Background(), request)
}

// ModifyLiveSnapshotTemplate
// 修改截图模板配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_COSBUCKETNOTEXIST = "FailedOperation.CosBucketNotExist"
//  FAILEDOPERATION_COSBUCKETNOTPERMISSION = "FailedOperation.CosBucketNotPermission"
//  FAILEDOPERATION_COSROLENOTEXISTS = "FailedOperation.CosRoleNotExists"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveSnapshotTemplateWithContext(ctx context.Context, request *ModifyLiveSnapshotTemplateRequest) (response *ModifyLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLiveSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveStreamMonitorRequest() (request *ModifyLiveStreamMonitorRequest) {
    request = &ModifyLiveStreamMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveStreamMonitor")
    
    
    return
}

func NewModifyLiveStreamMonitorResponse() (response *ModifyLiveStreamMonitorResponse) {
    response = &ModifyLiveStreamMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveStreamMonitor
// 该接口用来修改直播流监播任务的配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_INPUTSTREAMMIXTYPENOTACCESSIBLE = "FailedOperation.InputStreamMixTypeNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_RELATEDRUNNINGMONITORLIMITEXCEEDED = "FailedOperation.RelatedRunningMonitorLimitExceeded"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveStreamMonitor(request *ModifyLiveStreamMonitorRequest) (response *ModifyLiveStreamMonitorResponse, err error) {
    return c.ModifyLiveStreamMonitorWithContext(context.Background(), request)
}

// ModifyLiveStreamMonitor
// 该接口用来修改直播流监播任务的配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_INPUTSTREAMMIXTYPENOTACCESSIBLE = "FailedOperation.InputStreamMixTypeNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_RELATEDRUNNINGMONITORLIMITEXCEEDED = "FailedOperation.RelatedRunningMonitorLimitExceeded"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveStreamMonitorWithContext(ctx context.Context, request *ModifyLiveStreamMonitorRequest) (response *ModifyLiveStreamMonitorResponse, err error) {
    if request == nil {
        request = NewModifyLiveStreamMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLiveStreamMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveStreamMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveStreamMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveTimeShiftTemplateRequest() (request *ModifyLiveTimeShiftTemplateRequest) {
    request = &ModifyLiveTimeShiftTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveTimeShiftTemplate")
    
    
    return
}

func NewModifyLiveTimeShiftTemplateResponse() (response *ModifyLiveTimeShiftTemplateResponse) {
    response = &ModifyLiveTimeShiftTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveTimeShiftTemplate
// 修改直播时移模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveTimeShiftTemplate(request *ModifyLiveTimeShiftTemplateRequest) (response *ModifyLiveTimeShiftTemplateResponse, err error) {
    return c.ModifyLiveTimeShiftTemplateWithContext(context.Background(), request)
}

// ModifyLiveTimeShiftTemplate
// 修改直播时移模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveTimeShiftTemplateWithContext(ctx context.Context, request *ModifyLiveTimeShiftTemplateRequest) (response *ModifyLiveTimeShiftTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveTimeShiftTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLiveTimeShiftTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveTimeShiftTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveTimeShiftTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveTranscodeTemplateRequest() (request *ModifyLiveTranscodeTemplateRequest) {
    request = &ModifyLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveTranscodeTemplate")
    
    
    return
}

func NewModifyLiveTranscodeTemplateResponse() (response *ModifyLiveTranscodeTemplateResponse) {
    response = &ModifyLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveTranscodeTemplate
// 修改转码模板配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_PROCESSORALREADYEXIST = "InternalError.ProcessorAlreadyExist"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_GOPMUSTEQUALANDEXISTS = "InvalidParameter.GopMustEqualAndExists"
//  INVALIDPARAMETER_PROCESSORALREADYEXIST = "InvalidParameter.ProcessorAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveTranscodeTemplate(request *ModifyLiveTranscodeTemplateRequest) (response *ModifyLiveTranscodeTemplateResponse, err error) {
    return c.ModifyLiveTranscodeTemplateWithContext(context.Background(), request)
}

// ModifyLiveTranscodeTemplate
// 修改转码模板配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_PROCESSORALREADYEXIST = "InternalError.ProcessorAlreadyExist"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_GOPMUSTEQUALANDEXISTS = "InvalidParameter.GopMustEqualAndExists"
//  INVALIDPARAMETER_PROCESSORALREADYEXIST = "InvalidParameter.ProcessorAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveTranscodeTemplateWithContext(ctx context.Context, request *ModifyLiveTranscodeTemplateRequest) (response *ModifyLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyLiveTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPullStreamConfigRequest() (request *ModifyPullStreamConfigRequest) {
    request = &ModifyPullStreamConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyPullStreamConfig")
    
    
    return
}

func NewModifyPullStreamConfigResponse() (response *ModifyPullStreamConfigResponse) {
    response = &ModifyPullStreamConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPullStreamConfig
// 更新拉流配置。该接口为已下线接口，请使用新接口 ModifyLivePullStreamTask。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyPullStreamConfig(request *ModifyPullStreamConfigRequest) (response *ModifyPullStreamConfigResponse, err error) {
    return c.ModifyPullStreamConfigWithContext(context.Background(), request)
}

// ModifyPullStreamConfig
// 更新拉流配置。该接口为已下线接口，请使用新接口 ModifyLivePullStreamTask。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyPullStreamConfigWithContext(ctx context.Context, request *ModifyPullStreamConfigRequest) (response *ModifyPullStreamConfigResponse, err error) {
    if request == nil {
        request = NewModifyPullStreamConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyPullStreamConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPullStreamConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPullStreamConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPullStreamStatusRequest() (request *ModifyPullStreamStatusRequest) {
    request = &ModifyPullStreamStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyPullStreamStatus")
    
    
    return
}

func NewModifyPullStreamStatusResponse() (response *ModifyPullStreamStatusResponse) {
    response = &ModifyPullStreamStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPullStreamStatus
// 修改直播拉流配置的状态。该接口已下线,请使用新接口 ModifyLivePullStreamTask。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyPullStreamStatus(request *ModifyPullStreamStatusRequest) (response *ModifyPullStreamStatusResponse, err error) {
    return c.ModifyPullStreamStatusWithContext(context.Background(), request)
}

// ModifyPullStreamStatus
// 修改直播拉流配置的状态。该接口已下线,请使用新接口 ModifyLivePullStreamTask。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyPullStreamStatusWithContext(ctx context.Context, request *ModifyPullStreamStatusRequest) (response *ModifyPullStreamStatusResponse, err error) {
    if request == nil {
        request = NewModifyPullStreamStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ModifyPullStreamStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPullStreamStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPullStreamStatusResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseCasterRequest() (request *ReleaseCasterRequest) {
    request = &ReleaseCasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ReleaseCaster")
    
    
    return
}

func NewReleaseCasterResponse() (response *ReleaseCasterResponse) {
    response = &ReleaseCasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseCaster
// 调用该接口，释放导播台实例，但保留所有的配置。
//
// 执行该接口，预监与主监画面停止，第三方推流停止。
//
// 点播文件与直播地址将停止展示，客户自行推到导播台的流需要手动停止。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RELEASECASTERFAILED = "FailedOperation.ReleaseCasterFailed"
//  FAILEDOPERATION_STOPCASTERTASKFAILED = "FailedOperation.StopCasterTaskFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ReleaseCaster(request *ReleaseCasterRequest) (response *ReleaseCasterResponse, err error) {
    return c.ReleaseCasterWithContext(context.Background(), request)
}

// ReleaseCaster
// 调用该接口，释放导播台实例，但保留所有的配置。
//
// 执行该接口，预监与主监画面停止，第三方推流停止。
//
// 点播文件与直播地址将停止展示，客户自行推到导播台的流需要手动停止。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RELEASECASTERFAILED = "FailedOperation.ReleaseCasterFailed"
//  FAILEDOPERATION_STOPCASTERTASKFAILED = "FailedOperation.StopCasterTaskFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ReleaseCasterWithContext(ctx context.Context, request *ReleaseCasterRequest) (response *ReleaseCasterResponse, err error) {
    if request == nil {
        request = NewReleaseCasterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ReleaseCaster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseCaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseCasterResponse()
    err = c.Send(request, response)
    return
}

func NewRestartLivePullStreamTaskRequest() (request *RestartLivePullStreamTaskRequest) {
    request = &RestartLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "RestartLivePullStreamTask")
    
    
    return
}

func NewRestartLivePullStreamTaskResponse() (response *RestartLivePullStreamTaskResponse) {
    response = &RestartLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartLivePullStreamTask
// 将正在运行的拉流转推任务进行重启。
//
// 注意：
//
// 1. 重启任务会造成推流中断。
//
// 2. 点播源任务的重启，会根据VodRefreshType决定是续播还是从头开始播。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) RestartLivePullStreamTask(request *RestartLivePullStreamTaskRequest) (response *RestartLivePullStreamTaskResponse, err error) {
    return c.RestartLivePullStreamTaskWithContext(context.Background(), request)
}

// RestartLivePullStreamTask
// 将正在运行的拉流转推任务进行重启。
//
// 注意：
//
// 1. 重启任务会造成推流中断。
//
// 2. 点播源任务的重启，会根据VodRefreshType决定是续播还是从头开始播。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) RestartLivePullStreamTaskWithContext(ctx context.Context, request *RestartLivePullStreamTaskRequest) (response *RestartLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewRestartLivePullStreamTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "RestartLivePullStreamTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartLivePullStreamTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartLivePullStreamTaskResponse()
    err = c.Send(request, response)
    return
}

func NewResumeDelayLiveStreamRequest() (request *ResumeDelayLiveStreamRequest) {
    request = &ResumeDelayLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ResumeDelayLiveStream")
    
    
    return
}

func NewResumeDelayLiveStreamResponse() (response *ResumeDelayLiveStreamResponse) {
    response = &ResumeDelayLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeDelayLiveStream
// 取消直播流设置的延时配置，恢复实时直播画面。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeDelayLiveStream(request *ResumeDelayLiveStreamRequest) (response *ResumeDelayLiveStreamResponse, err error) {
    return c.ResumeDelayLiveStreamWithContext(context.Background(), request)
}

// ResumeDelayLiveStream
// 取消直播流设置的延时配置，恢复实时直播画面。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeDelayLiveStreamWithContext(ctx context.Context, request *ResumeDelayLiveStreamRequest) (response *ResumeDelayLiveStreamResponse, err error) {
    if request == nil {
        request = NewResumeDelayLiveStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ResumeDelayLiveStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeDelayLiveStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeDelayLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewResumeLiveStreamRequest() (request *ResumeLiveStreamRequest) {
    request = &ResumeLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ResumeLiveStream")
    
    
    return
}

func NewResumeLiveStreamResponse() (response *ResumeLiveStreamResponse) {
    response = &ResumeLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeLiveStream
// 恢复某条流的推流。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ResumeLiveStream(request *ResumeLiveStreamRequest) (response *ResumeLiveStreamResponse, err error) {
    return c.ResumeLiveStreamWithContext(context.Background(), request)
}

// ResumeLiveStream
// 恢复某条流的推流。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ResumeLiveStreamWithContext(ctx context.Context, request *ResumeLiveStreamRequest) (response *ResumeLiveStreamResponse, err error) {
    if request == nil {
        request = NewResumeLiveStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "ResumeLiveStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeLiveStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewSendLiveCloudEffectRequest() (request *SendLiveCloudEffectRequest) {
    request = &SendLiveCloudEffectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "SendLiveCloudEffect")
    
    
    return
}

func NewSendLiveCloudEffectResponse() (response *SendLiveCloudEffectResponse) {
    response = &SendLiveCloudEffectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendLiveCloudEffect
// 使用该接口发送云端特效到线上正活跃的直播流，观众可在播放端看到特效从直播流画面中展示。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SendLiveCloudEffect(request *SendLiveCloudEffectRequest) (response *SendLiveCloudEffectResponse, err error) {
    return c.SendLiveCloudEffectWithContext(context.Background(), request)
}

// SendLiveCloudEffect
// 使用该接口发送云端特效到线上正活跃的直播流，观众可在播放端看到特效从直播流画面中展示。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SendLiveCloudEffectWithContext(ctx context.Context, request *SendLiveCloudEffectRequest) (response *SendLiveCloudEffectResponse, err error) {
    if request == nil {
        request = NewSendLiveCloudEffectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "SendLiveCloudEffect")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendLiveCloudEffect require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendLiveCloudEffectResponse()
    err = c.Send(request, response)
    return
}

func NewStartLivePadStreamRequest() (request *StartLivePadStreamRequest) {
    request = &StartLivePadStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "StartLivePadStream")
    
    
    return
}

func NewStartLivePadStreamResponse() (response *StartLivePadStreamResponse) {
    response = &StartLivePadStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartLivePadStream
// 使用该接口将直播流开始切入垫片。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) StartLivePadStream(request *StartLivePadStreamRequest) (response *StartLivePadStreamResponse, err error) {
    return c.StartLivePadStreamWithContext(context.Background(), request)
}

// StartLivePadStream
// 使用该接口将直播流开始切入垫片。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) StartLivePadStreamWithContext(ctx context.Context, request *StartLivePadStreamRequest) (response *StartLivePadStreamResponse, err error) {
    if request == nil {
        request = NewStartLivePadStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "StartLivePadStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartLivePadStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartLivePadStreamResponse()
    err = c.Send(request, response)
    return
}

func NewStartLiveStreamMonitorRequest() (request *StartLiveStreamMonitorRequest) {
    request = &StartLiveStreamMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "StartLiveStreamMonitor")
    
    
    return
}

func NewStartLiveStreamMonitorResponse() (response *StartLiveStreamMonitorResponse) {
    response = &StartLiveStreamMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartLiveStreamMonitor
// 该接口用来启动直播流监播任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_INPUTSTREAMMIXTYPENOTACCESSIBLE = "FailedOperation.InputStreamMixTypeNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) StartLiveStreamMonitor(request *StartLiveStreamMonitorRequest) (response *StartLiveStreamMonitorResponse, err error) {
    return c.StartLiveStreamMonitorWithContext(context.Background(), request)
}

// StartLiveStreamMonitor
// 该接口用来启动直播流监播任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_INPUTSTREAMMIXTYPENOTACCESSIBLE = "FailedOperation.InputStreamMixTypeNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) StartLiveStreamMonitorWithContext(ctx context.Context, request *StartLiveStreamMonitorRequest) (response *StartLiveStreamMonitorResponse, err error) {
    if request == nil {
        request = NewStartLiveStreamMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "StartLiveStreamMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartLiveStreamMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartLiveStreamMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewStopCasterPgmRequest() (request *StopCasterPgmRequest) {
    request = &StopCasterPgmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "StopCasterPgm")
    
    
    return
}

func NewStopCasterPgmResponse() (response *StopCasterPgmResponse) {
    response = &StopCasterPgmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCasterPgm
// 该接口用来停止导播台的主监输出。
//
// 停止主监后，对应的推流到腾讯云直播源站和推流到其他第三方平台均将会停止。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_STOPCASTERTASKFAILED = "FailedOperation.StopCasterTaskFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopCasterPgm(request *StopCasterPgmRequest) (response *StopCasterPgmResponse, err error) {
    return c.StopCasterPgmWithContext(context.Background(), request)
}

// StopCasterPgm
// 该接口用来停止导播台的主监输出。
//
// 停止主监后，对应的推流到腾讯云直播源站和推流到其他第三方平台均将会停止。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_STOPCASTERTASKFAILED = "FailedOperation.StopCasterTaskFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopCasterPgmWithContext(ctx context.Context, request *StopCasterPgmRequest) (response *StopCasterPgmResponse, err error) {
    if request == nil {
        request = NewStopCasterPgmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "StopCasterPgm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCasterPgm require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCasterPgmResponse()
    err = c.Send(request, response)
    return
}

func NewStopCasterPvwRequest() (request *StopCasterPvwRequest) {
    request = &StopCasterPvwRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "StopCasterPvw")
    
    
    return
}

func NewStopCasterPvwResponse() (response *StopCasterPvwResponse) {
    response = &StopCasterPvwResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCasterPvw
// 该接口用来停止导播台的预监任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_STOPCASTERTASKFAILED = "FailedOperation.StopCasterTaskFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopCasterPvw(request *StopCasterPvwRequest) (response *StopCasterPvwResponse, err error) {
    return c.StopCasterPvwWithContext(context.Background(), request)
}

// StopCasterPvw
// 该接口用来停止导播台的预监任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CASTERNOTFOUND = "FailedOperation.CasterNotFound"
//  FAILEDOPERATION_STOPCASTERTASKFAILED = "FailedOperation.StopCasterTaskFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopCasterPvwWithContext(ctx context.Context, request *StopCasterPvwRequest) (response *StopCasterPvwResponse, err error) {
    if request == nil {
        request = NewStopCasterPvwRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "StopCasterPvw")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCasterPvw require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCasterPvwResponse()
    err = c.Send(request, response)
    return
}

func NewStopLivePadStreamRequest() (request *StopLivePadStreamRequest) {
    request = &StopLivePadStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "StopLivePadStream")
    
    
    return
}

func NewStopLivePadStreamResponse() (response *StopLivePadStreamResponse) {
    response = &StopLivePadStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopLivePadStream
// 使用该接口将直播流停止切入垫片。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) StopLivePadStream(request *StopLivePadStreamRequest) (response *StopLivePadStreamResponse, err error) {
    return c.StopLivePadStreamWithContext(context.Background(), request)
}

// StopLivePadStream
// 使用该接口将直播流停止切入垫片。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) StopLivePadStreamWithContext(ctx context.Context, request *StopLivePadStreamRequest) (response *StopLivePadStreamResponse, err error) {
    if request == nil {
        request = NewStopLivePadStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "StopLivePadStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopLivePadStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopLivePadStreamResponse()
    err = c.Send(request, response)
    return
}

func NewStopLiveRecordRequest() (request *StopLiveRecordRequest) {
    request = &StopLiveRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "StopLiveRecord")
    
    
    return
}

func NewStopLiveRecordResponse() (response *StopLiveRecordResponse) {
    response = &StopLiveRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopLiveRecord
// 说明：录制后的文件存放于点播平台。用户如需使用录制功能，需首先自行开通点播账号并确保账号可用。录制文件存放后，相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，请参考对应文档。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) StopLiveRecord(request *StopLiveRecordRequest) (response *StopLiveRecordResponse, err error) {
    return c.StopLiveRecordWithContext(context.Background(), request)
}

// StopLiveRecord
// 说明：录制后的文件存放于点播平台。用户如需使用录制功能，需首先自行开通点播账号并确保账号可用。录制文件存放后，相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，请参考对应文档。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) StopLiveRecordWithContext(ctx context.Context, request *StopLiveRecordRequest) (response *StopLiveRecordResponse, err error) {
    if request == nil {
        request = NewStopLiveRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "StopLiveRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopLiveRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopLiveRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopLiveStreamMonitorRequest() (request *StopLiveStreamMonitorRequest) {
    request = &StopLiveStreamMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "StopLiveStreamMonitor")
    
    
    return
}

func NewStopLiveStreamMonitorResponse() (response *StopLiveStreamMonitorResponse) {
    response = &StopLiveStreamMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopLiveStreamMonitor
// 该接口用来停止直播流监播任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) StopLiveStreamMonitor(request *StopLiveStreamMonitorRequest) (response *StopLiveStreamMonitorResponse, err error) {
    return c.StopLiveStreamMonitorWithContext(context.Background(), request)
}

// StopLiveStreamMonitor
// 该接口用来停止直播流监播任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTACTION = "FailedOperation.ConflictAction"
//  FAILEDOPERATION_DATABASENOTACCESSIBLE = "FailedOperation.DatabaseNotAccessible"
//  FAILEDOPERATION_MONITORISACTIVE = "FailedOperation.MonitorIsActive"
//  FAILEDOPERATION_MONITORLIMITEXCEEDED = "FailedOperation.MonitorLimitExceeded"
//  FAILEDOPERATION_MONITORNOTEXIST = "FailedOperation.MonitorNotExist"
//  FAILEDOPERATION_STARTTASKFAILED = "FailedOperation.StartTaskFailed"
//  FAILEDOPERATION_STOPTASKFAILED = "FailedOperation.StopTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTALLOWUSELVB = "InvalidParameter.NotAllowUseLVB"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) StopLiveStreamMonitorWithContext(ctx context.Context, request *StopLiveStreamMonitorRequest) (response *StopLiveStreamMonitorResponse, err error) {
    if request == nil {
        request = NewStopLiveStreamMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "StopLiveStreamMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopLiveStreamMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopLiveStreamMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewStopRecordTaskRequest() (request *StopRecordTaskRequest) {
    request = &StopRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "StopRecordTask")
    
    
    return
}

func NewStopRecordTaskResponse() (response *StopRecordTaskResponse) {
    response = &StopRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopRecordTask
// 提前结束录制，中止运行中的录制任务并生成录制文件。任务被成功终止后，本次任务将不再启动。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopRecordTask(request *StopRecordTaskRequest) (response *StopRecordTaskResponse, err error) {
    return c.StopRecordTaskWithContext(context.Background(), request)
}

// StopRecordTask
// 提前结束录制，中止运行中的录制任务并生成录制文件。任务被成功终止后，本次任务将不再启动。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopRecordTaskWithContext(ctx context.Context, request *StopRecordTaskRequest) (response *StopRecordTaskResponse, err error) {
    if request == nil {
        request = NewStopRecordTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "StopRecordTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRecordTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopRecordTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopScreenshotTaskRequest() (request *StopScreenshotTaskRequest) {
    request = &StopScreenshotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "StopScreenshotTask")
    
    
    return
}

func NewStopScreenshotTaskResponse() (response *StopScreenshotTaskResponse) {
    response = &StopScreenshotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopScreenshotTask
// 提前结束截图，中止运行中的截图任务。任务被成功终止后，本次任务将不再启动。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopScreenshotTask(request *StopScreenshotTaskRequest) (response *StopScreenshotTaskResponse, err error) {
    return c.StopScreenshotTaskWithContext(context.Background(), request)
}

// StopScreenshotTask
// 提前结束截图，中止运行中的截图任务。任务被成功终止后，本次任务将不再启动。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopScreenshotTaskWithContext(ctx context.Context, request *StopScreenshotTaskRequest) (response *StopScreenshotTaskResponse, err error) {
    if request == nil {
        request = NewStopScreenshotTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "StopScreenshotTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopScreenshotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopScreenshotTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchBackupStreamRequest() (request *SwitchBackupStreamRequest) {
    request = &SwitchBackupStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "SwitchBackupStream")
    
    
    return
}

func NewSwitchBackupStreamResponse() (response *SwitchBackupStreamResponse) {
    response = &SwitchBackupStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchBackupStream
// 调用该接口实现切换当前播放所使用的主备流。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) SwitchBackupStream(request *SwitchBackupStreamRequest) (response *SwitchBackupStreamResponse, err error) {
    return c.SwitchBackupStreamWithContext(context.Background(), request)
}

// SwitchBackupStream
// 调用该接口实现切换当前播放所使用的主备流。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) SwitchBackupStreamWithContext(ctx context.Context, request *SwitchBackupStreamRequest) (response *SwitchBackupStreamResponse, err error) {
    if request == nil {
        request = NewSwitchBackupStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "SwitchBackupStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchBackupStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchBackupStreamResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindLiveDomainCertRequest() (request *UnBindLiveDomainCertRequest) {
    request = &UnBindLiveDomainCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "UnBindLiveDomainCert")
    
    
    return
}

func NewUnBindLiveDomainCertResponse() (response *UnBindLiveDomainCertResponse) {
    response = &UnBindLiveDomainCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnBindLiveDomainCert
// 解绑域名证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) UnBindLiveDomainCert(request *UnBindLiveDomainCertRequest) (response *UnBindLiveDomainCertResponse, err error) {
    return c.UnBindLiveDomainCertWithContext(context.Background(), request)
}

// UnBindLiveDomainCert
// 解绑域名证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) UnBindLiveDomainCertWithContext(ctx context.Context, request *UnBindLiveDomainCertRequest) (response *UnBindLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewUnBindLiveDomainCertRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "UnBindLiveDomainCert")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindLiveDomainCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateLiveWatermarkRequest() (request *UpdateLiveWatermarkRequest) {
    request = &UpdateLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "UpdateLiveWatermark")
    
    
    return
}

func NewUpdateLiveWatermarkResponse() (response *UpdateLiveWatermarkResponse) {
    response = &UpdateLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateLiveWatermark
// 更新水印。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_WATERMARKEDITERROR = "InternalError.WatermarkEditError"
//  INTERNALERROR_WATERMARKNOTEXIST = "InternalError.WatermarkNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) UpdateLiveWatermark(request *UpdateLiveWatermarkRequest) (response *UpdateLiveWatermarkResponse, err error) {
    return c.UpdateLiveWatermarkWithContext(context.Background(), request)
}

// UpdateLiveWatermark
// 更新水印。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_WATERMARKEDITERROR = "InternalError.WatermarkEditError"
//  INTERNALERROR_WATERMARKNOTEXIST = "InternalError.WatermarkNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) UpdateLiveWatermarkWithContext(ctx context.Context, request *UpdateLiveWatermarkRequest) (response *UpdateLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewUpdateLiveWatermarkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "live", APIVersion, "UpdateLiveWatermark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateLiveWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}
