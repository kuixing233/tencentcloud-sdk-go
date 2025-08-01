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

package v20201029

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type DBInstanceDetail struct {
	// DB实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// DB实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// DB实例状态,-1:已隔离, 0:创建中, 1:流程中, 2:运行中, 3:未初始化
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// DB实例状态描述,-1:已隔离, 0:创建中, 1:流程中, 2:运行中, 3:未初始化
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// DB实例版本
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Vip信息
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Vip使用的端口号
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 字符串型的私有网络ID
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// 字符串型的私有网络子网ID
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// 是否为分布式版本,0:否,1:是
	Shard *int64 `json:"Shard,omitnil,omitempty" name:"Shard"`

	// DB实例节点数
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// CPU规格(单位:核数)
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存规格(单位:GB)
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘规格(单位:GB)
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// 分布式类型的实例的分片数
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Db所在主机列表, 格式: m1,s1|m2,s2
	DbHosts *string `json:"DbHosts,omitnil,omitempty" name:"DbHosts"`

	// 主机角色, 1:主, 2:从, 3:主+从
	HostRole *int64 `json:"HostRole,omitnil,omitempty" name:"HostRole"`

	// DB引擎，MySQL,Percona,MariaDB
	DbEngine *string `json:"DbEngine,omitnil,omitempty" name:"DbEngine"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群主机Id
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 实例类型,0:mariadb, 1:tdsql
	ShardType []*int64 `json:"ShardType,omitnil,omitempty" name:"ShardType"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群主机Id
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 实例类型,0:mariadb, 1:tdsql
	ShardType []*int64 `json:"ShardType,omitnil,omitempty" name:"ShardType"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "HostId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ShardType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// 独享集群内的DB实例列表
	Instances []*DBInstanceDetail `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 独享集群内的DB实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostListRequestParams struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分配状态过滤，0-可分配，1-禁止分配
	AssignStatus []*int64 `json:"AssignStatus,omitnil,omitempty" name:"AssignStatus"`
}

type DescribeHostListRequest struct {
	*tchttp.BaseRequest
	
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分配状态过滤，0-可分配，1-禁止分配
	AssignStatus []*int64 `json:"AssignStatus,omitnil,omitempty" name:"AssignStatus"`
}

