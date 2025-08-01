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

package v20170312

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccountQuota struct {
	// 后付费配额列表
	PostPaidQuotaSet []*PostPaidQuota `json:"PostPaidQuotaSet,omitnil,omitempty" name:"PostPaidQuotaSet"`

	// 预付费配额列表
	PrePaidQuotaSet []*PrePaidQuota `json:"PrePaidQuotaSet,omitnil,omitempty" name:"PrePaidQuotaSet"`

	// spot配额列表
	SpotPaidQuotaSet []*SpotPaidQuota `json:"SpotPaidQuotaSet,omitnil,omitempty" name:"SpotPaidQuotaSet"`

	// 镜像配额列表
	ImageQuotaSet []*ImageQuota `json:"ImageQuotaSet,omitnil,omitempty" name:"ImageQuotaSet"`

	// 置放群组配额列表
	DisasterRecoverGroupQuotaSet []*DisasterRecoverGroupQuota `json:"DisasterRecoverGroupQuotaSet,omitnil,omitempty" name:"DisasterRecoverGroupQuotaSet"`
}

type AccountQuotaOverview struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 配额数据
	AccountQuota *AccountQuota `json:"AccountQuota,omitnil,omitempty" name:"AccountQuota"`
}

type ActionTimer struct {
	// 定时器动作，目前仅支持销毁一个值：TerminateInstances。
	TimerAction *string `json:"TimerAction,omitnil,omitempty" name:"TimerAction"`

	// 执行时间，按照ISO8601标准表示，并且使用UTC时间。格式为 YYYY-MM-DDThh:mm:ssZ。例如 2018-05-29T11:26:40Z，执行时间必须大于当前时间5分钟。
	ActionTime *string `json:"ActionTime,omitnil,omitempty" name:"ActionTime"`

	// 扩展数据。仅做出参使用。
	Externals *Externals `json:"Externals,omitnil,omitempty" name:"Externals"`

	// 定时器ID。仅做出参使用。
	ActionTimerId *string `json:"ActionTimerId,omitnil,omitempty" name:"ActionTimerId"`

	// 定时器状态，仅做出参使用。取值范围：<br><li>UNDO：未执行</li> <li>DOING：正在执行</li><li>DONE：执行完成。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 定时器对应的实例ID。仅做出参使用。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type AllocateHostsRequestParams struct {
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 用于保证请求幂等性的字符串。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitnil,omitempty" name:"HostChargePrepaid"`

	// 实例计费类型。目前仅支持：PREPAID（预付费，即包年包月模式），默认为：'PREPAID'。
	HostChargeType *string `json:"HostChargeType,omitnil,omitempty" name:"HostChargeType"`

	// CDH实例机型，默认为：'HS1'。
	HostType *string `json:"HostType,omitnil,omitempty" name:"HostType"`

	// 购买CDH实例数量，默认为：1。
	HostCount *uint64 `json:"HostCount,omitnil,omitempty" name:"HostCount"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

type AllocateHostsRequest struct {
	*tchttp.BaseRequest
	
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 用于保证请求幂等性的字符串。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitnil,omitempty" name:"HostChargePrepaid"`

	// 实例计费类型。目前仅支持：PREPAID（预付费，即包年包月模式），默认为：'PREPAID'。
	HostChargeType *string `json:"HostChargeType,omitnil,omitempty" name:"HostChargeType"`

	// CDH实例机型，默认为：'HS1'。
	HostType *string `json:"HostType,omitnil,omitempty" name:"HostType"`

	// 购买CDH实例数量，默认为：1。
	HostCount *uint64 `json:"HostCount,omitnil,omitempty" name:"HostCount"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

func (r *AllocateHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "ClientToken")
	delete(f, "HostChargePrepaid")
	delete(f, "HostChargeType")
	delete(f, "HostType")
	delete(f, "HostCount")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateHostsResponseParams struct {
	// 新创建云子机的实例ID列表。
	HostIdSet []*string `json:"HostIdSet,omitnil,omitempty" name:"HostIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AllocateHostsResponse struct {
	*tchttp.BaseResponse
	Response *AllocateHostsResponseParams `json:"Response"`
}

func (r *AllocateHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateInstancesKeyPairsRequestParams struct {
	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。可以通过以下方式获取可用的实例ID：
	// <li>通过登录[控制台](https://console.cloud.tencent.com/cvm/index)查询实例ID。</li>
	// <li>通过调用接口 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) ，取返回信息中的`InstanceId`获取实例ID。</li>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 一个或多个待操作的密钥对ID，每次请求批量密钥对的上限为100。可以通过以下方式获取可用的密钥ID：
	// <li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥ID。</li>
	// <li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) ，取返回信息中的`KeyId`获取密钥对ID。</li>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 是否强制关机，默认值为 false。常规场景下，建议手动关机后绑定密钥。取值范围：
	// <li>true：先执行强制关机，再绑定密钥。</li>
	// <li>false：不执行强制关机，仅支持对已关机状态实例进行绑定操作。</li>
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`
}

type AssociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。可以通过以下方式获取可用的实例ID：
	// <li>通过登录[控制台](https://console.cloud.tencent.com/cvm/index)查询实例ID。</li>
	// <li>通过调用接口 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) ，取返回信息中的`InstanceId`获取实例ID。</li>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 一个或多个待操作的密钥对ID，每次请求批量密钥对的上限为100。可以通过以下方式获取可用的密钥ID：
	// <li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥ID。</li>
	// <li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) ，取返回信息中的`KeyId`获取密钥对ID。</li>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 是否强制关机，默认值为 false。常规场景下，建议手动关机后绑定密钥。取值范围：
	// <li>true：先执行强制关机，再绑定密钥。</li>
	// <li>false：不执行强制关机，仅支持对已关机状态实例进行绑定操作。</li>
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`
}

func (r *AssociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateInstancesKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "KeyIds")
	delete(f, "ForceStop")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateInstancesKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateInstancesKeyPairsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateInstancesKeyPairsResponseParams `json:"Response"`
}

