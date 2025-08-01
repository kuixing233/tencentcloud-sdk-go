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

package integration
import (
    "os"
    "testing"

    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	aaiv20180522 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aai/v20180522"
	acav20210323 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aca/v20210323"
	acpv20220105 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/acp/v20220105"
	advisorv20200721 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/advisor/v20200721"
	afv20200226 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/af/v20200226"
	afcv20200226 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/afc/v20200226"
	ai3dv20250513 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ai3d/v20250513"
	aiartv20221229 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aiart/v20221229"
	amev20190916 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ame/v20190916"
	amsv20200608 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ams/v20200608"
	amsv20201229 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ams/v20201229"
	anicloudv20220923 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/anicloud/v20220923"
	antiddosv20200309 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/antiddos/v20200309"
	apev20200513 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ape/v20200513"
	apiv20201106 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/api/v20201106"
	apigatewayv20180808 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apigateway/v20180808"
	apmv20210622 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apm/v20210622"
	asv20180419 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/as/v20180419"
	asrv20190614 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
	aswv20200722 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asw/v20200722"
	bav20200720 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ba/v20200720"
	batchv20170312 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/batch/v20170312"
	bdav20200324 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bda/v20200324"
	bhv20230418 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bh/v20230418"
	biv20220105 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bi/v20220105"
	billingv20180709 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	bizlivev20190313 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bizlive/v20190313"
	bmv20180423 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bm/v20180423"
	bmav20210624 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bma/v20210624"
	bmav20221115 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bma/v20221115"
	bmeipv20180625 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmeip/v20180625"
	bmlbv20180625 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmlb/v20180625"
	bmvpcv20180625 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmvpc/v20180625"
	bpaasv20181217 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bpaas/v20181217"
	briv20190328 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bri/v20190328"
	bscav20210811 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bsca/v20210811"
	btoev20210303 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/btoe/v20210303"
	btoev20210514 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/btoe/v20210514"
	cav20230228 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ca/v20230228"
	camv20190116 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
	captchav20190722 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha/v20190722"
	carv20220110 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/car/v20220110"
	catv20180409 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cat/v20180409"
	cbsv20170312 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
	cccv20200210 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ccc/v20200210"
	cdbv20170320 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	cdcv20201214 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdc/v20201214"
	cdnv20180606 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	cdsv20180420 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cds/v20180420"
	cdwchv20200915 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdwch/v20200915"
	cdwdorisv20211228 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdwdoris/v20211228"
	cdwpgv20201230 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdwpg/v20201230"
	cdzv20221123 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdz/v20221123"
	cfgv20210820 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfg/v20210820"
	cfsv20190719 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfs/v20190719"
	cfwv20190904 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfw/v20190904"
	chcv20230418 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/chc/v20230418"
	chdfsv20190718 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/chdfs/v20190718"
	chdfsv20201112 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/chdfs/v20201112"
	ciamv20220331 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ciam/v20220331"
	ciiv20201210 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cii/v20201210"
	ciiv20210408 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cii/v20210408"
	cimv20190318 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cim/v20190318"
	ckafkav20190819 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ckafka/v20190819"
	clbv20180317 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
	cloudappv20220530 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudapp/v20220530"
	cloudauditv20190319 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudaudit/v20190319"
	cloudhsmv20191112 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudhsm/v20191112"
	cloudstudiov20230508 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudstudio/v20230508"
	clsv20201016 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cls/v20201016"
	cmev20191029 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cme/v20191029"
	cmqv20190304 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cmq/v20190304"
	cmsv20190321 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cms/v20190321"
	configv20220802 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/config/v20220802"
	controlcenterv20230110 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/controlcenter/v20230110"
	cpdpv20190820 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cpdp/v20190820"
	csipv20221121 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/csip/v20221121"
	csxgv20230303 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/csxg/v20230303"
	ctemv20231128 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ctem/v20231128"
	ctsdbv20230202 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ctsdb/v20230202"
	cvmv20170312 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	cwpv20180228 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cwp/v20180228"
	cwsv20180312 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cws/v20180312"
	cynosdbv20190107 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cynosdb/v20190107"
	dasbv20191018 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dasb/v20191018"
	dayuv20180709 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dayu/v20180709"
	dbbrainv20191016 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbbrain/v20191016"
	dbbrainv20210527 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbbrain/v20210527"
	dbdcv20201029 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbdc/v20201029"
	dcv20180410 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dc/v20180410"
	dcdbv20180411 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dcdb/v20180411"
	dlcv20210125 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dlc/v20210125"
	dnspodv20210323 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	domainv20180808 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/domain/v20180808"
	drmv20181115 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/drm/v20181115"
	dsv20180523 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ds/v20180523"
	dsgcv20190723 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dsgc/v20190723"
	dtsv20180330 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dts/v20180330"
	dtsv20211206 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dts/v20211206"
	ebv20210416 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eb/v20210416"
	eccv20181213 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecc/v20181213"
	ecdnv20191012 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecdn/v20191012"
	ecmv20190719 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecm/v20190719"
	eiamv20210420 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eiam/v20210420"
	eisv20200715 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eis/v20200715"
	eisv20210601 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eis/v20210601"
	emrv20190103 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/emr/v20190103"
	esv20180416 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/es/v20180416"
	esv20250101 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/es/v20250101"
	essv20201111 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	essbasicv20201222 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20201222"
	essbasicv20210526 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
	facefusionv20181201 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/facefusion/v20181201"
	facefusionv20220927 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/facefusion/v20220927"
	faceidv20180301 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid/v20180301"
	fmuv20191213 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/fmu/v20191213"
	ftv20200304 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ft/v20200304"
	gaapv20180529 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap/v20180529"
	gmev20180711 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gme/v20180711"
	goosefsv20220519 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/goosefs/v20220519"
	gsv20191118 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gs/v20191118"
	gwlbv20240906 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gwlb/v20240906"
	habov20181203 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/habo/v20181203"
	haiv20230812 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hai/v20230812"
	hasimv20210716 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hasim/v20210716"
	hcmv20181106 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hcm/v20181106"
	hunyuanv20230901 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hunyuan/v20230901"
	iaiv20180301 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai/v20180301"
	iaiv20200303 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai/v20200303"
	iapv20240713 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iap/v20240713"
	icv20190307 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ic/v20190307"
	icrv20211014 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/icr/v20211014"
	iev20200304 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ie/v20200304"
	igv20210518 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ig/v20210518"
	igtmv20231024 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/igtm/v20231024"
	imsv20200713 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ims/v20200713"
	imsv20201229 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ims/v20201229"
	ioav20220601 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ioa/v20220601"
	iotv20180123 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iot/v20180123"
	iotcloudv20180614 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotcloud/v20180614"
	iotcloudv20210408 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotcloud/v20210408"
	iotexplorerv20190423 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotexplorer/v20190423"
	iotvideov20191126 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotvideo/v20191126"
	iotvideov20201215 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotvideo/v20201215"
	iotvideov20211125 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotvideo/v20211125"
	iotvideoindustryv20201201 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotvideoindustry/v20201201"
	irpv20220324 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/irp/v20220324"
	irpv20220805 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/irp/v20220805"
	issv20230517 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iss/v20230517"
	ivldv20210903 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ivld/v20210903"
	keewidbv20220308 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/keewidb/v20220308"
	kmsv20190118 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/kms/v20190118"
	lcicv20220817 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lcic/v20220817"
	lighthousev20200324 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
	livev20180801 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
	lkev20231130 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lke/v20231130"
	lkeapv20240522 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lkeap/v20240522"
	lowcodev20210108 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lowcode/v20210108"
	mallv20230518 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mall/v20230518"
	mariadbv20170312 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mariadb/v20170312"
	marketv20191010 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/market/v20191010"
	memcachedv20190318 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/memcached/v20190318"
	mmpsv20200710 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mmps/v20200710"
	mnav20210119 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mna/v20210119"
	mongodbv20180408 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb/v20180408"
	mongodbv20190725 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb/v20190725"
	monitorv20180724 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	monitorv20230616 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20230616"
	mpsv20190612 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mps/v20190612"
	mqttv20240516 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mqtt/v20240516"
	mrsv20200910 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mrs/v20200910"
	msv20180408 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ms/v20180408"
	mspv20180319 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/msp/v20180319"
	nlpv20190408 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/nlp/v20190408"
	nppv20190823 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/npp/v20190823"
	oceanusv20190422 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/oceanus/v20190422"
	ocrv20181119 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	omicsv20221128 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/omics/v20221128"
	organizationv20181225 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/organization/v20181225"
	organizationv20210331 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/organization/v20210331"
	partnersv20180321 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/partners/v20180321"
	postgresv20170312 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/postgres/v20170312"
	privatednsv20201028 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"
	ptsv20210728 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/pts/v20210728"
	rcev20201103 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rce/v20201103"
	redisv20180412 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"
	regionv20220627 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/region/v20220627"
	rumv20210622 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rum/v20210622"
	scfv20180416 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf/v20180416"
	securitylakev20240117 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/securitylake/v20240117"
	sesv20201002 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses/v20201002"
	smhv20210712 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/smh/v20210712"
	smopv20201203 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/smop/v20201203"
	smsv20190711 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
	smsv20210111 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	soev20180724 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/soe/v20180724"
	sqlserverv20180328 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sqlserver/v20180328"
	ssav20180608 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssa/v20180608"
	sslv20191205 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
	sslpodv20190605 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sslpod/v20190605"
	ssmv20190923 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm/v20190923"
	stsv20180813 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
	svpv20240125 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/svp/v20240125"
	tafv20200210 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/taf/v20200210"
	tagv20180813 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag/v20180813"
	tatv20201028 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tat/v20201028"
	tbaasv20180416 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbaas/v20180416"
	tbpv20190311 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbp/v20190311"
	tbpv20190627 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbp/v20190627"
	tcaplusdbv20190823 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcaplusdb/v20190823"
	tcbv20180608 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcb/v20180608"
	tcbrv20220217 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcbr/v20220217"
	tccatalogv20241024 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tccatalog/v20241024"
	tchdv20230306 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tchd/v20230306"
	tcmv20210413 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcm/v20210413"
	tcrv20190924 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcr/v20190924"
	tcssv20201101 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcss/v20201101"
	tdcpgv20211118 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdcpg/v20211118"
	tdidv20210519 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdid/v20210519"
	tdmqv20200217 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdmq/v20200217"
	tdsv20220801 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tds/v20220801"
	temv20201221 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tem/v20201221"
	temv20210701 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tem/v20210701"
	teov20220106 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220106"
	teov20220901 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220901"
	thpcv20211109 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/thpc/v20211109"
	thpcv20220401 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/thpc/v20220401"
	thpcv20230321 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/thpc/v20230321"
	tiav20180226 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tia/v20180226"
	tiiav20190529 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiia/v20190529"
	tionev20191022 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tione/v20191022"
	tionev20211111 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tione/v20211111"
	tiwv20190919 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiw/v20190919"
	tkev20180525 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"
	tkev20220501 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20220501"
	tkgdqv20190411 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tkgdq/v20190411"
	tmsv20200713 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tms/v20200713"
	tmsv20201229 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tms/v20201229"
	tmtv20180321 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
	tourismv20230215 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tourism/v20230215"
	trabbitv20230418 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trabbit/v20230418"
	trocketv20230308 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trocket/v20230308"
	trpv20210515 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trp/v20210515"
	trrov20220325 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trro/v20220325"
	trtcv20190722 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc/v20190722"
	tsev20201207 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tse/v20201207"
	tsfv20180326 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsf/v20180326"
	tsiv20210325 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsi/v20210325"
	tswv20200924 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsw/v20200924"
	tswv20210412 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsw/v20210412"
	ttsv20190823 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tts/v20190823"
	vcgv20240404 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vcg/v20240404"
	vclmv20240523 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vclm/v20240523"
	vcubev20220410 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vcube/v20220410"
	vdbv20230616 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vdb/v20230616"
	vmv20200709 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vm/v20200709"
	vmv20201229 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vm/v20201229"
	vmv20210922 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vm/v20210922"
	vmsv20200902 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vms/v20200902"
	vodv20180717 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
	vodv20240718 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20240718"
	vpcv20170312 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
	vrsv20200824 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vrs/v20200824"
	vtcv20240223 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vtc/v20240223"
	wafv20180125 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/waf/v20180125"
	wavv20210129 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wav/v20210129"
	wedatav20210820 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wedata/v20210820"
	weilingwithv20230427 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/weilingwith/v20230427"
	wsav20250508 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wsa/v20250508"
	wssv20180426 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wss/v20180426"
	yinsudav20220527 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yinsuda/v20220527"
	yunjingv20180228 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yunjing/v20180228"
	yunsouv20180504 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yunsou/v20180504"
	yunsouv20191115 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yunsou/v20191115"
)

