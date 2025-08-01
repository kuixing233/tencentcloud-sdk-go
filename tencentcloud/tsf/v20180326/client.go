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

package v20180326

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-26"

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


func NewAddClusterInstancesRequest() (request *AddClusterInstancesRequest) {
    request = &AddClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "AddClusterInstances")
    
    
    return
}

func NewAddClusterInstancesResponse() (response *AddClusterInstancesResponse) {
    response = &AddClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddClusterInstances
// 添加云主机节点至TSF集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_INSTANCERESETERROR = "FailedOperation.InstanceResetError"
//  FAILEDOPERATION_INSTANCERESETTIMEOUT = "FailedOperation.InstanceResetTimeout"
//  FAILEDOPERATION_INSTANCEUPDATEFAILED = "FailedOperation.InstanceUpdateFailed"
//  FAILEDOPERATION_TKECLUSTERQUERYFAILED = "FailedOperation.TkeClusterQueryFailed"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  INVALIDPARAMETERVALUE_INSTANCEINVALIDIMAGE = "InvalidParameterValue.InstanceInvalidImage"
//  RESOURCEINUSE_INSTANCEHASBEENUSED = "ResourceInUse.InstanceHasBeenUsed"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_TKECLUSTERNOTEXISTS = "ResourceNotFound.TkeClusterNotExists"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) AddClusterInstances(request *AddClusterInstancesRequest) (response *AddClusterInstancesResponse, err error) {
    return c.AddClusterInstancesWithContext(context.Background(), request)
}

// AddClusterInstances
// 添加云主机节点至TSF集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_INSTANCERESETERROR = "FailedOperation.InstanceResetError"
//  FAILEDOPERATION_INSTANCERESETTIMEOUT = "FailedOperation.InstanceResetTimeout"
//  FAILEDOPERATION_INSTANCEUPDATEFAILED = "FailedOperation.InstanceUpdateFailed"
//  FAILEDOPERATION_TKECLUSTERQUERYFAILED = "FailedOperation.TkeClusterQueryFailed"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  INVALIDPARAMETERVALUE_INSTANCEINVALIDIMAGE = "InvalidParameterValue.InstanceInvalidImage"
//  RESOURCEINUSE_INSTANCEHASBEENUSED = "ResourceInUse.InstanceHasBeenUsed"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_TKECLUSTERNOTEXISTS = "ResourceNotFound.TkeClusterNotExists"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) AddClusterInstancesWithContext(ctx context.Context, request *AddClusterInstancesRequest) (response *AddClusterInstancesResponse, err error) {
    if request == nil {
        request = NewAddClusterInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "AddClusterInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAddInstancesRequest() (request *AddInstancesRequest) {
    request = &AddInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "AddInstances")
    
    
    return
}

func NewAddInstancesResponse() (response *AddInstancesResponse) {
    response = &AddInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddInstances
// 添加云主机节点至TSF集群
//
// 可能返回的错误码:
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEINVALIDIMAGE = "InvalidParameterValue.InstanceInvalidImage"
//  MISSINGPARAMETER_CLUSTERIDREQUIRED = "MissingParameter.ClusterIdRequired"
//  MISSINGPARAMETER_INSTANCEIMPORTMODENULL = "MissingParameter.InstanceImportModeNull"
//  RESOURCEINUSE_INSTANCEHASBEENUSED = "ResourceInUse.InstanceHasBeenUsed"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
func (c *Client) AddInstances(request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    return c.AddInstancesWithContext(context.Background(), request)
}

// AddInstances
// 添加云主机节点至TSF集群
//
// 可能返回的错误码:
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEINVALIDIMAGE = "InvalidParameterValue.InstanceInvalidImage"
//  MISSINGPARAMETER_CLUSTERIDREQUIRED = "MissingParameter.ClusterIdRequired"
//  MISSINGPARAMETER_INSTANCEIMPORTMODENULL = "MissingParameter.InstanceImportModeNull"
//  RESOURCEINUSE_INSTANCEHASBEENUSED = "ResourceInUse.InstanceHasBeenUsed"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
func (c *Client) AddInstancesWithContext(ctx context.Context, request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    if request == nil {
        request = NewAddInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "AddInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateBusinessLogConfigRequest() (request *AssociateBusinessLogConfigRequest) {
    request = &AssociateBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "AssociateBusinessLogConfig")
    
    
    return
}

func NewAssociateBusinessLogConfigResponse() (response *AssociateBusinessLogConfigResponse) {
    response = &AssociateBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateBusinessLogConfig
// 关联日志配置项到应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_TSFAPMBUSILOGCFGAPPRELATIONMASTERERROR = "InternalError.TsfApmBusiLogCfgAppRelationMasterError"
//  INTERNALERROR_TSFAPMCOMMONERROR = "InternalError.TsfApmCommonError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) AssociateBusinessLogConfig(request *AssociateBusinessLogConfigRequest) (response *AssociateBusinessLogConfigResponse, err error) {
    return c.AssociateBusinessLogConfigWithContext(context.Background(), request)
}

// AssociateBusinessLogConfig
// 关联日志配置项到应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_TSFAPMBUSILOGCFGAPPRELATIONMASTERERROR = "InternalError.TsfApmBusiLogCfgAppRelationMasterError"
//  INTERNALERROR_TSFAPMCOMMONERROR = "InternalError.TsfApmCommonError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) AssociateBusinessLogConfigWithContext(ctx context.Context, request *AssociateBusinessLogConfigRequest) (response *AssociateBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewAssociateBusinessLogConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "AssociateBusinessLogConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateBusinessLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateConfigWithGroupRequest() (request *AssociateConfigWithGroupRequest) {
    request = &AssociateConfigWithGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "AssociateConfigWithGroup")
    
    
    return
}

func NewAssociateConfigWithGroupResponse() (response *AssociateConfigWithGroupResponse) {
    response = &AssociateConfigWithGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateConfigWithGroup
// 关联投递配置到部署组
//
// 可能返回的错误码:
//  INTERNALERROR_TSFAPMCOMMONERROR = "InternalError.TsfApmCommonError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGCLOUDPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgCloudParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) AssociateConfigWithGroup(request *AssociateConfigWithGroupRequest) (response *AssociateConfigWithGroupResponse, err error) {
    return c.AssociateConfigWithGroupWithContext(context.Background(), request)
}

// AssociateConfigWithGroup
// 关联投递配置到部署组
//
// 可能返回的错误码:
//  INTERNALERROR_TSFAPMCOMMONERROR = "InternalError.TsfApmCommonError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGCLOUDPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgCloudParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) AssociateConfigWithGroupWithContext(ctx context.Context, request *AssociateConfigWithGroupRequest) (response *AssociateConfigWithGroupResponse, err error) {
    if request == nil {
        request = NewAssociateConfigWithGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "AssociateConfigWithGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateConfigWithGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateConfigWithGroupResponse()
    err = c.Send(request, response)
    return
}

func NewBindApiGroupRequest() (request *BindApiGroupRequest) {
    request = &BindApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "BindApiGroup")
    
    
    return
}

func NewBindApiGroupResponse() (response *BindApiGroupResponse) {
    response = &BindApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindApiGroup
// 网关与API分组批量绑定
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) BindApiGroup(request *BindApiGroupRequest) (response *BindApiGroupResponse, err error) {
    return c.BindApiGroupWithContext(context.Background(), request)
}

// BindApiGroup
// 网关与API分组批量绑定
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) BindApiGroupWithContext(ctx context.Context, request *BindApiGroupRequest) (response *BindApiGroupResponse, err error) {
    if request == nil {
        request = NewBindApiGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "BindApiGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindApiGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewBindPluginRequest() (request *BindPluginRequest) {
    request = &BindPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "BindPlugin")
    
    
    return
}

func NewBindPluginResponse() (response *BindPluginResponse) {
    response = &BindPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindPlugin
// 插件与网关分组/API批量绑定
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYCOMMONERROR = "FailedOperation.GatewayCommonError"
//  INTERNALERROR_GATEWAYCOMMONERROR = "InternalError.GatewayCommonError"
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) BindPlugin(request *BindPluginRequest) (response *BindPluginResponse, err error) {
    return c.BindPluginWithContext(context.Background(), request)
}

