// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

func (cli *ZSClient) AckAlarmData(params param.AckAlarmDataParam) (*view.AlertDataAckInventoryView, error) {
	var resp view.AckAlertDataEventView
	if err := cli.Post("v1/zwatch/alarm-histories/acknowledgments", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) AckAlertData(params param.AckAlertDataParam) (*view.AlertDataAckInventoryView, error) {
	var resp view.AckAlertDataEventView
	if err := cli.Post("v1/zwatch/alert-histories/acknowledgments", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) AckEventData(params param.AckEventDataParam) (*view.AlertDataAckInventoryView, error) {
	var resp view.AckAlertDataEventView
	if err := cli.Post("v1/zwatch/event-histories/acknowledgments", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) AllocateHostResource(params param.AllocateHostResourceParam) (*view.AllocateHostResourceEventView, error) {
	resp := view.AllocateHostResourceEventView{}
	if err := cli.Post("v1/hosts/{uuid}/allocate-resource", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ApplyDRSAdvice(uuid string, params param.ApplyDRSAdviceParam) (*view.ApplyDRSAdviceEventView, error) {
	resp := view.ApplyDRSAdviceEventView{}
	if err := cli.Put("v1/clusters/drs/advice/{adviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ApplyMonitorTemplateToMonitorGroup(params param.ApplyMonitorTemplateToMonitorGroupParam) (*view.MonitorGroupTemplateRefInventoryView, error) {
	var resp view.ApplyMonitorTemplateToMonitorGroupEventView
	if err := cli.Post("v1/zwatch/monitortemplates/{templateUuid}/monitorgroups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ApplyRuleSetChanges(uuid string, params param.ApplyRuleSetChangesParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	var resp view.ApplyRuleSetChangesEventView
	if err := cli.Put("v1/vpcfirewalls/ruleSets/apply/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ApplyTemplateConfig(uuid string, params param.ApplyTemplateConfigParam) (*view.ApplyTemplateConfigEventView, error) {
	resp := view.ApplyTemplateConfigEventView{}
	if err := cli.Put("v1/template-configurations/{templateUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) BackupDatabaseToPublicCloud(params param.BackupDatabaseToPublicCloudParam) (*view.BackupDatabaseToPublicCloudEventView, error) {
	resp := view.BackupDatabaseToPublicCloudEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) BackupStorageMigrateImage(uuid string, params param.BackupStorageMigrateImageParam) (*view.ImageInventoryView, error) {
	var resp view.BackupStorageMigrateImageEventView
	if err := cli.Put("v1/backup-storage/images/{imageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) BatchAddBareMetal2IpmiChassis(params param.BatchAddBareMetal2IpmiChassisParam) (*view.LongJobInventoryView, error) {
	var resp view.BatchAddBareMetal2ChassisEventView
	if err := cli.Post("v1/baremetal2/chassis/ipmi/from-file", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) BatchCreateBaremetalChassis(params param.BatchCreateBaremetalChassisParam) (*view.LongJobInventoryView, error) {
	var resp view.BatchCreateBaremetalChassisEventView
	if err := cli.Post("v1/baremetal/chassis/from-file", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) BatchCreateIAM2VirtualIDFromConfigFile(params param.BatchCreateIAM2VirtualIDFromConfigFileParam) (*view.BatchCreateIAM2VirtualIDFromConfigFileEventView, error) {
	resp := view.BatchCreateIAM2VirtualIDFromConfigFileEventView{}
	if err := cli.Post("v1/iam2/virtual-ids/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) BatchDeleteVolumeSnapshot(uuid string, params param.BatchDeleteVolumeSnapshotParam) (*view.BatchDeleteVolumeSnapshotEventView, error) {
	resp := view.BatchDeleteVolumeSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/batch-delete", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) BatchQuery(params param.BatchQueryParam) (*view.BatchQueryView, error) {
	var resp view.BatchQueryView
	if err := cli.Get("v1/batch-queries", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) BatchSyncVolumeSize(params param.BatchSyncVolumeSizeParam) (*view.BatchSyncVolumeSizeView, error) {
	resp := view.BatchSyncVolumeSizeView{}
	if err := cli.Post("v1/volumes/batch-sync-volumes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) BindModelToService(params param.BindModelToServiceParam) (*view.ModelServiceInventoryView, error) {
	var resp view.BindModelToServiceEventView
	if err := cli.Post("v1/ai/models/{modelUuid}/model-services/{modelServiceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) BootstrapMiniHost(params param.BootstrapMiniHostParam) (*view.BootstrapMiniHostEventView, error) {
	resp := view.BootstrapMiniHostEventView{}
	if err := cli.Post("v1/mini-clusters/hosts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CalculateAccountBillingSpending(uuid string, params param.CalculateAccountBillingSpendingParam) (*view.CalculateAccountBillingSpendingView, error) {
	resp := view.CalculateAccountBillingSpendingView{}
	if err := cli.Put("v1/billings/accounts/{accountUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CalculateAccountSpending(uuid string, params param.CalculateAccountSpendingParam) (*view.CalculateAccountSpendingView, error) {
	resp := view.CalculateAccountSpendingView{}
	if err := cli.Put("v1/billings/accounts/{accountUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CalculateImageHash(uuid string, params param.CalculateImageHashParam) (*view.ImageInventoryView, error) {
	var resp view.CalculateImageHashEventView
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) CalculateResourceSpending(uuid string, params param.CalculateResourceSpendingParam) (*view.CalculateResourceSpendingView, error) {
	resp := view.CalculateResourceSpendingView{}
	if err := cli.Put("v1/billings/resources/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CancelLongJob(uuid string, params param.CancelLongJobParam) (*view.CancelLongJobEventView, error) {
	resp := view.CancelLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckApiPermission(uuid string, params param.CheckApiPermissionParam) (*view.MapView, error) {
	var resp view.CheckApiPermissionView
	if err := cli.Put("v1/accounts/permissions/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) CheckBareMetal2IpmiChassisConfigFile(params param.CheckBareMetal2IpmiChassisConfigFileParam) (*view.CheckBareMetal2ChassisConfigFileView, error) {
	resp := view.CheckBareMetal2ChassisConfigFileView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckBaremetalChassisConfigFile(params param.CheckBaremetalChassisConfigFileParam) (*view.CheckBaremetalChassisConfigFileView, error) {
	resp := view.CheckBaremetalChassisConfigFileView{}
	if err := cli.Post("v1/baremetal/chassis/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckBatchDataIntegrity(params param.CheckBatchDataIntegrityParam) (*view.CheckBatchDataIntegrityView, error) {
	var resp view.CheckBatchDataIntegrityView
	if err := cli.Get("v1/check/batch/data/integrity/", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckBuildAppParameters(params param.CheckBuildAppParametersParam) (*view.CheckBuildAppParametersView, error) {
	resp := view.CheckBuildAppParametersView{}
	if err := cli.Post("v1/appcenter/buildapp/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckElaborationContent(params param.CheckElaborationContentParam) (*view.CheckElaborationContentView, error) {
	resp := view.CheckElaborationContentView{}
	if err := cli.Post("v1/errorcode/elaborations/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckFirewallRuleConfigFile(params param.CheckFirewallRuleConfigFileParam) (*view.CheckFirewallRuleConfigFileView, error) {
	resp := view.CheckFirewallRuleConfigFileView{}
	if err := cli.Post("v1/vpcfirewalls/rules/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckIAM2OrganizationAvailability(params param.CheckIAM2OrganizationAvailabilityParam) (*view.CheckIAM2OrganizationAvailabilityView, error) {
	var resp view.CheckIAM2OrganizationAvailabilityView
	if err := cli.Get("v1/iam2/organizations/availabilities", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckIAM2VirtualIDConfigFile(uuid string, params param.CheckIAM2VirtualIDConfigFileParam) (*view.CheckIAM2VirtualIDConfigFileView, error) {
	resp := view.CheckIAM2VirtualIDConfigFileView{}
	if err := cli.Put("v1/iam2/virtual-ids/from-file", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckIpAvailability(params param.CheckIpAvailabilityParam) (*view.CheckIpAvailabilityView, error) {
	var resp view.CheckIpAvailabilityView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/ip/{ip}/availability", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckKVMHostConfigFile(params param.CheckKVMHostConfigFileParam) (*view.CheckHostConfigFileView, error) {
	resp := view.CheckHostConfigFileView{}
	if err := cli.Post("v1/hosts/kvm/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckNetworkReachable(params param.CheckNetworkReachableParam) (*view.CheckNetworkReachableView, error) {
	var resp view.CheckNetworkReachableView
	if err := cli.Get("v1/zops/check/network", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckResourcePermission(params param.CheckResourcePermissionParam) (*view.CheckResourcePermissionView, error) {
	var resp view.CheckResourcePermissionView
	if err := cli.Get("v1/accounts/resource/api-permissions", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckScsiLunClusterStatus(uuid string, params param.CheckScsiLunClusterStatusParam) (*view.ScsiLunClusterStatusInventoryView, error) {
	var resp view.CheckScsiLunClusterStatusView
	if err := cli.Put("v1/storage-devices/scsi-lun/{uuid}/cluster/{clusterUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) CheckStackTemplateParameters(params param.CheckStackTemplateParametersParam) (*view.CheckStackTemplateParametersView, error) {
	resp := view.CheckStackTemplateParametersView{}
	if err := cli.Post("v1/cloudformation/stack/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckStaticProvisionIp(params param.CheckStaticProvisionIpParam) (*view.CheckStaticProvisionIpView, error) {
	resp := view.CheckStaticProvisionIpView{}
	if err := cli.Post("v1/baremetal2/bm-instances/static/provision/ip/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckVipPortAvailability(params param.CheckVipPortAvailabilityParam) (*view.CheckVipPortAvailabilityView, error) {
	var resp view.CheckVipPortAvailabilityView
	if err := cli.Get("v1/vips/{vipUuid}/check-port-availability", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CheckVolumeSnapshotGroupAvailability(params param.CheckVolumeSnapshotGroupAvailabilityParam) (*view.CheckVolumeSnapshotGroupAvailabilityView, error) {
	var resp view.CheckVolumeSnapshotGroupAvailabilityView
	if err := cli.Get("v1/volume-snapshots/groups/availabilities", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanInvalidLdapBinding(uuid string, params param.CleanInvalidLdapBindingParam) (*view.AccountInventoryView, error) {
	resp := view.AccountInventoryView{}
	if err := cli.Put("v1/ldap/bindings/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanInvalidLdapIAM2Binding(uuid string, params param.CleanInvalidLdapIAM2BindingParam) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.Put("v1/iam2/ldap/bindings/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanLongJob(uuid string, params param.CleanLongJobParam) (*view.CleanLongJobEventView, error) {
	resp := view.CleanLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanQueue(uuid string, params param.CleanQueueParam) (*view.CleanQueueEventView, error) {
	resp := view.CleanQueueEventView{}
	if err := cli.Put("v1/clean/queue", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanUpBareMetal2Bonding(uuid string, params param.CleanUpBareMetal2BondingParam) (*view.CleanUpBaremetal2BondingEventView, error) {
	resp := view.CleanUpBaremetal2BondingEventView{}
	if err := cli.Put("v1/baremetal2/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanUpBaremetalChassisBonding(uuid string, params param.CleanUpBaremetalChassisBondingParam) (*view.CleanUpBaremetalChassisBondingEventView, error) {
	resp := view.CleanUpBaremetalChassisBondingEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanUpImageCacheOnPrimaryStorage(uuid string, params param.CleanUpImageCacheOnPrimaryStorageParam) (*view.CleanUpImageCacheOnPrimaryStorageEventView, error) {
	resp := view.CleanUpImageCacheOnPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanUpStorageTrashOnPrimaryStorage(uuid string, params param.CleanUpStorageTrashOnPrimaryStorageParam) (*view.CleanUpStorageTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpStorageTrashOnPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/storagetrash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanUpTrashOnBackupStorage(uuid string, params param.CleanUpTrashOnBackupStorageParam) (*view.CleanUpTrashOnBackupStorageEventView, error) {
	resp := view.CleanUpTrashOnBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/trash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanUpTrashOnPrimaryStorage(uuid string, params param.CleanUpTrashOnPrimaryStorageParam) (*view.CleanUpTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpTrashOnPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/trash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanV2VConversionCache(uuid string, params param.CleanV2VConversionCacheParam) (*view.CleanV2VConversionCacheEventView, error) {
	resp := view.CleanV2VConversionCacheEventView{}
	if err := cli.Put("v1/v2v/conversion/host/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) CleanupBillingUsage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/usage", uuid, string(deleteMode))
}

func (cli *ZSClient) ConvertVmFromForeignHypervisor(params param.ConvertVmFromForeignHypervisorParam) (*view.LongJobInventoryView, error) {
	var resp view.ConvertVmFromForeignHypervisorEventView
	if err := cli.Post("v1/v2vs", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) DebugSignal(params param.DebugSignalParam) (*view.DebugSignalEventView, error) {
	resp := view.DebugSignalEventView{}
	if err := cli.Post("v1/debug", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) DecodeStackTemplate(params param.DecodeStackTemplateParam) (*view.DecodeStackTemplateView, error) {
	resp := view.DecodeStackTemplateView{}
	if err := cli.Post("v1/cloudformation/stack/preview/resource", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) DegradeFromLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server", uuid, string(deleteMode))
}

func (cli *ZSClient) DeployAppDevelopmentService(uuid string, params param.DeployAppDevelopmentServiceParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.DeployAppDevelopmentServiceEventView
	if err := cli.Put("v1/ai/model-services/app/", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) DeployDistributedModelService(uuid string, params param.DeployDistributedModelServiceParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.DeployDistributedModelServiceEventView
	if err := cli.Put("v1/ai/model-services", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) DeployModelEvalService(uuid string, params param.DeployModelEvalServiceParam) (*view.ModelEvalServiceInstanceGroupInventoryView, error) {
	var resp view.DeployModelEvalServiceEventView
	if err := cli.Put("v1/ai/model-services/eval/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) DeployModelService(uuid string, params param.DeployModelServiceParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.DeployModelServiceEventView
	if err := cli.Put("v1/ai/model-services/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) DescribeVmInstanceRecoveryPoint(params param.DescribeVmInstanceRecoveryPointParam) (*view.DescribeVmInstanceRecoveryPointView, error) {
	var resp view.DescribeVmInstanceRecoveryPointView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/recovery-point", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) DisableCbtTask(params param.DisableCbtTaskParam) (*view.CbtTaskInventoryView, error) {
	var resp view.DisableCbtTaskEventView
	if err := cli.Post("v1/cbt-task/disable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) DisableCdpTask(params param.DisableCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	var resp view.DisableCdpTaskEventView
	if err := cli.Post("v1/cdp-task/disable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) DiscoverExternalPrimaryStorage(params param.DiscoverExternalPrimaryStorageParam) (*view.ExternalPrimaryStorageInventoryView, error) {
	var resp view.DiscoverExternalPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/addon/discover", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) DownloadBackupFileFromPublicCloud(params param.DownloadBackupFileFromPublicCloudParam) (*view.DownloadBackupFileFromPublicCloudEventView, error) {
	resp := view.DownloadBackupFileFromPublicCloudEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql/download", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) EjectZBox(uuid string, params param.EjectZBoxParam) (*view.ZBoxInventoryView, error) {
	var resp view.EjectZBoxEventView
	if err := cli.Put("v1/zbox/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) EnableCbtTask(params param.EnableCbtTaskParam) (*view.EnableCbtTaskEventView, error) {
	resp := view.EnableCbtTaskEventView{}
	if err := cli.Post("v1/cbt-task/enable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) EnableCdpTask(params param.EnableCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	var resp view.EnableCdpTaskEventView
	if err := cli.Post("v1/cdp-task/enable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ExecuteAutoScalingRule(uuid string, params param.ExecuteAutoScalingRuleParam) (*view.ExecuteAutoScalingRuleEventView, error) {
	resp := view.ExecuteAutoScalingRuleEventView{}
	if err := cli.Put("v1/autoscaling/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ExecuteDRSScheduling(uuid string, params param.ExecuteDRSSchedulingParam) (*view.ExecuteDRSSchedulingEventView, error) {
	resp := view.ExecuteDRSSchedulingEventView{}
	if err := cli.Put("v1/clusters/drs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ExecuteGuestVmCommand(params param.ExecuteGuestVmCommandParam) (*view.ExecuteGuestVmCommandEventView, error) {
	resp := view.ExecuteGuestVmCommandEventView{}
	if err := cli.Post("v1/vm-instances/commands/exec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ExecuteGuestVmScript(uuid string, params param.ExecuteGuestVmScriptParam) (*view.GuestVmScriptExecutedRecordInventoryView, error) {
	var resp view.ExecuteGuestVmScriptEventView
	if err := cli.Put("v1/scripts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ExportBuildApp(uuid string, params param.ExportBuildAppParam) (*view.BuildAppExportHistoryInventoryView, error) {
	var resp view.ExportBuildAppEventView
	if err := cli.Put("v1/appcenter/buildapp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ExportDatabaseBackupFromBackupStorage(uuid string, params param.ExportDatabaseBackupFromBackupStorageParam) (*view.ExportDatabaseBackupFromBackupStorageEventView, error) {
	resp := view.ExportDatabaseBackupFromBackupStorageEventView{}
	if err := cli.Put("v1/database-backups/{databaseBackupUuid}/backup-storage/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ExportImageFromBackupStorage(uuid string, params param.ExportImageFromBackupStorageParam) (*view.ExportImageFromBackupStorageEventView, error) {
	resp := view.ExportImageFromBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ExportNbdVolumes(params param.ExportNbdVolumesParam) (*view.ExportNbdVolumesEventView, error) {
	resp := view.ExportNbdVolumesEventView{}
	if err := cli.Post("v1/cbt-task/exportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ExportVmOvaPackage(params param.ExportVmOvaPackageParam) (*view.ImagePackageInventoryView, error) {
	var resp view.ExportVmOvaPackageEventView
	if err := cli.Post("v1/ovf/ova-packages", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) FailoverFaultToleranceVm(uuid string, params param.FailoverFaultToleranceVmParam) (*view.FailoverFaultToleranceVmEventView, error) {
	resp := view.FailoverFaultToleranceVmEventView{}
	if err := cli.Put("v1/vm-instances/fault-tolerance", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) FlattenVmInstance(uuid string, params param.FlattenVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.FlattenVmInstanceEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) FlattenVolume(uuid string, params param.FlattenVolumeParam) (*view.VolumeInventoryView, error) {
	var resp view.FlattenVolumeEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) FstrimVm(params param.FstrimVmParam) (*view.FstrimVmEventView, error) {
	resp := view.FstrimVmEventView{}
	if err := cli.Post("v1/vm-instances/{uuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) GCAliyunSnapshotRemote(params param.GCAliyunSnapshotRemoteParam) (*view.GCAliyunSnapshotRemoteEventView, error) {
	resp := view.GCAliyunSnapshotRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/snapshot/{dataCenterUuid}/gc", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) GenerateAccountBilling(uuid string, params param.GenerateAccountBillingParam) (*view.GenerateAccountBillingEventView, error) {
	resp := view.GenerateAccountBillingEventView{}
	if err := cli.Put("v1/billings/accounts/{accountUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) GenerateHygonMdevDevices(uuid string, params param.GenerateHygonMdevDevicesParam) (*view.GenerateHygonMdevDevicesEventView, error) {
	resp := view.GenerateHygonMdevDevicesEventView{}
	if err := cli.Put("v1/hygon-devices/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) GenerateMdevDevices(uuid string, params param.GenerateMdevDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) GenerateModelMetadata(uuid string, params param.GenerateModelMetadataParam) (*view.GenerateModelMetadataEventView, error) {
	resp := view.GenerateModelMetadataEventView{}
	if err := cli.Put("v1/ai/model/metadata/generate", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) GenerateSeMdevDevices(uuid string, params param.GenerateSeMdevDevicesParam) (*view.GenerateSeMdevDevicesEventView, error) {
	resp := view.GenerateSeMdevDevicesEventView{}
	if err := cli.Put("v1/mtty-devices/{mttyDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) GenerateSriovPciDevices(uuid string, params param.GenerateSriovPciDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) GenerateSshKeyPair(params param.GenerateSshKeyPairParam) (*view.SshPrivateKeyPairInventoryView, error) {
	var resp view.GenerateSshKeyPairView
	if err := cli.Post("v1/ssh-key-pair/generate", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) IdentifyHost(uuid string, params param.IdentifyHostParam) (*view.IdentifyHostEventView, error) {
	resp := view.IdentifyHostEventView{}
	if err := cli.Put("v1/hosts/kvm/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) InspectBareMetal2ChassisByInstance(uuid string, params param.InspectBareMetal2ChassisByInstanceParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.InspectBareMetal2ChassisByInstanceEventView
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) InspectBareMetal2Chassis(uuid string, params param.InspectBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.InspectBareMetal2ChassisEventView
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) InspectBaremetalChassis(uuid string, params param.InspectBaremetalChassisParam) (*view.BaremetalChassisInventoryView, error) {
	var resp view.InspectBaremetalChassisEventView
	if err := cli.Put("v1/baremetal/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) IsLicenseServer(params param.IsLicenseServerParam) (*view.IsLicenseServerView, error) {
	var resp view.IsLicenseServerView
	if err := cli.Get("v1/license-server/is-server", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) IsOpensourceVersion(params param.IsOpensourceVersionParam) (*view.IsOpensourceVersionView, error) {
	var resp view.IsOpensourceVersionView
	if err := cli.Get("v1/meta-data/opensource", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) IsReadyToGo(params param.IsReadyToGoParam) (*view.IsReadyToGoView, error) {
	var resp view.IsReadyToGoView
	if err := cli.Get("v1/management-nodes/ready", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) IsVfNicAvailableInL3Network(params param.IsVfNicAvailableInL3NetworkParam) (*view.IsVfNicAvailableInL3NetworkView, error) {
	var resp view.IsVfNicAvailableInL3NetworkView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/hosts/{hostUuid}/vfnicavailable", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) KvmRunShell(uuid string, params param.KvmRunShellParam) (*view.MapView, error) {
	var resp view.KvmRunShellEventView
	if err := cli.Put("v1/hosts/kvm/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ListVMsFromKVMHost(params param.ListVMsFromKVMHostParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post("v1/v2v", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ListVmSchedulingRulesFromExecuteState(params param.ListVmSchedulingRulesFromExecuteStateParam) (*view.ListVmSchedulingRulesFromExecuteStateView, error) {
	resp := view.ListVmSchedulingRulesFromExecuteStateView{}
	if err := cli.Post("v1/list/vmSchedulingRules/from/conflict/state", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ListVmsFromSchedulingState(params param.ListVmsFromSchedulingStateParam) (*view.ListVmsFromSchedulingStateView, error) {
	resp := view.ListVmsFromSchedulingStateView{}
	if err := cli.Post("v1/list/vms/from/executeState", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) LocalStorageGetVolumeMigratableHosts(params param.LocalStorageGetVolumeMigratableHostsParam) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.Get("v1/volumes/{volumeUuid}/migration-target-hosts", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) LocalStorageMigrateVolume(uuid string, params param.LocalStorageMigrateVolumeParam) (*view.LocalStorageResourceRefInventoryView, error) {
	var resp view.LocalStorageMigrateVolumeEventView
	if err := cli.Put("v1/primary-storage/local-storage/volumes/{volumeUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) LocateHostNetworkInterface(uuid string, params param.LocateHostNetworkInterfaceParam) (*view.LocateHostNetworkInterfaceEventView, error) {
	resp := view.LocateHostNetworkInterfaceEventView{}
	if err := cli.Put("v1/hosts/{hostUuid}/locate/network-interface", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) LocateLocalRaidPhysicalDrive(uuid string, params param.LocateLocalRaidPhysicalDriveParam) (*view.RaidPhysicalDriveInventoryView, error) {
	var resp view.LocateLocalRaidPhysicalDriveEventView
	if err := cli.Put("v1/storage-devices/local-raid/physical-drives/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) LogInByAccount(uuid string, params param.LogInByAccountParam) (*view.SessionInventoryView, error) {
	var resp view.LogInView
	if err := cli.Put("v1/accounts/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) LogInByLdap(uuid string, params param.LogInByLdapParam) (*view.SessionInventoryView, error) {
	var resp view.LogInByLdapView
	if err := cli.Put("v1/ldap/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) LogInByUser(uuid string, params param.LogInByUserParam) (*view.SessionInventoryView, error) {
	var resp view.LogInView
	if err := cli.Put("v1/accounts/users/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) LogIn(uuid string, params param.LogInParam) (*view.SessionInventoryView, error) {
	var resp view.LogInView
	if err := cli.Put("v1/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) LogOut(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/sessions/{sessionUuid}", uuid, string(deleteMode))
}

func (cli *ZSClient) LoginByCas(uuid string, params param.LoginByCasParam) (*view.SessionInventoryView, error) {
	var resp view.LoginByCasView
	if err := cli.Put("v1/cas/login/", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) LoginIAM2Platform(uuid string, params param.LoginIAM2PlatformParam) (*view.SessionInventoryView, error) {
	var resp view.LoginIAM2PlatformView
	if err := cli.Put("v1/iam2/platform/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) LoginIAM2Project(uuid string, params param.LoginIAM2ProjectParam) (*view.SessionInventoryView, error) {
	var resp view.LoginIAM2ProjectView
	if err := cli.Put("v1/iam2/projects/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) LoginIAM2VirtualID(uuid string, params param.LoginIAM2VirtualIDParam) (*view.SessionInventoryView, error) {
	var resp view.LoginIAM2VirtualIDView
	if err := cli.Put("v1/iam2/virtual-ids/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) LoginIAM2VirtualIDWithLdap(uuid string, params param.LoginIAM2VirtualIDWithLdapParam) (*view.LoginIAM2VirtualIDWithLdapView, error) {
	resp := view.LoginIAM2VirtualIDWithLdapView{}
	if err := cli.Put("v1/iam2/login/virtual-ids/ldap", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) MatchModelServiceTemplateWithModel(params param.MatchModelServiceTemplateWithModelParam) (*view.MatchModelServiceTemplateWithModelEventView, error) {
	resp := view.MatchModelServiceTemplateWithModelEventView{}
	if err := cli.Post("v1/ai/model-services/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) MergeDataOnBackupStorage(uuid string, params param.MergeDataOnBackupStorageParam) (*view.MergeDataOnBackupStorageEventView, error) {
	resp := view.MergeDataOnBackupStorageEventView{}
	if err := cli.Put("v1/cdp-task/mergedata/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) MountVmInstanceRecoveryPoint(params param.MountVmInstanceRecoveryPointParam) (*view.MountVmInstanceRecoveryPointEventView, error) {
	resp := view.MountVmInstanceRecoveryPointEventView{}
	if err := cli.Post("v1/cdp-backup-storage/mount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) MoveDirectory(uuid string, params param.MoveDirectoryParam) (*view.MoveDirectoryEventView, error) {
	resp := view.MoveDirectoryEventView{}
	if err := cli.Put("v1/move/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) MoveResourcesToDirectory(uuid string, params param.MoveResourcesToDirectoryParam) (*view.MoveResourcesToDirectoryEventView, error) {
	resp := view.MoveResourcesToDirectoryEventView{}
	if err := cli.Put("v1/move/resources/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ParseOvf(params param.ParseOvfParam) (*view.ParseOvfView, error) {
	resp := view.ParseOvfView{}
	if err := cli.Post("v1/ovf/parse", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PauseVmInstance(uuid string, params param.PauseVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.PauseVmInstanceEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) PowerOffBareMetal2Chassis(uuid string, params param.PowerOffBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.PowerOffBareMetal2ChassisEventView
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) PowerOffBaremetalChassis(uuid string, params param.PowerOffBaremetalChassisParam) (*view.PowerOffBaremetalChassisEventView, error) {
	resp := view.PowerOffBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PowerOffHost(uuid string, params param.PowerOffHostParam) (*view.PowerOffHostEventView, error) {
	resp := view.PowerOffHostEventView{}
	if err := cli.Put("v1/hosts/power-off/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PowerOnBareMetal2Chassis(uuid string, params param.PowerOnBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.PowerOnBareMetal2ChassisEventView
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) PowerOnBaremetalChassis(uuid string, params param.PowerOnBaremetalChassisParam) (*view.PowerOnBaremetalChassisEventView, error) {
	resp := view.PowerOnBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PowerOnHost(uuid string, params param.PowerOnHostParam) (*view.HostInventoryView, error) {
	var resp view.PowerOnHostEventView
	if err := cli.Put("v1/hosts/power/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) PowerResetBareMetal2Chassis(uuid string, params param.PowerResetBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.PowerResetBareMetal2ChassisEventView
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) PowerResetBaremetalChassis(uuid string, params param.PowerResetBaremetalChassisParam) (*view.PowerResetBaremetalChassisEventView, error) {
	resp := view.PowerResetBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PowerResetHost(uuid string, params param.PowerResetHostParam) (*view.HostInventoryView, error) {
	var resp view.PowerResetHostEventView
	if err := cli.Put("v1/hosts/power/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) PreviewResourceFromApp(params param.PreviewResourceFromAppParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post("v1/appcenter/app/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PreviewResourceStack(params param.PreviewResourceStackParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post("v1/cloudformation/stack/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PrimaryStorageMigrateVm(uuid string, params param.PrimaryStorageMigrateVmParam) (*view.VmInstanceInventoryView, error) {
	var resp view.PrimaryStorageMigrateVmEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) PrimaryStorageMigrateVolume(uuid string, params param.PrimaryStorageMigrateVolumeParam) (*view.VolumeInventoryView, error) {
	var resp view.PrimaryStorageMigrateVolumeEventView
	if err := cli.Put("v1/primary-storage/volumes/{volumeUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) PrometheusQueryLabelValues(params param.PrometheusQueryLabelValuesParam) (*view.PrometheusQueryLabelValuesView, error) {
	var resp view.PrometheusQueryLabelValuesView
	if err := cli.Get("v1/prometheus/labels", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PrometheusQueryMetadata(params param.PrometheusQueryMetadataParam) (*view.PrometheusQueryMetadataView, error) {
	var resp view.PrometheusQueryMetadataView
	if err := cli.Get("v1/prometheus/meta-data", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PrometheusQueryPassThrough(params param.PrometheusQueryPassThroughParam) (*view.PrometheusQueryPassThroughView, error) {
	var resp view.PrometheusQueryPassThroughView
	if err := cli.Get("v1/prometheus/all", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PrometheusQueryVmMonitoringData(params param.PrometheusQueryVmMonitoringDataParam) (*view.PrometheusQueryVmMonitoringDataView, error) {
	var resp view.PrometheusQueryVmMonitoringDataView
	if err := cli.Get("v1/prometheus/vm-instances", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ProtectVmInstanceRecoveryPoint(uuid string, params param.ProtectVmInstanceRecoveryPointParam) (*view.ProtectVmInstanceRecoveryPointEventView, error) {
	resp := view.ProtectVmInstanceRecoveryPointEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/protect-recovery-point", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ProvisionNfvInstConfig(uuid string, params param.ProvisionNfvInstConfigParam) (*view.ApplianceVmInventoryView, error) {
	var resp view.ProvisionNfvInstConfigEventView
	if err := cli.Put("v1/vm-instances/appliances/nfvinst/{vmInstanceUuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ProvisionNfvInstGroup(uuid string, params param.ProvisionNfvInstGroupParam) (*view.NfvInstGroupInventoryView, error) {
	var resp view.ProvisionNfvInstGroupEventView
	if err := cli.Put("v1/nfvinstgroup/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ProvisionSlbInstance(uuid string, params param.ProvisionSlbInstanceParam) (*view.SlbGroupInventoryView, error) {
	var resp view.ProvisionSlbGroupInstanceEventView
	if err := cli.Put("v1/load-balancers/slb/instances/{uuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ProvisionVirtualRouterConfig(uuid string, params param.ProvisionVirtualRouterConfigParam) (*view.ApplianceVmInventoryView, error) {
	var resp view.ProvisionVirtualRouterConfigEventView
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) PublishApp(params param.PublishAppParam) (*view.PublishAppInventoryView, error) {
	var resp view.PublishAppEventView
	if err := cli.Post("v1/appcenter/app", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) PullHuaweiIMasterController(uuid string, params param.PullHuaweiIMasterControllerParam) (*view.HuaweiIMasterSdnControllerInventoryView, error) {
	resp := view.HuaweiIMasterSdnControllerInventoryView{}
	if err := cli.Put("v1/sdn-controller/huawei-imaster/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PullSdnControllerTenant(uuid string, params param.PullSdnControllerTenantParam) (*view.H3cSdnControllerTenantInventoryView, error) {
	resp := view.H3cSdnControllerTenantInventoryView{}
	if err := cli.Put("v1/sdn-controllers/{uuid}/tenant/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PushLicenseAddOnsUsage(uuid string, params param.PushLicenseAddOnsUsageParam) (*view.PushLicenseAddOnsUsageEventView, error) {
	resp := view.PushLicenseAddOnsUsageEventView{}
	if err := cli.Put("v1/licenses/addons/usage", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) PutMetricData(params param.PutMetricDataParam) (*view.PutMetricDataEventView, error) {
	resp := view.PutMetricDataEventView{}
	if err := cli.Post("v1/zwatch/metrics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ReclaimSpaceFromImageStore(uuid string, params param.ReclaimSpaceFromImageStoreParam) (*view.ReclaimSpaceFromImageStoreEventView, error) {
	resp := view.ReclaimSpaceFromImageStoreEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ReconnectAppBuildSystem(uuid string, params param.ReconnectAppBuildSystemParam) (*view.AppBuildSystemInventoryView, error) {
	var resp view.ReconnectAppBuildSystemEventView
	if err := cli.Put("v1/appcenter/buildsystem/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectBackupStorage(uuid string, params param.ReconnectBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.ReconnectBackupStorageEventView
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectBareMetal2Gateway(uuid string, params param.ReconnectBareMetal2GatewayParam) (*view.BareMetal2GatewayInventoryView, error) {
	var resp view.ReconnectBareMetal2GatewayEventView
	if err := cli.Put("v1/baremetal2/gateways/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectBareMetal2Instance(uuid string, params param.ReconnectBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	var resp view.ReconnectBareMetal2InstanceEventView
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectBaremetalPxeServer(uuid string, params param.ReconnectBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.ReconnectBaremetalPxeServerEventView
	if err := cli.Put("v1/baremetal/pxeservers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectConsoleProxyAgent(uuid string, params param.ReconnectConsoleProxyAgentParam) (*view.MapView, error) {
	var resp view.ReconnectConsoleProxyAgentEventView
	if err := cli.Put("v1/consoles/agents", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectHost(uuid string, params param.ReconnectHostParam) (*view.HostInventoryView, error) {
	var resp view.ReconnectHostEventView
	if err := cli.Put("v1/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectIPsecConnection(uuid string, params param.ReconnectIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	var resp view.ReconnectIPsecConnectionEventView
	if err := cli.Put("v1/ipsec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectImageStoreBackupStorage(uuid string, params param.ReconnectImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	var resp view.ReconnectImageStoreBackupStorageEventView
	if err := cli.Put("v1/backup-storage/image-store/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectNfvInst(uuid string, params param.ReconnectNfvInstParam) (*view.ApplianceVmInventoryView, error) {
	var resp view.ReconnectNfvInstEventView
	if err := cli.Put("v1/vm-instances/appliances/nfvinst/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectPrimaryStorage(uuid string, params param.ReconnectPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.ReconnectPrimaryStorageEventView
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectSdnController(uuid string, params param.ReconnectSdnControllerParam) (*view.SdnControllerInventoryView, error) {
	var resp view.ReconnectSdnControllerEventView
	if err := cli.Put("v1/sdn-controllers/{sdnControllerUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectSftpBackupStorage(uuid string, params param.ReconnectSftpBackupStorageParam) (*view.SftpBackupStorageInventoryView, error) {
	var resp view.ReconnectSftpBackupStorageEventView
	if err := cli.Put("v1/backup-storage/sftp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectVirtualRouter(uuid string, params param.ReconnectVirtualRouterParam) (*view.ApplianceVmInventoryView, error) {
	var resp view.ReconnectVirtualRouterEventView
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReconnectZdfs(uuid string, params param.ReconnectZdfsParam) (*view.ZdfsInventoryView, error) {
	var resp view.ReconnectZdfsEventView
	if err := cli.Put("v1/zdfs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RefreshCaptcha(params param.RefreshCaptchaParam) (*view.RefreshCaptchaView, error) {
	var resp view.RefreshCaptchaView
	if err := cli.Get("v1/captcha/refresh", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RefreshFiberChannelStorage(params param.RefreshFiberChannelStorageParam) (*view.FiberChannelStorageInventoryView, error) {
	resp := view.FiberChannelStorageInventoryView{}
	if err := cli.Post("v1/storage-devices/fiber-channel/controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RefreshFirewall(uuid string, params param.RefreshFirewallParam) (*view.VpcFirewallInventoryView, error) {
	var resp view.RefreshFirewallEventView
	if err := cli.Put("v1/vpcfirewalls/refresh/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RefreshGuestOsMetadata(uuid string, params param.RefreshGuestOsMetadataParam) (*view.RefreshGuestOsMetadataEventView, error) {
	resp := view.RefreshGuestOsMetadataEventView{}
	if err := cli.Put("v1/guest-os/metadata/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RefreshIscsiServer(params param.RefreshIscsiServerParam) (*view.IscsiServerInventoryView, error) {
	var resp view.RefreshIscsiServerEventView
	if err := cli.Post("v1/storage-devices/iscsi/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RefreshLoadBalancer(uuid string, params param.RefreshLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	var resp view.RefreshLoadBalancerEventView
	if err := cli.Put("v1/load-balancers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RefreshLocalRaid(uuid string, params param.RefreshLocalRaidParam) (*view.RaidControllerInventoryView, error) {
	resp := view.RaidControllerInventoryView{}
	if err := cli.Put("v1/storage-devices/local-raid/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RefreshNvmeTarget(params param.RefreshNvmeTargetParam) (*view.NvmeTargetInventoryView, error) {
	resp := view.NvmeTargetInventoryView{}
	if err := cli.Post("v1/storage-devices/nvme/controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RefreshPluginDrivers(uuid string, params param.RefreshPluginDriversParam) (*view.RefreshPluginDriversEventView, error) {
	resp := view.RefreshPluginDriversEventView{}
	if err := cli.Put("v1/external/plugins", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RefreshSSOServerToken(uuid string, params param.RefreshSSOServerTokenParam) (*view.SSOServerTokenInventoryView, error) {
	var resp view.RefreshSSOServerTokenEventView
	if err := cli.Put("v1/sso/server/token/refresh", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RefreshSearchIndexes(params param.RefreshSearchIndexesParam) (*view.RefreshSearchIndexesView, error) {
	var resp view.RefreshSearchIndexesView
	if err := cli.Get("v1/search/indexes/refresh", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RefreshSharedblockDeviceCapacity(params param.RefreshSharedblockDeviceCapacityParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	var resp view.RefreshSharedBlockDeviceCapacityEventView
	if err := cli.Post("v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RegisterLicenseRequestedApplication(params param.RegisterLicenseRequestedApplicationParam) (*view.RegisterLicenseRequestedApplicationEventView, error) {
	resp := view.RegisterLicenseRequestedApplicationEventView{}
	if err := cli.Post("v1/licenses/applications", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RegisterLicenseServer(params param.RegisterLicenseServerParam) (*view.RegisterLicenseServerEventView, error) {
	resp := view.RegisterLicenseServerEventView{}
	if err := cli.Post("v1/license-server/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ReimageVmInstance(uuid string, params param.ReimageVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.ReimageVmInstanceEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ReloadElaboration(uuid string, params param.ReloadElaborationParam) (*view.ReloadElaborationEventView, error) {
	resp := view.ReloadElaborationEventView{}
	if err := cli.Put("v1/errorcode/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ReloadExternalService(uuid string, params param.ReloadExternalServiceParam) (*view.ReloadExternalServiceEventView, error) {
	resp := view.ReloadExternalServiceEventView{}
	if err := cli.Put("v1/external/services", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ReloadLicense(uuid string, params param.ReloadLicenseParam) (*view.LicenseInventoryView, error) {
	var resp view.ReloadLicenseView
	if err := cli.Put("v1/licenses/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RenewSession(uuid string, params param.RenewSessionParam) (*view.SessionInventoryView, error) {
	var resp view.RenewSessionEventView
	if err := cli.Put("v1/accounts/sessions/{sessionUuid}/renew", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RequestConsoleAccess(params param.RequestConsoleAccessParam) (*view.ConsoleInventoryView, error) {
	var resp view.RequestConsoleAccessEventView
	if err := cli.Post("v1/consoles", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RequestLicenseCapacity(params param.RequestLicenseCapacityParam) (*view.LicenseAuthorizedCapacityInventoryView, error) {
	var resp view.RequestLicenseCapacityEventView
	if err := cli.Post("v1/license-server/capacity-application", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RerunLongJob(uuid string, params param.RerunLongJobParam) (*view.LongJobInventoryView, error) {
	var resp view.RerunLongJobEventView
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ResetGlobalConfig(uuid string, params param.ResetGlobalConfigParam) (*view.ResetGlobalConfigEventView, error) {
	resp := view.ResetGlobalConfigEventView{}
	if err := cli.Put("v1/global-configurations/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ResetTemplateConfig(uuid string, params param.ResetTemplateConfigParam) (*view.ResetTemplateConfigEventView, error) {
	resp := view.ResetTemplateConfigEventView{}
	if err := cli.Put("v1/template-configurations/{templateUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ResetTwoFactorAuthenticationSecret(uuid string, params param.ResetTwoFactorAuthenticationSecretParam) (*view.TwoFactorAuthenticationSecretInventoryView, error) {
	var resp view.ResetTwoFactorAuthenticationSecretEventView
	if err := cli.Put("v1/twofactorauthentication/secrets", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RestartModelServiceGroups(uuid string, params param.RestartModelServiceGroupsParam) (*view.RestartModelServiceGroupsEventView, error) {
	resp := view.RestartModelServiceGroupsEventView{}
	if err := cli.Put("v1/model-service-instance-groups", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RestartResourceStack(uuid string, params param.RestartResourceStackParam) (*view.ResourceStackInventoryView, error) {
	var resp view.RestartResourceStackEventView
	if err := cli.Put("v1/cloudformation/stack/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ResumeLongJob(uuid string, params param.ResumeLongJobParam) (*view.LongJobInventoryView, error) {
	var resp view.ResumeLongJobEventView
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) ResumeVmInstance(uuid string, params param.ResumeVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.ResumeVmInstanceEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RevertTemplateConfig(uuid string, params param.RevertTemplateConfigParam) (*view.RevertTemplateConfigEventView, error) {
	resp := view.RevertTemplateConfigEventView{}
	if err := cli.Put("v1/template-configurations/{templateUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RevertVmFromCdpBackup(uuid string, params param.RevertVmFromCdpBackupParam) (*view.RevertVmFromCdpBackupEventView, error) {
	resp := view.RevertVmFromCdpBackupEventView{}
	if err := cli.Put("v1/cdp-backups/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RevertVmFromSnapshotGroup(uuid string, params param.RevertVmFromSnapshotGroupParam) (*view.RevertVmFromSnapshotGroupEventView, error) {
	resp := view.RevertVmFromSnapshotGroupEventView{}
	if err := cli.Put("v1/volume-snapshots/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RevertVmFromVmBackup(uuid string, params param.RevertVmFromVmBackupParam) (*view.RevertVmFromVmBackupEventView, error) {
	resp := view.RevertVmFromVmBackupEventView{}
	if err := cli.Put("v1/vm-backups/{groupUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RevertVolumeFromSnapshot(uuid string, params param.RevertVolumeFromSnapshotParam) (*view.RevertVolumeFromSnapshotEventView, error) {
	resp := view.RevertVolumeFromSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RevertVolumeFromVolumeBackup(uuid string, params param.RevertVolumeFromVolumeBackupParam) (*view.RevertVolumeFromVolumeBackupEventView, error) {
	resp := view.RevertVolumeFromVolumeBackupEventView{}
	if err := cli.Put("v1/volume-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RevokeMonitorTemplateFromMonitorGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitortemplates/{templateUuid}/monitorgroups/{groupUuid}", uuid, string(deleteMode))
}

func (cli *ZSClient) RevokeResourceSharing(uuid string, params param.RevokeResourceSharingParam) (*view.RevokeResourceSharingEventView, error) {
	resp := view.RevokeResourceSharingEventView{}
	if err := cli.Put("v1/accounts/resources/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) RunIAM2Script(params param.RunIAM2ScriptParam) (*view.LongJobInventoryView, error) {
	var resp view.RunIAM2ScriptEventView
	if err := cli.Post("v1/iam2/iam2-script/run", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) RunSchedulerTrigger(uuid string, params param.RunSchedulerTriggerParam) (*view.RunSchedulerTriggerEventView, error) {
	resp := view.RunSchedulerTriggerEventView{}
	if err := cli.Put("v1/scheduler/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SNSDingTalkTestConnection(params param.SNSDingTalkTestConnectionParam) (*view.SNSDingTalkTestConnectionEventView, error) {
	resp := view.SNSDingTalkTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/ding-talk/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SNSEmailTestConnection(params param.SNSEmailTestConnectionParam) (*view.SNSEmailTestConnectionEventView, error) {
	resp := view.SNSEmailTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/email/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SNSFeiShuTestConnection(params param.SNSFeiShuTestConnectionParam) (*view.SNSFeiShuTestConnectionEventView, error) {
	resp := view.SNSFeiShuTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/feishu/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SNSHttpTestConnection(params param.SNSHttpTestConnectionParam) (*view.SNSHttpTestConnectionEventView, error) {
	resp := view.SNSHttpTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/http/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SNSMicrosoftTeamsTestConnection(params param.SNSMicrosoftTeamsTestConnectionParam) (*view.SNSMicrosoftTeamsTestConnectionEventView, error) {
	resp := view.SNSMicrosoftTeamsTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/microsoft-teams/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SNSSnmpTestConnection(params param.SNSSnmpTestConnectionParam) (*view.SNSSnmpTestConnectionEventView, error) {
	resp := view.SNSSnmpTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/snmp/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SNSWeComTestConnection(params param.SNSWeComTestConnectionParam) (*view.SNSWeComTestConnectionEventView, error) {
	resp := view.SNSWeComTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/we-com/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SdnControllerAddHost(params param.SdnControllerAddHostParam) (*view.SdnControllerInventoryView, error) {
	var resp view.SdnControllerAddHostEventView
	if err := cli.Post("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SdnControllerChangeHost(uuid string, params param.SdnControllerChangeHostParam) (*view.SdnControllerInventoryView, error) {
	var resp view.SdnControllerChangeHostEventView
	if err := cli.Put("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SdnControllerRemoveHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}", uuid, string(deleteMode))
}

func (cli *ZSClient) SecurityMachineDetectSync(params param.SecurityMachineDetectSyncParam) (*view.SecurityMachineDetectSyncEventView, error) {
	resp := view.SecurityMachineDetectSyncEventView{}
	if err := cli.Post("v1/security-machine/{uuid}/detect/sync/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SecurityMachineEncrypt(params param.SecurityMachineEncryptParam) (*view.SecurityMachineEncryptEventView, error) {
	resp := view.SecurityMachineEncryptEventView{}
	if err := cli.Post("v1/security-machine/encrypt/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SelfTestLocalRaid(uuid string, params param.SelfTestLocalRaidParam) (*view.SelfTestLocalRaidEventView, error) {
	resp := view.SelfTestLocalRaidEventView{}
	if err := cli.Put("v1/storage-devices/local-raid/physical-drives/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetFlowMeterRouterId(params param.SetFlowMeterRouterIdParam) (*view.SetFlowMeterRouterIdEventView, error) {
	resp := view.SetFlowMeterRouterIdEventView{}
	if err := cli.Post("v1/flowmeters/{vRouterUuid}/routerid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetIAM2ProjectContainerCluster(uuid string, params param.SetIAM2ProjectContainerClusterParam) (*view.SetIAM2ProjectContainerClusterEventView, error) {
	resp := view.SetIAM2ProjectContainerClusterEventView{}
	if err := cli.Put("v1/iam2/projects/{uuid}/container/cluster/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetIAM2ProjectLoginExpired(uuid string, params param.SetIAM2ProjectLoginExpiredParam) (*view.SetIAM2ProjectLoginExpiredEventView, error) {
	resp := view.SetIAM2ProjectLoginExpiredEventView{}
	if err := cli.Put("v1/iam2/projects/add/login-expired/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetIAM2ProjectRetirePolicy(uuid string, params param.SetIAM2ProjectRetirePolicyParam) (*view.SetIAM2ProjectRetirePolicyEventView, error) {
	resp := view.SetIAM2ProjectRetirePolicyEventView{}
	if err := cli.Put("v1/iam2/projects/retire-policies/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetImageBootMode(uuid string, params param.SetImageBootModeParam) (*view.SetImageBootModeEventView, error) {
	resp := view.SetImageBootModeEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetImageQga(uuid string, params param.SetImageQgaParam) (*view.SetImageQgaEventView, error) {
	resp := view.SetImageQgaEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetImageSecurityLevel(uuid string, params param.SetImageSecurityLevelParam) (*view.SetImageSecurityLevelEventView, error) {
	resp := view.SetImageSecurityLevelEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetImageStoreBackupStorageQuota(uuid string, params param.SetImageStoreBackupStorageQuotaParam) (*view.SetImageStoreBackupStorageQuotaEventView, error) {
	resp := view.SetImageStoreBackupStorageQuotaEventView{}
	if err := cli.Put("v1/backup-storage/image-store/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetIpOnHostNetworkBonding(params param.SetIpOnHostNetworkBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	var resp view.SetIpOnHostNetworkBondingEventView
	if err := cli.Post("v1/hosts/bondings/{bondingUuid}/ip", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetIpOnHostNetworkInterface(params param.SetIpOnHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	var resp view.SetIpOnHostNetworkInterfaceEventView
	if err := cli.Post("v1/hosts/nics/{interfaceUuid}/ip", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetL3NetworkMtu(params param.SetL3NetworkMtuParam) (*view.SetL3NetworkMtuEventView, error) {
	resp := view.SetL3NetworkMtuEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/mtu", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetL3NetworkRouterInterfaceIp(params param.SetL3NetworkRouterInterfaceIpParam) (*view.SetL3NetworkRouterInterfaceIpEventView, error) {
	resp := view.SetL3NetworkRouterInterfaceIpEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/router-interface-ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetNicQos(uuid string, params param.SetNicQosParam) (*view.SetNicQosEventView, error) {
	resp := view.SetNicQosEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetOrganizationOperation(uuid string, params param.SetOrganizationOperationParam) (*view.SetOrganizationOperationEventView, error) {
	resp := view.SetOrganizationOperationEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/operation", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetOrganizationSupervisor(uuid string, params param.SetOrganizationSupervisorParam) (*view.SetOrganizationSupervisorEventView, error) {
	resp := view.SetOrganizationSupervisorEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetSecurityMachineKey(params param.SetSecurityMachineKeyParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.Post("v1/secret-resource-pool-token/set/{uuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetServiceTypeOnHostNetworkBonding(params param.SetServiceTypeOnHostNetworkBondingParam) (*view.ListView, error) {
	var resp view.SetServiceTypeOnHostNetworkBondingEventView
	if err := cli.Post("v1/hosts/bondings/service-types", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetServiceTypeOnHostNetworkInterface(params param.SetServiceTypeOnHostNetworkInterfaceParam) (*view.ListView, error) {
	var resp view.SetServiceTypeOnHostNetworkInterfaceEventView
	if err := cli.Post("v1/hosts/nics/service-types", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVRouterRouterId(params param.SetVRouterRouterIdParam) (*view.SetVRouterRouterIdEventView, error) {
	resp := view.SetVRouterRouterIdEventView{}
	if err := cli.Post("v1/routerArea/{vRouterUuid}/routerid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVipQos(uuid string, params param.SetVipQosParam) (*view.VipQosInventoryView, error) {
	var resp view.SetVipQosEventView
	if err := cli.Put("v1/vips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVmBootMode(uuid string, params param.SetVmBootModeParam) (*view.SetVmBootModeEventView, error) {
	resp := view.SetVmBootModeEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmBootOrder(uuid string, params param.SetVmBootOrderParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmBootOrderEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVmBootVolume(uuid string, params param.SetVmBootVolumeParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmBootVolumeEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVmCleanTraffic(uuid string, params param.SetVmCleanTrafficParam) (*view.SetVmCleanTrafficEventView, error) {
	resp := view.SetVmCleanTrafficEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmClockTrack(uuid string, params param.SetVmClockTrackParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmClockTrackEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVmConsoleMode(uuid string, params param.SetVmConsoleModeParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmConsoleModeEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVmConsolePassword(uuid string, params param.SetVmConsolePasswordParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmConsolePasswordEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVmEmulatorPinning(uuid string, params param.SetVmEmulatorPinningParam) (*view.SetVmEmulatorPinningEventView, error) {
	resp := view.SetVmEmulatorPinningEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmHostname(uuid string, params param.SetVmHostnameParam) (*view.SetVmHostnameEventView, error) {
	resp := view.SetVmHostnameEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmInstanceDefaultCdRom(uuid string, params param.SetVmInstanceDefaultCdRomParam) (*view.VmCdRomInventoryView, error) {
	var resp view.SetVmInstanceDefaultCdRomEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/cdroms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVmInstanceHaLevel(params param.SetVmInstanceHaLevelParam) (*view.SetVmInstanceHaLevelEventView, error) {
	resp := view.SetVmInstanceHaLevelEventView{}
	if err := cli.Post("v1/vm-instances/{uuid}/ha-levels", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmInstanceHygonMdev(params param.SetVmInstanceHygonMdevParam) (*view.SetVmInstanceHygonMdevEventView, error) {
	resp := view.SetVmInstanceHygonMdevEventView{}
	if err := cli.Post("v1/vm-instances/{uuid}/hygon-mdev", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmMonitorNumber(uuid string, params param.SetVmMonitorNumberParam) (*view.SetVmMonitorNumberEventView, error) {
	resp := view.SetVmMonitorNumberEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmNicSecurityGroup(uuid string, params param.SetVmNicSecurityGroupParam) (*view.ListView, error) {
	var resp view.SetVmNicSecurityGroupEventView
	if err := cli.Put("v1/security-groups/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVmNuma(uuid string, params param.SetVmNumaParam) (*view.SetVmNumaEventView, error) {
	resp := view.SetVmNumaEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmQga(uuid string, params param.SetVmQgaParam) (*view.SetVmQgaEventView, error) {
	resp := view.SetVmQgaEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmQxlMemory(uuid string, params param.SetVmQxlMemoryParam) (*view.SetVmQxlMemoryEventView, error) {
	resp := view.SetVmQxlMemoryEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmRDP(uuid string, params param.SetVmRDPParam) (*view.SetVmRDPEventView, error) {
	resp := view.SetVmRDPEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmSecurityLevel(uuid string, params param.SetVmSecurityLevelParam) (*view.SetVmSecurityLevelEventView, error) {
	resp := view.SetVmSecurityLevelEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmSoundType(uuid string, params param.SetVmSoundTypeParam) (*view.SetVmSoundTypeEventView, error) {
	resp := view.SetVmSoundTypeEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmSshKey(uuid string, params param.SetVmSshKeyParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmSshKeyEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVmStaticIp(uuid string, params param.SetVmStaticIpParam) (*view.SetVmStaticIpEventView, error) {
	resp := view.SetVmStaticIpEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmUsbRedirect(uuid string, params param.SetVmUsbRedirectParam) (*view.SetVmUsbRedirectEventView, error) {
	resp := view.SetVmUsbRedirectEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmUserDefinedXmlHookScript(uuid string, params param.SetVmUserDefinedXmlHookScriptParam) (*view.SetVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.SetVmUserDefinedXmlHookScriptEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVmUserDefinedXml(uuid string, params param.SetVmUserDefinedXmlParam) (*view.SetVmUserDefinedXmlEventView, error) {
	resp := view.SetVmUserDefinedXmlEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVolumeIoThreadPin(uuid string, params param.SetVolumeIoThreadPinParam) (*view.SetVolumeIoThreadPinEventView, error) {
	resp := view.SetVolumeIoThreadPinEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVolumeQos(uuid string, params param.SetVolumeQosParam) (*view.VolumeInventoryView, error) {
	var resp view.SetVolumeQosEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SetVpcVRouterDistributedRoutingEnabled(params param.SetVpcVRouterDistributedRoutingEnabledParam) (*view.SetVpcVRouterDistributedRoutingEnabledEventView, error) {
	resp := view.SetVpcVRouterDistributedRoutingEnabledEventView{}
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/distributed-routing", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SetVpcVRouterNetworkServiceState(params param.SetVpcVRouterNetworkServiceStateParam) (*view.SetVpcVRouterNetworkServiceStateEventView, error) {
	resp := view.SetVpcVRouterNetworkServiceStateEventView{}
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/networkservicestate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ShareResource(uuid string, params param.ShareResourceParam) (*view.ShareResourceEventView, error) {
	resp := view.ShareResourceEventView{}
	if err := cli.Put("v1/accounts/resources/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ShrinkVolumeSnapshot(uuid string, params param.ShrinkVolumeSnapshotParam) (*view.ShrinkVolumeSnapshotEventView, error) {
	resp := view.ShrinkVolumeSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/shrink/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ShutdownHost(uuid string, params param.ShutdownHostParam) (*view.HostInventoryView, error) {
	var resp view.ShutdownHostEventView
	if err := cli.Put("v1/hosts/power/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SsoClientPushData(uuid string, params param.SsoClientPushDataParam) (*view.SsoClientPushDataEventView, error) {
	resp := view.SsoClientPushDataEventView{}
	if err := cli.Put("v1/sso/resource/data/push", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SubmitLongJob(params param.SubmitLongJobParam) (*view.LongJobInventoryView, error) {
	var resp view.SubmitLongJobEventView
	if err := cli.Post("v1/longjobs", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SubscribeEvent(params param.SubscribeEventParam) (*view.EventSubscriptionInventoryView, error) {
	var resp view.SubscribeEventEventView
	if err := cli.Post("v1/zwatch/events/subscriptions", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SubscribeSNSTopic(params param.SubscribeSNSTopicParam) (*view.SubscribeSNSTopicEventView, error) {
	resp := view.SubscribeSNSTopicEventView{}
	if err := cli.Post("v1/sns/topics/{topicUuid}/endpoints/{endpointUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncAINginxConfiguration(params param.SyncAINginxConfigurationParam) (*view.SyncAINginxConfigurationView, error) {
	resp := view.SyncAINginxConfigurationView{}
	if err := cli.Post("v1/ai/nginx/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncAliyunRouteEntryFromRemote(uuid string, params param.SyncAliyunRouteEntryFromRemoteParam) (*view.VpcVirtualRouteEntryInventoryView, error) {
	resp := view.VpcVirtualRouteEntryInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/route-entry/{vRouterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncAliyunRouterInterfaceFromRemote(uuid string, params param.SyncAliyunRouterInterfaceFromRemoteParam) (*view.AliyunRouterInterfaceInventoryView, error) {
	resp := view.AliyunRouterInterfaceInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/router-interface/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncAliyunSnapshotRemote(params param.SyncAliyunSnapshotRemoteParam) (*view.AliyunSnapshotInventoryView, error) {
	resp := view.AliyunSnapshotInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/snapshot/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncAliyunVirtualRouterFromRemote(uuid string, params param.SyncAliyunVirtualRouterFromRemoteParam) (*view.VpcVirtualRouterInventoryView, error) {
	resp := view.VpcVirtualRouterInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/vrouter/{vpcUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncBackupFromImageStoreBackupStorage(uuid string, params param.SyncBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	var resp view.SyncBackupFromImageStoreBackupStorageEventView
	if err := cli.Put("v1/volume-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SyncChronyServers(uuid string, params param.SyncChronyServersParam) (*view.SyncChronyServersEventView, error) {
	resp := view.SyncChronyServersEventView{}
	if err := cli.Put("v1/zops/chrony/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncConnectionAccessPointFromRemote(uuid string, params param.SyncConnectionAccessPointFromRemoteParam) (*view.ConnectionAccessPointInventoryView, error) {
	resp := view.ConnectionAccessPointInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/access-point/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncContainerManagementEndpoint(uuid string, params param.SyncContainerManagementEndpointParam) (*view.ContainerManagementEndpointInventoryView, error) {
	var resp view.SyncContainerManagementEndpointEventView
	if err := cli.Put("v1/container/management/endpoint/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SyncDataCenterFromRemote(params param.SyncDataCenterFromRemoteParam) (*view.SyncDataCenterFromRemoteEventView, error) {
	var resp view.SyncDataCenterFromRemoteEventView
	if err := cli.Get("v1/hybrid/data-center/{uuid}/sync", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncDatabaseBackupFromImageStoreBackupStorage(uuid string, params param.SyncDatabaseBackupFromImageStoreBackupStorageParam) (*view.DatabaseBackupInventoryView, error) {
	var resp view.SyncDatabaseBackupFromImageStoreBackupStorageEventView
	if err := cli.Put("v1/database-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SyncDatabaseBackup(uuid string, params param.SyncDatabaseBackupParam) (*view.SyncDatabaseBackupEventView, error) {
	resp := view.SyncDatabaseBackupEventView{}
	if err := cli.Put("v1/database-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncDiskFromAliyunFromRemote(params param.SyncDiskFromAliyunFromRemoteParam) (*view.AliyunDiskInventoryView, error) {
	resp := view.AliyunDiskInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/disk/{identityUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncEcsImageFromRemote(params param.SyncEcsImageFromRemoteParam) (*view.EcsImageInventoryView, error) {
	resp := view.EcsImageInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/image/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncEcsInstanceFromRemote(params param.SyncEcsInstanceFromRemoteParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/ecs/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncEcsSecurityGroupFromRemote(uuid string, params param.SyncEcsSecurityGroupFromRemoteParam) (*view.EcsSecurityGroupInventoryView, error) {
	resp := view.EcsSecurityGroupInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/security-group/{ecsVpcUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncEcsSecurityGroupRuleFromRemote(uuid string, params param.SyncEcsSecurityGroupRuleFromRemoteParam) (*view.EcsSecurityGroupRuleInventoryView, error) {
	resp := view.EcsSecurityGroupRuleInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/security-group-rule/{uuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncEcsVSwitchFromRemote(params param.SyncEcsVSwitchFromRemoteParam) (*view.EcsVSwitchInventoryView, error) {
	resp := view.EcsVSwitchInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/vswitch/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncEcsVpcFromRemote(params param.SyncEcsVpcFromRemoteParam) (*view.EcsVpcInventoryView, error) {
	resp := view.EcsVpcInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/vpc/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncHybridEipFromRemote(uuid string, params param.SyncHybridEipFromRemoteParam) (*view.HybridEipAddressInventoryView, error) {
	resp := view.HybridEipAddressInventoryView{}
	if err := cli.Put("v1/hybrid/eip/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncIdentityFromRemote(params param.SyncIdentityFromRemoteParam) (*view.SyncIdentityFromRemoteEventView, error) {
	var resp view.SyncIdentityFromRemoteEventView
	if err := cli.Get("v1/hybrid/identity-zone/{uuid}/sync", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncImageFromImageStoreBackupStorage(uuid string, params param.SyncImageFromImageStoreBackupStorageParam) (*view.ImageInventoryView, error) {
	var resp view.SyncImageFromImageStoreBackupStorageEventView
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SyncImage(uuid string, params param.SyncImageParam) (*view.SyncImageEventView, error) {
	resp := view.SyncImageEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncImageSize(uuid string, params param.SyncImageSizeParam) (*view.ImageInventoryView, error) {
	var resp view.SyncImageSizeEventView
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SyncLdapServer(uuid string, params param.SyncLdapServerParam) (*view.LongJobInventoryView, error) {
	var resp view.SyncLdapServerEventView
	if err := cli.Put("v1/ldap/servers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SyncLicenseCapacity(uuid string, params param.SyncLicenseCapacityParam) (*view.SyncLicenseCapacityEventView, error) {
	resp := view.SyncLicenseCapacityEventView{}
	if err := cli.Put("v1/license-server/authorized-capacity/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncNfvInstGroup(uuid string, params param.SyncNfvInstGroupParam) (*view.NfvInstGroupInventoryView, error) {
	var resp view.SyncNfvInstGroupEventView
	if err := cli.Put("v1/nfvinstgroup/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SyncPrimaryStorageCapacity(uuid string, params param.SyncPrimaryStorageCapacityParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.SyncPrimaryStorageCapacityEventView
	if err := cli.Put("v1/primary-storage/{primaryStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SyncVCenter(uuid string, params param.SyncVCenterParam) (*view.SyncVCenterEventView, error) {
	resp := view.SyncVCenterEventView{}
	if err := cli.Put("v1/vcenters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncVirtualBorderRouterFromRemote(uuid string, params param.SyncVirtualBorderRouterFromRemoteParam) (*view.VirtualBorderRouterInventoryView, error) {
	resp := view.VirtualBorderRouterInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/border-router/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncVmBackupFromImageStoreBackupStorage(uuid string, params param.SyncVmBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.Put("v1/vm-backups/{groupUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncVmBackup(uuid string, params param.SyncVmBackupParam) (*view.SyncVmBackupEventView, error) {
	resp := view.SyncVmBackupEventView{}
	if err := cli.Put("v1/vm-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncVmClock(uuid string, params param.SyncVmClockParam) (*view.SyncVmClockEventView, error) {
	resp := view.SyncVmClockEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncVolumeBackup(uuid string, params param.SyncVolumeBackupParam) (*view.SyncVolumeBackupEventView, error) {
	resp := view.SyncVolumeBackupEventView{}
	if err := cli.Put("v1/volume-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncVolumeSize(uuid string, params param.SyncVolumeSizeParam) (*view.VolumeInventoryView, error) {
	var resp view.SyncVolumeSizeEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) SyncVpcUserVpnGatewayFromRemote(uuid string, params param.SyncVpcUserVpnGatewayFromRemoteParam) (*view.VpcUserVpnGatewayInventoryView, error) {
	resp := view.VpcUserVpnGatewayInventoryView{}
	if err := cli.Put("v1/hybrid/user-vpn/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncVpcVpnConnectionFromRemote(uuid string, params param.SyncVpcVpnConnectionFromRemoteParam) (*view.VpcVpnConnectionInventoryView, error) {
	resp := view.VpcVpnConnectionInventoryView{}
	if err := cli.Put("v1/hybrid/vpn-connection/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncVpcVpnGatewayFromRemote(uuid string, params param.SyncVpcVpnGatewayFromRemoteParam) (*view.VpcVpnGatewayInventoryView, error) {
	resp := view.VpcVpnGatewayInventoryView{}
	if err := cli.Put("v1/hybrid/vpc-vpn/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) SyncZBoxCapacity(uuid string, params param.SyncZBoxCapacityParam) (*view.ZBoxInventoryView, error) {
	var resp view.SyncZBoxCapacityEventView
	if err := cli.Put("v1/zbox/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) TakeVmConsoleScreenshot(uuid string, params param.TakeVmConsoleScreenshotParam) (*view.TakeVmConsoleScreenshotEventView, error) {
	resp := view.TakeVmConsoleScreenshotEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) TerminateVirtualBorderRouterRemote(uuid string, params param.TerminateVirtualBorderRouterRemoteParam) (*view.TerminateVirtualBorderRouterRemoteEventView, error) {
	resp := view.TerminateVirtualBorderRouterRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/border-router/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) TokenIntrospection(params param.TokenIntrospectionParam) (*view.TokenIntrospectionView, error) {
	resp := view.TokenIntrospectionView{}
	if err := cli.Post("v1/token/introspect", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) TriggerGCJob(uuid string, params param.TriggerGCJobParam) (*view.TriggerGCJobEventView, error) {
	resp := view.TriggerGCJobEventView{}
	if err := cli.Put("v1/gc-jobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UnbindModelFromService(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models/{modelUuid}/model-services/{modelServiceUuid}", uuid, string(deleteMode))
}

func (cli *ZSClient) UndoSnapshotCreation(uuid string, params param.UndoSnapshotCreationParam) (*view.VolumeInventoryView, error) {
	var resp view.UndoSnapshotCreationEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) UnexportNbdVolumes(params param.UnexportNbdVolumesParam) (*view.UnexportNbdVolumesEventView, error) {
	resp := view.UnexportNbdVolumesEventView{}
	if err := cli.Post("v1/cbt-task/unexportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UngenerateHygonMdevDevices(uuid string, params param.UngenerateHygonMdevDevicesParam) (*view.UngenerateHygonMdevDevicesEventView, error) {
	resp := view.UngenerateHygonMdevDevicesEventView{}
	if err := cli.Put("v1/hygon-devices/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UngenerateMdevDevices(uuid string, params param.UngenerateMdevDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UngenerateSeMdevDevices(uuid string, params param.UngenerateSeMdevDevicesParam) (*view.UngenerateSeMdevDevicesEventView, error) {
	resp := view.UngenerateSeMdevDevicesEventView{}
	if err := cli.Put("v1/mtty-devices/{mttyDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UngenerateSriovPciDevices(uuid string, params param.UngenerateSriovPciDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UngroupVolumeSnapshotGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots/ungroup/{uuid}", uuid, string(deleteMode))
}

func (cli *ZSClient) UnlockIdentity(params param.UnlockIdentityParam) (*view.UnlockIdentityView, error) {
	var resp view.UnlockIdentityView
	if err := cli.Get("v1/login/control/unlock", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UnmountVmInstanceRecoveryPoint(params param.UnmountVmInstanceRecoveryPointParam) (*view.UnmountVmInstanceRecoveryPointEventView, error) {
	resp := view.UnmountVmInstanceRecoveryPointEventView{}
	if err := cli.Post("v1/cdp-backup-storage/unmount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UnprotectVmInstanceRecoveryPoint(uuid string, params param.UnprotectVmInstanceRecoveryPointParam) (*view.UnprotectVmInstanceRecoveryPointEventView, error) {
	resp := view.UnprotectVmInstanceRecoveryPointEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/unprotect-recovery-point", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UnregisterLicenseRequestedApplication(uuid string, params param.UnregisterLicenseRequestedApplicationParam) (*view.UnregisterLicenseRequestedApplicationEventView, error) {
	resp := view.UnregisterLicenseRequestedApplicationEventView{}
	if err := cli.Put("v1/license/unregister-applications", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UnregisterLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server/client", uuid, string(deleteMode))
}

func (cli *ZSClient) UnsubscribeEvent(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions/{uuid}", uuid, string(deleteMode))
}

func (cli *ZSClient) UnsubscribeSNSTopic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/topics/{topicUuid}/endpoints/{endpointUuid}", uuid, string(deleteMode))
}

func (cli *ZSClient) UpgradeBackupStorageCdpTasks(uuid string, params param.UpgradeBackupStorageCdpTasksParam) (*view.UpgradeBackupStorageCdpTasksEventView, error) {
	resp := view.UpgradeBackupStorageCdpTasksEventView{}
	if err := cli.Put("v1/cdp-task/upgrade/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) UpgradeToLicenseServer(params param.UpgradeToLicenseServerParam) (*view.LicenseAuthorizedNodeInventoryView, error) {
	var resp view.UpgradeToLicenseServerEventView
	if err := cli.Post("v1/license-server", params, &resp); err != nil {
		return nil, err
	}
	return resp.Inventory, nil
}

func (cli *ZSClient) UploadFileToVm(params param.UploadFileToVmParam) (*view.UploadFileToVmEventView, error) {
	resp := view.UploadFileToVmEventView{}
	if err := cli.Post("v1/upload-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidateClusterSupportDRS(params param.ValidateClusterSupportDRSParam) (*view.ValidateClusterSupportDRSView, error) {
	var resp view.ValidateClusterSupportDRSView
	if err := cli.Get("v1/clusters/{clusterUuid}/drs/valid", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidateDiskOfferingUserConfig(uuid string, params param.ValidateDiskOfferingUserConfigParam) (*view.ValidateDiskOfferingUserConfigEventView, error) {
	resp := view.ValidateDiskOfferingUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidateInstanceOfferingUserConfig(uuid string, params param.ValidateInstanceOfferingUserConfigParam) (*view.ValidateInstanceOfferingUserConfigEventView, error) {
	resp := view.ValidateInstanceOfferingUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidatePassword(uuid string, params param.ValidatePasswordParam) (*view.ValidatePasswordView, error) {
	resp := view.ValidatePasswordView{}
	if err := cli.Put("v1/password/verify", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidatePriceUserConfig(uuid string, params param.ValidatePriceUserConfigParam) (*view.ValidatePriceUserConfigEventView, error) {
	resp := view.ValidatePriceUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidateSNSAliyunSmsEndpoint(uuid string, params param.ValidateSNSAliyunSmsEndpointParam) (*view.ValidateSNSAliyunSmsEndpointEventView, error) {
	resp := view.ValidateSNSAliyunSmsEndpointEventView{}
	if err := cli.Put("v1/sns/sms-endpoints/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidateSNSEmailPlatform(uuid string, params param.ValidateSNSEmailPlatformParam) (*view.ValidateSNSEmailPlatformEventView, error) {
	resp := view.ValidateSNSEmailPlatformEventView{}
	if err := cli.Put("v1/sns/application-platforms/email/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidateSNSUniversalSmsEndpoint(uuid string, params param.ValidateSNSUniversalSmsEndpointParam) (*view.ValidateSNSApplicationEndpointEventView, error) {
	resp := view.ValidateSNSApplicationEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/universal-sms/{uuid}/validate", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidateSecurityGroupRule(params param.ValidateSecurityGroupRuleParam) (*view.ValidateSecurityGroupRuleView, error) {
	var resp view.ValidateSecurityGroupRuleView
	if err := cli.Get("v1/security-groups/{securityGroupUuid}/rules/validation", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidateSession(params param.ValidateSessionParam) (*view.ValidateSessionView, error) {
	var resp view.ValidateSessionView
	if err := cli.Get("v1/accounts/sessions/{sessionUuid}/valid", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidateVmSchedulingRule(uuid string, params param.ValidateVmSchedulingRuleParam) (*view.ValidateVmSchedulingRuleView, error) {
	resp := view.ValidateVmSchedulingRuleView{}
	if err := cli.Put("v1/validate/vmSchedulingRule", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) ValidateVolumeSnapshotChain(uuid string, params param.ValidateVolumeSnapshotChainParam) (*view.ValidateVolumeSnapshotChainEventView, error) {
	resp := view.ValidateVolumeSnapshotChainEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) VerifyLicenseServer(params param.VerifyLicenseServerParam) (*view.VerifyLicenseServerEventView, error) {
	resp := view.VerifyLicenseServerEventView{}
	if err := cli.Post("v1/license-server/register-verify", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) WithdrawLicenseCapacityApplication(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server/capacity-application", uuid, string(deleteMode))
}

func (cli *ZSClient) ZQLQuery(params param.ZQLQueryParam) (*view.ZQLQueryView, error) {
	var resp view.ZQLQueryView
	if err := cli.Get("v1/zql", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