func TestAaiv20180522Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := aaiv20180522.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init aai_v20180522 client: %v", err)
    }
}

func TestAcav20210323Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := acav20210323.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init aca_v20210323 client: %v", err)
    }
}

func TestAcpv20220105Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := acpv20220105.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init acp_v20220105 client: %v", err)
    }
}

func TestAdvisorv20200721Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := advisorv20200721.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init advisor_v20200721 client: %v", err)
    }
}

func TestAfv20200226Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := afv20200226.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init af_v20200226 client: %v", err)
    }
}

func TestAfcv20200226Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := afcv20200226.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init afc_v20200226 client: %v", err)
    }
}

func TestAi3Dv20250513Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ai3dv20250513.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ai3d_v20250513 client: %v", err)
    }
}

func TestAiartv20221229Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := aiartv20221229.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init aiart_v20221229 client: %v", err)
    }
}

func TestAmev20190916Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := amev20190916.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ame_v20190916 client: %v", err)
    }
}

func TestAmsv20200608Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := amsv20200608.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ams_v20200608 client: %v", err)
    }
}

func TestAmsv20201229Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := amsv20201229.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ams_v20201229 client: %v", err)
    }
}

func TestAnicloudv20220923Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := anicloudv20220923.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init anicloud_v20220923 client: %v", err)
    }
}

