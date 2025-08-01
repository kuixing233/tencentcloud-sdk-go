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

package v20200910

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AdmissionConditionBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AdmissionDiagnosisBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Advice struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type AspectRatio struct {
	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 数值
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// 关系
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type BaseInfo struct {
	// 项目名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 标准值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 原文对应坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type BaseItem struct {
	// 类型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原始文本
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化后值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 四点坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type BaseItem2 struct {
	// 类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原始文本
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化后值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 四点坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type BaseItem3 struct {
	// 类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原始文本
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化后值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 四点坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`

	// 第几次
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`
}

type BiologicalProductInfo struct {
	// 药品名称，包括通用名和商品名
	Name *BiologicalProductInfoName `json:"Name,omitnil,omitempty" name:"Name"`

	// 成份和性状
	IngredientAndAppearance *BiologicalProductInfoIngredientAndAppearance `json:"IngredientAndAppearance,omitnil,omitempty" name:"IngredientAndAppearance"`

	// 接种对象
	VaccinationTarget *BiologicalProductInfoVaccinationTarget `json:"VaccinationTarget,omitnil,omitempty" name:"VaccinationTarget"`

	// 作用与用途
	Indications *BiologicalProductInfoIndications `json:"Indications,omitnil,omitempty" name:"Indications"`

	// 规格
	Brochure *BiologicalProductInfoBrochure `json:"Brochure,omitnil,omitempty" name:"Brochure"`

	// 免疫程序和剂量
	Dosage *BiologicalProductInfoDosage `json:"Dosage,omitnil,omitempty" name:"Dosage"`

	// 不良反应
	AdverseReaction *BiologicalProductInfoAdverseReaction `json:"AdverseReaction,omitnil,omitempty" name:"AdverseReaction"`

	// 禁忌情况
	Contraindications *BiologicalProductInfoContraindications `json:"Contraindications,omitnil,omitempty" name:"Contraindications"`

	// 注意事项
	Precautions *BiologicalProductInfoPrecautions `json:"Precautions,omitnil,omitempty" name:"Precautions"`

	// 储存条件
	Storage *BiologicalProductInfoStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 包装信息
	Packaging *BiologicalProductInfoPackaging `json:"Packaging,omitnil,omitempty" name:"Packaging"`

	// 有效期
	ValidityPeriod *BiologicalProductInfoValidityPeriod `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// 执行标准
	ExecutiveStandards *BiologicalProductInfoExecutiveStandards `json:"ExecutiveStandards,omitnil,omitempty" name:"ExecutiveStandards"`

	// 批准文号
	Approval *BiologicalProductInfoApproval `json:"Approval,omitnil,omitempty" name:"Approval"`

	// 生产企业名称和地址
	Manufacturer *BiologicalProductInfoManufacturer `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`
}

type BiologicalProductInfoAdverseReaction struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoApproval struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoBrochure struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoContraindications struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoDosage struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoExecutiveStandards struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoIndications struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoIngredientAndAppearance struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoManufacturer struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoName struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 通用名
	GenericName *string `json:"GenericName,omitnil,omitempty" name:"GenericName"`

	// 品牌名
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: BarndName is deprecated.
	BarndName *string `json:"BarndName,omitnil,omitempty" name:"BarndName"`

	// 英文名
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`

	// 拼音
	Pinyin *string `json:"Pinyin,omitnil,omitempty" name:"Pinyin"`
}

type BiologicalProductInfoPackaging struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoPrecautions struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoStorage struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoVaccinationTarget struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiologicalProductInfoValidityPeriod struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type BiopsyPart struct {
	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type BirthCert struct {
	// 新生儿信息
	NeonatalInfo *NeonatalInfo `json:"NeonatalInfo,omitnil,omitempty" name:"NeonatalInfo"`

	// 母亲信息
	MotherInfo *ParentInfo `json:"MotherInfo,omitnil,omitempty" name:"MotherInfo"`

	// 父亲信息
	FatherInfo *ParentInfo `json:"FatherInfo,omitnil,omitempty" name:"FatherInfo"`

	// 签发信息
	IssueInfo *IssueInfo `json:"IssueInfo,omitnil,omitempty" name:"IssueInfo"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type BirthPlaceBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Block struct {
	// 诊断信息
	Check []*Check `json:"Check,omitnil,omitempty" name:"Check"`

	// 病理报告
	Pathology []*PathologyReport `json:"Pathology,omitnil,omitempty" name:"Pathology"`

	// 医学资料
	MedDoc []*MedDoc `json:"MedDoc,omitnil,omitempty" name:"MedDoc"`

	// 诊断证明
	DiagCert []*DiagCert `json:"DiagCert,omitnil,omitempty" name:"DiagCert"`

	// 病案首页
	FirstPage []*FirstPage `json:"FirstPage,omitnil,omitempty" name:"FirstPage"`

	// 检验报告
	Indicator []*Indicator `json:"Indicator,omitnil,omitempty" name:"Indicator"`

	// 门诊病历信息
	MedicalRecordInfo []*MedicalRecordInfo `json:"MedicalRecordInfo,omitnil,omitempty" name:"MedicalRecordInfo"`

	// 出入院信息
	Hospitalization []*Hospitalization `json:"Hospitalization,omitnil,omitempty" name:"Hospitalization"`

	// 手术记录
	Surgery []*Surgery `json:"Surgery,omitnil,omitempty" name:"Surgery"`

	// 处方单
	Prescription []*Prescription `json:"Prescription,omitnil,omitempty" name:"Prescription"`

	// 免疫接种证明
	VaccineCertificate []*VaccineCertificate `json:"VaccineCertificate,omitnil,omitempty" name:"VaccineCertificate"`

	// 心电图
	Electrocardiogram []*Electrocardiogram `json:"Electrocardiogram,omitnil,omitempty" name:"Electrocardiogram"`

	// 病理报告v2
	PathologyV2 []*PathologyV2 `json:"PathologyV2,omitnil,omitempty" name:"PathologyV2"`

	// 内窥镜报告
	Endoscopy []*Endoscopy `json:"Endoscopy,omitnil,omitempty" name:"Endoscopy"`

	// C14检验报告
	C14 []*Indicator `json:"C14,omitnil,omitempty" name:"C14"`

	// 体检结论
	Exame []*Exame `json:"Exame,omitnil,omitempty" name:"Exame"`

	// 出入院结构体
	MedDocV2 []*DischargeInfoBlock `json:"MedDocV2,omitnil,omitempty" name:"MedDocV2"`

	// 检验报告v3
	IndicatorV3 []*IndicatorV3 `json:"IndicatorV3,omitnil,omitempty" name:"IndicatorV3"`

	// 孕产报告
	Maternity []*Maternity `json:"Maternity,omitnil,omitempty" name:"Maternity"`

	// 时间轴
	Timeline []*TimelineInformation `json:"Timeline,omitnil,omitempty" name:"Timeline"`

	// 核酸报告结论
	Covid []*CovidItemsInfo `json:"Covid,omitnil,omitempty" name:"Covid"`

	// 眼科报告结构体
	Eye []*EyeItemsInfo `json:"Eye,omitnil,omitempty" name:"Eye"`

	// 出生证明结构化信息
	BirthCert []*BirthCert `json:"BirthCert,omitnil,omitempty" name:"BirthCert"`

	// 文本类型列表
	TextTypeListBlocks []*TextTypeListBlock `json:"TextTypeListBlocks,omitnil,omitempty" name:"TextTypeListBlocks"`

	// 体检报告信息
	PhysicalExamination *PhysicalExaminationV1 `json:"PhysicalExamination,omitnil,omitempty" name:"PhysicalExamination"`

	// 内窥镜报告V2
	EndoscopyV2 []*Check `json:"EndoscopyV2,omitnil,omitempty" name:"EndoscopyV2"`
}

type BlockInfo struct {
	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 阳性
	Positive *string `json:"Positive,omitnil,omitempty" name:"Positive"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 大小
	Size []*Size `json:"Size,omitnil,omitempty" name:"Size"`
}

type BlockInfoV2 struct {
	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 疾病编码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

type BlockTitle struct {
	// name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// src
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type BloodPressureBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 舒张压
	NormDiastolic *string `json:"NormDiastolic,omitnil,omitempty" name:"NormDiastolic"`

	// 收缩压
	NormSystolic *string `json:"NormSystolic,omitnil,omitempty" name:"NormSystolic"`
}

type BloodPressureItem struct {
	// 类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目原文
	Item *PhysicalBaseItem `json:"Item,omitnil,omitempty" name:"Item"`

	// 数值
	Result *PhysicalBaseItem `json:"Result,omitnil,omitempty" name:"Result"`

	// 单位
	Unit *PhysicalBaseItem `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 第几次
	Times *PhysicalBaseItem `json:"Times,omitnil,omitempty" name:"Times"`

	// 左右手臂
	Location *PhysicalBaseItem `json:"Location,omitnil,omitempty" name:"Location"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type BodyExaminationBlock struct {
	// 体温
	BodyTemperature *BodyTemperatureBlock `json:"BodyTemperature,omitnil,omitempty" name:"BodyTemperature"`

	// 脉搏
	Pulse *BodyTemperatureBlock `json:"Pulse,omitnil,omitempty" name:"Pulse"`

	// 呼吸
	Breathe *BodyTemperatureBlock `json:"Breathe,omitnil,omitempty" name:"Breathe"`

	// 血压
	BloodPressure *BloodPressureBlock `json:"BloodPressure,omitnil,omitempty" name:"BloodPressure"`
}

type BodyTemperatureBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Check struct {
	// 描述
	Desc *Desc `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 结论
	Summary *Summary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 检查报告块标题
	BlockTitle []*BlockTitle `json:"BlockTitle,omitnil,omitempty" name:"BlockTitle"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type ChemicalProductInfo struct {
	// 药品名称，包括通用名和商品名
	Name *ChemicalProductInfoName `json:"Name,omitnil,omitempty" name:"Name"`

	// 活性成份消息定义，如果是复方制剂，可以不列出每个活性成份的详细信息
	ActiveIngredient *ChemicalProductInfoActiveIngredient `json:"ActiveIngredient,omitnil,omitempty" name:"ActiveIngredient"`

	// 性状
	Appearance *ChemicalProductInfoAppearance `json:"Appearance,omitnil,omitempty" name:"Appearance"`

	// 适应症描述
	Indications *ChemicalProductInfoIndications `json:"Indications,omitnil,omitempty" name:"Indications"`

	// 规格
	Brochure *ChemicalProductInfoBrochure `json:"Brochure,omitnil,omitempty" name:"Brochure"`

	// 用法用量
	Dosage *ChemicalProductInfoDosage `json:"Dosage,omitnil,omitempty" name:"Dosage"`

	// 不良反应
	AdverseReaction *ChemicalProductInfoAdverseReaction `json:"AdverseReaction,omitnil,omitempty" name:"AdverseReaction"`

	// 禁忌情况
	Contraindications *ChemicalProductInfoContraindications `json:"Contraindications,omitnil,omitempty" name:"Contraindications"`

	// 注意事项
	Precautions *ChemicalProductInfoPrecautions `json:"Precautions,omitnil,omitempty" name:"Precautions"`

	// 孕妇及哺乳期妇女用药
	PregnancyLactationUse *ChemicalProductInfoPregnancyLactationUse `json:"PregnancyLactationUse,omitnil,omitempty" name:"PregnancyLactationUse"`

	// 儿童用药
	PediatricUse *ChemicalProductInfoPediatricUse `json:"PediatricUse,omitnil,omitempty" name:"PediatricUse"`

	// 老年用药
	GeriatricUse *ChemicalProductInfoGeriatricUse `json:"GeriatricUse,omitnil,omitempty" name:"GeriatricUse"`

	// 药品的药物相互作用
	Interactions *ChemicalProductInfoInteractions `json:"Interactions,omitnil,omitempty" name:"Interactions"`

	// 药物过量
	Overdose *ChemicalProductInfoOverdose `json:"Overdose,omitnil,omitempty" name:"Overdose"`

	// 临床试验
	ClinicalTrial *ChemicalProductInfoClinicalTrial `json:"ClinicalTrial,omitnil,omitempty" name:"ClinicalTrial"`

	// 药理毒理
	PharmacologyToxicology *ChemicalProductInfoPharmacologyToxicology `json:"PharmacologyToxicology,omitnil,omitempty" name:"PharmacologyToxicology"`

	// 药代动力学
	Pharmacokinetics *ChemicalProductInfoPharmacokinetics `json:"Pharmacokinetics,omitnil,omitempty" name:"Pharmacokinetics"`

	// 储存条件
	Storage *ChemicalProductInfoStorage `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 包装信息
	Packaging *ChemicalProductInfoPackaging `json:"Packaging,omitnil,omitempty" name:"Packaging"`

	// 有效期
	ValidityPeriod *ChemicalProductInfoValidityPeriod `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// 执行标准
	ExecutiveStandards *ChemicalProductInfoExecutiveStandards `json:"ExecutiveStandards,omitnil,omitempty" name:"ExecutiveStandards"`

	// 批准文号
	Approval *ChemicalProductInfoApproval `json:"Approval,omitnil,omitempty" name:"Approval"`

	// 生产企业名称和地址
	Manufacturer *ChemicalProductInfoManufacturer `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`
}

type ChemicalProductInfoActiveIngredient struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 活性成份的化学名称
	ChemicalName *string `json:"ChemicalName,omitnil,omitempty" name:"ChemicalName"`

	// 活性成份的化学结构式
	ChemicalFormula *string `json:"ChemicalFormula,omitnil,omitempty" name:"ChemicalFormula"`

	// 活性成份的分子式
	MolecularFormula *string `json:"MolecularFormula,omitnil,omitempty" name:"MolecularFormula"`

	// 活性成份的分子量
	MolecularWeight *string `json:"MolecularWeight,omitnil,omitempty" name:"MolecularWeight"`
}

type ChemicalProductInfoAdverseReaction struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoAppearance struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoApproval struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoBrochure struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoClinicalTrial struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoContraindications struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoDosage struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoExecutiveStandards struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoGeriatricUse struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoIndications struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoInteractions struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoManufacturer struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 企业名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 生产地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 邮政编码
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// 电话，包含区号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 传真，包含区号
	Fax *string `json:"Fax,omitnil,omitempty" name:"Fax"`

	// 网址，如无则不填写
	Website *string `json:"Website,omitnil,omitempty" name:"Website"`
}

type ChemicalProductInfoName struct {
	// 文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 通用名
	GenericName *string `json:"GenericName,omitnil,omitempty" name:"GenericName"`

	// 品牌名
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: BarndName is deprecated.
	BarndName *string `json:"BarndName,omitnil,omitempty" name:"BarndName"`

	// 英文名
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`

	// 拼音
	Pinyin *string `json:"Pinyin,omitnil,omitempty" name:"Pinyin"`

	// 品牌名
	BrandName *string `json:"BrandName,omitnil,omitempty" name:"BrandName"`
}

type ChemicalProductInfoOverdose struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoPackaging struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoPediatricUse struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoPharmacokinetics struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoPharmacologyToxicology struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoPrecautions struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoPregnancyLactationUse struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoStorage struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChemicalProductInfoValidityPeriod struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ChestCircumferenceItem struct {
	// 类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目原文
	Item *PhysicalBaseItem `json:"Item,omitnil,omitempty" name:"Item"`

	// 数值
	Result *PhysicalBaseItem `json:"Result,omitnil,omitempty" name:"Result"`

	// 单位
	Unit *PhysicalBaseItem `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 呼吸状态
	State *PhysicalBaseItem `json:"State,omitnil,omitempty" name:"State"`
}

type ChiefComplaintBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 单位输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 主诉详情
	Detail []*ChiefComplaintDetailBlock `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type ChiefComplaintDetailBlock struct {
	// 疾病名称
	DiseaseName *string `json:"DiseaseName,omitnil,omitempty" name:"DiseaseName"`

	// 部位
	Part *string `json:"Part,omitnil,omitempty" name:"Part"`

	// 时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 时间类型
	TimeType *string `json:"TimeType,omitnil,omitempty" name:"TimeType"`
}

type ClinicalStaging struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Coord struct {
	// 坐标
	Points []*Point `json:"Points,omitnil,omitempty" name:"Points"`
}

type Coordinate struct {
	// 左上角x坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 左上角y坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 宽度，单位像素
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高度，单位像素
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type CovidItem struct {
	// 采样时间
	SampleTime *BaseItem `json:"SampleTime,omitnil,omitempty" name:"SampleTime"`

	// 检测时间
	TestTime *BaseItem `json:"TestTime,omitnil,omitempty" name:"TestTime"`

	// 检测机构
	TestOrganization *BaseItem `json:"TestOrganization,omitnil,omitempty" name:"TestOrganization"`

	// 检测结果
	TestResult *BaseItem `json:"TestResult,omitnil,omitempty" name:"TestResult"`

	// 健康码颜色
	CodeColor *BaseItem `json:"CodeColor,omitnil,omitempty" name:"CodeColor"`
}

type CovidItemsInfo struct {
	// 核酸报告结论
	CovidItems []*CovidItem `json:"CovidItems,omitnil,omitempty" name:"CovidItems"`

	// 版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type DeathDateBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 时间戳
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Desc struct {
	// 描述
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 器官
	Organ []*Organ `json:"Organ,omitnil,omitempty" name:"Organ"`

	// 结节
	Tuber []*TuberInfo `json:"Tuber,omitnil,omitempty" name:"Tuber"`

	// 坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type DescInfo struct {
	// 描述段落文本
	Text *BaseInfo `json:"Text,omitnil,omitempty" name:"Text"`

	// 描述段落详情
	Infos []*DetailInformation `json:"Infos,omitnil,omitempty" name:"Infos"`
}

type DetailInformation struct {
	// 部位
	Part *Part `json:"Part,omitnil,omitempty" name:"Part"`

	// 组织大小
	TissueSizes []*Size `json:"TissueSizes,omitnil,omitempty" name:"TissueSizes"`

	// 结节大小
	TuberSizes []*Size `json:"TuberSizes,omitnil,omitempty" name:"TuberSizes"`

	// 肿瘤大小
	CancerSizes []*Size `json:"CancerSizes,omitnil,omitempty" name:"CancerSizes"`

	// 组织学等级
	HistologyLevel *BaseInfo `json:"HistologyLevel,omitnil,omitempty" name:"HistologyLevel"`

	// 组织学类型
	HistologyType *HistologyTypeV2 `json:"HistologyType,omitnil,omitempty" name:"HistologyType"`

	// 侵犯
	Invasive []*InvasiveV2 `json:"Invasive,omitnil,omitempty" name:"Invasive"`

	// pTNM
	PTNM *PTNM `json:"PTNM,omitnil,omitempty" name:"PTNM"`

	// 浸润深度
	InfiltrationDepth *BaseInfo `json:"InfiltrationDepth,omitnil,omitempty" name:"InfiltrationDepth"`

	// 结节数量
	TuberNum *BaseInfo `json:"TuberNum,omitnil,omitempty" name:"TuberNum"`

	// 钙化
	Calcification *BaseInfo `json:"Calcification,omitnil,omitempty" name:"Calcification"`

	// 坏死
	Necrosis *BaseInfo `json:"Necrosis,omitnil,omitempty" name:"Necrosis"`

	// 异形
	Abnormity *BaseInfo `json:"Abnormity,omitnil,omitempty" name:"Abnormity"`

	// 断链
	Breaked *BaseInfo `json:"Breaked,omitnil,omitempty" name:"Breaked"`
}

type DiagCert struct {
	// 建议
	Advice *Advice `json:"Advice,omitnil,omitempty" name:"Advice"`

	// 诊断
	Diagnosis []*DiagCertItem `json:"Diagnosis,omitnil,omitempty" name:"Diagnosis"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type DiagCertItem struct {
	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 值
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DischargeConditionBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`
}

type DischargeDiagnosis struct {
	// 表格位置
	TableIndex *int64 `json:"TableIndex,omitnil,omitempty" name:"TableIndex"`

	// 出院诊断
	OutDiagnosis *BlockInfo `json:"OutDiagnosis,omitnil,omitempty" name:"OutDiagnosis"`

	// 疾病编码
	DiseaseCode *BlockInfo `json:"DiseaseCode,omitnil,omitempty" name:"DiseaseCode"`

	// 入院情况
	InStatus *BlockInfo `json:"InStatus,omitnil,omitempty" name:"InStatus"`

	// 出院情况
	OutStatus *BlockInfo `json:"OutStatus,omitnil,omitempty" name:"OutStatus"`
}

type DischargeDiagnosisBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DischargeInfoBlock struct {
	// 疾病史
	DiseaseHistory *DiseaseHistoryBlock `json:"DiseaseHistory,omitnil,omitempty" name:"DiseaseHistory"`

	// 个人史
	PersonalHistory *PersonalHistoryBlock `json:"PersonalHistory,omitnil,omitempty" name:"PersonalHistory"`

	// 药物史
	DrugHistory *DrugHistoryBlock `json:"DrugHistory,omitnil,omitempty" name:"DrugHistory"`

	// 治疗相关
	TreatmentRecord *TreatmentRecordBlock `json:"TreatmentRecord,omitnil,omitempty" name:"TreatmentRecord"`

	// 文本段落
	ParagraphBlock *ParagraphBlock `json:"ParagraphBlock,omitnil,omitempty" name:"ParagraphBlock"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type DiseaseHistoryBlock struct {
	// 主要病史
	MainDiseaseHistory *MainDiseaseHistoryBlock `json:"MainDiseaseHistory,omitnil,omitempty" name:"MainDiseaseHistory"`

	// 过敏史
	AllergyHistory *MainDiseaseHistoryBlock `json:"AllergyHistory,omitnil,omitempty" name:"AllergyHistory"`

	// 注射史
	InfectHistory *MainDiseaseHistoryBlock `json:"InfectHistory,omitnil,omitempty" name:"InfectHistory"`

	// 手术史
	SurgeryHistory *SurgeryHistoryBlock `json:"SurgeryHistory,omitnil,omitempty" name:"SurgeryHistory"`

	// 输血史
	TransfusionHistory *TransfusionHistoryBlock `json:"TransfusionHistory,omitnil,omitempty" name:"TransfusionHistory"`
}

type DiseaseMedicalHistory struct {
	// 主病史
	MainDiseaseHistory *string `json:"MainDiseaseHistory,omitnil,omitempty" name:"MainDiseaseHistory"`

	// 过敏史
	AllergyHistory *string `json:"AllergyHistory,omitnil,omitempty" name:"AllergyHistory"`

	// 传染疾病史
	InfectHistory *string `json:"InfectHistory,omitnil,omitempty" name:"InfectHistory"`

	// 手术史
	OperationHistory *string `json:"OperationHistory,omitnil,omitempty" name:"OperationHistory"`

	// 输血史
	TransfusionHistory *string `json:"TransfusionHistory,omitnil,omitempty" name:"TransfusionHistory"`

	// 现病史
	DiseasePresent *string `json:"DiseasePresent,omitnil,omitempty" name:"DiseasePresent"`
}

type DiseasePresentBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 归一化
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`
}

type DosageBlock struct {
	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 单次计量
	SingleMeasurement *string `json:"SingleMeasurement,omitnil,omitempty" name:"SingleMeasurement"`

	// 频次
	Frequency *string `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// 给药途径
	DrugDeliveryRoute *string `json:"DrugDeliveryRoute,omitnil,omitempty" name:"DrugDeliveryRoute"`
}

type DrugHistoryBlock struct {
	// 药品名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 药物列表
	DrugList []*DrugListBlock `json:"DrugList,omitnil,omitempty" name:"DrugList"`

	// 归一化值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type DrugInstructionObjectRequestParams struct {
	// 药品说明书PDF文件信息, 目前只支持传PDF文件的Base64编码字符(PDF文件不能超过10MB，如果超过建议先压缩PDF，再转成base64).
	PdfInfo *PdfInfo `json:"PdfInfo,omitnil,omitempty" name:"PdfInfo"`
}

type DrugInstructionObjectRequest struct {
	*tchttp.BaseRequest
	
	// 药品说明书PDF文件信息, 目前只支持传PDF文件的Base64编码字符(PDF文件不能超过10MB，如果超过建议先压缩PDF，再转成base64).
	PdfInfo *PdfInfo `json:"PdfInfo,omitnil,omitempty" name:"PdfInfo"`
}

func (r *DrugInstructionObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DrugInstructionObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PdfInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DrugInstructionObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DrugInstructionObjectResponseParams struct {
	// 药品说明书消息定义
	ChemicalProductInfo *ChemicalProductInfo `json:"ChemicalProductInfo,omitnil,omitempty" name:"ChemicalProductInfo"`

	// 预防用生物制品说明书
	BiologicalProductInfo *BiologicalProductInfo `json:"BiologicalProductInfo,omitnil,omitempty" name:"BiologicalProductInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DrugInstructionObjectResponse struct {
	*tchttp.BaseResponse
	Response *DrugInstructionObjectResponseParams `json:"Response"`
}

func (r *DrugInstructionObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DrugInstructionObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrugListBlock struct {
	// 通用名称
	CommonName *string `json:"CommonName,omitnil,omitempty" name:"CommonName"`

	// 商品名称
	TradeName *string `json:"TradeName,omitnil,omitempty" name:"TradeName"`

	// 用法用量
	Dosage *DosageBlock `json:"Dosage,omitnil,omitempty" name:"Dosage"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type EcgDescription struct {
	// 心率
	HeartRate *EcgItem `json:"HeartRate,omitnil,omitempty" name:"HeartRate"`

	// 心房率
	AuricularRate *EcgItem `json:"AuricularRate,omitnil,omitempty" name:"AuricularRate"`

	// 心室率
	VentricularRate *EcgItem `json:"VentricularRate,omitnil,omitempty" name:"VentricularRate"`

	// 节律
	Rhythm *EcgItem `json:"Rhythm,omitnil,omitempty" name:"Rhythm"`

	// P波时间
	PDuration *EcgItem `json:"PDuration,omitnil,omitempty" name:"PDuration"`

	// QRS时间
	QrsDuration *EcgItem `json:"QrsDuration,omitnil,omitempty" name:"QrsDuration"`

	// QRS电轴
	QrsAxis *EcgItem `json:"QrsAxis,omitnil,omitempty" name:"QrsAxis"`

	// P-R间期
	PRInterval *EcgItem `json:"PRInterval,omitnil,omitempty" name:"PRInterval"`

	// P-P间期
	PPInterval *EcgItem `json:"PPInterval,omitnil,omitempty" name:"PPInterval"`

	// R-R间期
	RRInterval *EcgItem `json:"RRInterval,omitnil,omitempty" name:"RRInterval"`

	// P-J间期
	PJInterval *EcgItem `json:"PJInterval,omitnil,omitempty" name:"PJInterval"`

	// Q-T间期
	QTInterval *EcgItem `json:"QTInterval,omitnil,omitempty" name:"QTInterval"`

	// qt/qtc间期
	QTCInterval *EcgItem `json:"QTCInterval,omitnil,omitempty" name:"QTCInterval"`

	// RV5/SV1振幅
	Rv5SV1Amplitude *EcgItem `json:"Rv5SV1Amplitude,omitnil,omitempty" name:"Rv5SV1Amplitude"`

	// RV5+SV1振幅
	Rv5AddSV1Amplitude *EcgItem `json:"Rv5AddSV1Amplitude,omitnil,omitempty" name:"Rv5AddSV1Amplitude"`

	// PRT电轴
	PRTAxis *EcgItem `json:"PRTAxis,omitnil,omitempty" name:"PRTAxis"`

	// RV5振幅
	Rv5Amplitude *EcgItem `json:"Rv5Amplitude,omitnil,omitempty" name:"Rv5Amplitude"`

	// SV1振幅
	SV1Amplitude *EcgItem `json:"SV1Amplitude,omitnil,omitempty" name:"SV1Amplitude"`

	// RV6/SV2
	RV6SV2 *EcgItem `json:"RV6SV2,omitnil,omitempty" name:"RV6SV2"`

	// P/QRS/T电轴
	PQRSTAxis *EcgItem `json:"PQRSTAxis,omitnil,omitempty" name:"PQRSTAxis"`
}

type EcgDiagnosis struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type EcgItem struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`
}

type Elastic struct {
	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 分数
	Score *string `json:"Score,omitnil,omitempty" name:"Score"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type Electrocardiogram struct {
	// 心电图详情
	EcgDescription *EcgDescription `json:"EcgDescription,omitnil,omitempty" name:"EcgDescription"`

	// 心电图诊断
	EcgDiagnosis *EcgDiagnosis `json:"EcgDiagnosis,omitnil,omitempty" name:"EcgDiagnosis"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type Endoscopy struct {
	// 活检部位
	BiopsyPart *BiopsyPart `json:"BiopsyPart,omitnil,omitempty" name:"BiopsyPart"`

	// 可见描述
	Desc *EndoscopyDesc `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 结论
	Summary *Summary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type EndoscopyDesc struct {
	// 描述内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 器官
	Organ []*EndoscopyOrgan `json:"Organ,omitnil,omitempty" name:"Organ"`

	// 坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type EndoscopyOrgan struct {
	// 部位
	Part *Part `json:"Part,omitnil,omitempty" name:"Part"`

	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 部位别名
	PartAlias *string `json:"PartAlias,omitnil,omitempty" name:"PartAlias"`

	// 症状描述
	SymDescList []*BlockInfo `json:"SymDescList,omitnil,omitempty" name:"SymDescList"`

	// 坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type Exame struct {
	// 结论段落
	OverView []*ResultInfo `json:"OverView,omitnil,omitempty" name:"OverView"`

	// 异常与建议段落
	Abnormality []*ResultInfo `json:"Abnormality,omitnil,omitempty" name:"Abnormality"`
}

type EyeChildItem struct {
	// 球镜
	Sph []*BaseItem3 `json:"Sph,omitnil,omitempty" name:"Sph"`

	// 柱镜
	Cyl []*BaseItem3 `json:"Cyl,omitnil,omitempty" name:"Cyl"`

	// 轴位
	Ax []*BaseItem3 `json:"Ax,omitnil,omitempty" name:"Ax"`

	// 等效球镜
	Se *BaseItem2 `json:"Se,omitnil,omitempty" name:"Se"`
}

type EyeItem struct {
	// 左眼
	Left *EyeChildItem `json:"Left,omitnil,omitempty" name:"Left"`

	// 右眼
	Right *EyeChildItem `json:"Right,omitnil,omitempty" name:"Right"`

	// 瞳距
	Pd *BaseItem2 `json:"Pd,omitnil,omitempty" name:"Pd"`
}

type EyeItemsInfo struct {
	// 眼科报告
	EyeItems *EyeItem `json:"EyeItems,omitnil,omitempty" name:"EyeItems"`

	// 版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type FamilyHistoryBlock struct {
	// 家庭成员
	RelativeHistory *RelativeHistoryBlock `json:"RelativeHistory,omitnil,omitempty" name:"RelativeHistory"`

	// 家族肿瘤史
	RelativeCancerHistory *RelativeCancerHistoryBlock `json:"RelativeCancerHistory,omitnil,omitempty" name:"RelativeCancerHistory"`

	// 家族遗传史
	GeneticHistory *GeneticHistoryBlock `json:"GeneticHistory,omitnil,omitempty" name:"GeneticHistory"`
}

type FamilyMedicalHistory struct {
	// 家族成员史
	RelativeHistory *string `json:"RelativeHistory,omitnil,omitempty" name:"RelativeHistory"`

	// 家族肿瘤史
	RelativeCancerHistory *string `json:"RelativeCancerHistory,omitnil,omitempty" name:"RelativeCancerHistory"`

	// 家族遗传史
	GeneticHistory *string `json:"GeneticHistory,omitnil,omitempty" name:"GeneticHistory"`
}

type FertilityHistoryBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 妊娠次数
	PregCount *string `json:"PregCount,omitnil,omitempty" name:"PregCount"`

	// 生产次数
	ProduCount *string `json:"ProduCount,omitnil,omitempty" name:"ProduCount"`
}

type Fetus struct {
	// 双顶径
	BPD *FieldInfo `json:"BPD,omitnil,omitempty" name:"BPD"`

	// 腹前后径
	APTD *FieldInfo `json:"APTD,omitnil,omitempty" name:"APTD"`

	// 腹左右径
	TTD *FieldInfo `json:"TTD,omitnil,omitempty" name:"TTD"`

	// 头臀径
	CRL *FieldInfo `json:"CRL,omitnil,omitempty" name:"CRL"`

	// 头围
	HC *FieldInfo `json:"HC,omitnil,omitempty" name:"HC"`

	// 腹围
	AC *FieldInfo `json:"AC,omitnil,omitempty" name:"AC"`

	// 股骨长
	FL *FieldInfo `json:"FL,omitnil,omitempty" name:"FL"`

	// 肱骨长
	HL *FieldInfo `json:"HL,omitnil,omitempty" name:"HL"`

	// 胎儿重量
	Weight *FieldInfo `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 颈项透明层
	NT *FieldInfo `json:"NT,omitnil,omitempty" name:"NT"`

	// 脐动脉血流
	UmbilicalCord *FieldInfo `json:"UmbilicalCord,omitnil,omitempty" name:"UmbilicalCord"`

	// 羊水最大深度
	WaterDeep *FieldInfo `json:"WaterDeep,omitnil,omitempty" name:"WaterDeep"`

	// 羊水四象限测量
	WaterQuad *FieldInfo `json:"WaterQuad,omitnil,omitempty" name:"WaterQuad"`

	// 羊水指数
	AFI *FieldInfo `json:"AFI,omitnil,omitempty" name:"AFI"`

	// 胎心
	FHR *FieldInfo `json:"FHR,omitnil,omitempty" name:"FHR"`

	// 胎动
	Movement *FieldInfo `json:"Movement,omitnil,omitempty" name:"Movement"`

	// 胎数
	Num *FieldInfo `json:"Num,omitnil,omitempty" name:"Num"`

	// 胎位
	Position *FieldInfo `json:"Position,omitnil,omitempty" name:"Position"`

	// 是否活胎
	Alive *FieldInfo `json:"Alive,omitnil,omitempty" name:"Alive"`

	// 胎盘位置
	PlacentaLocation *FieldInfo `json:"PlacentaLocation,omitnil,omitempty" name:"PlacentaLocation"`

	// 胎盘厚度
	PlacentaThickness *FieldInfo `json:"PlacentaThickness,omitnil,omitempty" name:"PlacentaThickness"`

	// 胎盘成熟度
	PlacentaGrade *FieldInfo `json:"PlacentaGrade,omitnil,omitempty" name:"PlacentaGrade"`

	// 妊娠时间
	GestationTime *FieldInfo `json:"GestationTime,omitnil,omitempty" name:"GestationTime"`

	// 妊娠周期
	GestationPeriod *FieldInfo `json:"GestationPeriod,omitnil,omitempty" name:"GestationPeriod"`

	// 绕颈
	AroundNeck *FieldInfo `json:"AroundNeck,omitnil,omitempty" name:"AroundNeck"`

	// 病变
	Sym []*FieldInfo `json:"Sym,omitnil,omitempty" name:"Sym"`

	// 原文内容
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`
}

type FieldInfo struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 数值
	Nums []*NumValue `json:"Nums,omitnil,omitempty" name:"Nums"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`
}

type FirstPage struct {
	// 出入院诊断
	DischargeDiagnosis []*DischargeDiagnosis `json:"DischargeDiagnosis,omitnil,omitempty" name:"DischargeDiagnosis"`

	// 病理诊断
	PathologicalDiagnosis *BlockInfo `json:"PathologicalDiagnosis,omitnil,omitempty" name:"PathologicalDiagnosis"`

	// 临床诊断
	ClinicalDiagnosis *BlockInfo `json:"ClinicalDiagnosis,omitnil,omitempty" name:"ClinicalDiagnosis"`

	// 受伤中毒的外部原因
	DamagePoi *BlockInfoV2 `json:"DamagePoi,omitnil,omitempty" name:"DamagePoi"`

	// 病案首页第二页
	Fp2NdItems []*Fp2NdItem `json:"Fp2NdItems,omitnil,omitempty" name:"Fp2NdItems"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type Fp2NdItem struct {
	// 手术编码
	Code *BaseItem `json:"Code,omitnil,omitempty" name:"Code"`

	// 手术名称
	Name *BaseItem `json:"Name,omitnil,omitempty" name:"Name"`

	// 手术开始时间
	StartTime *BaseItem `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 手术结束时间
	EndTime *BaseItem `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 手术等级
	Level *BaseItem `json:"Level,omitnil,omitempty" name:"Level"`

	// 手术类型
	Type *BaseItem `json:"Type,omitnil,omitempty" name:"Type"`

	// 醉愈合方式
	IncisionHealingGrade *BaseItem `json:"IncisionHealingGrade,omitnil,omitempty" name:"IncisionHealingGrade"`

	// 麻醉方法
	AnesthesiaMethod *BaseItem `json:"AnesthesiaMethod,omitnil,omitempty" name:"AnesthesiaMethod"`
}

type GeneralExaminationBaseItem struct {
	// 生命体征
	VitalSign *GeneralExaminationVitalSign `json:"VitalSign,omitnil,omitempty" name:"VitalSign"`

	// 其他
	Others *GeneralExaminationOthers `json:"Others,omitnil,omitempty" name:"Others"`

	// 小结
	BriefSummary *GeneralExaminationBriefSummary `json:"BriefSummary,omitnil,omitempty" name:"BriefSummary"`
}

type GeneralExaminationBriefSummary struct {
	// 一般检查小结
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type GeneralExaminationOthers struct {
	// 面容与表情
	Countenance *KeyValueItem `json:"Countenance,omitnil,omitempty" name:"Countenance"`

	// 精神状态
	MentalStatus *KeyValueItem `json:"MentalStatus,omitnil,omitempty" name:"MentalStatus"`

	// 发育及营养状况
	DevelopmentCondition *KeyValueItem `json:"DevelopmentCondition,omitnil,omitempty" name:"DevelopmentCondition"`

	// 记忆力
	Memory *KeyValueItem `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 臀围
	Hipline *ValueUnitItem `json:"Hipline,omitnil,omitempty" name:"Hipline"`

	// 腰臀比
	WaistHipRatio *ValueUnitItem `json:"WaistHipRatio,omitnil,omitempty" name:"WaistHipRatio"`

	// 生活嗜好
	Addiction *KeyValueItem `json:"Addiction,omitnil,omitempty" name:"Addiction"`

	// 生活能力评定
	AbilityOfLifeADL *KeyValueItem `json:"AbilityOfLifeADL,omitnil,omitempty" name:"AbilityOfLifeADL"`

	// 一般检查其他
	Others []*KeyValueItem `json:"Others,omitnil,omitempty" name:"Others"`

	// 胸围
	ChestCircumference *ChestCircumferenceItem `json:"ChestCircumference,omitnil,omitempty" name:"ChestCircumference"`
}

type GeneralExaminationVitalSign struct {
	// 生命体征总体描述
	Text *ValueUnitItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 体温
	BodyTemperature *ValueUnitItem `json:"BodyTemperature,omitnil,omitempty" name:"BodyTemperature"`

	// 脉率
	Pulse *ValueUnitItem `json:"Pulse,omitnil,omitempty" name:"Pulse"`

	// 心率
	HeartRate *ValueUnitItem `json:"HeartRate,omitnil,omitempty" name:"HeartRate"`

	// 呼吸频率
	BreathingRate *ValueUnitItem `json:"BreathingRate,omitnil,omitempty" name:"BreathingRate"`

	// 身高
	BodyHeight *ValueUnitItem `json:"BodyHeight,omitnil,omitempty" name:"BodyHeight"`

	// 体重
	BodyWeight *ValueUnitItem `json:"BodyWeight,omitnil,omitempty" name:"BodyWeight"`

	// 体质指数
	BodyMassIndex *ValueUnitItem `json:"BodyMassIndex,omitnil,omitempty" name:"BodyMassIndex"`

	// 腰围
	Waistline *ValueUnitItem `json:"Waistline,omitnil,omitempty" name:"Waistline"`

	// 血压
	BloodPressure *GeneralExaminationVitalSignBloodPressure `json:"BloodPressure,omitnil,omitempty" name:"BloodPressure"`
}

type GeneralExaminationVitalSignBloodPressure struct {
	// 血压
	Text *BloodPressureItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 收缩压/舒张压
	SystolicDiastolicPressure []*BloodPressureItem `json:"SystolicDiastolicPressure,omitnil,omitempty" name:"SystolicDiastolicPressure"`

	// 收缩压
	SystolicPressure []*BloodPressureItem `json:"SystolicPressure,omitnil,omitempty" name:"SystolicPressure"`

	// 舒张压
	DiastolicPressure []*BloodPressureItem `json:"DiastolicPressure,omitnil,omitempty" name:"DiastolicPressure"`
}

type GeneticHistoryBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 遗传列表
	GeneticList *string `json:"GeneticList,omitnil,omitempty" name:"GeneticList"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type GynaecologyAdnexal struct {
	// 子宫附件总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type GynaecologyBaseItem struct {
	// 外阴
	Vulva *GynaecologyVulva `json:"Vulva,omitnil,omitempty" name:"Vulva"`

	// 阴道
	Vagina *GynaecologyVagina `json:"Vagina,omitnil,omitempty" name:"Vagina"`

	// 子宫颈
	Cervix *GynaecologyCervix `json:"Cervix,omitnil,omitempty" name:"Cervix"`

	// 子宫
	Uterus *GynaecologyUterus `json:"Uterus,omitnil,omitempty" name:"Uterus"`

	// 子宫附件
	Adnexal *GynaecologyAdnexal `json:"Adnexal,omitnil,omitempty" name:"Adnexal"`

	// 盆腔
	PelvicCavity *GynaecologyPelvicCavity `json:"PelvicCavity,omitnil,omitempty" name:"PelvicCavity"`

	// 妇科其他
	Others []*KeyValueItem `json:"Others,omitnil,omitempty" name:"Others"`

	// 月经史
	MenstrualHistory *GynaecologyMenstrualHistory `json:"MenstrualHistory,omitnil,omitempty" name:"MenstrualHistory"`

	// 小结
	BriefSummary *GynaecologyBriefSummary `json:"BriefSummary,omitnil,omitempty" name:"BriefSummary"`
}

type GynaecologyBriefSummary struct {
	// 小结
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type GynaecologyCervix struct {
	// 子宫颈总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type GynaecologyMenstrualHistory struct {
	// 妇科月经史总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type GynaecologyPelvicCavity struct {
	// 盆腔总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type GynaecologyUterus struct {
	// 子宫总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type GynaecologyVagina struct {
	// 阴道总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type GynaecologyVulva struct {
	// 外阴总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type HandleParam struct {
	// ocr引擎
	OcrEngineType *int64 `json:"OcrEngineType,omitnil,omitempty" name:"OcrEngineType"`

	// 是否返回分行文本内容
	IsReturnText *bool `json:"IsReturnText,omitnil,omitempty" name:"IsReturnText"`

	// 顺时针旋转角度
	RotateTheAngle *float64 `json:"RotateTheAngle,omitnil,omitempty" name:"RotateTheAngle"`

	// 自动适配方向,仅支持优图引擎
	AutoFitDirection *bool `json:"AutoFitDirection,omitnil,omitempty" name:"AutoFitDirection"`

	// 坐标优化
	AutoOptimizeCoordinate *bool `json:"AutoOptimizeCoordinate,omitnil,omitempty" name:"AutoOptimizeCoordinate"`

	// 是否开启图片压缩，开启时imageOriginalSize必须正确填写
	IsScale *bool `json:"IsScale,omitnil,omitempty" name:"IsScale"`

	// 原始图片大小(单位byes),用来判断该图片是否需要压缩
	ImageOriginalSize *uint64 `json:"ImageOriginalSize,omitnil,omitempty" name:"ImageOriginalSize"`

	// 采用后台默认值(2048Kb)
	ScaleTargetSize *uint64 `json:"ScaleTargetSize,omitnil,omitempty" name:"ScaleTargetSize"`
}

type HearingItem struct {
	// 类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目原文
	Item *PhysicalBaseItem `json:"Item,omitnil,omitempty" name:"Item"`

	// 方位
	Location *PhysicalBaseItem `json:"Location,omitnil,omitempty" name:"Location"`

	// 描述
	Result *PhysicalBaseItem `json:"Result,omitnil,omitempty" name:"Result"`
}

type HistologyClass struct {
	// 项目名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 原文对应坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type HistologyLevel struct {
	// 等级
	Grade *string `json:"Grade,omitnil,omitempty" name:"Grade"`

	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`
}

type HistologyType struct {
	// 浸润
	Infiltration *string `json:"Infiltration,omitnil,omitempty" name:"Infiltration"`

	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type HistologyTypeV2 struct {
	// 浸润
	Infiltration *string `json:"Infiltration,omitnil,omitempty" name:"Infiltration"`

	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化后的组织学类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 项目名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文对应坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type Hospitalization struct {
	// 入院时间
	AdmissionTime *string `json:"AdmissionTime,omitnil,omitempty" name:"AdmissionTime"`

	// 出院时间
	DischargeTime *string `json:"DischargeTime,omitnil,omitempty" name:"DischargeTime"`

	// 住院天数
	AdmissionDays *string `json:"AdmissionDays,omitnil,omitempty" name:"AdmissionDays"`

	// 入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: AdmissionDignosis is deprecated.
	AdmissionDignosis *string `json:"AdmissionDignosis,omitnil,omitempty" name:"AdmissionDignosis"`

	// 入院情况
	AdmissionCondition *string `json:"AdmissionCondition,omitnil,omitempty" name:"AdmissionCondition"`

	// 诊疗经过
	DiagnosisTreatment *string `json:"DiagnosisTreatment,omitnil,omitempty" name:"DiagnosisTreatment"`

	// 出院诊断
	DischargeDiagnosis *string `json:"DischargeDiagnosis,omitnil,omitempty" name:"DischargeDiagnosis"`

	// 出院医嘱
	DischargeInstruction *string `json:"DischargeInstruction,omitnil,omitempty" name:"DischargeInstruction"`

	// 入院诊断
	AdmissionDiagnosis *string `json:"AdmissionDiagnosis,omitnil,omitempty" name:"AdmissionDiagnosis"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type IHCBlock struct {
	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 具体值
	Value *ValueBlock `json:"Value,omitnil,omitempty" name:"Value"`

	// 坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type IHCInfo struct {
	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value *Value `json:"Value,omitnil,omitempty" name:"Value"`
}

type IHCV2 struct {
	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// ihc归一化
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// ihc详情
	Value *Value `json:"Value,omitnil,omitempty" name:"Value"`

	// 原文对应坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type ImageInfo struct {
	// 图片id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 图片url(暂不支持传图片Url信息,请使用Base64字段传递图片的Base64编码)
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 图片base64编码
	Base64 *string `json:"Base64,omitnil,omitempty" name:"Base64"`
}

// Predefined struct for user
type ImageMaskAsyncGetResultRequestParams struct {
	// 异步任务ID
	// 1.因上传图片TaskID 仅 24 小时内有效，请于提交任务获取TaskID后24小时内调用ImageMaskAsyncGetResult接口获取结果.
	// 2.建议在获取到TaskID 后，5-10分钟后再调用 ImageMaskAsyncGetResult 接口获取脱敏结果。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

type ImageMaskAsyncGetResultRequest struct {
	*tchttp.BaseRequest
	
	// 异步任务ID
	// 1.因上传图片TaskID 仅 24 小时内有效，请于提交任务获取TaskID后24小时内调用ImageMaskAsyncGetResult接口获取结果.
	// 2.建议在获取到TaskID 后，5-10分钟后再调用 ImageMaskAsyncGetResult 接口获取脱敏结果。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

func (r *ImageMaskAsyncGetResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskAsyncGetResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageMaskAsyncGetResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageMaskAsyncGetResultResponseParams struct {
	// 脱敏后图片的base64编码
	MaskedImage *string `json:"MaskedImage,omitnil,omitempty" name:"MaskedImage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageMaskAsyncGetResultResponse struct {
	*tchttp.BaseResponse
	Response *ImageMaskAsyncGetResultResponseParams `json:"Response"`
}

func (r *ImageMaskAsyncGetResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskAsyncGetResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageMaskAsyncRequestParams struct {
	// 图片信息,目前只支持传图片base64
	Image *ImageInfo `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片脱敏选项, 不传默认都脱敏
	MaskFlag *ImageMaskFlags `json:"MaskFlag,omitnil,omitempty" name:"MaskFlag"`

	// 是否自动矫正图片方向
	AutoFixImageDirection *bool `json:"AutoFixImageDirection,omitnil,omitempty" name:"AutoFixImageDirection"`
}

type ImageMaskAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 图片信息,目前只支持传图片base64
	Image *ImageInfo `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片脱敏选项, 不传默认都脱敏
	MaskFlag *ImageMaskFlags `json:"MaskFlag,omitnil,omitempty" name:"MaskFlag"`

	// 是否自动矫正图片方向
	AutoFixImageDirection *bool `json:"AutoFixImageDirection,omitnil,omitempty" name:"AutoFixImageDirection"`
}

func (r *ImageMaskAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "MaskFlag")
	delete(f, "AutoFixImageDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageMaskAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageMaskAsyncResponseParams struct {
	// 加密任务ID
	// 1.因此TaskID 仅 24 小时内有效，请于提交任务获取TaskID后24小时内调用ImageMaskAsyncGetResult接口获取结果.
	// 2.建议在获取到TaskID 后，5-10分钟后再调用 ImageMaskAsyncGetResult 接口获取脱敏结果。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageMaskAsyncResponse struct {
	*tchttp.BaseResponse
	Response *ImageMaskAsyncResponseParams `json:"Response"`
}

func (r *ImageMaskAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageMaskFlags struct {
	// 是否对医院信息进行脱敏
	HospitalFlag *bool `json:"HospitalFlag,omitnil,omitempty" name:"HospitalFlag"`

	// 是否对医生信息进行脱敏
	DoctorFlag *bool `json:"DoctorFlag,omitnil,omitempty" name:"DoctorFlag"`

	// 是否对患者信息进行脱敏
	PatientFlag *bool `json:"PatientFlag,omitnil,omitempty" name:"PatientFlag"`

	// 是否对二维码信息进行脱敏
	BarFlag *bool `json:"BarFlag,omitnil,omitempty" name:"BarFlag"`
}

// Predefined struct for user
type ImageMaskRequestParams struct {
	// 图片信息,目前只支持传图片base64
	Image *ImageInfo `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片脱敏选项, 不传默认都脱敏
	MaskFlag *ImageMaskFlags `json:"MaskFlag,omitnil,omitempty" name:"MaskFlag"`

	// 是否自动矫正图片方向
	AutoFixImageDirection *bool `json:"AutoFixImageDirection,omitnil,omitempty" name:"AutoFixImageDirection"`
}

type ImageMaskRequest struct {
	*tchttp.BaseRequest
	
	// 图片信息,目前只支持传图片base64
	Image *ImageInfo `json:"Image,omitnil,omitempty" name:"Image"`

	// 图片脱敏选项, 不传默认都脱敏
	MaskFlag *ImageMaskFlags `json:"MaskFlag,omitnil,omitempty" name:"MaskFlag"`

	// 是否自动矫正图片方向
	AutoFixImageDirection *bool `json:"AutoFixImageDirection,omitnil,omitempty" name:"AutoFixImageDirection"`
}

func (r *ImageMaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "MaskFlag")
	delete(f, "AutoFixImageDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageMaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageMaskResponseParams struct {
	// 脱敏后图片的Base64信息
	MaskedImage *string `json:"MaskedImage,omitnil,omitempty" name:"MaskedImage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageMaskResponse struct {
	*tchttp.BaseResponse
	Response *ImageMaskResponseParams `json:"Response"`
}

func (r *ImageMaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToClassRequestParams struct {
	// 图片列表，允许传入多张图片，支持传入图片的base64编码，暂不支持图片url
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitnil,omitempty" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitnil,omitempty" name:"HandleParam"`

	// 不填，默认为0
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil,omitempty" name:"UserType"`
}

type ImageToClassRequest struct {
	*tchttp.BaseRequest
	
	// 图片列表，允许传入多张图片，支持传入图片的base64编码，暂不支持图片url
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitnil,omitempty" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitnil,omitempty" name:"HandleParam"`

	// 不填，默认为0
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil,omitempty" name:"UserType"`
}

func (r *ImageToClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfoList")
	delete(f, "HandleParam")
	delete(f, "Type")
	delete(f, "UserType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageToClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToClassResponseParams struct {
	// 分类结果
	TextTypeList []*TextType `json:"TextTypeList,omitnil,omitempty" name:"TextTypeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageToClassResponse struct {
	*tchttp.BaseResponse
	Response *ImageToClassResponseParams `json:"Response"`
}

func (r *ImageToClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToObjectRequestParams struct {
	// 报告类型，目前支持11（检验报告），12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为 False，则 Type 字段不能为 0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitnil,omitempty" name:"IsUsedClassify"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitnil,omitempty" name:"HandleParam"`

	// 图片列表，目前只支持传入一张图片，需要是图片base64编码，图片url暂不支持
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitnil,omitempty" name:"ImageInfoList"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *int64 `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病历记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitnil,omitempty" name:"ReportTypeVersion"`

	// 可选。 图片OCR信息列表，每一个元素是一张图片的OCR结果。适用于不想将医疗报告图片传入腾讯云的客户，客户可对图片OCR信息中的敏感信息去除之后再传入。与 ImageInfoList 二选一，同时存在则使用OcrInfoList
	OcrInfoList []*OcrInfo `json:"OcrInfoList,omitnil,omitempty" name:"OcrInfoList"`
}

type ImageToObjectRequest struct {
	*tchttp.BaseRequest
	
	// 报告类型，目前支持11（检验报告），12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为 False，则 Type 字段不能为 0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitnil,omitempty" name:"IsUsedClassify"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitnil,omitempty" name:"HandleParam"`

	// 图片列表，目前只支持传入一张图片，需要是图片base64编码，图片url暂不支持
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitnil,omitempty" name:"ImageInfoList"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *int64 `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病历记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitnil,omitempty" name:"ReportTypeVersion"`

	// 可选。 图片OCR信息列表，每一个元素是一张图片的OCR结果。适用于不想将医疗报告图片传入腾讯云的客户，客户可对图片OCR信息中的敏感信息去除之后再传入。与 ImageInfoList 二选一，同时存在则使用OcrInfoList
	OcrInfoList []*OcrInfo `json:"OcrInfoList,omitnil,omitempty" name:"OcrInfoList"`
}

func (r *ImageToObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "IsUsedClassify")
	delete(f, "HandleParam")
	delete(f, "ImageInfoList")
	delete(f, "UserType")
	delete(f, "ReportTypeVersion")
	delete(f, "OcrInfoList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageToObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToObjectResponseParams struct {
	// 报告结构化结果
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// 多级分类结果
	TextTypeList []*TextType `json:"TextTypeList,omitnil,omitempty" name:"TextTypeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageToObjectResponse struct {
	*tchttp.BaseResponse
	Response *ImageToObjectResponseParams `json:"Response"`
}

func (r *ImageToObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImmunohistochemistryBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 免疫组化详情
	Value []*IHCBlock `json:"Value,omitnil,omitempty" name:"Value"`
}

type Indicator struct {
	// 检验指标项
	Indicators []*IndicatorItem `json:"Indicators,omitnil,omitempty" name:"Indicators"`

	// 检验报告块标题
	BlockTitle []*BlockTitle `json:"BlockTitle,omitnil,omitempty" name:"BlockTitle"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type IndicatorItem struct {
	// 英文缩写
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 标准缩写
	Scode *string `json:"Scode,omitnil,omitempty" name:"Scode"`

	// 项目名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标准名
	Sname *string `json:"Sname,omitnil,omitempty" name:"Sname"`

	// 结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 参考范围
	Range *string `json:"Range,omitnil,omitempty" name:"Range"`

	// 上下箭头
	Arrow *string `json:"Arrow,omitnil,omitempty" name:"Arrow"`

	// 是否正常
	Normal *bool `json:"Normal,omitnil,omitempty" name:"Normal"`

	// 项目原文
	ItemString *string `json:"ItemString,omitnil,omitempty" name:"ItemString"`

	// 指标项ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 指标项坐标位置
	Coords *Coordinate `json:"Coords,omitnil,omitempty" name:"Coords"`

	// 推测结果是否异常
	InferNormal *string `json:"InferNormal,omitnil,omitempty" name:"InferNormal"`

	// 标本
	Sample *string `json:"Sample,omitnil,omitempty" name:"Sample"`

	// 检测方法
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`
}

type IndicatorItemV2 struct {
	// 项目名称
	Item *BaseItem `json:"Item,omitnil,omitempty" name:"Item"`

	// 英文编码
	Code *BaseItem `json:"Code,omitnil,omitempty" name:"Code"`

	// 结果
	Result *BaseItem `json:"Result,omitnil,omitempty" name:"Result"`

	// 单位
	Unit *BaseItem `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 参考范围
	Range *BaseItem `json:"Range,omitnil,omitempty" name:"Range"`

	// 上下箭头
	Arrow *BaseItem `json:"Arrow,omitnil,omitempty" name:"Arrow"`

	// 检测方法
	Method *BaseItem `json:"Method,omitnil,omitempty" name:"Method"`

	// 结果是否异常
	Normal *bool `json:"Normal,omitnil,omitempty" name:"Normal"`

	// ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 序号
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 推测结果是否异常
	InferNormal *string `json:"InferNormal,omitnil,omitempty" name:"InferNormal"`
}

type IndicatorV3 struct {
	// 检验报告V3结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: TableIndictors is deprecated.
	TableIndictors []*TableIndicators `json:"TableIndictors,omitnil,omitempty" name:"TableIndictors"`

	// 版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 检验报告V3结论
	TableIndicators []*TableIndicators `json:"TableIndicators,omitnil,omitempty" name:"TableIndicators"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type InternalMedicineAbdomen struct {
	// 内科腹部小结
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 肝脏
	Liver *InternalMedicineAbdomenLiver `json:"Liver,omitnil,omitempty" name:"Liver"`

	// 胆囊
	GallBladder *InternalMedicineAbdomenGallBladder `json:"GallBladder,omitnil,omitempty" name:"GallBladder"`

	// 胰腺
	Pancreas *InternalMedicineAbdomenPancreas `json:"Pancreas,omitnil,omitempty" name:"Pancreas"`

	// 脾脏
	Spleen *InternalMedicineAbdomenSpleen `json:"Spleen,omitnil,omitempty" name:"Spleen"`

	// 肾脏
	Kidney *InternalMedicineAbdomenKidney `json:"Kidney,omitnil,omitempty" name:"Kidney"`

	// 腹部其他
	Others []*KeyValueItem `json:"Others,omitnil,omitempty" name:"Others"`
}

type InternalMedicineAbdomenGallBladder struct {
	// 胆囊总体描述
	Src *KeyValueItem `json:"Src,omitnil,omitempty" name:"Src"`

	// 胆囊大小
	Size *KeyValueItem `json:"Size,omitnil,omitempty" name:"Size"`

	// 胆囊触诊
	Palpation *KeyValueItem `json:"Palpation,omitnil,omitempty" name:"Palpation"`

	// 胆囊叩诊
	Percussion *KeyValueItem `json:"Percussion,omitnil,omitempty" name:"Percussion"`

	// 胆囊压痛
	Tenderness *KeyValueItem `json:"Tenderness,omitnil,omitempty" name:"Tenderness"`

	// 胆囊质地
	Consistency *KeyValueItem `json:"Consistency,omitnil,omitempty" name:"Consistency"`
}

type InternalMedicineAbdomenKidney struct {
	// 肾脏总体描述
	Src *KeyValueItem `json:"Src,omitnil,omitempty" name:"Src"`

	// 肾脏大小
	Size *KeyValueItem `json:"Size,omitnil,omitempty" name:"Size"`

	// 肾脏触诊
	Palpation *KeyValueItem `json:"Palpation,omitnil,omitempty" name:"Palpation"`

	// 肾脏叩诊
	Percussion *KeyValueItem `json:"Percussion,omitnil,omitempty" name:"Percussion"`

	// 肾脏压痛
	Tenderness *KeyValueItem `json:"Tenderness,omitnil,omitempty" name:"Tenderness"`

	// 肾脏质地
	Consistency *KeyValueItem `json:"Consistency,omitnil,omitempty" name:"Consistency"`
}

type InternalMedicineAbdomenLiver struct {
	// 肝脏总体描述
	Src *KeyValueItem `json:"Src,omitnil,omitempty" name:"Src"`

	// 肝脏大小
	Size *KeyValueItem `json:"Size,omitnil,omitempty" name:"Size"`

	// 肝脏触诊
	Palpation *KeyValueItem `json:"Palpation,omitnil,omitempty" name:"Palpation"`

	// 肝脏叩诊
	Percussion *KeyValueItem `json:"Percussion,omitnil,omitempty" name:"Percussion"`

	// 肝脏压痛
	Tenderness *KeyValueItem `json:"Tenderness,omitnil,omitempty" name:"Tenderness"`

	// 肝脏质地
	Consistency *KeyValueItem `json:"Consistency,omitnil,omitempty" name:"Consistency"`
}

type InternalMedicineAbdomenPancreas struct {
	// 胰腺总体描述
	Src *KeyValueItem `json:"Src,omitnil,omitempty" name:"Src"`

	// 胰腺大小
	Size *KeyValueItem `json:"Size,omitnil,omitempty" name:"Size"`

	// 胰腺触诊
	Palpation *KeyValueItem `json:"Palpation,omitnil,omitempty" name:"Palpation"`

	// 胰腺叩诊
	Percussion *KeyValueItem `json:"Percussion,omitnil,omitempty" name:"Percussion"`

	// 肝脏压痛
	Tenderness *KeyValueItem `json:"Tenderness,omitnil,omitempty" name:"Tenderness"`

	// 胰腺质地
	Consistency *KeyValueItem `json:"Consistency,omitnil,omitempty" name:"Consistency"`
}

type InternalMedicineAbdomenSpleen struct {
	// 脾脏总体描述
	Src *KeyValueItem `json:"Src,omitnil,omitempty" name:"Src"`

	// 脾脏大小
	Size *KeyValueItem `json:"Size,omitnil,omitempty" name:"Size"`

	// 脾脏触诊
	Palpation *KeyValueItem `json:"Palpation,omitnil,omitempty" name:"Palpation"`

	// 脾脏叩诊
	Percussion *KeyValueItem `json:"Percussion,omitnil,omitempty" name:"Percussion"`

	// 脾脏压痛
	Tenderness *KeyValueItem `json:"Tenderness,omitnil,omitempty" name:"Tenderness"`

	// 脾脏质地
	Consistency *KeyValueItem `json:"Consistency,omitnil,omitempty" name:"Consistency"`
}

type InternalMedicineBaseItem struct {
	// 体检报告-内科-腹部
	Abdomen *InternalMedicineAbdomen `json:"Abdomen,omitnil,omitempty" name:"Abdomen"`

	// 体检报告-内科-心脏
	Heart *InternalMedicineHeart `json:"Heart,omitnil,omitempty" name:"Heart"`

	// 体检报告-内科-血管
	Vessel *InternalMedicineVessel `json:"Vessel,omitnil,omitempty" name:"Vessel"`

	// 体检报告-内科-呼吸系统
	RespiratorySystem *InternalMedicineRespiratorySystem `json:"RespiratorySystem,omitnil,omitempty" name:"RespiratorySystem"`

	// 体检报告-内科-内科其他
	Others []*KeyValueItem `json:"Others,omitnil,omitempty" name:"Others"`

	// 体检报告-内科-小结
	BriefSummary *InternalMedicineBriefSummary `json:"BriefSummary,omitnil,omitempty" name:"BriefSummary"`
}

type InternalMedicineBriefSummary struct {
	// 内科小结
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type InternalMedicineHeart struct {
	// 心脏总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 心律
	HeartRhythm *KeyValueItem `json:"HeartRhythm,omitnil,omitempty" name:"HeartRhythm"`

	// 心率
	HeartRate *ValueUnitItem `json:"HeartRate,omitnil,omitempty" name:"HeartRate"`

	// 心脏听诊
	HeartAuscultation *KeyValueItem `json:"HeartAuscultation,omitnil,omitempty" name:"HeartAuscultation"`
}

type InternalMedicineRespiratorySystem struct {
	// 呼吸系统总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 胸廓
	Thoracic *KeyValueItem `json:"Thoracic,omitnil,omitempty" name:"Thoracic"`

	// 痰量
	Sputum *KeyValueItem `json:"Sputum,omitnil,omitempty" name:"Sputum"`

	// 肺部叩诊
	LungPercussion *KeyValueItem `json:"LungPercussion,omitnil,omitempty" name:"LungPercussion"`

	// 肺部听诊其他
	LungAuscultation []*KeyValueItem `json:"LungAuscultation,omitnil,omitempty" name:"LungAuscultation"`
}

type InternalMedicineVessel struct {
	// 血管总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 血管杂音
	VascularMurmur *KeyValueItem `json:"VascularMurmur,omitnil,omitempty" name:"VascularMurmur"`

	// 外周血管
	PeripheralVessel *KeyValueItem `json:"PeripheralVessel,omitnil,omitempty" name:"PeripheralVessel"`
}

type Invas struct {
	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 部位
	Part *Part `json:"Part,omitnil,omitempty" name:"Part"`

	// 阳性
	Positive *string `json:"Positive,omitnil,omitempty" name:"Positive"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`
}

type InvasiveV2 struct {
	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 部位
	Part *Part `json:"Part,omitnil,omitempty" name:"Part"`

	// 阴性或阳性
	Positive *string `json:"Positive,omitnil,omitempty" name:"Positive"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 原文对应坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type IssueInfo struct {
	// 编号
	CertNumber *string `json:"CertNumber,omitnil,omitempty" name:"CertNumber"`

	// 签发机构
	IssuedAuthority *string `json:"IssuedAuthority,omitnil,omitempty" name:"IssuedAuthority"`

	// 签发日期
	IssuedDate *string `json:"IssuedDate,omitnil,omitempty" name:"IssuedDate"`
}

type KeyValueItem struct {
	// 类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目原文
	Item *PhysicalBaseItem `json:"Item,omitnil,omitempty" name:"Item"`

	// 结果
	Result *PhysicalBaseItem `json:"Result,omitnil,omitempty" name:"Result"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type LastMenstrualPeriodBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 时间戳
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Lymph struct {
	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 部位
	Part *Part `json:"Part,omitnil,omitempty" name:"Part"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 转移数
	TransferNum *int64 `json:"TransferNum,omitnil,omitempty" name:"TransferNum"`
}

type LymphNode struct {
	// 项目名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 部位
	Part *Part `json:"Part,omitnil,omitempty" name:"Part"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 总数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 转移数量
	TransferNum *int64 `json:"TransferNum,omitnil,omitempty" name:"TransferNum"`

	// 淋巴结大小
	Sizes []*int64 `json:"Sizes,omitnil,omitempty" name:"Sizes"`

	// 原文对应坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type LymphTotal struct {
	// 项目名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 转移数量
	TransferNum *int64 `json:"TransferNum,omitnil,omitempty" name:"TransferNum"`

	// 总数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文对应坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type MainDiseaseHistoryBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 状态
	State *bool `json:"State,omitnil,omitempty" name:"State"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 否定列表
	Neglist *NeglistBlock `json:"Neglist,omitnil,omitempty" name:"Neglist"`

	// 肯定列表
	Poslist *PoslistBlock `json:"Poslist,omitnil,omitempty" name:"Poslist"`
}

type Maternity struct {
	// 描述部分
	Desc *MaternityDesc `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 结论部分
	Summary *MaternitySummary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 报告原文
	OcrText *string `json:"OcrText,omitnil,omitempty" name:"OcrText"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type MaternityDesc struct {
	// 胎儿数据结构
	Fetus []*Fetus `json:"Fetus,omitnil,omitempty" name:"Fetus"`

	// 胎儿数量
	FetusNum *FieldInfo `json:"FetusNum,omitnil,omitempty" name:"FetusNum"`

	// 原文
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type MaternitySummary struct {
	// 胎儿数据结构
	Fetus []*Fetus `json:"Fetus,omitnil,omitempty" name:"Fetus"`

	// 胎儿数量
	FetusNum *FieldInfo `json:"FetusNum,omitnil,omitempty" name:"FetusNum"`

	// 病变
	Sym []*FieldInfo `json:"Sym,omitnil,omitempty" name:"Sym"`

	// 原文
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type MedDoc struct {
	// 建议
	Advice *Advice `json:"Advice,omitnil,omitempty" name:"Advice"`

	// 诊断结果
	Diagnosis []*DiagCertItem `json:"Diagnosis,omitnil,omitempty" name:"Diagnosis"`

	// 疾病史
	DiseaseMedicalHistory *DiseaseMedicalHistory `json:"DiseaseMedicalHistory,omitnil,omitempty" name:"DiseaseMedicalHistory"`

	// 个人史
	PersonalMedicalHistory *PersonalMedicalHistory `json:"PersonalMedicalHistory,omitnil,omitempty" name:"PersonalMedicalHistory"`

	// 婚孕史
	ObstericalMedicalHistory *ObstericalMedicalHistory `json:"ObstericalMedicalHistory,omitnil,omitempty" name:"ObstericalMedicalHistory"`

	// 家族史
	FamilyMedicalHistory *FamilyMedicalHistory `json:"FamilyMedicalHistory,omitnil,omitempty" name:"FamilyMedicalHistory"`

	// 月经史
	MenstrualMedicalHistory *MenstrualMedicalHistory `json:"MenstrualMedicalHistory,omitnil,omitempty" name:"MenstrualMedicalHistory"`

	// 诊疗记录
	TreatmentRecord *TreatmentRecord `json:"TreatmentRecord,omitnil,omitempty" name:"TreatmentRecord"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type MedicalRecordInfo struct {
	// 就诊日期
	DiagnosisTime *string `json:"DiagnosisTime,omitnil,omitempty" name:"DiagnosisTime"`

	// 就诊科室
	DiagnosisDepartmentName *string `json:"DiagnosisDepartmentName,omitnil,omitempty" name:"DiagnosisDepartmentName"`

	// 就诊医生
	DiagnosisDoctorName *string `json:"DiagnosisDoctorName,omitnil,omitempty" name:"DiagnosisDoctorName"`

	// 临床诊断
	ClinicalDiagnosis *string `json:"ClinicalDiagnosis,omitnil,omitempty" name:"ClinicalDiagnosis"`

	// 主述
	MainNarration *string `json:"MainNarration,omitnil,omitempty" name:"MainNarration"`

	// 体格检查
	PhysicalExamination *string `json:"PhysicalExamination,omitnil,omitempty" name:"PhysicalExamination"`

	// 检查结论
	InspectionFindings *string `json:"InspectionFindings,omitnil,omitempty" name:"InspectionFindings"`

	// 治疗意见
	TreatmentOpinion *string `json:"TreatmentOpinion,omitnil,omitempty" name:"TreatmentOpinion"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type Medicine struct {
	// 药品名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品名
	TradeName *string `json:"TradeName,omitnil,omitempty" name:"TradeName"`

	// 厂商
	Firm *string `json:"Firm,omitnil,omitempty" name:"Firm"`

	// 医保类型
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 规格
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 最小包装数量
	MinQuantity *string `json:"MinQuantity,omitnil,omitempty" name:"MinQuantity"`

	// 最小制剂单位
	DosageUnit *string `json:"DosageUnit,omitnil,omitempty" name:"DosageUnit"`

	// 最小包装单位
	PackingUnit *string `json:"PackingUnit,omitnil,omitempty" name:"PackingUnit"`
}

type MenstrualFlowBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type MenstrualHistoryBlock struct {
	// 末次月经
	LastMenstrualPeriod *LastMenstrualPeriodBlock `json:"LastMenstrualPeriod,omitnil,omitempty" name:"LastMenstrualPeriod"`

	// 月经量
	MenstrualFlow *MenstrualFlowBlock `json:"MenstrualFlow,omitnil,omitempty" name:"MenstrualFlow"`

	// 初潮时间
	MenarcheAge *LastMenstrualPeriodBlock `json:"MenarcheAge,omitnil,omitempty" name:"MenarcheAge"`

	// 是否绝经
	MenstruationOrNot *MenstruationOrNotBlock `json:"MenstruationOrNot,omitnil,omitempty" name:"MenstruationOrNot"`

	// 月经周期
	MenstrualCycles *LastMenstrualPeriodBlock `json:"MenstrualCycles,omitnil,omitempty" name:"MenstrualCycles"`

	// 月经经期
	MenstrualPeriod *MenstrualPeriodBlock `json:"MenstrualPeriod,omitnil,omitempty" name:"MenstrualPeriod"`
}

type MenstrualHistoryDetailBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 时间类型
	TimeType *string `json:"TimeType,omitnil,omitempty" name:"TimeType"`

	// 时间戳
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type MenstrualMedicalHistory struct {
	// 末次月经时间
	LastMenstrualPeriod *string `json:"LastMenstrualPeriod,omitnil,omitempty" name:"LastMenstrualPeriod"`

	// 经量
	MenstrualFlow *string `json:"MenstrualFlow,omitnil,omitempty" name:"MenstrualFlow"`

	// 月经初潮年龄
	MenarcheAge *string `json:"MenarcheAge,omitnil,omitempty" name:"MenarcheAge"`

	// 是否来月经
	MenstruationOrNot *string `json:"MenstruationOrNot,omitnil,omitempty" name:"MenstruationOrNot"`

	// 月经周期
	MenstrualCycles *string `json:"MenstrualCycles,omitnil,omitempty" name:"MenstrualCycles"`

	// 月经持续天数
	MenstrualPeriod *string `json:"MenstrualPeriod,omitnil,omitempty" name:"MenstrualPeriod"`
}

type MenstrualPeriodBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 时间戳
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type MenstruationOrNotBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 时间类型
	TimeType *string `json:"TimeType,omitnil,omitempty" name:"TimeType"`

	// 时间戳
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Molecular struct {
	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 基因名称标注化
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分子病理详细信息
	Value *MolecularValue `json:"Value,omitnil,omitempty" name:"Value"`

	// 原文对应坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type MolecularValue struct {
	// 外显子
	Exon *string `json:"Exon,omitnil,omitempty" name:"Exon"`

	// 点位
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 阳性或阴性
	Positive *string `json:"Positive,omitnil,omitempty" name:"Positive"`

	// 基因名称原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`
}

type Multiple struct {
	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type NeglistBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type NeonatalInfo struct {
	// 新生儿名字
	NeonatalName *string `json:"NeonatalName,omitnil,omitempty" name:"NeonatalName"`

	// 新生儿性别
	NeonatalGender *string `json:"NeonatalGender,omitnil,omitempty" name:"NeonatalGender"`

	// 出生身长
	BirthLength *string `json:"BirthLength,omitnil,omitempty" name:"BirthLength"`

	// 出生体重
	BirthWeight *string `json:"BirthWeight,omitnil,omitempty" name:"BirthWeight"`

	// 出生孕周
	GestationalAge *string `json:"GestationalAge,omitnil,omitempty" name:"GestationalAge"`

	// 出生时间
	BirthTime *string `json:"BirthTime,omitnil,omitempty" name:"BirthTime"`

	// 出生地点
	BirthPlace *string `json:"BirthPlace,omitnil,omitempty" name:"BirthPlace"`

	// 医疗机构
	MedicalInstitutions *string `json:"MedicalInstitutions,omitnil,omitempty" name:"MedicalInstitutions"`
}

type NormPart struct {
	// 部位值
	Part *string `json:"Part,omitnil,omitempty" name:"Part"`

	// 部位方向
	PartDirection *string `json:"PartDirection,omitnil,omitempty" name:"PartDirection"`

	// 组织值
	Tissue *string `json:"Tissue,omitnil,omitempty" name:"Tissue"`

	// 组织方向
	TissueDirection *string `json:"TissueDirection,omitnil,omitempty" name:"TissueDirection"`

	// 上级部位
	Upper *string `json:"Upper,omitnil,omitempty" name:"Upper"`

	// 部位详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: PartDetail is deprecated.
	PartDetail *PartDesc `json:"PartDetail,omitnil,omitempty" name:"PartDetail"`

	// 部位详情
	PartDetailList []*PartDesc `json:"PartDetailList,omitnil,omitempty" name:"PartDetailList"`
}

type NormSize struct {
	// 数量
	Number []*string `json:"Number,omitnil,omitempty" name:"Number"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 归一化值
	Impl *string `json:"Impl,omitnil,omitempty" name:"Impl"`
}

type NumValue struct {
	// 数值
	Num *string `json:"Num,omitnil,omitempty" name:"Num"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

type ObstericalMedicalHistory struct {
	// 婚史
	MarriageHistory *string `json:"MarriageHistory,omitnil,omitempty" name:"MarriageHistory"`

	// 孕史
	FertilityHistory *string `json:"FertilityHistory,omitnil,omitempty" name:"FertilityHistory"`
}

type ObstetricalHistoryBlock struct {
	// 婚姻史
	MarriageHistory *MenstrualHistoryDetailBlock `json:"MarriageHistory,omitnil,omitempty" name:"MarriageHistory"`

	// 婚育史
	FertilityHistory *FertilityHistoryBlock `json:"FertilityHistory,omitnil,omitempty" name:"FertilityHistory"`
}

type OcrInfo struct {
	// 图片进行OCR之后得到的所有包含字块的OCR信息
	Items []*OcrItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 图片进行OCR之后得到的所有字符
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type OcrItem struct {
	// 图片中文字的字符串
	Words *string `json:"Words,omitnil,omitempty" name:"Words"`

	// Words 中每个文字的坐标数组，顺序与Words中的字符顺序一致
	Coords []*Coordinate `json:"Coords,omitnil,omitempty" name:"Coords"`

	// 整个字符块的坐标信息
	WordCoords *Coordinate `json:"WordCoords,omitnil,omitempty" name:"WordCoords"`
}

type OphthalmologyBareEyeSight struct {
	// 左眼视力
	LeftEyeVisual *KeyValueItem `json:"LeftEyeVisual,omitnil,omitempty" name:"LeftEyeVisual"`

	// 裸眼视力
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 右眼视力
	RightEyeVisual *KeyValueItem `json:"RightEyeVisual,omitnil,omitempty" name:"RightEyeVisual"`
}

type OphthalmologyBaseItem struct {
	// 裸眼视力
	BareEyeSight *OphthalmologyBareEyeSight `json:"BareEyeSight,omitnil,omitempty" name:"BareEyeSight"`

	// 矫正视力
	CorrectedVisualAcuity *OphthalmologyCorrectedVisualAcuity `json:"CorrectedVisualAcuity,omitnil,omitempty" name:"CorrectedVisualAcuity"`

	// 色觉
	ColourVision *OphthalmologyColourVision `json:"ColourVision,omitnil,omitempty" name:"ColourVision"`

	// 眼底
	Fundoscopy *OphthalmologyFundoscopy `json:"Fundoscopy,omitnil,omitempty" name:"Fundoscopy"`

	// 眼科其他
	Others []*KeyValueItem `json:"Others,omitnil,omitempty" name:"Others"`

	// 眼科小结
	BriefSummary *OphthalmologyBriefSummary `json:"BriefSummary,omitnil,omitempty" name:"BriefSummary"`
}

type OphthalmologyBriefSummary struct {
	// 眼科小结
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type OphthalmologyColourVision struct {
	// 色觉总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type OphthalmologyCorrectedVisualAcuity struct {
	// 左眼矫正视力
	LeftEyeVisual *KeyValueItem `json:"LeftEyeVisual,omitnil,omitempty" name:"LeftEyeVisual"`

	// 矫正视力
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 右眼矫正视力
	RightEyeVisual *KeyValueItem `json:"RightEyeVisual,omitnil,omitempty" name:"RightEyeVisual"`
}

type OphthalmologyFundoscopy struct {
	// 眼底检查总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type Organ struct {
	// 部位
	Part *Part `json:"Part,omitnil,omitempty" name:"Part"`

	// 大小
	Size []*Size `json:"Size,omitnil,omitempty" name:"Size"`

	// 包膜
	Envelope *BlockInfo `json:"Envelope,omitnil,omitempty" name:"Envelope"`

	// 边缘
	Edge *BlockInfo `json:"Edge,omitnil,omitempty" name:"Edge"`

	// 内部回声
	InnerEcho *BlockInfo `json:"InnerEcho,omitnil,omitempty" name:"InnerEcho"`

	// 腺体
	Gland *BlockInfo `json:"Gland,omitnil,omitempty" name:"Gland"`

	// 形状
	Shape *BlockInfo `json:"Shape,omitnil,omitempty" name:"Shape"`

	// 厚度
	Thickness *BlockInfo `json:"Thickness,omitnil,omitempty" name:"Thickness"`

	// 形态
	ShapeAttr *BlockInfo `json:"ShapeAttr,omitnil,omitempty" name:"ShapeAttr"`

	// 血液cdfi
	CDFI *BlockInfo `json:"CDFI,omitnil,omitempty" name:"CDFI"`

	// 描述信息
	SymDesc *BlockInfo `json:"SymDesc,omitnil,omitempty" name:"SymDesc"`

	// 大小状态
	SizeStatus *BlockInfo `json:"SizeStatus,omitnil,omitempty" name:"SizeStatus"`

	// 轮廓
	Outline *BlockInfo `json:"Outline,omitnil,omitempty" name:"Outline"`

	// 结构
	Structure *BlockInfo `json:"Structure,omitnil,omitempty" name:"Structure"`

	// 密度
	Density *BlockInfo `json:"Density,omitnil,omitempty" name:"Density"`

	// 血管
	Vas *BlockInfo `json:"Vas,omitnil,omitempty" name:"Vas"`

	// 囊壁
	Cysticwall *BlockInfo `json:"Cysticwall,omitnil,omitempty" name:"Cysticwall"`

	// 被膜
	Capsule *BlockInfo `json:"Capsule,omitnil,omitempty" name:"Capsule"`

	// 峡部厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IsthmusThicknese is deprecated.
	IsthmusThicknese *Size `json:"IsthmusThicknese,omitnil,omitempty" name:"IsthmusThicknese"`

	// 内部回声分布
	InnerEchoDistribution *BlockInfo `json:"InnerEchoDistribution,omitnil,omitempty" name:"InnerEchoDistribution"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 透声度
	Transparent *BlockInfo `json:"Transparent,omitnil,omitempty" name:"Transparent"`

	// MRI ADC
	MriAdc *BlockInfo `json:"MriAdc,omitnil,omitempty" name:"MriAdc"`

	// MRI DWI
	MriDwi *BlockInfo `json:"MriDwi,omitnil,omitempty" name:"MriDwi"`

	// MRI T1信号
	MriT1 *BlockInfo `json:"MriT1,omitnil,omitempty" name:"MriT1"`

	// MRI T2信号
	MriT2 *BlockInfo `json:"MriT2,omitnil,omitempty" name:"MriT2"`

	// CT HU值
	CtHu *BlockInfo `json:"CtHu,omitnil,omitempty" name:"CtHu"`

	// SUmax值
	Suvmax *BlockInfo `json:"Suvmax,omitnil,omitempty" name:"Suvmax"`

	// 代谢情况
	Metabolism *BlockInfo `json:"Metabolism,omitnil,omitempty" name:"Metabolism"`

	// 放射性摄取
	RadioactiveUptake *BlockInfo `json:"RadioactiveUptake,omitnil,omitempty" name:"RadioactiveUptake"`

	// 淋巴结情况
	LymphEnlargement *BlockInfo `json:"LymphEnlargement,omitnil,omitempty" name:"LymphEnlargement"`

	// 影像特征
	ImageFeature *BlockInfo `json:"ImageFeature,omitnil,omitempty" name:"ImageFeature"`

	// 导管
	Duct *BlockInfo `json:"Duct,omitnil,omitempty" name:"Duct"`

	// 趋势
	Trend *BlockInfo `json:"Trend,omitnil,omitempty" name:"Trend"`

	// 手术情况
	Operation *BlockInfo `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 器官在报告图片中的坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`

	// 峡部厚度
	IsthmusThickness *Size `json:"IsthmusThickness,omitnil,omitempty" name:"IsthmusThickness"`
}

type OtherInfo struct {
	// 麻醉方法
	Anesthesia *SurgeryAttr `json:"Anesthesia,omitnil,omitempty" name:"Anesthesia"`

	// 术中出血
	BloodLoss *SurgeryAttr `json:"BloodLoss,omitnil,omitempty" name:"BloodLoss"`

	// 输血
	BloodTransfusion *SurgeryAttr `json:"BloodTransfusion,omitnil,omitempty" name:"BloodTransfusion"`

	// 手术用时
	Duration *SurgeryAttr `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 手术开始时间
	EndTime *SurgeryAttr `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 手术结束时间
	StartTime *SurgeryAttr `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type OtolaryngologyBaseItem struct {
	// 耳朵
	Ear *OtolaryngologyEar `json:"Ear,omitnil,omitempty" name:"Ear"`

	// 鼻
	Nose *OtolaryngologyNose `json:"Nose,omitnil,omitempty" name:"Nose"`

	// 喉
	Larynx *OtolaryngologyLarynx `json:"Larynx,omitnil,omitempty" name:"Larynx"`

	// 耳鼻喉其他
	Others []*KeyValueItem `json:"Others,omitnil,omitempty" name:"Others"`

	// 小结
	BriefSummary *OtolaryngologyBriefSummary `json:"BriefSummary,omitnil,omitempty" name:"BriefSummary"`
}

type OtolaryngologyBriefSummary struct {
	// 耳鼻喉小结
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type OtolaryngologyEar struct {
	// 耳总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 听力
	Hearing *HearingItem `json:"Hearing,omitnil,omitempty" name:"Hearing"`
}

type OtolaryngologyLarynx struct {
	// 喉总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type OtolaryngologyNose struct {
	// 鼻总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type PTNM struct {
	// 项目名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// pT
	PT *string `json:"PT,omitnil,omitempty" name:"PT"`

	// pN
	PN *string `json:"PN,omitnil,omitempty" name:"PN"`

	// pM
	PM *string `json:"PM,omitnil,omitempty" name:"PM"`

	// 原文对应坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type PTNMBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// PTNM分期
	PTNMM *string `json:"PTNMM,omitnil,omitempty" name:"PTNMM"`

	// PTNM分期
	PTNMN *string `json:"PTNMN,omitnil,omitempty" name:"PTNMN"`

	// PTNM分期
	PTNMT *string `json:"PTNMT,omitnil,omitempty" name:"PTNMT"`
}

type ParagraphBlock struct {
	// 切口愈合情况
	IncisionHealingText *string `json:"IncisionHealingText,omitnil,omitempty" name:"IncisionHealingText"`

	// 辅助检查
	AuxiliaryExaminationText *string `json:"AuxiliaryExaminationText,omitnil,omitempty" name:"AuxiliaryExaminationText"`

	// 特殊检查
	SpecialExamText *string `json:"SpecialExamText,omitnil,omitempty" name:"SpecialExamText"`

	// 门诊诊断
	OutpatientDiagnosisText *string `json:"OutpatientDiagnosisText,omitnil,omitempty" name:"OutpatientDiagnosisText"`

	// 入院情况
	AdmissionConditionText *string `json:"AdmissionConditionText,omitnil,omitempty" name:"AdmissionConditionText"`

	// 诊疗经过
	CheckAndTreatmentProcessText *string `json:"CheckAndTreatmentProcessText,omitnil,omitempty" name:"CheckAndTreatmentProcessText"`

	// 体征
	SymptomsAndSignsText *string `json:"SymptomsAndSignsText,omitnil,omitempty" name:"SymptomsAndSignsText"`

	// 出院医嘱
	DischargeInstructionsText *string `json:"DischargeInstructionsText,omitnil,omitempty" name:"DischargeInstructionsText"`

	// 入院诊断
	AdmissionDiagnosisText *string `json:"AdmissionDiagnosisText,omitnil,omitempty" name:"AdmissionDiagnosisText"`

	// 手术情况
	SurgeryConditionText *string `json:"SurgeryConditionText,omitnil,omitempty" name:"SurgeryConditionText"`

	// 病理诊断
	PathologicalDiagnosisText *string `json:"PathologicalDiagnosisText,omitnil,omitempty" name:"PathologicalDiagnosisText"`

	// 出院情况
	DischargeConditionText *string `json:"DischargeConditionText,omitnil,omitempty" name:"DischargeConditionText"`

	// 检查记录
	CheckRecordText *string `json:"CheckRecordText,omitnil,omitempty" name:"CheckRecordText"`

	// 主诉
	ChiefComplaintText *string `json:"ChiefComplaintText,omitnil,omitempty" name:"ChiefComplaintText"`

	// 出院诊断
	DischargeDiagnosisText *string `json:"DischargeDiagnosisText,omitnil,omitempty" name:"DischargeDiagnosisText"`

	// 既往史
	MainDiseaseHistoryText *string `json:"MainDiseaseHistoryText,omitnil,omitempty" name:"MainDiseaseHistoryText"`

	// 现病史
	DiseasePresentText *string `json:"DiseasePresentText,omitnil,omitempty" name:"DiseasePresentText"`

	// 个人史
	PersonalHistoryText *string `json:"PersonalHistoryText,omitnil,omitempty" name:"PersonalHistoryText"`

	// 月经史
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: MenstruallHistoryText is deprecated.
	MenstruallHistoryText *string `json:"MenstruallHistoryText,omitnil,omitempty" name:"MenstruallHistoryText"`

	// 婚育史
	ObstericalHistoryText *string `json:"ObstericalHistoryText,omitnil,omitempty" name:"ObstericalHistoryText"`

	// 家族史
	FamilyHistoryText *string `json:"FamilyHistoryText,omitnil,omitempty" name:"FamilyHistoryText"`

	// 过敏史
	AllergyHistoryText *string `json:"AllergyHistoryText,omitnil,omitempty" name:"AllergyHistoryText"`

	// 病史信息
	DiseaseHistoryText *string `json:"DiseaseHistoryText,omitnil,omitempty" name:"DiseaseHistoryText"`

	// 其它诊断
	OtherDiagnosisText *string `json:"OtherDiagnosisText,omitnil,omitempty" name:"OtherDiagnosisText"`

	// 体格检查
	BodyExaminationText *string `json:"BodyExaminationText,omitnil,omitempty" name:"BodyExaminationText"`

	// 专科检查
	SpecialistExaminationText *string `json:"SpecialistExaminationText,omitnil,omitempty" name:"SpecialistExaminationText"`

	// 治疗结果
	TreatmentResultText *string `json:"TreatmentResultText,omitnil,omitempty" name:"TreatmentResultText"`

	// 月经史
	MenstrualHistoryText *string `json:"MenstrualHistoryText,omitnil,omitempty" name:"MenstrualHistoryText"`
}

type ParentInfo struct {
	// 名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 年龄
	Age *string `json:"Age,omitnil,omitempty" name:"Age"`

	// 证件号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 民族
	Ethnicity *string `json:"Ethnicity,omitnil,omitempty" name:"Ethnicity"`

	// 国籍
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

type Part struct {
	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 部位
	NormPart *NormPart `json:"NormPart,omitnil,omitempty" name:"NormPart"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	ValueBrief *string `json:"ValueBrief,omitnil,omitempty" name:"ValueBrief"`
}

type PartDesc struct {
	// 主要部位
	MainDir *string `json:"MainDir,omitnil,omitempty" name:"MainDir"`

	// 部位
	Part *string `json:"Part,omitnil,omitempty" name:"Part"`

	// 次要部位
	SecondaryDir *string `json:"SecondaryDir,omitnil,omitempty" name:"SecondaryDir"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type PathologicalDiagnosisBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 病理详细
	Detail []*PathologicalDiagnosisDetailBlock `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PathologicalDiagnosisDetailBlock struct {
	// 部位
	Part *string `json:"Part,omitnil,omitempty" name:"Part"`

	// 类型
	HistologicalType *string `json:"HistologicalType,omitnil,omitempty" name:"HistologicalType"`

	// 等级
	HistologicalGrade *string `json:"HistologicalGrade,omitnil,omitempty" name:"HistologicalGrade"`
}

type PathologyReport struct {
	// 癌症部位
	CancerPart *Part `json:"CancerPart,omitnil,omitempty" name:"CancerPart"`

	// 癌症部位大小
	CancerSize []*Size `json:"CancerSize,omitnil,omitempty" name:"CancerSize"`

	// 描述文本
	DescText *string `json:"DescText,omitnil,omitempty" name:"DescText"`

	// 组织学等级
	HistologyLevel *HistologyLevel `json:"HistologyLevel,omitnil,omitempty" name:"HistologyLevel"`

	// 组织学类型
	HistologyType *HistologyType `json:"HistologyType,omitnil,omitempty" name:"HistologyType"`

	// IHC信息
	IHC []*IHCInfo `json:"IHC,omitnil,omitempty" name:"IHC"`

	// 浸润深度
	InfiltrationDepth *BlockInfo `json:"InfiltrationDepth,omitnil,omitempty" name:"InfiltrationDepth"`

	// 肿瘤扩散
	Invasive []*Invas `json:"Invasive,omitnil,omitempty" name:"Invasive"`

	// 淋巴结
	LymphNodes []*Lymph `json:"LymphNodes,omitnil,omitempty" name:"LymphNodes"`

	// PTNM信息
	PTNM *BlockInfo `json:"PTNM,omitnil,omitempty" name:"PTNM"`

	// 病理报告类型
	PathologicalReportType *BlockInfo `json:"PathologicalReportType,omitnil,omitempty" name:"PathologicalReportType"`

	// 报告原文
	ReportText *string `json:"ReportText,omitnil,omitempty" name:"ReportText"`

	// 标本类型
	SampleType *BlockInfo `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 结论文本
	SummaryText *string `json:"SummaryText,omitnil,omitempty" name:"SummaryText"`
}

type PathologyV2 struct {
	// 报告类型
	PathologicalReportType *Report `json:"PathologicalReportType,omitnil,omitempty" name:"PathologicalReportType"`

	// 描述段落
	Desc *DescInfo `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 诊断结论
	Summary *SummaryInfo `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 报告全文
	ReportText *string `json:"ReportText,omitnil,omitempty" name:"ReportText"`

	// 淋巴结总计转移信息
	LymphTotal []*LymphTotal `json:"LymphTotal,omitnil,omitempty" name:"LymphTotal"`

	// 单淋巴结转移信息
	LymphNodes []*LymphNode `json:"LymphNodes,omitnil,omitempty" name:"LymphNodes"`

	// ihc信息
	Ihc []*IHCV2 `json:"Ihc,omitnil,omitempty" name:"Ihc"`

	// 临床诊断
	Clinical *BaseInfo `json:"Clinical,omitnil,omitempty" name:"Clinical"`

	// 是否癌前病变
	Precancer *HistologyClass `json:"Precancer,omitnil,omitempty" name:"Precancer"`

	// 是否恶性肿瘤
	Malignant *HistologyClass `json:"Malignant,omitnil,omitempty" name:"Malignant"`

	// 是否良性肿瘤
	Benigntumor *HistologyClass `json:"Benigntumor,omitnil,omitempty" name:"Benigntumor"`

	// 送检材料
	SampleType *BaseInfo `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 淋巴结大小
	LymphSize []*Size `json:"LymphSize,omitnil,omitempty" name:"LymphSize"`

	// 分子病理
	Molecular []*Molecular `json:"Molecular,omitnil,omitempty" name:"Molecular"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type PatientInfo struct {
	// 患者姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 患者性别
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// 患者年龄
	Age *string `json:"Age,omitnil,omitempty" name:"Age"`

	// 患者手机号码
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 患者地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 患者身份证
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 健康卡号
	HealthCardNo *string `json:"HealthCardNo,omitnil,omitempty" name:"HealthCardNo"`

	// 社保卡号
	SocialSecurityCardNo *string `json:"SocialSecurityCardNo,omitnil,omitempty" name:"SocialSecurityCardNo"`

	// 出生日期
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// 民族
	Ethnicity *string `json:"Ethnicity,omitnil,omitempty" name:"Ethnicity"`

	// 婚姻状况
	Married *string `json:"Married,omitnil,omitempty" name:"Married"`

	// 职业
	Profession *string `json:"Profession,omitnil,omitempty" name:"Profession"`

	// 教育程度
	EducationBackground *string `json:"EducationBackground,omitnil,omitempty" name:"EducationBackground"`

	// 国籍
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// 籍贯
	BirthPlace *string `json:"BirthPlace,omitnil,omitempty" name:"BirthPlace"`

	// 医保类型
	MedicalInsuranceType *string `json:"MedicalInsuranceType,omitnil,omitempty" name:"MedicalInsuranceType"`

	// 标准化年龄
	AgeNorm *string `json:"AgeNorm,omitnil,omitempty" name:"AgeNorm"`

	// 民族。该字段已不再使用，请从Ethnicity取值
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Nation is deprecated.
	Nation *string `json:"Nation,omitnil,omitempty" name:"Nation"`

	// 婚姻代码
	MarriedCode *string `json:"MarriedCode,omitnil,omitempty" name:"MarriedCode"`

	// 职业代码
	ProfessionCode *string `json:"ProfessionCode,omitnil,omitempty" name:"ProfessionCode"`

	// 居民医保代码
	MedicalInsuranceTypeCode *string `json:"MedicalInsuranceTypeCode,omitnil,omitempty" name:"MedicalInsuranceTypeCode"`

	// 床号
	BedNo *string `json:"BedNo,omitnil,omitempty" name:"BedNo"`
}

type PdfInfo struct {
	// pdf文件url链接(暂不支持)
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// pdf文件base64编码字符串
	Base64 *string `json:"Base64,omitnil,omitempty" name:"Base64"`
}

type PersonalHistoryBlock struct {
	// 出生地
	BirthPlace *BirthPlaceBlock `json:"BirthPlace,omitnil,omitempty" name:"BirthPlace"`

	// 居住地
	LivePlace *BirthPlaceBlock `json:"LivePlace,omitnil,omitempty" name:"LivePlace"`

	// 职业
	Job *BirthPlaceBlock `json:"Job,omitnil,omitempty" name:"Job"`

	// 吸烟
	SmokeHistory *SmokeHistoryBlock `json:"SmokeHistory,omitnil,omitempty" name:"SmokeHistory"`

	// 喝酒
	AlcoholicHistory *SmokeHistoryBlock `json:"AlcoholicHistory,omitnil,omitempty" name:"AlcoholicHistory"`

	// 月经史
	MenstrualHistory *MenstrualHistoryBlock `json:"MenstrualHistory,omitnil,omitempty" name:"MenstrualHistory"`

	// 婚姻-生育史
	ObstericalHistory *ObstetricalHistoryBlock `json:"ObstericalHistory,omitnil,omitempty" name:"ObstericalHistory"`

	// 家族史
	FamilyHistory *FamilyHistoryBlock `json:"FamilyHistory,omitnil,omitempty" name:"FamilyHistory"`
}

type PersonalMedicalHistory struct {
	// 出生史
	BirthPlace *string `json:"BirthPlace,omitnil,omitempty" name:"BirthPlace"`

	// 居住史
	LivePlace *string `json:"LivePlace,omitnil,omitempty" name:"LivePlace"`

	// 工作史
	Job *string `json:"Job,omitnil,omitempty" name:"Job"`

	// 吸烟史
	SmokeHistory *string `json:"SmokeHistory,omitnil,omitempty" name:"SmokeHistory"`

	// 饮酒史
	AlcoholicHistory *string `json:"AlcoholicHistory,omitnil,omitempty" name:"AlcoholicHistory"`
}

type PhysicalBaseItem struct {
	// 类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原始文本
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 归一化后值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 四点坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type PhysicalExamination struct {
	// 一般检查
	GeneralExamination *GeneralExaminationBaseItem `json:"GeneralExamination,omitnil,omitempty" name:"GeneralExamination"`

	// 内科
	InternalMedicine *InternalMedicineBaseItem `json:"InternalMedicine,omitnil,omitempty" name:"InternalMedicine"`

	// 外科
	Surgery *SurgeryBaseItem `json:"Surgery,omitnil,omitempty" name:"Surgery"`

	// 口腔科
	Stomatology *StomatologyBaseItem `json:"Stomatology,omitnil,omitempty" name:"Stomatology"`

	// 眼科
	Ophthalmology *OphthalmologyBaseItem `json:"Ophthalmology,omitnil,omitempty" name:"Ophthalmology"`

	// 耳鼻喉科
	Otolaryngology *OtolaryngologyBaseItem `json:"Otolaryngology,omitnil,omitempty" name:"Otolaryngology"`

	// 妇科
	Gynaecology *GynaecologyBaseItem `json:"Gynaecology,omitnil,omitempty" name:"Gynaecology"`

	// 未标准化
	Unclassified []*KeyValueItem `json:"Unclassified,omitnil,omitempty" name:"Unclassified"`
}

type PhysicalExaminationV1 struct {
	// 体检报告信息
	PhysicalExaminationMulti *PhysicalExamination `json:"PhysicalExaminationMulti,omitnil,omitempty" name:"PhysicalExaminationMulti"`

	// 版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

type Point struct {
	// x坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// y坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`
}

type PoslistBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Prescription struct {
	// 药品列表
	MedicineList []*Medicine `json:"MedicineList,omitnil,omitempty" name:"MedicineList"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type Rectangle struct {
	// x坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// y坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 宽
	W *int64 `json:"W,omitnil,omitempty" name:"W"`

	// 高
	H *int64 `json:"H,omitnil,omitempty" name:"H"`
}

type RelapseDateBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 疾病名称
	DiseaseName *string `json:"DiseaseName,omitnil,omitempty" name:"DiseaseName"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 归一化值
	Norm *string `json:"Norm,omitnil,omitempty" name:"Norm"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 时间戳
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RelativeCancerHistoryBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 肿瘤史列表
	RelativeCancerList *string `json:"RelativeCancerList,omitnil,omitempty" name:"RelativeCancerList"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RelativeHistoryBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 成员列表
	Detail []*RelativeHistoryDetailBlock `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RelativeHistoryDetailBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关系
	Relation *string `json:"Relation,omitnil,omitempty" name:"Relation"`

	// 死亡时间
	TimeOfDeath *string `json:"TimeOfDeath,omitnil,omitempty" name:"TimeOfDeath"`

	// 时间类型
	TimeType *string `json:"TimeType,omitnil,omitempty" name:"TimeType"`
}

type Report struct {
	// 项目名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 索引
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 报告类型
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 原文对应坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type ReportInfo struct {
	// 医院名称
	Hospital *string `json:"Hospital,omitnil,omitempty" name:"Hospital"`

	// 科室名称
	DepartmentName *string `json:"DepartmentName,omitnil,omitempty" name:"DepartmentName"`

	// 申请时间
	BillingTime *string `json:"BillingTime,omitnil,omitempty" name:"BillingTime"`

	// 报告时间
	ReportTime *string `json:"ReportTime,omitnil,omitempty" name:"ReportTime"`

	// 检查时间
	InspectTime *string `json:"InspectTime,omitnil,omitempty" name:"InspectTime"`

	// 检查号
	CheckNum *string `json:"CheckNum,omitnil,omitempty" name:"CheckNum"`

	// 影像号
	ImageNum *string `json:"ImageNum,omitnil,omitempty" name:"ImageNum"`

	// 放射号
	RadiationNum *string `json:"RadiationNum,omitnil,omitempty" name:"RadiationNum"`

	// 检验号
	TestNum *string `json:"TestNum,omitnil,omitempty" name:"TestNum"`

	// 门诊号
	OutpatientNum *string `json:"OutpatientNum,omitnil,omitempty" name:"OutpatientNum"`

	// 病理号
	PathologyNum *string `json:"PathologyNum,omitnil,omitempty" name:"PathologyNum"`

	// 住院号
	InHospitalNum *string `json:"InHospitalNum,omitnil,omitempty" name:"InHospitalNum"`

	// 样本号
	SampleNum *string `json:"SampleNum,omitnil,omitempty" name:"SampleNum"`

	// 标本种类
	SampleType *string `json:"SampleType,omitnil,omitempty" name:"SampleType"`

	// 病历号
	MedicalRecordNum *string `json:"MedicalRecordNum,omitnil,omitempty" name:"MedicalRecordNum"`

	// 报告名称
	ReportName *string `json:"ReportName,omitnil,omitempty" name:"ReportName"`

	// 超声号
	UltraNum *string `json:"UltraNum,omitnil,omitempty" name:"UltraNum"`

	// 临床诊断
	Diagnose *string `json:"Diagnose,omitnil,omitempty" name:"Diagnose"`

	// 检查项目
	CheckItem *string `json:"CheckItem,omitnil,omitempty" name:"CheckItem"`

	// 检查方法
	CheckMethod *string `json:"CheckMethod,omitnil,omitempty" name:"CheckMethod"`

	// 诊断时间
	DiagnoseTime *string `json:"DiagnoseTime,omitnil,omitempty" name:"DiagnoseTime"`

	// 体检号
	HealthCheckupNum *string `json:"HealthCheckupNum,omitnil,omitempty" name:"HealthCheckupNum"`

	// 其它时间
	OtherTime *string `json:"OtherTime,omitnil,omitempty" name:"OtherTime"`

	// 打印时间
	PrintTime *string `json:"PrintTime,omitnil,omitempty" name:"PrintTime"`

	// 未归类时间
	Times []*Time `json:"Times,omitnil,omitempty" name:"Times"`

	// 床号
	BedNo *string `json:"BedNo,omitnil,omitempty" name:"BedNo"`
}

type ReportTypeVersion struct {
	// 检验报告
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 版本2
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`
}

type ResultInfo struct {
	// 段落文本
	Text *BaseInfo `json:"Text,omitnil,omitempty" name:"Text"`

	// 结论详情
	Items []*BaseInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type Size struct {
	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 标准大小
	NormSize *NormSize `json:"NormSize,omitnil,omitempty" name:"NormSize"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type SmokeHistoryBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 时间归一化
	TimeNorm *string `json:"TimeNorm,omitnil,omitempty" name:"TimeNorm"`

	// 吸烟量
	Amount *string `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 戒烟状态
	QuitState *bool `json:"QuitState,omitnil,omitempty" name:"QuitState"`

	// 是否吸烟
	State *bool `json:"State,omitnil,omitempty" name:"State"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type StomatologyBaseItem struct {
	// 龋齿
	ToothDecay *StomatologyToothDecay `json:"ToothDecay,omitnil,omitempty" name:"ToothDecay"`

	// 牙龈
	Gingiva *StomatologyGingiva `json:"Gingiva,omitnil,omitempty" name:"Gingiva"`

	// 牙周
	Periodontics *StomatologyPeriodontics `json:"Periodontics,omitnil,omitempty" name:"Periodontics"`

	// 口腔其他
	Others []*KeyValueItem `json:"Others,omitnil,omitempty" name:"Others"`

	// 小结
	BriefSummary *StomatologyBriefSummary `json:"BriefSummary,omitnil,omitempty" name:"BriefSummary"`
}

type StomatologyBriefSummary struct {
	// 口腔小结
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type StomatologyGingiva struct {
	// 牙龈总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type StomatologyPeriodontics struct {
	// 牙周总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type StomatologyToothDecay struct {
	// 龋齿总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type Summary struct {
	// 症状
	Symptom []*SymptomInfo `json:"Symptom,omitnil,omitempty" name:"Symptom"`

	// 文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type SummaryInfo struct {
	// 诊断结论文本
	Text *BaseInfo `json:"Text,omitnil,omitempty" name:"Text"`

	// 诊断结论详情
	Infos []*DetailInformation `json:"Infos,omitnil,omitempty" name:"Infos"`
}

type Surgery struct {
	// 手术史
	SurgeryHistory *SurgeryHistory `json:"SurgeryHistory,omitnil,omitempty" name:"SurgeryHistory"`

	// 其他信息
	OtherInfo *OtherInfo `json:"OtherInfo,omitnil,omitempty" name:"OtherInfo"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type SurgeryAnorectal struct {
	// 肛门直肠总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 直肠指检
	DigitalRectalExamination *KeyValueItem `json:"DigitalRectalExamination,omitnil,omitempty" name:"DigitalRectalExamination"`

	// 痔疮
	Hemorrhoid *KeyValueItem `json:"Hemorrhoid,omitnil,omitempty" name:"Hemorrhoid"`
}

type SurgeryAttr struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type SurgeryBaseItem struct {
	// 体检报告-外科-头颈部
	HeadNeck *SurgeryHeadNeck `json:"HeadNeck,omitnil,omitempty" name:"HeadNeck"`

	// 体检报告-外科-甲状腺
	Thyroid *SurgeryThyroid `json:"Thyroid,omitnil,omitempty" name:"Thyroid"`

	// 体检报告-外科-乳房
	Breast *SurgeryBreast `json:"Breast,omitnil,omitempty" name:"Breast"`

	// 体检报告-外科-浅表淋巴结
	LymphNode *SurgeryLymphNode `json:"LymphNode,omitnil,omitempty" name:"LymphNode"`

	// 体检报告-外科-脊柱
	SpinalExtremities *SurgerySpinalExtremities `json:"SpinalExtremities,omitnil,omitempty" name:"SpinalExtremities"`

	// 体检报告-外科-皮肤
	Skin *SurgerySkin `json:"Skin,omitnil,omitempty" name:"Skin"`

	// 体检报告-外科-肛门直肠
	Anorectal *SurgeryAnorectal `json:"Anorectal,omitnil,omitempty" name:"Anorectal"`

	// 体检报告-外科-泌尿生殖系统
	UrogenitalSystem *SurgeryUrogenitalSystem `json:"UrogenitalSystem,omitnil,omitempty" name:"UrogenitalSystem"`

	// 体检报告-外科-外科其他
	Others []*KeyValueItem `json:"Others,omitnil,omitempty" name:"Others"`

	// 体检报告-外科-小结
	BriefSummary *SurgeryBriefSummary `json:"BriefSummary,omitnil,omitempty" name:"BriefSummary"`
}

type SurgeryBreast struct {
	// 乳房总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type SurgeryBriefSummary struct {
	// 外科小结
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type SurgeryConditionBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 手术历史
	SurgeryList []*SurgeryListBlock `json:"SurgeryList,omitnil,omitempty" name:"SurgeryList"`

	// 对外输出值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type SurgeryHeadNeck struct {
	// 头颈部总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type SurgeryHistory struct {
	// 手术名称
	SurgeryName *SurgeryAttr `json:"SurgeryName,omitnil,omitempty" name:"SurgeryName"`

	// 手术日期
	SurgeryDate *SurgeryAttr `json:"SurgeryDate,omitnil,omitempty" name:"SurgeryDate"`

	// 术前诊断
	PreoperativePathology *SurgeryAttr `json:"PreoperativePathology,omitnil,omitempty" name:"PreoperativePathology"`

	// 术中诊断
	IntraoperativePathology *SurgeryAttr `json:"IntraoperativePathology,omitnil,omitempty" name:"IntraoperativePathology"`

	// 术后诊断
	PostoperativePathology *SurgeryAttr `json:"PostoperativePathology,omitnil,omitempty" name:"PostoperativePathology"`

	// 出院诊断
	DischargeDiagnosis *SurgeryAttr `json:"DischargeDiagnosis,omitnil,omitempty" name:"DischargeDiagnosis"`
}

type SurgeryHistoryBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 手术列表
	Surgerylist []*SurgeryListBlock `json:"Surgerylist,omitnil,omitempty" name:"Surgerylist"`
}

type SurgeryListBlock struct {
	// 时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 类型
	TimeType *string `json:"TimeType,omitnil,omitempty" name:"TimeType"`

	// 名称
	Name []*string `json:"Name,omitnil,omitempty" name:"Name"`

	// 部位
	Part *string `json:"Part,omitnil,omitempty" name:"Part"`
}

type SurgeryLymphNode struct {
	// 浅表淋巴结总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type SurgerySkin struct {
	// 皮肤总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type SurgerySpinalExtremities struct {
	// 脊柱四肢总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 脊柱
	SpinalColumn *KeyValueItem `json:"SpinalColumn,omitnil,omitempty" name:"SpinalColumn"`

	// 四肢和关节
	LimbJoint *KeyValueItem `json:"LimbJoint,omitnil,omitempty" name:"LimbJoint"`

	// 平跛足
	Foot *KeyValueItem `json:"Foot,omitnil,omitempty" name:"Foot"`

	// 骨骼
	Bone *KeyValueItem `json:"Bone,omitnil,omitempty" name:"Bone"`

	// 步态
	Gait *KeyValueItem `json:"Gait,omitnil,omitempty" name:"Gait"`

	// 残疾或畸形
	Deformity *KeyValueItem `json:"Deformity,omitnil,omitempty" name:"Deformity"`
}

type SurgeryThyroid struct {
	// 甲状腺总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`
}

type SurgeryUrogenitalSystem struct {
	// 泌尿生殖系统总体描述
	Text *KeyValueItem `json:"Text,omitnil,omitempty" name:"Text"`

	// 前列腺
	Prostate *KeyValueItem `json:"Prostate,omitnil,omitempty" name:"Prostate"`

	// 外生殖器（男性）
	ExternalReproductiveOrgans *KeyValueItem `json:"ExternalReproductiveOrgans,omitnil,omitempty" name:"ExternalReproductiveOrgans"`
}

type SymptomInfo struct {
	// 等级
	Grade *BlockInfo `json:"Grade,omitnil,omitempty" name:"Grade"`

	// 部位
	Part *Part `json:"Part,omitnil,omitempty" name:"Part"`

	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 病变
	Symptom *BlockInfo `json:"Symptom,omitnil,omitempty" name:"Symptom"`

	// 属性
	Attrs []*BlockInfo `json:"Attrs,omitnil,omitempty" name:"Attrs"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`
}

type TableIndicators struct {
	// 项目列表
	Indicators []*IndicatorItemV2 `json:"Indicators,omitnil,omitempty" name:"Indicators"`

	// 采样标本
	Sample *BaseItem `json:"Sample,omitnil,omitempty" name:"Sample"`
}

type Template struct {
	// 患者信息
	PatientInfo *PatientInfo `json:"PatientInfo,omitnil,omitempty" name:"PatientInfo"`

	// 报告信息
	ReportInfo *ReportInfo `json:"ReportInfo,omitnil,omitempty" name:"ReportInfo"`

	// 检查报告
	Check *Check `json:"Check,omitnil,omitempty" name:"Check"`

	// 病理报告
	Pathology *PathologyReport `json:"Pathology,omitnil,omitempty" name:"Pathology"`

	// 出院报告，入院报告，门诊病历
	MedDoc *MedDoc `json:"MedDoc,omitnil,omitempty" name:"MedDoc"`

	// 诊断证明
	DiagCert *DiagCert `json:"DiagCert,omitnil,omitempty" name:"DiagCert"`

	// 病案首页
	FirstPage *FirstPage `json:"FirstPage,omitnil,omitempty" name:"FirstPage"`

	// 检验报告
	Indicator *Indicator `json:"Indicator,omitnil,omitempty" name:"Indicator"`

	// 报告类型
	ReportType *string `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 门诊病历信息
	MedicalRecordInfo *MedicalRecordInfo `json:"MedicalRecordInfo,omitnil,omitempty" name:"MedicalRecordInfo"`

	// 出入院信息
	Hospitalization *Hospitalization `json:"Hospitalization,omitnil,omitempty" name:"Hospitalization"`

	// 手术记录
	Surgery *Surgery `json:"Surgery,omitnil,omitempty" name:"Surgery"`

	// 心电图报告
	Electrocardiogram *Electrocardiogram `json:"Electrocardiogram,omitnil,omitempty" name:"Electrocardiogram"`

	// 内窥镜报告
	Endoscopy *Endoscopy `json:"Endoscopy,omitnil,omitempty" name:"Endoscopy"`

	// 处方单
	Prescription *Prescription `json:"Prescription,omitnil,omitempty" name:"Prescription"`

	// 疫苗接种凭证
	VaccineCertificate *VaccineCertificate `json:"VaccineCertificate,omitnil,omitempty" name:"VaccineCertificate"`

	// OCR文本
	OcrText *string `json:"OcrText,omitnil,omitempty" name:"OcrText"`

	// OCR拼接后文本
	OcrResult *string `json:"OcrResult,omitnil,omitempty" name:"OcrResult"`

	// 报告类型
	ReportTypeDesc *string `json:"ReportTypeDesc,omitnil,omitempty" name:"ReportTypeDesc"`

	// 病理报告v2
	PathologyV2 *PathologyV2 `json:"PathologyV2,omitnil,omitempty" name:"PathologyV2"`

	// 碳14尿素呼气试验
	C14 *Indicator `json:"C14,omitnil,omitempty" name:"C14"`

	// 体检结论
	Exame *Exame `json:"Exame,omitnil,omitempty" name:"Exame"`

	// 出院报告v2，入院报告v2，门诊病历v2
	MedDocV2 *DischargeInfoBlock `json:"MedDocV2,omitnil,omitempty" name:"MedDocV2"`

	// 检验报告v3
	IndicatorV3 *IndicatorV3 `json:"IndicatorV3,omitnil,omitempty" name:"IndicatorV3"`

	// 核酸报告
	Covid *CovidItemsInfo `json:"Covid,omitnil,omitempty" name:"Covid"`

	// 孕产报告
	Maternity *Maternity `json:"Maternity,omitnil,omitempty" name:"Maternity"`

	// 眼科报告
	Eye *EyeItemsInfo `json:"Eye,omitnil,omitempty" name:"Eye"`

	// 出生证明
	BirthCert *BirthCert `json:"BirthCert,omitnil,omitempty" name:"BirthCert"`

	// 时间轴
	Timeline *TimelineInformation `json:"Timeline,omitnil,omitempty" name:"Timeline"`

	// 内窥镜报告V2
	EndoscopyV2 *Check `json:"EndoscopyV2,omitnil,omitempty" name:"EndoscopyV2"`
}

// Predefined struct for user
type TextToClassRequestParams struct {
	// 报告文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil,omitempty" name:"UserType"`
}

type TextToClassRequest struct {
	*tchttp.BaseRequest
	
	// 报告文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil,omitempty" name:"UserType"`
}

func (r *TextToClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "UserType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToClassResponseParams struct {
	// 分类结果
	TextTypeList []*TextType `json:"TextTypeList,omitnil,omitempty" name:"TextTypeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextToClassResponse struct {
	*tchttp.BaseResponse
	Response *TextToClassResponseParams `json:"Response"`
}

func (r *TextToClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToObjectRequestParams struct {
	// 报告文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 报告类型，目前支持12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为False，则Type字段不能为0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitnil,omitempty" name:"IsUsedClassify"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病历记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitnil,omitempty" name:"ReportTypeVersion"`
}

type TextToObjectRequest struct {
	*tchttp.BaseRequest
	
	// 报告文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 报告类型，目前支持12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为False，则Type字段不能为0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitnil,omitempty" name:"IsUsedClassify"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病历记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitnil,omitempty" name:"ReportTypeVersion"`
}

func (r *TextToObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "Type")
	delete(f, "IsUsedClassify")
	delete(f, "UserType")
	delete(f, "ReportTypeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToObjectResponseParams struct {
	// 报告结构化结果
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextToObjectResponse struct {
	*tchttp.BaseResponse
	Response *TextToObjectResponseParams `json:"Response"`
}

func (r *TextToObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextType struct {
	// 类别Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 类别层级
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 类别名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type TextTypeListBlock struct {
	// 文本类型列表
	TextTypeList []*TextType `json:"TextTypeList,omitnil,omitempty" name:"TextTypeList"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type Time struct {
	// 具体时间类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 时间值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TimelineEvent struct {
	// 事件类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 原文本
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 事件子类型
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 事件发生时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 事件值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 位置坐标
	Rectangle *Rectangle `json:"Rectangle,omitnil,omitempty" name:"Rectangle"`

	// 事件发生地点
	Place *string `json:"Place,omitnil,omitempty" name:"Place"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type TimelineInformation struct {
	// 时间轴
	Timeline []*TimelineEvent `json:"Timeline,omitnil,omitempty" name:"Timeline"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type TransfusionHistoryBlock struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 状态
	State *bool `json:"State,omitnil,omitempty" name:"State"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TreatmentRecord struct {
	// 入院
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: DmissionCondition is deprecated.
	DmissionCondition *string `json:"DmissionCondition,omitnil,omitempty" name:"DmissionCondition"`

	// 主诉
	ChiefComplaint *string `json:"ChiefComplaint,omitnil,omitempty" name:"ChiefComplaint"`

	// 现病史
	//
	// Deprecated: DiseasePresent is deprecated.
	DiseasePresent *string `json:"DiseasePresent,omitnil,omitempty" name:"DiseasePresent"`

	// 主要症状体征
	SymptomsAndSigns *string `json:"SymptomsAndSigns,omitnil,omitempty" name:"SymptomsAndSigns"`

	// 辅助检查
	AuxiliaryExamination *string `json:"AuxiliaryExamination,omitnil,omitempty" name:"AuxiliaryExamination"`

	// 体格检查
	BodyExamination *string `json:"BodyExamination,omitnil,omitempty" name:"BodyExamination"`

	// 专科检查
	SpecialistExamination *string `json:"SpecialistExamination,omitnil,omitempty" name:"SpecialistExamination"`

	// 精神检查
	MentalExamination *string `json:"MentalExamination,omitnil,omitempty" name:"MentalExamination"`

	// 检查记录
	CheckRecord *string `json:"CheckRecord,omitnil,omitempty" name:"CheckRecord"`

	// 化验结果
	InspectResult *string `json:"InspectResult,omitnil,omitempty" name:"InspectResult"`

	// 切口愈合情况
	IncisionHealing *string `json:"IncisionHealing,omitnil,omitempty" name:"IncisionHealing"`

	// 处理意见
	TreatmentSuggestion *string `json:"TreatmentSuggestion,omitnil,omitempty" name:"TreatmentSuggestion"`

	// 门诊随访要求
	FollowUpRequirements *string `json:"FollowUpRequirements,omitnil,omitempty" name:"FollowUpRequirements"`

	// 诊疗经过
	CheckAndTreatmentProcess *string `json:"CheckAndTreatmentProcess,omitnil,omitempty" name:"CheckAndTreatmentProcess"`

	// 手术经过
	SurgeryCondition *string `json:"SurgeryCondition,omitnil,omitempty" name:"SurgeryCondition"`

	// 入院情况
	ConditionChanges *string `json:"ConditionChanges,omitnil,omitempty" name:"ConditionChanges"`

	// 出院情况
	DischargeCondition *string `json:"DischargeCondition,omitnil,omitempty" name:"DischargeCondition"`

	// pTNM信息
	PTNM *string `json:"PTNM,omitnil,omitempty" name:"PTNM"`

	// pTNMM信息
	PTNMM *string `json:"PTNMM,omitnil,omitempty" name:"PTNMM"`

	// pTNMN信息
	PTNMN *string `json:"PTNMN,omitnil,omitempty" name:"PTNMN"`

	// pTNMT信息
	PTNMT *string `json:"PTNMT,omitnil,omitempty" name:"PTNMT"`

	// ECOG信息
	ECOG *string `json:"ECOG,omitnil,omitempty" name:"ECOG"`

	// NRS信息
	NRS *string `json:"NRS,omitnil,omitempty" name:"NRS"`

	// KPS信息
	KPS *string `json:"KPS,omitnil,omitempty" name:"KPS"`

	// 死亡日期
	DeathDate *string `json:"DeathDate,omitnil,omitempty" name:"DeathDate"`

	// 复发日期
	RelapseDate *string `json:"RelapseDate,omitnil,omitempty" name:"RelapseDate"`

	// 观测天数
	ObservationDays *string `json:"ObservationDays,omitnil,omitempty" name:"ObservationDays"`

	// 入院
	AdmissionCondition *string `json:"AdmissionCondition,omitnil,omitempty" name:"AdmissionCondition"`
}

type TreatmentRecordBlock struct {
	// 免疫组化
	Immunohistochemistry *ImmunohistochemistryBlock `json:"Immunohistochemistry,omitnil,omitempty" name:"Immunohistochemistry"`

	// 主诉
	ChiefComplaint *ChiefComplaintBlock `json:"ChiefComplaint,omitnil,omitempty" name:"ChiefComplaint"`

	// 入院情况
	AdmissionCondition *AdmissionConditionBlock `json:"AdmissionCondition,omitnil,omitempty" name:"AdmissionCondition"`

	// 查体
	BodyExamination *BodyExaminationBlock `json:"BodyExamination,omitnil,omitempty" name:"BodyExamination"`

	// 入院诊断
	AdmissionDiagnosis *AdmissionDiagnosisBlock `json:"AdmissionDiagnosis,omitnil,omitempty" name:"AdmissionDiagnosis"`

	// 入院中医诊断
	AdmissionTraditionalDiagnosis *AdmissionDiagnosisBlock `json:"AdmissionTraditionalDiagnosis,omitnil,omitempty" name:"AdmissionTraditionalDiagnosis"`

	// 入院西医诊断
	AdmissionModernDiagnosis *AdmissionDiagnosisBlock `json:"AdmissionModernDiagnosis,omitnil,omitempty" name:"AdmissionModernDiagnosis"`

	// 病理诊断
	PathologicalDiagnosis *PathologicalDiagnosisBlock `json:"PathologicalDiagnosis,omitnil,omitempty" name:"PathologicalDiagnosis"`

	// 现病史
	DiseasePresent *DiseasePresentBlock `json:"DiseasePresent,omitnil,omitempty" name:"DiseasePresent"`

	// 体征
	SymptomsAndSigns *DiseasePresentBlock `json:"SymptomsAndSigns,omitnil,omitempty" name:"SymptomsAndSigns"`

	// 辅助检查
	AuxiliaryExamination *DiseasePresentBlock `json:"AuxiliaryExamination,omitnil,omitempty" name:"AuxiliaryExamination"`

	// 特殊检查
	SpecialistExamination *DiseasePresentBlock `json:"SpecialistExamination,omitnil,omitempty" name:"SpecialistExamination"`

	// 精神检查
	MentalExamination *DiseasePresentBlock `json:"MentalExamination,omitnil,omitempty" name:"MentalExamination"`

	// 检查记录
	CheckRecord *DiseasePresentBlock `json:"CheckRecord,omitnil,omitempty" name:"CheckRecord"`

	// 检查结果
	InspectResult *DiseasePresentBlock `json:"InspectResult,omitnil,omitempty" name:"InspectResult"`

	// 治疗经过
	CheckAndTreatmentProcess *DiseasePresentBlock `json:"CheckAndTreatmentProcess,omitnil,omitempty" name:"CheckAndTreatmentProcess"`

	// 手术经过
	SurgeryCondition *SurgeryConditionBlock `json:"SurgeryCondition,omitnil,omitempty" name:"SurgeryCondition"`

	// 切口愈合
	IncisionHealing *DiseasePresentBlock `json:"IncisionHealing,omitnil,omitempty" name:"IncisionHealing"`

	// 出院诊断
	DischargeDiagnosis *DischargeDiagnosisBlock `json:"DischargeDiagnosis,omitnil,omitempty" name:"DischargeDiagnosis"`

	// 出院中医诊断
	DischargeTraditionalDiagnosis *DiseasePresentBlock `json:"DischargeTraditionalDiagnosis,omitnil,omitempty" name:"DischargeTraditionalDiagnosis"`

	// 出院西医诊断
	DischargeModernDiagnosis *DischargeDiagnosisBlock `json:"DischargeModernDiagnosis,omitnil,omitempty" name:"DischargeModernDiagnosis"`

	// 出院情况
	DischargeCondition *DischargeConditionBlock `json:"DischargeCondition,omitnil,omitempty" name:"DischargeCondition"`

	// 出院医嘱
	DischargeInstructions *DiseasePresentBlock `json:"DischargeInstructions,omitnil,omitempty" name:"DischargeInstructions"`

	// 治疗建议
	TreatmentSuggestion *DiseasePresentBlock `json:"TreatmentSuggestion,omitnil,omitempty" name:"TreatmentSuggestion"`

	// 随访
	FollowUpRequirements *DiseasePresentBlock `json:"FollowUpRequirements,omitnil,omitempty" name:"FollowUpRequirements"`

	// 治疗情况变化
	ConditionChanges *DiseasePresentBlock `json:"ConditionChanges,omitnil,omitempty" name:"ConditionChanges"`

	// 收缩压
	PulmonaryArterySystolicPressure *DiseasePresentBlock `json:"PulmonaryArterySystolicPressure,omitnil,omitempty" name:"PulmonaryArterySystolicPressure"`

	// bclc分期
	BCLC *DiseasePresentBlock `json:"BCLC,omitnil,omitempty" name:"BCLC"`

	// PTNM分期
	PTNM *PTNMBlock `json:"PTNM,omitnil,omitempty" name:"PTNM"`

	// ECOG评分
	ECOG *DiseasePresentBlock `json:"ECOG,omitnil,omitempty" name:"ECOG"`

	// NRS评分
	NRS *DiseasePresentBlock `json:"NRS,omitnil,omitempty" name:"NRS"`

	// kps评分
	KPS *DiseasePresentBlock `json:"KPS,omitnil,omitempty" name:"KPS"`

	// 癌症分期
	Cancerstaging *ClinicalStaging `json:"Cancerstaging,omitnil,omitempty" name:"Cancerstaging"`

	// 死亡时间
	DeathDate *DeathDateBlock `json:"DeathDate,omitnil,omitempty" name:"DeathDate"`

	// 复发日期
	RelapseDate *RelapseDateBlock `json:"RelapseDate,omitnil,omitempty" name:"RelapseDate"`

	// 观察日期
	ObservationDays *DeathDateBlock `json:"ObservationDays,omitnil,omitempty" name:"ObservationDays"`

	// 切口愈合情况
	IncisionHealingText *string `json:"IncisionHealingText,omitnil,omitempty" name:"IncisionHealingText"`

	// 辅助检查
	AuxiliaryExaminationText *string `json:"AuxiliaryExaminationText,omitnil,omitempty" name:"AuxiliaryExaminationText"`

	// 特殊检查
	SpecialExamText *string `json:"SpecialExamText,omitnil,omitempty" name:"SpecialExamText"`

	// 门诊诊断
	OutpatientDiagnosisText *string `json:"OutpatientDiagnosisText,omitnil,omitempty" name:"OutpatientDiagnosisText"`

	// 入院情况
	AdmissionConditionText *string `json:"AdmissionConditionText,omitnil,omitempty" name:"AdmissionConditionText"`

	// 诊疗经过
	CheckAndTreatmentProcessText *string `json:"CheckAndTreatmentProcessText,omitnil,omitempty" name:"CheckAndTreatmentProcessText"`

	// 体征
	SymptomsAndSignsText *string `json:"SymptomsAndSignsText,omitnil,omitempty" name:"SymptomsAndSignsText"`

	// 出院医嘱
	DischargeInstructionsText *string `json:"DischargeInstructionsText,omitnil,omitempty" name:"DischargeInstructionsText"`

	// 入院诊断
	AdmissionDiagnosisText *string `json:"AdmissionDiagnosisText,omitnil,omitempty" name:"AdmissionDiagnosisText"`

	// 手术情况
	SurgeryConditionText *string `json:"SurgeryConditionText,omitnil,omitempty" name:"SurgeryConditionText"`

	// 病理诊断
	PathologicalDiagnosisText *string `json:"PathologicalDiagnosisText,omitnil,omitempty" name:"PathologicalDiagnosisText"`

	// 出院情况
	DischargeConditionText *string `json:"DischargeConditionText,omitnil,omitempty" name:"DischargeConditionText"`

	// 检查记录
	CheckRecordText *string `json:"CheckRecordText,omitnil,omitempty" name:"CheckRecordText"`

	// 主诉
	ChiefComplaintText *string `json:"ChiefComplaintText,omitnil,omitempty" name:"ChiefComplaintText"`

	// 出院诊断
	DischargeDiagnosisText *string `json:"DischargeDiagnosisText,omitnil,omitempty" name:"DischargeDiagnosisText"`
}

type TuberInfo struct {
	// 类型
	Type *BlockInfo `json:"Type,omitnil,omitempty" name:"Type"`

	// 部位
	Part *Part `json:"Part,omitnil,omitempty" name:"Part"`

	// 大小
	Size []*Size `json:"Size,omitnil,omitempty" name:"Size"`

	// 多发
	Multiple *Multiple `json:"Multiple,omitnil,omitempty" name:"Multiple"`

	// 纵横比
	AspectRatio *AspectRatio `json:"AspectRatio,omitnil,omitempty" name:"AspectRatio"`

	// 边缘
	Edge *BlockInfo `json:"Edge,omitnil,omitempty" name:"Edge"`

	// 内部回声
	InnerEcho *BlockInfo `json:"InnerEcho,omitnil,omitempty" name:"InnerEcho"`

	// 外部回声
	RearEcho *BlockInfo `json:"RearEcho,omitnil,omitempty" name:"RearEcho"`

	// 弹性质地
	Elastic *Elastic `json:"Elastic,omitnil,omitempty" name:"Elastic"`

	// 形状
	Shape *BlockInfo `json:"Shape,omitnil,omitempty" name:"Shape"`

	// 形态
	ShapeAttr *BlockInfo `json:"ShapeAttr,omitnil,omitempty" name:"ShapeAttr"`

	// 皮髓质信息
	SkinMedulla *BlockInfo `json:"SkinMedulla,omitnil,omitempty" name:"SkinMedulla"`

	// 变化趋势
	Trend *BlockInfo `json:"Trend,omitnil,omitempty" name:"Trend"`

	// 钙化
	Calcification *BlockInfo `json:"Calcification,omitnil,omitempty" name:"Calcification"`

	// 包膜
	Envelope *BlockInfo `json:"Envelope,omitnil,omitempty" name:"Envelope"`

	// 强化
	Enhancement *BlockInfo `json:"Enhancement,omitnil,omitempty" name:"Enhancement"`

	// 淋巴结
	LymphEnlargement *BlockInfo `json:"LymphEnlargement,omitnil,omitempty" name:"LymphEnlargement"`

	// 淋巴门
	LymphDoor *BlockInfo `json:"LymphDoor,omitnil,omitempty" name:"LymphDoor"`

	// 活动度
	Activity *BlockInfo `json:"Activity,omitnil,omitempty" name:"Activity"`

	// 手术情况
	Operation *BlockInfo `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 血液cdfi
	CDFI *BlockInfo `json:"CDFI,omitnil,omitempty" name:"CDFI"`

	// 原文位置
	Index []*int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 大小状态
	SizeStatus *BlockInfo `json:"SizeStatus,omitnil,omitempty" name:"SizeStatus"`

	// 内部回声分布
	InnerEchoDistribution *BlockInfo `json:"InnerEchoDistribution,omitnil,omitempty" name:"InnerEchoDistribution"`

	// 内部回声类型
	InnerEchoType []*BlockInfo `json:"InnerEchoType,omitnil,omitempty" name:"InnerEchoType"`

	// 轮廓
	Outline *BlockInfo `json:"Outline,omitnil,omitempty" name:"Outline"`

	// 结构
	Structure *BlockInfo `json:"Structure,omitnil,omitempty" name:"Structure"`

	// 密度
	Density *BlockInfo `json:"Density,omitnil,omitempty" name:"Density"`

	// 血管
	Vas *BlockInfo `json:"Vas,omitnil,omitempty" name:"Vas"`

	// 囊壁
	Cysticwall *BlockInfo `json:"Cysticwall,omitnil,omitempty" name:"Cysticwall"`

	// 被膜
	Capsule *BlockInfo `json:"Capsule,omitnil,omitempty" name:"Capsule"`

	// 峡部厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IsthmusThicknese is deprecated.
	IsthmusThicknese *Size `json:"IsthmusThicknese,omitnil,omitempty" name:"IsthmusThicknese"`

	// 原文
	Src *string `json:"Src,omitnil,omitempty" name:"Src"`

	// 透声度
	Transparent *BlockInfo `json:"Transparent,omitnil,omitempty" name:"Transparent"`

	// MRI ADC
	MriAdc *BlockInfo `json:"MriAdc,omitnil,omitempty" name:"MriAdc"`

	// MRI DWI
	MriDwi *BlockInfo `json:"MriDwi,omitnil,omitempty" name:"MriDwi"`

	// MRI T1信号
	MriT1 *BlockInfo `json:"MriT1,omitnil,omitempty" name:"MriT1"`

	// MRI T2信号
	MriT2 *BlockInfo `json:"MriT2,omitnil,omitempty" name:"MriT2"`

	// CT HU值
	CtHu *BlockInfo `json:"CtHu,omitnil,omitempty" name:"CtHu"`

	// SUmax值
	Suvmax *BlockInfo `json:"Suvmax,omitnil,omitempty" name:"Suvmax"`

	// 代谢情况
	Metabolism *BlockInfo `json:"Metabolism,omitnil,omitempty" name:"Metabolism"`

	// 放射性摄取
	RadioactiveUptake *BlockInfo `json:"RadioactiveUptake,omitnil,omitempty" name:"RadioactiveUptake"`

	// 病变
	SymDesc *BlockInfo `json:"SymDesc,omitnil,omitempty" name:"SymDesc"`

	// 影像特征
	ImageFeature *BlockInfo `json:"ImageFeature,omitnil,omitempty" name:"ImageFeature"`

	// 在报告图片中的坐标
	Coords []*Coord `json:"Coords,omitnil,omitempty" name:"Coords"`

	// 峡部厚度
	IsthmusThickness *Size `json:"IsthmusThickness,omitnil,omitempty" name:"IsthmusThickness"`
}

// Predefined struct for user
type TurnPDFToObjectAsyncGetResultRequestParams struct {
	// 加密任务ID。在上一步通过TurnPDFToObjectAsync 接口返回的TaskID。
	// 1、建议在上一步调用TurnPDFToObjectAsync接口传入PDF之后，等5-10分钟再调用此接口获取 json 结果。如果任务还没完成，可以等待几分钟之后再重新调用此接口获取 json 结果。
	// 2、临时加密存储的 json 结果会 24 小时后定时自动删除，因此TaskID 仅 24 小时内有效。
	// 3、TaskID 与腾讯云的账号绑定，通过 TurnPDFToObjectAsync 传入 PDF 文件和通过 TurnPDFToObjectAsyncGetResult 获取 json 结果，必须是同一个腾讯云账号，否则无法获取到 json 结果。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

type TurnPDFToObjectAsyncGetResultRequest struct {
	*tchttp.BaseRequest
	
	// 加密任务ID。在上一步通过TurnPDFToObjectAsync 接口返回的TaskID。
	// 1、建议在上一步调用TurnPDFToObjectAsync接口传入PDF之后，等5-10分钟再调用此接口获取 json 结果。如果任务还没完成，可以等待几分钟之后再重新调用此接口获取 json 结果。
	// 2、临时加密存储的 json 结果会 24 小时后定时自动删除，因此TaskID 仅 24 小时内有效。
	// 3、TaskID 与腾讯云的账号绑定，通过 TurnPDFToObjectAsync 传入 PDF 文件和通过 TurnPDFToObjectAsyncGetResult 获取 json 结果，必须是同一个腾讯云账号，否则无法获取到 json 结果。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

func (r *TurnPDFToObjectAsyncGetResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectAsyncGetResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TurnPDFToObjectAsyncGetResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnPDFToObjectAsyncGetResultResponseParams struct {
	// 报告结构化结果
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// 多级分类结果
	TextTypeList []*TextType `json:"TextTypeList,omitnil,omitempty" name:"TextTypeList"`

	// 报告结构化结果(体检报告PDF结构化接口返回的 json 内容非常多，建议通过本地代码调用)
	Block *Block `json:"Block,omitnil,omitempty" name:"Block"`

	// 是否使用Block字段
	IsBlock *bool `json:"IsBlock,omitnil,omitempty" name:"IsBlock"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TurnPDFToObjectAsyncGetResultResponse struct {
	*tchttp.BaseResponse
	Response *TurnPDFToObjectAsyncGetResultResponseParams `json:"Response"`
}

func (r *TurnPDFToObjectAsyncGetResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectAsyncGetResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnPDFToObjectAsyncRequestParams struct {
	// 体检报告PDF文件信息, 目前只支持传PDF文件的Base64编码字符(PDF文件不能超过10MB，如果超过建议先压缩PDF，再转成base64)
	PdfInfo *PdfInfo `json:"PdfInfo,omitnil,omitempty" name:"PdfInfo"`

	// PDF文件中的文字是否为文本内容.
	// 如果该字段为true,那么就会自动判断是电子版还是图片，自动选择直接读取文字还是 OCR 方式.
	// 如果该字段为false, 那么始终采用 OCR 方式
	TextBasedPdfFlag *bool `json:"TextBasedPdfFlag,omitnil,omitempty" name:"TextBasedPdfFlag"`
}

type TurnPDFToObjectAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 体检报告PDF文件信息, 目前只支持传PDF文件的Base64编码字符(PDF文件不能超过10MB，如果超过建议先压缩PDF，再转成base64)
	PdfInfo *PdfInfo `json:"PdfInfo,omitnil,omitempty" name:"PdfInfo"`

	// PDF文件中的文字是否为文本内容.
	// 如果该字段为true,那么就会自动判断是电子版还是图片，自动选择直接读取文字还是 OCR 方式.
	// 如果该字段为false, 那么始终采用 OCR 方式
	TextBasedPdfFlag *bool `json:"TextBasedPdfFlag,omitnil,omitempty" name:"TextBasedPdfFlag"`
}

func (r *TurnPDFToObjectAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PdfInfo")
	delete(f, "TextBasedPdfFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TurnPDFToObjectAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnPDFToObjectAsyncResponseParams struct {
	// 加密任务ID。 
	// 1、此 ID 是经过加密生成，是用于获取 PDF 返回 json 的凭证，需要由客户存储该 TaskID。
	// 2、建议在获取到TaskID 后，5-10分钟后再调用 TurnPDFToObjectAsyncGetResult 接口获取 json 结果。
	// 3、使用此接口，腾讯不会存储传入的 PDF 文件，但是会临时加密存储对应的 json 结果。如果不希望腾讯临时加密存储 json 结果，请使用 TurnPDFToObject 接口。
	// 4、加密存储的 json 结果会24小时后定时自动删除，因此TaskID 仅 24 小时内有效，请在24小时内调用接口 TurnPDFToObjectAsyncGetResult 获取对应 json 结果。
	// 5、TaskID 与腾讯云的账号绑定，通过 TurnPDFToObjectAsync 传入PDF文件和通过 TurnPDFToObjectAsyncGetResult 获取 json 结果，必须是同一个腾讯云账号。即其它人就算获取到 TaskID 也无法获取到 json 结果。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TurnPDFToObjectAsyncResponse struct {
	*tchttp.BaseResponse
	Response *TurnPDFToObjectAsyncResponseParams `json:"Response"`
}

func (r *TurnPDFToObjectAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnPDFToObjectRequestParams struct {
	// 体检报告PDF文件信息, 目前只支持传PDF文件的Base64编码字符(PDF文件不能超过10MB，如果超过建议先压缩PDF，再转成base64)
	PdfInfo *PdfInfo `json:"PdfInfo,omitnil,omitempty" name:"PdfInfo"`

	// PDF文件中的文字是否为文本内容.
	// 如果该字段为true,那么就会自动判断是电子版还是图片，自动选择直接读取文字还是 OCR 方式.
	// 如果该字段为false, 那么始终采用 OCR 方式
	TextBasedPdfFlag *bool `json:"TextBasedPdfFlag,omitnil,omitempty" name:"TextBasedPdfFlag"`
}

type TurnPDFToObjectRequest struct {
	*tchttp.BaseRequest
	
	// 体检报告PDF文件信息, 目前只支持传PDF文件的Base64编码字符(PDF文件不能超过10MB，如果超过建议先压缩PDF，再转成base64)
	PdfInfo *PdfInfo `json:"PdfInfo,omitnil,omitempty" name:"PdfInfo"`

	// PDF文件中的文字是否为文本内容.
	// 如果该字段为true,那么就会自动判断是电子版还是图片，自动选择直接读取文字还是 OCR 方式.
	// 如果该字段为false, 那么始终采用 OCR 方式
	TextBasedPdfFlag *bool `json:"TextBasedPdfFlag,omitnil,omitempty" name:"TextBasedPdfFlag"`
}

func (r *TurnPDFToObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PdfInfo")
	delete(f, "TextBasedPdfFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TurnPDFToObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnPDFToObjectResponseParams struct {
	// 报告结构化结果
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// 多级分类结果
	TextTypeList []*TextType `json:"TextTypeList,omitnil,omitempty" name:"TextTypeList"`

	// 报告结构化结果(体检报告PDF结构化接口返回的 json 内容非常多，建议通过本地代码调用)
	Block *Block `json:"Block,omitnil,omitempty" name:"Block"`

	// 是否使用Block字段
	IsBlock *bool `json:"IsBlock,omitnil,omitempty" name:"IsBlock"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TurnPDFToObjectResponse struct {
	*tchttp.BaseResponse
	Response *TurnPDFToObjectResponseParams `json:"Response"`
}

func (r *TurnPDFToObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Vaccination struct {
	// 序号
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 疫苗名称
	Vaccine *string `json:"Vaccine,omitnil,omitempty" name:"Vaccine"`

	// 剂次
	Dose *string `json:"Dose,omitnil,omitempty" name:"Dose"`

	// 接种日期
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 疫苗批号
	LotNumber *string `json:"LotNumber,omitnil,omitempty" name:"LotNumber"`

	// 生产企业
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// 接种单位
	Clinic *string `json:"Clinic,omitnil,omitempty" name:"Clinic"`

	// 接种部位
	Site *string `json:"Site,omitnil,omitempty" name:"Site"`

	// 接种者
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 疫苗批号
	Lot *string `json:"Lot,omitnil,omitempty" name:"Lot"`
}

type VaccineCertificate struct {
	// 免疫接种列表
	VaccineList []*Vaccination `json:"VaccineList,omitnil,omitempty" name:"VaccineList"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type Value struct {
	// 等级
	Grade *string `json:"Grade,omitnil,omitempty" name:"Grade"`

	// 百分比
	Percent []*float64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 阳性
	Positive *string `json:"Positive,omitnil,omitempty" name:"Positive"`
}

type ValueBlock struct {
	// 等级
	Grade *string `json:"Grade,omitnil,omitempty" name:"Grade"`

	// 百分比
	Percent []*float64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 阳性阴性
	Positive *string `json:"Positive,omitnil,omitempty" name:"Positive"`
}

type ValueUnitItem struct {
	// 类型
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目原文
	Item *PhysicalBaseItem `json:"Item,omitnil,omitempty" name:"Item"`

	// 数值
	Result *PhysicalBaseItem `json:"Result,omitnil,omitempty" name:"Result"`

	// 单位
	Unit *PhysicalBaseItem `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 数据在原PDF文件中的第几页
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}