// BindPlugin
// 插件与网关分组/API批量绑定
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYCOMMONERROR = "FailedOperation.GatewayCommonError"
//  INTERNALERROR_GATEWAYCOMMONERROR = "InternalError.GatewayCommonError"
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) BindPluginWithContext(ctx context.Context, request *BindPluginRequest) (response *BindPluginResponse, err error) {
    if request == nil {
        request = NewBindPluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "BindPlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindPlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindPluginResponse()
    err = c.Send(request, response)
    return
}

func NewChangeApiUsableStatusRequest() (request *ChangeApiUsableStatusRequest) {
    request = &ChangeApiUsableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ChangeApiUsableStatus")
    
    
    return
}

func NewChangeApiUsableStatusResponse() (response *ChangeApiUsableStatusResponse) {
    response = &ChangeApiUsableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChangeApiUsableStatus
// 启用或禁用API
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) ChangeApiUsableStatus(request *ChangeApiUsableStatusRequest) (response *ChangeApiUsableStatusResponse, err error) {
    return c.ChangeApiUsableStatusWithContext(context.Background(), request)
}

// ChangeApiUsableStatus
// 启用或禁用API
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) ChangeApiUsableStatusWithContext(ctx context.Context, request *ChangeApiUsableStatusRequest) (response *ChangeApiUsableStatusResponse, err error) {
    if request == nil {
        request = NewChangeApiUsableStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ChangeApiUsableStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeApiUsableStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeApiUsableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewContinueRunFailedTaskBatchRequest() (request *ContinueRunFailedTaskBatchRequest) {
    request = &ContinueRunFailedTaskBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ContinueRunFailedTaskBatch")
    
    
    return
}

func NewContinueRunFailedTaskBatchResponse() (response *ContinueRunFailedTaskBatchResponse) {
    response = &ContinueRunFailedTaskBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ContinueRunFailedTaskBatch
// 对执行失败的任务批次执行续跑
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ContinueRunFailedTaskBatch(request *ContinueRunFailedTaskBatchRequest) (response *ContinueRunFailedTaskBatchResponse, err error) {
    return c.ContinueRunFailedTaskBatchWithContext(context.Background(), request)
}

// ContinueRunFailedTaskBatch
// 对执行失败的任务批次执行续跑
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ContinueRunFailedTaskBatchWithContext(ctx context.Context, request *ContinueRunFailedTaskBatchRequest) (response *ContinueRunFailedTaskBatchResponse, err error) {
    if request == nil {
        request = NewContinueRunFailedTaskBatchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ContinueRunFailedTaskBatch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ContinueRunFailedTaskBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewContinueRunFailedTaskBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAllGatewayApiAsyncRequest() (request *CreateAllGatewayApiAsyncRequest) {
    request = &CreateAllGatewayApiAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateAllGatewayApiAsync")
    
    
    return
}

func NewCreateAllGatewayApiAsyncResponse() (response *CreateAllGatewayApiAsyncResponse) {
    response = &CreateAllGatewayApiAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAllGatewayApiAsync
// 一键导入API分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) CreateAllGatewayApiAsync(request *CreateAllGatewayApiAsyncRequest) (response *CreateAllGatewayApiAsyncResponse, err error) {
    return c.CreateAllGatewayApiAsyncWithContext(context.Background(), request)
}

// CreateAllGatewayApiAsync
// 一键导入API分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) CreateAllGatewayApiAsyncWithContext(ctx context.Context, request *CreateAllGatewayApiAsyncRequest) (response *CreateAllGatewayApiAsyncResponse, err error) {
    if request == nil {
        request = NewCreateAllGatewayApiAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateAllGatewayApiAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAllGatewayApiAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAllGatewayApiAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiGroupRequest() (request *CreateApiGroupRequest) {
    request = &CreateApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApiGroup")
    
    
    return
}

func NewCreateApiGroupResponse() (response *CreateApiGroupResponse) {
    response = &CreateApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApiGroup
// 创建API分组
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateApiGroup(request *CreateApiGroupRequest) (response *CreateApiGroupResponse, err error) {
    return c.CreateApiGroupWithContext(context.Background(), request)
}

// CreateApiGroup
// 创建API分组
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateApiGroupWithContext(ctx context.Context, request *CreateApiGroupRequest) (response *CreateApiGroupResponse, err error) {
    if request == nil {
        request = NewCreateApiGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateApiGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiRateLimitRuleRequest() (request *CreateApiRateLimitRuleRequest) {
    request = &CreateApiRateLimitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApiRateLimitRule")
    
    
    return
}

func NewCreateApiRateLimitRuleResponse() (response *CreateApiRateLimitRuleResponse) {
    response = &CreateApiRateLimitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApiRateLimitRule
// 创建API限流规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateApiRateLimitRule(request *CreateApiRateLimitRuleRequest) (response *CreateApiRateLimitRuleResponse, err error) {
    return c.CreateApiRateLimitRuleWithContext(context.Background(), request)
}

// CreateApiRateLimitRule
// 创建API限流规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateApiRateLimitRuleWithContext(ctx context.Context, request *CreateApiRateLimitRuleRequest) (response *CreateApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewCreateApiRateLimitRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateApiRateLimitRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiRateLimitRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiRateLimitRuleWithDetailRespRequest() (request *CreateApiRateLimitRuleWithDetailRespRequest) {
    request = &CreateApiRateLimitRuleWithDetailRespRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApiRateLimitRuleWithDetailResp")
    
    
    return
}

func NewCreateApiRateLimitRuleWithDetailRespResponse() (response *CreateApiRateLimitRuleWithDetailRespResponse) {
    response = &CreateApiRateLimitRuleWithDetailRespResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApiRateLimitRuleWithDetailResp
// 创建API限流规则,并返回规则信息
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateApiRateLimitRuleWithDetailResp(request *CreateApiRateLimitRuleWithDetailRespRequest) (response *CreateApiRateLimitRuleWithDetailRespResponse, err error) {
    return c.CreateApiRateLimitRuleWithDetailRespWithContext(context.Background(), request)
}

// CreateApiRateLimitRuleWithDetailResp
// 创建API限流规则,并返回规则信息
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateApiRateLimitRuleWithDetailRespWithContext(ctx context.Context, request *CreateApiRateLimitRuleWithDetailRespRequest) (response *CreateApiRateLimitRuleWithDetailRespResponse, err error) {
    if request == nil {
        request = NewCreateApiRateLimitRuleWithDetailRespRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateApiRateLimitRuleWithDetailResp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiRateLimitRuleWithDetailResp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiRateLimitRuleWithDetailRespResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApplication")
    
    
    return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
    response = &CreateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplication
// 创建应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONCREATEESATUHERROR = "FailedOperation.ApplicationCreateEsAtuhError"
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  INTERNALERROR_APPLICATIONMASTERFEIGNERROR = "InternalError.ApplicationMasterFeignError"
//  INTERNALERROR_APPLICATIONMASTERNUKNOWNERROR = "InternalError.ApplicationMasterNuknownError"
//  INTERNALERROR_APPLICATIONSCALABLEINITERROR = "InternalError.ApplicationScalableInitError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INVALIDPARAMETERVALUE_APPLICATIONMICROTYPEINVALID = "InvalidParameterValue.ApplicationMicroTypeInvalid"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMEEXIST = "InvalidParameterValue.ApplicationNameExist"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMELENGTH = "InvalidParameterValue.ApplicationNameLength"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMENULL = "InvalidParameterValue.ApplicationNameNull"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMEREGXINVALID = "InvalidParameterValue.ApplicationNameRegxInvalid"
//  INVALIDPARAMETERVALUE_APPLICATIONTYPEINVALID = "InvalidParameterValue.ApplicationTypeInvalid"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  MISSINGPARAMETER_APPLICATIONTYPENULL = "MissingParameter.ApplicationTypeNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    return c.CreateApplicationWithContext(context.Background(), request)
}

// CreateApplication
// 创建应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONCREATEESATUHERROR = "FailedOperation.ApplicationCreateEsAtuhError"
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  INTERNALERROR_APPLICATIONMASTERFEIGNERROR = "InternalError.ApplicationMasterFeignError"
//  INTERNALERROR_APPLICATIONMASTERNUKNOWNERROR = "InternalError.ApplicationMasterNuknownError"
//  INTERNALERROR_APPLICATIONSCALABLEINITERROR = "InternalError.ApplicationScalableInitError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INVALIDPARAMETERVALUE_APPLICATIONMICROTYPEINVALID = "InvalidParameterValue.ApplicationMicroTypeInvalid"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMEEXIST = "InvalidParameterValue.ApplicationNameExist"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMELENGTH = "InvalidParameterValue.ApplicationNameLength"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMENULL = "InvalidParameterValue.ApplicationNameNull"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMEREGXINVALID = "InvalidParameterValue.ApplicationNameRegxInvalid"
//  INVALIDPARAMETERVALUE_APPLICATIONTYPEINVALID = "InvalidParameterValue.ApplicationTypeInvalid"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  MISSINGPARAMETER_APPLICATIONTYPENULL = "MissingParameter.ApplicationTypeNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCluster
// 创建集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_CLUSTERCREATEVPCFAIL = "FailedOperation.ClusterCreateVpcFail"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_TKECLUSTERCREATEFAILED = "FailedOperation.TkeClusterCreateFailed"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CLUSTERCOMMONERROR = "InternalError.ClusterCommonError"
//  INTERNALERROR_CLUSTERMASTERFEIGNERROR = "InternalError.ClusterMasterFeignError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CLUSTERCIDRCONFLICT = "InvalidParameterValue.ClusterCidrConflict"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEEXIST = "InvalidParameterValue.ClusterNameExist"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEPREFIXILLEGALERROR = "InvalidParameterValue.ClusterNamePrefixIllegalError"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEREQUIRED = "InvalidParameterValue.ClusterNameRequired"
//  INVALIDPARAMETERVALUE_CLUSTERREGIONINVALID = "InvalidParameterValue.ClusterRegionInvalid"
//  INVALIDPARAMETERVALUE_CLUSTERTYPEINVALID = "InvalidParameterValue.ClusterTypeInvalid"
//  INVALIDPARAMETERVALUE_CLUSTERZONEINVALID = "InvalidParameterValue.ClusterZoneInvalid"
//  LIMITEXCEEDED_TKECLUSTERNUMBEREXCEEDLIMIT = "LimitExceeded.TkeClusterNumberExceedLimit"
//  MISSINGPARAMETER_CLUSTERSUBNETREQUIRED = "MissingParameter.ClusterSubnetRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CLUSTERVPCNOTEXIST = "ResourceNotFound.ClusterVpcNotExist"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// 创建集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_CLUSTERCREATEVPCFAIL = "FailedOperation.ClusterCreateVpcFail"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_TKECLUSTERCREATEFAILED = "FailedOperation.TkeClusterCreateFailed"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CLUSTERCOMMONERROR = "InternalError.ClusterCommonError"
//  INTERNALERROR_CLUSTERMASTERFEIGNERROR = "InternalError.ClusterMasterFeignError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CLUSTERCIDRCONFLICT = "InvalidParameterValue.ClusterCidrConflict"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEEXIST = "InvalidParameterValue.ClusterNameExist"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEPREFIXILLEGALERROR = "InvalidParameterValue.ClusterNamePrefixIllegalError"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEREQUIRED = "InvalidParameterValue.ClusterNameRequired"
//  INVALIDPARAMETERVALUE_CLUSTERREGIONINVALID = "InvalidParameterValue.ClusterRegionInvalid"
//  INVALIDPARAMETERVALUE_CLUSTERTYPEINVALID = "InvalidParameterValue.ClusterTypeInvalid"
//  INVALIDPARAMETERVALUE_CLUSTERZONEINVALID = "InvalidParameterValue.ClusterZoneInvalid"
//  LIMITEXCEEDED_TKECLUSTERNUMBEREXCEEDLIMIT = "LimitExceeded.TkeClusterNumberExceedLimit"
//  MISSINGPARAMETER_CLUSTERSUBNETREQUIRED = "MissingParameter.ClusterSubnetRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CLUSTERVPCNOTEXIST = "ResourceNotFound.ClusterVpcNotExist"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigRequest() (request *CreateConfigRequest) {
    request = &CreateConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateConfig")
    
    
    return
}

func NewCreateConfigResponse() (response *CreateConfigResponse) {
    response = &CreateConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfig
// 创建配置项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGNAMEINVALID = "InvalidParameterValue.ConfigNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONDESCINVALID = "InvalidParameterValue.ConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_APPLICATIONIDREQUIRED = "MissingParameter.ApplicationIdRequired"
//  MISSINGPARAMETER_CONFIGNAMEREQUIRED = "MissingParameter.ConfigNameRequired"
//  MISSINGPARAMETER_CONFIGTYPEREQUIRED = "MissingParameter.ConfigTypeRequired"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  MISSINGPARAMETER_CONFIGVERSIONREQUIRED = "MissingParameter.ConfigVersionRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateConfig(request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
    return c.CreateConfigWithContext(context.Background(), request)
}

// CreateConfig
// 创建配置项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGNAMEINVALID = "InvalidParameterValue.ConfigNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONDESCINVALID = "InvalidParameterValue.ConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_APPLICATIONIDREQUIRED = "MissingParameter.ApplicationIdRequired"
//  MISSINGPARAMETER_CONFIGNAMEREQUIRED = "MissingParameter.ConfigNameRequired"
//  MISSINGPARAMETER_CONFIGTYPEREQUIRED = "MissingParameter.ConfigTypeRequired"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  MISSINGPARAMETER_CONFIGVERSIONREQUIRED = "MissingParameter.ConfigVersionRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateConfigWithContext(ctx context.Context, request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
    if request == nil {
        request = NewCreateConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigTemplateRequest() (request *CreateConfigTemplateRequest) {
    request = &CreateConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateConfigTemplate")
    
    
    return
}

func NewCreateConfigTemplateResponse() (response *CreateConfigTemplateResponse) {
    response = &CreateConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfigTemplate
// 创建参数模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"
//  FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"
//  INVALIDPARAMETER_CONFIGTEMPLATENAMEINVALID = "InvalidParameter.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMEINVALID = "InvalidParameterValue.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"
func (c *Client) CreateConfigTemplate(request *CreateConfigTemplateRequest) (response *CreateConfigTemplateResponse, err error) {
    return c.CreateConfigTemplateWithContext(context.Background(), request)
}

// CreateConfigTemplate
// 创建参数模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"
//  FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"
//  INVALIDPARAMETER_CONFIGTEMPLATENAMEINVALID = "InvalidParameter.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMEINVALID = "InvalidParameterValue.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"
func (c *Client) CreateConfigTemplateWithContext(ctx context.Context, request *CreateConfigTemplateRequest) (response *CreateConfigTemplateResponse, err error) {
    if request == nil {
        request = NewCreateConfigTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateConfigTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfigTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigTemplateWithDetailRespRequest() (request *CreateConfigTemplateWithDetailRespRequest) {
    request = &CreateConfigTemplateWithDetailRespRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateConfigTemplateWithDetailResp")
    
    
    return
}

func NewCreateConfigTemplateWithDetailRespResponse() (response *CreateConfigTemplateWithDetailRespResponse) {
    response = &CreateConfigTemplateWithDetailRespResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfigTemplateWithDetailResp
// 创建参数模板，并返回模板详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"
//  FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"
//  INVALIDPARAMETER_CONFIGTEMPLATENAMEINVALID = "InvalidParameter.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMEINVALID = "InvalidParameterValue.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"
func (c *Client) CreateConfigTemplateWithDetailResp(request *CreateConfigTemplateWithDetailRespRequest) (response *CreateConfigTemplateWithDetailRespResponse, err error) {
    return c.CreateConfigTemplateWithDetailRespWithContext(context.Background(), request)
}

// CreateConfigTemplateWithDetailResp
// 创建参数模板，并返回模板详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"
//  FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"
//  INVALIDPARAMETER_CONFIGTEMPLATENAMEINVALID = "InvalidParameter.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMEINVALID = "InvalidParameterValue.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"
func (c *Client) CreateConfigTemplateWithDetailRespWithContext(ctx context.Context, request *CreateConfigTemplateWithDetailRespRequest) (response *CreateConfigTemplateWithDetailRespResponse, err error) {
    if request == nil {
        request = NewCreateConfigTemplateWithDetailRespRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateConfigTemplateWithDetailResp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfigTemplateWithDetailResp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigTemplateWithDetailRespResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigWithDetailRespRequest() (request *CreateConfigWithDetailRespRequest) {
    request = &CreateConfigWithDetailRespRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateConfigWithDetailResp")
    
    
    return
}

func NewCreateConfigWithDetailRespResponse() (response *CreateConfigWithDetailRespResponse) {
    response = &CreateConfigWithDetailRespResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfigWithDetailResp
// 创建配置项，返回详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGNAMEINVALID = "InvalidParameterValue.ConfigNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONDESCINVALID = "InvalidParameterValue.ConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_APPLICATIONIDREQUIRED = "MissingParameter.ApplicationIdRequired"
//  MISSINGPARAMETER_CONFIGNAMEREQUIRED = "MissingParameter.ConfigNameRequired"
//  MISSINGPARAMETER_CONFIGTYPEREQUIRED = "MissingParameter.ConfigTypeRequired"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  MISSINGPARAMETER_CONFIGVERSIONREQUIRED = "MissingParameter.ConfigVersionRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateConfigWithDetailResp(request *CreateConfigWithDetailRespRequest) (response *CreateConfigWithDetailRespResponse, err error) {
    return c.CreateConfigWithDetailRespWithContext(context.Background(), request)
}

// CreateConfigWithDetailResp
// 创建配置项，返回详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGNAMEINVALID = "InvalidParameterValue.ConfigNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONDESCINVALID = "InvalidParameterValue.ConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_APPLICATIONIDREQUIRED = "MissingParameter.ApplicationIdRequired"
//  MISSINGPARAMETER_CONFIGNAMEREQUIRED = "MissingParameter.ConfigNameRequired"
//  MISSINGPARAMETER_CONFIGTYPEREQUIRED = "MissingParameter.ConfigTypeRequired"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  MISSINGPARAMETER_CONFIGVERSIONREQUIRED = "MissingParameter.ConfigVersionRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateConfigWithDetailRespWithContext(ctx context.Context, request *CreateConfigWithDetailRespRequest) (response *CreateConfigWithDetailRespResponse, err error) {
    if request == nil {
        request = NewCreateConfigWithDetailRespRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateConfigWithDetailResp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfigWithDetailResp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigWithDetailRespResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContainGroupRequest() (request *CreateContainGroupRequest) {
    request = &CreateContainGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateContainGroup")
    
    
    return
}

func NewCreateContainGroupResponse() (response *CreateContainGroupResponse) {
    response = &CreateContainGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateContainGroup
// （已废弃，请使用 CreateGroup 和 DeployContainerGroup 创建和部署容器部署组）创建容器部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_GROUPMASTERNUKNOWNERROR = "InternalError.GroupMasterNuknownError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMELEGNTH = "InvalidParameterValue.ContainergroupGroupnameLegnth"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMENULL = "InvalidParameterValue.ContainergroupGroupnameNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMEREGEXMATCHFALSE = "InvalidParameterValue.ContainergroupGroupnameRegexMatchFalse"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTINVALID = "InvalidParameterValue.ContainergroupPortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLINVALID = "InvalidParameterValue.ContainergroupProtocolInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLNULL = "InvalidParameterValue.ContainergroupProtocolNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"
//  INVALIDPARAMETERVALUE_GROUPCLUSTERNAMESPACENOTBOUND = "InvalidParameterValue.GroupClusterNamespaceNotBound"
//  INVALIDPARAMETERVALUE_GROUPCLUSTERTYPEMISMATCH = "InvalidParameterValue.GroupClusterTypeMismatch"
//  INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"
//  RESOURCENOTFOUND_CONTAINERGROUPCLUSTERNOTFOUND = "ResourceNotFound.ContainergroupClusterNotfound"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"
func (c *Client) CreateContainGroup(request *CreateContainGroupRequest) (response *CreateContainGroupResponse, err error) {
    return c.CreateContainGroupWithContext(context.Background(), request)
}

// CreateContainGroup
// （已废弃，请使用 CreateGroup 和 DeployContainerGroup 创建和部署容器部署组）创建容器部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_GROUPMASTERNUKNOWNERROR = "InternalError.GroupMasterNuknownError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMELEGNTH = "InvalidParameterValue.ContainergroupGroupnameLegnth"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMENULL = "InvalidParameterValue.ContainergroupGroupnameNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMEREGEXMATCHFALSE = "InvalidParameterValue.ContainergroupGroupnameRegexMatchFalse"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTINVALID = "InvalidParameterValue.ContainergroupPortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLINVALID = "InvalidParameterValue.ContainergroupProtocolInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLNULL = "InvalidParameterValue.ContainergroupProtocolNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"
//  INVALIDPARAMETERVALUE_GROUPCLUSTERNAMESPACENOTBOUND = "InvalidParameterValue.GroupClusterNamespaceNotBound"
//  INVALIDPARAMETERVALUE_GROUPCLUSTERTYPEMISMATCH = "InvalidParameterValue.GroupClusterTypeMismatch"
//  INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"
//  RESOURCENOTFOUND_CONTAINERGROUPCLUSTERNOTFOUND = "ResourceNotFound.ContainergroupClusterNotfound"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"
func (c *Client) CreateContainGroupWithContext(ctx context.Context, request *CreateContainGroupRequest) (response *CreateContainGroupResponse, err error) {
    if request == nil {
        request = NewCreateContainGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateContainGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateContainGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateContainGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileConfigRequest() (request *CreateFileConfigRequest) {
    request = &CreateFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateFileConfig")
    
    
    return
}

func NewCreateFileConfigResponse() (response *CreateFileConfigResponse) {
    response = &CreateFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFileConfig
// 创建文件配置项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_FILECONFIGCODEUNSUPPORTED = "InvalidParameterValue.FileConfigCodeUnsupported"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTS = "InvalidParameterValue.FileConfigExists"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATH = "InvalidParameterValue.FileConfigExistsPath"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATHOTHER = "InvalidParameterValue.FileConfigExistsPathOther"
//  INVALIDPARAMETERVALUE_FILECONFIGFILENAMEINVALID = "InvalidParameterValue.FileConfigFileNameInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGFILEPATHINVALID = "InvalidParameterValue.FileConfigFilePathInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGFIXEDFIELDUNCHANGEABLE = "InvalidParameterValue.FileConfigFixedFieldUnchangeable"
//  INVALIDPARAMETERVALUE_FILECONFIGNAMEINVALID = "InvalidParameterValue.FileConfigNameInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGVERSIONDESCINVALID = "InvalidParameterValue.FileConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGVERSIONINVALID = "InvalidParameterValue.FileConfigVersionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_APPLICATIONIDREQUIRED = "MissingParameter.ApplicationIdRequired"
//  MISSINGPARAMETER_FILECONFIGFILEVALUEREQUIRED = "MissingParameter.FileConfigFileValueRequired"
func (c *Client) CreateFileConfig(request *CreateFileConfigRequest) (response *CreateFileConfigResponse, err error) {
    return c.CreateFileConfigWithContext(context.Background(), request)
}

// CreateFileConfig
// 创建文件配置项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_FILECONFIGCODEUNSUPPORTED = "InvalidParameterValue.FileConfigCodeUnsupported"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTS = "InvalidParameterValue.FileConfigExists"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATH = "InvalidParameterValue.FileConfigExistsPath"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATHOTHER = "InvalidParameterValue.FileConfigExistsPathOther"
//  INVALIDPARAMETERVALUE_FILECONFIGFILENAMEINVALID = "InvalidParameterValue.FileConfigFileNameInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGFILEPATHINVALID = "InvalidParameterValue.FileConfigFilePathInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGFIXEDFIELDUNCHANGEABLE = "InvalidParameterValue.FileConfigFixedFieldUnchangeable"
//  INVALIDPARAMETERVALUE_FILECONFIGNAMEINVALID = "InvalidParameterValue.FileConfigNameInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGVERSIONDESCINVALID = "InvalidParameterValue.FileConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGVERSIONINVALID = "InvalidParameterValue.FileConfigVersionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_APPLICATIONIDREQUIRED = "MissingParameter.ApplicationIdRequired"
//  MISSINGPARAMETER_FILECONFIGFILEVALUEREQUIRED = "MissingParameter.FileConfigFileValueRequired"
func (c *Client) CreateFileConfigWithContext(ctx context.Context, request *CreateFileConfigRequest) (response *CreateFileConfigResponse, err error) {
    if request == nil {
        request = NewCreateFileConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateFileConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileConfigWithDetailRespRequest() (request *CreateFileConfigWithDetailRespRequest) {
    request = &CreateFileConfigWithDetailRespRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateFileConfigWithDetailResp")
    
    
    return
}

func NewCreateFileConfigWithDetailRespResponse() (response *CreateFileConfigWithDetailRespResponse) {
    response = &CreateFileConfigWithDetailRespResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFileConfigWithDetailResp
// 创建文件配置项，返回详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTS = "InvalidParameterValue.FileConfigExists"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATH = "InvalidParameterValue.FileConfigExistsPath"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATHOTHER = "InvalidParameterValue.FileConfigExistsPathOther"
//  INVALIDPARAMETERVALUE_FILECONFIGFILENAMEINVALID = "InvalidParameterValue.FileConfigFileNameInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGFILEPATHINVALID = "InvalidParameterValue.FileConfigFilePathInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGNAMEINVALID = "InvalidParameterValue.FileConfigNameInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGVERSIONDESCINVALID = "InvalidParameterValue.FileConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGVERSIONINVALID = "InvalidParameterValue.FileConfigVersionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_FILECONFIGFILEVALUEREQUIRED = "MissingParameter.FileConfigFileValueRequired"
func (c *Client) CreateFileConfigWithDetailResp(request *CreateFileConfigWithDetailRespRequest) (response *CreateFileConfigWithDetailRespResponse, err error) {
    return c.CreateFileConfigWithDetailRespWithContext(context.Background(), request)
}

// CreateFileConfigWithDetailResp
// 创建文件配置项，返回详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTS = "InvalidParameterValue.FileConfigExists"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATH = "InvalidParameterValue.FileConfigExistsPath"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATHOTHER = "InvalidParameterValue.FileConfigExistsPathOther"
//  INVALIDPARAMETERVALUE_FILECONFIGFILENAMEINVALID = "InvalidParameterValue.FileConfigFileNameInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGFILEPATHINVALID = "InvalidParameterValue.FileConfigFilePathInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGNAMEINVALID = "InvalidParameterValue.FileConfigNameInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGVERSIONDESCINVALID = "InvalidParameterValue.FileConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGVERSIONINVALID = "InvalidParameterValue.FileConfigVersionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_FILECONFIGFILEVALUEREQUIRED = "MissingParameter.FileConfigFileValueRequired"
func (c *Client) CreateFileConfigWithDetailRespWithContext(ctx context.Context, request *CreateFileConfigWithDetailRespRequest) (response *CreateFileConfigWithDetailRespResponse, err error) {
    if request == nil {
        request = NewCreateFileConfigWithDetailRespRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateFileConfigWithDetailResp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileConfigWithDetailResp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileConfigWithDetailRespResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGatewayApiRequest() (request *CreateGatewayApiRequest) {
    request = &CreateGatewayApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayApi")
    
    
    return
}

func NewCreateGatewayApiResponse() (response *CreateGatewayApiResponse) {
    response = &CreateGatewayApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGatewayApi
// 批量导入API至api分组(也支持新建API到分组)
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNSUPPORTEDOPERATION_GATEWAYTOOMANYREQUESTPARAMETER = "UnsupportedOperation.GatewayTooManyRequestParameter"
func (c *Client) CreateGatewayApi(request *CreateGatewayApiRequest) (response *CreateGatewayApiResponse, err error) {
    return c.CreateGatewayApiWithContext(context.Background(), request)
}

// CreateGatewayApi
// 批量导入API至api分组(也支持新建API到分组)
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNSUPPORTEDOPERATION_GATEWAYTOOMANYREQUESTPARAMETER = "UnsupportedOperation.GatewayTooManyRequestParameter"
func (c *Client) CreateGatewayApiWithContext(ctx context.Context, request *CreateGatewayApiRequest) (response *CreateGatewayApiResponse, err error) {
    if request == nil {
        request = NewCreateGatewayApiRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateGatewayApi")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGatewayApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGatewayApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGroup")
    
    
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGroup
// 创建虚拟机部署组
//
// 可能返回的错误码:
//  INTERNALERROR_GROUPMASTERNUKNOWNERROR = "InternalError.GroupMasterNuknownError"
//  INVALIDPARAMETERVALUE_GROUPCLUSTERNAMESPACENOTBOUND = "InvalidParameterValue.GroupClusterNamespaceNotBound"
//  INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"
//  INVALIDPARAMETERVALUE_GROUPNAMELENGTH = "InvalidParameterValue.GroupNameLength"
//  INVALIDPARAMETERVALUE_GROUPNAMEREGXMISMATCH = "InvalidParameterValue.GroupNameRegxMismatch"
//  MISSINGPARAMETER_GROUPAPPLICATIONNULL = "MissingParameter.GroupApplicationNull"
//  MISSINGPARAMETER_GROUPNAMESPACENULL = "MissingParameter.GroupNamespaceNull"
//  RESOURCENOTFOUND_CONTAINERGROUPCLUSTERNOTFOUND = "ResourceNotFound.ContainergroupClusterNotfound"
//  RESOURCENOTFOUND_CVMCAEMASTERRESOURCENOTFOUND = "ResourceNotFound.CvmcaeMasterResourceNotFound"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    return c.CreateGroupWithContext(context.Background(), request)
}

// CreateGroup
// 创建虚拟机部署组
//
// 可能返回的错误码:
//  INTERNALERROR_GROUPMASTERNUKNOWNERROR = "InternalError.GroupMasterNuknownError"
//  INVALIDPARAMETERVALUE_GROUPCLUSTERNAMESPACENOTBOUND = "InvalidParameterValue.GroupClusterNamespaceNotBound"
//  INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"
//  INVALIDPARAMETERVALUE_GROUPNAMELENGTH = "InvalidParameterValue.GroupNameLength"
//  INVALIDPARAMETERVALUE_GROUPNAMEREGXMISMATCH = "InvalidParameterValue.GroupNameRegxMismatch"
//  MISSINGPARAMETER_GROUPAPPLICATIONNULL = "MissingParameter.GroupApplicationNull"
//  MISSINGPARAMETER_GROUPNAMESPACENULL = "MissingParameter.GroupNamespaceNull"
//  RESOURCENOTFOUND_CONTAINERGROUPCLUSTERNOTFOUND = "ResourceNotFound.ContainergroupClusterNotfound"
//  RESOURCENOTFOUND_CVMCAEMASTERRESOURCENOTFOUND = "ResourceNotFound.CvmcaeMasterResourceNotFound"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLaneRequest() (request *CreateLaneRequest) {
    request = &CreateLaneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateLane")
    
    
    return
}

func NewCreateLaneResponse() (response *CreateLaneResponse) {
    response = &CreateLaneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLane
// 创建泳道
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANEINFORELEASEMESHFAILED = "FailedOperation.LaneInfoReleaseMeshFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateLane(request *CreateLaneRequest) (response *CreateLaneResponse, err error) {
    return c.CreateLaneWithContext(context.Background(), request)
}

// CreateLane
// 创建泳道
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANEINFORELEASEMESHFAILED = "FailedOperation.LaneInfoReleaseMeshFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateLaneWithContext(ctx context.Context, request *CreateLaneRequest) (response *CreateLaneResponse, err error) {
    if request == nil {
        request = NewCreateLaneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateLane")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLane require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLaneResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLaneRuleRequest() (request *CreateLaneRuleRequest) {
    request = &CreateLaneRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateLaneRule")
    
    
    return
}

func NewCreateLaneRuleResponse() (response *CreateLaneRuleResponse) {
    response = &CreateLaneRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLaneRule
// 创建泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGOPERATORINVALID = "InvalidParameterValue.LaneRuleTagOperatorInvalid"
//  INVALIDPARAMETERVALUE_LANERULETAGOPERATORNOTEMPTY = "InvalidParameterValue.LaneRuleTagOperatorNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUENOTEMPTY = "InvalidParameterValue.LaneRuleTagValueNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateLaneRule(request *CreateLaneRuleRequest) (response *CreateLaneRuleResponse, err error) {
    return c.CreateLaneRuleWithContext(context.Background(), request)
}

// CreateLaneRule
// 创建泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGOPERATORINVALID = "InvalidParameterValue.LaneRuleTagOperatorInvalid"
//  INVALIDPARAMETERVALUE_LANERULETAGOPERATORNOTEMPTY = "InvalidParameterValue.LaneRuleTagOperatorNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUENOTEMPTY = "InvalidParameterValue.LaneRuleTagValueNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateLaneRuleWithContext(ctx context.Context, request *CreateLaneRuleRequest) (response *CreateLaneRuleResponse, err error) {
    if request == nil {
        request = NewCreateLaneRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateLaneRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLaneRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLaneRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMicroserviceRequest() (request *CreateMicroserviceRequest) {
    request = &CreateMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateMicroservice")
    
    
    return
}

func NewCreateMicroserviceResponse() (response *CreateMicroserviceResponse) {
    response = &CreateMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMicroservice
// 新增微服务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_SERVICEDESCLENGTH = "InvalidParameterValue.ServiceDescLength"
//  INVALIDPARAMETERVALUE_SERVICENAMEREPEATED = "InvalidParameterValue.ServiceNameRepeated"
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateMicroservice(request *CreateMicroserviceRequest) (response *CreateMicroserviceResponse, err error) {
    return c.CreateMicroserviceWithContext(context.Background(), request)
}

// CreateMicroservice
// 新增微服务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_SERVICEDESCLENGTH = "InvalidParameterValue.ServiceDescLength"
//  INVALIDPARAMETERVALUE_SERVICENAMEREPEATED = "InvalidParameterValue.ServiceNameRepeated"
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateMicroserviceWithContext(ctx context.Context, request *CreateMicroserviceRequest) (response *CreateMicroserviceResponse, err error) {
    if request == nil {
        request = NewCreateMicroserviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateMicroservice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMicroservice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMicroserviceWithDetailRespRequest() (request *CreateMicroserviceWithDetailRespRequest) {
    request = &CreateMicroserviceWithDetailRespRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateMicroserviceWithDetailResp")
    
    
    return
}

func NewCreateMicroserviceWithDetailRespResponse() (response *CreateMicroserviceWithDetailRespResponse) {
    response = &CreateMicroserviceWithDetailRespResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMicroserviceWithDetailResp
// 新增微服务返回id
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_SERVICENAMEREPEATED = "InvalidParameterValue.ServiceNameRepeated"
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
func (c *Client) CreateMicroserviceWithDetailResp(request *CreateMicroserviceWithDetailRespRequest) (response *CreateMicroserviceWithDetailRespResponse, err error) {
    return c.CreateMicroserviceWithDetailRespWithContext(context.Background(), request)
}

// CreateMicroserviceWithDetailResp
// 新增微服务返回id
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_SERVICENAMEREPEATED = "InvalidParameterValue.ServiceNameRepeated"
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
func (c *Client) CreateMicroserviceWithDetailRespWithContext(ctx context.Context, request *CreateMicroserviceWithDetailRespRequest) (response *CreateMicroserviceWithDetailRespResponse, err error) {
    if request == nil {
        request = NewCreateMicroserviceWithDetailRespRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateMicroserviceWithDetailResp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMicroserviceWithDetailResp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMicroserviceWithDetailRespResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateNamespace")
    
    
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNamespace
// 创建命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION_NAMESPACECREATEFAILED = "FailedOperation.NamespaceCreateFailed"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_KUBERNETESAPICREATENAMESPACESERROR = "InternalError.KubernetesApiCreateNamespacesError"
//  INTERNALERROR_KUBERNETESAPICREATESECRETERROR = "InternalError.KubernetesApiCreateSecretError"
//  INVALIDPARAMETERVALUE_GLOBALNAMESPACENAMEEXIST = "InvalidParameterValue.GlobalNamespaceNameExist"
//  INVALIDPARAMETERVALUE_NAMESPACEALREADYBINDCLUSTER = "InvalidParameterValue.NamespaceAlreadyBindCluster"
//  INVALIDPARAMETERVALUE_NAMESPACEDESCINVALID = "InvalidParameterValue.NamespaceDescInvalid"
//  INVALIDPARAMETERVALUE_NAMESPACENAMEEXIST = "InvalidParameterValue.NamespaceNameExist"
//  INVALIDPARAMETERVALUE_NAMESPACENAMEINVALID = "InvalidParameterValue.NamespaceNameInvalid"
//  MISSINGPARAMETER_NAMESPACENAMEREQUIRED = "MissingParameter.NamespaceNameRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    return c.CreateNamespaceWithContext(context.Background(), request)
}

// CreateNamespace
// 创建命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION_NAMESPACECREATEFAILED = "FailedOperation.NamespaceCreateFailed"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_KUBERNETESAPICREATENAMESPACESERROR = "InternalError.KubernetesApiCreateNamespacesError"
//  INTERNALERROR_KUBERNETESAPICREATESECRETERROR = "InternalError.KubernetesApiCreateSecretError"
//  INVALIDPARAMETERVALUE_GLOBALNAMESPACENAMEEXIST = "InvalidParameterValue.GlobalNamespaceNameExist"
//  INVALIDPARAMETERVALUE_NAMESPACEALREADYBINDCLUSTER = "InvalidParameterValue.NamespaceAlreadyBindCluster"
//  INVALIDPARAMETERVALUE_NAMESPACEDESCINVALID = "InvalidParameterValue.NamespaceDescInvalid"
//  INVALIDPARAMETERVALUE_NAMESPACENAMEEXIST = "InvalidParameterValue.NamespaceNameExist"
//  INVALIDPARAMETERVALUE_NAMESPACENAMEINVALID = "InvalidParameterValue.NamespaceNameInvalid"
//  MISSINGPARAMETER_NAMESPACENAMEREQUIRED = "MissingParameter.NamespaceNameRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateNamespaceWithContext(ctx context.Context, request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePathRewritesRequest() (request *CreatePathRewritesRequest) {
    request = &CreatePathRewritesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreatePathRewrites")
    
    
    return
}

func NewCreatePathRewritesResponse() (response *CreatePathRewritesResponse) {
    response = &CreatePathRewritesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePathRewrites
// 创建路径重写
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreatePathRewrites(request *CreatePathRewritesRequest) (response *CreatePathRewritesResponse, err error) {
    return c.CreatePathRewritesWithContext(context.Background(), request)
}

// CreatePathRewrites
// 创建路径重写
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreatePathRewritesWithContext(ctx context.Context, request *CreatePathRewritesRequest) (response *CreatePathRewritesResponse, err error) {
    if request == nil {
        request = NewCreatePathRewritesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreatePathRewrites")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePathRewrites require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePathRewritesResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePathRewritesWithDetailRespRequest() (request *CreatePathRewritesWithDetailRespRequest) {
    request = &CreatePathRewritesWithDetailRespRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreatePathRewritesWithDetailResp")
    
    
    return
}

func NewCreatePathRewritesWithDetailRespResponse() (response *CreatePathRewritesWithDetailRespResponse) {
    response = &CreatePathRewritesWithDetailRespResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePathRewritesWithDetailResp
// 创建路径重写，并返回路径重写规则信息
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreatePathRewritesWithDetailResp(request *CreatePathRewritesWithDetailRespRequest) (response *CreatePathRewritesWithDetailRespResponse, err error) {
    return c.CreatePathRewritesWithDetailRespWithContext(context.Background(), request)
}

// CreatePathRewritesWithDetailResp
// 创建路径重写，并返回路径重写规则信息
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreatePathRewritesWithDetailRespWithContext(ctx context.Context, request *CreatePathRewritesWithDetailRespRequest) (response *CreatePathRewritesWithDetailRespResponse, err error) {
    if request == nil {
        request = NewCreatePathRewritesWithDetailRespRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreatePathRewritesWithDetailResp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePathRewritesWithDetailResp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePathRewritesWithDetailRespResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProgramRequest() (request *CreateProgramRequest) {
    request = &CreateProgramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateProgram")
    
    
    return
}

func NewCreateProgramResponse() (response *CreateProgramResponse) {
    response = &CreateProgramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProgram
// 创建数据集
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DUPLICATEPROGRAMNAME = "InvalidParameterValue.DuplicateProgramName"
//  INVALIDPARAMETERVALUE_INVALIDPROGRAMNAME = "InvalidParameterValue.InvalidProgramName"
func (c *Client) CreateProgram(request *CreateProgramRequest) (response *CreateProgramResponse, err error) {
    return c.CreateProgramWithContext(context.Background(), request)
}

// CreateProgram
// 创建数据集
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DUPLICATEPROGRAMNAME = "InvalidParameterValue.DuplicateProgramName"
//  INVALIDPARAMETERVALUE_INVALIDPROGRAMNAME = "InvalidParameterValue.InvalidProgramName"
func (c *Client) CreateProgramWithContext(ctx context.Context, request *CreateProgramRequest) (response *CreateProgramResponse, err error) {
    if request == nil {
        request = NewCreateProgramRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateProgram")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProgram require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProgramResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublicConfigRequest() (request *CreatePublicConfigRequest) {
    request = &CreatePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreatePublicConfig")
    
    
    return
}

func NewCreatePublicConfigResponse() (response *CreatePublicConfigResponse) {
    response = &CreatePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePublicConfig
// 创建公共配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreatePublicConfig(request *CreatePublicConfigRequest) (response *CreatePublicConfigResponse, err error) {
    return c.CreatePublicConfigWithContext(context.Background(), request)
}

// CreatePublicConfig
// 创建公共配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreatePublicConfigWithContext(ctx context.Context, request *CreatePublicConfigRequest) (response *CreatePublicConfigResponse, err error) {
    if request == nil {
        request = NewCreatePublicConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreatePublicConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePublicConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublicConfigWithDetailRespRequest() (request *CreatePublicConfigWithDetailRespRequest) {
    request = &CreatePublicConfigWithDetailRespRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreatePublicConfigWithDetailResp")
    
    
    return
}

func NewCreatePublicConfigWithDetailRespResponse() (response *CreatePublicConfigWithDetailRespResponse) {
    response = &CreatePublicConfigWithDetailRespResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePublicConfigWithDetailResp
// 创建公共配置项，并返回配置项详细信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreatePublicConfigWithDetailResp(request *CreatePublicConfigWithDetailRespRequest) (response *CreatePublicConfigWithDetailRespResponse, err error) {
    return c.CreatePublicConfigWithDetailRespWithContext(context.Background(), request)
}

// CreatePublicConfigWithDetailResp
// 创建公共配置项，并返回配置项详细信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreatePublicConfigWithDetailRespWithContext(ctx context.Context, request *CreatePublicConfigWithDetailRespRequest) (response *CreatePublicConfigWithDetailRespResponse, err error) {
    if request == nil {
        request = NewCreatePublicConfigWithDetailRespRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreatePublicConfigWithDetailResp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePublicConfigWithDetailResp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePublicConfigWithDetailRespResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRepositoryRequest() (request *CreateRepositoryRequest) {
    request = &CreateRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateRepository")
    
    
    return
}

func NewCreateRepositoryResponse() (response *CreateRepositoryResponse) {
    response = &CreateRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRepository
// 创建仓库
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINUSE_OBJECTEXIST = "ResourceInUse.ObjectExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) CreateRepository(request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
    return c.CreateRepositoryWithContext(context.Background(), request)
}

// CreateRepository
// 创建仓库
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINUSE_OBJECTEXIST = "ResourceInUse.ObjectExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) CreateRepositoryWithContext(ctx context.Context, request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
    if request == nil {
        request = NewCreateRepositoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateRepository")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTask
// 创建任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    return c.CreateTaskWithContext(context.Background(), request)
}

// CreateTask
// 创建任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskFlowRequest() (request *CreateTaskFlowRequest) {
    request = &CreateTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateTaskFlow")
    
    
    return
}

func NewCreateTaskFlowResponse() (response *CreateTaskFlowResponse) {
    response = &CreateTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTaskFlow
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateTaskFlow(request *CreateTaskFlowRequest) (response *CreateTaskFlowResponse, err error) {
    return c.CreateTaskFlowWithContext(context.Background(), request)
}

// CreateTaskFlow
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateTaskFlowWithContext(ctx context.Context, request *CreateTaskFlowRequest) (response *CreateTaskFlowResponse, err error) {
    if request == nil {
        request = NewCreateTaskFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateTaskFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUnitNamespacesRequest() (request *CreateUnitNamespacesRequest) {
    request = &CreateUnitNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateUnitNamespaces")
    
    
    return
}

func NewCreateUnitNamespacesResponse() (response *CreateUnitNamespacesResponse) {
    response = &CreateUnitNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUnitNamespaces
// 批量创建单元化命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreateUnitNamespaces(request *CreateUnitNamespacesRequest) (response *CreateUnitNamespacesResponse, err error) {
    return c.CreateUnitNamespacesWithContext(context.Background(), request)
}

// CreateUnitNamespaces
// 批量创建单元化命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreateUnitNamespacesWithContext(ctx context.Context, request *CreateUnitNamespacesRequest) (response *CreateUnitNamespacesResponse, err error) {
    if request == nil {
        request = NewCreateUnitNamespacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateUnitNamespaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUnitNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUnitNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUnitRuleRequest() (request *CreateUnitRuleRequest) {
    request = &CreateUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateUnitRule")
    
    
    return
}

func NewCreateUnitRuleResponse() (response *CreateUnitRuleResponse) {
    response = &CreateUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUnitRule
// 创建单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreateUnitRule(request *CreateUnitRuleRequest) (response *CreateUnitRuleResponse, err error) {
    return c.CreateUnitRuleWithContext(context.Background(), request)
}

// CreateUnitRule
// 创建单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreateUnitRuleWithContext(ctx context.Context, request *CreateUnitRuleRequest) (response *CreateUnitRuleResponse, err error) {
    if request == nil {
        request = NewCreateUnitRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateUnitRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUnitRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUnitRuleWithDetailRespRequest() (request *CreateUnitRuleWithDetailRespRequest) {
    request = &CreateUnitRuleWithDetailRespRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "CreateUnitRuleWithDetailResp")
    
    
    return
}

func NewCreateUnitRuleWithDetailRespResponse() (response *CreateUnitRuleWithDetailRespResponse) {
    response = &CreateUnitRuleWithDetailRespResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUnitRuleWithDetailResp
// 创建单元化规则, 并返回详细信息
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreateUnitRuleWithDetailResp(request *CreateUnitRuleWithDetailRespRequest) (response *CreateUnitRuleWithDetailRespResponse, err error) {
    return c.CreateUnitRuleWithDetailRespWithContext(context.Background(), request)
}

// CreateUnitRuleWithDetailResp
// 创建单元化规则, 并返回详细信息
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreateUnitRuleWithDetailRespWithContext(ctx context.Context, request *CreateUnitRuleWithDetailRespRequest) (response *CreateUnitRuleWithDetailRespResponse, err error) {
    if request == nil {
        request = NewCreateUnitRuleWithDetailRespRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "CreateUnitRuleWithDetailResp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUnitRuleWithDetailResp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUnitRuleWithDetailRespResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiGroupRequest() (request *DeleteApiGroupRequest) {
    request = &DeleteApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteApiGroup")
    
    
    return
}

func NewDeleteApiGroupResponse() (response *DeleteApiGroupResponse) {
    response = &DeleteApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApiGroup
// 删除Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) DeleteApiGroup(request *DeleteApiGroupRequest) (response *DeleteApiGroupResponse, err error) {
    return c.DeleteApiGroupWithContext(context.Background(), request)
}

// DeleteApiGroup
// 删除Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) DeleteApiGroupWithContext(ctx context.Context, request *DeleteApiGroupRequest) (response *DeleteApiGroupResponse, err error) {
    if request == nil {
        request = NewDeleteApiGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteApiGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApiGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiRateLimitRuleRequest() (request *DeleteApiRateLimitRuleRequest) {
    request = &DeleteApiRateLimitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteApiRateLimitRule")
    
    
    return
}

func NewDeleteApiRateLimitRuleResponse() (response *DeleteApiRateLimitRuleResponse) {
    response = &DeleteApiRateLimitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApiRateLimitRule
// 删除API限流规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERERROR = "InvalidParameterValue.GatewayParameterError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  RESOURCEINUSE_RATELIMITRULEEXISTERROR = "ResourceInUse.RatelimitRuleExistError"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) DeleteApiRateLimitRule(request *DeleteApiRateLimitRuleRequest) (response *DeleteApiRateLimitRuleResponse, err error) {
    return c.DeleteApiRateLimitRuleWithContext(context.Background(), request)
}

// DeleteApiRateLimitRule
// 删除API限流规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERERROR = "InvalidParameterValue.GatewayParameterError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  RESOURCEINUSE_RATELIMITRULEEXISTERROR = "ResourceInUse.RatelimitRuleExistError"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) DeleteApiRateLimitRuleWithContext(ctx context.Context, request *DeleteApiRateLimitRuleRequest) (response *DeleteApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewDeleteApiRateLimitRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteApiRateLimitRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApiRateLimitRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
    request = &DeleteApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteApplication")
    
    
    return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
    response = &DeleteApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplication
// 删除应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_IMAGEVALIDATEISUSED = "FailedOperation.ImageValidateIsUsed"
//  FAILEDOPERATION_IMAGEREPOREJECTDELERROR = "FailedOperation.ImagerepoRejectDelError"
//  INTERNALERROR_APPLICATIONREPODELETEPKG = "InternalError.ApplicationRepoDeletePkg"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_REMOTESERVICECALLERROR = "InternalError.RemoteServiceCallError"
//  RESOURCEINUSE_APPLICATIONCANNOTDELETE = "ResourceInUse.ApplicationCannotDelete"
//  RESOURCEINUSE_CVMCAEMASTERCANNOTDELETE = "ResourceInUse.CvmcaeMasterCannotDelete"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    return c.DeleteApplicationWithContext(context.Background(), request)
}

// DeleteApplication
// 删除应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_IMAGEVALIDATEISUSED = "FailedOperation.ImageValidateIsUsed"
//  FAILEDOPERATION_IMAGEREPOREJECTDELERROR = "FailedOperation.ImagerepoRejectDelError"
//  INTERNALERROR_APPLICATIONREPODELETEPKG = "InternalError.ApplicationRepoDeletePkg"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_REMOTESERVICECALLERROR = "InternalError.RemoteServiceCallError"
//  RESOURCEINUSE_APPLICATIONCANNOTDELETE = "ResourceInUse.ApplicationCannotDelete"
//  RESOURCEINUSE_CVMCAEMASTERCANNOTDELETE = "ResourceInUse.CvmcaeMasterCannotDelete"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
func (c *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteCluster")
    
    
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCluster
// 删除集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_CLUSTEROPERATEENVILLEGALERROR = "FailedOperation.ClusterOperateEnvIllegalError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESTKECLUSTERABNORMAL = "FailedOperation.ContainergroupKubernetesTkeClusterAbnormal"
//  FAILEDOPERATION_CPCLUSTERUNAVAILABLE = "FailedOperation.CpClusterUnavailable"
//  FAILEDOPERATION_TKECLUSTERDELETEFAILED = "FailedOperation.TkeClusterDeleteFailed"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINUSE_CVMCAEMASTERCANNOTDELETE = "ResourceInUse.CvmcaeMasterCannotDelete"
//  RESOURCEINUSE_GROUPEXISTS = "ResourceInUse.GroupExists"
//  RESOURCEINUSE_INSTANCEEXISTS = "ResourceInUse.InstanceExists"
//  RESOURCEINUSE_NONDEFAULTNAMESPACEEXISTS = "ResourceInUse.NonDefaultNamespaceExists"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_TKECLUSTERNOTEXISTS = "ResourceNotFound.TkeClusterNotExists"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    return c.DeleteClusterWithContext(context.Background(), request)
}

// DeleteCluster
// 删除集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_CLUSTEROPERATEENVILLEGALERROR = "FailedOperation.ClusterOperateEnvIllegalError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESTKECLUSTERABNORMAL = "FailedOperation.ContainergroupKubernetesTkeClusterAbnormal"
//  FAILEDOPERATION_CPCLUSTERUNAVAILABLE = "FailedOperation.CpClusterUnavailable"
//  FAILEDOPERATION_TKECLUSTERDELETEFAILED = "FailedOperation.TkeClusterDeleteFailed"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINUSE_CVMCAEMASTERCANNOTDELETE = "ResourceInUse.CvmcaeMasterCannotDelete"
//  RESOURCEINUSE_GROUPEXISTS = "ResourceInUse.GroupExists"
//  RESOURCEINUSE_INSTANCEEXISTS = "ResourceInUse.InstanceExists"
//  RESOURCEINUSE_NONDEFAULTNAMESPACEEXISTS = "ResourceInUse.NonDefaultNamespaceExists"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_TKECLUSTERNOTEXISTS = "ResourceNotFound.TkeClusterNotExists"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigRequest() (request *DeleteConfigRequest) {
    request = &DeleteConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteConfig")
    
    
    return
}

func NewDeleteConfigResponse() (response *DeleteConfigResponse) {
    response = &DeleteConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfig
// 删除配置项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_RELEASEDCONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedConfigCanNotBeDeleted"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteConfig(request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
    return c.DeleteConfigWithContext(context.Background(), request)
}

// DeleteConfig
// 删除配置项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_RELEASEDCONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedConfigCanNotBeDeleted"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteConfigWithContext(ctx context.Context, request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
    if request == nil {
        request = NewDeleteConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigTemplateRequest() (request *DeleteConfigTemplateRequest) {
    request = &DeleteConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteConfigTemplate")
    
    
    return
}

func NewDeleteConfigTemplateResponse() (response *DeleteConfigTemplateResponse) {
    response = &DeleteConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfigTemplate
// 删除模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"
//  FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"
//  INVALIDPARAMETER_CONFIGTEMPLATENAMEINVALID = "InvalidParameter.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"
//  MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"
func (c *Client) DeleteConfigTemplate(request *DeleteConfigTemplateRequest) (response *DeleteConfigTemplateResponse, err error) {
    return c.DeleteConfigTemplateWithContext(context.Background(), request)
}

// DeleteConfigTemplate
// 删除模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"
//  FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"
//  INVALIDPARAMETER_CONFIGTEMPLATENAMEINVALID = "InvalidParameter.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"
//  MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"
func (c *Client) DeleteConfigTemplateWithContext(ctx context.Context, request *DeleteConfigTemplateRequest) (response *DeleteConfigTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteConfigTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteConfigTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfigTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteContainerGroupRequest() (request *DeleteContainerGroupRequest) {
    request = &DeleteContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteContainerGroup")
    
    
    return
}

func NewDeleteContainerGroupResponse() (response *DeleteContainerGroupResponse) {
    response = &DeleteContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteContainerGroup
// 删除容器部署组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESTKECLUSTERABNORMAL = "FailedOperation.ContainergroupKubernetesTkeClusterAbnormal"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteContainerGroup(request *DeleteContainerGroupRequest) (response *DeleteContainerGroupResponse, err error) {
    return c.DeleteContainerGroupWithContext(context.Background(), request)
}

// DeleteContainerGroup
// 删除容器部署组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESTKECLUSTERABNORMAL = "FailedOperation.ContainergroupKubernetesTkeClusterAbnormal"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteContainerGroupWithContext(ctx context.Context, request *DeleteContainerGroupRequest) (response *DeleteContainerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteContainerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteContainerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteContainerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileConfigRequest() (request *DeleteFileConfigRequest) {
    request = &DeleteFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteFileConfig")
    
    
    return
}

func NewDeleteFileConfigResponse() (response *DeleteFileConfigResponse) {
    response = &DeleteFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFileConfig
// 删除文件配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_FILECONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.FileConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_RELEASEDFILECONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedFileConfigCanNotBeDeleted"
func (c *Client) DeleteFileConfig(request *DeleteFileConfigRequest) (response *DeleteFileConfigResponse, err error) {
    return c.DeleteFileConfigWithContext(context.Background(), request)
}

// DeleteFileConfig
// 删除文件配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_FILECONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.FileConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_RELEASEDFILECONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedFileConfigCanNotBeDeleted"
func (c *Client) DeleteFileConfigWithContext(ctx context.Context, request *DeleteFileConfigRequest) (response *DeleteFileConfigResponse, err error) {
    if request == nil {
        request = NewDeleteFileConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteFileConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFileConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGatewayApiRequest() (request *DeleteGatewayApiRequest) {
    request = &DeleteGatewayApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteGatewayApi")
    
    
    return
}

func NewDeleteGatewayApiResponse() (response *DeleteGatewayApiResponse) {
    response = &DeleteGatewayApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGatewayApi
// 批量删除API
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) DeleteGatewayApi(request *DeleteGatewayApiRequest) (response *DeleteGatewayApiResponse, err error) {
    return c.DeleteGatewayApiWithContext(context.Background(), request)
}

// DeleteGatewayApi
// 批量删除API
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) DeleteGatewayApiWithContext(ctx context.Context, request *DeleteGatewayApiRequest) (response *DeleteGatewayApiResponse, err error) {
    if request == nil {
        request = NewDeleteGatewayApiRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteGatewayApi")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGatewayApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGatewayApiResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGroup
// 删除容器部署组
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GROUPDELETECLUSTERTYPEMISMATCH = "InvalidParameterValue.GroupDeleteClusterTypeMismatch"
//  RESOURCEINUSE_GROUPCANNOTDELETE = "ResourceInUse.GroupCannotDelete"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// 删除容器部署组
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GROUPDELETECLUSTERTYPEMISMATCH = "InvalidParameterValue.GroupDeleteClusterTypeMismatch"
//  RESOURCEINUSE_GROUPCANNOTDELETE = "ResourceInUse.GroupCannotDelete"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageTagsRequest() (request *DeleteImageTagsRequest) {
    request = &DeleteImageTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteImageTags")
    
    
    return
}

func NewDeleteImageTagsResponse() (response *DeleteImageTagsResponse) {
    response = &DeleteImageTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteImageTags
// 批量删除镜像版本
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONTAINERGROUPIMAGETAGISINUSE = "InvalidParameterValue.ContainerGroupImageTagIsInUse"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DeleteImageTags(request *DeleteImageTagsRequest) (response *DeleteImageTagsResponse, err error) {
    return c.DeleteImageTagsWithContext(context.Background(), request)
}

// DeleteImageTags
// 批量删除镜像版本
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONTAINERGROUPIMAGETAGISINUSE = "InvalidParameterValue.ContainerGroupImageTagIsInUse"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DeleteImageTagsWithContext(ctx context.Context, request *DeleteImageTagsRequest) (response *DeleteImageTagsResponse, err error) {
    if request == nil {
        request = NewDeleteImageTagsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteImageTags")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLaneRequest() (request *DeleteLaneRequest) {
    request = &DeleteLaneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteLane")
    
    
    return
}

func NewDeleteLaneResponse() (response *DeleteLaneResponse) {
    response = &DeleteLaneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLane
// 删除泳道
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) DeleteLane(request *DeleteLaneRequest) (response *DeleteLaneResponse, err error) {
    return c.DeleteLaneWithContext(context.Background(), request)
}

// DeleteLane
// 删除泳道
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) DeleteLaneWithContext(ctx context.Context, request *DeleteLaneRequest) (response *DeleteLaneResponse, err error) {
    if request == nil {
        request = NewDeleteLaneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteLane")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLane require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLaneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLaneRuleRequest() (request *DeleteLaneRuleRequest) {
    request = &DeleteLaneRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteLaneRule")
    
    
    return
}

func NewDeleteLaneRuleResponse() (response *DeleteLaneRuleResponse) {
    response = &DeleteLaneRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLaneRule
// 删除泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) DeleteLaneRule(request *DeleteLaneRuleRequest) (response *DeleteLaneRuleResponse, err error) {
    return c.DeleteLaneRuleWithContext(context.Background(), request)
}

// DeleteLaneRule
// 删除泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) DeleteLaneRuleWithContext(ctx context.Context, request *DeleteLaneRuleRequest) (response *DeleteLaneRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLaneRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteLaneRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLaneRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLaneRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMicroserviceRequest() (request *DeleteMicroserviceRequest) {
    request = &DeleteMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteMicroservice")
    
    
    return
}

func NewDeleteMicroserviceResponse() (response *DeleteMicroserviceResponse) {
    response = &DeleteMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMicroservice
// 删除微服务
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_SERVICEIDREQUIRED = "MissingParameter.ServiceIdRequired"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteMicroservice(request *DeleteMicroserviceRequest) (response *DeleteMicroserviceResponse, err error) {
    return c.DeleteMicroserviceWithContext(context.Background(), request)
}

// DeleteMicroservice
// 删除微服务
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_SERVICEIDREQUIRED = "MissingParameter.ServiceIdRequired"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteMicroserviceWithContext(ctx context.Context, request *DeleteMicroserviceRequest) (response *DeleteMicroserviceResponse, err error) {
    if request == nil {
        request = NewDeleteMicroserviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteMicroservice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMicroservice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteNamespace")
    
    
    return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
    response = &DeleteNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNamespace
// 删除命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONCREATEESATUHERROR = "FailedOperation.ApplicationCreateEsAtuhError"
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLUSTERCREATEVPCFAIL = "FailedOperation.ClusterCreateVpcFail"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_CLUSTERUPDATEFAILED = "FailedOperation.ClusterUpdateFailed"
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  FAILEDOPERATION_GROUPEXISTS = "FailedOperation.GroupExists"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"
//  FAILEDOPERATION_INSTANCERESETTIMEOUT = "FailedOperation.InstanceResetTimeout"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_LANEINFORELEASEMESHFAILED = "FailedOperation.LaneInfoReleaseMeshFailed"
//  FAILEDOPERATION_NAMESPACECREATEFAILED = "FailedOperation.NamespaceCreateFailed"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_RATELIMITCONSULERROR = "FailedOperation.RatelimitConsulError"
//  FAILEDOPERATION_RATELIMITMESHAPISERVICEERROR = "FailedOperation.RatelimitMeshApiServiceError"
//  FAILEDOPERATION_RESOURCEOPFAILED = "FailedOperation.ResourceOpFailed"
//  FAILEDOPERATION_ROUTEAFFINITYMESHFAILED = "FailedOperation.RouteAffinityMeshFailed"
//  FAILEDOPERATION_ROUTEENABLECONSULFAILED = "FailedOperation.RouteEnableConsulFailed"
//  FAILEDOPERATION_ROUTENAMESPACEREQUESTERROR = "FailedOperation.RouteNamespaceRequestError"
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_TKECLUSTERDELETEFAILED = "FailedOperation.TkeClusterDeleteFailed"
//  FAILEDOPERATION_TKECLUSTERQUERYFAILED = "FailedOperation.TkeClusterQueryFailed"
//  FAILEDOPERATION_TSFAPMAGENTTASKQUERYERROR = "FailedOperation.TsfApmAgentTaskQueryError"
//  FAILEDOPERATION_TSFAPMAGENTTASKWRITEERROR = "FailedOperation.TsfApmAgentTaskWriteError"
//  FAILEDOPERATION_TSFAPMAPMAGENTNOCONNECTION = "FailedOperation.TsfApmApmAgentNoConnection"
//  FAILEDOPERATION_TSFAPMBUSILOGCFGAPPRELATIONWRITEERROR = "FailedOperation.TsfApmBusiLogCfgAppRelationWriteError"
//  FAILEDOPERATION_TSFAPMBUSILOGCFGQUERYERROR = "FailedOperation.TsfApmBusiLogCfgQueryError"
//  FAILEDOPERATION_TSFAPMBUSILOGCFGSCHEMAQUERYERROR = "FailedOperation.TsfApmBusiLogCfgSchemaQueryError"
//  FAILEDOPERATION_TSFAPMBUSILOGCFGSCHEMAWRITEERROR = "FailedOperation.TsfApmBusiLogCfgSchemaWriteError"
//  FAILEDOPERATION_TSFAPMBUSILOGCFGWRITEERROR = "FailedOperation.TsfApmBusiLogCfgWriteError"
//  FAILEDOPERATION_TSFAPMCALLMASTERINTERFACEFAILED = "FailedOperation.TsfApmCallMasterInterfaceFailed"
//  FAILEDOPERATION_TSFAPMSTATSSEARCHSERVICEQUERYERROR = "FailedOperation.TsfApmStatsSearchServiceQueryError"
//  FAILEDOPERATION_TSFASDBINSTERFAIL = "FailedOperation.TsfAsDbInsterFail"
//  FAILEDOPERATION_TSFASDBQUERYFAIL = "FailedOperation.TsfAsDbQueryFail"
//  FAILEDOPERATION_TSFASEXPANDCOUNTANDLIMITERROR = "FailedOperation.TsfAsExpandCountAndLimitError"
//  FAILEDOPERATION_TSFASEXPANDINDICATORSLESSSHRINK = "FailedOperation.TsfAsExpandIndicatorsLessShrink"
//  FAILEDOPERATION_TSFASEXPANDLIMITLESSSHRINKLIMIT = "FailedOperation.TsfAsExpandLimitLessShrinkLimit"
//  FAILEDOPERATION_TSFCMONITORCTSDBCLIENTREQUESTFAIL = "FailedOperation.TsfCmonitorCtsdbClientRequestFail"
//  FAILEDOPERATION_TSFMSSERVERERROR = "FailedOperation.TsfMsServerError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNAUTHORIZEDOPERATION = "FailedOperation.UnauthorizedOperation"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_DISPATCHCOMMONERROR = "InternalError.DispatchCommonError"
//  INTERNALERROR_KUBERNETESCALLERROR = "InternalError.KubernetesCallError"
//  RESOURCEINUSE_DEFAULTNAMEPSACECANNOTBEDELETED = "ResourceInUse.DefaultNamepsaceCannotBeDeleted"
//  RESOURCEINUSE_NAMESPACECANNOTDELETE = "ResourceInUse.NamespaceCannotDelete"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    return c.DeleteNamespaceWithContext(context.Background(), request)
}

// DeleteNamespace
// 删除命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONCREATEESATUHERROR = "FailedOperation.ApplicationCreateEsAtuhError"
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLUSTERCREATEVPCFAIL = "FailedOperation.ClusterCreateVpcFail"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_CLUSTERUPDATEFAILED = "FailedOperation.ClusterUpdateFailed"
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  FAILEDOPERATION_GROUPEXISTS = "FailedOperation.GroupExists"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"
//  FAILEDOPERATION_INSTANCERESETTIMEOUT = "FailedOperation.InstanceResetTimeout"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"
//  FAILEDOPERATION_LANEINFORELEASEMESHFAILED = "FailedOperation.LaneInfoReleaseMeshFailed"
//  FAILEDOPERATION_NAMESPACECREATEFAILED = "FailedOperation.NamespaceCreateFailed"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_RATELIMITCONSULERROR = "FailedOperation.RatelimitConsulError"
//  FAILEDOPERATION_RATELIMITMESHAPISERVICEERROR = "FailedOperation.RatelimitMeshApiServiceError"
//  FAILEDOPERATION_RESOURCEOPFAILED = "FailedOperation.ResourceOpFailed"
//  FAILEDOPERATION_ROUTEAFFINITYMESHFAILED = "FailedOperation.RouteAffinityMeshFailed"
//  FAILEDOPERATION_ROUTEENABLECONSULFAILED = "FailedOperation.RouteEnableConsulFailed"
//  FAILEDOPERATION_ROUTENAMESPACEREQUESTERROR = "FailedOperation.RouteNamespaceRequestError"
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_TKECLUSTERDELETEFAILED = "FailedOperation.TkeClusterDeleteFailed"
//  FAILEDOPERATION_TKECLUSTERQUERYFAILED = "FailedOperation.TkeClusterQueryFailed"
//  FAILEDOPERATION_TSFAPMAGENTTASKQUERYERROR = "FailedOperation.TsfApmAgentTaskQueryError"
//  FAILEDOPERATION_TSFAPMAGENTTASKWRITEERROR = "FailedOperation.TsfApmAgentTaskWriteError"
//  FAILEDOPERATION_TSFAPMAPMAGENTNOCONNECTION = "FailedOperation.TsfApmApmAgentNoConnection"
//  FAILEDOPERATION_TSFAPMBUSILOGCFGAPPRELATIONWRITEERROR = "FailedOperation.TsfApmBusiLogCfgAppRelationWriteError"
//  FAILEDOPERATION_TSFAPMBUSILOGCFGQUERYERROR = "FailedOperation.TsfApmBusiLogCfgQueryError"
//  FAILEDOPERATION_TSFAPMBUSILOGCFGSCHEMAQUERYERROR = "FailedOperation.TsfApmBusiLogCfgSchemaQueryError"
//  FAILEDOPERATION_TSFAPMBUSILOGCFGSCHEMAWRITEERROR = "FailedOperation.TsfApmBusiLogCfgSchemaWriteError"
//  FAILEDOPERATION_TSFAPMBUSILOGCFGWRITEERROR = "FailedOperation.TsfApmBusiLogCfgWriteError"
//  FAILEDOPERATION_TSFAPMCALLMASTERINTERFACEFAILED = "FailedOperation.TsfApmCallMasterInterfaceFailed"
//  FAILEDOPERATION_TSFAPMSTATSSEARCHSERVICEQUERYERROR = "FailedOperation.TsfApmStatsSearchServiceQueryError"
//  FAILEDOPERATION_TSFASDBINSTERFAIL = "FailedOperation.TsfAsDbInsterFail"
//  FAILEDOPERATION_TSFASDBQUERYFAIL = "FailedOperation.TsfAsDbQueryFail"
//  FAILEDOPERATION_TSFASEXPANDCOUNTANDLIMITERROR = "FailedOperation.TsfAsExpandCountAndLimitError"
//  FAILEDOPERATION_TSFASEXPANDINDICATORSLESSSHRINK = "FailedOperation.TsfAsExpandIndicatorsLessShrink"
//  FAILEDOPERATION_TSFASEXPANDLIMITLESSSHRINKLIMIT = "FailedOperation.TsfAsExpandLimitLessShrinkLimit"
//  FAILEDOPERATION_TSFCMONITORCTSDBCLIENTREQUESTFAIL = "FailedOperation.TsfCmonitorCtsdbClientRequestFail"
//  FAILEDOPERATION_TSFMSSERVERERROR = "FailedOperation.TsfMsServerError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNAUTHORIZEDOPERATION = "FailedOperation.UnauthorizedOperation"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_DISPATCHCOMMONERROR = "InternalError.DispatchCommonError"
//  INTERNALERROR_KUBERNETESCALLERROR = "InternalError.KubernetesCallError"
//  RESOURCEINUSE_DEFAULTNAMEPSACECANNOTBEDELETED = "ResourceInUse.DefaultNamepsaceCannotBeDeleted"
//  RESOURCEINUSE_NAMESPACECANNOTDELETE = "ResourceInUse.NamespaceCannotDelete"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteNamespaceWithContext(ctx context.Context, request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePathRewritesRequest() (request *DeletePathRewritesRequest) {
    request = &DeletePathRewritesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePathRewrites")
    
    
    return
}

func NewDeletePathRewritesResponse() (response *DeletePathRewritesResponse) {
    response = &DeletePathRewritesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePathRewrites
// 删除路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeletePathRewrites(request *DeletePathRewritesRequest) (response *DeletePathRewritesResponse, err error) {
    return c.DeletePathRewritesWithContext(context.Background(), request)
}

// DeletePathRewrites
// 删除路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeletePathRewritesWithContext(ctx context.Context, request *DeletePathRewritesRequest) (response *DeletePathRewritesResponse, err error) {
    if request == nil {
        request = NewDeletePathRewritesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeletePathRewrites")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePathRewrites require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePathRewritesResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePkgsRequest() (request *DeletePkgsRequest) {
    request = &DeletePkgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePkgs")
    
    
    return
}

func NewDeletePkgsResponse() (response *DeletePkgsResponse) {
    response = &DeletePkgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePkgs
// 从软件仓库批量删除程序包。
//
// 一次最多支持删除1000个包，数量超过1000，返回UpperDeleteLimit错误。
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PACKAGEINUSE = "InvalidParameter.PackageInUse"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_UPPERDELETELIMIT = "InvalidParameter.UpperDeleteLimit"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DeletePkgs(request *DeletePkgsRequest) (response *DeletePkgsResponse, err error) {
    return c.DeletePkgsWithContext(context.Background(), request)
}

// DeletePkgs
// 从软件仓库批量删除程序包。
//
// 一次最多支持删除1000个包，数量超过1000，返回UpperDeleteLimit错误。
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PACKAGEINUSE = "InvalidParameter.PackageInUse"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_UPPERDELETELIMIT = "InvalidParameter.UpperDeleteLimit"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DeletePkgsWithContext(ctx context.Context, request *DeletePkgsRequest) (response *DeletePkgsResponse, err error) {
    if request == nil {
        request = NewDeletePkgsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeletePkgs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePkgs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePkgsResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePublicConfigRequest() (request *DeletePublicConfigRequest) {
    request = &DeletePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePublicConfig")
    
    
    return
}

func NewDeletePublicConfigResponse() (response *DeletePublicConfigResponse) {
    response = &DeletePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePublicConfig
// 删除公共配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_RELEASEDCONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedConfigCanNotBeDeleted"
//  MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"
func (c *Client) DeletePublicConfig(request *DeletePublicConfigRequest) (response *DeletePublicConfigResponse, err error) {
    return c.DeletePublicConfigWithContext(context.Background(), request)
}

// DeletePublicConfig
// 删除公共配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_RELEASEDCONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedConfigCanNotBeDeleted"
//  MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"
func (c *Client) DeletePublicConfigWithContext(ctx context.Context, request *DeletePublicConfigRequest) (response *DeletePublicConfigResponse, err error) {
    if request == nil {
        request = NewDeletePublicConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeletePublicConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePublicConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRepositoryRequest() (request *DeleteRepositoryRequest) {
    request = &DeleteRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteRepository")
    
    
    return
}

func NewDeleteRepositoryResponse() (response *DeleteRepositoryResponse) {
    response = &DeleteRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRepository
// 删除仓库
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOSITORYNOTEMPTY = "InvalidParameter.RepositoryNotEmpty"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DeleteRepository(request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
    return c.DeleteRepositoryWithContext(context.Background(), request)
}

// DeleteRepository
// 删除仓库
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOSITORYNOTEMPTY = "InvalidParameter.RepositoryNotEmpty"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DeleteRepositoryWithContext(ctx context.Context, request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteRepository")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServerlessGroupRequest() (request *DeleteServerlessGroupRequest) {
    request = &DeleteServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteServerlessGroup")
    
    
    return
}

func NewDeleteServerlessGroupResponse() (response *DeleteServerlessGroupResponse) {
    response = &DeleteServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteServerlessGroup
// serverless 能力已下线。下线对应接口。
//
// 
//
// 删除Serverless部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INVALIDPARAMETERVALUE_GROUPDELETECLUSTERTYPEMISMATCH = "InvalidParameterValue.GroupDeleteClusterTypeMismatch"
//  INVALIDPARAMETERVALUE_GROUPIDNULL = "InvalidParameterValue.GroupIdNull"
//  RESOURCEINUSE_GROUPINOPERATION = "ResourceInUse.GroupInOperation"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DeleteServerlessGroup(request *DeleteServerlessGroupRequest) (response *DeleteServerlessGroupResponse, err error) {
    return c.DeleteServerlessGroupWithContext(context.Background(), request)
}

// DeleteServerlessGroup
// serverless 能力已下线。下线对应接口。
//
// 
//
// 删除Serverless部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INVALIDPARAMETERVALUE_GROUPDELETECLUSTERTYPEMISMATCH = "InvalidParameterValue.GroupDeleteClusterTypeMismatch"
//  INVALIDPARAMETERVALUE_GROUPIDNULL = "InvalidParameterValue.GroupIdNull"
//  RESOURCEINUSE_GROUPINOPERATION = "ResourceInUse.GroupInOperation"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DeleteServerlessGroupWithContext(ctx context.Context, request *DeleteServerlessGroupRequest) (response *DeleteServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDeleteServerlessGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteServerlessGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServerlessGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
    request = &DeleteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteTask")
    
    
    return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
    response = &DeleteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTask
// 删除任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKDELETEERROR = "FailedOperation.TaskDeleteError"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    return c.DeleteTaskWithContext(context.Background(), request)
}

// DeleteTask
// 删除任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKDELETEERROR = "FailedOperation.TaskDeleteError"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteTaskWithContext(ctx context.Context, request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUnitNamespacesRequest() (request *DeleteUnitNamespacesRequest) {
    request = &DeleteUnitNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteUnitNamespaces")
    
    
    return
}

func NewDeleteUnitNamespacesResponse() (response *DeleteUnitNamespacesResponse) {
    response = &DeleteUnitNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUnitNamespaces
// 删除单元化命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeleteUnitNamespaces(request *DeleteUnitNamespacesRequest) (response *DeleteUnitNamespacesResponse, err error) {
    return c.DeleteUnitNamespacesWithContext(context.Background(), request)
}

// DeleteUnitNamespaces
// 删除单元化命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeleteUnitNamespacesWithContext(ctx context.Context, request *DeleteUnitNamespacesRequest) (response *DeleteUnitNamespacesResponse, err error) {
    if request == nil {
        request = NewDeleteUnitNamespacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteUnitNamespaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUnitNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUnitNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUnitRuleRequest() (request *DeleteUnitRuleRequest) {
    request = &DeleteUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteUnitRule")
    
    
    return
}

func NewDeleteUnitRuleResponse() (response *DeleteUnitRuleResponse) {
    response = &DeleteUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUnitRule
// 删除单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeleteUnitRule(request *DeleteUnitRuleRequest) (response *DeleteUnitRuleResponse, err error) {
    return c.DeleteUnitRuleWithContext(context.Background(), request)
}

// DeleteUnitRule
// 删除单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeleteUnitRuleWithContext(ctx context.Context, request *DeleteUnitRuleRequest) (response *DeleteUnitRuleResponse, err error) {
    if request == nil {
        request = NewDeleteUnitRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeleteUnitRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUnitRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeployContainerApplicationRequest() (request *DeployContainerApplicationRequest) {
    request = &DeployContainerApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeployContainerApplication")
    
    
    return
}

func NewDeployContainerApplicationResponse() (response *DeployContainerApplicationResponse) {
    response = &DeployContainerApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployContainerApplication
// 部署容器应用-更新
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_TOKENQUERYFAILED = "FailedOperation.TokenQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPENVVALUENOTSET = "InvalidParameterValue.ContainergroupEnvValueNotSet"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTNULL = "InvalidParameterValue.ContainergroupPortNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTSREPEAT = "InvalidParameterValue.ContainergroupPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLINVALID = "InvalidParameterValue.ContainergroupProtocolInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLMIXERROR = "InvalidParameterValue.ContainergroupProtocolMixError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLNULL = "InvalidParameterValue.ContainergroupProtocolNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPREPONAMEINVALID = "InvalidParameterValue.ContainergroupReponameInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPRESOURCEAGENTVALUEINVALID = "InvalidParameterValue.ContainergroupResourceAgentValueInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTSREPEAT = "InvalidParameterValue.ContainergroupTargetPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATETYPEINVALID = "InvalidParameterValue.ContainergroupUpdatetypeInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPYAMLUSERCONTAINERNOTFOUND = "InvalidParameterValue.ContainergroupYamlUserContainerNotFound"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMENULL = "InvalidParameterValue.ImagerepoRepoNameNull"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"
//  INVALIDPARAMETERVALUE_IMAGEREPOTAGNAMENULL = "InvalidParameterValue.ImagerepoTagnameNull"
//  INVALIDPARAMETERVALUE_WRONGDONTSTARTVALUE = "InvalidParameterValue.WrongDontStartValue"
//  MISSINGPARAMETER_SYSTEMPARAMETERREQUIRED = "MissingParameter.SystemParameterRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeployContainerApplication(request *DeployContainerApplicationRequest) (response *DeployContainerApplicationResponse, err error) {
    return c.DeployContainerApplicationWithContext(context.Background(), request)
}

// DeployContainerApplication
// 部署容器应用-更新
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_TOKENQUERYFAILED = "FailedOperation.TokenQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPENVVALUENOTSET = "InvalidParameterValue.ContainergroupEnvValueNotSet"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTNULL = "InvalidParameterValue.ContainergroupPortNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTSREPEAT = "InvalidParameterValue.ContainergroupPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLINVALID = "InvalidParameterValue.ContainergroupProtocolInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLMIXERROR = "InvalidParameterValue.ContainergroupProtocolMixError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLNULL = "InvalidParameterValue.ContainergroupProtocolNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPREPONAMEINVALID = "InvalidParameterValue.ContainergroupReponameInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPRESOURCEAGENTVALUEINVALID = "InvalidParameterValue.ContainergroupResourceAgentValueInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTSREPEAT = "InvalidParameterValue.ContainergroupTargetPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATETYPEINVALID = "InvalidParameterValue.ContainergroupUpdatetypeInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPYAMLUSERCONTAINERNOTFOUND = "InvalidParameterValue.ContainergroupYamlUserContainerNotFound"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMENULL = "InvalidParameterValue.ImagerepoRepoNameNull"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"
//  INVALIDPARAMETERVALUE_IMAGEREPOTAGNAMENULL = "InvalidParameterValue.ImagerepoTagnameNull"
//  INVALIDPARAMETERVALUE_WRONGDONTSTARTVALUE = "InvalidParameterValue.WrongDontStartValue"
//  MISSINGPARAMETER_SYSTEMPARAMETERREQUIRED = "MissingParameter.SystemParameterRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeployContainerApplicationWithContext(ctx context.Context, request *DeployContainerApplicationRequest) (response *DeployContainerApplicationResponse, err error) {
    if request == nil {
        request = NewDeployContainerApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeployContainerApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployContainerApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployContainerApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeployContainerGroupRequest() (request *DeployContainerGroupRequest) {
    request = &DeployContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeployContainerGroup")
    
    
    return
}

func NewDeployContainerGroupResponse() (response *DeployContainerGroupResponse) {
    response = &DeployContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployContainerGroup
// 部署容器应用-更新
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_TOKENQUERYFAILED = "FailedOperation.TokenQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPENVVALUENOTSET = "InvalidParameterValue.ContainergroupEnvValueNotSet"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTNULL = "InvalidParameterValue.ContainergroupPortNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTSREPEAT = "InvalidParameterValue.ContainergroupPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLINVALID = "InvalidParameterValue.ContainergroupProtocolInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLMIXERROR = "InvalidParameterValue.ContainergroupProtocolMixError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLNULL = "InvalidParameterValue.ContainergroupProtocolNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPREPONAMEINVALID = "InvalidParameterValue.ContainergroupReponameInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPRESOURCEAGENTVALUEINVALID = "InvalidParameterValue.ContainergroupResourceAgentValueInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTSREPEAT = "InvalidParameterValue.ContainergroupTargetPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATETYPEINVALID = "InvalidParameterValue.ContainergroupUpdatetypeInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPYAMLUSERCONTAINERNOTFOUND = "InvalidParameterValue.ContainergroupYamlUserContainerNotFound"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMENULL = "InvalidParameterValue.ImagerepoRepoNameNull"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"
//  INVALIDPARAMETERVALUE_IMAGEREPOTAGNAMENULL = "InvalidParameterValue.ImagerepoTagnameNull"
//  INVALIDPARAMETERVALUE_WRONGDONTSTARTVALUE = "InvalidParameterValue.WrongDontStartValue"
//  MISSINGPARAMETER_SYSTEMPARAMETERREQUIRED = "MissingParameter.SystemParameterRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeployContainerGroup(request *DeployContainerGroupRequest) (response *DeployContainerGroupResponse, err error) {
    return c.DeployContainerGroupWithContext(context.Background(), request)
}

// DeployContainerGroup
// 部署容器应用-更新
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_TOKENQUERYFAILED = "FailedOperation.TokenQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPENVVALUENOTSET = "InvalidParameterValue.ContainergroupEnvValueNotSet"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTNULL = "InvalidParameterValue.ContainergroupPortNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTSREPEAT = "InvalidParameterValue.ContainergroupPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLINVALID = "InvalidParameterValue.ContainergroupProtocolInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLMIXERROR = "InvalidParameterValue.ContainergroupProtocolMixError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLNULL = "InvalidParameterValue.ContainergroupProtocolNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPREPONAMEINVALID = "InvalidParameterValue.ContainergroupReponameInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPRESOURCEAGENTVALUEINVALID = "InvalidParameterValue.ContainergroupResourceAgentValueInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTSREPEAT = "InvalidParameterValue.ContainergroupTargetPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATETYPEINVALID = "InvalidParameterValue.ContainergroupUpdatetypeInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPYAMLUSERCONTAINERNOTFOUND = "InvalidParameterValue.ContainergroupYamlUserContainerNotFound"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMENULL = "InvalidParameterValue.ImagerepoRepoNameNull"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"
//  INVALIDPARAMETERVALUE_IMAGEREPOTAGNAMENULL = "InvalidParameterValue.ImagerepoTagnameNull"
//  INVALIDPARAMETERVALUE_WRONGDONTSTARTVALUE = "InvalidParameterValue.WrongDontStartValue"
//  MISSINGPARAMETER_SYSTEMPARAMETERREQUIRED = "MissingParameter.SystemParameterRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeployContainerGroupWithContext(ctx context.Context, request *DeployContainerGroupRequest) (response *DeployContainerGroupResponse, err error) {
    if request == nil {
        request = NewDeployContainerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeployContainerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployContainerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeployGroupRequest() (request *DeployGroupRequest) {
    request = &DeployGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DeployGroup")
    
    
    return
}

func NewDeployGroupResponse() (response *DeployGroupResponse) {
    response = &DeployGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployGroup
// 部署虚拟机部署组应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAPMBINDREQUIRED = "FailedOperation.TapmBindRequired"
//  FAILEDOPERATION_TAPMEXPIREDOTEL = "FailedOperation.TapmExpiredOtel"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CVMCAEMASTERINTERNALERROR = "InternalError.CvmCaeMasterInternalError"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTNOTFOUND = "InvalidParameterValue.CvmCaeMasterAgentNotFound"
//  INVALIDPARAMETERVALUE_GROUPBATCHPARAMETERINVALID = "InvalidParameterValue.GroupBatchParameterInvalid"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeployGroup(request *DeployGroupRequest) (response *DeployGroupResponse, err error) {
    return c.DeployGroupWithContext(context.Background(), request)
}

// DeployGroup
// 部署虚拟机部署组应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAPMBINDREQUIRED = "FailedOperation.TapmBindRequired"
//  FAILEDOPERATION_TAPMEXPIREDOTEL = "FailedOperation.TapmExpiredOtel"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CVMCAEMASTERINTERNALERROR = "InternalError.CvmCaeMasterInternalError"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTNOTFOUND = "InvalidParameterValue.CvmCaeMasterAgentNotFound"
//  INVALIDPARAMETERVALUE_GROUPBATCHPARAMETERINVALID = "InvalidParameterValue.GroupBatchParameterInvalid"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeployGroupWithContext(ctx context.Context, request *DeployGroupRequest) (response *DeployGroupResponse, err error) {
    if request == nil {
        request = NewDeployGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DeployGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiDetailRequest() (request *DescribeApiDetailRequest) {
    request = &DescribeApiDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiDetail")
    
    
    return
}

func NewDescribeApiDetailResponse() (response *DescribeApiDetailResponse) {
    response = &DescribeApiDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiDetail
// 查询API详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIMETAPARSEFAILED = "FailedOperation.ApiMetaParseFailed"
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeApiDetail(request *DescribeApiDetailRequest) (response *DescribeApiDetailResponse, err error) {
    return c.DescribeApiDetailWithContext(context.Background(), request)
}

// DescribeApiDetail
// 查询API详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIMETAPARSEFAILED = "FailedOperation.ApiMetaParseFailed"
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeApiDetailWithContext(ctx context.Context, request *DescribeApiDetailRequest) (response *DescribeApiDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApiDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeApiDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiGroupRequest() (request *DescribeApiGroupRequest) {
    request = &DescribeApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiGroup")
    
    
    return
}

func NewDescribeApiGroupResponse() (response *DescribeApiGroupResponse) {
    response = &DescribeApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiGroup
// 查询API分组
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeApiGroup(request *DescribeApiGroupRequest) (response *DescribeApiGroupResponse, err error) {
    return c.DescribeApiGroupWithContext(context.Background(), request)
}

// DescribeApiGroup
// 查询API分组
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeApiGroupWithContext(ctx context.Context, request *DescribeApiGroupRequest) (response *DescribeApiGroupResponse, err error) {
    if request == nil {
        request = NewDescribeApiGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeApiGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiGroupsRequest() (request *DescribeApiGroupsRequest) {
    request = &DescribeApiGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiGroups")
    
    
    return
}

func NewDescribeApiGroupsResponse() (response *DescribeApiGroupsResponse) {
    response = &DescribeApiGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiGroups
// 查询API 分组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApiGroups(request *DescribeApiGroupsRequest) (response *DescribeApiGroupsResponse, err error) {
    return c.DescribeApiGroupsWithContext(context.Background(), request)
}

// DescribeApiGroups
// 查询API 分组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApiGroupsWithContext(ctx context.Context, request *DescribeApiGroupsRequest) (response *DescribeApiGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeApiGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeApiGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiRateLimitRulesRequest() (request *DescribeApiRateLimitRulesRequest) {
    request = &DescribeApiRateLimitRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiRateLimitRules")
    
    
    return
}

func NewDescribeApiRateLimitRulesResponse() (response *DescribeApiRateLimitRulesResponse) {
    response = &DescribeApiRateLimitRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiRateLimitRules
// 查询API限流规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
func (c *Client) DescribeApiRateLimitRules(request *DescribeApiRateLimitRulesRequest) (response *DescribeApiRateLimitRulesResponse, err error) {
    return c.DescribeApiRateLimitRulesWithContext(context.Background(), request)
}

// DescribeApiRateLimitRules
// 查询API限流规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
func (c *Client) DescribeApiRateLimitRulesWithContext(ctx context.Context, request *DescribeApiRateLimitRulesRequest) (response *DescribeApiRateLimitRulesResponse, err error) {
    if request == nil {
        request = NewDescribeApiRateLimitRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeApiRateLimitRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiRateLimitRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiRateLimitRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiUseDetailRequest() (request *DescribeApiUseDetailRequest) {
    request = &DescribeApiUseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiUseDetail")
    
    
    return
}

func NewDescribeApiUseDetailResponse() (response *DescribeApiUseDetailResponse) {
    response = &DescribeApiUseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiUseDetail
// 查询网关API监控明细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeApiUseDetail(request *DescribeApiUseDetailRequest) (response *DescribeApiUseDetailResponse, err error) {
    return c.DescribeApiUseDetailWithContext(context.Background(), request)
}

// DescribeApiUseDetail
// 查询网关API监控明细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeApiUseDetailWithContext(ctx context.Context, request *DescribeApiUseDetailRequest) (response *DescribeApiUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApiUseDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeApiUseDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiUseDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiUseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiVersionsRequest() (request *DescribeApiVersionsRequest) {
    request = &DescribeApiVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiVersions")
    
    
    return
}

func NewDescribeApiVersionsResponse() (response *DescribeApiVersionsResponse) {
    response = &DescribeApiVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiVersions
// 查询API 版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeApiVersions(request *DescribeApiVersionsRequest) (response *DescribeApiVersionsResponse, err error) {
    return c.DescribeApiVersionsWithContext(context.Background(), request)
}

// DescribeApiVersions
// 查询API 版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeApiVersionsWithContext(ctx context.Context, request *DescribeApiVersionsRequest) (response *DescribeApiVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeApiVersionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeApiVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationRequest() (request *DescribeApplicationRequest) {
    request = &DescribeApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplication")
    
    
    return
}

func NewDescribeApplicationResponse() (response *DescribeApplicationResponse) {
    response = &DescribeApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplication
// 获取应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_APPLICATIONIDNULL = "MissingParameter.ApplicationIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplication(request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    return c.DescribeApplicationWithContext(context.Background(), request)
}

// DescribeApplication
// 获取应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_APPLICATIONIDNULL = "MissingParameter.ApplicationIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplicationWithContext(ctx context.Context, request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationAttributeRequest() (request *DescribeApplicationAttributeRequest) {
    request = &DescribeApplicationAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationAttribute")
    
    
    return
}

func NewDescribeApplicationAttributeResponse() (response *DescribeApplicationAttributeResponse) {
    response = &DescribeApplicationAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationAttribute
// 获取应用列表其它字段，如实例数量信息等
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_APPLICATIONIDNULL = "MissingParameter.ApplicationIdNull"
//  RESOURCENOTFOUND_APPLICATIONPROJECTNOTMATCH = "ResourceNotFound.ApplicationProjectNotMatch"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplicationAttribute(request *DescribeApplicationAttributeRequest) (response *DescribeApplicationAttributeResponse, err error) {
    return c.DescribeApplicationAttributeWithContext(context.Background(), request)
}

// DescribeApplicationAttribute
// 获取应用列表其它字段，如实例数量信息等
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_APPLICATIONIDNULL = "MissingParameter.ApplicationIdNull"
//  RESOURCENOTFOUND_APPLICATIONPROJECTNOTMATCH = "ResourceNotFound.ApplicationProjectNotMatch"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplicationAttributeWithContext(ctx context.Context, request *DescribeApplicationAttributeRequest) (response *DescribeApplicationAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeApplicationAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsRequest() (request *DescribeApplicationsRequest) {
    request = &DescribeApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplications")
    
    
    return
}

func NewDescribeApplicationsResponse() (response *DescribeApplicationsResponse) {
    response = &DescribeApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplications
// 获取应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_APPLICATIONORDERTYPEINVALID = "InvalidParameterValue.ApplicationOrderTypeInvalid"
//  INVALIDPARAMETERVALUE_APPLICATIONPAGELIMITINVALID = "InvalidParameterValue.ApplicationPageLimitInvalid"
//  RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplications(request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    return c.DescribeApplicationsWithContext(context.Background(), request)
}

// DescribeApplications
// 获取应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_APPLICATIONORDERTYPEINVALID = "InvalidParameterValue.ApplicationOrderTypeInvalid"
//  INVALIDPARAMETERVALUE_APPLICATIONPAGELIMITINVALID = "InvalidParameterValue.ApplicationPageLimitInvalid"
//  RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplicationsWithContext(ctx context.Context, request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeApplications")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicResourceUsageRequest() (request *DescribeBasicResourceUsageRequest) {
    request = &DescribeBasicResourceUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeBasicResourceUsage")
    
    
    return
}

func NewDescribeBasicResourceUsageResponse() (response *DescribeBasicResourceUsageResponse) {
    response = &DescribeBasicResourceUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBasicResourceUsage
// TSF基本资源信息概览接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeBasicResourceUsage(request *DescribeBasicResourceUsageRequest) (response *DescribeBasicResourceUsageResponse, err error) {
    return c.DescribeBasicResourceUsageWithContext(context.Background(), request)
}

// DescribeBasicResourceUsage
// TSF基本资源信息概览接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeBasicResourceUsageWithContext(ctx context.Context, request *DescribeBasicResourceUsageRequest) (response *DescribeBasicResourceUsageResponse, err error) {
    if request == nil {
        request = NewDescribeBasicResourceUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeBasicResourceUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBasicResourceUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBasicResourceUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBusinessLogConfigRequest() (request *DescribeBusinessLogConfigRequest) {
    request = &DescribeBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeBusinessLogConfig")
    
    
    return
}

func NewDescribeBusinessLogConfigResponse() (response *DescribeBusinessLogConfigResponse) {
    response = &DescribeBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBusinessLogConfig
// 查询业务日志配置项信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeBusinessLogConfig(request *DescribeBusinessLogConfigRequest) (response *DescribeBusinessLogConfigResponse, err error) {
    return c.DescribeBusinessLogConfigWithContext(context.Background(), request)
}

// DescribeBusinessLogConfig
// 查询业务日志配置项信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeBusinessLogConfigWithContext(ctx context.Context, request *DescribeBusinessLogConfigRequest) (response *DescribeBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBusinessLogConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeBusinessLogConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBusinessLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBusinessLogConfigsRequest() (request *DescribeBusinessLogConfigsRequest) {
    request = &DescribeBusinessLogConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeBusinessLogConfigs")
    
    
    return
}

func NewDescribeBusinessLogConfigsResponse() (response *DescribeBusinessLogConfigsResponse) {
    response = &DescribeBusinessLogConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBusinessLogConfigs
// 查询日志配置项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMBUSILOGCFGWRITEERROR = "FailedOperation.TsfApmBusiLogCfgWriteError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGLIMITPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgLimitParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeBusinessLogConfigs(request *DescribeBusinessLogConfigsRequest) (response *DescribeBusinessLogConfigsResponse, err error) {
    return c.DescribeBusinessLogConfigsWithContext(context.Background(), request)
}

// DescribeBusinessLogConfigs
// 查询日志配置项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMBUSILOGCFGWRITEERROR = "FailedOperation.TsfApmBusiLogCfgWriteError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGLIMITPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgLimitParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeBusinessLogConfigsWithContext(ctx context.Context, request *DescribeBusinessLogConfigsRequest) (response *DescribeBusinessLogConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeBusinessLogConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeBusinessLogConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBusinessLogConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBusinessLogConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
    request = &DescribeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterInstances")
    
    
    return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
    response = &DescribeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterInstances
// 查询集群实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"
//  FAILEDOPERATION_INSTANCEUPDATEFAILED = "FailedOperation.InstanceUpdateFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CAMROLEREQUESTERROR = "InternalError.CamRoleRequestError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CLUSTERNOTEXISTORPRIVILEGEERROR = "InternalError.ClusterNotExistOrPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    return c.DescribeClusterInstancesWithContext(context.Background(), request)
}

// DescribeClusterInstances
// 查询集群实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"
//  FAILEDOPERATION_INSTANCEUPDATEFAILED = "FailedOperation.InstanceUpdateFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CAMROLEREQUESTERROR = "InternalError.CamRoleRequestError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CLUSTERNOTEXISTORPRIVILEGEERROR = "InternalError.ClusterNotExistOrPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeClusterInstancesWithContext(ctx context.Context, request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeClusterInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusters
// 获取集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CLUSTERPAGELIMITINVALID = "InvalidParameterValue.ClusterPageLimitInvalid"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// 获取集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CLUSTERPAGELIMITINVALID = "InvalidParameterValue.ClusterPageLimitInvalid"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
    request = &DescribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfig")
    
    
    return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
    response = &DescribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfig
// 查询配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    return c.DescribeConfigWithContext(context.Background(), request)
}

// DescribeConfig
// 查询配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeConfigWithContext(ctx context.Context, request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigReleaseLogsRequest() (request *DescribeConfigReleaseLogsRequest) {
    request = &DescribeConfigReleaseLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigReleaseLogs")
    
    
    return
}

func NewDescribeConfigReleaseLogsResponse() (response *DescribeConfigReleaseLogsResponse) {
    response = &DescribeConfigReleaseLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigReleaseLogs
// 查询配置发布历史
//
// 可能返回的错误码:
//  INTERNALERROR_SQLTOOMANYINITEM = "InternalError.SqlTooManyInItem"
func (c *Client) DescribeConfigReleaseLogs(request *DescribeConfigReleaseLogsRequest) (response *DescribeConfigReleaseLogsResponse, err error) {
    return c.DescribeConfigReleaseLogsWithContext(context.Background(), request)
}

// DescribeConfigReleaseLogs
// 查询配置发布历史
//
// 可能返回的错误码:
//  INTERNALERROR_SQLTOOMANYINITEM = "InternalError.SqlTooManyInItem"
func (c *Client) DescribeConfigReleaseLogsWithContext(ctx context.Context, request *DescribeConfigReleaseLogsRequest) (response *DescribeConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigReleaseLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeConfigReleaseLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigReleaseLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigReleasesRequest() (request *DescribeConfigReleasesRequest) {
    request = &DescribeConfigReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigReleases")
    
    
    return
}

func NewDescribeConfigReleasesResponse() (response *DescribeConfigReleasesResponse) {
    response = &DescribeConfigReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigReleases
// 查询配置发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  FAILEDOPERATION_CONFIGRELEASEQUERYFAILED = "FailedOperation.ConfigReleaseQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigReleases(request *DescribeConfigReleasesRequest) (response *DescribeConfigReleasesResponse, err error) {
    return c.DescribeConfigReleasesWithContext(context.Background(), request)
}

// DescribeConfigReleases
// 查询配置发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  FAILEDOPERATION_CONFIGRELEASEQUERYFAILED = "FailedOperation.ConfigReleaseQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigReleasesWithContext(ctx context.Context, request *DescribeConfigReleasesRequest) (response *DescribeConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigReleasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeConfigReleases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigReleases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigSummaryRequest() (request *DescribeConfigSummaryRequest) {
    request = &DescribeConfigSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigSummary")
    
    
    return
}

func NewDescribeConfigSummaryResponse() (response *DescribeConfigSummaryResponse) {
    response = &DescribeConfigSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigSummary
// 查询配置汇总列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigSummary(request *DescribeConfigSummaryRequest) (response *DescribeConfigSummaryResponse, err error) {
    return c.DescribeConfigSummaryWithContext(context.Background(), request)
}

// DescribeConfigSummary
// 查询配置汇总列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigSummaryWithContext(ctx context.Context, request *DescribeConfigSummaryRequest) (response *DescribeConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeConfigSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeConfigSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigTemplateRequest() (request *DescribeConfigTemplateRequest) {
    request = &DescribeConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigTemplate")
    
    
    return
}

func NewDescribeConfigTemplateResponse() (response *DescribeConfigTemplateResponse) {
    response = &DescribeConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigTemplate
// 导入配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"
//  FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"
//  INVALIDPARAMETER_CONFIGTEMPLATENAMEINVALID = "InvalidParameter.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"
//  MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigTemplate(request *DescribeConfigTemplateRequest) (response *DescribeConfigTemplateResponse, err error) {
    return c.DescribeConfigTemplateWithContext(context.Background(), request)
}

// DescribeConfigTemplate
// 导入配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"
//  FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"
//  INVALIDPARAMETER_CONFIGTEMPLATENAMEINVALID = "InvalidParameter.ConfigTemplateNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"
//  MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigTemplateWithContext(ctx context.Context, request *DescribeConfigTemplateRequest) (response *DescribeConfigTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeConfigTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeConfigTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
    request = &DescribeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigs")
    
    
    return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
    response = &DescribeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigs
// 查询配置项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  FAILEDOPERATION_CONFIGQUERYFAILED = "FailedOperation.ConfigQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    return c.DescribeConfigsWithContext(context.Background(), request)
}

// DescribeConfigs
// 查询配置项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  FAILEDOPERATION_CONFIGQUERYFAILED = "FailedOperation.ConfigQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigsWithContext(ctx context.Context, request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerEventsRequest() (request *DescribeContainerEventsRequest) {
    request = &DescribeContainerEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerEvents")
    
    
    return
}

func NewDescribeContainerEventsResponse() (response *DescribeContainerEventsResponse) {
    response = &DescribeContainerEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContainerEvents
// 获取容器事件列表
//
// 参数限制
//
// 
//
// - 当类型是 instance 时，GroupId是必填项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_RESOURCETYPEERROR = "InvalidParameterValue.ResourceTypeError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeContainerEvents(request *DescribeContainerEventsRequest) (response *DescribeContainerEventsResponse, err error) {
    return c.DescribeContainerEventsWithContext(context.Background(), request)
}

// DescribeContainerEvents
// 获取容器事件列表
//
// 参数限制
//
// 
//
// - 当类型是 instance 时，GroupId是必填项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_RESOURCETYPEERROR = "InvalidParameterValue.ResourceTypeError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeContainerEventsWithContext(ctx context.Context, request *DescribeContainerEventsRequest) (response *DescribeContainerEventsResponse, err error) {
    if request == nil {
        request = NewDescribeContainerEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeContainerEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupAttributeRequest() (request *DescribeContainerGroupAttributeRequest) {
    request = &DescribeContainerGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupAttribute")
    
    
    return
}

func NewDescribeContainerGroupAttributeResponse() (response *DescribeContainerGroupAttributeResponse) {
    response = &DescribeContainerGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContainerGroupAttribute
// 获取部署组其他字段-用于前端并发调用
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CDISTATUSFAILED = "FailedOperation.CdiStatusFailed"
//  FAILEDOPERATION_CDISTATUSINVALID = "FailedOperation.CdiStatusInvalid"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupAttribute(request *DescribeContainerGroupAttributeRequest) (response *DescribeContainerGroupAttributeResponse, err error) {
    return c.DescribeContainerGroupAttributeWithContext(context.Background(), request)
}

// DescribeContainerGroupAttribute
// 获取部署组其他字段-用于前端并发调用
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CDISTATUSFAILED = "FailedOperation.CdiStatusFailed"
//  FAILEDOPERATION_CDISTATUSINVALID = "FailedOperation.CdiStatusInvalid"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupAttributeWithContext(ctx context.Context, request *DescribeContainerGroupAttributeRequest) (response *DescribeContainerGroupAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeContainerGroupAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupDeployInfoRequest() (request *DescribeContainerGroupDeployInfoRequest) {
    request = &DescribeContainerGroupDeployInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupDeployInfo")
    
    
    return
}

func NewDescribeContainerGroupDeployInfoResponse() (response *DescribeContainerGroupDeployInfoResponse) {
    response = &DescribeContainerGroupDeployInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContainerGroupDeployInfo
// 获取部署组详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APMGETCONFIGSFAILED = "FailedOperation.ApmGetConfigsFailed"
//  FAILEDOPERATION_CDISTATUSFAILED = "FailedOperation.CdiStatusFailed"
//  FAILEDOPERATION_CDISTATUSINVALID = "FailedOperation.CdiStatusInvalid"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CDISTATUSDATAINVALID = "InternalError.CdiStatusDataInvalid"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  MISSINGPARAMETER_SYSTEMPARAMETERREQUIRED = "MissingParameter.SystemParameterRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupDeployInfo(request *DescribeContainerGroupDeployInfoRequest) (response *DescribeContainerGroupDeployInfoResponse, err error) {
    return c.DescribeContainerGroupDeployInfoWithContext(context.Background(), request)
}

// DescribeContainerGroupDeployInfo
// 获取部署组详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APMGETCONFIGSFAILED = "FailedOperation.ApmGetConfigsFailed"
//  FAILEDOPERATION_CDISTATUSFAILED = "FailedOperation.CdiStatusFailed"
//  FAILEDOPERATION_CDISTATUSINVALID = "FailedOperation.CdiStatusInvalid"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CDISTATUSDATAINVALID = "InternalError.CdiStatusDataInvalid"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  MISSINGPARAMETER_SYSTEMPARAMETERREQUIRED = "MissingParameter.SystemParameterRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupDeployInfoWithContext(ctx context.Context, request *DescribeContainerGroupDeployInfoRequest) (response *DescribeContainerGroupDeployInfoResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupDeployInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeContainerGroupDeployInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerGroupDeployInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerGroupDeployInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupDetailRequest() (request *DescribeContainerGroupDetailRequest) {
    request = &DescribeContainerGroupDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupDetail")
    
    
    return
}

func NewDescribeContainerGroupDetailResponse() (response *DescribeContainerGroupDetailResponse) {
    response = &DescribeContainerGroupDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContainerGroupDetail
// 容器部署组详情（已废弃，请使用  DescribeContainerGroupDeployInfo）
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  MISSINGPARAMETER_SYSTEMPARAMETERREQUIRED = "MissingParameter.SystemParameterRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupDetail(request *DescribeContainerGroupDetailRequest) (response *DescribeContainerGroupDetailResponse, err error) {
    return c.DescribeContainerGroupDetailWithContext(context.Background(), request)
}

// DescribeContainerGroupDetail
// 容器部署组详情（已废弃，请使用  DescribeContainerGroupDeployInfo）
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  MISSINGPARAMETER_SYSTEMPARAMETERREQUIRED = "MissingParameter.SystemParameterRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupDetailWithContext(ctx context.Context, request *DescribeContainerGroupDetailRequest) (response *DescribeContainerGroupDetailResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeContainerGroupDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerGroupDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerGroupDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupsRequest() (request *DescribeContainerGroupsRequest) {
    request = &DescribeContainerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroups")
    
    
    return
}

func NewDescribeContainerGroupsResponse() (response *DescribeContainerGroupsResponse) {
    response = &DescribeContainerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContainerGroups
// 容器部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPLIMITVALUEINVALID = "InvalidParameterValue.ContainergroupLimitValueInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroups(request *DescribeContainerGroupsRequest) (response *DescribeContainerGroupsResponse, err error) {
    return c.DescribeContainerGroupsWithContext(context.Background(), request)
}

// DescribeContainerGroups
// 容器部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPLIMITVALUEINVALID = "InvalidParameterValue.ContainergroupLimitValueInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupsWithContext(ctx context.Context, request *DescribeContainerGroupsRequest) (response *DescribeContainerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeContainerGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCreateGatewayApiStatusRequest() (request *DescribeCreateGatewayApiStatusRequest) {
    request = &DescribeCreateGatewayApiStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeCreateGatewayApiStatus")
    
    
    return
}

func NewDescribeCreateGatewayApiStatusResponse() (response *DescribeCreateGatewayApiStatusResponse) {
    response = &DescribeCreateGatewayApiStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCreateGatewayApiStatus
// 查询一键导入API分组任务的状态
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
func (c *Client) DescribeCreateGatewayApiStatus(request *DescribeCreateGatewayApiStatusRequest) (response *DescribeCreateGatewayApiStatusResponse, err error) {
    return c.DescribeCreateGatewayApiStatusWithContext(context.Background(), request)
}

// DescribeCreateGatewayApiStatus
// 查询一键导入API分组任务的状态
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
func (c *Client) DescribeCreateGatewayApiStatusWithContext(ctx context.Context, request *DescribeCreateGatewayApiStatusRequest) (response *DescribeCreateGatewayApiStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCreateGatewayApiStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeCreateGatewayApiStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCreateGatewayApiStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCreateGatewayApiStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeliveryConfigRequest() (request *DescribeDeliveryConfigRequest) {
    request = &DescribeDeliveryConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeDeliveryConfig")
    
    
    return
}

func NewDescribeDeliveryConfigResponse() (response *DescribeDeliveryConfigResponse) {
    response = &DescribeDeliveryConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeliveryConfig
// 获取单个投递项配置信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGCLOUDPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgCloudParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGIDPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgIdParamError"
func (c *Client) DescribeDeliveryConfig(request *DescribeDeliveryConfigRequest) (response *DescribeDeliveryConfigResponse, err error) {
    return c.DescribeDeliveryConfigWithContext(context.Background(), request)
}

// DescribeDeliveryConfig
// 获取单个投递项配置信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGCLOUDPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgCloudParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGIDPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgIdParamError"
func (c *Client) DescribeDeliveryConfigWithContext(ctx context.Context, request *DescribeDeliveryConfigRequest) (response *DescribeDeliveryConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDeliveryConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeDeliveryConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeliveryConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeliveryConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeliveryConfigByGroupIdRequest() (request *DescribeDeliveryConfigByGroupIdRequest) {
    request = &DescribeDeliveryConfigByGroupIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeDeliveryConfigByGroupId")
    
    
    return
}

func NewDescribeDeliveryConfigByGroupIdResponse() (response *DescribeDeliveryConfigByGroupIdResponse) {
    response = &DescribeDeliveryConfigByGroupIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeliveryConfigByGroupId
// 用部署组id获取绑定信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeDeliveryConfigByGroupId(request *DescribeDeliveryConfigByGroupIdRequest) (response *DescribeDeliveryConfigByGroupIdResponse, err error) {
    return c.DescribeDeliveryConfigByGroupIdWithContext(context.Background(), request)
}

// DescribeDeliveryConfigByGroupId
// 用部署组id获取绑定信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeDeliveryConfigByGroupIdWithContext(ctx context.Context, request *DescribeDeliveryConfigByGroupIdRequest) (response *DescribeDeliveryConfigByGroupIdResponse, err error) {
    if request == nil {
        request = NewDescribeDeliveryConfigByGroupIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeDeliveryConfigByGroupId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeliveryConfigByGroupId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeliveryConfigByGroupIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeliveryConfigsRequest() (request *DescribeDeliveryConfigsRequest) {
    request = &DescribeDeliveryConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeDeliveryConfigs")
    
    
    return
}

func NewDescribeDeliveryConfigsResponse() (response *DescribeDeliveryConfigsResponse) {
    response = &DescribeDeliveryConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeliveryConfigs
// 获取多个投递项配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMBUSILOGCFGSCHEMAQUERYERROR = "FailedOperation.TsfApmBusiLogCfgSchemaQueryError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeDeliveryConfigs(request *DescribeDeliveryConfigsRequest) (response *DescribeDeliveryConfigsResponse, err error) {
    return c.DescribeDeliveryConfigsWithContext(context.Background(), request)
}

// DescribeDeliveryConfigs
// 获取多个投递项配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMBUSILOGCFGSCHEMAQUERYERROR = "FailedOperation.TsfApmBusiLogCfgSchemaQueryError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeDeliveryConfigsWithContext(ctx context.Context, request *DescribeDeliveryConfigsRequest) (response *DescribeDeliveryConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeDeliveryConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeDeliveryConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeliveryConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeliveryConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDownloadInfoRequest() (request *DescribeDownloadInfoRequest) {
    request = &DescribeDownloadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeDownloadInfo")
    
    
    return
}

func NewDescribeDownloadInfoResponse() (response *DescribeDownloadInfoResponse) {
    response = &DescribeDownloadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDownloadInfo
// TSF上传的程序包存放在腾讯云对象存储（COS）中，通过该API可以获取从COS下载程序包需要的信息，包括包所在的桶、存储路径、鉴权信息等，之后使用COS API（或SDK）进行下载。
//
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeDownloadInfo(request *DescribeDownloadInfoRequest) (response *DescribeDownloadInfoResponse, err error) {
    return c.DescribeDownloadInfoWithContext(context.Background(), request)
}

// DescribeDownloadInfo
// TSF上传的程序包存放在腾讯云对象存储（COS）中，通过该API可以获取从COS下载程序包需要的信息，包括包所在的桶、存储路径、鉴权信息等，之后使用COS API（或SDK）进行下载。
//
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeDownloadInfoWithContext(ctx context.Context, request *DescribeDownloadInfoRequest) (response *DescribeDownloadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeDownloadInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDownloadInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDownloadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnabledUnitRuleRequest() (request *DescribeEnabledUnitRuleRequest) {
    request = &DescribeEnabledUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeEnabledUnitRule")
    
    
    return
}

func NewDescribeEnabledUnitRuleResponse() (response *DescribeEnabledUnitRuleResponse) {
    response = &DescribeEnabledUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnabledUnitRule
// 查询生效的单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeEnabledUnitRule(request *DescribeEnabledUnitRuleRequest) (response *DescribeEnabledUnitRuleResponse, err error) {
    return c.DescribeEnabledUnitRuleWithContext(context.Background(), request)
}

// DescribeEnabledUnitRule
// 查询生效的单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeEnabledUnitRuleWithContext(ctx context.Context, request *DescribeEnabledUnitRuleRequest) (response *DescribeEnabledUnitRuleResponse, err error) {
    if request == nil {
        request = NewDescribeEnabledUnitRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeEnabledUnitRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnabledUnitRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnabledUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileConfigReleasesRequest() (request *DescribeFileConfigReleasesRequest) {
    request = &DescribeFileConfigReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigReleases")
    
    
    return
}

func NewDescribeFileConfigReleasesResponse() (response *DescribeFileConfigReleasesResponse) {
    response = &DescribeFileConfigReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileConfigReleases
// 查询文件配置项发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeFileConfigReleases(request *DescribeFileConfigReleasesRequest) (response *DescribeFileConfigReleasesResponse, err error) {
    return c.DescribeFileConfigReleasesWithContext(context.Background(), request)
}

// DescribeFileConfigReleases
// 查询文件配置项发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeFileConfigReleasesWithContext(ctx context.Context, request *DescribeFileConfigReleasesRequest) (response *DescribeFileConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribeFileConfigReleasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeFileConfigReleases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileConfigReleases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileConfigsRequest() (request *DescribeFileConfigsRequest) {
    request = &DescribeFileConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigs")
    
    
    return
}

func NewDescribeFileConfigsResponse() (response *DescribeFileConfigsResponse) {
    response = &DescribeFileConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileConfigs
// 查询文件配置项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
func (c *Client) DescribeFileConfigs(request *DescribeFileConfigsRequest) (response *DescribeFileConfigsResponse, err error) {
    return c.DescribeFileConfigsWithContext(context.Background(), request)
}

// DescribeFileConfigs
// 查询文件配置项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
func (c *Client) DescribeFileConfigsWithContext(ctx context.Context, request *DescribeFileConfigsRequest) (response *DescribeFileConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeFileConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeFileConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowLastBatchStateRequest() (request *DescribeFlowLastBatchStateRequest) {
    request = &DescribeFlowLastBatchStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFlowLastBatchState")
    
    
    return
}

func NewDescribeFlowLastBatchStateResponse() (response *DescribeFlowLastBatchStateResponse) {
    response = &DescribeFlowLastBatchStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFlowLastBatchState
// 查询工作流最新一个批次的状态信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeFlowLastBatchState(request *DescribeFlowLastBatchStateRequest) (response *DescribeFlowLastBatchStateResponse, err error) {
    return c.DescribeFlowLastBatchStateWithContext(context.Background(), request)
}

// DescribeFlowLastBatchState
// 查询工作流最新一个批次的状态信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeFlowLastBatchStateWithContext(ctx context.Context, request *DescribeFlowLastBatchStateRequest) (response *DescribeFlowLastBatchStateResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLastBatchStateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeFlowLastBatchState")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowLastBatchState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowLastBatchStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayAllGroupApisRequest() (request *DescribeGatewayAllGroupApisRequest) {
    request = &DescribeGatewayAllGroupApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayAllGroupApis")
    
    
    return
}

func NewDescribeGatewayAllGroupApisResponse() (response *DescribeGatewayAllGroupApisResponse) {
    response = &DescribeGatewayAllGroupApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGatewayAllGroupApis
// 查询网关所有分组下Api列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_DEPLOYGROUPNOTEXISTS = "InvalidParameterValue.DeployGroupNotExists"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGatewayAllGroupApis(request *DescribeGatewayAllGroupApisRequest) (response *DescribeGatewayAllGroupApisResponse, err error) {
    return c.DescribeGatewayAllGroupApisWithContext(context.Background(), request)
}

// DescribeGatewayAllGroupApis
// 查询网关所有分组下Api列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_DEPLOYGROUPNOTEXISTS = "InvalidParameterValue.DeployGroupNotExists"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGatewayAllGroupApisWithContext(ctx context.Context, request *DescribeGatewayAllGroupApisRequest) (response *DescribeGatewayAllGroupApisResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayAllGroupApisRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGatewayAllGroupApis")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayAllGroupApis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayAllGroupApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayApisRequest() (request *DescribeGatewayApisRequest) {
    request = &DescribeGatewayApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayApis")
    
    
    return
}

func NewDescribeGatewayApisResponse() (response *DescribeGatewayApisResponse) {
    response = &DescribeGatewayApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGatewayApis
// 查询API分组下的Api列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribeGatewayApis(request *DescribeGatewayApisRequest) (response *DescribeGatewayApisResponse, err error) {
    return c.DescribeGatewayApisWithContext(context.Background(), request)
}

// DescribeGatewayApis
// 查询API分组下的Api列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribeGatewayApisWithContext(ctx context.Context, request *DescribeGatewayApisRequest) (response *DescribeGatewayApisResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayApisRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGatewayApis")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayApis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayMonitorOverviewRequest() (request *DescribeGatewayMonitorOverviewRequest) {
    request = &DescribeGatewayMonitorOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayMonitorOverview")
    
    
    return
}

func NewDescribeGatewayMonitorOverviewResponse() (response *DescribeGatewayMonitorOverviewResponse) {
    response = &DescribeGatewayMonitorOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGatewayMonitorOverview
// 查询网关监控概览
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGatewayMonitorOverview(request *DescribeGatewayMonitorOverviewRequest) (response *DescribeGatewayMonitorOverviewResponse, err error) {
    return c.DescribeGatewayMonitorOverviewWithContext(context.Background(), request)
}

// DescribeGatewayMonitorOverview
// 查询网关监控概览
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGatewayMonitorOverviewWithContext(ctx context.Context, request *DescribeGatewayMonitorOverviewRequest) (response *DescribeGatewayMonitorOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayMonitorOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGatewayMonitorOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayMonitorOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatewayMonitorOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
    request = &DescribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroup")
    
    
    return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
    response = &DescribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroup
// 查询虚拟机部署组详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    return c.DescribeGroupWithContext(context.Background(), request)
}

// DescribeGroup
// 查询虚拟机部署组详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupWithContext(ctx context.Context, request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupAttributeRequest() (request *DescribeGroupAttributeRequest) {
    request = &DescribeGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupAttribute")
    
    
    return
}

func NewDescribeGroupAttributeResponse() (response *DescribeGroupAttributeResponse) {
    response = &DescribeGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupAttribute
// 获取部署组其他属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribeGroupAttribute(request *DescribeGroupAttributeRequest) (response *DescribeGroupAttributeResponse, err error) {
    return c.DescribeGroupAttributeWithContext(context.Background(), request)
}

// DescribeGroupAttribute
// 获取部署组其他属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribeGroupAttributeWithContext(ctx context.Context, request *DescribeGroupAttributeRequest) (response *DescribeGroupAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeGroupAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGroupAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupBindedGatewaysRequest() (request *DescribeGroupBindedGatewaysRequest) {
    request = &DescribeGroupBindedGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupBindedGateways")
    
    
    return
}

func NewDescribeGroupBindedGatewaysResponse() (response *DescribeGroupBindedGatewaysResponse) {
    response = &DescribeGroupBindedGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupBindedGateways
// 查询某个API分组已绑定的网关部署组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGroupBindedGateways(request *DescribeGroupBindedGatewaysRequest) (response *DescribeGroupBindedGatewaysResponse, err error) {
    return c.DescribeGroupBindedGatewaysWithContext(context.Background(), request)
}

// DescribeGroupBindedGateways
// 查询某个API分组已绑定的网关部署组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGroupBindedGatewaysWithContext(ctx context.Context, request *DescribeGroupBindedGatewaysRequest) (response *DescribeGroupBindedGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeGroupBindedGatewaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGroupBindedGateways")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupBindedGateways require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupBindedGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupBusinessLogConfigsRequest() (request *DescribeGroupBusinessLogConfigsRequest) {
    request = &DescribeGroupBusinessLogConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupBusinessLogConfigs")
    
    
    return
}

func NewDescribeGroupBusinessLogConfigsResponse() (response *DescribeGroupBusinessLogConfigsResponse) {
    response = &DescribeGroupBusinessLogConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupBusinessLogConfigs
// 查询分组管理日志配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeGroupBusinessLogConfigs(request *DescribeGroupBusinessLogConfigsRequest) (response *DescribeGroupBusinessLogConfigsResponse, err error) {
    return c.DescribeGroupBusinessLogConfigsWithContext(context.Background(), request)
}

// DescribeGroupBusinessLogConfigs
// 查询分组管理日志配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeGroupBusinessLogConfigsWithContext(ctx context.Context, request *DescribeGroupBusinessLogConfigsRequest) (response *DescribeGroupBusinessLogConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupBusinessLogConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGroupBusinessLogConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupBusinessLogConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupBusinessLogConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupGatewaysRequest() (request *DescribeGroupGatewaysRequest) {
    request = &DescribeGroupGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupGateways")
    
    
    return
}

func NewDescribeGroupGatewaysResponse() (response *DescribeGroupGatewaysResponse) {
    response = &DescribeGroupGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupGateways
// 查询某个网关绑定的API 分组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeGroupGateways(request *DescribeGroupGatewaysRequest) (response *DescribeGroupGatewaysResponse, err error) {
    return c.DescribeGroupGatewaysWithContext(context.Background(), request)
}

// DescribeGroupGateways
// 查询某个网关绑定的API 分组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeGroupGatewaysWithContext(ctx context.Context, request *DescribeGroupGatewaysRequest) (response *DescribeGroupGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeGroupGatewaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGroupGateways")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupGateways require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupInstancesRequest() (request *DescribeGroupInstancesRequest) {
    request = &DescribeGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupInstances")
    
    
    return
}

func NewDescribeGroupInstancesResponse() (response *DescribeGroupInstancesResponse) {
    response = &DescribeGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupInstances
// 查询虚拟机部署组云主机列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CAMROLEREQUESTERROR = "InternalError.CamRoleRequestError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupInstances(request *DescribeGroupInstancesRequest) (response *DescribeGroupInstancesResponse, err error) {
    return c.DescribeGroupInstancesWithContext(context.Background(), request)
}

// DescribeGroupInstances
// 查询虚拟机部署组云主机列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CAMROLEREQUESTERROR = "InternalError.CamRoleRequestError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupInstancesWithContext(ctx context.Context, request *DescribeGroupInstancesRequest) (response *DescribeGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeGroupInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGroupInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupReleaseRequest() (request *DescribeGroupReleaseRequest) {
    request = &DescribeGroupReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupRelease")
    
    
    return
}

func NewDescribeGroupReleaseResponse() (response *DescribeGroupReleaseResponse) {
    response = &DescribeGroupReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupRelease
// 查询部署组相关的发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"
//  FAILEDOPERATION_CONFIGRELEASEQUERYFAILED = "FailedOperation.ConfigReleaseQueryFailed"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeGroupRelease(request *DescribeGroupReleaseRequest) (response *DescribeGroupReleaseResponse, err error) {
    return c.DescribeGroupReleaseWithContext(context.Background(), request)
}

// DescribeGroupRelease
// 查询部署组相关的发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"
//  FAILEDOPERATION_CONFIGRELEASEQUERYFAILED = "FailedOperation.ConfigReleaseQueryFailed"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeGroupReleaseWithContext(ctx context.Context, request *DescribeGroupReleaseRequest) (response *DescribeGroupReleaseResponse, err error) {
    if request == nil {
        request = NewDescribeGroupReleaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGroupRelease")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupUseDetailRequest() (request *DescribeGroupUseDetailRequest) {
    request = &DescribeGroupUseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupUseDetail")
    
    
    return
}

func NewDescribeGroupUseDetailResponse() (response *DescribeGroupUseDetailResponse) {
    response = &DescribeGroupUseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupUseDetail
// 查询网关分组监控明细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGroupUseDetail(request *DescribeGroupUseDetailRequest) (response *DescribeGroupUseDetailResponse, err error) {
    return c.DescribeGroupUseDetailWithContext(context.Background(), request)
}

// DescribeGroupUseDetail
// 查询网关分组监控明细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGroupUseDetailWithContext(ctx context.Context, request *DescribeGroupUseDetailRequest) (response *DescribeGroupUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGroupUseDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGroupUseDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupUseDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupUseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupsRequest() (request *DescribeGroupsRequest) {
    request = &DescribeGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroups")
    
    
    return
}

func NewDescribeGroupsResponse() (response *DescribeGroupsResponse) {
    response = &DescribeGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroups
// 获取虚拟机部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_GROUPPAGELIMITINVALID = "InvalidParameterValue.GroupPageLimitInvalid"
//  INVALIDPARAMETERVALUE_GROUPSTATUSINVALID = "InvalidParameterValue.GroupStatusInvalid"
//  MISSINGPARAMETER_CLUSTERIDREQUIRED = "MissingParameter.ClusterIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    return c.DescribeGroupsWithContext(context.Background(), request)
}

// DescribeGroups
// 获取虚拟机部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_GROUPPAGELIMITINVALID = "InvalidParameterValue.GroupPageLimitInvalid"
//  INVALIDPARAMETERVALUE_GROUPSTATUSINVALID = "InvalidParameterValue.GroupStatusInvalid"
//  MISSINGPARAMETER_CLUSTERIDREQUIRED = "MissingParameter.ClusterIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupsWithContext(ctx context.Context, request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupsWithPluginRequest() (request *DescribeGroupsWithPluginRequest) {
    request = &DescribeGroupsWithPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupsWithPlugin")
    
    
    return
}

func NewDescribeGroupsWithPluginResponse() (response *DescribeGroupsWithPluginResponse) {
    response = &DescribeGroupsWithPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupsWithPlugin
// 查询某个插件下绑定或未绑定的API分组
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_GROUPPAGELIMITINVALID = "InvalidParameterValue.GroupPageLimitInvalid"
//  INVALIDPARAMETERVALUE_GROUPSTATUSINVALID = "InvalidParameterValue.GroupStatusInvalid"
//  MISSINGPARAMETER_CLUSTERIDREQUIRED = "MissingParameter.ClusterIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupsWithPlugin(request *DescribeGroupsWithPluginRequest) (response *DescribeGroupsWithPluginResponse, err error) {
    return c.DescribeGroupsWithPluginWithContext(context.Background(), request)
}

// DescribeGroupsWithPlugin
// 查询某个插件下绑定或未绑定的API分组
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_GROUPPAGELIMITINVALID = "InvalidParameterValue.GroupPageLimitInvalid"
//  INVALIDPARAMETERVALUE_GROUPSTATUSINVALID = "InvalidParameterValue.GroupStatusInvalid"
//  MISSINGPARAMETER_CLUSTERIDREQUIRED = "MissingParameter.ClusterIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupsWithPluginWithContext(ctx context.Context, request *DescribeGroupsWithPluginRequest) (response *DescribeGroupsWithPluginResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsWithPluginRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeGroupsWithPlugin")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupsWithPlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupsWithPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRepositoryRequest() (request *DescribeImageRepositoryRequest) {
    request = &DescribeImageRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageRepository")
    
    
    return
}

func NewDescribeImageRepositoryResponse() (response *DescribeImageRepositoryResponse) {
    response = &DescribeImageRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageRepository
// 查询镜像仓库列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeImageRepository(request *DescribeImageRepositoryRequest) (response *DescribeImageRepositoryResponse, err error) {
    return c.DescribeImageRepositoryWithContext(context.Background(), request)
}

// DescribeImageRepository
// 查询镜像仓库列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeImageRepositoryWithContext(ctx context.Context, request *DescribeImageRepositoryRequest) (response *DescribeImageRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeImageRepositoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeImageRepository")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageTagsRequest() (request *DescribeImageTagsRequest) {
    request = &DescribeImageTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageTags")
    
    
    return
}

func NewDescribeImageTagsResponse() (response *DescribeImageTagsResponse) {
    response = &DescribeImageTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageTags
// 镜像版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CAMROLEREQUESTERROR = "InternalError.CamRoleRequestError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeImageTags(request *DescribeImageTagsRequest) (response *DescribeImageTagsResponse, err error) {
    return c.DescribeImageTagsWithContext(context.Background(), request)
}

// DescribeImageTags
// 镜像版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CAMROLEREQUESTERROR = "InternalError.CamRoleRequestError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeImageTagsWithContext(ctx context.Context, request *DescribeImageTagsRequest) (response *DescribeImageTagsResponse, err error) {
    if request == nil {
        request = NewDescribeImageTagsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeImageTags")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInovcationIndicatorsRequest() (request *DescribeInovcationIndicatorsRequest) {
    request = &DescribeInovcationIndicatorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeInovcationIndicators")
    
    
    return
}

func NewDescribeInovcationIndicatorsResponse() (response *DescribeInovcationIndicatorsResponse) {
    response = &DescribeInovcationIndicatorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInovcationIndicators
// 废弃接口
//
// 
//
// 废弃
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeInovcationIndicators(request *DescribeInovcationIndicatorsRequest) (response *DescribeInovcationIndicatorsResponse, err error) {
    return c.DescribeInovcationIndicatorsWithContext(context.Background(), request)
}

// DescribeInovcationIndicators
// 废弃接口
//
// 
//
// 废弃
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeInovcationIndicatorsWithContext(ctx context.Context, request *DescribeInovcationIndicatorsRequest) (response *DescribeInovcationIndicatorsResponse, err error) {
    if request == nil {
        request = NewDescribeInovcationIndicatorsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeInovcationIndicators")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInovcationIndicators require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInovcationIndicatorsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 无
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 无
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationMetricDataCurveRequest() (request *DescribeInvocationMetricDataCurveRequest) {
    request = &DescribeInvocationMetricDataCurveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationMetricDataCurve")
    
    
    return
}

func NewDescribeInvocationMetricDataCurveResponse() (response *DescribeInvocationMetricDataCurveResponse) {
    response = &DescribeInvocationMetricDataCurveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInvocationMetricDataCurve
// 查询调用指标数据变化曲线
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeInvocationMetricDataCurve(request *DescribeInvocationMetricDataCurveRequest) (response *DescribeInvocationMetricDataCurveResponse, err error) {
    return c.DescribeInvocationMetricDataCurveWithContext(context.Background(), request)
}

// DescribeInvocationMetricDataCurve
// 查询调用指标数据变化曲线
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeInvocationMetricDataCurveWithContext(ctx context.Context, request *DescribeInvocationMetricDataCurveRequest) (response *DescribeInvocationMetricDataCurveResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationMetricDataCurveRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeInvocationMetricDataCurve")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvocationMetricDataCurve require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvocationMetricDataCurveResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationMetricDataDimensionRequest() (request *DescribeInvocationMetricDataDimensionRequest) {
    request = &DescribeInvocationMetricDataDimensionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationMetricDataDimension")
    
    
    return
}

func NewDescribeInvocationMetricDataDimensionResponse() (response *DescribeInvocationMetricDataDimensionResponse) {
    response = &DescribeInvocationMetricDataDimensionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInvocationMetricDataDimension
// 查询维度
//
// 可能返回的错误码:
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
func (c *Client) DescribeInvocationMetricDataDimension(request *DescribeInvocationMetricDataDimensionRequest) (response *DescribeInvocationMetricDataDimensionResponse, err error) {
    return c.DescribeInvocationMetricDataDimensionWithContext(context.Background(), request)
}

// DescribeInvocationMetricDataDimension
// 查询维度
//
// 可能返回的错误码:
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
func (c *Client) DescribeInvocationMetricDataDimensionWithContext(ctx context.Context, request *DescribeInvocationMetricDataDimensionRequest) (response *DescribeInvocationMetricDataDimensionResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationMetricDataDimensionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeInvocationMetricDataDimension")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvocationMetricDataDimension require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvocationMetricDataDimensionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationMetricDataPointRequest() (request *DescribeInvocationMetricDataPointRequest) {
    request = &DescribeInvocationMetricDataPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationMetricDataPoint")
    
    
    return
}

func NewDescribeInvocationMetricDataPointResponse() (response *DescribeInvocationMetricDataPointResponse) {
    response = &DescribeInvocationMetricDataPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInvocationMetricDataPoint
// 查询单值指标维度
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMINTERNALERROR = "FailedOperation.TsfApmInternalError"
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeInvocationMetricDataPoint(request *DescribeInvocationMetricDataPointRequest) (response *DescribeInvocationMetricDataPointResponse, err error) {
    return c.DescribeInvocationMetricDataPointWithContext(context.Background(), request)
}

// DescribeInvocationMetricDataPoint
// 查询单值指标维度
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMINTERNALERROR = "FailedOperation.TsfApmInternalError"
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeInvocationMetricDataPointWithContext(ctx context.Context, request *DescribeInvocationMetricDataPointRequest) (response *DescribeInvocationMetricDataPointResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationMetricDataPointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeInvocationMetricDataPoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvocationMetricDataPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvocationMetricDataPointResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInvocationMetricScatterPlotRequest() (request *DescribeInvocationMetricScatterPlotRequest) {
    request = &DescribeInvocationMetricScatterPlotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationMetricScatterPlot")
    
    
    return
}

func NewDescribeInvocationMetricScatterPlotResponse() (response *DescribeInvocationMetricScatterPlotResponse) {
    response = &DescribeInvocationMetricScatterPlotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInvocationMetricScatterPlot
// 查询调用指标数据散点图
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeInvocationMetricScatterPlot(request *DescribeInvocationMetricScatterPlotRequest) (response *DescribeInvocationMetricScatterPlotResponse, err error) {
    return c.DescribeInvocationMetricScatterPlotWithContext(context.Background(), request)
}

// DescribeInvocationMetricScatterPlot
// 查询调用指标数据散点图
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeInvocationMetricScatterPlotWithContext(ctx context.Context, request *DescribeInvocationMetricScatterPlotRequest) (response *DescribeInvocationMetricScatterPlotResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationMetricScatterPlotRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeInvocationMetricScatterPlot")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvocationMetricScatterPlot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvocationMetricScatterPlotResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJvmMonitorRequest() (request *DescribeJvmMonitorRequest) {
    request = &DescribeJvmMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeJvmMonitor")
    
    
    return
}

func NewDescribeJvmMonitorResponse() (response *DescribeJvmMonitorResponse) {
    response = &DescribeJvmMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJvmMonitor
// 查询java实例jvm监控数据,返回数据可选
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFCMONITORCTSDBCLIENTREQUESTFAIL = "FailedOperation.TsfCmonitorCtsdbClientRequestFail"
//  FAILEDOPERATION_TSFMONITORWAITEDTIMEOUT = "FailedOperation.TsfMonitorWaitedTimeout"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFMONITORDATEPARSEFAILED = "InternalError.TsfMonitorDateParseFailed"
//  INTERNALERROR_TSFMONITORINTERNALERROR = "InternalError.TsfMonitorInternalError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_TSFMONITORREQUESTPARAMILLEGAL = "InvalidParameter.TsfMonitorRequestParamIllegal"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeJvmMonitor(request *DescribeJvmMonitorRequest) (response *DescribeJvmMonitorResponse, err error) {
    return c.DescribeJvmMonitorWithContext(context.Background(), request)
}

// DescribeJvmMonitor
// 查询java实例jvm监控数据,返回数据可选
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFCMONITORCTSDBCLIENTREQUESTFAIL = "FailedOperation.TsfCmonitorCtsdbClientRequestFail"
//  FAILEDOPERATION_TSFMONITORWAITEDTIMEOUT = "FailedOperation.TsfMonitorWaitedTimeout"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFMONITORDATEPARSEFAILED = "InternalError.TsfMonitorDateParseFailed"
//  INTERNALERROR_TSFMONITORINTERNALERROR = "InternalError.TsfMonitorInternalError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_TSFMONITORREQUESTPARAMILLEGAL = "InvalidParameter.TsfMonitorRequestParamIllegal"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeJvmMonitorWithContext(ctx context.Context, request *DescribeJvmMonitorRequest) (response *DescribeJvmMonitorResponse, err error) {
    if request == nil {
        request = NewDescribeJvmMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeJvmMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJvmMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJvmMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLaneRulesRequest() (request *DescribeLaneRulesRequest) {
    request = &DescribeLaneRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeLaneRules")
    
    
    return
}

func NewDescribeLaneRulesResponse() (response *DescribeLaneRulesResponse) {
    response = &DescribeLaneRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLaneRules
// 查询泳道规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeLaneRules(request *DescribeLaneRulesRequest) (response *DescribeLaneRulesResponse, err error) {
    return c.DescribeLaneRulesWithContext(context.Background(), request)
}

// DescribeLaneRules
// 查询泳道规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeLaneRulesWithContext(ctx context.Context, request *DescribeLaneRulesRequest) (response *DescribeLaneRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLaneRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeLaneRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLaneRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLaneRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLanesRequest() (request *DescribeLanesRequest) {
    request = &DescribeLanesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeLanes")
    
    
    return
}

func NewDescribeLanesResponse() (response *DescribeLanesResponse) {
    response = &DescribeLanesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLanes
// 查询泳道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeLanes(request *DescribeLanesRequest) (response *DescribeLanesResponse, err error) {
    return c.DescribeLanesWithContext(context.Background(), request)
}

// DescribeLanes
// 查询泳道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeLanesWithContext(ctx context.Context, request *DescribeLanesRequest) (response *DescribeLanesResponse, err error) {
    if request == nil {
        request = NewDescribeLanesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeLanes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLanes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLanesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMicroserviceRequest() (request *DescribeMicroserviceRequest) {
    request = &DescribeMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservice")
    
    
    return
}

func NewDescribeMicroserviceResponse() (response *DescribeMicroserviceResponse) {
    response = &DescribeMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMicroservice
// 查询微服务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_SERVICEIDREQUIRED = "MissingParameter.ServiceIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeMicroservice(request *DescribeMicroserviceRequest) (response *DescribeMicroserviceResponse, err error) {
    return c.DescribeMicroserviceWithContext(context.Background(), request)
}

// DescribeMicroservice
// 查询微服务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_SERVICEIDREQUIRED = "MissingParameter.ServiceIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeMicroserviceWithContext(ctx context.Context, request *DescribeMicroserviceRequest) (response *DescribeMicroserviceResponse, err error) {
    if request == nil {
        request = NewDescribeMicroserviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeMicroservice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMicroservice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMicroservicesRequest() (request *DescribeMicroservicesRequest) {
    request = &DescribeMicroservicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservices")
    
    
    return
}

func NewDescribeMicroservicesResponse() (response *DescribeMicroservicesResponse) {
    response = &DescribeMicroservicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMicroservices
// 获取微服务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_SERVICENOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ServiceNotExistsOrPermissionDenied"
//  MISSINGPARAMETER_NAMESPACEIDREQUIRED = "MissingParameter.NamespaceIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeMicroservices(request *DescribeMicroservicesRequest) (response *DescribeMicroservicesResponse, err error) {
    return c.DescribeMicroservicesWithContext(context.Background(), request)
}

// DescribeMicroservices
// 获取微服务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_SERVICENOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ServiceNotExistsOrPermissionDenied"
//  MISSINGPARAMETER_NAMESPACEIDREQUIRED = "MissingParameter.NamespaceIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeMicroservicesWithContext(ctx context.Context, request *DescribeMicroservicesRequest) (response *DescribeMicroservicesResponse, err error) {
    if request == nil {
        request = NewDescribeMicroservicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeMicroservices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMicroservices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMicroservicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMicroservicesByGroupIdsRequest() (request *DescribeMicroservicesByGroupIdsRequest) {
    request = &DescribeMicroservicesByGroupIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservicesByGroupIds")
    
    
    return
}

func NewDescribeMicroservicesByGroupIdsResponse() (response *DescribeMicroservicesByGroupIdsResponse) {
    response = &DescribeMicroservicesByGroupIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMicroservicesByGroupIds
// 通过部署组ID获取微服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFMSERROR = "FailedOperation.TsfMsError"
//  INTERNALERROR_TSFMSERROR = "InternalError.TsfMsError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_NAMESPACEIDREQUIRED = "MissingParameter.NamespaceIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeMicroservicesByGroupIds(request *DescribeMicroservicesByGroupIdsRequest) (response *DescribeMicroservicesByGroupIdsResponse, err error) {
    return c.DescribeMicroservicesByGroupIdsWithContext(context.Background(), request)
}

// DescribeMicroservicesByGroupIds
// 通过部署组ID获取微服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFMSERROR = "FailedOperation.TsfMsError"
//  INTERNALERROR_TSFMSERROR = "InternalError.TsfMsError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_NAMESPACEIDREQUIRED = "MissingParameter.NamespaceIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeMicroservicesByGroupIdsWithContext(ctx context.Context, request *DescribeMicroservicesByGroupIdsRequest) (response *DescribeMicroservicesByGroupIdsResponse, err error) {
    if request == nil {
        request = NewDescribeMicroservicesByGroupIdsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeMicroservicesByGroupIds")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMicroservicesByGroupIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMicroservicesByGroupIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsApiListRequest() (request *DescribeMsApiListRequest) {
    request = &DescribeMsApiListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMsApiList")
    
    
    return
}

func NewDescribeMsApiListResponse() (response *DescribeMsApiListResponse) {
    response = &DescribeMsApiListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMsApiList
// 查询服务API列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeMsApiList(request *DescribeMsApiListRequest) (response *DescribeMsApiListResponse, err error) {
    return c.DescribeMsApiListWithContext(context.Background(), request)
}

// DescribeMsApiList
// 查询服务API列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeMsApiListWithContext(ctx context.Context, request *DescribeMsApiListRequest) (response *DescribeMsApiListResponse, err error) {
    if request == nil {
        request = NewDescribeMsApiListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeMsApiList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMsApiList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMsApiListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewInvocationRequest() (request *DescribeOverviewInvocationRequest) {
    request = &DescribeOverviewInvocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeOverviewInvocation")
    
    
    return
}

func NewDescribeOverviewInvocationResponse() (response *DescribeOverviewInvocationResponse) {
    response = &DescribeOverviewInvocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOverviewInvocation
// 服务调用监控统计概览
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeOverviewInvocation(request *DescribeOverviewInvocationRequest) (response *DescribeOverviewInvocationResponse, err error) {
    return c.DescribeOverviewInvocationWithContext(context.Background(), request)
}

// DescribeOverviewInvocation
// 服务调用监控统计概览
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeOverviewInvocationWithContext(ctx context.Context, request *DescribeOverviewInvocationRequest) (response *DescribeOverviewInvocationResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewInvocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeOverviewInvocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewInvocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewInvocationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePathRewriteRequest() (request *DescribePathRewriteRequest) {
    request = &DescribePathRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePathRewrite")
    
    
    return
}

func NewDescribePathRewriteResponse() (response *DescribePathRewriteResponse) {
    response = &DescribePathRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePathRewrite
// 查询路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribePathRewrite(request *DescribePathRewriteRequest) (response *DescribePathRewriteResponse, err error) {
    return c.DescribePathRewriteWithContext(context.Background(), request)
}

// DescribePathRewrite
// 查询路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribePathRewriteWithContext(ctx context.Context, request *DescribePathRewriteRequest) (response *DescribePathRewriteResponse, err error) {
    if request == nil {
        request = NewDescribePathRewriteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePathRewrite")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePathRewrite require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePathRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePathRewritesRequest() (request *DescribePathRewritesRequest) {
    request = &DescribePathRewritesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePathRewrites")
    
    
    return
}

func NewDescribePathRewritesResponse() (response *DescribePathRewritesResponse) {
    response = &DescribePathRewritesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePathRewrites
// 查询路径重写列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePathRewrites(request *DescribePathRewritesRequest) (response *DescribePathRewritesResponse, err error) {
    return c.DescribePathRewritesWithContext(context.Background(), request)
}

// DescribePathRewrites
// 查询路径重写列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePathRewritesWithContext(ctx context.Context, request *DescribePathRewritesRequest) (response *DescribePathRewritesResponse, err error) {
    if request == nil {
        request = NewDescribePathRewritesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePathRewrites")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePathRewrites require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePathRewritesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePkgsRequest() (request *DescribePkgsRequest) {
    request = &DescribePkgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePkgs")
    
    
    return
}

func NewDescribePkgsResponse() (response *DescribePkgsResponse) {
    response = &DescribePkgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePkgs
// 无
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribePkgs(request *DescribePkgsRequest) (response *DescribePkgsResponse, err error) {
    return c.DescribePkgsWithContext(context.Background(), request)
}

// DescribePkgs
// 无
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribePkgsWithContext(ctx context.Context, request *DescribePkgsRequest) (response *DescribePkgsResponse, err error) {
    if request == nil {
        request = NewDescribePkgsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePkgs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePkgs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePkgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginInstancesRequest() (request *DescribePluginInstancesRequest) {
    request = &DescribePluginInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePluginInstances")
    
    
    return
}

func NewDescribePluginInstancesResponse() (response *DescribePluginInstancesResponse) {
    response = &DescribePluginInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePluginInstances
// 分页查询网关分组/API绑定（或未绑定）的插件列表
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribePluginInstances(request *DescribePluginInstancesRequest) (response *DescribePluginInstancesResponse, err error) {
    return c.DescribePluginInstancesWithContext(context.Background(), request)
}

// DescribePluginInstances
// 分页查询网关分组/API绑定（或未绑定）的插件列表
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribePluginInstancesWithContext(ctx context.Context, request *DescribePluginInstancesRequest) (response *DescribePluginInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePluginInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePluginInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePluginInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePodInstancesRequest() (request *DescribePodInstancesRequest) {
    request = &DescribePodInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePodInstances")
    
    
    return
}

func NewDescribePodInstancesResponse() (response *DescribePodInstancesResponse) {
    response = &DescribePodInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePodInstances
// 获取部署组实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPLIMITVALUEINVALID = "InvalidParameterValue.ContainergroupLimitValueInvalid"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePodInstances(request *DescribePodInstancesRequest) (response *DescribePodInstancesResponse, err error) {
    return c.DescribePodInstancesWithContext(context.Background(), request)
}

// DescribePodInstances
// 获取部署组实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPLIMITVALUEINVALID = "InvalidParameterValue.ContainergroupLimitValueInvalid"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePodInstancesWithContext(ctx context.Context, request *DescribePodInstancesRequest) (response *DescribePodInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePodInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePodInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePodInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePodInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProgramsRequest() (request *DescribeProgramsRequest) {
    request = &DescribeProgramsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePrograms")
    
    
    return
}

func NewDescribeProgramsResponse() (response *DescribeProgramsResponse) {
    response = &DescribeProgramsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrograms
// 查询数据集列表
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePrograms(request *DescribeProgramsRequest) (response *DescribeProgramsResponse, err error) {
    return c.DescribeProgramsWithContext(context.Background(), request)
}

// DescribePrograms
// 查询数据集列表
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeProgramsWithContext(ctx context.Context, request *DescribeProgramsRequest) (response *DescribeProgramsResponse, err error) {
    if request == nil {
        request = NewDescribeProgramsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePrograms")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrograms require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProgramsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigRequest() (request *DescribePublicConfigRequest) {
    request = &DescribePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfig")
    
    
    return
}

func NewDescribePublicConfigResponse() (response *DescribePublicConfigResponse) {
    response = &DescribePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicConfig
// 查询公共配置（单条）
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfig(request *DescribePublicConfigRequest) (response *DescribePublicConfigResponse, err error) {
    return c.DescribePublicConfigWithContext(context.Background(), request)
}

// DescribePublicConfig
// 查询公共配置（单条）
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfigWithContext(ctx context.Context, request *DescribePublicConfigRequest) (response *DescribePublicConfigResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePublicConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigReleaseLogsRequest() (request *DescribePublicConfigReleaseLogsRequest) {
    request = &DescribePublicConfigReleaseLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigReleaseLogs")
    
    
    return
}

func NewDescribePublicConfigReleaseLogsResponse() (response *DescribePublicConfigReleaseLogsResponse) {
    response = &DescribePublicConfigReleaseLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicConfigReleaseLogs
// 查询公共配置发布历史
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfigReleaseLogs(request *DescribePublicConfigReleaseLogsRequest) (response *DescribePublicConfigReleaseLogsResponse, err error) {
    return c.DescribePublicConfigReleaseLogsWithContext(context.Background(), request)
}

// DescribePublicConfigReleaseLogs
// 查询公共配置发布历史
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfigReleaseLogsWithContext(ctx context.Context, request *DescribePublicConfigReleaseLogsRequest) (response *DescribePublicConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigReleaseLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePublicConfigReleaseLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicConfigReleaseLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigReleasesRequest() (request *DescribePublicConfigReleasesRequest) {
    request = &DescribePublicConfigReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigReleases")
    
    
    return
}

func NewDescribePublicConfigReleasesResponse() (response *DescribePublicConfigReleasesResponse) {
    response = &DescribePublicConfigReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicConfigReleases
// 查询公共配置发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"
//  FAILEDOPERATION_CONFIGRELEASEQUERYFAILED = "FailedOperation.ConfigReleaseQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfigReleases(request *DescribePublicConfigReleasesRequest) (response *DescribePublicConfigReleasesResponse, err error) {
    return c.DescribePublicConfigReleasesWithContext(context.Background(), request)
}

// DescribePublicConfigReleases
// 查询公共配置发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"
//  FAILEDOPERATION_CONFIGRELEASEQUERYFAILED = "FailedOperation.ConfigReleaseQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfigReleasesWithContext(ctx context.Context, request *DescribePublicConfigReleasesRequest) (response *DescribePublicConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigReleasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePublicConfigReleases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicConfigReleases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigSummaryRequest() (request *DescribePublicConfigSummaryRequest) {
    request = &DescribePublicConfigSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigSummary")
    
    
    return
}

func NewDescribePublicConfigSummaryResponse() (response *DescribePublicConfigSummaryResponse) {
    response = &DescribePublicConfigSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicConfigSummary
// 查询公共配置汇总列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePublicConfigSummary(request *DescribePublicConfigSummaryRequest) (response *DescribePublicConfigSummaryResponse, err error) {
    return c.DescribePublicConfigSummaryWithContext(context.Background(), request)
}

// DescribePublicConfigSummary
// 查询公共配置汇总列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePublicConfigSummaryWithContext(ctx context.Context, request *DescribePublicConfigSummaryRequest) (response *DescribePublicConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePublicConfigSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicConfigSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigsRequest() (request *DescribePublicConfigsRequest) {
    request = &DescribePublicConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigs")
    
    
    return
}

func NewDescribePublicConfigsResponse() (response *DescribePublicConfigsResponse) {
    response = &DescribePublicConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicConfigs
// 查询公共配置项列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePublicConfigs(request *DescribePublicConfigsRequest) (response *DescribePublicConfigsResponse, err error) {
    return c.DescribePublicConfigsWithContext(context.Background(), request)
}

// DescribePublicConfigs
// 查询公共配置项列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePublicConfigsWithContext(ctx context.Context, request *DescribePublicConfigsRequest) (response *DescribePublicConfigsResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribePublicConfigs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleasedConfigRequest() (request *DescribeReleasedConfigRequest) {
    request = &DescribeReleasedConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasedConfig")
    
    
    return
}

func NewDescribeReleasedConfigResponse() (response *DescribeReleasedConfigResponse) {
    response = &DescribeReleasedConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReleasedConfig
// 查询group发布的配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeReleasedConfig(request *DescribeReleasedConfigRequest) (response *DescribeReleasedConfigResponse, err error) {
    return c.DescribeReleasedConfigWithContext(context.Background(), request)
}

// DescribeReleasedConfig
// 查询group发布的配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeReleasedConfigWithContext(ctx context.Context, request *DescribeReleasedConfigRequest) (response *DescribeReleasedConfigResponse, err error) {
    if request == nil {
        request = NewDescribeReleasedConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeReleasedConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReleasedConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleasedConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoriesRequest() (request *DescribeRepositoriesRequest) {
    request = &DescribeRepositoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRepositories")
    
    
    return
}

func NewDescribeRepositoriesResponse() (response *DescribeRepositoriesResponse) {
    response = &DescribeRepositoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRepositories
// 查询仓库列表
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeRepositories(request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
    return c.DescribeRepositoriesWithContext(context.Background(), request)
}

// DescribeRepositories
// 查询仓库列表
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeRepositoriesWithContext(ctx context.Context, request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeRepositories")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRepositories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRepositoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryRequest() (request *DescribeRepositoryRequest) {
    request = &DescribeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRepository")
    
    
    return
}

func NewDescribeRepositoryResponse() (response *DescribeRepositoryResponse) {
    response = &DescribeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRepository
// 查询仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeRepository(request *DescribeRepositoryRequest) (response *DescribeRepositoryResponse, err error) {
    return c.DescribeRepositoryWithContext(context.Background(), request)
}

// DescribeRepository
// 查询仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeRepositoryWithContext(ctx context.Context, request *DescribeRepositoryRequest) (response *DescribeRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeRepository")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTaskStatusRequest() (request *DescribeResourceTaskStatusRequest) {
    request = &DescribeResourceTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeResourceTaskStatus")
    
    
    return
}

func NewDescribeResourceTaskStatusResponse() (response *DescribeResourceTaskStatusResponse) {
    response = &DescribeResourceTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceTaskStatus
// 资源任务的执行状态描述接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CVMCAEMASTERQUERYERROR = "InternalError.CvmCaeMasterQueryError"
//  INVALIDPARAMETER_CVMCAEMASTERJSONDECODEFAIL = "InvalidParameter.CvmCaeMasterJsonDecodeFail"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERTASKNOTEXIST = "InvalidParameterValue.CvmCaeMasterTaskNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeResourceTaskStatus(request *DescribeResourceTaskStatusRequest) (response *DescribeResourceTaskStatusResponse, err error) {
    return c.DescribeResourceTaskStatusWithContext(context.Background(), request)
}

// DescribeResourceTaskStatus
// 资源任务的执行状态描述接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CVMCAEMASTERQUERYERROR = "InternalError.CvmCaeMasterQueryError"
//  INVALIDPARAMETER_CVMCAEMASTERJSONDECODEFAIL = "InvalidParameter.CvmCaeMasterJsonDecodeFail"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERTASKNOTEXIST = "InvalidParameterValue.CvmCaeMasterTaskNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeResourceTaskStatusWithContext(ctx context.Context, request *DescribeResourceTaskStatusRequest) (response *DescribeResourceTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTaskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeResourceTaskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleApplicationsRequest() (request *DescribeSimpleApplicationsRequest) {
    request = &DescribeSimpleApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleApplications")
    
    
    return
}

func NewDescribeSimpleApplicationsResponse() (response *DescribeSimpleApplicationsResponse) {
    response = &DescribeSimpleApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSimpleApplications
// 查询简单应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleApplications(request *DescribeSimpleApplicationsRequest) (response *DescribeSimpleApplicationsResponse, err error) {
    return c.DescribeSimpleApplicationsWithContext(context.Background(), request)
}

// DescribeSimpleApplications
// 查询简单应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleApplicationsWithContext(ctx context.Context, request *DescribeSimpleApplicationsRequest) (response *DescribeSimpleApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleApplicationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeSimpleApplications")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSimpleApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSimpleApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleClustersRequest() (request *DescribeSimpleClustersRequest) {
    request = &DescribeSimpleClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleClusters")
    
    
    return
}

func NewDescribeSimpleClustersResponse() (response *DescribeSimpleClustersResponse) {
    response = &DescribeSimpleClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSimpleClusters
// 查询简单集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleClusters(request *DescribeSimpleClustersRequest) (response *DescribeSimpleClustersResponse, err error) {
    return c.DescribeSimpleClustersWithContext(context.Background(), request)
}

// DescribeSimpleClusters
// 查询简单集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleClustersWithContext(ctx context.Context, request *DescribeSimpleClustersRequest) (response *DescribeSimpleClustersResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeSimpleClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSimpleClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSimpleClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleGroupsRequest() (request *DescribeSimpleGroupsRequest) {
    request = &DescribeSimpleGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleGroups")
    
    
    return
}

func NewDescribeSimpleGroupsResponse() (response *DescribeSimpleGroupsResponse) {
    response = &DescribeSimpleGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSimpleGroups
// 查询简单部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleGroups(request *DescribeSimpleGroupsRequest) (response *DescribeSimpleGroupsResponse, err error) {
    return c.DescribeSimpleGroupsWithContext(context.Background(), request)
}

// DescribeSimpleGroups
// 查询简单部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleGroupsWithContext(ctx context.Context, request *DescribeSimpleGroupsRequest) (response *DescribeSimpleGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeSimpleGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSimpleGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSimpleGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleNamespacesRequest() (request *DescribeSimpleNamespacesRequest) {
    request = &DescribeSimpleNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleNamespaces")
    
    
    return
}

func NewDescribeSimpleNamespacesResponse() (response *DescribeSimpleNamespacesResponse) {
    response = &DescribeSimpleNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSimpleNamespaces
// 查询简单命名空间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleNamespaces(request *DescribeSimpleNamespacesRequest) (response *DescribeSimpleNamespacesResponse, err error) {
    return c.DescribeSimpleNamespacesWithContext(context.Background(), request)
}

// DescribeSimpleNamespaces
// 查询简单命名空间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleNamespacesWithContext(ctx context.Context, request *DescribeSimpleNamespacesRequest) (response *DescribeSimpleNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleNamespacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeSimpleNamespaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSimpleNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSimpleNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStatisticsRequest() (request *DescribeStatisticsRequest) {
    request = &DescribeStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeStatistics")
    
    
    return
}

func NewDescribeStatisticsResponse() (response *DescribeStatisticsResponse) {
    response = &DescribeStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStatistics
// 服务统计页面：接口和服务维度
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLTSFDCFGFAILED = "FailedOperation.CallTsfDcfgFailed"
//  FAILEDOPERATION_FINDMETRICSEXCLUSIVEERROR = "FailedOperation.FindMetricsExclusiveError"
//  FAILEDOPERATION_TSFAPMCALLTSFMSFAILED = "FailedOperation.TsfApmCallTsfMsFailed"
//  FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"
//  FAILEDOPERATION_TSFAPMESCLIENTREQUESTERROR = "FailedOperation.TsfApmEsClientRequestError"
//  FAILEDOPERATION_TSFAPMINTERNALERROR = "FailedOperation.TsfApmInternalError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_TSFAPMCALLTSFMSFAILED = "InternalError.TsfApmCallTsfMsFailed"
//  INTERNALERROR_TSFAPMINTERNALERROR = "InternalError.TsfApmInternalError"
//  INVALIDPARAMETER_TSFAPMTRACESEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmTraceSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeStatistics(request *DescribeStatisticsRequest) (response *DescribeStatisticsResponse, err error) {
    return c.DescribeStatisticsWithContext(context.Background(), request)
}

// DescribeStatistics
// 服务统计页面：接口和服务维度
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLTSFDCFGFAILED = "FailedOperation.CallTsfDcfgFailed"
//  FAILEDOPERATION_FINDMETRICSEXCLUSIVEERROR = "FailedOperation.FindMetricsExclusiveError"
//  FAILEDOPERATION_TSFAPMCALLTSFMSFAILED = "FailedOperation.TsfApmCallTsfMsFailed"
//  FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"
//  FAILEDOPERATION_TSFAPMESCLIENTREQUESTERROR = "FailedOperation.TsfApmEsClientRequestError"
//  FAILEDOPERATION_TSFAPMINTERNALERROR = "FailedOperation.TsfApmInternalError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_TSFAPMCALLTSFMSFAILED = "InternalError.TsfApmCallTsfMsFailed"
//  INTERNALERROR_TSFAPMINTERNALERROR = "InternalError.TsfApmInternalError"
//  INVALIDPARAMETER_TSFAPMTRACESEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmTraceSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeStatisticsWithContext(ctx context.Context, request *DescribeStatisticsRequest) (response *DescribeStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskDetail")
    
    
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskDetail
// 查询任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    return c.DescribeTaskDetailWithContext(context.Background(), request)
}

// DescribeTaskDetail
// 查询任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLastStatusRequest() (request *DescribeTaskLastStatusRequest) {
    request = &DescribeTaskLastStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskLastStatus")
    
    
    return
}

func NewDescribeTaskLastStatusResponse() (response *DescribeTaskLastStatusResponse) {
    response = &DescribeTaskLastStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskLastStatus
// 查询任务最近一次执行状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskLastStatus(request *DescribeTaskLastStatusRequest) (response *DescribeTaskLastStatusResponse, err error) {
    return c.DescribeTaskLastStatusWithContext(context.Background(), request)
}

// DescribeTaskLastStatus
// 查询任务最近一次执行状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskLastStatusWithContext(ctx context.Context, request *DescribeTaskLastStatusRequest) (response *DescribeTaskLastStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLastStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeTaskLastStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskLastStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskLastStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRecordsRequest() (request *DescribeTaskRecordsRequest) {
    request = &DescribeTaskRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskRecords")
    
    
    return
}

func NewDescribeTaskRecordsResponse() (response *DescribeTaskRecordsResponse) {
    response = &DescribeTaskRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskRecords
// 翻页查询任务列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskRecords(request *DescribeTaskRecordsRequest) (response *DescribeTaskRecordsResponse, err error) {
    return c.DescribeTaskRecordsWithContext(context.Background(), request)
}

// DescribeTaskRecords
// 翻页查询任务列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskRecordsWithContext(ctx context.Context, request *DescribeTaskRecordsRequest) (response *DescribeTaskRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeTaskRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnitApiUseDetailRequest() (request *DescribeUnitApiUseDetailRequest) {
    request = &DescribeUnitApiUseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitApiUseDetail")
    
    
    return
}

func NewDescribeUnitApiUseDetailResponse() (response *DescribeUnitApiUseDetailResponse) {
    response = &DescribeUnitApiUseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnitApiUseDetail
// 查询网关API监控明细数据（仅单元化网关），非单元化网关使用DescribeApiUseDetail
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
func (c *Client) DescribeUnitApiUseDetail(request *DescribeUnitApiUseDetailRequest) (response *DescribeUnitApiUseDetailResponse, err error) {
    return c.DescribeUnitApiUseDetailWithContext(context.Background(), request)
}

// DescribeUnitApiUseDetail
// 查询网关API监控明细数据（仅单元化网关），非单元化网关使用DescribeApiUseDetail
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
func (c *Client) DescribeUnitApiUseDetailWithContext(ctx context.Context, request *DescribeUnitApiUseDetailRequest) (response *DescribeUnitApiUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeUnitApiUseDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeUnitApiUseDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnitApiUseDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnitApiUseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnitNamespacesRequest() (request *DescribeUnitNamespacesRequest) {
    request = &DescribeUnitNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitNamespaces")
    
    
    return
}

func NewDescribeUnitNamespacesResponse() (response *DescribeUnitNamespacesResponse) {
    response = &DescribeUnitNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnitNamespaces
// 查询单元化命名空间列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitNamespaces(request *DescribeUnitNamespacesRequest) (response *DescribeUnitNamespacesResponse, err error) {
    return c.DescribeUnitNamespacesWithContext(context.Background(), request)
}

// DescribeUnitNamespaces
// 查询单元化命名空间列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitNamespacesWithContext(ctx context.Context, request *DescribeUnitNamespacesRequest) (response *DescribeUnitNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeUnitNamespacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeUnitNamespaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnitNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnitNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnitRuleRequest() (request *DescribeUnitRuleRequest) {
    request = &DescribeUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitRule")
    
    
    return
}

func NewDescribeUnitRuleResponse() (response *DescribeUnitRuleResponse) {
    response = &DescribeUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnitRule
// 查询单元化规则详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitRule(request *DescribeUnitRuleRequest) (response *DescribeUnitRuleResponse, err error) {
    return c.DescribeUnitRuleWithContext(context.Background(), request)
}

// DescribeUnitRule
// 查询单元化规则详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitRuleWithContext(ctx context.Context, request *DescribeUnitRuleRequest) (response *DescribeUnitRuleResponse, err error) {
    if request == nil {
        request = NewDescribeUnitRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeUnitRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnitRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnitRulesRequest() (request *DescribeUnitRulesRequest) {
    request = &DescribeUnitRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitRules")
    
    
    return
}

func NewDescribeUnitRulesResponse() (response *DescribeUnitRulesResponse) {
    response = &DescribeUnitRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnitRules
// 查询单元化规则列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitRules(request *DescribeUnitRulesRequest) (response *DescribeUnitRulesResponse, err error) {
    return c.DescribeUnitRulesWithContext(context.Background(), request)
}

// DescribeUnitRules
// 查询单元化规则列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitRulesWithContext(ctx context.Context, request *DescribeUnitRulesRequest) (response *DescribeUnitRulesResponse, err error) {
    if request == nil {
        request = NewDescribeUnitRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeUnitRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnitRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnitRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnitRulesV2Request() (request *DescribeUnitRulesV2Request) {
    request = &DescribeUnitRulesV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitRulesV2")
    
    
    return
}

func NewDescribeUnitRulesV2Response() (response *DescribeUnitRulesV2Response) {
    response = &DescribeUnitRulesV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnitRulesV2
// 查询单元化规则列表V2
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitRulesV2(request *DescribeUnitRulesV2Request) (response *DescribeUnitRulesV2Response, err error) {
    return c.DescribeUnitRulesV2WithContext(context.Background(), request)
}

// DescribeUnitRulesV2
// 查询单元化规则列表V2
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitRulesV2WithContext(ctx context.Context, request *DescribeUnitRulesV2Request) (response *DescribeUnitRulesV2Response, err error) {
    if request == nil {
        request = NewDescribeUnitRulesV2Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeUnitRulesV2")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnitRulesV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnitRulesV2Response()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadInfoRequest() (request *DescribeUploadInfoRequest) {
    request = &DescribeUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUploadInfo")
    
    
    return
}

func NewDescribeUploadInfoResponse() (response *DescribeUploadInfoResponse) {
    response = &DescribeUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUploadInfo
// TSF会将软件包上传到腾讯云对象存储（COS）。调用此接口获取上传信息，如目标地域，桶，包Id，存储路径，鉴权信息等，之后请使用COS API（或SDK）进行上传。
//
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINUSE_OBJECTEXIST = "ResourceInUse.ObjectExist"
//  RESOURCEINSUFFICIENT_PACKAGESPACEFULL = "ResourceInsufficient.PackageSpaceFull"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeUploadInfo(request *DescribeUploadInfoRequest) (response *DescribeUploadInfoResponse, err error) {
    return c.DescribeUploadInfoWithContext(context.Background(), request)
}

// DescribeUploadInfo
// TSF会将软件包上传到腾讯云对象存储（COS）。调用此接口获取上传信息，如目标地域，桶，包Id，存储路径，鉴权信息等，之后请使用COS API（或SDK）进行上传。
//
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINUSE_OBJECTEXIST = "ResourceInUse.ObjectExist"
//  RESOURCEINSUFFICIENT_PACKAGESPACEFULL = "ResourceInsufficient.PackageSpaceFull"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeUploadInfoWithContext(ctx context.Context, request *DescribeUploadInfoRequest) (response *DescribeUploadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUploadInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeUploadInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUploadInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsableUnitNamespacesRequest() (request *DescribeUsableUnitNamespacesRequest) {
    request = &DescribeUsableUnitNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUsableUnitNamespaces")
    
    
    return
}

func NewDescribeUsableUnitNamespacesResponse() (response *DescribeUsableUnitNamespacesResponse) {
    response = &DescribeUsableUnitNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsableUnitNamespaces
// 查询可用于被导入的命名空间列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUsableUnitNamespaces(request *DescribeUsableUnitNamespacesRequest) (response *DescribeUsableUnitNamespacesResponse, err error) {
    return c.DescribeUsableUnitNamespacesWithContext(context.Background(), request)
}

// DescribeUsableUnitNamespaces
// 查询可用于被导入的命名空间列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUsableUnitNamespacesWithContext(ctx context.Context, request *DescribeUsableUnitNamespacesRequest) (response *DescribeUsableUnitNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeUsableUnitNamespacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DescribeUsableUnitNamespaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsableUnitNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsableUnitNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableLaneRuleRequest() (request *DisableLaneRuleRequest) {
    request = &DisableLaneRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DisableLaneRule")
    
    
    return
}

func NewDisableLaneRuleResponse() (response *DisableLaneRuleResponse) {
    response = &DisableLaneRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableLaneRule
// 禁用泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEIDINVALID = "InvalidParameterValue.LaneRuleIdInvalid"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) DisableLaneRule(request *DisableLaneRuleRequest) (response *DisableLaneRuleResponse, err error) {
    return c.DisableLaneRuleWithContext(context.Background(), request)
}

// DisableLaneRule
// 禁用泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEIDINVALID = "InvalidParameterValue.LaneRuleIdInvalid"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) DisableLaneRuleWithContext(ctx context.Context, request *DisableLaneRuleRequest) (response *DisableLaneRuleResponse, err error) {
    if request == nil {
        request = NewDisableLaneRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DisableLaneRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableLaneRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableLaneRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDisableTaskRequest() (request *DisableTaskRequest) {
    request = &DisableTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DisableTask")
    
    
    return
}

func NewDisableTaskResponse() (response *DisableTaskResponse) {
    response = &DisableTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableTask
// 停用任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DisableTask(request *DisableTaskRequest) (response *DisableTaskResponse, err error) {
    return c.DisableTaskWithContext(context.Background(), request)
}

// DisableTask
// 停用任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DisableTaskWithContext(ctx context.Context, request *DisableTaskRequest) (response *DisableTaskResponse, err error) {
    if request == nil {
        request = NewDisableTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DisableTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDisableTaskFlowRequest() (request *DisableTaskFlowRequest) {
    request = &DisableTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DisableTaskFlow")
    
    
    return
}

func NewDisableTaskFlowResponse() (response *DisableTaskFlowResponse) {
    response = &DisableTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableTaskFlow
// 停用工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DisableTaskFlow(request *DisableTaskFlowRequest) (response *DisableTaskFlowResponse, err error) {
    return c.DisableTaskFlowWithContext(context.Background(), request)
}

// DisableTaskFlow
// 停用工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DisableTaskFlowWithContext(ctx context.Context, request *DisableTaskFlowRequest) (response *DisableTaskFlowResponse, err error) {
    if request == nil {
        request = NewDisableTaskFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DisableTaskFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableTaskFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDisableUnitRouteRequest() (request *DisableUnitRouteRequest) {
    request = &DisableUnitRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DisableUnitRoute")
    
    
    return
}

func NewDisableUnitRouteResponse() (response *DisableUnitRouteResponse) {
    response = &DisableUnitRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableUnitRoute
// 禁用单元化路由
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DisableUnitRoute(request *DisableUnitRouteRequest) (response *DisableUnitRouteResponse, err error) {
    return c.DisableUnitRouteWithContext(context.Background(), request)
}

// DisableUnitRoute
// 禁用单元化路由
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DisableUnitRouteWithContext(ctx context.Context, request *DisableUnitRouteRequest) (response *DisableUnitRouteResponse, err error) {
    if request == nil {
        request = NewDisableUnitRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DisableUnitRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableUnitRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableUnitRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDisableUnitRuleRequest() (request *DisableUnitRuleRequest) {
    request = &DisableUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DisableUnitRule")
    
    
    return
}

func NewDisableUnitRuleResponse() (response *DisableUnitRuleResponse) {
    response = &DisableUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableUnitRule
// 禁用单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DisableUnitRule(request *DisableUnitRuleRequest) (response *DisableUnitRuleResponse, err error) {
    return c.DisableUnitRuleWithContext(context.Background(), request)
}

// DisableUnitRule
// 禁用单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DisableUnitRuleWithContext(ctx context.Context, request *DisableUnitRuleRequest) (response *DisableUnitRuleResponse, err error) {
    if request == nil {
        request = NewDisableUnitRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DisableUnitRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableUnitRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateBusinessLogConfigRequest() (request *DisassociateBusinessLogConfigRequest) {
    request = &DisassociateBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DisassociateBusinessLogConfig")
    
    
    return
}

func NewDisassociateBusinessLogConfigResponse() (response *DisassociateBusinessLogConfigResponse) {
    response = &DisassociateBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateBusinessLogConfig
// 取消关联业务日志配置项和应用
//
// 可能返回的错误码:
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"
func (c *Client) DisassociateBusinessLogConfig(request *DisassociateBusinessLogConfigRequest) (response *DisassociateBusinessLogConfigResponse, err error) {
    return c.DisassociateBusinessLogConfigWithContext(context.Background(), request)
}

// DisassociateBusinessLogConfig
// 取消关联业务日志配置项和应用
//
// 可能返回的错误码:
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"
func (c *Client) DisassociateBusinessLogConfigWithContext(ctx context.Context, request *DisassociateBusinessLogConfigRequest) (response *DisassociateBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewDisassociateBusinessLogConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DisassociateBusinessLogConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateBusinessLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateKafkaConfigRequest() (request *DisassociateKafkaConfigRequest) {
    request = &DisassociateKafkaConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DisassociateKafkaConfig")
    
    
    return
}

func NewDisassociateKafkaConfigResponse() (response *DisassociateKafkaConfigResponse) {
    response = &DisassociateKafkaConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateKafkaConfig
// 取消关联投递信息和部署组
//
// 可能返回的错误码:
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGCLOUDPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgCloudParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DisassociateKafkaConfig(request *DisassociateKafkaConfigRequest) (response *DisassociateKafkaConfigResponse, err error) {
    return c.DisassociateKafkaConfigWithContext(context.Background(), request)
}

// DisassociateKafkaConfig
// 取消关联投递信息和部署组
//
// 可能返回的错误码:
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGCLOUDPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgCloudParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DisassociateKafkaConfigWithContext(ctx context.Context, request *DisassociateKafkaConfigRequest) (response *DisassociateKafkaConfigResponse, err error) {
    if request == nil {
        request = NewDisassociateKafkaConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DisassociateKafkaConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateKafkaConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateKafkaConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDraftApiGroupRequest() (request *DraftApiGroupRequest) {
    request = &DraftApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "DraftApiGroup")
    
    
    return
}

func NewDraftApiGroupResponse() (response *DraftApiGroupResponse) {
    response = &DraftApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DraftApiGroup
// 下线Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) DraftApiGroup(request *DraftApiGroupRequest) (response *DraftApiGroupResponse, err error) {
    return c.DraftApiGroupWithContext(context.Background(), request)
}

// DraftApiGroup
// 下线Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) DraftApiGroupWithContext(ctx context.Context, request *DraftApiGroupRequest) (response *DraftApiGroupResponse, err error) {
    if request == nil {
        request = NewDraftApiGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "DraftApiGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DraftApiGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDraftApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewEnableLaneRuleRequest() (request *EnableLaneRuleRequest) {
    request = &EnableLaneRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "EnableLaneRule")
    
    
    return
}

func NewEnableLaneRuleResponse() (response *EnableLaneRuleResponse) {
    response = &EnableLaneRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableLaneRule
// 启用泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANEINFORELEASEMESHFAILED = "FailedOperation.LaneInfoReleaseMeshFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEIDINVALID = "InvalidParameterValue.LaneRuleIdInvalid"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) EnableLaneRule(request *EnableLaneRuleRequest) (response *EnableLaneRuleResponse, err error) {
    return c.EnableLaneRuleWithContext(context.Background(), request)
}

// EnableLaneRule
// 启用泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANEINFORELEASEMESHFAILED = "FailedOperation.LaneInfoReleaseMeshFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEIDINVALID = "InvalidParameterValue.LaneRuleIdInvalid"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) EnableLaneRuleWithContext(ctx context.Context, request *EnableLaneRuleRequest) (response *EnableLaneRuleResponse, err error) {
    if request == nil {
        request = NewEnableLaneRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "EnableLaneRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableLaneRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableLaneRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTaskRequest() (request *EnableTaskRequest) {
    request = &EnableTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "EnableTask")
    
    
    return
}

func NewEnableTaskResponse() (response *EnableTaskResponse) {
    response = &EnableTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableTask
// 启用任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) EnableTask(request *EnableTaskRequest) (response *EnableTaskResponse, err error) {
    return c.EnableTaskWithContext(context.Background(), request)
}

// EnableTask
// 启用任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) EnableTaskWithContext(ctx context.Context, request *EnableTaskRequest) (response *EnableTaskResponse, err error) {
    if request == nil {
        request = NewEnableTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "EnableTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableTaskResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTaskFlowRequest() (request *EnableTaskFlowRequest) {
    request = &EnableTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "EnableTaskFlow")
    
    
    return
}

func NewEnableTaskFlowResponse() (response *EnableTaskFlowResponse) {
    response = &EnableTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableTaskFlow
// 启用工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) EnableTaskFlow(request *EnableTaskFlowRequest) (response *EnableTaskFlowResponse, err error) {
    return c.EnableTaskFlowWithContext(context.Background(), request)
}

// EnableTaskFlow
// 启用工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) EnableTaskFlowWithContext(ctx context.Context, request *EnableTaskFlowRequest) (response *EnableTaskFlowResponse, err error) {
    if request == nil {
        request = NewEnableTaskFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "EnableTaskFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableTaskFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewEnableUnitRouteRequest() (request *EnableUnitRouteRequest) {
    request = &EnableUnitRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "EnableUnitRoute")
    
    
    return
}

func NewEnableUnitRouteResponse() (response *EnableUnitRouteResponse) {
    response = &EnableUnitRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableUnitRoute
// 启用单元化路由
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) EnableUnitRoute(request *EnableUnitRouteRequest) (response *EnableUnitRouteResponse, err error) {
    return c.EnableUnitRouteWithContext(context.Background(), request)
}

// EnableUnitRoute
// 启用单元化路由
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) EnableUnitRouteWithContext(ctx context.Context, request *EnableUnitRouteRequest) (response *EnableUnitRouteResponse, err error) {
    if request == nil {
        request = NewEnableUnitRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "EnableUnitRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableUnitRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableUnitRouteResponse()
    err = c.Send(request, response)
    return
}

func NewEnableUnitRuleRequest() (request *EnableUnitRuleRequest) {
    request = &EnableUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "EnableUnitRule")
    
    
    return
}

func NewEnableUnitRuleResponse() (response *EnableUnitRuleResponse) {
    response = &EnableUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableUnitRule
// 启用单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) EnableUnitRule(request *EnableUnitRuleRequest) (response *EnableUnitRuleResponse, err error) {
    return c.EnableUnitRuleWithContext(context.Background(), request)
}

// EnableUnitRule
// 启用单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) EnableUnitRuleWithContext(ctx context.Context, request *EnableUnitRuleRequest) (response *EnableUnitRuleResponse, err error) {
    if request == nil {
        request = NewEnableUnitRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "EnableUnitRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableUnitRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteTaskRequest() (request *ExecuteTaskRequest) {
    request = &ExecuteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ExecuteTask")
    
    
    return
}

func NewExecuteTaskResponse() (response *ExecuteTaskResponse) {
    response = &ExecuteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExecuteTask
// 手动执行一次任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) ExecuteTask(request *ExecuteTaskRequest) (response *ExecuteTaskResponse, err error) {
    return c.ExecuteTaskWithContext(context.Background(), request)
}

// ExecuteTask
// 手动执行一次任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) ExecuteTaskWithContext(ctx context.Context, request *ExecuteTaskRequest) (response *ExecuteTaskResponse, err error) {
    if request == nil {
        request = NewExecuteTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ExecuteTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteTaskFlowRequest() (request *ExecuteTaskFlowRequest) {
    request = &ExecuteTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ExecuteTaskFlow")
    
    
    return
}

func NewExecuteTaskFlowResponse() (response *ExecuteTaskFlowResponse) {
    response = &ExecuteTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExecuteTaskFlow
// 执行一次工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKPUSHERROR = "FailedOperation.TaskPushError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ExecuteTaskFlow(request *ExecuteTaskFlowRequest) (response *ExecuteTaskFlowResponse, err error) {
    return c.ExecuteTaskFlowWithContext(context.Background(), request)
}

// ExecuteTaskFlow
// 执行一次工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKPUSHERROR = "FailedOperation.TaskPushError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ExecuteTaskFlowWithContext(ctx context.Context, request *ExecuteTaskFlowRequest) (response *ExecuteTaskFlowResponse, err error) {
    if request == nil {
        request = NewExecuteTaskFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ExecuteTaskFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteTaskFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewExpandGroupRequest() (request *ExpandGroupRequest) {
    request = &ExpandGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ExpandGroup")
    
    
    return
}

func NewExpandGroupResponse() (response *ExpandGroupResponse) {
    response = &ExpandGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExpandGroup
// 虚拟机部署组添加实例
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INVALIDPARAMETER_CVMCAEMASTERUNKNOWNINSTANCESTATUS = "InvalidParameter.CvmCaeMasterUnknownInstanceStatus"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_GROUPVALIDINSTANCENULL = "InvalidParameterValue.GroupValidInstanceNull"
//  MISSINGPARAMETER_GROUPEXPANDSERVERIDNULL = "MissingParameter.GroupExpandServeridNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ExpandGroup(request *ExpandGroupRequest) (response *ExpandGroupResponse, err error) {
    return c.ExpandGroupWithContext(context.Background(), request)
}

// ExpandGroup
// 虚拟机部署组添加实例
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INVALIDPARAMETER_CVMCAEMASTERUNKNOWNINSTANCESTATUS = "InvalidParameter.CvmCaeMasterUnknownInstanceStatus"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_GROUPVALIDINSTANCENULL = "InvalidParameterValue.GroupValidInstanceNull"
//  MISSINGPARAMETER_GROUPEXPANDSERVERIDNULL = "MissingParameter.GroupExpandServeridNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ExpandGroupWithContext(ctx context.Context, request *ExpandGroupRequest) (response *ExpandGroupResponse, err error) {
    if request == nil {
        request = NewExpandGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ExpandGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExpandGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewExpandGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationRequest() (request *ModifyApplicationRequest) {
    request = &ModifyApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyApplication")
    
    
    return
}

func NewModifyApplicationResponse() (response *ModifyApplicationResponse) {
    response = &ModifyApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplication
// 修改应用
//
// 可能返回的错误码:
//  INVALIDPARAMETER_APPLICATIONDELETEFAILED = "InvalidParameter.ApplicationDeleteFailed"
//  INVALIDPARAMETERVALUE_APPLICATIONDESCLENGTH = "InvalidParameterValue.ApplicationDescLength"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    return c.ModifyApplicationWithContext(context.Background(), request)
}

// ModifyApplication
// 修改应用
//
// 可能返回的错误码:
//  INVALIDPARAMETER_APPLICATIONDELETEFAILED = "InvalidParameter.ApplicationDeleteFailed"
//  INVALIDPARAMETERVALUE_APPLICATIONDESCLENGTH = "InvalidParameterValue.ApplicationDescLength"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ModifyApplicationWithContext(ctx context.Context, request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterRequest() (request *ModifyClusterRequest) {
    request = &ModifyClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyCluster")
    
    
    return
}

func NewModifyClusterResponse() (response *ModifyClusterResponse) {
    response = &ModifyClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCluster
// 修改集群信息
//
// 可能返回的错误码:
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
func (c *Client) ModifyCluster(request *ModifyClusterRequest) (response *ModifyClusterResponse, err error) {
    return c.ModifyClusterWithContext(context.Background(), request)
}

// ModifyCluster
// 修改集群信息
//
// 可能返回的错误码:
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
func (c *Client) ModifyClusterWithContext(ctx context.Context, request *ModifyClusterRequest) (response *ModifyClusterResponse, err error) {
    if request == nil {
        request = NewModifyClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContainerGroupRequest() (request *ModifyContainerGroupRequest) {
    request = &ModifyContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyContainerGroup")
    
    
    return
}

func NewModifyContainerGroupResponse() (response *ModifyContainerGroupResponse) {
    response = &ModifyContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyContainerGroup
// 修改容器部署组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
func (c *Client) ModifyContainerGroup(request *ModifyContainerGroupRequest) (response *ModifyContainerGroupResponse, err error) {
    return c.ModifyContainerGroupWithContext(context.Background(), request)
}

// ModifyContainerGroup
// 修改容器部署组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
func (c *Client) ModifyContainerGroupWithContext(ctx context.Context, request *ModifyContainerGroupRequest) (response *ModifyContainerGroupResponse, err error) {
    if request == nil {
        request = NewModifyContainerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyContainerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyContainerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContainerReplicasRequest() (request *ModifyContainerReplicasRequest) {
    request = &ModifyContainerReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyContainerReplicas")
    
    
    return
}

func NewModifyContainerReplicasResponse() (response *ModifyContainerReplicasResponse) {
    response = &ModifyContainerReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyContainerReplicas
// 修改容器部署组实例数
//
// 可能返回的错误码:
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEDEPLOYMENTNOTFOUND = "InternalError.ContainergroupKuberneteDeploymentNotfound"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINSTANCENUMINVALID = "InvalidParameterValue.ContainergroupInstanceNumInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ModifyContainerReplicas(request *ModifyContainerReplicasRequest) (response *ModifyContainerReplicasResponse, err error) {
    return c.ModifyContainerReplicasWithContext(context.Background(), request)
}

// ModifyContainerReplicas
// 修改容器部署组实例数
//
// 可能返回的错误码:
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEDEPLOYMENTNOTFOUND = "InternalError.ContainergroupKuberneteDeploymentNotfound"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINSTANCENUMINVALID = "InvalidParameterValue.ContainergroupInstanceNumInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ModifyContainerReplicasWithContext(ctx context.Context, request *ModifyContainerReplicasRequest) (response *ModifyContainerReplicasResponse, err error) {
    if request == nil {
        request = NewModifyContainerReplicasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyContainerReplicas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyContainerReplicas require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyContainerReplicasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupRequest() (request *ModifyGroupRequest) {
    request = &ModifyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyGroup")
    
    
    return
}

func NewModifyGroupResponse() (response *ModifyGroupResponse) {
    response = &ModifyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGroup
// 更新部署组信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    return c.ModifyGroupWithContext(context.Background(), request)
}

// ModifyGroup
// 更新部署组信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ModifyGroupWithContext(ctx context.Context, request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    if request == nil {
        request = NewModifyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLaneRequest() (request *ModifyLaneRequest) {
    request = &ModifyLaneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyLane")
    
    
    return
}

func NewModifyLaneResponse() (response *ModifyLaneResponse) {
    response = &ModifyLaneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLane
// 更新泳道信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) ModifyLane(request *ModifyLaneRequest) (response *ModifyLaneResponse, err error) {
    return c.ModifyLaneWithContext(context.Background(), request)
}

// ModifyLane
// 更新泳道信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) ModifyLaneWithContext(ctx context.Context, request *ModifyLaneRequest) (response *ModifyLaneResponse, err error) {
    if request == nil {
        request = NewModifyLaneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyLane")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLane require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLaneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLaneRuleRequest() (request *ModifyLaneRuleRequest) {
    request = &ModifyLaneRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyLaneRule")
    
    
    return
}

func NewModifyLaneRuleResponse() (response *ModifyLaneRuleResponse) {
    response = &ModifyLaneRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLaneRule
// 更新泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGOPERATORINVALID = "InvalidParameterValue.LaneRuleTagOperatorInvalid"
//  INVALIDPARAMETERVALUE_LANERULETAGOPERATORNOTEMPTY = "InvalidParameterValue.LaneRuleTagOperatorNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUENOTEMPTY = "InvalidParameterValue.LaneRuleTagValueNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) ModifyLaneRule(request *ModifyLaneRuleRequest) (response *ModifyLaneRuleResponse, err error) {
    return c.ModifyLaneRuleWithContext(context.Background(), request)
}

// ModifyLaneRule
// 更新泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGOPERATORINVALID = "InvalidParameterValue.LaneRuleTagOperatorInvalid"
//  INVALIDPARAMETERVALUE_LANERULETAGOPERATORNOTEMPTY = "InvalidParameterValue.LaneRuleTagOperatorNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUENOTEMPTY = "InvalidParameterValue.LaneRuleTagValueNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) ModifyLaneRuleWithContext(ctx context.Context, request *ModifyLaneRuleRequest) (response *ModifyLaneRuleResponse, err error) {
    if request == nil {
        request = NewModifyLaneRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyLaneRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLaneRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLaneRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMicroserviceRequest() (request *ModifyMicroserviceRequest) {
    request = &ModifyMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyMicroservice")
    
    
    return
}

func NewModifyMicroserviceResponse() (response *ModifyMicroserviceResponse) {
    response = &ModifyMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMicroservice
// 修改微服务详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) ModifyMicroservice(request *ModifyMicroserviceRequest) (response *ModifyMicroserviceResponse, err error) {
    return c.ModifyMicroserviceWithContext(context.Background(), request)
}

// ModifyMicroservice
// 修改微服务详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) ModifyMicroserviceWithContext(ctx context.Context, request *ModifyMicroserviceRequest) (response *ModifyMicroserviceResponse, err error) {
    if request == nil {
        request = NewModifyMicroserviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyMicroservice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMicroservice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNamespaceRequest() (request *ModifyNamespaceRequest) {
    request = &ModifyNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyNamespace")
    
    
    return
}

func NewModifyNamespaceResponse() (response *ModifyNamespaceResponse) {
    response = &ModifyNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNamespace
// 修改命名空间
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) ModifyNamespace(request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
    return c.ModifyNamespaceWithContext(context.Background(), request)
}

// ModifyNamespace
// 修改命名空间
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) ModifyNamespaceWithContext(ctx context.Context, request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPathRewriteRequest() (request *ModifyPathRewriteRequest) {
    request = &ModifyPathRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyPathRewrite")
    
    
    return
}

func NewModifyPathRewriteResponse() (response *ModifyPathRewriteResponse) {
    response = &ModifyPathRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPathRewrite
// 修改路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) ModifyPathRewrite(request *ModifyPathRewriteRequest) (response *ModifyPathRewriteResponse, err error) {
    return c.ModifyPathRewriteWithContext(context.Background(), request)
}

// ModifyPathRewrite
// 修改路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) ModifyPathRewriteWithContext(ctx context.Context, request *ModifyPathRewriteRequest) (response *ModifyPathRewriteResponse, err error) {
    if request == nil {
        request = NewModifyPathRewriteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyPathRewrite")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPathRewrite require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPathRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProgramRequest() (request *ModifyProgramRequest) {
    request = &ModifyProgramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyProgram")
    
    
    return
}

func NewModifyProgramResponse() (response *ModifyProgramResponse) {
    response = &ModifyProgramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProgram
// 更新数据集
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DUPLICATEPROGRAMNAME = "InvalidParameterValue.DuplicateProgramName"
//  INVALIDPARAMETERVALUE_INVALIDPROGRAMNAME = "InvalidParameterValue.InvalidProgramName"
//  INVALIDPARAMETERVALUE_PROGRAMITEMNOTEXISTS = "InvalidParameterValue.ProgramItemNotExists"
func (c *Client) ModifyProgram(request *ModifyProgramRequest) (response *ModifyProgramResponse, err error) {
    return c.ModifyProgramWithContext(context.Background(), request)
}

// ModifyProgram
// 更新数据集
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DUPLICATEPROGRAMNAME = "InvalidParameterValue.DuplicateProgramName"
//  INVALIDPARAMETERVALUE_INVALIDPROGRAMNAME = "InvalidParameterValue.InvalidProgramName"
//  INVALIDPARAMETERVALUE_PROGRAMITEMNOTEXISTS = "InvalidParameterValue.ProgramItemNotExists"
func (c *Client) ModifyProgramWithContext(ctx context.Context, request *ModifyProgramRequest) (response *ModifyProgramResponse, err error) {
    if request == nil {
        request = NewModifyProgramRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyProgram")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProgram require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProgramResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskRequest() (request *ModifyTaskRequest) {
    request = &ModifyTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyTask")
    
    
    return
}

func NewModifyTaskResponse() (response *ModifyTaskResponse) {
    response = &ModifyTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTask
// 修改任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ModifyTask(request *ModifyTaskRequest) (response *ModifyTaskResponse, err error) {
    return c.ModifyTaskWithContext(context.Background(), request)
}

// ModifyTask
// 修改任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ModifyTaskWithContext(ctx context.Context, request *ModifyTaskRequest) (response *ModifyTaskResponse, err error) {
    if request == nil {
        request = NewModifyTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUploadInfoRequest() (request *ModifyUploadInfoRequest) {
    request = &ModifyUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyUploadInfo")
    
    
    return
}

func NewModifyUploadInfoResponse() (response *ModifyUploadInfoResponse) {
    response = &ModifyUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUploadInfo
// 调用该接口和COS的上传接口后，需要调用此接口更新TSF中保存的程序包状态。
//
// 调用此接口完成后，才标志上传包流程结束。
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) ModifyUploadInfo(request *ModifyUploadInfoRequest) (response *ModifyUploadInfoResponse, err error) {
    return c.ModifyUploadInfoWithContext(context.Background(), request)
}

// ModifyUploadInfo
// 调用该接口和COS的上传接口后，需要调用此接口更新TSF中保存的程序包状态。
//
// 调用此接口完成后，才标志上传包流程结束。
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) ModifyUploadInfoWithContext(ctx context.Context, request *ModifyUploadInfoRequest) (response *ModifyUploadInfoResponse, err error) {
    if request == nil {
        request = NewModifyUploadInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ModifyUploadInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUploadInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOperateApplicationTcrBindingRequest() (request *OperateApplicationTcrBindingRequest) {
    request = &OperateApplicationTcrBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "OperateApplicationTcrBinding")
    
    
    return
}

func NewOperateApplicationTcrBindingResponse() (response *OperateApplicationTcrBindingResponse) {
    response = &OperateApplicationTcrBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OperateApplicationTcrBinding
// 绑定解绑tcr仓库
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_IMAGEREPOTCRBINDERROR = "FailedOperation.ImagerepoTcrBindError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_IMAGEREPOTCRBINDERROR = "InternalError.ImagerepoTcrBindError"
//  INVALIDPARAMETER_IMAGEREPOTCRNAMESPACENOTFOUND = "InvalidParameter.ImagerepoTcrNamespaceNotFound"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
func (c *Client) OperateApplicationTcrBinding(request *OperateApplicationTcrBindingRequest) (response *OperateApplicationTcrBindingResponse, err error) {
    return c.OperateApplicationTcrBindingWithContext(context.Background(), request)
}

// OperateApplicationTcrBinding
// 绑定解绑tcr仓库
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"
//  FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"
//  FAILEDOPERATION_IMAGEREPOTCRBINDERROR = "FailedOperation.ImagerepoTcrBindError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_IMAGEREPOTCRBINDERROR = "InternalError.ImagerepoTcrBindError"
//  INVALIDPARAMETER_IMAGEREPOTCRNAMESPACENOTFOUND = "InvalidParameter.ImagerepoTcrNamespaceNotFound"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
func (c *Client) OperateApplicationTcrBindingWithContext(ctx context.Context, request *OperateApplicationTcrBindingRequest) (response *OperateApplicationTcrBindingResponse, err error) {
    if request == nil {
        request = NewOperateApplicationTcrBindingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "OperateApplicationTcrBinding")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OperateApplicationTcrBinding require credential")
    }

    request.SetContext(ctx)
    
    response = NewOperateApplicationTcrBindingResponse()
    err = c.Send(request, response)
    return
}

func NewReassociateBusinessLogConfigRequest() (request *ReassociateBusinessLogConfigRequest) {
    request = &ReassociateBusinessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ReassociateBusinessLogConfig")
    
    
    return
}

func NewReassociateBusinessLogConfigResponse() (response *ReassociateBusinessLogConfigResponse) {
    response = &ReassociateBusinessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReassociateBusinessLogConfig
// 后端服务已经删除这个接口,  API 接口下线处理
//
// 
//
// 重关联业务日志配置
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) ReassociateBusinessLogConfig(request *ReassociateBusinessLogConfigRequest) (response *ReassociateBusinessLogConfigResponse, err error) {
    return c.ReassociateBusinessLogConfigWithContext(context.Background(), request)
}

// ReassociateBusinessLogConfig
// 后端服务已经删除这个接口,  API 接口下线处理
//
// 
//
// 重关联业务日志配置
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) ReassociateBusinessLogConfigWithContext(ctx context.Context, request *ReassociateBusinessLogConfigRequest) (response *ReassociateBusinessLogConfigResponse, err error) {
    if request == nil {
        request = NewReassociateBusinessLogConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ReassociateBusinessLogConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReassociateBusinessLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewReassociateBusinessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskRequest() (request *RedoTaskRequest) {
    request = &RedoTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTask")
    
    
    return
}

func NewRedoTaskResponse() (response *RedoTaskResponse) {
    response = &RedoTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RedoTask
// 重新执行任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  FAILEDOPERATION_TASKTERMINATEFAILED = "FailedOperation.TaskTerminateFailed"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_TASKNOTSUPPORTED = "UnsupportedOperation.TaskNotSupported"
func (c *Client) RedoTask(request *RedoTaskRequest) (response *RedoTaskResponse, err error) {
    return c.RedoTaskWithContext(context.Background(), request)
}

// RedoTask
// 重新执行任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  FAILEDOPERATION_TASKTERMINATEFAILED = "FailedOperation.TaskTerminateFailed"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_TASKNOTSUPPORTED = "UnsupportedOperation.TaskNotSupported"
func (c *Client) RedoTaskWithContext(ctx context.Context, request *RedoTaskRequest) (response *RedoTaskResponse, err error) {
    if request == nil {
        request = NewRedoTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "RedoTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RedoTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRedoTaskResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskBatchRequest() (request *RedoTaskBatchRequest) {
    request = &RedoTaskBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskBatch")
    
    
    return
}

func NewRedoTaskBatchResponse() (response *RedoTaskBatchResponse) {
    response = &RedoTaskBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RedoTaskBatch
// 重新执行任务批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskBatch(request *RedoTaskBatchRequest) (response *RedoTaskBatchResponse, err error) {
    return c.RedoTaskBatchWithContext(context.Background(), request)
}

// RedoTaskBatch
// 重新执行任务批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskBatchWithContext(ctx context.Context, request *RedoTaskBatchRequest) (response *RedoTaskBatchResponse, err error) {
    if request == nil {
        request = NewRedoTaskBatchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "RedoTaskBatch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RedoTaskBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewRedoTaskBatchResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskExecuteRequest() (request *RedoTaskExecuteRequest) {
    request = &RedoTaskExecuteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskExecute")
    
    
    return
}

func NewRedoTaskExecuteResponse() (response *RedoTaskExecuteResponse) {
    response = &RedoTaskExecuteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RedoTaskExecute
// 重新执行在某个节点上执行任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskExecute(request *RedoTaskExecuteRequest) (response *RedoTaskExecuteResponse, err error) {
    return c.RedoTaskExecuteWithContext(context.Background(), request)
}

// RedoTaskExecute
// 重新执行在某个节点上执行任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskExecuteWithContext(ctx context.Context, request *RedoTaskExecuteRequest) (response *RedoTaskExecuteResponse, err error) {
    if request == nil {
        request = NewRedoTaskExecuteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "RedoTaskExecute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RedoTaskExecute require credential")
    }

    request.SetContext(ctx)
    
    response = NewRedoTaskExecuteResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskFlowBatchRequest() (request *RedoTaskFlowBatchRequest) {
    request = &RedoTaskFlowBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskFlowBatch")
    
    
    return
}

func NewRedoTaskFlowBatchResponse() (response *RedoTaskFlowBatchResponse) {
    response = &RedoTaskFlowBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RedoTaskFlowBatch
// 重新执行工作流批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskFlowBatch(request *RedoTaskFlowBatchRequest) (response *RedoTaskFlowBatchResponse, err error) {
    return c.RedoTaskFlowBatchWithContext(context.Background(), request)
}

// RedoTaskFlowBatch
// 重新执行工作流批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskFlowBatchWithContext(ctx context.Context, request *RedoTaskFlowBatchRequest) (response *RedoTaskFlowBatchResponse, err error) {
    if request == nil {
        request = NewRedoTaskFlowBatchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "RedoTaskFlowBatch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RedoTaskFlowBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewRedoTaskFlowBatchResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseApiGroupRequest() (request *ReleaseApiGroupRequest) {
    request = &ReleaseApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseApiGroup")
    
    
    return
}

func NewReleaseApiGroupResponse() (response *ReleaseApiGroupResponse) {
    response = &ReleaseApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseApiGroup
// 发布Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ReleaseApiGroup(request *ReleaseApiGroupRequest) (response *ReleaseApiGroupResponse, err error) {
    return c.ReleaseApiGroupWithContext(context.Background(), request)
}

// ReleaseApiGroup
// 发布Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ReleaseApiGroupWithContext(ctx context.Context, request *ReleaseApiGroupRequest) (response *ReleaseApiGroupResponse, err error) {
    if request == nil {
        request = NewReleaseApiGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ReleaseApiGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseApiGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseConfigRequest() (request *ReleaseConfigRequest) {
    request = &ReleaseConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseConfig")
    
    
    return
}

func NewReleaseConfigResponse() (response *ReleaseConfigResponse) {
    response = &ReleaseConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseConfig
// 发布配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CANNOTCONNCONSULSERVER = "InternalError.CanNotConnConsulServer"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGGROUPAPPLICATIONIDNOTMATCH = "InvalidParameterValue.ConfigGroupApplicationIdNotMatch"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"
//  MISSINGPARAMETER_GROUPIDREQUIRED = "MissingParameter.GroupIdRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ReleaseConfig(request *ReleaseConfigRequest) (response *ReleaseConfigResponse, err error) {
    return c.ReleaseConfigWithContext(context.Background(), request)
}

// ReleaseConfig
// 发布配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CANNOTCONNCONSULSERVER = "InternalError.CanNotConnConsulServer"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGGROUPAPPLICATIONIDNOTMATCH = "InvalidParameterValue.ConfigGroupApplicationIdNotMatch"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"
//  MISSINGPARAMETER_GROUPIDREQUIRED = "MissingParameter.GroupIdRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ReleaseConfigWithContext(ctx context.Context, request *ReleaseConfigRequest) (response *ReleaseConfigResponse, err error) {
    if request == nil {
        request = NewReleaseConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ReleaseConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseConfigResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseConfigWithDetailRespRequest() (request *ReleaseConfigWithDetailRespRequest) {
    request = &ReleaseConfigWithDetailRespRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseConfigWithDetailResp")
    
    
    return
}

func NewReleaseConfigWithDetailRespResponse() (response *ReleaseConfigWithDetailRespResponse) {
    response = &ReleaseConfigWithDetailRespResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseConfigWithDetailResp
// 发布配置,并且返回配置 ID
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CANNOTCONNCONSULSERVER = "InternalError.CanNotConnConsulServer"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGGROUPAPPLICATIONIDNOTMATCH = "InvalidParameterValue.ConfigGroupApplicationIdNotMatch"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"
//  MISSINGPARAMETER_GROUPIDREQUIRED = "MissingParameter.GroupIdRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ReleaseConfigWithDetailResp(request *ReleaseConfigWithDetailRespRequest) (response *ReleaseConfigWithDetailRespResponse, err error) {
    return c.ReleaseConfigWithDetailRespWithContext(context.Background(), request)
}

// ReleaseConfigWithDetailResp
// 发布配置,并且返回配置 ID
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CANNOTCONNCONSULSERVER = "InternalError.CanNotConnConsulServer"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGGROUPAPPLICATIONIDNOTMATCH = "InvalidParameterValue.ConfigGroupApplicationIdNotMatch"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"
//  MISSINGPARAMETER_GROUPIDREQUIRED = "MissingParameter.GroupIdRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ReleaseConfigWithDetailRespWithContext(ctx context.Context, request *ReleaseConfigWithDetailRespRequest) (response *ReleaseConfigWithDetailRespResponse, err error) {
    if request == nil {
        request = NewReleaseConfigWithDetailRespRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ReleaseConfigWithDetailResp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseConfigWithDetailResp require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseConfigWithDetailRespResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseFileConfigRequest() (request *ReleaseFileConfigRequest) {
    request = &ReleaseFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseFileConfig")
    
    
    return
}

func NewReleaseFileConfigResponse() (response *ReleaseFileConfigResponse) {
    response = &ReleaseFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseFileConfig
// 发布文件配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_FILECONFIGALREADYRELEASED = "InvalidParameterValue.FileConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_FILECONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.FileConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_FILECONFIGPATHEXISTS = "InvalidParameterValue.FileConfigPathExists"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) ReleaseFileConfig(request *ReleaseFileConfigRequest) (response *ReleaseFileConfigResponse, err error) {
    return c.ReleaseFileConfigWithContext(context.Background(), request)
}

// ReleaseFileConfig
// 发布文件配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_FILECONFIGALREADYRELEASED = "InvalidParameterValue.FileConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_FILECONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.FileConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_FILECONFIGPATHEXISTS = "InvalidParameterValue.FileConfigPathExists"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) ReleaseFileConfigWithContext(ctx context.Context, request *ReleaseFileConfigRequest) (response *ReleaseFileConfigResponse, err error) {
    if request == nil {
        request = NewReleaseFileConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ReleaseFileConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseFileConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewReleasePublicConfigRequest() (request *ReleasePublicConfigRequest) {
    request = &ReleasePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ReleasePublicConfig")
    
    
    return
}

func NewReleasePublicConfigResponse() (response *ReleasePublicConfigResponse) {
    response = &ReleasePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleasePublicConfig
// 发布公共配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"
//  INVALIDPARAMETERVALUE_NAMESPACENOTEXISTS = "InvalidParameterValue.NamespaceNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) ReleasePublicConfig(request *ReleasePublicConfigRequest) (response *ReleasePublicConfigResponse, err error) {
    return c.ReleasePublicConfigWithContext(context.Background(), request)
}

// ReleasePublicConfig
// 发布公共配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"
//  INVALIDPARAMETERVALUE_NAMESPACENOTEXISTS = "InvalidParameterValue.NamespaceNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) ReleasePublicConfigWithContext(ctx context.Context, request *ReleasePublicConfigRequest) (response *ReleasePublicConfigResponse, err error) {
    if request == nil {
        request = NewReleasePublicConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ReleasePublicConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleasePublicConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleasePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveInstancesRequest() (request *RemoveInstancesRequest) {
    request = &RemoveInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "RemoveInstances")
    
    
    return
}

func NewRemoveInstancesResponse() (response *RemoveInstancesResponse) {
    response = &RemoveInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveInstances
// 从 TSF 集群中批量移除云主机节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTEROPERATEENVILLEGALERROR = "FailedOperation.ClusterOperateEnvIllegalError"
//  FAILEDOPERATION_INSTANCEDELETEFAILED = "FailedOperation.InstanceDeleteFailed"
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RemoveInstances(request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
    return c.RemoveInstancesWithContext(context.Background(), request)
}

// RemoveInstances
// 从 TSF 集群中批量移除云主机节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTEROPERATEENVILLEGALERROR = "FailedOperation.ClusterOperateEnvIllegalError"
//  FAILEDOPERATION_INSTANCEDELETEFAILED = "FailedOperation.InstanceDeleteFailed"
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RemoveInstancesWithContext(ctx context.Context, request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
    if request == nil {
        request = NewRemoveInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "RemoveInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRevocationConfigRequest() (request *RevocationConfigRequest) {
    request = &RevocationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "RevocationConfig")
    
    
    return
}

func NewRevocationConfigResponse() (response *RevocationConfigResponse) {
    response = &RevocationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RevocationConfig
// 撤回已发布的配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_CONFIGRELEASENOTEXISTS = "InvalidParameterValue.ConfigReleaseNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_CONFIGRELEASEIDREQUIRED = "MissingParameter.ConfigReleaseIdRequired"
func (c *Client) RevocationConfig(request *RevocationConfigRequest) (response *RevocationConfigResponse, err error) {
    return c.RevocationConfigWithContext(context.Background(), request)
}

// RevocationConfig
// 撤回已发布的配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_CONFIGRELEASENOTEXISTS = "InvalidParameterValue.ConfigReleaseNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_CONFIGRELEASEIDREQUIRED = "MissingParameter.ConfigReleaseIdRequired"
func (c *Client) RevocationConfigWithContext(ctx context.Context, request *RevocationConfigRequest) (response *RevocationConfigResponse, err error) {
    if request == nil {
        request = NewRevocationConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "RevocationConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevocationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevocationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRevocationPublicConfigRequest() (request *RevocationPublicConfigRequest) {
    request = &RevocationPublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "RevocationPublicConfig")
    
    
    return
}

func NewRevocationPublicConfigResponse() (response *RevocationPublicConfigResponse) {
    response = &RevocationPublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RevocationPublicConfig
// 撤回已发布的公共配置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGRELEASENOTEXISTS = "InvalidParameterValue.ConfigReleaseNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) RevocationPublicConfig(request *RevocationPublicConfigRequest) (response *RevocationPublicConfigResponse, err error) {
    return c.RevocationPublicConfigWithContext(context.Background(), request)
}

// RevocationPublicConfig
// 撤回已发布的公共配置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGRELEASENOTEXISTS = "InvalidParameterValue.ConfigReleaseNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) RevocationPublicConfigWithContext(ctx context.Context, request *RevocationPublicConfigRequest) (response *RevocationPublicConfigResponse, err error) {
    if request == nil {
        request = NewRevocationPublicConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "RevocationPublicConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevocationPublicConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevocationPublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeFileConfigRequest() (request *RevokeFileConfigRequest) {
    request = &RevokeFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "RevokeFileConfig")
    
    
    return
}

func NewRevokeFileConfigResponse() (response *RevokeFileConfigResponse) {
    response = &RevokeFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RevokeFileConfig
// 撤回已发布的文件配置
//
// 可能返回的错误码:
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_FILECONFIGRELEASENOTEXISTS = "InvalidParameterValue.FileConfigReleaseNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) RevokeFileConfig(request *RevokeFileConfigRequest) (response *RevokeFileConfigResponse, err error) {
    return c.RevokeFileConfigWithContext(context.Background(), request)
}

// RevokeFileConfig
// 撤回已发布的文件配置
//
// 可能返回的错误码:
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_FILECONFIGRELEASENOTEXISTS = "InvalidParameterValue.FileConfigReleaseNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) RevokeFileConfigWithContext(ctx context.Context, request *RevokeFileConfigRequest) (response *RevokeFileConfigResponse, err error) {
    if request == nil {
        request = NewRevokeFileConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "RevokeFileConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevokeFileConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevokeFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackConfigRequest() (request *RollbackConfigRequest) {
    request = &RollbackConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "RollbackConfig")
    
    
    return
}

func NewRollbackConfigResponse() (response *RollbackConfigResponse) {
    response = &RollbackConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackConfig
// 回滚配置
//
// 可能返回的错误码:
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) RollbackConfig(request *RollbackConfigRequest) (response *RollbackConfigResponse, err error) {
    return c.RollbackConfigWithContext(context.Background(), request)
}

// RollbackConfig
// 回滚配置
//
// 可能返回的错误码:
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) RollbackConfigWithContext(ctx context.Context, request *RollbackConfigRequest) (response *RollbackConfigResponse, err error) {
    if request == nil {
        request = NewRollbackConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "RollbackConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackConfigResponse()
    err = c.Send(request, response)
    return
}

func NewSearchBusinessLogRequest() (request *SearchBusinessLogRequest) {
    request = &SearchBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "SearchBusinessLog")
    
    
    return
}

func NewSearchBusinessLogResponse() (response *SearchBusinessLogResponse) {
    response = &SearchBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchBusinessLog
// 业务日志搜索
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmBusiLogSearchRequestParamError"
//  INVALIDPARAMETERVALUE_TSFAPMBUSILOGSEARCHPARSERSPDATEFORMATERROR = "InvalidParameterValue.TsfApmBusiLogSearchParseRspDateFormatError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) SearchBusinessLog(request *SearchBusinessLogRequest) (response *SearchBusinessLogResponse, err error) {
    return c.SearchBusinessLogWithContext(context.Background(), request)
}

// SearchBusinessLog
// 业务日志搜索
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmBusiLogSearchRequestParamError"
//  INVALIDPARAMETERVALUE_TSFAPMBUSILOGSEARCHPARSERSPDATEFORMATERROR = "InvalidParameterValue.TsfApmBusiLogSearchParseRspDateFormatError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) SearchBusinessLogWithContext(ctx context.Context, request *SearchBusinessLogRequest) (response *SearchBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchBusinessLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "SearchBusinessLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchBusinessLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchStdoutLogRequest() (request *SearchStdoutLogRequest) {
    request = &SearchStdoutLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "SearchStdoutLog")
    
    
    return
}

func NewSearchStdoutLogResponse() (response *SearchStdoutLogResponse) {
    response = &SearchStdoutLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchStdoutLog
// 标准输出日志搜索
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMSTDOUTSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStdoutSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) SearchStdoutLog(request *SearchStdoutLogRequest) (response *SearchStdoutLogResponse, err error) {
    return c.SearchStdoutLogWithContext(context.Background(), request)
}

// SearchStdoutLog
// 标准输出日志搜索
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMSTDOUTSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStdoutSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) SearchStdoutLogWithContext(ctx context.Context, request *SearchStdoutLogRequest) (response *SearchStdoutLogResponse, err error) {
    if request == nil {
        request = NewSearchStdoutLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "SearchStdoutLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchStdoutLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchStdoutLogResponse()
    err = c.Send(request, response)
    return
}

func NewShrinkGroupRequest() (request *ShrinkGroupRequest) {
    request = &ShrinkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ShrinkGroup")
    
    
    return
}

func NewShrinkGroupResponse() (response *ShrinkGroupResponse) {
    response = &ShrinkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ShrinkGroup
// 下线部署组所有机器实例
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ShrinkGroup(request *ShrinkGroupRequest) (response *ShrinkGroupResponse, err error) {
    return c.ShrinkGroupWithContext(context.Background(), request)
}

// ShrinkGroup
// 下线部署组所有机器实例
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ShrinkGroupWithContext(ctx context.Context, request *ShrinkGroupRequest) (response *ShrinkGroupResponse, err error) {
    if request == nil {
        request = NewShrinkGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ShrinkGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ShrinkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewShrinkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewShrinkInstancesRequest() (request *ShrinkInstancesRequest) {
    request = &ShrinkInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "ShrinkInstances")
    
    
    return
}

func NewShrinkInstancesResponse() (response *ShrinkInstancesResponse) {
    response = &ShrinkInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ShrinkInstances
// 虚拟机部署组下线实例
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INTERNALERROR_CVMCAEMASTERNONALIVE = "InternalError.CvmCaeMasterNonAlive"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_GROUPVALIDINSTANCENULL = "InvalidParameterValue.GroupValidInstanceNull"
//  MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"
//  MISSINGPARAMETER_GROUPSHIRKSERVERIDNULL = "MissingParameter.GroupShirkServeridNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ShrinkInstances(request *ShrinkInstancesRequest) (response *ShrinkInstancesResponse, err error) {
    return c.ShrinkInstancesWithContext(context.Background(), request)
}

// ShrinkInstances
// 虚拟机部署组下线实例
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INTERNALERROR_CVMCAEMASTERNONALIVE = "InternalError.CvmCaeMasterNonAlive"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_GROUPVALIDINSTANCENULL = "InvalidParameterValue.GroupValidInstanceNull"
//  MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"
//  MISSINGPARAMETER_GROUPSHIRKSERVERIDNULL = "MissingParameter.GroupShirkServeridNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ShrinkInstancesWithContext(ctx context.Context, request *ShrinkInstancesRequest) (response *ShrinkInstancesResponse, err error) {
    if request == nil {
        request = NewShrinkInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "ShrinkInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ShrinkInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewShrinkInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartContainerGroupRequest() (request *StartContainerGroupRequest) {
    request = &StartContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "StartContainerGroup")
    
    
    return
}

func NewStartContainerGroupResponse() (response *StartContainerGroupResponse) {
    response = &StartContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartContainerGroup
// 启动容器部署组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPGROUPHASRUN = "FailedOperation.ContainergroupGroupHasrun"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPYAMLUSERCONTAINERNOTFOUND = "InvalidParameterValue.ContainergroupYamlUserContainerNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StartContainerGroup(request *StartContainerGroupRequest) (response *StartContainerGroupResponse, err error) {
    return c.StartContainerGroupWithContext(context.Background(), request)
}

// StartContainerGroup
// 启动容器部署组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPGROUPHASRUN = "FailedOperation.ContainergroupGroupHasrun"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPYAMLUSERCONTAINERNOTFOUND = "InvalidParameterValue.ContainergroupYamlUserContainerNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StartContainerGroupWithContext(ctx context.Context, request *StartContainerGroupRequest) (response *StartContainerGroupResponse, err error) {
    if request == nil {
        request = NewStartContainerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "StartContainerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartContainerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStartGroupRequest() (request *StartGroupRequest) {
    request = &StartGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "StartGroup")
    
    
    return
}

func NewStartGroupResponse() (response *StartGroupResponse) {
    response = &StartGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartGroup
// 启动分组
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) StartGroup(request *StartGroupRequest) (response *StartGroupResponse, err error) {
    return c.StartGroupWithContext(context.Background(), request)
}

// StartGroup
// 启动分组
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) StartGroupWithContext(ctx context.Context, request *StartGroupRequest) (response *StartGroupResponse, err error) {
    if request == nil {
        request = NewStartGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "StartGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopContainerGroupRequest() (request *StopContainerGroupRequest) {
    request = &StopContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "StopContainerGroup")
    
    
    return
}

func NewStopContainerGroupResponse() (response *StopContainerGroupResponse) {
    response = &StopContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopContainerGroup
// 停止容器部署组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPGROUPHASSTOP = "FailedOperation.ContainergroupGroupHasstop"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopContainerGroup(request *StopContainerGroupRequest) (response *StopContainerGroupResponse, err error) {
    return c.StopContainerGroupWithContext(context.Background(), request)
}

// StopContainerGroup
// 停止容器部署组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPGROUPHASSTOP = "FailedOperation.ContainergroupGroupHasstop"
//  FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopContainerGroupWithContext(ctx context.Context, request *StopContainerGroupRequest) (response *StopContainerGroupResponse, err error) {
    if request == nil {
        request = NewStopContainerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "StopContainerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopContainerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopGroupRequest() (request *StopGroupRequest) {
    request = &StopGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "StopGroup")
    
    
    return
}

func NewStopGroupResponse() (response *StopGroupResponse) {
    response = &StopGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopGroup
// 停止虚拟机部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERNONALIVE = "InternalError.CvmCaeMasterNonAlive"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERGROUPNOAGENT = "InvalidParameterValue.CvmCaeMasterGroupNoAgent"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopGroup(request *StopGroupRequest) (response *StopGroupResponse, err error) {
    return c.StopGroupWithContext(context.Background(), request)
}

// StopGroup
// 停止虚拟机部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERNONALIVE = "InternalError.CvmCaeMasterNonAlive"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERGROUPNOAGENT = "InvalidParameterValue.CvmCaeMasterGroupNoAgent"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopGroupWithContext(ctx context.Context, request *StopGroupRequest) (response *StopGroupResponse, err error) {
    if request == nil {
        request = NewStopGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "StopGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopTaskBatchRequest() (request *StopTaskBatchRequest) {
    request = &StopTaskBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "StopTaskBatch")
    
    
    return
}

func NewStopTaskBatchResponse() (response *StopTaskBatchResponse) {
    response = &StopTaskBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopTaskBatch
// 停止执行中的任务批次， 非运行中的任务不可调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopTaskBatch(request *StopTaskBatchRequest) (response *StopTaskBatchResponse, err error) {
    return c.StopTaskBatchWithContext(context.Background(), request)
}

// StopTaskBatch
// 停止执行中的任务批次， 非运行中的任务不可调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopTaskBatchWithContext(ctx context.Context, request *StopTaskBatchRequest) (response *StopTaskBatchResponse, err error) {
    if request == nil {
        request = NewStopTaskBatchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "StopTaskBatch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopTaskBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopTaskBatchResponse()
    err = c.Send(request, response)
    return
}

func NewStopTaskExecuteRequest() (request *StopTaskExecuteRequest) {
    request = &StopTaskExecuteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "StopTaskExecute")
    
    
    return
}

func NewStopTaskExecuteResponse() (response *StopTaskExecuteResponse) {
    response = &StopTaskExecuteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopTaskExecute
// 停止正在某个节点上执行的任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  FAILEDOPERATION_TASKTERMINATEFAILED = "FailedOperation.TaskTerminateFailed"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopTaskExecute(request *StopTaskExecuteRequest) (response *StopTaskExecuteResponse, err error) {
    return c.StopTaskExecuteWithContext(context.Background(), request)
}

// StopTaskExecute
// 停止正在某个节点上执行的任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  FAILEDOPERATION_TASKTERMINATEFAILED = "FailedOperation.TaskTerminateFailed"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopTaskExecuteWithContext(ctx context.Context, request *StopTaskExecuteRequest) (response *StopTaskExecuteResponse, err error) {
    if request == nil {
        request = NewStopTaskExecuteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "StopTaskExecute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopTaskExecute require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopTaskExecuteResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateTaskFlowBatchRequest() (request *TerminateTaskFlowBatchRequest) {
    request = &TerminateTaskFlowBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "TerminateTaskFlowBatch")
    
    
    return
}

func NewTerminateTaskFlowBatchResponse() (response *TerminateTaskFlowBatchResponse) {
    response = &TerminateTaskFlowBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateTaskFlowBatch
// 停止一个工作流批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) TerminateTaskFlowBatch(request *TerminateTaskFlowBatchRequest) (response *TerminateTaskFlowBatchResponse, err error) {
    return c.TerminateTaskFlowBatchWithContext(context.Background(), request)
}

// TerminateTaskFlowBatch
// 停止一个工作流批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) TerminateTaskFlowBatchWithContext(ctx context.Context, request *TerminateTaskFlowBatchRequest) (response *TerminateTaskFlowBatchResponse, err error) {
    if request == nil {
        request = NewTerminateTaskFlowBatchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "TerminateTaskFlowBatch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateTaskFlowBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateTaskFlowBatchResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindApiGroupRequest() (request *UnbindApiGroupRequest) {
    request = &UnbindApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "UnbindApiGroup")
    
    
    return
}

func NewUnbindApiGroupResponse() (response *UnbindApiGroupResponse) {
    response = &UnbindApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindApiGroup
// API分组批量与网关解绑
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
func (c *Client) UnbindApiGroup(request *UnbindApiGroupRequest) (response *UnbindApiGroupResponse, err error) {
    return c.UnbindApiGroupWithContext(context.Background(), request)
}

// UnbindApiGroup
// API分组批量与网关解绑
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
func (c *Client) UnbindApiGroupWithContext(ctx context.Context, request *UnbindApiGroupRequest) (response *UnbindApiGroupResponse, err error) {
    if request == nil {
        request = NewUnbindApiGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "UnbindApiGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindApiGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiGroupRequest() (request *UpdateApiGroupRequest) {
    request = &UpdateApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiGroup")
    
    
    return
}

func NewUpdateApiGroupResponse() (response *UpdateApiGroupResponse) {
    response = &UpdateApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateApiGroup
// 更新Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) UpdateApiGroup(request *UpdateApiGroupRequest) (response *UpdateApiGroupResponse, err error) {
    return c.UpdateApiGroupWithContext(context.Background(), request)
}

// UpdateApiGroup
// 更新Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) UpdateApiGroupWithContext(ctx context.Context, request *UpdateApiGroupRequest) (response *UpdateApiGroupResponse, err error) {
    if request == nil {
        request = NewUpdateApiGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "UpdateApiGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApiGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiRateLimitRuleRequest() (request *UpdateApiRateLimitRuleRequest) {
    request = &UpdateApiRateLimitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiRateLimitRule")
    
    
    return
}

func NewUpdateApiRateLimitRuleResponse() (response *UpdateApiRateLimitRuleResponse) {
    response = &UpdateApiRateLimitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateApiRateLimitRule
// 更新API限流规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERERROR = "InvalidParameterValue.GatewayParameterError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  RESOURCEINUSE_RATELIMITRULEEXISTERROR = "ResourceInUse.RatelimitRuleExistError"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) UpdateApiRateLimitRule(request *UpdateApiRateLimitRuleRequest) (response *UpdateApiRateLimitRuleResponse, err error) {
    return c.UpdateApiRateLimitRuleWithContext(context.Background(), request)
}

// UpdateApiRateLimitRule
// 更新API限流规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERERROR = "InvalidParameterValue.GatewayParameterError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  RESOURCEINUSE_RATELIMITRULEEXISTERROR = "ResourceInUse.RatelimitRuleExistError"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) UpdateApiRateLimitRuleWithContext(ctx context.Context, request *UpdateApiRateLimitRuleRequest) (response *UpdateApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewUpdateApiRateLimitRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "UpdateApiRateLimitRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApiRateLimitRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiRateLimitRulesRequest() (request *UpdateApiRateLimitRulesRequest) {
    request = &UpdateApiRateLimitRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiRateLimitRules")
    
    
    return
}

func NewUpdateApiRateLimitRulesResponse() (response *UpdateApiRateLimitRulesResponse) {
    response = &UpdateApiRateLimitRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateApiRateLimitRules
// 批量更新API限流规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RATELIMITCONSULERROR = "FailedOperation.RatelimitConsulError"
//  INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) UpdateApiRateLimitRules(request *UpdateApiRateLimitRulesRequest) (response *UpdateApiRateLimitRulesResponse, err error) {
    return c.UpdateApiRateLimitRulesWithContext(context.Background(), request)
}

// UpdateApiRateLimitRules
// 批量更新API限流规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RATELIMITCONSULERROR = "FailedOperation.RatelimitConsulError"
//  INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) UpdateApiRateLimitRulesWithContext(ctx context.Context, request *UpdateApiRateLimitRulesRequest) (response *UpdateApiRateLimitRulesResponse, err error) {
    if request == nil {
        request = NewUpdateApiRateLimitRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "UpdateApiRateLimitRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApiRateLimitRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApiRateLimitRulesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiTimeoutsRequest() (request *UpdateApiTimeoutsRequest) {
    request = &UpdateApiTimeoutsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiTimeouts")
    
    
    return
}

func NewUpdateApiTimeoutsResponse() (response *UpdateApiTimeoutsResponse) {
    response = &UpdateApiTimeoutsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateApiTimeouts
// 批量更新API超时
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateApiTimeouts(request *UpdateApiTimeoutsRequest) (response *UpdateApiTimeoutsResponse, err error) {
    return c.UpdateApiTimeoutsWithContext(context.Background(), request)
}

// UpdateApiTimeouts
// 批量更新API超时
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateApiTimeoutsWithContext(ctx context.Context, request *UpdateApiTimeoutsRequest) (response *UpdateApiTimeoutsResponse, err error) {
    if request == nil {
        request = NewUpdateApiTimeoutsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "UpdateApiTimeouts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApiTimeouts require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApiTimeoutsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateConfigTemplateRequest() (request *UpdateConfigTemplateRequest) {
    request = &UpdateConfigTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateConfigTemplate")
    
    
    return
}

func NewUpdateConfigTemplateResponse() (response *UpdateConfigTemplateResponse) {
    response = &UpdateConfigTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateConfigTemplate
// 更新参数模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"
//  FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"
//  MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"
func (c *Client) UpdateConfigTemplate(request *UpdateConfigTemplateRequest) (response *UpdateConfigTemplateResponse, err error) {
    return c.UpdateConfigTemplateWithContext(context.Background(), request)
}

// UpdateConfigTemplate
// 更新参数模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"
//  FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"
//  FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"
//  MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"
//  MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"
func (c *Client) UpdateConfigTemplateWithContext(ctx context.Context, request *UpdateConfigTemplateRequest) (response *UpdateConfigTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateConfigTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "UpdateConfigTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateConfigTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateConfigTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGatewayApiRequest() (request *UpdateGatewayApiRequest) {
    request = &UpdateGatewayApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayApi")
    
    
    return
}

func NewUpdateGatewayApiResponse() (response *UpdateGatewayApiResponse) {
    response = &UpdateGatewayApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGatewayApi
// 更新API
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateGatewayApi(request *UpdateGatewayApiRequest) (response *UpdateGatewayApiResponse, err error) {
    return c.UpdateGatewayApiWithContext(context.Background(), request)
}

// UpdateGatewayApi
// 更新API
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateGatewayApiWithContext(ctx context.Context, request *UpdateGatewayApiRequest) (response *UpdateGatewayApiResponse, err error) {
    if request == nil {
        request = NewUpdateGatewayApiRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "UpdateGatewayApi")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGatewayApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGatewayApiResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateHealthCheckSettingsRequest() (request *UpdateHealthCheckSettingsRequest) {
    request = &UpdateHealthCheckSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateHealthCheckSettings")
    
    
    return
}

func NewUpdateHealthCheckSettingsResponse() (response *UpdateHealthCheckSettingsResponse) {
    response = &UpdateHealthCheckSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateHealthCheckSettings
// 更新健康检查配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CVMCAEMASTERHEALTHCHECKCONFIGERROR = "FailedOperation.CvmCaeMasterHealthCheckConfigError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) UpdateHealthCheckSettings(request *UpdateHealthCheckSettingsRequest) (response *UpdateHealthCheckSettingsResponse, err error) {
    return c.UpdateHealthCheckSettingsWithContext(context.Background(), request)
}

// UpdateHealthCheckSettings
// 更新健康检查配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CVMCAEMASTERHEALTHCHECKCONFIGERROR = "FailedOperation.CvmCaeMasterHealthCheckConfigError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) UpdateHealthCheckSettingsWithContext(ctx context.Context, request *UpdateHealthCheckSettingsRequest) (response *UpdateHealthCheckSettingsResponse, err error) {
    if request == nil {
        request = NewUpdateHealthCheckSettingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "UpdateHealthCheckSettings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateHealthCheckSettings require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateHealthCheckSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRepositoryRequest() (request *UpdateRepositoryRequest) {
    request = &UpdateRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateRepository")
    
    
    return
}

func NewUpdateRepositoryResponse() (response *UpdateRepositoryResponse) {
    response = &UpdateRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRepository
// 更新仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) UpdateRepository(request *UpdateRepositoryRequest) (response *UpdateRepositoryResponse, err error) {
    return c.UpdateRepositoryWithContext(context.Background(), request)
}

// UpdateRepository
// 更新仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) UpdateRepositoryWithContext(ctx context.Context, request *UpdateRepositoryRequest) (response *UpdateRepositoryResponse, err error) {
    if request == nil {
        request = NewUpdateRepositoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "UpdateRepository")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUnitRuleRequest() (request *UpdateUnitRuleRequest) {
    request = &UpdateUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateUnitRule")
    
    
    return
}

func NewUpdateUnitRuleResponse() (response *UpdateUnitRuleResponse) {
    response = &UpdateUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUnitRule
// 更新单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateUnitRule(request *UpdateUnitRuleRequest) (response *UpdateUnitRuleResponse, err error) {
    return c.UpdateUnitRuleWithContext(context.Background(), request)
}

// UpdateUnitRule
// 更新单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateUnitRuleWithContext(ctx context.Context, request *UpdateUnitRuleRequest) (response *UpdateUnitRuleResponse, err error) {
    if request == nil {
        request = NewUpdateUnitRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tsf", APIVersion, "UpdateUnitRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUnitRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUnitRuleResponse()
    err = c.Send(request, response)
    return
}