func TestAntiddosv20200309Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := antiddosv20200309.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init antiddos_v20200309 client: %v", err)
    }
}

func TestApev20200513Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := apev20200513.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ape_v20200513 client: %v", err)
    }
}

func TestApiv20201106Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := apiv20201106.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init api_v20201106 client: %v", err)
    }
}

func TestApigatewayv20180808Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := apigatewayv20180808.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init apigateway_v20180808 client: %v", err)
    }
}

func TestApmv20210622Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := apmv20210622.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init apm_v20210622 client: %v", err)
    }
}

func TestAsv20180419Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := asv20180419.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init as_v20180419 client: %v", err)
    }
}

func TestAsrv20190614Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := asrv20190614.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init asr_v20190614 client: %v", err)
    }
}

func TestAswv20200722Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := aswv20200722.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init asw_v20200722 client: %v", err)
    }
}

func TestBav20200720Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bav20200720.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ba_v20200720 client: %v", err)
    }
}

func TestBatchv20170312Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := batchv20170312.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init batch_v20170312 client: %v", err)
    }
}

func TestBdav20200324Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bdav20200324.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bda_v20200324 client: %v", err)
    }
}

func TestBhv20230418Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bhv20230418.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bh_v20230418 client: %v", err)
    }
}

func TestBiv20220105Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := biv20220105.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bi_v20220105 client: %v", err)
    }
}

func TestBillingv20180709Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := billingv20180709.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init billing_v20180709 client: %v", err)
    }
}

func TestBizlivev20190313Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bizlivev20190313.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bizlive_v20190313 client: %v", err)
    }
}

func TestBmv20180423Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bmv20180423.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bm_v20180423 client: %v", err)
    }
}

func TestBmav20210624Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bmav20210624.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bma_v20210624 client: %v", err)
    }
}