func (r *AssociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateInstancesKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 要绑定的`安全组ID`，类似sg-efil73jd，只支持绑定单个安全组。可通过 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/215/15808) 接口返回值中的`SecurityGroupId`获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 被绑定的`实例ID`，类似ins-lesecurk，支持指定多个实例，每次请求批量实例的上限为100。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要绑定的`安全组ID`，类似sg-efil73jd，只支持绑定单个安全组。可通过 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/215/15808) 接口返回值中的`SecurityGroupId`获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 被绑定的`实例ID`，类似ins-lesecurk，支持指定多个实例，每次请求批量实例的上限为100。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Attribute struct {
	// 实例的自定义数据。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

type ChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。取值范围：<li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li><br>默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type ChcDeployExtraConfig struct {
	// minos类型。是使用腾讯云的minios，还是客户自己的minios。
	MiniOsType *string `json:"MiniOsType,omitnil,omitempty" name:"MiniOsType"`

	// 服务器的架构和启动方式。取值为x86_legacy, arm_uefi。
	BootType *string `json:"BootType,omitnil,omitempty" name:"BootType"`

	// PXE使用的引导文件。默认为pxelinux.0。
	BootFile *string `json:"BootFile,omitnil,omitempty" name:"BootFile"`

	// tftp服务器的ip地址。
	NextServerAddress *string `json:"NextServerAddress,omitnil,omitempty" name:"NextServerAddress"`
}

type ChcHost struct {
	// CHC物理服务器ID。
	ChcId *string `json:"ChcId,omitnil,omitempty" name:"ChcId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 服务器序列号。
	SerialNumber *string `json:"SerialNumber,omitnil,omitempty" name:"SerialNumber"`

	// CHC的状态<br/>
	// <ul>
	// <li>INIT: 设备已录入。还未配置带外和部署网络</li>
	// <li>READY: 已配置带外和部署网络</li>
	// <li>PREPARED: 可分配云主机</li>
	// <li>ONLINE: 已分配云主机</li>
	// <li>OPERATING: 设备操作中，如正在配置带外网络等。</li>
	// <li>CLEAR_NETWORK_FAILED: 清理带外和部署网络失败</li>
	// </ul>
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 设备类型。
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 所属可用区
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 带外网络。
	BmcVirtualPrivateCloud *VirtualPrivateCloud `json:"BmcVirtualPrivateCloud,omitnil,omitempty" name:"BmcVirtualPrivateCloud"`

	// 带外网络Ip。
	BmcIp *string `json:"BmcIp,omitnil,omitempty" name:"BmcIp"`

	// 带外网络安全组Id。
	BmcSecurityGroupIds []*string `json:"BmcSecurityGroupIds,omitnil,omitempty" name:"BmcSecurityGroupIds"`

	// 部署网络。
	DeployVirtualPrivateCloud *VirtualPrivateCloud `json:"DeployVirtualPrivateCloud,omitnil,omitempty" name:"DeployVirtualPrivateCloud"`

	// 部署网络Ip。
	DeployIp *string `json:"DeployIp,omitnil,omitempty" name:"DeployIp"`

	// 部署网络安全组Id。
	DeploySecurityGroupIds []*string `json:"DeploySecurityGroupIds,omitnil,omitempty" name:"DeploySecurityGroupIds"`

	// 关联的云主机Id。
	CvmInstanceId *string `json:"CvmInstanceId,omitnil,omitempty" name:"CvmInstanceId"`

	// 服务器导入的时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 机型的硬件描述，分别为CPU核数，内存容量和磁盘容量
	HardwareDescription *string `json:"HardwareDescription,omitnil,omitempty" name:"HardwareDescription"`

	// CHC物理服务器的CPU核数
	CPU *int64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// CHC物理服务器的内存大小，单位为GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// CHC物理服务器的磁盘信息
	Disk *string `json:"Disk,omitnil,omitempty" name:"Disk"`

	// 带外网络下分配的MAC地址
	BmcMAC *string `json:"BmcMAC,omitnil,omitempty" name:"BmcMAC"`

	// 部署网络下分配的MAC地址
	DeployMAC *string `json:"DeployMAC,omitnil,omitempty" name:"DeployMAC"`

	// 设备托管类型。
	// HOSTING: 托管
	// TENANT: 租赁
	TenantType *string `json:"TenantType,omitnil,omitempty" name:"TenantType"`

	// chc dhcp选项，用于minios调试。
	DeployExtraConfig *ChcDeployExtraConfig `json:"DeployExtraConfig,omitnil,omitempty" name:"DeployExtraConfig"`

	// GPU型号。
	Gpu *string `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// 网卡型号。主要指RDMA网卡。
	NetworkCard *string `json:"NetworkCard,omitnil,omitempty" name:"NetworkCard"`

	// 是否是预定义机型。
	IsPredefinedType *bool `json:"IsPredefinedType,omitnil,omitempty" name:"IsPredefinedType"`

	// CHC云主机机型。
	ChcInstanceType *string `json:"ChcInstanceType,omitnil,omitempty" name:"ChcInstanceType"`

	// CHC云主机机型簇。
	ChcInstanceFamily *string `json:"ChcInstanceFamily,omitnil,omitempty" name:"ChcInstanceFamily"`

	// CHC云主机机型簇名称。
	ChcInstanceFamilyName *string `json:"ChcInstanceFamilyName,omitnil,omitempty" name:"ChcInstanceFamilyName"`

	// 转售客户的AppId。
	ResaleAppId *string `json:"ResaleAppId,omitnil,omitempty" name:"ResaleAppId"`

	// 转售客户的账号ID。
	ResaleAccountId *string `json:"ResaleAccountId,omitnil,omitempty" name:"ResaleAccountId"`

	// 售卖状态。<br/>
	// <ul>
	// <li>NOT_FOR_SALE:不可售卖</li>
	// <li>AVAILABLE: 可售卖</li>
	// <li>SOLD: 已售卖</li>
	// </ul>
	SaleStatus *string `json:"SaleStatus,omitnil,omitempty" name:"SaleStatus"`

	// CHC物理服务器关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 最近操作
	LatestOperation *string `json:"LatestOperation,omitnil,omitempty" name:"LatestOperation"`

	// 最近操作错误码
	LatestOperationErrorCode *string `json:"LatestOperationErrorCode,omitnil,omitempty" name:"LatestOperationErrorCode"`

	// 最近操作错误详情和建议项
	LatestOperationErrorMsg *string `json:"LatestOperationErrorMsg,omitnil,omitempty" name:"LatestOperationErrorMsg"`

	// 最近操作名称
	LatestOperationName *string `json:"LatestOperationName,omitnil,omitempty" name:"LatestOperationName"`

	// 最近操作状态
	LatestOperationState *string `json:"LatestOperationState,omitnil,omitempty" name:"LatestOperationState"`
}

type ChcHostDeniedActions struct {
	// CHC物理服务器的实例id
	ChcId *string `json:"ChcId,omitnil,omitempty" name:"ChcId"`

	// CHC物理服务器的状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 当前CHC物理服务器禁止做的操作
	DenyActions []*string `json:"DenyActions,omitnil,omitempty" name:"DenyActions"`
}

// Predefined struct for user
type ConfigureChcAssistVpcRequestParams struct {
	// CHC物理服务器的实例Id。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`

	// 带外网络信息。
	BmcVirtualPrivateCloud *VirtualPrivateCloud `json:"BmcVirtualPrivateCloud,omitnil,omitempty" name:"BmcVirtualPrivateCloud"`

	// 带外网络的安全组列表
	BmcSecurityGroupIds []*string `json:"BmcSecurityGroupIds,omitnil,omitempty" name:"BmcSecurityGroupIds"`

	// 部署网络信息。
	DeployVirtualPrivateCloud *VirtualPrivateCloud `json:"DeployVirtualPrivateCloud,omitnil,omitempty" name:"DeployVirtualPrivateCloud"`

	// 部署网络的安全组列表
	DeploySecurityGroupIds []*string `json:"DeploySecurityGroupIds,omitnil,omitempty" name:"DeploySecurityGroupIds"`

	// 部署网络的附加参数，用于指定minios类型、bios引导模式等
	ChcDeployExtraConfig *ChcDeployExtraConfig `json:"ChcDeployExtraConfig,omitnil,omitempty" name:"ChcDeployExtraConfig"`
}

type ConfigureChcAssistVpcRequest struct {
	*tchttp.BaseRequest
	
	// CHC物理服务器的实例Id。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`

	// 带外网络信息。
	BmcVirtualPrivateCloud *VirtualPrivateCloud `json:"BmcVirtualPrivateCloud,omitnil,omitempty" name:"BmcVirtualPrivateCloud"`

	// 带外网络的安全组列表
	BmcSecurityGroupIds []*string `json:"BmcSecurityGroupIds,omitnil,omitempty" name:"BmcSecurityGroupIds"`

	// 部署网络信息。
	DeployVirtualPrivateCloud *VirtualPrivateCloud `json:"DeployVirtualPrivateCloud,omitnil,omitempty" name:"DeployVirtualPrivateCloud"`

	// 部署网络的安全组列表
	DeploySecurityGroupIds []*string `json:"DeploySecurityGroupIds,omitnil,omitempty" name:"DeploySecurityGroupIds"`

	// 部署网络的附加参数，用于指定minios类型、bios引导模式等
	ChcDeployExtraConfig *ChcDeployExtraConfig `json:"ChcDeployExtraConfig,omitnil,omitempty" name:"ChcDeployExtraConfig"`
}

func (r *ConfigureChcAssistVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureChcAssistVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	delete(f, "BmcVirtualPrivateCloud")
	delete(f, "BmcSecurityGroupIds")
	delete(f, "DeployVirtualPrivateCloud")
	delete(f, "DeploySecurityGroupIds")
	delete(f, "ChcDeployExtraConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfigureChcAssistVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureChcAssistVpcResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfigureChcAssistVpcResponse struct {
	*tchttp.BaseResponse
	Response *ConfigureChcAssistVpcResponseParams `json:"Response"`
}

func (r *ConfigureChcAssistVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureChcAssistVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureChcDeployVpcRequestParams struct {
	// CHC物理服务器的实例Id。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`

	// 部署网络信息。
	DeployVirtualPrivateCloud *VirtualPrivateCloud `json:"DeployVirtualPrivateCloud,omitnil,omitempty" name:"DeployVirtualPrivateCloud"`

	// 部署网络的安全组列表。
	DeploySecurityGroupIds []*string `json:"DeploySecurityGroupIds,omitnil,omitempty" name:"DeploySecurityGroupIds"`

	// 部署所需要的dhcp选项参数
	ChcDeployExtraConfig *ChcDeployExtraConfig `json:"ChcDeployExtraConfig,omitnil,omitempty" name:"ChcDeployExtraConfig"`
}

type ConfigureChcDeployVpcRequest struct {
	*tchttp.BaseRequest
	
	// CHC物理服务器的实例Id。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`

	// 部署网络信息。
	DeployVirtualPrivateCloud *VirtualPrivateCloud `json:"DeployVirtualPrivateCloud,omitnil,omitempty" name:"DeployVirtualPrivateCloud"`

	// 部署网络的安全组列表。
	DeploySecurityGroupIds []*string `json:"DeploySecurityGroupIds,omitnil,omitempty" name:"DeploySecurityGroupIds"`

	// 部署所需要的dhcp选项参数
	ChcDeployExtraConfig *ChcDeployExtraConfig `json:"ChcDeployExtraConfig,omitnil,omitempty" name:"ChcDeployExtraConfig"`
}

func (r *ConfigureChcDeployVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureChcDeployVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	delete(f, "DeployVirtualPrivateCloud")
	delete(f, "DeploySecurityGroupIds")
	delete(f, "ChcDeployExtraConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfigureChcDeployVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureChcDeployVpcResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConfigureChcDeployVpcResponse struct {
	*tchttp.BaseResponse
	Response *ConfigureChcDeployVpcResponseParams `json:"Response"`
}

func (r *ConfigureChcDeployVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureChcDeployVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConvertOperatingSystemsRequestParams struct {
	// 执行操作系统转换的实例 ID。
	// 可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	// 仅支持操作系统为 CentOS 7、CentOS 8 的实例执行转换。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否最小规模转换。
	// 默认值：false
	MinimalConversion *bool `json:"MinimalConversion,omitnil,omitempty" name:"MinimalConversion"`

	// 是否只预检。
	// 默认值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 转换的目标操作系统类型。仅支持 TencentOS。
	// 默认值：TencentOS
	TargetOSType *string `json:"TargetOSType,omitnil,omitempty" name:"TargetOSType"`
}

type ConvertOperatingSystemsRequest struct {
	*tchttp.BaseRequest
	
	// 执行操作系统转换的实例 ID。
	// 可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	// 仅支持操作系统为 CentOS 7、CentOS 8 的实例执行转换。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否最小规模转换。
	// 默认值：false
	MinimalConversion *bool `json:"MinimalConversion,omitnil,omitempty" name:"MinimalConversion"`

	// 是否只预检。
	// 默认值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 转换的目标操作系统类型。仅支持 TencentOS。
	// 默认值：TencentOS
	TargetOSType *string `json:"TargetOSType,omitnil,omitempty" name:"TargetOSType"`
}

func (r *ConvertOperatingSystemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConvertOperatingSystemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "MinimalConversion")
	delete(f, "DryRun")
	delete(f, "TargetOSType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConvertOperatingSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConvertOperatingSystemsResponseParams struct {
	// 转换的目标操系统信息，仅在入参 DryRun 为 true 时返回。
	SupportTargetOSList []*TargetOS `json:"SupportTargetOSList,omitnil,omitempty" name:"SupportTargetOSList"`

	// 操作系统转换的任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConvertOperatingSystemsResponse struct {
	*tchttp.BaseResponse
	Response *ConvertOperatingSystemsResponseParams `json:"Response"`
}

func (r *ConvertOperatingSystemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConvertOperatingSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CpuTopology struct {
	// 决定启用的CPU物理核心数。
	CoreCount *int64 `json:"CoreCount,omitnil,omitempty" name:"CoreCount"`

	// 每核心线程数。该参数决定是否开启或关闭超线程。<br><li>1 表示关闭超线程 </li><br><li>2 表示开启超线程</li>
	//  不设置时，实例使用默认的超线程策略。开关超线程请参考文档：[开启与关闭超线程](https://cloud.tencent.com/document/product/213/103798)。
	ThreadPerCore *int64 `json:"ThreadPerCore,omitnil,omitempty" name:"ThreadPerCore"`
}

// Predefined struct for user
type CreateDisasterRecoverGroupRequestParams struct {
	// 分散置放群组名称，长度1-60个字符，支持中、英文。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分散置放群组类型，取值范围：<br><li>HOST：物理机</li><li>SW：交换机</li><li>RACK：机架</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 置放群组的亲和度，在置放群组的实例会按该亲和度分布，亲和度的取值范围是：1-10，默认为1
	Affinity *int64 `json:"Affinity,omitnil,omitempty" name:"Affinity"`

	// 标签描述列表。通过指定该参数可以绑定标签到置放群组。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

type CreateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分散置放群组名称，长度1-60个字符，支持中、英文。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分散置放群组类型，取值范围：<br><li>HOST：物理机</li><li>SW：交换机</li><li>RACK：机架</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 置放群组的亲和度，在置放群组的实例会按该亲和度分布，亲和度的取值范围是：1-10，默认为1
	Affinity *int64 `json:"Affinity,omitnil,omitempty" name:"Affinity"`

	// 标签描述列表。通过指定该参数可以绑定标签到置放群组。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

func (r *CreateDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisasterRecoverGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "ClientToken")
	delete(f, "Affinity")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisasterRecoverGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisasterRecoverGroupResponseParams struct {
	// 分散置放群组ID列表。
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitnil,omitempty" name:"DisasterRecoverGroupId"`

	// 分散置放群组类型，取值范围：<br><li>HOST：物理机</li><li>SW：交换机</li><li>RACK：机架</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分散置放群组名称，长度1-60个字符，支持中、英文。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 置放群组内可容纳的云服务器数量。
	CvmQuotaTotal *int64 `json:"CvmQuotaTotal,omitnil,omitempty" name:"CvmQuotaTotal"`

	// 置放群组内已有的云服务器数量。
	CurrentNum *int64 `json:"CurrentNum,omitnil,omitempty" name:"CurrentNum"`

	// 置放群组创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDisasterRecoverGroupResponseParams `json:"Response"`
}

func (r *CreateDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHpcClusterRequestParams struct {
	// 可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 高性能计算集群名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 高性能计算集群备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 高性能计算集群类型。
	HpcClusterType *string `json:"HpcClusterType,omitnil,omitempty" name:"HpcClusterType"`

	// 高性能计算集群对应的业务场景标识，当前只支持CDC。
	HpcClusterBusinessId *string `json:"HpcClusterBusinessId,omitnil,omitempty" name:"HpcClusterBusinessId"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的HPC高性能集群。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

type CreateHpcClusterRequest struct {
	*tchttp.BaseRequest
	
	// 可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 高性能计算集群名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 高性能计算集群备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 高性能计算集群类型。
	HpcClusterType *string `json:"HpcClusterType,omitnil,omitempty" name:"HpcClusterType"`

	// 高性能计算集群对应的业务场景标识，当前只支持CDC。
	HpcClusterBusinessId *string `json:"HpcClusterBusinessId,omitnil,omitempty" name:"HpcClusterBusinessId"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的HPC高性能集群。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

func (r *CreateHpcClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHpcClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "HpcClusterType")
	delete(f, "HpcClusterBusinessId")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHpcClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHpcClusterResponseParams struct {
	// 高性能计算集群信息。
	HpcClusterSet []*HpcClusterInfo `json:"HpcClusterSet,omitnil,omitempty" name:"HpcClusterSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateHpcClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateHpcClusterResponseParams `json:"Response"`
}

func (r *CreateHpcClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHpcClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageRequestParams struct {
	// 镜像名称。
	// 最多支持60个字符。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 需要制作镜像的实例ID。基于实例创建镜像时，为必填参数。
	// InstanceId 和 SnapshotIds 为二选一必填参数。
	// 可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 镜像描述。
	// 最多支持 256 个字符。
	ImageDescription *string `json:"ImageDescription,omitnil,omitempty" name:"ImageDescription"`

	// 是否执行强制关机以制作镜像。
	// 取值范围：<br><li>true：表示关机之后制作镜像</li><br><li>false：表示开机状态制作镜像</li><br><br>默认取值：false。<br><br>开机状态制作镜像，可能导致部分数据未备份，影响数据安全。
	ForcePoweroff *string `json:"ForcePoweroff,omitnil,omitempty" name:"ForcePoweroff"`

	// 创建Windows镜像时是否启用Sysprep。
	// 取值范围：true或false，传true表示启用Sysprep，传false表示不启用，默认取值为false。
	// 
	// 关于Sysprep的详情请参考[链接](https://cloud.tencent.com/document/product/213/43498)。
	Sysprep *string `json:"Sysprep,omitnil,omitempty" name:"Sysprep"`

	// 基于实例创建整机镜像时，指定包含在镜像里的数据盘ID。
	// DataDiskIds 只能在指定 InstanceId 实例所包含的数据盘范围内指定。
	// 可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的 `DataDisks` 获取。
	DataDiskIds []*string `json:"DataDiskIds,omitnil,omitempty" name:"DataDiskIds"`

	// 基于快照创建镜像，指定快照ID，必须包含一个系统盘快照。不可与 InstanceId 同时传入。
	// InstanceId 和 SnapshotIds 为二选一必填参数。
	// 可通过 [DescribeSnapshots](https://cloud.tencent.com/document/product/362/15647) 接口返回值中的`SnapshotId`获取。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`

	// 检测本次请求的是否成功，但不会对操作的资源产生任何影响。默认取值为false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到自定义镜像。
	// 可通过 [DescribeTags](https://cloud.tencent.com/document/api/651/35316) 接口返回值中的 `TagKey` 和 `TagValue` 获取。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 镜像族
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`
}

type CreateImageRequest struct {
	*tchttp.BaseRequest
	
	// 镜像名称。
	// 最多支持60个字符。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 需要制作镜像的实例ID。基于实例创建镜像时，为必填参数。
	// InstanceId 和 SnapshotIds 为二选一必填参数。
	// 可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 镜像描述。
	// 最多支持 256 个字符。
	ImageDescription *string `json:"ImageDescription,omitnil,omitempty" name:"ImageDescription"`

	// 是否执行强制关机以制作镜像。
	// 取值范围：<br><li>true：表示关机之后制作镜像</li><br><li>false：表示开机状态制作镜像</li><br><br>默认取值：false。<br><br>开机状态制作镜像，可能导致部分数据未备份，影响数据安全。
	ForcePoweroff *string `json:"ForcePoweroff,omitnil,omitempty" name:"ForcePoweroff"`

	// 创建Windows镜像时是否启用Sysprep。
	// 取值范围：true或false，传true表示启用Sysprep，传false表示不启用，默认取值为false。
	// 
	// 关于Sysprep的详情请参考[链接](https://cloud.tencent.com/document/product/213/43498)。
	Sysprep *string `json:"Sysprep,omitnil,omitempty" name:"Sysprep"`

	// 基于实例创建整机镜像时，指定包含在镜像里的数据盘ID。
	// DataDiskIds 只能在指定 InstanceId 实例所包含的数据盘范围内指定。
	// 可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的 `DataDisks` 获取。
	DataDiskIds []*string `json:"DataDiskIds,omitnil,omitempty" name:"DataDiskIds"`

	// 基于快照创建镜像，指定快照ID，必须包含一个系统盘快照。不可与 InstanceId 同时传入。
	// InstanceId 和 SnapshotIds 为二选一必填参数。
	// 可通过 [DescribeSnapshots](https://cloud.tencent.com/document/product/362/15647) 接口返回值中的`SnapshotId`获取。
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`

	// 检测本次请求的是否成功，但不会对操作的资源产生任何影响。默认取值为false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到自定义镜像。
	// 可通过 [DescribeTags](https://cloud.tencent.com/document/api/651/35316) 接口返回值中的 `TagKey` 和 `TagValue` 获取。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 镜像族
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`
}

func (r *CreateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageName")
	delete(f, "InstanceId")
	delete(f, "ImageDescription")
	delete(f, "ForcePoweroff")
	delete(f, "Sysprep")
	delete(f, "DataDiskIds")
	delete(f, "SnapshotIds")
	delete(f, "DryRun")
	delete(f, "TagSpecification")
	delete(f, "ImageFamily")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageResponseParams struct {
	// 镜像ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateImageResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageResponseParams `json:"Response"`
}

func (r *CreateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyPairRequestParams struct {
	// 密钥对名称，可由数字、字母和下划线组成，长度不超过25个字符。密钥对名称不能和已经存在的密钥对名称重复。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 密钥对创建后所属的项目ID，ProjectId为0表示默认项目。
	// 可以通过以下方式获取项目ID：
	// <li>通过项目列表查询项目ID。</li>
	// <li>通过调用接口 [DescribeProjects](https://cloud.tencent.com/document/api/651/78725)，取返回信息中的`projectId `获取项目ID。</li>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到密钥对。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

type CreateKeyPairRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对名称，可由数字、字母和下划线组成，长度不超过25个字符。密钥对名称不能和已经存在的密钥对名称重复。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 密钥对创建后所属的项目ID，ProjectId为0表示默认项目。
	// 可以通过以下方式获取项目ID：
	// <li>通过项目列表查询项目ID。</li>
	// <li>通过调用接口 [DescribeProjects](https://cloud.tencent.com/document/api/651/78725)，取返回信息中的`projectId `获取项目ID。</li>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到密钥对。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

func (r *CreateKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyPairRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyName")
	delete(f, "ProjectId")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyPairResponseParams struct {
	// 密钥对信息。
	KeyPair *KeyPair `json:"KeyPair,omitnil,omitempty" name:"KeyPair"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *CreateKeyPairResponseParams `json:"Response"`
}

func (r *CreateKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaunchTemplateRequestParams struct {
	// 实例启动模板名称。长度为2~128个英文或中文字符。
	LaunchTemplateName *string `json:"LaunchTemplateName,omitnil,omitempty" name:"LaunchTemplateName"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，所属宿主机（在专用宿主机上创建子机时指定）等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，传入InstanceType获取当前机型支持的镜像列表，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例启动模板版本描述。长度为2~256个英文或中文字符，默认为空字符。
	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitnil,omitempty" name:"LaunchTemplateVersionDescription"`

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID\_BY\_HOUR的实例创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则系统将根据当前地域的资源售卖情况动态指定默认机型。</li><li>对于付费模式为CDHPAID的实例创建，该参数以"CDH_"为前缀，根据CPU和内存配置生成，具体形式为：CDH_XCXG，例如对于创建CPU为1核，内存为1G大小的专用宿主机的实例，该参数应该为CDH_1C1G。</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘，详情请参考[ 云硬盘使用限制](https://cloud.tencent.com/document/product/362/5145)。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络IP，即表示每个实例的主网卡IP；同时，InstanceCount参数必须与私有网络IP的个数一致且不能大于20。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 购买实例数量。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见[CVM实例购买限制](https://cloud.tencent.com/document/product/213/2664)。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。</li><li>最多支持128个字符（包含模式串）。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例登录设置。通过该参数可以设置实例的登录方式为密钥或保持镜像的原始登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的 `SecurityGroupId` 字段来获取。若不指定该参数，则绑定指定项目下的默认安全组，如默认安全组不存在则将自动创建。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务；自定义镜像与镜像市场镜像默认不开启云监控，云安全服务，而使用镜像里保留的服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><li>其他类型（Linux 等）实例：字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer,omitnil,omitempty" name:"ActionTimer"`

	// 置放群组id，仅支持指定一个。该参数可以通过调用 [ DescribeDisasterRecoverGroups ](https://cloud.tencent.com/document/api/213/17810) 的返回值中的 `DisasterRecoverGroupId` 字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云服务器实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526)和[Linux](https://cloud.tencent.com/document/product/213/17525)启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// CAM角色名称。可通过[ DescribeRoleList ](https://cloud.tencent.com/document/product/598/36223)接口返回值中的`RoleName `获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 高性能计算集群ID。若创建的实例为高性能计算实例，需指定实例放置的集群，否则不可指定。该参数可以通过调用 [DescribeHpcClusters](https://cloud.tencent.com/document/api/213/83220) 的返回值中的 `HpcClusterId` 字段来获取。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：按小时后付费</li><li>CDHPAID：独享子机（基于专用宿主机创建，宿主机部分的资源不收费）</li><li>SPOTPAID：竞价付费</li>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>TRUE：表示开启实例保护，不允许通过api接口删除实例</li><li>FALSE：表示关闭实例保护，允许通过api接口删除实例<br></li>默认取值：FALSE。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`

	// 标签描述列表。通过指定该参数可以绑定标签到实例启动模板。
	LaunchTemplateTagSpecification []*TagSpecification `json:"LaunchTemplateTagSpecification,omitnil,omitempty" name:"LaunchTemplateTagSpecification"`

	// 自定义metadata，支持创建 CVM 时添加自定义元数据键值对。
	// **注：内测中**。
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 只允许传递 Update 和 Replace 参数，在模板使用自定义 Metadata 且在 RunInstances 也传递 Metadata 时生效。默认采用 Replace。
	// 
	// - Update：设模板 t含本参数值为Update、 metadata=[k1:v1, k2:v2] ，则RunInstances（给metadata=[k2:v3]）+ t 创建的 cvm 使用metadata=[k1:v1, k2:v3] 
	// - Replace：模板 t含本参数值为Replace、 metadata=[k1:v1, k2:v2] ，则RunInstances（给metadata=[k2:v3]）+ t 创建的 cvm 使用metadata=[k2:v3] 
	// **注：内测中**。
	TemplateDataModifyAction *string `json:"TemplateDataModifyAction,omitnil,omitempty" name:"TemplateDataModifyAction"`
}

type CreateLaunchTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 实例启动模板名称。长度为2~128个英文或中文字符。
	LaunchTemplateName *string `json:"LaunchTemplateName,omitnil,omitempty" name:"LaunchTemplateName"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，所属宿主机（在专用宿主机上创建子机时指定）等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，传入InstanceType获取当前机型支持的镜像列表，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例启动模板版本描述。长度为2~256个英文或中文字符，默认为空字符。
	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitnil,omitempty" name:"LaunchTemplateVersionDescription"`

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID\_BY\_HOUR的实例创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则系统将根据当前地域的资源售卖情况动态指定默认机型。</li><li>对于付费模式为CDHPAID的实例创建，该参数以"CDH_"为前缀，根据CPU和内存配置生成，具体形式为：CDH_XCXG，例如对于创建CPU为1核，内存为1G大小的专用宿主机的实例，该参数应该为CDH_1C1G。</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘，详情请参考[ 云硬盘使用限制](https://cloud.tencent.com/document/product/362/5145)。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络IP，即表示每个实例的主网卡IP；同时，InstanceCount参数必须与私有网络IP的个数一致且不能大于20。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 购买实例数量。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见[CVM实例购买限制](https://cloud.tencent.com/document/product/213/2664)。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。</li><li>最多支持128个字符（包含模式串）。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例登录设置。通过该参数可以设置实例的登录方式为密钥或保持镜像的原始登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的 `SecurityGroupId` 字段来获取。若不指定该参数，则绑定指定项目下的默认安全组，如默认安全组不存在则将自动创建。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务；自定义镜像与镜像市场镜像默认不开启云监控，云安全服务，而使用镜像里保留的服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><li>其他类型（Linux 等）实例：字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer,omitnil,omitempty" name:"ActionTimer"`

	// 置放群组id，仅支持指定一个。该参数可以通过调用 [ DescribeDisasterRecoverGroups ](https://cloud.tencent.com/document/api/213/17810) 的返回值中的 `DisasterRecoverGroupId` 字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云服务器实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526)和[Linux](https://cloud.tencent.com/document/product/213/17525)启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// CAM角色名称。可通过[ DescribeRoleList ](https://cloud.tencent.com/document/product/598/36223)接口返回值中的`RoleName `获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 高性能计算集群ID。若创建的实例为高性能计算实例，需指定实例放置的集群，否则不可指定。该参数可以通过调用 [DescribeHpcClusters](https://cloud.tencent.com/document/api/213/83220) 的返回值中的 `HpcClusterId` 字段来获取。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：按小时后付费</li><li>CDHPAID：独享子机（基于专用宿主机创建，宿主机部分的资源不收费）</li><li>SPOTPAID：竞价付费</li>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>TRUE：表示开启实例保护，不允许通过api接口删除实例</li><li>FALSE：表示关闭实例保护，允许通过api接口删除实例<br></li>默认取值：FALSE。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`

	// 标签描述列表。通过指定该参数可以绑定标签到实例启动模板。
	LaunchTemplateTagSpecification []*TagSpecification `json:"LaunchTemplateTagSpecification,omitnil,omitempty" name:"LaunchTemplateTagSpecification"`

	// 自定义metadata，支持创建 CVM 时添加自定义元数据键值对。
	// **注：内测中**。
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 只允许传递 Update 和 Replace 参数，在模板使用自定义 Metadata 且在 RunInstances 也传递 Metadata 时生效。默认采用 Replace。
	// 
	// - Update：设模板 t含本参数值为Update、 metadata=[k1:v1, k2:v2] ，则RunInstances（给metadata=[k2:v3]）+ t 创建的 cvm 使用metadata=[k1:v1, k2:v3] 
	// - Replace：模板 t含本参数值为Replace、 metadata=[k1:v1, k2:v2] ，则RunInstances（给metadata=[k2:v3]）+ t 创建的 cvm 使用metadata=[k2:v3] 
	// **注：内测中**。
	TemplateDataModifyAction *string `json:"TemplateDataModifyAction,omitnil,omitempty" name:"TemplateDataModifyAction"`
}

func (r *CreateLaunchTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaunchTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateName")
	delete(f, "Placement")
	delete(f, "ImageId")
	delete(f, "LaunchTemplateVersionDescription")
	delete(f, "InstanceType")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "VirtualPrivateCloud")
	delete(f, "InternetAccessible")
	delete(f, "InstanceCount")
	delete(f, "InstanceName")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "EnhancedService")
	delete(f, "ClientToken")
	delete(f, "HostName")
	delete(f, "ActionTimer")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "TagSpecification")
	delete(f, "InstanceMarketOptions")
	delete(f, "UserData")
	delete(f, "DryRun")
	delete(f, "CamRoleName")
	delete(f, "HpcClusterId")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "DisableApiTermination")
	delete(f, "LaunchTemplateTagSpecification")
	delete(f, "Metadata")
	delete(f, "TemplateDataModifyAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLaunchTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaunchTemplateResponseParams struct {
	// 当通过本接口来创建实例启动模板时会返回该参数，表示创建成功的实例启动模板`ID`。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLaunchTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateLaunchTemplateResponseParams `json:"Response"`
}

func (r *CreateLaunchTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaunchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaunchTemplateVersionRequestParams struct {
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，所属宿主机（在专用宿主机上创建子机时指定）等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 启动模板ID，新版本将基于该实例启动模板ID创建。可通过 [DescribeLaunchTemplates](https://cloud.tencent.com/document/api/213/66322) 接口返回值中的`LaunchTemplateId`获取。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 若给定，新实例启动模板将基于给定的版本号创建。若未指定则使用默认版本,可以通过 [DescribeLaunchTemplateVersions](https://cloud.tencent.com/document/api/213/66323)查询默认版本。
	LaunchTemplateVersion *int64 `json:"LaunchTemplateVersion,omitnil,omitempty" name:"LaunchTemplateVersion"`

	// 实例启动模板版本描述。长度为2~256个英文或中文字符，不指定该参数时默认为空字符。
	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitnil,omitempty" name:"LaunchTemplateVersionDescription"`

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID\_BY\_HOUR的实例创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则系统将根据当前地域的资源售卖情况动态指定默认机型。</li><br><li>对于付费模式为CDHPAID的实例创建，该参数以"CDH_"为前缀，根据CPU和内存配置生成，具体形式为：CDH_XCXG，例如对于创建CPU为1核，内存为1G大小的专用宿主机的实例，该参数应该为CDH_1C1G。</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>云镜像市场</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`云镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，传入InstanceType获取当前机型支持的镜像列表，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络IP，即表示每个实例的主网卡IP；同时，InstanceCount参数必须与私有网络IP的个数一致且不能大于20。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 购买实例数量。具体配额相关限制详见[CVM实例购买限制](https://cloud.tencent.com/document/product/213/2664)。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。</li><li>最多支持128个字符（包含模式串）。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例登录设置。通过该参数可以设置实例的登录方式为密钥或保持镜像的原始登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务；自定义镜像与云镜像市场镜像默认不开启云监控，云安全服务，而使用镜像里保留的服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><br><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><br><li>其他类型（Linux 等）实例：字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer,omitnil,omitempty" name:"ActionTimer"`

	// 置放群组id，仅支持指定一个。可使用[DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/api/213/17810)接口获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云服务器实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526)和[Linux](https://cloud.tencent.com/document/product/213/17525)启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// CAM角色名称。可通过[ DescribeRoleList ](https://cloud.tencent.com/document/product/598/13887)接口返回值中的`roleName`获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 高性能计算集群ID。若创建的实例为高性能计算实例，需指定实例放置的集群，否则不可指定。该参数可以通过调用 [DescribeHpcClusters](https://cloud.tencent.com/document/api/213/83220) 的返回值中的 `HpcClusterId` 字段来获取。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：按小时后付费</li><li>CDHPAID：独享子机（基于专用宿主机创建，宿主机部分的资源不收费）</li><li>SPOTPAID：竞价付费</li>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>TRUE：表示开启实例保护，不允许通过api接口删除实例</li><br><li>FALSE：表示关闭实例保护，允许通过api接口删除实例</li><br><br>默认取值：FALSE。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`

	// 自定义metadata，支持创建 CVM 时添加自定义元数据键值对。
	// **注：内测中**。
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 只允许传递 Update 和 Replace 参数，在模板使用自定义 Metadata 且在 RunInstances 也传递 Metadata 时生效。默认采用 Replace。
	// 
	// - Update：设模板 t含本参数值为Update、 metadata=[k1:v1, k2:v2] ，则RunInstances（给metadata=[k2:v3]）+ t 创建的 cvm 使用metadata=[k1:v1, k2:v3] 
	// - Replace：模板 t含本参数值为Replace、 metadata=[k1:v1, k2:v2] ，则RunInstances（给metadata=[k2:v3]）+ t 创建的 cvm 使用metadata=[k2:v3] 
	// **注：内测中**。
	TemplateDataModifyAction *string `json:"TemplateDataModifyAction,omitnil,omitempty" name:"TemplateDataModifyAction"`
}

type CreateLaunchTemplateVersionRequest struct {
	*tchttp.BaseRequest
	
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，所属宿主机（在专用宿主机上创建子机时指定）等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 启动模板ID，新版本将基于该实例启动模板ID创建。可通过 [DescribeLaunchTemplates](https://cloud.tencent.com/document/api/213/66322) 接口返回值中的`LaunchTemplateId`获取。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 若给定，新实例启动模板将基于给定的版本号创建。若未指定则使用默认版本,可以通过 [DescribeLaunchTemplateVersions](https://cloud.tencent.com/document/api/213/66323)查询默认版本。
	LaunchTemplateVersion *int64 `json:"LaunchTemplateVersion,omitnil,omitempty" name:"LaunchTemplateVersion"`

	// 实例启动模板版本描述。长度为2~256个英文或中文字符，不指定该参数时默认为空字符。
	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitnil,omitempty" name:"LaunchTemplateVersionDescription"`

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID\_BY\_HOUR的实例创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则系统将根据当前地域的资源售卖情况动态指定默认机型。</li><br><li>对于付费模式为CDHPAID的实例创建，该参数以"CDH_"为前缀，根据CPU和内存配置生成，具体形式为：CDH_XCXG，例如对于创建CPU为1核，内存为1G大小的专用宿主机的实例，该参数应该为CDH_1C1G。</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>云镜像市场</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`云镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，传入InstanceType获取当前机型支持的镜像列表，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络IP，即表示每个实例的主网卡IP；同时，InstanceCount参数必须与私有网络IP的个数一致且不能大于20。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 购买实例数量。具体配额相关限制详见[CVM实例购买限制](https://cloud.tencent.com/document/product/213/2664)。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。</li><li>最多支持128个字符（包含模式串）。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例登录设置。通过该参数可以设置实例的登录方式为密钥或保持镜像的原始登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务；自定义镜像与云镜像市场镜像默认不开启云监控，云安全服务，而使用镜像里保留的服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><br><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><br><li>其他类型（Linux 等）实例：字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer,omitnil,omitempty" name:"ActionTimer"`

	// 置放群组id，仅支持指定一个。可使用[DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/api/213/17810)接口获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云服务器实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526)和[Linux](https://cloud.tencent.com/document/product/213/17525)启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// CAM角色名称。可通过[ DescribeRoleList ](https://cloud.tencent.com/document/product/598/13887)接口返回值中的`roleName`获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 高性能计算集群ID。若创建的实例为高性能计算实例，需指定实例放置的集群，否则不可指定。该参数可以通过调用 [DescribeHpcClusters](https://cloud.tencent.com/document/api/213/83220) 的返回值中的 `HpcClusterId` 字段来获取。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：按小时后付费</li><li>CDHPAID：独享子机（基于专用宿主机创建，宿主机部分的资源不收费）</li><li>SPOTPAID：竞价付费</li>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>TRUE：表示开启实例保护，不允许通过api接口删除实例</li><br><li>FALSE：表示关闭实例保护，允许通过api接口删除实例</li><br><br>默认取值：FALSE。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`

	// 自定义metadata，支持创建 CVM 时添加自定义元数据键值对。
	// **注：内测中**。
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 只允许传递 Update 和 Replace 参数，在模板使用自定义 Metadata 且在 RunInstances 也传递 Metadata 时生效。默认采用 Replace。
	// 
	// - Update：设模板 t含本参数值为Update、 metadata=[k1:v1, k2:v2] ，则RunInstances（给metadata=[k2:v3]）+ t 创建的 cvm 使用metadata=[k1:v1, k2:v3] 
	// - Replace：模板 t含本参数值为Replace、 metadata=[k1:v1, k2:v2] ，则RunInstances（给metadata=[k2:v3]）+ t 创建的 cvm 使用metadata=[k2:v3] 
	// **注：内测中**。
	TemplateDataModifyAction *string `json:"TemplateDataModifyAction,omitnil,omitempty" name:"TemplateDataModifyAction"`
}

func (r *CreateLaunchTemplateVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaunchTemplateVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "LaunchTemplateId")
	delete(f, "LaunchTemplateVersion")
	delete(f, "LaunchTemplateVersionDescription")
	delete(f, "InstanceType")
	delete(f, "ImageId")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "VirtualPrivateCloud")
	delete(f, "InternetAccessible")
	delete(f, "InstanceCount")
	delete(f, "InstanceName")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "EnhancedService")
	delete(f, "ClientToken")
	delete(f, "HostName")
	delete(f, "ActionTimer")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "TagSpecification")
	delete(f, "InstanceMarketOptions")
	delete(f, "UserData")
	delete(f, "DryRun")
	delete(f, "CamRoleName")
	delete(f, "HpcClusterId")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "DisableApiTermination")
	delete(f, "Metadata")
	delete(f, "TemplateDataModifyAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLaunchTemplateVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaunchTemplateVersionResponseParams struct {
	// 新创建的实例启动模板版本号。
	LaunchTemplateVersionNumber *int64 `json:"LaunchTemplateVersionNumber,omitnil,omitempty" name:"LaunchTemplateVersionNumber"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLaunchTemplateVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreateLaunchTemplateVersionResponseParams `json:"Response"`
}

func (r *CreateLaunchTemplateVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaunchTemplateVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// 数据盘大小，单位：GiB。最小调整步长为10GiB，不同数据盘类型取值范围不同，具体限制详见：[存储概述](https://cloud.tencent.com/document/product/213/4952)。默认值为0，表示不购买数据盘。更多限制详见产品文档。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 数据盘类型。数据盘类型限制详见[存储概述](https://cloud.tencent.com/document/product/213/4952)。取值范围：<br /><li>LOCAL_BASIC：本地硬盘 </li> <li>LOCAL_SSD：本地SSD硬盘</li><li>LOCAL_NVME：本地NVME硬盘，与InstanceType强相关，不支持指定</li><li>LOCAL_PRO：本地HDD硬盘，与InstanceType强相关，不支持指定</li><li>CLOUD_BASIC：普通云硬盘</li><li> CLOUD_PREMIUM：高性能云硬盘</li><li>CLOUD_SSD：SSD云硬盘</li><li> CLOUD_HSSD：增强型SSD云硬盘</li> <li>CLOUD_TSSD：极速型SSD云硬盘</li><li>CLOUD_BSSD：通用型SSD云硬盘</li><br />默认取值：LOCAL_BASIC<br/><br />该参数对`ResizeInstanceDisk`接口无效。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 数据盘ID。
	// 该参数目前仅用于`DescribeInstances`等查询类接口的返回参数，不可用于`RunInstances`等写接口的入参。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 数据盘是否随子机销毁。取值范围：<li>true：子机销毁时，销毁数据盘，只支持按小时后付费云盘</li><li>false：子机销毁时，保留数据盘</li><br/>默认取值：true <br/>该参数目前仅用于 `RunInstances` 接口。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// 数据盘快照ID。选择的数据盘快照大小需小于数据盘大小。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 数据盘是否加密。取值范围：<li>true：加密</li><li>false：不加密</li><br/>默认取值：false<br/>该参数目前仅用于 `RunInstances` 接口。
	Encrypt *bool `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 自定义CMK对应的ID，取值为UUID或者类似kms-abcd1234。用于加密云盘。
	// 
	// 该参数目前仅用于 `RunInstances` 接口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// 云硬盘性能，单位：MiB/s。使用此参数可给云硬盘购买额外的性能。
	// 当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThroughputPerformance *int64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 所属的独享集群ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// 突发性能
	//  <b>注：内测中。</b>
	BurstPerformance *bool `json:"BurstPerformance,omitnil,omitempty" name:"BurstPerformance"`

	// 磁盘名称，长度不超过128 个字符。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`
}

// Predefined struct for user
type DeleteDisasterRecoverGroupsRequestParams struct {
	// 分散置放群组ID列表，可通过[DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/api/213/17810)接口获取。每次请求允许操作的分散置放群组数量上限是10。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`
}

type DeleteDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 分散置放群组ID列表，可通过[DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/api/213/17810)接口获取。每次请求允许操作的分散置放群组数量上限是10。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`
}

func (r *DeleteDisasterRecoverGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDisasterRecoverGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisasterRecoverGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDisasterRecoverGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDisasterRecoverGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDisasterRecoverGroupsResponseParams `json:"Response"`
}

func (r *DeleteDisasterRecoverGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDisasterRecoverGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHpcClustersRequestParams struct {
	// 高性能计算集群ID列表。
	HpcClusterIds []*string `json:"HpcClusterIds,omitnil,omitempty" name:"HpcClusterIds"`
}

type DeleteHpcClustersRequest struct {
	*tchttp.BaseRequest
	
	// 高性能计算集群ID列表。
	HpcClusterIds []*string `json:"HpcClusterIds,omitnil,omitempty" name:"HpcClusterIds"`
}

func (r *DeleteHpcClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHpcClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HpcClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHpcClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHpcClustersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteHpcClustersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHpcClustersResponseParams `json:"Response"`
}

func (r *DeleteHpcClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHpcClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImagesRequestParams struct {
	// 删除的镜像 ID 列表。
	// 可通过 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) 接口返回值中的`ImageId`获取。
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// 是否删除镜像关联的快照。
	// 默认值：false
	DeleteBindedSnap *bool `json:"DeleteBindedSnap,omitnil,omitempty" name:"DeleteBindedSnap"`

	// 检测是否支持删除镜像。
	// 默认值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type DeleteImagesRequest struct {
	*tchttp.BaseRequest
	
	// 删除的镜像 ID 列表。
	// 可通过 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) 接口返回值中的`ImageId`获取。
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// 是否删除镜像关联的快照。
	// 默认值：false
	DeleteBindedSnap *bool `json:"DeleteBindedSnap,omitnil,omitempty" name:"DeleteBindedSnap"`

	// 检测是否支持删除镜像。
	// 默认值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *DeleteImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIds")
	delete(f, "DeleteBindedSnap")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImagesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImagesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImagesResponseParams `json:"Response"`
}

func (r *DeleteImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstancesActionTimerRequestParams struct {
	// 定时任务ID列表，可以通过[ DescribeInstancesActionTimer ](https://cloud.tencent.com/document/product/213/103950)接口查询。只能删除未执行的定时任务。
	ActionTimerIds []*string `json:"ActionTimerIds,omitnil,omitempty" name:"ActionTimerIds"`
}

type DeleteInstancesActionTimerRequest struct {
	*tchttp.BaseRequest
	
	// 定时任务ID列表，可以通过[ DescribeInstancesActionTimer ](https://cloud.tencent.com/document/product/213/103950)接口查询。只能删除未执行的定时任务。
	ActionTimerIds []*string `json:"ActionTimerIds,omitnil,omitempty" name:"ActionTimerIds"`
}

func (r *DeleteInstancesActionTimerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstancesActionTimerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActionTimerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstancesActionTimerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstancesActionTimerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstancesActionTimerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstancesActionTimerResponseParams `json:"Response"`
}

func (r *DeleteInstancesActionTimerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstancesActionTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKeyPairsRequestParams struct {
	// 一个或多个待操作的密钥对ID。每次请求批量密钥对的上限为100。<br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥ID。</li><br><li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) ，取返回信息中的 `KeyId` 获取密钥对ID。</li>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type DeleteKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的密钥对ID。每次请求批量密钥对的上限为100。<br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥ID。</li><br><li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) ，取返回信息中的 `KeyId` 获取密钥对ID。</li>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *DeleteKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKeyPairsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteKeyPairsResponseParams `json:"Response"`
}

func (r *DeleteKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchTemplateRequestParams struct {
	// 启动模板ID。可通过 [DescribeLaunchTemplates](https://cloud.tencent.com/document/api/213/66322) 接口返回值中的`LaunchTemplateId`获取。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`
}

type DeleteLaunchTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 启动模板ID。可通过 [DescribeLaunchTemplates](https://cloud.tencent.com/document/api/213/66322) 接口返回值中的`LaunchTemplateId`获取。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`
}

func (r *DeleteLaunchTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLaunchTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLaunchTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLaunchTemplateResponseParams `json:"Response"`
}

func (r *DeleteLaunchTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchTemplateVersionsRequestParams struct {
	// 启动模板ID。可通过 [DescribeLaunchTemplates](https://cloud.tencent.com/document/api/213/66322) 接口返回值中的`LaunchTemplateId`获取。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 实例启动模板版本列表。可通过 [DescribeLaunchTemplateVersions](https://cloud.tencent.com/document/api/213/66323) 接口返回值中的`LaunchTemplateVersion`获取。
	LaunchTemplateVersions []*int64 `json:"LaunchTemplateVersions,omitnil,omitempty" name:"LaunchTemplateVersions"`
}

type DeleteLaunchTemplateVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 启动模板ID。可通过 [DescribeLaunchTemplates](https://cloud.tencent.com/document/api/213/66322) 接口返回值中的`LaunchTemplateId`获取。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 实例启动模板版本列表。可通过 [DescribeLaunchTemplateVersions](https://cloud.tencent.com/document/api/213/66323) 接口返回值中的`LaunchTemplateVersion`获取。
	LaunchTemplateVersions []*int64 `json:"LaunchTemplateVersions,omitnil,omitempty" name:"LaunchTemplateVersions"`
}

func (r *DeleteLaunchTemplateVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchTemplateVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateId")
	delete(f, "LaunchTemplateVersions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLaunchTemplateVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchTemplateVersionsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLaunchTemplateVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLaunchTemplateVersionsResponseParams `json:"Response"`
}

func (r *DeleteLaunchTemplateVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchTemplateVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountQuotaRequestParams struct {
	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>quota-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>配额类型</strong>】进行过滤。配额类型形如：PostPaidQuotaSet。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：</p><p style="padding-left: 30px;">PostPaidQuotaSet: 后付费配额</p><p style="padding-left: 30px;">PrePaidQuotaSet: 预付费配额</p><p style="padding-left: 30px;">SpotPaidQuotaSet: 竞价配额</p><p style="padding-left: 30px;">ImageQuotaSet: 镜像配额</p><p style="padding-left: 30px;">DisasterRecoverGroupQuotaSet: 置放群组配额</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAccountQuotaRequest struct {
	*tchttp.BaseRequest
	
	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>quota-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>配额类型</strong>】进行过滤。配额类型形如：PostPaidQuotaSet。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：</p><p style="padding-left: 30px;">PostPaidQuotaSet: 后付费配额</p><p style="padding-left: 30px;">PrePaidQuotaSet: 预付费配额</p><p style="padding-left: 30px;">SpotPaidQuotaSet: 竞价配额</p><p style="padding-left: 30px;">ImageQuotaSet: 镜像配额</p><p style="padding-left: 30px;">DisasterRecoverGroupQuotaSet: 置放群组配额</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAccountQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountQuotaResponseParams struct {
	// 用户appid
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 配额数据
	AccountQuotaOverview *AccountQuotaOverview `json:"AccountQuotaOverview,omitnil,omitempty" name:"AccountQuotaOverview"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountQuotaResponseParams `json:"Response"`
}

func (r *DescribeAccountQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChcDeniedActionsRequestParams struct {
	// CHC物理服务器实例id
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`
}

type DescribeChcDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// CHC物理服务器实例id
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`
}

func (r *DescribeChcDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChcDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChcDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChcDeniedActionsResponseParams struct {
	// CHC实例禁止操作信息
	ChcHostDeniedActionSet []*ChcHostDeniedActions `json:"ChcHostDeniedActionSet,omitnil,omitempty" name:"ChcHostDeniedActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChcDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChcDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeChcDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChcDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChcHostsRequestParams struct {
	// CHC物理服务器实例ID。每次请求的实例的上限为100。参数不支持同时指定`ChcIds`和`Filters`。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`

	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>instance-name</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例名称</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>instance-state</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例状态</strong>】进行过滤。状态类型详见[实例状态表](https://cloud.tencent.com/document/api/213/15753#InstanceStatus)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>device-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>设备类型</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>vpc-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>私有网络唯一ID</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>subnet-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>私有子网唯一ID</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>sn</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>设备SN</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeChcHostsRequest struct {
	*tchttp.BaseRequest
	
	// CHC物理服务器实例ID。每次请求的实例的上限为100。参数不支持同时指定`ChcIds`和`Filters`。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`

	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>instance-name</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例名称</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>instance-state</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例状态</strong>】进行过滤。状态类型详见[实例状态表](https://cloud.tencent.com/document/api/213/15753#InstanceStatus)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>device-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>设备类型</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>vpc-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>私有网络唯一ID</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>subnet-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>私有子网唯一ID</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>sn</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>设备SN</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeChcHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChcHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChcHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChcHostsResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的实例列表
	ChcHostSet []*ChcHost `json:"ChcHostSet,omitnil,omitempty" name:"ChcHostSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChcHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChcHostsResponseParams `json:"Response"`
}

func (r *DescribeChcHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChcHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverGroupQuotaRequestParams struct {

}

type DescribeDisasterRecoverGroupQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDisasterRecoverGroupQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverGroupQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisasterRecoverGroupQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverGroupQuotaResponseParams struct {
	// 可创建置放群组数量的上限。
	GroupQuota *int64 `json:"GroupQuota,omitnil,omitempty" name:"GroupQuota"`

	// 当前用户已经创建的置放群组数量。
	CurrentNum *int64 `json:"CurrentNum,omitnil,omitempty" name:"CurrentNum"`

	// 物理机类型容灾组内实例的配额数。
	CvmInHostGroupQuota *int64 `json:"CvmInHostGroupQuota,omitnil,omitempty" name:"CvmInHostGroupQuota"`

	// 交换机类型容灾组内实例的配额数。
	//
	// Deprecated: CvmInSwGroupQuota is deprecated.
	CvmInSwGroupQuota *int64 `json:"CvmInSwGroupQuota,omitnil,omitempty" name:"CvmInSwGroupQuota"`

	// 机架类型容灾组内实例的配额数。
	CvmInRackGroupQuota *int64 `json:"CvmInRackGroupQuota,omitnil,omitempty" name:"CvmInRackGroupQuota"`

	// 交换机类型容灾组内实例的配额数。
	CvmInSwitchGroupQuota *int64 `json:"CvmInSwitchGroupQuota,omitnil,omitempty" name:"CvmInSwitchGroupQuota"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDisasterRecoverGroupQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisasterRecoverGroupQuotaResponseParams `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverGroupQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverGroupsRequestParams struct {
	// 分散置放群组ID列表。每次请求允许操作的分散置放群组数量上限是10。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 分散置放群组名称，支持模糊匹配。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <li><strong>tag-key</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-value</strong></li> <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag:tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 分散置放群组ID列表。每次请求允许操作的分散置放群组数量上限是10。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 分散置放群组名称，支持模糊匹配。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <li><strong>tag-key</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-value</strong></li> <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag:tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDisasterRecoverGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisasterRecoverGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverGroupsResponseParams struct {
	// 分散置放群组信息列表。
	DisasterRecoverGroupSet []*DisasterRecoverGroup `json:"DisasterRecoverGroupSet,omitnil,omitempty" name:"DisasterRecoverGroupSet"`

	// 用户置放群组总量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisasterRecoverGroupsResponseParams `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsRequestParams struct {
	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>project-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>项目ID</strong>】进行过滤，可通过调用[DescribeProject](https://cloud.tencent.com/document/api/651/78725)查询已创建的项目列表或登录[控制台](https://console.cloud.tencent.com/cvm/index)进行查看；也可以调用[AddProject](https://cloud.tencent.com/document/api/651/81952)创建新的项目。项目ID形如：1002189。</p><p style="padding-left: 30px;">类型：Integer</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>host-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>[CDH](https://cloud.tencent.com/document/product/416) ID</strong>】进行过滤。[CDH](https://cloud.tencent.com/document/product/416) ID形如：host-xxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>host-name</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>CDH实例名称</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>host-state</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>CDH实例状态</strong>】进行过滤。（PENDING：创建中 | LAUNCH_FAILURE：创建失败 | RUNNING：运行中 | EXPIRED：已过期）</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>host-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>CDH机型</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag-value</strong></li> <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag:tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。使用请参考示例。</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest
	
	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>project-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>项目ID</strong>】进行过滤，可通过调用[DescribeProject](https://cloud.tencent.com/document/api/651/78725)查询已创建的项目列表或登录[控制台](https://console.cloud.tencent.com/cvm/index)进行查看；也可以调用[AddProject](https://cloud.tencent.com/document/api/651/81952)创建新的项目。项目ID形如：1002189。</p><p style="padding-left: 30px;">类型：Integer</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>host-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>[CDH](https://cloud.tencent.com/document/product/416) ID</strong>】进行过滤。[CDH](https://cloud.tencent.com/document/product/416) ID形如：host-xxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>host-name</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>CDH实例名称</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>host-state</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>CDH实例状态</strong>】进行过滤。（PENDING：创建中 | LAUNCH_FAILURE：创建失败 | RUNNING：运行中 | EXPIRED：已过期）</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>host-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>CDH机型</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag-value</strong></li> <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag:tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。使用请参考示例。</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsResponseParams struct {
	// 符合查询条件的cdh实例总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// cdh实例详细信息列表
	HostSet []*HostItem `json:"HostSet,omitnil,omitempty" name:"HostSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostsResponseParams `json:"Response"`
}

func (r *DescribeHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHpcClustersRequestParams struct {
	// 一个或多个待操作的高性能计算集群ID。集群ID信息可通过 [查询高性能集群信息](https://cloud.tencent.com/document/api/213/83220) 接口获取。每次请求高性能计算集群信息的批量上限为100，默认配合Limit传参数，不能超过Limit值，Limit默认20。
	HpcClusterIds []*string `json:"HpcClusterIds,omitnil,omitempty" name:"HpcClusterIds"`

	// 高性能计算集群名称，长度限制[1-60]。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 可用区信息。可用区信息可通过 [查询可用区信息](https://cloud.tencent.com/document/api/213/15707) 接口获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 偏移量, 默认值0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 本次请求量, 默认值20，范围限制为[1-100]。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 高性能计算集群类型，当前有三个取值：分别是CDC/CHC/STANDARD，其中STANDARD是默认的标准模式。
	HpcClusterType *string `json:"HpcClusterType,omitnil,omitempty" name:"HpcClusterType"`

	// 高性能计算集群对应的业务场景标识，当前只支持CDC场景类型。	
	HpcClusterBusinessId *string `json:"HpcClusterBusinessId,omitnil,omitempty" name:"HpcClusterBusinessId"`

	// 高性能计算集群实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <li><strong>tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag-value</strong></li> <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag:tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeHpcClustersRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的高性能计算集群ID。集群ID信息可通过 [查询高性能集群信息](https://cloud.tencent.com/document/api/213/83220) 接口获取。每次请求高性能计算集群信息的批量上限为100，默认配合Limit传参数，不能超过Limit值，Limit默认20。
	HpcClusterIds []*string `json:"HpcClusterIds,omitnil,omitempty" name:"HpcClusterIds"`

	// 高性能计算集群名称，长度限制[1-60]。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 可用区信息。可用区信息可通过 [查询可用区信息](https://cloud.tencent.com/document/api/213/15707) 接口获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 偏移量, 默认值0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 本次请求量, 默认值20，范围限制为[1-100]。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 高性能计算集群类型，当前有三个取值：分别是CDC/CHC/STANDARD，其中STANDARD是默认的标准模式。
	HpcClusterType *string `json:"HpcClusterType,omitnil,omitempty" name:"HpcClusterType"`

	// 高性能计算集群对应的业务场景标识，当前只支持CDC场景类型。	
	HpcClusterBusinessId *string `json:"HpcClusterBusinessId,omitnil,omitempty" name:"HpcClusterBusinessId"`

	// 高性能计算集群实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <li><strong>tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag-value</strong></li> <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag:tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeHpcClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHpcClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HpcClusterIds")
	delete(f, "Name")
	delete(f, "Zone")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "HpcClusterType")
	delete(f, "HpcClusterBusinessId")
	delete(f, "InstanceType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHpcClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHpcClustersResponseParams struct {
	// 高性能计算集群信息。
	HpcClusterSet []*HpcClusterInfo `json:"HpcClusterSet,omitnil,omitempty" name:"HpcClusterSet"`

	// 高性能计算集群总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHpcClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHpcClustersResponseParams `json:"Response"`
}

func (r *DescribeHpcClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHpcClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageFromFamilyRequestParams struct {
	// 镜像族
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`
}

type DescribeImageFromFamilyRequest struct {
	*tchttp.BaseRequest
	
	// 镜像族
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`
}

func (r *DescribeImageFromFamilyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageFromFamilyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageFamily")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageFromFamilyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageFromFamilyResponseParams struct {
	// 镜像信息，没有可用镜像是返回为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageFromFamilyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageFromFamilyResponseParams `json:"Response"`
}

func (r *DescribeImageFromFamilyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageFromFamilyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageQuotaRequestParams struct {

}

type DescribeImageQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImageQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageQuotaResponseParams struct {
	// 账户的镜像配额
	ImageNumQuota *int64 `json:"ImageNumQuota,omitnil,omitempty" name:"ImageNumQuota"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageQuotaResponseParams `json:"Response"`
}

func (r *DescribeImageQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSharePermissionRequestParams struct {
	// 需要共享的镜像 ID，可通过 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) 接口返回的`ImageId`获取。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

type DescribeImageSharePermissionRequest struct {
	*tchttp.BaseRequest
	
	// 需要共享的镜像 ID，可通过 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) 接口返回的`ImageId`获取。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

func (r *DescribeImageSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSharePermissionResponseParams struct {
	// 镜像共享信息
	SharePermissionSet []*SharePermission `json:"SharePermissionSet,omitnil,omitempty" name:"SharePermissionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageSharePermissionResponseParams `json:"Response"`
}

func (r *DescribeImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesRequestParams struct {
	// 镜像ID列表 。镜像ID如：`img-gvbnzy6f`。array型参数的格式可以参考[API简介](https://cloud.tencent.com/document/api/213/15688)。镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取。</li><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。</li>
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// 过滤条件，每次请求的`Filters`的上限为10，`Filters.Values`的上限为5。参数不可以同时指定`ImageIds`和`Filters`。详细的过滤条件如下：
	// 
	// <li><strong>image-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>镜像ID</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>image-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>镜像类型</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：</p><p style="padding-left: 30px;">PRIVATE_IMAGE: 自定义镜像 (本账户创建的镜像)</p><p style="padding-left: 30px;">PUBLIC_IMAGE: 公共镜像 (腾讯云官方镜像)</p><p style="padding-left: 30px;">SHARED_IMAGE: 共享镜像(其他账户共享给本账户的镜像)</p>
	// <li><strong>image-name</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>镜像名称</strong>】进行过滤。支持模糊查询。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>platform</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>镜像平台</strong>】进行过滤，如 CentOS，支持模糊匹配。可通过 <a href="https://cloud.tencent.com/document/api/213/15715" target="_blank">DescribeImages</a> 接口返回的<code> Platform </code>获取。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-key</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。可通过 <a href="https://cloud.tencent.com/document/product/651/72275" target="_blank"> GetTags </a> 接口返回的<code> TagKey </code>获取。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-value</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。可通过 <a href="https://cloud.tencent.com/document/product/651/72275" target="_blank"> GetTags </a> 接口返回的<code> TagValue </code>获取。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag:tag-key</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。可通过 <a href="https://cloud.tencent.com/document/product/651/72275" target="_blank"> GetTags </a> 接口返回的<code> TagKey 和 TagValue </code>获取。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset详见[API简介](/document/api/213/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。关于Limit详见[API简介](/document/api/213/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例类型，如 `S1.SMALL1`。可通过 [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/product/213/15749) 接口返回的 `InstanceType` 获取。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID列表 。镜像ID如：`img-gvbnzy6f`。array型参数的格式可以参考[API简介](https://cloud.tencent.com/document/api/213/15688)。镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取。</li><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。</li>
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// 过滤条件，每次请求的`Filters`的上限为10，`Filters.Values`的上限为5。参数不可以同时指定`ImageIds`和`Filters`。详细的过滤条件如下：
	// 
	// <li><strong>image-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>镜像ID</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>image-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>镜像类型</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：</p><p style="padding-left: 30px;">PRIVATE_IMAGE: 自定义镜像 (本账户创建的镜像)</p><p style="padding-left: 30px;">PUBLIC_IMAGE: 公共镜像 (腾讯云官方镜像)</p><p style="padding-left: 30px;">SHARED_IMAGE: 共享镜像(其他账户共享给本账户的镜像)</p>
	// <li><strong>image-name</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>镜像名称</strong>】进行过滤。支持模糊查询。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>platform</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>镜像平台</strong>】进行过滤，如 CentOS，支持模糊匹配。可通过 <a href="https://cloud.tencent.com/document/api/213/15715" target="_blank">DescribeImages</a> 接口返回的<code> Platform </code>获取。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-key</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。可通过 <a href="https://cloud.tencent.com/document/product/651/72275" target="_blank"> GetTags </a> 接口返回的<code> TagKey </code>获取。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-value</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。可通过 <a href="https://cloud.tencent.com/document/product/651/72275" target="_blank"> GetTags </a> 接口返回的<code> TagValue </code>获取。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag:tag-key</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。可通过 <a href="https://cloud.tencent.com/document/product/651/72275" target="_blank"> GetTags </a> 接口返回的<code> TagKey 和 TagValue </code>获取。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset详见[API简介](/document/api/213/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量限制，默认为20，最大值为100。关于Limit详见[API简介](/document/api/213/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例类型，如 `S1.SMALL1`。可通过 [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/product/213/15749) 接口返回的 `InstanceType` 获取。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// 一个关于镜像详细信息的结构体，主要包括镜像的主要状态与属性。
	ImageSet []*Image `json:"ImageSet,omitnil,omitempty" name:"ImageSet"`

	// 符合要求的镜像数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImportImageOsRequestParams struct {

}

type DescribeImportImageOsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImportImageOsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImportImageOsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImportImageOsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImportImageOsResponseParams struct {
	// 支持的导入镜像的操作系统类型。
	ImportImageOsListSupported *ImageOsList `json:"ImportImageOsListSupported,omitnil,omitempty" name:"ImportImageOsListSupported"`

	// 支持的导入镜像的操作系统版本。
	ImportImageOsVersionSet []*OsVersion `json:"ImportImageOsVersionSet,omitnil,omitempty" name:"ImportImageOsVersionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImportImageOsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImportImageOsResponseParams `json:"Response"`
}

func (r *DescribeImportImageOsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImportImageOsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceFamilyConfigsRequestParams struct {

}

type DescribeInstanceFamilyConfigsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInstanceFamilyConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceFamilyConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceFamilyConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceFamilyConfigsResponseParams struct {
	// 实例机型组配置的列表信息
	InstanceFamilyConfigSet []*InstanceFamilyConfig `json:"InstanceFamilyConfigSet,omitnil,omitempty" name:"InstanceFamilyConfigSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceFamilyConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceFamilyConfigsResponseParams `json:"Response"`
}

func (r *DescribeInstanceFamilyConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceFamilyConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceInternetBandwidthConfigsRequestParams struct {
	// 待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceInternetBandwidthConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceInternetBandwidthConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceInternetBandwidthConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceInternetBandwidthConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceInternetBandwidthConfigsResponseParams struct {
	// 带宽配置信息列表。
	InternetBandwidthConfigSet []*InternetBandwidthConfig `json:"InternetBandwidthConfigSet,omitnil,omitempty" name:"InternetBandwidthConfigSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceInternetBandwidthConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceInternetBandwidthConfigsResponseParams `json:"Response"`
}

func (r *DescribeInstanceInternetBandwidthConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceInternetBandwidthConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTypeConfigsRequestParams struct {
	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>instance-family</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例机型系列</strong>】进行过滤。实例机型系列形如：S1、I1、M1等。具体取值参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>instance-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例类型</strong>】进行过滤。实例类型形如：S5.12XLARGE128、S5.12XLARGE96等。具体取值参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为1。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeInstanceTypeConfigsRequest struct {
	*tchttp.BaseRequest
	
	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>instance-family</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例机型系列</strong>】进行过滤。实例机型系列形如：S1、I1、M1等。具体取值参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>instance-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例类型</strong>】进行过滤。实例类型形如：S5.12XLARGE128、S5.12XLARGE96等。具体取值参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为1。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeInstanceTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypeConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceTypeConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTypeConfigsResponseParams struct {
	// 实例机型配置列表。
	InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitnil,omitempty" name:"InstanceTypeConfigSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceTypeConfigsResponseParams `json:"Response"`
}

func (r *DescribeInstanceTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceVncUrlRequestParams struct {
	// 一个操作的实例ID。可通过[ DescribeInstances ](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceVncUrlRequest struct {
	*tchttp.BaseRequest
	
	// 一个操作的实例ID。可通过[ DescribeInstances ](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceVncUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceVncUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceVncUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceVncUrlResponseParams struct {
	// 实例的管理终端地址。
	InstanceVncUrl *string `json:"InstanceVncUrl,omitnil,omitempty" name:"InstanceVncUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceVncUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceVncUrlResponseParams `json:"Response"`
}

func (r *DescribeInstanceVncUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceVncUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesActionTimerRequestParams struct {
	// 定时任务ID列表。
	ActionTimerIds []*string `json:"ActionTimerIds,omitnil,omitempty" name:"ActionTimerIds"`

	// 按照一个或者多个实例ID查询。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 定时器动作，目前仅支持销毁一个值：TerminateInstances。
	TimerAction *string `json:"TimerAction,omitnil,omitempty" name:"TimerAction"`

	// 定时任务执行时间的结束范围，用于条件筛选，格式如2018-05-01 19:00:00。
	EndActionTime *string `json:"EndActionTime,omitnil,omitempty" name:"EndActionTime"`

	// 定时任务执行时间的开始范围，用于条件筛选，格式如2018-05-01 19:00:00。
	StartActionTime *string `json:"StartActionTime,omitnil,omitempty" name:"StartActionTime"`

	// 定时任务状态列表。<br><li>UNDO：未执行</li> <li>DOING：正在执行</li><li>DONE：执行完成。</li>
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`
}

type DescribeInstancesActionTimerRequest struct {
	*tchttp.BaseRequest
	
	// 定时任务ID列表。
	ActionTimerIds []*string `json:"ActionTimerIds,omitnil,omitempty" name:"ActionTimerIds"`

	// 按照一个或者多个实例ID查询。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 定时器动作，目前仅支持销毁一个值：TerminateInstances。
	TimerAction *string `json:"TimerAction,omitnil,omitempty" name:"TimerAction"`

	// 定时任务执行时间的结束范围，用于条件筛选，格式如2018-05-01 19:00:00。
	EndActionTime *string `json:"EndActionTime,omitnil,omitempty" name:"EndActionTime"`

	// 定时任务执行时间的开始范围，用于条件筛选，格式如2018-05-01 19:00:00。
	StartActionTime *string `json:"StartActionTime,omitnil,omitempty" name:"StartActionTime"`

	// 定时任务状态列表。<br><li>UNDO：未执行</li> <li>DOING：正在执行</li><li>DONE：执行完成。</li>
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`
}

func (r *DescribeInstancesActionTimerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesActionTimerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActionTimerIds")
	delete(f, "InstanceIds")
	delete(f, "TimerAction")
	delete(f, "EndActionTime")
	delete(f, "StartActionTime")
	delete(f, "StatusList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesActionTimerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesActionTimerResponseParams struct {
	// 定时任务信息列表。
	ActionTimers []*ActionTimer `json:"ActionTimers,omitnil,omitempty" name:"ActionTimers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesActionTimerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesActionTimerResponseParams `json:"Response"`
}

func (r *DescribeInstancesActionTimerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesActionTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesAttributesRequestParams struct {
	// 需要获取的实例属性。可选值：
	// UserData: 实例自定义数据
	Attributes []*string `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// 实例ID列表。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeInstancesAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取的实例属性。可选值：
	// UserData: 实例自定义数据
	Attributes []*string `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// 实例ID列表。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Attributes")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesAttributesResponseParams struct {
	// 指定的实例属性列表
	InstanceSet []*InstanceAttribute `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesAttributesResponseParams `json:"Response"`
}

func (r *DescribeInstancesAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesModificationRequestParams struct {
	// 一个或多个待查询的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为20。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <li><strong>status</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>配置规格状态</strong>】进行过滤。配置规格状态形如：SELL(表示实例可售卖)、UNAVAILABLE(表示实例不可用)。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为2。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeInstancesModificationRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待查询的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为20。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <li><strong>status</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>配置规格状态</strong>】进行过滤。配置规格状态形如：SELL(表示实例可售卖)、UNAVAILABLE(表示实例不可用)。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为2。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeInstancesModificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesModificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesModificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesModificationResponseParams struct {
	// 实例调整的机型配置的数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例支持调整的机型配置列表。
	InstanceTypeConfigStatusSet []*InstanceTypeConfigStatus `json:"InstanceTypeConfigStatusSet,omitnil,omitempty" name:"InstanceTypeConfigStatusSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesModificationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesModificationResponseParams `json:"Response"`
}

func (r *DescribeInstancesModificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesModificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesOperationLimitRequestParams struct {
	// 按照一个或者多个实例ID查询，可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728)API返回值中的InstanceId获取。实例ID形如：ins-xxxxxxxx。（此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的ids.N一节）。每次请求的实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例操作。
	// <li> INSTANCE_DEGRADE：实例降配操作</li><li> INTERNET_CHARGE_TYPE_CHANGE：实例调整带宽付费模式操作</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

type DescribeInstancesOperationLimitRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个实例ID查询，可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728)API返回值中的InstanceId获取。实例ID形如：ins-xxxxxxxx。（此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的ids.N一节）。每次请求的实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例操作。
	// <li> INSTANCE_DEGRADE：实例降配操作</li><li> INTERNET_CHARGE_TYPE_CHANGE：实例调整带宽付费模式操作</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

func (r *DescribeInstancesOperationLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesOperationLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesOperationLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesOperationLimitResponseParams struct {
	// 该参数表示调整配置操作（降配）限制次数查询。
	InstanceOperationLimitSet []*OperationCountLimit `json:"InstanceOperationLimitSet,omitnil,omitempty" name:"InstanceOperationLimitSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesOperationLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesOperationLimitResponseParams `json:"Response"`
}

func (r *DescribeInstancesOperationLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesOperationLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 按照一个或者多个实例ID查询。实例ID例如：`ins-xxxxxxxx`。（此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的`ids.N`一节）。每次请求的实例的上限为100。参数不支持同时指定`InstanceIds`和`Filters`。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <li><strong>zone</strong></li> <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区例如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p> <li><strong>project-id</strong></li> <p style="padding-left: 30px;">按照【<strong>项目ID</strong>】进行过滤，可通过调用[DescribeProjects](https://cloud.tencent.com/document/api/651/78725)查询已创建的项目列表或登录[控制台](https://console.cloud.tencent.com/cvm/index)进行查看；也可以调用[AddProject](https://cloud.tencent.com/document/api/651/81952)创建新的项目。项目ID例如：1002189。</p><p style="padding-left: 30px;">类型：Integer</p><p style="padding-left: 30px;">必选：否</p> <li><strong>host-id</strong></li> <p style="padding-left: 30px;">按照【<strong>[CDH](https://cloud.tencent.com/document/product/416) ID</strong>】进行过滤。[CDH](https://cloud.tencent.com/document/product/416) ID例如：host-xxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>dedicated-cluster-id</strong></li> <p style="padding-left: 30px;">按照【<strong>[CDC](https://cloud.tencent.com/document/product/1346) ID</strong>】进行过滤。[CDC](https://cloud.tencent.com/document/product/1346) ID例如：cluster-xxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>vpc-id</strong></li> <p style="padding-left: 30px;">按照【<strong>VPC ID</strong>】进行过滤。VPC ID例如：vpc-xxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>subnet-id</strong></li> <p style="padding-left: 30px;">按照【<strong>子网ID</strong>】进行过滤。子网ID例如：subnet-xxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>instance-id</strong></li> <p style="padding-left: 30px;">按照【<strong>实例ID</strong>】进行过滤。实例ID例如：ins-xxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>uuid</strong></li> <p style="padding-left: 30px;">按照【<strong>实例UUID</strong>】进行过滤。实例UUID例如：xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>security-group-id</strong></li> <p style="padding-left: 30px;">按照【<strong>安全组ID</strong>】进行过滤。安全组ID例如: sg-8jlk3f3r。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>instance-name</strong></li> <p style="padding-left: 30px;">按照【<strong>实例名称</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>instance-charge-type</strong></li> <p style="padding-left: 30px;">按照【<strong>实例计费模式</strong>】进行过滤。(PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费 | CDHPAID：表示[CDH](https://cloud.tencent.com/document/product/416)付费，即只对[CDH](https://cloud.tencent.com/document/product/416)计费，不对[CDH](https://cloud.tencent.com/document/product/416)上的实例计费。)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>instance-state</strong></li> <p style="padding-left: 30px;">按照【<strong>实例状态</strong>】进行过滤。状态类型详见[实例状态表](https://cloud.tencent.com/document/api/213/15753#InstanceStatus)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>private-ip-address</strong></li> <p style="padding-left: 30px;">按照【<strong>实例主网卡的内网IP</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>public-ip-address</strong></li> <p style="padding-left: 30px;">按照【<strong>实例主网卡的公网IP</strong>】进行过滤，包含实例创建时自动分配的IP和实例创建后手动绑定的弹性IP。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>ipv6-address</strong></li> <p style="padding-left: 30px;">按照【<strong>实例的IPv6地址</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag-value</strong></li> <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag:tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。使用请参考示例2。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><li><strong>creation-start-time</strong></li> <p style="padding-left: 30px;">按照【<strong>创建实例请求的起始时间</strong>】进行过滤。例如：2023-06-01 00:00:00。
	// </p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>creation-end-time</strong></li> <p style="padding-left: 30px;">按照【<strong>创建实例请求的截止时间</strong>】进行过滤。例如：2023-06-01 00:00:00。
	// </p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个实例ID查询。实例ID例如：`ins-xxxxxxxx`。（此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的`ids.N`一节）。每次请求的实例的上限为100。参数不支持同时指定`InstanceIds`和`Filters`。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <li><strong>zone</strong></li> <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区例如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p> <li><strong>project-id</strong></li> <p style="padding-left: 30px;">按照【<strong>项目ID</strong>】进行过滤，可通过调用[DescribeProjects](https://cloud.tencent.com/document/api/651/78725)查询已创建的项目列表或登录[控制台](https://console.cloud.tencent.com/cvm/index)进行查看；也可以调用[AddProject](https://cloud.tencent.com/document/api/651/81952)创建新的项目。项目ID例如：1002189。</p><p style="padding-left: 30px;">类型：Integer</p><p style="padding-left: 30px;">必选：否</p> <li><strong>host-id</strong></li> <p style="padding-left: 30px;">按照【<strong>[CDH](https://cloud.tencent.com/document/product/416) ID</strong>】进行过滤。[CDH](https://cloud.tencent.com/document/product/416) ID例如：host-xxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>dedicated-cluster-id</strong></li> <p style="padding-left: 30px;">按照【<strong>[CDC](https://cloud.tencent.com/document/product/1346) ID</strong>】进行过滤。[CDC](https://cloud.tencent.com/document/product/1346) ID例如：cluster-xxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>vpc-id</strong></li> <p style="padding-left: 30px;">按照【<strong>VPC ID</strong>】进行过滤。VPC ID例如：vpc-xxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>subnet-id</strong></li> <p style="padding-left: 30px;">按照【<strong>子网ID</strong>】进行过滤。子网ID例如：subnet-xxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>instance-id</strong></li> <p style="padding-left: 30px;">按照【<strong>实例ID</strong>】进行过滤。实例ID例如：ins-xxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>uuid</strong></li> <p style="padding-left: 30px;">按照【<strong>实例UUID</strong>】进行过滤。实例UUID例如：xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>security-group-id</strong></li> <p style="padding-left: 30px;">按照【<strong>安全组ID</strong>】进行过滤。安全组ID例如: sg-8jlk3f3r。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>instance-name</strong></li> <p style="padding-left: 30px;">按照【<strong>实例名称</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>instance-charge-type</strong></li> <p style="padding-left: 30px;">按照【<strong>实例计费模式</strong>】进行过滤。(PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费 | CDHPAID：表示[CDH](https://cloud.tencent.com/document/product/416)付费，即只对[CDH](https://cloud.tencent.com/document/product/416)计费，不对[CDH](https://cloud.tencent.com/document/product/416)上的实例计费。)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>instance-state</strong></li> <p style="padding-left: 30px;">按照【<strong>实例状态</strong>】进行过滤。状态类型详见[实例状态表](https://cloud.tencent.com/document/api/213/15753#InstanceStatus)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>private-ip-address</strong></li> <p style="padding-left: 30px;">按照【<strong>实例主网卡的内网IP</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>public-ip-address</strong></li> <p style="padding-left: 30px;">按照【<strong>实例主网卡的公网IP</strong>】进行过滤，包含实例创建时自动分配的IP和实例创建后手动绑定的弹性IP。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>ipv6-address</strong></li> <p style="padding-left: 30px;">按照【<strong>实例的IPv6地址</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag-value</strong></li> <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> <li><strong>tag:tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。使用请参考示例2。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><li><strong>creation-start-time</strong></li> <p style="padding-left: 30px;">按照【<strong>创建实例请求的起始时间</strong>】进行过滤。例如：2023-06-01 00:00:00。
	// </p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>creation-end-time</strong></li> <p style="padding-left: 30px;">按照【<strong>创建实例请求的截止时间</strong>】进行过滤。例如：2023-06-01 00:00:00。
	// </p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p> 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息列表。
	InstanceSet []*Instance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesStatusRequestParams struct {
	// 按照一个或者多个实例ID查询。实例ID形如：`ins-11112222`。此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的`ids.N`一节）。每次请求的实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstancesStatusRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个实例ID查询。实例ID形如：`ins-11112222`。此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的`ids.N`一节）。每次请求的实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeInstancesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesStatusResponseParams struct {
	// 符合条件的实例状态数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// [实例状态](https://cloud.tencent.com/document/api/213/15753#InstanceStatus) 列表。
	InstanceStatusSet []*InstanceStatus `json:"InstanceStatusSet,omitnil,omitempty" name:"InstanceStatusSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesStatusResponseParams `json:"Response"`
}

func (r *DescribeInstancesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetChargeTypeConfigsRequestParams struct {

}

type DescribeInternetChargeTypeConfigsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInternetChargeTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetChargeTypeConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternetChargeTypeConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetChargeTypeConfigsResponseParams struct {
	// 网络计费类型配置。
	InternetChargeTypeConfigSet []*InternetChargeTypeConfig `json:"InternetChargeTypeConfigSet,omitnil,omitempty" name:"InternetChargeTypeConfigSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInternetChargeTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternetChargeTypeConfigsResponseParams `json:"Response"`
}

func (r *DescribeInternetChargeTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetChargeTypeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeyPairsRequestParams struct {
	// 密钥对ID，密钥对ID形如：`skey-11112222`（此接口支持同时传入多个ID进行过滤。此参数的具体格式可参考 API [简介](https://cloud.tencent.com/document/api/213/15688)的 `id.N` 一节）。参数不支持同时指定 `KeyIds` 和 `Filters`。密钥对ID可以通过登录[控制台](https://console.cloud.tencent.com/cvm/index)查询。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 过滤条件。
	// <li> project-id - Integer - 是否必填：否 -（过滤条件）按照项目ID过滤。可以通过[项目列表](https://console.cloud.tencent.com/project)查询项目ID，或者调用接口 [DescribeProject](https://cloud.tencent.com/document/api/378/4400)，取返回信息中的projectId获取项目ID。</li>
	// <li> key-name - String - 是否必填：否 -（过滤条件）按照密钥对名称过滤。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对过滤。tag-key使用具体的标签键进行替换。</li>
	// 参数不支持同时指定 `KeyIds` 和 `Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对ID，密钥对ID形如：`skey-11112222`（此接口支持同时传入多个ID进行过滤。此参数的具体格式可参考 API [简介](https://cloud.tencent.com/document/api/213/15688)的 `id.N` 一节）。参数不支持同时指定 `KeyIds` 和 `Filters`。密钥对ID可以通过登录[控制台](https://console.cloud.tencent.com/cvm/index)查询。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 过滤条件。
	// <li> project-id - Integer - 是否必填：否 -（过滤条件）按照项目ID过滤。可以通过[项目列表](https://console.cloud.tencent.com/project)查询项目ID，或者调用接口 [DescribeProject](https://cloud.tencent.com/document/api/378/4400)，取返回信息中的projectId获取项目ID。</li>
	// <li> key-name - String - 是否必填：否 -（过滤条件）按照密钥对名称过滤。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对过滤。tag-key使用具体的标签键进行替换。</li>
	// 参数不支持同时指定 `KeyIds` 和 `Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeyPairsResponseParams struct {
	// 符合条件的密钥对数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 密钥对详细信息列表。
	KeyPairSet []*KeyPair `json:"KeyPairSet,omitnil,omitempty" name:"KeyPairSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeyPairsResponseParams `json:"Response"`
}

func (r *DescribeKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchTemplateVersionsRequestParams struct {
	// 启动模板ID。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 实例启动模板列表。
	LaunchTemplateVersions []*uint64 `json:"LaunchTemplateVersions,omitnil,omitempty" name:"LaunchTemplateVersions"`

	// 通过范围指定版本时的最小版本号，默认为0。
	MinVersion *uint64 `json:"MinVersion,omitnil,omitempty" name:"MinVersion"`

	// 过范围指定版本时的最大版本号，默认为30。
	MaxVersion *uint64 `json:"MaxVersion,omitnil,omitempty" name:"MaxVersion"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否查询默认版本。该参数不可与LaunchTemplateVersions同时指定。
	DefaultVersion *bool `json:"DefaultVersion,omitnil,omitempty" name:"DefaultVersion"`
}

type DescribeLaunchTemplateVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 启动模板ID。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 实例启动模板列表。
	LaunchTemplateVersions []*uint64 `json:"LaunchTemplateVersions,omitnil,omitempty" name:"LaunchTemplateVersions"`

	// 通过范围指定版本时的最小版本号，默认为0。
	MinVersion *uint64 `json:"MinVersion,omitnil,omitempty" name:"MinVersion"`

	// 过范围指定版本时的最大版本号，默认为30。
	MaxVersion *uint64 `json:"MaxVersion,omitnil,omitempty" name:"MaxVersion"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否查询默认版本。该参数不可与LaunchTemplateVersions同时指定。
	DefaultVersion *bool `json:"DefaultVersion,omitnil,omitempty" name:"DefaultVersion"`
}

func (r *DescribeLaunchTemplateVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchTemplateVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateId")
	delete(f, "LaunchTemplateVersions")
	delete(f, "MinVersion")
	delete(f, "MaxVersion")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DefaultVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLaunchTemplateVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchTemplateVersionsResponseParams struct {
	// 实例启动模板总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例启动模板版本集合。
	LaunchTemplateVersionSet []*LaunchTemplateVersionInfo `json:"LaunchTemplateVersionSet,omitnil,omitempty" name:"LaunchTemplateVersionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLaunchTemplateVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLaunchTemplateVersionsResponseParams `json:"Response"`
}

func (r *DescribeLaunchTemplateVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchTemplateVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchTemplatesRequestParams struct {
	// 启动模板ID，一个或者多个启动模板ID。若未指定，则显示用户所有模板。
	LaunchTemplateIds []*string `json:"LaunchTemplateIds,omitnil,omitempty" name:"LaunchTemplateIds"`

	// <li><strong>LaunchTemplateName</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例启动模板名称</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-key</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-value</strong></li> <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag:tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`LaunchTemplateIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeLaunchTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 启动模板ID，一个或者多个启动模板ID。若未指定，则显示用户所有模板。
	LaunchTemplateIds []*string `json:"LaunchTemplateIds,omitnil,omitempty" name:"LaunchTemplateIds"`

	// <li><strong>LaunchTemplateName</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例启动模板名称</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-key</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>标签键</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag-value</strong></li> <p style="padding-left: 30px;">按照【<strong>标签值</strong>】进行过滤。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>tag:tag-key</strong></li> <p style="padding-left: 30px;">按照【<strong>标签键值对</strong>】进行过滤。tag-key使用具体的标签键进行替换。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`LaunchTemplateIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeLaunchTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLaunchTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchTemplatesResponseParams struct {
	// 符合条件的实例模板数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息列表。
	LaunchTemplateSet []*LaunchTemplateInfo `json:"LaunchTemplateSet,omitnil,omitempty" name:"LaunchTemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLaunchTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLaunchTemplatesResponseParams `json:"Response"`
}

func (r *DescribeLaunchTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// 地域数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 地域列表信息。
	RegionSet []*RegionInfo `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstancesConfigInfosRequestParams struct {
	// zone
	// 按照预留实例计费可购买的可用区进行过滤。形如：ap-guangzhou-1。
	// 类型：String
	// 必选：否
	// 可选项：各地域可用区列表
	// 
	// product-description
	// 按照预留实例计费的平台描述（即操作系统）进行过滤。形如：linux。
	// 类型：String
	// 必选：否
	// 可选项：linux
	// 
	// duration
	// 按照预留实例计费有效期，即预留实例计费购买时长进行过滤。形如：31536000。
	// 类型：Integer
	// 计量单位：秒
	// 必选：否
	// 可选项：31536000 (1年)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeReservedInstancesConfigInfosRequest struct {
	*tchttp.BaseRequest
	
	// zone
	// 按照预留实例计费可购买的可用区进行过滤。形如：ap-guangzhou-1。
	// 类型：String
	// 必选：否
	// 可选项：各地域可用区列表
	// 
	// product-description
	// 按照预留实例计费的平台描述（即操作系统）进行过滤。形如：linux。
	// 类型：String
	// 必选：否
	// 可选项：linux
	// 
	// duration
	// 按照预留实例计费有效期，即预留实例计费购买时长进行过滤。形如：31536000。
	// 类型：Integer
	// 计量单位：秒
	// 必选：否
	// 可选项：31536000 (1年)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeReservedInstancesConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstancesConfigInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReservedInstancesConfigInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstancesConfigInfosResponseParams struct {
	// 预留实例静态配置信息列表。
	ReservedInstanceConfigInfos []*ReservedInstanceConfigInfoItem `json:"ReservedInstanceConfigInfos,omitnil,omitempty" name:"ReservedInstanceConfigInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReservedInstancesConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReservedInstancesConfigInfosResponseParams `json:"Response"`
}

func (r *DescribeReservedInstancesConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstancesConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstancesOfferingsRequestParams struct {
	// 试运行, 默认为 false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 以最大有效期作为过滤参数。
	// 计量单位: 秒
	// 默认为 94608000。
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// 以最小有效期作为过滤参数。
	// 计量单位: 秒
	// 默认为 2592000。
	MinDuration *int64 `json:"MinDuration,omitnil,omitempty" name:"MinDuration"`

	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照预留实例计费可购买的【<strong>可用区</strong>】进行过滤。形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>duration</strong></li>
	// <p style="padding-left: 30px;">按照预留实例计费【<strong>有效期</strong>】即预留实例计费购买时长进行过滤。形如：31536000。</p><p style="padding-left: 30px;">类型：Integer</p><p style="padding-left: 30px;">计量单位：秒</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：31536000 (1年) | 94608000（3年）</p>
	// <li><strong>instance-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>预留实例计费类型</strong>】进行过滤。形如：S3.MEDIUM4。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/11518">预留实例计费类型列表</a></p>
	// <li><strong>offering-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>付款类型</strong>】进行过滤。形如：All Upfront (预付全部费用)。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：All Upfront (预付全部费用)</p>
	// <li><strong>product-description</strong></li>
	// <p style="padding-left: 30px;">按照预留实例计费的【<strong>平台描述</strong>】（即操作系统）进行过滤。形如：linux。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：linux</p>
	// <li><strong>reserved-instances-offering-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>预留实例计费配置ID</strong>】进行过滤。形如：650c138f-ae7e-4750-952a-96841d6e9fc1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeReservedInstancesOfferingsRequest struct {
	*tchttp.BaseRequest
	
	// 试运行, 默认为 false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 以最大有效期作为过滤参数。
	// 计量单位: 秒
	// 默认为 94608000。
	MaxDuration *int64 `json:"MaxDuration,omitnil,omitempty" name:"MaxDuration"`

	// 以最小有效期作为过滤参数。
	// 计量单位: 秒
	// 默认为 2592000。
	MinDuration *int64 `json:"MinDuration,omitnil,omitempty" name:"MinDuration"`

	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照预留实例计费可购买的【<strong>可用区</strong>】进行过滤。形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>duration</strong></li>
	// <p style="padding-left: 30px;">按照预留实例计费【<strong>有效期</strong>】即预留实例计费购买时长进行过滤。形如：31536000。</p><p style="padding-left: 30px;">类型：Integer</p><p style="padding-left: 30px;">计量单位：秒</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：31536000 (1年) | 94608000（3年）</p>
	// <li><strong>instance-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>预留实例计费类型</strong>】进行过滤。形如：S3.MEDIUM4。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/11518">预留实例计费类型列表</a></p>
	// <li><strong>offering-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>付款类型</strong>】进行过滤。形如：All Upfront (预付全部费用)。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：All Upfront (预付全部费用)</p>
	// <li><strong>product-description</strong></li>
	// <p style="padding-left: 30px;">按照预留实例计费的【<strong>平台描述</strong>】（即操作系统）进行过滤。形如：linux。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：linux</p>
	// <li><strong>reserved-instances-offering-id</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>预留实例计费配置ID</strong>】进行过滤。形如：650c138f-ae7e-4750-952a-96841d6e9fc1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeReservedInstancesOfferingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstancesOfferingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DryRun")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MaxDuration")
	delete(f, "MinDuration")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReservedInstancesOfferingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstancesOfferingsResponseParams struct {
	// 符合条件的预留实例计费数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的预留实例计费列表。
	ReservedInstancesOfferingsSet []*ReservedInstancesOffering `json:"ReservedInstancesOfferingsSet,omitnil,omitempty" name:"ReservedInstancesOfferingsSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReservedInstancesOfferingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReservedInstancesOfferingsResponseParams `json:"Response"`
}

func (r *DescribeReservedInstancesOfferingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstancesOfferingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoRequestParams struct {
	// 返回数量，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 按照指定的产品类型查询，支持取值：
	// 
	// - `CVM`：云服务器
	// - `CDH`：专用宿主机
	// - `CPM2.0`：裸金属云服务器
	// 
	// 未传入或为空时，默认查询全部产品类型。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 按照一个或多个任务状态ID进行过滤。
	// `TaskStatus`（任务状态ID）与任务状态中文名的对应关系如下：
	// 
	// - `1`：待授权
	// - `2`：处理中
	// - `3`：已结束
	// - `4`：已预约
	// - `5`：已取消
	// - `6`：已避免
	// 
	// 各任务状态的具体含义，可参考 [任务状态](https://cloud.tencent.com/document/product/213/67789#.E4.BB.BB.E5.8A.A1.E7.8A.B6.E6.80.81)。
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 按照一个或多个任务类型ID进行过滤。
	// 
	// `TaskTypeId`（任务类型ID）与任务类型中文名的对应关系如下：
	// 
	// - `101`：实例运行隐患
	// - `102`：实例运行异常
	// - `103`：实例硬盘异常
	// - `104`：实例网络连接异常
	// - `105`：实例运行预警
	// - `106`：实例硬盘预警
	// - `107`：实例维护升级
	// 
	// 各任务类型的具体含义，可参考 [维修任务分类](https://cloud.tencent.com/document/product/213/67789#.E7.BB.B4.E4.BF.AE.E4.BB.BB.E5.8A.A1.E5.88.86.E7.B1.BB)。
	TaskTypeIds []*int64 `json:"TaskTypeIds,omitnil,omitempty" name:"TaskTypeIds"`

	// 按照一个或者多个任务ID查询。任务ID形如：`rep-xxxxxxxx`。
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 按照一个或者多个实例ID查询。实例ID形如：`ins-xxxxxxxx`，可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 按照一个或者多个实例名称查询。
	Aliases []*string `json:"Aliases,omitnil,omitempty" name:"Aliases"`

	// 时间查询区间的起始位置，会根据任务创建时间`CreateTime`进行过滤，格式为`YYYY-MM-DD hh:mm:ss`。未传入时默认为当天`00:00:00`。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 时间查询区间的终止位置，会根据任务创建时间`CreateTime`进行过滤，格式为`YYYY-MM-DD hh:mm:ss`。未传入时默认为当前时刻。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 指定返回维修任务列表的排序字段，目前支持：
	// 
	// - `CreateTime`：任务创建时间
	// - `AuthTime`：任务授权时间
	// - `EndTime`：任务结束时间
	// 
	// 未传入时或为空时，默认按`CreateTime`字段进行排序。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式，目前支持：
	// 
	// - `0`：升序（默认）
	// - `1`：降序
	// 
	// 未传入或为空时，默认按升序排序。
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 按照指定的产品类型查询，支持取值：
	// 
	// - `CVM`：云服务器
	// - `CDH`：专用宿主机
	// - `CPM2.0`：裸金属云服务器
	// 
	// 未传入或为空时，默认查询全部产品类型。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 按照一个或多个任务状态ID进行过滤。
	// `TaskStatus`（任务状态ID）与任务状态中文名的对应关系如下：
	// 
	// - `1`：待授权
	// - `2`：处理中
	// - `3`：已结束
	// - `4`：已预约
	// - `5`：已取消
	// - `6`：已避免
	// 
	// 各任务状态的具体含义，可参考 [任务状态](https://cloud.tencent.com/document/product/213/67789#.E4.BB.BB.E5.8A.A1.E7.8A.B6.E6.80.81)。
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 按照一个或多个任务类型ID进行过滤。
	// 
	// `TaskTypeId`（任务类型ID）与任务类型中文名的对应关系如下：
	// 
	// - `101`：实例运行隐患
	// - `102`：实例运行异常
	// - `103`：实例硬盘异常
	// - `104`：实例网络连接异常
	// - `105`：实例运行预警
	// - `106`：实例硬盘预警
	// - `107`：实例维护升级
	// 
	// 各任务类型的具体含义，可参考 [维修任务分类](https://cloud.tencent.com/document/product/213/67789#.E7.BB.B4.E4.BF.AE.E4.BB.BB.E5.8A.A1.E5.88.86.E7.B1.BB)。
	TaskTypeIds []*int64 `json:"TaskTypeIds,omitnil,omitempty" name:"TaskTypeIds"`

	// 按照一个或者多个任务ID查询。任务ID形如：`rep-xxxxxxxx`。
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 按照一个或者多个实例ID查询。实例ID形如：`ins-xxxxxxxx`，可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 按照一个或者多个实例名称查询。
	Aliases []*string `json:"Aliases,omitnil,omitempty" name:"Aliases"`

	// 时间查询区间的起始位置，会根据任务创建时间`CreateTime`进行过滤，格式为`YYYY-MM-DD hh:mm:ss`。未传入时默认为当天`00:00:00`。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 时间查询区间的终止位置，会根据任务创建时间`CreateTime`进行过滤，格式为`YYYY-MM-DD hh:mm:ss`。未传入时默认为当前时刻。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 指定返回维修任务列表的排序字段，目前支持：
	// 
	// - `CreateTime`：任务创建时间
	// - `AuthTime`：任务授权时间
	// - `EndTime`：任务结束时间
	// 
	// 未传入时或为空时，默认按`CreateTime`字段进行排序。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式，目前支持：
	// 
	// - `0`：升序（默认）
	// - `1`：降序
	// 
	// 未传入或为空时，默认按升序排序。
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Product")
	delete(f, "TaskStatus")
	delete(f, "TaskTypeIds")
	delete(f, "TaskIds")
	delete(f, "InstanceIds")
	delete(f, "Aliases")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoResponseParams struct {
	// 查询返回的维修任务总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询返回的维修任务列表。
	RepairTaskInfoSet []*RepairTaskInfo `json:"RepairTaskInfoSet,omitnil,omitempty" name:"RepairTaskInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInfoResponseParams `json:"Response"`
}

func (r *DescribeTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneInstanceConfigInfosRequestParams struct {
	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>instance-family</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例机型系列</strong>】进行过滤。实例机型系列形如：S1、I1、M1等。具体取值参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>instance-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例机型</strong>】进行过滤。不同实例机型指定了不同的资源规格，具体取值可通过调用接口 [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/product/213/15749) 来获得最新的规格表或参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则默认查询筛选条件下所有机型。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>instance-charge-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例计费模式</strong>】进行过滤。(PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费  | CDHPAID：表示独享子机 | SPOTPAID：表示竞价付费 | CDCPAID：表示专用集群付费)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>sort-keys</strong></li>
	// <p style="padding-left: 30px;">按关键字进行排序,格式为排序字段加排序方式，中间用冒号分隔。 例如： 按cpu数逆序排序 "cpu:desc", 按mem大小顺序排序 "mem:asc"</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest
	
	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>可用区</strong>】进行过滤。可用区形如：ap-guangzhou-1。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">可选项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a></p>
	// <li><strong>instance-family</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例机型系列</strong>】进行过滤。实例机型系列形如：S1、I1、M1等。具体取值参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>instance-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例机型</strong>】进行过滤。不同实例机型指定了不同的资源规格，具体取值可通过调用接口 [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/product/213/15749) 来获得最新的规格表或参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则默认查询筛选条件下所有机型。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>instance-charge-type</strong></li>
	// <p style="padding-left: 30px;">按照【<strong>实例计费模式</strong>】进行过滤。(PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费  | CDHPAID：表示独享子机 | SPOTPAID：表示竞价付费 | CDCPAID：表示专用集群付费)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// <li><strong>sort-keys</strong></li>
	// <p style="padding-left: 30px;">按关键字进行排序,格式为排序字段加排序方式，中间用冒号分隔。 例如： 按cpu数逆序排序 "cpu:desc", 按mem大小顺序排序 "mem:asc"</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeZoneInstanceConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneInstanceConfigInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneInstanceConfigInfosResponseParams struct {
	// 可用区机型配置列表。
	InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"InstanceTypeQuotaSet,omitnil,omitempty" name:"InstanceTypeQuotaSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneInstanceConfigInfosResponseParams `json:"Response"`
}

func (r *DescribeZoneInstanceConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {

}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 可用区数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 可用区列表信息。
	ZoneSet []*ZoneInfo `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateInstancesKeyPairsRequestParams struct {
	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。<br>可以通过以下方式获取可用的实例ID：
	// <li>通过登录[控制台](https://console.cloud.tencent.com/cvm/index)查询实例ID。</li>
	// <li>通过调用接口 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) ，取返回信息中的 `InstanceId` 获取实例ID。</li>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 密钥对ID列表，列表长度上限为100。可以通过以下方式获取可用的密钥ID：
	// <li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥ID。</li>
	// <li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) ，取返回信息中的 `KeyId` 获取密钥对ID。</li>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 是否强制关机，默认值为 false。常规场景下，建议手动关机后解绑密钥。取值范围：
	// <li>true：先执行强制关机，再解绑密钥。</li>
	// <li>false：不执行强制关机，仅支持对已关机状态实例进行解绑操作。</li>
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`
}

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。<br>可以通过以下方式获取可用的实例ID：
	// <li>通过登录[控制台](https://console.cloud.tencent.com/cvm/index)查询实例ID。</li>
	// <li>通过调用接口 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) ，取返回信息中的 `InstanceId` 获取实例ID。</li>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 密钥对ID列表，列表长度上限为100。可以通过以下方式获取可用的密钥ID：
	// <li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥ID。</li>
	// <li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) ，取返回信息中的 `KeyId` 获取密钥对ID。</li>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 是否强制关机，默认值为 false。常规场景下，建议手动关机后解绑密钥。取值范围：
	// <li>true：先执行强制关机，再解绑密钥。</li>
	// <li>false：不执行强制关机，仅支持对已关机状态实例进行解绑操作。</li>
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`
}

func (r *DisassociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateInstancesKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "KeyIds")
	delete(f, "ForceStop")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateInstancesKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateInstancesKeyPairsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateInstancesKeyPairsResponseParams `json:"Response"`
}

func (r *DisassociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateInstancesKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// 要解绑的`安全组ID`，类似sg-efil73jd，只支持解绑单个安全组。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`SecurityGroupIds`获取实例绑定的安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 被解绑的`实例ID`，类似ins-lesecurk，支持指定多个实例，每次请求批量实例的上限为100。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要解绑的`安全组ID`，类似sg-efil73jd，只支持解绑单个安全组。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`SecurityGroupIds`获取实例绑定的安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 被解绑的`实例ID`，类似ins-lesecurk，支持指定多个实例，每次请求批量实例的上限为100。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisasterRecoverGroup struct {
	// 分散置放群组id。
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitnil,omitempty" name:"DisasterRecoverGroupId"`

	// 分散置放群组名称，长度1-60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分散置放群组类型，取值范围：<br>
	// <li>HOST：物理机<br></li>
	// <li>SW：交换机<br></li>
	// <li>RACK：机架</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分散置放群组内最大容纳云服务器数量。
	CvmQuotaTotal *int64 `json:"CvmQuotaTotal,omitnil,omitempty" name:"CvmQuotaTotal"`

	// 分散置放群组内云服务器当前数量。
	CurrentNum *int64 `json:"CurrentNum,omitnil,omitempty" name:"CurrentNum"`

	// 分散置放群组内，云服务器id列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 分散置放群组创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 置放群组亲和度
	Affinity *int64 `json:"Affinity,omitnil,omitempty" name:"Affinity"`

	// 置放群组关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DisasterRecoverGroupQuota struct {
	// 可创建置放群组数量的上限。
	GroupQuota *int64 `json:"GroupQuota,omitnil,omitempty" name:"GroupQuota"`

	// 当前用户已经创建的置放群组数量。
	CurrentNum *int64 `json:"CurrentNum,omitnil,omitempty" name:"CurrentNum"`

	// 物理机类型容灾组内实例的配额数。
	CvmInHostGroupQuota *int64 `json:"CvmInHostGroupQuota,omitnil,omitempty" name:"CvmInHostGroupQuota"`

	// 交换机类型容灾组内实例的配额数。
	CvmInSwitchGroupQuota *int64 `json:"CvmInSwitchGroupQuota,omitnil,omitempty" name:"CvmInSwitchGroupQuota"`

	// 机架类型容灾组内实例的配额数。
	CvmInRackGroupQuota *int64 `json:"CvmInRackGroupQuota,omitnil,omitempty" name:"CvmInRackGroupQuota"`
}

type EnhancedService struct {
	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitnil,omitempty" name:"SecurityService"`

	// 开启云监控服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitnil,omitempty" name:"MonitorService"`

	// 开启云自动化助手服务（TencentCloud Automation Tools，TAT）。若不指定该参数，则公共镜像默认开启云自动化助手服务，其他镜像默认不开启云自动化助手服务。
	AutomationService *RunAutomationServiceEnabled `json:"AutomationService,omitnil,omitempty" name:"AutomationService"`
}

// Predefined struct for user
type EnterRescueModeRequestParams struct {
	// 需要进入救援模式的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 救援模式下系统密码。不同操作系统类型密码复杂度限制不一样，具体如下：<li>Linux实例密码必须8到30位，至少包括两项[a-z]，[A-Z]、[0-9] 和 [( ) \` ~ ! @ # $ % ^ & *  - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。</li><li>Windows实例密码必须12到30位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) \` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? /]中的特殊符号。</li>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 救援模式下系统用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 是否强制关机。本参数已弃用，推荐使用StopType，不可以与参数StopType同时使用。
	//
	// Deprecated: ForceStop is deprecated.
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`

	// 实例的关闭模式。取值范围：<br><li>SOFT_FIRST：表示在正常关闭失败后进行强制关闭</li><br><li>HARD：直接强制关闭</li><br><li>SOFT：仅软关机</li><br>默认取值：SOFT。
	StopType *string `json:"StopType,omitnil,omitempty" name:"StopType"`
}

type EnterRescueModeRequest struct {
	*tchttp.BaseRequest
	
	// 需要进入救援模式的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 救援模式下系统密码。不同操作系统类型密码复杂度限制不一样，具体如下：<li>Linux实例密码必须8到30位，至少包括两项[a-z]，[A-Z]、[0-9] 和 [( ) \` ~ ! @ # $ % ^ & *  - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。</li><li>Windows实例密码必须12到30位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) \` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? /]中的特殊符号。</li>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 救援模式下系统用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 是否强制关机。本参数已弃用，推荐使用StopType，不可以与参数StopType同时使用。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`

	// 实例的关闭模式。取值范围：<br><li>SOFT_FIRST：表示在正常关闭失败后进行强制关闭</li><br><li>HARD：直接强制关闭</li><br><li>SOFT：仅软关机</li><br>默认取值：SOFT。
	StopType *string `json:"StopType,omitnil,omitempty" name:"StopType"`
}

func (r *EnterRescueModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnterRescueModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	delete(f, "Username")
	delete(f, "ForceStop")
	delete(f, "StopType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnterRescueModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnterRescueModeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnterRescueModeResponse struct {
	*tchttp.BaseResponse
	Response *EnterRescueModeResponseParams `json:"Response"`
}

func (r *EnterRescueModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnterRescueModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExitRescueModeRequestParams struct {
	// 退出救援模式的实例id。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ExitRescueModeRequest struct {
	*tchttp.BaseRequest
	
	// 退出救援模式的实例id。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ExitRescueModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExitRescueModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExitRescueModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExitRescueModeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExitRescueModeResponse struct {
	*tchttp.BaseResponse
	Response *ExitRescueModeResponseParams `json:"Response"`
}

func (r *ExitRescueModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExitRescueModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportImagesRequestParams struct {
	// COS存储桶名称。
	// 可通过 [List Buckets](https://cloud.tencent.com/document/product/436/8291) 接口查询请求者名下的所有存储桶列表或特定地域下的存储桶列表。
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 镜像ID列表。调用 ExportImages 接口时，参数 ImageIds 和 SnapshotIds 为二选一必填参数，目前参数 SnapshotIds 暂未对外开放。
	// 可通过 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) 接口返回值中的`ImageId`获取。
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// 镜像文件导出格式。取值范围：RAW，QCOW2，VHD，VMDK。默认为RAW
	ExportFormat *string `json:"ExportFormat,omitnil,omitempty" name:"ExportFormat"`

	// 导出文件的名称前缀列表。
	// 默认导出文件无名称前缀。
	FileNamePrefixList []*string `json:"FileNamePrefixList,omitnil,omitempty" name:"FileNamePrefixList"`

	// 是否只导出系统盘。
	// 默认值：false
	OnlyExportRootDisk *bool `json:"OnlyExportRootDisk,omitnil,omitempty" name:"OnlyExportRootDisk"`

	// 检测镜像是否支持导出。
	// 默认值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 角色名称。默认为CVM_QcsRole，发起请求前请确认是否存在该角色，以及是否已正确配置COS写入权限。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type ExportImagesRequest struct {
	*tchttp.BaseRequest
	
	// COS存储桶名称。
	// 可通过 [List Buckets](https://cloud.tencent.com/document/product/436/8291) 接口查询请求者名下的所有存储桶列表或特定地域下的存储桶列表。
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 镜像ID列表。调用 ExportImages 接口时，参数 ImageIds 和 SnapshotIds 为二选一必填参数，目前参数 SnapshotIds 暂未对外开放。
	// 可通过 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) 接口返回值中的`ImageId`获取。
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// 镜像文件导出格式。取值范围：RAW，QCOW2，VHD，VMDK。默认为RAW
	ExportFormat *string `json:"ExportFormat,omitnil,omitempty" name:"ExportFormat"`

	// 导出文件的名称前缀列表。
	// 默认导出文件无名称前缀。
	FileNamePrefixList []*string `json:"FileNamePrefixList,omitnil,omitempty" name:"FileNamePrefixList"`

	// 是否只导出系统盘。
	// 默认值：false
	OnlyExportRootDisk *bool `json:"OnlyExportRootDisk,omitnil,omitempty" name:"OnlyExportRootDisk"`

	// 检测镜像是否支持导出。
	// 默认值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 角色名称。默认为CVM_QcsRole，发起请求前请确认是否存在该角色，以及是否已正确配置COS写入权限。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *ExportImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BucketName")
	delete(f, "ImageIds")
	delete(f, "ExportFormat")
	delete(f, "FileNamePrefixList")
	delete(f, "OnlyExportRootDisk")
	delete(f, "DryRun")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportImagesResponseParams struct {
	// 导出镜像任务ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 导出镜像的COS文件名列表。其中，文件名格式如下。
	// * 系统盘：前缀名_镜像ID_system_快照ID.镜像格式
	// * 数据盘：前缀名_镜像ID_data_快照ID.镜像格式
	CosPaths []*string `json:"CosPaths,omitnil,omitempty" name:"CosPaths"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportImagesResponse struct {
	*tchttp.BaseResponse
	Response *ExportImagesResponseParams `json:"Response"`
}

func (r *ExportImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Externals struct {
	// 释放地址
	ReleaseAddress *bool `json:"ReleaseAddress,omitnil,omitempty" name:"ReleaseAddress"`

	// 不支持的网络类型，取值范围：<br><li>BASIC：基础网络</li><li>VPC1.0：私有网络VPC1.0</li>
	UnsupportNetworks []*string `json:"UnsupportNetworks,omitnil,omitempty" name:"UnsupportNetworks"`

	// HDD本地存储属性
	StorageBlockAttr *StorageBlock `json:"StorageBlockAttr,omitnil,omitempty" name:"StorageBlockAttr"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type GPUInfo struct {
	// 实例GPU个数。值小于1代表VGPU类型，大于1代表GPU直通类型。
	GPUCount *float64 `json:"GPUCount,omitnil,omitempty" name:"GPUCount"`

	// 实例GPU地址。
	GPUId []*string `json:"GPUId,omitnil,omitempty" name:"GPUId"`

	// 实例GPU类型。
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`
}

type HostItem struct {
	// 专用宿主机实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 专用宿主机实例ID
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 专用宿主机实例类型
	HostType *string `json:"HostType,omitnil,omitempty" name:"HostType"`

	// 专用宿主机实例名称
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 专用宿主机实例付费模式
	HostChargeType *string `json:"HostChargeType,omitnil,omitempty" name:"HostChargeType"`

	// 专用宿主机实例自动续费标记
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 专用宿主机实例创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 专用宿主机实例过期时间
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 专用宿主机实例上已创建云子机的实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 专用宿主机实例状态
	HostState *string `json:"HostState,omitnil,omitempty" name:"HostState"`

	// 专用宿主机实例IP
	HostIp *string `json:"HostIp,omitnil,omitempty" name:"HostIp"`

	// 专用宿主机实例资源信息
	HostResource *HostResource `json:"HostResource,omitnil,omitempty" name:"HostResource"`

	// 专用宿主机所属的围笼ID。该字段仅对金融专区围笼内的专用宿主机有效。
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// 专用宿主机关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type HostPriceInfo struct {
	// 描述了cdh实例相关的价格信息
	HostPrice *ItemPrice `json:"HostPrice,omitnil,omitempty" name:"HostPrice"`
}

type HostResource struct {
	// 专用宿主机实例总CPU核数
	CpuTotal *uint64 `json:"CpuTotal,omitnil,omitempty" name:"CpuTotal"`

	// 专用宿主机实例可用CPU核数
	CpuAvailable *uint64 `json:"CpuAvailable,omitnil,omitempty" name:"CpuAvailable"`

	// 专用宿主机实例总内存大小（单位为:GiB）
	MemTotal *float64 `json:"MemTotal,omitnil,omitempty" name:"MemTotal"`

	// 专用宿主机实例可用内存大小（单位为:GiB）
	MemAvailable *float64 `json:"MemAvailable,omitnil,omitempty" name:"MemAvailable"`

	// 专用宿主机实例总磁盘大小（单位为:GiB）
	DiskTotal *uint64 `json:"DiskTotal,omitnil,omitempty" name:"DiskTotal"`

	// 专用宿主机实例可用磁盘大小（单位为:GiB）
	DiskAvailable *uint64 `json:"DiskAvailable,omitnil,omitempty" name:"DiskAvailable"`

	// 专用宿主机实例磁盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 专用宿主机实例总GPU卡数
	GpuTotal *uint64 `json:"GpuTotal,omitnil,omitempty" name:"GpuTotal"`

	// 专用宿主机实例可用GPU卡数
	GpuAvailable *uint64 `json:"GpuAvailable,omitnil,omitempty" name:"GpuAvailable"`

	// CDH owner
	ExclusiveOwner *string `json:"ExclusiveOwner,omitnil,omitempty" name:"ExclusiveOwner"`
}

type HpcClusterInfo struct {
	// 高性能计算集群ID
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 高性能计算集群名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 高性能计算集群备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 集群下设备容量
	CvmQuotaTotal *uint64 `json:"CvmQuotaTotal,omitnil,omitempty" name:"CvmQuotaTotal"`

	// 集群所在可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群当前已有设备量
	CurrentNum *uint64 `json:"CurrentNum,omitnil,omitempty" name:"CurrentNum"`

	// 集群创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 集群内实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 高性能计算集群类型。
	HpcClusterType *string `json:"HpcClusterType,omitnil,omitempty" name:"HpcClusterType"`

	// 高性能计算集群对应的业务场景标识，当前只支持CDC。	
	HpcClusterBusinessId *string `json:"HpcClusterBusinessId,omitnil,omitempty" name:"HpcClusterBusinessId"`

	// 高性能计算集群网络模式
	HpcClusterNetMode *uint64 `json:"HpcClusterNetMode,omitnil,omitempty" name:"HpcClusterNetMode"`

	// 高性能计算集群关联的标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type Image struct {
	// 镜像ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 镜像操作系统
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 镜像类型。镜像类型返回值包括：
	// * `PUBLIC_IMAGE` 公共镜像
	// * `PRIVATE_IMAGE` 自定义镜像
	// * `SHARED_IMAGE` 共享镜像
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// 镜像创建时间。
	// 按照 ISO8601 标准表示，并且使用 UTC 时间，格式为：YYYY-MM-DDThh:mm:ssZ。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitnil,omitempty" name:"ImageDescription"`

	// 镜像大小，单位 GiB。
	ImageSize *int64 `json:"ImageSize,omitnil,omitempty" name:"ImageSize"`

	// 镜像架构。镜像架构返回值包括：
	// * `x86_64`
	// * `arm`
	// * `i386`
	Architecture *string `json:"Architecture,omitnil,omitempty" name:"Architecture"`

	// 镜像状态:
	// CREATING-创建中
	// NORMAL-正常
	// CREATEFAILED-创建失败
	// SYNCING-复制中
	// IMPORTING-导入中
	// IMPORTFAILED-导入失败
	ImageState *string `json:"ImageState,omitnil,omitempty" name:"ImageState"`

	// 镜像来源平台，包括如TencentOS、 CentOS、 Windows、 Ubuntu、 Debian、Fedora等。
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 镜像创建者
	ImageCreator *string `json:"ImageCreator,omitnil,omitempty" name:"ImageCreator"`

	// 镜像来源。镜像来源返回值包括：
	// * `OFFICIAL` 官方镜像
	// * `CREATE_IMAGE` 用户自建镜像
	// * `EXTERNAL_IMPORT` 用户外部导入镜像
	ImageSource *string `json:"ImageSource,omitnil,omitempty" name:"ImageSource"`

	// 同步百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	SyncPercent *int64 `json:"SyncPercent,omitnil,omitempty" name:"SyncPercent"`

	// 镜像是否支持cloud-init
	IsSupportCloudinit *bool `json:"IsSupportCloudinit,omitnil,omitempty" name:"IsSupportCloudinit"`

	// 镜像关联的快照信息
	SnapshotSet []*Snapshot `json:"SnapshotSet,omitnil,omitempty" name:"SnapshotSet"`

	// 镜像关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 镜像许可类型。镜像许可类型返回值包括：
	// * `TencentCloud` 腾讯云官方许可
	// * `BYOL` 用户自带许可
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 镜像族
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`

	// 镜像是否废弃
	ImageDeprecated *bool `json:"ImageDeprecated,omitnil,omitempty" name:"ImageDeprecated"`

	// CDC镜像缓存状态
	CdcCacheStatus *string `json:"CdcCacheStatus,omitnil,omitempty" name:"CdcCacheStatus"`
}

type ImageOsList struct {
	// 支持的Windows操作系统。
	Windows []*string `json:"Windows,omitnil,omitempty" name:"Windows"`

	// 支持的Linux操作系统
	Linux []*string `json:"Linux,omitnil,omitempty" name:"Linux"`
}

type ImageQuota struct {
	// 已使用配额
	UsedQuota *uint64 `json:"UsedQuota,omitnil,omitempty" name:"UsedQuota"`

	// 总配额
	TotalQuota *uint64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`
}

type ImportImageDataDisk struct {
	// 数据盘镜像 COS 链接
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

// Predefined struct for user
type ImportImageRequestParams struct {
	// 导入镜像的操作系统架构。
	// 取值范围包括：`x86_64` 、`i386`、`arm_64`
	Architecture *string `json:"Architecture,omitnil,omitempty" name:"Architecture"`

	// 导入镜像的操作系统类型 。
	// 可通过 [DescribeImportImageOs](https://cloud.tencent.com/document/api/213/15718) 接口返回值中的`ImportImageOsListSupported`获取。
	OsType *string `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 导入镜像的操作系统版本。
	// 可通过 [DescribeImportImageOs](https://cloud.tencent.com/document/api/213/15718) 接口返回值中的`ImportImageOsVersionSet`获取。
	OsVersion *string `json:"OsVersion,omitnil,omitempty" name:"OsVersion"`

	// 导入镜像存放的cos地址
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 镜像名称。
	// 最多支持 60 个字符。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像描述。
	// 最多支持 256 个字符。
	ImageDescription *string `json:"ImageDescription,omitnil,omitempty" name:"ImageDescription"`

	// 只检查参数，不执行任务。
	// 默认值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 是否强制导入，参考[强制导入镜像](https://cloud.tencent.com/document/product/213/12849)
	// 默认值：false
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到自定义镜像。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 导入镜像后，激活操作系统采用的许可证类型。
	// 默认值：TencentCloud
	// 取值范围：
	// TencentCloud: 腾讯云官方许可
	// BYOL: 自带许可（Bring Your Own License）
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 启动模式。
	// 取值范围：`Legacy BIOS`、`UEFI`
	// 默认值：Legacy BIOS
	BootMode *string `json:"BootMode,omitnil,omitempty" name:"BootMode"`

	//  镜像族
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`

	// 导入的数据盘列表
	ImportImageDataDiskList []*ImportImageDataDisk `json:"ImportImageDataDiskList,omitnil,omitempty" name:"ImportImageDataDiskList"`
}

type ImportImageRequest struct {
	*tchttp.BaseRequest
	
	// 导入镜像的操作系统架构。
	// 取值范围包括：`x86_64` 、`i386`、`arm_64`
	Architecture *string `json:"Architecture,omitnil,omitempty" name:"Architecture"`

	// 导入镜像的操作系统类型 。
	// 可通过 [DescribeImportImageOs](https://cloud.tencent.com/document/api/213/15718) 接口返回值中的`ImportImageOsListSupported`获取。
	OsType *string `json:"OsType,omitnil,omitempty" name:"OsType"`

	// 导入镜像的操作系统版本。
	// 可通过 [DescribeImportImageOs](https://cloud.tencent.com/document/api/213/15718) 接口返回值中的`ImportImageOsVersionSet`获取。
	OsVersion *string `json:"OsVersion,omitnil,omitempty" name:"OsVersion"`

	// 导入镜像存放的cos地址
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 镜像名称。
	// 最多支持 60 个字符。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像描述。
	// 最多支持 256 个字符。
	ImageDescription *string `json:"ImageDescription,omitnil,omitempty" name:"ImageDescription"`

	// 只检查参数，不执行任务。
	// 默认值：false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 是否强制导入，参考[强制导入镜像](https://cloud.tencent.com/document/product/213/12849)
	// 默认值：false
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到自定义镜像。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 导入镜像后，激活操作系统采用的许可证类型。
	// 默认值：TencentCloud
	// 取值范围：
	// TencentCloud: 腾讯云官方许可
	// BYOL: 自带许可（Bring Your Own License）
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 启动模式。
	// 取值范围：`Legacy BIOS`、`UEFI`
	// 默认值：Legacy BIOS
	BootMode *string `json:"BootMode,omitnil,omitempty" name:"BootMode"`

	//  镜像族
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`

	// 导入的数据盘列表
	ImportImageDataDiskList []*ImportImageDataDisk `json:"ImportImageDataDiskList,omitnil,omitempty" name:"ImportImageDataDiskList"`
}

func (r *ImportImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Architecture")
	delete(f, "OsType")
	delete(f, "OsVersion")
	delete(f, "ImageUrl")
	delete(f, "ImageName")
	delete(f, "ImageDescription")
	delete(f, "DryRun")
	delete(f, "Force")
	delete(f, "TagSpecification")
	delete(f, "LicenseType")
	delete(f, "BootMode")
	delete(f, "ImageFamily")
	delete(f, "ImportImageDataDiskList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportImageResponseParams struct {
	// 镜像 ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportImageResponse struct {
	*tchttp.BaseResponse
	Response *ImportImageResponseParams `json:"Response"`
}

func (r *ImportImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportInstancesActionTimerRequestParams struct {
	// 实例id列表，可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 定时器任务信息，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer,omitnil,omitempty" name:"ActionTimer"`
}

type ImportInstancesActionTimerRequest struct {
	*tchttp.BaseRequest
	
	// 实例id列表，可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 定时器任务信息，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer,omitnil,omitempty" name:"ActionTimer"`
}

func (r *ImportInstancesActionTimerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportInstancesActionTimerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ActionTimer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportInstancesActionTimerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportInstancesActionTimerResponseParams struct {
	// 定时器id列表
	ActionTimerIds []*string `json:"ActionTimerIds,omitnil,omitempty" name:"ActionTimerIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportInstancesActionTimerResponse struct {
	*tchttp.BaseResponse
	Response *ImportInstancesActionTimerResponseParams `json:"Response"`
}

func (r *ImportInstancesActionTimerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportInstancesActionTimerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportKeyPairRequestParams struct {
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 密钥对创建后所属的[项目](https://cloud.tencent.com/document/product/378/10861)ID。<br>可以通过以下方式获取项目ID：<li>通过[项目列表](https://console.cloud.tencent.com/project)查询项目ID。</li><li>通过调用接口 [DescribeProjects](https://cloud.tencent.com/document/api/651/78725)，取返回信息中的 `projectId ` 获取项目ID。</li>如果是默认项目，直接填0就可以。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 密钥对的公钥内容，`OpenSSH RSA` 格式。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到密钥对。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

type ImportKeyPairRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 密钥对创建后所属的[项目](https://cloud.tencent.com/document/product/378/10861)ID。<br>可以通过以下方式获取项目ID：<li>通过[项目列表](https://console.cloud.tencent.com/project)查询项目ID。</li><li>通过调用接口 [DescribeProjects](https://cloud.tencent.com/document/api/651/78725)，取返回信息中的 `projectId ` 获取项目ID。</li>如果是默认项目，直接填0就可以。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 密钥对的公钥内容，`OpenSSH RSA` 格式。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到密钥对。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

func (r *ImportKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportKeyPairRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyName")
	delete(f, "ProjectId")
	delete(f, "PublicKey")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportKeyPairResponseParams struct {
	// 密钥对ID。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *ImportKeyPairResponseParams `json:"Response"`
}

func (r *ImportKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePricePurchaseReservedInstancesOfferingRequestParams struct {
	// 购买预留实例计费数量
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 预留实例计费配置ID
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitnil,omitempty" name:"ReservedInstancesOfferingId"`

	// 试运行
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 预留实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>最多支持60个字符（包含模式串）。</li>
	ReservedInstanceName *string `json:"ReservedInstanceName,omitnil,omitempty" name:"ReservedInstanceName"`
}

type InquirePricePurchaseReservedInstancesOfferingRequest struct {
	*tchttp.BaseRequest
	
	// 购买预留实例计费数量
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 预留实例计费配置ID
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitnil,omitempty" name:"ReservedInstancesOfferingId"`

	// 试运行
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 预留实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>最多支持60个字符（包含模式串）。</li>
	ReservedInstanceName *string `json:"ReservedInstanceName,omitnil,omitempty" name:"ReservedInstanceName"`
}

func (r *InquirePricePurchaseReservedInstancesOfferingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePricePurchaseReservedInstancesOfferingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceCount")
	delete(f, "ReservedInstancesOfferingId")
	delete(f, "DryRun")
	delete(f, "ClientToken")
	delete(f, "ReservedInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePricePurchaseReservedInstancesOfferingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePricePurchaseReservedInstancesOfferingResponseParams struct {
	// 该参数表示对应配置预留实例的价格。
	Price *ReservedInstancePrice `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePricePurchaseReservedInstancesOfferingResponse struct {
	*tchttp.BaseResponse
	Response *InquirePricePurchaseReservedInstancesOfferingResponseParams `json:"Response"`
}

func (r *InquirePricePurchaseReservedInstancesOfferingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePricePurchaseReservedInstancesOfferingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceModifyInstancesChargeTypeRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 修改后的实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月。</li><li>POSTPAID_BY_HOUR：后付费，即按量付费。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。<dx-alert infotype="explain" title="">若指定修改后实例的付费模式为预付费则该参数必传。</dx-alert>
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 是否同时切换弹性数据云盘计费模式。取值范围：<br><li>true：表示切换弹性数据云盘计费模式</li><li>false：表示不切换弹性数据云盘计费模式</li><br>默认取值：false。
	ModifyPortableDataDisk *bool `json:"ModifyPortableDataDisk,omitnil,omitempty" name:"ModifyPortableDataDisk"`
}

type InquiryPriceModifyInstancesChargeTypeRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 修改后的实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月。</li><li>POSTPAID_BY_HOUR：后付费，即按量付费。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。<dx-alert infotype="explain" title="">若指定修改后实例的付费模式为预付费则该参数必传。</dx-alert>
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 是否同时切换弹性数据云盘计费模式。取值范围：<br><li>true：表示切换弹性数据云盘计费模式</li><li>false：表示不切换弹性数据云盘计费模式</li><br>默认取值：false。
	ModifyPortableDataDisk *bool `json:"ModifyPortableDataDisk,omitnil,omitempty" name:"ModifyPortableDataDisk"`
}

func (r *InquiryPriceModifyInstancesChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceModifyInstancesChargeTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "ModifyPortableDataDisk")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceModifyInstancesChargeTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceModifyInstancesChargeTypeResponseParams struct {
	// 该参数表示对应配置实例转换计费模式的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceModifyInstancesChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceModifyInstancesChargeTypeResponseParams `json:"Response"`
}

func (r *InquiryPriceModifyInstancesChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceModifyInstancesChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewHostsRequestParams struct {
	// 一个或多个待操作的`CDH`实例`ID`。可通过[`DescribeHosts`](https://cloud.tencent.com/document/api/213/16474)接口返回值中的`HostId`获取。每次请求批量实例的上限为100。
	HostIds []*string `json:"HostIds,omitnil,omitempty" name:"HostIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitnil,omitempty" name:"HostChargePrepaid"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type InquiryPriceRenewHostsRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的`CDH`实例`ID`。可通过[`DescribeHosts`](https://cloud.tencent.com/document/api/213/16474)接口返回值中的`HostId`获取。每次请求批量实例的上限为100。
	HostIds []*string `json:"HostIds,omitnil,omitempty" name:"HostIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitnil,omitempty" name:"HostChargePrepaid"`

	// 是否只预检此次请求。true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。如果检查不通过，则返回对应错误码；如果检查通过，则返回RequestId.false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *InquiryPriceRenewHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostIds")
	delete(f, "HostChargePrepaid")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewHostsResponseParams struct {
	// CDH实例续费价格信息
	Price *HostPriceInfo `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRenewHostsResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewHostsResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewInstancesRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 试运行，测试使用，不执行具体逻辑。取值范围：<br><li>true：跳过执行逻辑</li><li>false：执行逻辑<br><br>默认取值：false。</li>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 是否续费弹性数据盘。取值范围：<br><li>true：表示续费包年包月实例同时续费其挂载的弹性数据盘</li><li>false：表示续费包年包月实例同时不再续费其挂载的弹性数据盘</li><br>默认取值：true。
	RenewPortableDataDisk *bool `json:"RenewPortableDataDisk,omitnil,omitempty" name:"RenewPortableDataDisk"`
}

type InquiryPriceRenewInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 试运行，测试使用，不执行具体逻辑。取值范围：<br><li>true：跳过执行逻辑</li><li>false：执行逻辑<br><br>默认取值：false。</li>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 是否续费弹性数据盘。取值范围：<br><li>true：表示续费包年包月实例同时续费其挂载的弹性数据盘</li><li>false：表示续费包年包月实例同时不再续费其挂载的弹性数据盘</li><br>默认取值：true。
	RenewPortableDataDisk *bool `json:"RenewPortableDataDisk,omitnil,omitempty" name:"RenewPortableDataDisk"`
}

func (r *InquiryPriceRenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	delete(f, "DryRun")
	delete(f, "RenewPortableDataDisk")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewInstancesResponseParams struct {
	// 该参数表示对应配置实例的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewInstancesResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstanceRequestParams struct {
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定有效的[镜像](/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作，若不指定则默认系统盘大小保持不变。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526)和[Linux](https://cloud.tencent.com/document/product/213/17525)启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

type InquiryPriceResetInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定有效的[镜像](/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作，若不指定则默认系统盘大小保持不变。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526)和[Linux](https://cloud.tencent.com/document/product/213/17525)启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

func (r *InquiryPriceResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ImageId")
	delete(f, "SystemDisk")
	delete(f, "LoginSettings")
	delete(f, "EnhancedService")
	delete(f, "UserData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResetInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstanceResponseParams struct {
	// 该参数表示重装成对应配置实例的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResetInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstancesInternetMaxBandwidthRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。当调整 `BANDWIDTH_PREPAID` 和 `BANDWIDTH_POSTPAID_BY_HOUR` 计费方式的带宽时，只支持一个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持`InternetMaxBandwidthOut`参数。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 带宽生效的起始时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 带宽生效的终止时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`ExpiredTime`获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type InquiryPriceResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。当调整 `BANDWIDTH_PREPAID` 和 `BANDWIDTH_POSTPAID_BY_HOUR` 计费方式的带宽时，只支持一个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持`InternetMaxBandwidthOut`参数。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 带宽生效的起始时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 带宽生效的终止时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`ExpiredTime`获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InternetAccessible")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResetInstancesInternetMaxBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstancesInternetMaxBandwidthResponseParams struct {
	// 该参数表示带宽调整为对应大小之后的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResetInstancesInternetMaxBandwidthResponseParams `json:"Response"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstancesTypeRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。本接口每次请求批量实例的上限为1。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可参见附表[实例资源规格](https://cloud.tencent.com/document/product/213/11518)对照表，也可以调用查询[实例资源规格列表](https://cloud.tencent.com/document/product/213/15749)接口获得最新的规格表。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type InquiryPriceResetInstancesTypeRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。本接口每次请求批量实例的上限为1。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可参见附表[实例资源规格](https://cloud.tencent.com/document/product/213/11518)对照表，也可以调用查询[实例资源规格列表](https://cloud.tencent.com/document/product/213/15749)接口获得最新的规格表。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

func (r *InquiryPriceResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstancesTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResetInstancesTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstancesTypeResponseParams struct {
	// 该参数表示调整成对应机型实例的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResetInstancesTypeResponseParams `json:"Response"`
}

func (r *InquiryPriceResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstancesTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResizeInstanceDisksRequestParams struct {
	// 待操作的实例ID。可通过[ DescribeInstances ](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待扩容的数据盘配置信息。只支持扩容非弹性数据盘[ DescribeDisks ](https://cloud.tencent.com/document/api/362/16315)接口返回值中的`Portable`为`false`表示非弹性）。数据盘容量单位：GB。最小扩容步长：10G。关于数据盘类型的选择请参考硬盘产品简介。可选数据盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因数据盘类型的不同而有所差异。
	// <dx-alert infotype="explain" title="">您必须指定参数DataDisks与SystemDisk的其中一个，但不能同时指定。</dx-alert>
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>true：表示在正常关机失败后进行强制关机</li><br><li>false：表示在正常关机失败后不进行强制关机</li><br>默认取值：false。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`

	// 待扩容的系统盘配置信息。只支持扩容云盘。
	// <dx-alert infotype="explain" title="">您必须指定参数DataDisks与SystemDisk的其中一个，但不能同时指定。</dx-alert>
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 是否在线扩容
	ResizeOnline *bool `json:"ResizeOnline,omitnil,omitempty" name:"ResizeOnline"`
}

type InquiryPriceResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的实例ID。可通过[ DescribeInstances ](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待扩容的数据盘配置信息。只支持扩容非弹性数据盘[ DescribeDisks ](https://cloud.tencent.com/document/api/362/16315)接口返回值中的`Portable`为`false`表示非弹性）。数据盘容量单位：GB。最小扩容步长：10G。关于数据盘类型的选择请参考硬盘产品简介。可选数据盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因数据盘类型的不同而有所差异。
	// <dx-alert infotype="explain" title="">您必须指定参数DataDisks与SystemDisk的其中一个，但不能同时指定。</dx-alert>
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>true：表示在正常关机失败后进行强制关机</li><br><li>false：表示在正常关机失败后不进行强制关机</li><br>默认取值：false。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`

	// 待扩容的系统盘配置信息。只支持扩容云盘。
	// <dx-alert infotype="explain" title="">您必须指定参数DataDisks与SystemDisk的其中一个，但不能同时指定。</dx-alert>
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 是否在线扩容
	ResizeOnline *bool `json:"ResizeOnline,omitnil,omitempty" name:"ResizeOnline"`
}

func (r *InquiryPriceResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResizeInstanceDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DataDisks")
	delete(f, "ForceStop")
	delete(f, "SystemDisk")
	delete(f, "ResizeOnline")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResizeInstanceDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResizeInstanceDisksResponseParams struct {
	// 该参数表示磁盘扩容成对应配置的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResizeInstanceDisksResponseParams `json:"Response"`
}

func (r *InquiryPriceResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResizeInstanceDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRunInstancesRequestParams struct {
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	//  <b>注：如果您不指定LaunchTemplate参数，则Placement为必选参数。若同时传递Placement和LaunchTemplate，则默认覆盖LaunchTemplate中对应的Placement的值。</b>
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	//  <b>注：如果您不指定LaunchTemplate参数，则ImageId为必选参数。若同时传递ImageId和LaunchTemplate，则默认覆盖LaunchTemplate中对应的ImageId的值。</b>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月</li><br><li>POSTPAID_BY_HOUR：按小时后付费</li><br><li>SPOTPAID：竞价付费</li><br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID_BY_HOUR的实例创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则系统将根据当前地域的资源售卖情况动态指定默认机型。</li><br><li>对于付费模式为CDHPAID的实例创建，该参数以"CDH_"为前缀，根据CPU和内存配置生成，具体形式为：CDH_XCXG，例如对于创建CPU为1核，内存为1G大小的专用宿主机的实例，该参数应该为CDH_1C1G。</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络IP，那么InstanceCount参数只能为1。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 购买实例数量。取值范围：[1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见[CVM实例购买限制](https://cloud.tencent.com/document/product/213/2664)。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。</li><li>最多支持128个字符（包含模式串）。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的 `SecurityGroupId` 字段来获取。若不指定该参数，则绑定指定项目下的默认安全组，如默认安全组不存在则将自动创建。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><br><li>Windows 实例：主机名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><br><li>其他类型（Linux 等）实例：主机名字符长度为[2, 30]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云服务器实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例的市场相关选项，如竞价实例相关参数
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 自定义metadata，支持创建 CVM 时添加自定义元数据键值对。
	// 
	// **注：内测中**。
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 高性能计算集群ID。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 描述了实例CPU拓扑结构的相关信息。若不指定该参数，则按系统资源情况决定。
	CpuTopology *CpuTopology `json:"CpuTopology,omitnil,omitempty" name:"CpuTopology"`

	// 实例启动模板。
	LaunchTemplate *LaunchTemplate `json:"LaunchTemplate,omitnil,omitempty" name:"LaunchTemplate"`
}

type InquiryPriceRunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	//  <b>注：如果您不指定LaunchTemplate参数，则Placement为必选参数。若同时传递Placement和LaunchTemplate，则默认覆盖LaunchTemplate中对应的Placement的值。</b>
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	//  <b>注：如果您不指定LaunchTemplate参数，则ImageId为必选参数。若同时传递ImageId和LaunchTemplate，则默认覆盖LaunchTemplate中对应的ImageId的值。</b>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月</li><br><li>POSTPAID_BY_HOUR：按小时后付费</li><br><li>SPOTPAID：竞价付费</li><br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID_BY_HOUR的实例创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则系统将根据当前地域的资源售卖情况动态指定默认机型。</li><br><li>对于付费模式为CDHPAID的实例创建，该参数以"CDH_"为前缀，根据CPU和内存配置生成，具体形式为：CDH_XCXG，例如对于创建CPU为1核，内存为1G大小的专用宿主机的实例，该参数应该为CDH_1C1G。</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络IP，那么InstanceCount参数只能为1。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 购买实例数量。取值范围：[1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见[CVM实例购买限制](https://cloud.tencent.com/document/product/213/2664)。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。</li><li>最多支持128个字符（包含模式串）。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的 `SecurityGroupId` 字段来获取。若不指定该参数，则绑定指定项目下的默认安全组，如默认安全组不存在则将自动创建。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><br><li>Windows 实例：主机名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><br><li>其他类型（Linux 等）实例：主机名字符长度为[2, 30]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云服务器实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例的市场相关选项，如竞价实例相关参数
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 自定义metadata，支持创建 CVM 时添加自定义元数据键值对。
	// 
	// **注：内测中**。
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 高性能计算集群ID。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 描述了实例CPU拓扑结构的相关信息。若不指定该参数，则按系统资源情况决定。
	CpuTopology *CpuTopology `json:"CpuTopology,omitnil,omitempty" name:"CpuTopology"`

	// 实例启动模板。
	LaunchTemplate *LaunchTemplate `json:"LaunchTemplate,omitnil,omitempty" name:"LaunchTemplate"`
}

func (r *InquiryPriceRunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRunInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "ImageId")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "InstanceType")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "VirtualPrivateCloud")
	delete(f, "InternetAccessible")
	delete(f, "InstanceCount")
	delete(f, "InstanceName")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "EnhancedService")
	delete(f, "ClientToken")
	delete(f, "HostName")
	delete(f, "TagSpecification")
	delete(f, "InstanceMarketOptions")
	delete(f, "Metadata")
	delete(f, "HpcClusterId")
	delete(f, "CpuTopology")
	delete(f, "LaunchTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRunInstancesResponseParams struct {
	// 该参数表示对应配置实例的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRunInstancesResponseParams `json:"Response"`
}

func (r *InquiryPriceRunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceTerminateInstancesRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type InquiryPriceTerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *InquiryPriceTerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceTerminateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceTerminateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceTerminateInstancesResponseParams struct {
	// 退款详情。
	InstanceRefundsSet []*InstanceRefund `json:"InstanceRefundsSet,omitnil,omitempty" name:"InstanceRefundsSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceTerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceTerminateInstancesResponseParams `json:"Response"`
}

func (r *InquiryPriceTerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceTerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {
	// 实例所在的位置。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 实例`ID`。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例的CPU核数，单位：核。
	CPU *int64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 实例内存容量，单位：`GiB`。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例业务状态。取值范围：<br><li>NORMAL：表示正常状态的实例</li><li>EXPIRED：表示过期的实例</li><li>PROTECTIVELY_ISOLATED：表示被安全隔离的实例。</li>
	RestrictState *string `json:"RestrictState,omitnil,omitempty" name:"RestrictState"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。取值范围：<br><li>PREPAID：预付费，即包年包月</li><br><li>POSTPAID_BY_HOUR：按小时后付费</li><br><li>CDHPAID：独享子机（基于专用宿主机创建，宿主机部分的资源不收费）</li><br><li>SPOTPAID：竞价付费</li><br><li>CDCPAID：专用集群付费</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘信息。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 实例主网卡的内网`IP`列表。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 实例主网卡的公网`IP`列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 实例带宽信息。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 实例所属虚拟私有网络信息。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 生产实例所使用的镜像`ID`。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 自动续费标识。取值范围：<br><li>`NOTIFY_AND_MANUAL_RENEW`：表示通知即将过期，但不自动续费</li><li>`NOTIFY_AND_AUTO_RENEW`：表示通知即将过期，而且自动续费</li><li>`DISABLE_NOTIFY_AND_MANUAL_RENEW`：表示不通知即将过期，也不自动续费。
	// </li><li>注意：后付费模式本项为null</li>
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 到期时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。注意：后付费模式本项为null
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 操作系统名称。
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例登录设置。目前只返回实例所关联的密钥。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例状态。状态类型详见[实例状态表](https://cloud.tencent.com/document/api/213/15753#InstanceStatus)
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 实例关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例的关机计费模式。
	// 取值范围：<br><li>KEEP_CHARGING：关机继续收费</li><li>STOP_CHARGING：关机停止收费</li><li>NOT_APPLICABLE：实例处于非关机状态或者不适用关机停止计费的条件</li>
	StopChargingMode *string `json:"StopChargingMode,omitnil,omitempty" name:"StopChargingMode"`

	// 实例全局唯一ID
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 实例的最新操作。例：StopInstances、ResetInstance。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperation *string `json:"LatestOperation,omitnil,omitempty" name:"LatestOperation"`

	// 实例的最新操作状态。取值范围：<br><li>SUCCESS：表示操作成功</li><li>OPERATING：表示操作执行中</li><li>FAILED：表示操作失败</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationState *string `json:"LatestOperationState,omitnil,omitempty" name:"LatestOperationState"`

	// 实例最新操作的唯一请求 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitnil,omitempty" name:"LatestOperationRequestId"`

	// 分散置放群组ID。
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitnil,omitempty" name:"DisasterRecoverGroupId"`

	// 实例的IPv6地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6Addresses []*string `json:"IPv6Addresses,omitnil,omitempty" name:"IPv6Addresses"`

	// CAM角色名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 高性能计算集群`ID`。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 高性能计算集群`IP`列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RdmaIpAddresses []*string `json:"RdmaIpAddresses,omitnil,omitempty" name:"RdmaIpAddresses"`

	// 实例所在的专用集群`ID`。
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 实例隔离类型。取值范围：<br><li>ARREAR：表示欠费隔离<br></li><li>EXPIRE：表示到期隔离<br></li><li>MANMADE：表示主动退还隔离<br></li><li>NOTISOLATED：表示未隔离<br></li>
	IsolatedSource *string `json:"IsolatedSource,omitnil,omitempty" name:"IsolatedSource"`

	// GPU信息。如果是gpu类型子机，该值会返回GPU信息，如果是其他类型子机则不返回。
	GPUInfo *GPUInfo `json:"GPUInfo,omitnil,omitempty" name:"GPUInfo"`

	// 实例的操作系统许可类型，默认为TencentCloud
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>true：表示开启实例保护，不允许通过api接口删除实例</li><li>false：表示关闭实例保护，允许通过api接口删除实例</li><br>默认取值：false。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`

	// 默认登录用户。
	DefaultLoginUser *string `json:"DefaultLoginUser,omitnil,omitempty" name:"DefaultLoginUser"`

	// 默认登录端口。
	DefaultLoginPort *int64 `json:"DefaultLoginPort,omitnil,omitempty" name:"DefaultLoginPort"`

	// 实例的最新操作错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationErrorMsg *string `json:"LatestOperationErrorMsg,omitnil,omitempty" name:"LatestOperationErrorMsg"`

	// 实例绑定的公网IPv6地址。
	PublicIPv6Addresses []*string `json:"PublicIPv6Addresses,omitnil,omitempty" name:"PublicIPv6Addresses"`
}

type InstanceAttribute struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例属性信息。
	Attributes *Attribute `json:"Attributes,omitnil,omitempty" name:"Attributes"`
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li><br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li><br><br>默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceFamilyConfig struct {
	// 机型族名称的中文全称。
	InstanceFamilyName *string `json:"InstanceFamilyName,omitnil,omitempty" name:"InstanceFamilyName"`

	// 机型族名称的英文简称。
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`
}

type InstanceMarketOptionsRequest struct {
	// 竞价相关选项
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitnil,omitempty" name:"SpotOptions"`

	// 市场选项类型，当前只支持取值：spot
	MarketType *string `json:"MarketType,omitnil,omitempty" name:"MarketType"`
}

type InstanceRefund struct {
	// 实例Id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 退款数额。
	Refunds *float64 `json:"Refunds,omitnil,omitempty" name:"Refunds"`

	// 退款详情。
	PriceDetail *string `json:"PriceDetail,omitnil,omitempty" name:"PriceDetail"`
}

type InstanceStatus struct {
	// 实例`ID`。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例状态。取值范围：<br><li>PENDING：表示创建中<br></li><li>LAUNCH_FAILED：表示创建失败<br></li><li>RUNNING：表示运行中<br></li><li>STOPPED：表示关机<br></li><li>STARTING：表示开机中<br></li><li>STOPPING：表示关机中<br></li><li>REBOOTING：表示重启中<br></li><li>SHUTDOWN：表示停止待销毁<br></li><li>TERMINATING：表示销毁中<br></li><li>ENTER_RESCUE_MODE：表示进入救援模式<br></li><li>RESCUE_MODE：表示在救援模式中<br></li><li>EXIT_RESCUE_MODE：表示退出救援模式<br></li><li>ENTER_SERVICE_LIVE_MIGRATE：表示进入在线服务迁移<br></li><li>SERVICE_LIVE_MIGRATE：表示在线服务迁移中<br></li><li>EXIT_SERVICE_LIVE_MIGRATE：表示退出在线服务迁移。<br></li>
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`
}

type InstanceTypeConfig struct {
	// 可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例机型系列。
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// GPU核数，单位：核。
	GPU *int64 `json:"GPU,omitnil,omitempty" name:"GPU"`

	// CPU核数，单位：核。
	CPU *int64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 内存容量，单位：`GiB`。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// FPGA核数，单位：核。
	FPGA *int64 `json:"FPGA,omitnil,omitempty" name:"FPGA"`

	// 实例机型映射的物理GPU卡数，单位：卡。vGPU卡型小于1，直通卡型大于等于1。vGPU是通过分片虚拟化技术，将物理GPU卡重新划分，同一块GPU卡经虚拟化分割后可分配至不同的实例使用。直通卡型会将GPU设备直接挂载给实例使用。
	GpuCount *float64 `json:"GpuCount,omitnil,omitempty" name:"GpuCount"`
}

type InstanceTypeConfigStatus struct {
	// 状态描述
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 配置信息
	InstanceTypeConfig *InstanceTypeConfig `json:"InstanceTypeConfig,omitnil,omitempty" name:"InstanceTypeConfig"`
}

type InstanceTypeQuotaItem struct {
	// 可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例计费模式。取值范围： <br><li>PREPAID：表示预付费，即包年包月<br></li><li>POSTPAID_BY_HOUR：表示后付费，即按量计费</li><li>CDHPAID：表示[专用宿主机](https://cloud.tencent.com/document/product/416)付费，即只对`专用宿主机`计费，不对`专用宿主机`上的实例计费。<br></li><li>SPOTPAID：表示竞价实例付费。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 网卡类型，例如：25代表25G网卡
	NetworkCard *int64 `json:"NetworkCard,omitnil,omitempty" name:"NetworkCard"`

	// 扩展属性。
	Externals *Externals `json:"Externals,omitnil,omitempty" name:"Externals"`

	// 实例的CPU核数，单位：核。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存容量，单位：`GB`。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例机型系列。
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// 机型名称。
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 本地磁盘规格列表。当该参数返回为空值时，表示当前情况下无法创建本地盘。
	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitnil,omitempty" name:"LocalDiskTypeList"`

	// 实例是否售卖。取值范围： <br><li>SELL：表示实例可购买<br></li><li>SOLD_OUT：表示实例已售罄。</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例的售卖价格。
	Price *ItemPrice `json:"Price,omitnil,omitempty" name:"Price"`

	// 售罄原因。
	SoldOutReason *string `json:"SoldOutReason,omitnil,omitempty" name:"SoldOutReason"`

	// 内网带宽，单位Gbps。
	InstanceBandwidth *float64 `json:"InstanceBandwidth,omitnil,omitempty" name:"InstanceBandwidth"`

	// 网络收发包能力，单位万PPS。
	InstancePps *int64 `json:"InstancePps,omitnil,omitempty" name:"InstancePps"`

	// 本地存储块数量。
	StorageBlockAmount *int64 `json:"StorageBlockAmount,omitnil,omitempty" name:"StorageBlockAmount"`

	// 处理器型号。
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// 实例的GPU数量。
	Gpu *int64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// 实例的FPGA数量。
	Fpga *int64 `json:"Fpga,omitnil,omitempty" name:"Fpga"`

	// 实例备注信息。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例机型映射的物理GPU卡数，单位：卡。vGPU卡型小于1，直通卡型大于等于1。vGPU是通过分片虚拟化技术，将物理GPU卡重新划分，同一块GPU卡经虚拟化分割后可分配至不同的实例使用。直通卡型会将GPU设备直接挂载给实例使用。
	GpuCount *float64 `json:"GpuCount,omitnil,omitempty" name:"GpuCount"`

	// 实例的CPU主频信息
	Frequency *string `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// 描述库存情况。取值范围：
	// <li> EnoughStock：表示对应库存非常充足</li> <li>NormalStock：表示对应库存供应有保障</li><li> UnderStock：表示对应库存即将售罄</li> <li>WithoutStock：表示对应库存已经售罄</li>
	StatusCategory *string `json:"StatusCategory,omitnil,omitempty" name:"StatusCategory"`
}

type InternetAccessible struct {
	// 网络计费类型。取值范围：<br><li>BANDWIDTH_PREPAID：预付费按带宽结算</li><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费</li><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费</li><li>BANDWIDTH_PACKAGE：带宽包用户</li>默认取值：非带宽包用户默认与子机付费类型保持一致，比如子机付费类型为预付费，网络计费类型默认为预付费；子机付费类型为后付费，网络计费类型默认为后付费。
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见[购买网络带宽](https://cloud.tencent.com/document/product/213/12523)。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否分配公网IP。取值范围：<br><li>true：表示分配公网IP</li><li>false：表示不分配公网IP</li><br>当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。该参数仅在RunInstances接口中作为入参使用。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitnil,omitempty" name:"PublicIpAssigned"`

	// 带宽包ID。可通过[ DescribeBandwidthPackages ](https://cloud.tencent.com/document/api/215/19209)接口返回值中的`BandwidthPackageId`获取。该参数仅在RunInstances接口中作为入参使用。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// 线路类型。各种线路类型详情可参考：[EIP 的 IP 地址类型](https://cloud.tencent.com/document/product/1199/41646)。默认值：BGP。
	// 
	// - BGP：常规 BGP 线路
	// 
	// 已开通静态单线IP白名单的用户，可选值：
	// 
	//  - CMCC：中国移动
	//  - CTCC：中国电信
	//  - CUCC：中国联通
	// 
	// 注意：仅部分地域支持静态单线IP。
	// 示例值：BGP
	InternetServiceProvider *string `json:"InternetServiceProvider,omitnil,omitempty" name:"InternetServiceProvider"`

	// 公网 IP 类型。
	// 
	// - WanIP：普通公网IP。
	// - HighQualityEIP：精品 IP。仅新加坡和中国香港支持精品IP。
	// - AntiDDoSEIP：高防 IP。仅部分地域支持高防IP，详情可见[弹性公网IP产品概述](https://cloud.tencent.com/document/product/1199/41646)。
	// 
	// 如需为资源分配公网IPv4地址，请指定公网IPv4地址类型。
	// 
	// 示例值：WanIP
	// 
	// 此功能仅部分地区灰度开放，如需使用[请提交工单咨询](https://console.cloud.tencent.com/workorder/category)
	IPv4AddressType *string `json:"IPv4AddressType,omitnil,omitempty" name:"IPv4AddressType"`

	// 弹性公网 IPv6 类型。
	// - EIPv6：弹性公网 IPv6。
	// - HighQualityEIPv6：精品 IPv6。仅中国香港支持精品IPv6。
	// 
	// 如需为资源分配IPv6地址，请指定弹性公网IPv6类型。
	// 示例值：EIPv6
	// 
	// 此功能仅部分地区灰度开放，如需使用[请提交工单咨询](https://console.cloud.tencent.com/workorder/category)
	IPv6AddressType *string `json:"IPv6AddressType,omitnil,omitempty" name:"IPv6AddressType"`

	// 高防包唯一ID，申请高防IP时，该字段必传。
	// 示例值：bgp-12345678
	AntiDDoSPackageId *string `json:"AntiDDoSPackageId,omitnil,omitempty" name:"AntiDDoSPackageId"`
}

type InternetBandwidthConfig struct {
	// 开始时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例带宽信息。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`
}

type InternetChargeTypeConfig struct {
	// 网络计费类型。取值范围：<br><li>BANDWIDTH_PREPAID：预付费按带宽结算</li><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费</li><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费</li><li>BANDWIDTH_PACKAGE：带宽包用户</li>默认取值：非带宽包用户默认与子机付费类型保持一致，比如子机付费类型为预付费，网络计费类型默认为预付费；子机付费类型为后付费，网络计费类型默认为后付费。
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// 网络计费模式描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ItemPrice struct {
	// 后续合计费用的原价，后付费模式使用，单位：元。<br><li>如返回了其他时间区间项，如UnitPriceSecondStep，则本项代表时间区间在(0, 96)小时；若未返回其他时间区间项，则本项代表全时段，即(0, ∞)小时</li>
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 后续计价单元，后付费模式使用，可取值范围： <br><li>HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）：</li><li>GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。</li>
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// 预支合计费用的原价，预付费模式使用，单位：元。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 预支合计费用的折扣价，预付费模式使用，单位：元。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 折扣，如20.0代表2折。
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 后续合计费用的折扣价，后付费模式使用，单位：元<br><li>如返回了其他时间区间项，如UnitPriceDiscountSecondStep，则本项代表时间区间在(0, 96)小时；若未返回其他时间区间项，则本项代表全时段，即(0, ∞)小时</li>
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 使用时间区间在(96, 360)小时的后续合计费用的原价，后付费模式使用，单位：元。
	UnitPriceSecondStep *float64 `json:"UnitPriceSecondStep,omitnil,omitempty" name:"UnitPriceSecondStep"`

	// 使用时间区间在(96, 360)小时的后续合计费用的折扣价，后付费模式使用，单位：元
	UnitPriceDiscountSecondStep *float64 `json:"UnitPriceDiscountSecondStep,omitnil,omitempty" name:"UnitPriceDiscountSecondStep"`

	// 使用时间区间在(360, ∞)小时的后续合计费用的原价，后付费模式使用，单位：元。
	UnitPriceThirdStep *float64 `json:"UnitPriceThirdStep,omitnil,omitempty" name:"UnitPriceThirdStep"`

	// 使用时间区间在(360, ∞)小时的后续合计费用的折扣价，后付费模式使用，单位：元
	UnitPriceDiscountThirdStep *float64 `json:"UnitPriceDiscountThirdStep,omitnil,omitempty" name:"UnitPriceDiscountThirdStep"`

	// 预支三年合计费用的原价，预付费模式使用，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPriceThreeYear *float64 `json:"OriginalPriceThreeYear,omitnil,omitempty" name:"OriginalPriceThreeYear"`

	// 预支三年合计费用的折扣价，预付费模式使用，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPriceThreeYear *float64 `json:"DiscountPriceThreeYear,omitnil,omitempty" name:"DiscountPriceThreeYear"`

	// 预支三年应用的折扣，如20.0代表2折。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountThreeYear *float64 `json:"DiscountThreeYear,omitnil,omitempty" name:"DiscountThreeYear"`

	// 预支五年合计费用的原价，预付费模式使用，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPriceFiveYear *float64 `json:"OriginalPriceFiveYear,omitnil,omitempty" name:"OriginalPriceFiveYear"`

	// 预支五年合计费用的折扣价，预付费模式使用，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPriceFiveYear *float64 `json:"DiscountPriceFiveYear,omitnil,omitempty" name:"DiscountPriceFiveYear"`

	// 预支五年应用的折扣，如20.0代表2折。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountFiveYear *float64 `json:"DiscountFiveYear,omitnil,omitempty" name:"DiscountFiveYear"`

	// 预支一年合计费用的原价，预付费模式使用，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPriceOneYear *float64 `json:"OriginalPriceOneYear,omitnil,omitempty" name:"OriginalPriceOneYear"`

	// 预支一年合计费用的折扣价，预付费模式使用，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPriceOneYear *float64 `json:"DiscountPriceOneYear,omitnil,omitempty" name:"DiscountPriceOneYear"`

	// 预支一年应用的折扣，如20.0代表2折。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountOneYear *float64 `json:"DiscountOneYear,omitnil,omitempty" name:"DiscountOneYear"`
}

type KeyPair struct {
	// 密钥对的`ID`，是密钥对的唯一标识。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 密钥对名称。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 密钥对所属的项目 `ID`，ProjectId 为 0 时表示默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 密钥对描述信息。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 密钥对的纯文本公钥。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 密钥对的纯文本私钥。腾讯云不会保管私钥，请用户自行妥善保存。
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// 密钥关联的实例`ID`列表。
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitnil,omitempty" name:"AssociatedInstanceIds"`

	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 密钥关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type LaunchTemplate struct {
	// 实例启动模板ID，通过该参数可使用实例模板中的预设参数创建实例。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 实例启动模板版本号，若给定，新实例启动模板将基于给定的版本号创建
	LaunchTemplateVersion *uint64 `json:"LaunchTemplateVersion,omitnil,omitempty" name:"LaunchTemplateVersion"`
}

type LaunchTemplateInfo struct {
	// 实例启动模版本号。
	LatestVersionNumber *uint64 `json:"LatestVersionNumber,omitnil,omitempty" name:"LatestVersionNumber"`

	// 实例启动模板ID。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 实例启动模板名。
	LaunchTemplateName *string `json:"LaunchTemplateName,omitnil,omitempty" name:"LaunchTemplateName"`

	// 实例启动模板默认版本号。
	DefaultVersionNumber *uint64 `json:"DefaultVersionNumber,omitnil,omitempty" name:"DefaultVersionNumber"`

	// 实例启动模板包含的版本总数量。
	LaunchTemplateVersionCount *uint64 `json:"LaunchTemplateVersionCount,omitnil,omitempty" name:"LaunchTemplateVersionCount"`

	// 创建该模板的用户UIN。
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// 创建该模板的时间。
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`
}

type LaunchTemplateVersionData struct {
	// 实例所在的位置。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <li>`PREPAID`：表示预付费，即包年包月</li>
	// <li>`POSTPAID_BY_HOUR`：表示后付费，即按量计费</li>
	// <li>`CDHPAID`：`专用宿主机`付费，即只对`专用宿主机`计费，不对`专用宿主机`上的实例计费。</li>
	// <li>`SPOTPAID`：表示竞价实例付费。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘信息。只包含随实例购买的数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 实例带宽信息。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 实例所属虚拟私有网络信息。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 生产实例所使用的镜像`ID`。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例登录设置。目前只返回实例所关联的密钥。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// CAM角色名。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 高性能计算集群`ID`。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 购买实例数量。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 增强服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 置放群组ID，仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer,omitnil,omitempty" name:"ActionTimer"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 云服务器的主机名。
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 用于保证请求幂等性的字符串。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 预付费模式，即包年包月相关参数设置。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的云服务器、云硬盘实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：
	// 
	// true：表示开启实例保护，不允许通过api接口删除实例
	// false：表示关闭实例保护，允许通过api接口删除实例
	// 
	// 默认取值：false。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`
}

type LaunchTemplateVersionInfo struct {
	// 实例启动模板版本号。
	LaunchTemplateVersion *uint64 `json:"LaunchTemplateVersion,omitnil,omitempty" name:"LaunchTemplateVersion"`

	// 实例启动模板版本数据详情。
	LaunchTemplateVersionData *LaunchTemplateVersionData `json:"LaunchTemplateVersionData,omitnil,omitempty" name:"LaunchTemplateVersionData"`

	// 实例启动模板版本创建时间。
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 实例启动模板ID。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 是否为默认启动模板版本。
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitnil,omitempty" name:"IsDefaultVersion"`

	// 实例启动模板版本描述信息。
	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitnil,omitempty" name:"LaunchTemplateVersionDescription"`

	// 创建者的AppId。
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`
}

type LocalDiskType struct {
	// 本地磁盘类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 本地磁盘属性。
	PartitionType *string `json:"PartitionType,omitnil,omitempty" name:"PartitionType"`

	// 本地磁盘最小值。
	MinSize *int64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 本地磁盘最大值。
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 购买时本地盘是否为必选。取值范围：<br><li>REQUIRED：表示必选</li><li>OPTIONAL：表示可选。</li>
	Required *string `json:"Required,omitnil,omitempty" name:"Required"`
}

type LoginSettings struct {
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<li>Linux实例密码必须8到30位，至少包括两项[a-z]，[A-Z]、[0-9] 和 [( ) \` ~ ! @ # $ % ^ & *  - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。</li><li>Windows实例密码必须12到30位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) \` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? /]中的特殊符号。</li>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口[DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699)获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为true。取值范围：<li>true：表示保持镜像的登录设置</li><li>false：表示不保持镜像的登录设置</li>默认取值：false。
	KeepImageLogin *string `json:"KeepImageLogin,omitnil,omitempty" name:"KeepImageLogin"`
}

type Metadata struct {
	// 自定义metadata键值对列表。
	Items []*MetadataItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type MetadataItem struct {
	// 自定义metadata键，需符合正则 ^[a-zA-Z0-9_-]+$，长度 ≤128 字节（大小写敏感）；
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 自定义metadata值，支持任意数据（含二进制），大小 ≤256 KB（大小写敏感）；
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ModifyChcAttributeRequestParams struct {
	// CHC物理服务器ID。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`

	// CHC物理服务器名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 服务器类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 合法字符为字母,数字, 横线和下划线
	BmcUser *string `json:"BmcUser,omitnil,omitempty" name:"BmcUser"`

	// 密码8-16位字符, 允许数字，字母， 和特殊字符()`~!@#$%^&*-+=_|{}[]:;'<>,.?/
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// bmc网络的安全组列表
	BmcSecurityGroupIds []*string `json:"BmcSecurityGroupIds,omitnil,omitempty" name:"BmcSecurityGroupIds"`
}

type ModifyChcAttributeRequest struct {
	*tchttp.BaseRequest
	
	// CHC物理服务器ID。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`

	// CHC物理服务器名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 服务器类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 合法字符为字母,数字, 横线和下划线
	BmcUser *string `json:"BmcUser,omitnil,omitempty" name:"BmcUser"`

	// 密码8-16位字符, 允许数字，字母， 和特殊字符()`~!@#$%^&*-+=_|{}[]:;'<>,.?/
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// bmc网络的安全组列表
	BmcSecurityGroupIds []*string `json:"BmcSecurityGroupIds,omitnil,omitempty" name:"BmcSecurityGroupIds"`
}

func (r *ModifyChcAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChcAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	delete(f, "InstanceName")
	delete(f, "DeviceType")
	delete(f, "BmcUser")
	delete(f, "Password")
	delete(f, "BmcSecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyChcAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyChcAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyChcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyChcAttributeResponseParams `json:"Response"`
}

func (r *ModifyChcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChcAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisasterRecoverGroupAttributeRequestParams struct {
	// 分散置放群组ID，可使用[DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/api/213/17810)接口获取。
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitnil,omitempty" name:"DisasterRecoverGroupId"`

	// 分散置放群组名称，长度1-60个字符，支持中、英文。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ModifyDisasterRecoverGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 分散置放群组ID，可使用[DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/api/213/17810)接口获取。
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitnil,omitempty" name:"DisasterRecoverGroupId"`

	// 分散置放群组名称，长度1-60个字符，支持中、英文。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *ModifyDisasterRecoverGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisasterRecoverGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisasterRecoverGroupId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDisasterRecoverGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisasterRecoverGroupAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDisasterRecoverGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDisasterRecoverGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyDisasterRecoverGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisasterRecoverGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsAttributeRequestParams struct {
	// 一个或多个待操作的CDH实例ID。
	HostIds []*string `json:"HostIds,omitnil,omitempty" name:"HostIds"`

	// CDH实例显示名称。可任意命名，但不得超过60个字符。
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li><br>若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 项目ID。项目可以使用[AddProject](https://cloud.tencent.com/document/api/651/81952)接口创建。可通过[DescribeProject](https://cloud.tencent.com/document/product/378/4400) API返回值中的`projectId`获取。后续使用[DescribeHosts](https://cloud.tencent.com/document/api/213/16474)接口查询实例时，项目ID可用于过滤结果。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyHostsAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的CDH实例ID。
	HostIds []*string `json:"HostIds,omitnil,omitempty" name:"HostIds"`

	// CDH实例显示名称。可任意命名，但不得超过60个字符。
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li><br>若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 项目ID。项目可以使用[AddProject](https://cloud.tencent.com/document/api/651/81952)接口创建。可通过[DescribeProject](https://cloud.tencent.com/document/product/378/4400) API返回值中的`projectId`获取。后续使用[DescribeHosts](https://cloud.tencent.com/document/api/213/16474)接口查询实例时，项目ID可用于过滤结果。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ModifyHostsAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostsAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostIds")
	delete(f, "HostName")
	delete(f, "RenewFlag")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostsAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyHostsAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostsAttributeResponseParams `json:"Response"`
}

func (r *ModifyHostsAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostsAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHpcClusterAttributeRequestParams struct {
	// 高性能计算集群ID。集群ID可通过 [查询高性能集群信息](https://cloud.tencent.com/document/api/213/83220) 接口获取。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 高性能计算集群新名称，长度限制[1-60]。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 高性能计算集群新备注，长度[1-256]。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyHpcClusterAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 高性能计算集群ID。集群ID可通过 [查询高性能集群信息](https://cloud.tencent.com/document/api/213/83220) 接口获取。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 高性能计算集群新名称，长度限制[1-60]。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 高性能计算集群新备注，长度[1-256]。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyHpcClusterAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHpcClusterAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HpcClusterId")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHpcClusterAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHpcClusterAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyHpcClusterAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHpcClusterAttributeResponseParams `json:"Response"`
}

func (r *ModifyHpcClusterAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHpcClusterAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageAttributeRequestParams struct {
	// 镜像ID，形如`img-gvbnzy6f`。镜像ID可以通过如下方式获取：<li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取。</li><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 设置新的镜像名称；必须满足下列限制 <li> 不得超过60个字符。</li><li> 镜像名称不能与已有镜像重复。</li>
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 设置新的镜像描述；必须满足下列限制： <li> 不得超过 256 个字符。</li>
	ImageDescription *string `json:"ImageDescription,omitnil,omitempty" name:"ImageDescription"`

	// 设置镜像族；
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`

	// 设置镜像是否废弃；
	ImageDeprecated *bool `json:"ImageDeprecated,omitnil,omitempty" name:"ImageDeprecated"`
}

type ModifyImageAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID，形如`img-gvbnzy6f`。镜像ID可以通过如下方式获取：<li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取。</li><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 设置新的镜像名称；必须满足下列限制 <li> 不得超过60个字符。</li><li> 镜像名称不能与已有镜像重复。</li>
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 设置新的镜像描述；必须满足下列限制： <li> 不得超过 256 个字符。</li>
	ImageDescription *string `json:"ImageDescription,omitnil,omitempty" name:"ImageDescription"`

	// 设置镜像族；
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`

	// 设置镜像是否废弃；
	ImageDeprecated *bool `json:"ImageDeprecated,omitnil,omitempty" name:"ImageDeprecated"`
}

func (r *ModifyImageAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	delete(f, "ImageName")
	delete(f, "ImageDescription")
	delete(f, "ImageFamily")
	delete(f, "ImageDeprecated")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyImageAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImageAttributeResponseParams `json:"Response"`
}

func (r *ModifyImageAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSharePermissionRequestParams struct {
	// 镜像ID，形如`img-gvbnzy6f`。镜像Id可以通过如下方式获取：<br><li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取。</li><br><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。</li> <br>镜像ID必须指定为状态为`NORMAL`的镜像。镜像状态请参考[镜像数据表](https://cloud.tencent.com/document/product/213/15753#Image)。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 接收共享镜像的主账号ID列表，array型参数的格式可以参考[API简介](/document/api/213/568)。账号ID不同于QQ号，查询用户主账号ID请查看[账号信息](https://console.cloud.tencent.com/developer)中的账号ID栏。
	AccountIds []*string `json:"AccountIds,omitnil,omitempty" name:"AccountIds"`

	// 操作，包括 `SHARE`，`CANCEL`。其中`SHARE`代表共享操作，`CANCEL`代表取消共享操作。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`
}

type ModifyImageSharePermissionRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID，形如`img-gvbnzy6f`。镜像Id可以通过如下方式获取：<br><li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取。</li><br><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。</li> <br>镜像ID必须指定为状态为`NORMAL`的镜像。镜像状态请参考[镜像数据表](https://cloud.tencent.com/document/product/213/15753#Image)。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 接收共享镜像的主账号ID列表，array型参数的格式可以参考[API简介](/document/api/213/568)。账号ID不同于QQ号，查询用户主账号ID请查看[账号信息](https://console.cloud.tencent.com/developer)中的账号ID栏。
	AccountIds []*string `json:"AccountIds,omitnil,omitempty" name:"AccountIds"`

	// 操作，包括 `SHARE`，`CANCEL`。其中`SHARE`代表共享操作，`CANCEL`代表取消共享操作。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`
}

func (r *ModifyImageSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	delete(f, "AccountIds")
	delete(f, "Permission")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSharePermissionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImageSharePermissionResponseParams `json:"Response"`
}

func (r *ModifyImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceDiskTypeRequestParams struct {
	// 待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例数据盘配置信息，只需要指定要转换的目标云硬盘的介质类型，指定DiskType的值，当前只支持一个数据盘转化。只支持CDHPAID类型实例指定CdcId参数。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 实例系统盘配置信息，只需要指定要转换的目标云硬盘的介质类型，指定DiskType的值。只支持CDHPAID类型实例指定CdcId参数。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`
}

type ModifyInstanceDiskTypeRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例数据盘配置信息，只需要指定要转换的目标云硬盘的介质类型，指定DiskType的值，当前只支持一个数据盘转化。只支持CDHPAID类型实例指定CdcId参数。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 实例系统盘配置信息，只需要指定要转换的目标云硬盘的介质类型，指定DiskType的值。只支持CDHPAID类型实例指定CdcId参数。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`
}

func (r *ModifyInstanceDiskTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceDiskTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DataDisks")
	delete(f, "SystemDisk")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceDiskTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceDiskTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceDiskTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceDiskTypeResponseParams `json:"Response"`
}

func (r *ModifyInstanceDiskTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceDiskTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesAttributeRequestParams struct {
	// 一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 修改后实例名称。可任意命名，但不得超过60个字符。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16 KB。关于获取此参数的详细介绍，请参阅 [Windows](https://cloud.tencent.com/document/product/213/17526) 和 [Linux](https://cloud.tencent.com/document/product/213/17525) 启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 指定实例的修改后的安全组Id列表，子机将重新关联指定列表的安全组，原本关联的安全组会被解绑。
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// 给实例绑定用户角色，传空值为解绑操作
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 修改后实例的主机名。<li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><li>Windows 实例：主机名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><li>其他类型（Linux 等）实例：主机名字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li>注意点：修改主机名后实例会立即重启，重启后新的主机名生效。
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<li>true：表示开启实例保护，不允许通过api接口删除实例</li><li>false：表示关闭实例保护，允许通过api接口删除实例</li>默认取值：false。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`

	// 角色类别，与CamRoleName搭配使用，该值可从CAM [ DescribeRoleList ](https://cloud.tencent.com/document/product/598/36223)或[ GetRole ](https://cloud.tencent.com/document/product/598/36221)接口返回RoleType字段获取，当前只接受user、system和service_linked三种类别。
	// 举例：一般CamRoleName中包含“LinkedRoleIn”（如TKE_QCSLinkedRoleInPrometheusService）时，DescribeRoleList和GetRole返回的RoleType为service_linked，则本参数也需要传递service_linked。
	// 该参数默认值为user，若CameRoleName为非service_linked类型，本参数可不传递。
	CamRoleType *string `json:"CamRoleType,omitnil,omitempty" name:"CamRoleType"`

	// 修改实例主机名是否自动重启实例，不传默认自动重启。
	// - true: 修改主机名，并自动重启实例；
	// - false: 修改主机名，不自动重启实例，需要手动重启使新主机名生效。
	// 注意点：本参数仅对修改主机名生效。
	AutoReboot *bool `json:"AutoReboot,omitnil,omitempty" name:"AutoReboot"`
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 修改后实例名称。可任意命名，但不得超过60个字符。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16 KB。关于获取此参数的详细介绍，请参阅 [Windows](https://cloud.tencent.com/document/product/213/17526) 和 [Linux](https://cloud.tencent.com/document/product/213/17525) 启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 指定实例的修改后的安全组Id列表，子机将重新关联指定列表的安全组，原本关联的安全组会被解绑。
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// 给实例绑定用户角色，传空值为解绑操作
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 修改后实例的主机名。<li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><li>Windows 实例：主机名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><li>其他类型（Linux 等）实例：主机名字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li>注意点：修改主机名后实例会立即重启，重启后新的主机名生效。
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<li>true：表示开启实例保护，不允许通过api接口删除实例</li><li>false：表示关闭实例保护，允许通过api接口删除实例</li>默认取值：false。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`

	// 角色类别，与CamRoleName搭配使用，该值可从CAM [ DescribeRoleList ](https://cloud.tencent.com/document/product/598/36223)或[ GetRole ](https://cloud.tencent.com/document/product/598/36221)接口返回RoleType字段获取，当前只接受user、system和service_linked三种类别。
	// 举例：一般CamRoleName中包含“LinkedRoleIn”（如TKE_QCSLinkedRoleInPrometheusService）时，DescribeRoleList和GetRole返回的RoleType为service_linked，则本参数也需要传递service_linked。
	// 该参数默认值为user，若CameRoleName为非service_linked类型，本参数可不传递。
	CamRoleType *string `json:"CamRoleType,omitnil,omitempty" name:"CamRoleType"`

	// 修改实例主机名是否自动重启实例，不传默认自动重启。
	// - true: 修改主机名，并自动重启实例；
	// - false: 修改主机名，不自动重启实例，需要手动重启使新主机名生效。
	// 注意点：本参数仅对修改主机名生效。
	AutoReboot *bool `json:"AutoReboot,omitnil,omitempty" name:"AutoReboot"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceName")
	delete(f, "UserData")
	delete(f, "SecurityGroups")
	delete(f, "CamRoleName")
	delete(f, "HostName")
	delete(f, "DisableApiTermination")
	delete(f, "CamRoleType")
	delete(f, "AutoReboot")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesAttributeResponseParams `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesChargeTypeRequestParams struct {
	// 一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为30。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 修改后实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<li> PREPAID：预付费，即包年包月。</li><li> POSTPAID_BY_HOUR：后付费，即按量付费。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 修改后预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。<dx-alert infotype="explain" title="">若指定实例的付费模式为预付费则该参数必传。</dx-alert>
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 是否同时切换弹性数据云盘计费模式。取值范围：<li> true：表示切换弹性数据云盘计费模式</li><li> false：表示不切换弹性数据云盘计费模式</li>默认取值：false。
	ModifyPortableDataDisk *bool `json:"ModifyPortableDataDisk,omitnil,omitempty" name:"ModifyPortableDataDisk"`
}

type ModifyInstancesChargeTypeRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为30。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 修改后实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<li> PREPAID：预付费，即包年包月。</li><li> POSTPAID_BY_HOUR：后付费，即按量付费。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 修改后预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。<dx-alert infotype="explain" title="">若指定实例的付费模式为预付费则该参数必传。</dx-alert>
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 是否同时切换弹性数据云盘计费模式。取值范围：<li> true：表示切换弹性数据云盘计费模式</li><li> false：表示不切换弹性数据云盘计费模式</li>默认取值：false。
	ModifyPortableDataDisk *bool `json:"ModifyPortableDataDisk,omitnil,omitempty" name:"ModifyPortableDataDisk"`
}

func (r *ModifyInstancesChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesChargeTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "ModifyPortableDataDisk")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesChargeTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesChargeTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancesChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesChargeTypeResponseParams `json:"Response"`
}

func (r *ModifyInstancesChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesDisasterRecoverGroupRequestParams struct {
	// 一个或多个待操作的实例ID。可通过[ DescribeInstances ](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 分散置放群组ID，可使用[DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/api/213/17810)接口获取
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitnil,omitempty" name:"DisasterRecoverGroupId"`

	// 是否强制更换实例宿主机。取值范围：<br><li>true：表示允许实例更换宿主机，允许重启实例。本地盘子机不支持指定此参数。</li><br><li>false：不允许实例更换宿主机，只在当前宿主机上加入置放群组。这可能导致更换置放群组失败。</li><br><br>默认取值：false
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type ModifyInstancesDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过[ DescribeInstances ](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 分散置放群组ID，可使用[DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/api/213/17810)接口获取
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitnil,omitempty" name:"DisasterRecoverGroupId"`

	// 是否强制更换实例宿主机。取值范围：<br><li>true：表示允许实例更换宿主机，允许重启实例。本地盘子机不支持指定此参数。</li><br><li>false：不允许实例更换宿主机，只在当前宿主机上加入置放群组。这可能导致更换置放群组失败。</li><br><br>默认取值：false
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

func (r *ModifyInstancesDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesDisasterRecoverGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "DisasterRecoverGroupId")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesDisasterRecoverGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesDisasterRecoverGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancesDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesDisasterRecoverGroupResponseParams `json:"Response"`
}

func (r *ModifyInstancesDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesProjectRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 项目ID。项目可以使用[ AddProject ](https://cloud.tencent.com/document/api/651/81952)接口创建。可通过[ DescribeProject ](https://cloud.tencent.com/document/api/651/78725) 接口返回值中的`projectId`获取。后续使用[DescribeInstances](https://cloud.tencent.com/document/api/213/15728)接口查询实例时，项目ID可用于过滤结果。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyInstancesProjectRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 项目ID。项目可以使用[ AddProject ](https://cloud.tencent.com/document/api/651/81952)接口创建。可通过[ DescribeProject ](https://cloud.tencent.com/document/api/651/78725) 接口返回值中的`projectId`获取。后续使用[DescribeInstances](https://cloud.tencent.com/document/api/213/15728)接口查询实例时，项目ID可用于过滤结果。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ModifyInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesProjectResponseParams `json:"Response"`
}

func (r *ModifyInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesRenewFlagRequestParams struct {
	// 一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li><br>若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type ModifyInstancesRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li><br>若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

func (r *ModifyInstancesRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesRenewFlagResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancesRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyInstancesRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesVpcAttributeRequestParams struct {
	// 待操作的实例ID数组。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 私有网络相关信息配置，通过该参数指定私有网络的ID，子网ID，私有网络ip等信息。<br><li>当指定私有网络ID和子网ID（子网必须在实例所在的可用区）与指定实例所在私有网络不一致时，会将实例迁移至指定的私有网络的子网下。</li><li>可通过`PrivateIpAddresses`指定私有网络子网IP，若需指定则所有已指定的实例均需要指定子网IP，此时`InstanceIds`与`PrivateIpAddresses`一一对应。</li><li>不指定`PrivateIpAddresses`时随机分配私有网络子网IP。</li>
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 是否对运行中的实例选择强制关机。默认为true。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`

	// 是否保留主机名。默认为false。
	ReserveHostName *bool `json:"ReserveHostName,omitnil,omitempty" name:"ReserveHostName"`
}

type ModifyInstancesVpcAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的实例ID数组。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 私有网络相关信息配置，通过该参数指定私有网络的ID，子网ID，私有网络ip等信息。<br><li>当指定私有网络ID和子网ID（子网必须在实例所在的可用区）与指定实例所在私有网络不一致时，会将实例迁移至指定的私有网络的子网下。</li><li>可通过`PrivateIpAddresses`指定私有网络子网IP，若需指定则所有已指定的实例均需要指定子网IP，此时`InstanceIds`与`PrivateIpAddresses`一一对应。</li><li>不指定`PrivateIpAddresses`时随机分配私有网络子网IP。</li>
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 是否对运行中的实例选择强制关机。默认为true。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`

	// 是否保留主机名。默认为false。
	ReserveHostName *bool `json:"ReserveHostName,omitnil,omitempty" name:"ReserveHostName"`
}

func (r *ModifyInstancesVpcAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesVpcAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "VirtualPrivateCloud")
	delete(f, "ForceStop")
	delete(f, "ReserveHostName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesVpcAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesVpcAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancesVpcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesVpcAttributeResponseParams `json:"Response"`
}

func (r *ModifyInstancesVpcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesVpcAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKeyPairAttributeRequestParams struct {
	// 密钥对ID。可以通过以下方式获取可用的密钥 ID：
	// <li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥 ID。</li>
	// <li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/9403) ，取返回信息中的 `KeyId` 获取密钥对 ID。</li>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 修改后的密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 修改后的密钥对描述信息。可任意命名，但不得超过60个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyKeyPairAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对ID。可以通过以下方式获取可用的密钥 ID：
	// <li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥 ID。</li>
	// <li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/9403) ，取返回信息中的 `KeyId` 获取密钥对 ID。</li>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 修改后的密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 修改后的密钥对描述信息。可任意命名，但不得超过60个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyKeyPairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKeyPairAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "KeyName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKeyPairAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKeyPairAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKeyPairAttributeResponseParams `json:"Response"`
}

func (r *ModifyKeyPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKeyPairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaunchTemplateDefaultVersionRequestParams struct {
	// 启动模板ID。可通过 [DescribeLaunchTemplates](https://cloud.tencent.com/document/api/213/66322) 接口返回值中的`LaunchTemplateId `获取。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 待设置的默认版本号。可通过 [DescribeLaunchTemplateVersions](https://cloud.tencent.com/document/api/213/66323) 接口返回值中的`LaunchTemplateVersion`获取。
	DefaultVersion *int64 `json:"DefaultVersion,omitnil,omitempty" name:"DefaultVersion"`
}

type ModifyLaunchTemplateDefaultVersionRequest struct {
	*tchttp.BaseRequest
	
	// 启动模板ID。可通过 [DescribeLaunchTemplates](https://cloud.tencent.com/document/api/213/66322) 接口返回值中的`LaunchTemplateId `获取。
	LaunchTemplateId *string `json:"LaunchTemplateId,omitnil,omitempty" name:"LaunchTemplateId"`

	// 待设置的默认版本号。可通过 [DescribeLaunchTemplateVersions](https://cloud.tencent.com/document/api/213/66323) 接口返回值中的`LaunchTemplateVersion`获取。
	DefaultVersion *int64 `json:"DefaultVersion,omitnil,omitempty" name:"DefaultVersion"`
}

func (r *ModifyLaunchTemplateDefaultVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaunchTemplateDefaultVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateId")
	delete(f, "DefaultVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLaunchTemplateDefaultVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaunchTemplateDefaultVersionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLaunchTemplateDefaultVersionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLaunchTemplateDefaultVersionResponseParams `json:"Response"`
}

func (r *ModifyLaunchTemplateDefaultVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaunchTemplateDefaultVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationCountLimit struct {
	// 实例操作。取值范围：<br><li>`INSTANCE_DEGRADE`：降配操作</li><li>`INTERNET_CHARGE_TYPE_CHANGE`：修改网络带宽计费模式</li>
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 当前已使用次数，如果返回值为-1表示该操作无次数限制。
	CurrentCount *int64 `json:"CurrentCount,omitnil,omitempty" name:"CurrentCount"`

	// 操作次数最高额度，如果返回值为-1表示该操作无次数限制，如果返回值为0表示不支持调整配置。
	LimitCount *int64 `json:"LimitCount,omitnil,omitempty" name:"LimitCount"`
}

type OsVersion struct {
	// 操作系统类型
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 支持的操作系统版本
	OsVersions []*string `json:"OsVersions,omitnil,omitempty" name:"OsVersions"`

	// 支持的操作系统架构
	Architecture []*string `json:"Architecture,omitnil,omitempty" name:"Architecture"`
}

type Placement struct {
	// 实例所属的可用区名称。该参数可以通过调用  [DescribeZones](https://cloud.tencent.com/document/product/213/15707) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 `ProjectId` 字段来获取。默认取值0，表示默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例所属的专用宿主机ID列表，仅用于入参。如果您有购买专用宿主机并且指定了该参数，则您购买的实例就会随机的部署在这些专用宿主机上。该参数可以通过调用 [DescribeHosts](https://cloud.tencent.com/document/api/213/16474) 的返回值中的 `HostId` 字段来获取。
	HostIds []*string `json:"HostIds,omitnil,omitempty" name:"HostIds"`

	// 实例所属的专用宿主机ID，仅用于出参。
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`
}

type PostPaidQuota struct {
	// 累计已使用配额
	UsedQuota *uint64 `json:"UsedQuota,omitnil,omitempty" name:"UsedQuota"`

	// 剩余配额
	RemainingQuota *uint64 `json:"RemainingQuota,omitnil,omitempty" name:"RemainingQuota"`

	// 总配额
	TotalQuota *uint64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type PrePaidQuota struct {
	// 当月已使用配额
	UsedQuota *uint64 `json:"UsedQuota,omitnil,omitempty" name:"UsedQuota"`

	// 单次购买最大数量
	OnceQuota *uint64 `json:"OnceQuota,omitnil,omitempty" name:"OnceQuota"`

	// 剩余配额
	RemainingQuota *uint64 `json:"RemainingQuota,omitnil,omitempty" name:"RemainingQuota"`

	// 总配额
	TotalQuota *uint64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type Price struct {
	// 描述了实例价格。
	InstancePrice *ItemPrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 描述了网络价格。
	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitnil,omitempty" name:"BandwidthPrice"`
}

// Predefined struct for user
type ProgramFpgaImageRequestParams struct {
	// 实例的ID信息。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// FPGA镜像文件的COS URL地址。
	FPGAUrl *string `json:"FPGAUrl,omitnil,omitempty" name:"FPGAUrl"`

	// 实例上FPGA卡的DBDF号，不填默认烧录FPGA镜像到实例所拥有的所有FPGA卡。
	DBDFs []*string `json:"DBDFs,omitnil,omitempty" name:"DBDFs"`

	// 试运行，不会执行实际的烧录动作，默认为False。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type ProgramFpgaImageRequest struct {
	*tchttp.BaseRequest
	
	// 实例的ID信息。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// FPGA镜像文件的COS URL地址。
	FPGAUrl *string `json:"FPGAUrl,omitnil,omitempty" name:"FPGAUrl"`

	// 实例上FPGA卡的DBDF号，不填默认烧录FPGA镜像到实例所拥有的所有FPGA卡。
	DBDFs []*string `json:"DBDFs,omitnil,omitempty" name:"DBDFs"`

	// 试运行，不会执行实际的烧录动作，默认为False。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *ProgramFpgaImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProgramFpgaImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FPGAUrl")
	delete(f, "DBDFs")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProgramFpgaImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProgramFpgaImageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ProgramFpgaImageResponse struct {
	*tchttp.BaseResponse
	Response *ProgramFpgaImageResponseParams `json:"Response"`
}

func (r *ProgramFpgaImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProgramFpgaImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PurchaseReservedInstancesOfferingRequestParams struct {
	// 购买预留实例计费数量
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 预留实例计费配置ID
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitnil,omitempty" name:"ReservedInstancesOfferingId"`

	// 试运行
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 预留实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>最多支持60个字符（包含模式串）。</li>
	ReservedInstanceName *string `json:"ReservedInstanceName,omitnil,omitempty" name:"ReservedInstanceName"`
}

type PurchaseReservedInstancesOfferingRequest struct {
	*tchttp.BaseRequest
	
	// 购买预留实例计费数量
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 预留实例计费配置ID
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitnil,omitempty" name:"ReservedInstancesOfferingId"`

	// 试运行
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 预留实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>最多支持60个字符（包含模式串）。</li>
	ReservedInstanceName *string `json:"ReservedInstanceName,omitnil,omitempty" name:"ReservedInstanceName"`
}

func (r *PurchaseReservedInstancesOfferingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurchaseReservedInstancesOfferingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceCount")
	delete(f, "ReservedInstancesOfferingId")
	delete(f, "DryRun")
	delete(f, "ClientToken")
	delete(f, "ReservedInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PurchaseReservedInstancesOfferingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PurchaseReservedInstancesOfferingResponseParams struct {
	// 已购买预留实例计费ID
	ReservedInstanceId *string `json:"ReservedInstanceId,omitnil,omitempty" name:"ReservedInstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PurchaseReservedInstancesOfferingResponse struct {
	*tchttp.BaseResponse
	Response *PurchaseReservedInstancesOfferingResponseParams `json:"Response"`
}

func (r *PurchaseReservedInstancesOfferingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurchaseReservedInstancesOfferingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootInstancesRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 本参数已弃用，推荐使用StopType，不可以与参数StopType同时使用。表示是否在正常重启失败后选择强制重启实例。取值范围：<br><li>true：表示在正常重启失败后进行强制重启</li><li>false：表示在正常重启失败后不进行强制重启</li><br>默认取值：false。
	//
	// Deprecated: ForceReboot is deprecated.
	ForceReboot *bool `json:"ForceReboot,omitnil,omitempty" name:"ForceReboot"`

	// 关机类型。取值范围：<br><li>SOFT：表示软关机</li><li>HARD：表示硬关机</li><li>SOFT_FIRST：表示优先软关机，失败再执行硬关机</li><br>默认取值：SOFT。
	StopType *string `json:"StopType,omitnil,omitempty" name:"StopType"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 本参数已弃用，推荐使用StopType，不可以与参数StopType同时使用。表示是否在正常重启失败后选择强制重启实例。取值范围：<br><li>true：表示在正常重启失败后进行强制重启</li><li>false：表示在正常重启失败后不进行强制重启</li><br>默认取值：false。
	ForceReboot *bool `json:"ForceReboot,omitnil,omitempty" name:"ForceReboot"`

	// 关机类型。取值范围：<br><li>SOFT：表示软关机</li><li>HARD：表示硬关机</li><li>SOFT_FIRST：表示优先软关机，失败再执行硬关机</li><br>默认取值：SOFT。
	StopType *string `json:"StopType,omitnil,omitempty" name:"StopType"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ForceReboot")
	delete(f, "StopType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebootInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RebootInstancesResponseParams `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域名称，例如，ap-guangzhou
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域描述，例如，华南地区(广州)
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地域是否可用状态
	RegionState *string `json:"RegionState,omitnil,omitempty" name:"RegionState"`
}

// Predefined struct for user
type RemoveChcAssistVpcRequestParams struct {
	// CHC物理服务器Id。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`
}

type RemoveChcAssistVpcRequest struct {
	*tchttp.BaseRequest
	
	// CHC物理服务器Id。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`
}

func (r *RemoveChcAssistVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveChcAssistVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveChcAssistVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveChcAssistVpcResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveChcAssistVpcResponse struct {
	*tchttp.BaseResponse
	Response *RemoveChcAssistVpcResponseParams `json:"Response"`
}

func (r *RemoveChcAssistVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveChcAssistVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveChcDeployVpcRequestParams struct {
	// CHC物理服务器Id。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`
}

type RemoveChcDeployVpcRequest struct {
	*tchttp.BaseRequest
	
	// CHC物理服务器Id。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`
}

func (r *RemoveChcDeployVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveChcDeployVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveChcDeployVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveChcDeployVpcResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveChcDeployVpcResponse struct {
	*tchttp.BaseResponse
	Response *RemoveChcDeployVpcResponseParams `json:"Response"`
}

func (r *RemoveChcDeployVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveChcDeployVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewHostsRequestParams struct {
	// 一个或多个待操作的CDH实例ID。每次请求的CDH实例的上限为100。
	HostIds []*string `json:"HostIds,omitnil,omitempty" name:"HostIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitnil,omitempty" name:"HostChargePrepaid"`
}

type RenewHostsRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的CDH实例ID。每次请求的CDH实例的上限为100。
	HostIds []*string `json:"HostIds,omitnil,omitempty" name:"HostIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitnil,omitempty" name:"HostChargePrepaid"`
}

func (r *RenewHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostIds")
	delete(f, "HostChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewHostsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewHostsResponse struct {
	*tchttp.BaseResponse
	Response *RenewHostsResponseParams `json:"Response"`
}

func (r *RenewHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstancesRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。<dx-alert infotype="explain" title="">
	// 包年包月实例该参数为必传参数。</dx-alert>
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 是否续费弹性数据盘。取值范围：<br><li>true：表示续费包年包月实例同时续费其挂载的弹性数据盘</li><li>false：表示续费包年包月实例同时不再续费其挂载的弹性数据盘</li><br>默认取值：true。
	RenewPortableDataDisk *bool `json:"RenewPortableDataDisk,omitnil,omitempty" name:"RenewPortableDataDisk"`
}

type RenewInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。<dx-alert infotype="explain" title="">
	// 包年包月实例该参数为必传参数。</dx-alert>
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 是否续费弹性数据盘。取值范围：<br><li>true：表示续费包年包月实例同时续费其挂载的弹性数据盘</li><li>false：表示续费包年包月实例同时不再续费其挂载的弹性数据盘</li><br>默认取值：true。
	RenewPortableDataDisk *bool `json:"RenewPortableDataDisk,omitnil,omitempty" name:"RenewPortableDataDisk"`
}

func (r *RenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	delete(f, "RenewPortableDataDisk")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstancesResponseParams `json:"Response"`
}

func (r *RenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RepairTaskControlRequestParams struct {
	// 待授权任务实例对应的产品类型，支持取值：
	// 
	// - `CVM`：云服务器
	// - `CDH`：专用宿主机
	// - `CPM2.0`：裸金属云服务器
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 指定待操作的实例ID列表，仅允许对列表中的实例ID相关的维修任务发起授权，可通过 [DescribeTaskInfo](https://cloud.tencent.com/document/api/213/87933) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 指定待操作的维修任务ID，可通过 [DescribeTaskInfo](https://cloud.tencent.com/document/api/213/87933) 接口返回值中的`TaskId`获取。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作类型，当前只支持传入`AuthorizeRepair`。
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 预约授权时间，形如`2023-01-01 12:00:00`。预约时间需晚于当前时间至少5分钟，且在48小时之内。
	OrderAuthTime *string `json:"OrderAuthTime,omitnil,omitempty" name:"OrderAuthTime"`

	// 附加的授权处理策略，不传或为空时，按默认授权方式进行处理。对于支持弃盘迁移授权的维修任务，当且仅当传入`LossyLocal`时，代表本次授权可允许发起弃盘迁移。
	// 
	// 注意：
	// 1. 指定`TaskSubMethod`为`LossyLocal`调用接口发起**弃盘迁移授权**时，本地盘实例的**所有本地盘数据都会清空**，相当于**重新部署本地盘实例**。
	// 2. 对于非本地盘实例，或不支持弃盘迁移选项的任务，指定`TaskSubMethod`为`LossyLocal`时接口不会报错，不过后端会自动忽略该参数。
	// 3. 特别的：如果本地盘实例系统盘是CBS云盘，并且`/etc/fstab`里之前配置了本地盘的自动挂载项，建议可根据业务侧的实际需求，评估是否在对应挂载项追加上`nofail`参数（代表对应挂载点挂载失败不阻塞开机流程）或注释对应的挂载路径。否则授权弃盘迁移后，对应本地盘数据已清空，自动挂载失败会导致实例开机流程失败进入救援模式。具体可参考 [Linux 实例：/etc/fstab 配置错误导致无法登录](https://cloud.tencent.com/document/product/213/72039)。
	TaskSubMethod *string `json:"TaskSubMethod,omitnil,omitempty" name:"TaskSubMethod"`
}

type RepairTaskControlRequest struct {
	*tchttp.BaseRequest
	
	// 待授权任务实例对应的产品类型，支持取值：
	// 
	// - `CVM`：云服务器
	// - `CDH`：专用宿主机
	// - `CPM2.0`：裸金属云服务器
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 指定待操作的实例ID列表，仅允许对列表中的实例ID相关的维修任务发起授权，可通过 [DescribeTaskInfo](https://cloud.tencent.com/document/api/213/87933) 接口返回值中的`InstanceId`获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 指定待操作的维修任务ID，可通过 [DescribeTaskInfo](https://cloud.tencent.com/document/api/213/87933) 接口返回值中的`TaskId`获取。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作类型，当前只支持传入`AuthorizeRepair`。
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 预约授权时间，形如`2023-01-01 12:00:00`。预约时间需晚于当前时间至少5分钟，且在48小时之内。
	OrderAuthTime *string `json:"OrderAuthTime,omitnil,omitempty" name:"OrderAuthTime"`

	// 附加的授权处理策略，不传或为空时，按默认授权方式进行处理。对于支持弃盘迁移授权的维修任务，当且仅当传入`LossyLocal`时，代表本次授权可允许发起弃盘迁移。
	// 
	// 注意：
	// 1. 指定`TaskSubMethod`为`LossyLocal`调用接口发起**弃盘迁移授权**时，本地盘实例的**所有本地盘数据都会清空**，相当于**重新部署本地盘实例**。
	// 2. 对于非本地盘实例，或不支持弃盘迁移选项的任务，指定`TaskSubMethod`为`LossyLocal`时接口不会报错，不过后端会自动忽略该参数。
	// 3. 特别的：如果本地盘实例系统盘是CBS云盘，并且`/etc/fstab`里之前配置了本地盘的自动挂载项，建议可根据业务侧的实际需求，评估是否在对应挂载项追加上`nofail`参数（代表对应挂载点挂载失败不阻塞开机流程）或注释对应的挂载路径。否则授权弃盘迁移后，对应本地盘数据已清空，自动挂载失败会导致实例开机流程失败进入救援模式。具体可参考 [Linux 实例：/etc/fstab 配置错误导致无法登录](https://cloud.tencent.com/document/product/213/72039)。
	TaskSubMethod *string `json:"TaskSubMethod,omitnil,omitempty" name:"TaskSubMethod"`
}

func (r *RepairTaskControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RepairTaskControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceIds")
	delete(f, "TaskId")
	delete(f, "Operate")
	delete(f, "OrderAuthTime")
	delete(f, "TaskSubMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RepairTaskControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RepairTaskControlResponseParams struct {
	// 已完成授权的维修任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RepairTaskControlResponse struct {
	*tchttp.BaseResponse
	Response *RepairTaskControlResponseParams `json:"Response"`
}

func (r *RepairTaskControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RepairTaskControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RepairTaskInfo struct {
	// 维修任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 任务类型ID，与任务类型中文名的对应关系如下：
	// 
	// - `101`：实例运行隐患
	// - `102`：实例运行异常
	// - `103`：实例硬盘异常
	// - `104`：实例网络连接异常
	// - `105`：实例运行预警
	// - `106`：实例硬盘预警
	// - `107`：实例维护升级
	// 
	// 各任务类型的具体含义，可参考 [维修任务分类](https://cloud.tencent.com/document/product/213/67789#.E7.BB.B4.E4.BF.AE.E4.BB.BB.E5.8A.A1.E5.88.86.E7.B1.BB)。
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 任务类型中文名
	TaskTypeName *string `json:"TaskTypeName,omitnil,omitempty" name:"TaskTypeName"`

	// 任务状态ID，与任务状态中文名的对应关系如下：
	// 
	// - `1`：待授权
	// - `2`：处理中
	// - `3`：已结束
	// - `4`：已预约
	// - `5`：已取消
	// - `6`：已避免
	// 
	// 各任务状态的具体含义，可参考 [任务状态](https://cloud.tencent.com/document/product/213/67789#.E4.BB.BB.E5.8A.A1.E7.8A.B6.E6.80.81)。
	TaskStatus *uint64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 设备状态ID，与设备状态中文名的对应关系如下：
	// 
	// - `1`：故障中
	// - `2`：处理中
	// - `3`：正常
	// - `4`：已预约
	// - `5`：已取消
	// - `6`：已避免
	DeviceStatus *uint64 `json:"DeviceStatus,omitnil,omitempty" name:"DeviceStatus"`

	// 操作状态ID，与操作状态中文名的对应关系如下：
	// 
	// - `1`：未授权
	// - `2`：已授权
	// - `3`：已处理
	// - `4`：已预约
	// - `5`：已取消
	// - `6`：已避免
	OperateStatus *uint64 `json:"OperateStatus,omitnil,omitempty" name:"OperateStatus"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务授权时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthTime *string `json:"AuthTime,omitnil,omitempty" name:"AuthTime"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskDetail *string `json:"TaskDetail,omitnil,omitempty" name:"TaskDetail"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 所在私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 所在私有网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 所在子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 所在子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 实例公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitnil,omitempty" name:"WanIp"`

	// 实例内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIp *string `json:"LanIp,omitnil,omitempty" name:"LanIp"`

	// 产品类型，支持取值：
	// 
	// - `CVM`：云服务器
	// - `CDH`：专用宿主机
	// - `CPM2.0`：裸金属云服务器
	// 注意：此字段可能返回 null，表示取不到有效值。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 任务子类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskSubType *string `json:"TaskSubType,omitnil,omitempty" name:"TaskSubType"`

	// 任务授权类型，当前`AuthType`和维修任务提供的授权选项的对应关系如下：
	// 
	// - `"1"`：仅提供【在线迁移授权】
	// - `"2"`：仅提供【停机授权】
	// - `"3"`：仅提供【在线换盘授权】
	// - `"4"`：提供【停机换盘授权】（默认）、【弃盘迁移授权】（可选）
	// - `"5"`：提供【停机授权】（默认）、【弃盘迁移授权】（可选）
	// - `"6"`：仅提供【在线维护授权】
	// - `"7"`：提供【在线维护授权】（默认）、【停机授权】（可选）
	// - `"8"`：仅提供【弃盘迁移授权】
	AuthType *uint64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 授权渠道，支持取值：
	// 
	// - `Waiting_auth`：待授权
	// - `Customer_auth`：客户操作授权
	// - `System_mandatory_auth`：系统默认授权
	// - `Pre_policy_auth`：预置策略授权
	AuthSource *string `json:"AuthSource,omitnil,omitempty" name:"AuthSource"`
}

type ReservedInstanceConfigInfoItem struct {
	// 实例规格。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例规格名称。
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 优先级。
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 实例族信息列表。
	InstanceFamilies []*ReservedInstanceFamilyItem `json:"InstanceFamilies,omitnil,omitempty" name:"InstanceFamilies"`
}

type ReservedInstanceFamilyItem struct {
	// 实例族。
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// 优先级。
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 实例类型信息列表。
	InstanceTypes []*ReservedInstanceTypeItem `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`
}

type ReservedInstancePrice struct {
	// 预支合计费用的原价，单位：元。
	OriginalFixedPrice *float64 `json:"OriginalFixedPrice,omitnil,omitempty" name:"OriginalFixedPrice"`

	// 预支合计费用的折扣价，单位：元。
	DiscountFixedPrice *float64 `json:"DiscountFixedPrice,omitnil,omitempty" name:"DiscountFixedPrice"`

	// 后续合计费用的原价，单位：元/小时
	OriginalUsagePrice *float64 `json:"OriginalUsagePrice,omitnil,omitempty" name:"OriginalUsagePrice"`

	// 后续合计费用的折扣价，单位：元/小时
	DiscountUsagePrice *float64 `json:"DiscountUsagePrice,omitnil,omitempty" name:"DiscountUsagePrice"`

	// 预支费用的折扣，如20.0代表2折。 注意：此字段可能返回 null，表示取不到有效值。
	FixedPriceDiscount *float64 `json:"FixedPriceDiscount,omitnil,omitempty" name:"FixedPriceDiscount"`

	// 后续费用的折扣，如20.0代表2折。 注意：此字段可能返回 null，表示取不到有效值。
	UsagePriceDiscount *float64 `json:"UsagePriceDiscount,omitnil,omitempty" name:"UsagePriceDiscount"`
}

type ReservedInstancePriceItem struct {
	// 付费类型，如："All Upfront","Partial Upfront","No Upfront"
	OfferingType *string `json:"OfferingType,omitnil,omitempty" name:"OfferingType"`

	// 预支合计费用，单位：元。
	FixedPrice *float64 `json:"FixedPrice,omitnil,omitempty" name:"FixedPrice"`

	// 后续合计费用，单位：元/小时
	UsagePrice *float64 `json:"UsagePrice,omitnil,omitempty" name:"UsagePrice"`

	// 预留实例配置ID
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitnil,omitempty" name:"ReservedInstancesOfferingId"`

	// 预留实例计费可购买的可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 预留实例计费【有效期】即预留实例计费购买时长。形如：31536000。
	// 计量单位：秒
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 预留实例计费的平台描述（即操作系统）。形如：Linux。
	// 返回项： Linux 。
	ProductDescription *string `json:"ProductDescription,omitnil,omitempty" name:"ProductDescription"`

	// 预支合计费用，单位：元。
	DiscountUsagePrice *float64 `json:"DiscountUsagePrice,omitnil,omitempty" name:"DiscountUsagePrice"`

	// 后续合计费用的折扣价，单位：元/小时
	DiscountFixedPrice *float64 `json:"DiscountFixedPrice,omitnil,omitempty" name:"DiscountFixedPrice"`
}

type ReservedInstanceTypeItem struct {
	// 实例类型。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// CPU核数。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// GPU数量。
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// FPGA数量。
	Fpga *uint64 `json:"Fpga,omitnil,omitempty" name:"Fpga"`

	// 本地存储块数量。
	StorageBlock *uint64 `json:"StorageBlock,omitnil,omitempty" name:"StorageBlock"`

	// 网卡数。
	NetworkCard *uint64 `json:"NetworkCard,omitnil,omitempty" name:"NetworkCard"`

	// 最大带宽。
	MaxBandwidth *float64 `json:"MaxBandwidth,omitnil,omitempty" name:"MaxBandwidth"`

	// 主频。
	Frequency *string `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// CPU型号名称。
	CpuModelName *string `json:"CpuModelName,omitnil,omitempty" name:"CpuModelName"`

	// 包转发率。
	Pps *uint64 `json:"Pps,omitnil,omitempty" name:"Pps"`

	// 外部信息。
	Externals *Externals `json:"Externals,omitnil,omitempty" name:"Externals"`

	// 备注信息。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 预留实例配置价格信息。
	Prices []*ReservedInstancePriceItem `json:"Prices,omitnil,omitempty" name:"Prices"`
}

type ReservedInstancesOffering struct {
	// 预留实例计费可购买的可用区。形如：ap-guangzhou-1。
	// 返回项：<a href="https://cloud.tencent.com/document/product/213/6091">可用区列表</a>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可购买的预留实例计费类型的结算货币，使用ISO 4217标准货币代码。
	// 返回项：USD（美元）。
	CurrencyCode *string `json:"CurrencyCode,omitnil,omitempty" name:"CurrencyCode"`

	// 预留实例计费【有效期】即预留实例计费购买时长。形如：31536000。
	// 计量单位：秒
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 预留实例计费的购买价格。形如：4000.0。
	// 计量单位：与 currencyCode 一致，目前支持 USD（美元）
	FixedPrice *float64 `json:"FixedPrice,omitnil,omitempty" name:"FixedPrice"`

	// 预留实例计费的实例类型。形如：S3.MEDIUM4。
	// 返回项：<a href="https://cloud.tencent.com/product/cvm/instances">预留实例计费类型列表</a>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 预留实例计费的付款类型。形如：All Upfront。
	// 返回项： All Upfront (预付全部费用)。
	OfferingType *string `json:"OfferingType,omitnil,omitempty" name:"OfferingType"`

	// 可购买的预留实例计费配置ID。形如：650c138f-ae7e-4750-952a-96841d6e9fc1。
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitnil,omitempty" name:"ReservedInstancesOfferingId"`

	// 预留实例计费的平台描述（即操作系统）。形如：linux。
	// 返回项： linux 。
	ProductDescription *string `json:"ProductDescription,omitnil,omitempty" name:"ProductDescription"`

	// 扣除预付费之后的使用价格 (按小时计费)。形如：0.0。
	// 目前，因为只支持 All Upfront 付款类型，所以默认为 0元/小时。
	// 计量单位：元/小时，货币单位与 currencyCode 一致，目前支持 USD（美元）
	UsagePrice *float64 `json:"UsagePrice,omitnil,omitempty" name:"UsagePrice"`
}

// Predefined struct for user
type ResetInstanceRequestParams struct {
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	// 默认取值：默认使用当前镜像。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务；自定义镜像与镜像市场镜像默认不开启云监控，云安全服务，而使用镜像里保留的服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 重装系统时，可以指定修改实例的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><br><li>Windows 实例：主机名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><br><li>其他类型（Linux 等）实例：主机名字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526)和[Linux](https://cloud.tencent.com/document/product/213/17525)启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	// 默认取值：默认使用当前镜像。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务；自定义镜像与镜像市场镜像默认不开启云监控，云安全服务，而使用镜像里保留的服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 重装系统时，可以指定修改实例的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><br><li>Windows 实例：主机名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><br><li>其他类型（Linux 等）实例：主机名字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526)和[Linux](https://cloud.tencent.com/document/product/213/17525)启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

func (r *ResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ImageId")
	delete(f, "SystemDisk")
	delete(f, "LoginSettings")
	delete(f, "EnhancedService")
	delete(f, "HostName")
	delete(f, "UserData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstanceResponseParams `json:"Response"`
}

func (r *ResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesInternetMaxBandwidthRequestParams struct {
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/9388)接口返回值中的 `InstanceId` 获取。 每次请求批量实例的上限为100。当调整 `BANDWIDTH_PREPAID` 和 `BANDWIDTH_POSTPAID_BY_HOUR` 计费方式的带宽时，只支持一个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持 `InternetMaxBandwidthOut` 参数。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 带宽生效的起始时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 带宽生效的终止时间。格式： `YYYY-MM-DD` ，例如：`2016-10-30` 。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/9388)接口返回值中的`ExpiredTime`获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/9388)接口返回值中的 `InstanceId` 获取。 每次请求批量实例的上限为100。当调整 `BANDWIDTH_PREPAID` 和 `BANDWIDTH_POSTPAID_BY_HOUR` 计费方式的带宽时，只支持一个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持 `InternetMaxBandwidthOut` 参数。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 带宽生效的起始时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 带宽生效的终止时间。格式： `YYYY-MM-DD` ，例如：`2016-10-30` 。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/9388)接口返回值中的`ExpiredTime`获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InternetAccessible")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesInternetMaxBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesInternetMaxBandwidthResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesInternetMaxBandwidthResponseParams `json:"Response"`
}

func (r *ResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesPasswordRequestParams struct {
	// 一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 重置后的实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：Linux 实例密码必须8-30位，推荐使用12位以上密码，不能以“/”开头，至少包含以下字符中的三种不同字符，字符种类：<br><li>小写字母：[a-z]</li><li>大写字母：[A-Z]</li><li>数字：0-9</li><li>特殊字符： ()\`\~!@#$%^&\*-+=\_|{}[]:;'<>,.?/Windows 实例密码必须12\~30位，不能以“/”开头且不包括用户名，至少包含以下字符中的三种不同字符</li><li>小写字母：[a-z]</li><li>大写字母：[A-Z]</li><li>数字： 0-9</li><li>特殊字符：()\`\~!@#$%^&\*-+=\_|{}[]:;' <>,.?/</li><li>如果实例即包含 `Linux` 实例又包含 `Windows` 实例，则密码复杂度限制按照 `Windows` 实例的限制。</li>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 待重置密码的实例操作系统的用户名。不得超过64个字符。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>true：表示在正常关机失败后进行强制关机</li><li>false：表示在正常关机失败后不进行强制关机</li>默认取值：false。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 重置后的实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：Linux 实例密码必须8-30位，推荐使用12位以上密码，不能以“/”开头，至少包含以下字符中的三种不同字符，字符种类：<br><li>小写字母：[a-z]</li><li>大写字母：[A-Z]</li><li>数字：0-9</li><li>特殊字符： ()\`\~!@#$%^&\*-+=\_|{}[]:;'<>,.?/Windows 实例密码必须12\~30位，不能以“/”开头且不包括用户名，至少包含以下字符中的三种不同字符</li><li>小写字母：[a-z]</li><li>大写字母：[A-Z]</li><li>数字： 0-9</li><li>特殊字符：()\`\~!@#$%^&\*-+=\_|{}[]:;' <>,.?/</li><li>如果实例即包含 `Linux` 实例又包含 `Windows` 实例，则密码复杂度限制按照 `Windows` 实例的限制。</li>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 待重置密码的实例操作系统的用户名。不得超过64个字符。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>true：表示在正常关机失败后进行强制关机</li><li>false：表示在正常关机失败后不进行强制关机</li>默认取值：false。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`
}

func (r *ResetInstancesPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Password")
	delete(f, "UserName")
	delete(f, "ForceStop")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesPasswordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetInstancesPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesPasswordResponseParams `json:"Response"`
}

func (r *ResetInstancesPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesTypeRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。本接口目前仅支持每次操作1个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 调整后的实例机型。不同实例机型指定了不同的资源规格，具体取值可通过调用接口[ DescribeInstanceTypeConfigs ](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机。取值范围：<br><li>true：表示在正常关机失败后进行强制关机</li><br><li>false：表示在正常关机失败后不进行强制关机</li><br><br>默认取值：false。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`
}

type ResetInstancesTypeRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。本接口目前仅支持每次操作1个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 调整后的实例机型。不同实例机型指定了不同的资源规格，具体取值可通过调用接口[ DescribeInstanceTypeConfigs ](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机。取值范围：<br><li>true：表示在正常关机失败后进行强制关机</li><br><li>false：表示在正常关机失败后不进行强制关机</li><br><br>默认取值：false。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`
}

func (r *ResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceType")
	delete(f, "ForceStop")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesTypeResponseParams `json:"Response"`
}

func (r *ResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeInstanceDisksRequestParams struct {
	// 待操作的实例ID。可通过[ DescribeInstances ](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待扩容的数据盘配置信息，仅支持指定待扩容盘的目的容量。只支持扩容非弹性数据盘（[ DescribeDisks ](https://cloud.tencent.com/document/api/362/16315)接口返回值中的`Portable`为`false`表示非弹性）。数据盘容量单位：GiB。最小扩容步长：10GiB。关于数据盘类型的选择请参考[硬盘产品简介](https://cloud.tencent.com/document/product/362/2353)。可选数据盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因数据盘类型的不同而有所差异。
	// <dx-alert infotype="explain" title="">您必须指定参数DataDisks与SystemDisk的其中一个，但不能同时指定。</dx-alert>
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再扩容实例磁盘。取值范围：<br><li>true：表示在正常关机失败后进行强制关机</li><br><li>false：表示在正常关机失败后不进行强制关机</li><br><br>默认取值：false。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`

	// 待扩容的系统盘配置信息，仅支持指定待扩容盘的目的容量。只支持扩容云盘。
	// <dx-alert infotype="explain" title="">您必须指定参数DataDisks与SystemDisk的其中一个，但不能同时指定。</dx-alert>
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 扩容云盘的方式是否为在线扩容。
	ResizeOnline *bool `json:"ResizeOnline,omitnil,omitempty" name:"ResizeOnline"`
}

type ResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的实例ID。可通过[ DescribeInstances ](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待扩容的数据盘配置信息，仅支持指定待扩容盘的目的容量。只支持扩容非弹性数据盘（[ DescribeDisks ](https://cloud.tencent.com/document/api/362/16315)接口返回值中的`Portable`为`false`表示非弹性）。数据盘容量单位：GiB。最小扩容步长：10GiB。关于数据盘类型的选择请参考[硬盘产品简介](https://cloud.tencent.com/document/product/362/2353)。可选数据盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因数据盘类型的不同而有所差异。
	// <dx-alert infotype="explain" title="">您必须指定参数DataDisks与SystemDisk的其中一个，但不能同时指定。</dx-alert>
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再扩容实例磁盘。取值范围：<br><li>true：表示在正常关机失败后进行强制关机</li><br><li>false：表示在正常关机失败后不进行强制关机</li><br><br>默认取值：false。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`

	// 待扩容的系统盘配置信息，仅支持指定待扩容盘的目的容量。只支持扩容云盘。
	// <dx-alert infotype="explain" title="">您必须指定参数DataDisks与SystemDisk的其中一个，但不能同时指定。</dx-alert>
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 扩容云盘的方式是否为在线扩容。
	ResizeOnline *bool `json:"ResizeOnline,omitnil,omitempty" name:"ResizeOnline"`
}

func (r *ResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeInstanceDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DataDisks")
	delete(f, "ForceStop")
	delete(f, "SystemDisk")
	delete(f, "ResizeOnline")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeInstanceDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeInstanceDisksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *ResizeInstanceDisksResponseParams `json:"Response"`
}

func (r *ResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeInstanceDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunAutomationServiceEnabled struct {
	// 是否开启云自动化助手。取值范围：<br><li>true：表示开启云自动化助手服务<br><li>false：表示不开启云自动化助手服务<br><br>默认取值：false。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

// Predefined struct for user
type RunInstancesRequestParams struct {
	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月</li><br><li>POSTPAID_BY_HOUR：按小时后付费</li><br><li>CDHPAID：独享子机（基于专用宿主机创建，宿主机部分的资源不收费）</li><br><li>SPOTPAID：竞价付费</li><br><li>CDCPAID：专用集群付费</li><br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，所属宿主机（在专用宿主机上创建子机时指定）等属性。
	//  <b>注：如果您不指定LaunchTemplate参数，则Placement为必选参数。若同时传递Placement和LaunchTemplate，则默认覆盖LaunchTemplate中对应的Placement的值。</b>
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID\_BY\_HOUR的实例创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则系统将根据当前地域的资源售卖情况动态指定默认机型。</li><br><li>对于付费模式为CDHPAID的实例创建，该参数以"CDH_"为前缀，根据CPU和内存配置生成，具体形式为：CDH_XCXG，例如对于创建CPU为1核，内存为1G大小的专用宿主机的实例，该参数应该为CDH_1C1G。</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，传入InstanceType获取当前机型支持的镜像列表，取返回信息中的`ImageId`字段。</li>
	//  <b>注：如果您不指定LaunchTemplate参数，则ImageId为必选参数。若同时传递ImageId和LaunchTemplate，则默认覆盖LaunchTemplate中对应的ImageId的值。</b>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若在此参数中指定了私有网络IP，即表示每个实例的主网卡IP；同时，InstanceCount参数必须与私有网络IP的个数一致且不能大于20。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 购买实例数量。包年包月实例取值范围：[1，500]，按量计费实例取值范围：[1，500]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见[CVM实例购买限制](https://cloud.tencent.com/document/product/213/2664)。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。</li><li>最多支持128个字符（包含模式串）。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的 `SecurityGroupId` 字段来获取。若不指定该参数，则绑定指定项目下的默认安全组，如默认安全组不存在则将自动创建。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务；自定义镜像与镜像市场镜像默认不开启云监控，云安全服务，而使用镜像里保留的服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 实例主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><br><li>Windows 实例：主机名名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><br><li>其他类型（Linux 等）实例：主机名字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li><br><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server{R:3}`，购买1台时，实例主机名为`server3`；购买2台时，实例主机名分别为`server3`，`server4`。支持指定多个模式串`{R:x}`。</li><br><li>购买多台实例，如果不指定模式串，则在实例主机名添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server`，购买2台时，实例主机名分别为`server1`，`server2`。</li>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer,omitnil,omitempty" name:"ActionTimer"`

	// 置放群组id，仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的云服务器、云硬盘实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费但没有传递该参数时，默认按当前固定折扣价格出价。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526)和[Linux](https://cloud.tencent.com/document/product/213/17525)启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 自定义metadata，支持创建 CVM 时添加自定义元数据键值对。
	// **注：内测中**。
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 描述了实例CPU拓扑结构的相关信息。若不指定该参数，则按系统资源情况决定。
	CpuTopology *CpuTopology `json:"CpuTopology,omitnil,omitempty" name:"CpuTopology"`

	// CAM角色名称。可通过[ DescribeRoleList ](https://cloud.tencent.com/document/product/598/36223)接口返回值中的`RoleName `获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 高性能计算集群ID。若创建的实例为高性能计算实例，需指定实例放置的集群，否则不可指定。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 实例启动模板。
	LaunchTemplate *LaunchTemplate `json:"LaunchTemplate,omitnil,omitempty" name:"LaunchTemplate"`

	// 指定专用集群创建。
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 指定CHC物理服务器来创建CHC云主机。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`

	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>true：表示开启实例保护，不允许通过api接口删除实例</li><br><li>false：表示关闭实例保护，允许通过api接口删除实例</li><br><br>默认取值：false。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`

	// 实例是否开启巨型帧，取值范围：<br><li/> true：表示实例开启巨型帧，只有支持巨型帧的机型可设置为true。<br><li/>false：表示实例关闭巨型帧，只有支持巨型帧的机型可设置为false。<br> 支持巨型帧的实例规格： [实例规格](https://cloud.tencent.com/document/product/213/11518)
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitnil,omitempty" name:"EnableJumboFrame"`
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月</li><br><li>POSTPAID_BY_HOUR：按小时后付费</li><br><li>CDHPAID：独享子机（基于专用宿主机创建，宿主机部分的资源不收费）</li><br><li>SPOTPAID：竞价付费</li><br><li>CDCPAID：专用集群付费</li><br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，所属宿主机（在专用宿主机上创建子机时指定）等属性。
	//  <b>注：如果您不指定LaunchTemplate参数，则Placement为必选参数。若同时传递Placement和LaunchTemplate，则默认覆盖LaunchTemplate中对应的Placement的值。</b>
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID\_BY\_HOUR的实例创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则系统将根据当前地域的资源售卖情况动态指定默认机型。</li><br><li>对于付费模式为CDHPAID的实例创建，该参数以"CDH_"为前缀，根据CPU和内存配置生成，具体形式为：CDH_XCXG，例如对于创建CPU为1核，内存为1G大小的专用宿主机的实例，该参数应该为CDH_1C1G。</li>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，传入InstanceType获取当前机型支持的镜像列表，取返回信息中的`ImageId`字段。</li>
	//  <b>注：如果您不指定LaunchTemplate参数，则ImageId为必选参数。若同时传递ImageId和LaunchTemplate，则默认覆盖LaunchTemplate中对应的ImageId的值。</b>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若在此参数中指定了私有网络IP，即表示每个实例的主网卡IP；同时，InstanceCount参数必须与私有网络IP的个数一致且不能大于20。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 购买实例数量。包年包月实例取值范围：[1，500]，按量计费实例取值范围：[1，500]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见[CVM实例购买限制](https://cloud.tencent.com/document/product/213/2664)。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。</li><li>最多支持128个字符（包含模式串）。</li>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的 `SecurityGroupId` 字段来获取。若不指定该参数，则绑定指定项目下的默认安全组，如默认安全组不存在则将自动创建。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务；自定义镜像与镜像市场镜像默认不开启云监控，云安全服务，而使用镜像里保留的服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 实例主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li><br><li>Windows 实例：主机名名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。</li><br><li>其他类型（Linux 等）实例：主机名字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li><br><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server{R:3}`，购买1台时，实例主机名为`server3`；购买2台时，实例主机名分别为`server3`，`server4`。支持指定多个模式串`{R:x}`。</li><br><li>购买多台实例，如果不指定模式串，则在实例主机名添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server`，购买2台时，实例主机名分别为`server1`，`server2`。</li>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer,omitnil,omitempty" name:"ActionTimer"`

	// 置放群组id，仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的云服务器、云硬盘实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费但没有传递该参数时，默认按当前固定折扣价格出价。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅[Windows](https://cloud.tencent.com/document/product/213/17526)和[Linux](https://cloud.tencent.com/document/product/213/17525)启动时运行命令。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 自定义metadata，支持创建 CVM 时添加自定义元数据键值对。
	// **注：内测中**。
	Metadata *Metadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 描述了实例CPU拓扑结构的相关信息。若不指定该参数，则按系统资源情况决定。
	CpuTopology *CpuTopology `json:"CpuTopology,omitnil,omitempty" name:"CpuTopology"`

	// CAM角色名称。可通过[ DescribeRoleList ](https://cloud.tencent.com/document/product/598/36223)接口返回值中的`RoleName `获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 高性能计算集群ID。若创建的实例为高性能计算实例，需指定实例放置的集群，否则不可指定。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// 实例启动模板。
	LaunchTemplate *LaunchTemplate `json:"LaunchTemplate,omitnil,omitempty" name:"LaunchTemplate"`

	// 指定专用集群创建。
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 指定CHC物理服务器来创建CHC云主机。
	ChcIds []*string `json:"ChcIds,omitnil,omitempty" name:"ChcIds"`

	// 实例销毁保护标志，表示是否允许通过api接口删除实例。取值范围：<br><li>true：表示开启实例保护，不允许通过api接口删除实例</li><br><li>false：表示关闭实例保护，允许通过api接口删除实例</li><br><br>默认取值：false。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil,omitempty" name:"DisableApiTermination"`

	// 实例是否开启巨型帧，取值范围：<br><li/> true：表示实例开启巨型帧，只有支持巨型帧的机型可设置为true。<br><li/>false：表示实例关闭巨型帧，只有支持巨型帧的机型可设置为false。<br> 支持巨型帧的实例规格： [实例规格](https://cloud.tencent.com/document/product/213/11518)
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitnil,omitempty" name:"EnableJumboFrame"`
}

func (r *RunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "Placement")
	delete(f, "InstanceType")
	delete(f, "ImageId")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "VirtualPrivateCloud")
	delete(f, "InternetAccessible")
	delete(f, "InstanceCount")
	delete(f, "InstanceName")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "EnhancedService")
	delete(f, "ClientToken")
	delete(f, "HostName")
	delete(f, "ActionTimer")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "TagSpecification")
	delete(f, "InstanceMarketOptions")
	delete(f, "UserData")
	delete(f, "Metadata")
	delete(f, "DryRun")
	delete(f, "CpuTopology")
	delete(f, "CamRoleName")
	delete(f, "HpcClusterId")
	delete(f, "LaunchTemplate")
	delete(f, "DedicatedClusterId")
	delete(f, "ChcIds")
	delete(f, "DisableApiTermination")
	delete(f, "EnableJumboFrame")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunInstancesResponseParams struct {
	// 当通过本接口来创建实例时会返回该参数，表示一个或多个实例`ID`。返回实例`ID`列表并不代表实例创建成功，可根据 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口查询返回的InstancesSet中对应实例的`ID`的状态来判断创建是否完成；如果实例状态由“PENDING(创建中)”变为“RUNNING(运行中)”，则为创建成功。
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RunInstancesResponseParams `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {
	// 是否开启[云监控](/document/product/248)服务。取值范围：<br><li>true：表示开启云监控服务</li><li>false：表示不开启云监控服务</li><br>默认取值：true。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// 是否开启[云安全](/document/product/296)服务。取值范围：<br><li>true：表示开启云安全服务<br><li>false：表示不开启云安全服务<br><br>默认取值：true。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type SharePermission struct {
	// 镜像分享时间。
	// 按照 ISO8601 标准表示，并且使用 UTC 时间，格式为：YYYY-MM-DDThh:mm:ssZ。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 镜像分享的账户ID
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`
}

type Snapshot struct {
	// 快照Id。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 创建此快照的云硬盘类型。取值范围：
	// SYSTEM_DISK：系统盘
	// DATA_DISK：数据盘。
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 创建此快照的云硬盘大小，单位 GiB。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type SpotMarketOptions struct {
	// 竞价出价
	MaxPrice *string `json:"MaxPrice,omitnil,omitempty" name:"MaxPrice"`

	// 竞价请求类型，当前仅支持类型：one-time
	SpotInstanceType *string `json:"SpotInstanceType,omitnil,omitempty" name:"SpotInstanceType"`
}

type SpotPaidQuota struct {
	// 已使用配额，单位：vCPU核心数
	UsedQuota *uint64 `json:"UsedQuota,omitnil,omitempty" name:"UsedQuota"`

	// 剩余配额，单位：vCPU核心数
	RemainingQuota *uint64 `json:"RemainingQuota,omitnil,omitempty" name:"RemainingQuota"`

	// 总配额，单位：vCPU核心数
	TotalQuota *uint64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

// Predefined struct for user
type StartInstancesRequestParams struct {
	// 一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过[DescribeInstances](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *StartInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StartInstancesResponseParams `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 本参数已弃用，推荐使用StopType，不可以与参数StopType同时使用。表示是否在正常关闭失败后选择强制关闭实例。取值范围：<br><li>true：表示在正常关闭失败后进行强制关闭</li><li>false：表示在正常关闭失败后不进行强制关闭</li><br>默认取值：false。
	//
	// Deprecated: ForceStop is deprecated.
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`

	// 实例的关闭模式。取值范围：<br><li>SOFT_FIRST：表示在正常关闭失败后进行强制关闭</li><li>HARD：直接强制关闭</li><li>SOFT：仅软关机</li>默认取值：SOFT。
	StopType *string `json:"StopType,omitnil,omitempty" name:"StopType"`

	// 按量计费实例关机收费模式。
	// 取值范围：<br><li>KEEP_CHARGING：关机继续收费</li><li>STOP_CHARGING：关机停止收费</li><br>默认取值：KEEP_CHARGING。
	// 该参数只针对部分按量计费实例生效，详情参考[按量计费实例关机不收费说明](https://cloud.tencent.com/document/product/213/19918)
	StoppedMode *string `json:"StoppedMode,omitnil,omitempty" name:"StoppedMode"`
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 本参数已弃用，推荐使用StopType，不可以与参数StopType同时使用。表示是否在正常关闭失败后选择强制关闭实例。取值范围：<br><li>true：表示在正常关闭失败后进行强制关闭</li><li>false：表示在正常关闭失败后不进行强制关闭</li><br>默认取值：false。
	ForceStop *bool `json:"ForceStop,omitnil,omitempty" name:"ForceStop"`

	// 实例的关闭模式。取值范围：<br><li>SOFT_FIRST：表示在正常关闭失败后进行强制关闭</li><li>HARD：直接强制关闭</li><li>SOFT：仅软关机</li>默认取值：SOFT。
	StopType *string `json:"StopType,omitnil,omitempty" name:"StopType"`

	// 按量计费实例关机收费模式。
	// 取值范围：<br><li>KEEP_CHARGING：关机继续收费</li><li>STOP_CHARGING：关机停止收费</li><br>默认取值：KEEP_CHARGING。
	// 该参数只针对部分按量计费实例生效，详情参考[按量计费实例关机不收费说明](https://cloud.tencent.com/document/product/213/19918)
	StoppedMode *string `json:"StoppedMode,omitnil,omitempty" name:"StoppedMode"`
}

func (r *StopInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ForceStop")
	delete(f, "StopType")
	delete(f, "StoppedMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StopInstancesResponseParams `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageBlock struct {
	// HDD本地存储类型，值为：LOCAL_PRO.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// HDD本地存储的最小容量。单位：GiB。
	MinSize *int64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// HDD本地存储的最大容量。单位：GiB。
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

type SyncImage struct {
	// 镜像ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type SyncImagesRequestParams struct {
	// 镜像ID列表 ，镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取，镜像ID对应的镜像状态必须为`NORMAL`。</li><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。</li>
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// 目的同步地域列表，必须满足如下限制：<br><li>必须是一个合法的Region。</li><li>如果是自定义镜像，则目标同步地域不能为源地域。</li><li>如果是共享镜像，则目的同步地域仅支持源地域，表示将共享镜像复制为源地域的自定义镜像。</li><li>如果是自定义镜像复制为加密自定义镜像，则目的同步地域仅支持源地域，表示将自定义镜像复制为源地域的加密自定义镜像。</li><li>暂不支持部分地域同步，请参考[复制镜像](https://cloud.tencent.com/document/product/213/4943#.E5.A4.8D.E5.88.B6.E8.AF.B4.E6.98.8E)。</li>具体地域参数请参考[Region](https://cloud.tencent.com/document/product/213/6091)。
	DestinationRegions []*string `json:"DestinationRegions,omitnil,omitempty" name:"DestinationRegions"`

	// 检测是否支持发起同步镜像。
	// 默认值: false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 目标镜像名称。默认使用源镜像名称。
	// 最多支持 60 个字符。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 是否需要返回目的地域的镜像ID。
	// 默认值: false
	ImageSetRequired *bool `json:"ImageSetRequired,omitnil,omitempty" name:"ImageSetRequired"`

	// 是否复制为加密自定义镜像。
	// 默认值为 false。
	// 复制加密自定义镜像仅支持同地域， 即 DestinationRegions 仅支持填写指定镜像所在地域。
	Encrypt *bool `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 加密自定义镜像使用的 KMS 密钥 ID。
	// 仅当复制加密镜像时，即 Encrypt 为 true 时，此参数有效；
	// 不指定 KmsKeyId，默认使用 CBS 云产品 KMS 密钥。
	// KMS 密钥 ID 通过[KMS 控制台](https://console.cloud.tencent.com/kms2)获取。
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`
}

type SyncImagesRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID列表 ，镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取，镜像ID对应的镜像状态必须为`NORMAL`。</li><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。</li>
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// 目的同步地域列表，必须满足如下限制：<br><li>必须是一个合法的Region。</li><li>如果是自定义镜像，则目标同步地域不能为源地域。</li><li>如果是共享镜像，则目的同步地域仅支持源地域，表示将共享镜像复制为源地域的自定义镜像。</li><li>如果是自定义镜像复制为加密自定义镜像，则目的同步地域仅支持源地域，表示将自定义镜像复制为源地域的加密自定义镜像。</li><li>暂不支持部分地域同步，请参考[复制镜像](https://cloud.tencent.com/document/product/213/4943#.E5.A4.8D.E5.88.B6.E8.AF.B4.E6.98.8E)。</li>具体地域参数请参考[Region](https://cloud.tencent.com/document/product/213/6091)。
	DestinationRegions []*string `json:"DestinationRegions,omitnil,omitempty" name:"DestinationRegions"`

	// 检测是否支持发起同步镜像。
	// 默认值: false
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 目标镜像名称。默认使用源镜像名称。
	// 最多支持 60 个字符。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 是否需要返回目的地域的镜像ID。
	// 默认值: false
	ImageSetRequired *bool `json:"ImageSetRequired,omitnil,omitempty" name:"ImageSetRequired"`

	// 是否复制为加密自定义镜像。
	// 默认值为 false。
	// 复制加密自定义镜像仅支持同地域， 即 DestinationRegions 仅支持填写指定镜像所在地域。
	Encrypt *bool `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 加密自定义镜像使用的 KMS 密钥 ID。
	// 仅当复制加密镜像时，即 Encrypt 为 true 时，此参数有效；
	// 不指定 KmsKeyId，默认使用 CBS 云产品 KMS 密钥。
	// KMS 密钥 ID 通过[KMS 控制台](https://console.cloud.tencent.com/kms2)获取。
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`
}

func (r *SyncImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIds")
	delete(f, "DestinationRegions")
	delete(f, "DryRun")
	delete(f, "ImageName")
	delete(f, "ImageSetRequired")
	delete(f, "Encrypt")
	delete(f, "KmsKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncImagesResponseParams struct {
	// 目的地域的镜像ID信息。
	ImageSet []*SyncImage `json:"ImageSet,omitnil,omitempty" name:"ImageSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncImagesResponse struct {
	*tchttp.BaseResponse
	Response *SyncImagesResponseParams `json:"Response"`
}

func (r *SyncImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {
	// 系统盘类型。系统盘类型限制详见[存储概述](https://cloud.tencent.com/document/product/213/4952)。取值范围：<br>
	// <li>LOCAL_BASIC：本地硬盘</li>
	// <li>LOCAL_SSD：本地SSD硬盘</li>
	// <li>CLOUD_BASIC：普通云硬盘</li>
	// <li>CLOUD_SSD：SSD云硬盘</li>
	// <li>CLOUD_PREMIUM：高性能云硬盘</li>
	// <li>CLOUD_BSSD：通用型SSD云硬盘</li>
	// <li>CLOUD_HSSD：增强型SSD云硬盘</li>
	// <li>CLOUD_TSSD：极速型SSD云硬盘</li><br>
	// 默认取值：当前有库存的硬盘类型。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 系统盘ID。
	// 该参数目前仅用于`DescribeInstances`等查询类接口的返回参数，不可用于`RunInstances`等写接口的入参。
	DiskId *string `json:"DiskId,omitnil,omitempty" name:"DiskId"`

	// 系统盘大小，单位：GiB。默认值为 50
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 所属的独享集群ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// 磁盘名称，长度不超过128 个字符。
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TagSpecification struct {
	// 标签绑定的资源类型，云服务器为“instance”，专用宿主机为“host”，镜像为“image”，密钥为“keypair”，置放群组为“ps”，高性能计算集群为“hpc”。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 标签对列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type TargetOS struct {
	// 目标操作系统类型
	TargetOSType *string `json:"TargetOSType,omitnil,omitempty" name:"TargetOSType"`

	// 目标操作系统版本
	TargetOSVersion *string `json:"TargetOSVersion,omitnil,omitempty" name:"TargetOSVersion"`
}

// Predefined struct for user
type TerminateInstancesRequestParams struct {
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 释放弹性IP。EIP2.0下，仅提供主网卡下首个EIP，EIP类型限定在HighQualityEIP、AntiDDoSEIP、EIPv6、HighQualityEIPv6这几种类型。默认行为不释放。
	// 
	// 示例值：true
	// 默认值：false
	ReleaseAddress *bool `json:"ReleaseAddress,omitnil,omitempty" name:"ReleaseAddress"`

	// 释放实例挂载的包年包月数据盘。true表示销毁实例同时释放包年包月数据盘，false表示只销毁实例。
	// 默认值：false
	ReleasePrepaidDataDisks *bool `json:"ReleasePrepaidDataDisks,omitnil,omitempty" name:"ReleasePrepaidDataDisks"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 释放弹性IP。EIP2.0下，仅提供主网卡下首个EIP，EIP类型限定在HighQualityEIP、AntiDDoSEIP、EIPv6、HighQualityEIPv6这几种类型。默认行为不释放。
	// 
	// 示例值：true
	// 默认值：false
	ReleaseAddress *bool `json:"ReleaseAddress,omitnil,omitempty" name:"ReleaseAddress"`

	// 释放实例挂载的包年包月数据盘。true表示销毁实例同时释放包年包月数据盘，false表示只销毁实例。
	// 默认值：false
	ReleasePrepaidDataDisks *bool `json:"ReleasePrepaidDataDisks,omitnil,omitempty" name:"ReleasePrepaidDataDisks"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ReleaseAddress")
	delete(f, "ReleasePrepaidDataDisks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateInstancesResponseParams `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualPrivateCloud struct {
	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台](https://console.cloud.tencent.com/vpc/vpc?rid=1)查询；也可以调用接口 [DescribeVpcs](https://cloud.tencent.com/document/product/215/15778) ，从接口返回中的`VpcId `字段获取。若在创建子机时VpcId与SubnetId同时传入`DEFAULT`，则强制使用默认vpc网络。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台](https://console.cloud.tencent.com/vpc/subnet?rid=1)查询；也可以调用接口  [DescribeSubnets](https://cloud.tencent.com/document/product/215/15784) ，从接口返回中的`SubnetId `字段获取。若在创建子机时SubnetId与VpcId同时传入`DEFAULT`，则强制使用默认vpc网络。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：<li>true：表示用作公网网关</li><li>false：表示不作为公网网关</li>默认取值：false。
	AsVpcGateway *bool `json:"AsVpcGateway,omitnil,omitempty" name:"AsVpcGateway"`

	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 为弹性网卡指定随机生成的 IPv6 地址数量。
	Ipv6AddressCount *uint64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`
}

type ZoneInfo struct {
	// 可用区名称，例如，ap-guangzhou-3
	// 全网可用区名称如下：
	// <li> ap-chongqing-1 </li>
	// <li> ap-seoul-1 </li>
	// <li> ap-seoul-2 </li>
	// <li> ap-chengdu-1 </li>
	// <li> ap-chengdu-2 </li>
	// <li> ap-hongkong-1（售罄） </li>
	// <li> ap-hongkong-2 </li>
	// <li> ap-hongkong-3 </li>
	// <li> ap-shenzhen-fsi-1 </li>
	// <li> ap-shenzhen-fsi-2 </li>
	// <li> ap-shenzhen-fsi-3（售罄） </li>
	// <li> ap-guangzhou-1（售罄）</li>
	// <li> ap-guangzhou-2（售罄）</li>
	// <li> ap-guangzhou-3 </li>
	// <li> ap-guangzhou-4 </li>
	// <li> ap-guangzhou-6 </li>
	// <li> ap-guangzhou-7 </li>
	// <li> ap-tokyo-1 </li>
	// <li> ap-tokyo-2 </li>
	// <li> ap-singapore-1 </li>
	// <li> ap-singapore-2 </li>
	// <li> ap-singapore-3 </li>
	// <li>ap-singapore-4 </li>
	// <li> ap-shanghai-fsi-1 </li>
	// <li> ap-shanghai-fsi-2 </li>
	// <li> ap-shanghai-fsi-3 </li>
	// <li> ap-bangkok-1 </li>
	// <li> ap-bangkok-2 </li>
	// <li> ap-shanghai-1（售罄） </li>
	// <li> ap-shanghai-2 </li>
	// <li> ap-shanghai-3 </li>
	// <li> ap-shanghai-4 </li>
	// <li> ap-shanghai-5 </li>
	// <li> ap-shanghai-8 </li>
	// <li> ap-mumbai-1 </li>
	// <li> ap-mumbai-2 </li>
	// <li> ap-beijing-1（售罄）</li>
	// <li> ap-beijing-2 </li>
	// <li> ap-beijing-3 </li>
	// <li> ap-beijing-4 </li>
	// <li> ap-beijing-5 </li>
	// <li> ap-beijing-6 </li>
	// <li> ap-beijing-7 </li>
	// <li> na-siliconvalley-1 </li>
	// <li> na-siliconvalley-2 </li>
	// <li> eu-frankfurt-1 </li>
	// <li> eu-frankfurt-2 </li>
	// <li> na-ashburn-1 </li>
	// <li> na-ashburn-2 </li>
	// <li> ap-nanjing-1 </li>
	// <li> ap-nanjing-2 </li>
	// <li> ap-nanjing-3 </li>
	// <li> sa-saopaulo-1</li>
	// <li> ap-jakarta-1 </li>
	// <li> ap-jakarta-2 </li>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区描述，例如，广州三区
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区状态，包含AVAILABLE和UNAVAILABLE。AVAILABLE代表可用，UNAVAILABLE代表不可用。
	ZoneState *string `json:"ZoneState,omitnil,omitempty" name:"ZoneState"`
}