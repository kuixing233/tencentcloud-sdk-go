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

type Activity struct {
	// 活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 计算节点ID
	ComputeNodeId *string `json:"ComputeNodeId,omitnil,omitempty" name:"ComputeNodeId"`

	// 计算节点活动类型，创建或者销毁
	ComputeNodeActivityType *string `json:"ComputeNodeActivityType,omitnil,omitempty" name:"ComputeNodeActivityType"`

	// 计算环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 起因
	Cause *string `json:"Cause,omitnil,omitempty" name:"Cause"`

	// 活动状态。取值范围：<br><li>SUBMITTED：已提交</li><li>PROCESSING：处理中</li><li>SUCCEED：成功</li><li>FAILED：失败</li>
	ActivityState *string `json:"ActivityState,omitnil,omitempty" name:"ActivityState"`

	// 状态原因
	StateReason *string `json:"StateReason,omitnil,omitempty" name:"StateReason"`

	// 活动开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 活动结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 云服务器实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type AgentRunningMode struct {
	// 场景类型，支持WINDOWS
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 运行Agent的User
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 运行Agent的Session
	Session *string `json:"Session,omitnil,omitempty" name:"Session"`
}

type AnonymousComputeEnv struct {
	// 计算环境管理类型
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 计算环境具体参数
	EnvData *EnvData `json:"EnvData,omitnil,omitempty" name:"EnvData"`

	// 数据盘挂载选项
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitnil,omitempty" name:"MountDataDisks"`

	// agent运行模式，适用于Windows系统
	AgentRunningMode *AgentRunningMode `json:"AgentRunningMode,omitnil,omitempty" name:"AgentRunningMode"`
}

type Application struct {
	// 应用程序的交付方式，包括PACKAGE、LOCAL 两种取值，分别指远程存储的软件包、计算环境本地。
	DeliveryForm *string `json:"DeliveryForm,omitnil,omitempty" name:"DeliveryForm"`

	// 松耦合任务执行命令。与Commands不能同时指定，一般使用Command字段提交任务。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 应用程序软件包的远程存储路径
	PackagePath *string `json:"PackagePath,omitnil,omitempty" name:"PackagePath"`

	// 应用使用Docker的相关配置。在使用Docker配置的情况下，DeliveryForm 为 LOCAL 表示直接使用Docker镜像内部的应用软件，通过Docker方式运行；DeliveryForm 为 PACKAGE，表示将远程应用包注入到Docker镜像后，通过Docker方式运行。为避免Docker不同版本的兼容性问题，Docker安装包及相关依赖由Batch统一负责，对于已安装Docker的自定义镜像，请卸载后再使用Docker特性。
	Docker *Docker `json:"Docker,omitnil,omitempty" name:"Docker"`

	// 紧耦合任务执行命令信息。与Command不能同时指定。Command和Commands必须指定一个。
	Commands []*CommandLine `json:"Commands,omitnil,omitempty" name:"Commands"`
}

// Predefined struct for user
type AttachInstancesRequestParams struct {
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 加入计算环境实例列表，每次请求的实例的上限为100。
	Instances []*Instance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type AttachInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 加入计算环境实例列表，每次请求的实例的上限为100。
	Instances []*Instance `json:"Instances,omitnil,omitempty" name:"Instances"`
}

func (r *AttachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AttachInstancesResponseParams `json:"Response"`
}

