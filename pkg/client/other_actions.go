// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ChangeIAM2OrganizationState changes IAM2OrganizationState
func (cli *ZSClient) ChangeIAM2OrganizationState(uuid string, params param.ChangeIAM2OrganizationStateParam) (*view.IAM2OrganizationInventoryView, error) {
	var resp view.ChangeIAM2OrganizationStateEventView
	if err := cli.Put("v1/iam2/organizations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateAutoScalingGroupAddingNewInstanceRule creates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupAddingNewInstanceRule(params param.CreateAutoScalingGroupAddingNewInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	var resp view.CreateAutoScalingRuleEventView
	if err := cli.Post("v1/autoscaling/rules/adding-new-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetServiceTypeOnHostNetworkBonding operates on ServiceTypeOnHostNetworkBonding
func (cli *ZSClient) SetServiceTypeOnHostNetworkBonding(params param.SetServiceTypeOnHostNetworkBondingParam) (*view.HostNetworkBondingServiceRefInventoryView, error) {
	resp := view.HostNetworkBondingServiceRefInventoryView{}
	if err := cli.Post("v1/hosts/bondings/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryPciDevicePciDeviceOffering queries PciDevicePciDeviceOffering list
func (cli *ZSClient) QueryPciDevicePciDeviceOffering(params *param.QueryParam) ([]view.PciDevicePciDeviceOfferingRefInventoryView, error) {
	var resp []view.PciDevicePciDeviceOfferingRefInventoryView
	return resp, cli.List("v1/pci-devices/pci-devices/pci-device-offerings", params, &resp)
}

// AddAttributesToIAM2Organization adds AttributesToIAM2Organization
func (cli *ZSClient) AddAttributesToIAM2Organization(params param.AddAttributesToIAM2OrganizationParam) (*view.AddAttributesToIAM2OrganizationEventView, error) {
	resp := view.AddAttributesToIAM2OrganizationEventView{}
	if err := cli.Post("v1/iam2/organizations/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCreateEcsImageProgress gets CreateEcsImageProgress by uuid
func (cli *ZSClient) GetCreateEcsImageProgress(uuid string) (*view.GetCreateEcsImageProgressView, error) {
	var resp view.GetCreateEcsImageProgressView
	if err := cli.Get("v1/hybrid/aliyun/image/{dataCenterUuid}/{imageUuid}/progress", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAccessControlListToLoadBalancer adds AccessControlListToLoadBalancer
func (cli *ZSClient) AddAccessControlListToLoadBalancer(params param.AddAccessControlListToLoadBalancerParam) (*view.LoadBalancerListenerInventoryView, error) {
	var resp view.AddAccessControlListToLoadBalancerEventView
	if err := cli.Post("v1/load-balancers/listeners/{listenerUuid}/access-control-lists", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// LogOut operates on LogOut
func (cli *ZSClient) LogOut(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/sessions/{sessionUuid}", uuid, string(deleteMode))
}

// GetVmXmlHookScript gets VmXmlHookScript by uuid
func (cli *ZSClient) GetVmXmlHookScript(uuid string) (*view.GetVmXmlHookScriptView, error) {
	var resp view.GetVmXmlHookScriptView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/xml-hook-script", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachHybridKey operates on HybridKey
func (cli *ZSClient) AttachHybridKey(uuid string, params param.AttachHybridKeyParam) (*view.AttachHybridKeyEventView, error) {
	resp := view.AttachHybridKeyEventView{}
	if err := cli.Put("v1/hybrid/hybrid/key/{uuid}/attach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImageQga gets ImageQga by uuid
func (cli *ZSClient) GetImageQga(uuid string) (*view.GetImageQgaView, error) {
	var resp view.GetImageQgaView
	if err := cli.Get("v1/images/{uuid}/qga", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInterdependentL3NetworksBackupStorages gets InterdependentL3NetworksBackupStorages by uuid
func (cli *ZSClient) GetInterdependentL3NetworksBackupStorages(uuid string) (*view.GetInterdependentL3NetworksBackupStoragesView, error) {
	var resp view.GetInterdependentL3NetworksBackupStoragesView
	if err := cli.Get("v1/backupStorage-l3networks/dependencies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBackupFileInPublic deletes BackupFileInPublic
func (cli *ZSClient) DeleteBackupFileInPublic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/backup-mysql", uuid, string(deleteMode))
}

// BatchCreateIAM2VirtualIDFromConfigFile operates on CreateIAM2VirtualIDFromConfigFile
func (cli *ZSClient) BatchCreateIAM2VirtualIDFromConfigFile(params param.BatchCreateIAM2VirtualIDFromConfigFileParam) (*view.BatchCreateIAM2VirtualIDFromConfigFileEventView, error) {
	resp := view.BatchCreateIAM2VirtualIDFromConfigFileEventView{}
	if err := cli.Post("v1/iam2/virtual-ids/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmClockTrack operates on VmClockTrack
func (cli *ZSClient) SetVmClockTrack(uuid string, params param.SetVmClockTrackParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmClockTrackEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateEmailMonitorTriggerAction updates EmailMonitorTrigger
func (cli *ZSClient) UpdateEmailMonitorTriggerAction(uuid string, params param.UpdateEmailMonitorTriggerActionParam) (*view.MonitorTriggerActionInventoryView, error) {
	var resp view.UpdateMonitorTriggerActionEventView
	if err := cli.Put("v1/monitoring/trigger-actions/emails/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryLocalRaidController queries LocalRaidController list
func (cli *ZSClient) QueryLocalRaidController(params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, error) {
	var resp []view.RaidPhysicalDriveInventoryView
	return resp, cli.List("v1/storage-devices/local-raid/controllers", params, &resp)
}

// SyncDataCenterFromRemote operates on DataCenterFromRemote
func (cli *ZSClient) SyncDataCenterFromRemote(params param.SyncDataCenterFromRemoteParam) (*view.SyncDataCenterFromRemoteEventView, error) {
	var resp view.SyncDataCenterFromRemoteEventView
	if err := cli.Get("v1/hybrid/data-center/{uuid}/sync", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBackupStorageState changes BackupStorageState
func (cli *ZSClient) ChangeBackupStorageState(uuid string, params param.ChangeBackupStorageStateParam) (*view.BackupStorageInventoryView, error) {
	var resp view.ChangeBackupStorageStateEventView
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetVmInstanceHygonMdev operates on VmInstanceHygonMdev
func (cli *ZSClient) SetVmInstanceHygonMdev(params param.SetVmInstanceHygonMdevParam) (*view.SetVmInstanceHygonMdevEventView, error) {
	resp := view.SetVmInstanceHygonMdevEventView{}
	if err := cli.Post("v1/vm-instances/{uuid}/hygon-mdev", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateIsoForAttachingVm gets CandidateIsoForAttachingVm by uuid
func (cli *ZSClient) GetCandidateIsoForAttachingVm(uuid string) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/iso-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SecurityMachineDetectSync operates on MachineDetectSync
func (cli *ZSClient) SecurityMachineDetectSync(params param.SecurityMachineDetectSyncParam) (*view.SecurityMachineDetectSyncEventView, error) {
	resp := view.SecurityMachineDetectSyncEventView{}
	if err := cli.Post("v1/security-machine/{uuid}/detect/sync/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSecurityGroupState changes SecurityGroupState
func (cli *ZSClient) ChangeSecurityGroupState(uuid string, params param.ChangeSecurityGroupStateParam) (*view.SecurityGroupInventoryView, error) {
	var resp view.ChangeSecurityGroupStateEventView
	if err := cli.Put("v1/security-groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeBareMetal2ChassisOfferingState changes BareMetal2ChassisOfferingState
func (cli *ZSClient) ChangeBareMetal2ChassisOfferingState(uuid string, params param.ChangeBareMetal2ChassisOfferingStateParam) (*view.BareMetal2ChassisOfferingInventoryView, error) {
	var resp view.ChangeBareMetal2ChassisOfferingStateEventView
	if err := cli.Put("v1/baremetal2/chassis/offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetPrometheusMetricLabelValue gets PrometheusMetricLabelValue by uuid
func (cli *ZSClient) GetPrometheusMetricLabelValue(uuid string) (*view.GetPrometheusMetricLabelValueView, error) {
	var resp view.GetPrometheusMetricLabelValueView
	if err := cli.Get("v1/zwatch/metrics/prometheus/label-values", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryConnectionAccessPointFromLocal queries ConnectionAccessPointFromLocal list
func (cli *ZSClient) QueryConnectionAccessPointFromLocal(params *param.QueryParam) ([]view.ConnectionAccessPointInventoryView, error) {
	var resp []view.ConnectionAccessPointInventoryView
	return resp, cli.List("v1/hybrid/aliyun/access-point", params, &resp)
}

// UpdateAlarmData updates AlarmData
func (cli *ZSClient) UpdateAlarmData(uuid string, params param.UpdateAlarmDataParam) (*view.UpdateAlarmDataEventView, error) {
	resp := view.UpdateAlarmDataEventView{}
	if err := cli.Put("v1/zwatch/alarm-histories/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LoginIAM2VirtualIDWithLdap operates on IAM2VirtualIDWithLdap
func (cli *ZSClient) LoginIAM2VirtualIDWithLdap(uuid string, params param.LoginIAM2VirtualIDWithLdapParam) (*view.LoginIAM2VirtualIDWithLdapView, error) {
	resp := view.LoginIAM2VirtualIDWithLdapView{}
	if err := cli.Put("v1/iam2/login/virtual-ids/ldap", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpnIpsecConfig creates VpnIpsecConfig
func (cli *ZSClient) CreateVpnIpsecConfig(params param.CreateVpnIpsecConfigParam) (*view.VpcVpnIpSecConfigInventoryView, error) {
	var resp view.CreateVpnIpsecConfigEventView
	if err := cli.Post("v1/hybrid/vpn-connection/ipsec", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SNSEmailTestConnection operates on EmailTestConnection
func (cli *ZSClient) SNSEmailTestConnection(params param.SNSEmailTestConnectionParam) (*view.SNSEmailTestConnectionEventView, error) {
	resp := view.SNSEmailTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/email/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RegisterLicenseServer operates on LicenseServer
func (cli *ZSClient) RegisterLicenseServer(params param.RegisterLicenseServerParam) (*view.RegisterLicenseServerEventView, error) {
	resp := view.RegisterLicenseServerEventView{}
	if err := cli.Post("v1/license-server/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAutoScalingGroupState changes AutoScalingGroupState
func (cli *ZSClient) ChangeAutoScalingGroupState(uuid string, params param.ChangeAutoScalingGroupStateParam) (*view.AutoScalingGroupInventoryView, error) {
	var resp view.ChangeAutoScalingGroupStateEventView
	if err := cli.Put("v1/autoscaling/groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateAutoScalingGroupRemovalInstanceRule creates AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupRemovalInstanceRule(params param.CreateAutoScalingGroupRemovalInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	var resp view.CreateAutoScalingRuleEventView
	if err := cli.Post("v1/autoscaling/rules/removal-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeSlbGroupMonitorIps changes SlbGroupMonitorIps
func (cli *ZSClient) ChangeSlbGroupMonitorIps(uuid string, params param.ChangeSlbGroupMonitorIpsParam) (*view.SlbGroupInventoryView, error) {
	var resp view.ChangeSlbGroupMonitorIpsEventView
	if err := cli.Put("v1/load-balancers/slb/groups/{slbGroupUuid}/monitorIps", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteModelEvaluationTasks deletes ModelEvaluationTasks
func (cli *ZSClient) DeleteModelEvaluationTasks(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/model-evaluation-tasks", uuid, string(deleteMode))
}

// AttachL3NetworkToVm operates on L3NetworkToVm
func (cli *ZSClient) AttachL3NetworkToVm(params param.AttachL3NetworkToVmParam) (*view.VmInstanceInventoryView, error) {
	var resp view.AttachL3NetworkToVmEventView
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachPrimaryStorageToCluster operates on PrimaryStorageToCluster
func (cli *ZSClient) AttachPrimaryStorageToCluster(params param.AttachPrimaryStorageToClusterParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AttachPrimaryStorageToClusterEventView
	if err := cli.Post("v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachL2NetworkToCluster operates on L2NetworkToCluster
func (cli *ZSClient) AttachL2NetworkToCluster(params param.AttachL2NetworkToClusterParam) (*view.L2NetworkInventoryView, error) {
	var resp view.AttachL2NetworkToClusterEventView
	if err := cli.Post("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeVmNicType changes VmNicType
func (cli *ZSClient) ChangeVmNicType(uuid string, params param.ChangeVmNicTypeParam) (*view.VmNicInventoryView, error) {
	var resp view.ChangeVmNicTypeEventView
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeFirewallRuleState changes FirewallRuleState
func (cli *ZSClient) ChangeFirewallRuleState(uuid string, params param.ChangeFirewallRuleStateParam) (*view.VpcFirewallRuleInventoryView, error) {
	var resp view.ChangeFirewallRuleStateEventView
	if err := cli.Put("v1/vpcfirewalls/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetMdevDeviceCandidates gets MdevDeviceCandidates by uuid
func (cli *ZSClient) GetMdevDeviceCandidates(uuid string) (*view.MdevDeviceInventoryView, error) {
	var resp view.MdevDeviceInventoryView
	if err := cli.Get("v1/mdev-devices/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTwoFactorAuthenticationState gets TwoFactorAuthenticationState by uuid
func (cli *ZSClient) GetTwoFactorAuthenticationState(uuid string) (*view.GetTwoFactorAuthenticationStateView, error) {
	var resp view.GetTwoFactorAuthenticationStateView
	if err := cli.Get("v1/twofactorauthentication/state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BootstrapMiniHost operates on MiniHost
func (cli *ZSClient) BootstrapMiniHost(params param.BootstrapMiniHostParam) (*view.BootstrapMiniHostEventView, error) {
	resp := view.BootstrapMiniHostEventView{}
	if err := cli.Post("v1/mini-clusters/hosts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveActionFromAlarm removes ActionFromAlarm
func (cli *ZSClient) RemoveActionFromAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/{alarmUuid}/actions/{actionUuid}", uuid, string(deleteMode))
}

// ChangeEipState changes EipState
func (cli *ZSClient) ChangeEipState(uuid string, params param.ChangeEipStateParam) (*view.EipInventoryView, error) {
	var resp view.ChangeEipStateEventView
	if err := cli.Put("v1/eips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachSshKeyPairFromVmInstance operates on SshKeyPairFromVmInstance
func (cli *ZSClient) DetachSshKeyPairFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ssh-key-pair/{sshKeyPairUuid}/vm-instance/{vmInstanceUuid}", uuid, string(deleteMode))
}

// GetPrimaryStorageCandidatesForVmMigration gets PrimaryStorageCandidatesForVmMigration by uuid
func (cli *ZSClient) GetPrimaryStorageCandidatesForVmMigration(uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/storage-migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVolume operates on PrimaryStorageMigrateVolume
func (cli *ZSClient) PrimaryStorageMigrateVolume(uuid string, params param.PrimaryStorageMigrateVolumeParam) (*view.VolumeInventoryView, error) {
	var resp view.PrimaryStorageMigrateVolumeEventView
	if err := cli.Put("v1/primary-storage/volumes/{volumeUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteHybridEipRemote deletes HybridEipRemote
func (cli *ZSClient) DeleteHybridEipRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/eip/{uuid}/remote", uuid, string(deleteMode))
}

// DeleteModelServiceInstanceGroups deletes ModelServiceInstanceGroups
func (cli *ZSClient) DeleteModelServiceInstanceGroups(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-services/instances/groups", uuid, string(deleteMode))
}

// GetVmBootOrder gets VmBootOrder by uuid
func (cli *ZSClient) GetVmBootOrder(uuid string) (*view.GetVmBootOrderView, error) {
	var resp view.GetVmBootOrderView
	if err := cli.Get("v1/vm-instances/{uuid}/boot-orders", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmBootOrder operates on VmBootOrder
func (cli *ZSClient) SetVmBootOrder(uuid string, params param.SetVmBootOrderParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmBootOrderEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryThirdpartyAlert queries ThirdpartyAlert list
func (cli *ZSClient) QueryThirdpartyAlert(params *param.QueryParam) ([]view.ThirdpartyOriginalAlertInventoryView, error) {
	var resp []view.ThirdpartyOriginalAlertInventoryView
	return resp, cli.List("v1/zwatch/third-party/alerts", params, &resp)
}

// GetDatabaseBackupFromImageStore gets DatabaseBackupFromImageStore by uuid
func (cli *ZSClient) GetDatabaseBackupFromImageStore(uuid string) (*view.GetDatabaseBackupFromImageStoreView, error) {
	var resp view.GetDatabaseBackupFromImageStoreView
	if err := cli.Get("v1/database-backups/image-store", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsVSwitchFromRemote operates on EcsVSwitchFromRemote
func (cli *ZSClient) SyncEcsVSwitchFromRemote(params param.SyncEcsVSwitchFromRemoteParam) (*view.EcsVSwitchInventoryView, error) {
	resp := view.EcsVSwitchInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/vswitch/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LocateLocalRaidPhysicalDrive operates on LocalRaidPhysicalDrive
func (cli *ZSClient) LocateLocalRaidPhysicalDrive(uuid string, params param.LocateLocalRaidPhysicalDriveParam) (*view.RaidPhysicalDriveInventoryView, error) {
	var resp view.LocateLocalRaidPhysicalDriveEventView
	if err := cli.Put("v1/storage-devices/local-raid/physical-drives/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CleanUpBaremetalChassisBonding operates on UpBaremetalChassisBonding
func (cli *ZSClient) CleanUpBaremetalChassisBonding(uuid string, params param.CleanUpBaremetalChassisBondingParam) (*view.CleanUpBaremetalChassisBondingEventView, error) {
	resp := view.CleanUpBaremetalChassisBondingEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemovePciDeviceSpecFromVmInstance removes PciDeviceSpecFromVmInstance
func (cli *ZSClient) RemovePciDeviceSpecFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device-specs/{pciSpecUuid}/vm-instances/{vmInstanceUuid}", uuid, string(deleteMode))
}

// AddIAM2VirtualIDGroupToProjects adds IAM2VirtualIDGroupToProjects
func (cli *ZSClient) AddIAM2VirtualIDGroupToProjects(params param.AddIAM2VirtualIDGroupToProjectsParam) (*view.AddIAM2VirtualIDGroupToProjectsEventView, error) {
	resp := view.AddIAM2VirtualIDGroupToProjectsEventView{}
	if err := cli.Post("v1/iam2/projects/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveServerGroupFromLoadBalancerListener removes ServerGroupFromLoadBalancerListener
func (cli *ZSClient) RemoveServerGroupFromLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/servergroups", uuid, string(deleteMode))
}

// AddSharedBlockToSharedBlockGroup adds SharedBlockToSharedBlockGroup
func (cli *ZSClient) AddSharedBlockToSharedBlockGroup(params param.AddSharedBlockToSharedBlockGroupParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	var resp view.AddSharedBlockToSharedBlockGroupEventView
	if err := cli.Post("v1/primary-storage/sharedblockgroup/{uuid}/sharedblocks", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RefreshCaptcha operates on Captcha
func (cli *ZSClient) RefreshCaptcha(params param.RefreshCaptchaParam) (*view.RefreshCaptchaView, error) {
	var resp view.RefreshCaptchaView
	if err := cli.Get("v1/captcha/refresh", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsVSwitchInLocal deletes EcsVSwitchInLocal
func (cli *ZSClient) DeleteEcsVSwitchInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vswitch/{uuid}", uuid, string(deleteMode))
}

// DeleteTag deletes Tag
func (cli *ZSClient) DeleteTag(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tags/{uuid}", uuid, string(deleteMode))
}

// AddIAM2VirtualIDsToOrganization adds IAM2VirtualIDsToOrganization
func (cli *ZSClient) AddIAM2VirtualIDsToOrganization(params param.AddIAM2VirtualIDsToOrganizationParam) (*view.AddIAM2VirtualIDsToOrganizationEventView, error) {
	resp := view.AddIAM2VirtualIDsToOrganizationEventView{}
	if err := cli.Post("v1/iam2/organizations/{organizationUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachProvisionNicToBonding operates on ProvisionNicToBonding
func (cli *ZSClient) AttachProvisionNicToBonding(params param.AttachProvisionNicToBondingParam) (*view.BareMetal2InstanceInventoryView, error) {
	var resp view.AttachProvisionNicToBondingEventView
	if err := cli.Post("v1/baremetal2/bm-instances/{uuid}/bm2-bondings/{bondingUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ExportNbdVolumes operates on NbdVolumes
func (cli *ZSClient) ExportNbdVolumes(params param.ExportNbdVolumesParam) (*view.ExportNbdVolumesEventView, error) {
	resp := view.ExportNbdVolumesEventView{}
	if err := cli.Post("v1/cbt-task/exportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelfTestLocalRaid operates on LocalRaid
func (cli *ZSClient) SelfTestLocalRaid(uuid string, params param.SelfTestLocalRaidParam) (*view.SelfTestLocalRaidEventView, error) {
	resp := view.SelfTestLocalRaidEventView{}
	if err := cli.Put("v1/storage-devices/local-raid/physical-drives/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSNSApplicationPlatformState changes SNSApplicationPlatformState
func (cli *ZSClient) ChangeSNSApplicationPlatformState(uuid string, params param.ChangeSNSApplicationPlatformStateParam) (*view.SNSApplicationPlatformInventoryView, error) {
	var resp view.ChangeSNSApplicationPlatformStateEventView
	if err := cli.Put("v1/sns/application-platforms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PowerOffBareMetal2Chassis operates on PowerOffBareMetal2Chassis
func (cli *ZSClient) PowerOffBareMetal2Chassis(uuid string, params param.PowerOffBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.PowerOffBareMetal2ChassisEventView
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SdnControllerChangeHost operates on SdnControllerChangeHost
func (cli *ZSClient) SdnControllerChangeHost(uuid string, params param.SdnControllerChangeHostParam) (*view.SdnControllerInventoryView, error) {
	var resp view.SdnControllerChangeHostEventView
	if err := cli.Put("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateResourcePrice updates ResourcePrice
func (cli *ZSClient) UpdateResourcePrice(uuid string, params param.UpdateResourcePriceParam) (*view.PriceInventoryView, error) {
	var resp view.UpdateResourcePriceEventView
	if err := cli.Put("v1/billings/prices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachTagFromResources operates on TagFromResources
func (cli *ZSClient) DetachTagFromResources(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tags/{tagUuid}/resources", uuid, string(deleteMode))
}

// ChangeHostState changes HostState
func (cli *ZSClient) ChangeHostState(uuid string, params param.ChangeHostStateParam) (*view.HostInventoryView, error) {
	var resp view.ChangeHostStateEventView
	if err := cli.Put("v1/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateVmNicMac updates VmNicMac
func (cli *ZSClient) UpdateVmNicMac(uuid string, params param.UpdateVmNicMacParam) (*view.VmNicInventoryView, error) {
	var resp view.UpdateVmNicMacEventView
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteVmInstanceHaLevel deletes VmInstanceHaLevel
func (cli *ZSClient) DeleteVmInstanceHaLevel(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/ha-levels", uuid, string(deleteMode))
}

// DeleteResourcePrice deletes ResourcePrice
func (cli *ZSClient) DeleteResourcePrice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/prices/{uuid}", uuid, string(deleteMode))
}

// CleanUpBareMetal2Bonding operates on UpBareMetal2Bonding
func (cli *ZSClient) CleanUpBareMetal2Bonding(uuid string, params param.CleanUpBareMetal2BondingParam) (*view.CleanUpBaremetal2BondingEventView, error) {
	resp := view.CleanUpBaremetal2BondingEventView{}
	if err := cli.Put("v1/baremetal2/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteMetricData deletes MetricData
func (cli *ZSClient) DeleteMetricData(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/metrics", uuid, string(deleteMode))
}

// AddLabelToAlarm adds LabelToAlarm
func (cli *ZSClient) AddLabelToAlarm(params param.AddLabelToAlarmParam) (*view.AlarmLabelInventoryView, error) {
	var resp view.AddLabelToAlarmEventView
	if err := cli.Post("v1/zwatch/alarms/{alarmUuid}/labels", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SyncAliyunRouterInterfaceFromRemote operates on AliyunRouterInterfaceFromRemote
func (cli *ZSClient) SyncAliyunRouterInterfaceFromRemote(uuid string, params param.SyncAliyunRouterInterfaceFromRemoteParam) (*view.AliyunRouterInterfaceInventoryView, error) {
	resp := view.AliyunRouterInterfaceInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/router-interface/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportVmOvaPackage operates on VmOvaPackage
func (cli *ZSClient) ExportVmOvaPackage(params param.ExportVmOvaPackageParam) (*view.ImagePackageInventoryView, error) {
	var resp view.ExportVmOvaPackageEventView
	if err := cli.Post("v1/ovf/ova-packages", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RevertVmFromCdpBackup operates on VmFromCdpBackup
func (cli *ZSClient) RevertVmFromCdpBackup(uuid string, params param.RevertVmFromCdpBackupParam) (*view.RevertVmFromCdpBackupEventView, error) {
	resp := view.RevertVmFromCdpBackupEventView{}
	if err := cli.Put("v1/cdp-backups/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSFeiShuTestConnection operates on FeiShuTestConnection
func (cli *ZSClient) SNSFeiShuTestConnection(params param.SNSFeiShuTestConnectionParam) (*view.SNSFeiShuTestConnectionEventView, error) {
	resp := view.SNSFeiShuTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/feishu/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSchedulerExecutionReport gets SchedulerExecutionReport by uuid
func (cli *ZSClient) GetSchedulerExecutionReport(uuid string) (*view.GetSchedulerExecutionReportView, error) {
	var resp view.GetSchedulerExecutionReportView
	if err := cli.Get("v1/scheduler/report", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRuleFromConfigFile creates FirewallRuleFromConfigFile
func (cli *ZSClient) CreateFirewallRuleFromConfigFile(params param.CreateFirewallRuleFromConfigFileParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.Post("v1/vpcfirewalls/rules/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSupportedIdentityModels gets SupportedIdentityModels by uuid
func (cli *ZSClient) GetSupportedIdentityModels(uuid string) (*view.GetSupportedIdentityModelsView, error) {
	var resp view.GetSupportedIdentityModelsView
	if err := cli.Get("v1/identity-models", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddUserToGroup adds UserToGroup
func (cli *ZSClient) AddUserToGroup(params param.AddUserToGroupParam) (*view.AddUserToGroupEventView, error) {
	resp := view.AddUserToGroupEventView{}
	if err := cli.Post("v1/accounts/groups/{groupUuid}/users", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVRouterOspfArea updates VRouterOspfArea
func (cli *ZSClient) UpdateVRouterOspfArea(uuid string, params param.UpdateVRouterOspfAreaParam) (*view.RouterAreaInventoryView, error) {
	var resp view.UpdateVRouterOspfAreaEventView
	if err := cli.Put("v1/routerArea/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetPrimaryStorageTypes gets PrimaryStorageTypes by uuid
func (cli *ZSClient) GetPrimaryStorageTypes(uuid string) (*view.GetPrimaryStorageTypesView, error) {
	var resp view.GetPrimaryStorageTypesView
	if err := cli.Get("v1/primary-storage/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIPSecConnection queries IPSecConnection list
func (cli *ZSClient) QueryIPSecConnection(params *param.QueryParam) ([]view.IPsecConnectionInventoryView, error) {
	var resp []view.IPsecConnectionInventoryView
	return resp, cli.List("v1/ipsec", params, &resp)
}

// BatchDeleteVolumeSnapshot operates on DeleteVolumeSnapshot
func (cli *ZSClient) BatchDeleteVolumeSnapshot(uuid string, params param.BatchDeleteVolumeSnapshotParam) (*view.BatchDeleteVolumeSnapshotEventView, error) {
	resp := view.BatchDeleteVolumeSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/batch-delete", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReloadLicense operates on ReloadLicense
func (cli *ZSClient) ReloadLicense(uuid string, params param.ReloadLicenseParam) (*view.LicenseInventoryView, error) {
	var resp view.ReloadLicenseView
	if err := cli.Put("v1/licenses/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteNicQos deletes NicQos
func (cli *ZSClient) DeleteNicQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/nic-qos", uuid, string(deleteMode))
}

// ChangeL2NetworkVlanId changes L2NetworkVlanId
func (cli *ZSClient) ChangeL2NetworkVlanId(uuid string, params param.ChangeL2NetworkVlanIdParam) (*view.L2NetworkInventoryView, error) {
	var resp view.ChangeL2NetworkVlanIdEventView
	if err := cli.Put("v1/l2-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetResourceStackVmStatus gets ResourceStackVmStatus by uuid
func (cli *ZSClient) GetResourceStackVmStatus(uuid string) (*view.GetResourceStackVmStatusView, error) {
	var resp view.GetResourceStackVmStatusView
	if err := cli.Get("v1/cloudformation/stack/monitor/vmstatus", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachHybridKey operates on HybridKey
func (cli *ZSClient) DetachHybridKey(uuid string, params param.DetachHybridKeyParam) (*view.DetachHybridKeyEventView, error) {
	resp := view.DetachHybridKeyEventView{}
	if err := cli.Put("v1/hybrid/hybrid/key/{uuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveDnsFromVpcRouter removes DnsFromVpcRouter
func (cli *ZSClient) RemoveDnsFromVpcRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpc/virtual-routers/{uuid}/dns", uuid, string(deleteMode))
}

// DeleteHybridEipFromLocal deletes HybridEipFromLocal
func (cli *ZSClient) DeleteHybridEipFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/eip/{uuid}", uuid, string(deleteMode))
}

// GetAvailableTriggers gets AvailableTriggers by uuid
func (cli *ZSClient) GetAvailableTriggers(uuid string) (*view.SchedulerTriggerInventoryView, error) {
	var resp view.SchedulerTriggerInventoryView
	if err := cli.Get("v1/scheduler/triggers/available", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReimageVmInstance operates on ReimageVmInstance
func (cli *ZSClient) ReimageVmInstance(uuid string, params param.ReimageVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.ReimageVmInstanceEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateDatasets updates Datasets
func (cli *ZSClient) UpdateDatasets(uuid string, params param.UpdateDatasetsParam) (*view.UpdateDatasetsEventView, error) {
	resp := view.UpdateDatasetsEventView{}
	if err := cli.Put("v1/ai/datasets", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsSecurityGroupRuleFromRemote operates on EcsSecurityGroupRuleFromRemote
func (cli *ZSClient) SyncEcsSecurityGroupRuleFromRemote(uuid string, params param.SyncEcsSecurityGroupRuleFromRemoteParam) (*view.EcsSecurityGroupRuleInventoryView, error) {
	resp := view.EcsSecurityGroupRuleInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/security-group-rule/{uuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallRuleTemplate queries FirewallRuleTemplate list
func (cli *ZSClient) QueryFirewallRuleTemplate(params *param.QueryParam) ([]view.VpcFirewallRuleTemplateInventoryView, error) {
	var resp []view.VpcFirewallRuleTemplateInventoryView
	return resp, cli.List("v1/vpcfirewalls/rules/templates", params, &resp)
}

// SyncIdentityFromRemote operates on IdentityFromRemote
func (cli *ZSClient) SyncIdentityFromRemote(params param.SyncIdentityFromRemoteParam) (*view.SyncIdentityFromRemoteEventView, error) {
	var resp view.SyncIdentityFromRemoteEventView
	if err := cli.Get("v1/hybrid/identity-zone/{uuid}/sync", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageStoreBackupStorageQuota operates on ImageStoreBackupStorageQuota
func (cli *ZSClient) SetImageStoreBackupStorageQuota(uuid string, params param.SetImageStoreBackupStorageQuotaParam) (*view.SetImageStoreBackupStorageQuotaEventView, error) {
	resp := view.SetImageStoreBackupStorageQuotaEventView{}
	if err := cli.Put("v1/backup-storage/image-store/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeClusterState changes ClusterState
func (cli *ZSClient) ChangeClusterState(uuid string, params param.ChangeClusterStateParam) (*view.ClusterInventoryView, error) {
	var resp view.ChangeClusterStateEventView
	if err := cli.Put("v1/clusters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeVfNicHaState changes VfNicHaState
func (cli *ZSClient) ChangeVfNicHaState(uuid string, params param.ChangeVfNicHaStateParam) (*view.VmVfNicInventoryView, error) {
	var resp view.ChangeVfNicHaStateEventView
	if err := cli.Put("v1/vm-instances/nics/{vfNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateOvnControllerOffering creates OvnControllerOffering
func (cli *ZSClient) CreateOvnControllerOffering(params param.CreateOvnControllerOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	var resp view.CreateInstanceOfferingEventView
	if err := cli.Post("v1/instance-offerings/ovn", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetIAM2OrganizationVirtualIDNumber gets IAM2OrganizationVirtualIDNumber by uuid
func (cli *ZSClient) GetIAM2OrganizationVirtualIDNumber(uuid string) (*view.GetIAM2OrganizationVirtualIDNumberView, error) {
	var resp view.GetIAM2OrganizationVirtualIDNumberView
	if err := cli.Get("v1/iam2/organizations/{uuid}/virtualIDNumber", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsInstanceLocal deletes EcsInstanceLocal
func (cli *ZSClient) DeleteEcsInstanceLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/ecs/{uuid}", uuid, string(deleteMode))
}

// ChangePortMirrorState changes PortMirrorState
func (cli *ZSClient) ChangePortMirrorState(uuid string, params param.ChangePortMirrorStateParam) (*view.PortMirrorInventoryView, error) {
	var resp view.ChangePortMirrorStateEventView
	if err := cli.Put("v1/port-mirrors/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UnsubscribeSNSTopic operates on UnsubscribeSNSTopic
func (cli *ZSClient) UnsubscribeSNSTopic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/topics/{topicUuid}/endpoints/{endpointUuid}", uuid, string(deleteMode))
}

// SetNicQos operates on NicQos
func (cli *ZSClient) SetNicQos(uuid string, params param.SetNicQosParam) (*view.SetNicQosEventView, error) {
	resp := view.SetNicQosEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelLongJob operates on CancelLongJob
func (cli *ZSClient) CancelLongJob(uuid string, params param.CancelLongJobParam) (*view.CancelLongJobEventView, error) {
	resp := view.CancelLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRouteTableVpcVRouterCandidate gets RouteTableVpcVRouterCandidate by uuid
func (cli *ZSClient) GetRouteTableVpcVRouterCandidate(uuid string) (*view.VpcRouterVmInventoryView, error) {
	var resp view.VpcRouterVmInventoryView
	if err := cli.Get("v1/vpc/virtual-routers/get-vpc-candidate", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateAccountBilling operates on AccountBilling
func (cli *ZSClient) GenerateAccountBilling(uuid string, params param.GenerateAccountBillingParam) (*view.GenerateAccountBillingEventView, error) {
	resp := view.GenerateAccountBillingEventView{}
	if err := cli.Put("v1/billings/accounts/{accountUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAliyunVirtualRouterFromRemote operates on AliyunVirtualRouterFromRemote
func (cli *ZSClient) SyncAliyunVirtualRouterFromRemote(uuid string, params param.SyncAliyunVirtualRouterFromRemoteParam) (*view.VpcVirtualRouterInventoryView, error) {
	resp := view.VpcVirtualRouterInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/vrouter/{vpcUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsVpcFromLocal queries EcsVpcFromLocal list
func (cli *ZSClient) QueryEcsVpcFromLocal(params *param.QueryParam) ([]view.EcsVpcInventoryView, error) {
	var resp []view.EcsVpcInventoryView
	return resp, cli.List("v1/hybrid/aliyun/vpc", params, &resp)
}

// GetInvocationRecords gets InvocationRecords by uuid
func (cli *ZSClient) GetInvocationRecords(uuid string) (*view.InvocationRecordView, error) {
	var resp view.InvocationRecordView
	if err := cli.Get("v1/scripts/aliyun-invocations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeRoleState changes RoleState
func (cli *ZSClient) ChangeRoleState(uuid string, params param.ChangeRoleStateParam) (*view.RoleInventoryView, error) {
	var resp view.ChangeRoleStateEventView
	if err := cli.Put("v1/identities/roles/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVRouterFlowCounter gets VRouterFlowCounter by uuid
func (cli *ZSClient) GetVRouterFlowCounter(uuid string) (*view.GetVRouterFlowCounterView, error) {
	var resp view.GetVRouterFlowCounterView
	if err := cli.Get("v1/flowmeters/{vRouterUuid}/counter", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAliyunRouterInterfaceRemote creates AliyunRouterInterfaceRemote
func (cli *ZSClient) CreateAliyunRouterInterfaceRemote(params param.CreateAliyunRouterInterfaceRemoteParam) (*view.AliyunRouterInterfaceInventoryView, error) {
	var resp view.CreateAliyunRouterInterfaceRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/router-interface", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetBareMetal2SupportedBootMode gets BareMetal2SupportedBootMode by uuid
func (cli *ZSClient) GetBareMetal2SupportedBootMode(uuid string) (*view.GetBareMetal2SupportedBootModeView, error) {
	var resp view.GetBareMetal2SupportedBootModeView
	if err := cli.Get("v1/baremetal2/chassis/supported-boot-modes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostPowerStatus gets HostPowerStatus by uuid
func (cli *ZSClient) GetHostPowerStatus(uuid string) (*view.HostIpmiInventoryView, error) {
	var resp view.GetHostPowerStatusEventView
	if err := cli.Get("v1/hosts/power/{uuid}/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetChainTask gets ChainTask by uuid
func (cli *ZSClient) GetChainTask(uuid string) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.Get("v1/core/task-details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch updates ConnectionBetweenL3NetWorkAndAliyunVSwitch
func (cli *ZSClient) UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch(uuid string, params param.UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParam) (*view.ConnectionRelationShipInventoryView, error) {
	var resp view.UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchEventView
	if err := cli.Put("v1/hybrid/aliyun/connections/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeHostPassword changes HostPassword
func (cli *ZSClient) ChangeHostPassword(uuid string, params param.ChangeHostPasswordParam) (*view.ChangeHostPasswordEventView, error) {
	resp := view.ChangeHostPasswordEventView{}
	if err := cli.Put("v1/hosts/kvm/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSlbInstance creates SlbInstance
func (cli *ZSClient) CreateSlbInstance(params param.CreateSlbInstanceParam) (*view.SlbVmInstanceInventoryView, error) {
	var resp view.CreateSlbInstanceEventView
	if err := cli.Post("v1/load-balancers/slb/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangePortForwardingRuleState changes PortForwardingRuleState
func (cli *ZSClient) ChangePortForwardingRuleState(uuid string, params param.ChangePortForwardingRuleStateParam) (*view.PortForwardingRuleInventoryView, error) {
	var resp view.ChangePortForwardingRuleStateEventView
	if err := cli.Put("v1/port-forwarding/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// IsLicenseServer operates on IsLicenseServer
func (cli *ZSClient) IsLicenseServer(params param.IsLicenseServerParam) (*view.IsLicenseServerView, error) {
	var resp view.IsLicenseServerView
	if err := cli.Get("v1/license-server/is-server", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryLabelValues operates on PrometheusQueryLabelValues
func (cli *ZSClient) PrometheusQueryLabelValues(params param.PrometheusQueryLabelValuesParam) (*view.PrometheusQueryLabelValuesView, error) {
	var resp view.PrometheusQueryLabelValuesView
	if err := cli.Get("v1/prometheus/labels", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateClusterSupportDRS operates on ClusterSupportDRS
func (cli *ZSClient) ValidateClusterSupportDRS(params param.ValidateClusterSupportDRSParam) (*view.ValidateClusterSupportDRSView, error) {
	var resp view.ValidateClusterSupportDRSView
	if err := cli.Get("v1/clusters/{clusterUuid}/drs/valid", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShrinkVolumeSnapshot operates on ShrinkVolumeSnapshot
func (cli *ZSClient) ShrinkVolumeSnapshot(uuid string, params param.ShrinkVolumeSnapshotParam) (*view.ShrinkVolumeSnapshotEventView, error) {
	resp := view.ShrinkVolumeSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/shrink/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddHostToHostSchedulingRuleGroup adds HostToHostSchedulingRuleGroup
func (cli *ZSClient) AddHostToHostSchedulingRuleGroup(params param.AddHostToHostSchedulingRuleGroupParam) (*view.AddHostToHostSchedulingRuleGroupEventView, error) {
	resp := view.AddHostToHostSchedulingRuleGroupEventView{}
	if err := cli.Post("v1/hostSchedulingRuleGroup/{hostGroupUuid}/host/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBuildApp creates BuildApp
func (cli *ZSClient) CreateBuildApp(params param.CreateBuildAppParam) (*view.BuildApplicationInventoryView, error) {
	var resp view.CreateBuildAppEventView
	if err := cli.Post("v1/appcenter/buildapp", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryIdentityZoneFromLocal queries IdentityZoneFromLocal list
func (cli *ZSClient) QueryIdentityZoneFromLocal(params *param.QueryParam) ([]view.IdentityZoneInventoryView, error) {
	var resp []view.IdentityZoneInventoryView
	return resp, cli.List("v1/hybrid/identity-zone", params, &resp)
}

// GetVmNicAttachedNetworkService gets VmNicAttachedNetworkService by uuid
func (cli *ZSClient) GetVmNicAttachedNetworkService(uuid string) (*view.GetVmNicAttachedNetworkServiceView, error) {
	var resp view.GetVmNicAttachedNetworkServiceView
	if err := cli.Get("v1/vm-instances/nics/{vmNicUuid}/attached-networkservices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmHostname gets VmHostname by uuid
func (cli *ZSClient) GetVmHostname(uuid string) (*view.GetVmHostnameView, error) {
	var resp view.GetVmHostnameView
	if err := cli.Get("v1/vm-instances/{uuid}/hostnames", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSchedulerJobsToSchedulerJobGroup adds SchedulerJobsToSchedulerJobGroup
func (cli *ZSClient) AddSchedulerJobsToSchedulerJobGroup(params param.AddSchedulerJobsToSchedulerJobGroupParam) (*view.SchedulerJobGroupJobRefInventoryView, error) {
	resp := view.SchedulerJobGroupJobRefInventoryView{}
	if err := cli.Post("v1/scheduler/jobgroups/{schedulerJobGroupUuid}/job", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL3NetworkFromVm operates on L3NetworkFromVm
func (cli *ZSClient) DetachL3NetworkFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/nics/{vmNicUuid}", uuid, string(deleteMode))
}

// DeleteVpcUserVpnGatewayLocal deletes VpcUserVpnGatewayLocal
func (cli *ZSClient) DeleteVpcUserVpnGatewayLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/user-gateway/{uuid}", uuid, string(deleteMode))
}

// CreateVRouterOspfArea creates VRouterOspfArea
func (cli *ZSClient) CreateVRouterOspfArea(params param.CreateVRouterOspfAreaParam) (*view.RouterAreaInventoryView, error) {
	var resp view.CreateVRouterOspfAreaEventView
	if err := cli.Post("v1/routerArea", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetSecurityMachineKey operates on SecurityMachineKey
func (cli *ZSClient) SetSecurityMachineKey(params param.SetSecurityMachineKeyParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.Post("v1/secret-resource-pool-token/set/{uuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOAuthClient creates OAuthClient
func (cli *ZSClient) CreateOAuthClient(params param.CreateOAuthClientParam) (*view.OAuth2ClientInventoryView, error) {
	var resp view.CreateOAuthClientEventView
	if err := cli.Post("v1/create/oauth2/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVpcAttachedEip gets VpcAttachedEip by uuid
func (cli *ZSClient) GetVpcAttachedEip(uuid string) (*view.EipInventoryView, error) {
	var resp view.EipInventoryView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-eip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveSchedulerJobFromSchedulerTrigger removes SchedulerJobFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobFromSchedulerTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobs/{schedulerJobUuid}/scheduler/triggers/{schedulerTriggerUuid}", uuid, string(deleteMode))
}

// ChangeMediaState changes MediaState
func (cli *ZSClient) ChangeMediaState(uuid string, params param.ChangeMediaStateParam) (*view.MediaInventoryView, error) {
	var resp view.ChangeMediaStateEventView
	if err := cli.Put("v1/media/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeIPSecConnectionState changes IPSecConnectionState
func (cli *ZSClient) ChangeIPSecConnectionState(uuid string, params param.ChangeIPSecConnectionStateParam) (*view.IPsecConnectionInventoryView, error) {
	var resp view.ChangeIPSecConnectionStateEventView
	if err := cli.Put("v1/ipsec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryAliyunDiskFromLocal queries AliyunDiskFromLocal list
func (cli *ZSClient) QueryAliyunDiskFromLocal(params *param.QueryParam) ([]view.AliyunDiskInventoryView, error) {
	var resp []view.AliyunDiskInventoryView
	return resp, cli.List("v1/hybrid/aliyun/disk", params, &resp)
}

// QueryEcsSecurityGroupRuleFromLocal queries EcsSecurityGroupRuleFromLocal list
func (cli *ZSClient) QueryEcsSecurityGroupRuleFromLocal(params *param.QueryParam) ([]view.EcsSecurityGroupRuleInventoryView, error) {
	var resp []view.EcsSecurityGroupRuleInventoryView
	return resp, cli.List("v1/hybrid/aliyun/security-group-rule", params, &resp)
}

// StopAllResourcesInIAM2Project stops AllResourcesInIAM2Project
func (cli *ZSClient) StopAllResourcesInIAM2Project(uuid string, params param.StopAllResourcesInIAM2ProjectParam) (*view.StopAllResourcesInIAM2ProjectEventView, error) {
	resp := view.StopAllResourcesInIAM2ProjectEventView{}
	if err := cli.Put("v1/iam2/projects/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmNetworkConfig updates VmNetworkConfig
func (cli *ZSClient) UpdateVmNetworkConfig(uuid string, params param.UpdateVmNetworkConfigParam) (*view.UpdateVmNetworkConfigEventView, error) {
	resp := view.UpdateVmNetworkConfigEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/update-nic-config", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveTicketTypesFromTicketFlowCollection removes TicketTypesFromTicketFlowCollection
func (cli *ZSClient) RemoveTicketTypesFromTicketFlowCollection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tickets/flow-collections/{ticketFlowCollectionUuid}/ticket-types", uuid, string(deleteMode))
}

// DeleteEcsVSwitchRemote deletes EcsVSwitchRemote
func (cli *ZSClient) DeleteEcsVSwitchRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vswitch/remote/{uuid}", uuid, string(deleteMode))
}

// SetVmStaticIp operates on VmStaticIp
func (cli *ZSClient) SetVmStaticIp(uuid string, params param.SetVmStaticIpParam) (*view.SetVmStaticIpEventView, error) {
	resp := view.SetVmStaticIpEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmSshKey gets VmSshKey by uuid
func (cli *ZSClient) GetVmSshKey(uuid string) (*view.GetVmSshKeyView, error) {
	var resp view.GetVmSshKeyView
	if err := cli.Get("v1/vm-instances/{uuid}/ssh-keys", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmGuestToolsInfo gets VmGuestToolsInfo by uuid
func (cli *ZSClient) GetVmGuestToolsInfo(uuid string) (*view.GetVmGuestToolsInfoView, error) {
	var resp view.GetVmGuestToolsInfoView
	if err := cli.Get("v1/vm-instances/{uuid}/guest-tools-infos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateDiskOfferingUserConfig operates on DiskOfferingUserConfig
func (cli *ZSClient) ValidateDiskOfferingUserConfig(uuid string, params param.ValidateDiskOfferingUserConfigParam) (*view.ValidateDiskOfferingUserConfigEventView, error) {
	resp := view.ValidateDiskOfferingUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcVpnGatewayLocal deletes VpcVpnGatewayLocal
func (cli *ZSClient) DeleteVpcVpnGatewayLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-gateway/{uuid}", uuid, string(deleteMode))
}

// SetVmRDP operates on VmRDP
func (cli *ZSClient) SetVmRDP(uuid string, params param.SetVmRDPParam) (*view.SetVmRDPEventView, error) {
	resp := view.SetVmRDPEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunSchedulerTrigger operates on RunSchedulerTrigger
func (cli *ZSClient) RunSchedulerTrigger(uuid string, params param.RunSchedulerTriggerParam) (*view.RunSchedulerTriggerEventView, error) {
	resp := view.RunSchedulerTriggerEventView{}
	if err := cli.Put("v1/scheduler/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAliyunVpcVirtualRouterEntryRemote creates AliyunVpcVirtualRouterEntryRemote
func (cli *ZSClient) CreateAliyunVpcVirtualRouterEntryRemote(params param.CreateAliyunVpcVirtualRouterEntryRemoteParam) (*view.VpcVirtualRouteEntryInventoryView, error) {
	var resp view.CreateAliyunVpcVirtualRouterEntryRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/route-entry", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PowerOnHost operates on PowerOnHost
func (cli *ZSClient) PowerOnHost(uuid string, params param.PowerOnHostParam) (*view.HostInventoryView, error) {
	var resp view.PowerOnHostEventView
	if err := cli.Put("v1/hosts/power/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteAliyunSnapshotFromRemote deletes AliyunSnapshotFromRemote
func (cli *ZSClient) DeleteAliyunSnapshotFromRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/snapshot/{uuid}/remote", uuid, string(deleteMode))
}

// RemoveCertificateFromLoadBalancerListener removes CertificateFromLoadBalancerListener
func (cli *ZSClient) RemoveCertificateFromLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/certificate", uuid, string(deleteMode))
}

// GetPortForwardingAttachableVmNics gets PortForwardingAttachableVmNics by uuid
func (cli *ZSClient) GetPortForwardingAttachableVmNics(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.Get("v1/port-forwarding/{ruleUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveRendezvousPointFromMulticastRouter removes RendezvousPointFromMulticastRouter
func (cli *ZSClient) RemoveRendezvousPointFromMulticastRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/multicast/virtual-routers/{uuid}/RendezvousPoint", uuid, string(deleteMode))
}

// AddIAM2VirtualIDsToProject adds IAM2VirtualIDsToProject
func (cli *ZSClient) AddIAM2VirtualIDsToProject(params param.AddIAM2VirtualIDsToProjectParam) (*view.AddIAM2VirtualIDsToProjectEventView, error) {
	resp := view.AddIAM2VirtualIDsToProjectEventView{}
	if err := cli.Post("v1/iam2/projects/{projectUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SubscribeEvent operates on SubscribeEvent
func (cli *ZSClient) SubscribeEvent(params param.SubscribeEventParam) (*view.EventSubscriptionInventoryView, error) {
	var resp view.SubscribeEventEventView
	if err := cli.Post("v1/zwatch/events/subscriptions", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetPrimaryStorageCandidatesForVolumeMigration gets PrimaryStorageCandidatesForVolumeMigration by uuid
func (cli *ZSClient) GetPrimaryStorageCandidatesForVolumeMigration(uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.Get("v1/primary-storage/volumes/{volumeUuid}/migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradeBackupStorageCdpTasks operates on UpgradeBackupStorageCdpTasks
func (cli *ZSClient) UpgradeBackupStorageCdpTasks(uuid string, params param.UpgradeBackupStorageCdpTasksParam) (*view.UpgradeBackupStorageCdpTasksEventView, error) {
	resp := view.UpgradeBackupStorageCdpTasksEventView{}
	if err := cli.Put("v1/cdp-task/upgrade/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVxlanL2Network deletes VxlanL2Network
func (cli *ZSClient) DeleteVxlanL2Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/vxlan/{uuid}", uuid, string(deleteMode))
}

// RemoveVmFromAffinityGroup removes VmFromAffinityGroup
func (cli *ZSClient) RemoveVmFromAffinityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/affinity-groups/{affinityGroupUuid}/vm-instances", uuid, string(deleteMode))
}

// SetVolumeIoThreadPin operates on VolumeIoThreadPin
func (cli *ZSClient) SetVolumeIoThreadPin(uuid string, params param.SetVolumeIoThreadPinParam) (*view.SetVolumeIoThreadPinEventView, error) {
	resp := view.SetVolumeIoThreadPinEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePriorityConfig updates PriorityConfig
func (cli *ZSClient) UpdatePriorityConfig(uuid string, params param.UpdatePriorityConfigParam) (*view.UpdatePriorityConfigEventView, error) {
	resp := view.UpdatePriorityConfigEventView{}
	if err := cli.Put("v1/vm-priority-config/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IdentifyHost operates on IdentifyHost
func (cli *ZSClient) IdentifyHost(uuid string, params param.IdentifyHostParam) (*view.IdentifyHostEventView, error) {
	resp := view.IdentifyHostEventView{}
	if err := cli.Put("v1/hosts/kvm/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromVolumeBackup creates RootVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeBackup(params param.CreateRootVolumeTemplateFromVolumeBackupParam) (*view.ImageInventoryView, error) {
	var resp view.CreateRootVolumeTemplateFromVolumeBackupEventView
	if err := cli.Post("v1/images/root-volume-templates/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CheckFirewallRuleConfigFile operates on FirewallRuleConfigFile
func (cli *ZSClient) CheckFirewallRuleConfigFile(params param.CheckFirewallRuleConfigFileParam) (*view.CheckFirewallRuleConfigFileView, error) {
	resp := view.CheckFirewallRuleConfigFileView{}
	if err := cli.Post("v1/vpcfirewalls/rules/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmConsoleAddress gets VmConsoleAddress by uuid
func (cli *ZSClient) GetVmConsoleAddress(uuid string) (*view.GetVmConsoleAddressView, error) {
	var resp view.GetVmConsoleAddressView
	if err := cli.Get("v1/vm-instances/{uuid}/console-addresses", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLoadBalancerListenerACLEntries gets LoadBalancerListenerACLEntries by uuid
func (cli *ZSClient) GetLoadBalancerListenerACLEntries(uuid string) (*view.StringView, error) {
	var resp view.StringView
	if err := cli.Get("v1/load-balancers/listeners/access-control-lists/entries", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHostIommuState updates HostIommuState
func (cli *ZSClient) UpdateHostIommuState(uuid string, params param.UpdateHostIommuStateParam) (*view.UpdateHostIommuStateEventView, error) {
	resp := view.UpdateHostIommuStateEventView{}
	if err := cli.Put("v1/pci-device/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsubscribeEvent operates on UnsubscribeEvent
func (cli *ZSClient) UnsubscribeEvent(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions/{uuid}", uuid, string(deleteMode))
}

// CreateObservabilityServer creates ObservabilityServer
func (cli *ZSClient) CreateObservabilityServer(params param.CreateObservabilityServerParam) (*view.VmInstanceInventoryView, error) {
	var resp view.CreateObservabilityServerEventView
	if err := cli.Post("v1/observability-servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveMonFromCephPrimaryStorage removes MonFromCephPrimaryStorage
func (cli *ZSClient) RemoveMonFromCephPrimaryStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/ceph/{uuid}/mons", uuid, string(deleteMode))
}

// GetVmsSchedulingStateFromSchedulingRule gets VmsSchedulingStateFromSchedulingRule by uuid
func (cli *ZSClient) GetVmsSchedulingStateFromSchedulingRule(uuid string) (*view.GetVmsSchedulingStateFromSchedulingRuleView, error) {
	var resp view.GetVmsSchedulingStateFromSchedulingRuleView
	if err := cli.Get("v1/get/vms/schedulingState/from/SchedulingRule", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAlarmState changes AlarmState
func (cli *ZSClient) ChangeAlarmState(uuid string, params param.ChangeAlarmStateParam) (*view.AlarmInventoryView, error) {
	var resp view.ChangeAlarmStateEventView
	if err := cli.Put("v1/zwatch/alarms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetLocalStorageHostDiskCapacity gets LocalStorageHostDiskCapacity by uuid
func (cli *ZSClient) GetLocalStorageHostDiskCapacity(uuid string) (*view.HostDiskCapacityView, error) {
	var resp view.HostDiskCapacityView
	if err := cli.Get("v1/primary-storage/local-storage/{primaryStorageUuid}/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmSshKey deletes VmSshKey
func (cli *ZSClient) DeleteVmSshKey(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/ssh-keys", uuid, string(deleteMode))
}

// GetPolicyRouteRuleSetFromVirtualRouter gets PolicyRouteRuleSetFromVirtualRouter by uuid
func (cli *ZSClient) GetPolicyRouteRuleSetFromVirtualRouter(uuid string) (*view.PolicyRouteRuleSetInventoryView, error) {
	var resp view.PolicyRouteRuleSetInventoryView
	if err := cli.Get("v1/policy-routes/rulesets/virtualrouter/{vmInstanceUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVxlanPoolRemoteVtep deletes VxlanPoolRemoteVtep
func (cli *ZSClient) DeleteVxlanPoolRemoteVtep(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}/delete/remote-vtep-ip", uuid, string(deleteMode))
}

// RemoveAttributesFromIAM2Project removes AttributesFromIAM2Project
func (cli *ZSClient) RemoveAttributesFromIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{uuid}/attributes", uuid, string(deleteMode))
}

// RecoverDataVolume operates on DataVolume
func (cli *ZSClient) RecoverDataVolume(uuid string, params param.RecoverDataVolumeParam) (*view.VolumeInventoryView, error) {
	var resp view.RecoverDataVolumeEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveIAM2VirtualIDsFromGroup removes IAM2VirtualIDsFromGroup
func (cli *ZSClient) RemoveIAM2VirtualIDsFromGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups/{groupUuid}/virtual-ids", uuid, string(deleteMode))
}

// QueryEventRecord queries EventRecord list
func (cli *ZSClient) QueryEventRecord(params *param.QueryParam) ([]view.EventRecordsInventoryView, error) {
	var resp []view.EventRecordsInventoryView
	return resp, cli.List("v1/zwatch/event-records", params, &resp)
}

// AttachBareMetal2ProvisionNetworkToCluster operates on BareMetal2ProvisionNetworkToCluster
func (cli *ZSClient) AttachBareMetal2ProvisionNetworkToCluster(params param.AttachBareMetal2ProvisionNetworkToClusterParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	var resp view.AttachBareMetal2ProvisionNetworkToClusterEventView
	if err := cli.Post("v1/baremetal2/clusters/{clusterUuid}/provision-networks/{networkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryIAM2LdapBinding queries IAM2LdapBinding list
func (cli *ZSClient) QueryIAM2LdapBinding(params *param.QueryParam) ([]view.LdapResourceRefInventoryView, error) {
	var resp []view.LdapResourceRefInventoryView
	return resp, cli.List("v1/iam2/ldap/bindings", params, &resp)
}

// ProvisionSlbInstance operates on ProvisionSlbInstance
func (cli *ZSClient) ProvisionSlbInstance(uuid string, params param.ProvisionSlbInstanceParam) (*view.SlbGroupInventoryView, error) {
	var resp view.ProvisionSlbGroupInstanceEventView
	if err := cli.Put("v1/load-balancers/slb/instances/{uuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetVmUserDefinedXmlHookScript operates on VmUserDefinedXmlHookScript
func (cli *ZSClient) SetVmUserDefinedXmlHookScript(uuid string, params param.SetVmUserDefinedXmlHookScriptParam) (*view.SetVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.SetVmUserDefinedXmlHookScriptEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachProvisionNicFromBonding operates on ProvisionNicFromBonding
func (cli *ZSClient) DetachProvisionNicFromBonding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/bm-instances/bm2-bondings/{uuid}", uuid, string(deleteMode))
}

// LoginIAM2Platform operates on IAM2Platform
func (cli *ZSClient) LoginIAM2Platform(uuid string, params param.LoginIAM2PlatformParam) (*view.SessionInventoryView, error) {
	var resp view.LoginIAM2PlatformView
	if err := cli.Put("v1/iam2/platform/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetHostAllocatorStrategies gets HostAllocatorStrategies by uuid
func (cli *ZSClient) GetHostAllocatorStrategies(uuid string) (*view.GetHostAllocatorStrategiesView, error) {
	var resp view.GetHostAllocatorStrategiesView
	if err := cli.Get("v1/hosts/allocators/strategies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInterfaceServiceTypeStatistic gets InterfaceServiceTypeStatistic by uuid
func (cli *ZSClient) GetInterfaceServiceTypeStatistic(uuid string) (*view.GetInterfaceServiceTypeStatisticView, error) {
	var resp view.GetInterfaceServiceTypeStatisticView
	if err := cli.Get("v1/hosts/hosts-network-interfaces/service-type-statistic", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartConnectionBetweenAliyunRouterInterface starts ConnectionBetweenAliyunRouterInterface
func (cli *ZSClient) StartConnectionBetweenAliyunRouterInterface(uuid string, params param.StartConnectionBetweenAliyunRouterInterfaceParam) (*view.AliyunRouterInterfaceInventoryView, error) {
	var resp view.StartConnectionBetweenAliyunRouterInterfaceEventView
	if err := cli.Put("v1/hybrid/aliyun/router-interface/{vbrInterfaceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteModels deletes Models
func (cli *ZSClient) DeleteModels(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models", uuid, string(deleteMode))
}

// ListVmsFromSchedulingState operates on ListVmsFromSchedulingState
func (cli *ZSClient) ListVmsFromSchedulingState(params param.ListVmsFromSchedulingStateParam) (*view.ListVmsFromSchedulingStateView, error) {
	resp := view.ListVmsFromSchedulingStateView{}
	if err := cli.Post("v1/list/vms/from/executeState", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromVolumeSnapshot creates RootVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeSnapshot(params param.CreateRootVolumeTemplateFromVolumeSnapshotParam) (*view.ImageInventoryView, error) {
	var resp view.CreateRootVolumeTemplateFromVolumeSnapshotEventView
	if err := cli.Post("v1/images/root-volume-templates/from/volume-snapshots/{snapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AllocateHostResource operates on HostResource
func (cli *ZSClient) AllocateHostResource(params param.AllocateHostResourceParam) (*view.AllocateHostResourceEventView, error) {
	resp := view.AllocateHostResourceEventView{}
	if err := cli.Post("v1/hosts/{uuid}/allocate-resource", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateLdapEntryForBinding gets CandidateLdapEntryForBinding by uuid
func (cli *ZSClient) GetCandidateLdapEntryForBinding(uuid string) (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.Get("v1/ldap/entries/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckElaborationContent operates on ElaborationContent
func (cli *ZSClient) CheckElaborationContent(params param.CheckElaborationContentParam) (*view.CheckElaborationContentView, error) {
	resp := view.CheckElaborationContentView{}
	if err := cli.Post("v1/errorcode/elaborations/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmConsolePassword deletes VmConsolePassword
func (cli *ZSClient) DeleteVmConsolePassword(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/console-password", uuid, string(deleteMode))
}

// CreateVmBackup creates VmBackup
func (cli *ZSClient) CreateVmBackup(params param.CreateVmBackupParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.Post("v1/volumes/{rootVolumeUuid}/vm-backups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageLicenseInfo gets PrimaryStorageLicenseInfo by uuid
func (cli *ZSClient) GetPrimaryStorageLicenseInfo(uuid string) (*view.GetPrimaryStorageLicenseInfoView, error) {
	var resp view.GetPrimaryStorageLicenseInfoView
	if err := cli.Get("v1/primary-storage/{uuid}/license", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEncryptedField gets EncryptedField by uuid
func (cli *ZSClient) GetEncryptedField(uuid string) (*view.GetEncryptedFieldView, error) {
	var resp view.GetEncryptedFieldView
	if err := cli.Get("v1/encrypted/fields", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveIAM2ProjectLoginExpired removes IAM2ProjectLoginExpired
func (cli *ZSClient) RemoveIAM2ProjectLoginExpired(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/remove/login-expired/{uuid}/actions", uuid, string(deleteMode))
}

// CleanInvalidLdapBinding operates on InvalidLdapBinding
func (cli *ZSClient) CleanInvalidLdapBinding(uuid string, params param.CleanInvalidLdapBindingParam) (*view.AccountInventoryView, error) {
	resp := view.AccountInventoryView{}
	if err := cli.Put("v1/ldap/bindings/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachBaremetalPxeServerToCluster operates on BaremetalPxeServerToCluster
func (cli *ZSClient) AttachBaremetalPxeServerToCluster(params param.AttachBaremetalPxeServerToClusterParam) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.AttachBaremetalPxeServerToClusterEventView
	if err := cli.Post("v1/clusters/{clusterUuid}/pxeservers/{pxeServerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVmStartingCandidateClustersHosts gets VmStartingCandidateClustersHosts by uuid
func (cli *ZSClient) GetVmStartingCandidateClustersHosts(uuid string) (*view.GetVmStartingCandidateClustersHostsView, error) {
	var resp view.GetVmStartingCandidateClustersHostsView
	if err := cli.Get("v1/vm-instances/{uuid}/starting-target-hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoverVmBackupFromImageStoreBackupStorage operates on VmBackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverVmBackupFromImageStoreBackupStorage(uuid string, params param.RecoverVmBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.Put("v1/vm-backups/{groupUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachIAM2ProjectFromIAM2Organization operates on IAM2ProjectFromIAM2Organization
func (cli *ZSClient) DetachIAM2ProjectFromIAM2Organization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{projectUuid}/iam2/organizations", uuid, string(deleteMode))
}

// DiscoverExternalPrimaryStorage operates on DiscoverExternalPrimaryStorage
func (cli *ZSClient) DiscoverExternalPrimaryStorage(params param.DiscoverExternalPrimaryStorageParam) (*view.ExternalPrimaryStorageInventoryView, error) {
	var resp view.DiscoverExternalPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/addon/discover", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVolumeIoThreadPin gets VolumeIoThreadPin by uuid
func (cli *ZSClient) GetVolumeIoThreadPin(uuid string) (*view.GetVolumeIoThreadPinView, error) {
	var resp view.GetVolumeIoThreadPinView
	if err := cli.Get("v1/volumes/{uuid}/io-thread-pin", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConnectionAccessPointFromRemote gets ConnectionAccessPointFromRemote by uuid
func (cli *ZSClient) GetConnectionAccessPointFromRemote(uuid string) (*view.ConnectionAccessPointInventoryView, error) {
	var resp view.ConnectionAccessPointInventoryView
	if err := cli.Get("v1/hybrid/aliyun/access-point{dataCenterUuid}/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedOspf gets VpcAttachedOspf by uuid
func (cli *ZSClient) GetVpcAttachedOspf(uuid string) (*view.NetworkRouterAreaRefInventoryView, error) {
	var resp view.NetworkRouterAreaRefInventoryView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-ospf", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOffHost operates on PowerOffHost
func (cli *ZSClient) PowerOffHost(uuid string, params param.PowerOffHostParam) (*view.PowerOffHostEventView, error) {
	resp := view.PowerOffHostEventView{}
	if err := cli.Put("v1/hosts/power-off/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveIAM2VirtualIDGroupFromProjects removes IAM2VirtualIDGroupFromProjects
func (cli *ZSClient) RemoveIAM2VirtualIDGroupFromProjects(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups", uuid, string(deleteMode))
}

// UpdateVmUserDefinedXmlHookScript updates VmUserDefinedXmlHookScript
func (cli *ZSClient) UpdateVmUserDefinedXmlHookScript(uuid string, params param.UpdateVmUserDefinedXmlHookScriptParam) (*view.XmlHookInventoryView, error) {
	var resp view.UpdateVmUserDefinedXmlHookScriptEventView
	if err := cli.Put("v1/vm-instances/xml-hook-script", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetIAM2ProjectContainerImageTags gets IAM2ProjectContainerImageTags by uuid
func (cli *ZSClient) GetIAM2ProjectContainerImageTags(uuid string) (*view.ContainerImageTagInventoryView, error) {
	var resp view.ContainerImageTagInventoryView
	if err := cli.Get("v1/iam2/project/{projectId}/repository/{repositoryId}/image/{imageName}/tag", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunDiskFromRemote deletes AliyunDiskFromRemote
func (cli *ZSClient) DeleteAliyunDiskFromRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/disk/{uuid}/remote", uuid, string(deleteMode))
}

// GetVersion gets Version by uuid
func (cli *ZSClient) GetVersion(uuid string) (*view.GetVersionView, error) {
	var resp view.GetVersionView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateBackupStorageForCreatingImage gets CandidateBackupStorageForCreatingImage by uuid
func (cli *ZSClient) GetCandidateBackupStorageForCreatingImage(uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.Get("v1/images/candidate-backup-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachAutoScalingTemplateToGroup operates on AutoScalingTemplateToGroup
func (cli *ZSClient) AttachAutoScalingTemplateToGroup(params param.AttachAutoScalingTemplateToGroupParam) (*view.AutoScalingGroupInventoryView, error) {
	var resp view.AttachAutoScalingTemplateToGroupEventView
	if err := cli.Post("v1/autoscaling/template/{uuid}/groups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetCpuMemoryCapacity gets CpuMemoryCapacity by uuid
func (cli *ZSClient) GetCpuMemoryCapacity(uuid string) (*view.GetCpuMemoryCapacityView, error) {
	var resp view.GetCpuMemoryCapacityView
	if err := cli.Get("v1/hosts/capacities/cpu-memory", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIntegrityResource adds IntegrityResource
func (cli *ZSClient) AddIntegrityResource(params param.AddIntegrityResourceParam) (*view.AddIntegrityResourceEventView, error) {
	resp := view.AddIntegrityResourceEventView{}
	if err := cli.Post("v1/integrity/resource/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckVipPortAvailability operates on VipPortAvailability
func (cli *ZSClient) CheckVipPortAvailability(params param.CheckVipPortAvailabilityParam) (*view.CheckVipPortAvailabilityView, error) {
	var resp view.CheckVipPortAvailabilityView
	if err := cli.Get("v1/vips/{vipUuid}/check-port-availability", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateClustersForAttachingL2Network gets CandidateClustersForAttachingL2Network by uuid
func (cli *ZSClient) GetCandidateClustersForAttachingL2Network(uuid string) (*view.ClusterInventoryView, error) {
	var resp view.ClusterInventoryView
	if err := cli.Get("v1/l2-networks/{l2NetworkUuid}/cluster-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckScsiLunClusterStatus operates on ScsiLunClusterStatus
func (cli *ZSClient) CheckScsiLunClusterStatus(uuid string, params param.CheckScsiLunClusterStatusParam) (*view.ScsiLunClusterStatusInventoryView, error) {
	var resp view.CheckScsiLunClusterStatusView
	if err := cli.Put("v1/storage-devices/scsi-lun/{uuid}/cluster/{clusterUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CheckBatchDataIntegrity operates on BatchDataIntegrity
func (cli *ZSClient) CheckBatchDataIntegrity(params param.CheckBatchDataIntegrityParam) (*view.CheckBatchDataIntegrityView, error) {
	var resp view.CheckBatchDataIntegrityView
	if err := cli.Get("v1/check/batch/data/integrity/", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAutoScalingGroupRemovalInstanceRule updates AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupRemovalInstanceRule(uuid string, params param.UpdateAutoScalingGroupRemovalInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	var resp view.UpdateAutoScalingRuleEventView
	if err := cli.Put("v1/autoscaling/rules/removal-instance/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryHybridKeySecret queries HybridKeySecret list
func (cli *ZSClient) QueryHybridKeySecret(params *param.QueryParam) ([]view.HybridAccountInventoryView, error) {
	var resp []view.HybridAccountInventoryView
	return resp, cli.List("v1/hybrid/hybrid/key", params, &resp)
}

// UploadFileToVm operates on UploadFileToVm
func (cli *ZSClient) UploadFileToVm(params param.UploadFileToVmParam) (*view.UploadFileToVmEventView, error) {
	resp := view.UploadFileToVmEventView{}
	if err := cli.Post("v1/upload-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcVpnGatewayFromLocal queries VpcVpnGatewayFromLocal list
func (cli *ZSClient) QueryVpcVpnGatewayFromLocal(params *param.QueryParam) ([]view.VpcVpnGatewayInventoryView, error) {
	var resp []view.VpcVpnGatewayInventoryView
	return resp, cli.List("v1/hybrid/vpc-vpn", params, &resp)
}

// ChangeL3NetworkDhcpIpAddress changes L3NetworkDhcpIpAddress
func (cli *ZSClient) ChangeL3NetworkDhcpIpAddress(uuid string, params param.ChangeL3NetworkDhcpIpAddressParam) (*view.ChangeL3NetworkDhcpIpAddressEventView, error) {
	resp := view.ChangeL3NetworkDhcpIpAddressEventView{}
	if err := cli.Put("v1/l3-networks/{l3NetworkUuid}/dhcp-ip", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckVolumeSnapshotGroupAvailability operates on VolumeSnapshotGroupAvailability
func (cli *ZSClient) CheckVolumeSnapshotGroupAvailability(params param.CheckVolumeSnapshotGroupAvailabilityParam) (*view.CheckVolumeSnapshotGroupAvailabilityView, error) {
	var resp view.CheckVolumeSnapshotGroupAvailabilityView
	if err := cli.Get("v1/volume-snapshots/groups/availabilities", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SsoClientPushData operates on SsoClientPushData
func (cli *ZSClient) SsoClientPushData(uuid string, params param.SsoClientPushDataParam) (*view.SsoClientPushDataEventView, error) {
	resp := view.SsoClientPushDataEventView{}
	if err := cli.Put("v1/sso/resource/data/push", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddEmailAddressToSNSEmailEndpoint adds EmailAddressToSNSEmailEndpoint
func (cli *ZSClient) AddEmailAddressToSNSEmailEndpoint(params param.AddEmailAddressToSNSEmailEndpointParam) (*view.SNSEmailAddressInventoryView, error) {
	var resp view.AddEmailAddressToSNSEmailEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/emails/email-addresses", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryVmNicInSecurityGroup queries VmNicInSecurityGroup list
func (cli *ZSClient) QueryVmNicInSecurityGroup(params *param.QueryParam) ([]view.VmNicSecurityGroupRefInventoryView, error) {
	var resp []view.VmNicSecurityGroupRefInventoryView
	return resp, cli.List("v1/security-groups/vm-instances/nics", params, &resp)
}

// BackupDatabaseToPublicCloud operates on DatabaseToPublicCloud
func (cli *ZSClient) BackupDatabaseToPublicCloud(params param.BackupDatabaseToPublicCloudParam) (*view.BackupDatabaseToPublicCloudEventView, error) {
	resp := view.BackupDatabaseToPublicCloudEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoveryImageFromImageStoreBackupStorage operates on yImageFromImageStoreBackupStorage
func (cli *ZSClient) RecoveryImageFromImageStoreBackupStorage(uuid string, params param.RecoveryImageFromImageStoreBackupStorageParam) (*view.ImageInventoryView, error) {
	var resp view.RecoveryImageFromImageStoreBackupStorageEventView
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryEventFromResourceStack queries EventFromResourceStack list
func (cli *ZSClient) QueryEventFromResourceStack(params *param.QueryParam) ([]view.CloudFormationStackEventInventoryView, error) {
	var resp []view.CloudFormationStackEventInventoryView
	return resp, cli.List("v1/cloudformation/event", params, &resp)
}

// LogInByUser operates on LogInByUser
func (cli *ZSClient) LogInByUser(uuid string, params param.LogInByUserParam) (*view.SessionInventoryView, error) {
	var resp view.LogInView
	if err := cli.Put("v1/accounts/users/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RevertVmFromSnapshotGroup operates on VmFromSnapshotGroup
func (cli *ZSClient) RevertVmFromSnapshotGroup(uuid string, params param.RevertVmFromSnapshotGroupParam) (*view.RevertVmFromSnapshotGroupEventView, error) {
	resp := view.RevertVmFromSnapshotGroupEventView{}
	if err := cli.Put("v1/volume-snapshots/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachFirewallRuleSetFromL3 operates on FirewallRuleSetFromL3
func (cli *ZSClient) DetachFirewallRuleSetFromL3(params param.DetachFirewallRuleSetFromL3Param) (*view.DetachFirewallRuleSetFromL3EventView, error) {
	resp := view.DetachFirewallRuleSetFromL3EventView{}
	if err := cli.Post("v1/vpcfirewalls/l3networks/{l3Uuid}/ruleSets/{ruleSetUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLoginCaptcha gets LoginCaptcha by uuid
func (cli *ZSClient) GetLoginCaptcha(uuid string) (*view.GetLoginCaptchaView, error) {
	var resp view.GetLoginCaptchaView
	if err := cli.Get("v1/login/control/captcha", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVmSchedulingRulesFromExecuteState operates on ListVmSchedulingRulesFromExecuteState
func (cli *ZSClient) ListVmSchedulingRulesFromExecuteState(params param.ListVmSchedulingRulesFromExecuteStateParam) (*view.ListVmSchedulingRulesFromExecuteStateView, error) {
	resp := view.ListVmSchedulingRulesFromExecuteStateView{}
	if err := cli.Post("v1/list/vmSchedulingRules/from/conflict/state", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmUserDefinedXml operates on VmUserDefinedXml
func (cli *ZSClient) SetVmUserDefinedXml(uuid string, params param.SetVmUserDefinedXmlParam) (*view.SetVmUserDefinedXmlEventView, error) {
	resp := view.SetVmUserDefinedXmlEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageQga operates on ImageQga
func (cli *ZSClient) SetImageQga(uuid string, params param.SetImageQgaParam) (*view.SetImageQgaEventView, error) {
	resp := view.SetImageQgaEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVMsFromKVMHost operates on ListVMsFromKVMHost
func (cli *ZSClient) ListVMsFromKVMHost(params param.ListVMsFromKVMHostParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post("v1/v2v", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// TakeVmConsoleScreenshot operates on TakeVmConsoleScreenshot
func (cli *ZSClient) TakeVmConsoleScreenshot(uuid string, params param.TakeVmConsoleScreenshotParam) (*view.TakeVmConsoleScreenshotEventView, error) {
	resp := view.TakeVmConsoleScreenshotEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVRouterNetworksFromOspfArea removes VRouterNetworksFromOspfArea
func (cli *ZSClient) RemoveVRouterNetworksFromOspfArea(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/routerArea/networks", uuid, string(deleteMode))
}

// GetAliyunNasMountTargetRemote gets AliyunNasMountTargetRemote by uuid
func (cli *ZSClient) GetAliyunNasMountTargetRemote(uuid string) (*view.AliyunNasMountTargetPropertyView, error) {
	var resp view.AliyunNasMountTargetPropertyView
	if err := cli.Get("v1/nas/aliyun/mount/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateImageGroupFromVmInstance creates ImageGroupFromVmInstance
func (cli *ZSClient) CreateImageGroupFromVmInstance(params param.CreateImageGroupFromVmInstanceParam) (*view.ImageGroupInventoryView, error) {
	var resp view.CreateImageGroupFromVmInstanceEventView
	if err := cli.Post("v1/images/groups/from/vm-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// TerminateVirtualBorderRouterRemote operates on TerminateVirtualBorderRouterRemote
func (cli *ZSClient) TerminateVirtualBorderRouterRemote(uuid string, params param.TerminateVirtualBorderRouterRemoteParam) (*view.TerminateVirtualBorderRouterRemoteEventView, error) {
	resp := view.TerminateVirtualBorderRouterRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/border-router/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmBackup deletes VmBackup
func (cli *ZSClient) DeleteVmBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-backups/{groupUuid}", uuid, string(deleteMode))
}

// SetVmSecurityLevel operates on VmSecurityLevel
func (cli *ZSClient) SetVmSecurityLevel(uuid string, params param.SetVmSecurityLevelParam) (*view.SetVmSecurityLevelEventView, error) {
	resp := view.SetVmSecurityLevelEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveMdevDeviceSpecFromVmInstance removes MdevDeviceSpecFromVmInstance
func (cli *ZSClient) RemoveMdevDeviceSpecFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/mdev-device-specs/{mdevSpecUuid}/vm-instances/{vmInstanceUuid}", uuid, string(deleteMode))
}

// SyncVolumeSize operates on VolumeSize
func (cli *ZSClient) SyncVolumeSize(uuid string, params param.SyncVolumeSizeParam) (*view.VolumeInventoryView, error) {
	var resp view.SyncVolumeSizeEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetTrashOnBackupStorage gets TrashOnBackupStorage by uuid
func (cli *ZSClient) GetTrashOnBackupStorage(uuid string) (*view.InstallPathRecycleInventoryView, error) {
	var resp view.InstallPathRecycleInventoryView
	if err := cli.Get("v1/backup-storage/trash", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeDiskOfferingState changes DiskOfferingState
func (cli *ZSClient) ChangeDiskOfferingState(uuid string, params param.ChangeDiskOfferingStateParam) (*view.DiskOfferingInventoryView, error) {
	var resp view.ChangeDiskOfferingStateEventView
	if err := cli.Put("v1/disk-offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RequestConsoleAccess operates on RequestConsoleAccess
func (cli *ZSClient) RequestConsoleAccess(params param.RequestConsoleAccessParam) (*view.ConsoleInventoryView, error) {
	var resp view.RequestConsoleAccessEventView
	if err := cli.Post("v1/consoles", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeIAM2VirtualIDGroupState changes IAM2VirtualIDGroupState
func (cli *ZSClient) ChangeIAM2VirtualIDGroupState(uuid string, params param.ChangeIAM2VirtualIDGroupStateParam) (*view.IAM2VirtualIDGroupInventoryView, error) {
	var resp view.ChangeIAM2VirtualIDGroupStateEventView
	if err := cli.Put("v1/iam2/projects/groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateEventData updates EventData
func (cli *ZSClient) UpdateEventData(uuid string, params param.UpdateEventDataParam) (*view.UpdateEventDataEventView, error) {
	resp := view.UpdateEventDataEventView{}
	if err := cli.Put("v1/zwatch/events/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncHybridEipFromRemote operates on HybridEipFromRemote
func (cli *ZSClient) SyncHybridEipFromRemote(uuid string, params param.SyncHybridEipFromRemoteParam) (*view.HybridEipAddressInventoryView, error) {
	resp := view.HybridEipAddressInventoryView{}
	if err := cli.Put("v1/hybrid/eip/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunRouteEntryRemote deletes AliyunRouteEntryRemote
func (cli *ZSClient) DeleteAliyunRouteEntryRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/route-entry/{uuid}", uuid, string(deleteMode))
}

// UngenerateSriovPciDevices operates on UngenerateSriovPciDevices
func (cli *ZSClient) UngenerateSriovPciDevices(uuid string, params param.UngenerateSriovPciDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmStaticIp deletes VmStaticIp
func (cli *ZSClient) DeleteVmStaticIp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/static-ips", uuid, string(deleteMode))
}

// AttachMonitorTriggerActionToTrigger operates on MonitorTriggerActionToTrigger
func (cli *ZSClient) AttachMonitorTriggerActionToTrigger(params param.AttachMonitorTriggerActionToTriggerParam) (*view.AttachMonitorTriggerActionToTriggerEventView, error) {
	resp := view.AttachMonitorTriggerActionToTriggerEventView{}
	if err := cli.Post("v1/monitoring/triggers/{triggerUuid}/trigger-actions/{actionUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOrganizationQuota updates OrganizationQuota
func (cli *ZSClient) UpdateOrganizationQuota(uuid string, params param.UpdateOrganizationQuotaParam) (*view.QuotaInventoryView, error) {
	var resp view.UpdateOrganizationQuotaEventView
	if err := cli.Put("v1/iam2/Organization/quotas/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetAliyunNasFileSystemRemote gets AliyunNasFileSystemRemote by uuid
func (cli *ZSClient) GetAliyunNasFileSystemRemote(uuid string) (*view.AliyunNasFileSystemPropertyView, error) {
	var resp view.AliyunNasFileSystemPropertyView
	if err := cli.Get("v1/nas/aliyun/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePreconfigurationTemplateState changes PreconfigurationTemplateState
func (cli *ZSClient) ChangePreconfigurationTemplateState(uuid string, params param.ChangePreconfigurationTemplateStateParam) (*view.PreconfigurationTemplateInventoryView, error) {
	var resp view.ChangePreconfigurationTemplateStateEventView
	if err := cli.Put("v1/baremetal/preconfigurations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetOrganizationSupervisor operates on OrganizationSupervisor
func (cli *ZSClient) SetOrganizationSupervisor(uuid string, params param.SetOrganizationSupervisorParam) (*view.SetOrganizationSupervisorEventView, error) {
	resp := view.SetOrganizationSupervisorEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL3NetworksToIPsecConnection operates on L3NetworksToIPsecConnection
func (cli *ZSClient) AttachL3NetworksToIPsecConnection(params param.AttachL3NetworksToIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	var resp view.AttachL3NetworksToIPsecConnectionEventView
	if err := cli.Post("v1/ipsec/{uuid}/l3networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ExecuteGuestVmScript operates on ExecuteGuestVmScript
func (cli *ZSClient) ExecuteGuestVmScript(uuid string, params param.ExecuteGuestVmScriptParam) (*view.GuestVmScriptExecutedRecordInventoryView, error) {
	var resp view.ExecuteGuestVmScriptEventView
	if err := cli.Put("v1/scripts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddNfsPrimaryStorage adds NfsPrimaryStorage
func (cli *ZSClient) AddNfsPrimaryStorage(params param.AddNfsPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/nfs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetIAM2ProjectContainerClusterCandidates gets IAM2ProjectContainerClusterCandidates by uuid
func (cli *ZSClient) GetIAM2ProjectContainerClusterCandidates(uuid string) (*view.ContainerClusterInventoryView, error) {
	var resp view.ContainerClusterInventoryView
	if err := cli.Get("v1/iam2/projects/container/cluster/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachTagToResources operates on TagToResources
func (cli *ZSClient) AttachTagToResources(params param.AttachTagToResourcesParam) (*view.AttachTagToResourcesEventView, error) {
	resp := view.AttachTagToResourcesEventView{}
	if err := cli.Post("v1/tags/{tagUuid}/resources", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePrimaryStorageState changes PrimaryStorageState
func (cli *ZSClient) ChangePrimaryStorageState(uuid string, params param.ChangePrimaryStorageStateParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.ChangePrimaryStorageStateEventView
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVpcAttachedNetflow gets VpcAttachedNetflow by uuid
func (cli *ZSClient) GetVpcAttachedNetflow(uuid string) (*view.FlowMeterInventoryView, error) {
	var resp view.FlowMeterInventoryView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-netflow", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAuditData gets AuditData by uuid
func (cli *ZSClient) GetAuditData(uuid string) (*view.GetAuditDataView, error) {
	var resp view.GetAuditDataView
	if err := cli.Get("v1/zwatch/audits", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSpiceCertificates gets SpiceCertificates by uuid
func (cli *ZSClient) GetSpiceCertificates(uuid string) (*view.GetSpiceCertificatesView, error) {
	var resp view.GetSpiceCertificatesView
	if err := cli.Get("v1/spice/certificates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveUserFromGroup removes UserFromGroup
func (cli *ZSClient) RemoveUserFromGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/groups/{groupUuid}/users/{userUuid}", uuid, string(deleteMode))
}

// DeleteEcsVpcRemote deletes EcsVpcRemote
func (cli *ZSClient) DeleteEcsVpcRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vpc/remote/{uuid}", uuid, string(deleteMode))
}

// SyncDatabaseBackupFromImageStoreBackupStorage operates on DatabaseBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncDatabaseBackupFromImageStoreBackupStorage(uuid string, params param.SyncDatabaseBackupFromImageStoreBackupStorageParam) (*view.DatabaseBackupInventoryView, error) {
	var resp view.SyncDatabaseBackupFromImageStoreBackupStorageEventView
	if err := cli.Put("v1/database-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteFirewallIpSetTemplate deletes FirewallIpSetTemplate
func (cli *ZSClient) DeleteFirewallIpSetTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/ipset/templates/{uuid}", uuid, string(deleteMode))
}

// SNSDingTalkTestConnection operates on DingTalkTestConnection
func (cli *ZSClient) SNSDingTalkTestConnection(params param.SNSDingTalkTestConnectionParam) (*view.SNSDingTalkTestConnectionEventView, error) {
	resp := view.SNSDingTalkTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/ding-talk/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportImageFromBackupStorage operates on ImageFromBackupStorage
func (cli *ZSClient) ExportImageFromBackupStorage(uuid string, params param.ExportImageFromBackupStorageParam) (*view.ExportImageFromBackupStorageEventView, error) {
	resp := view.ExportImageFromBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetModelCenterServices gets ModelCenterServices by uuid
func (cli *ZSClient) GetModelCenterServices(uuid string) (*view.GetModelCenterServicesView, error) {
	var resp view.GetModelCenterServicesView
	if err := cli.Get("v1/ai/model-centers/services", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallIpSetTemplate creates FirewallIpSetTemplate
func (cli *ZSClient) CreateFirewallIpSetTemplate(params param.CreateFirewallIpSetTemplateParam) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	var resp view.CreateFirewallIpSetTemplateEventView
	if err := cli.Post("v1/vpcfirewalls/ipset/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachMonitorTriggerActionFromTrigger operates on MonitorTriggerActionFromTrigger
func (cli *ZSClient) DetachMonitorTriggerActionFromTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/triggers/{triggerUuid}/trigger-actions/{actionUuid}", uuid, string(deleteMode))
}

// DetachPolicyRouteRuleSetFromL3 operates on PolicyRouteRuleSetFromL3
func (cli *ZSClient) DetachPolicyRouteRuleSetFromL3(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}", uuid, string(deleteMode))
}

// CreateL2TfNetwork creates L2TfNetwork
func (cli *ZSClient) CreateL2TfNetwork(params param.CreateL2TfNetworkParam) (*view.L2NetworkInventoryView, error) {
	var resp view.CreateL2NetworkEventView
	if err := cli.Post("v1/l2-networks/tf", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetInterdependentL3NetworksImages gets InterdependentL3NetworksImages by uuid
func (cli *ZSClient) GetInterdependentL3NetworksImages(uuid string) (*view.GetInterdependentL3NetworkImageView, error) {
	var resp view.GetInterdependentL3NetworkImageView
	if err := cli.Get("v1/images-l3networks/dependencies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateVolumeSnapshotChain operates on VolumeSnapshotChain
func (cli *ZSClient) ValidateVolumeSnapshotChain(uuid string, params param.ValidateVolumeSnapshotChainParam) (*view.ValidateVolumeSnapshotChainEventView, error) {
	resp := view.ValidateVolumeSnapshotChainEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeHostNetworkInterfaceLldpMode changes HostNetworkInterfaceLldpMode
func (cli *ZSClient) ChangeHostNetworkInterfaceLldpMode(uuid string, params param.ChangeHostNetworkInterfaceLldpModeParam) (*view.HostNetworkInterfaceLldpInventoryView, error) {
	resp := view.HostNetworkInterfaceLldpInventoryView{}
	if err := cli.Put("v1/hostNetworkInterface/lldp/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGuestOsMetadata gets GuestOsMetadata by uuid
func (cli *ZSClient) GetGuestOsMetadata(uuid string) (*view.GuestOsCharacterInventoryView, error) {
	var resp view.GuestOsCharacterInventoryView
	if err := cli.Get("v1/guest-os/metadata", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicsForLoadBalancerServerGroup gets CandidateVmNicsForLoadBalancerServerGroup by uuid
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancerServerGroup(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.Get("v1/load-balancers/servergroups/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachIscsiServerToCluster operates on IscsiServerToCluster
func (cli *ZSClient) AttachIscsiServerToCluster(params param.AttachIscsiServerToClusterParam) (*view.IscsiServerInventoryView, error) {
	var resp view.AttachIscsiServerToClusterEventView
	if err := cli.Post("v1/clusters/{clusterUuid}/storage-devices/iscsi/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachRoleToAccount operates on RoleToAccount
func (cli *ZSClient) AttachRoleToAccount(params param.AttachRoleToAccountParam) (*view.AttachRoleToAccountEventView, error) {
	resp := view.AttachRoleToAccountEventView{}
	if err := cli.Post("v1/identities/accounts/{accountUuid}/roles/{roleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryBuildApp queries BuildApp list
func (cli *ZSClient) QueryBuildApp(params *param.QueryParam) ([]view.BuildApplicationInventoryView, error) {
	var resp []view.BuildApplicationInventoryView
	return resp, cli.List("v1/appcenter/buildapp", params, &resp)
}

// AttachIsoToVmInstance operates on IsoToVmInstance
func (cli *ZSClient) AttachIsoToVmInstance(params param.AttachIsoToVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.AttachIsoToVmInstanceEventView
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/iso/{isoUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// LoginByCas operates on ByCas
func (cli *ZSClient) LoginByCas(uuid string, params param.LoginByCasParam) (*view.SessionInventoryView, error) {
	var resp view.LoginByCasView
	if err := cli.Put("v1/cas/login/", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetVRouterRouterId operates on VRouterRouterId
func (cli *ZSClient) SetVRouterRouterId(params param.SetVRouterRouterIdParam) (*view.SetVRouterRouterIdEventView, error) {
	resp := view.SetVRouterRouterIdEventView{}
	if err := cli.Post("v1/routerArea/{vRouterUuid}/routerid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExpungeVmUserDefinedXmlHookScript operates on VmUserDefinedXmlHookScript
func (cli *ZSClient) ExpungeVmUserDefinedXmlHookScript(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/xml-hook-script/{uuid}", uuid, string(deleteMode))
}

// DeleteCdpTaskData deletes CdpTaskData
func (cli *ZSClient) DeleteCdpTaskData(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cdp-task/{uuid}/data", uuid, string(deleteMode))
}

// CheckApiPermission operates on ApiPermission
func (cli *ZSClient) CheckApiPermission(uuid string, params param.CheckApiPermissionParam) (*view.CheckApiPermissionView, error) {
	resp := view.CheckApiPermissionView{}
	if err := cli.Put("v1/accounts/permissions/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTextTemplateArg gets TextTemplateArg by uuid
func (cli *ZSClient) GetTextTemplateArg(uuid string) (*view.GetTextTemplateArgView, error) {
	var resp view.GetTextTemplateArgView
	if err := cli.Get("v1/zwatch/textTemplateArg", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFirewall deletes Firewall
func (cli *ZSClient) DeleteFirewall(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/{uuid}", uuid, string(deleteMode))
}

// GetVmCapabilities gets VmCapabilities by uuid
func (cli *ZSClient) GetVmCapabilities(uuid string) (*view.GetVmCapabilitiesView, error) {
	var resp view.GetVmCapabilitiesView
	if err := cli.Get("v1/vm-instances/{uuid}/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAccessKeyState changes AccessKeyState
func (cli *ZSClient) ChangeAccessKeyState(uuid string, params param.ChangeAccessKeyStateParam) (*view.AccessKeyInventoryView, error) {
	var resp view.ChangeAccessKeyStateEventView
	if err := cli.Put("v1/accesskeys/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeployDistributedModelService operates on DeployDistributedModelService
func (cli *ZSClient) DeployDistributedModelService(uuid string, params param.DeployDistributedModelServiceParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.DeployDistributedModelServiceEventView
	if err := cli.Put("v1/ai/model-services", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetIAM2SystemAttributes gets IAM2SystemAttributes by uuid
func (cli *ZSClient) GetIAM2SystemAttributes(uuid string) (*view.IAM2AttributeInventoryView, error) {
	var resp view.IAM2AttributeInventoryView
	if err := cli.Get("v1/iam2/attributes/system", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeInstanceOfferingState changes InstanceOfferingState
func (cli *ZSClient) ChangeInstanceOfferingState(uuid string, params param.ChangeInstanceOfferingStateParam) (*view.InstanceOfferingInventoryView, error) {
	var resp view.ChangeInstanceOfferingStateEventView
	if err := cli.Put("v1/instance-offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryVirtualBorderRouterFromLocal queries VirtualBorderRouterFromLocal list
func (cli *ZSClient) QueryVirtualBorderRouterFromLocal(params *param.QueryParam) ([]view.VirtualBorderRouterInventoryView, error) {
	var resp []view.VirtualBorderRouterInventoryView
	return resp, cli.List("v1/hybrid/aliyun/border-router", params, &resp)
}

// GetBackupStorageCapacity gets BackupStorageCapacity by uuid
func (cli *ZSClient) GetBackupStorageCapacity(uuid string) (*view.GetBackupStorageCapacityView, error) {
	var resp view.GetBackupStorageCapacityView
	if err := cli.Get("v1/backup-storage/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateSeMdevDevices operates on SeMdevDevices
func (cli *ZSClient) GenerateSeMdevDevices(uuid string, params param.GenerateSeMdevDevicesParam) (*view.GenerateSeMdevDevicesEventView, error) {
	resp := view.GenerateSeMdevDevicesEventView{}
	if err := cli.Put("v1/mtty-devices/{mttyDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateMiniCluster creates MiniCluster
func (cli *ZSClient) CreateMiniCluster(params param.CreateMiniClusterParam) (*view.ClusterInventoryView, error) {
	var resp view.CreateMiniClusterEventView
	if err := cli.Post("v1/mini-clusters", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SyncImageFromImageStoreBackupStorage operates on ImageFromImageStoreBackupStorage
func (cli *ZSClient) SyncImageFromImageStoreBackupStorage(uuid string, params param.SyncImageFromImageStoreBackupStorageParam) (*view.ImageInventoryView, error) {
	var resp view.SyncImageFromImageStoreBackupStorageEventView
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeVipState changes VipState
func (cli *ZSClient) ChangeVipState(uuid string, params param.ChangeVipStateParam) (*view.VipInventoryView, error) {
	var resp view.ChangeVipStateEventView
	if err := cli.Put("v1/vips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UndoSnapshotCreation operates on UndoSnapshotCreation
func (cli *ZSClient) UndoSnapshotCreation(uuid string, params param.UndoSnapshotCreationParam) (*view.VolumeInventoryView, error) {
	var resp view.UndoSnapshotCreationEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddBuildApp adds BuildApp
func (cli *ZSClient) AddBuildApp(params param.AddBuildAppParam) (*view.BuildApplicationInventoryView, error) {
	var resp view.AddBuildAppEventView
	if err := cli.Post("v1/appcenter/buildapp/add", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateVmFromVolumeBackup creates VmFromVolumeBackup
func (cli *ZSClient) CreateVmFromVolumeBackup(params param.CreateVmFromVolumeBackupParam) (*view.VmInstanceInventoryView, error) {
	var resp view.CreateVmFromVolumeBackupEventView
	if err := cli.Post("v1/vm-instances/from/vm-backup/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetIdentityZoneFromRemote gets IdentityZoneFromRemote by uuid
func (cli *ZSClient) GetIdentityZoneFromRemote(uuid string) (*view.IdentityZonePropertyView, error) {
	var resp view.IdentityZonePropertyView
	if err := cli.Get("v1/hybrid/identity-zone/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEcsInstanceVncUrl gets EcsInstanceVncUrl by uuid
func (cli *ZSClient) GetEcsInstanceVncUrl(uuid string) (*view.GetEcsInstanceVncUrlView, error) {
	var resp view.GetEcsInstanceVncUrlView
	if err := cli.Get("v1/hybrid/aliyun/ecs-vnc/{uuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddMonToCephPrimaryStorage adds MonToCephPrimaryStorage
func (cli *ZSClient) AddMonToCephPrimaryStorage(params param.AddMonToCephPrimaryStorageParam) (*view.CephPrimaryStorageInventoryView, error) {
	var resp view.AddMonToCephPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/ceph/{uuid}/mons", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryLocalRaidPhysicalDrive queries LocalRaidPhysicalDrive list
func (cli *ZSClient) QueryLocalRaidPhysicalDrive(params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, error) {
	var resp []view.RaidPhysicalDriveInventoryView
	return resp, cli.List("v1/storage-devices/local-raid/physical-drives", params, &resp)
}

// RemoveHostRouteFromL3Network removes HostRouteFromL3Network
func (cli *ZSClient) RemoveHostRouteFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/{l3NetworkUuid}/hostroute", uuid, string(deleteMode))
}

// BackupStorageMigrateImage operates on StorageMigrateImage
func (cli *ZSClient) BackupStorageMigrateImage(uuid string, params param.BackupStorageMigrateImageParam) (*view.ImageInventoryView, error) {
	var resp view.BackupStorageMigrateImageEventView
	if err := cli.Put("v1/backup-storage/images/{imageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryArchiveTicketHistory queries ArchiveTicketHistory list
func (cli *ZSClient) QueryArchiveTicketHistory(params *param.QueryParam) ([]view.ArchiveTicketStatusHistoryInventoryView, error) {
	var resp []view.ArchiveTicketStatusHistoryInventoryView
	return resp, cli.List("v1/tickets/histories/archives", params, &resp)
}

// ChangeIAM2VirtualIDType changes IAM2VirtualIDType
func (cli *ZSClient) ChangeIAM2VirtualIDType(uuid string, params param.ChangeIAM2VirtualIDTypeParam) (*view.IAM2VirtualIDInventoryView, error) {
	var resp view.ChangeIAM2VirtualIDTypeEventView
	if err := cli.Put("v1/iam2/virtual-ids/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveBackendServerFromServerGroup removes BackendServerFromServerGroup
func (cli *ZSClient) RemoveBackendServerFromServerGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/servergroups/{serverGroupUuid}/backendservers/actions", uuid, string(deleteMode))
}

// GetVpcAttachedVip gets VpcAttachedVip by uuid
func (cli *ZSClient) GetVpcAttachedVip(uuid string) (*view.VipInventoryView, error) {
	var resp view.VipInventoryView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-vip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIpv6Range adds Ipv6Range
func (cli *ZSClient) AddIpv6Range(params param.AddIpv6RangeParam) (*view.IpRangeInventoryView, error) {
	var resp view.AddIpRangeEventView
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/ipv6-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CheckBaremetalChassisConfigFile operates on BaremetalChassisConfigFile
func (cli *ZSClient) CheckBaremetalChassisConfigFile(params param.CheckBaremetalChassisConfigFileParam) (*view.CheckBaremetalChassisConfigFileView, error) {
	resp := view.CheckBaremetalChassisConfigFileView{}
	if err := cli.Post("v1/baremetal/chassis/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOssBucketFileRemote deletes OssBucketFileRemote
func (cli *ZSClient) DeleteOssBucketFileRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/oss-bucket-file/remote/{uuid}", uuid, string(deleteMode))
}

// ChangeMulticastRouterState changes MulticastRouterState
func (cli *ZSClient) ChangeMulticastRouterState(uuid string, params param.ChangeMulticastRouterStateParam) (*view.MulticastRouterInventoryView, error) {
	var resp view.ChangeMulticastRouterStateEventView
	if err := cli.Put("v1/multicast/virtual-routers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetMaaSUsage gets MaaSUsage by uuid
func (cli *ZSClient) GetMaaSUsage(uuid string) (*view.GetMaaSUsageView, error) {
	var resp view.GetMaaSUsageView
	if err := cli.Get("v1/maas/usage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFreeIp gets FreeIp by uuid
func (cli *ZSClient) GetFreeIp(uuid string) (*view.FreeIpInventoryView, error) {
	var resp view.FreeIpInventoryView
	if err := cli.Get("v1/l3-networks/ip/free", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOssBucketRemote deletes OssBucketRemote
func (cli *ZSClient) DeleteOssBucketRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/oss-bucket/remote/{uuid}", uuid, string(deleteMode))
}

// LogInByLdap operates on LogInByLdap
func (cli *ZSClient) LogInByLdap(uuid string, params param.LogInByLdapParam) (*view.SessionInventoryView, error) {
	var resp view.LogInByLdapView
	if err := cli.Put("v1/ldap/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateL2PortGroup creates L2PortGroup
func (cli *ZSClient) CreateL2PortGroup(params param.CreateL2PortGroupParam) (*view.CreateL2PortGroupEventView, error) {
	resp := view.CreateL2PortGroupEventView{}
	if err := cli.Post("v1/l2-networks/port-group", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateInstanceOfferingUserConfig operates on InstanceOfferingUserConfig
func (cli *ZSClient) ValidateInstanceOfferingUserConfig(uuid string, params param.ValidateInstanceOfferingUserConfigParam) (*view.ValidateInstanceOfferingUserConfigEventView, error) {
	resp := view.ValidateInstanceOfferingUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryTag queries Tag list
func (cli *ZSClient) QueryTag(params *param.QueryParam) ([]view.TagPatternInventoryView, error) {
	var resp []view.TagPatternInventoryView
	return resp, cli.List("v1/tags", params, &resp)
}

// SetVmHostname operates on VmHostname
func (cli *ZSClient) SetVmHostname(uuid string, params param.SetVmHostnameParam) (*view.SetVmHostnameEventView, error) {
	resp := view.SetVmHostnameEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// TriggerGCJob operates on TriggerGCJob
func (cli *ZSClient) TriggerGCJob(uuid string, params param.TriggerGCJobParam) (*view.TriggerGCJobEventView, error) {
	resp := view.TriggerGCJobEventView{}
	if err := cli.Put("v1/gc-jobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckBareMetal2IpmiChassisConfigFile operates on BareMetal2IpmiChassisConfigFile
func (cli *ZSClient) CheckBareMetal2IpmiChassisConfigFile(params param.CheckBareMetal2IpmiChassisConfigFileParam) (*view.CheckBareMetal2ChassisConfigFileView, error) {
	resp := view.CheckBareMetal2ChassisConfigFileView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVirtualRouterLocal deletes VirtualRouterLocal
func (cli *ZSClient) DeleteVirtualRouterLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vrouter/{uuid}", uuid, string(deleteMode))
}

// DeleteVpcIkeConfigLocal deletes VpcIkeConfigLocal
func (cli *ZSClient) DeleteVpcIkeConfigLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/ike/{uuid}", uuid, string(deleteMode))
}

// CreateOssBucketRemote creates OssBucketRemote
func (cli *ZSClient) CreateOssBucketRemote(params param.CreateOssBucketRemoteParam) (*view.OssBucketInventoryView, error) {
	var resp view.CreateOssBucketRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/oss-bucket/remote", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddSimulatorPrimaryStorage adds SimulatorPrimaryStorage
func (cli *ZSClient) AddSimulatorPrimaryStorage(params param.AddSimulatorPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachVRouterRouteTableFromVRouter operates on VRouterRouteTableFromVRouter
func (cli *ZSClient) DetachVRouterRouteTableFromVRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vrouter-route-tables/{routeTableUuid}/detach/{virtualRouterVmUuid}", uuid, string(deleteMode))
}

// GetVipUsedPorts gets VipUsedPorts by uuid
func (cli *ZSClient) GetVipUsedPorts(uuid string) (*view.VipPortRangeInventoryView, error) {
	var resp view.VipPortRangeInventoryView
	if err := cli.Get("v1/vips/{uuid}/usedports", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmConsolePassword operates on VmConsolePassword
func (cli *ZSClient) SetVmConsolePassword(uuid string, params param.SetVmConsolePasswordParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmConsolePasswordEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachFirewallRuleSetToL3 operates on FirewallRuleSetToL3
func (cli *ZSClient) AttachFirewallRuleSetToL3(params param.AttachFirewallRuleSetToL3Param) (*view.VpcFirewallRuleSetL3RefInventoryView, error) {
	var resp view.AttachFirewallRuleSetToL3EventView
	if err := cli.Post("v1/vpcfirewalls/ruleSets/{ruleSetUuid}/l3networks/{l3Uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CleanUpStorageTrashOnPrimaryStorage operates on UpStorageTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpStorageTrashOnPrimaryStorage(uuid string, params param.CleanUpStorageTrashOnPrimaryStorageParam) (*view.CleanUpStorageTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpStorageTrashOnPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/storagetrash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVRouterFlowMeterNetwork queries VRouterFlowMeterNetwork list
func (cli *ZSClient) QueryVRouterFlowMeterNetwork(params *param.QueryParam) ([]view.NetworkRouterFlowMeterRefInventoryView, error) {
	var resp []view.NetworkRouterFlowMeterRefInventoryView
	return resp, cli.List("v1/flowmeters/networks", params, &resp)
}

// GetManagementNodeDirCapacity gets ManagementNodeDirCapacity by uuid
func (cli *ZSClient) GetManagementNodeDirCapacity(uuid string) (*view.GetManagementNodeDirCapacityView, error) {
	var resp view.GetManagementNodeDirCapacityView
	if err := cli.Get("v1/zwatch/mn", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGpuDeviceSpecCandidates gets GpuDeviceSpecCandidates by uuid
func (cli *ZSClient) GetGpuDeviceSpecCandidates(uuid string) (*view.GpuDeviceSpecInventoryView, error) {
	var resp view.GpuDeviceSpecInventoryView
	if err := cli.Get("v1/gpu-device-specs/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallRuleSetL3Ref queries FirewallRuleSetL3Ref list
func (cli *ZSClient) QueryFirewallRuleSetL3Ref(params *param.QueryParam) ([]view.VpcFirewallRuleSetL3RefInventoryView, error) {
	var resp []view.VpcFirewallRuleSetL3RefInventoryView
	return resp, cli.List("v1/vpcfirewalls/l3networks/rulesets/refs", params, &resp)
}

// UngroupVolumeSnapshotGroup operates on UngroupVolumeSnapshotGroup
func (cli *ZSClient) UngroupVolumeSnapshotGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots/ungroup/{uuid}", uuid, string(deleteMode))
}

// SubscribeSNSTopic operates on SubscribeSNSTopic
func (cli *ZSClient) SubscribeSNSTopic(params param.SubscribeSNSTopicParam) (*view.SubscribeSNSTopicEventView, error) {
	resp := view.SubscribeSNSTopicEventView{}
	if err := cli.Post("v1/sns/topics/{topicUuid}/endpoints/{endpointUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicForSecurityGroup gets CandidateVmNicForSecurityGroup by uuid
func (cli *ZSClient) GetCandidateVmNicForSecurityGroup(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.Get("v1/security-groups/{securityGroupUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmRDP gets VmRDP by uuid
func (cli *ZSClient) GetVmRDP(uuid string) (*view.GetVmRDPView, error) {
	var resp view.GetVmRDPView
	if err := cli.Get("v1/vm-instances/{uuid}/rdp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPciDeviceToVm operates on PciDeviceToVm
func (cli *ZSClient) AttachPciDeviceToVm(params param.AttachPciDeviceToVmParam) (*view.PciDeviceInventoryView, error) {
	var resp view.AttachPciDeviceToVmEventView
	if err := cli.Post("v1/pci-device/pci-devices/{pciDeviceUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CleanupBillingUsage operates on upBillingUsage
func (cli *ZSClient) CleanupBillingUsage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/usage", uuid, string(deleteMode))
}

// GetLdapEntry gets LdapEntry by uuid
func (cli *ZSClient) GetLdapEntry(uuid string) (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.Get("v1/ldap/entry", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL2NetworksForAttachingCluster gets CandidateL2NetworksForAttachingCluster by uuid
func (cli *ZSClient) GetCandidateL2NetworksForAttachingCluster(uuid string) (*view.L2NetworkDataView, error) {
	var resp view.L2NetworkDataView
	if err := cli.Get("v1/cluster/{clusterUuid}/l2-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsVfNicAvailableInL3Network operates on IsVfNicAvailableInL3Network
func (cli *ZSClient) IsVfNicAvailableInL3Network(params param.IsVfNicAvailableInL3NetworkParam) (*view.IsVfNicAvailableInL3NetworkView, error) {
	var resp view.IsVfNicAvailableInL3NetworkView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/hosts/{hostUuid}/vfnicavailable", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAllMetricMetadata gets AllMetricMetadata by uuid
func (cli *ZSClient) GetAllMetricMetadata(uuid string) (*view.GetAllMetricMetadataView, error) {
	var resp view.GetAllMetricMetadataView
	if err := cli.Get("v1/zwatch/metrics/meta-data", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddOssBucketFromRemote adds OssBucketFromRemote
func (cli *ZSClient) AddOssBucketFromRemote(params param.AddOssBucketFromRemoteParam) (*view.OssBucketInventoryView, error) {
	var resp view.AddOssBucketFromRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/oss-bucket", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SyncVmBackup operates on VmBackup
func (cli *ZSClient) SyncVmBackup(uuid string, params param.SyncVmBackupParam) (*view.SyncVmBackupEventView, error) {
	resp := view.SyncVmBackupEventView{}
	if err := cli.Put("v1/vm-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshGuestOsMetadata operates on GuestOsMetadata
func (cli *ZSClient) RefreshGuestOsMetadata(uuid string, params param.RefreshGuestOsMetadataParam) (*view.RefreshGuestOsMetadataEventView, error) {
	resp := view.RefreshGuestOsMetadataEventView{}
	if err := cli.Put("v1/guest-os/metadata/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GCAliyunSnapshotRemote operates on GCAliyunSnapshotRemote
func (cli *ZSClient) GCAliyunSnapshotRemote(params param.GCAliyunSnapshotRemoteParam) (*view.GCAliyunSnapshotRemoteEventView, error) {
	resp := view.GCAliyunSnapshotRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/snapshot/{dataCenterUuid}/gc", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadBackupFileFromPublicCloud operates on DownloadBackupFileFromPublicCloud
func (cli *ZSClient) DownloadBackupFileFromPublicCloud(params param.DownloadBackupFileFromPublicCloudParam) (*view.DownloadBackupFileFromPublicCloudEventView, error) {
	resp := view.DownloadBackupFileFromPublicCloudEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql/download", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIAM2VirtualIDsToProjects adds IAM2VirtualIDsToProjects
func (cli *ZSClient) AddIAM2VirtualIDsToProjects(params param.AddIAM2VirtualIDsToProjectsParam) (*view.AddIAM2VirtualIDsToProjectsEventView, error) {
	resp := view.AddIAM2VirtualIDsToProjectsEventView{}
	if err := cli.Post("v1/iam2/projects/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2ProjectTemplateFromProject creates IAM2ProjectTemplateFromProject
func (cli *ZSClient) CreateIAM2ProjectTemplateFromProject(params param.CreateIAM2ProjectTemplateFromProjectParam) (*view.IAM2ProjectTemplateInventoryView, error) {
	var resp view.CreateIAM2ProjectTemplateFromProjectEventView
	if err := cli.Post("v1/iam2/projects/templates/from/projects/{projectUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateTag creates Tag
func (cli *ZSClient) CreateTag(params param.CreateTagParam) (*view.TagPatternInventoryView, error) {
	var resp view.CreateTagEventView
	if err := cli.Post("v1/tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateConsolePassword updates ConsolePassword
func (cli *ZSClient) UpdateConsolePassword(uuid string, params param.UpdateConsolePasswordParam) (*view.VmInstanceInventoryView, error) {
	var resp view.UpdateConsolePasswordEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateVmInstanceFromVolumeSnapshotGroup creates VmInstanceFromVolumeSnapshotGroup
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshotGroup(params param.CreateVmInstanceFromVolumeSnapshotGroupParam) (*view.VmInstanceInventoryView, error) {
	var resp view.CreateVmInstanceFromVolumeSnapshotGroupEventView
	if err := cli.Post("v1/vm-instances/from/volume-snapshots/group/{volumeSnapshotGroupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetIAM2ProjectRetirePolicy operates on IAM2ProjectRetirePolicy
func (cli *ZSClient) SetIAM2ProjectRetirePolicy(uuid string, params param.SetIAM2ProjectRetirePolicyParam) (*view.SetIAM2ProjectRetirePolicyEventView, error) {
	resp := view.SetIAM2ProjectRetirePolicyEventView{}
	if err := cli.Put("v1/iam2/projects/retire-policies/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunIAM2Script operates on RunIAM2Script
func (cli *ZSClient) RunIAM2Script(params param.RunIAM2ScriptParam) (*view.LongJobInventoryView, error) {
	var resp view.RunIAM2ScriptEventView
	if err := cli.Post("v1/iam2/iam2-script/run", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachServiceToObservabilityServer operates on ServiceToObservabilityServer
func (cli *ZSClient) AttachServiceToObservabilityServer(params param.AttachServiceToObservabilityServerParam) (*view.ObservabilityServerVmInventoryView, error) {
	var resp view.AttachServiceToObservabilityServerEventView
	if err := cli.Post("v1/observability-server/{observabilityServerUuid}/service", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteHostNetworkServiceType deletes HostNetworkServiceType
func (cli *ZSClient) DeleteHostNetworkServiceType(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts/service-types/{uuid}", uuid, string(deleteMode))
}

// CreateIAM2ProjectFromTemplate creates IAM2ProjectFromTemplate
func (cli *ZSClient) CreateIAM2ProjectFromTemplate(params param.CreateIAM2ProjectFromTemplateParam) (*view.IAM2ProjectInventoryView, error) {
	var resp view.CreateIAM2ProjectFromTemplateEventView
	if err := cli.Post("v1/iam2/projects/from/templates/{templateUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryVpcIpSecConfigFromLocal queries VpcIpSecConfigFromLocal list
func (cli *ZSClient) QueryVpcIpSecConfigFromLocal(params *param.QueryParam) ([]view.VpcVpnIpSecConfigInventoryView, error) {
	var resp []view.VpcVpnIpSecConfigInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection/ipsec", params, &resp)
}

// AddConnectionAccessPointFromRemote adds ConnectionAccessPointFromRemote
func (cli *ZSClient) AddConnectionAccessPointFromRemote(params param.AddConnectionAccessPointFromRemoteParam) (*view.ConnectionAccessPointInventoryView, error) {
	var resp view.AddConnectionAccessPointFromRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/access-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachSshKeyPairToVmInstance operates on SshKeyPairToVmInstance
func (cli *ZSClient) AttachSshKeyPairToVmInstance(params param.AttachSshKeyPairToVmInstanceParam) (*view.SshKeyPairInventoryView, error) {
	var resp view.AttachSshKeyPairToVmInstanceEventView
	if err := cli.Post("v1/ssh-key-pair/{sshKeyPairUuid}/vm-instance/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryEmailTriggerAction queries EmailTrigger list
func (cli *ZSClient) QueryEmailTriggerAction(params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, error) {
	var resp []view.MonitorTriggerActionInventoryView
	return resp, cli.List("v1/monitoring/trigger-actions/emails", params, &resp)
}

// DetachBareMetal2GatewayFromCluster operates on BareMetal2GatewayFromCluster
func (cli *ZSClient) DetachBareMetal2GatewayFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/clusters/{clusterUuid}/gateways/{gatewayUuid}", uuid, string(deleteMode))
}

// ReloadElaboration operates on ReloadElaboration
func (cli *ZSClient) ReloadElaboration(uuid string, params param.ReloadElaborationParam) (*view.ReloadElaborationEventView, error) {
	resp := view.ReloadElaborationEventView{}
	if err := cli.Put("v1/errorcode/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReconnectVirtualRouter operates on VirtualRouter
func (cli *ZSClient) ReconnectVirtualRouter(uuid string, params param.ReconnectVirtualRouterParam) (*view.ApplianceVmInventoryView, error) {
	var resp view.ReconnectVirtualRouterEventView
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ConvertVmFromForeignHypervisor operates on ConvertVmFromForeignHypervisor
func (cli *ZSClient) ConvertVmFromForeignHypervisor(params param.ConvertVmFromForeignHypervisorParam) (*view.LongJobInventoryView, error) {
	var resp view.ConvertVmFromForeignHypervisorEventView
	if err := cli.Post("v1/v2vs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch deletes ConnectionBetweenL3NetWorkAndAliyunVSwitch
func (cli *ZSClient) DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/connections/{uuid}", uuid, string(deleteMode))
}

// RestartResourceStack operates on RestartResourceStack
func (cli *ZSClient) RestartResourceStack(uuid string, params param.RestartResourceStackParam) (*view.ResourceStackInventoryView, error) {
	var resp view.RestartResourceStackEventView
	if err := cli.Put("v1/cloudformation/stack/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryHygonDevice queries HygonDevice list
func (cli *ZSClient) QueryHygonDevice(params *param.QueryParam) ([]view.HygonCcpDeviceInventoryView, error) {
	var resp []view.HygonCcpDeviceInventoryView
	return resp, cli.List("v1/hygon-devices", params, &resp)
}

// SyncEcsImageFromRemote operates on EcsImageFromRemote
func (cli *ZSClient) SyncEcsImageFromRemote(params param.SyncEcsImageFromRemoteParam) (*view.EcsImageInventoryView, error) {
	resp := view.EcsImageInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/image/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPoliciesToUser operates on PoliciesToUser
func (cli *ZSClient) AttachPoliciesToUser(params param.AttachPoliciesToUserParam) (*view.AttachPoliciesToUserEventView, error) {
	resp := view.AttachPoliciesToUserEventView{}
	if err := cli.Post("v1/accounts/users/{userUuid}/policy-collection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachBackupStorageToZone operates on BackupStorageToZone
func (cli *ZSClient) AttachBackupStorageToZone(params param.AttachBackupStorageToZoneParam) (*view.BackupStorageInventoryView, error) {
	var resp view.AttachBackupStorageToZoneEventView
	if err := cli.Post("v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddPciDeviceSpecToVmInstance adds PciDeviceSpecToVmInstance
func (cli *ZSClient) AddPciDeviceSpecToVmInstance(params param.AddPciDeviceSpecToVmInstanceParam) (*view.VmInstancePciDeviceSpecRefInventoryView, error) {
	var resp view.AddPciDeviceSpecToVmInstanceEventView
	if err := cli.Post("v1/pci-device-specs/{pciSpecUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ResizeRootVolume operates on RootVolume
func (cli *ZSClient) ResizeRootVolume(uuid string, params param.ResizeRootVolumeParam) (*view.VolumeInventoryView, error) {
	var resp view.ResizeRootVolumeEventView
	if err := cli.Put("v1/volumes/resize/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVpcVpnConfigurationFromRemote gets VpcVpnConfigurationFromRemote by uuid
func (cli *ZSClient) GetVpcVpnConfigurationFromRemote(uuid string) (*view.GetVpcVpnConfigurationFromRemoteView, error) {
	var resp view.GetVpcVpnConfigurationFromRemoteView
	if err := cli.Get("v1/hybrid/vpn-conf/{uuid}/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateImageGroupFromImage creates ImageGroupFromImage
func (cli *ZSClient) CreateImageGroupFromImage(params param.CreateImageGroupFromImageParam) (*view.ImageGroupInventoryView, error) {
	var resp view.CreateImageGroupFromImageEventView
	if err := cli.Post("v1/imagegroup/from/image/{rootVolumeTemplateUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// TokenIntrospection operates on TokenIntrospection
func (cli *ZSClient) TokenIntrospection(params param.TokenIntrospectionParam) (*view.TokenIntrospectionView, error) {
	resp := view.TokenIntrospectionView{}
	if err := cli.Post("v1/token/introspect", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVmBackupFromImageStoreBackupStorage operates on VmBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncVmBackupFromImageStoreBackupStorage(uuid string, params param.SyncVmBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.Put("v1/vm-backups/{groupUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddCertificateToLoadBalancerListener adds CertificateToLoadBalancerListener
func (cli *ZSClient) AddCertificateToLoadBalancerListener(params param.AddCertificateToLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	var resp view.AddCertificateToLoadBalancerListenerEventView
	if err := cli.Post("v1/load-balancers/listeners/{listenerUuid}/certificate", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddRolesToIAM2VirtualID adds RolesToIAM2VirtualID
func (cli *ZSClient) AddRolesToIAM2VirtualID(params param.AddRolesToIAM2VirtualIDParam) (*view.AddRolesToIAM2VirtualIDEventView, error) {
	resp := view.AddRolesToIAM2VirtualIDEventView{}
	if err := cli.Post("v1/iam2/projects/virtual-ids/{virtualIDUuid}/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryTicketHistory queries TicketHistory list
func (cli *ZSClient) QueryTicketHistory(params *param.QueryParam) ([]view.TicketStatusHistoryInventoryView, error) {
	var resp []view.TicketStatusHistoryInventoryView
	return resp, cli.List("v1/tickets/histories", params, &resp)
}

// CreateFaultToleranceVmInstance creates FaultToleranceVmInstance
func (cli *ZSClient) CreateFaultToleranceVmInstance(params param.CreateFaultToleranceVmInstanceParam) (*view.CreateFaultToleranceVmInstanceEventView, error) {
	resp := view.CreateFaultToleranceVmInstanceEventView{}
	if err := cli.Post("v1/vm-instances/fault-tolerance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteResourceStackVmPortMonitor deletes ResourceStackVmPortMonitor
func (cli *ZSClient) DeleteResourceStackVmPortMonitor(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/stack/monitor/delvm", uuid, string(deleteMode))
}

// DeleteGCJob deletes GCJob
func (cli *ZSClient) DeleteGCJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/gc-jobs/{uuid}", uuid, string(deleteMode))
}

// DeleteEmailAddressOfSNSEmailEndpoint deletes EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) DeleteEmailAddressOfSNSEmailEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/emails/{endpointUuid}/email-addresses/{emailAddressUuid}", uuid, string(deleteMode))
}

// CleanInvalidLdapIAM2Binding operates on InvalidLdapIAM2Binding
func (cli *ZSClient) CleanInvalidLdapIAM2Binding(uuid string, params param.CleanInvalidLdapIAM2BindingParam) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.Put("v1/iam2/ldap/bindings/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHybridEip updates HybridEip
func (cli *ZSClient) UpdateHybridEip(uuid string, params param.UpdateHybridEipParam) (*view.HybridEipAddressInventoryView, error) {
	var resp view.UpdateHybridEipEventView
	if err := cli.Put("v1/hybrid/eip/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVpcAttachedIpsec gets VpcAttachedIpsec by uuid
func (cli *ZSClient) GetVpcAttachedIpsec(uuid string) (*view.IPsecConnectionInventoryView, error) {
	var resp view.IPsecConnectionInventoryView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-ipsec", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImagesFromImageStoreBackupStorage gets ImagesFromImageStoreBackupStorage by uuid
func (cli *ZSClient) GetImagesFromImageStoreBackupStorage(uuid string) (*view.GetImagesFromImageStoreBackupStorageView, error) {
	var resp view.GetImagesFromImageStoreBackupStorageView
	if err := cli.Get("v1/backup-storage/{uuid}/image-store", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetElaborationCategories gets ElaborationCategories by uuid
func (cli *ZSClient) GetElaborationCategories(uuid string) (*view.GetElaborationCategoriesView, error) {
	var resp view.GetElaborationCategoriesView
	if err := cli.Get("v1/errorcode/elaborations/categories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetScsiLunCandidatesForAttachingVm gets ScsiLunCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetScsiLunCandidatesForAttachingVm(uuid string) (*view.ScsiLunInventoryView, error) {
	var resp view.ScsiLunInventoryView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/candidate-storage-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostMultipathTopology gets HostMultipathTopology by uuid
func (cli *ZSClient) GetHostMultipathTopology(uuid string) (*view.GetHostMultipathTopologyView, error) {
	var resp view.GetHostMultipathTopologyView
	if err := cli.Get("v1/storage-devices/multipath/topology", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallIpSetTemplate queries FirewallIpSetTemplate list
func (cli *ZSClient) QueryFirewallIpSetTemplate(params *param.QueryParam) ([]view.VpcFirewallIpSetTemplateInventoryView, error) {
	var resp []view.VpcFirewallIpSetTemplateInventoryView
	return resp, cli.List("v1/vpcfirewalls/ipset/templates", params, &resp)
}

// DeleteEcsImageRemote deletes EcsImageRemote
func (cli *ZSClient) DeleteEcsImageRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/image/remote/{uuid}", uuid, string(deleteMode))
}

// GetHostNetworkFacts gets HostNetworkFacts by uuid
func (cli *ZSClient) GetHostNetworkFacts(uuid string) (*view.GetHostNetworkFactsView, error) {
	var resp view.GetHostNetworkFactsView
	if err := cli.Get("v1/hosts/network-facts/{hostUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpTrashOnBackupStorage operates on UpTrashOnBackupStorage
func (cli *ZSClient) CleanUpTrashOnBackupStorage(uuid string, params param.CleanUpTrashOnBackupStorageParam) (*view.CleanUpTrashOnBackupStorageEventView, error) {
	resp := view.CleanUpTrashOnBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/trash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateConnectionBetweenL3NetworkAndAliyunVSwitch creates ConnectionBetweenL3NetworkAndAliyunVSwitch
func (cli *ZSClient) CreateConnectionBetweenL3NetworkAndAliyunVSwitch(params param.CreateConnectionBetweenL3NetworkAndAliyunVSwitchParam) (*view.ConnectionRelationShipInventoryView, error) {
	var resp view.CreateConnectionBetweenL3NetworkAndAliyunVSwitchEventView
	if err := cli.Post("v1/hybrid/aliyun/connections", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachPriceTableFromAccount operates on PriceTableFromAccount
func (cli *ZSClient) DetachPriceTableFromAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}", uuid, string(deleteMode))
}

// AddVRouterNetworksToFlowMeter adds VRouterNetworksToFlowMeter
func (cli *ZSClient) AddVRouterNetworksToFlowMeter(params param.AddVRouterNetworksToFlowMeterParam) (*view.NetworkRouterFlowMeterRefInventoryView, error) {
	resp := view.NetworkRouterFlowMeterRefInventoryView{}
	if err := cli.Post("v1/flowmeters/{flowMeterUuid}/router/{vRouterUuid}/addnetworks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2VirtualIDInGroup gets IAM2VirtualIDInGroup by uuid
func (cli *ZSClient) GetIAM2VirtualIDInGroup(uuid string) (*view.IAM2VirtualIDInventoryView, error) {
	var resp view.IAM2VirtualIDInventoryView
	if err := cli.Get("v1/iam2/IAM2VirtualIDGroup/IAM2VirtualID", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnlockIdentity operates on UnlockIdentity
func (cli *ZSClient) UnlockIdentity(params param.UnlockIdentityParam) (*view.UnlockIdentityView, error) {
	var resp view.UnlockIdentityView
	if err := cli.Get("v1/login/control/unlock", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmSchedulingRuleState changes VmSchedulingRuleState
func (cli *ZSClient) ChangeVmSchedulingRuleState(uuid string, params param.ChangeVmSchedulingRuleStateParam) (*view.VmSchedulingRuleInventoryView, error) {
	var resp view.ChangeVmSchedulingRuleStateEventView
	if err := cli.Put("v1/vmSchedulingRule/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetCandidateVmNicsForPortMirror gets CandidateVmNicsForPortMirror by uuid
func (cli *ZSClient) GetCandidateVmNicsForPortMirror(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.Get("v1/port-mirrors/{portMirrorUuid}/vm-instances/candidate-nics/{type}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRule creates FirewallRule
func (cli *ZSClient) CreateFirewallRule(params param.CreateFirewallRuleParam) (*view.VpcFirewallRuleInventoryView, error) {
	var resp view.CreateFirewallRuleEventView
	if err := cli.Post("v1/vpcfirewalls/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetIAM2ProjectLoginExpired operates on IAM2ProjectLoginExpired
func (cli *ZSClient) SetIAM2ProjectLoginExpired(uuid string, params param.SetIAM2ProjectLoginExpiredParam) (*view.SetIAM2ProjectLoginExpiredEventView, error) {
	resp := view.SetIAM2ProjectLoginExpiredEventView{}
	if err := cli.Put("v1/iam2/projects/add/login-expired/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmEmulatorPinning gets VmEmulatorPinning by uuid
func (cli *ZSClient) GetVmEmulatorPinning(uuid string) (*view.GetVmEmulatorPinningView, error) {
	var resp view.GetVmEmulatorPinningView
	if err := cli.Get("v1/vm-instances/{uuid}/emulator-pinning", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDataVolumeAttachableVm gets DataVolumeAttachableVm by uuid
func (cli *ZSClient) GetDataVolumeAttachableVm(uuid string) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.Get("v1/volumes/{volumeUuid}/candidate-vm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIpRangeByNetworkCidr adds IpRangeByNetworkCidr
func (cli *ZSClient) AddIpRangeByNetworkCidr(params param.AddIpRangeByNetworkCidrParam) (*view.IpRangeInventoryView, error) {
	var resp view.AddIpRangeByNetworkCidrEventView
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/ip-ranges/by-cidr", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateL2NoVlanNetwork creates L2NoVlanNetwork
func (cli *ZSClient) CreateL2NoVlanNetwork(params param.CreateL2NoVlanNetworkParam) (*view.L2NetworkInventoryView, error) {
	var resp view.CreateL2NetworkEventView
	if err := cli.Post("v1/l2-networks/no-vlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddMonToCephBackupStorage adds MonToCephBackupStorage
func (cli *ZSClient) AddMonToCephBackupStorage(params param.AddMonToCephBackupStorageParam) (*view.CephBackupStorageInventoryView, error) {
	var resp view.AddMonToCephBackupStorageEventView
	if err := cli.Post("v1/backup-storage/ceph/{uuid}/mons", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachBareMetal2ProvisionNetworkFromCluster operates on BareMetal2ProvisionNetworkFromCluster
func (cli *ZSClient) DetachBareMetal2ProvisionNetworkFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/clusters/{clusterUuid}/provision-networks/{networkUuid}", uuid, string(deleteMode))
}

// DeleteAliyunDiskFromLocal deletes AliyunDiskFromLocal
func (cli *ZSClient) DeleteAliyunDiskFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/disk/{uuid}", uuid, string(deleteMode))
}

// GetResourceNames gets ResourceNames by uuid
func (cli *ZSClient) GetResourceNames(uuid string) (*view.ResourceInventoryView, error) {
	var resp view.ResourceInventoryView
	if err := cli.Get("v1/resources/names", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2VirtualIDAPIPermission gets IAM2VirtualIDAPIPermission by uuid
func (cli *ZSClient) GetIAM2VirtualIDAPIPermission(uuid string) (*view.GetIAM2VirtualIDAPIPermissionView, error) {
	var resp view.GetIAM2VirtualIDAPIPermissionView
	if err := cli.Get("v1/iam2/virtual-ids/api-permissions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOrganizationQuotaUsage gets OrganizationQuotaUsage by uuid
func (cli *ZSClient) GetOrganizationQuotaUsage(uuid string) (*view.GetOrganizationQuotaUsageView, error) {
	var resp view.GetOrganizationQuotaUsageView
	if err := cli.Get("v1/iam2/organizations/quota/{uuid}/usages", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceConfigs gets ResourceConfigs by uuid
func (cli *ZSClient) GetResourceConfigs(uuid string) (*view.GetResourceConfigsView, error) {
	var resp view.GetResourceConfigsView
	if err := cli.Get("v1/resource-configurations/{resourceUuid}/{category}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVpcUserVpnGatewayFromRemote operates on VpcUserVpnGatewayFromRemote
func (cli *ZSClient) SyncVpcUserVpnGatewayFromRemote(uuid string, params param.SyncVpcUserVpnGatewayFromRemoteParam) (*view.VpcUserVpnGatewayInventoryView, error) {
	resp := view.VpcUserVpnGatewayInventoryView{}
	if err := cli.Put("v1/hybrid/user-vpn/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPrimaryStorageFromCluster operates on PrimaryStorageFromCluster
func (cli *ZSClient) DetachPrimaryStorageFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}", uuid, string(deleteMode))
}

// CheckStackTemplateParameters operates on StackTemplateParameters
func (cli *ZSClient) CheckStackTemplateParameters(params param.CheckStackTemplateParametersParam) (*view.CheckStackTemplateParametersView, error) {
	resp := view.CheckStackTemplateParametersView{}
	if err := cli.Post("v1/cloudformation/stack/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFactoryModeState gets FactoryModeState by uuid
func (cli *ZSClient) GetFactoryModeState(uuid string) (*view.GetFactoryModeStateView, error) {
	var resp view.GetFactoryModeStateView
	if err := cli.Get("v1/management-nodes/factory-mode-state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddServerGroupToLoadBalancerListener adds ServerGroupToLoadBalancerListener
func (cli *ZSClient) AddServerGroupToLoadBalancerListener(params param.AddServerGroupToLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	var resp view.AddServerGroupToLoadBalancerListenerEventView
	if err := cli.Post("v1/load-balancers/listeners/{listenerUuid}/servergroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetActiveAlarmStatus gets ActiveAlarmStatus by uuid
func (cli *ZSClient) GetActiveAlarmStatus(uuid string) (*view.GetActiveAlarmStatusView, error) {
	var resp view.GetActiveAlarmStatusView
	if err := cli.Get("v1/zwatch/activealarms/status", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployModelEvalService operates on DeployModelEvalService
func (cli *ZSClient) DeployModelEvalService(uuid string, params param.DeployModelEvalServiceParam) (*view.ModelEvalServiceInstanceGroupInventoryView, error) {
	var resp view.DeployModelEvalServiceEventView
	if err := cli.Put("v1/ai/model-services/eval/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachNvmeServerToCluster operates on NvmeServerToCluster
func (cli *ZSClient) AttachNvmeServerToCluster(params param.AttachNvmeServerToClusterParam) (*view.NvmeServerInventoryView, error) {
	var resp view.AttachNvmeServerToClusterEventView
	if err := cli.Post("v1/clusters/{clusterUuid}/storage-devices/nvme/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetHostResourceAllocation gets HostResourceAllocation by uuid
func (cli *ZSClient) GetHostResourceAllocation(uuid string) (*view.GetHostResourceAllocationEventView, error) {
	var resp view.GetHostResourceAllocationEventView
	if err := cli.Get("v1/hosts/{uuid}/resource-allocation", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachUsbDeviceToVm operates on UsbDeviceToVm
func (cli *ZSClient) AttachUsbDeviceToVm(params param.AttachUsbDeviceToVmParam) (*view.UsbDeviceInventoryView, error) {
	var resp view.AttachUsbDeviceToVmEventView
	if err := cli.Post("v1/usb-device/usb-devices/{usbDeviceUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetLicenseAddOns gets LicenseAddOns by uuid
func (cli *ZSClient) GetLicenseAddOns(uuid string) (*view.GetLicenseAddOnsView, error) {
	var resp view.GetLicenseAddOnsView
	if err := cli.Get("v1/licenses/addons", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAliyunSnapshotRemote operates on AliyunSnapshotRemote
func (cli *ZSClient) SyncAliyunSnapshotRemote(params param.SyncAliyunSnapshotRemoteParam) (*view.AliyunSnapshotInventoryView, error) {
	resp := view.AliyunSnapshotInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/snapshot/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunRouterInterfaceFromLocal queries AliyunRouterInterfaceFromLocal list
func (cli *ZSClient) QueryAliyunRouterInterfaceFromLocal(params *param.QueryParam) ([]view.AliyunRouterInterfaceInventoryView, error) {
	var resp []view.AliyunRouterInterfaceInventoryView
	return resp, cli.List("v1/hybrid/aliyun/router-interface", params, &resp)
}

// UpdateTicketRequest updates TicketRequest
func (cli *ZSClient) UpdateTicketRequest(uuid string, params param.UpdateTicketRequestParam) (*view.TicketInventoryView, error) {
	var resp view.UpdateTicketRequestEventView
	if err := cli.Put("v1/tickets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVpcIPsecLog gets VpcIPsecLog by uuid
func (cli *ZSClient) GetVpcIPsecLog(uuid string) (*view.GetVpcIPsecLogView, error) {
	var resp view.GetVpcIPsecLogView
	if err := cli.Get("v1/vpc/virtual-routers/ipseclog", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromOvf creates VmInstanceFromOvf
func (cli *ZSClient) CreateVmInstanceFromOvf(params param.CreateVmInstanceFromOvfParam) (*view.VmInstanceInventoryView, error) {
	var resp view.CreateVmInstanceFromOvfEventView
	if err := cli.Post("v1/ovf/create-vm-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryLdapBinding queries LdapBinding list
func (cli *ZSClient) QueryLdapBinding(params *param.QueryParam) ([]view.LdapAccountRefInventoryView, error) {
	var resp []view.LdapAccountRefInventoryView
	return resp, cli.List("v1/ldap/bindings", params, &resp)
}

// ChangeVmImage changes VmImage
func (cli *ZSClient) ChangeVmImage(uuid string, params param.ChangeVmImageParam) (*view.VmInstanceInventoryView, error) {
	var resp view.ChangeVmImageEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddResourcesToDirectory adds ResourcesToDirectory
func (cli *ZSClient) AddResourcesToDirectory(params param.AddResourcesToDirectoryParam) (*view.AddResourcesToDirectoryEventView, error) {
	resp := view.AddResourcesToDirectoryEventView{}
	if err := cli.Post("v1/add/resources/directory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachGuestToolsIsoToVm operates on GuestToolsIsoToVm
func (cli *ZSClient) AttachGuestToolsIsoToVm(uuid string, params param.AttachGuestToolsIsoToVmParam) (*view.AttachGuestToolsIsoToVmEventView, error) {
	resp := view.AttachGuestToolsIsoToVmEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachAliyunKey operates on AliyunKey
func (cli *ZSClient) DetachAliyunKey(uuid string, params param.DetachAliyunKeyParam) (*view.DetachAliyunKeyEventView, error) {
	resp := view.DetachAliyunKeyEventView{}
	if err := cli.Put("v1/hybrid/aliyun/key/{uuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBuildApp deletes BuildApp
func (cli *ZSClient) DeleteBuildApp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/buildapp/{uuid}", uuid, string(deleteMode))
}

// ChangeBareMetal2InstancePassword changes BareMetal2InstancePassword
func (cli *ZSClient) ChangeBareMetal2InstancePassword(uuid string, params param.ChangeBareMetal2InstancePasswordParam) (*view.BareMetal2InstanceInventoryView, error) {
	var resp view.ChangeBareMetal2InstancePasswordEventView
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/action", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetResourceFromPublishApp gets ResourceFromPublishApp by uuid
func (cli *ZSClient) GetResourceFromPublishApp(uuid string) (*view.GetResourceFromPublishAppView, error) {
	var resp view.GetResourceFromPublishAppView
	if err := cli.Get("v1/appcenter/app/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeResourceOwner changes ResourceOwner
func (cli *ZSClient) ChangeResourceOwner(uuid string, params param.ChangeResourceOwnerParam) (*view.AccountResourceRefInventoryView, error) {
	var resp view.ChangeResourceOwnerEventView
	if err := cli.Put("v1/account/{accountUuid}/resources", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetHostIommuState gets HostIommuState by uuid
func (cli *ZSClient) GetHostIommuState(uuid string) (*view.GetHostIommuStateView, error) {
	var resp view.GetHostIommuStateView
	if err := cli.Get("v1/pci-device/hosts/{uuid}/state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVirtualBorderRouterLocal deletes VirtualBorderRouterLocal
func (cli *ZSClient) DeleteVirtualBorderRouterLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/border-router/{uuid}", uuid, string(deleteMode))
}

// GetMetricData gets MetricData by uuid
func (cli *ZSClient) GetMetricData(uuid string) (*view.GetMetricDataView, error) {
	var resp view.GetMetricDataView
	if err := cli.Get("v1/zwatch/metrics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunSnapshotFromLocal queries AliyunSnapshotFromLocal list
func (cli *ZSClient) QueryAliyunSnapshotFromLocal(params *param.QueryParam) ([]view.AliyunSnapshotInventoryView, error) {
	var resp []view.AliyunSnapshotInventoryView
	return resp, cli.List("v1/hybrid/aliyun/snapshot", params, &resp)
}

// EnableCbtTask operates on EnableCbtTask
func (cli *ZSClient) EnableCbtTask(params param.EnableCbtTaskParam) (*view.EnableCbtTaskEventView, error) {
	resp := view.EnableCbtTaskEventView{}
	if err := cli.Post("v1/cbt-task/enable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAliyunNasAccessGroupRemote gets AliyunNasAccessGroupRemote by uuid
func (cli *ZSClient) GetAliyunNasAccessGroupRemote(uuid string) (*view.AliyunNasAccessGroupPropertyView, error) {
	var resp view.AliyunNasAccessGroupPropertyView
	if err := cli.Get("v1/nas/aliyun/access/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckBuildAppParameters operates on BuildAppParameters
func (cli *ZSClient) CheckBuildAppParameters(params param.CheckBuildAppParametersParam) (*view.CheckBuildAppParametersView, error) {
	resp := view.CheckBuildAppParametersView{}
	if err := cli.Post("v1/appcenter/buildapp/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddLabelToEventSubscription adds LabelToEventSubscription
func (cli *ZSClient) AddLabelToEventSubscription(params param.AddLabelToEventSubscriptionParam) (*view.EventSubscriptionLabelInventoryView, error) {
	var resp view.AddLabelToEventSubscriptionEventView
	if err := cli.Post("v1/zwatch/events/subscriptions/{subscriptionUuid}/labels", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVpcVRouterDistributedRoutingConnections gets VpcVRouterDistributedRoutingConnections by uuid
func (cli *ZSClient) GetVpcVRouterDistributedRoutingConnections(uuid string) (*view.StringView, error) {
	var resp view.StringView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/tracked-connections", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateThirdpartyAlerts updates ThirdpartyAlerts
func (cli *ZSClient) UpdateThirdpartyAlerts(uuid string, params param.UpdateThirdpartyAlertsParam) (*view.UpdateThirdpartyAlertsEventView, error) {
	resp := view.UpdateThirdpartyAlertsEventView{}
	if err := cli.Put("v1/zwatch/third-party/alerts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PullSdnControllerTenant operates on PullSdnControllerTenant
func (cli *ZSClient) PullSdnControllerTenant(uuid string, params param.PullSdnControllerTenantParam) (*view.H3cSdnControllerTenantInventoryView, error) {
	resp := view.H3cSdnControllerTenantInventoryView{}
	if err := cli.Put("v1/sdn-controllers/{uuid}/tenant/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachServiceFromObservabilityServer operates on ServiceFromObservabilityServer
func (cli *ZSClient) DetachServiceFromObservabilityServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/observability-server/{observabilityServerUuid}/service", uuid, string(deleteMode))
}

// GenerateHygonMdevDevices operates on HygonMdevDevices
func (cli *ZSClient) GenerateHygonMdevDevices(uuid string, params param.GenerateHygonMdevDevicesParam) (*view.GenerateHygonMdevDevicesEventView, error) {
	resp := view.GenerateHygonMdevDevicesEventView{}
	if err := cli.Put("v1/hygon-devices/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmUsbRedirect operates on VmUsbRedirect
func (cli *ZSClient) SetVmUsbRedirect(uuid string, params param.SetVmUsbRedirectParam) (*view.SetVmUsbRedirectEventView, error) {
	resp := view.SetVmUsbRedirectEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostCandidatesForVmMigration gets HostCandidatesForVmMigration by uuid
func (cli *ZSClient) GetHostCandidatesForVmMigration(uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.Get("v1/primary-storage/hosts/{vmInstanceUuid}/migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmNicAttachableEips gets VmNicAttachableEips by uuid
func (cli *ZSClient) GetVmNicAttachableEips(uuid string) (*view.EipInventoryView, error) {
	var resp view.EipInventoryView
	if err := cli.Get("v1/vm-instances/nics/{vmNicUuid}/candidate-eips", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFactoryModeState updates FactoryModeState
func (cli *ZSClient) UpdateFactoryModeState(uuid string, params param.UpdateFactoryModeStateParam) (*view.UpdateFactoryModeStateEventView, error) {
	resp := view.UpdateFactoryModeStateEventView{}
	if err := cli.Put("v1/management-nodes/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateChronyServers updates ChronyServers
func (cli *ZSClient) UpdateChronyServers(uuid string, params param.UpdateChronyServersParam) (*view.UpdateChronyServersEventView, error) {
	resp := view.UpdateChronyServersEventView{}
	if err := cli.Put("v1/zops/chrony/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPolicyRouteRuleSetToL3 operates on PolicyRouteRuleSetToL3
func (cli *ZSClient) AttachPolicyRouteRuleSetToL3(params param.AttachPolicyRouteRuleSetToL3Param) (*view.AttachPolicyRouteRuleSetToL3EventView, error) {
	resp := view.AttachPolicyRouteRuleSetToL3EventView{}
	if err := cli.Post("v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOAuthClient updates OAuthClient
func (cli *ZSClient) UpdateOAuthClient(uuid string, params param.UpdateOAuthClientParam) (*view.OAuth2ClientInventoryView, error) {
	var resp view.UpdateOAuthClientEventView
	if err := cli.Put("v1/update/oauth2/client", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetZWatchAlertHistogram gets ZWatchAlertHistogram by uuid
func (cli *ZSClient) GetZWatchAlertHistogram(uuid string) (*view.GetZWatchAlertHistogramView, error) {
	var resp view.GetZWatchAlertHistogramView
	if err := cli.Get("v1/zwatch/alert-histories/histogram", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunRouterInterfaceRemote deletes AliyunRouterInterfaceRemote
func (cli *ZSClient) DeleteAliyunRouterInterfaceRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/router-interface/remote/{uuid}", uuid, string(deleteMode))
}

// SetImageBootMode operates on ImageBootMode
func (cli *ZSClient) SetImageBootMode(uuid string, params param.SetImageBootModeParam) (*view.SetImageBootModeEventView, error) {
	resp := view.SetImageBootModeEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachAutoScalingTemplateFromGroup operates on AutoScalingTemplateFromGroup
func (cli *ZSClient) DetachAutoScalingTemplateFromGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/template/{templateUuid}/groups/{groupUuid}", uuid, string(deleteMode))
}

// UpdateVirtualBorderRouterRemote updates VirtualBorderRouterRemote
func (cli *ZSClient) UpdateVirtualBorderRouterRemote(uuid string, params param.UpdateVirtualBorderRouterRemoteParam) (*view.VirtualBorderRouterInventoryView, error) {
	var resp view.UpdateVirtualBorderRouterRemoteEventView
	if err := cli.Put("v1/hybrid/aliyun/border-router/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVmsCapabilities gets VmsCapabilities by uuid
func (cli *ZSClient) GetVmsCapabilities(uuid string) (*view.GetVmsCapabilitiesView, error) {
	var resp view.GetVmsCapabilitiesView
	if err := cli.Get("v1/vm-instances/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPolicyToUserGroup operates on PolicyToUserGroup
func (cli *ZSClient) AttachPolicyToUserGroup(params param.AttachPolicyToUserGroupParam) (*view.AttachPolicyToUserGroupEventView, error) {
	resp := view.AttachPolicyToUserGroupEventView{}
	if err := cli.Post("v1/accounts/groups/{groupUuid}/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevokeMonitorTemplateFromMonitorGroup operates on RevokeMonitorTemplateFromMonitorGroup
func (cli *ZSClient) RevokeMonitorTemplateFromMonitorGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitortemplates/{templateUuid}/monitorgroups/{groupUuid}", uuid, string(deleteMode))
}

// DeleteFirewallRule deletes FirewallRule
func (cli *ZSClient) DeleteFirewallRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/rules/{uuid}", uuid, string(deleteMode))
}

// ShareResource operates on ShareResource
func (cli *ZSClient) ShareResource(uuid string, params param.ShareResourceParam) (*view.ShareResourceEventView, error) {
	resp := view.ShareResourceEventView{}
	if err := cli.Put("v1/accounts/resources/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsVpcRemote creates EcsVpcRemote
func (cli *ZSClient) CreateEcsVpcRemote(params param.CreateEcsVpcRemoteParam) (*view.EcsVpcInventoryView, error) {
	var resp view.CreateEcsVpcRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/vpc", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetAccountQuotaUsage gets AccountQuotaUsage by uuid
func (cli *ZSClient) GetAccountQuotaUsage(uuid string) (*view.GetAccountQuotaUsageView, error) {
	var resp view.GetAccountQuotaUsageView
	if err := cli.Get("v1/accounts/quota/{uuid}/usages", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryPluginDrivers queries PluginDrivers list
func (cli *ZSClient) QueryPluginDrivers(params *param.QueryParam) ([]view.PluginDriverInventoryView, error) {
	var resp []view.PluginDriverInventoryView
	return resp, cli.List("v1/external/plugins", params, &resp)
}

// RemoveIAM2VirtualIDsFromProjects removes IAM2VirtualIDsFromProjects
func (cli *ZSClient) RemoveIAM2VirtualIDsFromProjects(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/virtual-ids", uuid, string(deleteMode))
}

// GetCandidateL3NetworksForServerGroup gets CandidateL3NetworksForServerGroup by uuid
func (cli *ZSClient) GetCandidateL3NetworksForServerGroup(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.Get("v1/load-balancers/servergroups/candidate-l3network", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmFromCdpBackup creates VmFromCdpBackup
func (cli *ZSClient) CreateVmFromCdpBackup(params param.CreateVmFromCdpBackupParam) (*view.VmInstanceInventoryView, error) {
	var resp view.CreateVmFromCdpBackupEventView
	if err := cli.Post("v1/cdp-backups/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SyncVpcVpnConnectionFromRemote operates on VpcVpnConnectionFromRemote
func (cli *ZSClient) SyncVpcVpnConnectionFromRemote(uuid string, params param.SyncVpcVpnConnectionFromRemoteParam) (*view.VpcVpnConnectionInventoryView, error) {
	resp := view.VpcVpnConnectionInventoryView{}
	if err := cli.Put("v1/hybrid/vpn-connection/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpnIkeConfig creates VpnIkeConfig
func (cli *ZSClient) CreateVpnIkeConfig(params param.CreateVpnIkeConfigParam) (*view.VpcVpnIkeConfigInventoryView, error) {
	var resp view.CreateVpnIkeConfigEventView
	if err := cli.Post("v1/hybrid/vpn-connection/ike", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SubmitLongJob operates on SubmitLongJob
func (cli *ZSClient) SubmitLongJob(params param.SubmitLongJobParam) (*view.LongJobInventoryView, error) {
	var resp view.SubmitLongJobEventView
	if err := cli.Post("v1/longjobs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateDataVolumeTemplateFromVolumeBackup creates DataVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeBackup(params param.CreateDataVolumeTemplateFromVolumeBackupParam) (*view.ImageInventoryView, error) {
	var resp view.CreateDataVolumeTemplateFromVolumeBackupEventView
	if err := cli.Post("v1/images/data-volume-templates/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DegradeFromLicenseServer operates on DegradeFromLicenseServer
func (cli *ZSClient) DegradeFromLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server", uuid, string(deleteMode))
}

// UpdateNfvInstProvisionConfig updates NfvInstProvisionConfig
func (cli *ZSClient) UpdateNfvInstProvisionConfig(uuid string, params param.UpdateNfvInstProvisionConfigParam) (*view.ApplianceVmInventoryView, error) {
	var resp view.UpdateNfvInstProvisionConfigEventView
	if err := cli.Put("v1/vm-instances/appliances/nfvinst/{vmInstanceUuid}/provision/update", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetDebugSignal gets DebugSignal by uuid
func (cli *ZSClient) GetDebugSignal(uuid string) (*view.GetDebugSignalView, error) {
	var resp view.GetDebugSignalView
	if err := cli.Get("v1/debug", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAliyunKeySecret updates AliyunKeySecret
func (cli *ZSClient) UpdateAliyunKeySecret(uuid string, params param.UpdateAliyunKeySecretParam) (*view.HybridAccountInventoryView, error) {
	var resp view.UpdateAliyunKeySecretEventView
	if err := cli.Put("v1/hybrid/aliyun/{uuid}/key", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SyncLicenseCapacity operates on LicenseCapacity
func (cli *ZSClient) SyncLicenseCapacity(uuid string, params param.SyncLicenseCapacityParam) (*view.SyncLicenseCapacityEventView, error) {
	resp := view.SyncLicenseCapacityEventView{}
	if err := cli.Put("v1/license-server/authorized-capacity/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachDataVolumeToHost operates on DataVolumeToHost
func (cli *ZSClient) AttachDataVolumeToHost(params param.AttachDataVolumeToHostParam) (*view.AttachDataVolumeToHostEventView, error) {
	resp := view.AttachDataVolumeToHostEventView{}
	if err := cli.Post("v1/volumes/{volumeUuid}/hosts/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SecurityMachineEncrypt operates on MachineEncrypt
func (cli *ZSClient) SecurityMachineEncrypt(params param.SecurityMachineEncryptParam) (*view.SecurityMachineEncryptEventView, error) {
	resp := view.SecurityMachineEncryptEventView{}
	if err := cli.Post("v1/security-machine/encrypt/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAppBuildSystemState changes AppBuildSystemState
func (cli *ZSClient) ChangeAppBuildSystemState(uuid string, params param.ChangeAppBuildSystemStateParam) (*view.AppBuildSystemInventoryView, error) {
	var resp view.ChangeAppBuildSystemStateEventView
	if err := cli.Put("v1/appcenter/buildsystem/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetMemorySnapshotGroupReference gets MemorySnapshotGroupReference by uuid
func (cli *ZSClient) GetMemorySnapshotGroupReference(uuid string) (*view.VolumeSnapshotGroupInventoryView, error) {
	var resp view.VolumeSnapshotGroupInventoryView
	if err := cli.Get("v1/memory-snapshots/group/reference", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcVRouterNetworkServiceState gets VpcVRouterNetworkServiceState by uuid
func (cli *ZSClient) GetVpcVRouterNetworkServiceState(uuid string) (*view.GetVpcVRouterNetworkServiceStateView, error) {
	var resp view.GetVpcVRouterNetworkServiceStateView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/networkservicestate", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachNetworkServiceFromL3Network operates on NetworkServiceFromL3Network
func (cli *ZSClient) DetachNetworkServiceFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/{l3NetworkUuid}/network-services", uuid, string(deleteMode))
}

// CreateDataVolumeFromVolumeBackup creates DataVolumeFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeFromVolumeBackup(params param.CreateDataVolumeFromVolumeBackupParam) (*view.VolumeInventoryView, error) {
	var resp view.CreateDataVolumeFromVolumeBackupEventView
	if err := cli.Post("v1/volumes/data-volume/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteContainerResourceFromEndpoint deletes ContainerResourceFromEndpoint
func (cli *ZSClient) DeleteContainerResourceFromEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/container/management/endpoint/{uuid}/resources/delete", uuid, string(deleteMode))
}

// GetSupportAPIs gets SupportAPIs by uuid
func (cli *ZSClient) GetSupportAPIs(uuid string) (*view.GetSupportAPIsView, error) {
	var resp view.GetSupportAPIsView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSharedMountPointPrimaryStorage adds SharedMountPointPrimaryStorage
func (cli *ZSClient) AddSharedMountPointPrimaryStorage(params param.AddSharedMountPointPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/smp", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetTrashOnPrimaryStorage gets TrashOnPrimaryStorage by uuid
func (cli *ZSClient) GetTrashOnPrimaryStorage(uuid string) (*view.InstallPathRecycleInventoryView, error) {
	var resp view.InstallPathRecycleInventoryView
	if err := cli.Get("v1/primary-storage/trash", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicsForLoadBalancer gets CandidateVmNicsForLoadBalancer by uuid
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancer(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.Get("v1/load-balancers/listeners/{listenerUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachBaremetalPxeServerFromCluster operates on BaremetalPxeServerFromCluster
func (cli *ZSClient) DetachBaremetalPxeServerFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/pxeservers/{pxeServerUuid}", uuid, string(deleteMode))
}

// DeleteVpcUserVpnGatewayRemote deletes VpcUserVpnGatewayRemote
func (cli *ZSClient) DeleteVpcUserVpnGatewayRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/user-gateway/{uuid}/remote", uuid, string(deleteMode))
}

// ChangeAccessControlListRedirectRule changes AccessControlListRedirectRule
func (cli *ZSClient) ChangeAccessControlListRedirectRule(uuid string, params param.ChangeAccessControlListRedirectRuleParam) (*view.AccessControlListEntryInventoryView, error) {
	var resp view.ChangeAccessControlListRedirectRuleEventView
	if err := cli.Put("v1/access-control-lists/redirectRules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddResourceStackVmPortMonitor adds ResourceStackVmPortMonitor
func (cli *ZSClient) AddResourceStackVmPortMonitor(params param.AddResourceStackVmPortMonitorParam) (*view.AddResourceStackVmPortMonitorEventView, error) {
	resp := view.AddResourceStackVmPortMonitorEventView{}
	if err := cli.Post("v1/cloudformation/stack/monitor/addvm", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSNSApplicationEndpointState changes SNSApplicationEndpointState
func (cli *ZSClient) ChangeSNSApplicationEndpointState(uuid string, params param.ChangeSNSApplicationEndpointStateParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.ChangeSNSApplicationEndpointStateEventView
	if err := cli.Put("v1/sns/application-endpoints/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVpcAttachedLoadBalancer gets VpcAttachedLoadBalancer by uuid
func (cli *ZSClient) GetVpcAttachedLoadBalancer(uuid string) (*view.LoadBalancerInventoryView, error) {
	var resp view.LoadBalancerInventoryView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-lb", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpcVpnConnectionRemote creates VpcVpnConnectionRemote
func (cli *ZSClient) CreateVpcVpnConnectionRemote(params param.CreateVpcVpnConnectionRemoteParam) (*view.VpcVpnConnectionInventoryView, error) {
	var resp view.CreateVpcVpnConnectionRemoteEventView
	if err := cli.Post("v1/hybrid/vpn-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVpcAttachedPortForwardingRules gets VpcAttachedPortForwardingRules by uuid
func (cli *ZSClient) GetVpcAttachedPortForwardingRules(uuid string) (*view.PortForwardingRuleInventoryView, error) {
	var resp view.PortForwardingRuleInventoryView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-pf", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVpcVRouterNetworkServiceState operates on VpcVRouterNetworkServiceState
func (cli *ZSClient) SetVpcVRouterNetworkServiceState(params param.SetVpcVRouterNetworkServiceStateParam) (*view.SetVpcVRouterNetworkServiceStateEventView, error) {
	resp := view.SetVpcVRouterNetworkServiceStateEventView{}
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/networkservicestate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachNfvInstFromGroup operates on NfvInstFromGroup
func (cli *ZSClient) DetachNfvInstFromGroup(uuid string, params param.DetachNfvInstFromGroupParam) (*view.NfvInstGroupInventoryView, error) {
	var resp view.DetachNfvInstFromGroupEventView
	if err := cli.Put("v1/nfvinstgroup/group/{groupUuid}/instances/{nfvInstUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddDnsToVpcRouter adds DnsToVpcRouter
func (cli *ZSClient) AddDnsToVpcRouter(params param.AddDnsToVpcRouterParam) (*view.VpcRouterVmInventoryView, error) {
	var resp view.AddDnsToVpcRouterEventView
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/dns", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryAccountBilling queries AccountBilling list
func (cli *ZSClient) QueryAccountBilling(params *param.QueryParam) ([]view.BillingInventoryView, error) {
	var resp []view.BillingInventoryView
	return resp, cli.List("v1/billing/billings", params, &resp)
}

// GetVmXml gets VmXml by uuid
func (cli *ZSClient) GetVmXml(uuid string) (*view.GetVmXmlView, error) {
	var resp view.GetVmXmlView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/xml", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceFirstBootDevice gets VmInstanceFirstBootDevice by uuid
func (cli *ZSClient) GetVmInstanceFirstBootDevice(uuid string) (*view.GetVmInstanceFirstBootDeviceView, error) {
	var resp view.GetVmInstanceFirstBootDeviceView
	if err := cli.Get("v1/vm-instances/{uuid}/first-boot-device", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOvnControllerVm creates OvnControllerVm
func (cli *ZSClient) CreateOvnControllerVm(params param.CreateOvnControllerVmParam) (*view.ApplianceVmInventoryView, error) {
	var resp view.CreateOvnControllerVmEventView
	if err := cli.Post("v1/ovn/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteIpAddress deletes IpAddress
func (cli *ZSClient) DeleteIpAddress(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/{l3NetworkUuid}/ip-address", uuid, string(deleteMode))
}

// DeleteVpcVpnConnectionRemote deletes VpcVpnConnectionRemote
func (cli *ZSClient) DeleteVpcVpnConnectionRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/{uuid}/remote", uuid, string(deleteMode))
}

// AttachOssBucketToEcsDataCenter operates on OssBucketToEcsDataCenter
func (cli *ZSClient) AttachOssBucketToEcsDataCenter(uuid string, params param.AttachOssBucketToEcsDataCenterParam) (*view.OssBucketInventoryView, error) {
	var resp view.AttachOssBucketToEcsDataCenterEventView
	if err := cli.Put("v1/hybrid/aliyun/oss-bucket/{ossBucketUuid}/attach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CheckIAM2OrganizationAvailability operates on IAM2OrganizationAvailability
func (cli *ZSClient) CheckIAM2OrganizationAvailability(params param.CheckIAM2OrganizationAvailabilityParam) (*view.CheckIAM2OrganizationAvailabilityView, error) {
	var resp view.CheckIAM2OrganizationAvailabilityView
	if err := cli.Get("v1/iam2/organizations/availabilities", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnmountVmInstanceRecoveryPoint operates on UnmountVmInstanceRecoveryPoint
func (cli *ZSClient) UnmountVmInstanceRecoveryPoint(params param.UnmountVmInstanceRecoveryPointParam) (*view.UnmountVmInstanceRecoveryPointEventView, error) {
	resp := view.UnmountVmInstanceRecoveryPointEventView{}
	if err := cli.Post("v1/cdp-backup-storage/unmount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemovePolicyStatementsFromRole removes PolicyStatementsFromRole
func (cli *ZSClient) RemovePolicyStatementsFromRole(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/roles/{uuid}/policy-statements", uuid, string(deleteMode))
}

// GenerateModelMetadata operates on ModelMetadata
func (cli *ZSClient) GenerateModelMetadata(uuid string, params param.GenerateModelMetadataParam) (*view.GenerateModelMetadataEventView, error) {
	resp := view.GenerateModelMetadataEventView{}
	if err := cli.Put("v1/ai/model/metadata/generate", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsReadyToGo operates on IsReadyToGo
func (cli *ZSClient) IsReadyToGo(params param.IsReadyToGoParam) (*view.IsReadyToGoView, error) {
	var resp view.IsReadyToGoView
	if err := cli.Get("v1/management-nodes/ready", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostIommuStatus gets HostIommuStatus by uuid
func (cli *ZSClient) GetHostIommuStatus(uuid string) (*view.GetHostIommuStatusView, error) {
	var resp view.GetHostIommuStatusView
	if err := cli.Get("v1/pci-device/hosts/{uuid}/status", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DescribeVmInstanceRecoveryPoint operates on DescribeVmInstanceRecoveryPoint
func (cli *ZSClient) DescribeVmInstanceRecoveryPoint(params param.DescribeVmInstanceRecoveryPointParam) (*view.DescribeVmInstanceRecoveryPointView, error) {
	var resp view.DescribeVmInstanceRecoveryPointView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/recovery-point", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceCandidatesForAttachingVm gets PciDeviceCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetPciDeviceCandidatesForAttachingVm(uuid string) (*view.PciDeviceInventoryView, error) {
	var resp view.PciDeviceInventoryView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/candidate-pci-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeMonitorTriggerState changes MonitorTriggerState
func (cli *ZSClient) ChangeMonitorTriggerState(uuid string, params param.ChangeMonitorTriggerStateParam) (*view.MonitorTriggerInventoryView, error) {
	var resp view.ChangeMonitorTriggerStateEventView
	if err := cli.Put("v1/monitoring/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetBareMetal2ChassisPowerStatus gets BareMetal2ChassisPowerStatus by uuid
func (cli *ZSClient) GetBareMetal2ChassisPowerStatus(uuid string) (*view.GetBareMetal2ChassisPowerStatusView, error) {
	var resp view.GetBareMetal2ChassisPowerStatusView
	if err := cli.Get("v1/baremetal2/chassis/{uuid}/powerstatus", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTaskProgress gets TaskProgress by uuid
func (cli *ZSClient) GetTaskProgress(uuid string) (*view.TaskProgressInventoryView, error) {
	var resp view.TaskProgressInventoryView
	if err := cli.Get("v1/task-progresses/{apiId}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartDataProtection starts DataProtection
func (cli *ZSClient) StartDataProtection(params param.StartDataProtectionParam) (*view.StartDataProtectionEventView, error) {
	resp := view.StartDataProtectionEventView{}
	if err := cli.Post("v1/start/data/protection/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateSession operates on Session
func (cli *ZSClient) ValidateSession(params param.ValidateSessionParam) (*view.ValidateSessionView, error) {
	var resp view.ValidateSessionView
	if err := cli.Get("v1/accounts/sessions/{sessionUuid}/valid", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeActiveAlarmState changes ActiveAlarmState
func (cli *ZSClient) ChangeActiveAlarmState(uuid string, params param.ChangeActiveAlarmStateParam) (*view.ChangeActiveAlarmStateEventView, error) {
	resp := view.ChangeActiveAlarmStateEventView{}
	if err := cli.Put("v1/zwatch/activealarms/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmCleanTraffic operates on VmCleanTraffic
func (cli *ZSClient) SetVmCleanTraffic(uuid string, params param.SetVmCleanTrafficParam) (*view.SetVmCleanTrafficEventView, error) {
	resp := view.SetVmCleanTrafficEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmBootMode operates on VmBootMode
func (cli *ZSClient) SetVmBootMode(uuid string, params param.SetVmBootModeParam) (*view.SetVmBootModeEventView, error) {
	resp := view.SetVmBootModeEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncImageSize operates on ImageSize
func (cli *ZSClient) SyncImageSize(uuid string, params param.SyncImageSizeParam) (*view.ImageInventoryView, error) {
	var resp view.SyncImageSizeEventView
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetNoTriggerSchedulerJobs gets NoTriggerSchedulerJobs by uuid
func (cli *ZSClient) GetNoTriggerSchedulerJobs(uuid string) (*view.SchedulerJobInventoryView, error) {
	var resp view.SchedulerJobInventoryView
	if err := cli.Get("v1/scheduler/jobs/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddProxyToResource adds ProxyToResource
func (cli *ZSClient) AddProxyToResource(params param.AddProxyToResourceParam) (*view.UserProxyConfigResourceRefInventoryView, error) {
	var resp view.AddProxyToResourceEventView
	if err := cli.Post("v1/proxy/{proxyUuid}/resource/{resourceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ProtectVmInstanceRecoveryPoint operates on ProtectVmInstanceRecoveryPoint
func (cli *ZSClient) ProtectVmInstanceRecoveryPoint(uuid string, params param.ProtectVmInstanceRecoveryPointParam) (*view.ProtectVmInstanceRecoveryPointEventView, error) {
	resp := view.ProtectVmInstanceRecoveryPointEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/protect-recovery-point", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteConnectionAccessPointLocal deletes ConnectionAccessPointLocal
func (cli *ZSClient) DeleteConnectionAccessPointLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/access-point/{uuid}", uuid, string(deleteMode))
}

// QueryPhysicalDriveSelfTestHistory queries PhysicalDriveSelfTestHistory list
func (cli *ZSClient) QueryPhysicalDriveSelfTestHistory(params *param.QueryParam) ([]view.PhysicalDriveSmartSelfTestHistoryInventoryView, error) {
	var resp []view.PhysicalDriveSmartSelfTestHistoryInventoryView
	return resp, cli.List("v1/storage-devices/local-raid/physical-drives/self-test", params, &resp)
}

// RemoveIAM2VirtualIDsFromProject removes IAM2VirtualIDsFromProject
func (cli *ZSClient) RemoveIAM2VirtualIDsFromProject(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{projectUuid}/virtual-ids", uuid, string(deleteMode))
}

// CreateEcsImageFromEcsSnapshot creates EcsImageFromEcsSnapshot
func (cli *ZSClient) CreateEcsImageFromEcsSnapshot(params param.CreateEcsImageFromEcsSnapshotParam) (*view.EcsImageInventoryView, error) {
	var resp view.CreateEcsImageFromEcsSnapshotEventView
	if err := cli.Post("v1/hybrid/aliyun/image/snapshot", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateResourceStackFromApp creates ResourceStackFromApp
func (cli *ZSClient) CreateResourceStackFromApp(params param.CreateResourceStackFromAppParam) (*view.ResourceStackInventoryView, error) {
	var resp view.CreateResourceStackEventView
	if err := cli.Post("v1/appcenter/app/stack", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetSharedBlockCandidate gets SharedBlockCandidate by uuid
func (cli *ZSClient) GetSharedBlockCandidate(uuid string) (*view.GetSharedBlockCandidateView, error) {
	var resp view.GetSharedBlockCandidateView
	if err := cli.Get("v1/primary-storage/sharedblockgroup/sharedblock-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsSecurityGroupFromRemote operates on EcsSecurityGroupFromRemote
func (cli *ZSClient) SyncEcsSecurityGroupFromRemote(uuid string, params param.SyncEcsSecurityGroupFromRemoteParam) (*view.EcsSecurityGroupInventoryView, error) {
	resp := view.EcsSecurityGroupInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/security-group/{ecsVpcUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportBuildApp operates on BuildApp
func (cli *ZSClient) ExportBuildApp(uuid string, params param.ExportBuildAppParam) (*view.BuildAppExportHistoryInventoryView, error) {
	var resp view.ExportBuildAppEventView
	if err := cli.Put("v1/appcenter/buildapp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ReclaimSpaceFromImageStore operates on ReclaimSpaceFromImageStore
func (cli *ZSClient) ReclaimSpaceFromImageStore(uuid string, params param.ReclaimSpaceFromImageStoreParam) (*view.ReclaimSpaceFromImageStoreEventView, error) {
	resp := view.ReclaimSpaceFromImageStoreEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAllEventMetadata gets AllEventMetadata by uuid
func (cli *ZSClient) GetAllEventMetadata(uuid string) (*view.GetAllEventMetadataView, error) {
	var resp view.GetAllEventMetadataView
	if err := cli.Get("v1/zwatch/events/meta-data", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmForAttachingIso gets CandidateVmForAttachingIso by uuid
func (cli *ZSClient) GetCandidateVmForAttachingIso(uuid string) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.Get("v1/images/iso/{isoUuid}/vm-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachDataVolumeToVm operates on DataVolumeToVm
func (cli *ZSClient) AttachDataVolumeToVm(params param.AttachDataVolumeToVmParam) (*view.VolumeInventoryView, error) {
	var resp view.AttachDataVolumeToVmEventView
	if err := cli.Post("v1/volumes/{volumeUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateAliyunVirtualRouter updates AliyunVirtualRouter
func (cli *ZSClient) UpdateAliyunVirtualRouter(uuid string, params param.UpdateAliyunVirtualRouterParam) (*view.VpcVirtualRouterInventoryView, error) {
	var resp view.UpdateAliyunVirtualRouterEventView
	if err := cli.Put("v1/hybrid/aliyun/vrouter/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteDataVolume deletes DataVolume
func (cli *ZSClient) DeleteDataVolume(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{uuid}", uuid, string(deleteMode))
}

// GetUploadImageJobDetails gets UploadImageJobDetails by uuid
func (cli *ZSClient) GetUploadImageJobDetails(uuid string) (*view.GetUploadImageJobDetailsView, error) {
	var resp view.GetUploadImageJobDetailsView
	if err := cli.Get("v1/images/upload-job/details/{imageId}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachIscsiServerFromCluster operates on IscsiServerFromCluster
func (cli *ZSClient) DetachIscsiServerFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/storage-devices/iscsi/servers/{uuid}", uuid, string(deleteMode))
}

// SetVolumeQos operates on VolumeQos
func (cli *ZSClient) SetVolumeQos(uuid string, params param.SetVolumeQosParam) (*view.VolumeInventoryView, error) {
	var resp view.SetVolumeQosEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachHybridEipFromEcs operates on HybridEipFromEcs
func (cli *ZSClient) DetachHybridEipFromEcs(params param.DetachHybridEipFromEcsParam) (*view.DetachHybridEipFromEcsEventView, error) {
	resp := view.DetachHybridEipFromEcsEventView{}
	if err := cli.Post("v1/hybrid/eip/{eipUuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeCapabilities gets VolumeCapabilities by uuid
func (cli *ZSClient) GetVolumeCapabilities(uuid string) (*view.GetVolumeCapabilitiesView, error) {
	var resp view.GetVolumeCapabilitiesView
	if err := cli.Get("v1/volumes/{uuid}/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2GatewayCluster changes BareMetal2GatewayCluster
func (cli *ZSClient) ChangeBareMetal2GatewayCluster(uuid string, params param.ChangeBareMetal2GatewayClusterParam) (*view.BareMetal2GatewayInventoryView, error) {
	var resp view.ChangeBareMetal2GatewayClusterEventView
	if err := cli.Put("v1/baremetal2/gateways/{gatewayUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryVpcIkeConfigFromLocal queries VpcIkeConfigFromLocal list
func (cli *ZSClient) QueryVpcIkeConfigFromLocal(params *param.QueryParam) ([]view.VpcVpnIkeConfigInventoryView, error) {
	var resp []view.VpcVpnIkeConfigInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection/ike", params, &resp)
}

// SetVmInstanceHaLevel operates on VmInstanceHaLevel
func (cli *ZSClient) SetVmInstanceHaLevel(params param.SetVmInstanceHaLevelParam) (*view.SetVmInstanceHaLevelEventView, error) {
	resp := view.SetVmInstanceHaLevelEventView{}
	if err := cli.Post("v1/vm-instances/{uuid}/ha-levels", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVRouterNetworksFromFlowMeter removes VRouterNetworksFromFlowMeter
func (cli *ZSClient) RemoveVRouterNetworksFromFlowMeter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/flowmeters/networks", uuid, string(deleteMode))
}

// GetCandidateL3NetworksForChangeVmNicNetwork gets CandidateL3NetworksForChangeVmNicNetwork by uuid
func (cli *ZSClient) GetCandidateL3NetworksForChangeVmNicNetwork(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.Get("v1/vm-instances/nics/{vmNicUuid}/l3-networks-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteHybridKeySecret deletes HybridKeySecret
func (cli *ZSClient) DeleteHybridKeySecret(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/hybrid/key/{uuid}", uuid, string(deleteMode))
}

// GetCandidatePrimaryStoragesForCreatingVm gets CandidatePrimaryStoragesForCreatingVm by uuid
func (cli *ZSClient) GetCandidatePrimaryStoragesForCreatingVm(uuid string) (*view.GetCandidatePrimaryStoragesForCreatingVmView, error) {
	var resp view.GetCandidatePrimaryStoragesForCreatingVmView
	if err := cli.Get("v1/vm-instances/candidate-storages", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmConsolePassword gets VmConsolePassword by uuid
func (cli *ZSClient) GetVmConsolePassword(uuid string) (*view.GetVmConsolePasswordView, error) {
	var resp view.GetVmConsolePasswordView
	if err := cli.Get("v1/vm-instances/{uuid}/console-passwords", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceBindableConfig gets ResourceBindableConfig by uuid
func (cli *ZSClient) GetResourceBindableConfig(uuid string) (*view.GetResourceBindableConfigView, error) {
	var resp view.GetResourceBindableConfigView
	if err := cli.Get("v1/resource-configurations/bindable", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceHaLevel gets VmInstanceHaLevel by uuid
func (cli *ZSClient) GetVmInstanceHaLevel(uuid string) (*view.GetVmInstanceHaLevelView, error) {
	var resp view.GetVmInstanceHaLevelView
	if err := cli.Get("v1/vm-instances/{uuid}/ha-levels", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateLdapEntryForIAM2Binding gets CandidateLdapEntryForIAM2Binding by uuid
func (cli *ZSClient) GetCandidateLdapEntryForIAM2Binding(uuid string) (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.Get("v1/iam2/ldap/entries/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAudit queries Audit list
func (cli *ZSClient) QueryAudit(params *param.QueryParam) ([]view.AuditsInventoryView, error) {
	var resp []view.AuditsInventoryView
	return resp, cli.List("v1/zwatch/audit-records", params, &resp)
}

// RemoveResourcesFromDirectory removes ResourcesFromDirectory
func (cli *ZSClient) RemoveResourcesFromDirectory(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/remove/resources/directory", uuid, string(deleteMode))
}

// CreateVmFromVmBackup creates VmFromVmBackup
func (cli *ZSClient) CreateVmFromVmBackup(params param.CreateVmFromVmBackupParam) (*view.VmInstanceInventoryView, error) {
	var resp view.CreateVmFromVmBackupEventView
	if err := cli.Post("v1/vm-instances/from/vm-backups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteExportedDatabaseBackupFromBackupStorage deletes ExportedDatabaseBackupFromBackupStorage
func (cli *ZSClient) DeleteExportedDatabaseBackupFromBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/exported-database-backup/{databaseBackupUuid}/backup-storage/{backupStorageUuid}", uuid, string(deleteMode))
}

// AttachNetworkServiceToL3Network operates on NetworkServiceToL3Network
func (cli *ZSClient) AttachNetworkServiceToL3Network(params param.AttachNetworkServiceToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	var resp view.AttachNetworkServiceToL3NetworkEventView
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/network-services", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UnexportNbdVolumes operates on UnexportNbdVolumes
func (cli *ZSClient) UnexportNbdVolumes(params param.UnexportNbdVolumesParam) (*view.UnexportNbdVolumesEventView, error) {
	resp := view.UnexportNbdVolumesEventView{}
	if err := cli.Post("v1/cbt-task/unexportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoveryVirtualBorderRouterRemote operates on yVirtualBorderRouterRemote
func (cli *ZSClient) RecoveryVirtualBorderRouterRemote(uuid string, params param.RecoveryVirtualBorderRouterRemoteParam) (*view.RecoveryVirtualBorderRouterRemoteEventView, error) {
	resp := view.RecoveryVirtualBorderRouterRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/border-router/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcRouter queries VpcRouter list
func (cli *ZSClient) QueryVpcRouter(params *param.QueryParam) ([]view.VpcRouterVmInventoryView, error) {
	var resp []view.VpcRouterVmInventoryView
	return resp, cli.List("v1/vpc/virtual-routers", params, &resp)
}

// ExecuteAutoScalingRule operates on ExecuteAutoScalingRule
func (cli *ZSClient) ExecuteAutoScalingRule(uuid string, params param.ExecuteAutoScalingRuleParam) (*view.ExecuteAutoScalingRuleEventView, error) {
	resp := view.ExecuteAutoScalingRuleEventView{}
	if err := cli.Put("v1/autoscaling/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSHttpTestConnection operates on HttpTestConnection
func (cli *ZSClient) SNSHttpTestConnection(params param.SNSHttpTestConnectionParam) (*view.SNSHttpTestConnectionEventView, error) {
	resp := view.SNSHttpTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/http/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageSecurityLevel operates on ImageSecurityLevel
func (cli *ZSClient) SetImageSecurityLevel(uuid string, params param.SetImageSecurityLevelParam) (*view.SetImageSecurityLevelEventView, error) {
	resp := view.SetImageSecurityLevelEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2ChassisState changes BareMetal2ChassisState
func (cli *ZSClient) ChangeBareMetal2ChassisState(uuid string, params param.ChangeBareMetal2ChassisStateParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.ChangeBareMetal2ChassisStateEventView
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryHybridEipFromLocal queries HybridEipFromLocal list
func (cli *ZSClient) QueryHybridEipFromLocal(params *param.QueryParam) ([]view.HybridEipAddressInventoryView, error) {
	var resp []view.HybridEipAddressInventoryView
	return resp, cli.List("v1/hybrid/eip", params, &resp)
}

// AddHybridKeySecret adds HybridKeySecret
func (cli *ZSClient) AddHybridKeySecret(params param.AddHybridKeySecretParam) (*view.HybridAccountInventoryView, error) {
	var resp view.AddHybridKeySecretEventView
	if err := cli.Post("v1/hybrid/hybrid/key", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryAliyunRouteEntryFromLocal queries AliyunRouteEntryFromLocal list
func (cli *ZSClient) QueryAliyunRouteEntryFromLocal(params *param.QueryParam) ([]view.VpcVirtualRouteEntryInventoryView, error) {
	var resp []view.VpcVirtualRouteEntryInventoryView
	return resp, cli.List("v1/hybrid/aliyun/route-entry", params, &resp)
}

// DetachVmFromVmSchedulingRuleGroup operates on VmFromVmSchedulingRuleGroup
func (cli *ZSClient) DetachVmFromVmSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRuleGroup/{vmGroupUuid}/vmInstance/", uuid, string(deleteMode))
}

// AddVRouterNetworksToOspfArea adds VRouterNetworksToOspfArea
func (cli *ZSClient) AddVRouterNetworksToOspfArea(params param.AddVRouterNetworksToOspfAreaParam) (*view.NetworkRouterAreaRefInventoryView, error) {
	resp := view.NetworkRouterAreaRefInventoryView{}
	if err := cli.Post("v1/routerArea/{routerAreaUuid}/router/{vRouterUuid}/addnetworks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QuerySNSSmsEndpoint queries SNSSmsEndpoint list
func (cli *ZSClient) QuerySNSSmsEndpoint(params *param.QueryParam) ([]view.SNSAliyunSmsEndpointInventoryView, error) {
	var resp []view.SNSAliyunSmsEndpointInventoryView
	return resp, cli.List("v1/sns/sms-endpoints", params, &resp)
}

// AddRolesToIAM2VirtualIDGroup adds RolesToIAM2VirtualIDGroup
func (cli *ZSClient) AddRolesToIAM2VirtualIDGroup(params param.AddRolesToIAM2VirtualIDGroupParam) (*view.AddRolesToIAM2VirtualIDGroupEventView, error) {
	resp := view.AddRolesToIAM2VirtualIDGroupEventView{}
	if err := cli.Post("v1/iam2/projects/groups/{groupUuid}/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckStaticProvisionIp operates on StaticProvisionIp
func (cli *ZSClient) CheckStaticProvisionIp(params param.CheckStaticProvisionIpParam) (*view.CheckStaticProvisionIpView, error) {
	resp := view.CheckStaticProvisionIpView{}
	if err := cli.Post("v1/baremetal2/bm-instances/static/provision/ip/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeEventSubscriptionState changes EventSubscriptionState
func (cli *ZSClient) ChangeEventSubscriptionState(uuid string, params param.ChangeEventSubscriptionStateParam) (*view.EventSubscriptionInventoryView, error) {
	var resp view.ChangeEventSubscriptionStateEventView
	if err := cli.Put("v1/zwatch/change/eventSubscription/{uuid}/state", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PushLicenseAddOnsUsage operates on PushLicenseAddOnsUsage
func (cli *ZSClient) PushLicenseAddOnsUsage(uuid string, params param.PushLicenseAddOnsUsageParam) (*view.PushLicenseAddOnsUsageEventView, error) {
	resp := view.PushLicenseAddOnsUsageEventView{}
	if err := cli.Put("v1/licenses/addons/usage", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachHybridEipToEcs operates on HybridEipToEcs
func (cli *ZSClient) AttachHybridEipToEcs(params param.AttachHybridEipToEcsParam) (*view.HybridEipAddressInventoryView, error) {
	var resp view.AttachHybridEipToEcsEventView
	if err := cli.Post("v1/hybrid/eip/{eipUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateEcsImageFromLocalImage creates EcsImageFromLocalImage
func (cli *ZSClient) CreateEcsImageFromLocalImage(params param.CreateEcsImageFromLocalImageParam) (*view.EcsImageInventoryView, error) {
	var resp view.CreateEcsImageFromLocalImageEventView
	if err := cli.Post("v1/hybrid/aliyun/image", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddHostRouteToL3Network adds HostRouteToL3Network
func (cli *ZSClient) AddHostRouteToL3Network(params param.AddHostRouteToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	var resp view.AddHostRouteToL3NetworkEventView
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/hostroute", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddInstanceToMonitorGroup adds InstanceToMonitorGroup
func (cli *ZSClient) AddInstanceToMonitorGroup(params param.AddInstanceToMonitorGroupParam) (*view.MonitorGroupInstanceInventoryView, error) {
	var resp view.AddInstanceToMonitorGroupEventView
	if err := cli.Post("v1/zwatch/monitorgroups/{groupUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetBareMetal2ProvisionNetworkIpAddressCapacity gets BareMetal2ProvisionNetworkIpAddressCapacity by uuid
func (cli *ZSClient) GetBareMetal2ProvisionNetworkIpAddressCapacity(uuid string) (*view.GetBareMetal2ProvisionNetworkIpAddressCapacityView, error) {
	var resp view.GetBareMetal2ProvisionNetworkIpAddressCapacityView
	if err := cli.Get("v1/baremetal2/provision-networks/ip-capacity", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachMdevDeviceToVm operates on MdevDeviceToVm
func (cli *ZSClient) AttachMdevDeviceToVm(params param.AttachMdevDeviceToVmParam) (*view.MdevDeviceInventoryView, error) {
	var resp view.AttachMdevDeviceToVmEventView
	if err := cli.Post("v1/mdev-devices/{mdevDeviceUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DecodeStackTemplate operates on DecodeStackTemplate
func (cli *ZSClient) DecodeStackTemplate(params param.DecodeStackTemplateParam) (*view.DecodeStackTemplateView, error) {
	resp := view.DecodeStackTemplateView{}
	if err := cli.Post("v1/cloudformation/stack/preview/resource", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVirtualRouter updates VirtualRouter
func (cli *ZSClient) UpdateVirtualRouter(uuid string, params param.UpdateVirtualRouterParam) (*view.VirtualRouterVmInventoryView, error) {
	var resp view.UpdateVirtualRouterEventView
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVSwitchTypes gets VSwitchTypes by uuid
func (cli *ZSClient) GetVSwitchTypes(uuid string) (*view.GetVSwitchTypesView, error) {
	var resp view.GetVSwitchTypesView
	if err := cli.Get("v1/l2-networks/vSwitchTypes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsInstanceFromLocal queries EcsInstanceFromLocal list
func (cli *ZSClient) QueryEcsInstanceFromLocal(params *param.QueryParam) ([]view.EcsInstanceInventoryView, error) {
	var resp []view.EcsInstanceInventoryView
	return resp, cli.List("v1/hybrid/aliyun/ecs", params, &resp)
}

// CreateL2HardwareVxlanNetworkPool creates L2HardwareVxlanNetworkPool
func (cli *ZSClient) CreateL2HardwareVxlanNetworkPool(params param.CreateL2HardwareVxlanNetworkPoolParam) (*view.CreateL2HardwareVxlanNetworkPoolEventView, error) {
	resp := view.CreateL2HardwareVxlanNetworkPoolEventView{}
	if err := cli.Post("v1/l2-networks/hardware-vxlan-pool", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLdapServerAvailableAttributes gets LdapServerAvailableAttributes by uuid
func (cli *ZSClient) GetLdapServerAvailableAttributes(uuid string) (*view.GetLdapServerAvailableAttributesView, error) {
	var resp view.GetLdapServerAvailableAttributesView
	if err := cli.Get("v1/ldap/server/attributes/{uuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResizeDataVolume operates on DataVolume
func (cli *ZSClient) ResizeDataVolume(uuid string, params param.ResizeDataVolumeParam) (*view.VolumeInventoryView, error) {
	var resp view.ResizeDataVolumeEventView
	if err := cli.Put("v1/volumes/data/resize/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetEipAttachableVmNics gets EipAttachableVmNics by uuid
func (cli *ZSClient) GetEipAttachableVmNics(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.Get("v1/eips/{eipUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIpv6RangeByNetworkCidr adds Ipv6RangeByNetworkCidr
func (cli *ZSClient) AddIpv6RangeByNetworkCidr(params param.AddIpv6RangeByNetworkCidrParam) (*view.IpRangeInventoryView, error) {
	var resp view.AddIpRangeByNetworkCidrEventView
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/ipv6-ranges/by-cidr", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// BatchQuery operates on Query
func (cli *ZSClient) BatchQuery(params param.BatchQueryParam) (*view.BatchQueryView, error) {
	var resp view.BatchQueryView
	if err := cli.Get("v1/batch-queries", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReloadExternalService operates on ReloadExternalService
func (cli *ZSClient) ReloadExternalService(uuid string, params param.ReloadExternalServiceParam) (*view.ReloadExternalServiceEventView, error) {
	resp := view.ReloadExternalServiceEventView{}
	if err := cli.Put("v1/external/services", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIAM2VirtualIDsToGroup adds IAM2VirtualIDsToGroup
func (cli *ZSClient) AddIAM2VirtualIDsToGroup(params param.AddIAM2VirtualIDsToGroupParam) (*view.AddIAM2VirtualIDToGroupEventView, error) {
	resp := view.AddIAM2VirtualIDToGroupEventView{}
	if err := cli.Post("v1/iam2/projects/groups/{groupUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2VirtualIDLdapBinding creates IAM2VirtualIDLdapBinding
func (cli *ZSClient) CreateIAM2VirtualIDLdapBinding(params param.CreateIAM2VirtualIDLdapBindingParam) (*view.LdapResourceRefInventoryView, error) {
	var resp view.CreateIAM2VirtualIDLdapBindingEventView
	if err := cli.Post("v1/iam2/ldap/bindings", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetVmNicSecurityGroup operates on VmNicSecurityGroup
func (cli *ZSClient) SetVmNicSecurityGroup(uuid string, params param.SetVmNicSecurityGroupParam) (*view.VmNicSecurityGroupRefInventoryView, error) {
	resp := view.VmNicSecurityGroupRefInventoryView{}
	if err := cli.Put("v1/security-groups/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryResourcePrice queries ResourcePrice list
func (cli *ZSClient) QueryResourcePrice(params *param.QueryParam) ([]view.PriceInventoryView, error) {
	var resp []view.PriceInventoryView
	return resp, cli.List("v1/billings/prices", params, &resp)
}

// AddIdentityZoneFromRemote adds IdentityZoneFromRemote
func (cli *ZSClient) AddIdentityZoneFromRemote(params param.AddIdentityZoneFromRemoteParam) (*view.IdentityZoneInventoryView, error) {
	var resp view.AddIdentityZoneFromRemoteEventView
	if err := cli.Post("v1/hybrid/identity-zone", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVolumeSnapshotSize gets VolumeSnapshotSize by uuid
func (cli *ZSClient) GetVolumeSnapshotSize(uuid string) (*view.GetVolumeSnapshotSizeEventView, error) {
	var resp view.GetVolumeSnapshotSizeEventView
	if err := cli.Get("v1/volume-snapshots/{uuid}/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchSyncVolumeSize operates on SyncVolumeSize
func (cli *ZSClient) BatchSyncVolumeSize(params param.BatchSyncVolumeSizeParam) (*view.BatchSyncVolumeSizeView, error) {
	resp := view.BatchSyncVolumeSizeView{}
	if err := cli.Post("v1/volumes/batch-sync-volumes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHypervisorTypes gets HypervisorTypes by uuid
func (cli *ZSClient) GetHypervisorTypes(uuid string) (*view.GetHypervisorTypesView, error) {
	var resp view.GetHypervisorTypesView
	if err := cli.Get("v1/hosts/hypervisor-types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmAttachableDataVolume gets VmAttachableDataVolume by uuid
func (cli *ZSClient) GetVmAttachableDataVolume(uuid string) (*view.VolumeInventoryView, error) {
	var resp view.VolumeInventoryView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/data-volume-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmMonitorNumber gets VmMonitorNumber by uuid
func (cli *ZSClient) GetVmMonitorNumber(uuid string) (*view.GetVmMonitorNumberView, error) {
	var resp view.GetVmMonitorNumberView
	if err := cli.Get("v1/vm-instances/{uuid}/monitorNumber", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2VirtualIDFromLdapUid creates IAM2VirtualIDFromLdapUid
func (cli *ZSClient) CreateIAM2VirtualIDFromLdapUid(params param.CreateIAM2VirtualIDFromLdapUidParam) (*view.LdapResourceRefInventoryView, error) {
	var resp view.CreateIAM2VirtualIDFromLdapUidEventView
	if err := cli.Post("v1/iam2/virtual-id/ldap/uid", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ValidatePriceUserConfig operates on PriceUserConfig
func (cli *ZSClient) ValidatePriceUserConfig(uuid string, params param.ValidatePriceUserConfigParam) (*view.ValidatePriceUserConfigEventView, error) {
	resp := view.ValidatePriceUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2GatewayState changes BareMetal2GatewayState
func (cli *ZSClient) ChangeBareMetal2GatewayState(uuid string, params param.ChangeBareMetal2GatewayStateParam) (*view.BareMetal2GatewayInventoryView, error) {
	var resp view.ChangeBareMetal2GatewayStateEventView
	if err := cli.Put("v1/baremetal2/gateways/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveActionFromEventSubscription removes ActionFromEventSubscription
func (cli *ZSClient) RemoveActionFromEventSubscription(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions/{subscriptionUuid}/actions/{actionUuid}", uuid, string(deleteMode))
}

// CheckKVMHostConfigFile operates on KVMHostConfigFile
func (cli *ZSClient) CheckKVMHostConfigFile(params param.CheckKVMHostConfigFileParam) (*view.CheckHostConfigFileView, error) {
	resp := view.CheckHostConfigFileView{}
	if err := cli.Post("v1/hosts/kvm/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetContainerUsage gets ContainerUsage by uuid
func (cli *ZSClient) GetContainerUsage(uuid string) (*view.GetContainerUsageView, error) {
	var resp view.GetContainerUsageView
	if err := cli.Get("v1/container/usage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSSnmpTestConnection operates on SnmpTestConnection
func (cli *ZSClient) SNSSnmpTestConnection(params param.SNSSnmpTestConnectionParam) (*view.SNSSnmpTestConnectionEventView, error) {
	resp := view.SNSSnmpTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/snmp/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDataCenterFromRemote gets DataCenterFromRemote by uuid
func (cli *ZSClient) GetDataCenterFromRemote(uuid string) (*view.DataCenterPropertyView, error) {
	var resp view.DataCenterPropertyView
	if err := cli.Get("v1/hybrid/data-center/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryGCJob queries GCJob list
func (cli *ZSClient) QueryGCJob(params *param.QueryParam) ([]view.GarbageCollectorInventoryView, error) {
	var resp []view.GarbageCollectorInventoryView
	return resp, cli.List("v1/gc-jobs", params, &resp)
}

// CreateHostNetworkServiceType creates HostNetworkServiceType
func (cli *ZSClient) CreateHostNetworkServiceType(params param.CreateHostNetworkServiceTypeParam) (*view.HostNetworkLabelInventoryView, error) {
	var resp view.CreateHostNetworkServiceTypeEventView
	if err := cli.Post("v1/hosts/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteEcsImageLocal deletes EcsImageLocal
func (cli *ZSClient) DeleteEcsImageLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/image/{uuid}", uuid, string(deleteMode))
}

// DetachNvmeServerFromCluster operates on NvmeServerFromCluster
func (cli *ZSClient) DetachNvmeServerFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/storage-devices/nvme/servers/{uuid}", uuid, string(deleteMode))
}

// GetBackupStorageTypes gets BackupStorageTypes by uuid
func (cli *ZSClient) GetBackupStorageTypes(uuid string) (*view.GetBackupStorageTypesView, error) {
	var resp view.GetBackupStorageTypesView
	if err := cli.Get("v1/backup-storage/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeQos gets VolumeQos by uuid
func (cli *ZSClient) GetVolumeQos(uuid string) (*view.GetVolumeQosView, error) {
	var resp view.GetVolumeQosView
	if err := cli.Get("v1/volumes/{uuid}/qos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRemoteCidrsToIPsecConnection adds RemoteCidrsToIPsecConnection
func (cli *ZSClient) AddRemoteCidrsToIPsecConnection(params param.AddRemoteCidrsToIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	var resp view.AddRemoteCidrsToIPsecConnectionEventView
	if err := cli.Post("v1/ipsec/{uuid}/remote-cidrs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PowerOnBaremetalChassis operates on PowerOnBaremetalChassis
func (cli *ZSClient) PowerOnBaremetalChassis(uuid string, params param.PowerOnBaremetalChassisParam) (*view.PowerOnBaremetalChassisEventView, error) {
	resp := view.PowerOnBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RequestLicenseCapacity operates on RequestLicenseCapacity
func (cli *ZSClient) RequestLicenseCapacity(params param.RequestLicenseCapacityParam) (*view.LicenseAuthorizedCapacityInventoryView, error) {
	var resp view.RequestLicenseCapacityEventView
	if err := cli.Post("v1/license-server/capacity-application", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateDataVolumeFromVolumeSnapshot creates DataVolumeFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeFromVolumeSnapshot(params param.CreateDataVolumeFromVolumeSnapshotParam) (*view.VolumeInventoryView, error) {
	var resp view.CreateDataVolumeFromVolumeSnapshotEventView
	if err := cli.Post("v1/volumes/data/from/volume-snapshots/{volumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachIsoFromVmInstance operates on IsoFromVmInstance
func (cli *ZSClient) DetachIsoFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/iso", uuid, string(deleteMode))
}

// DetachSecurityGroupFromL3Network operates on SecurityGroupFromL3Network
func (cli *ZSClient) DetachSecurityGroupFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups/{securityGroupUuid}/l3-networks/{l3NetworkUuid}", uuid, string(deleteMode))
}

// GetVirtualizerInfo gets VirtualizerInfo by uuid
func (cli *ZSClient) GetVirtualizerInfo(uuid string) (*view.VirtualizerInfoInventoryView, error) {
	var resp view.VirtualizerInfoInventoryView
	if err := cli.Get("v1/vm-instances/virtualizer-info", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkIpStatistic gets L3NetworkIpStatistic by uuid
func (cli *ZSClient) GetL3NetworkIpStatistic(uuid string) (*view.GetL3NetworkIpStatisticView, error) {
	var resp view.GetL3NetworkIpStatisticView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/ip-statistic", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImageCandidatesForVmToChange gets ImageCandidatesForVmToChange by uuid
func (cli *ZSClient) GetImageCandidatesForVmToChange(uuid string) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/image-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeImageState changes ImageState
func (cli *ZSClient) ChangeImageState(uuid string, params param.ChangeImageStateParam) (*view.ImageInventoryView, error) {
	var resp view.ChangeImageStateEventView
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// KvmRunShell operates on KvmRunShell
func (cli *ZSClient) KvmRunShell(uuid string, params param.KvmRunShellParam) (*view.KvmRunShellEventView, error) {
	resp := view.KvmRunShellEventView{}
	if err := cli.Put("v1/hosts/kvm/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAliyunNasAccessGroupRule creates AliyunNasAccessGroupRule
func (cli *ZSClient) CreateAliyunNasAccessGroupRule(params param.CreateAliyunNasAccessGroupRuleParam) (*view.AliyunNasAccessRuleInventoryView, error) {
	var resp view.CreateAliyunNasAccessGroupRuleEventView
	if err := cli.Post("v1/nas/aliyun/rule", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RecoverBackupFromImageStoreBackupStorage operates on BackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverBackupFromImageStoreBackupStorage(uuid string, params param.RecoverBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	var resp view.RecoverBackupFromImageStoreBackupStorageEventView
	if err := cli.Put("v1/volume-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeTicketFlowCollectionState changes TicketFlowCollectionState
func (cli *ZSClient) ChangeTicketFlowCollectionState(uuid string, params param.ChangeTicketFlowCollectionStateParam) (*view.TicketFlowCollectionInventoryView, error) {
	var resp view.ChangeTicketFlowCollectionStateEventView
	if err := cli.Put("v1/tickets/flow-collections/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ExpungeDataVolume operates on DataVolume
func (cli *ZSClient) ExpungeDataVolume(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{uuid}/actions", uuid, string(deleteMode))
}

// AddActionToEventSubscription adds ActionToEventSubscription
func (cli *ZSClient) AddActionToEventSubscription(params param.AddActionToEventSubscriptionParam) (*view.EventSubscriptionInventoryView, error) {
	var resp view.AddActionToEventSubscriptionEventView
	if err := cli.Post("v1/zwatch/events/subscriptions/{subscriptionUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVRouterRouterId gets VRouterRouterId by uuid
func (cli *ZSClient) GetVRouterRouterId(uuid string) (*view.GetVRouterRouterIdView, error) {
	var resp view.GetVRouterRouterIdView
	if err := cli.Get("v1/routerArea/{vRouterUuid}/routerid", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetZBoxBackupDetails gets ZBoxBackupDetails by uuid
func (cli *ZSClient) GetZBoxBackupDetails(uuid string) (*view.GetZBoxBackupDetailsView, error) {
	var resp view.GetZBoxBackupDetailsView
	if err := cli.Get("v1/externalbackup/zbox/{uuid}/details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetExternalServices gets ExternalServices by uuid
func (cli *ZSClient) GetExternalServices(uuid string) (*view.ExternalServiceInventoryView, error) {
	var resp view.ExternalServiceInventoryView
	if err := cli.Get("v1/external/services", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2ProjectRepository gets IAM2ProjectRepository by uuid
func (cli *ZSClient) GetIAM2ProjectRepository(uuid string) (*view.ProjectRepositoryInventoryView, error) {
	var resp view.ProjectRepositoryInventoryView
	if err := cli.Get("v1/iam2/projects/repositories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateNetworkInterfaces gets CandidateNetworkInterfaces by uuid
func (cli *ZSClient) GetCandidateNetworkInterfaces(uuid string) (*view.GetCandidateNetworkInterfacesView, error) {
	var resp view.GetCandidateNetworkInterfacesView
	if err := cli.Get("v1/cluster/hosts-network-interfaces", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAccessControlListServerGroup changes AccessControlListServerGroup
func (cli *ZSClient) ChangeAccessControlListServerGroup(uuid string, params param.ChangeAccessControlListServerGroupParam) (*view.LoadBalancerListerAclView, error) {
	var resp view.ChangeAccessControlListServerGroupEventView
	if err := cli.Put("v1/load-balancers/listener/acl/{aclUuid}/servergroup/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SyncVirtualBorderRouterFromRemote operates on VirtualBorderRouterFromRemote
func (cli *ZSClient) SyncVirtualBorderRouterFromRemote(uuid string, params param.SyncVirtualBorderRouterFromRemoteParam) (*view.VirtualBorderRouterInventoryView, error) {
	resp := view.VirtualBorderRouterInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/border-router/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAtPersonOfAtFeiShuEndpoint updates AtPersonOfAtFeiShuEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtFeiShuEndpoint(uuid string, params param.UpdateAtPersonOfAtFeiShuEndpointParam) (*view.SNSFeiShuAtPersonInventoryView, error) {
	var resp view.UpdateAtPersonOfFeiShuEndpointEventView
	if err := cli.Put("v1/sns/application-endpoints/feishu/at-persons/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateL2HardwareVxlanNetwork creates L2HardwareVxlanNetwork
func (cli *ZSClient) CreateL2HardwareVxlanNetwork(params param.CreateL2HardwareVxlanNetworkParam) (*view.CreateL2HardwareVxlanNetworkEventView, error) {
	resp := view.CreateL2HardwareVxlanNetworkEventView{}
	if err := cli.Post("v1/l2-networks/hardware-vxlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGlobalConfigOptions gets GlobalConfigOptions by uuid
func (cli *ZSClient) GetGlobalConfigOptions(uuid string) (*view.GetGlobalConfigOptionsView, error) {
	var resp view.GetGlobalConfigOptionsView
	if err := cli.Get("v1/global-configurations/{category}/{name}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateHybridEip creates HybridEip
func (cli *ZSClient) CreateHybridEip(params param.CreateHybridEipParam) (*view.HybridEipAddressInventoryView, error) {
	var resp view.CreateHybridEipEventView
	if err := cli.Post("v1/hybrid/eip", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ApplyMonitorTemplateToMonitorGroup operates on MonitorTemplateToMonitorGroup
func (cli *ZSClient) ApplyMonitorTemplateToMonitorGroup(params param.ApplyMonitorTemplateToMonitorGroupParam) (*view.MonitorGroupTemplateRefInventoryView, error) {
	var resp view.ApplyMonitorTemplateToMonitorGroupEventView
	if err := cli.Post("v1/zwatch/monitortemplates/{templateUuid}/monitorgroups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PutMetricData operates on PutMetricData
func (cli *ZSClient) PutMetricData(params param.PutMetricDataParam) (*view.PutMetricDataEventView, error) {
	resp := view.PutMetricDataEventView{}
	if err := cli.Post("v1/zwatch/metrics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAttachablePublicL3ForVRouter gets AttachablePublicL3ForVRouter by uuid
func (cli *ZSClient) GetAttachablePublicL3ForVRouter(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.Get("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/attachable-public-l3s", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RerunLongJob operates on RerunLongJob
func (cli *ZSClient) RerunLongJob(uuid string, params param.RerunLongJobParam) (*view.LongJobInventoryView, error) {
	var resp view.RerunLongJobEventView
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryFirewallRuleSet queries FirewallRuleSet list
func (cli *ZSClient) QueryFirewallRuleSet(params *param.QueryParam) ([]view.VpcFirewallRuleSetInventoryView, error) {
	var resp []view.VpcFirewallRuleSetInventoryView
	return resp, cli.List("v1/vpcfirewalls/ruleSets", params, &resp)
}

// DeleteExportedImageFromBackupStorage deletes ExportedImageFromBackupStorage
func (cli *ZSClient) DeleteExportedImageFromBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/backup-storage/{backupStorageUuid}/exported-images/{imageUuid}", uuid, string(deleteMode))
}

// UpdateClusterOS updates ClusterOS
func (cli *ZSClient) UpdateClusterOS(uuid string, params param.UpdateClusterOSParam) (*view.LongJobInventoryView, error) {
	var resp view.UpdateClusterOSEventView
	if err := cli.Put("v1/clusters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVmUsbRedirect gets VmUsbRedirect by uuid
func (cli *ZSClient) GetVmUsbRedirect(uuid string) (*view.GetVmUsbRedirectView, error) {
	var resp view.GetVmUsbRedirectView
	if err := cli.Get("v1/vm-instances/{uuid}/usbredirect", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateImageGroupFromSnapshot creates ImageGroupFromSnapshot
func (cli *ZSClient) CreateImageGroupFromSnapshot(params param.CreateImageGroupFromSnapshotParam) (*view.ImageGroupInventoryView, error) {
	var resp view.CreateImageGroupFromSnapshotEventView
	if err := cli.Post("v1/imagegroup/from/snapshot/{rootVolumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetOssBucketFileFromRemote gets OssBucketFileFromRemote by uuid
func (cli *ZSClient) GetOssBucketFileFromRemote(uuid string) (*view.GetOssBucketFileFromRemoteView, error) {
	var resp view.GetOssBucketFileFromRemoteView
	if err := cli.Get("v1/hybrid/oss/file/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVipToVpcSharedQos operates on VipToVpcSharedQos
func (cli *ZSClient) AttachVipToVpcSharedQos(params param.AttachVipToVpcSharedQosParam) (*view.AttachVipToVpcSharedQosEventView, error) {
	resp := view.AttachVipToVpcSharedQosEventView{}
	if err := cli.Post("v1/vips/sharedqos/{sharedQosUuid}/vips", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEventData gets EventData by uuid
func (cli *ZSClient) GetEventData(uuid string) (*view.GetEventDataView, error) {
	var resp view.GetEventDataView
	if err := cli.Get("v1/zwatch/events", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIpAvailability operates on IpAvailability
func (cli *ZSClient) CheckIpAvailability(params param.CheckIpAvailabilityParam) (*view.CheckIpAvailabilityView, error) {
	var resp view.CheckIpAvailabilityView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/ip/{ip}/availability", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVmNicFromLoadBalancer removes VmNicFromLoadBalancer
func (cli *ZSClient) RemoveVmNicFromLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics", uuid, string(deleteMode))
}

// RemoveRolesFromIAM2VirtualID removes RolesFromIAM2VirtualID
func (cli *ZSClient) RemoveRolesFromIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/virtual-ids/{virtualIDUuid}/roles", uuid, string(deleteMode))
}

// CalculateResourceSpending operates on ResourceSpending
func (cli *ZSClient) CalculateResourceSpending(uuid string, params param.CalculateResourceSpendingParam) (*view.CalculateResourceSpendingView, error) {
	resp := view.CalculateResourceSpendingView{}
	if err := cli.Put("v1/billings/resources/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAlarmRecord queries AlarmRecord list
func (cli *ZSClient) QueryAlarmRecord(params *param.QueryParam) ([]view.AlarmRecordsInventoryView, error) {
	var resp []view.AlarmRecordsInventoryView
	return resp, cli.List("v1/zwatch/alarm-records", params, &resp)
}

// DetachBackupStorageFromZone operates on BackupStorageFromZone
func (cli *ZSClient) DetachBackupStorageFromZone(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}", uuid, string(deleteMode))
}

// UpdateCCSCertificateUserState updates CCSCertificateUserState
func (cli *ZSClient) UpdateCCSCertificateUserState(uuid string, params param.UpdateCCSCertificateUserStateParam) (*view.CCSCertificateInventoryView, error) {
	var resp view.UpdateCCSCertificateUserStateEventView
	if err := cli.Put("v1/crypto/ccs-certificate/update-state/{userUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PowerResetBaremetalChassis operates on PowerResetBaremetalChassis
func (cli *ZSClient) PowerResetBaremetalChassis(uuid string, params param.PowerResetBaremetalChassisParam) (*view.PowerResetBaremetalChassisEventView, error) {
	resp := view.PowerResetBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpTrashOnPrimaryStorage operates on UpTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpTrashOnPrimaryStorage(uuid string, params param.CleanUpTrashOnPrimaryStorageParam) (*view.CleanUpTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpTrashOnPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/trash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDisasterImageStoreBackupStorage adds DisasterImageStoreBackupStorage
func (cli *ZSClient) AddDisasterImageStoreBackupStorage(params param.AddDisasterImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	var resp view.AddImageStoreBackupStorageEventView
	if err := cli.Post("v1/backup-storage/image-store/disaster", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVmSchedulingRulesExecuteState gets VmSchedulingRulesExecuteState by uuid
func (cli *ZSClient) GetVmSchedulingRulesExecuteState(uuid string) (*view.GetVmSchedulingRulesExecuteStateView, error) {
	var resp view.GetVmSchedulingRulesExecuteStateView
	if err := cli.Get("v1/get/vmSchedulingRules/conflict/state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVolumesSnapshot creates VolumesSnapshot
func (cli *ZSClient) CreateVolumesSnapshot(params param.CreateVolumesSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	resp := view.VolumeSnapshotInventoryView{}
	if err := cli.Post("v1/volumes/volume-snapshots", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIpAddressCapacity gets IpAddressCapacity by uuid
func (cli *ZSClient) GetIpAddressCapacity(uuid string) (*view.GetIpAddressCapacityView, error) {
	var resp view.GetIpAddressCapacityView
	if err := cli.Get("v1/ip-capacity", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIAM2ProjectContainerCluster operates on IAM2ProjectContainerCluster
func (cli *ZSClient) SetIAM2ProjectContainerCluster(uuid string, params param.SetIAM2ProjectContainerClusterParam) (*view.SetIAM2ProjectContainerClusterEventView, error) {
	resp := view.SetIAM2ProjectContainerClusterEventView{}
	if err := cli.Put("v1/iam2/projects/{uuid}/container/cluster/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployAppDevelopmentService operates on DeployAppDevelopmentService
func (cli *ZSClient) DeployAppDevelopmentService(uuid string, params param.DeployAppDevelopmentServiceParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.DeployAppDevelopmentServiceEventView
	if err := cli.Put("v1/ai/model-services/app/", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RefreshPluginDrivers operates on PluginDrivers
func (cli *ZSClient) RefreshPluginDrivers(uuid string, params param.RefreshPluginDriversParam) (*view.RefreshPluginDriversEventView, error) {
	resp := view.RefreshPluginDriversEventView{}
	if err := cli.Put("v1/external/plugins", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PauseVmInstance operates on PauseVmInstance
func (cli *ZSClient) PauseVmInstance(uuid string, params param.PauseVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.PauseVmInstanceEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachUserDefinedXmlHookScriptFromVm operates on UserDefinedXmlHookScriptFromVm
func (cli *ZSClient) DetachUserDefinedXmlHookScriptFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/xmlhook/vm-instances/{vmInstanceUuid}/detach", uuid, string(deleteMode))
}

// GetSignatureServerEncryptPublicKey gets SignatureServerEncryptPublicKey by uuid
func (cli *ZSClient) GetSignatureServerEncryptPublicKey(uuid string) (*view.GetSignatureServerEncryptPublicKeyView, error) {
	var resp view.GetSignatureServerEncryptPublicKeyView
	if err := cli.Get("v1/secret-resource-pool-token/signature-server-encrypt-public-key", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAliyunKeySecret adds AliyunKeySecret
func (cli *ZSClient) AddAliyunKeySecret(params param.AddAliyunKeySecretParam) (*view.HybridAccountInventoryView, error) {
	var resp view.AddAliyunKeySecretEventView
	if err := cli.Post("v1/hybrid/aliyun/key", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddBackupStoragesToReplicationGroup adds BackupStoragesToReplicationGroup
func (cli *ZSClient) AddBackupStoragesToReplicationGroup(params param.AddBackupStoragesToReplicationGroupParam) (*view.ImageReplicationGroupBackupStorageRefInventoryView, error) {
	resp := view.ImageReplicationGroupBackupStorageRefInventoryView{}
	if err := cli.Post("v1/image-replication-groups/{replicationGroupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDataCenterFromRemote adds DataCenterFromRemote
func (cli *ZSClient) AddDataCenterFromRemote(params param.AddDataCenterFromRemoteParam) (*view.DataCenterInventoryView, error) {
	var resp view.AddDataCenterFromRemoteEventView
	if err := cli.Post("v1/hybrid/data-center", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteFirewallRuleSet deletes FirewallRuleSet
func (cli *ZSClient) DeleteFirewallRuleSet(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/ruleSets/{uuid}", uuid, string(deleteMode))
}

// BatchAddBareMetal2IpmiChassis operates on AddBareMetal2IpmiChassis
func (cli *ZSClient) BatchAddBareMetal2IpmiChassis(params param.BatchAddBareMetal2IpmiChassisParam) (*view.LongJobInventoryView, error) {
	var resp view.BatchAddBareMetal2ChassisEventView
	if err := cli.Post("v1/baremetal2/chassis/ipmi/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// LocalStorageMigrateVolume operates on LocalStorageMigrateVolume
func (cli *ZSClient) LocalStorageMigrateVolume(uuid string, params param.LocalStorageMigrateVolumeParam) (*view.LocalStorageResourceRefInventoryView, error) {
	var resp view.LocalStorageMigrateVolumeEventView
	if err := cli.Put("v1/primary-storage/local-storage/volumes/{volumeUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachNicToBonding operates on NicToBonding
func (cli *ZSClient) AttachNicToBonding(uuid string, params param.AttachNicToBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	var resp view.AttachNicToBondingEventView
	if err := cli.Put("v1/hosts/bondings/{uuid}/attach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetOrganizationOperation operates on OrganizationOperation
func (cli *ZSClient) SetOrganizationOperation(uuid string, params param.SetOrganizationOperationParam) (*view.SetOrganizationOperationEventView, error) {
	resp := view.SetOrganizationOperationEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/operation", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolume creates DataVolumeTemplateFromVolume
func (cli *ZSClient) CreateDataVolumeTemplateFromVolume(params param.CreateDataVolumeTemplateFromVolumeParam) (*view.ImageInventoryView, error) {
	var resp view.CreateDataVolumeTemplateFromVolumeEventView
	if err := cli.Post("v1/images/data-volume-templates/from/volumes/{volumeUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveIAM2VirtualIDsFromOrganization removes IAM2VirtualIDsFromOrganization
func (cli *ZSClient) RemoveIAM2VirtualIDsFromOrganization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/organizations/{organizationUuid}/virtual-ids", uuid, string(deleteMode))
}

// ExportDatabaseBackupFromBackupStorage operates on DatabaseBackupFromBackupStorage
func (cli *ZSClient) ExportDatabaseBackupFromBackupStorage(uuid string, params param.ExportDatabaseBackupFromBackupStorageParam) (*view.ExportDatabaseBackupFromBackupStorageEventView, error) {
	resp := view.ExportDatabaseBackupFromBackupStorageEventView{}
	if err := cli.Put("v1/database-backups/{databaseBackupUuid}/backup-storage/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachIAM2ProjectToIAM2Organization operates on IAM2ProjectToIAM2Organization
func (cli *ZSClient) AttachIAM2ProjectToIAM2Organization(params param.AttachIAM2ProjectToIAM2OrganizationParam) (*view.IAM2ProjectInventoryView, error) {
	var resp view.AttachIAM2ProjectToIAM2OrganizationEventView
	if err := cli.Post("v1/iam2/projects/{projectUuid}/iam2/organizations/{organizationUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateEmailMonitorTriggerAction creates EmailMonitorTrigger
func (cli *ZSClient) CreateEmailMonitorTriggerAction(params param.CreateEmailMonitorTriggerActionParam) (*view.MonitorTriggerActionInventoryView, error) {
	var resp view.CreateMonitorTriggerActionEventView
	if err := cli.Post("v1/monitoring/trigger-actions/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetVpcVRouterDistributedRoutingEnabled operates on VpcVRouterDistributedRoutingEnabled
func (cli *ZSClient) SetVpcVRouterDistributedRoutingEnabled(params param.SetVpcVRouterDistributedRoutingEnabledParam) (*view.SetVpcVRouterDistributedRoutingEnabledEventView, error) {
	resp := view.SetVpcVRouterDistributedRoutingEnabledEventView{}
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/distributed-routing", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOnBareMetal2Chassis operates on PowerOnBareMetal2Chassis
func (cli *ZSClient) PowerOnBareMetal2Chassis(uuid string, params param.PowerOnBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.PowerOnBareMetal2ChassisEventView
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetLocalRaidPhysicalDriveSmart gets LocalRaidPhysicalDriveSmart by uuid
func (cli *ZSClient) GetLocalRaidPhysicalDriveSmart(uuid string) (*view.GetLocalRaidPhysicalDriveSmartView, error) {
	var resp view.GetLocalRaidPhysicalDriveSmartView
	if err := cli.Get("v1/storage-devices/local-raid/physical-drives/{uuid}/smart", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHybridKeySecret updates HybridKeySecret
func (cli *ZSClient) UpdateHybridKeySecret(uuid string, params param.UpdateHybridKeySecretParam) (*view.HybridAccountInventoryView, error) {
	var resp view.UpdateHybridKeySecretEventView
	if err := cli.Put("v1/hybrid/hybrid/{uuid}/key", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PullHuaweiIMasterController operates on PullHuaweiIMasterController
func (cli *ZSClient) PullHuaweiIMasterController(uuid string, params param.PullHuaweiIMasterControllerParam) (*view.HuaweiIMasterSdnControllerInventoryView, error) {
	resp := view.HuaweiIMasterSdnControllerInventoryView{}
	if err := cli.Put("v1/sdn-controller/huawei-imaster/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveRolesFromIAM2VirtualIDGroup removes RolesFromIAM2VirtualIDGroup
func (cli *ZSClient) RemoveRolesFromIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups/{groupUuid}/roles", uuid, string(deleteMode))
}

// AckAlarmData operates on AlarmData
func (cli *ZSClient) AckAlarmData(params param.AckAlarmDataParam) (*view.AlertDataAckInventoryView, error) {
	var resp view.AckAlertDataEventView
	if err := cli.Post("v1/zwatch/alarm-histories/acknowledgments", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveDnsFromL3Network removes DnsFromL3Network
func (cli *ZSClient) RemoveDnsFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/{l3NetworkUuid}/dns/{dns}", uuid, string(deleteMode))
}

// ChangeIAM2OrganizationParent changes IAM2OrganizationParent
func (cli *ZSClient) ChangeIAM2OrganizationParent(uuid string, params param.ChangeIAM2OrganizationParentParam) (*view.ChangeIAM2OrganizationParentEventView, error) {
	resp := view.ChangeIAM2OrganizationParentEventView{}
	if err := cli.Put("v1/iam2/organizations/{parentUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSWeComTestConnection operates on WeComTestConnection
func (cli *ZSClient) SNSWeComTestConnection(params param.SNSWeComTestConnectionParam) (*view.SNSWeComTestConnectionEventView, error) {
	resp := view.SNSWeComTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/we-com/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProvisionVirtualRouterConfig operates on ProvisionVirtualRouterConfig
func (cli *ZSClient) ProvisionVirtualRouterConfig(uuid string, params param.ProvisionVirtualRouterConfigParam) (*view.ApplianceVmInventoryView, error) {
	var resp view.ProvisionVirtualRouterConfigEventView
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetVmQga operates on VmQga
func (cli *ZSClient) SetVmQga(uuid string, params param.SetVmQgaParam) (*view.SetVmQgaEventView, error) {
	resp := view.SetVmQgaEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidatePassword operates on Password
func (cli *ZSClient) ValidatePassword(uuid string, params param.ValidatePasswordParam) (*view.ValidatePasswordView, error) {
	resp := view.ValidatePasswordView{}
	if err := cli.Put("v1/password/verify", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetChronyServers gets ChronyServers by uuid
func (cli *ZSClient) GetChronyServers(uuid string) (*view.GetChronyServersView, error) {
	var resp view.GetChronyServersView
	if err := cli.Get("v1/zops/chrony/servers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL3NetworkToVmNic operates on L3NetworkToVmNic
func (cli *ZSClient) AttachL3NetworkToVmNic(params param.AttachL3NetworkToVmNicParam) (*view.VmNicInventoryView, error) {
	var resp view.AttachL3NetworkToVmNicEventView
	if err := cli.Post("v1/nics/{vmNicUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeSecurityMachineState changes SecurityMachineState
func (cli *ZSClient) ChangeSecurityMachineState(uuid string, params param.ChangeSecurityMachineStateParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.ChangeSecurityMachineStateEventView
	if err := cli.Put("v1/security-machines/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetVmQxlMemory operates on VmQxlMemory
func (cli *ZSClient) SetVmQxlMemory(uuid string, params param.SetVmQxlMemoryParam) (*view.SetVmQxlMemoryEventView, error) {
	resp := view.SetVmQxlMemoryEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddLocalPrimaryStorage adds LocalPrimaryStorage
func (cli *ZSClient) AddLocalPrimaryStorage(params param.AddLocalPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/local-storage", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVolumeFormat gets VolumeFormat by uuid
func (cli *ZSClient) GetVolumeFormat(uuid string) (*view.GetVolumeFormatView, error) {
	var resp view.GetVolumeFormatView
	if err := cli.Get("v1/volumes/formats", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAtPersonOfAtDingTalkEndpoint updates AtPersonOfAtDingTalkEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtDingTalkEndpoint(uuid string, params param.UpdateAtPersonOfAtDingTalkEndpointParam) (*view.SNSDingTalkAtPersonInventoryView, error) {
	var resp view.UpdateAtPersonOfDingTalkEndpointEventView
	if err := cli.Put("v1/sns/application-endpoints/ding-talk/at-persons/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateAliyunMountTarget updates AliyunMountTarget
func (cli *ZSClient) UpdateAliyunMountTarget(uuid string, params param.UpdateAliyunMountTargetParam) (*view.NasMountTargetInventoryView, error) {
	var resp view.UpdateNasMountTargetEventView
	if err := cli.Put("v1/nas/aliyun/mount", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetResourceAccount gets ResourceAccount by uuid
func (cli *ZSClient) GetResourceAccount(uuid string) (*view.StringView, error) {
	var resp view.StringView
	if err := cli.Get("v1/resources/accounts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSecretResourcePoolState changes SecretResourcePoolState
func (cli *ZSClient) ChangeSecretResourcePoolState(uuid string, params param.ChangeSecretResourcePoolStateParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.ChangeSecretResourcePoolStateEventView
	if err := cli.Put("v1/secret-resource-pools/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddSimulatorBackupStorage adds SimulatorBackupStorage
func (cli *ZSClient) AddSimulatorBackupStorage(params param.AddSimulatorBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.AddBackupStorageEventView
	if err := cli.Post("v1/backup-storage/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// BindModelToService operates on BindModelToService
func (cli *ZSClient) BindModelToService(params param.BindModelToServiceParam) (*view.ModelServiceInventoryView, error) {
	var resp view.BindModelToServiceEventView
	if err := cli.Post("v1/ai/models/{modelUuid}/model-services/{modelServiceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetCandidateAffinityGroupForCreatingVm gets CandidateAffinityGroupForCreatingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForCreatingVm(uuid string) (*view.AffinityGroupInventoryView, error) {
	var resp view.AffinityGroupInventoryView
	if err := cli.Get("v1/vm-instances/candidate-affinityGroup", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckNetworkReachable operates on NetworkReachable
func (cli *ZSClient) CheckNetworkReachable(params param.CheckNetworkReachableParam) (*view.CheckNetworkReachableView, error) {
	var resp view.CheckNetworkReachableView
	if err := cli.Get("v1/zops/check/network", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetFlowMeterRouterId operates on FlowMeterRouterId
func (cli *ZSClient) SetFlowMeterRouterId(params param.SetFlowMeterRouterIdParam) (*view.SetFlowMeterRouterIdEventView, error) {
	resp := view.SetFlowMeterRouterIdEventView{}
	if err := cli.Post("v1/flowmeters/{vRouterUuid}/routerid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddStorageProtocol adds StorageProtocol
func (cli *ZSClient) AddStorageProtocol(params param.AddStorageProtocolParam) (*view.AddStorageProtocolEventView, error) {
	resp := view.AddStorageProtocolEventView{}
	if err := cli.Post("v1/primary-storage/protocol", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployModelService operates on DeployModelService
func (cli *ZSClient) DeployModelService(uuid string, params param.DeployModelServiceParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.DeployModelServiceEventView
	if err := cli.Put("v1/ai/model-services/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetMonitorItem gets MonitorItem by uuid
func (cli *ZSClient) GetMonitorItem(uuid string) (*view.ItemInventoryView, error) {
	var resp view.ItemInventoryView
	if err := cli.Get("v1/monitoring/items", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseRecords gets LicenseRecords by uuid
func (cli *ZSClient) GetLicenseRecords(uuid string) (*view.LicenseInventoryView, error) {
	var resp view.LicenseInventoryView
	if err := cli.Get("v1/licenses/records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnregisterLicenseRequestedApplication operates on LicenseRequestedApplication
func (cli *ZSClient) UnregisterLicenseRequestedApplication(uuid string, params param.UnregisterLicenseRequestedApplicationParam) (*view.UnregisterLicenseRequestedApplicationEventView, error) {
	resp := view.UnregisterLicenseRequestedApplicationEventView{}
	if err := cli.Put("v1/license/unregister-applications", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachSecurityGroupToL3Network operates on SecurityGroupToL3Network
func (cli *ZSClient) AttachSecurityGroupToL3Network(params param.AttachSecurityGroupToL3NetworkParam) (*view.SecurityGroupInventoryView, error) {
	var resp view.AttachSecurityGroupToL3NetworkEventView
	if err := cli.Post("v1/security-groups/{securityGroupUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateVmNicDriver updates VmNicDriver
func (cli *ZSClient) UpdateVmNicDriver(uuid string, params param.UpdateVmNicDriverParam) (*view.VmNicInventoryView, error) {
	var resp view.UpdateVmNicDriverEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetIpOnHostNetworkInterface operates on IpOnHostNetworkInterface
func (cli *ZSClient) SetIpOnHostNetworkInterface(params param.SetIpOnHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	var resp view.SetIpOnHostNetworkInterfaceEventView
	if err := cli.Post("v1/hosts/nics/{interfaceUuid}/ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ProvisionNfvInstGroup operates on ProvisionNfvInstGroup
func (cli *ZSClient) ProvisionNfvInstGroup(uuid string, params param.ProvisionNfvInstGroupParam) (*view.NfvInstGroupInventoryView, error) {
	var resp view.ProvisionNfvInstGroupEventView
	if err := cli.Put("v1/nfvinstgroup/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachNicFromBonding operates on NicFromBonding
func (cli *ZSClient) DetachNicFromBonding(uuid string, params param.DetachNicFromBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	var resp view.DetachNicFromBondingEventView
	if err := cli.Put("v1/hosts/bondings/{uuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeMonitorTriggerActionState changes MonitorTriggerActionState
func (cli *ZSClient) ChangeMonitorTriggerActionState(uuid string, params param.ChangeMonitorTriggerActionStateParam) (*view.MonitorTriggerActionInventoryView, error) {
	var resp view.ChangeMonitorTriggerActionStateEventView
	if err := cli.Put("v1/monitoring/trigger-actions/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// MigrateVm operates on Vm
func (cli *ZSClient) MigrateVm(uuid string, params param.MigrateVmParam) (*view.VmInstanceInventoryView, error) {
	var resp view.MigrateVmEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeVmPassword changes VmPassword
func (cli *ZSClient) ChangeVmPassword(uuid string, params param.ChangeVmPasswordParam) (*view.ChangeVmPasswordEventView, error) {
	resp := view.ChangeVmPasswordEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVmInstance operates on FlattenVmInstance
func (cli *ZSClient) FlattenVmInstance(uuid string, params param.FlattenVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.FlattenVmInstanceEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteAllEcsInstancesFromDataCenter deletes AllEcsInstancesFromDataCenter
func (cli *ZSClient) DeleteAllEcsInstancesFromDataCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/dc-ecs/{uuid}", uuid, string(deleteMode))
}

// GetVpcMulticastRoute gets VpcMulticastRoute by uuid
func (cli *ZSClient) GetVpcMulticastRoute(uuid string) (*view.MulticastRouteInventoryView, error) {
	var resp view.MulticastRouteInventoryView
	if err := cli.Get("v1/multicast/virtual-routers/{uuid}/routes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmUserDefinedXmlHookScript deletes VmUserDefinedXmlHookScript
func (cli *ZSClient) DeleteVmUserDefinedXmlHookScript(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/xml-hook-script", uuid, string(deleteMode))
}

// AddL3NetworkToGroup adds L3NetworkToGroup
func (cli *ZSClient) AddL3NetworkToGroup(params param.AddL3NetworkToGroupParam) (*view.AddL3NetworkToGroupEventView, error) {
	resp := view.AddL3NetworkToGroupEventView{}
	if err := cli.Post("v1/nfvinstgroup/group/{nfvInstGroupUuid}/service/{networkServiceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncZBoxCapacity operates on ZBoxCapacity
func (cli *ZSClient) SyncZBoxCapacity(uuid string, params param.SyncZBoxCapacityParam) (*view.ZBoxInventoryView, error) {
	var resp view.SyncZBoxCapacityEventView
	if err := cli.Put("v1/zbox/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AckEventData operates on EventData
func (cli *ZSClient) AckEventData(params param.AckEventDataParam) (*view.AlertDataAckInventoryView, error) {
	var resp view.AckAlertDataEventView
	if err := cli.Post("v1/zwatch/event-histories/acknowledgments", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CheckResourcePermission operates on ResourcePermission
func (cli *ZSClient) CheckResourcePermission(params param.CheckResourcePermissionParam) (*view.CheckResourcePermissionView, error) {
	var resp view.CheckResourcePermissionView
	if err := cli.Get("v1/accounts/resource/api-permissions", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProvisionNfvInstConfig operates on ProvisionNfvInstConfig
func (cli *ZSClient) ProvisionNfvInstConfig(uuid string, params param.ProvisionNfvInstConfigParam) (*view.ApplianceVmInventoryView, error) {
	var resp view.ProvisionNfvInstConfigEventView
	if err := cli.Put("v1/vm-instances/appliances/nfvinst/{vmInstanceUuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetCandidateMiniHosts gets CandidateMiniHosts by uuid
func (cli *ZSClient) GetCandidateMiniHosts(uuid string) (*view.GetCandidateMiniHostsView, error) {
	var resp view.GetCandidateMiniHostsView
	if err := cli.Get("v1/mini-clusters/candidate-hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDatasets deletes Datasets
func (cli *ZSClient) DeleteDatasets(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/datasets", uuid, string(deleteMode))
}

// RevokeResourceSharing operates on RevokeResourceSharing
func (cli *ZSClient) RevokeResourceSharing(uuid string, params param.RevokeResourceSharingParam) (*view.RevokeResourceSharingEventView, error) {
	resp := view.RevokeResourceSharingEventView{}
	if err := cli.Put("v1/accounts/resources/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteModelServices deletes ModelServices
func (cli *ZSClient) DeleteModelServices(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-services/", uuid, string(deleteMode))
}

// ChangeL3NetworkState changes L3NetworkState
func (cli *ZSClient) ChangeL3NetworkState(uuid string, params param.ChangeL3NetworkStateParam) (*view.L3NetworkInventoryView, error) {
	var resp view.ChangeL3NetworkStateEventView
	if err := cli.Put("v1/l3-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetHostNUMATopology gets HostNUMATopology by uuid
func (cli *ZSClient) GetHostNUMATopology(uuid string) (*view.GetHostNUMATopologyEventView, error) {
	var resp view.GetHostNUMATopologyEventView
	if err := cli.Get("v1/hosts/{uuid}/numa", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateL2VirtualSwitch creates L2VirtualSwitch
func (cli *ZSClient) CreateL2VirtualSwitch(params param.CreateL2VirtualSwitchParam) (*view.CreateL2VirtualSwitchEventView, error) {
	resp := view.CreateL2VirtualSwitchEventView{}
	if err := cli.Post("v1/l2-networks/virtual-switch", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmNicToLoadBalancer adds VmNicToLoadBalancer
func (cli *ZSClient) AddVmNicToLoadBalancer(params param.AddVmNicToLoadBalancerParam) (*view.LoadBalancerListenerInventoryView, error) {
	var resp view.AddVmNicToLoadBalancerEventView
	if err := cli.Post("v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateBuildApp updates BuildApp
func (cli *ZSClient) UpdateBuildApp(uuid string, params param.UpdateBuildAppParam) (*view.BuildApplicationInventoryView, error) {
	var resp view.UpdateBuildAppEventView
	if err := cli.Put("v1/appcenter/buildapp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetClusterDRSStatus gets ClusterDRSStatus by uuid
func (cli *ZSClient) GetClusterDRSStatus(uuid string) (*view.GetClusterDRSStatusView, error) {
	var resp view.GetClusterDRSStatusView
	if err := cli.Get("v1/clusters/drs/status", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAliyunNasPrimaryStorage adds AliyunNasPrimaryStorage
func (cli *ZSClient) AddAliyunNasPrimaryStorage(params param.AddAliyunNasPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/aliyun/nas", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVmNuma gets VmNuma by uuid
func (cli *ZSClient) GetVmNuma(uuid string) (*view.GetVmNumaView, error) {
	var resp view.GetVmNumaView
	if err := cli.Get("v1/vm-instances/{uuid}/vnuma", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeZoneState changes ZoneState
func (cli *ZSClient) ChangeZoneState(uuid string, params param.ChangeZoneStateParam) (*view.ZoneInventoryView, error) {
	var resp view.ChangeZoneStateEventView
	if err := cli.Put("v1/zones/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachAppBuildSystemToZone operates on AppBuildSystemToZone
func (cli *ZSClient) AttachAppBuildSystemToZone(params param.AttachAppBuildSystemToZoneParam) (*view.AppBuildSystemZoneRefInventoryView, error) {
	var resp view.AttachAppBuildSystemToZoneEventView
	if err := cli.Post("v1/zones/{zoneUuid}/buildsystem/{buildSystemUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateDataVolume creates DataVolume
func (cli *ZSClient) CreateDataVolume(params param.CreateDataVolumeParam) (*view.VolumeInventoryView, error) {
	var resp view.CreateDataVolumeEventView
	if err := cli.Post("v1/volumes/data", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UngenerateHygonMdevDevices operates on UngenerateHygonMdevDevices
func (cli *ZSClient) UngenerateHygonMdevDevices(uuid string, params param.UngenerateHygonMdevDevicesParam) (*view.UngenerateHygonMdevDevicesEventView, error) {
	resp := view.UngenerateHygonMdevDevicesEventView{}
	if err := cli.Put("v1/hygon-devices/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePluginDrivers deletes PluginDrivers
func (cli *ZSClient) DeletePluginDrivers(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/external/plugins/{uuid}", uuid, string(deleteMode))
}

// BatchCreateBaremetalChassis operates on CreateBaremetalChassis
func (cli *ZSClient) BatchCreateBaremetalChassis(params param.BatchCreateBaremetalChassisParam) (*view.LongJobInventoryView, error) {
	var resp view.BatchCreateBaremetalChassisEventView
	if err := cli.Post("v1/baremetal/chassis/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddSchedulerJobToSchedulerTrigger adds SchedulerJobToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobToSchedulerTrigger(params param.AddSchedulerJobToSchedulerTriggerParam) (*view.SchedulerJobSchedulerTriggerInventoryView, error) {
	var resp view.AddSchedulerJobToSchedulerTriggerEventView
	if err := cli.Post("v1/scheduler/jobs/{schedulerJobUuid}/scheduler/triggers/{schedulerTriggerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachPolicyFromRole operates on PolicyFromRole
func (cli *ZSClient) DetachPolicyFromRole(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/policies/{policyUuid}/roles/{roleUuid}", uuid, string(deleteMode))
}

// RestartModelServiceGroups operates on RestartModelServiceGroups
func (cli *ZSClient) RestartModelServiceGroups(uuid string, params param.RestartModelServiceGroupsParam) (*view.RestartModelServiceGroupsEventView, error) {
	resp := view.RestartModelServiceGroupsEventView{}
	if err := cli.Put("v1/model-service-instance-groups", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLoadBalancerOwner gets LoadBalancerOwner by uuid
func (cli *ZSClient) GetLoadBalancerOwner(uuid string) (*view.GetLoadBalancerOwnerView, error) {
	var resp view.GetLoadBalancerOwnerView
	if err := cli.Get("v1/load-balancers/{loadBalancerUuid}/owner", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNicQos gets NicQos by uuid
func (cli *ZSClient) GetNicQos(uuid string) (*view.GetNicQosView, error) {
	var resp view.GetNicQosView
	if err := cli.Get("v1/vm-instances/{uuid}/nic-qos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmNicNetwork changes VmNicNetwork
func (cli *ZSClient) ChangeVmNicNetwork(uuid string, params param.ChangeVmNicNetworkParam) (*view.VmNicInventoryView, error) {
	var resp view.ChangeVmNicNetworkEventView
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/l3-networks/{destL3NetworkUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateBareMetal2IpmiChassisHardwareInfo creates BareMetal2IpmiChassisHardwareInfo
func (cli *ZSClient) CreateBareMetal2IpmiChassisHardwareInfo(params param.CreateBareMetal2IpmiChassisHardwareInfoParam) (*view.CreateBareMetal2ChassisHardwareView, error) {
	resp := view.CreateBareMetal2ChassisHardwareView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/hardwareinfos", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveLabelFromAlarm removes LabelFromAlarm
func (cli *ZSClient) RemoveLabelFromAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/labels/{uuid}", uuid, string(deleteMode))
}

// DeleteIAM2VirtualIDLdapBinding deletes IAM2VirtualIDLdapBinding
func (cli *ZSClient) DeleteIAM2VirtualIDLdapBinding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/ldap/bindings/{uuid}", uuid, string(deleteMode))
}

// UpdateVmPriority updates VmPriority
func (cli *ZSClient) UpdateVmPriority(uuid string, params param.UpdateVmPriorityParam) (*view.UpdateVmPriorityEventView, error) {
	resp := view.UpdateVmPriorityEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachMdevDeviceFromVm operates on MdevDeviceFromVm
func (cli *ZSClient) DetachMdevDeviceFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/mdev-devices/{mdevDeviceUuid}/vm-instances/{vmInstanceUuid}", uuid, string(deleteMode))
}

// DeleteVmHostname deletes VmHostname
func (cli *ZSClient) DeleteVmHostname(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/hostnames", uuid, string(deleteMode))
}

// GetLicenseCapabilities gets LicenseCapabilities by uuid
func (cli *ZSClient) GetLicenseCapabilities(uuid string) (*view.GetLicenseCapabilitiesView, error) {
	var resp view.GetLicenseCapabilitiesView
	if err := cli.Get("v1/licenses/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRuleTemplate creates FirewallRuleTemplate
func (cli *ZSClient) CreateFirewallRuleTemplate(params param.CreateFirewallRuleTemplateParam) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	var resp view.CreateFirewallRuleTemplateEventView
	if err := cli.Post("v1/vpcfirewalls/rules/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeIAM2ProjectState changes IAM2ProjectState
func (cli *ZSClient) ChangeIAM2ProjectState(uuid string, params param.ChangeIAM2ProjectStateParam) (*view.IAM2ProjectInventoryView, error) {
	var resp view.ChangeIAM2ProjectStateEventView
	if err := cli.Put("v1/iam2/projects/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetVmSoundType operates on VmSoundType
func (cli *ZSClient) SetVmSoundType(uuid string, params param.SetVmSoundTypeParam) (*view.SetVmSoundTypeEventView, error) {
	resp := view.SetVmSoundTypeEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MergeDataOnBackupStorage operates on MergeDataOnBackupStorage
func (cli *ZSClient) MergeDataOnBackupStorage(uuid string, params param.MergeDataOnBackupStorageParam) (*view.MergeDataOnBackupStorageEventView, error) {
	resp := view.MergeDataOnBackupStorageEventView{}
	if err := cli.Put("v1/cdp-task/mergedata/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCdpBackupStorageRequirement gets CdpBackupStorageRequirement by uuid
func (cli *ZSClient) GetCdpBackupStorageRequirement(uuid string) (*view.GetCdpBackupStorageRequirementView, error) {
	var resp view.GetCdpBackupStorageRequirementView
	if err := cli.Get("v1/cdp-backup-storage/{backupStorageUuid}/requirement", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAttributesToIAM2VirtualIDGroup adds AttributesToIAM2VirtualIDGroup
func (cli *ZSClient) AddAttributesToIAM2VirtualIDGroup(params param.AddAttributesToIAM2VirtualIDGroupParam) (*view.AddAttributesToIAM2VirtualIDGroupEventView, error) {
	resp := view.AddAttributesToIAM2VirtualIDGroupEventView{}
	if err := cli.Post("v1/iam2/projects/groups/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAffinityGroupState changes AffinityGroupState
func (cli *ZSClient) ChangeAffinityGroupState(uuid string, params param.ChangeAffinityGroupStateParam) (*view.AffinityGroupInventoryView, error) {
	var resp view.ChangeAffinityGroupStateEventView
	if err := cli.Put("v1/affinity-groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeSecurityGroupRuleState changes SecurityGroupRuleState
func (cli *ZSClient) ChangeSecurityGroupRuleState(uuid string, params param.ChangeSecurityGroupRuleStateParam) (*view.SecurityGroupInventoryView, error) {
	var resp view.ChangeSecurityGroupRuleStateEventView
	if err := cli.Put("v1/security-groups/{securityGroupUuid}/rules/state/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddVmNicToSecurityGroup adds VmNicToSecurityGroup
func (cli *ZSClient) AddVmNicToSecurityGroup(params param.AddVmNicToSecurityGroupParam) (*view.AddVmNicToSecurityGroupEventView, error) {
	resp := view.AddVmNicToSecurityGroupEventView{}
	if err := cli.Post("v1/security-groups/{securityGroupUuid}/vm-instances/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAliyunRouteEntryFromRemote operates on AliyunRouteEntryFromRemote
func (cli *ZSClient) SyncAliyunRouteEntryFromRemote(uuid string, params param.SyncAliyunRouteEntryFromRemoteParam) (*view.VpcVirtualRouteEntryInventoryView, error) {
	resp := view.VpcVirtualRouteEntryInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/route-entry/{vRouterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEmailAddressOfSNSEmailEndpoint updates EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) UpdateEmailAddressOfSNSEmailEndpoint(uuid string, params param.UpdateEmailAddressOfSNSEmailEndpointParam) (*view.SNSEmailAddressInventoryView, error) {
	var resp view.UpdateEmailAddressOfSNSEmailEndpointEventView
	if err := cli.Put("v1/sns/application-endpoints/emails/email-addresses", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetMetricLabelValue gets MetricLabelValue by uuid
func (cli *ZSClient) GetMetricLabelValue(uuid string) (*view.GetMetricLabelValueView, error) {
	var resp view.GetMetricLabelValueView
	if err := cli.Get("v1/zwatch/metrics/label-values", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateZonesClustersHostsForCreatingVm gets CandidateZonesClustersHostsForCreatingVm by uuid
func (cli *ZSClient) GetCandidateZonesClustersHostsForCreatingVm(uuid string) (*view.GetCandidateZonesClustersHostsForCreatingVmView, error) {
	var resp view.GetCandidateZonesClustersHostsForCreatingVmView
	if err := cli.Get("v1/vm-instances/candidate-destinations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateResourcePrice creates ResourcePrice
func (cli *ZSClient) CreateResourcePrice(params param.CreateResourcePriceParam) (*view.PriceInventoryView, error) {
	var resp view.CreateResourcePriceEventView
	if err := cli.Post("v1/billings/prices", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveSchedulerJobGroupFromSchedulerTrigger removes SchedulerJobGroupFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobGroupFromSchedulerTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobgroups/{schedulerJobGroupUuid}/scheduler/triggers/{schedulerTriggerUuid}", uuid, string(deleteMode))
}

// ChangeAccountPriceTableBinding changes AccountPriceTableBinding
func (cli *ZSClient) ChangeAccountPriceTableBinding(uuid string, params param.ChangeAccountPriceTableBindingParam) (*view.PriceTableInventoryView, error) {
	var resp view.ChangeAccountPriceTableBindingEventView
	if err := cli.Put("v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteVolumeQos deletes VolumeQos
func (cli *ZSClient) DeleteVolumeQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{uuid}/qos", uuid, string(deleteMode))
}

// GetL3NetworkDhcpIpAddress gets L3NetworkDhcpIpAddress by uuid
func (cli *ZSClient) GetL3NetworkDhcpIpAddress(uuid string) (*view.GetL3NetworkDhcpIpAddressView, error) {
	var resp view.GetL3NetworkDhcpIpAddressView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/dhcp-ip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRuleSet creates FirewallRuleSet
func (cli *ZSClient) CreateFirewallRuleSet(params param.CreateFirewallRuleSetParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	var resp view.CreateFirewallRuleSetEventView
	if err := cli.Post("v1/vpcfirewalls/ruleSets", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateJitSecretResourcePool creates JitSecretResourcePool
func (cli *ZSClient) CreateJitSecretResourcePool(params param.CreateJitSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.CreateSecretResourcePoolEventView
	if err := cli.Post("v1/secret-resource-pool/jit", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetBaremetalChassisPowerStatus gets BaremetalChassisPowerStatus by uuid
func (cli *ZSClient) GetBaremetalChassisPowerStatus(uuid string) (*view.GetBaremetalChassisPowerStatusView, error) {
	var resp view.GetBaremetalChassisPowerStatusView
	if err := cli.Get("v1/baremetal/chassis/{uuid}/powerstatus", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVmUserDefinedXmlHookScript queries VmUserDefinedXmlHookScript list
func (cli *ZSClient) QueryVmUserDefinedXmlHookScript(params *param.QueryParam) ([]view.XmlHookInventoryView, error) {
	var resp []view.XmlHookInventoryView
	return resp, cli.List("v1/vm-instances/xml-hook-script", params, &resp)
}

// RefreshFirewall operates on Firewall
func (cli *ZSClient) RefreshFirewall(uuid string, params param.RefreshFirewallParam) (*view.VpcFirewallInventoryView, error) {
	var resp view.RefreshFirewallEventView
	if err := cli.Put("v1/vpcfirewalls/refresh/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachL3NetworksFromIPsecConnection operates on L3NetworksFromIPsecConnection
func (cli *ZSClient) DetachL3NetworksFromIPsecConnection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ipsec/{uuid}/l3networks", uuid, string(deleteMode))
}

// UpdateAutoScalingGroupAddingNewInstanceRule updates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupAddingNewInstanceRule(uuid string, params param.UpdateAutoScalingGroupAddingNewInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	var resp view.UpdateAutoScalingRuleEventView
	if err := cli.Put("v1/autoscaling/rules/adding-new-instance/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetFaultToleranceVms gets FaultToleranceVms by uuid
func (cli *ZSClient) GetFaultToleranceVms(uuid string) (*view.GetFaultToleranceVmsView, error) {
	var resp view.GetFaultToleranceVmsView
	if err := cli.Get("v1/vm-instances/fault-tolerance/sub-vms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunKeySecret deletes AliyunKeySecret
func (cli *ZSClient) DeleteAliyunKeySecret(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/key/{uuid}", uuid, string(deleteMode))
}

// CreateVmInstanceFromVolumeSnapshot creates VmInstanceFromVolumeSnapshot
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshot(params param.CreateVmInstanceFromVolumeSnapshotParam) (*view.VmInstanceInventoryView, error) {
	var resp view.CreateVmInstanceFromVolumeSnapshotEventView
	if err := cli.Post("v1/vm-instances/from/volume-snapshots/{volumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PowerResetBareMetal2Chassis operates on PowerResetBareMetal2Chassis
func (cli *ZSClient) PowerResetBareMetal2Chassis(uuid string, params param.PowerResetBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.PowerResetBareMetal2ChassisEventView
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PrometheusQueryVmMonitoringData operates on PrometheusQueryVmMonitoringData
func (cli *ZSClient) PrometheusQueryVmMonitoringData(params param.PrometheusQueryVmMonitoringDataParam) (*view.PrometheusQueryVmMonitoringDataView, error) {
	var resp view.PrometheusQueryVmMonitoringDataView
	if err := cli.Get("v1/prometheus/vm-instances", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateResourceConfigs updates ResourceConfigs
func (cli *ZSClient) UpdateResourceConfigs(uuid string, params param.UpdateResourceConfigsParam) (*view.ResourceConfigStructView, error) {
	resp := view.ResourceConfigStructView{}
	if err := cli.Put("v1/resource-configurations/{resourceUuid}/resource-configs/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LogInByAccount operates on LogInByAccount
func (cli *ZSClient) LogInByAccount(uuid string, params param.LogInByAccountParam) (*view.SessionInventoryView, error) {
	var resp view.LogInView
	if err := cli.Put("v1/accounts/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryVpcUserVpnGatewayFromLocal queries VpcUserVpnGatewayFromLocal list
func (cli *ZSClient) QueryVpcUserVpnGatewayFromLocal(params *param.QueryParam) ([]view.VpcUserVpnGatewayInventoryView, error) {
	var resp []view.VpcUserVpnGatewayInventoryView
	return resp, cli.List("v1/hybrid/user-vpn", params, &resp)
}

// RevertVolumeFromSnapshot operates on VolumeFromSnapshot
func (cli *ZSClient) RevertVolumeFromSnapshot(uuid string, params param.RevertVolumeFromSnapshotParam) (*view.RevertVolumeFromSnapshotEventView, error) {
	resp := view.RevertVolumeFromSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBlockPrimaryStorageMetadata gets BlockPrimaryStorageMetadata by uuid
func (cli *ZSClient) GetBlockPrimaryStorageMetadata(uuid string) (*view.BlockPrimaryStorageInventoryView, error) {
	var resp view.BlockPrimaryStorageInventoryView
	if err := cli.Get("v1/primary-storage/block/metadata", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBonding updates Bonding
func (cli *ZSClient) UpdateBonding(uuid string, params param.UpdateBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	var resp view.UpdateBondingEventView
	if err := cli.Put("v1/hosts/bondings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetManagementNodeArch gets ManagementNodeArch by uuid
func (cli *ZSClient) GetManagementNodeArch(uuid string) (*view.GetManagementNodeArchView, error) {
	var resp view.GetManagementNodeArchView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachScsiLunFromHost operates on ScsiLunFromHost
func (cli *ZSClient) DetachScsiLunFromHost(uuid string, params param.DetachScsiLunFromHostParam) (*view.ScsiLunInventoryView, error) {
	var resp view.DetachScsiLunFromHostEventView
	if err := cli.Put("v1/storage-devices/scsi-lun/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DisableCbtTask operates on DisableCbtTask
func (cli *ZSClient) DisableCbtTask(params param.DisableCbtTaskParam) (*view.CbtTaskInventoryView, error) {
	var resp view.DisableCbtTaskEventView
	if err := cli.Post("v1/cbt-task/disable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RefreshLocalRaid operates on LocalRaid
func (cli *ZSClient) RefreshLocalRaid(uuid string, params param.RefreshLocalRaidParam) (*view.RaidControllerInventoryView, error) {
	resp := view.RaidControllerInventoryView{}
	if err := cli.Put("v1/storage-devices/local-raid/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSubscribeEvent updates SubscribeEvent
func (cli *ZSClient) UpdateSubscribeEvent(uuid string, params param.UpdateSubscribeEventParam) (*view.EventSubscriptionInventoryView, error) {
	var resp view.UpdateSubscribeEventEventView
	if err := cli.Put("v1/zwatch/events/subscriptions/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetVmSshKey operates on VmSshKey
func (cli *ZSClient) SetVmSshKey(uuid string, params param.SetVmSshKeyParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmSshKeyEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// FailoverFaultToleranceVm operates on FailoverFaultToleranceVm
func (cli *ZSClient) FailoverFaultToleranceVm(uuid string, params param.FailoverFaultToleranceVmParam) (*view.FailoverFaultToleranceVmEventView, error) {
	resp := view.FailoverFaultToleranceVmEventView{}
	if err := cli.Put("v1/vm-instances/fault-tolerance", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EjectZBox operates on EjectZBox
func (cli *ZSClient) EjectZBox(uuid string, params param.EjectZBoxParam) (*view.ZBoxInventoryView, error) {
	var resp view.EjectZBoxEventView
	if err := cli.Put("v1/zbox/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PrometheusQueryMetadata operates on PrometheusQueryMetadata
func (cli *ZSClient) PrometheusQueryMetadata(params param.PrometheusQueryMetadataParam) (*view.PrometheusQueryMetadataView, error) {
	var resp view.PrometheusQueryMetadataView
	if err := cli.Get("v1/prometheus/meta-data", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIpAddress queries IpAddress list
func (cli *ZSClient) QueryIpAddress(params *param.QueryParam) ([]view.UsedIpInventoryView, error) {
	var resp []view.UsedIpInventoryView
	return resp, cli.List("v1/l3-networks/ip-address", params, &resp)
}

// DeleteFirewallRuleTemplate deletes FirewallRuleTemplate
func (cli *ZSClient) DeleteFirewallRuleTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/rules/templates/{uuid}", uuid, string(deleteMode))
}

// DetachPciDeviceFromVm operates on PciDeviceFromVm
func (cli *ZSClient) DetachPciDeviceFromVm(params param.DetachPciDeviceFromVmParam) (*view.PciDeviceInventoryView, error) {
	var resp view.DetachPciDeviceFromVmEventView
	if err := cli.Post("v1/pci-device/pci-devices/{pciDeviceUuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ExecuteGuestVmCommand operates on ExecuteGuestVmCommand
func (cli *ZSClient) ExecuteGuestVmCommand(params param.ExecuteGuestVmCommandParam) (*view.ExecuteGuestVmCommandEventView, error) {
	resp := view.ExecuteGuestVmCommandEventView{}
	if err := cli.Post("v1/vm-instances/commands/exec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVpcSharedQosBandwidth changes VpcSharedQosBandwidth
func (cli *ZSClient) ChangeVpcSharedQosBandwidth(uuid string, params param.ChangeVpcSharedQosBandwidthParam) (*view.VpcSharedQosInventoryView, error) {
	var resp view.ChangeVpcSharedQosBandwidthEventView
	if err := cli.Put("v1/vips/sharedqos/{sharedQosUuid}/bandwidth/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveAttributesFromIAM2VirtualIDGroup removes AttributesFromIAM2VirtualIDGroup
func (cli *ZSClient) RemoveAttributesFromIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups/{uuid}/attributes", uuid, string(deleteMode))
}

// AddAttributesToIAM2VirtualID adds AttributesToIAM2VirtualID
func (cli *ZSClient) AddAttributesToIAM2VirtualID(params param.AddAttributesToIAM2VirtualIDParam) (*view.AddAttributesToIAM2VirtualIDEventView, error) {
	resp := view.AddAttributesToIAM2VirtualIDEventView{}
	if err := cli.Post("v1/iam2/virtual-ids/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVolume operates on FlattenVolume
func (cli *ZSClient) FlattenVolume(uuid string, params param.FlattenVolumeParam) (*view.VolumeInventoryView, error) {
	var resp view.FlattenVolumeEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateAliyunDiskFromRemote creates AliyunDiskFromRemote
func (cli *ZSClient) CreateAliyunDiskFromRemote(params param.CreateAliyunDiskFromRemoteParam) (*view.AliyunDiskInventoryView, error) {
	var resp view.CreateAliyunDiskFromRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/disk", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteEcsSecurityGroupRuleRemote deletes EcsSecurityGroupRuleRemote
func (cli *ZSClient) DeleteEcsSecurityGroupRuleRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/security-group-rule/remote/{uuid}", uuid, string(deleteMode))
}

// DetachAliyunDiskFromEcs operates on AliyunDiskFromEcs
func (cli *ZSClient) DetachAliyunDiskFromEcs(params param.DetachAliyunDiskFromEcsParam) (*view.DetachAliyunDiskFromEcsEventView, error) {
	resp := view.DetachAliyunDiskFromEcsEventView{}
	if err := cli.Post("v1/hybrid/aliyun/disk/{uuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateAffinityGroupForAttachingVm gets CandidateAffinityGroupForAttachingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForAttachingVm(uuid string) (*view.AffinityGroupInventoryView, error) {
	var resp view.AffinityGroupInventoryView
	if err := cli.Get("v1/affinityGroup/attachingVm", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFirewallIpSetTemplate updates FirewallIpSetTemplate
func (cli *ZSClient) UpdateFirewallIpSetTemplate(uuid string, params param.UpdateFirewallIpSetTemplateParam) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	var resp view.UpdateFirewallIpSetTemplateEventView
	if err := cli.Put("v1/vpcfirewalls/ipset/templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddAccessControlListRedirectRule adds AccessControlListRedirectRule
func (cli *ZSClient) AddAccessControlListRedirectRule(params param.AddAccessControlListRedirectRuleParam) (*view.AccessControlListEntryInventoryView, error) {
	var resp view.AddAccessControlListEntryEventView
	if err := cli.Post("v1/access-control-lists/{aclUuid}/redirectRules", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachHostFromHostSchedulingRuleGroup operates on HostFromHostSchedulingRuleGroup
func (cli *ZSClient) DetachHostFromHostSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hostSchedulingRuleGroup/{hostGroupUuid}/host", uuid, string(deleteMode))
}

// UpdateAliyunRouteInterfaceRemote updates AliyunRouteInterfaceRemote
func (cli *ZSClient) UpdateAliyunRouteInterfaceRemote(uuid string, params param.UpdateAliyunRouteInterfaceRemoteParam) (*view.UpdateAliyunRouteInterfaceRemoteEventView, error) {
	resp := view.UpdateAliyunRouteInterfaceRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/router-interface/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceSpecCandidates gets PciDeviceSpecCandidates by uuid
func (cli *ZSClient) GetPciDeviceSpecCandidates(uuid string) (*view.PciDeviceSpecInventoryView, error) {
	var resp view.PciDeviceSpecInventoryView
	if err := cli.Get("v1/pci-device-specs/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryPassThrough operates on PrometheusQueryPassThrough
func (cli *ZSClient) PrometheusQueryPassThrough(params param.PrometheusQueryPassThroughParam) (*view.PrometheusQueryPassThroughView, error) {
	var resp view.PrometheusQueryPassThroughView
	if err := cli.Get("v1/prometheus/all", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVmNicToVm operates on VmNicToVm
func (cli *ZSClient) AttachVmNicToVm(params param.AttachVmNicToVmParam) (*view.VmInstanceInventoryView, error) {
	var resp view.AttachVmNicToVmEventView
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/nices/{vmNicUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveMonFromCephBackupStorage removes MonFromCephBackupStorage
func (cli *ZSClient) RemoveMonFromCephBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/backup-storage/ceph/{uuid}/mons", uuid, string(deleteMode))
}

// GetVmDeviceAddress gets VmDeviceAddress by uuid
func (cli *ZSClient) GetVmDeviceAddress(uuid string) (*view.GetVmDeviceAddressView, error) {
	var resp view.GetVmDeviceAddressView
	if err := cli.Get("v1/vm-instances/devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveInstanceFromMonitorGroup removes InstanceFromMonitorGroup
func (cli *ZSClient) RemoveInstanceFromMonitorGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitorgroups/{groupUuid}/actions/{instanceUuid}", uuid, string(deleteMode))
}

// CleanQueue operates on Queue
func (cli *ZSClient) CleanQueue(uuid string, params param.CleanQueueParam) (*view.CleanQueueEventView, error) {
	resp := view.CleanQueueEventView{}
	if err := cli.Put("v1/clean/queue", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveAccessControlListFromLoadBalancer removes AccessControlListFromLoadBalancer
func (cli *ZSClient) RemoveAccessControlListFromLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/access-control-lists", uuid, string(deleteMode))
}

// RemoveLabelFromEventSubscription removes LabelFromEventSubscription
func (cli *ZSClient) RemoveLabelFromEventSubscription(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions/labels/{uuid}", uuid, string(deleteMode))
}

// SdnControllerRemoveHost operates on SdnControllerRemoveHost
func (cli *ZSClient) SdnControllerRemoveHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}", uuid, string(deleteMode))
}

// DetachCCSCertificateFromUser operates on CCSCertificateFromUser
func (cli *ZSClient) DetachCCSCertificateFromUser(params param.DetachCCSCertificateFromUserParam) (*view.DetachCCSCertificateFromUserEventView, error) {
	resp := view.DetachCCSCertificateFromUserEventView{}
	if err := cli.Post("v1/crypto/ccs-certificate/detach-user/{userUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetManagementNodeOS gets ManagementNodeOS by uuid
func (cli *ZSClient) GetManagementNodeOS(uuid string) (*view.GetManagementNodeOSView, error) {
	var resp view.GetManagementNodeOSView
	if err := cli.Get("v1/management/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateLdapBinding creates LdapBinding
func (cli *ZSClient) CreateLdapBinding(params param.CreateLdapBindingParam) (*view.LdapAccountRefInventoryView, error) {
	var resp view.CreateLdapBindingEventView
	if err := cli.Post("v1/ldap/bindings", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ExecuteDRSScheduling operates on ExecuteDRSScheduling
func (cli *ZSClient) ExecuteDRSScheduling(uuid string, params param.ExecuteDRSSchedulingParam) (*view.ExecuteDRSSchedulingEventView, error) {
	resp := view.ExecuteDRSSchedulingEventView{}
	if err := cli.Put("v1/clusters/drs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsSecurityGroupRuleRemote creates EcsSecurityGroupRuleRemote
func (cli *ZSClient) CreateEcsSecurityGroupRuleRemote(params param.CreateEcsSecurityGroupRuleRemoteParam) (*view.EcsSecurityGroupRuleInventoryView, error) {
	var resp view.CreateEcsSecurityGroupRuleRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/security-group-rule", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVmQga gets VmQga by uuid
func (cli *ZSClient) GetVmQga(uuid string) (*view.GetVmQgaView, error) {
	var resp view.GetVmQgaView
	if err := cli.Get("v1/vm-instances/{uuid}/qga", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PreviewResourceStack operates on PreviewResourceStack
func (cli *ZSClient) PreviewResourceStack(params param.PreviewResourceStackParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post("v1/cloudformation/stack/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmvNUMATopology gets VmvNUMATopology by uuid
func (cli *ZSClient) GetVmvNUMATopology(uuid string) (*view.GetVmvNUMATopologyView, error) {
	var resp view.GetVmvNUMATopologyView
	if err := cli.Get("v1/vm-instances/{uuid}/vnuma-topology", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveSchedulerJobsFromSchedulerJobGroup removes SchedulerJobsFromSchedulerJobGroup
func (cli *ZSClient) RemoveSchedulerJobsFromSchedulerJobGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scheduler/jobgroups/{schedulerJobGroupUuid}/job", uuid, string(deleteMode))
}

// ChangeTicketStatus changes TicketStatus
func (cli *ZSClient) ChangeTicketStatus(uuid string, params param.ChangeTicketStatusParam) (*view.TicketInventoryView, error) {
	var resp view.ChangeTicketStatusEventView
	if err := cli.Put("v1/tickets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetHostPhysicalMemoryFacts gets HostPhysicalMemoryFacts by uuid
func (cli *ZSClient) GetHostPhysicalMemoryFacts(uuid string) (*view.HostPhysicalMemoryInventoryView, error) {
	var resp view.HostPhysicalMemoryInventoryView
	if err := cli.Get("v1/hosts/physical-memory-facts/{hostUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseInfo gets LicenseInfo by uuid
func (cli *ZSClient) GetLicenseInfo(uuid string) (*view.LicenseInventoryView, error) {
	var resp view.GetLicenseInfoView
	if err := cli.Get("v1/licenses", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeSchedulerState changes SchedulerState
func (cli *ZSClient) ChangeSchedulerState(uuid string, params param.ChangeSchedulerStateParam) (*view.SchedulerJobInventoryView, error) {
	var resp view.ChangeSchedulerStateEventView
	if err := cli.Put("v1/schedulers/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachPriceTableToAccount operates on PriceTableToAccount
func (cli *ZSClient) AttachPriceTableToAccount(params param.AttachPriceTableToAccountParam) (*view.PriceTableInventoryView, error) {
	var resp view.AttachPriceTableToAccountEventView
	if err := cli.Post("v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GenerateMdevDevices operates on MdevDevices
func (cli *ZSClient) GenerateMdevDevices(uuid string, params param.GenerateMdevDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PreviewResourceFromApp operates on PreviewResourceFromApp
func (cli *ZSClient) PreviewResourceFromApp(params param.PreviewResourceFromAppParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post("v1/appcenter/app/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSNSTopicState changes SNSTopicState
func (cli *ZSClient) ChangeSNSTopicState(uuid string, params param.ChangeSNSTopicStateParam) (*view.SNSTopicInventoryView, error) {
	var resp view.ChangeSNSTopicStateEventView
	if err := cli.Put("v1/zwatch/topics/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachScsiLunToVmInstance operates on ScsiLunToVmInstance
func (cli *ZSClient) AttachScsiLunToVmInstance(params param.AttachScsiLunToVmInstanceParam) (*view.ScsiLunInventoryView, error) {
	var resp view.AttachScsiLunToVmInstanceEventView
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/scsi-lun/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveRemoteCidrsFromIPsecConnection removes RemoteCidrsFromIPsecConnection
func (cli *ZSClient) RemoveRemoteCidrsFromIPsecConnection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ipsec/{uuid}/remote-cidrs", uuid, string(deleteMode))
}

// GetIAM2ProjectsOfVirtualID gets IAM2ProjectsOfVirtualID by uuid
func (cli *ZSClient) GetIAM2ProjectsOfVirtualID(uuid string) (*view.IAM2ProjectInventoryView, error) {
	var resp view.IAM2ProjectInventoryView
	if err := cli.Get("v1/iam2/virtual-ids/projects", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmToAffinityGroup adds VmToAffinityGroup
func (cli *ZSClient) AddVmToAffinityGroup(params param.AddVmToAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	var resp view.AddVmToAffinityGroupEventView
	if err := cli.Post("v1/affinity-groups/{affinityGroupUuid}/vm-instances/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetPrimaryStorageAllocatorStrategies gets PrimaryStorageAllocatorStrategies by uuid
func (cli *ZSClient) GetPrimaryStorageAllocatorStrategies(uuid string) (*view.GetPrimaryStorageAllocatorStrategiesView, error) {
	var resp view.GetPrimaryStorageAllocatorStrategiesView
	if err := cli.Get("v1/primary-storage/allocators/strategies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPlatformTimeZone gets PlatformTimeZone by uuid
func (cli *ZSClient) GetPlatformTimeZone(uuid string) (*view.GetPlatformTimeZoneView, error) {
	var resp view.GetPlatformTimeZoneView
	if err := cli.Get("v1/management-nodes/platform-timezone", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPolicyFromUser operates on PolicyFromUser
func (cli *ZSClient) DetachPolicyFromUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/users/{userUuid}/policies/{policyUuid}", uuid, string(deleteMode))
}

// SetVmInstanceDefaultCdRom operates on VmInstanceDefaultCdRom
func (cli *ZSClient) SetVmInstanceDefaultCdRom(uuid string, params param.SetVmInstanceDefaultCdRomParam) (*view.VmCdRomInventoryView, error) {
	var resp view.SetVmInstanceDefaultCdRomEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/cdroms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RefreshSharedblockDeviceCapacity operates on SharedblockDeviceCapacity
func (cli *ZSClient) RefreshSharedblockDeviceCapacity(params param.RefreshSharedblockDeviceCapacityParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	var resp view.RefreshSharedBlockDeviceCapacityEventView
	if err := cli.Post("v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// FstrimVm operates on FstrimVm
func (cli *ZSClient) FstrimVm(params param.FstrimVmParam) (*view.FstrimVmEventView, error) {
	resp := view.FstrimVmEventView{}
	if err := cli.Post("v1/vm-instances/{uuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL2NetworkFromCluster operates on L2NetworkFromCluster
func (cli *ZSClient) DetachL2NetworkFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}", uuid, string(deleteMode))
}

// SyncAINginxConfiguration operates on AINginxConfiguration
func (cli *ZSClient) SyncAINginxConfiguration(params param.SyncAINginxConfigurationParam) (*view.SyncAINginxConfigurationView, error) {
	resp := view.SyncAINginxConfigurationView{}
	if err := cli.Post("v1/ai/nginx/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MatchModelServiceTemplateWithModel operates on MatchModelServiceTemplateWithModel
func (cli *ZSClient) MatchModelServiceTemplateWithModel(params param.MatchModelServiceTemplateWithModelParam) (*view.MatchModelServiceTemplateWithModelEventView, error) {
	resp := view.MatchModelServiceTemplateWithModelEventView{}
	if err := cli.Post("v1/ai/model-services/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmNicState changes VmNicState
func (cli *ZSClient) ChangeVmNicState(uuid string, params param.ChangeVmNicStateParam) (*view.VmInstanceInventoryView, error) {
	var resp view.ChangeVmNicStateEventView
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddPolicyStatementsToRole adds PolicyStatementsToRole
func (cli *ZSClient) AddPolicyStatementsToRole(params param.AddPolicyStatementsToRoleParam) (*view.AddPolicyStatementsToRoleEventView, error) {
	resp := view.AddPolicyStatementsToRoleEventView{}
	if err := cli.Post("v1/identities/roles/{uuid}/policy-statements", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnprotectVmInstanceRecoveryPoint operates on UnprotectVmInstanceRecoveryPoint
func (cli *ZSClient) UnprotectVmInstanceRecoveryPoint(uuid string, params param.UnprotectVmInstanceRecoveryPointParam) (*view.UnprotectVmInstanceRecoveryPointEventView, error) {
	resp := view.UnprotectVmInstanceRecoveryPointEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/unprotect-recovery-point", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachVipFromVpcSharedQos operates on VipFromVpcSharedQos
func (cli *ZSClient) DetachVipFromVpcSharedQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips/sharedqos/{sharedQosUuid}/vips", uuid, string(deleteMode))
}

// ApplyRuleSetChanges operates on RuleSetChanges
func (cli *ZSClient) ApplyRuleSetChanges(uuid string, params param.ApplyRuleSetChangesParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	var resp view.ApplyRuleSetChangesEventView
	if err := cli.Put("v1/vpcfirewalls/ruleSets/apply/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// PrimaryStorageMigrateVm operates on PrimaryStorageMigrateVm
func (cli *ZSClient) PrimaryStorageMigrateVm(uuid string, params param.PrimaryStorageMigrateVmParam) (*view.VmInstanceInventoryView, error) {
	var resp view.PrimaryStorageMigrateVmEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RecoverDatabaseFromBackup operates on DatabaseFromBackup
func (cli *ZSClient) RecoverDatabaseFromBackup(uuid string, params param.RecoverDatabaseFromBackupParam) (*view.RecoverDatabaseFromBackupEventView, error) {
	resp := view.RecoverDatabaseFromBackupEventView{}
	if err := cli.Put("v1/database-backups/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2TickFlowCollection creates IAM2TickFlowCollection
func (cli *ZSClient) CreateIAM2TickFlowCollection(params param.CreateIAM2TickFlowCollectionParam) (*view.TicketFlowCollectionInventoryView, error) {
	var resp view.CreateTickFlowCollectionEventView
	if err := cli.Post("v1/tickets/flow-collections", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UngenerateMdevDevices operates on UngenerateMdevDevices
func (cli *ZSClient) UngenerateMdevDevices(uuid string, params param.UngenerateMdevDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveDirectory operates on MoveDirectory
func (cli *ZSClient) MoveDirectory(uuid string, params param.MoveDirectoryParam) (*view.MoveDirectoryEventView, error) {
	resp := view.MoveDirectoryEventView{}
	if err := cli.Put("v1/move/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVRouterOspfNeighbor gets VRouterOspfNeighbor by uuid
func (cli *ZSClient) GetVRouterOspfNeighbor(uuid string) (*view.GetVRouterOspfNeighborView, error) {
	var resp view.GetVRouterOspfNeighborView
	if err := cli.Get("v1/routerArea/{vRouterUuid}/neighbor", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpcVRouter creates VpcVRouter
func (cli *ZSClient) CreateVpcVRouter(params param.CreateVpcVRouterParam) (*view.VmInstanceInventoryView, error) {
	var resp view.CreateVpcVRouterEventView
	if err := cli.Post("v1/vpc/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SyncEcsInstanceFromRemote operates on EcsInstanceFromRemote
func (cli *ZSClient) SyncEcsInstanceFromRemote(params param.SyncEcsInstanceFromRemoteParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/ecs/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMdevDeviceSpecCandidates gets MdevDeviceSpecCandidates by uuid
func (cli *ZSClient) GetMdevDeviceSpecCandidates(uuid string) (*view.MdevDeviceSpecInventoryView, error) {
	var resp view.MdevDeviceSpecInventoryView
	if err := cli.Get("v1/mdev-device-specs/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFlowMeterRouterId gets FlowMeterRouterId by uuid
func (cli *ZSClient) GetFlowMeterRouterId(uuid string) (*view.GetFlowMeterRouterIdView, error) {
	var resp view.GetFlowMeterRouterIdView
	if err := cli.Get("v1/flowmeters/{vRouterUuid}/routerid", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceCandidatesForNewCreateVm gets PciDeviceCandidatesForNewCreateVm by uuid
func (cli *ZSClient) GetPciDeviceCandidatesForNewCreateVm(uuid string) (*view.PciDeviceInventoryView, error) {
	var resp view.PciDeviceInventoryView
	if err := cli.Get("v1/pci-device/candidate-pci-devices-for-new-create-vm", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryDataCenterFromLocal queries DataCenterFromLocal list
func (cli *ZSClient) QueryDataCenterFromLocal(params *param.QueryParam) ([]view.DataCenterInventoryView, error) {
	var resp []view.DataCenterInventoryView
	return resp, cli.List("v1/hybrid/data-center", params, &resp)
}

// GetHostTask gets HostTask by uuid
func (cli *ZSClient) GetHostTask(uuid string) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.Get("v1/hosts/task-details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddResourceToIAM2Project adds ResourceToIAM2Project
func (cli *ZSClient) AddResourceToIAM2Project(params param.AddResourceToIAM2ProjectParam) (*view.AddResourceToIAM2ProjectEventView, error) {
	resp := view.AddResourceToIAM2ProjectEventView{}
	if err := cli.Post("v1/iam2/projects/add/resource/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAlarmData gets AlarmData by uuid
func (cli *ZSClient) GetAlarmData(uuid string) (*view.GetAlarmDataView, error) {
	var resp view.GetAlarmDataView
	if err := cli.Get("v1/zwatch/alarm-histories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeV2VConversionHostState changes V2VConversionHostState
func (cli *ZSClient) ChangeV2VConversionHostState(uuid string, params param.ChangeV2VConversionHostStateParam) (*view.V2VConversionHostInventoryView, error) {
	var resp view.ChangeV2VConversionHostStateEventView
	if err := cli.Put("v1/v2v-conversion-hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RecoverResourceSplitBrain operates on ResourceSplitBrain
func (cli *ZSClient) RecoverResourceSplitBrain(uuid string, params param.RecoverResourceSplitBrainParam) (*view.RecoverResourceSplitBrainEventView, error) {
	resp := view.RecoverResourceSplitBrainEventView{}
	if err := cli.Put("v1/primary-storage/mini/{resourceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsOpensourceVersion operates on IsOpensourceVersion
func (cli *ZSClient) IsOpensourceVersion(params param.IsOpensourceVersionParam) (*view.IsOpensourceVersionView, error) {
	var resp view.IsOpensourceVersionView
	if err := cli.Get("v1/meta-data/opensource", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsInstanceFromEcsImage creates EcsInstanceFromEcsImage
func (cli *ZSClient) CreateEcsInstanceFromEcsImage(params param.CreateEcsInstanceFromEcsImageParam) (*view.EcsInstanceInventoryView, error) {
	var resp view.CreateEcsInstanceFromEcsImageEventView
	if err := cli.Post("v1/hybrid/aliyun/ecs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetResourceFromResourceStack gets ResourceFromResourceStack by uuid
func (cli *ZSClient) GetResourceFromResourceStack(uuid string) (*view.GetResourceFromResourceStackView, error) {
	var resp view.GetResourceFromResourceStackView
	if err := cli.Get("v1/cloudformation/stack/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveResourcesToDirectory operates on MoveResourcesToDirectory
func (cli *ZSClient) MoveResourcesToDirectory(uuid string, params param.MoveResourcesToDirectoryParam) (*view.MoveResourcesToDirectoryEventView, error) {
	resp := view.MoveResourcesToDirectoryEventView{}
	if err := cli.Put("v1/move/resources/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSupportedCloudFormationResources gets SupportedCloudFormationResources by uuid
func (cli *ZSClient) GetSupportedCloudFormationResources(uuid string) (*view.GetSupportedCloudFormationResourcesView, error) {
	var resp view.GetSupportedCloudFormationResourcesView
	if err := cli.Get("v1/cloudformation/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIdentityZoneInLocal deletes IdentityZoneInLocal
func (cli *ZSClient) DeleteIdentityZoneInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/identity-zone/{uuid}", uuid, string(deleteMode))
}

// QueryVRouterOspfNetwork queries VRouterOspfNetwork list
func (cli *ZSClient) QueryVRouterOspfNetwork(params *param.QueryParam) ([]view.NetworkRouterAreaRefInventoryView, error) {
	var resp []view.NetworkRouterAreaRefInventoryView
	return resp, cli.List("v1/routerArea/network", params, &resp)
}

// ChangeIAM2VirtualIDState changes IAM2VirtualIDState
func (cli *ZSClient) ChangeIAM2VirtualIDState(uuid string, params param.ChangeIAM2VirtualIDStateParam) (*view.IAM2VirtualIDInventoryView, error) {
	var resp view.ChangeIAM2VirtualIDStateEventView
	if err := cli.Put("v1/iam2/virtual-ids/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UnregisterLicenseServer operates on LicenseServer
func (cli *ZSClient) UnregisterLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server/client", uuid, string(deleteMode))
}

// CreateVmUserDefinedXmlHookScript creates VmUserDefinedXmlHookScript
func (cli *ZSClient) CreateVmUserDefinedXmlHookScript(params param.CreateVmUserDefinedXmlHookScriptParam) (*view.XmlHookInventoryView, error) {
	var resp view.CreateVmUserDefinedXmlHookScriptEventView
	if err := cli.Post("v1/vm-instances/xml-hook-script", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpgradeToLicenseServer operates on UpgradeToLicenseServer
func (cli *ZSClient) UpgradeToLicenseServer(params param.UpgradeToLicenseServerParam) (*view.LicenseAuthorizedNodeInventoryView, error) {
	var resp view.UpgradeToLicenseServerEventView
	if err := cli.Post("v1/license-server", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachAppBuildSystemToZone operates on AppBuildSystemToZone
func (cli *ZSClient) DetachAppBuildSystemToZone(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zones/{zoneUuid}/buildsystem/{buildSystemUuid}", uuid, string(deleteMode))
}

// GetAppBuildSystemCapacity gets AppBuildSystemCapacity by uuid
func (cli *ZSClient) GetAppBuildSystemCapacity(uuid string) (*view.GetAppBuildSystemCapacityView, error) {
	var resp view.GetAppBuildSystemCapacityView
	if err := cli.Get("v1/appcenter/buildsystem/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAttachableVpcL3Network gets AttachableVpcL3Network by uuid
func (cli *ZSClient) GetAttachableVpcL3Network(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attachable-vpc-l3s", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBaremetalChassisState changes BaremetalChassisState
func (cli *ZSClient) ChangeBaremetalChassisState(uuid string, params param.ChangeBaremetalChassisStateParam) (*view.BaremetalChassisInventoryView, error) {
	var resp view.ChangeBaremetalChassisStateEventView
	if err := cli.Put("v1/baremetal/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeNfvInstGroupOperationMode changes NfvInstGroupOperationMode
func (cli *ZSClient) ChangeNfvInstGroupOperationMode(uuid string, params param.ChangeNfvInstGroupOperationModeParam) (*view.NfvInstGroupInventoryView, error) {
	var resp view.ChangeNfvInstGroupOperationModeEventView
	if err := cli.Put("v1/nfvinstgroup/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetL3NetworkMtu gets L3NetworkMtu by uuid
func (cli *ZSClient) GetL3NetworkMtu(uuid string) (*view.GetL3NetworkMtuView, error) {
	var resp view.GetL3NetworkMtuView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/mtu", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVipToLoadBalancer operates on VipToLoadBalancer
func (cli *ZSClient) AttachVipToLoadBalancer(params param.AttachVipToLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	var resp view.AttachVipToLoadBalancerEventView
	if err := cli.Post("v1/load-balancers/{loadBalancerUuid}/vip/{vipUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateSecurityGroupRulePriority updates SecurityGroupRulePriority
func (cli *ZSClient) UpdateSecurityGroupRulePriority(uuid string, params param.UpdateSecurityGroupRulePriorityParam) (*view.SecurityGroupInventoryView, error) {
	var resp view.UpdateSecurityGroupRulePriorityEventView
	if err := cli.Put("v1/security-groups/{securityGroupUuid}/rules/priority/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryConnectionBetweenL3NetworkAndAliyunVSwitch queries ConnectionBetweenL3NetworkAndAliyunVSwitch list
func (cli *ZSClient) QueryConnectionBetweenL3NetworkAndAliyunVSwitch(params *param.QueryParam) ([]view.ConnectionRelationShipInventoryView, error) {
	var resp []view.ConnectionRelationShipInventoryView
	return resp, cli.List("v1/hybrid/aliyun/relationships", params, &resp)
}

// AddDnsToL3Network adds DnsToL3Network
func (cli *ZSClient) AddDnsToL3Network(params param.AddDnsToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	var resp view.AddDnsToL3NetworkEventView
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/dns", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryPortMirrorNetworkUsedIp queries PortMirrorNetworkUsedIp list
func (cli *ZSClient) QueryPortMirrorNetworkUsedIp(params *param.QueryParam) ([]view.MirrorNetworkUsedIpInventoryView, error) {
	var resp []view.MirrorNetworkUsedIpInventoryView
	return resp, cli.List("v1/port-mirrors/networks/usedIps", params, &resp)
}

// SetVmMonitorNumber operates on VmMonitorNumber
func (cli *ZSClient) SetVmMonitorNumber(uuid string, params param.SetVmMonitorNumberParam) (*view.SetVmMonitorNumberEventView, error) {
	resp := view.SetVmMonitorNumberEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeLoadBalancerBackendServer changes LoadBalancerBackendServer
func (cli *ZSClient) ChangeLoadBalancerBackendServer(uuid string, params param.ChangeLoadBalancerBackendServerParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	var resp view.ChangeLoadBalancerBackendServerEventView
	if err := cli.Put("v1/load-balancers/servergroups/{serverGroupUuid}/backendserver/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateEcsVSwitchRemote creates EcsVSwitchRemote
func (cli *ZSClient) CreateEcsVSwitchRemote(params param.CreateEcsVSwitchRemoteParam) (*view.EcsVSwitchInventoryView, error) {
	var resp view.CreateEcsVSwitchRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/vswitch", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVmMigrationCandidateHosts gets VmMigrationCandidateHosts by uuid
func (cli *ZSClient) GetVmMigrationCandidateHosts(uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/migration-target-hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForIpSecConnection gets CandidateL3NetworksForIpSecConnection by uuid
func (cli *ZSClient) GetCandidateL3NetworksForIpSecConnection(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.Get("v1/ipsec/candidatesL3Networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHostNetworkServiceType updates HostNetworkServiceType
func (cli *ZSClient) UpdateHostNetworkServiceType(uuid string, params param.UpdateHostNetworkServiceTypeParam) (*view.HostNetworkLabelInventoryView, error) {
	var resp view.UpdateHostNetworkServiceTypeEventView
	if err := cli.Put("v1/hosts/service-types/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SNSMicrosoftTeamsTestConnection operates on MicrosoftTeamsTestConnection
func (cli *ZSClient) SNSMicrosoftTeamsTestConnection(params param.SNSMicrosoftTeamsTestConnectionParam) (*view.SNSMicrosoftTeamsTestConnectionEventView, error) {
	resp := view.SNSMicrosoftTeamsTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/microsoft-teams/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLatestGuestToolsForVm gets LatestGuestToolsForVm by uuid
func (cli *ZSClient) GetLatestGuestToolsForVm(uuid string) (*view.GuestToolsInventoryView, error) {
	var resp view.GetLatestGuestToolsForVmView
	if err := cli.Get("v1/vm-instances/{uuid}/latest-guest-tools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateVpcUserVpnGatewayRemote creates VpcUserVpnGatewayRemote
func (cli *ZSClient) CreateVpcUserVpnGatewayRemote(params param.CreateVpcUserVpnGatewayRemoteParam) (*view.VpcUserVpnGatewayInventoryView, error) {
	var resp view.CreateVpcUserVpnGatewayRemoteEventView
	if err := cli.Post("v1/hybrid/user-vpn", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateOssBackupBucketRemote creates OssBackupBucketRemote
func (cli *ZSClient) CreateOssBackupBucketRemote(params param.CreateOssBackupBucketRemoteParam) (*view.CreateOssBackupBucketRemoteEventView, error) {
	resp := view.CreateOssBackupBucketRemoteEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql/oss", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOffBaremetalChassis operates on PowerOffBaremetalChassis
func (cli *ZSClient) PowerOffBaremetalChassis(uuid string, params param.PowerOffBaremetalChassisParam) (*view.PowerOffBaremetalChassisEventView, error) {
	resp := view.PowerOffBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateInterfaceVlanIds gets CandidateInterfaceVlanIds by uuid
func (cli *ZSClient) GetCandidateInterfaceVlanIds(uuid string) (*view.GetCandidateInterfaceVlanIdsView, error) {
	var resp view.GetCandidateInterfaceVlanIdsView
	if err := cli.Get("v1/host/network-interface-vlan-ids", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNetworkServiceTypes gets NetworkServiceTypes by uuid
func (cli *ZSClient) GetNetworkServiceTypes(uuid string) (*view.GetNetworkServiceTypesView, error) {
	var resp view.GetNetworkServiceTypesView
	if err := cli.Get("v1/network-services/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmUserDefinedXml deletes VmUserDefinedXml
func (cli *ZSClient) DeleteVmUserDefinedXml(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/xml", uuid, string(deleteMode))
}

// GetAvailableVpcL3Network gets AvailableVpcL3Network by uuid
func (cli *ZSClient) GetAvailableVpcL3Network(uuid string) (*view.GetAvailableVpcL3NetworkView, error) {
	var resp view.GetAvailableVpcL3NetworkView
	if err := cli.Get("v1/vpc/virtual-routers/available-vpc-l3s", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCurrentTime gets CurrentTime by uuid
func (cli *ZSClient) GetCurrentTime(uuid string) (*view.GetCurrentTimeView, error) {
	var resp view.GetCurrentTimeView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateAccountSpending operates on AccountSpending
func (cli *ZSClient) CalculateAccountSpending(uuid string, params param.CalculateAccountSpendingParam) (*view.CalculateAccountSpendingView, error) {
	resp := view.CalculateAccountSpendingView{}
	if err := cli.Put("v1/billings/accounts/{accountUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmAttachableL3Network gets VmAttachableL3Network by uuid
func (cli *ZSClient) GetVmAttachableL3Network(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/l3-networks-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEcsInstanceVncPassword updates EcsInstanceVncPassword
func (cli *ZSClient) UpdateEcsInstanceVncPassword(uuid string, params param.UpdateEcsInstanceVncPasswordParam) (*view.UpdateEcsInstanceVncPasswordEventView, error) {
	resp := view.UpdateEcsInstanceVncPasswordEventView{}
	if err := cli.Put("v1/hybrid/aliyun/{uuid}/ecs-vnc", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncChronyServers operates on ChronyServers
func (cli *ZSClient) SyncChronyServers(uuid string, params param.SyncChronyServersParam) (*view.SyncChronyServersEventView, error) {
	resp := view.SyncChronyServersEventView{}
	if err := cli.Put("v1/zops/chrony/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceProtectedRecoveryPoints gets VmInstanceProtectedRecoveryPoints by uuid
func (cli *ZSClient) GetVmInstanceProtectedRecoveryPoints(uuid string) (*view.GetVmInstanceProtectedRecoveryPointsView, error) {
	var resp view.GetVmInstanceProtectedRecoveryPointsView
	if err := cli.Get("v1/vm-instances/{uuid}/protected-recovery-points", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmToVmSchedulingRuleGroup adds VmToVmSchedulingRuleGroup
func (cli *ZSClient) AddVmToVmSchedulingRuleGroup(params param.AddVmToVmSchedulingRuleGroupParam) (*view.AddVmToVmSchedulingRuleGroupEventView, error) {
	resp := view.AddVmToVmSchedulingRuleGroupEventView{}
	if err := cli.Post("v1/vmSchedulingRuleGroup/{vmGroupUuid}/vmInstance/{vmUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncBackupFromImageStoreBackupStorage operates on BackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncBackupFromImageStoreBackupStorage(uuid string, params param.SyncBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	var resp view.SyncBackupFromImageStoreBackupStorageEventView
	if err := cli.Put("v1/volume-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetHostWebSshUrl gets HostWebSshUrl by uuid
func (cli *ZSClient) GetHostWebSshUrl(uuid string) (*view.GetHostWebSshUrlEventView, error) {
	var resp view.GetHostWebSshUrlEventView
	if err := cli.Get("v1/hosts/webssh", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetL3NetworkMtu operates on L3NetworkMtu
func (cli *ZSClient) SetL3NetworkMtu(params param.SetL3NetworkMtuParam) (*view.SetL3NetworkMtuEventView, error) {
	resp := view.SetL3NetworkMtuEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/mtu", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkRouterInterfaceIp gets L3NetworkRouterInterfaceIp by uuid
func (cli *ZSClient) GetL3NetworkRouterInterfaceIp(uuid string) (*view.GetL3NetworkRouterInterfaceIpView, error) {
	var resp view.GetL3NetworkRouterInterfaceIpView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/router-interface-ip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVmClock operates on VmClock
func (cli *ZSClient) SyncVmClock(uuid string, params param.SyncVmClockParam) (*view.SyncVmClockEventView, error) {
	resp := view.SyncVmClockEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcVpnConnectionFromLocal queries VpcVpnConnectionFromLocal list
func (cli *ZSClient) QueryVpcVpnConnectionFromLocal(params *param.QueryParam) ([]view.VpcVpnConnectionInventoryView, error) {
	var resp []view.VpcVpnConnectionInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection", params, &resp)
}

// CreateSNSSnmpEndpoint creates SNSSnmpEndpoint
func (cli *ZSClient) CreateSNSSnmpEndpoint(params param.CreateSNSSnmpEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.CreateSNSApplicationEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/snmp", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SdnControllerAddHost operates on SdnControllerAddHost
func (cli *ZSClient) SdnControllerAddHost(params param.SdnControllerAddHostParam) (*view.SdnControllerInventoryView, error) {
	var resp view.SdnControllerAddHostEventView
	if err := cli.Post("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetLicenseNodeUsageDetails gets LicenseNodeUsageDetails by uuid
func (cli *ZSClient) GetLicenseNodeUsageDetails(uuid string) (*view.GetLicenseNodeUsageDetailsView, error) {
	var resp view.GetLicenseNodeUsageDetailsView
	if err := cli.Get("v1/license/node/usage/details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAliyunSnapshotRemote creates AliyunSnapshotRemote
func (cli *ZSClient) CreateAliyunSnapshotRemote(params param.CreateAliyunSnapshotRemoteParam) (*view.AliyunSnapshotInventoryView, error) {
	var resp view.CreateAliyunSnapshotRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/snapshot", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetVmBootVolume operates on VmBootVolume
func (cli *ZSClient) SetVmBootVolume(uuid string, params param.SetVmBootVolumeParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmBootVolumeEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeVpcHaGroupMonitorIps changes VpcHaGroupMonitorIps
func (cli *ZSClient) ChangeVpcHaGroupMonitorIps(uuid string, params param.ChangeVpcHaGroupMonitorIpsParam) (*view.VpcHaGroupInventoryView, error) {
	var resp view.ChangeVpcHaGroupMonitorIpsEventView
	if err := cli.Put("v1/vpc/hagroups/{uuid}/monitorIps", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RenewSession operates on RenewSession
func (cli *ZSClient) RenewSession(uuid string, params param.RenewSessionParam) (*view.SessionInventoryView, error) {
	var resp view.RenewSessionEventView
	if err := cli.Put("v1/accounts/sessions/{sessionUuid}/renew", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteDataCenterInLocal deletes DataCenterInLocal
func (cli *ZSClient) DeleteDataCenterInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/data-center/{uuid}", uuid, string(deleteMode))
}

// SetVmConsoleMode operates on VmConsoleMode
func (cli *ZSClient) SetVmConsoleMode(uuid string, params param.SetVmConsoleModeParam) (*view.VmInstanceInventoryView, error) {
	var resp view.SetVmConsoleModeEventView
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachPolicyToUser operates on PolicyToUser
func (cli *ZSClient) AttachPolicyToUser(params param.AttachPolicyToUserParam) (*view.AttachPolicyToUserEventView, error) {
	resp := view.AttachPolicyToUserEventView{}
	if err := cli.Post("v1/accounts/users/{userUuid}/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVfPciDeviceAvailableInL2Network gets VfPciDeviceAvailableInL2Network by uuid
func (cli *ZSClient) GetVfPciDeviceAvailableInL2Network(uuid string) (*view.GetVfPciDeviceAvailableInL2NetworkView, error) {
	var resp view.GetVfPciDeviceAvailableInL2NetworkView
	if err := cli.Get("v1/l2-networks/vf-pci-devices-available", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAttributesToIAM2Project adds AttributesToIAM2Project
func (cli *ZSClient) AddAttributesToIAM2Project(params param.AddAttributesToIAM2ProjectParam) (*view.AddAttributesToIAM2ProjectEventView, error) {
	resp := view.AddAttributesToIAM2ProjectEventView{}
	if err := cli.Post("v1/iam2/projects/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UngenerateSeMdevDevices operates on UngenerateSeMdevDevices
func (cli *ZSClient) UngenerateSeMdevDevices(uuid string, params param.UngenerateSeMdevDevicesParam) (*view.UngenerateSeMdevDevicesEventView, error) {
	resp := view.UngenerateSeMdevDevicesEventView{}
	if err := cli.Put("v1/mtty-devices/{mttyDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmEmulatorPinning operates on VmEmulatorPinning
func (cli *ZSClient) SetVmEmulatorPinning(uuid string, params param.SetVmEmulatorPinningParam) (*view.SetVmEmulatorPinningEventView, error) {
	resp := view.SetVmEmulatorPinningEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanV2VConversionCache operates on V2VConversionCache
func (cli *ZSClient) CleanV2VConversionCache(uuid string, params param.CleanV2VConversionCacheParam) (*view.CleanV2VConversionCacheEventView, error) {
	resp := view.CleanV2VConversionCacheEventView{}
	if err := cli.Put("v1/v2v/conversion/host/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnbindModelFromService operates on UnbindModelFromService
func (cli *ZSClient) UnbindModelFromService(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models/{modelUuid}/model-services/{modelServiceUuid}", uuid, string(deleteMode))
}

// GetEcsInstanceType gets EcsInstanceType by uuid
func (cli *ZSClient) GetEcsInstanceType(uuid string) (*view.GetEcsInstanceTypeView, error) {
	var resp view.GetEcsInstanceTypeView
	if err := cli.Get("v1/hybrid/ecs/type", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseUKeyStatus gets LicenseUKeyStatus by uuid
func (cli *ZSClient) GetLicenseUKeyStatus(uuid string) (*view.UKeyInventoryView, error) {
	var resp view.UKeyInventoryView
	if err := cli.Get("v1/licenses/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddTicketTypesToTicketFlowCollection adds TicketTypesToTicketFlowCollection
func (cli *ZSClient) AddTicketTypesToTicketFlowCollection(params param.AddTicketTypesToTicketFlowCollectionParam) (*view.AddTicketTypesToTicketFlowCollectionEventView, error) {
	resp := view.AddTicketTypesToTicketFlowCollectionEventView{}
	if err := cli.Post("v1/tickets/flow-collections/{ticketFlowCollectionUuid}/ticket-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetL3NetworkRouterInterfaceIp operates on L3NetworkRouterInterfaceIp
func (cli *ZSClient) SetL3NetworkRouterInterfaceIp(params param.SetL3NetworkRouterInterfaceIpParam) (*view.SetL3NetworkRouterInterfaceIpEventView, error) {
	resp := view.SetL3NetworkRouterInterfaceIpEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/router-interface-ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConnectionBetweenL3NetworkAndAliyunVSwitch gets ConnectionBetweenL3NetworkAndAliyunVSwitch by uuid
func (cli *ZSClient) GetConnectionBetweenL3NetworkAndAliyunVSwitch(uuid string) (*view.ConnectionRelationShipPropertyView, error) {
	var resp view.ConnectionRelationShipPropertyView
	if err := cli.Get("v1/hybrid/aliyun/relationships", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEthernetVF queries EthernetVF list
func (cli *ZSClient) QueryEthernetVF(params *param.QueryParam) ([]view.EthernetVfPciDeviceInventoryView, error) {
	var resp []view.EthernetVfPciDeviceInventoryView
	return resp, cli.List("v1/pci-device/ethernet-vfs", params, &resp)
}

// GetBareMetal2GatewayAllocatorStrategies gets BareMetal2GatewayAllocatorStrategies by uuid
func (cli *ZSClient) GetBareMetal2GatewayAllocatorStrategies(uuid string) (*view.GetBareMetal2GatewayAllocatorStrategiesView, error) {
	var resp view.GetBareMetal2GatewayAllocatorStrategiesView
	if err := cli.Get("v1/baremetal2/gateways/allocators/strategies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsSecurityGroupFromLocal queries EcsSecurityGroupFromLocal list
func (cli *ZSClient) QueryEcsSecurityGroupFromLocal(params *param.QueryParam) ([]view.EcsSecurityGroupInventoryView, error) {
	var resp []view.EcsSecurityGroupInventoryView
	return resp, cli.List("v1/hybrid/aliyun/security-group", params, &resp)
}

// UpdateFirewallRuleTemplate updates FirewallRuleTemplate
func (cli *ZSClient) UpdateFirewallRuleTemplate(uuid string, params param.UpdateFirewallRuleTemplateParam) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	var resp view.UpdateFirewallRuleTemplateEventView
	if err := cli.Put("v1/vpcfirewalls/rules/template/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetUsbDeviceCandidatesForAttachingVm gets UsbDeviceCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetUsbDeviceCandidatesForAttachingVm(uuid string) (*view.UsbDeviceInventoryView, error) {
	var resp view.UsbDeviceInventoryView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/candidate-usb-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForLoadBalancer gets CandidateL3NetworksForLoadBalancer by uuid
func (cli *ZSClient) GetCandidateL3NetworksForLoadBalancer(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.Get("v1/load-balancers/listeners/{listenerUuid}/networks/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// WithdrawLicenseCapacityApplication operates on WithdrawLicenseCapacityApplication
func (cli *ZSClient) WithdrawLicenseCapacityApplication(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server/capacity-application", uuid, string(deleteMode))
}

// PowerResetHost operates on PowerResetHost
func (cli *ZSClient) PowerResetHost(uuid string, params param.PowerResetHostParam) (*view.HostInventoryView, error) {
	var resp view.PowerResetHostEventView
	if err := cli.Put("v1/hosts/power/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryFirewallRule queries FirewallRule list
func (cli *ZSClient) QueryFirewallRule(params *param.QueryParam) ([]view.VpcFirewallRuleInventoryView, error) {
	var resp []view.VpcFirewallRuleInventoryView
	return resp, cli.List("v1/vpcfirewalls/rules", params, &resp)
}

// RevertVmFromVmBackup operates on VmFromVmBackup
func (cli *ZSClient) RevertVmFromVmBackup(uuid string, params param.RevertVmFromVmBackupParam) (*view.RevertVmFromVmBackupEventView, error) {
	resp := view.RevertVmFromVmBackupEventView{}
	if err := cli.Put("v1/vm-backups/{groupUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachCCSCertificateToUser operates on CCSCertificateToUser
func (cli *ZSClient) AttachCCSCertificateToUser(params param.AttachCCSCertificateToUserParam) (*view.CCSCertificateInventoryView, error) {
	var resp view.AttachCCSCertificateToUserEventView
	if err := cli.Post("v1/crypto/ccs-certificate/attach-user/{userUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryOssBucketFileName queries OssBucketFileName list
func (cli *ZSClient) QueryOssBucketFileName(params *param.QueryParam) ([]view.OssBucketInventoryView, error) {
	var resp []view.OssBucketInventoryView
	return resp, cli.List("v1/hybrid/aliyun/oss-bucket", params, &resp)
}

// SetVmNuma operates on VmNuma
func (cli *ZSClient) SetVmNuma(uuid string, params param.SetVmNumaParam) (*view.SetVmNumaEventView, error) {
	resp := view.SetVmNumaEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LogIn operates on LogIn
func (cli *ZSClient) LogIn(uuid string, params param.LogInParam) (*view.SessionInventoryView, error) {
	var resp view.LogInView
	if err := cli.Put("v1/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QueryVRouterOspfArea queries VRouterOspfArea list
func (cli *ZSClient) QueryVRouterOspfArea(params *param.QueryParam) ([]view.RouterAreaInventoryView, error) {
	var resp []view.RouterAreaInventoryView
	return resp, cli.List("v1/routerArea", params, &resp)
}

// DeleteAliyunRouterInterfaceLocal deletes AliyunRouterInterfaceLocal
func (cli *ZSClient) DeleteAliyunRouterInterfaceLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/router-interface/{uuid}", uuid, string(deleteMode))
}

// UpdateFirewallRuleSet updates FirewallRuleSet
func (cli *ZSClient) UpdateFirewallRuleSet(uuid string, params param.UpdateFirewallRuleSetParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	var resp view.UpdateFirewallRuleSetEventView
	if err := cli.Put("v1/vpcfirewalls/ruleSets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachAliyunKey operates on AliyunKey
func (cli *ZSClient) AttachAliyunKey(uuid string, params param.AttachAliyunKeyParam) (*view.AttachAliyunKeyEventView, error) {
	resp := view.AttachAliyunKeyEventView{}
	if err := cli.Put("v1/hybrid/aliyun/key/{uuid}/attach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshSearchIndexes operates on SearchIndexes
func (cli *ZSClient) RefreshSearchIndexes(params param.RefreshSearchIndexesParam) (*view.RefreshSearchIndexesView, error) {
	var resp view.RefreshSearchIndexesView
	if err := cli.Get("v1/search/indexes/refresh", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateImageHash operates on ImageHash
func (cli *ZSClient) CalculateImageHash(uuid string, params param.CalculateImageHashParam) (*view.ImageInventoryView, error) {
	var resp view.CalculateImageHashEventView
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetL2NetworkTypes gets L2NetworkTypes by uuid
func (cli *ZSClient) GetL2NetworkTypes(uuid string) (*view.GetL2NetworkTypesView, error) {
	var resp view.GetL2NetworkTypesView
	if err := cli.Get("v1/l2-networks/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShutdownHost operates on ShutdownHost
func (cli *ZSClient) ShutdownHost(uuid string, params param.ShutdownHostParam) (*view.HostInventoryView, error) {
	var resp view.ShutdownHostEventView
	if err := cli.Put("v1/hosts/power/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateVpcVpnConnectionRemote updates VpcVpnConnectionRemote
func (cli *ZSClient) UpdateVpcVpnConnectionRemote(uuid string, params param.UpdateVpcVpnConnectionRemoteParam) (*view.VpcVpnConnectionInventoryView, error) {
	var resp view.UpdateVpcVpnConnectionRemoteEventView
	if err := cli.Put("v1/hybrid/vpn-connection/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVmTask gets VmTask by uuid
func (cli *ZSClient) GetVmTask(uuid string) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.Get("v1/vm-instances/task-details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableCdpTask operates on DisableCdpTask
func (cli *ZSClient) DisableCdpTask(params param.DisableCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	var resp view.DisableCdpTaskEventView
	if err := cli.Post("v1/cdp-task/disable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SetIpOnHostNetworkBonding operates on IpOnHostNetworkBonding
func (cli *ZSClient) SetIpOnHostNetworkBonding(params param.SetIpOnHostNetworkBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	var resp view.SetIpOnHostNetworkBondingEventView
	if err := cli.Post("v1/hosts/bondings/{bondingUuid}/ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveAttributesFromIAM2VirtualID removes AttributesFromIAM2VirtualID
func (cli *ZSClient) RemoveAttributesFromIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/virtual-ids/{uuid}/attributes", uuid, string(deleteMode))
}

// CreateBonding creates Bonding
func (cli *ZSClient) CreateBonding(params param.CreateBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.Post("v1/hosts/bondings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachUsbDeviceFromVm operates on UsbDeviceFromVm
func (cli *ZSClient) DetachUsbDeviceFromVm(params param.DetachUsbDeviceFromVmParam) (*view.UsbDeviceInventoryView, error) {
	var resp view.DetachUsbDeviceFromVmEventView
	if err := cli.Post("v1/usb-device/usb-devices/{usbDeviceUuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateDataVolumeTemplateFromVolumeSnapshot creates DataVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeSnapshot(params param.CreateDataVolumeTemplateFromVolumeSnapshotParam) (*view.ImageInventoryView, error) {
	var resp view.CreateDataVolumeTemplateFromVolumeSnapshotEventView
	if err := cli.Post("v1/images/data-volume-templates/from/volume-snapshots/{snapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachRoleFromAccount operates on RoleFromAccount
func (cli *ZSClient) DetachRoleFromAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/accounts/{accountUuid}/roles/{roleUuid}", uuid, string(deleteMode))
}

// AddRendezvousPointToMulticastRouter adds RendezvousPointToMulticastRouter
func (cli *ZSClient) AddRendezvousPointToMulticastRouter(params param.AddRendezvousPointToMulticastRouterParam) (*view.MulticastRouterInventoryView, error) {
	var resp view.AddRendezvousPointToMulticastRouterEventView
	if err := cli.Post("v1/multicast/virtual-routers/{uuid}/RendezvousPoint", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// QuerySNSTopicSubscriber queries SNSTopicSubscriber list
func (cli *ZSClient) QuerySNSTopicSubscriber(params *param.QueryParam) ([]view.SNSSubscriberInventoryView, error) {
	var resp []view.SNSSubscriberInventoryView
	return resp, cli.List("v1/sns/topics/subscribers", params, &resp)
}

// DeleteLdapBinding deletes LdapBinding
func (cli *ZSClient) DeleteLdapBinding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ldap/bindings/{uuid}", uuid, string(deleteMode))
}

// AttachNfvInstToGroup operates on NfvInstToGroup
func (cli *ZSClient) AttachNfvInstToGroup(uuid string, params param.AttachNfvInstToGroupParam) (*view.NfvInstGroupInventoryView, error) {
	var resp view.AttachNfvInstToGroupEventView
	if err := cli.Put("v1/nfvinstgroup/group/{groupUuid}/instances/{nfvInstUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DebugSignal operates on DebugSignal
func (cli *ZSClient) DebugSignal(params param.DebugSignalParam) (*view.DebugSignalEventView, error) {
	resp := view.DebugSignalEventView{}
	if err := cli.Post("v1/debug", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromVolume creates VmInstanceFromVolume
func (cli *ZSClient) CreateVmInstanceFromVolume(params param.CreateVmInstanceFromVolumeParam) (*view.VmInstanceInventoryView, error) {
	var resp view.CreateVmInstanceFromVolumeEventView
	if err := cli.Post("v1/vm-instances/from/volume", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetVpcVRouterDistributedRoutingEnabled gets VpcVRouterDistributedRoutingEnabled by uuid
func (cli *ZSClient) GetVpcVRouterDistributedRoutingEnabled(uuid string) (*view.GetVpcVRouterDistributedRoutingEnabledView, error) {
	var resp view.GetVpcVRouterDistributedRoutingEnabledView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/distributed-routing", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsSecurityGroupRemote creates EcsSecurityGroupRemote
func (cli *ZSClient) CreateEcsSecurityGroupRemote(params param.CreateEcsSecurityGroupRemoteParam) (*view.EcsSecurityGroupInventoryView, error) {
	var resp view.CreateEcsSecurityGroupRemoteEventView
	if err := cli.Post("v1/hybrid/aliyun/security-group/remote", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// RemoveAttributesFromIAM2Organization removes AttributesFromIAM2Organization
func (cli *ZSClient) RemoveAttributesFromIAM2Organization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/organizations/{uuid}/attributes", uuid, string(deleteMode))
}

// DeleteAliyunSnapshotFromLocal deletes AliyunSnapshotFromLocal
func (cli *ZSClient) DeleteAliyunSnapshotFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/snapshot/{uuid}", uuid, string(deleteMode))
}

// GetIAM2ProjectContainerImages gets IAM2ProjectContainerImages by uuid
func (cli *ZSClient) GetIAM2ProjectContainerImages(uuid string) (*view.ZakuImageInventoryView, error) {
	var resp view.ZakuImageInventoryView
	if err := cli.Get("v1/iam2/project/{projectId}/repository/{repositoryId}/image", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachDataVolumeFromVm operates on DataVolumeFromVm
func (cli *ZSClient) DetachDataVolumeFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{uuid}/vm-instances", uuid, string(deleteMode))
}

// QueryEcsVSwitchFromLocal queries EcsVSwitchFromLocal list
func (cli *ZSClient) QueryEcsVSwitchFromLocal(params *param.QueryParam) ([]view.EcsVSwitchInventoryView, error) {
	var resp []view.EcsVSwitchInventoryView
	return resp, cli.List("v1/hybrid/aliyun/vswitch", params, &resp)
}

// CreateRootVolumeTemplateFromRootVolume creates RootVolumeTemplateFromRootVolume
func (cli *ZSClient) CreateRootVolumeTemplateFromRootVolume(params param.CreateRootVolumeTemplateFromRootVolumeParam) (*view.ImageInventoryView, error) {
	var resp view.CreateRootVolumeTemplateFromRootVolumeEventView
	if err := cli.Post("v1/images/root-volume-templates/from/volumes/{rootVolumeUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachAliyunDiskToEcs operates on AliyunDiskToEcs
func (cli *ZSClient) AttachAliyunDiskToEcs(params param.AttachAliyunDiskToEcsParam) (*view.AliyunDiskInventoryView, error) {
	var resp view.AttachAliyunDiskToEcsEventView
	if err := cli.Post("v1/hybrid/aliyun/disk/{diskUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteOssBucketNameLocal deletes OssBucketNameLocal
func (cli *ZSClient) DeleteOssBucketNameLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/oss-bucket/{uuid}", uuid, string(deleteMode))
}

// QueryEcsImageFromLocal queries EcsImageFromLocal list
func (cli *ZSClient) QueryEcsImageFromLocal(params *param.QueryParam) ([]view.EcsImageInventoryView, error) {
	var resp []view.EcsImageInventoryView
	return resp, cli.List("v1/hybrid/aliyun/image", params, &resp)
}

// GetObservabilityServerServiceData gets ObservabilityServerServiceData by uuid
func (cli *ZSClient) GetObservabilityServerServiceData(uuid string) (*view.ObservabilityServerServiceDataInventoryView, error) {
	var resp view.ObservabilityServerServiceDataInventoryView
	if err := cli.Get("v1/observability-server/{observabilityServerUuid}/service-data", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunVirtualRouterFromLocal queries AliyunVirtualRouterFromLocal list
func (cli *ZSClient) QueryAliyunVirtualRouterFromLocal(params *param.QueryParam) ([]view.VpcVirtualRouterInventoryView, error) {
	var resp []view.VpcVirtualRouterInventoryView
	return resp, cli.List("v1/hybrid/aliyun/vrouter", params, &resp)
}

// VerifyLicenseServer operates on VerifyLicenseServer
func (cli *ZSClient) VerifyLicenseServer(params param.VerifyLicenseServerParam) (*view.VerifyLicenseServerEventView, error) {
	resp := view.VerifyLicenseServerEventView{}
	if err := cli.Post("v1/license-server/register-verify", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachBareMetal2GatewayToCluster operates on BareMetal2GatewayToCluster
func (cli *ZSClient) AttachBareMetal2GatewayToCluster(params param.AttachBareMetal2GatewayToClusterParam) (*view.BareMetal2GatewayInventoryView, error) {
	var resp view.AttachBareMetal2GatewayToClusterEventView
	if err := cli.Post("v1/baremetal2/clusters/{clusterUuid}/gateways/{gatewayUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateAtPersonOfAtWeComEndpoint updates AtPersonOfAtWeComEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtWeComEndpoint(uuid string, params param.UpdateAtPersonOfAtWeComEndpointParam) (*view.SNSWeComAtPersonInventoryView, error) {
	var resp view.UpdateAtPersonOfWeComEndpointEventView
	if err := cli.Put("v1/sns/application-endpoints/we-com/at-persons/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeSlbGroupDeployType changes SlbGroupDeployType
func (cli *ZSClient) ChangeSlbGroupDeployType(uuid string, params param.ChangeSlbGroupDeployTypeParam) (*view.SlbGroupInventoryView, error) {
	var resp view.ChangeSlbGroupDeployTypeEventView
	if err := cli.Put("v1/load-balancers/slb/groups/{slbGroupUuid}/deployType", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteEcsSecurityGroupInLocal deletes EcsSecurityGroupInLocal
func (cli *ZSClient) DeleteEcsSecurityGroupInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/security-group/{uuid}", uuid, string(deleteMode))
}

// DetachDataVolumeFromHost operates on DataVolumeFromHost
func (cli *ZSClient) DetachDataVolumeFromHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{volumeUuid}/hosts", uuid, string(deleteMode))
}

// GetVmInstanceRecoveryPoints gets VmInstanceRecoveryPoints by uuid
func (cli *ZSClient) GetVmInstanceRecoveryPoints(uuid string) (*view.GetVmInstanceRecoveryPointsView, error) {
	var resp view.GetVmInstanceRecoveryPointsView
	if err := cli.Get("v1/vm-instances/{uuid}/recovery-points", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSystemTags creates SystemTags
func (cli *ZSClient) CreateSystemTags(params param.CreateSystemTagsParam) (*view.SystemTagInventoryView, error) {
	resp := view.SystemTagInventoryView{}
	if err := cli.Post("v1/system-tags/{resourceUuid}/tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOssBackupBucketFromRemote gets OssBackupBucketFromRemote by uuid
func (cli *ZSClient) GetOssBackupBucketFromRemote(uuid string) (*view.GetOssBackupBucketFromRemoteView, error) {
	var resp view.GetOssBackupBucketFromRemoteView
	if err := cli.Get("v1/hybrid/backup-mysql/oss", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkTypes gets L3NetworkTypes by uuid
func (cli *ZSClient) GetL3NetworkTypes(uuid string) (*view.GetL3NetworkTypesView, error) {
	var resp view.GetL3NetworkTypesView
	if err := cli.Get("v1/l3-networks/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPoliciesFromUser operates on PoliciesFromUser
func (cli *ZSClient) DetachPoliciesFromUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/users/{userUuid}/policies", uuid, string(deleteMode))
}

// CleanUpImageCacheOnPrimaryStorage operates on UpImageCacheOnPrimaryStorage
func (cli *ZSClient) CleanUpImageCacheOnPrimaryStorage(uuid string, params param.CleanUpImageCacheOnPrimaryStorageParam) (*view.CleanUpImageCacheOnPrimaryStorageEventView, error) {
	resp := view.CleanUpImageCacheOnPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddKVMHostFromConfigFile adds KVMHostFromConfigFile
func (cli *ZSClient) AddKVMHostFromConfigFile(params param.AddKVMHostFromConfigFileParam) (*view.AddHostFromConfigFileEventView, error) {
	resp := view.AddHostFromConfigFileEventView{}
	if err := cli.Post("v1/hosts/kvm/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// InspectBareMetal2ChassisByInstance operates on BareMetal2ChassisByInstance
func (cli *ZSClient) InspectBareMetal2ChassisByInstance(uuid string, params param.InspectBareMetal2ChassisByInstanceParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.InspectBareMetal2ChassisByInstanceEventView
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteVmBootMode deletes VmBootMode
func (cli *ZSClient) DeleteVmBootMode(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/bootmode", uuid, string(deleteMode))
}

// GetCandidateVMForAttachingAffinityGroup gets CandidateVMForAttachingAffinityGroup by uuid
func (cli *ZSClient) GetCandidateVMForAttachingAffinityGroup(uuid string) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.Get("v1/VM/attachingGroup", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcVpnConnectionLocal deletes VpcVpnConnectionLocal
func (cli *ZSClient) DeleteVpcVpnConnectionLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/{uuid}", uuid, string(deleteMode))
}

// DetachPolicyFromUserGroup operates on PolicyFromUserGroup
func (cli *ZSClient) DetachPolicyFromUserGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/groups/{groupUuid}/policies/{policyUuid}", uuid, string(deleteMode))
}

// AddActionToAlarm adds ActionToAlarm
func (cli *ZSClient) AddActionToAlarm(params param.AddActionToAlarmParam) (*view.AlarmInventoryView, error) {
	var resp view.AddActionToAlarmEventView
	if err := cli.Post("v1/zwatch/alarms/{alarmUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UpdateFirewallRule updates FirewallRule
func (cli *ZSClient) UpdateFirewallRule(uuid string, params param.UpdateFirewallRuleParam) (*view.VpcFirewallRuleInventoryView, error) {
	var resp view.UpdateFirewallRuleEventView
	if err := cli.Put("v1/vpcfirewalls/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ZQLQuery operates on ZQLQuery
func (cli *ZSClient) ZQLQuery(params param.ZQLQueryParam) (*view.ZQLQueryView, error) {
	var resp view.ZQLQueryView
	if err := cli.Get("v1/zql", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetElaborations gets Elaborations by uuid
func (cli *ZSClient) GetElaborations(uuid string) (*view.GetElaborationsView, error) {
	var resp view.GetElaborationsView
	if err := cli.Get("v1/errorcode/elaborations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAccessPath gets AccessPath by uuid
func (cli *ZSClient) GetAccessPath(uuid string) (*view.GetAccessPathView, error) {
	var resp view.GetAccessPathView
	if err := cli.Get("v1/block-volumes/access/path", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageUsageReport gets PrimaryStorageUsageReport by uuid
func (cli *ZSClient) GetPrimaryStorageUsageReport(uuid string) (*view.GetPrimaryStorageUsageReportView, error) {
	var resp view.GetPrimaryStorageUsageReportView
	if err := cli.Get("v1/primary-storage/{primaryStorageUuid}/usage/report", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVolumeFromVolumeBackup operates on VolumeFromVolumeBackup
func (cli *ZSClient) RevertVolumeFromVolumeBackup(uuid string, params param.RevertVolumeFromVolumeBackupParam) (*view.RevertVolumeFromVolumeBackupEventView, error) {
	resp := view.RevertVolumeFromVolumeBackupEventView{}
	if err := cli.Put("v1/volume-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeFromVolumeTemplate creates DataVolumeFromVolumeTemplate
func (cli *ZSClient) CreateDataVolumeFromVolumeTemplate(params param.CreateDataVolumeFromVolumeTemplateParam) (*view.VolumeInventoryView, error) {
	var resp view.CreateDataVolumeFromVolumeTemplateEventView
	if err := cli.Post("v1/volumes/data/from/data-volume-templates/{imageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// LocalStorageGetVolumeMigratableHosts operates on LocalStorageGetVolumeMigratableHosts
func (cli *ZSClient) LocalStorageGetVolumeMigratableHosts(params param.LocalStorageGetVolumeMigratableHostsParam) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.Get("v1/volumes/{volumeUuid}/migration-target-hosts", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOssBucketNameFromRemote gets OssBucketNameFromRemote by uuid
func (cli *ZSClient) GetOssBucketNameFromRemote(uuid string) (*view.OssBucketPropertyView, error) {
	var resp view.OssBucketPropertyView
	if err := cli.Get("v1/hybrid/oss/{dataCenterUuid}/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsVpcFromRemote operates on EcsVpcFromRemote
func (cli *ZSClient) SyncEcsVpcFromRemote(params param.SyncEcsVpcFromRemoteParam) (*view.EcsVpcInventoryView, error) {
	resp := view.EcsVpcInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/vpc/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetServiceTypeOnHostNetworkInterface operates on ServiceTypeOnHostNetworkInterface
func (cli *ZSClient) SetServiceTypeOnHostNetworkInterface(params param.SetServiceTypeOnHostNetworkInterfaceParam) (*view.HostNetworkInterfaceServiceRefInventoryView, error) {
	resp := view.HostNetworkInterfaceServiceRefInventoryView{}
	if err := cli.Post("v1/hosts/nics/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddBackendServerToServerGroup adds BackendServerToServerGroup
func (cli *ZSClient) AddBackendServerToServerGroup(params param.AddBackendServerToServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	var resp view.AddBackendServerToServerGroupEventView
	if err := cli.Post("v1/load-balancers/servergroups/{serverGroupUuid}/backendservers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachUserDefinedXmlHookScriptToVm operates on UserDefinedXmlHookScriptToVm
func (cli *ZSClient) AttachUserDefinedXmlHookScriptToVm(params param.AttachUserDefinedXmlHookScriptToVmParam) (*view.AttachUserDefinedXmlHookScriptToVmEventView, error) {
	resp := view.AttachUserDefinedXmlHookScriptToVmEventView{}
	if err := cli.Post("v1/xmlhook/{xmlHookUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPolicyToRole operates on PolicyToRole
func (cli *ZSClient) AttachPolicyToRole(params param.AttachPolicyToRoleParam) (*view.AttachPolicyToRoleEventView, error) {
	resp := view.AttachPolicyToRoleEventView{}
	if err := cli.Post("v1/identities/policies/{policyUuid}/roles/{roleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2ProvisionNetworkState changes BareMetal2ProvisionNetworkState
func (cli *ZSClient) ChangeBareMetal2ProvisionNetworkState(uuid string, params param.ChangeBareMetal2ProvisionNetworkStateParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	var resp view.ChangeBareMetal2ProvisionNetworkStateEventView
	if err := cli.Put("v1/baremetal2/provision-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetBackupStorageCandidatesForImageMigration gets BackupStorageCandidatesForImageMigration by uuid
func (cli *ZSClient) GetBackupStorageCandidatesForImageMigration(uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.Get("v1/backup-storage/{srcBackupStorageUuid}/migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcIpSecConfigLocal deletes VpcIpSecConfigLocal
func (cli *ZSClient) DeleteVpcIpSecConfigLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/ipsec/{uuid}", uuid, string(deleteMode))
}

// GenerateSriovPciDevices operates on SriovPciDevices
func (cli *ZSClient) GenerateSriovPciDevices(uuid string, params param.GenerateSriovPciDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateAccountBillingSpending operates on AccountBillingSpending
func (cli *ZSClient) CalculateAccountBillingSpending(uuid string, params param.CalculateAccountBillingSpendingParam) (*view.CalculateAccountBillingSpendingView, error) {
	resp := view.CalculateAccountBillingSpendingView{}
	if err := cli.Put("v1/billings/accounts/{accountUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVRouterOspfArea deletes VRouterOspfArea
func (cli *ZSClient) DeleteVRouterOspfArea(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/routerArea/{uuid}", uuid, string(deleteMode))
}

// GetVipAvailablePort gets VipAvailablePort by uuid
func (cli *ZSClient) GetVipAvailablePort(uuid string) (*view.GetVipAvailablePortView, error) {
	var resp view.GetVipAvailablePortView
	if err := cli.Get("v1/vips/{vipUuid}/get-port-availability", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncDiskFromAliyunFromRemote operates on DiskFromAliyunFromRemote
func (cli *ZSClient) SyncDiskFromAliyunFromRemote(params param.SyncDiskFromAliyunFromRemoteParam) (*view.AliyunDiskInventoryView, error) {
	resp := view.AliyunDiskInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/disk/{identityUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVolumeState changes VolumeState
func (cli *ZSClient) ChangeVolumeState(uuid string, params param.ChangeVolumeStateParam) (*view.VolumeInventoryView, error) {
	var resp view.ChangeVolumeStateEventView
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// MountVmInstanceRecoveryPoint operates on MountVmInstanceRecoveryPoint
func (cli *ZSClient) MountVmInstanceRecoveryPoint(params param.MountVmInstanceRecoveryPointParam) (*view.MountVmInstanceRecoveryPointEventView, error) {
	resp := view.MountVmInstanceRecoveryPointEventView{}
	if err := cli.Post("v1/cdp-backup-storage/mount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVxlanPoolRemoteVtep creates VxlanPoolRemoteVtep
func (cli *ZSClient) CreateVxlanPoolRemoteVtep(params param.CreateVxlanPoolRemoteVtepParam) (*view.RemoteVtepInventoryView, error) {
	var resp view.CreateVxlanPoolRemoteVtepEventView
	if err := cli.Post("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}/remote-vtep-ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetResourceStackFromResource gets ResourceStackFromResource by uuid
func (cli *ZSClient) GetResourceStackFromResource(uuid string) (*view.GetResourceStackFromResourceView, error) {
	var resp view.GetResourceStackFromResourceView
	if err := cli.Get("v1/cloudformation/resources/stack", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIAM2VirtualIDConfigFile operates on IAM2VirtualIDConfigFile
func (cli *ZSClient) CheckIAM2VirtualIDConfigFile(uuid string, params param.CheckIAM2VirtualIDConfigFileParam) (*view.CheckIAM2VirtualIDConfigFileView, error) {
	resp := view.CheckIAM2VirtualIDConfigFileView{}
	if err := cli.Put("v1/iam2/virtual-ids/from-file", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetClusterHostNetworkFacts gets ClusterHostNetworkFacts by uuid
func (cli *ZSClient) GetClusterHostNetworkFacts(uuid string) (*view.GetClusterHostNetworkFactsView, error) {
	var resp view.GetClusterHostNetworkFactsView
	if err := cli.Get("v1/cluster/hosts-network-facts/{clusterUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachOssBucketFromEcsDataCenter operates on OssBucketFromEcsDataCenter
func (cli *ZSClient) DetachOssBucketFromEcsDataCenter(uuid string, params param.DetachOssBucketFromEcsDataCenterParam) (*view.OssBucketInventoryView, error) {
	var resp view.DetachOssBucketFromEcsDataCenterEventView
	if err := cli.Put("v1/hybrid/aliyun/oss-bucket/{ossBucketUuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ParseOvf operates on ParseOvf
func (cli *ZSClient) ParseOvf(params param.ParseOvfParam) (*view.ParseOvfView, error) {
	resp := view.ParseOvfView{}
	if err := cli.Post("v1/ovf/parse", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFaultToleranceVm queries FaultToleranceVm list
func (cli *ZSClient) QueryFaultToleranceVm(params *param.QueryParam) ([]view.FaultToleranceVmGroupInventoryView, error) {
	var resp []view.FaultToleranceVmGroupInventoryView
	return resp, cli.List("v1/vm-instances/fault-tolerance", params, &resp)
}

// AddSchedulerJobGroupToSchedulerTrigger adds SchedulerJobGroupToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobGroupToSchedulerTrigger(params param.AddSchedulerJobGroupToSchedulerTriggerParam) (*view.SchedulerJobGroupSchedulerTriggerRefInventoryView, error) {
	var resp view.AddSchedulerJobGroupToSchedulerTriggerEventView
	if err := cli.Post("v1/scheduler/jobgroups/{schedulerJobGroupUuid}/scheduler/triggers/{schedulerTriggerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DeleteAliyunNasAccessGroupRule deletes AliyunNasAccessGroupRule
func (cli *ZSClient) DeleteAliyunNasAccessGroupRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nas/access/rule/{uuid}", uuid, string(deleteMode))
}

// GetLoginProcedures gets LoginProcedures by uuid
func (cli *ZSClient) GetLoginProcedures(uuid string) (*view.GetLoginProceduresView, error) {
	var resp view.GetLoginProceduresView
	if err := cli.Get("v1/login/procedures", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBonding deletes Bonding
func (cli *ZSClient) DeleteBonding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts/bondings/{uuid}", uuid, string(deleteMode))
}

// DeleteEcsSecurityGroupRemote deletes EcsSecurityGroupRemote
func (cli *ZSClient) DeleteEcsSecurityGroupRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/security-group/remote/{uuid}", uuid, string(deleteMode))
}

// DeleteVmNicFromSecurityGroup deletes VmNicFromSecurityGroup
func (cli *ZSClient) DeleteVmNicFromSecurityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups/{securityGroupUuid}/vm-instances/nics", uuid, string(deleteMode))
}

// UpdateTag updates Tag
func (cli *ZSClient) UpdateTag(uuid string, params param.UpdateTagParam) (*view.TagPatternInventoryView, error) {
	var resp view.UpdateTagEventView
	if err := cli.Put("v1/tags/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AttachVRouterRouteTableToVRouter operates on VRouterRouteTableToVRouter
func (cli *ZSClient) AttachVRouterRouteTableToVRouter(params param.AttachVRouterRouteTableToVRouterParam) (*view.VRouterRouteTableInventoryView, error) {
	var resp view.AttachVRouterRouteTableToVRouterEventView
	if err := cli.Post("v1/vrouter-route-tables/{routeTableUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateVxlanVtep creates VxlanVtep
func (cli *ZSClient) CreateVxlanVtep(params param.CreateVxlanVtepParam) (*view.VtepInventoryView, error) {
	var resp view.CreateVxlanVtepEventView
	if err := cli.Post("v1/l2-networks/vxlan/vteps", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddMdevDeviceSpecToVmInstance adds MdevDeviceSpecToVmInstance
func (cli *ZSClient) AddMdevDeviceSpecToVmInstance(params param.AddMdevDeviceSpecToVmInstanceParam) (*view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp view.AddMdevDeviceSpecToVmInstanceEventView
	if err := cli.Post("v1/mdev-device-specs/{mdevSpecUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// DetachScsiLunFromVmInstance operates on ScsiLunFromVmInstance
func (cli *ZSClient) DetachScsiLunFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/scsi-lun/{uuid}", uuid, string(deleteMode))
}

// EnableCdpTask operates on EnableCdpTask
func (cli *ZSClient) EnableCdpTask(params param.EnableCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	var resp view.EnableCdpTaskEventView
	if err := cli.Post("v1/cdp-task/enable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SyncConnectionAccessPointFromRemote operates on ConnectionAccessPointFromRemote
func (cli *ZSClient) SyncConnectionAccessPointFromRemote(uuid string, params param.SyncConnectionAccessPointFromRemoteParam) (*view.ConnectionAccessPointInventoryView, error) {
	resp := view.ConnectionAccessPointInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/access-point/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RegisterLicenseRequestedApplication operates on LicenseRequestedApplication
func (cli *ZSClient) RegisterLicenseRequestedApplication(params param.RegisterLicenseRequestedApplicationParam) (*view.RegisterLicenseRequestedApplicationEventView, error) {
	resp := view.RegisterLicenseRequestedApplicationEventView{}
	if err := cli.Post("v1/licenses/applications", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVpcVpnGatewayFromRemote operates on VpcVpnGatewayFromRemote
func (cli *ZSClient) SyncVpcVpnGatewayFromRemote(uuid string, params param.SyncVpcVpnGatewayFromRemoteParam) (*view.VpcVpnGatewayInventoryView, error) {
	resp := view.VpcVpnGatewayInventoryView{}
	if err := cli.Put("v1/hybrid/vpc-vpn/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsVpcInLocal deletes EcsVpcInLocal
func (cli *ZSClient) DeleteEcsVpcInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vpc/{uuid}", uuid, string(deleteMode))
}

