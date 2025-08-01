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

package v20230901

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ActivateServiceRequestParams struct {
	// 开通之后，是否关闭后付费；默认为0，不关闭；1为关闭
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type ActivateServiceRequest struct {
	*tchttp.BaseRequest
	
	// 开通之后，是否关闭后付费；默认为0，不关闭；1为关闭
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

func (r *ActivateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActivateServiceResponse struct {
	*tchttp.BaseResponse
	Response *ActivateServiceResponseParams `json:"Response"`
}

func (r *ActivateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Approximate struct {
	// 表示 ISO 国家代码
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 表示城市名称
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 表示区域名称
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 表示IANA时区
	Timezone *string `json:"Timezone,omitnil,omitempty" name:"Timezone"`

	// 表示详细地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

type Character struct {
	// 人物名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 人物对应SystemPrompt
	SystemPrompt *string `json:"SystemPrompt,omitnil,omitempty" name:"SystemPrompt"`
}

// Predefined struct for user
type ChatCompletionsRequestParams struct {
	// 模型名称，可选值参考 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中混元生文模型列表。
	// 示例值：hunyuan-turbos-latest
	// 各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。注意：不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message.Role 可选值：system、user、assistant、 tool（functioncall场景）。
	// 其中，system 角色可选，如存在则必须位于列表的最开始。user（tool） 和 assistant 需交替出现（一问一答），以 user 提问开始，user（tool）提问结束，其中tool可以连续出现多次，且 Content 不能为空。Role 的顺序示例：[system（可选） user assistant user（tool tool ...） assistant user（tool tool ...） ...]。
	// 3. Messages 中 Content 总长度不能超过模型输入长度上限（可参考 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 文档），超过则会截断最前面的内容，只保留尾部内容。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为非流式调用（false）。
	// 2. 流式调用时以 SSE 协议增量返回结果（返回值取 Choices[n].Delta 中的值，需要拼接增量数据才能获得完整结果）。
	// 3. 非流式调用时：
	// 调用方式与普通 HTTP 请求无异。
	// 接口响应耗时较长，**如需更低时延建议设置为 true**。
	// 只返回一次最终结果（返回值取 Choices[n].Message 中的值）。
	// 
	// 注意：
	// 通过 SDK 调用时，流式和非流式调用需用**不同的方式**获取返回值，具体参考 SDK 中的注释或示例（在各语言 SDK 代码仓库的 examples/hunyuan/v20230901/ 目录中）。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 流式输出审核开关。
	// 说明：
	// 1. 当使用流式输出（Stream 字段值为 true）时，该字段生效。
	// 2. 输出审核有流式和同步两种模式，**流式模式首包响应更快**。未传值时默认为流式模式（true）。
	// 3. 如果值为 true，将对输出内容进行分段审核，审核通过的内容流式输出返回。如果出现审核不过，响应中的 FinishReason 值为 sensitive。
	// 4. 如果值为 false，则不使用流式输出审核，需要审核完所有输出内容后再返回结果。
	// 
	// 注意：
	// 当选择流式输出审核时，可能会出现部分内容已输出，但中间某一段响应中的 FinishReason 值为 sensitive，此时说明安全审核未通过。如果业务场景有实时文字上屏的需求，需要自行撤回已上屏的内容，并建议自定义替换为一条提示语，如 “这个问题我不方便回答，不如我们换个话题试试”，以保障终端体验。
	StreamModeration *bool `json:"StreamModeration,omitnil,omitempty" name:"StreamModeration"`

	// 说明：
	// 1. 影响输出文本的多样性。模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。
	// 2. 取值区间为 [0.0, 1.0]。取值越大，生成文本的多样性越强。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 说明：
	// 1. 影响模型输出多样性，模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。
	// 2. 取值区间为 [0.0, 2.0]。较高的数值会使输出更加多样化和不可预测，而较低的数值会使其更加集中和确定。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 功能增强（如搜索）开关。
	// 说明：
	// 1. hunyuan-lite 无功能增强（如搜索）能力，该参数对 hunyuan-lite 版本不生效。
	// 2. 未传值时默认关闭开关。
	// 3. 关闭时将直接由主模型生成回复内容，可以降低响应时延（对于流式输出时的首字时延尤为明显）。但在少数场景里，回复效果可能会下降。
	// 4. 安全审核能力不属于功能增强范围，不受此字段影响。
	// 5. 2025-04-20 00:00:00起，由默认开启状态转为默认关闭状态。
	EnableEnhancement *bool `json:"EnableEnhancement,omitnil,omitempty" name:"EnableEnhancement"`

	// 可调用的工具列表，仅对 hunyuan-turbos、hunyuan-t1、hunyuan-functioncall 模型生效。
	Tools []*Tool `json:"Tools,omitnil,omitempty" name:"Tools"`

	// 工具使用选项，可选值包括 none、auto、custom。说明：1. 仅对 hunyuan-turbo、hunyuan-functioncall 模型生效。2. none：不调用工具；auto：模型自行选择生成回复或调用工具；custom：强制模型调用指定的工具。3. 未设置时，默认值为auto
	ToolChoice *string `json:"ToolChoice,omitnil,omitempty" name:"ToolChoice"`

	// 强制模型调用指定的工具，当参数ToolChoice为custom时，此参数为必填
	CustomTool *Tool `json:"CustomTool,omitnil,omitempty" name:"CustomTool"`

	// 默认是false，在值为true且命中搜索时，接口会返回SearchInfo
	SearchInfo *bool `json:"SearchInfo,omitnil,omitempty" name:"SearchInfo"`

	// 搜索引文角标开关。
	// 说明：
	// 1. 配合EnableEnhancement和SearchInfo参数使用。打开后，回答中命中搜索的结果会在片段后增加角标标志，对应SearchInfo列表中的链接。
	// 2. false：开关关闭，true：开关打开。
	// 3. 未传值时默认开关关闭（false）。
	Citation *bool `json:"Citation,omitnil,omitempty" name:"Citation"`

	// 是否开启极速版搜索，默认false，不开启；在开启且命中搜索时，会启用极速版搜索，流式输出首字返回更快。
	EnableSpeedSearch *bool `json:"EnableSpeedSearch,omitnil,omitempty" name:"EnableSpeedSearch"`

	// 多媒体开关。
	// 详细介绍请阅读 [多媒体介绍](https://cloud.tencent.com/document/product/1729/111178) 中的说明。
	// 说明：
	// 1. 该参数目前仅对白名单内用户生效，如您想体验该功能请 [联系我们](https://cloud.tencent.com/act/event/Online_service)。
	// 2. 该参数仅在功能增强（如搜索）开关开启（EnableEnhancement=true）并且极速版搜索开关关闭（EnableSpeedSearch=false）时生效。
	// 3. hunyuan-lite 无多媒体能力，该参数对 hunyuan-lite 版本不生效。
	// 4. 未传值时默认关闭。
	// 5. 开启并搜索到对应的多媒体信息时，会输出对应的多媒体地址，可以定制个性化的图文消息。
	EnableMultimedia *bool `json:"EnableMultimedia,omitnil,omitempty" name:"EnableMultimedia"`

	// 是否开启深度研究该问题，默认是false，在值为true且命中深度研究该问题时，会返回深度研究该问题信息。
	EnableDeepSearch *bool `json:"EnableDeepSearch,omitnil,omitempty" name:"EnableDeepSearch"`

	// 说明： 1. 确保模型的输出是可复现的。 2. 取值区间为非0正整数，最大值10000。 3. 非必要不建议使用，不合理的取值会影响效果。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`

	// 强制搜索增强开关。
	// 说明：
	// 1. 未传值时默认关闭。
	// 2. 开启后，将强制走AI搜索，当AI搜索结果为空时，由大模型回复兜底话术。
	ForceSearchEnhancement *bool `json:"ForceSearchEnhancement,omitnil,omitempty" name:"ForceSearchEnhancement"`

	// 自定义结束生成字符串
	// 
	// 调用 OpenAI 的接口时，如果您指定了 `stop` 参数, 模型会停止在匹配到 `stop` 的内容之前。
	// 在调用混元接口时，会停止在匹配到 `stop` 的内容之后。
	// 
	// **说明：**
	// 未来我们可能会修改此行为以便和 OpenAI 保持一致。
	// 但是目前有使用该参数的情况下，开发者需要注意该参数是否会对应用造成影响，以及未来该行为调整时带来的影响。
	Stop []*string `json:"Stop,omitnil,omitempty" name:"Stop"`

	// 推荐问答开关。
	// 说明：
	// 1. 未传值时默认关闭。
	// 2. 开启后，在返回值的最后一个包中会增加 RecommendedQuestions 字段表示推荐问答， 最多返回3条。
	EnableRecommendedQuestions *bool `json:"EnableRecommendedQuestions,omitnil,omitempty" name:"EnableRecommendedQuestions"`

	// 是否开启深度阅读，默认是false，在值为true时，会返回深度阅读的结果信息。说明:1.深度阅读需要开启插件增强,即设置EnableEnhancement为true,当设置EnableDeepRead为true时EnableEnhancement默认为true；2.目前暂时只支持单文档单轮的深度阅读；3.深度阅读功能的文件上传可以使用FilesUploads接口，具体参数详见FilesUploads接口文档
	EnableDeepRead *bool `json:"EnableDeepRead,omitnil,omitempty" name:"EnableDeepRead"`

	// 知识注入相关的参数信息
	WebSearchOptions *WebSearchOptions `json:"WebSearchOptions,omitnil,omitempty" name:"WebSearchOptions"`

	// 用户传入Topic
	TopicChoice *string `json:"TopicChoice,omitnil,omitempty" name:"TopicChoice"`

	// 模型思维链开关 说明： 1. 未传值时默认开启，打开模型思维链推理能力。 2. 关闭后，关闭模型思维链推理能力。  开关当前仅对hunyuan-a13b模型生效 示例值：ture
	EnableThinking *bool `json:"EnableThinking,omitnil,omitempty" name:"EnableThinking"`
}

type ChatCompletionsRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称，可选值参考 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中混元生文模型列表。
	// 示例值：hunyuan-turbos-latest
	// 各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。注意：不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。
	// 说明：
	// 1. 长度最多为 40，按对话时间从旧到新在数组中排列。
	// 2. Message.Role 可选值：system、user、assistant、 tool（functioncall场景）。
	// 其中，system 角色可选，如存在则必须位于列表的最开始。user（tool） 和 assistant 需交替出现（一问一答），以 user 提问开始，user（tool）提问结束，其中tool可以连续出现多次，且 Content 不能为空。Role 的顺序示例：[system（可选） user assistant user（tool tool ...） assistant user（tool tool ...） ...]。
	// 3. Messages 中 Content 总长度不能超过模型输入长度上限（可参考 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 文档），超过则会截断最前面的内容，只保留尾部内容。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为非流式调用（false）。
	// 2. 流式调用时以 SSE 协议增量返回结果（返回值取 Choices[n].Delta 中的值，需要拼接增量数据才能获得完整结果）。
	// 3. 非流式调用时：
	// 调用方式与普通 HTTP 请求无异。
	// 接口响应耗时较长，**如需更低时延建议设置为 true**。
	// 只返回一次最终结果（返回值取 Choices[n].Message 中的值）。
	// 
	// 注意：
	// 通过 SDK 调用时，流式和非流式调用需用**不同的方式**获取返回值，具体参考 SDK 中的注释或示例（在各语言 SDK 代码仓库的 examples/hunyuan/v20230901/ 目录中）。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 流式输出审核开关。
	// 说明：
	// 1. 当使用流式输出（Stream 字段值为 true）时，该字段生效。
	// 2. 输出审核有流式和同步两种模式，**流式模式首包响应更快**。未传值时默认为流式模式（true）。
	// 3. 如果值为 true，将对输出内容进行分段审核，审核通过的内容流式输出返回。如果出现审核不过，响应中的 FinishReason 值为 sensitive。
	// 4. 如果值为 false，则不使用流式输出审核，需要审核完所有输出内容后再返回结果。
	// 
	// 注意：
	// 当选择流式输出审核时，可能会出现部分内容已输出，但中间某一段响应中的 FinishReason 值为 sensitive，此时说明安全审核未通过。如果业务场景有实时文字上屏的需求，需要自行撤回已上屏的内容，并建议自定义替换为一条提示语，如 “这个问题我不方便回答，不如我们换个话题试试”，以保障终端体验。
	StreamModeration *bool `json:"StreamModeration,omitnil,omitempty" name:"StreamModeration"`

	// 说明：
	// 1. 影响输出文本的多样性。模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。
	// 2. 取值区间为 [0.0, 1.0]。取值越大，生成文本的多样性越强。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 说明：
	// 1. 影响模型输出多样性，模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。
	// 2. 取值区间为 [0.0, 2.0]。较高的数值会使输出更加多样化和不可预测，而较低的数值会使其更加集中和确定。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 功能增强（如搜索）开关。
	// 说明：
	// 1. hunyuan-lite 无功能增强（如搜索）能力，该参数对 hunyuan-lite 版本不生效。
	// 2. 未传值时默认关闭开关。
	// 3. 关闭时将直接由主模型生成回复内容，可以降低响应时延（对于流式输出时的首字时延尤为明显）。但在少数场景里，回复效果可能会下降。
	// 4. 安全审核能力不属于功能增强范围，不受此字段影响。
	// 5. 2025-04-20 00:00:00起，由默认开启状态转为默认关闭状态。
	EnableEnhancement *bool `json:"EnableEnhancement,omitnil,omitempty" name:"EnableEnhancement"`

	// 可调用的工具列表，仅对 hunyuan-turbos、hunyuan-t1、hunyuan-functioncall 模型生效。
	Tools []*Tool `json:"Tools,omitnil,omitempty" name:"Tools"`

	// 工具使用选项，可选值包括 none、auto、custom。说明：1. 仅对 hunyuan-turbo、hunyuan-functioncall 模型生效。2. none：不调用工具；auto：模型自行选择生成回复或调用工具；custom：强制模型调用指定的工具。3. 未设置时，默认值为auto
	ToolChoice *string `json:"ToolChoice,omitnil,omitempty" name:"ToolChoice"`

	// 强制模型调用指定的工具，当参数ToolChoice为custom时，此参数为必填
	CustomTool *Tool `json:"CustomTool,omitnil,omitempty" name:"CustomTool"`

	// 默认是false，在值为true且命中搜索时，接口会返回SearchInfo
	SearchInfo *bool `json:"SearchInfo,omitnil,omitempty" name:"SearchInfo"`

	// 搜索引文角标开关。
	// 说明：
	// 1. 配合EnableEnhancement和SearchInfo参数使用。打开后，回答中命中搜索的结果会在片段后增加角标标志，对应SearchInfo列表中的链接。
	// 2. false：开关关闭，true：开关打开。
	// 3. 未传值时默认开关关闭（false）。
	Citation *bool `json:"Citation,omitnil,omitempty" name:"Citation"`

	// 是否开启极速版搜索，默认false，不开启；在开启且命中搜索时，会启用极速版搜索，流式输出首字返回更快。
	EnableSpeedSearch *bool `json:"EnableSpeedSearch,omitnil,omitempty" name:"EnableSpeedSearch"`

	// 多媒体开关。
	// 详细介绍请阅读 [多媒体介绍](https://cloud.tencent.com/document/product/1729/111178) 中的说明。
	// 说明：
	// 1. 该参数目前仅对白名单内用户生效，如您想体验该功能请 [联系我们](https://cloud.tencent.com/act/event/Online_service)。
	// 2. 该参数仅在功能增强（如搜索）开关开启（EnableEnhancement=true）并且极速版搜索开关关闭（EnableSpeedSearch=false）时生效。
	// 3. hunyuan-lite 无多媒体能力，该参数对 hunyuan-lite 版本不生效。
	// 4. 未传值时默认关闭。
	// 5. 开启并搜索到对应的多媒体信息时，会输出对应的多媒体地址，可以定制个性化的图文消息。
	EnableMultimedia *bool `json:"EnableMultimedia,omitnil,omitempty" name:"EnableMultimedia"`

	// 是否开启深度研究该问题，默认是false，在值为true且命中深度研究该问题时，会返回深度研究该问题信息。
	EnableDeepSearch *bool `json:"EnableDeepSearch,omitnil,omitempty" name:"EnableDeepSearch"`

	// 说明： 1. 确保模型的输出是可复现的。 2. 取值区间为非0正整数，最大值10000。 3. 非必要不建议使用，不合理的取值会影响效果。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`

	// 强制搜索增强开关。
	// 说明：
	// 1. 未传值时默认关闭。
	// 2. 开启后，将强制走AI搜索，当AI搜索结果为空时，由大模型回复兜底话术。
	ForceSearchEnhancement *bool `json:"ForceSearchEnhancement,omitnil,omitempty" name:"ForceSearchEnhancement"`

	// 自定义结束生成字符串
	// 
	// 调用 OpenAI 的接口时，如果您指定了 `stop` 参数, 模型会停止在匹配到 `stop` 的内容之前。
	// 在调用混元接口时，会停止在匹配到 `stop` 的内容之后。
	// 
	// **说明：**
	// 未来我们可能会修改此行为以便和 OpenAI 保持一致。
	// 但是目前有使用该参数的情况下，开发者需要注意该参数是否会对应用造成影响，以及未来该行为调整时带来的影响。
	Stop []*string `json:"Stop,omitnil,omitempty" name:"Stop"`

	// 推荐问答开关。
	// 说明：
	// 1. 未传值时默认关闭。
	// 2. 开启后，在返回值的最后一个包中会增加 RecommendedQuestions 字段表示推荐问答， 最多返回3条。
	EnableRecommendedQuestions *bool `json:"EnableRecommendedQuestions,omitnil,omitempty" name:"EnableRecommendedQuestions"`

	// 是否开启深度阅读，默认是false，在值为true时，会返回深度阅读的结果信息。说明:1.深度阅读需要开启插件增强,即设置EnableEnhancement为true,当设置EnableDeepRead为true时EnableEnhancement默认为true；2.目前暂时只支持单文档单轮的深度阅读；3.深度阅读功能的文件上传可以使用FilesUploads接口，具体参数详见FilesUploads接口文档
	EnableDeepRead *bool `json:"EnableDeepRead,omitnil,omitempty" name:"EnableDeepRead"`

	// 知识注入相关的参数信息
	WebSearchOptions *WebSearchOptions `json:"WebSearchOptions,omitnil,omitempty" name:"WebSearchOptions"`

	// 用户传入Topic
	TopicChoice *string `json:"TopicChoice,omitnil,omitempty" name:"TopicChoice"`

	// 模型思维链开关 说明： 1. 未传值时默认开启，打开模型思维链推理能力。 2. 关闭后，关闭模型思维链推理能力。  开关当前仅对hunyuan-a13b模型生效 示例值：ture
	EnableThinking *bool `json:"EnableThinking,omitnil,omitempty" name:"EnableThinking"`
}

func (r *ChatCompletionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Messages")
	delete(f, "Stream")
	delete(f, "StreamModeration")
	delete(f, "TopP")
	delete(f, "Temperature")
	delete(f, "EnableEnhancement")
	delete(f, "Tools")
	delete(f, "ToolChoice")
	delete(f, "CustomTool")
	delete(f, "SearchInfo")
	delete(f, "Citation")
	delete(f, "EnableSpeedSearch")
	delete(f, "EnableMultimedia")
	delete(f, "EnableDeepSearch")
	delete(f, "Seed")
	delete(f, "ForceSearchEnhancement")
	delete(f, "Stop")
	delete(f, "EnableRecommendedQuestions")
	delete(f, "EnableDeepRead")
	delete(f, "WebSearchOptions")
	delete(f, "TopicChoice")
	delete(f, "EnableThinking")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatCompletionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatCompletionsResponseParams struct {
	// Unix 时间戳，单位为秒。
	Created *int64 `json:"Created,omitnil,omitempty" name:"Created"`

	// Token 统计信息。
	// 按照总 Token 数量计费。
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 免责声明。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 本次请求的 RequestId。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 回复内容。
	Choices []*Choice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// 错误信息。
	// 如果流式返回中服务处理异常，返回该错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *ErrorMsg `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 多轮会话风险审核，值为1时，表明存在信息安全风险，建议终止客户多轮会话。
	//
	// Deprecated: ModerationLevel is deprecated.
	ModerationLevel *string `json:"ModerationLevel,omitnil,omitempty" name:"ModerationLevel"`

	// 搜索结果信息
	SearchInfo *SearchInfo `json:"SearchInfo,omitnil,omitempty" name:"SearchInfo"`

	// 多媒体信息。
	// 说明：
	// 1. 可以用多媒体信息替换回复内容里的占位符，得到完整的消息。
	// 2. 可能会出现回复内容里存在占位符，但是因为审核等原因没有返回多媒体信息。
	Replaces []*Replace `json:"Replaces,omitnil,omitempty" name:"Replaces"`

	// 推荐问答。
	RecommendedQuestions []*string `json:"RecommendedQuestions,omitnil,omitempty" name:"RecommendedQuestions"`

	// AI搜索返回状态
	Processes *Processes `json:"Processes,omitnil,omitempty" name:"Processes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChatCompletionsResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ChatCompletionsResponseParams `json:"Response"`
}

func (r *ChatCompletionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatTranslationsRequestParams struct {
	// 模型名称，可选值包括 hunyuan-translation、hunyuan-translation-lite。
	// 各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。
	// 
	// 注意：
	// 不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为非流式调用（false）。
	// 2. 流式调用时以 SSE 协议增量返回结果（返回值取 Choices[n].Delta 中的值，需要拼接增量数据才能获得完整结果）。
	// 3. 非流式调用时：
	// 调用方式与普通 HTTP 请求无异。
	// 接口响应耗时较长，**如需更低时延建议设置为 true**。
	// 只返回一次最终结果（返回值取 Choices[n].Message 中的值）。
	// 
	// 注意：
	// 通过 SDK 调用时，流式和非流式调用需用**不同的方式**获取返回值，具体参考 SDK 中的注释或示例（在各语言 SDK 代码仓库的 examples/hunyuan/v20230901/ 目录中）。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 待翻译的文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 源语言。
	// 支持语言列表: 1. 简体中文：zh，2. 粤语：yue，3. 英语：en，4. 法语：fr，5. 葡萄牙语：pt，6. 西班牙语：es，7. 日语：ja，8. 土耳其语：tr，9. 俄语：ru，10. 阿拉伯语：ar，11. 韩语：ko，12. 泰语：th，13. 意大利语：it，14. 德语：de，15. 越南语：vi，16. 马来语：ms，17. 印尼语：id
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言。
	// 支持语言列表: 1. 简体中文：zh，2. 粤语：yue，3. 英语：en，4. 法语：fr，5. 葡萄牙语：pt，6. 西班牙语：es，7. 日语：ja，8. 土耳其语：tr，9. 俄语：ru，10. 阿拉伯语：ar，11. 韩语：ko，12. 泰语：th，13. 意大利语：it，14. 德语：de，15. 越南语：vi，16. 马来语：ms，17. 印尼语：id
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 待翻译文本所属领域，例如游戏剧情等
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 参考示例，最多10个
	References []*Reference `json:"References,omitnil,omitempty" name:"References"`
}

type ChatTranslationsRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称，可选值包括 hunyuan-translation、hunyuan-translation-lite。
	// 各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。
	// 
	// 注意：
	// 不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为非流式调用（false）。
	// 2. 流式调用时以 SSE 协议增量返回结果（返回值取 Choices[n].Delta 中的值，需要拼接增量数据才能获得完整结果）。
	// 3. 非流式调用时：
	// 调用方式与普通 HTTP 请求无异。
	// 接口响应耗时较长，**如需更低时延建议设置为 true**。
	// 只返回一次最终结果（返回值取 Choices[n].Message 中的值）。
	// 
	// 注意：
	// 通过 SDK 调用时，流式和非流式调用需用**不同的方式**获取返回值，具体参考 SDK 中的注释或示例（在各语言 SDK 代码仓库的 examples/hunyuan/v20230901/ 目录中）。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 待翻译的文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 源语言。
	// 支持语言列表: 1. 简体中文：zh，2. 粤语：yue，3. 英语：en，4. 法语：fr，5. 葡萄牙语：pt，6. 西班牙语：es，7. 日语：ja，8. 土耳其语：tr，9. 俄语：ru，10. 阿拉伯语：ar，11. 韩语：ko，12. 泰语：th，13. 意大利语：it，14. 德语：de，15. 越南语：vi，16. 马来语：ms，17. 印尼语：id
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言。
	// 支持语言列表: 1. 简体中文：zh，2. 粤语：yue，3. 英语：en，4. 法语：fr，5. 葡萄牙语：pt，6. 西班牙语：es，7. 日语：ja，8. 土耳其语：tr，9. 俄语：ru，10. 阿拉伯语：ar，11. 韩语：ko，12. 泰语：th，13. 意大利语：it，14. 德语：de，15. 越南语：vi，16. 马来语：ms，17. 印尼语：id
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 待翻译文本所属领域，例如游戏剧情等
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 参考示例，最多10个
	References []*Reference `json:"References,omitnil,omitempty" name:"References"`
}

func (r *ChatTranslationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatTranslationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Stream")
	delete(f, "Text")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "Field")
	delete(f, "References")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatTranslationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatTranslationsResponseParams struct {
	// 本次请求的 RequestId。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 免责声明。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Unix 时间戳，单位为秒。
	Created *int64 `json:"Created,omitnil,omitempty" name:"Created"`

	// Token 统计信息。
	// 按照总 Token 数量计费。
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 回复内容。
	Choices []*TranslationChoice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// 错误信息。
	// 如果流式返回中服务处理异常，返回该错误信息。
	ErrorMsg *ErrorMsg `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChatTranslationsResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ChatTranslationsResponseParams `json:"Response"`
}

func (r *ChatTranslationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatTranslationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Choice struct {
	// 结束标志位，可能为 stop、 sensitive或者tool_calls。
	// stop 表示输出正常结束。
	// sensitive 只在开启流式输出审核时会出现，表示安全审核未通过。
	// tool_calls 标识函数调用。
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// 增量返回值，流式调用时使用该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Delta *Delta `json:"Delta,omitnil,omitempty" name:"Delta"`

	// 返回值，非流式调用时使用该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *Message `json:"Message,omitnil,omitempty" name:"Message"`

	// 索引值，流式调用时使用该字段。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 多轮会话风险审核，值为1时，表明存在信息安全风险，建议终止客户多轮会话。
	ModerationLevel *string `json:"ModerationLevel,omitnil,omitempty" name:"ModerationLevel"`
}

type Content struct {
	// 内容类型
	// 注意：
	// 需包含至少一个 Type 为"text"的参数。
	// 参数值可选范围：[text", "image_url"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当 Type 为 text 时使用，表示具体的文本内容。当 Type 为 image_url 时，当前字段内容需保持为空，传递内容不生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 图片的url，当 Type 为 image_url 时使用，表示具体的图片内容
	// 如"https://example.com/1.png" 或 图片的base64（注意 "data:image/jpeg;base64," 为必要部分）："data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAA......"。当 Type 为 text 时，当前字段内容需保持为空，传递内容不生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *ImageUrl `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

// Predefined struct for user
type CreateThreadRequestParams struct {

}

type CreateThreadRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateThreadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateThreadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateThreadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateThreadResponseParams struct {
	// 会话 ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 对象类型
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 创建时间，Unix 时间戳，单位为秒。
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 提供给工具的资源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolResources *ThreadToolResources `json:"ToolResources,omitnil,omitempty" name:"ToolResources"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateThreadResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *CreateThreadResponseParams `json:"Response"`
}

func (r *CreateThreadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateThreadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Delta struct {
	// 角色名称。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容详情。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 模型生成的工具调用，仅 hunyuan-functioncall 模型支持
	// 说明：
	// 对于每一次的输出值应该以Id为标识对Type、Name、Arguments字段进行合并。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolCalls []*ToolCall `json:"ToolCalls,omitnil,omitempty" name:"ToolCalls"`

	// 思维链内容。用于展示模型思考过程，仅 Hunyuan-T1 系列模型可用。注意：在进行多轮对话时，请不要将此字段拼接到 messages 中。请求 messages 的请求参数中包含 reasoning_content，接口将报错。
	ReasoningContent *string `json:"ReasoningContent,omitnil,omitempty" name:"ReasoningContent"`
}

type EmbeddingData struct {
	// Embedding 信息，目前为 1024 维浮点数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Embedding []*float64 `json:"Embedding,omitnil,omitempty" name:"Embedding"`

	// 下标，目前不支持批量，因此固定为 0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 目前固定为 "embedding"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`
}

type EmbeddingUsage struct {
	// 输入 Token 数。
	PromptTokens *int64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// 总 Token 数。
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type ErrorMsg struct {
	// 错误提示信息。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 错误码。
	// 4000 服务内部异常。
	// 4001 请求模型超时。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`
}

type FileObject struct {
	// 文件标识符，可在各个API中引用。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 对象类型，始终为 file。
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 文件大小，单位为字节。
	Bytes *uint64 `json:"Bytes,omitnil,omitempty" name:"Bytes"`

	// 文件创建时的 Unix 时间戳（秒）。
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 文件名。
	Filename *string `json:"Filename,omitnil,omitempty" name:"Filename"`

	// 上传文件的用途。
	Purpose *string `json:"Purpose,omitnil,omitempty" name:"Purpose"`
}

// Predefined struct for user
type FilesDeletionsRequestParams struct {
	// 文件标识符。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type FilesDeletionsRequest struct {
	*tchttp.BaseRequest
	
	// 文件标识符。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *FilesDeletionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FilesDeletionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FilesDeletionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FilesDeletionsResponseParams struct {
	// 文件标识符。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 对象类型，始终为 file。
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 是否删除成功。
	Deleted *bool `json:"Deleted,omitnil,omitempty" name:"Deleted"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FilesDeletionsResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *FilesDeletionsResponseParams `json:"Response"`
}

func (r *FilesDeletionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FilesDeletionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FilesListRequestParams struct {
	// 分页偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，最大 100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type FilesListRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，最大 100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *FilesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FilesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FilesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FilesListResponseParams struct {
	// 文件数量。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 对象类型，始终为 list。
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// FileObject 列表。
	Data []*FileObject `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FilesListResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *FilesListResponseParams `json:"Response"`
}

func (r *FilesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FilesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FilesUploadsRequestParams struct {
	// 文件名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件链接。目前支持 csv, doc, docx, pdf, ppt, pptx, txt, xls, xlsx 格式，单文件大小限制为100M。
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`
}

type FilesUploadsRequest struct {
	*tchttp.BaseRequest
	
	// 文件名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件链接。目前支持 csv, doc, docx, pdf, ppt, pptx, txt, xls, xlsx 格式，单文件大小限制为100M。
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`
}

func (r *FilesUploadsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FilesUploadsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "URL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FilesUploadsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FilesUploadsResponseParams struct {
	// 文件标识符，可在各个API中引用。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 对象类型，始终为 file。
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 文件大小，单位为字节。
	Bytes *int64 `json:"Bytes,omitnil,omitempty" name:"Bytes"`

	// 文件创建时的 Unix 时间戳（秒）。
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 文件名。
	Filename *string `json:"Filename,omitnil,omitempty" name:"Filename"`

	// 上传文件的用途。
	Purpose *string `json:"Purpose,omitnil,omitempty" name:"Purpose"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FilesUploadsResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *FilesUploadsResponseParams `json:"Response"`
}

func (r *FilesUploadsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FilesUploadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmbeddingRequestParams struct {
	// 输入文本。总长度不超过 1024 个 Token，超过则会截断最后面的内容。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 输入文本数组。输入数组总长度不超过 50 。
	InputList []*string `json:"InputList,omitnil,omitempty" name:"InputList"`
}

type GetEmbeddingRequest struct {
	*tchttp.BaseRequest
	
	// 输入文本。总长度不超过 1024 个 Token，超过则会截断最后面的内容。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 输入文本数组。输入数组总长度不超过 50 。
	InputList []*string `json:"InputList,omitnil,omitempty" name:"InputList"`
}

func (r *GetEmbeddingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmbeddingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Input")
	delete(f, "InputList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEmbeddingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmbeddingResponseParams struct {
	// 返回的 Embedding 信息。
	Data []*EmbeddingData `json:"Data,omitnil,omitempty" name:"Data"`

	// Token 使用计数，按照总 Token 数量收费。
	Usage *EmbeddingUsage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetEmbeddingResponse struct {
	*tchttp.BaseResponse
	Response *GetEmbeddingResponseParams `json:"Response"`
}

func (r *GetEmbeddingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmbeddingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetThreadMessageListRequestParams struct {
	// 会话 ID
	ThreadID *string `json:"ThreadID,omitnil,omitempty" name:"ThreadID"`

	// 返回的消息条数，1 - 100 条
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式，按创建时间升序（asc）或降序（desc），默认为 desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type GetThreadMessageListRequest struct {
	*tchttp.BaseRequest
	
	// 会话 ID
	ThreadID *string `json:"ThreadID,omitnil,omitempty" name:"ThreadID"`

	// 返回的消息条数，1 - 100 条
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式，按创建时间升序（asc）或降序（desc），默认为 desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *GetThreadMessageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetThreadMessageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ThreadID")
	delete(f, "Limit")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetThreadMessageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetThreadMessageListResponseParams struct {
	// 消息列表
	Data []*ThreadMessage `json:"Data,omitnil,omitempty" name:"Data"`

	// 第一条消息 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: FirstID is deprecated.
	FirstID *string `json:"FirstID,omitnil,omitempty" name:"FirstID"`

	// 已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: LastID is deprecated.
	LastID *int64 `json:"LastID,omitnil,omitempty" name:"LastID"`

	// 是否还有更多消息
	HasMore *bool `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// 对象类型
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 第一条消息 ID
	FirstMsgID *string `json:"FirstMsgID,omitnil,omitempty" name:"FirstMsgID"`

	// 最后一条消息 ID
	LastMsgID *string `json:"LastMsgID,omitnil,omitempty" name:"LastMsgID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetThreadMessageListResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *GetThreadMessageListResponseParams `json:"Response"`
}

func (r *GetThreadMessageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetThreadMessageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetThreadMessageRequestParams struct {
	// 会话 ID
	ThreadID *string `json:"ThreadID,omitnil,omitempty" name:"ThreadID"`

	// 消息 ID
	MessageID *string `json:"MessageID,omitnil,omitempty" name:"MessageID"`
}

type GetThreadMessageRequest struct {
	*tchttp.BaseRequest
	
	// 会话 ID
	ThreadID *string `json:"ThreadID,omitnil,omitempty" name:"ThreadID"`

	// 消息 ID
	MessageID *string `json:"MessageID,omitnil,omitempty" name:"MessageID"`
}

func (r *GetThreadMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetThreadMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ThreadID")
	delete(f, "MessageID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetThreadMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetThreadMessageResponseParams struct {
	// 消息 ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 对象类型
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 创建时间
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 会话 ID
	ThreadID *string `json:"ThreadID,omitnil,omitempty" name:"ThreadID"`

	// 状态，处理中 in_progress，已完成 completed，未完成 incomplete。 
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 未完成原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCompleteDetails *ThreadMessageInCompleteDetailsObject `json:"InCompleteDetails,omitnil,omitempty" name:"InCompleteDetails"`

	// 完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompletedAt *int64 `json:"CompletedAt,omitnil,omitempty" name:"CompletedAt"`

	// 未完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCompleteAt *int64 `json:"InCompleteAt,omitnil,omitempty" name:"InCompleteAt"`

	// 角色
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 助手 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssistantID *string `json:"AssistantID,omitnil,omitempty" name:"AssistantID"`

	// 运行 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunID *string `json:"RunID,omitnil,omitempty" name:"RunID"`

	// 附件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attachments []*ThreadMessageAttachmentObject `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetThreadMessageResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *GetThreadMessageResponseParams `json:"Response"`
}

func (r *GetThreadMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetThreadMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetThreadRequestParams struct {
	// 会话 ID
	ThreadID *string `json:"ThreadID,omitnil,omitempty" name:"ThreadID"`
}

type GetThreadRequest struct {
	*tchttp.BaseRequest
	
	// 会话 ID
	ThreadID *string `json:"ThreadID,omitnil,omitempty" name:"ThreadID"`
}

func (r *GetThreadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetThreadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ThreadID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetThreadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetThreadResponseParams struct {
	// 会话 ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 对象类型
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 创建时间，Unix 时间戳，单位为秒。
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 提供给工具的资源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolResources *ThreadToolResources `json:"ToolResources,omitnil,omitempty" name:"ToolResources"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetThreadResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *GetThreadResponseParams `json:"Response"`
}

func (r *GetThreadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetThreadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTokenCountRequestParams struct {
	// 输入文本
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`
}

type GetTokenCountRequest struct {
	*tchttp.BaseRequest
	
	// 输入文本
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`
}

func (r *GetTokenCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTokenCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTokenCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTokenCountResponseParams struct {
	// token计数
	TokenCount *int64 `json:"TokenCount,omitnil,omitempty" name:"TokenCount"`

	// 字符计数
	CharacterCount *int64 `json:"CharacterCount,omitnil,omitempty" name:"CharacterCount"`

	// 切分后的列表
	Tokens []*string `json:"Tokens,omitnil,omitempty" name:"Tokens"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTokenCountResponse struct {
	*tchttp.BaseResponse
	Response *GetTokenCountResponseParams `json:"Response"`
}

func (r *GetTokenCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTokenCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupChatCompletionsRequestParams struct {
	// 模型名称，可选值包括 hunyuan-large-role-group。各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。注意：不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。
	Messages []*GroupMessage `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为非流式调用（false）。
	// 2. 流式调用时以 SSE 协议增量返回结果（返回值取 Choices[n].Delta 中的值，需要拼接增量数据才能获得完整结果）。
	// 3. 非流式调用时：
	// 调用方式与普通 HTTP 请求无异。
	// 接口响应耗时较长，**如需更低时延建议设置为 true**。
	// 只返回一次最终结果（返回值取 Choices[n].Message 中的值）。
	// 
	// 注意：
	// 通过 SDK 调用时，流式和非流式调用需用**不同的方式**获取返回值，具体参考 SDK 中的注释或示例（在各语言 SDK 代码仓库的 examples/hunyuan/v20230901/ 目录中）。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 目标人物名称
	TargetCharacterName *string `json:"TargetCharacterName,omitnil,omitempty" name:"TargetCharacterName"`

	// 角色描述
	GroupChatConfig *GroupChatConfig `json:"GroupChatConfig,omitnil,omitempty" name:"GroupChatConfig"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 对话接口
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type GroupChatCompletionsRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称，可选值包括 hunyuan-large-role-group。各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。注意：不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。
	Messages []*GroupMessage `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为非流式调用（false）。
	// 2. 流式调用时以 SSE 协议增量返回结果（返回值取 Choices[n].Delta 中的值，需要拼接增量数据才能获得完整结果）。
	// 3. 非流式调用时：
	// 调用方式与普通 HTTP 请求无异。
	// 接口响应耗时较长，**如需更低时延建议设置为 true**。
	// 只返回一次最终结果（返回值取 Choices[n].Message 中的值）。
	// 
	// 注意：
	// 通过 SDK 调用时，流式和非流式调用需用**不同的方式**获取返回值，具体参考 SDK 中的注释或示例（在各语言 SDK 代码仓库的 examples/hunyuan/v20230901/ 目录中）。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 目标人物名称
	TargetCharacterName *string `json:"TargetCharacterName,omitnil,omitempty" name:"TargetCharacterName"`

	// 角色描述
	GroupChatConfig *GroupChatConfig `json:"GroupChatConfig,omitnil,omitempty" name:"GroupChatConfig"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 对话接口
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *GroupChatCompletionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupChatCompletionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Messages")
	delete(f, "Stream")
	delete(f, "TargetCharacterName")
	delete(f, "GroupChatConfig")
	delete(f, "UserId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GroupChatCompletionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupChatCompletionsResponseParams struct {
	// Unix 时间戳，单位为秒。
	Created *int64 `json:"Created,omitnil,omitempty" name:"Created"`

	// Token 统计信息。
	// 按照总 Token 数量计费。
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 免责声明。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 本次请求的 RequestId。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 回复内容。
	Choices []*Choice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// 错误信息。
	// 如果流式返回中服务处理异常，返回该错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *ErrorMsg `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 搜索结果信息
	SearchInfo *SearchInfo `json:"SearchInfo,omitnil,omitempty" name:"SearchInfo"`

	// 多媒体信息。
	// 说明：
	// 1. 可以用多媒体信息替换回复内容里的占位符，得到完整的消息。
	// 2. 可能会出现回复内容里存在占位符，但是因为审核等原因没有返回多媒体信息。
	Replaces []*Replace `json:"Replaces,omitnil,omitempty" name:"Replaces"`

	// 推荐问答。
	RecommendedQuestions []*string `json:"RecommendedQuestions,omitnil,omitempty" name:"RecommendedQuestions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GroupChatCompletionsResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *GroupChatCompletionsResponseParams `json:"Response"`
}

func (r *GroupChatCompletionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupChatCompletionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupChatConfig struct {
	// 人物名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// ### 主题：\n武道修炼与科技创新的碰撞\n\n### 地点：\n布尔玛的实验室\n\n### 故事背景：\n布尔玛正在研发一种新型的龙珠雷达，旨在更精确地定位龙珠的位置。她邀请了孙悟空、天津饭、饺子和雅木茶前来测试新设备。然而，这些武道家们对科技的理解有限，导致了一系列有趣的误解和互动。\n\n### 人物关系：\n- **布尔玛**：天才科学家，负责研发和解释新设备。\n- **孙悟空**：纯粹的战斗狂，对科技一窍不通，但对新事物充满好奇。\n- **天津饭**：严肃自律的武道家，对科技持谨慎态度，但愿意尝试。\n- **饺子**：内向单纯，依赖天津饭，对科技感到困惑和害怕。\n- **雅木茶**：幽默自嘲的前沙漠强盗，对科技有一定了解，但更倾向于轻松调侃。\n\n### 人物目标：\n- **布尔玛**：希望新龙珠雷达能够得到有效测试，并得到反馈以改进。\n- **孙悟空**：希望通过新设备找到更强的对手进行战斗。\n- **天津饭**：希望通过测试新设备提升自己的武道修炼。\n- **饺子**：希望在不引起麻烦的情况下完成任务，并得到天津饭的认可。\n- **雅木茶**：希望通过参与测试证明自己的价值，同时享受与朋友们的互动。\n\n### 场景描述：\n布尔玛在实验室中展示她的新龙珠雷达，解释其工作原理和优势。孙悟空对设备的功能感到兴奋，但完全无法理解技术细节，不断提出天真的问题。天津饭则严肃地询问设备的安全性和可靠性，表现出对科技的谨慎态度。饺子对复杂的设备感到害怕，不断向天津饭寻求确认和安慰。雅木茶则在一旁调侃大家的反应，试图缓解紧张气氛。布尔玛在解释过程中不断被孙悟空的问题打断，感到无奈，但也被他的热情所感染。最终，大家决定一起外出测试新设备，展开一场充满欢笑和冒险的旅程。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 角色描述
	Characters []*Character `json:"Characters,omitnil,omitempty" name:"Characters"`
}

type GroupMessage struct {
	// 角色，可选值包括 system、user、assistant、 tool。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 文本内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 角色名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type History struct {
	// 对话的 ID，用于唯一标识一轮对话
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 原始输入的 Prompt 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 扩写后的 Prompt 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	RevisedPrompt *string `json:"RevisedPrompt,omitnil,omitempty" name:"RevisedPrompt"`

	// 生成图的随机种子
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`
}

type Image struct {
	// 图片Url。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 图片Base64。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`
}

type ImageMessage struct {
	// 角色，可选值包括 system、user、assistant。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 文本内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 多种类型内容（目前支持图片和文本），仅 hunyuan-vision 和 hunyuan-turbo-vision 模型支持
	Contents []*Content `json:"Contents,omitnil,omitempty" name:"Contents"`
}

// Predefined struct for user
type ImageQuestionRequestParams struct {
	// 模型名称，可选值包括 hunyuan-vision-image-question。各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。注意：不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。说明：1. 长度最多为 40，按对话时间从旧到新在数组中排列。2. Message.Role 可选值：system、user、assistant。其中，system 角色可选，如存在则必须位于列表的最开始。user 和 assistant 需交替出现（一问一答），以 user 提问开始，user提问结束，且 Content 不能为空。Role 的顺序示例：[system（可选） user assistant user assistant user ...]。3. Messages 中 Content 总长度不能超过模型输入长度上限（可参考 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 文档），超过则会截断最前面的内容，只保留尾部内容。
	Messages []*ImageMessage `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为非流式调用（false）。
	// 2. 流式调用时以 SSE 协议增量返回结果（返回值取 Choices[n].Delta 中的值，需要拼接增量数据才能获得完整结果）。
	// 3. 非流式调用时：
	// 调用方式与普通 HTTP 请求无异。
	// 接口响应耗时较长，**如需更低时延建议设置为 true**。
	// 只返回一次最终结果（返回值取 Choices[n].Message 中的值）。
	// 
	// 注意：
	// 通过 SDK 调用时，流式和非流式调用需用**不同的方式**获取返回值，具体参考 SDK 中的注释或示例（在各语言 SDK 代码仓库的 examples/hunyuan/v20230901/ 目录中）。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`
}

type ImageQuestionRequest struct {
	*tchttp.BaseRequest
	
	// 模型名称，可选值包括 hunyuan-vision-image-question。各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。注意：不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 聊天上下文信息。说明：1. 长度最多为 40，按对话时间从旧到新在数组中排列。2. Message.Role 可选值：system、user、assistant。其中，system 角色可选，如存在则必须位于列表的最开始。user 和 assistant 需交替出现（一问一答），以 user 提问开始，user提问结束，且 Content 不能为空。Role 的顺序示例：[system（可选） user assistant user assistant user ...]。3. Messages 中 Content 总长度不能超过模型输入长度上限（可参考 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 文档），超过则会截断最前面的内容，只保留尾部内容。
	Messages []*ImageMessage `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 流式调用开关。
	// 说明：
	// 1. 未传值时默认为非流式调用（false）。
	// 2. 流式调用时以 SSE 协议增量返回结果（返回值取 Choices[n].Delta 中的值，需要拼接增量数据才能获得完整结果）。
	// 3. 非流式调用时：
	// 调用方式与普通 HTTP 请求无异。
	// 接口响应耗时较长，**如需更低时延建议设置为 true**。
	// 只返回一次最终结果（返回值取 Choices[n].Message 中的值）。
	// 
	// 注意：
	// 通过 SDK 调用时，流式和非流式调用需用**不同的方式**获取返回值，具体参考 SDK 中的注释或示例（在各语言 SDK 代码仓库的 examples/hunyuan/v20230901/ 目录中）。
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`
}

func (r *ImageQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Messages")
	delete(f, "Stream")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageQuestionResponseParams struct {
	// Unix 时间戳，单位为秒。
	Created *int64 `json:"Created,omitnil,omitempty" name:"Created"`

	// Token 统计信息。
	// 按照总 Token 数量计费。
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 免责声明。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 本次请求的 RequestId。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 回复内容。
	Choices []*Choice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// 错误信息。
	// 如果流式返回中服务处理异常，返回该错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *ErrorMsg `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageQuestionResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ImageQuestionResponseParams `json:"Response"`
}

func (r *ImageQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageUrl struct {
	// 图片的 Url（以 http:// 或 https:// 开头）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type Knowledge struct {
	// 表示具体的知识信息文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type LogoParam struct {
	// 水印url
	LogoUrl *string `json:"LogoUrl,omitnil,omitempty" name:"LogoUrl"`

	// 水印base64，url和base64二选一传入
	LogoImage *string `json:"LogoImage,omitnil,omitempty" name:"LogoImage"`

	// 水印图片位于融合结果图中的坐标，将按照坐标对标识图片进行位置和大小的拉伸匹配
	LogoRect *LogoRect `json:"LogoRect,omitnil,omitempty" name:"LogoRect"`
}

type LogoRect struct {
	// 左上角X坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 左上角Y坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 方框宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 方框高度
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type Message struct {
	// 角色，可选值包括 system、user、assistant、 tool。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 文本内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 多种类型内容（目前支持图片和文本），仅 hunyuan-vision 和 hunyuan-turbo-vision 模型支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contents []*Content `json:"Contents,omitnil,omitempty" name:"Contents"`

	// 当role为tool时传入，标识具体的函数调用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolCallId *string `json:"ToolCallId,omitnil,omitempty" name:"ToolCallId"`

	// 模型生成的工具调用，仅 hunyuan-pro 或者 hunyuan-functioncall 模型支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolCalls []*ToolCall `json:"ToolCalls,omitnil,omitempty" name:"ToolCalls"`

	// 文件标识符。单次最大 50 个文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileIDs []*string `json:"FileIDs,omitnil,omitempty" name:"FileIDs"`

	// 思维链内容。用于展示模型思考过程，仅 Hunyuan-T1 系列模型可用。注意：在进行多轮对话时，请**不要**将此字段拼接到 messages 中。请求 messages 的请求参数中包含 reasoning_content，接口将报错。
	ReasoningContent *string `json:"ReasoningContent,omitnil,omitempty" name:"ReasoningContent"`
}

type Mindmap struct {
	// 脑图缩略图链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThumbUrl *string `json:"ThumbUrl,omitnil,omitempty" name:"ThumbUrl"`

	// 脑图图片链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type Multimedia struct {
	// 多媒体类型，可选值包括 image、music、album、playlist。
	// 说明：
	// 1. image：图片；music：单曲，类型为单曲时，会返回详细歌手和歌曲信息；album：专辑；playlist：歌单。
	// 2. 当 type 为 music、album、playlist 时，需要配合 [QQ音乐SDK](https://developer.y.qq.com/) 使用。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 多媒体地址。
	// 说明：
	// 1. type 为 image 时，地址为图片的预览地址；其他类型时，地址为封面图地址。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 如果Url为图片地址，标识图片宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 如果Url为图片地址，标识图片高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 多媒体详情地址。
	// 说明：
	// 1. 仅 type 为 image 时，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 缩略图地址。
	ThumbURL *string `json:"ThumbURL,omitnil,omitempty" name:"ThumbURL"`

	// 缩略图宽度
	ThumbWidth *int64 `json:"ThumbWidth,omitnil,omitempty" name:"ThumbWidth"`

	// 缩略图高度
	ThumbHeight *int64 `json:"ThumbHeight,omitnil,omitempty" name:"ThumbHeight"`

	// 名称。
	// 说明：
	// 1. type 为 image 时，该字段为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 歌手名称。
	// 说明：
	// 1. 仅 type 为 music 时，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Singer *string `json:"Singer,omitnil,omitempty" name:"Singer"`

	// 歌曲详情。
	// 说明：
	// 1. 仅 type 为 music 时，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ext *SongExt `json:"Ext,omitnil,omitempty" name:"Ext"`

	// 发布时间。
	PublishTime *string `json:"PublishTime,omitnil,omitempty" name:"PublishTime"`

	// 站点名称
	SiteName *string `json:"SiteName,omitnil,omitempty" name:"SiteName"`

	// 站点图标
	SiteIcon *string `json:"SiteIcon,omitnil,omitempty" name:"SiteIcon"`
}

type Processes struct {
	// 输出信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// plan:开始获取资料…
	// recall:找到 n 篇相关资料
	// quote:引用 n 篇资料作为参考
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 当状态是recall和quote，会给出来相关数量
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`
}

// Predefined struct for user
type QueryHunyuanImageChatJobRequestParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuanImageChatJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuanImageChatJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanImageChatJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuanImageChatJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanImageChatJobResponseParams struct {
	// 当前任务状态码：
	// 1：等待中、2：运行中、4：处理失败、5：处理完成。
	JobStatusCode *string `json:"JobStatusCode,omitnil,omitempty" name:"JobStatusCode"`

	// 当前任务状态：排队中、处理中、处理失败或者处理完成。
	JobStatusMsg *string `json:"JobStatusMsg,omitnil,omitempty" name:"JobStatusMsg"`

	// 任务处理失败错误码。
	JobErrorCode *string `json:"JobErrorCode,omitnil,omitempty" name:"JobErrorCode"`

	// 任务处理失败错误信息。
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`

	// 本轮对话的 ChatId，ChatId 用于唯一标识一轮对话。
	// 一个对话组中，最多支持进行100轮对话。
	// 每轮对话数据有效期为7天，到期后 ChatId 失效，有效期内的历史对话数据可通过 History 查询，如有长期使用需求请及时保存输入输出数据。
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 生成图 URL 列表，有效期7天，请及时保存。
	ResultImage []*string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// 结果 detail 数组，Success 代表成功。
	ResultDetails []*string `json:"ResultDetails,omitnil,omitempty" name:"ResultDetails"`

	// 本轮对话前置的历史对话数据（不含生成图）。
	History []*History `json:"History,omitnil,omitempty" name:"History"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuanImageChatJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuanImageChatJobResponseParams `json:"Response"`
}

func (r *QueryHunyuanImageChatJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanImageChatJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanImageJobRequestParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuanImageJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuanImageJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanImageJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuanImageJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanImageJobResponseParams struct {
	// 当前任务状态码：
	// 1：等待中、2：运行中、4：处理失败、5：处理完成。
	JobStatusCode *string `json:"JobStatusCode,omitnil,omitempty" name:"JobStatusCode"`

	// 当前任务状态：排队中、处理中、处理失败或者处理完成。
	JobStatusMsg *string `json:"JobStatusMsg,omitnil,omitempty" name:"JobStatusMsg"`

	// 任务处理失败错误码。
	JobErrorCode *string `json:"JobErrorCode,omitnil,omitempty" name:"JobErrorCode"`

	// 任务处理失败错误信息。
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`

	// 生成图 URL 列表，有效期1小时，请及时保存。
	ResultImage []*string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// 结果 detail 数组，Success 代表成功。
	ResultDetails []*string `json:"ResultDetails,omitnil,omitempty" name:"ResultDetails"`

	// 对应 SubmitHunyuanImageJob 接口中 Revise 参数。开启扩写时，返回扩写后的 prompt 文本。 如果关闭扩写，将直接返回原始输入的 prompt。
	RevisedPrompt []*string `json:"RevisedPrompt,omitnil,omitempty" name:"RevisedPrompt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuanImageJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuanImageJobResponseParams `json:"Response"`
}

func (r *QueryHunyuanImageJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanImageJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Reference struct {
	// 翻译文本类型，枚举"sentence"表示句子, "term"表示术语
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 原文
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 译文
	Translation *string `json:"Translation,omitnil,omitempty" name:"Translation"`
}

type RelevantEntity struct {
	// 相关组织及人物名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 相关组织及人物内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 相关事件引用文章标号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reference []*int64 `json:"Reference,omitnil,omitempty" name:"Reference"`
}

type RelevantEvent struct {
	// 相关事件标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 相关事件内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 相关事件时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Datetime *string `json:"Datetime,omitnil,omitempty" name:"Datetime"`

	// 相关事件引用文章标号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reference []*int64 `json:"Reference,omitnil,omitempty" name:"Reference"`
}

type Replace struct {
	// 占位符序号
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 多媒体详情
	Multimedia []*Multimedia `json:"Multimedia,omitnil,omitempty" name:"Multimedia"`
}

// Predefined struct for user
type RunThreadRequestParams struct {
	// 会话 ID
	ThreadID *string `json:"ThreadID,omitnil,omitempty" name:"ThreadID"`

	// 助手 ID（目前未使用，留空）
	AssistantID *string `json:"AssistantID,omitnil,omitempty" name:"AssistantID"`

	// 模型名称，可选值包括 hunyuan-lite、hunyuan-standard、hunyuan-standard-256K、hunyuan-pro、 hunyuan-code、 hunyuan-role、 hunyuan-functioncall、 hunyuan-vision、 hunyuan-turbo。各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。注意：不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 附加消息
	AdditionalMessages []*ThreadAdditionalMessage `json:"AdditionalMessages,omitnil,omitempty" name:"AdditionalMessages"`

	// 说明：1. 影响模型输出多样性，模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。2. 取值区间为 [0.0, 2.0]。较高的数值会使输出更加多样化和不可预测，而较低的数值会使其更加集中和确定。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 说明：1. 影响输出文本的多样性。模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。2. 取值区间为 [0.0, 1.0]。取值越大，生成文本的多样性越强。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 是否流式输出，当前只允许 true
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 运行过程中可使用的 token 最大数量。
	MaxPromptTokens *int64 `json:"MaxPromptTokens,omitnil,omitempty" name:"MaxPromptTokens"`

	// 运行过程中可使用的完成 token 的最大数量。
	MaxCompletionTokens *int64 `json:"MaxCompletionTokens,omitnil,omitempty" name:"MaxCompletionTokens"`

	// 可调用的工具列表，仅对 hunyuan-pro、hunyuan-turbo、hunyuan-functioncall 模型生效。
	Tools []*Tool `json:"Tools,omitnil,omitempty" name:"Tools"`

	// 工具使用选项，可选值包括 none、auto、custom。说明：1. 仅对 hunyuan-pro、hunyuan-turbo、hunyuan-functioncall 模型生效。2. none：不调用工具；auto：模型自行选择生成回复或调用工具；custom：强制模型调用指定的工具。3. 未设置时，默认值为auto
	ToolChoice *string `json:"ToolChoice,omitnil,omitempty" name:"ToolChoice"`
}

type RunThreadRequest struct {
	*tchttp.BaseRequest
	
	// 会话 ID
	ThreadID *string `json:"ThreadID,omitnil,omitempty" name:"ThreadID"`

	// 助手 ID（目前未使用，留空）
	AssistantID *string `json:"AssistantID,omitnil,omitempty" name:"AssistantID"`

	// 模型名称，可选值包括 hunyuan-lite、hunyuan-standard、hunyuan-standard-256K、hunyuan-pro、 hunyuan-code、 hunyuan-role、 hunyuan-functioncall、 hunyuan-vision、 hunyuan-turbo。各模型介绍请阅读 [产品概述](https://cloud.tencent.com/document/product/1729/104753) 中的说明。注意：不同的模型计费不同，请根据 [购买指南](https://cloud.tencent.com/document/product/1729/97731) 按需调用。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 附加消息
	AdditionalMessages []*ThreadAdditionalMessage `json:"AdditionalMessages,omitnil,omitempty" name:"AdditionalMessages"`

	// 说明：1. 影响模型输出多样性，模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。2. 取值区间为 [0.0, 2.0]。较高的数值会使输出更加多样化和不可预测，而较低的数值会使其更加集中和确定。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 说明：1. 影响输出文本的多样性。模型已有默认参数，不传值时使用各模型推荐值，不推荐用户修改。2. 取值区间为 [0.0, 1.0]。取值越大，生成文本的多样性越强。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 是否流式输出，当前只允许 true
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// 运行过程中可使用的 token 最大数量。
	MaxPromptTokens *int64 `json:"MaxPromptTokens,omitnil,omitempty" name:"MaxPromptTokens"`

	// 运行过程中可使用的完成 token 的最大数量。
	MaxCompletionTokens *int64 `json:"MaxCompletionTokens,omitnil,omitempty" name:"MaxCompletionTokens"`

	// 可调用的工具列表，仅对 hunyuan-pro、hunyuan-turbo、hunyuan-functioncall 模型生效。
	Tools []*Tool `json:"Tools,omitnil,omitempty" name:"Tools"`

	// 工具使用选项，可选值包括 none、auto、custom。说明：1. 仅对 hunyuan-pro、hunyuan-turbo、hunyuan-functioncall 模型生效。2. none：不调用工具；auto：模型自行选择生成回复或调用工具；custom：强制模型调用指定的工具。3. 未设置时，默认值为auto
	ToolChoice *string `json:"ToolChoice,omitnil,omitempty" name:"ToolChoice"`
}

func (r *RunThreadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunThreadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ThreadID")
	delete(f, "AssistantID")
	delete(f, "Model")
	delete(f, "AdditionalMessages")
	delete(f, "Temperature")
	delete(f, "TopP")
	delete(f, "Stream")
	delete(f, "MaxPromptTokens")
	delete(f, "MaxCompletionTokens")
	delete(f, "Tools")
	delete(f, "ToolChoice")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunThreadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunThreadResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。本接口为流式响应接口，当请求成功时，RequestId 会被放在 HTTP 响应的 Header "X-TC-RequestId" 中。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunThreadResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *RunThreadResponseParams `json:"Response"`
}

func (r *RunThreadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunThreadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchInfo struct {
	// 搜索引文信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchResults []*SearchResult `json:"SearchResults,omitnil,omitempty" name:"SearchResults"`

	// 脑图（回复中不一定存在，流式协议中，仅在最后一条流式数据中返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mindmap *Mindmap `json:"Mindmap,omitnil,omitempty" name:"Mindmap"`

	// 相关事件（回复中不一定存在，流式协议中，仅在最后一条流式数据中返回，深度模式下返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevantEvents []*RelevantEvent `json:"RelevantEvents,omitnil,omitempty" name:"RelevantEvents"`

	// 相关组织及人物（回复中不一定存在，流式协议中，仅在最后一条流式数据中返回，深度模式下返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevantEntities []*RelevantEntity `json:"RelevantEntities,omitnil,omitempty" name:"RelevantEntities"`

	// 时间线（回复中不一定存在，流式协议中，仅在最后一条流式数据中返回，深度模式下返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeline []*Timeline `json:"Timeline,omitnil,omitempty" name:"Timeline"`

	// 是否命中搜索深度模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportDeepSearch *bool `json:"SupportDeepSearch,omitnil,omitempty" name:"SupportDeepSearch"`

	// 搜索回复大纲（深度模式下返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Outline []*string `json:"Outline,omitnil,omitempty" name:"Outline"`
}

type SearchResult struct {
	// 搜索引文序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 搜索引文标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 搜索引文链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 搜索引文站点名
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 搜索引文图标
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`
}

// Predefined struct for user
type SetPayModeRequestParams struct {
	// 设置后付费状态，0：后付费打开；1：后付费关闭
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type SetPayModeRequest struct {
	*tchttp.BaseRequest
	
	// 设置后付费状态，0：后付费打开；1：后付费关闭
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

func (r *SetPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPayModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetPayModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetPayModeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetPayModeResponse struct {
	*tchttp.BaseResponse
	Response *SetPayModeResponseParams `json:"Response"`
}

func (r *SetPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SongExt struct {
	// 歌曲id
	SongId *int64 `json:"SongId,omitnil,omitempty" name:"SongId"`

	// 歌曲mid
	SongMid *string `json:"SongMid,omitnil,omitempty" name:"SongMid"`

	// 歌曲是否为vip。1：vip歌曲； 0：普通歌曲。
	Vip *int64 `json:"Vip,omitnil,omitempty" name:"Vip"`
}

// Predefined struct for user
type SubmitHunyuanImageChatJobRequestParams struct {
	// 本轮对话的文本描述。
	// 提交一个任务请求对应发起一轮生图对话，每轮对话中可输入一条 Prompt，生成一张图像，支持通过多轮输入 Prompt 来不断调整图像内容。
	// 推荐使用中文，最多可传1024个 utf-8 字符。
	// 输入示例：
	// <li> 第一轮对话：一颗红色的苹果 </li>
	// <li> 第二轮对话：将苹果改为绿色 </li>
	// <li> 第三轮对话：苹果放在桌子上 </li>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 上传上一轮对话的 ChatId，本轮对话将在指定的上一轮对话结果基础上继续生成图像。
	// 如果不传代表新建一个对话组，重新开启一轮新的对话。
	// 一个对话组中，最多支持进行100轮对话。
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitHunyuanImageChatJobRequest struct {
	*tchttp.BaseRequest
	
	// 本轮对话的文本描述。
	// 提交一个任务请求对应发起一轮生图对话，每轮对话中可输入一条 Prompt，生成一张图像，支持通过多轮输入 Prompt 来不断调整图像内容。
	// 推荐使用中文，最多可传1024个 utf-8 字符。
	// 输入示例：
	// <li> 第一轮对话：一颗红色的苹果 </li>
	// <li> 第二轮对话：将苹果改为绿色 </li>
	// <li> 第三轮对话：苹果放在桌子上 </li>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 上传上一轮对话的 ChatId，本轮对话将在指定的上一轮对话结果基础上继续生成图像。
	// 如果不传代表新建一个对话组，重新开启一轮新的对话。
	// 一个对话组中，最多支持进行100轮对话。
	ChatId *string `json:"ChatId,omitnil,omitempty" name:"ChatId"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitHunyuanImageChatJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanImageChatJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "ChatId")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanImageChatJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanImageChatJobResponseParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanImageChatJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanImageChatJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanImageChatJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanImageChatJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanImageJobRequestParams struct {
	// 文本描述。 
	// 算法将根据输入的文本智能生成与之相关的图像。 
	// 不能为空，推荐使用中文。最多可传1024个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 反向提示词。 
	// 推荐使用中文。最多可传1024个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在 [混元生图风格列表](https://cloud.tencent.com/document/product/1729/105846) 中选择期望的风格，传入风格编号。
	// 不传默认不指定风格。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3），不传默认使用1024:1024。
	// 如果上传 ContentImage 参考图，分辨率仅支持：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1），不传将自动适配分辨率。如果参考图被用于做风格转换，将生成保持原图长宽比例且长边为1024的图片，指定的分辨率不生效。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 图片生成数量。
	// 支持1 ~ 4张，默认生成1张。
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 超分选项，默认不做超分，可选开启。
	//  x2：2倍超分
	//  x4：4倍超分
	// 在 Resolution 的基础上按比例提高分辨率，例如1024:1024开启2倍超分后将得到2048:2048。
	Clarity *string `json:"Clarity,omitnil,omitempty" name:"Clarity"`

	// 用于引导内容的参考图。
	// 图片限制：单边分辨率小于5000，转成 Base64 字符串后小于 8MB，格式支持 jpg、jpeg、png、bmp、tiff、webp。
	ContentImage *Image `json:"ContentImage,omitnil,omitempty" name:"ContentImage"`

	// prompt 扩写开关。1为开启，0为关闭，不传默认开启。
	// 开启扩写后，将自动扩写原始输入的 prompt 并使用扩写后的 prompt 生成图片，返回生成图片结果时将一并返回扩写后的 prompt 文本。
	// 如果关闭扩写，将直接使用原始输入的 prompt 生成图片。如果上传了参考图，扩写关闭不生效，将保持开启。
	// 建议开启，在多数场景下可提升生成图片效果、丰富生成图片细节。
	Revise *int64 `json:"Revise,omitnil,omitempty" name:"Revise"`

	// 随机种子，默认随机。
	// 不传：随机种子生成。
	// 正数：固定种子生成。
	// 扩写开启时固定种子不生效，将保持随机。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitHunyuanImageJobRequest struct {
	*tchttp.BaseRequest
	
	// 文本描述。 
	// 算法将根据输入的文本智能生成与之相关的图像。 
	// 不能为空，推荐使用中文。最多可传1024个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 反向提示词。 
	// 推荐使用中文。最多可传1024个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在 [混元生图风格列表](https://cloud.tencent.com/document/product/1729/105846) 中选择期望的风格，传入风格编号。
	// 不传默认不指定风格。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3），不传默认使用1024:1024。
	// 如果上传 ContentImage 参考图，分辨率仅支持：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1），不传将自动适配分辨率。如果参考图被用于做风格转换，将生成保持原图长宽比例且长边为1024的图片，指定的分辨率不生效。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 图片生成数量。
	// 支持1 ~ 4张，默认生成1张。
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 超分选项，默认不做超分，可选开启。
	//  x2：2倍超分
	//  x4：4倍超分
	// 在 Resolution 的基础上按比例提高分辨率，例如1024:1024开启2倍超分后将得到2048:2048。
	Clarity *string `json:"Clarity,omitnil,omitempty" name:"Clarity"`

	// 用于引导内容的参考图。
	// 图片限制：单边分辨率小于5000，转成 Base64 字符串后小于 8MB，格式支持 jpg、jpeg、png、bmp、tiff、webp。
	ContentImage *Image `json:"ContentImage,omitnil,omitempty" name:"ContentImage"`

	// prompt 扩写开关。1为开启，0为关闭，不传默认开启。
	// 开启扩写后，将自动扩写原始输入的 prompt 并使用扩写后的 prompt 生成图片，返回生成图片结果时将一并返回扩写后的 prompt 文本。
	// 如果关闭扩写，将直接使用原始输入的 prompt 生成图片。如果上传了参考图，扩写关闭不生效，将保持开启。
	// 建议开启，在多数场景下可提升生成图片效果、丰富生成图片细节。
	Revise *int64 `json:"Revise,omitnil,omitempty" name:"Revise"`

	// 随机种子，默认随机。
	// 不传：随机种子生成。
	// 正数：固定种子生成。
	// 扩写开启时固定种子不生效，将保持随机。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitHunyuanImageJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanImageJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "NegativePrompt")
	delete(f, "Style")
	delete(f, "Resolution")
	delete(f, "Num")
	delete(f, "Clarity")
	delete(f, "ContentImage")
	delete(f, "Revise")
	delete(f, "Seed")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanImageJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanImageJobResponseParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanImageJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanImageJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanImageJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanImageJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToImageLiteRequestParams struct {
	// 文本描述。
	// 算法将根据输入的文本智能生成与之相关的图像。建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。
	// 不能为空，推荐使用中文。最多可传256个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传256个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在 [文生图轻量版风格列表](https://cloud.tencent.com/document/product/1729/108992) 中选择期望的风格，传入风格编号。不传默认使用201（日系动漫风格）。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3）、1080:1920（9:16）、1920:1080（16:9），不传默认使用768:768。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按0处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 返回图像方式（base64 或 url) ，二选一，默认为 base64。url 有效期为1小时。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

type TextToImageLiteRequest struct {
	*tchttp.BaseRequest
	
	// 文本描述。
	// 算法将根据输入的文本智能生成与之相关的图像。建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。
	// 不能为空，推荐使用中文。最多可传256个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传256个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在 [文生图轻量版风格列表](https://cloud.tencent.com/document/product/1729/108992) 中选择期望的风格，传入风格编号。不传默认使用201（日系动漫风格）。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3）、1080:1920（9:16）、1920:1080（16:9），不传默认使用768:768。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按0处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 返回图像方式（base64 或 url) ，二选一，默认为 base64。url 有效期为1小时。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

func (r *TextToImageLiteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToImageLiteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "NegativePrompt")
	delete(f, "Style")
	delete(f, "Resolution")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToImageLiteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToImageLiteResponseParams struct {
	// 根据入参 RspImgType 填入不同，返回不同的内容。
	// 如果传入 base64 则返回生成图 Base64 编码。
	// 如果传入 url 则返回的生成图 URL , 有效期1小时，请及时保存。
	ResultImage *string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextToImageLiteResponse struct {
	*tchttp.BaseResponse
	Response *TextToImageLiteResponseParams `json:"Response"`
}

func (r *TextToImageLiteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToImageLiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ThreadAdditionalMessage struct {
	// 角色
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 附件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attachments []*ThreadMessageAttachmentObject `json:"Attachments,omitnil,omitempty" name:"Attachments"`
}

type ThreadMessage struct {
	// 消息 ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 对象类型
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 创建时间
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 会话 ID
	ThreadID *string `json:"ThreadID,omitnil,omitempty" name:"ThreadID"`

	// 状态，处理中 in_progress，已完成 completed，未完成 incomplete。 
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 未完成原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCompleteDetails *ThreadMessageInCompleteDetailsObject `json:"InCompleteDetails,omitnil,omitempty" name:"InCompleteDetails"`

	// 完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompletedAt *int64 `json:"CompletedAt,omitnil,omitempty" name:"CompletedAt"`

	// 未完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCompleteAt *int64 `json:"InCompleteAt,omitnil,omitempty" name:"InCompleteAt"`

	// 角色
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 助手 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssistantID *string `json:"AssistantID,omitnil,omitempty" name:"AssistantID"`

	// 运行 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunID *string `json:"RunID,omitnil,omitempty" name:"RunID"`

	// 附件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attachments []*ThreadMessageAttachmentObject `json:"Attachments,omitnil,omitempty" name:"Attachments"`
}

type ThreadMessageAttachmentObject struct {
	// 文件 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileID *string `json:"FileID,omitnil,omitempty" name:"FileID"`
}

type ThreadMessageInCompleteDetailsObject struct {
	// 会话消息未完成原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type ThreadToolResources struct {
	// 文件 ID 列表
	CodeInterpreter []*string `json:"CodeInterpreter,omitnil,omitempty" name:"CodeInterpreter"`

	// 向量存储 ID 列表
	VectorStoreIDs []*string `json:"VectorStoreIDs,omitnil,omitempty" name:"VectorStoreIDs"`
}

type Timeline struct {
	// 标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Datetime *string `json:"Datetime,omitnil,omitempty" name:"Datetime"`

	// 相关网页链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type Tool struct {
	// 工具类型，当前只支持function
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 具体要调用的function
	Function *ToolFunction `json:"Function,omitnil,omitempty" name:"Function"`
}

type ToolCall struct {
	// 工具调用id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 工具调用类型，当前只支持function
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 具体的function调用
	Function *ToolCallFunction `json:"Function,omitnil,omitempty" name:"Function"`

	// 索引值
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

type ToolCallFunction struct {
	// function名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// function参数，一般为json字符串
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`
}

type ToolFunction struct {
	// function名称，只能包含a-z，A-Z，0-9，\_或-
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// function参数，一般为json字符串
	Parameters *string `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// function的简单描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type TranslationChoice struct {
	// 结束标志位，可能为 stop、 sensitive。
	// stop 表示输出正常结束。
	// sensitive 只在开启流式输出审核时会出现，表示安全审核未通过。
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// 索引值，流式调用时使用该字段。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 增量返回值，流式调用时使用该字段。
	Delta *TranslationDelta `json:"Delta,omitnil,omitempty" name:"Delta"`

	// 返回值，非流式调用时使用该字段。
	Message *TranslationMessage `json:"Message,omitnil,omitempty" name:"Message"`
}

type TranslationDelta struct {
	// 角色名称。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 内容详情。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type TranslationMessage struct {
	// 角色，可选值包括 system、user、assistant、 tool。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 文本内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type Usage struct {
	// 输入 Token 数量。
	PromptTokens *int64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// 输出 Token 数量。
	CompletionTokens *int64 `json:"CompletionTokens,omitnil,omitempty" name:"CompletionTokens"`

	// 总 Token 数量。
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type UserLocation struct {
	// 表示位置类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 用户近似位置的详细信息
	Approximate *Approximate `json:"Approximate,omitnil,omitempty" name:"Approximate"`
}

type WebSearchOptions struct {
	// 表示用户注入的知识信息
	Knowledge []*Knowledge `json:"Knowledge,omitnil,omitempty" name:"Knowledge"`

	// 用户位置详细信息
	UserLocation *UserLocation `json:"UserLocation,omitnil,omitempty" name:"UserLocation"`

	// 打开开关，会返回搜索状态
	Processes *bool `json:"Processes,omitnil,omitempty" name:"Processes"`
}