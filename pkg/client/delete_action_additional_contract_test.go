// Copyright (c) ZStack.io, Inc.

package client

import (
	"net/url"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestDeleteActionAdditionalContracts(t *testing.T) {
	option := "option"
	service := "DHCP"
	enforcing := "Enforcing"
	disabled := false
	port := 443
	startTime := int64(1700000000000)
	endTime := int64(1700000001000)
	tests := []deleteActionContract{
		{
			name:  "DeleteAlert",
			path:  "/zstack/v1/monitoring/alerts",
			query: url.Values{"deleteMode": {"Permissive"}, "uuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteAlertWithParams(param.DeleteModePermissive, param.DeleteAlertParamDetail{Uuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "DeleteBackupFileInPublic",
			path:  "/zstack/v1/hybrid/backup-mysql",
			query: url.Values{"deleteMode": {"Permissive"}, "file": {"folder/file & name.sql"}, "regionId": {"selector"}, "type": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteBackupFileInPublicWithParams(param.DeleteModePermissive, param.DeleteBackupFileInPublicParamDetail{File: "folder/file & name.sql", RegionId: "selector", Type: "selector"})
			},
		},
		{
			name:  "DeleteBilling",
			path:  "/zstack/v1/billings/billings",
			query: url.Values{"deleteMode": {"Permissive"}, "accountUuid": {"option"}, "endTime": {"1700000001000"}, "startTime": {"1700000000000"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteBillingWithParams(param.DeleteModePermissive, param.DeleteBillingParamDetail{AccountUuid: &option, EndTime: &endTime, StartTime: &startTime})
			},
		},
		{
			name:  "DeleteDatasets",
			path:  "/zstack/v1/ai/datasets",
			query: url.Values{"deleteMode": {"Permissive"}, "uuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteDatasetsWithParams(param.DeleteModePermissive, param.DeleteDatasetsParamDetail{Uuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "DeleteHybridEipRemote",
			path:  "/zstack/v1/hybrid/eip/parent-id/remote",
			query: url.Values{"deleteMode": {"Permissive"}, "type": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteHybridEipRemoteWithParams("parent-id", param.DeleteModePermissive, param.DeleteHybridEipRemoteParamDetail{Type: "selector"})
			},
		},
		{
			name:  "DeleteLicense",
			path:  "/zstack/v1/licenses/mn/parent-id/actions",
			query: url.Values{"deleteMode": {"Permissive"}, "module": {"option"}, "uuid": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteLicenseWithParams("parent-id", param.DeleteModePermissive, param.DeleteLicenseParamDetail{Module: &option, Uuid: "selector"})
			},
		},
		{
			name:  "DeleteLogConfiguration",
			path:  "/zstack/v1/log/configurations/log4j2",
			query: url.Values{"deleteMode": {"Permissive"}, "configId": {"922337203685477580"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteLogConfigurationWithParams(param.DeleteModePermissive, param.DeleteLogConfigurationParamDetail{ConfigId: 922337203685477580})
			},
		},
		{
			name:  "DeleteMetricData",
			path:  "/zstack/v1/zwatch/metrics",
			query: url.Values{"deleteMode": {"Permissive"}, "labels": {"value-a", "value-b"}, "metricName": {"selector"}, "namespace": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteMetricDataWithParams(param.DeleteModePermissive, param.DeleteMetricDataParamDetail{Labels: []string{"value-a", "value-b"}, MetricName: "selector", Namespace: "selector"})
			},
		},
		{
			name:  "DeleteModelEvaluationTasks",
			path:  "/zstack/v1/model-evaluation-tasks",
			query: url.Values{"deleteMode": {"Permissive"}, "uuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteModelEvaluationTasksWithParams(param.DeleteModePermissive, param.DeleteModelEvaluationTasksParamDetail{Uuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "DeleteModelServiceInstanceGroups",
			path:  "/zstack/v1/ai/model-services/instances/groups",
			query: url.Values{"deleteMode": {"Permissive"}, "uuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteModelServiceInstanceGroupsWithParams(param.DeleteModePermissive, param.DeleteModelServiceInstanceGroupsParamDetail{Uuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "DeleteModelServices",
			path:  "/zstack/v1/ai/model-services/",
			query: url.Values{"deleteMode": {"Permissive"}, "uuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteModelServicesWithParams(param.DeleteModePermissive, param.DeleteModelServicesParamDetail{Uuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "DeleteModels",
			path:  "/zstack/v1/ai/models",
			query: url.Values{"deleteMode": {"Permissive"}, "uuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteModelsWithParams(param.DeleteModePermissive, param.DeleteModelsParamDetail{Uuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "DeleteResourceStackVmPortMonitor",
			path:  "/zstack/v1/cloudformation/stack/monitor/delvm",
			query: url.Values{"deleteMode": {"Permissive"}, "port": {"443"}, "stackUuid": {"option"}, "vmInstanceUuid": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteResourceStackVmPortMonitorWithParams(param.DeleteModePermissive, param.DeleteResourceStackVmPortMonitorParamDetail{Port: &port, StackUuid: &option, VmInstanceUuid: "selector"})
			},
		},
		{
			name:  "DetachHostFromHostSchedulingRuleGroup",
			path:  "/zstack/v1/hostSchedulingRuleGroup/parent-id/host",
			query: url.Values{"deleteMode": {"Permissive"}, "hostUuid": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DetachHostFromHostSchedulingRuleGroupWithParams("parent-id", param.DeleteModePermissive, param.DetachHostFromHostSchedulingRuleGroupParamDetail{HostUuid: "selector"})
			},
		},
		{
			name:  "DetachServiceFromObservabilityServer",
			path:  "/zstack/v1/observability-server/parent-id/service",
			query: url.Values{"deleteMode": {"Permissive"}, "serviceType": {"selector"}, "serviceUuid": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DetachServiceFromObservabilityServerWithParams("parent-id", param.DeleteModePermissive, param.DetachServiceFromObservabilityServerParamDetail{ServiceType: "selector", ServiceUuid: "selector"})
			},
		},
		{
			name:  "DetachUserDefinedXmlHookScriptFromVm",
			path:  "/zstack/v1/xmlhook/vm-instances/parent-id/detach",
			query: url.Values{"deleteMode": {"Permissive"}, "startupStrategy": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DetachUserDefinedXmlHookScriptFromVmWithParams("parent-id", param.DeleteModePermissive, param.DetachUserDefinedXmlHookScriptFromVmParamDetail{StartupStrategy: &option})
			},
		},
		{
			name:  "DetachVipFromVpcSharedQos",
			path:  "/zstack/v1/vips/sharedqos/parent-id/vips",
			query: url.Values{"deleteMode": {"Permissive"}, "vipLists": {"value-a", "value-b"}, "vipUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.DetachVipFromVpcSharedQosWithParams("parent-id", param.DeleteModePermissive, param.DetachVipFromVpcSharedQosParamDetail{VipLists: []string{"value-a", "value-b"}, VipUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "DetachVmFromVmSchedulingRuleGroup",
			path:  "/zstack/v1/vmSchedulingRuleGroup/parent-id/vmInstance/",
			query: url.Values{"deleteMode": {"Permissive"}, "vmUuid": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DetachVmFromVmSchedulingRuleGroupWithParams("parent-id", param.DeleteModePermissive, param.DetachVmFromVmSchedulingRuleGroupParamDetail{VmUuid: "selector"})
			},
		},
		{
			name:  "RemoveAccessControlListFromLoadBalancer",
			path:  "/zstack/v1/load-balancers/listeners/parent-id/access-control-lists",
			query: url.Values{"deleteMode": {"Permissive"}, "aclUuids": {"value-a", "value-b"}, "serverGroupUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveAccessControlListFromLoadBalancerWithParams("parent-id", param.DeleteModePermissive, param.RemoveAccessControlListFromLoadBalancerParamDetail{AclUuids: []string{"value-a", "value-b"}, ServerGroupUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "RemoveAttributesFromIAM2Organization",
			path:  "/zstack/v1/iam2/organizations/parent-id/attributes",
			query: url.Values{"deleteMode": {"Permissive"}, "attributeUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveAttributesFromIAM2OrganizationWithParams("parent-id", param.DeleteModePermissive, param.RemoveAttributesFromIAM2OrganizationParamDetail{AttributeUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "RemoveAttributesFromIAM2Project",
			path:  "/zstack/v1/iam2/projects/parent-id/attributes",
			query: url.Values{"deleteMode": {"Permissive"}, "attributeUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveAttributesFromIAM2ProjectWithParams("parent-id", param.DeleteModePermissive, param.RemoveAttributesFromIAM2ProjectParamDetail{AttributeUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "RemoveAttributesFromIAM2VirtualID",
			path:  "/zstack/v1/iam2/virtual-ids/parent-id/attributes",
			query: url.Values{"deleteMode": {"Permissive"}, "attributeUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveAttributesFromIAM2VirtualIDWithParams("parent-id", param.DeleteModePermissive, param.RemoveAttributesFromIAM2VirtualIDParamDetail{AttributeUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "RemoveAttributesFromIAM2VirtualIDGroup",
			path:  "/zstack/v1/iam2/projects/groups/parent-id/attributes",
			query: url.Values{"deleteMode": {"Permissive"}, "attributeUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveAttributesFromIAM2VirtualIDGroupWithParams("parent-id", param.DeleteModePermissive, param.RemoveAttributesFromIAM2VirtualIDGroupParamDetail{AttributeUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "RemoveDnsFromVpcRouter",
			path:  "/zstack/v1/vpc/virtual-routers/parent-id/dns",
			query: url.Values{"deleteMode": {"Permissive"}, "dns": {"2001:db8::53"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveDnsFromVpcRouterWithParams("parent-id", param.DeleteModePermissive, param.RemoveDnsFromVpcRouterParamDetail{Dns: "2001:db8::53"})
			},
		},
		{
			name:  "RemoveIAM2VirtualIDGroupFromProjects",
			path:  "/zstack/v1/iam2/projects/groups",
			query: url.Values{"deleteMode": {"Permissive"}, "groupUuids": {"value-a", "value-b"}, "projectUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveIAM2VirtualIDGroupFromProjectsWithParams(param.DeleteModePermissive, param.RemoveIAM2VirtualIDGroupFromProjectsParamDetail{GroupUuids: []string{"value-a", "value-b"}, ProjectUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "RemoveRendezvousPointFromMulticastRouter",
			path:  "/zstack/v1/multicast/virtual-routers/parent-id/RendezvousPoint",
			query: url.Values{"deleteMode": {"Permissive"}, "groupAddress": {"selector"}, "rpAddress": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveRendezvousPointFromMulticastRouterWithParams("parent-id", param.DeleteModePermissive, param.RemoveRendezvousPointFromMulticastRouterParamDetail{GroupAddress: "selector", RpAddress: "selector"})
			},
		},
		{
			name:  "RemoveSchedulerJobsFromSchedulerJobGroup",
			path:  "/zstack/v1/scheduler/jobgroups/parent-id/job",
			query: url.Values{"deleteMode": {"Permissive"}, "schedulerJobUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveSchedulerJobsFromSchedulerJobGroupWithParams("parent-id", param.DeleteModePermissive, param.RemoveSchedulerJobsFromSchedulerJobGroupParamDetail{SchedulerJobUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "RemoveTicketTypesFromTicketFlowCollection",
			path:  "/zstack/v1/tickets/flow-collections/parent-id/ticket-types",
			query: url.Values{"deleteMode": {"Permissive"}, "ticketTypeUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveTicketTypesFromTicketFlowCollectionWithParams("parent-id", param.DeleteModePermissive, param.RemoveTicketTypesFromTicketFlowCollectionParamDetail{TicketTypeUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "RemoveVRouterNetworksFromFlowMeter",
			path:  "/zstack/v1/flowmeters/networks",
			query: url.Values{"deleteMode": {"Permissive"}, "uuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveVRouterNetworksFromFlowMeterWithParams(param.DeleteModePermissive, param.RemoveVRouterNetworksFromFlowMeterParamDetail{Uuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "RemoveVRouterNetworksFromOspfArea",
			path:  "/zstack/v1/routerArea/networks",
			query: url.Values{"deleteMode": {"Permissive"}, "uuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveVRouterNetworksFromOspfAreaWithParams(param.DeleteModePermissive, param.RemoveVRouterNetworksFromOspfAreaParamDetail{Uuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "UnregisterLicenseServer",
			path:  "/zstack/v1/license-server/client",
			query: url.Values{"deleteMode": {"Permissive"}, "clientAuthorizedNodeUuid": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.UnregisterLicenseServerWithParams(param.DeleteModePermissive, param.UnregisterLicenseServerParamDetail{ClientAuthorizedNodeUuid: &option})
			},
		},
		{
			name:  "WithdrawLicenseCapacityApplication",
			path:  "/zstack/v1/license-server/capacity-application",
			query: url.Values{"deleteMode": {"Permissive"}, "clientAuthorizedNodeUuid": {"selector"}, "licenseType": {"selector"}, "resourceUuidList": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.WithdrawLicenseCapacityApplicationWithParams(param.DeleteModePermissive, param.WithdrawLicenseCapacityApplicationParamDetail{ClientAuthorizedNodeUuid: "selector", LicenseType: "selector", ResourceUuidList: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "DeleteAliyunRouterInterfaceRemote",
			path:  "/zstack/v1/hybrid/aliyun/router-interface/remote/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "vRouterType": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteAliyunRouterInterfaceRemoteWithParams("parent-id", param.DeleteModePermissive, param.DeleteAliyunRouterInterfaceRemoteParamDetail{VRouterType: "selector"})
			},
		},
		{
			name:  "DeleteHybridEipFromLocal",
			path:  "/zstack/v1/hybrid/eip/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "type": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteHybridEipFromLocalWithParams("parent-id", param.DeleteModePermissive, param.DeleteHybridEipFromLocalParamDetail{Type: "selector"})
			},
		},
		{
			name:  "DeleteOssBucketFileRemote",
			path:  "/zstack/v1/hybrid/aliyun/oss-bucket-file/remote/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "fileName": {"folder/file & name"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteOssBucketFileRemoteWithParams("parent-id", param.DeleteModePermissive, param.DeleteOssBucketFileRemoteParamDetail{FileName: "folder/file & name"})
			},
		},
		{
			name:  "DeleteVxlanPoolRemoteVtep",
			path:  "/zstack/v1/l2-networks/parent-id/clusters/cluster-id/delete/remote-vtep-ip",
			query: url.Values{"deleteMode": {"Permissive"}, "remoteVtepIp": {"192.0.2.10"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteVxlanPoolRemoteVtepWithParams("parent-id", "cluster-id", param.DeleteModePermissive, param.DeleteVxlanPoolRemoteVtepParamDetail{RemoteVtepIp: "192.0.2.10"})
			},
		},
		{
			name:  "DetachProvisionNicFromBonding",
			path:  "/zstack/v1/baremetal2/bm-instances/bm2-bondings/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "provisionNicUuid": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.DetachProvisionNicFromBondingWithParams("parent-id", param.DeleteModePermissive, param.DetachProvisionNicFromBondingParamDetail{ProvisionNicUuid: "selector"})
			},
		},
		{
			name:  "DeleteBuildAppExportHistory",
			path:  "/zstack/v1/buildapp/exports/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "buildSystemUuid": {"option"}, "exportId": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteBuildAppExportHistoryWithParams("parent-id", param.DeleteModePermissive, param.DeleteBuildAppExportHistoryParamDetail{BuildSystemUuid: &option, ExportId: &option})
			},
		},
		{
			name:  "DeleteCbtTask",
			path:  "/zstack/v1/cbt-task/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "force": {"false"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteCbtTaskWithParams("parent-id", param.DeleteModePermissive, param.DeleteCbtTaskParamDetail{Force: &disabled})
			},
		},
		{
			name:  "DeleteCdpTask",
			path:  "/zstack/v1/cdp-task/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "force": {"false"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteCdpTaskWithParams("parent-id", param.DeleteModePermissive, param.DeleteCdpTaskParamDetail{Force: &disabled})
			},
		},
		{
			name:  "DeleteDatabaseBackup",
			path:  "/zstack/v1/database-backups/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "backupStorageUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteDatabaseBackupWithParams("parent-id", param.DeleteModePermissive, param.DeleteDatabaseBackupParamDetail{BackupStorageUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "DeleteHuaweiIMasterFabric",
			path:  "/zstack/v1/sdn-controller/huawei-imaster/fabrics/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "sdnControllerUuid": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteHuaweiIMasterFabricWithParams("parent-id", param.DeleteModePermissive, param.DeleteHuaweiIMasterFabricParamDetail{SdnControllerUuid: &option})
			},
		},
		{
			name:  "DeleteHuaweiIMasterTenant",
			path:  "/zstack/v1/sdn-controller/huawei-imaster/tenants/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "sdnControllerUuid": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteHuaweiIMasterTenantWithParams("parent-id", param.DeleteModePermissive, param.DeleteHuaweiIMasterTenantParamDetail{SdnControllerUuid: &option})
			},
		},
		{
			name:  "DeleteHuaweiIMasterVRouter",
			path:  "/zstack/v1/sdn-controller/huawei-imaster/vrouters/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "sdnControllerUuid": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteHuaweiIMasterVRouterWithParams("parent-id", param.DeleteModePermissive, param.DeleteHuaweiIMasterVRouterParamDetail{SdnControllerUuid: &option})
			},
		},
		{
			name:  "DeleteHuaweiIMasterVpc",
			path:  "/zstack/v1/sdn-controller/huawei-imaster/vpcs/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "sdnControllerUuid": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteHuaweiIMasterVpcWithParams("parent-id", param.DeleteModePermissive, param.DeleteHuaweiIMasterVpcParamDetail{SdnControllerUuid: &option})
			},
		},
		{
			name:  "DeleteImage",
			path:  "/zstack/v1/images/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "backupStorageUuids": {"value-a", "value-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteImageWithParams("parent-id", param.DeleteModePermissive, param.DeleteImageParamDetail{BackupStorageUuids: []string{"value-a", "value-b"}})
			},
		},
		{
			name:  "DeleteResourcePrice",
			path:  "/zstack/v1/billings/prices/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "cutoffPrice": {"false"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteResourcePriceWithParams("parent-id", param.DeleteModePermissive, param.DeleteResourcePriceParamDetail{CutoffPrice: &disabled})
			},
		},
		{
			name:  "DeleteVmBackup",
			path:  "/zstack/v1/vm-backups/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "backupStorageUuids": {"value-a", "value-b"}, "handleDependency": {"false"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteVmBackupWithParams("parent-id", param.DeleteModePermissive, param.DeleteVmBackupParamDetail{BackupStorageUuids: []string{"value-a", "value-b"}, HandleDependency: &disabled})
			},
		},
		{
			name:  "DeleteVolumeBackup",
			path:  "/zstack/v1/volume-backups/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "backupStorageUuids": {"value-a", "value-b"}, "handleDependency": {"false"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteVolumeBackupWithParams("parent-id", param.DeleteModePermissive, param.DeleteVolumeBackupParamDetail{BackupStorageUuids: []string{"value-a", "value-b"}, HandleDependency: &disabled})
			},
		},
		{
			name:  "DeleteVolumeSnapshot",
			path:  "/zstack/v1/volume-snapshots/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "direction": {"option"}, "scope": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteVolumeSnapshotWithParams("parent-id", param.DeleteModePermissive, param.DeleteVolumeSnapshotParamDetail{Direction: &option, Scope: &option})
			},
		},
		{
			name:  "DeleteVolumeSnapshotGroup",
			path:  "/zstack/v1/volume-snapshots/group/parent-id",
			query: url.Values{"deleteMode": {"Permissive"}, "direction": {"option"}, "scope": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteVolumeSnapshotGroupWithParams("parent-id", param.DeleteModePermissive, param.DeleteVolumeSnapshotGroupParamDetail{Direction: &option, Scope: &option})
			},
		},
		{
			name:  "RemoveSNSSmsReceiver",
			path:  "/zstack/v1/sns/sms-endpoints/parent-id/receivers/+8613800000010",
			query: url.Values{"deleteMode": {"Permissive"}, "phoneNumberList": {"+8613800000010", "+8613800000011"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveSNSSmsReceiverWithParams("parent-id", "+8613800000010", param.DeleteModePermissive, param.RemoveSNSSmsReceiverParamDetail{PhoneNumberList: []string{"+8613800000010", "+8613800000011"}})
			},
		},
		{
			name:  "SdnControllerRemoveHost",
			path:  "/zstack/v1/sdn-controllers/parent-id/hosts/host-id",
			query: url.Values{"deleteMode": {"Permissive"}, "vSwitchType": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.SdnControllerRemoveHostWithParams("parent-id", "host-id", param.DeleteModePermissive, param.SdnControllerRemoveHostParamDetail{VSwitchType: &option})
			},
		},
		{
			name:  "network service map",
			path:  "/zstack/v1/l3-networks/network-id/network-services",
			query: url.Values{"deleteMode": {"Permissive"}, "networkServices.provider-a": {"DHCP", "DNS"}, "networkServices.provider-b": {"SNAT"}},
			call: func(cli *ZSClient) error {
				return cli.DetachNetworkServiceFromL3NetworkWithParams("network-id", param.DeleteModePermissive, param.DetachNetworkServiceFromL3NetworkParamDetail{NetworkServices: map[string]any{"provider-a": []string{"DHCP", "DNS"}, "provider-b": []any{"SNAT"}}})
			},
		},
		{
			name:  "network service by name",
			path:  "/zstack/v1/l3-networks/network-id/network-services/DHCP",
			query: url.Values{"deleteMode": {"Permissive"}},
			call: func(cli *ZSClient) error {
				return cli.DetachNetworkServiceFromL3NetworkWithParams("network-id", param.DeleteModePermissive, param.DetachNetworkServiceFromL3NetworkParamDetail{Service: &service})
			},
		},
		{
			name:  "logout client info",
			path:  "/zstack/v1/accounts/sessions/session-id",
			query: url.Values{"deleteMode": {"Permissive"}, "clientInfo.clientBrowser": {"Browser + & details"}, "clientInfo.clientIp": {"192.0.2.1"}},
			call: func(cli *ZSClient) error {
				return cli.LogOutWithParams("session-id", param.DeleteModePermissive, param.LogOutParamDetail{ClientInfo: map[string]string{"clientBrowser": "Browser + & details", "clientIp": "192.0.2.1"}})
			},
		},
		{
			name:  "single SMS receiver",
			path:  "/zstack/v1/sns/sms-endpoints/endpoint-id/receivers/+8613800000010",
			query: url.Values{"deleteMode": {"Permissive"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveSNSSmsReceiver("endpoint-id", "+8613800000010", param.DeleteModePermissive)
			},
		},
		{
			name:  "SMS list without path number",
			path:  "/zstack/v1/sns/sms-endpoints/endpoint-id/receivers/+8613800000010",
			query: url.Values{"deleteMode": {"Permissive"}, "phoneNumberList": {"+8613800000010", "+8613800000011"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveSNSSmsReceiverWithParams("endpoint-id", "", param.DeleteModePermissive, param.RemoveSNSSmsReceiverParamDetail{PhoneNumberList: []string{"+8613800000010", "+8613800000011"}})
			},
		},
		{
			name:  "single model",
			path:  "/zstack/v1/ai/models",
			query: url.Values{"deleteMode": {"Permissive"}, "uuids": {"model-id"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteModels("model-id", param.DeleteModePermissive)
			},
		},
		{
			name:  "legacy log configuration",
			path:  "/zstack/v1/log/configurations/log4j2",
			query: url.Values{"deleteMode": {"Permissive"}, "configId": {"922337203685477580"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteLogConfiguration("922337203685477580", param.DeleteModePermissive)
			},
		},
		{
			name:  "default image delete",
			path:  "/zstack/v1/images/image-id",
			query: url.Values{"deleteMode": {"Permissive"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteImage("image-id", param.DeleteModePermissive)
			},
		},
		{
			name:  "detail delete mode",
			path:  "/zstack/v1/images/image-id",
			query: url.Values{"deleteMode": {"Enforcing"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteImageWithParams("image-id", "", param.DeleteImageParamDetail{DeleteMode: &enforcing})
			},
		},
		{
			name:  "explicit delete mode precedence",
			path:  "/zstack/v1/images/image-id",
			query: url.Values{"deleteMode": {"Permissive"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteImageWithParams("image-id", param.DeleteModePermissive, param.DeleteImageParamDetail{DeleteMode: &enforcing})
			},
		},
		{
			name:  "license by module",
			path:  "/zstack/v1/licenses/mn/node-id/actions",
			query: url.Values{"deleteMode": {"Permissive"}, "module": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteLicenseWithParams("node-id", param.DeleteModePermissive, param.DeleteLicenseParamDetail{Module: &option})
			},
		},
		{
			name:  "whole billing explicitly requested",
			path:  "/zstack/v1/billings/billings",
			query: url.Values{},
			call: func(cli *ZSClient) error {
				return cli.DeleteAllBilling("")
			},
		},
		{
			name:  "unregister server",
			path:  "/zstack/v1/license-server/client",
			query: url.Values{},
			call: func(cli *ZSClient) error {
				return cli.UnregisterLicenseServer("", "")
			},
		},
		{
			name:  "unregister client node",
			path:  "/zstack/v1/license-server/client",
			query: url.Values{"clientAuthorizedNodeUuid": {"node-id"}},
			call: func(cli *ZSClient) error {
				return cli.UnregisterLicenseServer("node-id", "")
			},
		},
	}
	assertDeleteActionContracts(t, tests)
}