func TestBmav20221115Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bmav20221115.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bma_v20221115 client: %v", err)
    }
}

func TestBmeipv20180625Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bmeipv20180625.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bmeip_v20180625 client: %v", err)
    }
}

func TestBmlbv20180625Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bmlbv20180625.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bmlb_v20180625 client: %v", err)
    }
}

func TestBmvpcv20180625Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bmvpcv20180625.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bmvpc_v20180625 client: %v", err)
    }
}

func TestBpaasv20181217Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bpaasv20181217.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bpaas_v20181217 client: %v", err)
    }
}

func TestBriv20190328Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := briv20190328.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bri_v20190328 client: %v", err)
    }
}

func TestBscav20210811Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := bscav20210811.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init bsca_v20210811 client: %v", err)
    }
}

func TestBtoev20210303Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := btoev20210303.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init btoe_v20210303 client: %v", err)
    }
}

func TestBtoev20210514Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := btoev20210514.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init btoe_v20210514 client: %v", err)
    }
}

func TestCav20230228Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cav20230228.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ca_v20230228 client: %v", err)
    }
}

func TestCamv20190116Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := camv20190116.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cam_v20190116 client: %v", err)
    }
}

func TestCaptchav20190722Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := captchav20190722.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init captcha_v20190722 client: %v", err)
    }
}

func TestCarv20220110Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := carv20220110.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init car_v20220110 client: %v", err)
    }
}

func TestCatv20180409Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := catv20180409.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cat_v20180409 client: %v", err)
    }
}

func TestCbsv20170312Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cbsv20170312.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cbs_v20170312 client: %v", err)
    }
}

func TestCccv20200210Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cccv20200210.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ccc_v20200210 client: %v", err)
    }
}

func TestCdbv20170320Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cdbv20170320.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cdb_v20170320 client: %v", err)
    }
}

func TestCdcv20201214Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cdcv20201214.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cdc_v20201214 client: %v", err)
    }
}

func TestCdnv20180606Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cdnv20180606.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cdn_v20180606 client: %v", err)
    }
}

func TestCdsv20180420Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cdsv20180420.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cds_v20180420 client: %v", err)
    }
}

func TestCdwchv20200915Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cdwchv20200915.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cdwch_v20200915 client: %v", err)
    }
}

func TestCdwdorisv20211228Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cdwdorisv20211228.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cdwdoris_v20211228 client: %v", err)
    }
}

func TestCdwpgv20201230Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cdwpgv20201230.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cdwpg_v20201230 client: %v", err)
    }
}

func TestCdzv20221123Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cdzv20221123.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cdz_v20221123 client: %v", err)
    }
}

func TestCfgv20210820Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cfgv20210820.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cfg_v20210820 client: %v", err)
    }
}

func TestCfsv20190719Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cfsv20190719.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cfs_v20190719 client: %v", err)
    }
}

func TestCfwv20190904Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cfwv20190904.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cfw_v20190904 client: %v", err)
    }
}

func TestChcv20230418Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := chcv20230418.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init chc_v20230418 client: %v", err)
    }
}

func TestChdfsv20190718Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := chdfsv20190718.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init chdfs_v20190718 client: %v", err)
    }
}

func TestChdfsv20201112Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := chdfsv20201112.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init chdfs_v20201112 client: %v", err)
    }
}

func TestCiamv20220331Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ciamv20220331.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ciam_v20220331 client: %v", err)
    }
}

func TestCiiv20201210Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ciiv20201210.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cii_v20201210 client: %v", err)
    }
}

func TestCiiv20210408Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ciiv20210408.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cii_v20210408 client: %v", err)
    }
}

func TestCimv20190318Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cimv20190318.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cim_v20190318 client: %v", err)
    }
}

func TestCkafkav20190819Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ckafkav20190819.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ckafka_v20190819 client: %v", err)
    }
}

func TestClbv20180317Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := clbv20180317.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init clb_v20180317 client: %v", err)
    }
}

func TestCloudappv20220530Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cloudappv20220530.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cloudapp_v20220530 client: %v", err)
    }
}

func TestCloudauditv20190319Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cloudauditv20190319.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cloudaudit_v20190319 client: %v", err)
    }
}

func TestCloudhsmv20191112Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cloudhsmv20191112.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cloudhsm_v20191112 client: %v", err)
    }
}

func TestCloudstudiov20230508Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cloudstudiov20230508.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cloudstudio_v20230508 client: %v", err)
    }
}

func TestClsv20201016Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := clsv20201016.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cls_v20201016 client: %v", err)
    }
}

func TestCmev20191029Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cmev20191029.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cme_v20191029 client: %v", err)
    }
}

func TestCmqv20190304Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cmqv20190304.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cmq_v20190304 client: %v", err)
    }
}

func TestCmsv20190321Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cmsv20190321.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cms_v20190321 client: %v", err)
    }
}

func TestConfigv20220802Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := configv20220802.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init config_v20220802 client: %v", err)
    }
}

func TestControlcenterv20230110Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := controlcenterv20230110.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init controlcenter_v20230110 client: %v", err)
    }
}

func TestCpdpv20190820Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cpdpv20190820.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cpdp_v20190820 client: %v", err)
    }
}

func TestCsipv20221121Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := csipv20221121.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init csip_v20221121 client: %v", err)
    }
}

func TestCsxgv20230303Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := csxgv20230303.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init csxg_v20230303 client: %v", err)
    }
}

func TestCtemv20231128Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ctemv20231128.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ctem_v20231128 client: %v", err)
    }
}