func (r *DescribeHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AssignStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostListResponseParams struct {
	// 主机总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主机详情
	Hosts []*HostDetail `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHostListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostListResponseParams `json:"Response"`
}

func (r *DescribeHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetail struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 产品ID, 0:CDB, 1:TDSQL
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 集群类型, 0:公有云, 1:金融围笼, 2:CDC集群
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 主机类型, 0:物理机, 1:CVM机型, 2:CDC机型
	HostType *int64 `json:"HostType,omitnil,omitempty" name:"HostType"`

	// 自动续费标志, 0:未设置, 1:自动续费, 2:到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 集群状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 主机数
	HostNum *int64 `json:"HostNum,omitnil,omitempty" name:"HostNum"`

	// DB实例数
	DbNum *int64 `json:"DbNum,omitnil,omitempty" name:"DbNum"`

	// 分配策略, 0:紧凑, 1:均匀
	AssignStrategy *int64 `json:"AssignStrategy,omitnil,omitempty" name:"AssignStrategy"`

	// 总主机CPU(单位:核数)
	CpuSpec *int64 `json:"CpuSpec,omitnil,omitempty" name:"CpuSpec"`

	// 总已分配CPU(单位:核数)
	CpuAssigned *int64 `json:"CpuAssigned,omitnil,omitempty" name:"CpuAssigned"`

	// 总可分配CPU(单位:核数)
	CpuAssignable *int64 `json:"CpuAssignable,omitnil,omitempty" name:"CpuAssignable"`

	// 总主机内存(单位:GB)
	MemorySpec *int64 `json:"MemorySpec,omitnil,omitempty" name:"MemorySpec"`

	// 总已分配内存(单位:GB)
	MemoryAssigned *int64 `json:"MemoryAssigned,omitnil,omitempty" name:"MemoryAssigned"`

	// 总可分配内存(单位:GB)
	MemoryAssignable *int64 `json:"MemoryAssignable,omitnil,omitempty" name:"MemoryAssignable"`

	// 总机器磁盘(单位:GB)
	DiskSpec *int64 `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// 总已分配磁盘(单位:GB)
	DiskAssigned *int64 `json:"DiskAssigned,omitnil,omitempty" name:"DiskAssigned"`

	// 总可分配磁盘(单位:GB)
	DiskAssignable *int64 `json:"DiskAssignable,omitnil,omitempty" name:"DiskAssignable"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 金融围笼ID
	FenceId *string `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 所属集群ID(默认集群为空)
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

// Predefined struct for user
type DescribeInstanceDetailRequestParams struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDetailResponseParams struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 产品ID, 0:CDB, 1:TDSQL
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 集群类型, 0:公有云, 1:金融围笼
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 主机类型, 0:物理机, 1:cvm本地盘, 2:cvm云盘
	HostType *int64 `json:"HostType,omitnil,omitempty" name:"HostType"`

	// 自动续费标志, 0:未设置, 1:自动续费, 2:到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 集群状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 主机数
	HostNum *int64 `json:"HostNum,omitnil,omitempty" name:"HostNum"`

	// Db实例数
	DbNum *int64 `json:"DbNum,omitnil,omitempty" name:"DbNum"`

	// 分配策略, 0:紧凑, 1:均匀
	AssignStrategy *int64 `json:"AssignStrategy,omitnil,omitempty" name:"AssignStrategy"`

	// 总主机CPU(单位:核)
	CpuSpec *int64 `json:"CpuSpec,omitnil,omitempty" name:"CpuSpec"`

	// 总已分配CPU(单位:核)
	CpuAssigned *int64 `json:"CpuAssigned,omitnil,omitempty" name:"CpuAssigned"`

	// 总可分配CPU(单位:核)
	CpuAssignable *int64 `json:"CpuAssignable,omitnil,omitempty" name:"CpuAssignable"`

	// 总主机内存(单位:GB)
	MemorySpec *int64 `json:"MemorySpec,omitnil,omitempty" name:"MemorySpec"`

	// 总已分配内存(单位:GB)
	MemoryAssigned *int64 `json:"MemoryAssigned,omitnil,omitempty" name:"MemoryAssigned"`

	// 总可分配内存(单位:GB)
	MemoryAssignable *int64 `json:"MemoryAssignable,omitnil,omitempty" name:"MemoryAssignable"`

	// 总机器磁盘(单位:GB)
	DiskSpec *int64 `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// 总已分配磁盘(单位:GB)
	DiskAssigned *int64 `json:"DiskAssigned,omitnil,omitempty" name:"DiskAssigned"`

	// 总可分配磁盘(单位:GB)
	DiskAssignable *int64 `json:"DiskAssignable,omitnil,omitempty" name:"DiskAssignable"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 金融围笼ID
	FenceId *string `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 所属集群ID(默认集群为空)
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListRequestParams struct {
	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，createTime,instancename两者之一
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序规则，desc,asc两者之一
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 按产品过滤，0:CDB, 1:TDSQL
	ProductId []*int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 按实例ID过滤
	InstanceId []*string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按实例名称过滤
	InstanceName []*string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 按金融围笼ID过滤
	FenceId []*string `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 按实例状态过滤, -1:已隔离, 0:创建中, 1:运行中, 2:扩容中, 3:删除中
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 按所属集群ID过滤
	ClusterId []*string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，createTime,instancename两者之一
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序规则，desc,asc两者之一
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 按产品过滤，0:CDB, 1:TDSQL
	ProductId []*int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 按实例ID过滤
	InstanceId []*string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按实例名称过滤
	InstanceName []*string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 按金融围笼ID过滤
	FenceId []*string `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 按实例状态过滤, -1:已隔离, 0:创建中, 1:运行中, 2:扩容中, 3:删除中
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 按所属集群ID过滤
	ClusterId []*string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "SortBy")
	delete(f, "ProductId")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "FenceId")
	delete(f, "Status")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListResponseParams struct {
	// 独享集群列表
	Instances []*DescribeInstanceDetail `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 独享集群实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceListResponseParams `json:"Response"`
}

func (r *DescribeInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 集群类型: 0 一主一备, 1 一主两备...N-1 一主N备
	InstanceTypes []*int64 `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 产品ID:  0 MYSQL，1 TDSQL
	ProductIds []*int64 `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`

	// 集群uuid: 如 dbdc-q810131s
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否按金融围笼标志搜索
	FenceFlag *bool `json:"FenceFlag,omitnil,omitempty" name:"FenceFlag"`

	// 按实例名字模糊匹配
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 每页数目, 整型
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码, 整型
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 排序字段，枚举：createtime,groupname
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式: asc升序, desc降序
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 集群状态: -2 已删除, -1 已隔离, 0 创建中, 1 运行中, 2 扩容中, 3 删除中
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群类型: 0 一主一备, 1 一主两备...N-1 一主N备
	InstanceTypes []*int64 `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 产品ID:  0 MYSQL，1 TDSQL
	ProductIds []*int64 `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`

	// 集群uuid: 如 dbdc-q810131s
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否按金融围笼标志搜索
	FenceFlag *bool `json:"FenceFlag,omitnil,omitempty" name:"FenceFlag"`

	// 按实例名字模糊匹配
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 每页数目, 整型
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码, 整型
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 排序字段，枚举：createtime,groupname
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式: asc升序, desc降序
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 集群状态: -2 已删除, -1 已隔离, 0 创建中, 1 运行中, 2 扩容中, 3 删除中
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
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
	delete(f, "InstanceTypes")
	delete(f, "ProductIds")
	delete(f, "InstanceIds")
	delete(f, "FenceFlag")
	delete(f, "InstanceName")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "InstanceStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 集群数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群扩展信息
	Instances []*InstanceExpand `json:"Instances,omitnil,omitempty" name:"Instances"`

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

type DeviceInfo struct {
	// 设备ID
	DeviceId *int64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备No
	DeviceNo *string `json:"DeviceNo,omitnil,omitempty" name:"DeviceNo"`

	// 设备类型
	DevClass *string `json:"DevClass,omitnil,omitempty" name:"DevClass"`

	// 设备总内存，单位GB
	MaxMemory *float64 `json:"MaxMemory,omitnil,omitempty" name:"MaxMemory"`

	// 设备总磁盘，单位GB
	MaxDisk *float64 `json:"MaxDisk,omitnil,omitempty" name:"MaxDisk"`

	// 设备剩余内存，单位GB
	RestMemory *float64 `json:"RestMemory,omitnil,omitempty" name:"RestMemory"`

	// 设备剩余磁盘，单位GB
	RestDisk *float64 `json:"RestDisk,omitnil,omitempty" name:"RestDisk"`

	// 设备机器个数
	RawDeviceNum *uint64 `json:"RawDeviceNum,omitnil,omitempty" name:"RawDeviceNum"`

	// 数据库实例个数
	InstanceNum *uint64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`
}

type HostDetail struct {
	// 主机Id
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 主机名称
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 主机状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 分配DB实例状态,0:可分配,1:不可分配
	AssignStatus *int64 `json:"AssignStatus,omitnil,omitempty" name:"AssignStatus"`

	// 主机类型, 0:物理机, 1:cvm本地盘, 2:cvm云盘
	HostType *int64 `json:"HostType,omitnil,omitempty" name:"HostType"`

	// DB实例数
	DbNum *int64 `json:"DbNum,omitnil,omitempty" name:"DbNum"`

	// 主机CPU(单位:核数)
	CpuSpec *int64 `json:"CpuSpec,omitnil,omitempty" name:"CpuSpec"`

	// 已分配CPU(单位:核数)
	CpuAssigned *int64 `json:"CpuAssigned,omitnil,omitempty" name:"CpuAssigned"`

	// 可分配CPU(单位:核数)
	CpuAssignable *int64 `json:"CpuAssignable,omitnil,omitempty" name:"CpuAssignable"`

	// 主机内存(单位:GB)
	MemorySpec *int64 `json:"MemorySpec,omitnil,omitempty" name:"MemorySpec"`

	// 已分配内存(单位:GB)
	MemoryAssigned *int64 `json:"MemoryAssigned,omitnil,omitempty" name:"MemoryAssigned"`

	// 可分配内存(单位:GB)
	MemoryAssignable *int64 `json:"MemoryAssignable,omitnil,omitempty" name:"MemoryAssignable"`

	// 主机磁盘(单位:GB)
	DiskSpec *int64 `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// 已分配磁盘(单位:GB)
	DiskAssigned *int64 `json:"DiskAssigned,omitnil,omitempty" name:"DiskAssigned"`

	// 可分配磁盘(GB)
	DiskAssignable *int64 `json:"DiskAssignable,omitnil,omitempty" name:"DiskAssignable"`

	// CPU分配比
	CpuRatio *float64 `json:"CpuRatio,omitnil,omitempty" name:"CpuRatio"`

	// 内存分配比
	MemoryRatio *float64 `json:"MemoryRatio,omitnil,omitempty" name:"MemoryRatio"`

	// 磁盘分配比
	DiskRatio *float64 `json:"DiskRatio,omitnil,omitempty" name:"DiskRatio"`

	// 机型名称
	MachineName *string `json:"MachineName,omitnil,omitempty" name:"MachineName"`

	// 机型类别
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// 计费标签
	PidTag *string `json:"PidTag,omitnil,omitempty" name:"PidTag"`

	// 计费ID
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type InstanceDetail struct {
	// 集群状态，0：运行中，1：不在运行
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 读写集群剩余内存容量，单位GB
	ReadWriteTotalLeaveMemory *float64 `json:"ReadWriteTotalLeaveMemory,omitnil,omitempty" name:"ReadWriteTotalLeaveMemory"`

	// 读写集群剩余磁盘容量，单位GB
	ReadWriteTotalLeaveDisk *float64 `json:"ReadWriteTotalLeaveDisk,omitnil,omitempty" name:"ReadWriteTotalLeaveDisk"`

	// 读写集群总内存容量，单位GB
	ReadWriteTotalMemory *float64 `json:"ReadWriteTotalMemory,omitnil,omitempty" name:"ReadWriteTotalMemory"`

	// 读写集群总磁盘容量，单位GB
	ReadWriteTotalDisk *float64 `json:"ReadWriteTotalDisk,omitnil,omitempty" name:"ReadWriteTotalDisk"`

	// 只读集群剩余内存容量，单位GB
	ReadOnlyTotalLeaveMemory *float64 `json:"ReadOnlyTotalLeaveMemory,omitnil,omitempty" name:"ReadOnlyTotalLeaveMemory"`

	// 只读集群剩余磁盘容量，单位GB
	ReadOnlyTotalLeaveDisk *float64 `json:"ReadOnlyTotalLeaveDisk,omitnil,omitempty" name:"ReadOnlyTotalLeaveDisk"`

	// 只读集群总内存容量，单位GB
	ReadOnlyTotalMemory *float64 `json:"ReadOnlyTotalMemory,omitnil,omitempty" name:"ReadOnlyTotalMemory"`

	// 只读集群总磁盘容量，单位GB
	ReadOnlyTotalDisk *float64 `json:"ReadOnlyTotalDisk,omitnil,omitempty" name:"ReadOnlyTotalDisk"`

	// 集群设备详情
	InstanceDeviceInfos []*InstanceDeviceInfo `json:"InstanceDeviceInfos,omitnil,omitempty" name:"InstanceDeviceInfos"`
}

type InstanceDeviceInfo struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 读写设备组
	ReadWriteDevice []*DeviceInfo `json:"ReadWriteDevice,omitnil,omitempty" name:"ReadWriteDevice"`

	// 只读设备组
	ReadOnlyDevice []*DeviceInfo `json:"ReadOnlyDevice,omitnil,omitempty" name:"ReadOnlyDevice"`

	// 空闲设备组
	FreeDevice []*DeviceInfo `json:"FreeDevice,omitnil,omitempty" name:"FreeDevice"`
}

type InstanceExpand struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用户ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群类型： 0：一主一备，1：一主两备
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 集群状态: 0 集群创建中, 1 集群有效, 2 集群扩容中, 3 集群删除中, 4 集群缩容中 -1 集群已隔离 -2 集群已删除
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例自动续费标识： 0正常续费 1自动续费 2到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 机型
	Machine *string `json:"Machine,omitnil,omitempty" name:"Machine"`

	// 过期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 集群信息
	InstanceDetail *InstanceDetail `json:"InstanceDetail,omitnil,omitempty" name:"InstanceDetail"`

	// 计费侧的产品ID
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`
}

// Predefined struct for user
type ModifyInstanceNameRequestParams struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}