func (r *AttachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Authentication struct {
	// 授权场景，例如COS
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// SecretId
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// SecretKey
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`
}

type CommandLine struct {
	// 任务执行命令。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`
}

type ComputeEnvCreateInfo struct {
	// 计算环境 ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 计算环境名称
	EnvName *string `json:"EnvName,omitnil,omitempty" name:"EnvName"`

	// 计算环境描述
	EnvDescription *string `json:"EnvDescription,omitnil,omitempty" name:"EnvDescription"`

	// 计算环境类型，仅支持“MANAGED”类型
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 计算环境参数
	EnvData *EnvData `json:"EnvData,omitnil,omitempty" name:"EnvData"`

	// 数据盘挂载选项
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitnil,omitempty" name:"MountDataDisks"`

	// 输入映射
	InputMappings []*InputMapping `json:"InputMappings,omitnil,omitempty" name:"InputMappings"`

	// 授权信息
	Authentications []*Authentication `json:"Authentications,omitnil,omitempty" name:"Authentications"`

	// 通知信息
	Notifications []*Notification `json:"Notifications,omitnil,omitempty" name:"Notifications"`

	// 计算节点期望个数
	DesiredComputeNodeCount *uint64 `json:"DesiredComputeNodeCount,omitnil,omitempty" name:"DesiredComputeNodeCount"`

	// 计算环境标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ComputeEnvData struct {
	// CVM实例类型列表
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`
}

type ComputeEnvView struct {
	// 计算环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 计算环境名称
	EnvName *string `json:"EnvName,omitnil,omitempty" name:"EnvName"`

	// 位置信息
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 计算节点统计指标
	ComputeNodeMetrics *ComputeNodeMetrics `json:"ComputeNodeMetrics,omitnil,omitempty" name:"ComputeNodeMetrics"`

	// 计算环境类型
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 计算节点期望个数
	DesiredComputeNodeCount *uint64 `json:"DesiredComputeNodeCount,omitnil,omitempty" name:"DesiredComputeNodeCount"`

	// 计算环境资源类型，当前为CVM和CPM（黑石）
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 下一步动作
	NextAction *string `json:"NextAction,omitnil,omitempty" name:"NextAction"`

	// 用户添加到计算环境中的计算节点个数
	AttachedComputeNodeCount *uint64 `json:"AttachedComputeNodeCount,omitnil,omitempty" name:"AttachedComputeNodeCount"`

	// 计算环境绑定的标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ComputeNode struct {
	// 计算节点ID
	ComputeNodeId *string `json:"ComputeNodeId,omitnil,omitempty" name:"ComputeNodeId"`

	// 计算节点实例ID，对于CVM场景，即为CVM的InstanceId
	ComputeNodeInstanceId *string `json:"ComputeNodeInstanceId,omitnil,omitempty" name:"ComputeNodeInstanceId"`

	// 计算节点状态。取值范围：<br><li>PENDING：表示创建中</li><li>SUBMITTED：表示已提交创建</li><li>CREATING：表示创建中</li><li>CREATED：表示创建完成</li><li>CREATION_FAILED：表示创建失败。</li><li>RUNNING：表示运行中。</li><li>ABNORMAL：表示节点异常。</li><li>DELETING：表示删除中。</li>
	ComputeNodeState *string `json:"ComputeNodeState,omitnil,omitempty" name:"ComputeNodeState"`

	// CPU核数
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存容量，单位GiB
	Mem *uint64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 资源创建完成时间
	ResourceCreatedTime *string `json:"ResourceCreatedTime,omitnil,omitempty" name:"ResourceCreatedTime"`

	// 计算节点运行  TaskInstance 可用容量。0表示计算节点忙碌。
	TaskInstanceNumAvailable *uint64 `json:"TaskInstanceNumAvailable,omitnil,omitempty" name:"TaskInstanceNumAvailable"`

	// Batch Agent 版本
	AgentVersion *string `json:"AgentVersion,omitnil,omitempty" name:"AgentVersion"`

	// 实例内网IP
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 实例公网IP
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 计算环境资源类型，当前为CVM和CPM（黑石）
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 计算环境资源来源。<br>BATCH_CREATED：由批量计算创建的实例资源。<br>
	// USER_ATTACHED：用户添加到计算环境中的实例资源。
	ResourceOrigin *string `json:"ResourceOrigin,omitnil,omitempty" name:"ResourceOrigin"`
}

type ComputeNodeMetrics struct {
	// 已经完成提交的计算节点数量
	SubmittedCount *uint64 `json:"SubmittedCount,omitnil,omitempty" name:"SubmittedCount"`

	// 创建中的计算节点数量
	CreatingCount *uint64 `json:"CreatingCount,omitnil,omitempty" name:"CreatingCount"`

	// 创建失败的计算节点数量
	CreationFailedCount *uint64 `json:"CreationFailedCount,omitnil,omitempty" name:"CreationFailedCount"`

	// 完成创建的计算节点数量
	CreatedCount *uint64 `json:"CreatedCount,omitnil,omitempty" name:"CreatedCount"`

	// 运行中的计算节点数量
	RunningCount *uint64 `json:"RunningCount,omitnil,omitempty" name:"RunningCount"`

	// 销毁中的计算节点数量
	DeletingCount *uint64 `json:"DeletingCount,omitnil,omitempty" name:"DeletingCount"`

	// 异常的计算节点数量
	AbnormalCount *uint64 `json:"AbnormalCount,omitnil,omitempty" name:"AbnormalCount"`
}

// Predefined struct for user
type CreateComputeEnvRequestParams struct {
	// 计算环境信息
	ComputeEnv *NamedComputeEnv `json:"ComputeEnv,omitnil,omitempty" name:"ComputeEnv"`

	// 位置信息
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 用于保证请求幂等性的字符串。该字符串由用户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

type CreateComputeEnvRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境信息
	ComputeEnv *NamedComputeEnv `json:"ComputeEnv,omitnil,omitempty" name:"ComputeEnv"`

	// 位置信息
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 用于保证请求幂等性的字符串。该字符串由用户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

func (r *CreateComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateComputeEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ComputeEnv")
	delete(f, "Placement")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateComputeEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateComputeEnvResponseParams struct {
	// 计算环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *CreateComputeEnvResponseParams `json:"Response"`
}

func (r *CreateComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateComputeEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskTemplateRequestParams struct {
	// 任务模板名称，最大长度限制60个字符。
	TaskTemplateName *string `json:"TaskTemplateName,omitnil,omitempty" name:"TaskTemplateName"`

	// 任务模板内容，参数要求与任务一致
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitnil,omitempty" name:"TaskTemplateInfo"`

	// 任务模板描述，最大长度限制200个字符。
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitnil,omitempty" name:"TaskTemplateDescription"`

	// 标签列表。通过指定该参数可以支持绑定标签到任务模板。每个任务模板最多绑定10个标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateTaskTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 任务模板名称，最大长度限制60个字符。
	TaskTemplateName *string `json:"TaskTemplateName,omitnil,omitempty" name:"TaskTemplateName"`

	// 任务模板内容，参数要求与任务一致
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitnil,omitempty" name:"TaskTemplateInfo"`

	// 任务模板描述，最大长度限制200个字符。
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitnil,omitempty" name:"TaskTemplateDescription"`

	// 标签列表。通过指定该参数可以支持绑定标签到任务模板。每个任务模板最多绑定10个标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateTaskTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskTemplateName")
	delete(f, "TaskTemplateInfo")
	delete(f, "TaskTemplateDescription")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskTemplateResponseParams struct {
	// 任务模板ID
	TaskTemplateId *string `json:"TaskTemplateId,omitnil,omitempty" name:"TaskTemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskTemplateResponseParams `json:"Response"`
}

func (r *CreateTaskTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskTemplateResponse) FromJsonString(s string) error {
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

type DataPointView struct {
	// 监控数据采集的时间
	Timestamps []*int64 `json:"Timestamps,omitnil,omitempty" name:"Timestamps"`

	// 监控指标数据; 如果涉及到多个实例的监控数据的间隙时间，取值会为null
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type DeleteComputeEnvRequestParams struct {
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取，不能对状态处于删除中或者存在计算实例未销毁的环境发起删除动作。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DeleteComputeEnvRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取，不能对状态处于删除中或者存在计算实例未销毁的环境发起删除动作。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DeleteComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteComputeEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteComputeEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteComputeEnvResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *DeleteComputeEnvResponseParams `json:"Response"`
}

func (r *DeleteComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteComputeEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobRequestParams struct {
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DeleteJobRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DeleteJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteJobResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteJobResponse struct {
	*tchttp.BaseResponse
	Response *DeleteJobResponseParams `json:"Response"`
}

func (r *DeleteJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskTemplatesRequestParams struct {
	// 用于删除任务模板信息，最大数量上限100，环境模版ID通过调用接口 [DescribeTaskTemplates](https://cloud.tencent.com/document/api/599/15902)获取。
	TaskTemplateIds []*string `json:"TaskTemplateIds,omitnil,omitempty" name:"TaskTemplateIds"`
}

type DeleteTaskTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 用于删除任务模板信息，最大数量上限100，环境模版ID通过调用接口 [DescribeTaskTemplates](https://cloud.tencent.com/document/api/599/15902)获取。
	TaskTemplateIds []*string `json:"TaskTemplateIds,omitnil,omitempty" name:"TaskTemplateIds"`
}

func (r *DeleteTaskTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskTemplatesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTaskTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskTemplatesResponseParams `json:"Response"`
}

func (r *DeleteTaskTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dependence struct {
	// 依赖关系的起点任务名称
	StartTask *string `json:"StartTask,omitnil,omitempty" name:"StartTask"`

	// 依赖关系的终点任务名称
	EndTask *string `json:"EndTask,omitnil,omitempty" name:"EndTask"`
}

// Predefined struct for user
type DescribeAvailableCvmInstanceTypesRequestParams struct {
	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照[可用区](https://cloud.tencent.com/document/product/213/15707)过滤。</li>
	// <li> instance-family String - 是否必填：否 -（过滤条件）按照[机型系列](https://cloud.tencent.com/document/product/213/15748)过滤。实例机型系列形如：S1、I1、M1等。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAvailableCvmInstanceTypesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照[可用区](https://cloud.tencent.com/document/product/213/15707)过滤。</li>
	// <li> instance-family String - 是否必填：否 -（过滤条件）按照[机型系列](https://cloud.tencent.com/document/product/213/15748)过滤。实例机型系列形如：S1、I1、M1等。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAvailableCvmInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableCvmInstanceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableCvmInstanceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableCvmInstanceTypesResponseParams struct {
	// 机型配置数组
	InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitnil,omitempty" name:"InstanceTypeConfigSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAvailableCvmInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableCvmInstanceTypesResponseParams `json:"Response"`
}

func (r *DescribeAvailableCvmInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableCvmInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvActivitiesRequestParams struct {
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 偏移量，默认为0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认值20，最大值100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件<li> compute-node-id - String - 是否必填：否 -（过滤条件）按照计算节点ID过滤，节点ID通过调用接口 [DescribeComputeEnv](https://cloud.tencent.com/document/api/599/15892)获取。</li>
	Filters *Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeComputeEnvActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 偏移量，默认为0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认值20，最大值100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件<li> compute-node-id - String - 是否必填：否 -（过滤条件）按照计算节点ID过滤，节点ID通过调用接口 [DescribeComputeEnv](https://cloud.tencent.com/document/api/599/15892)获取。</li>
	Filters *Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeComputeEnvActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComputeEnvActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvActivitiesResponseParams struct {
	// 计算环境中的活动列表
	ActivitySet []*Activity `json:"ActivitySet,omitnil,omitempty" name:"ActivitySet"`

	// 活动数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeComputeEnvActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComputeEnvActivitiesResponseParams `json:"Response"`
}

func (r *DescribeComputeEnvActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvCreateInfoRequestParams struct {
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeComputeEnvCreateInfoRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeComputeEnvCreateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvCreateInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComputeEnvCreateInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvCreateInfoResponseParams struct {
	// 计算环境 ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 计算环境名称
	EnvName *string `json:"EnvName,omitnil,omitempty" name:"EnvName"`

	// 计算环境描述
	EnvDescription *string `json:"EnvDescription,omitnil,omitempty" name:"EnvDescription"`

	// 计算环境类型，仅支持“MANAGED”类型
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 计算环境参数
	EnvData *EnvData `json:"EnvData,omitnil,omitempty" name:"EnvData"`

	// 数据盘挂载选项
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitnil,omitempty" name:"MountDataDisks"`

	// 输入映射
	InputMappings []*InputMapping `json:"InputMappings,omitnil,omitempty" name:"InputMappings"`

	// 授权信息
	Authentications []*Authentication `json:"Authentications,omitnil,omitempty" name:"Authentications"`

	// 通知信息
	Notifications []*Notification `json:"Notifications,omitnil,omitempty" name:"Notifications"`

	// 计算节点期望个数
	DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitnil,omitempty" name:"DesiredComputeNodeCount"`

	// 计算环境绑定的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeComputeEnvCreateInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComputeEnvCreateInfoResponseParams `json:"Response"`
}

func (r *DescribeComputeEnvCreateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvCreateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvCreateInfosRequestParams struct {
	// 计算环境ID列表，与Filters参数不能同时指定，最大限制100，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvIds []*string `json:"EnvIds,omitnil,omitempty" name:"EnvIds"`

	// 过滤条件<li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤，可用区通过调用接口 [DescribeZones](https://cloud.tencent.com/document/api/213/15707)获取。</li><li> env-id - String - 是否必填：否 -（过滤条件）按照计算环境ID过滤。</li><li> env-name - String - 是否必填：否 -（过滤条件）按照计算环境名称过滤。</li>与EnvIds参数不能同时指定。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认值20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeComputeEnvCreateInfosRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID列表，与Filters参数不能同时指定，最大限制100，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvIds []*string `json:"EnvIds,omitnil,omitempty" name:"EnvIds"`

	// 过滤条件<li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤，可用区通过调用接口 [DescribeZones](https://cloud.tencent.com/document/api/213/15707)获取。</li><li> env-id - String - 是否必填：否 -（过滤条件）按照计算环境ID过滤。</li><li> env-name - String - 是否必填：否 -（过滤条件）按照计算环境名称过滤。</li>与EnvIds参数不能同时指定。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认值20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeComputeEnvCreateInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvCreateInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComputeEnvCreateInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvCreateInfosResponseParams struct {
	// 计算环境数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 计算环境创建信息列表
	ComputeEnvCreateInfoSet []*ComputeEnvCreateInfo `json:"ComputeEnvCreateInfoSet,omitnil,omitempty" name:"ComputeEnvCreateInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeComputeEnvCreateInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComputeEnvCreateInfosResponseParams `json:"Response"`
}

func (r *DescribeComputeEnvCreateInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvCreateInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvRequestParams struct {
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeComputeEnvRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComputeEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvResponseParams struct {
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnv](https://cloud.tencent.com/document/api/599/15892)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 计算环境名称
	EnvName *string `json:"EnvName,omitnil,omitempty" name:"EnvName"`

	// 位置信息
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 计算环境创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 计算节点列表信息
	ComputeNodeSet []*ComputeNode `json:"ComputeNodeSet,omitnil,omitempty" name:"ComputeNodeSet"`

	// 计算节点统计指标
	ComputeNodeMetrics *ComputeNodeMetrics `json:"ComputeNodeMetrics,omitnil,omitempty" name:"ComputeNodeMetrics"`

	// 计算节点期望个数
	DesiredComputeNodeCount *uint64 `json:"DesiredComputeNodeCount,omitnil,omitempty" name:"DesiredComputeNodeCount"`

	// 计算环境管理类型，枚举如下： MANAGED: 由客户在Batch平台主动创建； THPC_QUEUE: 由thpc平台创建，关联thpc平台集群队列。
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 计算环境资源类型，当前为CVM和CPM（黑石）
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 下一步的动作，枚举如下： DELETING: 删除中
	NextAction *string `json:"NextAction,omitnil,omitempty" name:"NextAction"`

	// 用户添加到计算环境中的计算节点个数
	AttachedComputeNodeCount *uint64 `json:"AttachedComputeNodeCount,omitnil,omitempty" name:"AttachedComputeNodeCount"`

	// 计算环境绑定的标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComputeEnvResponseParams `json:"Response"`
}

func (r *DescribeComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvsRequestParams struct {
	// 计算环境ID列表，与Filters参数不能同时指定。最大数量上限100。
	EnvIds []*string `json:"EnvIds,omitnil,omitempty" name:"EnvIds"`

	// 过滤条件<li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤，可用区通过调用接口 [DescribeZones](https://cloud.tencent.com/document/api/213/15707)获取。</li><li> env-id - String - 是否必填：否 -（过滤条件）按照计算环境ID过滤。</li><li> env-name - String - 是否必填：否 -（过滤条件）按照计算环境名称过滤。</li><li> resource-type - String - 是否必填：否 -（过滤条件）按照计算资源类型过滤，取值CVM或者CPM(黑石)。</li><li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li><li>tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li><li>tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。</li>与EnvIds参数不能同时指定。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认值20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeComputeEnvsRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID列表，与Filters参数不能同时指定。最大数量上限100。
	EnvIds []*string `json:"EnvIds,omitnil,omitempty" name:"EnvIds"`

	// 过滤条件<li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤，可用区通过调用接口 [DescribeZones](https://cloud.tencent.com/document/api/213/15707)获取。</li><li> env-id - String - 是否必填：否 -（过滤条件）按照计算环境ID过滤。</li><li> env-name - String - 是否必填：否 -（过滤条件）按照计算环境名称过滤。</li><li> resource-type - String - 是否必填：否 -（过滤条件）按照计算资源类型过滤，取值CVM或者CPM(黑石)。</li><li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li><li>tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li><li>tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。</li>与EnvIds参数不能同时指定。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认值20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeComputeEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComputeEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComputeEnvsResponseParams struct {
	// 计算环境列表
	ComputeEnvSet []*ComputeEnvView `json:"ComputeEnvSet,omitnil,omitempty" name:"ComputeEnvSet"`

	// 计算环境数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeComputeEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComputeEnvsResponseParams `json:"Response"`
}

func (r *DescribeComputeEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComputeEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmZoneInstanceConfigInfosRequestParams struct {
	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照[可用区](https://cloud.tencent.com/document/product/213/15707)过滤。</li>
	// <li> instance-family String - 是否必填：否 -（过滤条件）按照[机型系列](https://cloud.tencent.com/document/product/213/15748)过滤。实例机型系列形如：S1、I1、M1等。</li>
	// <li> instance-type - String - 是否必填：否 - （过滤条件）按照[机型](https://cloud.tencent.com/document/product/213/15749)过滤。实例机型形如：：S5.12XLARGE128、S5.12XLARGE96等。</li>
	// <li> instance-charge-type - String - 是否必填：否 -（过滤条件）按照实例计费模式过滤。 ( POSTPAID_BY_HOUR：表示后付费，即按量计费机型 | SPOTPAID：表示竞价付费机型。 )  </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCvmZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照[可用区](https://cloud.tencent.com/document/product/213/15707)过滤。</li>
	// <li> instance-family String - 是否必填：否 -（过滤条件）按照[机型系列](https://cloud.tencent.com/document/product/213/15748)过滤。实例机型系列形如：S1、I1、M1等。</li>
	// <li> instance-type - String - 是否必填：否 - （过滤条件）按照[机型](https://cloud.tencent.com/document/product/213/15749)过滤。实例机型形如：：S5.12XLARGE128、S5.12XLARGE96等。</li>
	// <li> instance-charge-type - String - 是否必填：否 -（过滤条件）按照实例计费模式过滤。 ( POSTPAID_BY_HOUR：表示后付费，即按量计费机型 | SPOTPAID：表示竞价付费机型。 )  </li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeCvmZoneInstanceConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCvmZoneInstanceConfigInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmZoneInstanceConfigInfosResponseParams struct {
	// 可用区机型配置列表。
	InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"InstanceTypeQuotaSet,omitnil,omitempty" name:"InstanceTypeQuotaSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCvmZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCvmZoneInstanceConfigInfosResponseParams `json:"Response"`
}

func (r *DescribeCvmZoneInstanceConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmZoneInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCategoriesRequestParams struct {

}

type DescribeInstanceCategoriesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInstanceCategoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCategoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceCategoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCategoriesResponseParams struct {
	// CVM实例分类列表
	InstanceCategorySet []*InstanceCategoryItem `json:"InstanceCategorySet,omitnil,omitempty" name:"InstanceCategorySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceCategoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceCategoriesResponseParams `json:"Response"`
}

func (r *DescribeInstanceCategoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCategoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobMonitorDataRequestParams struct {
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业的Task名称，详见[作业详情](https://cloud.tencent.com/document/product/599/15904)。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 作业任务实例的序号，详见[任务详情](https://cloud.tencent.com/document/product/599/15905)
	TaskInstanceIndex *int64 `json:"TaskInstanceIndex,omitnil,omitempty" name:"TaskInstanceIndex"`

	// 支持查询的指标；当前支持查询的任务指标；
	// 
	// - CpuUsage：cpu利用率，单位：%
	// - MemUsage：内存利用率，单位：%
	// - LanOuttraffic：内网出带宽，单位：Bytes/s
	// - LanIntraffic：内网入带宽，单位：Bytes/s
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 查询任务实例的起始时间；如果未传入查询起始时间或传入的时间小于任务实例的创建时间（任务实例创建时间详见[任务详情](https://cloud.tencent.com/document/product/599/15905)），会自动将查询时间调整到任务实例的创建时间。传入时间格式只支持零时区格式。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询任务实例的终止时间；如果未传入查询终止时间或传入的时间大于任务实例的终止时间（任务实例终止时间详见[任务详情](https://cloud.tencent.com/document/product/599/15905)），并且任务实例已经结束，会自动将查询终止时间调整到任务实例的终止时间；如果任务实例未结束，会自动将查询终止时间调整到当前时间。传入时间格式只支持零时区格式。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeJobMonitorDataRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业的Task名称，详见[作业详情](https://cloud.tencent.com/document/product/599/15904)。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 作业任务实例的序号，详见[任务详情](https://cloud.tencent.com/document/product/599/15905)
	TaskInstanceIndex *int64 `json:"TaskInstanceIndex,omitnil,omitempty" name:"TaskInstanceIndex"`

	// 支持查询的指标；当前支持查询的任务指标；
	// 
	// - CpuUsage：cpu利用率，单位：%
	// - MemUsage：内存利用率，单位：%
	// - LanOuttraffic：内网出带宽，单位：Bytes/s
	// - LanIntraffic：内网入带宽，单位：Bytes/s
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 查询任务实例的起始时间；如果未传入查询起始时间或传入的时间小于任务实例的创建时间（任务实例创建时间详见[任务详情](https://cloud.tencent.com/document/product/599/15905)），会自动将查询时间调整到任务实例的创建时间。传入时间格式只支持零时区格式。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询任务实例的终止时间；如果未传入查询终止时间或传入的时间大于任务实例的终止时间（任务实例终止时间详见[任务详情](https://cloud.tencent.com/document/product/599/15905)），并且任务实例已经结束，会自动将查询终止时间调整到任务实例的终止时间；如果任务实例未结束，会自动将查询终止时间调整到当前时间。传入时间格式只支持零时区格式。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeJobMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobMonitorDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "TaskName")
	delete(f, "TaskInstanceIndex")
	delete(f, "MetricName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobMonitorDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobMonitorDataResponseParams struct {
	// 监控数据粒度，单位:秒；时间粒度随着查询的时间范围变化，查询时间范围越小，时间粒度越小。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 监控采集的数据。时间戳和对应的值一一对应；如果查询的任务重试，采集时间段涉及多个实例的话，某些时间段内的值为null, 表示对应时间点没有实例存在，也不存在对应的监控数据；相邻监控时间段之间的空值数量最多为10。
	DataPoints *DataPointView `json:"DataPoints,omitnil,omitempty" name:"DataPoints"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobMonitorDataResponseParams `json:"Response"`
}

func (r *DescribeJobMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobRequestParams struct {
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeJobRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobResponseParams struct {
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业名称
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 可用区信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 作业优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 作业状态
	JobState *string `json:"JobState,omitnil,omitempty" name:"JobState"`

	// 创建时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 结束时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务视图信息
	TaskSet []*TaskView `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 任务间依赖信息
	DependenceSet []*Dependence `json:"DependenceSet,omitnil,omitempty" name:"DependenceSet"`

	// 任务统计指标
	TaskMetrics *TaskMetrics `json:"TaskMetrics,omitnil,omitempty" name:"TaskMetrics"`

	// 任务实例统计指标
	TaskInstanceMetrics *TaskInstanceMetrics `json:"TaskInstanceMetrics,omitnil,omitempty" name:"TaskInstanceMetrics"`

	// 作业失败原因
	StateReason *string `json:"StateReason,omitnil,omitempty" name:"StateReason"`

	// 作业绑定的标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 下一步动作
	NextAction *string `json:"NextAction,omitnil,omitempty" name:"NextAction"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobResponseParams `json:"Response"`
}

func (r *DescribeJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobSubmitInfoRequestParams struct {
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeJobSubmitInfoRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeJobSubmitInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobSubmitInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobSubmitInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobSubmitInfoResponseParams struct {
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业名称
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 作业描述
	JobDescription *string `json:"JobDescription,omitnil,omitempty" name:"JobDescription"`

	// 作业优先级，任务（Task）和任务实例（TaskInstance）会继承作业优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 任务信息
	Tasks []*Task `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 依赖信息
	Dependences []*Dependence `json:"Dependences,omitnil,omitempty" name:"Dependences"`

	// 作业绑定的标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobSubmitInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobSubmitInfoResponseParams `json:"Response"`
}

func (r *DescribeJobSubmitInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobSubmitInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobsRequestParams struct {
	// 作业ID列表，与Filters参数不能同时指定。
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// 过滤条件
	// <li> job-id - String - 是否必填：否 -（过滤条件）按照作业ID过滤。</li>
	// <li> job-name - String - 是否必填：否 -（过滤条件）按照作业名称过滤。</li>
	// <li> job-state - String - 是否必填：否 -（过滤条件）按照作业状态过滤。</li>
	// 
	//     - SUBMITTED：已提交；
	//     - PENDING：等待中；
	//     - RUNNABLE：可运行；
	//     - STARTING：启动中；
	//     - RUNNING：运行中；
	//     - SUCCEED：成功；
	//     - FAILED：失败；
	//     - FAILED_INTERRUPTED：失败后保留实例。
	// 
	// <li> zone - String - 是否必填：否 -（过滤条件）按照[可用区](https://cloud.tencent.com/document/product/213/15707)过滤。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。</li>
	// 与JobIds参数不能同时指定。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回job数量限制，最大值: 100，默认值: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeJobsRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID列表，与Filters参数不能同时指定。
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`

	// 过滤条件
	// <li> job-id - String - 是否必填：否 -（过滤条件）按照作业ID过滤。</li>
	// <li> job-name - String - 是否必填：否 -（过滤条件）按照作业名称过滤。</li>
	// <li> job-state - String - 是否必填：否 -（过滤条件）按照作业状态过滤。</li>
	// 
	//     - SUBMITTED：已提交；
	//     - PENDING：等待中；
	//     - RUNNABLE：可运行；
	//     - STARTING：启动中；
	//     - RUNNING：运行中；
	//     - SUCCEED：成功；
	//     - FAILED：失败；
	//     - FAILED_INTERRUPTED：失败后保留实例。
	// 
	// <li> zone - String - 是否必填：否 -（过滤条件）按照[可用区](https://cloud.tencent.com/document/product/213/15707)过滤。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。</li>
	// 与JobIds参数不能同时指定。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回job数量限制，最大值: 100，默认值: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobsResponseParams struct {
	// 作业列表
	JobSet []*JobView `json:"JobSet,omitnil,omitempty" name:"JobSet"`

	// 符合条件的作业数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobsResponseParams `json:"Response"`
}

func (r *DescribeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogsRequestParams struct {
	// 作业ID。JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务实例集合；与Offset不能同时指定。
	TaskInstanceIndexes []*uint64 `json:"TaskInstanceIndexes,omitnil,omitempty" name:"TaskInstanceIndexes"`

	// 起始任务实例。与TaskInstanceIndexes参数不能同时指定。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大任务实例数；默认值为5， 最大值为10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTaskLogsRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID。JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务实例集合；与Offset不能同时指定。
	TaskInstanceIndexes []*uint64 `json:"TaskInstanceIndexes,omitnil,omitempty" name:"TaskInstanceIndexes"`

	// 起始任务实例。与TaskInstanceIndexes参数不能同时指定。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大任务实例数；默认值为5， 最大值为10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTaskLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "TaskName")
	delete(f, "TaskInstanceIndexes")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogsResponseParams struct {
	// 任务实例总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务实例日志详情集合
	TaskInstanceLogSet []*TaskInstanceLog `json:"TaskInstanceLogSet,omitnil,omitempty" name:"TaskInstanceLogSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLogsResponseParams `json:"Response"`
}

func (r *DescribeTaskLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRequestParams struct {
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量。默认取值100，最大取值1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，详情如下：
	// task-instance-state     - String - 是否必填： 否 - 按照任务实例状态进行过滤（
	// 
	// - SUBMITTED：已提交；
	// - PENDING：等待中；
	// - RUNNABLE：可运行；
	// - STARTING：启动中；
	// - RUNNING：运行中；
	// - SUCCEED：成功；
	// - FAILED：失败；
	// - FAILED_INTERRUPTED：失败后保留实例）。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量。默认取值100，最大取值1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，详情如下：
	// task-instance-state     - String - 是否必填： 否 - 按照任务实例状态进行过滤（
	// 
	// - SUBMITTED：已提交；
	// - PENDING：等待中；
	// - RUNNABLE：可运行；
	// - STARTING：启动中；
	// - RUNNING：运行中；
	// - SUCCEED：成功；
	// - FAILED：失败；
	// - FAILED_INTERRUPTED：失败后保留实例）。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "TaskName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResponseParams struct {
	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务状态
	TaskState *string `json:"TaskState,omitnil,omitempty" name:"TaskState"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务实例总数
	TaskInstanceTotalCount *int64 `json:"TaskInstanceTotalCount,omitnil,omitempty" name:"TaskInstanceTotalCount"`

	// 任务实例信息
	TaskInstanceSet []*TaskInstanceView `json:"TaskInstanceSet,omitnil,omitempty" name:"TaskInstanceSet"`

	// 任务实例统计指标
	TaskInstanceMetrics *TaskInstanceMetrics `json:"TaskInstanceMetrics,omitnil,omitempty" name:"TaskInstanceMetrics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResponseParams `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskTemplatesRequestParams struct {
	// 任务模板ID列表，与Filters参数不能同时指定。
	TaskTemplateIds []*string `json:"TaskTemplateIds,omitnil,omitempty" name:"TaskTemplateIds"`

	// 过滤条件
	// <li> task-template-name - String - 是否必填：否 -（过滤条件）按照任务模板名称过滤。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。</li>
	// 与TaskTemplateIds参数不能同时指定。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTaskTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 任务模板ID列表，与Filters参数不能同时指定。
	TaskTemplateIds []*string `json:"TaskTemplateIds,omitnil,omitempty" name:"TaskTemplateIds"`

	// 过滤条件
	// <li> task-template-name - String - 是否必填：否 -（过滤条件）按照任务模板名称过滤。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。</li>
	// 与TaskTemplateIds参数不能同时指定。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTaskTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskTemplateIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskTemplatesResponseParams struct {
	// 任务模板列表
	TaskTemplateSet []*TaskTemplateView `json:"TaskTemplateSet,omitnil,omitempty" name:"TaskTemplateSet"`

	// 任务模板数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskTemplatesResponseParams `json:"Response"`
}

func (r *DescribeTaskTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachInstancesRequestParams struct {
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 实例ID列表，实例ID通过调用接口 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728)获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DetachInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 实例ID列表，实例ID通过调用接口 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728)获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DetachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DetachInstancesResponseParams `json:"Response"`
}

func (r *DetachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Docker struct {
	// Docker Hub填写“[user/repo]:[tag]”，Tencent Registry填写“ccr.ccs.tencentyun.com/[namespace/repo]:[tag]”
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Docker Hub 用户名或 Tencent Registry 用户名；公共镜像可不填写此参数。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Docker Hub 密码或 Tencent Registry 密码；公共镜像可不填写此参数。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Docker Hub 可以不填，但确保具有公网访问能力。或者是 Tencent Registry 服务地址“ccr.ccs.tencentyun.com”
	Server *string `json:"Server,omitnil,omitempty" name:"Server"`

	// 拉取Docker镜像重试次数。默认值：0。
	MaxRetryCount *uint64 `json:"MaxRetryCount,omitnil,omitempty" name:"MaxRetryCount"`

	// 拉取Docker镜像失败时延迟时间。单位：秒。
	DelayOnRetry *uint64 `json:"DelayOnRetry,omitnil,omitempty" name:"DelayOnRetry"`

	// Docker命令运行参数。
	DockerRunOption *string `json:"DockerRunOption,omitnil,omitempty" name:"DockerRunOption"`
}

type EnhancedService struct {
	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitnil,omitempty" name:"SecurityService"`

	// 开启云监控服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitnil,omitempty" name:"MonitorService"`

	// 开启云自动化助手服务（TencentCloud Automation Tools，TAT）。若不指定该参数，则公共镜像默认开启云自动化助手服务，其他镜像默认不开启云自动化助手服务。
	AutomationService *RunAutomationServiceEnabled `json:"AutomationService,omitnil,omitempty" name:"AutomationService"`
}

type EnvData struct {
	// CVM实例类型，不能与InstanceTypes和InstanceTypeOptions同时出现。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// CVM镜像ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例系统盘配置信息
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 私有网络相关信息配置，与Zones和VirtualPrivateClouds不能同时指定。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil,omitempty" name:"VirtualPrivateCloud"`

	// 公网带宽相关信息设置
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// CVM实例显示名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例登录设置
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// CVM实例计费类型<br><li>POSTPAID_BY_HOUR：按小时后付费</li><li>SPOTPAID：竞价付费</li><br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例的市场相关选项，如竞价实例相关参数
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// CVM实例类型列表，不能与InstanceType和InstanceTypeOptions同时出现。指定该字段后，计算节点按照机型先后顺序依次尝试创建，直到实例创建成功，结束遍历过程。最多支持10个机型。
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// CVM实例机型配置。不能与InstanceType和InstanceTypes同时出现。
	InstanceTypeOptions *InstanceTypeOptions `json:"InstanceTypeOptions,omitnil,omitempty" name:"InstanceTypeOptions"`

	// 可用区列表，支持跨可用区创建CVM实例。与VirtualPrivateCloud和VirtualPrivateClouds不能同时指定。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 私有网络列表，支持跨私有网络创建CVM实例。与VirtualPrivateCloud和Zones不能同时指定。
	VirtualPrivateClouds []*VirtualPrivateCloud `json:"VirtualPrivateClouds,omitnil,omitempty" name:"VirtualPrivateClouds"`
}

type EnvVar struct {
	// 环境变量名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境变量取值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type EventConfig struct {
	// 事件类型，包括：<br/><li>“JOB_RUNNING”：作业运行，适用于"SubmitJob"。</li><li>“JOB_SUCCEED”：作业成功，适用于"SubmitJob"。</li><li>“JOB_FAILED”：作业失败，适用于"SubmitJob"。</li><li>“JOB_FAILED_INTERRUPTED”：作业失败，保留实例，适用于"SubmitJob"。</li><li>“TASK_RUNNING”：任务运行，适用于"SubmitJob"。</li><li>“TASK_SUCCEED”：任务成功，适用于"SubmitJob"。</li><li>“TASK_FAILED”：任务失败，适用于"SubmitJob"。</li><li>“TASK_FAILED_INTERRUPTED”：任务失败，保留实例，适用于"SubmitJob"。</li><li>“TASK_INSTANCE_RUNNING”：任务实例运行，适用于"SubmitJob"。</li><li>“TASK_INSTANCE_SUCCEED”：任务实例成功，适用于"SubmitJob"。</li><li>“TASK_INSTANCE_FAILED”：任务实例失败，适用于"SubmitJob"。</li><li>“TASK_INSTANCE_FAILED_INTERRUPTED”：任务实例失败，保留实例，适用于"SubmitJob"。</li><li>“COMPUTE_ENV_CREATED”：计算环境已创建，适用于"CreateComputeEnv"。</li><li>“COMPUTE_ENV_DELETED”：计算环境已删除，适用于"CreateComputeEnv"。</li><li>“COMPUTE_NODE_CREATED”：计算节点已创建，适用于"CreateComputeEnv"和"SubmitJob"。</li><li>“COMPUTE_NODE_CREATION_FAILED”：计算节点创建失败，适用于"CreateComputeEnv"和"SubmitJob"。</li><li>“COMPUTE_NODE_RUNNING”：计算节点运行中，适用于"CreateComputeEnv"和"SubmitJob"。</li><li>“COMPUTE_NODE_ABNORMAL”：计算节点异常，适用于"CreateComputeEnv"和"SubmitJob"。</li><li>“COMPUTE_NODE_DELETING”：计算节点已删除，适用于"CreateComputeEnv"和"SubmitJob"。</li>
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 自定义键值对
	EventVars []*EventVar `json:"EventVars,omitnil,omitempty" name:"EventVars"`
}

type EventVar struct {
	// 自定义键
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
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

type InputMapping struct {
	// 源端路径
	SourcePath *string `json:"SourcePath,omitnil,omitempty" name:"SourcePath"`

	// 目的端路径
	DestinationPath *string `json:"DestinationPath,omitnil,omitempty" name:"DestinationPath"`

	// 挂载配置项参数
	MountOptionParameter *string `json:"MountOptionParameter,omitnil,omitempty" name:"MountOptionParameter"`
}

type Instance struct {
	// 实例ID，可通过调用接口[DescribeInstances](https://cloud.tencent.com/document/product/213/15728)获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 镜像ID，可通过调用接口[DescribeImages](https://cloud.tencent.com/document/product/213/15715)获取。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`
}

type InstanceCategoryItem struct {
	// 实例类型名
	InstanceCategory *string `json:"InstanceCategory,omitnil,omitempty" name:"InstanceCategory"`

	// 实例族列表
	InstanceFamilySet []*string `json:"InstanceFamilySet,omitnil,omitempty" name:"InstanceFamilySet"`
}

type InstanceMarketOptionsRequest struct {
	// 竞价相关选项
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitnil,omitempty" name:"SpotOptions"`

	// 市场选项类型，当前只支持取值：spot
	MarketType *string `json:"MarketType,omitnil,omitempty" name:"MarketType"`
}

type InstanceTypeConfig struct {
	// 内存容量，单位：`GB`。
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// CPU核数，单位：核。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例机型系列。
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`
}

type InstanceTypeOptions struct {
	// CPU核数。
	CPU *uint64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 内存值，单位GB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例机型类别，可选参数：“ALL”、“GENERAL”、“GENERAL_2”、“GENERAL_3”、“COMPUTE”、“COMPUTE_2”和“COMPUTE_3”。默认值“ALL”。
	InstanceCategories []*string `json:"InstanceCategories,omitnil,omitempty" name:"InstanceCategories"`
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

type Job struct {
	// 任务信息
	Tasks []*Task `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 作业名称; 字符串长度限制60.
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 作业描述；字符串长度限制200.
	JobDescription *string `json:"JobDescription,omitnil,omitempty" name:"JobDescription"`

	// 作业优先级，任务（Task）和任务实例（TaskInstance）会继承作业优先级；范围0～100，数值越大，优先级越高。
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 依赖信息
	Dependences []*Dependence `json:"Dependences,omitnil,omitempty" name:"Dependences"`

	// 通知信息
	Notifications []*Notification `json:"Notifications,omitnil,omitempty" name:"Notifications"`

	// 对于存在依赖关系的任务中，后序任务执行对于前序任务的依赖条件。取值范围包括 PRE_TASK_SUCCEED，PRE_TASK_AT_LEAST_PARTLY_SUCCEED，PRE_TASK_FINISHED，默认值为PRE_TASK_SUCCEED。
	TaskExecutionDependOn *string `json:"TaskExecutionDependOn,omitnil,omitempty" name:"TaskExecutionDependOn"`

	// 表示创建 CVM 失败按照何种策略处理。取值范围包括 FAILED，RUNNABLE。FAILED 表示创建 CVM 失败按照一次执行失败处理，RUNNABLE 表示创建 CVM 失败按照继续等待处理。默认值为FAILED。StateIfCreateCvmFailed对于提交的指定计算环境的作业无效。
	StateIfCreateCvmFailed *string `json:"StateIfCreateCvmFailed,omitnil,omitempty" name:"StateIfCreateCvmFailed"`

	// 标签列表。通过指定该参数可以支持绑定标签到作业。每个作业最多绑定10个标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 表示通知信息的通知目标类型。
	// 取值范围：CMQ，TDMQ_CMQ。
	// CMQ:表示向腾讯云CMQ发送消息。
	// TDMQ_CMQ：表示向腾讯云TDMQ_CMQ发送消息。<br/>默认值为CMQ。<br/>注：腾讯云计划于2022年6月前正式下线消息队列 CMQ，建议使用TDMQ_CMQ。参考文档：[CMQ迁移到TDMQ_CMQ](https://cloud.tencent.com/document/product/406/60860)
	NotificationTarget *string `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`
}

type JobView struct {
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 作业名称
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 作业状态:
	// - SUBMITTED：已提交；
	// - PENDING：等待中；
	// - RUNNABLE：可运行；
	// - STARTING：启动中；
	// - RUNNING：运行中；
	// - SUCCEED：成功；
	// - FAILED：失败；
	// - FAILED_INTERRUPTED：失败后保留实例。
	JobState *string `json:"JobState,omitnil,omitempty" name:"JobState"`

	// 作业优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 位置信息
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 创建时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 结束时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务统计指标
	TaskMetrics *TaskMetrics `json:"TaskMetrics,omitnil,omitempty" name:"TaskMetrics"`

	// 作业绑定的标签列表。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? \/ ]中的特殊符号。</li> 
	// <li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? \/]中的特殊符号。</li><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<li>TRUE：表示保持镜像的登录设置</li><li>FALSE：表示不保持镜像的登录设置</li>默认取值：FALSE。
	KeepImageLogin *string `json:"KeepImageLogin,omitnil,omitempty" name:"KeepImageLogin"`
}

// Predefined struct for user
type ModifyComputeEnvRequestParams struct {
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 计算节点期望个数，最大上限2000。
	DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitnil,omitempty" name:"DesiredComputeNodeCount"`

	// 计算环境名称
	EnvName *string `json:"EnvName,omitnil,omitempty" name:"EnvName"`

	// 计算环境描述
	EnvDescription *string `json:"EnvDescription,omitnil,omitempty" name:"EnvDescription"`

	// 计算环境属性数据
	EnvData *ComputeEnvData `json:"EnvData,omitnil,omitempty" name:"EnvData"`
}

type ModifyComputeEnvRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 计算节点期望个数，最大上限2000。
	DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitnil,omitempty" name:"DesiredComputeNodeCount"`

	// 计算环境名称
	EnvName *string `json:"EnvName,omitnil,omitempty" name:"EnvName"`

	// 计算环境描述
	EnvDescription *string `json:"EnvDescription,omitnil,omitempty" name:"EnvDescription"`

	// 计算环境属性数据
	EnvData *ComputeEnvData `json:"EnvData,omitnil,omitempty" name:"EnvData"`
}

func (r *ModifyComputeEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyComputeEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "DesiredComputeNodeCount")
	delete(f, "EnvName")
	delete(f, "EnvDescription")
	delete(f, "EnvData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyComputeEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyComputeEnvResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyComputeEnvResponse struct {
	*tchttp.BaseResponse
	Response *ModifyComputeEnvResponseParams `json:"Response"`
}

func (r *ModifyComputeEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyComputeEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskTemplateRequestParams struct {
	// 任务模板ID; 详见[任务模版](https://cloud.tencent.com/document/product/599/15902)。
	TaskTemplateId *string `json:"TaskTemplateId,omitnil,omitempty" name:"TaskTemplateId"`

	// 任务模板名称；字节长度限制60。
	TaskTemplateName *string `json:"TaskTemplateName,omitnil,omitempty" name:"TaskTemplateName"`

	// 任务模板描述；字节长度限制200。
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitnil,omitempty" name:"TaskTemplateDescription"`

	// 任务模板信息
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitnil,omitempty" name:"TaskTemplateInfo"`
}

type ModifyTaskTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 任务模板ID; 详见[任务模版](https://cloud.tencent.com/document/product/599/15902)。
	TaskTemplateId *string `json:"TaskTemplateId,omitnil,omitempty" name:"TaskTemplateId"`

	// 任务模板名称；字节长度限制60。
	TaskTemplateName *string `json:"TaskTemplateName,omitnil,omitempty" name:"TaskTemplateName"`

	// 任务模板描述；字节长度限制200。
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitnil,omitempty" name:"TaskTemplateDescription"`

	// 任务模板信息
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitnil,omitempty" name:"TaskTemplateInfo"`
}

func (r *ModifyTaskTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskTemplateId")
	delete(f, "TaskTemplateName")
	delete(f, "TaskTemplateDescription")
	delete(f, "TaskTemplateInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTaskTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskTemplateResponseParams `json:"Response"`
}

func (r *ModifyTaskTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountDataDisk struct {
	// 挂载点，Linux系统合法路径，或Windows系统盘符,比如"H:\\"
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`

	// 文件系统类型，Linux系统下支持"EXT3"和"EXT4"两种，默认"EXT3"；Windows系统下仅支持"NTFS"
	FileSystemType *string `json:"FileSystemType,omitnil,omitempty" name:"FileSystemType"`
}

type NamedComputeEnv struct {
	// 计算环境名称
	EnvName *string `json:"EnvName,omitnil,omitempty" name:"EnvName"`

	// 计算节点期望个数，最大上限2000.
	DesiredComputeNodeCount *int64 `json:"DesiredComputeNodeCount,omitnil,omitempty" name:"DesiredComputeNodeCount"`

	// 计算环境描述
	EnvDescription *string `json:"EnvDescription,omitnil,omitempty" name:"EnvDescription"`

	// 计算环境管理类型，枚举如下：
	// MANAGED: 由客户在Batch平台主动创建；
	// THPC_QUEUE: 由THPC平台创建，关联THPC平台的集群队列。
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 计算环境具体参数
	EnvData *EnvData `json:"EnvData,omitnil,omitempty" name:"EnvData"`

	// 数据盘挂载选项
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitnil,omitempty" name:"MountDataDisks"`

	// 授权信息
	Authentications []*Authentication `json:"Authentications,omitnil,omitempty" name:"Authentications"`

	// 输入映射信息
	InputMappings []*InputMapping `json:"InputMappings,omitnil,omitempty" name:"InputMappings"`

	// agent运行模式，适用于Windows系统
	AgentRunningMode *AgentRunningMode `json:"AgentRunningMode,omitnil,omitempty" name:"AgentRunningMode"`

	// 通知信息
	Notifications []*Notification `json:"Notifications,omitnil,omitempty" name:"Notifications"`

	// 非活跃节点处理策略，默认“RECREATE”，即对于实例创建失败或异常退还的计算节点，定期重新创建实例资源。
	ActionIfComputeNodeInactive *string `json:"ActionIfComputeNodeInactive,omitnil,omitempty" name:"ActionIfComputeNodeInactive"`

	// 对于实例创建失败或异常退还的计算节点，定期重新创建实例资源的最大重试次数，最大值100，如果不设置的话，系统会设置一个默认值，当前为7
	ResourceMaxRetryCount *int64 `json:"ResourceMaxRetryCount,omitnil,omitempty" name:"ResourceMaxRetryCount"`

	// 标签列表。通过指定该参数可以支持绑定标签到计算环境。每个计算环境最多绑定10个标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 表示通知信息的通知目标类型。
	// 取值范围：CMQ，TDMQ_CMQ。
	// CMQ:表示向腾讯云CMQ发送消息。
	// TDMQ_CMQ：表示向腾讯云TDMQ_CMQ发送消息。<br/>默认值为CMQ。<br/>注：腾讯云计划于2022年6月前正式下线消息队列 CMQ，建议使用TDMQ_CMQ。参考文档：[CMQ迁移到TDMQ_CMQ](https://cloud.tencent.com/document/product/406/60860)
	NotificationTarget *string `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`
}

type Notification struct {
	// CMQ主题名字，要求主题名有效且关联订阅
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 事件配置
	EventConfigs []*EventConfig `json:"EventConfigs,omitnil,omitempty" name:"EventConfigs"`
}

type OutputMapping struct {
	// 源端路径
	SourcePath *string `json:"SourcePath,omitnil,omitempty" name:"SourcePath"`

	// 目的端路径
	DestinationPath *string `json:"DestinationPath,omitnil,omitempty" name:"DestinationPath"`

	// 输出映射选项
	OutputMappingOption *OutputMappingOption `json:"OutputMappingOption,omitnil,omitempty" name:"OutputMappingOption"`
}

type OutputMappingConfig struct {
	// 存储类型，仅支持COS
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 并行worker数量
	WorkerNum *int64 `json:"WorkerNum,omitnil,omitempty" name:"WorkerNum"`

	// worker分块大小，单位MB
	WorkerPartSize *int64 `json:"WorkerPartSize,omitnil,omitempty" name:"WorkerPartSize"`
}

type OutputMappingOption struct {
	// 容器场景下,输出选项从实例映射到容器内的实例侧的工作空间。
	// BATCH_WORKSPACE: 工作空间为BATCH在实例内定义的工作空间，BATCH侧保证作业之间的隔离。（默认）
	// GLOBAL_WORKSPACE: 工作空间为实例操作系统空间。
	Workspace *string `json:"Workspace,omitnil,omitempty" name:"Workspace"`
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

type RedirectInfo struct {
	// 标准输出重定向路径; 
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitnil,omitempty" name:"StdoutRedirectPath"`

	// 标准错误重定向路径
	StderrRedirectPath *string `json:"StderrRedirectPath,omitnil,omitempty" name:"StderrRedirectPath"`

	// 标准输出重定向文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StdoutRedirectFileName *string `json:"StdoutRedirectFileName,omitnil,omitempty" name:"StdoutRedirectFileName"`

	// 标准错误重定向文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StderrRedirectFileName *string `json:"StderrRedirectFileName,omitnil,omitempty" name:"StderrRedirectFileName"`
}

type RedirectLocalInfo struct {
	// 标准输出重定向本地路径
	StdoutLocalPath *string `json:"StdoutLocalPath,omitnil,omitempty" name:"StdoutLocalPath"`

	// 标准错误重定向本地路径
	StderrLocalPath *string `json:"StderrLocalPath,omitnil,omitempty" name:"StderrLocalPath"`

	// 标准输出重定向本地文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StdoutLocalFileName *string `json:"StdoutLocalFileName,omitnil,omitempty" name:"StdoutLocalFileName"`

	// 标准错误重定向本地文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StderrLocalFileName *string `json:"StderrLocalFileName,omitnil,omitempty" name:"StderrLocalFileName"`
}

// Predefined struct for user
type RetryJobsRequestParams struct {
	// 作业ID列表。最大重试作业数100；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)。
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`
}

type RetryJobsRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID列表。最大重试作业数100；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)。
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`
}

func (r *RetryJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryJobsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryJobsResponse struct {
	*tchttp.BaseResponse
	Response *RetryJobsResponseParams `json:"Response"`
}

func (r *RetryJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunAutomationServiceEnabled struct {
	// 是否开启云自动化助手。取值范围：<br><li>true：表示开启云自动化助手服务<br><li>false：表示不开启云自动化助手服务<br><br>默认取值：false。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RunMonitorServiceEnabled struct {
	// 是否开启[云监控](/document/product/248)服务。取值范围：<br><li>true：表示开启云监控服务</li><li>false：表示不开启云监控服务</li><br>默认取值：true。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// 是否开启[云安全](/document/product/296)服务。取值范围：<br><li>true：表示开启云安全服务<br><li>false：表示不开启云安全服务<br><br>默认取值：true。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type SpotMarketOptions struct {
	// 竞价出价
	MaxPrice *string `json:"MaxPrice,omitnil,omitempty" name:"MaxPrice"`

	// 竞价请求类型，当前仅支持类型：one-time
	SpotInstanceType *string `json:"SpotInstanceType,omitnil,omitempty" name:"SpotInstanceType"`
}

type StorageBlock struct {
	// HDD本地存储类型，值为：LOCAL_PRO.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// HDD本地存储的最小容量。单位：GiB。
	MinSize *int64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// HDD本地存储的最大容量。单位：GiB。
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

// Predefined struct for user
type SubmitJobRequestParams struct {
	// 作业所提交的位置信息。通过该参数可以指定作业关联CVM所属可用区等信息。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 作业信息
	Job *Job `json:"Job,omitnil,omitempty" name:"Job"`

	// 用于保证请求幂等性的字符串。该字符串由用户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

type SubmitJobRequest struct {
	*tchttp.BaseRequest
	
	// 作业所提交的位置信息。通过该参数可以指定作业关联CVM所属可用区等信息。
	Placement *Placement `json:"Placement,omitnil,omitempty" name:"Placement"`

	// 作业信息
	Job *Job `json:"Job,omitnil,omitempty" name:"Job"`

	// 用于保证请求幂等性的字符串。该字符串由用户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

func (r *SubmitJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "Job")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitJobResponseParams struct {
	// 当通过本接口来提交作业时会返回该参数，表示一个作业ID。返回作业ID列表并不代表作业解析/运行成功，可根据 DescribeJob 接口查询其状态。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitJobResponseParams `json:"Response"`
}

func (r *SubmitJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitJobResponse) FromJsonString(s string) error {
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
	// 标签键。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Task struct {
	// 应用程序信息
	Application *Application `json:"Application,omitnil,omitempty" name:"Application"`

	// 任务名称，在一个作业内部唯一
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务实例运行个数，默认为1
	TaskInstanceNum *uint64 `json:"TaskInstanceNum,omitnil,omitempty" name:"TaskInstanceNum"`

	// 运行环境信息，ComputeEnv 和 EnvId 必须指定一个（且只有一个）参数。
	ComputeEnv *AnonymousComputeEnv `json:"ComputeEnv,omitnil,omitempty" name:"ComputeEnv"`

	// 计算环境ID，ComputeEnv 和 EnvId 必须指定一个（且只有一个）参数。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 重定向信息
	RedirectInfo *RedirectInfo `json:"RedirectInfo,omitnil,omitempty" name:"RedirectInfo"`

	// 重定向本地信息
	RedirectLocalInfo *RedirectLocalInfo `json:"RedirectLocalInfo,omitnil,omitempty" name:"RedirectLocalInfo"`

	// 输入映射
	InputMappings []*InputMapping `json:"InputMappings,omitnil,omitempty" name:"InputMappings"`

	// 输出映射
	OutputMappings []*OutputMapping `json:"OutputMappings,omitnil,omitempty" name:"OutputMappings"`

	// 输出映射配置
	OutputMappingConfigs []*OutputMappingConfig `json:"OutputMappingConfigs,omitnil,omitempty" name:"OutputMappingConfigs"`

	// 自定义环境变量
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// 授权信息
	Authentications []*Authentication `json:"Authentications,omitnil,omitempty" name:"Authentications"`

	// TaskInstance失败后处理方式，取值包括
	// 
	// - TERMINATE：销毁计算实例（默认）、
	// - INTERRUPT：中断任务，保留计算实例、
	// - FAST_INTERRUPT： 快速中断任务， 保留计算实例。
	FailedAction *string `json:"FailedAction,omitnil,omitempty" name:"FailedAction"`

	// 任务失败后的最大重试次数，默认为0
	MaxRetryCount *uint64 `json:"MaxRetryCount,omitnil,omitempty" name:"MaxRetryCount"`

	// 任务启动后的超时时间，单位秒，默认为86400秒
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 任务最大并发数限制，默认没有限制。
	MaxConcurrentNum *uint64 `json:"MaxConcurrentNum,omitnil,omitempty" name:"MaxConcurrentNum"`

	// 任务完成后，重启计算节点。适用于指定计算环境执行任务。
	RestartComputeNode *bool `json:"RestartComputeNode,omitnil,omitempty" name:"RestartComputeNode"`

	// 启动任务过程中，创建计算资源如CVM失败后的最大重试次数，默认为0。最大值100。
	// 计算资源创建重试的等待时间间隔策略设置如下：
	// [1, 3]: 等待600 s发起重试；
	// [4, 10]: 等待900 s发起重试；
	// [11, 50]: 等待1800 s发起重试；
	// [51, 100]: 等待3600 s发起重试；
	// [a, b]表示重试次数区间，每次重试的等待时间随着重试次数的增加而递增。
	// 例如，计算资源创建重试8次的耗时为：3*600 + 5*900 = 6300 s
	ResourceMaxRetryCount *uint64 `json:"ResourceMaxRetryCount,omitnil,omitempty" name:"ResourceMaxRetryCount"`
}

type TaskInstanceLog struct {
	// 任务实例
	TaskInstanceIndex *uint64 `json:"TaskInstanceIndex,omitnil,omitempty" name:"TaskInstanceIndex"`

	// 标准输出日志（Base64编码，解码后最大日志长度2048字节）
	StdoutLog *string `json:"StdoutLog,omitnil,omitempty" name:"StdoutLog"`

	// 标准错误日志（Base64编码，解码后最大日志长度2048字节）
	StderrLog *string `json:"StderrLog,omitnil,omitempty" name:"StderrLog"`

	// 标准输出重定向路径
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitnil,omitempty" name:"StdoutRedirectPath"`

	// 标准错误重定向路径
	StderrRedirectPath *string `json:"StderrRedirectPath,omitnil,omitempty" name:"StderrRedirectPath"`

	// 标准输出重定向文件名
	StdoutRedirectFileName *string `json:"StdoutRedirectFileName,omitnil,omitempty" name:"StdoutRedirectFileName"`

	// 标准错误重定向文件名
	StderrRedirectFileName *string `json:"StderrRedirectFileName,omitnil,omitempty" name:"StderrRedirectFileName"`
}

type TaskInstanceMetrics struct {
	// Submitted个数
	SubmittedCount *int64 `json:"SubmittedCount,omitnil,omitempty" name:"SubmittedCount"`

	// Pending个数
	PendingCount *int64 `json:"PendingCount,omitnil,omitempty" name:"PendingCount"`

	// Runnable个数
	RunnableCount *int64 `json:"RunnableCount,omitnil,omitempty" name:"RunnableCount"`

	// Starting个数
	StartingCount *int64 `json:"StartingCount,omitnil,omitempty" name:"StartingCount"`

	// Running个数
	RunningCount *int64 `json:"RunningCount,omitnil,omitempty" name:"RunningCount"`

	// Succeed个数
	SucceedCount *int64 `json:"SucceedCount,omitnil,omitempty" name:"SucceedCount"`

	// FailedInterrupted个数
	FailedInterruptedCount *int64 `json:"FailedInterruptedCount,omitnil,omitempty" name:"FailedInterruptedCount"`

	// Failed个数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`
}

type TaskInstanceView struct {
	// 任务实例索引
	TaskInstanceIndex *int64 `json:"TaskInstanceIndex,omitnil,omitempty" name:"TaskInstanceIndex"`

	// 任务实例状态: 
	// - PENDING：等待中；
	// - RUNNABLE：可运行；
	// - STARTING：启动中；
	// - RUNNING：运行中；
	// - SUCCEED：成功；
	// - FAILED：失败；
	// - FAILED_INTERRUPTED：失败后保留实例。
	TaskInstanceState *string `json:"TaskInstanceState,omitnil,omitempty" name:"TaskInstanceState"`

	// 应用程序执行结束的exit code
	ExitCode *int64 `json:"ExitCode,omitnil,omitempty" name:"ExitCode"`

	// 任务实例状态原因，任务实例失败时，会记录失败原因
	StateReason *string `json:"StateReason,omitnil,omitempty" name:"StateReason"`

	// 任务实例运行时所在计算节点（例如CVM）的InstanceId。任务实例未运行或者完结时，本字段为空。任务实例重试时，本字段会随之变化
	ComputeNodeInstanceId *string `json:"ComputeNodeInstanceId,omitnil,omitempty" name:"ComputeNodeInstanceId"`

	// 创建时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 启动时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	LaunchTime *string `json:"LaunchTime,omitnil,omitempty" name:"LaunchTime"`

	// 开始运行时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	RunningTime *string `json:"RunningTime,omitnil,omitempty" name:"RunningTime"`

	// 结束时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 重定向信息
	RedirectInfo *RedirectInfo `json:"RedirectInfo,omitnil,omitempty" name:"RedirectInfo"`

	// 任务实例状态原因详情，任务实例失败时，会记录失败原因
	StateDetailedReason *string `json:"StateDetailedReason,omitnil,omitempty" name:"StateDetailedReason"`
}

type TaskMetrics struct {
	// Submitted个数
	SubmittedCount *int64 `json:"SubmittedCount,omitnil,omitempty" name:"SubmittedCount"`

	// Pending个数
	PendingCount *int64 `json:"PendingCount,omitnil,omitempty" name:"PendingCount"`

	// Runnable个数
	RunnableCount *int64 `json:"RunnableCount,omitnil,omitempty" name:"RunnableCount"`

	// Starting个数
	StartingCount *int64 `json:"StartingCount,omitnil,omitempty" name:"StartingCount"`

	// Running个数
	RunningCount *int64 `json:"RunningCount,omitnil,omitempty" name:"RunningCount"`

	// Succeed个数
	SucceedCount *int64 `json:"SucceedCount,omitnil,omitempty" name:"SucceedCount"`

	// FailedInterrupted个数
	FailedInterruptedCount *int64 `json:"FailedInterruptedCount,omitnil,omitempty" name:"FailedInterruptedCount"`

	// Failed个数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`
}

type TaskTemplateView struct {
	// 任务模板ID
	TaskTemplateId *string `json:"TaskTemplateId,omitnil,omitempty" name:"TaskTemplateId"`

	// 任务模板名称
	TaskTemplateName *string `json:"TaskTemplateName,omitnil,omitempty" name:"TaskTemplateName"`

	// 任务模板描述
	TaskTemplateDescription *string `json:"TaskTemplateDescription,omitnil,omitempty" name:"TaskTemplateDescription"`

	// 任务模板信息
	TaskTemplateInfo *Task `json:"TaskTemplateInfo,omitnil,omitempty" name:"TaskTemplateInfo"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务模板绑定的标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type TaskView struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务状态:
	// - PENDING：等待中；
	// - RUNNABLE：可运行；
	// - STARTING：启动中；
	// - RUNNING：运行中；
	// - SUCCEED：成功；
	// - FAILED：失败；
	// - FAILED_INTERRUPTED：失败后保留实例。
	TaskState *string `json:"TaskState,omitnil,omitempty" name:"TaskState"`

	// 开始时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 结束时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type TerminateComputeNodeRequestParams struct {
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 计算节点ID，节点ID通过调用接口 [DescribeComputeEnv](https://cloud.tencent.com/document/api/599/15892)获取。
	ComputeNodeId *string `json:"ComputeNodeId,omitnil,omitempty" name:"ComputeNodeId"`
}

type TerminateComputeNodeRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnvs](https://cloud.tencent.com/document/api/599/15893)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 计算节点ID，节点ID通过调用接口 [DescribeComputeEnv](https://cloud.tencent.com/document/api/599/15892)获取。
	ComputeNodeId *string `json:"ComputeNodeId,omitnil,omitempty" name:"ComputeNodeId"`
}

func (r *TerminateComputeNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateComputeNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ComputeNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateComputeNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateComputeNodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateComputeNodeResponse struct {
	*tchttp.BaseResponse
	Response *TerminateComputeNodeResponseParams `json:"Response"`
}

func (r *TerminateComputeNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateComputeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateComputeNodesRequestParams struct {
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnv](https://cloud.tencent.com/document/api/599/15892)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 计算节点ID列表，最大数量上限100，节点ID通过调用接口 [DescribeComputeEnv](https://cloud.tencent.com/document/api/599/15892)获取。
	ComputeNodeIds []*string `json:"ComputeNodeIds,omitnil,omitempty" name:"ComputeNodeIds"`
}

type TerminateComputeNodesRequest struct {
	*tchttp.BaseRequest
	
	// 计算环境ID，环境ID通过调用接口 [DescribeComputeEnv](https://cloud.tencent.com/document/api/599/15892)获取。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 计算节点ID列表，最大数量上限100，节点ID通过调用接口 [DescribeComputeEnv](https://cloud.tencent.com/document/api/599/15892)获取。
	ComputeNodeIds []*string `json:"ComputeNodeIds,omitnil,omitempty" name:"ComputeNodeIds"`
}

func (r *TerminateComputeNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateComputeNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ComputeNodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateComputeNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateComputeNodesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateComputeNodesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateComputeNodesResponseParams `json:"Response"`
}

func (r *TerminateComputeNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateComputeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateJobRequestParams struct {
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type TerminateJobRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID；JobId详见[作业列表](https://cloud.tencent.com/document/product/599/15909)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *TerminateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateJobResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateJobResponse struct {
	*tchttp.BaseResponse
	Response *TerminateJobResponseParams `json:"Response"`
}

func (r *TerminateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTaskInstanceRequestParams struct {
	// 作业ID；详见[作业列表](https://cloud.tencent.com/document/product/599/15909)。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务名称；详见[作业提交信息](https://cloud.tencent.com/document/product/599/15910)
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务实例索引
	TaskInstanceIndex *int64 `json:"TaskInstanceIndex,omitnil,omitempty" name:"TaskInstanceIndex"`
}

type TerminateTaskInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 作业ID；详见[作业列表](https://cloud.tencent.com/document/product/599/15909)。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务名称；详见[作业提交信息](https://cloud.tencent.com/document/product/599/15910)
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务实例索引
	TaskInstanceIndex *int64 `json:"TaskInstanceIndex,omitnil,omitempty" name:"TaskInstanceIndex"`
}

func (r *TerminateTaskInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTaskInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "TaskName")
	delete(f, "TaskInstanceIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateTaskInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTaskInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateTaskInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateTaskInstanceResponseParams `json:"Response"`
}

func (r *TerminateTaskInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTaskInstanceResponse) FromJsonString(s string) error {
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