func TestCtsdbv20230202Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ctsdbv20230202.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ctsdb_v20230202 client: %v", err)
    }
}

func TestCvmv20170312Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cvmv20170312.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cvm_v20170312 client: %v", err)
    }
}

func TestCwpv20180228Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cwpv20180228.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cwp_v20180228 client: %v", err)
    }
}

func TestCwsv20180312Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cwsv20180312.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cws_v20180312 client: %v", err)
    }
}

func TestCynosdbv20190107Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := cynosdbv20190107.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init cynosdb_v20190107 client: %v", err)
    }
}

func TestDasbv20191018Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dasbv20191018.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dasb_v20191018 client: %v", err)
    }
}

func TestDayuv20180709Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dayuv20180709.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dayu_v20180709 client: %v", err)
    }
}

func TestDbbrainv20191016Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dbbrainv20191016.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dbbrain_v20191016 client: %v", err)
    }
}

func TestDbbrainv20210527Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dbbrainv20210527.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dbbrain_v20210527 client: %v", err)
    }
}

func TestDbdcv20201029Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dbdcv20201029.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dbdc_v20201029 client: %v", err)
    }
}

func TestDcv20180410Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dcv20180410.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dc_v20180410 client: %v", err)
    }
}

func TestDcdbv20180411Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dcdbv20180411.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dcdb_v20180411 client: %v", err)
    }
}

func TestDlcv20210125Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dlcv20210125.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dlc_v20210125 client: %v", err)
    }
}

func TestDnspodv20210323Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dnspodv20210323.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dnspod_v20210323 client: %v", err)
    }
}

func TestDomainv20180808Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := domainv20180808.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init domain_v20180808 client: %v", err)
    }
}

func TestDrmv20181115Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := drmv20181115.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init drm_v20181115 client: %v", err)
    }
}

func TestDsv20180523Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dsv20180523.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ds_v20180523 client: %v", err)
    }
}

func TestDsgcv20190723Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dsgcv20190723.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dsgc_v20190723 client: %v", err)
    }
}

func TestDtsv20180330Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dtsv20180330.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dts_v20180330 client: %v", err)
    }
}

func TestDtsv20211206Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := dtsv20211206.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init dts_v20211206 client: %v", err)
    }
}

func TestEbv20210416Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ebv20210416.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init eb_v20210416 client: %v", err)
    }
}

func TestEccv20181213Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := eccv20181213.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ecc_v20181213 client: %v", err)
    }
}

func TestEcdnv20191012Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ecdnv20191012.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ecdn_v20191012 client: %v", err)
    }
}

func TestEcmv20190719Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ecmv20190719.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ecm_v20190719 client: %v", err)
    }
}

func TestEiamv20210420Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := eiamv20210420.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init eiam_v20210420 client: %v", err)
    }
}

func TestEisv20200715Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := eisv20200715.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init eis_v20200715 client: %v", err)
    }
}

func TestEisv20210601Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := eisv20210601.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init eis_v20210601 client: %v", err)
    }
}

func TestEmrv20190103Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := emrv20190103.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init emr_v20190103 client: %v", err)
    }
}

func TestEsv20180416Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := esv20180416.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init es_v20180416 client: %v", err)
    }
}

func TestEsv20250101Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := esv20250101.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init es_v20250101 client: %v", err)
    }
}

func TestEssv20201111Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := essv20201111.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ess_v20201111 client: %v", err)
    }
}

func TestEssbasicv20201222Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := essbasicv20201222.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init essbasic_v20201222 client: %v", err)
    }
}

func TestEssbasicv20210526Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := essbasicv20210526.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init essbasic_v20210526 client: %v", err)
    }
}

func TestFacefusionv20181201Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := facefusionv20181201.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init facefusion_v20181201 client: %v", err)
    }
}

func TestFacefusionv20220927Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := facefusionv20220927.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init facefusion_v20220927 client: %v", err)
    }
}

func TestFaceidv20180301Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := faceidv20180301.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init faceid_v20180301 client: %v", err)
    }
}

func TestFmuv20191213Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := fmuv20191213.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init fmu_v20191213 client: %v", err)
    }
}

func TestFtv20200304Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ftv20200304.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ft_v20200304 client: %v", err)
    }
}

func TestGaapv20180529Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := gaapv20180529.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init gaap_v20180529 client: %v", err)
    }
}

func TestGmev20180711Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := gmev20180711.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init gme_v20180711 client: %v", err)
    }
}

func TestGoosefsv20220519Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := goosefsv20220519.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init goosefs_v20220519 client: %v", err)
    }
}

func TestGsv20191118Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := gsv20191118.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init gs_v20191118 client: %v", err)
    }
}

func TestGwlbv20240906Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := gwlbv20240906.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init gwlb_v20240906 client: %v", err)
    }
}

func TestHabov20181203Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := habov20181203.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init habo_v20181203 client: %v", err)
    }
}

func TestHaiv20230812Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := haiv20230812.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init hai_v20230812 client: %v", err)
    }
}

func TestHasimv20210716Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := hasimv20210716.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init hasim_v20210716 client: %v", err)
    }
}

func TestHcmv20181106Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := hcmv20181106.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init hcm_v20181106 client: %v", err)
    }
}

func TestHunyuanv20230901Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := hunyuanv20230901.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init hunyuan_v20230901 client: %v", err)
    }
}

func TestIaiv20180301Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iaiv20180301.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iai_v20180301 client: %v", err)
    }
}

