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

package v20190118

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AlgorithmInfo struct {
	// 算法的标识
	KeyUsage *string `json:"KeyUsage,omitnil,omitempty" name:"KeyUsage"`

	// 算法的名称
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`
}

// Predefined struct for user
type ArchiveKeyRequestParams struct {
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type ArchiveKeyRequest struct {
	*tchttp.BaseRequest
	
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *ArchiveKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ArchiveKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ArchiveKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ArchiveKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ArchiveKeyResponse struct {
	*tchttp.BaseResponse
	Response *ArchiveKeyResponseParams `json:"Response"`
}

func (r *ArchiveKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ArchiveKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AsymmetricRsaDecryptRequestParams struct {
	// CMK的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 使用PublicKey加密的密文，Base64编码
	Ciphertext *string `json:"Ciphertext,omitnil,omitempty" name:"Ciphertext"`

	// 在使用公钥加密时对应的算法：当前支持RSAES_PKCS1_V1_5、RSAES_OAEP_SHA_1、RSAES_OAEP_SHA_256
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`
}

type AsymmetricRsaDecryptRequest struct {
	*tchttp.BaseRequest
	
	// CMK的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 使用PublicKey加密的密文，Base64编码
	Ciphertext *string `json:"Ciphertext,omitnil,omitempty" name:"Ciphertext"`

	// 在使用公钥加密时对应的算法：当前支持RSAES_PKCS1_V1_5、RSAES_OAEP_SHA_1、RSAES_OAEP_SHA_256
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`
}

func (r *AsymmetricRsaDecryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AsymmetricRsaDecryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "Ciphertext")
	delete(f, "Algorithm")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AsymmetricRsaDecryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AsymmetricRsaDecryptResponseParams struct {
	// CMK的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 解密后的明文，base64编码
	Plaintext *string `json:"Plaintext,omitnil,omitempty" name:"Plaintext"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AsymmetricRsaDecryptResponse struct {
	*tchttp.BaseResponse
	Response *AsymmetricRsaDecryptResponseParams `json:"Response"`
}

func (r *AsymmetricRsaDecryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AsymmetricRsaDecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AsymmetricSm2DecryptRequestParams struct {
	// CMK的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 使用PublicKey加密的密文，Base64编码，原始密文格式需要为C1C3C2_ASN1。原始密文长度不能超过256字节。
	Ciphertext *string `json:"Ciphertext,omitnil,omitempty" name:"Ciphertext"`
}

type AsymmetricSm2DecryptRequest struct {
	*tchttp.BaseRequest
	
	// CMK的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 使用PublicKey加密的密文，Base64编码，原始密文格式需要为C1C3C2_ASN1。原始密文长度不能超过256字节。
	Ciphertext *string `json:"Ciphertext,omitnil,omitempty" name:"Ciphertext"`
}

func (r *AsymmetricSm2DecryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AsymmetricSm2DecryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "Ciphertext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AsymmetricSm2DecryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AsymmetricSm2DecryptResponseParams struct {
	// CMK的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 解密后的明文，base64编码
	Plaintext *string `json:"Plaintext,omitnil,omitempty" name:"Plaintext"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AsymmetricSm2DecryptResponse struct {
	*tchttp.BaseResponse
	Response *AsymmetricSm2DecryptResponseParams `json:"Response"`
}

func (r *AsymmetricSm2DecryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AsymmetricSm2DecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudResourceRequestParams struct {
	// cmk的ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 云产品的唯一性标识符
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 资源/实例ID，由调用方根据自己的云产品特征来定义，以字符串形式做存储。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type BindCloudResourceRequest struct {
	*tchttp.BaseRequest
	
	// cmk的ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 云产品的唯一性标识符
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 资源/实例ID，由调用方根据自己的云产品特征来定义，以字符串形式做存储。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *BindCloudResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "ProductId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindCloudResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudResourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindCloudResourceResponse struct {
	*tchttp.BaseResponse
	Response *BindCloudResourceResponseParams `json:"Response"`
}

func (r *BindCloudResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelDataKeyDeletionRequestParams struct {
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

type CancelDataKeyDeletionRequest struct {
	*tchttp.BaseRequest
	
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

func (r *CancelDataKeyDeletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDataKeyDeletionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelDataKeyDeletionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelDataKeyDeletionResponseParams struct {
	// 唯一标志被计划删除的数据密钥
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelDataKeyDeletionResponse struct {
	*tchttp.BaseResponse
	Response *CancelDataKeyDeletionResponseParams `json:"Response"`
}

func (r *CancelDataKeyDeletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDataKeyDeletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKeyArchiveRequestParams struct {
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type CancelKeyArchiveRequest struct {
	*tchttp.BaseRequest
	
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *CancelKeyArchiveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKeyArchiveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelKeyArchiveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKeyArchiveResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelKeyArchiveResponse struct {
	*tchttp.BaseResponse
	Response *CancelKeyArchiveResponseParams `json:"Response"`
}

func (r *CancelKeyArchiveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKeyArchiveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKeyDeletionRequestParams struct {
	// 需要被取消删除的CMK的唯一标志
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type CancelKeyDeletionRequest struct {
	*tchttp.BaseRequest
	
	// 需要被取消删除的CMK的唯一标志
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *CancelKeyDeletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKeyDeletionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelKeyDeletionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKeyDeletionResponseParams struct {
	// 唯一标志被取消删除的CMK。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelKeyDeletionResponse struct {
	*tchttp.BaseResponse
	Response *CancelKeyDeletionResponseParams `json:"Response"`
}

func (r *CancelKeyDeletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKeyDeletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyRequestParams struct {
	// 作为密钥更容易辨识，更容易被人看懂的别名， 不可为空，1-60个字母数字 - _ 的组合，首字符必须为字母或者数字。以 kms- 作为前缀的用于云产品使用，Alias 不可重复。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// CMK 的描述，最大1024字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 指定key的用途，默认为  "ENCRYPT_DECRYPT" 表示创建对称加解密密钥，其它支持用途 “ASYMMETRIC_DECRYPT_RSA_2048” 表示创建用于加解密的RSA2048非对称密钥，“ASYMMETRIC_DECRYPT_SM2” 表示创建用于加解密的SM2非对称密钥，“ASYMMETRIC_SIGN_VERIFY_SM2” 表示创建用于签名验签的SM2非对称密钥，“ASYMMETRIC_SIGN_VERIFY_ECC” 表示创建用于签名验签的ECC非对称密钥，“ASYMMETRIC_SIGN_VERIFY_RSA_2048” 表示创建用于签名验签的RSA_2048非对称密钥，“ASYMMETRIC_SIGN_VERIFY_ECDSA384”表示创建用于签名验签的 ECDSA384 非对称密钥。完整的密钥用途与算法支持列表可通过 ListAlgorithms 接口获取。
	KeyUsage *string `json:"KeyUsage,omitnil,omitempty" name:"KeyUsage"`

	// 指定key类型，默认为1，1表示默认类型，由KMS创建CMK密钥，2 表示EXTERNAL 类型，该类型需要用户导入密钥材料，参考 GetParametersForImport 和 ImportKeyMaterial 接口
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// KMS 高级版对应的 HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

type CreateKeyRequest struct {
	*tchttp.BaseRequest
	
	// 作为密钥更容易辨识，更容易被人看懂的别名， 不可为空，1-60个字母数字 - _ 的组合，首字符必须为字母或者数字。以 kms- 作为前缀的用于云产品使用，Alias 不可重复。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// CMK 的描述，最大1024字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 指定key的用途，默认为  "ENCRYPT_DECRYPT" 表示创建对称加解密密钥，其它支持用途 “ASYMMETRIC_DECRYPT_RSA_2048” 表示创建用于加解密的RSA2048非对称密钥，“ASYMMETRIC_DECRYPT_SM2” 表示创建用于加解密的SM2非对称密钥，“ASYMMETRIC_SIGN_VERIFY_SM2” 表示创建用于签名验签的SM2非对称密钥，“ASYMMETRIC_SIGN_VERIFY_ECC” 表示创建用于签名验签的ECC非对称密钥，“ASYMMETRIC_SIGN_VERIFY_RSA_2048” 表示创建用于签名验签的RSA_2048非对称密钥，“ASYMMETRIC_SIGN_VERIFY_ECDSA384”表示创建用于签名验签的 ECDSA384 非对称密钥。完整的密钥用途与算法支持列表可通过 ListAlgorithms 接口获取。
	KeyUsage *string `json:"KeyUsage,omitnil,omitempty" name:"KeyUsage"`

	// 指定key类型，默认为1，1表示默认类型，由KMS创建CMK密钥，2 表示EXTERNAL 类型，该类型需要用户导入密钥材料，参考 GetParametersForImport 和 ImportKeyMaterial 接口
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// KMS 高级版对应的 HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

func (r *CreateKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alias")
	delete(f, "Description")
	delete(f, "KeyUsage")
	delete(f, "Type")
	delete(f, "Tags")
	delete(f, "HsmClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyResponseParams struct {
	// CMK的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 作为密钥更容易辨识，更容易被人看懂的别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 密钥创建时间，unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// CMK的描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// CMK的状态
	KeyState *string `json:"KeyState,omitnil,omitempty" name:"KeyState"`

	// CMK的用途
	KeyUsage *string `json:"KeyUsage,omitnil,omitempty" name:"KeyUsage"`

	// 标签操作的返回码. 0: 成功；1: 内部错误；2: 业务处理错误
	TagCode *uint64 `json:"TagCode,omitnil,omitempty" name:"TagCode"`

	// 标签操作的返回信息
	TagMsg *string `json:"TagMsg,omitnil,omitempty" name:"TagMsg"`

	// HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateKeyResponseParams `json:"Response"`
}

func (r *CreateKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWhiteBoxKeyRequestParams struct {
	// 作为密钥更容易辨识，更容易被人看懂的别名， 不可为空，1-60个字母数字 - _ 的组合，首字符必须为字母或者数字。Alias不可重复。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 创建密钥所有的算法类型，支持的取值：AES_256,SM4
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 密钥的描述，最大1024字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest
	
	// 作为密钥更容易辨识，更容易被人看懂的别名， 不可为空，1-60个字母数字 - _ 的组合，首字符必须为字母或者数字。Alias不可重复。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 创建密钥所有的算法类型，支持的取值：AES_256,SM4
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 密钥的描述，最大1024字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateWhiteBoxKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWhiteBoxKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alias")
	delete(f, "Algorithm")
	delete(f, "Description")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWhiteBoxKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWhiteBoxKeyResponseParams struct {
	// 用于加密的密钥，base64编码
	EncryptKey *string `json:"EncryptKey,omitnil,omitempty" name:"EncryptKey"`

	// 用于解密的密钥，base64编码
	DecryptKey *string `json:"DecryptKey,omitnil,omitempty" name:"DecryptKey"`

	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 标签操作的返回码. 0: 成功；1: 内部错误；2: 业务处理错误
	TagCode *uint64 `json:"TagCode,omitnil,omitempty" name:"TagCode"`

	// 标签操作的返回信息
	TagMsg *string `json:"TagMsg,omitnil,omitempty" name:"TagMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateWhiteBoxKeyResponseParams `json:"Response"`
}

func (r *CreateWhiteBoxKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataKey struct {
	// DataKey的全局唯一标识。
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

type DataKeyMetadata struct {
	// DataKey的全局唯一标识
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// CMK的全局唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 作为密钥更容易辨识，更容易被人看懂的数据密钥名称
	DataKeyName *string `json:"DataKeyName,omitnil,omitempty" name:"DataKeyName"`

	// 数据密钥的长度,单位字节
	NumberOfBytes *uint64 `json:"NumberOfBytes,omitnil,omitempty" name:"NumberOfBytes"`

	// 密钥创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// DataKey的描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// DataKey的状态， 取值为：Enabled | Disabled | PendingDelete
	KeyState *string `json:"KeyState,omitnil,omitempty" name:"KeyState"`

	// 创建者
	CreatorUin *uint64 `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 数据密钥的创建者，用户创建的为 user，授权各云产品自动创建的为对应的产品名
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 计划删除的时间
	DeletionDate *uint64 `json:"DeletionDate,omitnil,omitempty" name:"DeletionDate"`

	// DataKey 密钥材料类型，由KMS创建的为： TENCENT_KMS， 由用户导入的类型为：EXTERNAL
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`

	// 资源ID，格式：creatorUin/$creatorUin/$dataKeyId
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 密钥是否是主副本。0:主本，1:同步副本。
	IsSyncReplica *int64 `json:"IsSyncReplica,omitnil,omitempty" name:"IsSyncReplica"`

	// 同步的原始地域
	SourceRegion *string `json:"SourceRegion,omitnil,omitempty" name:"SourceRegion"`

	// 密钥同步的状态，0:未同步，1:同步成功，2:同步失败，3:同步中。
	SyncStatus *int64 `json:"SyncStatus,omitnil,omitempty" name:"SyncStatus"`

	// 同步的结果描述
	SyncMessages *string `json:"SyncMessages,omitnil,omitempty" name:"SyncMessages"`

	// 同步的开始时间
	SyncStartTime *uint64 `json:"SyncStartTime,omitnil,omitempty" name:"SyncStartTime"`

	// 同步的结束时间
	SyncEndTime *uint64 `json:"SyncEndTime,omitnil,omitempty" name:"SyncEndTime"`

	// 同步的原始集群，如果为空，是公有云公共集群
	SourceHsmClusterId *string `json:"SourceHsmClusterId,omitnil,omitempty" name:"SourceHsmClusterId"`
}

// Predefined struct for user
type DecryptRequestParams struct {
	// 待解密的密文数据
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// key/value对的json字符串，如果Encrypt指定了该参数，则在调用Decrypt API时需要提供同样的参数，最大支持1024字符
	EncryptionContext *string `json:"EncryptionContext,omitnil,omitempty" name:"EncryptionContext"`

	// PEM 格式公钥字符串，支持 RSA2048 和 SM2 公钥，用于对返回数据中的 Plaintext 值进行加密。若为空，则不对 Plaintext 值加密。
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitnil,omitempty" name:"EncryptionPublicKey"`

	// 非对称加密算法，配合 EncryptionPublicKey 对返回数据进行加密。目前支持：SM2（以 C1C3C2 格式返回密文），SM2_C1C3C2_ASN1 （以 C1C3C2 ASN1 格式返回密文），RSAES_PKCS1_V1_5，RSAES_OAEP_SHA_1，RSAES_OAEP_SHA_256。若为空，则默认为 SM2。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitnil,omitempty" name:"EncryptionAlgorithm"`
}

type DecryptRequest struct {
	*tchttp.BaseRequest
	
	// 待解密的密文数据
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// key/value对的json字符串，如果Encrypt指定了该参数，则在调用Decrypt API时需要提供同样的参数，最大支持1024字符
	EncryptionContext *string `json:"EncryptionContext,omitnil,omitempty" name:"EncryptionContext"`

	// PEM 格式公钥字符串，支持 RSA2048 和 SM2 公钥，用于对返回数据中的 Plaintext 值进行加密。若为空，则不对 Plaintext 值加密。
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitnil,omitempty" name:"EncryptionPublicKey"`

	// 非对称加密算法，配合 EncryptionPublicKey 对返回数据进行加密。目前支持：SM2（以 C1C3C2 格式返回密文），SM2_C1C3C2_ASN1 （以 C1C3C2 ASN1 格式返回密文），RSAES_PKCS1_V1_5，RSAES_OAEP_SHA_1，RSAES_OAEP_SHA_256。若为空，则默认为 SM2。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitnil,omitempty" name:"EncryptionAlgorithm"`
}

func (r *DecryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DecryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CiphertextBlob")
	delete(f, "EncryptionContext")
	delete(f, "EncryptionPublicKey")
	delete(f, "EncryptionAlgorithm")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DecryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DecryptResponseParams struct {
	// CMK的全局唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 若调用时未提供 EncryptionPublicKey，该字段值为 Base64 编码的明文，需进行 Base64 解码以获取明文。
	// 若调用时提供了 EncryptionPublicKey，则该字段值为使用 EncryptionPublicKey 公钥进行非对称加密后的 Base64 编码的密文。需在 Base64 解码后，使用用户上传的公钥对应的私钥进行进一步解密，以获取明文。
	Plaintext *string `json:"Plaintext,omitnil,omitempty" name:"Plaintext"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DecryptResponse struct {
	*tchttp.BaseResponse
	Response *DecryptResponseParams `json:"Response"`
}

func (r *DecryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImportedKeyMaterialRequestParams struct {
	// 指定需要删除密钥材料的EXTERNAL CMK。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DeleteImportedKeyMaterialRequest struct {
	*tchttp.BaseRequest
	
	// 指定需要删除密钥材料的EXTERNAL CMK。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DeleteImportedKeyMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImportedKeyMaterialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImportedKeyMaterialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImportedKeyMaterialResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImportedKeyMaterialResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImportedKeyMaterialResponseParams `json:"Response"`
}

func (r *DeleteImportedKeyMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImportedKeyMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhiteBoxKeyRequestParams struct {
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DeleteWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest
	
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DeleteWhiteBoxKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhiteBoxKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWhiteBoxKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhiteBoxKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWhiteBoxKeyResponseParams `json:"Response"`
}

func (r *DeleteWhiteBoxKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataKeyRequestParams struct {
	// 数据密钥全局唯一标识符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

type DescribeDataKeyRequest struct {
	*tchttp.BaseRequest
	
	// 数据密钥全局唯一标识符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

func (r *DescribeDataKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataKeyResponseParams struct {
	// 数据密钥属性信息
	DataKeyMetadata *DataKeyMetadata `json:"DataKeyMetadata,omitnil,omitempty" name:"DataKeyMetadata"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataKeyResponseParams `json:"Response"`
}

func (r *DescribeDataKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataKeysRequestParams struct {
	// 查询DataKey的ID列表，批量查询一次最多支持100个DataKeyId
	DataKeyIds []*string `json:"DataKeyIds,omitnil,omitempty" name:"DataKeyIds"`
}

type DescribeDataKeysRequest struct {
	*tchttp.BaseRequest
	
	// 查询DataKey的ID列表，批量查询一次最多支持100个DataKeyId
	DataKeyIds []*string `json:"DataKeyIds,omitnil,omitempty" name:"DataKeyIds"`
}

func (r *DescribeDataKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataKeysResponseParams struct {
	// 返回数据密钥属性信息列表
	DataKeyMetadatas []*DataKeyMetadata `json:"DataKeyMetadatas,omitnil,omitempty" name:"DataKeyMetadatas"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataKeysResponseParams `json:"Response"`
}

func (r *DescribeDataKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeyRequestParams struct {
	// CMK全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DescribeKeyRequest struct {
	*tchttp.BaseRequest
	
	// CMK全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DescribeKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeyResponseParams struct {
	// 密钥属性信息
	KeyMetadata *KeyMetadata `json:"KeyMetadata,omitnil,omitempty" name:"KeyMetadata"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeyResponseParams `json:"Response"`
}

func (r *DescribeKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeysRequestParams struct {
	// 查询CMK的ID列表，批量查询一次最多支持100个KeyId
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type DescribeKeysRequest struct {
	*tchttp.BaseRequest
	
	// 查询CMK的ID列表，批量查询一次最多支持100个KeyId
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *DescribeKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeysResponseParams struct {
	// 返回的属性信息列表
	KeyMetadatas []*KeyMetadata `json:"KeyMetadatas,omitnil,omitempty" name:"KeyMetadatas"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeysResponseParams `json:"Response"`
}

func (r *DescribeKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxDecryptKeyRequestParams struct {
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DescribeWhiteBoxDecryptKeyRequest struct {
	*tchttp.BaseRequest
	
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DescribeWhiteBoxDecryptKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxDecryptKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoxDecryptKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxDecryptKeyResponseParams struct {
	// 白盒解密密钥，base64编码
	DecryptKey *string `json:"DecryptKey,omitnil,omitempty" name:"DecryptKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWhiteBoxDecryptKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoxDecryptKeyResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoxDecryptKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxDecryptKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxDeviceFingerprintsRequestParams struct {
	// 白盒密钥ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DescribeWhiteBoxDeviceFingerprintsRequest struct {
	*tchttp.BaseRequest
	
	// 白盒密钥ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DescribeWhiteBoxDeviceFingerprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxDeviceFingerprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoxDeviceFingerprintsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxDeviceFingerprintsResponseParams struct {
	// 设备指纹列表
	DeviceFingerprints []*DeviceFingerprint `json:"DeviceFingerprints,omitnil,omitempty" name:"DeviceFingerprints"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWhiteBoxDeviceFingerprintsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoxDeviceFingerprintsResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoxDeviceFingerprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxDeviceFingerprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxKeyDetailsRequestParams struct {
	// 过滤条件：密钥的状态，0：disabled，1：enabled
	KeyStatus *int64 `json:"KeyStatus,omitnil,omitempty" name:"KeyStatus"`

	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为0, 表示不分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeWhiteBoxKeyDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件：密钥的状态，0：disabled，1：enabled
	KeyStatus *int64 `json:"KeyStatus,omitnil,omitempty" name:"KeyStatus"`

	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为0, 表示不分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeWhiteBoxKeyDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxKeyDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyStatus")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoxKeyDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxKeyDetailsResponseParams struct {
	// 白盒密钥信息列表。
	KeyInfos []*WhiteboxKeyInfo `json:"KeyInfos,omitnil,omitempty" name:"KeyInfos"`

	// 白盒密钥总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWhiteBoxKeyDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoxKeyDetailsResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoxKeyDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxKeyDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxKeyRequestParams struct {
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DescribeWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest
	
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DescribeWhiteBoxKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoxKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxKeyResponseParams struct {
	// 白盒密钥信息
	KeyInfo *WhiteboxKeyInfo `json:"KeyInfo,omitnil,omitempty" name:"KeyInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoxKeyResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoxKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxServiceStatusRequestParams struct {

}

type DescribeWhiteBoxServiceStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWhiteBoxServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxServiceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoxServiceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxServiceStatusResponseParams struct {
	// 用户的白盒密钥服务是否可用
	ServiceEnabled *bool `json:"ServiceEnabled,omitnil,omitempty" name:"ServiceEnabled"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWhiteBoxServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoxServiceStatusResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoxServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestinationSyncConfig struct {
	// 同步任务的目标地域
	DestinationRegion *string `json:"DestinationRegion,omitnil,omitempty" name:"DestinationRegion"`

	// HsmClusterId为空表示公有云共享版，如果不为空表示地域下独享版集群。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

type DeviceFingerprint struct {
	// 指纹信息，由设备指纹采集工具采集获得，格式满足正则表达式：^[0-9a-f]{8}[\-][0-9a-f]{14}[\-][0-9a-f]{14}[\-][0-9a-f]{14}[\-][0-9a-f]{16}$
	Identity *string `json:"Identity,omitnil,omitempty" name:"Identity"`

	// 描述信息，如：IP，设备名称等，最大1024字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type DisableDataKeyRequestParams struct {
	// 数据密钥唯一标识符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

type DisableDataKeyRequest struct {
	*tchttp.BaseRequest
	
	// 数据密钥唯一标识符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

func (r *DisableDataKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableDataKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableDataKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableDataKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableDataKeyResponse struct {
	*tchttp.BaseResponse
	Response *DisableDataKeyResponseParams `json:"Response"`
}

func (r *DisableDataKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableDataKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableDataKeysRequestParams struct {
	// 需要批量禁用的DataKey Id 列表，数据密钥数量最大支持100
	DataKeyIds []*string `json:"DataKeyIds,omitnil,omitempty" name:"DataKeyIds"`
}

type DisableDataKeysRequest struct {
	*tchttp.BaseRequest
	
	// 需要批量禁用的DataKey Id 列表，数据密钥数量最大支持100
	DataKeyIds []*string `json:"DataKeyIds,omitnil,omitempty" name:"DataKeyIds"`
}

func (r *DisableDataKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableDataKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableDataKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableDataKeysResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableDataKeysResponse struct {
	*tchttp.BaseResponse
	Response *DisableDataKeysResponseParams `json:"Response"`
}

func (r *DisableDataKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableDataKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeyRequestParams struct {
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DisableKeyRequest struct {
	*tchttp.BaseRequest
	
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DisableKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableKeyResponse struct {
	*tchttp.BaseResponse
	Response *DisableKeyResponseParams `json:"Response"`
}

func (r *DisableKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeyRotationRequestParams struct {
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DisableKeyRotationRequest struct {
	*tchttp.BaseRequest
	
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DisableKeyRotationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeyRotationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableKeyRotationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeyRotationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableKeyRotationResponse struct {
	*tchttp.BaseResponse
	Response *DisableKeyRotationResponseParams `json:"Response"`
}

func (r *DisableKeyRotationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeyRotationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeysRequestParams struct {
	// 需要批量禁用的CMK Id 列表，CMK数量最大支持100
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type DisableKeysRequest struct {
	*tchttp.BaseRequest
	
	// 需要批量禁用的CMK Id 列表，CMK数量最大支持100
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *DisableKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeysResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableKeysResponse struct {
	*tchttp.BaseResponse
	Response *DisableKeysResponseParams `json:"Response"`
}

func (r *DisableKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWhiteBoxKeyRequestParams struct {
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DisableWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest
	
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DisableWhiteBoxKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWhiteBoxKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableWhiteBoxKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWhiteBoxKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse
	Response *DisableWhiteBoxKeyResponseParams `json:"Response"`
}

func (r *DisableWhiteBoxKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWhiteBoxKeysRequestParams struct {
	// 白盒密钥的全局唯一标识符列表。注意：要确保所有提供的KeyId是格式有效的，没有重复，个数不超过50个，并且都是有效存在的。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type DisableWhiteBoxKeysRequest struct {
	*tchttp.BaseRequest
	
	// 白盒密钥的全局唯一标识符列表。注意：要确保所有提供的KeyId是格式有效的，没有重复，个数不超过50个，并且都是有效存在的。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *DisableWhiteBoxKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWhiteBoxKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableWhiteBoxKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWhiteBoxKeysResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableWhiteBoxKeysResponse struct {
	*tchttp.BaseResponse
	Response *DisableWhiteBoxKeysResponseParams `json:"Response"`
}

func (r *DisableWhiteBoxKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWhiteBoxKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableDataKeyRequestParams struct {
	// 数据密钥唯一标识符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

type EnableDataKeyRequest struct {
	*tchttp.BaseRequest
	
	// 数据密钥唯一标识符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

func (r *EnableDataKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableDataKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableDataKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableDataKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableDataKeyResponse struct {
	*tchttp.BaseResponse
	Response *EnableDataKeyResponseParams `json:"Response"`
}

func (r *EnableDataKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableDataKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableDataKeysRequestParams struct {
	// 需要批量启用的DataKey Id 列表， 数据密钥数量最大支持100
	DataKeyIds []*string `json:"DataKeyIds,omitnil,omitempty" name:"DataKeyIds"`
}

type EnableDataKeysRequest struct {
	*tchttp.BaseRequest
	
	// 需要批量启用的DataKey Id 列表， 数据密钥数量最大支持100
	DataKeyIds []*string `json:"DataKeyIds,omitnil,omitempty" name:"DataKeyIds"`
}

func (r *EnableDataKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableDataKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableDataKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableDataKeysResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableDataKeysResponse struct {
	*tchttp.BaseResponse
	Response *EnableDataKeysResponseParams `json:"Response"`
}

func (r *EnableDataKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableDataKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeyRequestParams struct {
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type EnableKeyRequest struct {
	*tchttp.BaseRequest
	
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *EnableKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableKeyResponse struct {
	*tchttp.BaseResponse
	Response *EnableKeyResponseParams `json:"Response"`
}

func (r *EnableKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeyRotationRequestParams struct {
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 密钥轮转周期，单位天，允许范围 7 ~ 365，默认值 365。
	RotateDays *uint64 `json:"RotateDays,omitnil,omitempty" name:"RotateDays"`
}

type EnableKeyRotationRequest struct {
	*tchttp.BaseRequest
	
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 密钥轮转周期，单位天，允许范围 7 ~ 365，默认值 365。
	RotateDays *uint64 `json:"RotateDays,omitnil,omitempty" name:"RotateDays"`
}

func (r *EnableKeyRotationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeyRotationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "RotateDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableKeyRotationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeyRotationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableKeyRotationResponse struct {
	*tchttp.BaseResponse
	Response *EnableKeyRotationResponseParams `json:"Response"`
}

func (r *EnableKeyRotationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeyRotationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeysRequestParams struct {
	// 需要批量启用的CMK Id 列表， CMK数量最大支持100
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type EnableKeysRequest struct {
	*tchttp.BaseRequest
	
	// 需要批量启用的CMK Id 列表， CMK数量最大支持100
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *EnableKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeysResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableKeysResponse struct {
	*tchttp.BaseResponse
	Response *EnableKeysResponseParams `json:"Response"`
}

func (r *EnableKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWhiteBoxKeyRequestParams struct {
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type EnableWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest
	
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *EnableWhiteBoxKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWhiteBoxKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableWhiteBoxKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWhiteBoxKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse
	Response *EnableWhiteBoxKeyResponseParams `json:"Response"`
}

func (r *EnableWhiteBoxKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWhiteBoxKeysRequestParams struct {
	// 白盒密钥的全局唯一标识符列表。注意：要确保所有提供的KeyId是格式有效的，没有重复，个数不超过50个，并且都是有效存在的。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type EnableWhiteBoxKeysRequest struct {
	*tchttp.BaseRequest
	
	// 白盒密钥的全局唯一标识符列表。注意：要确保所有提供的KeyId是格式有效的，没有重复，个数不超过50个，并且都是有效存在的。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *EnableWhiteBoxKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWhiteBoxKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableWhiteBoxKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWhiteBoxKeysResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableWhiteBoxKeysResponse struct {
	*tchttp.BaseResponse
	Response *EnableWhiteBoxKeysResponseParams `json:"Response"`
}

func (r *EnableWhiteBoxKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWhiteBoxKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EncryptByWhiteBoxRequestParams struct {
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 待加密的文本， base64编码，文本的原始长度最大不超过4KB
	PlainText *string `json:"PlainText,omitnil,omitempty" name:"PlainText"`

	// 初始化向量，大小为 16 Bytes，加密算法会使用到, base64编码；如果不传，则由后端服务随机生成。用户需要自行保存该值，作为解密的参数。
	InitializationVector *string `json:"InitializationVector,omitnil,omitempty" name:"InitializationVector"`
}

type EncryptByWhiteBoxRequest struct {
	*tchttp.BaseRequest
	
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 待加密的文本， base64编码，文本的原始长度最大不超过4KB
	PlainText *string `json:"PlainText,omitnil,omitempty" name:"PlainText"`

	// 初始化向量，大小为 16 Bytes，加密算法会使用到, base64编码；如果不传，则由后端服务随机生成。用户需要自行保存该值，作为解密的参数。
	InitializationVector *string `json:"InitializationVector,omitnil,omitempty" name:"InitializationVector"`
}

func (r *EncryptByWhiteBoxRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EncryptByWhiteBoxRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "PlainText")
	delete(f, "InitializationVector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EncryptByWhiteBoxRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EncryptByWhiteBoxResponseParams struct {
	// 初始化向量，加密算法会使用到, base64编码。如果由调用方在入参中传入，则原样返回。如果调用方没有传入，则后端服务随机生成，并返回
	InitializationVector *string `json:"InitializationVector,omitnil,omitempty" name:"InitializationVector"`

	// 加密后的密文，base64编码
	CipherText *string `json:"CipherText,omitnil,omitempty" name:"CipherText"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EncryptByWhiteBoxResponse struct {
	*tchttp.BaseResponse
	Response *EncryptByWhiteBoxResponseParams `json:"Response"`
}

func (r *EncryptByWhiteBoxResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EncryptByWhiteBoxResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EncryptRequestParams struct {
	// 调用CreateKey生成的CMK全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 被加密的明文数据，该字段必须使用base64编码，原文最大长度支持4K
	Plaintext *string `json:"Plaintext,omitnil,omitempty" name:"Plaintext"`

	// key/value对的json字符串，如果指定了该参数，则在调用Decrypt API时需要提供同样的参数，最大支持1024个字符
	EncryptionContext *string `json:"EncryptionContext,omitnil,omitempty" name:"EncryptionContext"`
}

type EncryptRequest struct {
	*tchttp.BaseRequest
	
	// 调用CreateKey生成的CMK全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 被加密的明文数据，该字段必须使用base64编码，原文最大长度支持4K
	Plaintext *string `json:"Plaintext,omitnil,omitempty" name:"Plaintext"`

	// key/value对的json字符串，如果指定了该参数，则在调用Decrypt API时需要提供同样的参数，最大支持1024个字符
	EncryptionContext *string `json:"EncryptionContext,omitnil,omitempty" name:"EncryptionContext"`
}

func (r *EncryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EncryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "Plaintext")
	delete(f, "EncryptionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EncryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EncryptResponseParams struct {
	// 加密后的密文，base64编码。注意：本字段中打包了密文和密钥的相关信息，不是对明文的直接加密结果，只有将该字段作为Decrypt接口的输入参数，才可以解密出原文。
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// 加密使用的CMK的全局唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EncryptResponse struct {
	*tchttp.BaseResponse
	Response *EncryptResponseParams `json:"Response"`
}

func (r *EncryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExclusiveHSM struct {
	// 独享集群Id
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`

	// 独享集群名称
	HsmClusterName *string `json:"HsmClusterName,omitnil,omitempty" name:"HsmClusterName"`
}

// Predefined struct for user
type GenerateDataKeyRequestParams struct {
	// CMK全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 指定生成Datakey的加密算法以及Datakey大小，AES_128或者AES_256。KeySpec 和 NumberOfBytes 必须指定一个
	KeySpec *string `json:"KeySpec,omitnil,omitempty" name:"KeySpec"`

	// 生成的DataKey的长度，同时指定NumberOfBytes和KeySpec时，以NumberOfBytes为准。最小值为1， 最大值为1024。KeySpec 和 NumberOfBytes 必须指定一个
	NumberOfBytes *uint64 `json:"NumberOfBytes,omitnil,omitempty" name:"NumberOfBytes"`

	// key/value对的json字符串，如果使用该字段，则返回的DataKey在解密时需要填入相同的字符串
	EncryptionContext *string `json:"EncryptionContext,omitnil,omitempty" name:"EncryptionContext"`

	// PEM 格式公钥字符串，支持 RSA2048 和 SM2 公钥，用于对返回数据中的 Plaintext 值进行加密。若为空，则不对 Plaintext 值加密。
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitnil,omitempty" name:"EncryptionPublicKey"`

	// 非对称加密算法，配合 EncryptionPublicKey 对返回数据进行加密。目前支持：SM2（以 C1C3C2 格式返回密文），SM2_C1C3C2_ASN1 （以 C1C3C2 ASN1 格式返回密文），RSAES_PKCS1_V1_5，RSAES_OAEP_SHA_1，RSAES_OAEP_SHA_256。若为空，则默认为 SM2。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitnil,omitempty" name:"EncryptionAlgorithm"`

	// 表示生成的数据密钥是否被KMS托管。1:表示被KMS托管保存,0:表示KMS不托管。
	IsHostedByKms *uint64 `json:"IsHostedByKms,omitnil,omitempty" name:"IsHostedByKms"`

	// 数据密钥的名称，当IsHostedByKms为1时,必须填写。当IsHostedByKms为0时,可以不填，KMS不托管。
	DataKeyName *string `json:"DataKeyName,omitnil,omitempty" name:"DataKeyName"`

	// 数据密钥 的描述，最大100字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// KMS 独享版对应的 HSM 集群 ID。如果指定HsmClusterId，表明根密钥在此集群里，会校验KeyId是否和HsmClusterId对应。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

type GenerateDataKeyRequest struct {
	*tchttp.BaseRequest
	
	// CMK全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 指定生成Datakey的加密算法以及Datakey大小，AES_128或者AES_256。KeySpec 和 NumberOfBytes 必须指定一个
	KeySpec *string `json:"KeySpec,omitnil,omitempty" name:"KeySpec"`

	// 生成的DataKey的长度，同时指定NumberOfBytes和KeySpec时，以NumberOfBytes为准。最小值为1， 最大值为1024。KeySpec 和 NumberOfBytes 必须指定一个
	NumberOfBytes *uint64 `json:"NumberOfBytes,omitnil,omitempty" name:"NumberOfBytes"`

	// key/value对的json字符串，如果使用该字段，则返回的DataKey在解密时需要填入相同的字符串
	EncryptionContext *string `json:"EncryptionContext,omitnil,omitempty" name:"EncryptionContext"`

	// PEM 格式公钥字符串，支持 RSA2048 和 SM2 公钥，用于对返回数据中的 Plaintext 值进行加密。若为空，则不对 Plaintext 值加密。
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitnil,omitempty" name:"EncryptionPublicKey"`

	// 非对称加密算法，配合 EncryptionPublicKey 对返回数据进行加密。目前支持：SM2（以 C1C3C2 格式返回密文），SM2_C1C3C2_ASN1 （以 C1C3C2 ASN1 格式返回密文），RSAES_PKCS1_V1_5，RSAES_OAEP_SHA_1，RSAES_OAEP_SHA_256。若为空，则默认为 SM2。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitnil,omitempty" name:"EncryptionAlgorithm"`

	// 表示生成的数据密钥是否被KMS托管。1:表示被KMS托管保存,0:表示KMS不托管。
	IsHostedByKms *uint64 `json:"IsHostedByKms,omitnil,omitempty" name:"IsHostedByKms"`

	// 数据密钥的名称，当IsHostedByKms为1时,必须填写。当IsHostedByKms为0时,可以不填，KMS不托管。
	DataKeyName *string `json:"DataKeyName,omitnil,omitempty" name:"DataKeyName"`

	// 数据密钥 的描述，最大100字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// KMS 独享版对应的 HSM 集群 ID。如果指定HsmClusterId，表明根密钥在此集群里，会校验KeyId是否和HsmClusterId对应。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

func (r *GenerateDataKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateDataKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "KeySpec")
	delete(f, "NumberOfBytes")
	delete(f, "EncryptionContext")
	delete(f, "EncryptionPublicKey")
	delete(f, "EncryptionAlgorithm")
	delete(f, "IsHostedByKms")
	delete(f, "DataKeyName")
	delete(f, "Description")
	delete(f, "HsmClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateDataKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateDataKeyResponseParams struct {
	// CMK的全局唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 若调用时未提供 EncryptionPublicKey，该字段值为生成的数据密钥 DataKey 的 Base64 编码的明文，需进行 Base64 解码以获取 DataKey 明文。
	// 若调用时提供了 EncryptionPublicKey，则该字段值为使用 EncryptionPublicKey 公钥进行非对称加密后的 Base64 编码的密文。需在 Base64 解码后，使用用户上传的公钥对应的私钥进行进一步解密，以获取 DataKey 明文。
	Plaintext *string `json:"Plaintext,omitnil,omitempty" name:"Plaintext"`

	// 数据密钥DataKey加密后的密文，用户需要自行保存该密文，KMS不托管用户的数据密钥。可以通过Decrypt接口从CiphertextBlob中获取数据密钥DataKey明文
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// DataKey的全局唯一标识,当KMS托管数据密钥时返回。
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateDataKeyResponse struct {
	*tchttp.BaseResponse
	Response *GenerateDataKeyResponseParams `json:"Response"`
}

func (r *GenerateDataKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateDataKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateRandomRequestParams struct {
	// 生成的随机数的长度。最小值为1， 最大值为1024。
	NumberOfBytes *uint64 `json:"NumberOfBytes,omitnil,omitempty" name:"NumberOfBytes"`
}

type GenerateRandomRequest struct {
	*tchttp.BaseRequest
	
	// 生成的随机数的长度。最小值为1， 最大值为1024。
	NumberOfBytes *uint64 `json:"NumberOfBytes,omitnil,omitempty" name:"NumberOfBytes"`
}

func (r *GenerateRandomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateRandomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NumberOfBytes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateRandomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateRandomResponseParams struct {
	// 生成的随机数的明文，该明文使用base64编码，用户需要使用base64解码得到明文。
	Plaintext *string `json:"Plaintext,omitnil,omitempty" name:"Plaintext"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateRandomResponse struct {
	*tchttp.BaseResponse
	Response *GenerateRandomResponseParams `json:"Response"`
}

func (r *GenerateRandomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateRandomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataKeyCiphertextBlobRequestParams struct {
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

type GetDataKeyCiphertextBlobRequest struct {
	*tchttp.BaseRequest
	
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`
}

func (r *GetDataKeyCiphertextBlobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataKeyCiphertextBlobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDataKeyCiphertextBlobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataKeyCiphertextBlobResponseParams struct {
	// 数据密钥的密文
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDataKeyCiphertextBlobResponse struct {
	*tchttp.BaseResponse
	Response *GetDataKeyCiphertextBlobResponseParams `json:"Response"`
}

func (r *GetDataKeyCiphertextBlobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataKeyCiphertextBlobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataKeyPlaintextRequestParams struct {
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// PEM 格式公钥字符串，支持 RSA2048 和 SM2 公钥，用于对返回数据中的 Plaintext 值进行加密。若为空，则不对 Plaintext 值加密。
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitnil,omitempty" name:"EncryptionPublicKey"`

	// 非对称加密算法，配合 EncryptionPublicKey 对返回数据进行加密。目前支持：SM2（以 C1C3C2 格式返回密文），SM2_C1C3C2_ASN1 （以 C1C3C2 ASN1 格式返回密文），RSAES_PKCS1_V1_5，RSAES_OAEP_SHA_1，RSAES_OAEP_SHA_256。若为空，则默认为 SM2。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitnil,omitempty" name:"EncryptionAlgorithm"`
}

type GetDataKeyPlaintextRequest struct {
	*tchttp.BaseRequest
	
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// PEM 格式公钥字符串，支持 RSA2048 和 SM2 公钥，用于对返回数据中的 Plaintext 值进行加密。若为空，则不对 Plaintext 值加密。
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitnil,omitempty" name:"EncryptionPublicKey"`

	// 非对称加密算法，配合 EncryptionPublicKey 对返回数据进行加密。目前支持：SM2（以 C1C3C2 格式返回密文），SM2_C1C3C2_ASN1 （以 C1C3C2 ASN1 格式返回密文），RSAES_PKCS1_V1_5，RSAES_OAEP_SHA_1，RSAES_OAEP_SHA_256。若为空，则默认为 SM2。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitnil,omitempty" name:"EncryptionAlgorithm"`
}

func (r *GetDataKeyPlaintextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataKeyPlaintextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyId")
	delete(f, "EncryptionPublicKey")
	delete(f, "EncryptionAlgorithm")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDataKeyPlaintextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataKeyPlaintextResponseParams struct {
	// 若调用时未提供 EncryptionPublicKey，该字段值为 Base64 编码的明文，需进行 Base64 解码以获取明文。 若调用时提供了 EncryptionPublicKey，则该字段值为使用 EncryptionPublicKey 公钥进行非对称加密后的 Base64 编码的密文。需在 Base64 解码后，使用用户上传的公钥对应的私钥进行进一步解密，以获取明文。
	Plaintext *string `json:"Plaintext,omitnil,omitempty" name:"Plaintext"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDataKeyPlaintextResponse struct {
	*tchttp.BaseResponse
	Response *GetDataKeyPlaintextResponseParams `json:"Response"`
}

func (r *GetDataKeyPlaintextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataKeyPlaintextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetKeyRotationStatusRequestParams struct {
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type GetKeyRotationStatusRequest struct {
	*tchttp.BaseRequest
	
	// CMK唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *GetKeyRotationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetKeyRotationStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetKeyRotationStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetKeyRotationStatusResponseParams struct {
	// 密钥轮换是否开启
	KeyRotationEnabled *bool `json:"KeyRotationEnabled,omitnil,omitempty" name:"KeyRotationEnabled"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetKeyRotationStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetKeyRotationStatusResponseParams `json:"Response"`
}

func (r *GetKeyRotationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetKeyRotationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetParametersForImportRequestParams struct {
	// CMK的唯一标识，获取密钥参数的CMK必须是EXTERNAL类型，即在CreateKey时指定Type=2 类型的CMK。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 指定加密密钥材料的算法，目前支持RSAES_PKCS1_V1_5、RSAES_OAEP_SHA_1、RSAES_OAEP_SHA_256
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitnil,omitempty" name:"WrappingAlgorithm"`

	// 指定加密密钥材料的类型，目前只支持RSA_2048
	WrappingKeySpec *string `json:"WrappingKeySpec,omitnil,omitempty" name:"WrappingKeySpec"`
}

type GetParametersForImportRequest struct {
	*tchttp.BaseRequest
	
	// CMK的唯一标识，获取密钥参数的CMK必须是EXTERNAL类型，即在CreateKey时指定Type=2 类型的CMK。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 指定加密密钥材料的算法，目前支持RSAES_PKCS1_V1_5、RSAES_OAEP_SHA_1、RSAES_OAEP_SHA_256
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitnil,omitempty" name:"WrappingAlgorithm"`

	// 指定加密密钥材料的类型，目前只支持RSA_2048
	WrappingKeySpec *string `json:"WrappingKeySpec,omitnil,omitempty" name:"WrappingKeySpec"`
}

func (r *GetParametersForImportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetParametersForImportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "WrappingAlgorithm")
	delete(f, "WrappingKeySpec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetParametersForImportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetParametersForImportResponseParams struct {
	// CMK的唯一标识，用于指定目标导入密钥材料的CMK。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 导入密钥材料需要的token，用于作为 ImportKeyMaterial 的参数。
	ImportToken *string `json:"ImportToken,omitnil,omitempty" name:"ImportToken"`

	// 用于加密密钥材料的RSA公钥，base64编码。使用PublicKey base64解码后的公钥将导入密钥进行加密后作为 ImportKeyMaterial 的参数。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 该导出token和公钥的有效期，超过该时间后无法导入，需要重新调用GetParametersForImport获取。
	ParametersValidTo *uint64 `json:"ParametersValidTo,omitnil,omitempty" name:"ParametersValidTo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetParametersForImportResponse struct {
	*tchttp.BaseResponse
	Response *GetParametersForImportResponseParams `json:"Response"`
}

func (r *GetParametersForImportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetParametersForImportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPublicKeyRequestParams struct {
	// CMK的唯一标识。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type GetPublicKeyRequest struct {
	*tchttp.BaseRequest
	
	// CMK的唯一标识。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *GetPublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPublicKeyResponseParams struct {
	// CMK的唯一标识。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 经过base64编码的公钥内容。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// PEM格式的公钥内容。
	PublicKeyPem *string `json:"PublicKeyPem,omitnil,omitempty" name:"PublicKeyPem"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *GetPublicKeyResponseParams `json:"Response"`
}

func (r *GetPublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRegionsRequestParams struct {

}

type GetRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRegionsResponseParams struct {
	// 可用region列表
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRegionsResponse struct {
	*tchttp.BaseResponse
	Response *GetRegionsResponseParams `json:"Response"`
}

func (r *GetRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetServiceStatusRequestParams struct {

}

type GetServiceStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetServiceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetServiceStatusResponseParams struct {
	// KMS服务是否开通， true 表示已开通
	ServiceEnabled *bool `json:"ServiceEnabled,omitnil,omitempty" name:"ServiceEnabled"`

	// 服务不可用类型： 0-未购买，1-正常， 2-欠费停服， 3-资源释放
	InvalidType *int64 `json:"InvalidType,omitnil,omitempty" name:"InvalidType"`

	// 0-普通版，1-旗舰版
	UserLevel *uint64 `json:"UserLevel,omitnil,omitempty" name:"UserLevel"`

	// 旗舰版到期时间（Epoch Unix Timestamp）。
	ProExpireTime *uint64 `json:"ProExpireTime,omitnil,omitempty" name:"ProExpireTime"`

	// 旗舰版是否自动续费：0-不自动续费，1-自动续费
	ProRenewFlag *uint64 `json:"ProRenewFlag,omitnil,omitempty" name:"ProRenewFlag"`

	// 旗舰版购买记录的唯一性标识。如果未开通旗舰版，则返回值为空
	ProResourceId *string `json:"ProResourceId,omitnil,omitempty" name:"ProResourceId"`

	// 是否开通 KMS 托管版
	ExclusiveVSMEnabled *bool `json:"ExclusiveVSMEnabled,omitnil,omitempty" name:"ExclusiveVSMEnabled"`

	// 是否开通 KMS 独享版
	ExclusiveHSMEnabled *bool `json:"ExclusiveHSMEnabled,omitnil,omitempty" name:"ExclusiveHSMEnabled"`

	// KMS 订阅信息。
	SubscriptionInfo *string `json:"SubscriptionInfo,omitnil,omitempty" name:"SubscriptionInfo"`

	// 返回KMS用户密钥使用数量
	CmkUserCount *uint64 `json:"CmkUserCount,omitnil,omitempty" name:"CmkUserCount"`

	// 返回KMS用户密钥规格数量
	CmkLimit *uint64 `json:"CmkLimit,omitnil,omitempty" name:"CmkLimit"`

	// 返回独享集群组
	ExclusiveHSMList []*ExclusiveHSM `json:"ExclusiveHSMList,omitnil,omitempty" name:"ExclusiveHSMList"`

	// 是否支持数据密钥托管。1:支持，0:不支持。
	IsAllowedDataKeyHosted *bool `json:"IsAllowedDataKeyHosted,omitnil,omitempty" name:"IsAllowedDataKeyHosted"`

	// IsAllowedDataKeyHosted为1时有效，数据密钥的购买额度
	DataKeyLimit *uint64 `json:"DataKeyLimit,omitnil,omitempty" name:"DataKeyLimit"`

	// IsAllowedDataKeyHosted为1时有效，数据密钥免费额度。
	FreeDataKeyLimit *uint64 `json:"FreeDataKeyLimit,omitnil,omitempty" name:"FreeDataKeyLimit"`

	// IsAllowedDataKeyHosted为1时有效，已使用的数据密钥数量。
	DataKeyUsedCount *uint64 `json:"DataKeyUsedCount,omitnil,omitempty" name:"DataKeyUsedCount"`

	// 同步任务的目标地域信息
	SyncTaskList []*DestinationSyncConfig `json:"SyncTaskList,omitnil,omitempty" name:"SyncTaskList"`

	// 是否支持同步任务。true:支持，false:不支持。
	IsAllowedSync *bool `json:"IsAllowedSync,omitnil,omitempty" name:"IsAllowedSync"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetServiceStatusResponseParams `json:"Response"`
}

func (r *GetServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportDataKeyRequestParams struct {
	// 数据密钥的名称
	DataKeyName *string `json:"DataKeyName,omitnil,omitempty" name:"DataKeyName"`

	// 如果导入的是明文数据密钥，则是base64 转换后的明文数据密钥，  如果导入的是密文数据密钥，则是由KMS GenerateDataKey接口生成的密文数据密钥。
	ImportKeyMaterial *string `json:"ImportKeyMaterial,omitnil,omitempty" name:"ImportKeyMaterial"`

	// 1:密文导入(由KMS接口生成的密文数据密钥)，2:明文导入。
	ImportType *uint64 `json:"ImportType,omitnil,omitempty" name:"ImportType"`

	// 数据密钥 的描述，最大100字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 当导入密文数据密钥时，无需传入根密钥,如果传入会校验此KeyId是否和密文中一致。
	// 当导入明文数据密钥，KeyId 不能为空，会根据指定的根密钥加密数据密钥。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// KMS 独享版对应的 HSM 集群 ID。如果指定HsmClusterId，表明根密钥在此集群里，会校验KeyId是否和HsmClusterId对应。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

type ImportDataKeyRequest struct {
	*tchttp.BaseRequest
	
	// 数据密钥的名称
	DataKeyName *string `json:"DataKeyName,omitnil,omitempty" name:"DataKeyName"`

	// 如果导入的是明文数据密钥，则是base64 转换后的明文数据密钥，  如果导入的是密文数据密钥，则是由KMS GenerateDataKey接口生成的密文数据密钥。
	ImportKeyMaterial *string `json:"ImportKeyMaterial,omitnil,omitempty" name:"ImportKeyMaterial"`

	// 1:密文导入(由KMS接口生成的密文数据密钥)，2:明文导入。
	ImportType *uint64 `json:"ImportType,omitnil,omitempty" name:"ImportType"`

	// 数据密钥 的描述，最大100字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 当导入密文数据密钥时，无需传入根密钥,如果传入会校验此KeyId是否和密文中一致。
	// 当导入明文数据密钥，KeyId 不能为空，会根据指定的根密钥加密数据密钥。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// KMS 独享版对应的 HSM 集群 ID。如果指定HsmClusterId，表明根密钥在此集群里，会校验KeyId是否和HsmClusterId对应。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

func (r *ImportDataKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportDataKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyName")
	delete(f, "ImportKeyMaterial")
	delete(f, "ImportType")
	delete(f, "Description")
	delete(f, "KeyId")
	delete(f, "HsmClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportDataKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportDataKeyResponseParams struct {
	// CMK的全局唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// DataKey的全局唯一标识  否  官网/国内&国际站展示
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportDataKeyResponse struct {
	*tchttp.BaseResponse
	Response *ImportDataKeyResponseParams `json:"Response"`
}

func (r *ImportDataKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportDataKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportKeyMaterialRequestParams struct {
	// 使用GetParametersForImport 返回的PublicKey加密后的密钥材料base64编码。对于国密版本region的KMS，导入的密钥材料长度要求为 128 bit，FIPS版本region的KMS， 导入的密钥材料长度要求为 256 bit。
	EncryptedKeyMaterial *string `json:"EncryptedKeyMaterial,omitnil,omitempty" name:"EncryptedKeyMaterial"`

	// 通过调用GetParametersForImport获得的导入令牌。
	ImportToken *string `json:"ImportToken,omitnil,omitempty" name:"ImportToken"`

	// 指定导入密钥材料的CMK，需要和GetParametersForImport 指定的CMK相同。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 密钥材料过期时间 unix 时间戳，不指定或者 0 表示密钥材料不会过期，若指定过期时间，需要大于当前时间点，最大支持 2147443200。
	ValidTo *uint64 `json:"ValidTo,omitnil,omitempty" name:"ValidTo"`
}

type ImportKeyMaterialRequest struct {
	*tchttp.BaseRequest
	
	// 使用GetParametersForImport 返回的PublicKey加密后的密钥材料base64编码。对于国密版本region的KMS，导入的密钥材料长度要求为 128 bit，FIPS版本region的KMS， 导入的密钥材料长度要求为 256 bit。
	EncryptedKeyMaterial *string `json:"EncryptedKeyMaterial,omitnil,omitempty" name:"EncryptedKeyMaterial"`

	// 通过调用GetParametersForImport获得的导入令牌。
	ImportToken *string `json:"ImportToken,omitnil,omitempty" name:"ImportToken"`

	// 指定导入密钥材料的CMK，需要和GetParametersForImport 指定的CMK相同。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 密钥材料过期时间 unix 时间戳，不指定或者 0 表示密钥材料不会过期，若指定过期时间，需要大于当前时间点，最大支持 2147443200。
	ValidTo *uint64 `json:"ValidTo,omitnil,omitempty" name:"ValidTo"`
}

func (r *ImportKeyMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportKeyMaterialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EncryptedKeyMaterial")
	delete(f, "ImportToken")
	delete(f, "KeyId")
	delete(f, "ValidTo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportKeyMaterialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportKeyMaterialResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportKeyMaterialResponse struct {
	*tchttp.BaseResponse
	Response *ImportKeyMaterialResponseParams `json:"Response"`
}

func (r *ImportKeyMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportKeyMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Key struct {
	// CMK的全局唯一标识。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type KeyMetadata struct {
	// CMK的全局唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 作为密钥更容易辨识，更容易被人看懂的别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 密钥创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// CMK的描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// CMK的状态， 取值为：Enabled | Disabled | PendingDelete | PendingImport | Archived
	KeyState *string `json:"KeyState,omitnil,omitempty" name:"KeyState"`

	// CMK用途，取值为: ENCRYPT_DECRYPT | ASYMMETRIC_DECRYPT_RSA_2048 | ASYMMETRIC_DECRYPT_SM2 | ASYMMETRIC_SIGN_VERIFY_SM2 | ASYMMETRIC_SIGN_VERIFY_RSA_2048 | ASYMMETRIC_SIGN_VERIFY_ECC
	KeyUsage *string `json:"KeyUsage,omitnil,omitempty" name:"KeyUsage"`

	// CMK类型，2 表示符合FIPS标准，4表示符合国密标准
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 创建者
	CreatorUin *uint64 `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 是否开启了密钥轮换功能
	KeyRotationEnabled *bool `json:"KeyRotationEnabled,omitnil,omitempty" name:"KeyRotationEnabled"`

	// CMK的创建者，用户创建的为 user，授权各云产品自动创建的为对应的产品名
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 在密钥轮换开启状态下，下次轮换的时间
	NextRotateTime *uint64 `json:"NextRotateTime,omitnil,omitempty" name:"NextRotateTime"`

	// 计划删除的时间
	DeletionDate *uint64 `json:"DeletionDate,omitnil,omitempty" name:"DeletionDate"`

	// CMK 密钥材料类型，由KMS创建的为： TENCENT_KMS， 由用户导入的类型为：EXTERNAL
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 在Origin为  EXTERNAL 时有效，表示密钥材料的有效日期， 0 表示不过期
	ValidTo *uint64 `json:"ValidTo,omitnil,omitempty" name:"ValidTo"`

	// 资源ID，格式：creatorUin/$creatorUin/$keyId
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`

	// 密钥轮转周期（天）
	RotateDays *uint64 `json:"RotateDays,omitnil,omitempty" name:"RotateDays"`

	// 上次乱转时间（Unix timestamp）
	LastRotateTime *uint64 `json:"LastRotateTime,omitnil,omitempty" name:"LastRotateTime"`

	//  密钥是否是主副本。0:主本，1:同步副本。
	IsSyncReplica *int64 `json:"IsSyncReplica,omitnil,omitempty" name:"IsSyncReplica"`

	// 同步的原始地域
	SourceRegion *string `json:"SourceRegion,omitnil,omitempty" name:"SourceRegion"`

	// 密钥同步的状态，0:未同步,1:同步成功,2:同步失败,3:同步中。
	SyncStatus *int64 `json:"SyncStatus,omitnil,omitempty" name:"SyncStatus"`

	// 同步的结果描述
	SyncMessages *string `json:"SyncMessages,omitnil,omitempty" name:"SyncMessages"`

	// 同步的开始时间
	SyncStartTime *uint64 `json:"SyncStartTime,omitnil,omitempty" name:"SyncStartTime"`

	// 同步的结束时间
	SyncEndTime *uint64 `json:"SyncEndTime,omitnil,omitempty" name:"SyncEndTime"`

	// 同步的原始集群，如果为空，是公有云公共集群
	SourceHsmClusterId *string `json:"SourceHsmClusterId,omitnil,omitempty" name:"SourceHsmClusterId"`
}

// Predefined struct for user
type ListAlgorithmsRequestParams struct {

}

type ListAlgorithmsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListAlgorithmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAlgorithmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAlgorithmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAlgorithmsResponseParams struct {
	// 本地区支持的对称加密算法
	SymmetricAlgorithms []*AlgorithmInfo `json:"SymmetricAlgorithms,omitnil,omitempty" name:"SymmetricAlgorithms"`

	// 本地区支持的非对称加密算法
	AsymmetricAlgorithms []*AlgorithmInfo `json:"AsymmetricAlgorithms,omitnil,omitempty" name:"AsymmetricAlgorithms"`

	// 本地区支持的非对称签名验签算法
	AsymmetricSignVerifyAlgorithms []*AlgorithmInfo `json:"AsymmetricSignVerifyAlgorithms,omitnil,omitempty" name:"AsymmetricSignVerifyAlgorithms"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAlgorithmsResponse struct {
	*tchttp.BaseResponse
	Response *ListAlgorithmsResponseParams `json:"Response"`
}

func (r *ListAlgorithmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAlgorithmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDataKeyDetailRequestParams struct {
	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为10，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据创建者角色筛选，默认 0 表示用户自己创建的数据密钥， 1 表示授权其它云产品自动创建的数据密钥
	Role *uint64 `json:"Role,omitnil,omitempty" name:"Role"`

	// 根据DataKey创建时间排序， 0 表示按照降序排序，1表示按照升序排序
	OrderType *uint64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 根据DataKey状态筛选， 0表示全部DataKey， 1 表示仅查询Enabled DataKey， 2 表示仅查询Disabled DataKey，3 表示查询PendingDelete 状态的DataKey(处于计划删除状态的Key)。
	KeyState *uint64 `json:"KeyState,omitnil,omitempty" name:"KeyState"`

	// 根据DataKeyId或者DataKeyName进行模糊匹配查询
	SearchKeyAlias *string `json:"SearchKeyAlias,omitnil,omitempty" name:"SearchKeyAlias"`

	// 根据DateKey类型筛选， "TENCENT_KMS" 表示筛选密钥材料由KMS创建的数据密钥， "EXTERNAL" 表示筛选密钥材料需要用户导入的 EXTERNAL类型数据密钥，"ALL" 或者不设置表示两种类型都查询，大小写敏感。
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// KMS 高级版对应的 HSM 集群 ID。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`

	// 根密钥全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 数据密钥的长度
	DataKeyLen *uint64 `json:"DataKeyLen,omitnil,omitempty" name:"DataKeyLen"`
}

type ListDataKeyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为10，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据创建者角色筛选，默认 0 表示用户自己创建的数据密钥， 1 表示授权其它云产品自动创建的数据密钥
	Role *uint64 `json:"Role,omitnil,omitempty" name:"Role"`

	// 根据DataKey创建时间排序， 0 表示按照降序排序，1表示按照升序排序
	OrderType *uint64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 根据DataKey状态筛选， 0表示全部DataKey， 1 表示仅查询Enabled DataKey， 2 表示仅查询Disabled DataKey，3 表示查询PendingDelete 状态的DataKey(处于计划删除状态的Key)。
	KeyState *uint64 `json:"KeyState,omitnil,omitempty" name:"KeyState"`

	// 根据DataKeyId或者DataKeyName进行模糊匹配查询
	SearchKeyAlias *string `json:"SearchKeyAlias,omitnil,omitempty" name:"SearchKeyAlias"`

	// 根据DateKey类型筛选， "TENCENT_KMS" 表示筛选密钥材料由KMS创建的数据密钥， "EXTERNAL" 表示筛选密钥材料需要用户导入的 EXTERNAL类型数据密钥，"ALL" 或者不设置表示两种类型都查询，大小写敏感。
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// KMS 高级版对应的 HSM 集群 ID。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`

	// 根密钥全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 数据密钥的长度
	DataKeyLen *uint64 `json:"DataKeyLen,omitnil,omitempty" name:"DataKeyLen"`
}

func (r *ListDataKeyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDataKeyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Role")
	delete(f, "OrderType")
	delete(f, "KeyState")
	delete(f, "SearchKeyAlias")
	delete(f, "Origin")
	delete(f, "HsmClusterId")
	delete(f, "KeyId")
	delete(f, "DataKeyLen")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDataKeyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDataKeyDetailResponseParams struct {
	// 返回的属性信息列表。
	DataKeyMetadatas []*DataKeyMetadata `json:"DataKeyMetadatas,omitnil,omitempty" name:"DataKeyMetadatas"`

	// DataKey的总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDataKeyDetailResponse struct {
	*tchttp.BaseResponse
	Response *ListDataKeyDetailResponseParams `json:"Response"`
}

func (r *ListDataKeyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDataKeyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDataKeysRequestParams struct {
	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为10，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据创建者角色筛选，默认 0 表示用户自己创建的数据密钥， 1 表示授权其它云产品自动创建的数据密钥
	Role *uint64 `json:"Role,omitnil,omitempty" name:"Role"`

	// KMS 高级版对应的 HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

type ListDataKeysRequest struct {
	*tchttp.BaseRequest
	
	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为10，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据创建者角色筛选，默认 0 表示用户自己创建的数据密钥， 1 表示授权其它云产品自动创建的数据密钥
	Role *uint64 `json:"Role,omitnil,omitempty" name:"Role"`

	// KMS 高级版对应的 HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

func (r *ListDataKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDataKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Role")
	delete(f, "HsmClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDataKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDataKeysResponseParams struct {
	// 数据密钥Id列表数组
	DataKeys []*DataKey `json:"DataKeys,omitnil,omitempty" name:"DataKeys"`

	// 数据密钥的总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDataKeysResponse struct {
	*tchttp.BaseResponse
	Response *ListDataKeysResponseParams `json:"Response"`
}

func (r *ListDataKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDataKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListKeyDetailRequestParams struct {
	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为10，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据创建者角色筛选，默认 0 表示用户自己创建的cmk， 1 表示授权其它云产品自动创建的cmk
	Role *uint64 `json:"Role,omitnil,omitempty" name:"Role"`

	// 根据CMK创建时间排序， 0 表示按照降序排序，1表示按照升序排序
	OrderType *uint64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 根据CMK状态筛选， 0表示全部CMK， 1 表示仅查询Enabled CMK， 2 表示仅查询Disabled CMK，3 表示查询PendingDelete 状态的CMK(处于计划删除状态的Key)，4 表示查询 PendingImport 状态的CMK，5 表示查询 Archived 状态的 CMK
	KeyState *uint64 `json:"KeyState,omitnil,omitempty" name:"KeyState"`

	// 根据KeyId或者Alias进行模糊匹配查询
	SearchKeyAlias *string `json:"SearchKeyAlias,omitnil,omitempty" name:"SearchKeyAlias"`

	// 根据CMK类型筛选， "TENCENT_KMS" 表示筛选密钥材料由KMS创建的CMK， "EXTERNAL" 表示筛选密钥材料需要用户导入的 EXTERNAL类型CMK，"ALL" 或者不设置表示两种类型都查询，大小写敏感。
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 根据CMK的KeyUsage筛选，ALL表示筛选全部，可使用的参数为：ALL 或 ENCRYPT_DECRYPT 或 ASYMMETRIC_DECRYPT_RSA_2048 或 ASYMMETRIC_DECRYPT_SM2 或 ASYMMETRIC_SIGN_VERIFY_SM2 或 ASYMMETRIC_SIGN_VERIFY_RSA_2048 或 ASYMMETRIC_SIGN_VERIFY_ECC，为空则默认筛选ENCRYPT_DECRYPT类型
	KeyUsage *string `json:"KeyUsage,omitnil,omitempty" name:"KeyUsage"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// KMS 高级版对应的 HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

type ListKeyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为10，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据创建者角色筛选，默认 0 表示用户自己创建的cmk， 1 表示授权其它云产品自动创建的cmk
	Role *uint64 `json:"Role,omitnil,omitempty" name:"Role"`

	// 根据CMK创建时间排序， 0 表示按照降序排序，1表示按照升序排序
	OrderType *uint64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 根据CMK状态筛选， 0表示全部CMK， 1 表示仅查询Enabled CMK， 2 表示仅查询Disabled CMK，3 表示查询PendingDelete 状态的CMK(处于计划删除状态的Key)，4 表示查询 PendingImport 状态的CMK，5 表示查询 Archived 状态的 CMK
	KeyState *uint64 `json:"KeyState,omitnil,omitempty" name:"KeyState"`

	// 根据KeyId或者Alias进行模糊匹配查询
	SearchKeyAlias *string `json:"SearchKeyAlias,omitnil,omitempty" name:"SearchKeyAlias"`

	// 根据CMK类型筛选， "TENCENT_KMS" 表示筛选密钥材料由KMS创建的CMK， "EXTERNAL" 表示筛选密钥材料需要用户导入的 EXTERNAL类型CMK，"ALL" 或者不设置表示两种类型都查询，大小写敏感。
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 根据CMK的KeyUsage筛选，ALL表示筛选全部，可使用的参数为：ALL 或 ENCRYPT_DECRYPT 或 ASYMMETRIC_DECRYPT_RSA_2048 或 ASYMMETRIC_DECRYPT_SM2 或 ASYMMETRIC_SIGN_VERIFY_SM2 或 ASYMMETRIC_SIGN_VERIFY_RSA_2048 或 ASYMMETRIC_SIGN_VERIFY_ECC，为空则默认筛选ENCRYPT_DECRYPT类型
	KeyUsage *string `json:"KeyUsage,omitnil,omitempty" name:"KeyUsage"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// KMS 高级版对应的 HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

func (r *ListKeyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Role")
	delete(f, "OrderType")
	delete(f, "KeyState")
	delete(f, "SearchKeyAlias")
	delete(f, "Origin")
	delete(f, "KeyUsage")
	delete(f, "TagFilters")
	delete(f, "HsmClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListKeyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListKeyDetailResponseParams struct {
	// CMK的总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的属性信息列表。
	KeyMetadatas []*KeyMetadata `json:"KeyMetadatas,omitnil,omitempty" name:"KeyMetadatas"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListKeyDetailResponse struct {
	*tchttp.BaseResponse
	Response *ListKeyDetailResponseParams `json:"Response"`
}

func (r *ListKeyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListKeysRequestParams struct {
	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次获最多获取 Limit 个元素。缺省值为10，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据创建者角色筛选，默认 0 表示用户自己创建的cmk， 1 表示授权其它云产品自动创建的cmk
	Role *uint64 `json:"Role,omitnil,omitempty" name:"Role"`

	// KMS 高级版对应的 HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

type ListKeysRequest struct {
	*tchttp.BaseRequest
	
	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次获最多获取 Limit 个元素。缺省值为10，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据创建者角色筛选，默认 0 表示用户自己创建的cmk， 1 表示授权其它云产品自动创建的cmk
	Role *uint64 `json:"Role,omitnil,omitempty" name:"Role"`

	// KMS 高级版对应的 HSM 集群 ID（仅对 KMS 独占版/托管版服务实例有效）。
	HsmClusterId *string `json:"HsmClusterId,omitnil,omitempty" name:"HsmClusterId"`
}

func (r *ListKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Role")
	delete(f, "HsmClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListKeysResponseParams struct {
	// CMK列表数组
	Keys []*Key `json:"Keys,omitnil,omitempty" name:"Keys"`

	// CMK的总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListKeysResponse struct {
	*tchttp.BaseResponse
	Response *ListKeysResponseParams `json:"Response"`
}

func (r *ListKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OverwriteWhiteBoxDeviceFingerprintsRequestParams struct {
	// 白盒密钥ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 设备指纹列表，如果列表为空，则表示删除该密钥对应的所有指纹信息。列表最大长度不超过200。
	DeviceFingerprints []*DeviceFingerprint `json:"DeviceFingerprints,omitnil,omitempty" name:"DeviceFingerprints"`
}

type OverwriteWhiteBoxDeviceFingerprintsRequest struct {
	*tchttp.BaseRequest
	
	// 白盒密钥ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 设备指纹列表，如果列表为空，则表示删除该密钥对应的所有指纹信息。列表最大长度不超过200。
	DeviceFingerprints []*DeviceFingerprint `json:"DeviceFingerprints,omitnil,omitempty" name:"DeviceFingerprints"`
}

func (r *OverwriteWhiteBoxDeviceFingerprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OverwriteWhiteBoxDeviceFingerprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "DeviceFingerprints")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OverwriteWhiteBoxDeviceFingerprintsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OverwriteWhiteBoxDeviceFingerprintsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OverwriteWhiteBoxDeviceFingerprintsResponse struct {
	*tchttp.BaseResponse
	Response *OverwriteWhiteBoxDeviceFingerprintsResponseParams `json:"Response"`
}

func (r *OverwriteWhiteBoxDeviceFingerprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OverwriteWhiteBoxDeviceFingerprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PostQuantumCryptoDecryptRequestParams struct {
	// 待解密的密文数据
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// PEM 格式公钥字符串，支持 RSA2048 和 SM2 公钥，用于对返回数据中的 Plaintext 值进行加密。若为空，则不对 Plaintext 值加密。
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitnil,omitempty" name:"EncryptionPublicKey"`

	// 非对称加密算法，配合 EncryptionPublicKey 对返回数据进行加密。目前支持：SM2（以 C1C3C2 格式返回密文），SM2_C1C3C2_ASN1 （以 C1C3C2 ASN1 格式返回密文），RSAES_PKCS1_V1_5，RSAES_OAEP_SHA_1，RSAES_OAEP_SHA_256。若为空，则默认为 SM2。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitnil,omitempty" name:"EncryptionAlgorithm"`
}

type PostQuantumCryptoDecryptRequest struct {
	*tchttp.BaseRequest
	
	// 待解密的密文数据
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// PEM 格式公钥字符串，支持 RSA2048 和 SM2 公钥，用于对返回数据中的 Plaintext 值进行加密。若为空，则不对 Plaintext 值加密。
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitnil,omitempty" name:"EncryptionPublicKey"`

	// 非对称加密算法，配合 EncryptionPublicKey 对返回数据进行加密。目前支持：SM2（以 C1C3C2 格式返回密文），SM2_C1C3C2_ASN1 （以 C1C3C2 ASN1 格式返回密文），RSAES_PKCS1_V1_5，RSAES_OAEP_SHA_1，RSAES_OAEP_SHA_256。若为空，则默认为 SM2。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitnil,omitempty" name:"EncryptionAlgorithm"`
}

func (r *PostQuantumCryptoDecryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostQuantumCryptoDecryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CiphertextBlob")
	delete(f, "EncryptionPublicKey")
	delete(f, "EncryptionAlgorithm")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PostQuantumCryptoDecryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PostQuantumCryptoDecryptResponseParams struct {
	// CMK的全局唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 若调用时未提供 EncryptionPublicKey，该字段值为 Base64 编码的明文，需进行 Base64 解码以获取明文。
	// 若调用时提供了 EncryptionPublicKey，则该字段值为使用 EncryptionPublicKey 公钥进行非对称加密后的 Base64 编码的密文。需在 Base64 解码后，使用用户上传的公钥对应的私钥进行进一步解密，以获取明文。
	PlainText *string `json:"PlainText,omitnil,omitempty" name:"PlainText"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PostQuantumCryptoDecryptResponse struct {
	*tchttp.BaseResponse
	Response *PostQuantumCryptoDecryptResponseParams `json:"Response"`
}

func (r *PostQuantumCryptoDecryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostQuantumCryptoDecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PostQuantumCryptoEncryptRequestParams struct {
	// 调用CreateKey生成的CMK全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 被加密的明文数据，该字段必须使用base64编码，原文最大长度支持4K
	PlainText *string `json:"PlainText,omitnil,omitempty" name:"PlainText"`
}

type PostQuantumCryptoEncryptRequest struct {
	*tchttp.BaseRequest
	
	// 调用CreateKey生成的CMK全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 被加密的明文数据，该字段必须使用base64编码，原文最大长度支持4K
	PlainText *string `json:"PlainText,omitnil,omitempty" name:"PlainText"`
}

func (r *PostQuantumCryptoEncryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostQuantumCryptoEncryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "PlainText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PostQuantumCryptoEncryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PostQuantumCryptoEncryptResponseParams struct {
	// 加密后的密文，base64编码。注意：本字段中打包了密文和密钥的相关信息，不是对明文的直接加密结果，只有将该字段作为PostQuantumCryptoDecrypt接口的输入参数，才可以解密出原文。
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// 加密使用的CMK的全局唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PostQuantumCryptoEncryptResponse struct {
	*tchttp.BaseResponse
	Response *PostQuantumCryptoEncryptResponseParams `json:"Response"`
}

func (r *PostQuantumCryptoEncryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostQuantumCryptoEncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PostQuantumCryptoSignRequestParams struct {
	// Base64 编码的消息原文。消息原文的长度（Base64编码前的长度）不超过4096字节。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 密钥的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type PostQuantumCryptoSignRequest struct {
	*tchttp.BaseRequest
	
	// Base64 编码的消息原文。消息原文的长度（Base64编码前的长度）不超过4096字节。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 密钥的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *PostQuantumCryptoSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostQuantumCryptoSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Message")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PostQuantumCryptoSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PostQuantumCryptoSignResponseParams struct {
	// 签名值，Base64编码。可使用 PostQuantumCryptoVerify接口对签名值进行验证。
	Signature *string `json:"Signature,omitnil,omitempty" name:"Signature"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PostQuantumCryptoSignResponse struct {
	*tchttp.BaseResponse
	Response *PostQuantumCryptoSignResponseParams `json:"Response"`
}

func (r *PostQuantumCryptoSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostQuantumCryptoSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PostQuantumCryptoVerifyRequestParams struct {
	// 密钥的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 签名值，通过调用KMS PostQuantumCryptoSign签名接口生成
	SignatureValue *string `json:"SignatureValue,omitnil,omitempty" name:"SignatureValue"`

	// Base64 编码的消息原文，消息原文的长度（Base64编码前的长度）不超过4096字节。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type PostQuantumCryptoVerifyRequest struct {
	*tchttp.BaseRequest
	
	// 密钥的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 签名值，通过调用KMS PostQuantumCryptoSign签名接口生成
	SignatureValue *string `json:"SignatureValue,omitnil,omitempty" name:"SignatureValue"`

	// Base64 编码的消息原文，消息原文的长度（Base64编码前的长度）不超过4096字节。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

func (r *PostQuantumCryptoVerifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostQuantumCryptoVerifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "SignatureValue")
	delete(f, "Message")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PostQuantumCryptoVerifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PostQuantumCryptoVerifyResponseParams struct {
	// 签名是否有效。true：签名有效，false：签名无效。
	SignatureValid *bool `json:"SignatureValid,omitnil,omitempty" name:"SignatureValid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PostQuantumCryptoVerifyResponse struct {
	*tchttp.BaseResponse
	Response *PostQuantumCryptoVerifyResponseParams `json:"Response"`
}

func (r *PostQuantumCryptoVerifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostQuantumCryptoVerifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReEncryptRequestParams struct {
	// 需要重新加密的密文
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// 重新加密使用的CMK，如果为空，则使用密文原有的CMK重新加密（若密钥没有轮换则密文不会刷新）
	DestinationKeyId *string `json:"DestinationKeyId,omitnil,omitempty" name:"DestinationKeyId"`

	// CiphertextBlob 密文加密时使用的key/value对的json字符串。如果加密时未使用，则为空
	SourceEncryptionContext *string `json:"SourceEncryptionContext,omitnil,omitempty" name:"SourceEncryptionContext"`

	// 重新加密使用的key/value对的json字符串，如果使用该字段，则返回的新密文在解密时需要填入相同的字符串
	DestinationEncryptionContext *string `json:"DestinationEncryptionContext,omitnil,omitempty" name:"DestinationEncryptionContext"`
}

type ReEncryptRequest struct {
	*tchttp.BaseRequest
	
	// 需要重新加密的密文
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// 重新加密使用的CMK，如果为空，则使用密文原有的CMK重新加密（若密钥没有轮换则密文不会刷新）
	DestinationKeyId *string `json:"DestinationKeyId,omitnil,omitempty" name:"DestinationKeyId"`

	// CiphertextBlob 密文加密时使用的key/value对的json字符串。如果加密时未使用，则为空
	SourceEncryptionContext *string `json:"SourceEncryptionContext,omitnil,omitempty" name:"SourceEncryptionContext"`

	// 重新加密使用的key/value对的json字符串，如果使用该字段，则返回的新密文在解密时需要填入相同的字符串
	DestinationEncryptionContext *string `json:"DestinationEncryptionContext,omitnil,omitempty" name:"DestinationEncryptionContext"`
}

func (r *ReEncryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReEncryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CiphertextBlob")
	delete(f, "DestinationKeyId")
	delete(f, "SourceEncryptionContext")
	delete(f, "DestinationEncryptionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReEncryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReEncryptResponseParams struct {
	// 重新加密后的密文
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// 重新加密使用的CMK
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 重新加密前密文使用的CMK
	SourceKeyId *string `json:"SourceKeyId,omitnil,omitempty" name:"SourceKeyId"`

	// true表示密文已经重新加密。同一个CMK进行重加密，在密钥没有发生轮换的情况下不会进行实际重新加密操作，返回原密文
	ReEncrypted *bool `json:"ReEncrypted,omitnil,omitempty" name:"ReEncrypted"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReEncryptResponse struct {
	*tchttp.BaseResponse
	Response *ReEncryptResponseParams `json:"Response"`
}

func (r *ReEncryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReEncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScheduleDataKeyDeletionRequestParams struct {
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// 计划删除时间区间[7,30]
	PendingWindowInDays *uint64 `json:"PendingWindowInDays,omitnil,omitempty" name:"PendingWindowInDays"`
}

type ScheduleDataKeyDeletionRequest struct {
	*tchttp.BaseRequest
	
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// 计划删除时间区间[7,30]
	PendingWindowInDays *uint64 `json:"PendingWindowInDays,omitnil,omitempty" name:"PendingWindowInDays"`
}

func (r *ScheduleDataKeyDeletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScheduleDataKeyDeletionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyId")
	delete(f, "PendingWindowInDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScheduleDataKeyDeletionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScheduleDataKeyDeletionResponseParams struct {
	// 计划删除执行时间
	DeletionDate *uint64 `json:"DeletionDate,omitnil,omitempty" name:"DeletionDate"`

	// 唯一标志被计划删除的数据密钥
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScheduleDataKeyDeletionResponse struct {
	*tchttp.BaseResponse
	Response *ScheduleDataKeyDeletionResponseParams `json:"Response"`
}

func (r *ScheduleDataKeyDeletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScheduleDataKeyDeletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScheduleKeyDeletionRequestParams struct {
	// CMK的唯一标志
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 计划删除时间区间[7,30]
	PendingWindowInDays *uint64 `json:"PendingWindowInDays,omitnil,omitempty" name:"PendingWindowInDays"`
}

type ScheduleKeyDeletionRequest struct {
	*tchttp.BaseRequest
	
	// CMK的唯一标志
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 计划删除时间区间[7,30]
	PendingWindowInDays *uint64 `json:"PendingWindowInDays,omitnil,omitempty" name:"PendingWindowInDays"`
}

func (r *ScheduleKeyDeletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScheduleKeyDeletionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "PendingWindowInDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScheduleKeyDeletionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScheduleKeyDeletionResponseParams struct {
	// 计划删除执行时间
	DeletionDate *uint64 `json:"DeletionDate,omitnil,omitempty" name:"DeletionDate"`

	// 唯一标志被计划删除的CMK
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScheduleKeyDeletionResponse struct {
	*tchttp.BaseResponse
	Response *ScheduleKeyDeletionResponseParams `json:"Response"`
}

func (r *ScheduleKeyDeletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScheduleKeyDeletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignByAsymmetricKeyRequestParams struct {
	// 签名算法，支持的算法：SM2DSA，ECC_P256_R1，RSA_PSS_SHA_256，RSA_PKCS1_SHA_256 等。更多支持的算法可通过 ListAlgorithms 接口进行查询。
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 消息原文或消息摘要。如果提供的是消息原文，则消息原文的长度（Base64编码前的长度）不超过4096字节。如果提供的是消息摘要，消息摘要长度（Base64编码前的长度）必须等于32字节
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 密钥的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 消息类型：RAW，DIGEST，如果不传，默认为RAW，表示消息原文。
	MessageType *string `json:"MessageType,omitnil,omitempty" name:"MessageType"`
}

type SignByAsymmetricKeyRequest struct {
	*tchttp.BaseRequest
	
	// 签名算法，支持的算法：SM2DSA，ECC_P256_R1，RSA_PSS_SHA_256，RSA_PKCS1_SHA_256 等。更多支持的算法可通过 ListAlgorithms 接口进行查询。
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 消息原文或消息摘要。如果提供的是消息原文，则消息原文的长度（Base64编码前的长度）不超过4096字节。如果提供的是消息摘要，消息摘要长度（Base64编码前的长度）必须等于32字节
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 密钥的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 消息类型：RAW，DIGEST，如果不传，默认为RAW，表示消息原文。
	MessageType *string `json:"MessageType,omitnil,omitempty" name:"MessageType"`
}

func (r *SignByAsymmetricKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignByAsymmetricKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Algorithm")
	delete(f, "Message")
	delete(f, "KeyId")
	delete(f, "MessageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SignByAsymmetricKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignByAsymmetricKeyResponseParams struct {
	// 签名，Base64编码
	Signature *string `json:"Signature,omitnil,omitempty" name:"Signature"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SignByAsymmetricKeyResponse struct {
	*tchttp.BaseResponse
	Response *SignByAsymmetricKeyResponseParams `json:"Response"`
}

func (r *SignByAsymmetricKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignByAsymmetricKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UnbindCloudResourceRequestParams struct {
	// cmk的ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 云产品的唯一性标识符
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 资源/实例ID，由调用方根据自己的云产品特征来定义，以字符串形式做存储。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type UnbindCloudResourceRequest struct {
	*tchttp.BaseRequest
	
	// cmk的ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 云产品的唯一性标识符
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 资源/实例ID，由调用方根据自己的云产品特征来定义，以字符串形式做存储。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *UnbindCloudResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCloudResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "ProductId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindCloudResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindCloudResourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindCloudResourceResponse struct {
	*tchttp.BaseResponse
	Response *UnbindCloudResourceResponseParams `json:"Response"`
}

func (r *UnbindCloudResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCloudResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAliasRequestParams struct {
	// 新的别名，1-60个字符或数字的组合
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// CMK的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type UpdateAliasRequest struct {
	*tchttp.BaseRequest
	
	// 新的别名，1-60个字符或数字的组合
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// CMK的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *UpdateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alias")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAliasResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAliasResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAliasResponseParams `json:"Response"`
}

func (r *UpdateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataKeyDescriptionRequestParams struct {
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// 数据密钥 的描述，最大100字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateDataKeyDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// 数据密钥 的描述，最大100字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateDataKeyDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataKeyDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDataKeyDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataKeyDescriptionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDataKeyDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDataKeyDescriptionResponseParams `json:"Response"`
}

func (r *UpdateDataKeyDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataKeyDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataKeyNameRequestParams struct {
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// 数据密钥的名称
	DataKeyName *string `json:"DataKeyName,omitnil,omitempty" name:"DataKeyName"`
}

type UpdateDataKeyNameRequest struct {
	*tchttp.BaseRequest
	
	// 数据密钥的唯一标志符
	DataKeyId *string `json:"DataKeyId,omitnil,omitempty" name:"DataKeyId"`

	// 数据密钥的名称
	DataKeyName *string `json:"DataKeyName,omitnil,omitempty" name:"DataKeyName"`
}

func (r *UpdateDataKeyNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataKeyNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataKeyId")
	delete(f, "DataKeyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDataKeyNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataKeyNameResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDataKeyNameResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDataKeyNameResponseParams `json:"Response"`
}

func (r *UpdateDataKeyNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataKeyNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateKeyDescriptionRequestParams struct {
	// 新的描述信息，最大支持1024字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 需要修改描述信息的CMK ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type UpdateKeyDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 新的描述信息，最大支持1024字节
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 需要修改描述信息的CMK ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *UpdateKeyDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateKeyDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Description")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateKeyDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateKeyDescriptionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateKeyDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateKeyDescriptionResponseParams `json:"Response"`
}

func (r *UpdateKeyDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateKeyDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyByAsymmetricKeyRequestParams struct {
	// 密钥的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 签名值，通过调用KMS签名接口生成
	SignatureValue *string `json:"SignatureValue,omitnil,omitempty" name:"SignatureValue"`

	// 消息原文或消息摘要。如果提供的是消息原文，则消息原文的长度（Base64编码前的长度）不超过4096字节。如果提供的是消息摘要，则消息摘要长度（Base64编码前的长度）必须等于32字节
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 签名算法，支持的算法：SM2DSA，ECC_P256_R1，RSA_PSS_SHA_256，RSA_PKCS1_SHA_256 等。更多支持的算法可通过 ListAlgorithms 接口进行查询。
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 消息类型：RAW，DIGEST，如果不传，默认为RAW，表示消息原文。
	MessageType *string `json:"MessageType,omitnil,omitempty" name:"MessageType"`
}

type VerifyByAsymmetricKeyRequest struct {
	*tchttp.BaseRequest
	
	// 密钥的唯一标识
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 签名值，通过调用KMS签名接口生成
	SignatureValue *string `json:"SignatureValue,omitnil,omitempty" name:"SignatureValue"`

	// 消息原文或消息摘要。如果提供的是消息原文，则消息原文的长度（Base64编码前的长度）不超过4096字节。如果提供的是消息摘要，则消息摘要长度（Base64编码前的长度）必须等于32字节
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 签名算法，支持的算法：SM2DSA，ECC_P256_R1，RSA_PSS_SHA_256，RSA_PKCS1_SHA_256 等。更多支持的算法可通过 ListAlgorithms 接口进行查询。
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 消息类型：RAW，DIGEST，如果不传，默认为RAW，表示消息原文。
	MessageType *string `json:"MessageType,omitnil,omitempty" name:"MessageType"`
}

func (r *VerifyByAsymmetricKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyByAsymmetricKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "SignatureValue")
	delete(f, "Message")
	delete(f, "Algorithm")
	delete(f, "MessageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyByAsymmetricKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyByAsymmetricKeyResponseParams struct {
	// 签名是否有效。true：签名有效，false：签名无效。
	SignatureValid *bool `json:"SignatureValid,omitnil,omitempty" name:"SignatureValid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyByAsymmetricKeyResponse struct {
	*tchttp.BaseResponse
	Response *VerifyByAsymmetricKeyResponseParams `json:"Response"`
}

func (r *VerifyByAsymmetricKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyByAsymmetricKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WhiteboxKeyInfo struct {
	// 白盒密钥的全局唯一标识符
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 作为密钥更容易辨识，更容易被人看懂的别名， 不可为空，1-60个字母数字 - _ 的组合，首字符必须为字母或者数字. 不可重复
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 创建者
	CreatorUin *uint64 `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 密钥的描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 密钥创建时间，Unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 白盒密钥的状态， 取值为：Enabled | Disabled
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建者
	OwnerUin *uint64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 密钥所用的算法类型
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 白盒加密密钥，base64编码
	EncryptKey *string `json:"EncryptKey,omitnil,omitempty" name:"EncryptKey"`

	// 白盒解密密钥，base64编码
	DecryptKey *string `json:"DecryptKey,omitnil,omitempty" name:"DecryptKey"`

	// 资源ID，格式：creatorUin/$creatorUin/$keyId
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 是否有设备指纹与当前密钥绑定
	DeviceFingerprintBind *bool `json:"DeviceFingerprintBind,omitnil,omitempty" name:"DeviceFingerprintBind"`
}