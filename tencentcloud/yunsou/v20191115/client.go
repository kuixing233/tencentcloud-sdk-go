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

package v20191115

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-11-15"

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


func NewDataManipulationRequest() (request *DataManipulationRequest) {
    request = &DataManipulationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunsou", APIVersion, "DataManipulation")
    
    
    return
}

func NewDataManipulationResponse() (response *DataManipulationResponse) {
    response = &DataManipulationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DataManipulation
// 上传云搜数据的API接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTINFO = "FailedOperation.AccountInfo"
//  FAILEDOPERATION_APPINFO = "FailedOperation.AppInfo"
//  FAILEDOPERATION_SEARCH = "FailedOperation.Search"
//  FAILEDOPERATION_UPLOADDATAAPIFAIL = "FailedOperation.UploadDataApiFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATACONTENT = "InvalidParameter.DataContent"
//  LIMITEXCEEDED_SEARCH = "LimitExceeded.Search"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DataManipulation(request *DataManipulationRequest) (response *DataManipulationResponse, err error) {
    return c.DataManipulationWithContext(context.Background(), request)
}

// DataManipulation
// 上传云搜数据的API接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTINFO = "FailedOperation.AccountInfo"
//  FAILEDOPERATION_APPINFO = "FailedOperation.AppInfo"
//  FAILEDOPERATION_SEARCH = "FailedOperation.Search"
//  FAILEDOPERATION_UPLOADDATAAPIFAIL = "FailedOperation.UploadDataApiFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATACONTENT = "InvalidParameter.DataContent"
//  LIMITEXCEEDED_SEARCH = "LimitExceeded.Search"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DataManipulationWithContext(ctx context.Context, request *DataManipulationRequest) (response *DataManipulationResponse, err error) {
    if request == nil {
        request = NewDataManipulationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "yunsou", APIVersion, "DataManipulation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DataManipulation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDataManipulationResponse()
    err = c.Send(request, response)
    return
}

func NewDataSearchRequest() (request *DataSearchRequest) {
    request = &DataSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yunsou", APIVersion, "DataSearch")
    
    
    return
}

func NewDataSearchResponse() (response *DataSearchResponse) {
    response = &DataSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DataSearch
// 用于检索云搜中的数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTINFO = "FailedOperation.AccountInfo"
//  FAILEDOPERATION_APPINFO = "FailedOperation.AppInfo"
//  FAILEDOPERATION_SEARCH = "FailedOperation.Search"
//  FAILEDOPERATION_UPLOADDATAAPIFAIL = "FailedOperation.UploadDataApiFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATACONTENT = "InvalidParameter.DataContent"
//  LIMITEXCEEDED_SEARCH = "LimitExceeded.Search"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DataSearch(request *DataSearchRequest) (response *DataSearchResponse, err error) {
    return c.DataSearchWithContext(context.Background(), request)
}

// DataSearch
// 用于检索云搜中的数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTINFO = "FailedOperation.AccountInfo"
//  FAILEDOPERATION_APPINFO = "FailedOperation.AppInfo"
//  FAILEDOPERATION_SEARCH = "FailedOperation.Search"
//  FAILEDOPERATION_UPLOADDATAAPIFAIL = "FailedOperation.UploadDataApiFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATACONTENT = "InvalidParameter.DataContent"
//  LIMITEXCEEDED_SEARCH = "LimitExceeded.Search"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DataSearchWithContext(ctx context.Context, request *DataSearchRequest) (response *DataSearchResponse, err error) {
    if request == nil {
        request = NewDataSearchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "yunsou", APIVersion, "DataSearch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DataSearch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDataSearchResponse()
    err = c.Send(request, response)
    return
}