func TestIaiv20200303Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iaiv20200303.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iai_v20200303 client: %v", err)
    }
}

func TestIapv20240713Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iapv20240713.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iap_v20240713 client: %v", err)
    }
}

func TestIcv20190307Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := icv20190307.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ic_v20190307 client: %v", err)
    }
}

func TestIcrv20211014Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := icrv20211014.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init icr_v20211014 client: %v", err)
    }
}

func TestIev20200304Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iev20200304.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ie_v20200304 client: %v", err)
    }
}

func TestIgv20210518Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := igv20210518.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ig_v20210518 client: %v", err)
    }
}

func TestIgtmv20231024Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := igtmv20231024.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init igtm_v20231024 client: %v", err)
    }
}

func TestImsv20200713Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := imsv20200713.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ims_v20200713 client: %v", err)
    }
}

func TestImsv20201229Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := imsv20201229.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ims_v20201229 client: %v", err)
    }
}

func TestIoav20220601Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ioav20220601.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ioa_v20220601 client: %v", err)
    }
}

func TestIotv20180123Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iotv20180123.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iot_v20180123 client: %v", err)
    }
}

func TestIotcloudv20180614Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iotcloudv20180614.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iotcloud_v20180614 client: %v", err)
    }
}

func TestIotcloudv20210408Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iotcloudv20210408.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iotcloud_v20210408 client: %v", err)
    }
}

func TestIotexplorerv20190423Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iotexplorerv20190423.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iotexplorer_v20190423 client: %v", err)
    }
}

func TestIotvideov20191126Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iotvideov20191126.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iotvideo_v20191126 client: %v", err)
    }
}

func TestIotvideov20201215Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iotvideov20201215.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iotvideo_v20201215 client: %v", err)
    }
}

func TestIotvideov20211125Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iotvideov20211125.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iotvideo_v20211125 client: %v", err)
    }
}

func TestIotvideoindustryv20201201Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := iotvideoindustryv20201201.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iotvideoindustry_v20201201 client: %v", err)
    }
}

func TestIrpv20220324Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := irpv20220324.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init irp_v20220324 client: %v", err)
    }
}

func TestIrpv20220805Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := irpv20220805.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init irp_v20220805 client: %v", err)
    }
}

func TestIssv20230517Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := issv20230517.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init iss_v20230517 client: %v", err)
    }
}

func TestIvldv20210903Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ivldv20210903.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ivld_v20210903 client: %v", err)
    }
}

func TestKeewidbv20220308Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := keewidbv20220308.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init keewidb_v20220308 client: %v", err)
    }
}

func TestKmsv20190118Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := kmsv20190118.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init kms_v20190118 client: %v", err)
    }
}

func TestLcicv20220817Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := lcicv20220817.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init lcic_v20220817 client: %v", err)
    }
}

func TestLighthousev20200324Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := lighthousev20200324.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init lighthouse_v20200324 client: %v", err)
    }
}

func TestLivev20180801Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := livev20180801.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init live_v20180801 client: %v", err)
    }
}

func TestLkev20231130Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := lkev20231130.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init lke_v20231130 client: %v", err)
    }
}

func TestLkeapv20240522Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := lkeapv20240522.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init lkeap_v20240522 client: %v", err)
    }
}

func TestLowcodev20210108Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := lowcodev20210108.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init lowcode_v20210108 client: %v", err)
    }
}

func TestMallv20230518Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := mallv20230518.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init mall_v20230518 client: %v", err)
    }
}

func TestMariadbv20170312Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := mariadbv20170312.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init mariadb_v20170312 client: %v", err)
    }
}

func TestMarketv20191010Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := marketv20191010.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init market_v20191010 client: %v", err)
    }
}

func TestMemcachedv20190318Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := memcachedv20190318.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init memcached_v20190318 client: %v", err)
    }
}

func TestMmpsv20200710Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := mmpsv20200710.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init mmps_v20200710 client: %v", err)
    }
}

func TestMnav20210119Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := mnav20210119.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init mna_v20210119 client: %v", err)
    }
}

func TestMongodbv20180408Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := mongodbv20180408.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init mongodb_v20180408 client: %v", err)
    }
}

func TestMongodbv20190725Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := mongodbv20190725.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init mongodb_v20190725 client: %v", err)
    }
}

func TestMonitorv20180724Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := monitorv20180724.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init monitor_v20180724 client: %v", err)
    }
}

func TestMonitorv20230616Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := monitorv20230616.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init monitor_v20230616 client: %v", err)
    }
}

func TestMpsv20190612Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := mpsv20190612.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init mps_v20190612 client: %v", err)
    }
}

func TestMqttv20240516Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := mqttv20240516.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init mqtt_v20240516 client: %v", err)
    }
}

func TestMrsv20200910Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := mrsv20200910.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init mrs_v20200910 client: %v", err)
    }
}

func TestMsv20180408Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := msv20180408.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ms_v20180408 client: %v", err)
    }
}

func TestMspv20180319Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := mspv20180319.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init msp_v20180319 client: %v", err)
    }
}

func TestNlpv20190408Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := nlpv20190408.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init nlp_v20190408 client: %v", err)
    }
}

func TestNppv20190823Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := nppv20190823.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init npp_v20190823 client: %v", err)
    }
}

func TestOceanusv20190422Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := oceanusv20190422.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init oceanus_v20190422 client: %v", err)
    }
}

