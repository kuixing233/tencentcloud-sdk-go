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

package v20221123

const (
	// 此产品的特有错误码

	// 方法不存在
	FAILEDOPERATION_INVALIDACTION = "FailedOperation.InvalidAction"

	// HostUuids参数和InstancesIds参数不能同时传递。
	INVALIDPARAMETER_HOSTUUIDSANDINSIDSCANNOTAPPEARSAMETIME = "InvalidParameter.HostUuidsAndInsIdsCannotAppearSameTime"

	// 该专属可用区不存在
	RESOURCENOTFOUND_CDZIDNOTFOUND = "ResourceNotFound.CdzIdNotFound"
)
