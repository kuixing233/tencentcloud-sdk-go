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

package v20240713

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 身份提供商名称已经使用。
	INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"

	// 身份提供商元数据文档错误。
	INVALIDPARAMETER_METADATAERROR = "InvalidParameter.MetadataError"

	// 参数错误
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// OIDC签名公钥错误。
	INVALIDPARAMETERVALUE_IDENTITYKEYERROR = "InvalidParameterValue.IdentityKeyError"

	// 身份提供商URL错误。
	INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"

	// 身份提供商名称错误。
	INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"

	// 身份提供商已达到上限。
	LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"

	// 身份提供商不存在。
	RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"

	// 没有数据
	RESOURCENOTFOUND_RECORDNOTEXISTS = "ResourceNotFound.RecordNotExists"
)