func TestOcrv20181119Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ocrv20181119.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ocr_v20181119 client: %v", err)
    }
}

func TestOmicsv20221128Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := omicsv20221128.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init omics_v20221128 client: %v", err)
    }
}

func TestOrganizationv20181225Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := organizationv20181225.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init organization_v20181225 client: %v", err)
    }
}

func TestOrganizationv20210331Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := organizationv20210331.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init organization_v20210331 client: %v", err)
    }
}

func TestPartnersv20180321Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := partnersv20180321.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init partners_v20180321 client: %v", err)
    }
}

func TestPostgresv20170312Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := postgresv20170312.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init postgres_v20170312 client: %v", err)
    }
}

func TestPrivatednsv20201028Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := privatednsv20201028.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init privatedns_v20201028 client: %v", err)
    }
}

func TestPtsv20210728Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ptsv20210728.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init pts_v20210728 client: %v", err)
    }
}

func TestRcev20201103Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := rcev20201103.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init rce_v20201103 client: %v", err)
    }
}

func TestRedisv20180412Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := redisv20180412.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init redis_v20180412 client: %v", err)
    }
}

func TestRegionv20220627Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := regionv20220627.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init region_v20220627 client: %v", err)
    }
}

func TestRumv20210622Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := rumv20210622.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init rum_v20210622 client: %v", err)
    }
}

func TestScfv20180416Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := scfv20180416.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init scf_v20180416 client: %v", err)
    }
}

func TestSecuritylakev20240117Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := securitylakev20240117.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init securitylake_v20240117 client: %v", err)
    }
}

func TestSesv20201002Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := sesv20201002.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ses_v20201002 client: %v", err)
    }
}

func TestSmhv20210712Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := smhv20210712.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init smh_v20210712 client: %v", err)
    }
}

func TestSmopv20201203Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := smopv20201203.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init smop_v20201203 client: %v", err)
    }
}

func TestSmsv20190711Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := smsv20190711.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init sms_v20190711 client: %v", err)
    }
}

func TestSmsv20210111Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := smsv20210111.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init sms_v20210111 client: %v", err)
    }
}

func TestSoev20180724Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := soev20180724.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init soe_v20180724 client: %v", err)
    }
}

func TestSqlserverv20180328Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := sqlserverv20180328.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init sqlserver_v20180328 client: %v", err)
    }
}

func TestSsav20180608Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ssav20180608.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ssa_v20180608 client: %v", err)
    }
}

func TestSslv20191205Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := sslv20191205.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ssl_v20191205 client: %v", err)
    }
}

func TestSslpodv20190605Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := sslpodv20190605.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init sslpod_v20190605 client: %v", err)
    }
}

func TestSsmv20190923Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ssmv20190923.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init ssm_v20190923 client: %v", err)
    }
}

func TestStsv20180813Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := stsv20180813.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init sts_v20180813 client: %v", err)
    }
}

func TestSvpv20240125Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := svpv20240125.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init svp_v20240125 client: %v", err)
    }
}

func TestTafv20200210Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tafv20200210.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init taf_v20200210 client: %v", err)
    }
}

func TestTagv20180813Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tagv20180813.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tag_v20180813 client: %v", err)
    }
}

func TestTatv20201028Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tatv20201028.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tat_v20201028 client: %v", err)
    }
}

func TestTbaasv20180416Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tbaasv20180416.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tbaas_v20180416 client: %v", err)
    }
}

func TestTbpv20190311Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tbpv20190311.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tbp_v20190311 client: %v", err)
    }
}

func TestTbpv20190627Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tbpv20190627.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tbp_v20190627 client: %v", err)
    }
}

func TestTcaplusdbv20190823Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tcaplusdbv20190823.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tcaplusdb_v20190823 client: %v", err)
    }
}

func TestTcbv20180608Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tcbv20180608.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tcb_v20180608 client: %v", err)
    }
}

func TestTcbrv20220217Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tcbrv20220217.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tcbr_v20220217 client: %v", err)
    }
}

func TestTccatalogv20241024Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tccatalogv20241024.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tccatalog_v20241024 client: %v", err)
    }
}

func TestTchdv20230306Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tchdv20230306.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tchd_v20230306 client: %v", err)
    }
}

func TestTcmv20210413Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tcmv20210413.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tcm_v20210413 client: %v", err)
    }
}

func TestTcrv20190924Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tcrv20190924.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tcr_v20190924 client: %v", err)
    }
}

func TestTcssv20201101Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tcssv20201101.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tcss_v20201101 client: %v", err)
    }
}

func TestTdcpgv20211118Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tdcpgv20211118.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tdcpg_v20211118 client: %v", err)
    }
}

func TestTdidv20210519Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tdidv20210519.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tdid_v20210519 client: %v", err)
    }
}

func TestTdmqv20200217Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tdmqv20200217.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tdmq_v20200217 client: %v", err)
    }
}

func TestTdsv20220801Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tdsv20220801.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tds_v20220801 client: %v", err)
    }
}

func TestTemv20201221Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := temv20201221.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tem_v20201221 client: %v", err)
    }
}

func TestTemv20210701Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := temv20210701.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tem_v20210701 client: %v", err)
    }
}

func TestTeov20220106Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := teov20220106.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init teo_v20220106 client: %v", err)
    }
}

func TestTeov20220901Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := teov20220901.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init teo_v20220901 client: %v", err)
    }
}

func TestThpcv20211109Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := thpcv20211109.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init thpc_v20211109 client: %v", err)
    }
}

