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

const (
	// 此产品的特有错误码

	// API元数据解析失败。
	FAILEDOPERATION_APIMETAPARSEFAILED = "FailedOperation.ApiMetaParseFailed"

	// 调用TSF-APM服务获取配置列表失败
	FAILEDOPERATION_APMGETCONFIGSFAILED = "FailedOperation.ApmGetConfigsFailed"

	// 创建应用，获取ES鉴权信息失败。
	FAILEDOPERATION_APPLICATIONCREATEESATUHERROR = "FailedOperation.ApplicationCreateEsAtuhError"

	// 应用查询失败。
	FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"

	// TSF应用性能管理调用tsf-dcfg模块失败
	FAILEDOPERATION_CALLTSFDCFGFAILED = "FailedOperation.CallTsfDcfgFailed"

	// 调用CDI状态接口失败
	FAILEDOPERATION_CDISTATUSFAILED = "FailedOperation.CdiStatusFailed"

	// 调用CDI状态返回数据异常
	FAILEDOPERATION_CDISTATUSINVALID = "FailedOperation.CdiStatusInvalid"

	// TSF云API请求调用失败。
	FAILEDOPERATION_CLOUDAPIPROXYERROR = "FailedOperation.CloudApiProxyError"

	// 创建集群，开通VPC网络权限失败。
	FAILEDOPERATION_CLUSTERCREATEVPCFAIL = "FailedOperation.ClusterCreateVpcFail"

	// 本环境仅允许操作名称 %s 的集群
	FAILEDOPERATION_CLUSTEROPERATEENVILLEGALERROR = "FailedOperation.ClusterOperateEnvIllegalError"

	// 查询集群失败。
	FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"

	// 更新集群失败。
	FAILEDOPERATION_CLUSTERUPDATEFAILED = "FailedOperation.ClusterUpdateFailed"

	// 应用查询失败。
	FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"

	// 配置项创建失败。
	FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"

	// 部署组查询失败。
	FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"

	// 命名空间查询失败。
	FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"

	// 配置项查询失败。
	FAILEDOPERATION_CONFIGQUERYFAILED = "FailedOperation.ConfigQueryFailed"

	// 配置项发布信息查询失败。
	FAILEDOPERATION_CONFIGRELEASEQUERYFAILED = "FailedOperation.ConfigReleaseQueryFailed"

	// 配置模板创建失败。
	FAILEDOPERATION_CONFIGTEMPLATECREATEFAILED = "FailedOperation.ConfigTemplateCreateFailed"

	// 配置模板删除失败。
	FAILEDOPERATION_CONFIGTEMPLATEDELETEFAILED = "FailedOperation.ConfigTemplateDeleteFailed"

	// 配置模板导入失败。
	FAILEDOPERATION_CONFIGTEMPLATEIMPORTFAILED = "FailedOperation.ConfigTemplateImportFailed"

	// 配置模板分页查询失败。
	FAILEDOPERATION_CONFIGTEMPLATESEARCHLISTFAILED = "FailedOperation.ConfigTemplateSearchListFailed"

	// 配置模板更新失败。
	FAILEDOPERATION_CONFIGTEMPLATEUPDATEFAILED = "FailedOperation.ConfigTemplateUpdateFailed"

	// 部署组处于运行状态，无法启动。
	FAILEDOPERATION_CONTAINERGROUPGROUPHASRUN = "FailedOperation.ContainergroupGroupHasrun"

	// 部署组处于停止状态，无法执行此操作。
	FAILEDOPERATION_CONTAINERGROUPGROUPHASSTOP = "FailedOperation.ContainergroupGroupHasstop"

	// 调用 kube-api-server 失败。
	FAILEDOPERATION_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "FailedOperation.ContainergroupKuberneteApiInvokeError"

	// 连接 kube-api-server 失败。
	FAILEDOPERATION_CONTAINERGROUPKUBERNETECONNECTERROR = "FailedOperation.ContainergroupKuberneteConnectError"

	// 调用 kube-api-server 失败。
	FAILEDOPERATION_CONTAINERGROUPKUBERNETESAPIINVOKEERROR = "FailedOperation.ContainergroupKubernetesApiInvokeError"

	// 调用 kube-api-server 失败。
	FAILEDOPERATION_CONTAINERGROUPKUBERNETESCONNECTERROR = "FailedOperation.ContainergroupKubernetesConnectError"

	// k8s集群可能处于异常状态，请删除或恢复k8s集群后，再进行操作
	FAILEDOPERATION_CONTAINERGROUPKUBERNETESTKECLUSTERABNORMAL = "FailedOperation.ContainergroupKubernetesTkeClusterAbnormal"

	// 容器平台集群不可用，当前状态 %s
	FAILEDOPERATION_CPCLUSTERUNAVAILABLE = "FailedOperation.CpClusterUnavailable"

	// 健康检查配置失败。
	FAILEDOPERATION_CVMCAEMASTERHEALTHCHECKCONFIGERROR = "FailedOperation.CvmCaeMasterHealthCheckConfigError"

	// TSF暂时不能响应请求。
	FAILEDOPERATION_DISPATCHCOMMONERROR = "FailedOperation.DispatchCommonError"

	// 寻找独占配置中心相关指标异常
	FAILEDOPERATION_FINDMETRICSEXCLUSIVEERROR = "FailedOperation.FindMetricsExclusiveError"

	// 网关通用异常:%s。
	FAILEDOPERATION_GATEWAYCOMMONERROR = "FailedOperation.GatewayCommonError"

	// 远端访问错误: %s。
	FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"

	// 命名空间下存在部署组。
	FAILEDOPERATION_GROUPEXISTS = "FailedOperation.GroupExists"

	// 部署组查询失败。
	FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"

	// 镜像仓库已关联至部署组，请移除后重试
	FAILEDOPERATION_IMAGEVALIDATEISUSED = "FailedOperation.ImageValidateIsUsed"

	// 禁止直接删除应用镜像仓库, 需要先删除下面的镜像tag
	FAILEDOPERATION_IMAGEREPOREJECTDELERROR = "FailedOperation.ImagerepoRejectDelError"

	// tcr仓库绑定失败。
	FAILEDOPERATION_IMAGEREPOTCRBINDERROR = "FailedOperation.ImagerepoTcrBindError"

	// 机器实例删除失败。
	FAILEDOPERATION_INSTANCEDELETEFAILED = "FailedOperation.InstanceDeleteFailed"

	// 查询机器实例部分失败。
	FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"

	// 重装系统失败，请稍后重试。若无法解决，请联系智能客服或提交工单。
	FAILEDOPERATION_INSTANCERESETERROR = "FailedOperation.InstanceResetError"

	// 重装系统，请求超时。
	FAILEDOPERATION_INSTANCERESETTIMEOUT = "FailedOperation.InstanceResetTimeout"

	// 机器实例更新失败。
	FAILEDOPERATION_INSTANCEUPDATEFAILED = "FailedOperation.InstanceUpdateFailed"

	// 内部错误。
	FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"

	// 非法参数。
	FAILEDOPERATION_INVALIDPARAMETER = "FailedOperation.InvalidParameter"

	// 泳道从consul删除失败。
	FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"

	// 新增关联部署组不能为空。
	FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"

	// 泳道同步到consul失败。
	FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"

	// 泳道发布到mesh失败。
	FAILEDOPERATION_LANEINFORELEASEMESHFAILED = "FailedOperation.LaneInfoReleaseMeshFailed"

	// 全链路灰度规则启用失败。
	FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"

	// 用户全链路灰度规则最大100条。
	FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"

	// 无法创建命名空间。
	FAILEDOPERATION_NAMESPACECREATEFAILED = "FailedOperation.NamespaceCreateFailed"

	// 命名空间查询失败。
	FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"

	// 访问配置中心失败。
	FAILEDOPERATION_RATELIMITCONSULERROR = "FailedOperation.RatelimitConsulError"

	// 调用 Mesh API Server 失败。
	FAILEDOPERATION_RATELIMITMESHAPISERVICEERROR = "FailedOperation.RatelimitMeshApiServiceError"

	// 资源操作失败。
	FAILEDOPERATION_RESOURCEOPFAILED = "FailedOperation.ResourceOpFailed"

	// 路由就近访问策略启停用调用MESHAPI失败。
	FAILEDOPERATION_ROUTEAFFINITYMESHFAILED = "FailedOperation.RouteAffinityMeshFailed"

	// 服务路由规则启用生效失败。
	FAILEDOPERATION_ROUTEENABLECONSULFAILED = "FailedOperation.RouteEnableConsulFailed"

	// 路由就近访问，查询命名空间失败。
	FAILEDOPERATION_ROUTENAMESPACEREQUESTERROR = "FailedOperation.RouteNamespaceRequestError"

	// 服务数据库入库失败。
	FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"

	// 服务查询失败。
	FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"

	// 可观测性支持依赖TAPM业务系统绑定
	FAILEDOPERATION_TAPMBINDREQUIRED = "FailedOperation.TapmBindRequired"

	// 请使用TAPM提供的Otel探针
	FAILEDOPERATION_TAPMEXPIREDOTEL = "FailedOperation.TapmExpiredOtel"

	// 任务创建异常。
	FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"

	// 任务删除异常。
	FAILEDOPERATION_TASKDELETEERROR = "FailedOperation.TaskDeleteError"

	// 操作失败。
	FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"

	// 禁止操作。
	FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"

	// 任务下发异常。
	FAILEDOPERATION_TASKPUSHERROR = "FailedOperation.TaskPushError"

	// 任务查询异常。
	FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"

	// 停止任务失败。
	FAILEDOPERATION_TASKTERMINATEFAILED = "FailedOperation.TaskTerminateFailed"

	// 任务更新异常。
	FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"

	// TKE 集群创建失败，%s。
	FAILEDOPERATION_TKECLUSTERCREATEFAILED = "FailedOperation.TkeClusterCreateFailed"

	// TKE 集群删除失败。
	FAILEDOPERATION_TKECLUSTERDELETEFAILED = "FailedOperation.TkeClusterDeleteFailed"

	// TKE 集群查询失败。
	FAILEDOPERATION_TKECLUSTERQUERYFAILED = "FailedOperation.TkeClusterQueryFailed"

	// TOKEN查询失败。
	FAILEDOPERATION_TOKENQUERYFAILED = "FailedOperation.TokenQueryFailed"

	// TSF应用性能管理任务数据库查询失败。
	FAILEDOPERATION_TSFAPMAGENTTASKQUERYERROR = "FailedOperation.TsfApmAgentTaskQueryError"

	// TSF应用性能管理任务数据库写入失败。
	FAILEDOPERATION_TSFAPMAGENTTASKWRITEERROR = "FailedOperation.TsfApmAgentTaskWriteError"

	// TSF应用性能管理apm-agent无法与该实例建立连接。
	FAILEDOPERATION_TSFAPMAPMAGENTNOCONNECTION = "FailedOperation.TsfApmApmAgentNoConnection"

	// TSF应用性能管理业务日志配置与应用关联数据库写入失败。
	FAILEDOPERATION_TSFAPMBUSILOGCFGAPPRELATIONWRITEERROR = "FailedOperation.TsfApmBusiLogCfgAppRelationWriteError"

	// TSF应用性能管理业务日志配置数据库查询失败。
	FAILEDOPERATION_TSFAPMBUSILOGCFGQUERYERROR = "FailedOperation.TsfApmBusiLogCfgQueryError"

	// TSF应用性能管理业务日志配置解析规则查询失败。
	FAILEDOPERATION_TSFAPMBUSILOGCFGSCHEMAQUERYERROR = "FailedOperation.TsfApmBusiLogCfgSchemaQueryError"

	// TSF应用性能管理业务日志配置解析规则写入失败。
	FAILEDOPERATION_TSFAPMBUSILOGCFGSCHEMAWRITEERROR = "FailedOperation.TsfApmBusiLogCfgSchemaWriteError"

	// TSF应用性能管理业务日志配置数据库写入失败。
	FAILEDOPERATION_TSFAPMBUSILOGCFGWRITEERROR = "FailedOperation.TsfApmBusiLogCfgWriteError"

	// 调用master接口更新部署组失败。
	FAILEDOPERATION_TSFAPMCALLMASTERINTERFACEFAILED = "FailedOperation.TsfApmCallMasterInterfaceFailed"

	// TSF应用性能管理调用tsf-ms模块失败。
	FAILEDOPERATION_TSFAPMCALLTSFMSFAILED = "FailedOperation.TsfApmCallTsfMsFailed"

	// TSF应用性能管理CTSDB客户端调用失败。
	FAILEDOPERATION_TSFAPMCTSDBCLIENTREQUESTERROR = "FailedOperation.TsfApmCtsdbClientRequestError"

	// TSF应用性能管理ES客户端调用失败。
	FAILEDOPERATION_TSFAPMESCLIENTREQUESTERROR = "FailedOperation.TsfApmEsClientRequestError"

	// TSF应用性能管理内部异常, 请稍后重试。
	FAILEDOPERATION_TSFAPMINTERNALERROR = "FailedOperation.TsfApmInternalError"

	// TSF应用性能管理运行状态统计查询服务查询失败。
	FAILEDOPERATION_TSFAPMSTATSSEARCHSERVICEQUERYERROR = "FailedOperation.TsfApmStatsSearchServiceQueryError"

	// 数据库插入规则失败。
	FAILEDOPERATION_TSFASDBINSTERFAIL = "FailedOperation.TsfAsDbInsterFail"

	// 查询数据库失败。
	FAILEDOPERATION_TSFASDBQUERYFAIL = "FailedOperation.TsfAsDbQueryFail"

	// 扩容规则每次扩容机器数目不能大于机器最大数目。
	FAILEDOPERATION_TSFASEXPANDCOUNTANDLIMITERROR = "FailedOperation.TsfAsExpandCountAndLimitError"

	// 扩容规则指标不能小于缩容规则。
	FAILEDOPERATION_TSFASEXPANDINDICATORSLESSSHRINK = "FailedOperation.TsfAsExpandIndicatorsLessShrink"

	// 扩容最大实例数不能小于缩容最小实例数。
	FAILEDOPERATION_TSFASEXPANDLIMITLESSSHRINKLIMIT = "FailedOperation.TsfAsExpandLimitLessShrinkLimit"

	// ctsdb数据库请求失败。
	FAILEDOPERATION_TSFCMONITORCTSDBCLIENTREQUESTFAIL = "FailedOperation.TsfCmonitorCtsdbClientRequestFail"

	// TSF监控统计等待超时, 请稍后重试。
	FAILEDOPERATION_TSFMONITORWAITEDTIMEOUT = "FailedOperation.TsfMonitorWaitedTimeout"

	// TSF服务管理通用异常。
	FAILEDOPERATION_TSFMSERROR = "FailedOperation.TsfMsError"

	// 调用ms server失败。
	FAILEDOPERATION_TSFMSSERVERERROR = "FailedOperation.TsfMsServerError"

	// TSF权限模块异常，请联系系统管理员。。
	FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"

	// 未授权。
	FAILEDOPERATION_UNAUTHORIZEDOPERATION = "FailedOperation.UnauthorizedOperation"

	// 模块未处理异常。
	FAILEDOPERATION_UNHANDLEDEXCEPTION = "FailedOperation.UnhandledException"

	// 应用操作请求MASTER FEIGN失败。
	INTERNALERROR_APPLICATIONMASTERFEIGNERROR = "InternalError.ApplicationMasterFeignError"

	// 应用操作请求MASTER 操作失败。
	INTERNALERROR_APPLICATIONMASTERNUKNOWNERROR = "InternalError.ApplicationMasterNuknownError"

	// 删除应用程序包请求仓库失败。
	INTERNALERROR_APPLICATIONREPODELETEPKG = "InternalError.ApplicationRepoDeletePkg"

	// 创建应用初始化tsf-scalable请求失败。
	INTERNALERROR_APPLICATIONSCALABLEINITERROR = "InternalError.ApplicationScalableInitError"

	// TSF云API调用申请角色临时凭证调用请求失败。
	INTERNALERROR_CAMROLEREQUESTERROR = "InternalError.CamRoleRequestError"

	// 配置发布失败：无法连接配置中心服务器。
	INTERNALERROR_CANNOTCONNCONSULSERVER = "InternalError.CanNotConnConsulServer"

	// 调用CDI状态接口返回数据异常
	INTERNALERROR_CDISTATUSDATAINVALID = "InternalError.CdiStatusDataInvalid"

	// TSF云API请求调用失败。
	INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"

	// 集群通用错误。
	INTERNALERROR_CLUSTERCOMMONERROR = "InternalError.ClusterCommonError"

	// 虚拟机集群请求MASTER FEIGN失败。
	INTERNALERROR_CLUSTERMASTERFEIGNERROR = "InternalError.ClusterMasterFeignError"

	// 无法找到部署组，或相应集群/命名空间/应用的权限不足。
	INTERNALERROR_CLUSTERNOTEXISTORPRIVILEGEERROR = "InternalError.ClusterNotExistOrPrivilegeError"

	// 配置发布失败：配置中心服务器处理失败。
	INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"

	// 访问TKE服务失败。
	INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"

	// 连接TKE服务失败。
	INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"

	// Kubernetes deployment 未找到。
	INTERNALERROR_CONTAINERGROUPKUBERNETEDEPLOYMENTNOTFOUND = "InternalError.ContainergroupKuberneteDeploymentNotfound"

	// 容器应用SQL错误。
	INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"

	// 容器平台集群不可用，当前状态 %s。
	INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"

	// 命令下放失败。
	INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"

	// TSF MASTER 内部执行错误。
	INTERNALERROR_CVMCAEMASTERINTERNALERROR = "InternalError.CvmCaeMasterInternalError"

	// MASTER通道查询失败。
	INTERNALERROR_CVMCAEMASTERNONALIVE = "InternalError.CvmCaeMasterNonAlive"

	// 数据查询失败。
	INTERNALERROR_CVMCAEMASTERQUERYERROR = "InternalError.CvmCaeMasterQueryError"

	// TSF暂时不能响应请求。。
	INTERNALERROR_DISPATCHCOMMONERROR = "InternalError.DispatchCommonError"

	// 网关通用异常:%s。
	INTERNALERROR_GATEWAYCOMMONERROR = "InternalError.GatewayCommonError"

	// 数据一致性异常:%s。
	INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"

	// 配置中心访问异常。
	INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"

	// 网关数据异常。
	INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"

	// 部署组通用异常。
	INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"

	// 部署组操作请求MASTER 操作失败。
	INTERNALERROR_GROUPMASTERNUKNOWNERROR = "InternalError.GroupMasterNuknownError"

	// tcr仓库绑定失败。
	INTERNALERROR_IMAGEREPOTCRBINDERROR = "InternalError.ImagerepoTcrBindError"

	// TSF节点管理通用错误信息。
	INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"

	// 创建kubernetes命名空间失败。
	INTERNALERROR_KUBERNETESAPICREATENAMESPACESERROR = "InternalError.KubernetesApiCreateNamespacesError"

	// 创建kubernetes密钥失败。
	INTERNALERROR_KUBERNETESAPICREATESECRETERROR = "InternalError.KubernetesApiCreateSecretError"

	// kubernetes api 调用失败。
	INTERNALERROR_KUBERNETESCALLERROR = "InternalError.KubernetesCallError"

	// 远程调用失败。
	INTERNALERROR_REMOTESERVICECALLERROR = "InternalError.RemoteServiceCallError"

	// 仓库内部错误。
	INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"

	// IN子句中超过1000个候选项。
	INTERNALERROR_SQLTOOMANYINITEM = "InternalError.SqlTooManyInItem"

	// 任务内部异常。
	INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"

	// 调用 TKE 接口失败，%s。
	INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"

	// TSF应用性能管理业务日志配置与应用关联处理错误。
	INTERNALERROR_TSFAPMBUSILOGCFGAPPRELATIONMASTERERROR = "InternalError.TsfApmBusiLogCfgAppRelationMasterError"

	// TSF应用性能管理调用tsf-ms模块失败。
	INTERNALERROR_TSFAPMCALLTSFMSFAILED = "InternalError.TsfApmCallTsfMsFailed"

	// TSF应用性能管理通用异常。
	INTERNALERROR_TSFAPMCOMMONERROR = "InternalError.TsfApmCommonError"

	// TSF应用性能管理ES客户端响应状态异常。
	INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"

	// TSF应用性能管理内部异常, 请稍后重试。
	INTERNALERROR_TSFAPMINTERNALERROR = "InternalError.TsfApmInternalError"

	// TSF监控统计时间日期解析失败。
	INTERNALERROR_TSFMONITORDATEPARSEFAILED = "InternalError.TsfMonitorDateParseFailed"

	// TSF监控统计内部异常, 请稍后重试。
	INTERNALERROR_TSFMONITORINTERNALERROR = "InternalError.TsfMonitorInternalError"

	// TSF服务管理通用异常。
	INTERNALERROR_TSFMSERROR = "InternalError.TsfMsError"

	// [%s]模块未处理异常。。
	INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 删除应用失败。
	INVALIDPARAMETER_APPLICATIONDELETEFAILED = "InvalidParameter.ApplicationDeleteFailed"

	// [%s]模块接口[%s]请求不正确（400 BAD REQUEST）。。
	INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"

	// 配置模板名称不符合规范。
	INVALIDPARAMETER_CONFIGTEMPLATENAMEINVALID = "InvalidParameter.ConfigTemplateNameInvalid"

	// TSF MASTER 解包失败。
	INVALIDPARAMETER_CVMCAEMASTERJSONDECODEFAIL = "InvalidParameter.CvmCaeMasterJsonDecodeFail"

	// TSF MASTER 实例状态异常。
	INVALIDPARAMETER_CVMCAEMASTERUNKNOWNINSTANCESTATUS = "InvalidParameter.CvmCaeMasterUnknownInstanceStatus"

	// 未找到 TCR 实例或命名空间。
	INVALIDPARAMETER_IMAGEREPOTCRNAMESPACENOTFOUND = "InvalidParameter.ImagerepoTcrNamespaceNotFound"

	// 参数错误。
	INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"

	// 已经绑定灰度规则，无法删除。
	INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"

	// 存在同名的泳道。
	INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"

	// 泳道名称格式有误。
	INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"

	// 泳道名称不能为空。
	INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"

	// 泳道名称不能超过60个字符。
	INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"

	// 泳道不存在。
	INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"

	// 泳道没有设置任何入口应用。
	INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"

	// 泳道备注不能超过200个字符。
	INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"

	// 泳道规则中的泳道不存在。
	INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"

	// 存在同名的泳道规则名称。
	INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"

	// 泳道规则名称格式有误。
	INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"

	// 泳道规则名称不能为空。
	INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"

	// 泳道规则名称不能超过60个字符。
	INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"

	// 泳道规则不存在。
	INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"

	// 泳道规则备注不能超过200个字符。
	INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"

	// 泳道规则标签名不能为空。
	INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"

	// 泳道规则标签名不能超过32个字符。
	INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"

	// 泳道规则必须设置至少一个标签。
	INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"

	// 泳道规则标签值不能超过128个字符。
	INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"

	// 泳道规则总标签值不能超过200个字符。
	INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"

	// 包正在被使用，请先解除占用。
	INVALIDPARAMETER_PACKAGEINUSE = "InvalidParameter.PackageInUse"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 请求参数有误。
	INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"

	// 仓库中存在软件包，请先删除软件包。
	INVALIDPARAMETER_REPOSITORYNOTEMPTY = "InvalidParameter.RepositoryNotEmpty"

	// TSF应用性能管理业务日志应用标识参数错误。
	INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppParamError"

	// TSF应用性能管理业务日志配置与应用关联参数错误。
	INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"

	// TSF应用性能管理业务日志配置云账户参数错误。
	INVALIDPARAMETER_TSFAPMBUSILOGCFGCLOUDPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgCloudParamError"

	// TSF应用性能管理业务日志配置标识参数错误。
	INVALIDPARAMETER_TSFAPMBUSILOGCFGIDPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgIdParamError"

	// TSF应用性能管理业务日志配置数目参数错误。
	INVALIDPARAMETER_TSFAPMBUSILOGCFGLIMITPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgLimitParamError"

	// TSF应用性能管理业务日志搜索请求参数错误。
	INVALIDPARAMETER_TSFAPMBUSILOGSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmBusiLogSearchRequestParamError"

	// TSF应用性能管理运行状态统计查询请求参数错误。
	INVALIDPARAMETER_TSFAPMSTATSSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStatsSearchRequestParamError"

	// TSF应用性能管理标准输出日志搜索请求参数错误。
	INVALIDPARAMETER_TSFAPMSTDOUTSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStdoutSearchRequestParamError"

	// TSF应用性能管理调用链搜索请求参数错误。
	INVALIDPARAMETER_TSFAPMTRACESEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmTraceSearchRequestParamError"

	// TSF监控统计请求参数[%s]非法。
	INVALIDPARAMETER_TSFMONITORREQUESTPARAMILLEGAL = "InvalidParameter.TsfMonitorRequestParamIllegal"

	// 仓库批量删除包数量超过单次允许上限。
	INVALIDPARAMETER_UPPERDELETELIMIT = "InvalidParameter.UpperDeleteLimit"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 应用描述不能大于200个字符。
	INVALIDPARAMETERVALUE_APPLICATIONDESCLENGTH = "InvalidParameterValue.ApplicationDescLength"

	// 无效的微服务类型。
	INVALIDPARAMETERVALUE_APPLICATIONMICROTYPEINVALID = "InvalidParameterValue.ApplicationMicroTypeInvalid"

	// 应用名称已存在，请更换其他名称。
	INVALIDPARAMETERVALUE_APPLICATIONNAMEEXIST = "InvalidParameterValue.ApplicationNameExist"

	// 应用名称不能大于60字符。
	INVALIDPARAMETERVALUE_APPLICATIONNAMELENGTH = "InvalidParameterValue.ApplicationNameLength"

	// 应用名称不能为空。
	INVALIDPARAMETERVALUE_APPLICATIONNAMENULL = "InvalidParameterValue.ApplicationNameNull"

	// 应用名称格式不正确,只能包含小写字母、数字及分隔符("_"、"-")，且不能以分隔符开头或结尾。
	INVALIDPARAMETERVALUE_APPLICATIONNAMEREGXINVALID = "InvalidParameterValue.ApplicationNameRegxInvalid"

	// 无法获取应用。
	INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"

	// 无效的应用排序类型。
	INVALIDPARAMETERVALUE_APPLICATIONORDERTYPEINVALID = "InvalidParameterValue.ApplicationOrderTypeInvalid"

	// 无效的应用分页参数。
	INVALIDPARAMETERVALUE_APPLICATIONPAGELIMITINVALID = "InvalidParameterValue.ApplicationPageLimitInvalid"

	// 无效的应用类型。
	INVALIDPARAMETERVALUE_APPLICATIONTYPEINVALID = "InvalidParameterValue.ApplicationTypeInvalid"

	// 与同VPC其它集群CIDR冲突。
	INVALIDPARAMETERVALUE_CLUSTERCIDRCONFLICT = "InvalidParameterValue.ClusterCidrConflict"

	// 集群命名已存在，请更换其他名称。
	INVALIDPARAMETERVALUE_CLUSTERNAMEEXIST = "InvalidParameterValue.ClusterNameExist"

	// 本环境集群名称禁止以 %s 开头
	INVALIDPARAMETERVALUE_CLUSTERNAMEPREFIXILLEGALERROR = "InvalidParameterValue.ClusterNamePrefixIllegalError"

	// 集群命名不能为空。
	INVALIDPARAMETERVALUE_CLUSTERNAMEREQUIRED = "InvalidParameterValue.ClusterNameRequired"

	// 集群无效的分页参数。
	INVALIDPARAMETERVALUE_CLUSTERPAGELIMITINVALID = "InvalidParameterValue.ClusterPageLimitInvalid"

	// 创建集群，无效的地域字段。
	INVALIDPARAMETERVALUE_CLUSTERREGIONINVALID = "InvalidParameterValue.ClusterRegionInvalid"

	// 非法集群类型。
	INVALIDPARAMETERVALUE_CLUSTERTYPEINVALID = "InvalidParameterValue.ClusterTypeInvalid"

	// 创建集群，无效的可用区字段。
	INVALIDPARAMETERVALUE_CLUSTERZONEINVALID = "InvalidParameterValue.ClusterZoneInvalid"

	// 配置项已经发布过。
	INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"

	// 配置项已存在。
	INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"

	// 配置项和部署组所属应用不一致。
	INVALIDPARAMETERVALUE_CONFIGGROUPAPPLICATIONIDNOTMATCH = "InvalidParameterValue.ConfigGroupApplicationIdNotMatch"

	// 配置项名称不合规。
	INVALIDPARAMETERVALUE_CONFIGNAMEINVALID = "InvalidParameterValue.ConfigNameInvalid"

	// 无法获取配置项或无权限访问。
	INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"

	// 无法获取配置项发布信息。
	INVALIDPARAMETERVALUE_CONFIGRELEASENOTEXISTS = "InvalidParameterValue.ConfigReleaseNotExists"

	// 配置模板描述过长。
	INVALIDPARAMETERVALUE_CONFIGTEMPLATEDESCTOOLONG = "InvalidParameterValue.ConfigTemplateDescTooLong"

	// 配置模板名称不合规。
	INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMEINVALID = "InvalidParameterValue.ConfigTemplateNameInvalid"

	// 配置模板内容过长。
	INVALIDPARAMETERVALUE_CONFIGTEMPLATENAMETOOLONG = "InvalidParameterValue.ConfigTemplateNameTooLong"

	// 配置模板类型不合规。
	INVALIDPARAMETERVALUE_CONFIGTEMPLATETYPEINVALID = "InvalidParameterValue.ConfigTemplateTypeInvalid"

	// 配置格式不符合YAML要求。
	INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"

	// 配置项值内容大小（%s）超过限制。
	INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"

	// 配置项版本描述不合规。
	INVALIDPARAMETERVALUE_CONFIGVERSIONDESCINVALID = "InvalidParameterValue.ConfigVersionDescInvalid"

	// 配置项版本不合规。
	INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"

	// 该镜像被占用中。
	INVALIDPARAMETERVALUE_CONTAINERGROUPIMAGETAGISINUSE = "InvalidParameterValue.ContainerGroupImageTagIsInUse"

	// 服务访问方式不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"

	// 所属应用ID不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"

	// 集群 CPU 资源不足。
	INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"

	// 容器 Env 的 Value 和 ValueFrom 至少要有一个。
	INVALIDPARAMETERVALUE_CONTAINERGROUPENVVALUENOTSET = "InvalidParameterValue.ContainergroupEnvValueNotSet"

	// 部署组ID不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"

	// 部署组名不能大于60个字符。
	INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMELEGNTH = "InvalidParameterValue.ContainergroupGroupnameLegnth"

	// 部署组名不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMENULL = "InvalidParameterValue.ContainergroupGroupnameNull"

	// 部署组名称格式不正确,只能包含小写字母、数字及分隔符("-"),且必须以小写字母开头，数字或小写字母结尾。
	INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMEREGEXMATCHFALSE = "InvalidParameterValue.ContainergroupGroupnameRegexMatchFalse"

	// 实例数量不能为空或不合法。
	INVALIDPARAMETERVALUE_CONTAINERGROUPINSTANCENUMINVALID = "InvalidParameterValue.ContainergroupInstanceNumInvalid"

	// CPU limit 和 request 不能同时为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"

	// 内存 limit 和 request 不能同时为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"

	// limit最大数量，默认 20, 最大值 50。
	INVALIDPARAMETERVALUE_CONTAINERGROUPLIMITVALUEINVALID = "InvalidParameterValue.ContainergroupLimitValueInvalid"

	// 集群内存资源不足。
	INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"

	// 主机端口值非法。
	INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"

	// 服务端口值非法。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPORTINVALID = "InvalidParameterValue.ContainergroupPortInvalid"

	// 服务端口不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPORTNULL = "InvalidParameterValue.ContainergroupPortNull"

	// 服务端口不允许重复映射。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPORTSREPEAT = "InvalidParameterValue.ContainergroupPortsRepeat"

	// 协议值非法,限定:TCP/UDP。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLINVALID = "InvalidParameterValue.ContainergroupProtocolInvalid"

	// 公网访问方式下，协议需要一致。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLMIXERROR = "InvalidParameterValue.ContainergroupProtocolMixError"

	// 协议不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLNULL = "InvalidParameterValue.ContainergroupProtocolNull"

	// 协议端口不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"

	// 镜像仓库名与应用名不匹配。
	INVALIDPARAMETERVALUE_CONTAINERGROUPREPONAMEINVALID = "InvalidParameterValue.ContainergroupReponameInvalid"

	// agent 容器资源值非法 , %s。
	INVALIDPARAMETERVALUE_CONTAINERGROUPRESOURCEAGENTVALUEINVALID = "InvalidParameterValue.ContainergroupResourceAgentValueInvalid"

	// 容器端口不允许重复映射。
	INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTSREPEAT = "InvalidParameterValue.ContainergroupTargetPortsRepeat"

	// 容器端口不能为空。
	INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"

	// 更新间隔不能为空或者数值非法。
	INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"

	// updateType参数不合法，值必须为0、1。
	INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATETYPEINVALID = "InvalidParameterValue.ContainergroupUpdatetypeInvalid"

	// 找不到业务容器。
	INVALIDPARAMETERVALUE_CONTAINERGROUPYAMLUSERCONTAINERNOTFOUND = "InvalidParameterValue.ContainergroupYamlUserContainerNotFound"

	// TSF MASTER 正在执行任务，请等待任务执行完成再下发新任务。
	INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"

	// 无可用实例。
	INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTNOTFOUND = "InvalidParameterValue.CvmCaeMasterAgentNotFound"

	// TSF MASTER 部署组中无云主机。
	INVALIDPARAMETERVALUE_CVMCAEMASTERGROUPNOAGENT = "InvalidParameterValue.CvmCaeMasterGroupNoAgent"

	// 任务不存在。
	INVALIDPARAMETERVALUE_CVMCAEMASTERTASKNOTEXIST = "InvalidParameterValue.CvmCaeMasterTaskNotExist"

	// 部署组不存在。
	INVALIDPARAMETERVALUE_DEPLOYGROUPNOTEXISTS = "InvalidParameterValue.DeployGroupNotExists"

	// 重复数据集名。
	INVALIDPARAMETERVALUE_DUPLICATEPROGRAMNAME = "InvalidParameterValue.DuplicateProgramName"

	// 文件配置项已经发布。
	INVALIDPARAMETERVALUE_FILECONFIGALREADYRELEASED = "InvalidParameterValue.FileConfigAlreadyReleased"

	// 文件配置项编码方式不支持。
	INVALIDPARAMETERVALUE_FILECONFIGCODEUNSUPPORTED = "InvalidParameterValue.FileConfigCodeUnsupported"

	// 文件配置项已存在。
	INVALIDPARAMETERVALUE_FILECONFIGEXISTS = "InvalidParameterValue.FileConfigExists"

	// 配置文件路径重复。
	INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATH = "InvalidParameterValue.FileConfigExistsPath"

	// 其他用户已发布此配置文件路径。
	INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATHOTHER = "InvalidParameterValue.FileConfigExistsPathOther"

	// 文件配置项文件名称不合规。
	INVALIDPARAMETERVALUE_FILECONFIGFILENAMEINVALID = "InvalidParameterValue.FileConfigFileNameInvalid"

	// 文件配置项文件路径不合规。
	INVALIDPARAMETERVALUE_FILECONFIGFILEPATHINVALID = "InvalidParameterValue.FileConfigFilePathInvalid"

	// 文件配置项固定字段不可变更。
	INVALIDPARAMETERVALUE_FILECONFIGFIXEDFIELDUNCHANGEABLE = "InvalidParameterValue.FileConfigFixedFieldUnchangeable"

	// 文件配置项名称不合规。
	INVALIDPARAMETERVALUE_FILECONFIGNAMEINVALID = "InvalidParameterValue.FileConfigNameInvalid"

	// 无法获取文件配置项或无权限访问。
	INVALIDPARAMETERVALUE_FILECONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.FileConfigNotExistsOrPermissionDenied"

	// 同一部署组禁止配置文件重复(文件路径+文件名)。
	INVALIDPARAMETERVALUE_FILECONFIGPATHEXISTS = "InvalidParameterValue.FileConfigPathExists"

	// 无法获取文件配置项发布信息。
	INVALIDPARAMETERVALUE_FILECONFIGRELEASENOTEXISTS = "InvalidParameterValue.FileConfigReleaseNotExists"

	// 文件配置项版本描述不合规。
	INVALIDPARAMETERVALUE_FILECONFIGVERSIONDESCINVALID = "InvalidParameterValue.FileConfigVersionDescInvalid"

	// 文件配置项版本不合规。
	INVALIDPARAMETERVALUE_FILECONFIGVERSIONINVALID = "InvalidParameterValue.FileConfigVersionInvalid"

	// 请求参数异常:%s。
	INVALIDPARAMETERVALUE_GATEWAYPARAMETERERROR = "InvalidParameterValue.GatewayParameterError"

	// 无效请求参数:%s。
	INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"

	// 全局命名空间已经存在，只能创建一个全局命名空间。
	INVALIDPARAMETERVALUE_GLOBALNAMESPACENAMEEXIST = "InvalidParameterValue.GlobalNamespaceNameExist"

	// 部署相关请求参数校验失败。
	INVALIDPARAMETERVALUE_GROUPBATCHPARAMETERINVALID = "InvalidParameterValue.GroupBatchParameterInvalid"

	// 部署组的集群未绑定该命名空间。
	INVALIDPARAMETERVALUE_GROUPCLUSTERNAMESPACENOTBOUND = "InvalidParameterValue.GroupClusterNamespaceNotBound"

	// 创建分组， 集群类型不匹配。
	INVALIDPARAMETERVALUE_GROUPCLUSTERTYPEMISMATCH = "InvalidParameterValue.GroupClusterTypeMismatch"

	// 删除分组，集群类型不匹配。
	INVALIDPARAMETERVALUE_GROUPDELETECLUSTERTYPEMISMATCH = "InvalidParameterValue.GroupDeleteClusterTypeMismatch"

	// 部署组ID不能为空。
	INVALIDPARAMETERVALUE_GROUPIDNULL = "InvalidParameterValue.GroupIdNull"

	// 部署组名称已存在，请更换其他名称。
	INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"

	// 部署组名不能大于50个字符。
	INVALIDPARAMETERVALUE_GROUPNAMELENGTH = "InvalidParameterValue.GroupNameLength"

	// 部署组名称格式不正确,只能包含小写字母、数字及分隔符("-"),且必须以小写字母开头，数字或小写字母结尾。
	INVALIDPARAMETERVALUE_GROUPNAMEREGXMISMATCH = "InvalidParameterValue.GroupNameRegxMismatch"

	// 无法获取部署组。
	INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"

	// 分组无效的分业参数。
	INVALIDPARAMETERVALUE_GROUPPAGELIMITINVALID = "InvalidParameterValue.GroupPageLimitInvalid"

	// 无效的部署组状态过滤字段。
	INVALIDPARAMETERVALUE_GROUPSTATUSINVALID = "InvalidParameterValue.GroupStatusInvalid"

	// 分组操作，无有效机器。
	INVALIDPARAMETERVALUE_GROUPVALIDINSTANCENULL = "InvalidParameterValue.GroupValidInstanceNull"

	// 镜像仓库名不能为空。
	INVALIDPARAMETERVALUE_IMAGEREPOREPONAMENULL = "InvalidParameterValue.ImagerepoRepoNameNull"

	// 镜像仓库名不合法,示例:tsf-repo/nginx。
	INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"

	// imageTags不能为空。
	INVALIDPARAMETERVALUE_IMAGEREPOTAGNAMENULL = "InvalidParameterValue.ImagerepoTagnameNull"

	// 重装系统，无效的镜像id。
	INVALIDPARAMETERVALUE_INSTANCEINVALIDIMAGE = "InvalidParameterValue.InstanceInvalidImage"

	// 参数 %s 取值错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"

	// 参数格式异常。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"

	// 数据集名非法。
	INVALIDPARAMETERVALUE_INVALIDPROGRAMNAME = "InvalidParameterValue.InvalidProgramName"

	// 该泳道已关联全链路灰度发布规则,请先从规则中移除后删除
	INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"

	// 存在同名的泳道。
	INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"

	// 泳道名称格式有误。
	INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"

	// 泳道名称不能为空。
	INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"

	// 泳道名称不能超过60个字符。
	INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"

	// 泳道不存在。
	INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"

	// 泳道没有设置任何入口应用。
	INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"

	// 泳道备注不能超过200个字符。
	INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"

	// 全链路灰度规则ID错误
	INVALIDPARAMETERVALUE_LANERULEIDINVALID = "InvalidParameterValue.LaneRuleIdInvalid"

	// 全链路灰度规则中的泳道不存在。
	INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"

	// 存在同名的全链路灰度规则。
	INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"

	// 全链路灰度规则名称格式有误。
	INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"

	// 全链路灰度规则名称不能为空。
	INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"

	// 全链路灰度规则名称不能超过60个字符。
	INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"

	// 全链路灰度规则不存在。
	INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"

	// 全链路灰度规则备注不能超过200个字符。
	INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"

	// 全链路灰度规则标签名不能为空。
	INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"

	// 全链路灰度规则标签名不能超过32个字符。
	INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"

	// 全链路灰度规则必须设置至少一个标签。
	INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"

	// 全链路灰度规则逻辑关系错误
	INVALIDPARAMETERVALUE_LANERULETAGOPERATORINVALID = "InvalidParameterValue.LaneRuleTagOperatorInvalid"

	// 全链路灰度规则逻辑关系不能为空
	INVALIDPARAMETERVALUE_LANERULETAGOPERATORNOTEMPTY = "InvalidParameterValue.LaneRuleTagOperatorNotEmpty"

	// 全链路灰度规则标签值不能为空
	INVALIDPARAMETERVALUE_LANERULETAGVALUENOTEMPTY = "InvalidParameterValue.LaneRuleTagValueNotEmpty"

	// 全链路灰度规则标签值不能超过128个字符。
	INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"

	// 全链路灰度规则总标签值不能超过200个字符。
	INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"

	// 集群已关联该命名空间。
	INVALIDPARAMETERVALUE_NAMESPACEALREADYBINDCLUSTER = "InvalidParameterValue.NamespaceAlreadyBindCluster"

	// 命名空间描述格式不正确。
	INVALIDPARAMETERVALUE_NAMESPACEDESCINVALID = "InvalidParameterValue.NamespaceDescInvalid"

	// 命名空间名称已存在，请更换其他名称。
	INVALIDPARAMETERVALUE_NAMESPACENAMEEXIST = "InvalidParameterValue.NamespaceNameExist"

	// 命名空间名称格式不正确。
	INVALIDPARAMETERVALUE_NAMESPACENAMEINVALID = "InvalidParameterValue.NamespaceNameInvalid"

	// 无法获取命名空间。
	INVALIDPARAMETERVALUE_NAMESPACENOTEXISTS = "InvalidParameterValue.NamespaceNotExists"

	// 数据集项不存在。
	INVALIDPARAMETERVALUE_PROGRAMITEMNOTEXISTS = "InvalidParameterValue.ProgramItemNotExists"

	// 配置项已经发布，不允许删除。
	INVALIDPARAMETERVALUE_RELEASEDCONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedConfigCanNotBeDeleted"

	// 文件配置项已经发布，不允许删除。
	INVALIDPARAMETERVALUE_RELEASEDFILECONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedFileConfigCanNotBeDeleted"

	// 无权限操作资源%s。
	INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"

	// ResourceType 不支持。
	INVALIDPARAMETERVALUE_RESOURCETYPEERROR = "InvalidParameterValue.ResourceTypeError"

	// 服务描述不能大于200字符。
	INVALIDPARAMETERVALUE_SERVICEDESCLENGTH = "InvalidParameterValue.ServiceDescLength"

	// 服务名称重复。
	INVALIDPARAMETERVALUE_SERVICENAMEREPEATED = "InvalidParameterValue.ServiceNameRepeated"

	// 服务不存在或权限不足。
	INVALIDPARAMETERVALUE_SERVICENOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ServiceNotExistsOrPermissionDenied"

	// 无效请求参数。
	INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"

	// TSF应用性能管理业务日志搜索解析时间格式错误。
	INVALIDPARAMETERVALUE_TSFAPMBUSILOGSEARCHPARSERSPDATEFORMATERROR = "InvalidParameterValue.TsfApmBusiLogSearchParseRspDateFormatError"

	// 仅有停止状态下的部署组才可以不启动。
	INVALIDPARAMETERVALUE_WRONGDONTSTARTVALUE = "InvalidParameterValue.WrongDontStartValue"

	// 命名空间数达到上限。
	LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"

	// 仓库达到上限。
	LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"

	// 最多支持创建五个容器集群，当前已经超过使用上限。
	LIMITEXCEEDED_TKECLUSTERNUMBEREXCEEDLIMIT = "LimitExceeded.TkeClusterNumberExceedLimit"

	// 应用ID不能为空。
	MISSINGPARAMETER_APPLICATIONIDNULL = "MissingParameter.ApplicationIdNull"

	// 应用ID未填写。
	MISSINGPARAMETER_APPLICATIONIDREQUIRED = "MissingParameter.ApplicationIdRequired"

	// 应用类型不能为空。
	MISSINGPARAMETER_APPLICATIONTYPENULL = "MissingParameter.ApplicationTypeNull"

	// 集群ID未填写。
	MISSINGPARAMETER_CLUSTERIDREQUIRED = "MissingParameter.ClusterIdRequired"

	// 集群所属子网不能为空。
	MISSINGPARAMETER_CLUSTERSUBNETREQUIRED = "MissingParameter.ClusterSubnetRequired"

	// 配置项ID未填写。
	MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"

	// 配置项名称未填写。
	MISSINGPARAMETER_CONFIGNAMEREQUIRED = "MissingParameter.ConfigNameRequired"

	// 配置项发布信息ID未填写。
	MISSINGPARAMETER_CONFIGRELEASEIDREQUIRED = "MissingParameter.ConfigReleaseIdRequired"

	// 缺少配置模板id。
	MISSINGPARAMETER_CONFIGTEMPLATEIDREQUIRED = "MissingParameter.ConfigTemplateIdRequired"

	// 配置模板名称未填写。
	MISSINGPARAMETER_CONFIGTEMPLATENAMEREQUIRED = "MissingParameter.ConfigTemplateNameRequired"

	// 配置模板类型未填写。
	MISSINGPARAMETER_CONFIGTEMPLATETYPEREQUIRED = "MissingParameter.ConfigTemplateTypeRequired"

	// 配置项类型未填写。
	MISSINGPARAMETER_CONFIGTYPEREQUIRED = "MissingParameter.ConfigTypeRequired"

	// 配置项值未填写。
	MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"

	// 配置项版本未填写。
	MISSINGPARAMETER_CONFIGVERSIONREQUIRED = "MissingParameter.ConfigVersionRequired"

	// 文件配置项文件内容未填写。
	MISSINGPARAMETER_FILECONFIGFILEVALUEREQUIRED = "MissingParameter.FileConfigFileValueRequired"

	// 缺少请求参数:%s。
	MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"

	// 分组所属应用不能为空。
	MISSINGPARAMETER_GROUPAPPLICATIONNULL = "MissingParameter.GroupApplicationNull"

	// 分组扩容操作，机器列表为空。
	MISSINGPARAMETER_GROUPEXPANDSERVERIDNULL = "MissingParameter.GroupExpandServeridNull"

	// 分组ID不能为空。
	MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"

	// 部署组ID未填写。
	MISSINGPARAMETER_GROUPIDREQUIRED = "MissingParameter.GroupIdRequired"

	// 分组所属命名空间不能为空。
	MISSINGPARAMETER_GROUPNAMESPACENULL = "MissingParameter.GroupNamespaceNull"

	// 分组缩容操作，机器列表为空。
	MISSINGPARAMETER_GROUPSHIRKSERVERIDNULL = "MissingParameter.GroupShirkServeridNull"

	// 虚拟机集群导入云主机导入方式为空。
	MISSINGPARAMETER_INSTANCEIMPORTMODENULL = "MissingParameter.InstanceImportModeNull"

	// 命名空间ID不能为空。
	MISSINGPARAMETER_NAMESPACEIDREQUIRED = "MissingParameter.NamespaceIdRequired"

	// 命名空间名称未填写。
	MISSINGPARAMETER_NAMESPACENAMEREQUIRED = "MissingParameter.NamespaceNameRequired"

	// %s缺失。
	MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"

	// 未填写服务Id。
	MISSINGPARAMETER_SERVICEIDREQUIRED = "MissingParameter.ServiceIdRequired"

	// 未指定%s。。
	MISSINGPARAMETER_SYSTEMPARAMETERREQUIRED = "MissingParameter.SystemParameterRequired"

	// 缺少必填参数。
	MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"

	// 此应用下存在资源，无法执行删除操作。
	RESOURCEINUSE_APPLICATIONCANNOTDELETE = "ResourceInUse.ApplicationCannotDelete"

	// 资源仍在使用中 无法删除。
	RESOURCEINUSE_CVMCAEMASTERCANNOTDELETE = "ResourceInUse.CvmcaeMasterCannotDelete"

	// 默认命名空间不能被删除。
	RESOURCEINUSE_DEFAULTNAMEPSACECANNOTBEDELETED = "ResourceInUse.DefaultNamepsaceCannotBeDeleted"

	// 此分组下存在资源，无法执行删除操作。
	RESOURCEINUSE_GROUPCANNOTDELETE = "ResourceInUse.GroupCannotDelete"

	// 集群下存在分组。
	RESOURCEINUSE_GROUPEXISTS = "ResourceInUse.GroupExists"

	// 部署组在更新中 请稍后再执行该操作。
	RESOURCEINUSE_GROUPINOPERATION = "ResourceInUse.GroupInOperation"

	// 集群内有云主机,不允许删除集群。
	RESOURCEINUSE_INSTANCEEXISTS = "ResourceInUse.InstanceExists"

	// 机器实例已经被使用。
	RESOURCEINUSE_INSTANCEHASBEENUSED = "ResourceInUse.InstanceHasBeenUsed"

	// 此命名空间下存在资源，无法执行删除操作。
	RESOURCEINUSE_NAMESPACECANNOTDELETE = "ResourceInUse.NamespaceCannotDelete"

	// 集群下存在非默认命名空间。
	RESOURCEINUSE_NONDEFAULTNAMESPACEEXISTS = "ResourceInUse.NonDefaultNamespaceExists"

	// 资源对象已存在。
	RESOURCEINUSE_OBJECTEXIST = "ResourceInUse.ObjectExist"

	// 限流规则已存在，请检查规则名和规则配置。
	RESOURCEINUSE_RATELIMITRULEEXISTERROR = "ResourceInUse.RatelimitRuleExistError"

	// 仓库空间达到上限。
	RESOURCEINSUFFICIENT_PACKAGESPACEFULL = "ResourceInsufficient.PackageSpaceFull"

	// 无法获取应用信息。
	RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"

	// 无法获取应用或应用不属于当前项目。
	RESOURCENOTFOUND_APPLICATIONPROJECTNOTMATCH = "ResourceNotFound.ApplicationProjectNotMatch"

	// 无法获取命名空间所属集群。
	RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"

	// 集群所属私有网络不存在。
	RESOURCENOTFOUND_CLUSTERVPCNOTEXIST = "ResourceNotFound.ClusterVpcNotExist"

	// 找不到集群。
	RESOURCENOTFOUND_CONTAINERGROUPCLUSTERNOTFOUND = "ResourceNotFound.ContainergroupClusterNotfound"

	// 无法找到该部署组所属集群和命名空间。
	RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"

	// 无法找到该部署组。
	RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"

	// TSF MASTER 资源不存在。
	RESOURCENOTFOUND_CVMCAEMASTERRESOURCENOTFOUND = "ResourceNotFound.CvmcaeMasterResourceNotFound"

	// 镜像仓库不存在。
	RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"

	// 用户错误。
	RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"

	// 无法获取分组所属应用。
	RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"

	// 无法获取分组所属命名空间。
	RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"

	// 此部署组不存在，无法执行该操作。
	RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"

	// 无法获取机器信息。
	RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"

	// [%s]模块未提供该接口[%s]。。
	RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"

	// 无法找到License服务器。
	RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"

	// 目标微服务已离线[%s]。。
	RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"

	// 无法获取命名空间。
	RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"

	// 资源对象不存在。
	RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"

	// 无法获取服务，无法执行该操作。
	RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"

	// 任务不存在。
	RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"

	// TKE 中不存在该集群。
	RESOURCENOTFOUND_TKECLUSTERNOTEXISTS = "ResourceNotFound.TkeClusterNotExists"

	// 访问 CAM 系统出错，%s。
	UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"

	// 协作者身份未授权，需要主账号授予协作者权限，参考 TSF 官网文档「快速入门/准备工作」。
	UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"

	// 当前主账号未创建TSF_QCSRole或未对子账号授予预设策略QcloudCamSubaccountsAuthorizeRoleFullAccess。请参考产品文档主账号协作者使用说明。。
	UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"

	// License未激活。。
	UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"

	// 您所购买的服务不支持该操作。
	UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"

	// 缺少License。。
	UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"

	// 用户无权限访问该接口。
	UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"

	// 批量操作数量超过限制:%s。
	UNSUPPORTEDOPERATION_GATEWAYTOOMANYREQUESTPARAMETER = "UnsupportedOperation.GatewayTooManyRequestParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION_TASKNOTSUPPORTED = "UnsupportedOperation.TaskNotSupported"

	// 不支持的ACTION。。
	UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
)