func TestThpcv20220401Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := thpcv20220401.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init thpc_v20220401 client: %v", err)
    }
}

func TestThpcv20230321Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := thpcv20230321.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init thpc_v20230321 client: %v", err)
    }
}

func TestTiav20180226Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tiav20180226.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tia_v20180226 client: %v", err)
    }
}

func TestTiiav20190529Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tiiav20190529.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tiia_v20190529 client: %v", err)
    }
}

func TestTionev20191022Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tionev20191022.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tione_v20191022 client: %v", err)
    }
}

func TestTionev20211111Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tionev20211111.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tione_v20211111 client: %v", err)
    }
}

func TestTiwv20190919Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tiwv20190919.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tiw_v20190919 client: %v", err)
    }
}

func TestTkev20180525Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tkev20180525.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tke_v20180525 client: %v", err)
    }
}

func TestTkev20220501Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tkev20220501.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tke_v20220501 client: %v", err)
    }
}

func TestTkgdqv20190411Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tkgdqv20190411.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tkgdq_v20190411 client: %v", err)
    }
}

func TestTmsv20200713Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tmsv20200713.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tms_v20200713 client: %v", err)
    }
}

func TestTmsv20201229Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tmsv20201229.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tms_v20201229 client: %v", err)
    }
}

func TestTmtv20180321Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tmtv20180321.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tmt_v20180321 client: %v", err)
    }
}

func TestTourismv20230215Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tourismv20230215.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tourism_v20230215 client: %v", err)
    }
}

func TestTrabbitv20230418Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := trabbitv20230418.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init trabbit_v20230418 client: %v", err)
    }
}

func TestTrocketv20230308Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := trocketv20230308.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init trocket_v20230308 client: %v", err)
    }
}

func TestTrpv20210515Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := trpv20210515.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init trp_v20210515 client: %v", err)
    }
}

func TestTrrov20220325Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := trrov20220325.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init trro_v20220325 client: %v", err)
    }
}

func TestTrtcv20190722Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := trtcv20190722.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init trtc_v20190722 client: %v", err)
    }
}

func TestTsev20201207Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tsev20201207.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tse_v20201207 client: %v", err)
    }
}

func TestTsfv20180326Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tsfv20180326.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tsf_v20180326 client: %v", err)
    }
}

func TestTsiv20210325Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tsiv20210325.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tsi_v20210325 client: %v", err)
    }
}

func TestTswv20200924Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tswv20200924.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tsw_v20200924 client: %v", err)
    }
}

func TestTswv20210412Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := tswv20210412.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tsw_v20210412 client: %v", err)
    }
}

func TestTtsv20190823Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := ttsv20190823.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init tts_v20190823 client: %v", err)
    }
}

func TestVcgv20240404Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vcgv20240404.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vcg_v20240404 client: %v", err)
    }
}

func TestVclmv20240523Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vclmv20240523.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vclm_v20240523 client: %v", err)
    }
}

func TestVcubev20220410Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vcubev20220410.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vcube_v20220410 client: %v", err)
    }
}

func TestVdbv20230616Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vdbv20230616.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vdb_v20230616 client: %v", err)
    }
}

func TestVmv20200709Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vmv20200709.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vm_v20200709 client: %v", err)
    }
}

func TestVmv20201229Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vmv20201229.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vm_v20201229 client: %v", err)
    }
}

func TestVmv20210922Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vmv20210922.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vm_v20210922 client: %v", err)
    }
}

func TestVmsv20200902Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vmsv20200902.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vms_v20200902 client: %v", err)
    }
}

func TestVodv20180717Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vodv20180717.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vod_v20180717 client: %v", err)
    }
}

func TestVodv20240718Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vodv20240718.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vod_v20240718 client: %v", err)
    }
}

func TestVpcv20170312Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vpcv20170312.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vpc_v20170312 client: %v", err)
    }
}

func TestVrsv20200824Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vrsv20200824.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vrs_v20200824 client: %v", err)
    }
}

func TestVtcv20240223Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := vtcv20240223.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init vtc_v20240223 client: %v", err)
    }
}

func TestWafv20180125Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := wafv20180125.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init waf_v20180125 client: %v", err)
    }
}

func TestWavv20210129Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := wavv20210129.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init wav_v20210129 client: %v", err)
    }
}

func TestWedatav20210820Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := wedatav20210820.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init wedata_v20210820 client: %v", err)
    }
}

func TestWeilingwithv20230427Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := weilingwithv20230427.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init weilingwith_v20230427 client: %v", err)
    }
}

func TestWsav20250508Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := wsav20250508.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init wsa_v20250508 client: %v", err)
    }
}

func TestWssv20180426Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := wssv20180426.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init wss_v20180426 client: %v", err)
    }
}

func TestYinsudav20220527Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := yinsudav20220527.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init yinsuda_v20220527 client: %v", err)
    }
}

func TestYunjingv20180228Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := yunjingv20180228.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init yunjing_v20180228 client: %v", err)
    }
}

func TestYunsouv20180504Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := yunsouv20180504.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init yunsou_v20180504 client: %v", err)
    }
}

func TestYunsouv20191115Import(t *testing.T) {
    credential := common.NewCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"))
    client, err := yunsouv20191115.NewClient(
        credential,
        regions.Guangzhou,
        profile.NewClientProfile())
    if err != nil || client == nil {
        t.Errorf("fail to init yunsou_v20191115 client: %v", err)
    }
}
