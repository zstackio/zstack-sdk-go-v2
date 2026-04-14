// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ChangeIAM2OrganizationState changes IAM2OrganizationState
func (cli *ZSClient) ChangeIAM2OrganizationState(ctx context.Context, uuid string, params param.ChangeIAM2OrganizationStateParam) (*view.IAM2OrganizationInventoryView, error) {
	resp := view.IAM2OrganizationInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/organizations", uuid, "", map[string]interface{}{
		"changeIAM2OrganizationState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAutoScalingGroupAddingNewInstanceRule creates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupAddingNewInstanceRule(ctx context.Context, params param.CreateAutoScalingGroupAddingNewInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.Post(ctx, "v1/autoscaling/rules/adding-new-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetServiceTypeOnHostNetworkBonding operates on ServiceTypeOnHostNetworkBonding
func (cli *ZSClient) SetServiceTypeOnHostNetworkBonding(ctx context.Context, params param.SetServiceTypeOnHostNetworkBondingParam) (*view.HostNetworkBondingServiceRefInventoryView, error) {
	resp := view.HostNetworkBondingServiceRefInventoryView{}
	if err := cli.Post(ctx, "v1/hosts/bondings/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryPciDevicePciDeviceOffering queries PciDevicePciDeviceOffering list
func (cli *ZSClient) QueryPciDevicePciDeviceOffering(ctx context.Context, params *param.QueryParam) ([]view.PciDevicePciDeviceOfferingRefInventoryView, error) {
	var resp []view.PciDevicePciDeviceOfferingRefInventoryView
	return resp, cli.List(ctx, "v1/pci-devices/pci-devices/pci-device-offerings", params, &resp)
}

func (cli *ZSClient) GetPciDevicePciDeviceOffering(ctx context.Context, uuid string) (*view.PciDevicePciDeviceOfferingRefInventoryView, error) {
	var resp view.PciDevicePciDeviceOfferingRefInventoryView
	if err := cli.Get(ctx, "v1/pci-devices/pci-devices/pci-device-offerings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePciDevicePciDeviceOffering Pagination
func (cli *ZSClient) PagePciDevicePciDeviceOffering(ctx context.Context, params *param.QueryParam) ([]view.PciDevicePciDeviceOfferingRefInventoryView, int, error) {
	var pciDevicePciDeviceOfferings []view.PciDevicePciDeviceOfferingRefInventoryView
	total, err := cli.Page(ctx, "v1/pci-devices/pci-devices/pci-device-offerings", params, &pciDevicePciDeviceOfferings)
	return pciDevicePciDeviceOfferings, total, err
}

// AddAttributesToIAM2Organization adds AttributesToIAM2Organization
func (cli *ZSClient) AddAttributesToIAM2Organization(ctx context.Context, params param.AddAttributesToIAM2OrganizationParam) (*view.AddAttributesToIAM2OrganizationEventView, error) {
	resp := view.AddAttributesToIAM2OrganizationEventView{}
	if err := cli.Post(ctx, "v1/iam2/organizations/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCreateEcsImageProgress gets CreateEcsImageProgress by uuid
func (cli *ZSClient) GetCreateEcsImageProgress(ctx context.Context, dataCenterUuid string, imageUuid string) (*view.GetCreateEcsImageProgressView, error) {
	var resp view.GetCreateEcsImageProgressView
	err := cli.GetWithSpec(ctx, "v1/hybrid/aliyun/image", dataCenterUuid, fmt.Sprintf("%s/progress", imageUuid), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAccessControlListToLoadBalancer adds AccessControlListToLoadBalancer
func (cli *ZSClient) AddAccessControlListToLoadBalancer(ctx context.Context, params param.AddAccessControlListToLoadBalancerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.Post(ctx, "v1/load-balancers/listeners/{listenerUuid}/access-control-lists", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LogOut operates on LogOut
func (cli *ZSClient) LogOut(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/accounts/sessions", uuid, string(deleteMode))
}

// GetVmXmlHookScript gets VmXmlHookScript by uuid
func (cli *ZSClient) GetVmXmlHookScript(ctx context.Context, uuid string) (*view.GetVmXmlHookScriptView, error) {
	var resp view.GetVmXmlHookScriptView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachHybridKey operates on HybridKey
func (cli *ZSClient) AttachHybridKey(ctx context.Context, uuid string, params param.AttachHybridKeyParam) (*view.AttachHybridKeyEventView, error) {
	resp := view.AttachHybridKeyEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/hybrid/key", uuid, "", map[string]interface{}{
		"attachHybridKey": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImageQga gets ImageQga by uuid
func (cli *ZSClient) GetImageQga(ctx context.Context, uuid string) (*view.GetImageQgaView, error) {
	var resp view.GetImageQgaView
	if err := cli.GetWithRespKey(ctx, "v1/images", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInterdependentL3NetworksBackupStorages gets InterdependentL3NetworksBackupStorages by uuid
func (cli *ZSClient) GetInterdependentL3NetworksBackupStorages(ctx context.Context) (*view.GetInterdependentL3NetworksBackupStoragesView, error) {
	var resp view.GetInterdependentL3NetworksBackupStoragesView
	if err := cli.GetWithRespKey(ctx, "v1/backupStorage-l3networks/dependencies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBackupFileInPublic deletes BackupFileInPublic
func (cli *ZSClient) DeleteBackupFileInPublic(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/backup-mysql", uuid, string(deleteMode))
}

// BatchCreateIAM2VirtualIDFromConfigFile operates on CreateIAM2VirtualIDFromConfigFile
func (cli *ZSClient) BatchCreateIAM2VirtualIDFromConfigFile(ctx context.Context, params param.BatchCreateIAM2VirtualIDFromConfigFileParam) (*view.BatchCreateIAM2VirtualIDFromConfigFileEventView, error) {
	resp := view.BatchCreateIAM2VirtualIDFromConfigFileEventView{}
	if err := cli.Post(ctx, "v1/iam2/virtual-ids/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmClockTrack operates on VmClockTrack
func (cli *ZSClient) SetVmClockTrack(ctx context.Context, uuid string, params param.SetVmClockTrackParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmClockTrack": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEmailMonitorTriggerAction updates EmailMonitorTrigger
func (cli *ZSClient) UpdateEmailMonitorTriggerAction(ctx context.Context, uuid string, params param.UpdateEmailMonitorTriggerActionParam) (*view.MonitorTriggerActionInventoryView, error) {
	resp := view.MonitorTriggerActionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/monitoring/trigger-actions/emails", uuid, "", map[string]interface{}{
		"updateEmailMonitorTriggerAction": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryLocalRaidController queries LocalRaidController list
func (cli *ZSClient) QueryLocalRaidController(ctx context.Context, params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, error) {
	var resp []view.RaidPhysicalDriveInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/local-raid/controllers", params, &resp)
}

func (cli *ZSClient) GetLocalRaidController(ctx context.Context, uuid string) (*view.RaidPhysicalDriveInventoryView, error) {
	var resp view.RaidPhysicalDriveInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/local-raid/controllers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLocalRaidController Pagination
func (cli *ZSClient) PageLocalRaidController(ctx context.Context, params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, int, error) {
	var localRaidControllers []view.RaidPhysicalDriveInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/local-raid/controllers", params, &localRaidControllers)
	return localRaidControllers, total, err
}

// SyncDataCenterFromRemote operates on DataCenterFromRemote
func (cli *ZSClient) SyncDataCenterFromRemote(ctx context.Context, uuid string) (*view.SyncDataCenterFromRemoteEventView, error) {
	var resp view.SyncDataCenterFromRemoteEventView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/data-center", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBackupStorageState changes BackupStorageState
func (cli *ZSClient) ChangeBackupStorageState(ctx context.Context, uuid string, params param.ChangeBackupStorageStateParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage", uuid, "", map[string]interface{}{
		"changeBackupStorageState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateIsoForAttachingVm gets CandidateIsoForAttachingVm by uuid
func (cli *ZSClient) GetCandidateIsoForAttachingVm(ctx context.Context, uuid string) (*view.GetCandidateIsoForAttachingVmView, error) {
	var resp view.GetCandidateIsoForAttachingVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SecurityMachineDetectSync operates on MachineDetectSync
func (cli *ZSClient) SecurityMachineDetectSync(ctx context.Context, params param.SecurityMachineDetectSyncParam) (*view.SecurityMachineDetectSyncEventView, error) {
	resp := view.SecurityMachineDetectSyncEventView{}
	if err := cli.Post(ctx, "v1/security-machine/{uuid}/detect/sync/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSecurityGroupState changes SecurityGroupState
func (cli *ZSClient) ChangeSecurityGroupState(ctx context.Context, uuid string, params param.ChangeSecurityGroupStateParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-groups", uuid, "", map[string]interface{}{
		"changeSecurityGroupState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2ChassisOfferingState changes BareMetal2ChassisOfferingState
func (cli *ZSClient) ChangeBareMetal2ChassisOfferingState(ctx context.Context, uuid string, params param.ChangeBareMetal2ChassisOfferingStateParam) (*view.BareMetal2ChassisOfferingInventoryView, error) {
	resp := view.BareMetal2ChassisOfferingInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/chassis/offerings", uuid, "", map[string]interface{}{
		"changeBareMetal2ChassisOfferingState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrometheusMetricLabelValue gets PrometheusMetricLabelValue by uuid
func (cli *ZSClient) GetPrometheusMetricLabelValue(ctx context.Context) (*view.GetPrometheusMetricLabelValueView, error) {
	var resp view.GetPrometheusMetricLabelValueView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/metrics/prometheus/label-values", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryConnectionAccessPointFromLocal queries ConnectionAccessPointFromLocal list
func (cli *ZSClient) QueryConnectionAccessPointFromLocal(ctx context.Context, params *param.QueryParam) ([]view.ConnectionAccessPointInventoryView, error) {
	var resp []view.ConnectionAccessPointInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/access-point", params, &resp)
}

func (cli *ZSClient) GetConnectionAccessPointFromLocal(ctx context.Context, uuid string) (*view.ConnectionAccessPointInventoryView, error) {
	var resp view.ConnectionAccessPointInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/access-point", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageConnectionAccessPointFromLocal Pagination
func (cli *ZSClient) PageConnectionAccessPointFromLocal(ctx context.Context, params *param.QueryParam) ([]view.ConnectionAccessPointInventoryView, int, error) {
	var connectionAccessPointFromLocals []view.ConnectionAccessPointInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/access-point", params, &connectionAccessPointFromLocals)
	return connectionAccessPointFromLocals, total, err
}

// UpdateAlarmData updates AlarmData
func (cli *ZSClient) UpdateAlarmData(ctx context.Context, params param.UpdateAlarmDataParam) (*view.UpdateAlarmDataEventView, error) {
	resp := view.UpdateAlarmDataEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/alarm-histories/actions", "", "", map[string]interface{}{
		"updateAlarmData": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpnIpsecConfig creates VpnIpsecConfig
func (cli *ZSClient) CreateVpnIpsecConfig(ctx context.Context, params param.CreateVpnIpsecConfigParam) (*view.VpcVpnIpSecConfigInventoryView, error) {
	resp := view.VpcVpnIpSecConfigInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/vpn-connection/ipsec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSEmailTestConnection operates on EmailTestConnection
func (cli *ZSClient) SNSEmailTestConnection(ctx context.Context, params param.SNSEmailTestConnectionParam) (*view.SNSEmailTestConnectionEventView, error) {
	resp := view.SNSEmailTestConnectionEventView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/email/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RegisterLicenseServer operates on LicenseServer
func (cli *ZSClient) RegisterLicenseServer(ctx context.Context, params param.RegisterLicenseServerParam) (*view.RegisterLicenseServerEventView, error) {
	resp := view.RegisterLicenseServerEventView{}
	if err := cli.Post(ctx, "v1/license-server/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAutoScalingGroupState changes AutoScalingGroupState
func (cli *ZSClient) ChangeAutoScalingGroupState(ctx context.Context, uuid string, params param.ChangeAutoScalingGroupStateParam) (*view.AutoScalingGroupInventoryView, error) {
	resp := view.AutoScalingGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/autoscaling/groups", uuid, "", map[string]interface{}{
		"changeAutoScalingGroupState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAutoScalingGroupRemovalInstanceRule creates AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupRemovalInstanceRule(ctx context.Context, params param.CreateAutoScalingGroupRemovalInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.Post(ctx, "v1/autoscaling/rules/removal-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSlbGroupMonitorIps changes SlbGroupMonitorIps
func (cli *ZSClient) ChangeSlbGroupMonitorIps(ctx context.Context, slbGroupUuid string, params param.ChangeSlbGroupMonitorIpsParam) (*view.SlbGroupInventoryView, error) {
	resp := view.SlbGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/load-balancers/slb/groups", slbGroupUuid, "", map[string]interface{}{
		"changeSlbGroupMonitorIps": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteModelEvaluationTasks deletes ModelEvaluationTasks
func (cli *ZSClient) DeleteModelEvaluationTasks(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/model-evaluation-tasks", uuid, string(deleteMode))
}

// AttachL3NetworkToVm operates on L3NetworkToVm
func (cli *ZSClient) AttachL3NetworkToVm(ctx context.Context, params param.AttachL3NetworkToVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/{vmInstanceUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL2NetworkToCluster operates on L2NetworkToCluster
func (cli *ZSClient) AttachL2NetworkToCluster(ctx context.Context, params param.AttachL2NetworkToClusterParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.Post(ctx, "v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPrimaryStorageToCluster operates on PrimaryStorageToCluster
func (cli *ZSClient) AttachPrimaryStorageToCluster(ctx context.Context, params param.AttachPrimaryStorageToClusterParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmNicType changes VmNicType
func (cli *ZSClient) ChangeVmNicType(ctx context.Context, vmNicUuid string, params param.ChangeVmNicTypeParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/nics", vmNicUuid, "", map[string]interface{}{
		"changeVmNicType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeFirewallRuleState changes FirewallRuleState
func (cli *ZSClient) ChangeFirewallRuleState(ctx context.Context, uuid string, params param.ChangeFirewallRuleStateParam) (*view.VpcFirewallRuleInventoryView, error) {
	resp := view.VpcFirewallRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vpcfirewalls/rules", uuid, "", map[string]interface{}{
		"changeFirewallRuleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMdevDeviceCandidates gets MdevDeviceCandidates by uuid
func (cli *ZSClient) GetMdevDeviceCandidates(ctx context.Context) (*view.GetMdevDeviceCandidatesView, error) {
	var resp view.GetMdevDeviceCandidatesView
	if err := cli.GetWithRespKey(ctx, "v1/mdev-devices/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTwoFactorAuthenticationState gets TwoFactorAuthenticationState by uuid
func (cli *ZSClient) GetTwoFactorAuthenticationState(ctx context.Context) (*view.GetTwoFactorAuthenticationStateView, error) {
	var resp view.GetTwoFactorAuthenticationStateView
	if err := cli.GetWithRespKey(ctx, "v1/twofactorauthentication/state", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BootstrapMiniHost operates on MiniHost
func (cli *ZSClient) BootstrapMiniHost(ctx context.Context, params param.BootstrapMiniHostParam) (*view.BootstrapMiniHostEventView, error) {
	resp := view.BootstrapMiniHostEventView{}
	if err := cli.Post(ctx, "v1/mini-clusters/hosts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveActionFromAlarm removes ActionFromAlarm
func (cli *ZSClient) RemoveActionFromAlarm(ctx context.Context, alarmUuid string, actionUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zwatch/alarms", alarmUuid, fmt.Sprintf("actions/%s", actionUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ChangeEipState changes EipState
func (cli *ZSClient) ChangeEipState(ctx context.Context, uuid string, params param.ChangeEipStateParam) (*view.EipInventoryView, error) {
	resp := view.EipInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/eips", uuid, "", map[string]interface{}{
		"changeEipState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachSshKeyPairFromVmInstance operates on SshKeyPairFromVmInstance
func (cli *ZSClient) DetachSshKeyPairFromVmInstance(ctx context.Context, sshKeyPairUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/ssh-key-pair", sshKeyPairUuid, fmt.Sprintf("vm-instance/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetPrimaryStorageCandidatesForVmMigration gets PrimaryStorageCandidatesForVmMigration by uuid
func (cli *ZSClient) GetPrimaryStorageCandidatesForVmMigration(ctx context.Context, uuid string) (*view.GetPrimaryStorageCandidatesForVmMigrationView, error) {
	var resp view.GetPrimaryStorageCandidatesForVmMigrationView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVolume operates on PrimaryStorageMigrateVolume
func (cli *ZSClient) PrimaryStorageMigrateVolume(ctx context.Context, volumeUuid string, params param.PrimaryStorageMigrateVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage/volumes", volumeUuid, "", map[string]interface{}{
		"primaryStorageMigrateVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVolumeAsync Async
func (cli *ZSClient) PrimaryStorageMigrateVolumeAsync(ctx context.Context, params param.PrimaryStorageMigrateVolumeParam) (string, error) {

	resource := "v1/primary-storage/volumes/{volumeUuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// DeleteHybridEipRemote deletes HybridEipRemote
func (cli *ZSClient) DeleteHybridEipRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/eip", uuid, string(deleteMode))
}

// DeleteModelServiceInstanceGroups deletes ModelServiceInstanceGroups
func (cli *ZSClient) DeleteModelServiceInstanceGroups(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ai/model-services/instances/groups", uuid, string(deleteMode))
}

// GetVmBootOrder gets VmBootOrder by uuid
func (cli *ZSClient) GetVmBootOrder(ctx context.Context, uuid string) (*view.GetVmBootOrderView, error) {
	var resp view.GetVmBootOrderView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmBootOrder operates on VmBootOrder
func (cli *ZSClient) SetVmBootOrder(ctx context.Context, uuid string, params param.SetVmBootOrderParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmBootOrder": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryThirdpartyAlert queries ThirdpartyAlert list
func (cli *ZSClient) QueryThirdpartyAlert(ctx context.Context, params *param.QueryParam) ([]view.ThirdpartyOriginalAlertInventoryView, error) {
	var resp []view.ThirdpartyOriginalAlertInventoryView
	return resp, cli.List(ctx, "v1/zwatch/third-party/alerts", params, &resp)
}

func (cli *ZSClient) GetThirdpartyAlert(ctx context.Context, uuid string) (*view.ThirdpartyOriginalAlertInventoryView, error) {
	var resp view.ThirdpartyOriginalAlertInventoryView
	if err := cli.Get(ctx, "v1/zwatch/third-party/alerts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageThirdpartyAlert Pagination
func (cli *ZSClient) PageThirdpartyAlert(ctx context.Context, params *param.QueryParam) ([]view.ThirdpartyOriginalAlertInventoryView, int, error) {
	var thirdpartyAlerts []view.ThirdpartyOriginalAlertInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/third-party/alerts", params, &thirdpartyAlerts)
	return thirdpartyAlerts, total, err
}

// GetDatabaseBackupFromImageStore gets DatabaseBackupFromImageStore by uuid
func (cli *ZSClient) GetDatabaseBackupFromImageStore(ctx context.Context) (*view.GetDatabaseBackupFromImageStoreView, error) {
	var resp view.GetDatabaseBackupFromImageStoreView
	if err := cli.GetWithRespKey(ctx, "v1/database-backups/image-store", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsVSwitchFromRemote operates on EcsVSwitchFromRemote
func (cli *ZSClient) SyncEcsVSwitchFromRemote(ctx context.Context, params param.SyncEcsVSwitchFromRemoteParam) (*view.EcsVSwitchInventoryView, error) {
	resp := view.EcsVSwitchInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/vswitch/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LocateLocalRaidPhysicalDrive operates on LocalRaidPhysicalDrive
func (cli *ZSClient) LocateLocalRaidPhysicalDrive(ctx context.Context, uuid string, params param.LocateLocalRaidPhysicalDriveParam) (*view.RaidPhysicalDriveInventoryView, error) {
	resp := view.RaidPhysicalDriveInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/storage-devices/local-raid/physical-drives", uuid, "", map[string]interface{}{
		"locateLocalRaidPhysicalDrive": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpBaremetalChassisBonding operates on UpBaremetalChassisBonding
func (cli *ZSClient) CleanUpBaremetalChassisBonding(ctx context.Context, chassisUuid string, params param.CleanUpBaremetalChassisBondingParam) (*view.CleanUpBaremetalChassisBondingEventView, error) {
	resp := view.CleanUpBaremetalChassisBondingEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal/chassis", chassisUuid, "", map[string]interface{}{
		"cleanUpBaremetalChassisBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemovePciDeviceSpecFromVmInstance removes PciDeviceSpecFromVmInstance
func (cli *ZSClient) RemovePciDeviceSpecFromVmInstance(ctx context.Context, pciSpecUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/pci-device-specs", pciSpecUuid, fmt.Sprintf("vm-instances/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddIAM2VirtualIDGroupToProjects adds IAM2VirtualIDGroupToProjects
func (cli *ZSClient) AddIAM2VirtualIDGroupToProjects(ctx context.Context, params param.AddIAM2VirtualIDGroupToProjectsParam) (*view.AddIAM2VirtualIDGroupToProjectsEventView, error) {
	resp := view.AddIAM2VirtualIDGroupToProjectsEventView{}
	if err := cli.Post(ctx, "v1/iam2/projects/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveServerGroupFromLoadBalancerListener removes ServerGroupFromLoadBalancerListener
func (cli *ZSClient) RemoveServerGroupFromLoadBalancerListener(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/listeners", uuid, string(deleteMode))
}

// AddSharedBlockToSharedBlockGroup adds SharedBlockToSharedBlockGroup
func (cli *ZSClient) AddSharedBlockToSharedBlockGroup(ctx context.Context, params param.AddSharedBlockToSharedBlockGroupParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	resp := view.SharedBlockGroupPrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/sharedblockgroup/{uuid}/sharedblocks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshCaptcha operates on Captcha
func (cli *ZSClient) RefreshCaptcha(ctx context.Context) (*view.RefreshCaptchaView, error) {
	var resp view.RefreshCaptchaView
	if err := cli.GetWithRespKey(ctx, "v1/captcha/refresh", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsVSwitchInLocal deletes EcsVSwitchInLocal
func (cli *ZSClient) DeleteEcsVSwitchInLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/vswitch", uuid, string(deleteMode))
}

// DeleteTag deletes Tag
func (cli *ZSClient) DeleteTag(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/tags", uuid, string(deleteMode))
}

// AddIAM2VirtualIDsToOrganization adds IAM2VirtualIDsToOrganization
func (cli *ZSClient) AddIAM2VirtualIDsToOrganization(ctx context.Context, params param.AddIAM2VirtualIDsToOrganizationParam) (*view.AddIAM2VirtualIDsToOrganizationEventView, error) {
	resp := view.AddIAM2VirtualIDsToOrganizationEventView{}
	if err := cli.Post(ctx, "v1/iam2/organizations/{organizationUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachProvisionNicToBonding operates on ProvisionNicToBonding
func (cli *ZSClient) AttachProvisionNicToBonding(ctx context.Context, params param.AttachProvisionNicToBondingParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.Post(ctx, "v1/baremetal2/bm-instances/{uuid}/bm2-bondings/{bondingUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportNbdVolumes operates on NbdVolumes
func (cli *ZSClient) ExportNbdVolumes(ctx context.Context, params param.ExportNbdVolumesParam) (*view.ExportNbdVolumesEventView, error) {
	resp := view.ExportNbdVolumesEventView{}
	if err := cli.Post(ctx, "v1/cbt-task/exportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelfTestLocalRaid operates on LocalRaid
func (cli *ZSClient) SelfTestLocalRaid(ctx context.Context, uuid string, params param.SelfTestLocalRaidParam) (*view.SelfTestLocalRaidEventView, error) {
	resp := view.SelfTestLocalRaidEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/storage-devices/local-raid/physical-drives", uuid, "", map[string]interface{}{
		"selfTestLocalRaid": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSNSApplicationPlatformState changes SNSApplicationPlatformState
func (cli *ZSClient) ChangeSNSApplicationPlatformState(ctx context.Context, uuid string, params param.ChangeSNSApplicationPlatformStateParam) (*view.SNSApplicationPlatformInventoryView, error) {
	resp := view.SNSApplicationPlatformInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-platforms", uuid, "", map[string]interface{}{
		"changeSNSApplicationPlatformState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOffBareMetal2Chassis operates on PowerOffBareMetal2Chassis
func (cli *ZSClient) PowerOffBareMetal2Chassis(ctx context.Context, uuid string, params param.PowerOffBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"powerOffBareMetal2Chassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SdnControllerChangeHost operates on SdnControllerChangeHost
func (cli *ZSClient) SdnControllerChangeHost(ctx context.Context, sdnControllerUuid string, hostUuid string, params param.SdnControllerChangeHostParam) (*view.SdnControllerInventoryView, error) {
	resp := view.SdnControllerInventoryView{}
	err := cli.PutWithSpec(ctx, "v1/sdn-controllers", sdnControllerUuid, fmt.Sprintf("hosts/%s/actions", hostUuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateResourcePrice updates ResourcePrice
func (cli *ZSClient) UpdateResourcePrice(ctx context.Context, uuid string, params param.UpdateResourcePriceParam) (*view.PriceInventoryView, error) {
	resp := view.PriceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/prices", uuid, "", map[string]interface{}{
		"updateResourcePrice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachTagFromResources operates on TagFromResources
func (cli *ZSClient) DetachTagFromResources(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/tags", uuid, string(deleteMode))
}

// ChangeHostState changes HostState
func (cli *ZSClient) ChangeHostState(ctx context.Context, uuid string, params param.ChangeHostStateParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts", uuid, "", map[string]interface{}{
		"changeHostState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmNicMac updates VmNicMac
func (cli *ZSClient) UpdateVmNicMac(ctx context.Context, vmNicUuid string, params param.UpdateVmNicMacParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/nics", vmNicUuid, "", map[string]interface{}{
		"updateVmNicMac": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmInstanceHaLevel deletes VmInstanceHaLevel
func (cli *ZSClient) DeleteVmInstanceHaLevel(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// DeleteResourcePrice deletes ResourcePrice
func (cli *ZSClient) DeleteResourcePrice(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/billings/prices", uuid, string(deleteMode))
}

// CleanUpBareMetal2Bonding operates on UpBareMetal2Bonding
func (cli *ZSClient) CleanUpBareMetal2Bonding(ctx context.Context, chassisUuid string, params param.CleanUpBareMetal2BondingParam) (*view.CleanUpBaremetal2BondingEventView, error) {
	resp := view.CleanUpBaremetal2BondingEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/chassis", chassisUuid, "", map[string]interface{}{
		"cleanUpBareMetal2Bonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteMetricData deletes MetricData
func (cli *ZSClient) DeleteMetricData(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/metrics", uuid, string(deleteMode))
}

// AddLabelToAlarm adds LabelToAlarm
func (cli *ZSClient) AddLabelToAlarm(ctx context.Context, params param.AddLabelToAlarmParam) (*view.AlarmLabelInventoryView, error) {
	resp := view.AlarmLabelInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/alarms/{alarmUuid}/labels", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAliyunRouterInterfaceFromRemote operates on AliyunRouterInterfaceFromRemote
func (cli *ZSClient) SyncAliyunRouterInterfaceFromRemote(ctx context.Context, dataCenterUuid string, params param.SyncAliyunRouterInterfaceFromRemoteParam) (*view.AliyunRouterInterfaceInventoryView, error) {
	resp := view.AliyunRouterInterfaceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/router-interface", dataCenterUuid, "", map[string]interface{}{
		"syncAliyunRouterInterfaceFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportVmOvaPackage operates on VmOvaPackage
func (cli *ZSClient) ExportVmOvaPackage(ctx context.Context, params param.ExportVmOvaPackageParam) (*view.ImagePackageInventoryView, error) {
	resp := view.ImagePackageInventoryView{}
	if err := cli.Post(ctx, "v1/ovf/ova-packages", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportVmOvaPackageAsync Async
func (cli *ZSClient) ExportVmOvaPackageAsync(ctx context.Context, params param.ExportVmOvaPackageParam) (string, error) {

	resource := "v1/ovf/ova-packages"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// RevertVmFromCdpBackup operates on VmFromCdpBackup
func (cli *ZSClient) RevertVmFromCdpBackup(ctx context.Context, vmInstanceUuid string, params param.RevertVmFromCdpBackupParam) (*view.RevertVmFromCdpBackupEventView, error) {
	resp := view.RevertVmFromCdpBackupEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/cdp-backups", vmInstanceUuid, "", map[string]interface{}{
		"revertVmFromCdpBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVmFromCdpBackupAsync Async
func (cli *ZSClient) RevertVmFromCdpBackupAsync(ctx context.Context, params param.RevertVmFromCdpBackupParam) (string, error) {

	resource := "v1/cdp-backups/{vmInstanceUuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// SNSFeiShuTestConnection operates on FeiShuTestConnection
func (cli *ZSClient) SNSFeiShuTestConnection(ctx context.Context, params param.SNSFeiShuTestConnectionParam) (*view.SNSFeiShuTestConnectionEventView, error) {
	resp := view.SNSFeiShuTestConnectionEventView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/feishu/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSchedulerExecutionReport gets SchedulerExecutionReport by uuid
func (cli *ZSClient) GetSchedulerExecutionReport(ctx context.Context) (*view.GetSchedulerExecutionReportView, error) {
	var resp view.GetSchedulerExecutionReportView
	if err := cli.GetWithRespKey(ctx, "v1/scheduler/report", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRuleFromConfigFile creates FirewallRuleFromConfigFile
func (cli *ZSClient) CreateFirewallRuleFromConfigFile(ctx context.Context, params param.CreateFirewallRuleFromConfigFileParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.Post(ctx, "v1/vpcfirewalls/rules/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSupportedIdentityModels gets SupportedIdentityModels by uuid
func (cli *ZSClient) GetSupportedIdentityModels(ctx context.Context) (*view.GetSupportedIdentityModelsView, error) {
	var resp view.GetSupportedIdentityModelsView
	if err := cli.GetWithRespKey(ctx, "v1/identity-models", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddUserToGroup adds UserToGroup
func (cli *ZSClient) AddUserToGroup(ctx context.Context, params param.AddUserToGroupParam) (*view.AddUserToGroupEventView, error) {
	resp := view.AddUserToGroupEventView{}
	if err := cli.Post(ctx, "v1/accounts/groups/{groupUuid}/users", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVRouterOspfArea updates VRouterOspfArea
func (cli *ZSClient) UpdateVRouterOspfArea(ctx context.Context, uuid string, params param.UpdateVRouterOspfAreaParam) (*view.RouterAreaInventoryView, error) {
	resp := view.RouterAreaInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/routerArea", uuid, "", map[string]interface{}{
		"updateVRouterOspfArea": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageTypes gets PrimaryStorageTypes by uuid
func (cli *ZSClient) GetPrimaryStorageTypes(ctx context.Context) (*view.GetPrimaryStorageTypesView, error) {
	var resp view.GetPrimaryStorageTypesView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIPSecConnection queries IPSecConnection list
func (cli *ZSClient) QueryIPSecConnection(ctx context.Context, params *param.QueryParam) ([]view.IPsecConnectionInventoryView, error) {
	var resp []view.IPsecConnectionInventoryView
	return resp, cli.List(ctx, "v1/ipsec", params, &resp)
}

func (cli *ZSClient) GetIPSecConnection(ctx context.Context, uuid string) (*view.IPsecConnectionInventoryView, error) {
	var resp view.IPsecConnectionInventoryView
	if err := cli.Get(ctx, "v1/ipsec", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIPSecConnection Pagination
func (cli *ZSClient) PageIPSecConnection(ctx context.Context, params *param.QueryParam) ([]view.IPsecConnectionInventoryView, int, error) {
	var iPSecConnections []view.IPsecConnectionInventoryView
	total, err := cli.Page(ctx, "v1/ipsec", params, &iPSecConnections)
	return iPSecConnections, total, err
}

// BatchDeleteVolumeSnapshot operates on DeleteVolumeSnapshot
func (cli *ZSClient) BatchDeleteVolumeSnapshot(ctx context.Context, params param.BatchDeleteVolumeSnapshotParam) (*view.BatchDeleteVolumeSnapshotEventView, error) {
	resp := view.BatchDeleteVolumeSnapshotEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/volume-snapshots/batch-delete", "", "", map[string]interface{}{
		"batchDeleteVolumeSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReloadLicense operates on ReloadLicense
func (cli *ZSClient) ReloadLicense(ctx context.Context, params param.ReloadLicenseParam) (*view.LicenseInventoryView, error) {
	resp := view.LicenseInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/licenses/actions", "", "", map[string]interface{}{
		"reloadLicense": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNicQos deletes NicQos
func (cli *ZSClient) DeleteNicQos(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// ChangeL2NetworkVlanId changes L2NetworkVlanId
func (cli *ZSClient) ChangeL2NetworkVlanId(ctx context.Context, uuid string, params param.ChangeL2NetworkVlanIdParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/l2-networks", uuid, "", map[string]interface{}{
		"changeL2NetworkVlanId": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceStackVmStatus gets ResourceStackVmStatus by uuid
func (cli *ZSClient) GetResourceStackVmStatus(ctx context.Context) (*view.GetResourceStackVmStatusView, error) {
	var resp view.GetResourceStackVmStatusView
	if err := cli.GetWithRespKey(ctx, "v1/cloudformation/stack/monitor/vmstatus", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachHybridKey operates on HybridKey
func (cli *ZSClient) DetachHybridKey(ctx context.Context, uuid string, params param.DetachHybridKeyParam) (*view.DetachHybridKeyEventView, error) {
	resp := view.DetachHybridKeyEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/hybrid/key", uuid, "", map[string]interface{}{
		"detachHybridKey": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveDnsFromVpcRouter removes DnsFromVpcRouter
func (cli *ZSClient) RemoveDnsFromVpcRouter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpc/virtual-routers", uuid, string(deleteMode))
}

// DeleteHybridEipFromLocal deletes HybridEipFromLocal
func (cli *ZSClient) DeleteHybridEipFromLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/eip", uuid, string(deleteMode))
}

// GetAvailableTriggers gets AvailableTriggers by uuid
func (cli *ZSClient) GetAvailableTriggers(ctx context.Context) (*view.GetAvailableTriggersView, error) {
	var resp view.GetAvailableTriggersView
	if err := cli.GetWithRespKey(ctx, "v1/scheduler/triggers/available", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReimageVmInstance operates on ReimageVmInstance
func (cli *ZSClient) ReimageVmInstance(ctx context.Context, vmInstanceUuid string, params param.ReimageVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"reimageVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDatasets updates Datasets
func (cli *ZSClient) UpdateDatasets(ctx context.Context, params param.UpdateDatasetsParam) (*view.UpdateDatasetsEventView, error) {
	resp := view.UpdateDatasetsEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/ai/datasets", "", "", map[string]interface{}{
		"updateDatasets": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsSecurityGroupRuleFromRemote operates on EcsSecurityGroupRuleFromRemote
func (cli *ZSClient) SyncEcsSecurityGroupRuleFromRemote(ctx context.Context, uuid string, params param.SyncEcsSecurityGroupRuleFromRemoteParam) (*view.EcsSecurityGroupRuleInventoryView, error) {
	resp := view.EcsSecurityGroupRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/security-group-rule", uuid, "", map[string]interface{}{
		"syncEcsSecurityGroupRuleFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallRuleTemplate queries FirewallRuleTemplate list
func (cli *ZSClient) QueryFirewallRuleTemplate(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleTemplateInventoryView, error) {
	var resp []view.VpcFirewallRuleTemplateInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/rules/templates", params, &resp)
}

func (cli *ZSClient) GetFirewallRuleTemplate(ctx context.Context, uuid string) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	var resp view.VpcFirewallRuleTemplateInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls/rules/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRuleTemplate Pagination
func (cli *ZSClient) PageFirewallRuleTemplate(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleTemplateInventoryView, int, error) {
	var firewallRuleTemplates []view.VpcFirewallRuleTemplateInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/rules/templates", params, &firewallRuleTemplates)
	return firewallRuleTemplates, total, err
}

// SyncIdentityFromRemote operates on IdentityFromRemote
func (cli *ZSClient) SyncIdentityFromRemote(ctx context.Context, uuid string) (*view.SyncIdentityFromRemoteEventView, error) {
	var resp view.SyncIdentityFromRemoteEventView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/identity-zone", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageStoreBackupStorageQuota operates on ImageStoreBackupStorageQuota
func (cli *ZSClient) SetImageStoreBackupStorageQuota(ctx context.Context, params param.SetImageStoreBackupStorageQuotaParam) (*view.SetImageStoreBackupStorageQuotaEventView, error) {
	resp := view.SetImageStoreBackupStorageQuotaEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage/image-store/actions", "", "", map[string]interface{}{
		"setImageStoreBackupStorageQuota": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeClusterState changes ClusterState
func (cli *ZSClient) ChangeClusterState(ctx context.Context, uuid string, params param.ChangeClusterStateParam) (*view.ClusterInventoryView, error) {
	resp := view.ClusterInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/clusters", uuid, "", map[string]interface{}{
		"changeClusterState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVfNicHaState changes VfNicHaState
func (cli *ZSClient) ChangeVfNicHaState(ctx context.Context, vfNicUuid string, params param.ChangeVfNicHaStateParam) (*view.VmVfNicInventoryView, error) {
	resp := view.VmVfNicInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/nics", vfNicUuid, "", map[string]interface{}{
		"changeVfNicHaState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOvnControllerOffering creates OvnControllerOffering
func (cli *ZSClient) CreateOvnControllerOffering(ctx context.Context, params param.CreateOvnControllerOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.Post(ctx, "v1/instance-offerings/ovn", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2OrganizationVirtualIDNumber gets IAM2OrganizationVirtualIDNumber by uuid
func (cli *ZSClient) GetIAM2OrganizationVirtualIDNumber(ctx context.Context, uuid string) (*view.GetIAM2OrganizationVirtualIDNumberView, error) {
	var resp view.GetIAM2OrganizationVirtualIDNumberView
	if err := cli.GetWithRespKey(ctx, "v1/iam2/organizations", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsInstanceLocal deletes EcsInstanceLocal
func (cli *ZSClient) DeleteEcsInstanceLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/ecs", uuid, string(deleteMode))
}

// ChangePortMirrorState changes PortMirrorState
func (cli *ZSClient) ChangePortMirrorState(ctx context.Context, uuid string, params param.ChangePortMirrorStateParam) (*view.PortMirrorInventoryView, error) {
	resp := view.PortMirrorInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/port-mirrors", uuid, "", map[string]interface{}{
		"changePortMirrorState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsubscribeSNSTopic operates on UnsubscribeSNSTopic
func (cli *ZSClient) UnsubscribeSNSTopic(ctx context.Context, topicUuid string, endpointUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/sns/topics", topicUuid, fmt.Sprintf("endpoints/%s", endpointUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// SetNicQos operates on NicQos
func (cli *ZSClient) SetNicQos(ctx context.Context, uuid string, params param.SetNicQosParam) (*view.SetNicQosEventView, error) {
	resp := view.SetNicQosEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setNicQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelLongJob operates on CancelLongJob
func (cli *ZSClient) CancelLongJob(ctx context.Context, uuid string, params param.CancelLongJobParam) (*view.CancelLongJobEventView, error) {
	resp := view.CancelLongJobEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/longjobs", uuid, "", map[string]interface{}{
		"cancelLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRouteTableVpcVRouterCandidate gets RouteTableVpcVRouterCandidate by uuid
func (cli *ZSClient) GetRouteTableVpcVRouterCandidate(ctx context.Context) (*view.GetRouteTableVpcVRouterCandidateView, error) {
	var resp view.GetRouteTableVpcVRouterCandidateView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers/get-vpc-candidate", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateAccountBilling operates on AccountBilling
func (cli *ZSClient) GenerateAccountBilling(ctx context.Context, accountUuid string, params param.GenerateAccountBillingParam) (*view.GenerateAccountBillingEventView, error) {
	resp := view.GenerateAccountBillingEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/accounts", accountUuid, "", map[string]interface{}{
		"generateAccountBilling": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAliyunVirtualRouterFromRemote operates on AliyunVirtualRouterFromRemote
func (cli *ZSClient) SyncAliyunVirtualRouterFromRemote(ctx context.Context, vpcUuid string, params param.SyncAliyunVirtualRouterFromRemoteParam) (*view.VpcVirtualRouterInventoryView, error) {
	resp := view.VpcVirtualRouterInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/vrouter", vpcUuid, "", map[string]interface{}{
		"syncAliyunVirtualRouterFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsVpcFromLocal queries EcsVpcFromLocal list
func (cli *ZSClient) QueryEcsVpcFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsVpcInventoryView, error) {
	var resp []view.EcsVpcInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/vpc", params, &resp)
}

func (cli *ZSClient) GetEcsVpcFromLocal(ctx context.Context, uuid string) (*view.EcsVpcInventoryView, error) {
	var resp view.EcsVpcInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/vpc", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsVpcFromLocal Pagination
func (cli *ZSClient) PageEcsVpcFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsVpcInventoryView, int, error) {
	var ecsVpcFromLocals []view.EcsVpcInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/vpc", params, &ecsVpcFromLocals)
	return ecsVpcFromLocals, total, err
}

// GetInvocationRecords gets InvocationRecords by uuid
func (cli *ZSClient) GetInvocationRecords(ctx context.Context) (*view.GetInvocationRecordsView, error) {
	var resp view.GetInvocationRecordsView
	if err := cli.GetWithRespKey(ctx, "v1/scripts/aliyun-invocations", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeRoleState changes RoleState
func (cli *ZSClient) ChangeRoleState(ctx context.Context, uuid string, params param.ChangeRoleStateParam) (*view.RoleInventoryView, error) {
	resp := view.RoleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/identities/roles", uuid, "", map[string]interface{}{
		"changeRoleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVRouterFlowCounter gets VRouterFlowCounter by uuid
func (cli *ZSClient) GetVRouterFlowCounter(ctx context.Context, uuid string) (*view.GetVRouterFlowCounterView, error) {
	var resp view.GetVRouterFlowCounterView
	if err := cli.GetWithRespKey(ctx, "v1/flowmeters", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAliyunRouterInterfaceRemote creates AliyunRouterInterfaceRemote
func (cli *ZSClient) CreateAliyunRouterInterfaceRemote(ctx context.Context, params param.CreateAliyunRouterInterfaceRemoteParam) (*view.AliyunRouterInterfaceInventoryView, error) {
	resp := view.AliyunRouterInterfaceInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/router-interface", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBareMetal2SupportedBootMode gets BareMetal2SupportedBootMode by uuid
func (cli *ZSClient) GetBareMetal2SupportedBootMode(ctx context.Context) (*view.GetBareMetal2SupportedBootModeView, error) {
	var resp view.GetBareMetal2SupportedBootModeView
	if err := cli.GetWithRespKey(ctx, "v1/baremetal2/chassis/supported-boot-modes", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostPowerStatus gets HostPowerStatus by uuid
func (cli *ZSClient) GetHostPowerStatus(ctx context.Context, uuid string, params param.GetHostPowerStatusParam) (*view.HostIpmiInventoryView, error) {
	resp := view.HostIpmiInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/power", uuid, "", map[string]interface{}{
		"getHostPowerStatus": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetChainTask gets ChainTask by uuid
func (cli *ZSClient) GetChainTask(ctx context.Context) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.GetWithRespKey(ctx, "v1/core/task-details", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch updates ConnectionBetweenL3NetWorkAndAliyunVSwitch
func (cli *ZSClient) UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch(ctx context.Context, uuid string, params param.UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParam) (*view.ConnectionRelationShipInventoryView, error) {
	resp := view.ConnectionRelationShipInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/connections", uuid, "", map[string]interface{}{
		"updateConnectionBetweenL3NetWorkAndAliyunVSwitch": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeHostPassword changes HostPassword
func (cli *ZSClient) ChangeHostPassword(ctx context.Context, hostUuid string, params param.ChangeHostPasswordParam) (*view.ChangeHostPasswordEventView, error) {
	resp := view.ChangeHostPasswordEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/kvm", hostUuid, "", map[string]interface{}{
		"changeHostPassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSlbInstance creates SlbInstance
func (cli *ZSClient) CreateSlbInstance(ctx context.Context, params param.CreateSlbInstanceParam) (*view.SlbVmInstanceInventoryView, error) {
	resp := view.SlbVmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/load-balancers/slb/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePortForwardingRuleState changes PortForwardingRuleState
func (cli *ZSClient) ChangePortForwardingRuleState(ctx context.Context, uuid string, params param.ChangePortForwardingRuleStateParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/port-forwarding", uuid, "", map[string]interface{}{
		"changePortForwardingRuleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsLicenseServer operates on IsLicenseServer
func (cli *ZSClient) IsLicenseServer(ctx context.Context) (*view.IsLicenseServerView, error) {
	var resp view.IsLicenseServerView
	if err := cli.GetWithRespKey(ctx, "v1/license-server/is-server", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryLabelValues operates on PrometheusQueryLabelValues
func (cli *ZSClient) PrometheusQueryLabelValues(ctx context.Context) (*view.PrometheusQueryLabelValuesView, error) {
	var resp view.PrometheusQueryLabelValuesView
	if err := cli.GetWithRespKey(ctx, "v1/prometheus/labels", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateClusterSupportDRS operates on ClusterSupportDRS
func (cli *ZSClient) ValidateClusterSupportDRS(ctx context.Context, uuid string) (*view.ValidateClusterSupportDRSView, error) {
	var resp view.ValidateClusterSupportDRSView
	if err := cli.GetWithRespKey(ctx, "v1/clusters", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShrinkVolumeSnapshot operates on ShrinkVolumeSnapshot
func (cli *ZSClient) ShrinkVolumeSnapshot(ctx context.Context, uuid string, params param.ShrinkVolumeSnapshotParam) (*view.ShrinkVolumeSnapshotEventView, error) {
	resp := view.ShrinkVolumeSnapshotEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/volume-snapshots/shrink", uuid, "", map[string]interface{}{
		"shrinkVolumeSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddHostToHostSchedulingRuleGroup adds HostToHostSchedulingRuleGroup
func (cli *ZSClient) AddHostToHostSchedulingRuleGroup(ctx context.Context, params param.AddHostToHostSchedulingRuleGroupParam) (*view.AddHostToHostSchedulingRuleGroupEventView, error) {
	resp := view.AddHostToHostSchedulingRuleGroupEventView{}
	if err := cli.Post(ctx, "v1/hostSchedulingRuleGroup/{hostGroupUuid}/host/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBuildApp creates BuildApp
func (cli *ZSClient) CreateBuildApp(ctx context.Context, params param.CreateBuildAppParam) (*view.BuildApplicationInventoryView, error) {
	resp := view.BuildApplicationInventoryView{}
	if err := cli.Post(ctx, "v1/appcenter/buildapp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIdentityZoneFromLocal queries IdentityZoneFromLocal list
func (cli *ZSClient) QueryIdentityZoneFromLocal(ctx context.Context, params *param.QueryParam) ([]view.IdentityZoneInventoryView, error) {
	var resp []view.IdentityZoneInventoryView
	return resp, cli.List(ctx, "v1/hybrid/identity-zone", params, &resp)
}

func (cli *ZSClient) GetIdentityZoneFromLocal(ctx context.Context, uuid string) (*view.IdentityZoneInventoryView, error) {
	var resp view.IdentityZoneInventoryView
	if err := cli.Get(ctx, "v1/hybrid/identity-zone", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIdentityZoneFromLocal Pagination
func (cli *ZSClient) PageIdentityZoneFromLocal(ctx context.Context, params *param.QueryParam) ([]view.IdentityZoneInventoryView, int, error) {
	var identityZoneFromLocals []view.IdentityZoneInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/identity-zone", params, &identityZoneFromLocals)
	return identityZoneFromLocals, total, err
}

// GetVmNicAttachedNetworkService gets VmNicAttachedNetworkService by uuid
func (cli *ZSClient) GetVmNicAttachedNetworkService(ctx context.Context, uuid string) (*view.GetVmNicAttachedNetworkServiceView, error) {
	var resp view.GetVmNicAttachedNetworkServiceView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/nics", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmHostname gets VmHostname by uuid
func (cli *ZSClient) GetVmHostname(ctx context.Context, uuid string) (*view.GetVmHostnameView, error) {
	var resp view.GetVmHostnameView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSchedulerJobsToSchedulerJobGroup adds SchedulerJobsToSchedulerJobGroup
func (cli *ZSClient) AddSchedulerJobsToSchedulerJobGroup(ctx context.Context, params param.AddSchedulerJobsToSchedulerJobGroupParam) (*view.SchedulerJobGroupJobRefInventoryView, error) {
	resp := view.SchedulerJobGroupJobRefInventoryView{}
	if err := cli.Post(ctx, "v1/scheduler/jobgroups/{schedulerJobGroupUuid}/job", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL3NetworkFromVm operates on L3NetworkFromVm
func (cli *ZSClient) DetachL3NetworkFromVm(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances/nics", uuid, string(deleteMode))
}

// DeleteVpcUserVpnGatewayLocal deletes VpcUserVpnGatewayLocal
func (cli *ZSClient) DeleteVpcUserVpnGatewayLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/user-gateway", uuid, string(deleteMode))
}

// CreateVRouterOspfArea creates VRouterOspfArea
func (cli *ZSClient) CreateVRouterOspfArea(ctx context.Context, params param.CreateVRouterOspfAreaParam) (*view.RouterAreaInventoryView, error) {
	resp := view.RouterAreaInventoryView{}
	if err := cli.Post(ctx, "v1/routerArea", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetSecurityMachineKey operates on SecurityMachineKey
func (cli *ZSClient) SetSecurityMachineKey(ctx context.Context, params param.SetSecurityMachineKeyParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.Post(ctx, "v1/secret-resource-pool-token/set/{uuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOAuthClient creates OAuthClient
func (cli *ZSClient) CreateOAuthClient(ctx context.Context, params param.CreateOAuthClientParam) (*view.OAuth2ClientInventoryView, error) {
	resp := view.OAuth2ClientInventoryView{}
	if err := cli.Post(ctx, "v1/create/oauth2/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedEip gets VpcAttachedEip by uuid
func (cli *ZSClient) GetVpcAttachedEip(ctx context.Context, params param.GetVpcAttachedEipParam) (*view.EipInventoryView, error) {
	resp := view.EipInventoryView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/attached-eip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveSchedulerJobFromSchedulerTrigger removes SchedulerJobFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobFromSchedulerTrigger(ctx context.Context, schedulerJobUuid string, schedulerTriggerUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/scheduler/jobs", schedulerJobUuid, fmt.Sprintf("scheduler/triggers/%s", schedulerTriggerUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ChangeMediaState changes MediaState
func (cli *ZSClient) ChangeMediaState(ctx context.Context, uuid string, params param.ChangeMediaStateParam) (*view.MediaInventoryView, error) {
	resp := view.MediaInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/media", uuid, "", map[string]interface{}{
		"changeMediaState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeIPSecConnectionState changes IPSecConnectionState
func (cli *ZSClient) ChangeIPSecConnectionState(ctx context.Context, uuid string, params param.ChangeIPSecConnectionStateParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/ipsec", uuid, "", map[string]interface{}{
		"changeIPSecConnectionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunDiskFromLocal queries AliyunDiskFromLocal list
func (cli *ZSClient) QueryAliyunDiskFromLocal(ctx context.Context, params *param.QueryParam) ([]view.AliyunDiskInventoryView, error) {
	var resp []view.AliyunDiskInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/disk", params, &resp)
}

func (cli *ZSClient) GetAliyunDiskFromLocal(ctx context.Context, uuid string) (*view.AliyunDiskInventoryView, error) {
	var resp view.AliyunDiskInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/disk", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunDiskFromLocal Pagination
func (cli *ZSClient) PageAliyunDiskFromLocal(ctx context.Context, params *param.QueryParam) ([]view.AliyunDiskInventoryView, int, error) {
	var aliyunDiskFromLocals []view.AliyunDiskInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/disk", params, &aliyunDiskFromLocals)
	return aliyunDiskFromLocals, total, err
}

// QueryEcsSecurityGroupRuleFromLocal queries EcsSecurityGroupRuleFromLocal list
func (cli *ZSClient) QueryEcsSecurityGroupRuleFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsSecurityGroupRuleInventoryView, error) {
	var resp []view.EcsSecurityGroupRuleInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/security-group-rule", params, &resp)
}

func (cli *ZSClient) GetEcsSecurityGroupRuleFromLocal(ctx context.Context, uuid string) (*view.EcsSecurityGroupRuleInventoryView, error) {
	var resp view.EcsSecurityGroupRuleInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/security-group-rule", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsSecurityGroupRuleFromLocal Pagination
func (cli *ZSClient) PageEcsSecurityGroupRuleFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsSecurityGroupRuleInventoryView, int, error) {
	var ecsSecurityGroupRuleFromLocals []view.EcsSecurityGroupRuleInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/security-group-rule", params, &ecsSecurityGroupRuleFromLocals)
	return ecsSecurityGroupRuleFromLocals, total, err
}

// StopAllResourcesInIAM2Project stops AllResourcesInIAM2Project
func (cli *ZSClient) StopAllResourcesInIAM2Project(ctx context.Context, uuid string, params param.StopAllResourcesInIAM2ProjectParam) (*view.StopAllResourcesInIAM2ProjectEventView, error) {
	resp := view.StopAllResourcesInIAM2ProjectEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/projects", uuid, "", map[string]interface{}{
		"stopAllResourcesInIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmNetworkConfig updates VmNetworkConfig
func (cli *ZSClient) UpdateVmNetworkConfig(ctx context.Context, vmInstanceUuid string, params param.UpdateVmNetworkConfigParam) (*view.UpdateVmNetworkConfigEventView, error) {
	resp := view.UpdateVmNetworkConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"updateVmNetworkConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveTicketTypesFromTicketFlowCollection removes TicketTypesFromTicketFlowCollection
func (cli *ZSClient) RemoveTicketTypesFromTicketFlowCollection(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/tickets/flow-collections", uuid, string(deleteMode))
}

// DeleteEcsVSwitchRemote deletes EcsVSwitchRemote
func (cli *ZSClient) DeleteEcsVSwitchRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/vswitch/remote", uuid, string(deleteMode))
}

// SetVmStaticIp operates on VmStaticIp
func (cli *ZSClient) SetVmStaticIp(ctx context.Context, vmInstanceUuid string, params param.SetVmStaticIpParam) (*view.SetVmStaticIpEventView, error) {
	resp := view.SetVmStaticIpEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"setVmStaticIp": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmSshKey gets VmSshKey by uuid
func (cli *ZSClient) GetVmSshKey(ctx context.Context, uuid string) (*view.GetVmSshKeyView, error) {
	var resp view.GetVmSshKeyView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmGuestToolsInfo gets VmGuestToolsInfo by uuid
func (cli *ZSClient) GetVmGuestToolsInfo(ctx context.Context, uuid string) (*view.GetVmGuestToolsInfoView, error) {
	var resp view.GetVmGuestToolsInfoView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateDiskOfferingUserConfig operates on DiskOfferingUserConfig
func (cli *ZSClient) ValidateDiskOfferingUserConfig(ctx context.Context, params param.ValidateDiskOfferingUserConfigParam) (*view.ValidateDiskOfferingUserConfigEventView, error) {
	resp := view.ValidateDiskOfferingUserConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/accounts/actions", "", "", map[string]interface{}{
		"validateDiskOfferingUserConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcVpnGatewayLocal deletes VpcVpnGatewayLocal
func (cli *ZSClient) DeleteVpcVpnGatewayLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/vpn-gateway", uuid, string(deleteMode))
}

// SetVmRDP operates on VmRDP
func (cli *ZSClient) SetVmRDP(ctx context.Context, uuid string, params param.SetVmRDPParam) (*view.SetVmRDPEventView, error) {
	resp := view.SetVmRDPEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmRDP": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunSchedulerTrigger operates on RunSchedulerTrigger
func (cli *ZSClient) RunSchedulerTrigger(ctx context.Context, uuid string, params param.RunSchedulerTriggerParam) (*view.RunSchedulerTriggerEventView, error) {
	resp := view.RunSchedulerTriggerEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/scheduler/triggers", uuid, "", map[string]interface{}{
		"runSchedulerTrigger": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAliyunVpcVirtualRouterEntryRemote creates AliyunVpcVirtualRouterEntryRemote
func (cli *ZSClient) CreateAliyunVpcVirtualRouterEntryRemote(ctx context.Context, params param.CreateAliyunVpcVirtualRouterEntryRemoteParam) (*view.VpcVirtualRouteEntryInventoryView, error) {
	resp := view.VpcVirtualRouteEntryInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/route-entry", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOnHost operates on PowerOnHost
func (cli *ZSClient) PowerOnHost(ctx context.Context, uuid string, params param.PowerOnHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/power", uuid, "", map[string]interface{}{
		"powerOnHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunSnapshotFromRemote deletes AliyunSnapshotFromRemote
func (cli *ZSClient) DeleteAliyunSnapshotFromRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/snapshot", uuid, string(deleteMode))
}

// RemoveCertificateFromLoadBalancerListener removes CertificateFromLoadBalancerListener
func (cli *ZSClient) RemoveCertificateFromLoadBalancerListener(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/listeners", uuid, string(deleteMode))
}

// GetPortForwardingAttachableVmNics gets PortForwardingAttachableVmNics by uuid
func (cli *ZSClient) GetPortForwardingAttachableVmNics(ctx context.Context, uuid string) (*view.GetPortForwardingAttachableVmNicsView, error) {
	var resp view.GetPortForwardingAttachableVmNicsView
	if err := cli.GetWithRespKey(ctx, "v1/port-forwarding", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveRendezvousPointFromMulticastRouter removes RendezvousPointFromMulticastRouter
func (cli *ZSClient) RemoveRendezvousPointFromMulticastRouter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/multicast/virtual-routers", uuid, string(deleteMode))
}

// AddIAM2VirtualIDsToProject adds IAM2VirtualIDsToProject
func (cli *ZSClient) AddIAM2VirtualIDsToProject(ctx context.Context, params param.AddIAM2VirtualIDsToProjectParam) (*view.AddIAM2VirtualIDsToProjectEventView, error) {
	resp := view.AddIAM2VirtualIDsToProjectEventView{}
	if err := cli.Post(ctx, "v1/iam2/projects/{projectUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SubscribeEvent operates on SubscribeEvent
func (cli *ZSClient) SubscribeEvent(ctx context.Context, params param.SubscribeEventParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/events/subscriptions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageCandidatesForVolumeMigration gets PrimaryStorageCandidatesForVolumeMigration by uuid
func (cli *ZSClient) GetPrimaryStorageCandidatesForVolumeMigration(ctx context.Context, uuid string) (*view.GetPrimaryStorageCandidatesForVolumeMigrationView, error) {
	var resp view.GetPrimaryStorageCandidatesForVolumeMigrationView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradeBackupStorageCdpTasks operates on UpgradeBackupStorageCdpTasks
func (cli *ZSClient) UpgradeBackupStorageCdpTasks(ctx context.Context, backupStorageUuid string, params param.UpgradeBackupStorageCdpTasksParam) (*view.UpgradeBackupStorageCdpTasksEventView, error) {
	resp := view.UpgradeBackupStorageCdpTasksEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/cdp-task/upgrade", backupStorageUuid, "", map[string]interface{}{
		"upgradeBackupStorageCdpTasks": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVxlanL2Network deletes VxlanL2Network
func (cli *ZSClient) DeleteVxlanL2Network(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l2-networks/vxlan", uuid, string(deleteMode))
}

// RemoveVmFromAffinityGroup removes VmFromAffinityGroup
func (cli *ZSClient) RemoveVmFromAffinityGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/affinity-groups", uuid, string(deleteMode))
}

// SetVolumeIoThreadPin operates on VolumeIoThreadPin
func (cli *ZSClient) SetVolumeIoThreadPin(ctx context.Context, uuid string, params param.SetVolumeIoThreadPinParam) (*view.SetVolumeIoThreadPinEventView, error) {
	resp := view.SetVolumeIoThreadPinEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/volumes", uuid, "", map[string]interface{}{
		"setVolumeIoThreadPin": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePriorityConfig updates PriorityConfig
func (cli *ZSClient) UpdatePriorityConfig(ctx context.Context, uuid string, params param.UpdatePriorityConfigParam) (*view.UpdatePriorityConfigEventView, error) {
	resp := view.UpdatePriorityConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-priority-config", uuid, "", map[string]interface{}{
		"updatePriorityConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IdentifyHost operates on IdentifyHost
func (cli *ZSClient) IdentifyHost(ctx context.Context, uuid string, params param.IdentifyHostParam) (*view.IdentifyHostEventView, error) {
	resp := view.IdentifyHostEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/kvm", uuid, "", map[string]interface{}{
		"identifyHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromVolumeBackup creates RootVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeBackup(ctx context.Context, params param.CreateRootVolumeTemplateFromVolumeBackupParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(ctx, "v1/images/root-volume-templates/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckFirewallRuleConfigFile operates on FirewallRuleConfigFile
func (cli *ZSClient) CheckFirewallRuleConfigFile(ctx context.Context, params param.CheckFirewallRuleConfigFileParam) (*view.CheckFirewallRuleConfigFileView, error) {
	resp := view.CheckFirewallRuleConfigFileView{}
	if err := cli.Post(ctx, "v1/vpcfirewalls/rules/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmConsoleAddress gets VmConsoleAddress by uuid
func (cli *ZSClient) GetVmConsoleAddress(ctx context.Context, uuid string) (*view.GetVmConsoleAddressView, error) {
	var resp view.GetVmConsoleAddressView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLoadBalancerListenerACLEntries gets LoadBalancerListenerACLEntries by uuid
func (cli *ZSClient) GetLoadBalancerListenerACLEntries(ctx context.Context) (*view.GetLoadBalancerListenerACLEntriesView, error) {
	var resp view.GetLoadBalancerListenerACLEntriesView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers/listeners/access-control-lists/entries", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHostIommuState updates HostIommuState
func (cli *ZSClient) UpdateHostIommuState(ctx context.Context, uuid string, params param.UpdateHostIommuStateParam) (*view.UpdateHostIommuStateEventView, error) {
	resp := view.UpdateHostIommuStateEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/pci-device/hosts", uuid, "", map[string]interface{}{
		"updateHostIommuState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsubscribeEvent operates on UnsubscribeEvent
func (cli *ZSClient) UnsubscribeEvent(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/events/subscriptions", uuid, string(deleteMode))
}

// CreateObservabilityServer creates ObservabilityServer
func (cli *ZSClient) CreateObservabilityServer(ctx context.Context, params param.CreateObservabilityServerParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/observability-servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveMonFromCephPrimaryStorage removes MonFromCephPrimaryStorage
func (cli *ZSClient) RemoveMonFromCephPrimaryStorage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/primary-storage/ceph", uuid, string(deleteMode))
}

// GetVmsSchedulingStateFromSchedulingRule gets VmsSchedulingStateFromSchedulingRule by uuid
func (cli *ZSClient) GetVmsSchedulingStateFromSchedulingRule(ctx context.Context, params param.GetVmsSchedulingStateFromSchedulingRuleParam) (*view.GetVmsSchedulingStateFromSchedulingRuleView, error) {
	resp := view.GetVmsSchedulingStateFromSchedulingRuleView{}
	if err := cli.Post(ctx, "v1/get/vms/schedulingState/from/SchedulingRule", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAlarmState changes AlarmState
func (cli *ZSClient) ChangeAlarmState(ctx context.Context, uuid string, params param.ChangeAlarmStateParam) (*view.AlarmInventoryView, error) {
	resp := view.AlarmInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/alarms", uuid, "", map[string]interface{}{
		"changeAlarmState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLocalStorageHostDiskCapacity gets LocalStorageHostDiskCapacity by uuid
func (cli *ZSClient) GetLocalStorageHostDiskCapacity(ctx context.Context, uuid string) (*view.GetLocalStorageHostDiskCapacityView, error) {
	var resp view.GetLocalStorageHostDiskCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/local-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmSshKey deletes VmSshKey
func (cli *ZSClient) DeleteVmSshKey(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// GetPolicyRouteRuleSetFromVirtualRouter gets PolicyRouteRuleSetFromVirtualRouter by uuid
func (cli *ZSClient) GetPolicyRouteRuleSetFromVirtualRouter(ctx context.Context, uuid string) (*view.GetPolicyRouteRuleSetFromVirtualRouterView, error) {
	var resp view.GetPolicyRouteRuleSetFromVirtualRouterView
	if err := cli.GetWithRespKey(ctx, "v1/policy-routes/rulesets/virtualrouter", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVxlanPoolRemoteVtep deletes VxlanPoolRemoteVtep
func (cli *ZSClient) DeleteVxlanPoolRemoteVtep(ctx context.Context, l2NetworkUuid string, clusterUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/l2-networks", l2NetworkUuid, fmt.Sprintf("clusters/%s/delete/remote-vtep-ip", clusterUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// RemoveAttributesFromIAM2Project removes AttributesFromIAM2Project
func (cli *ZSClient) RemoveAttributesFromIAM2Project(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects", uuid, string(deleteMode))
}

// RecoverDataVolume operates on DataVolume
func (cli *ZSClient) RecoverDataVolume(ctx context.Context, uuid string, params param.RecoverDataVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/volumes", uuid, "", map[string]interface{}{
		"recoverDataVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveIAM2VirtualIDsFromGroup removes IAM2VirtualIDsFromGroup
func (cli *ZSClient) RemoveIAM2VirtualIDsFromGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects/groups", uuid, string(deleteMode))
}

// QueryEventRecord queries EventRecord list
func (cli *ZSClient) QueryEventRecord(ctx context.Context, params *param.QueryParam) ([]view.EventRecordsInventoryView, error) {
	var resp []view.EventRecordsInventoryView
	return resp, cli.List(ctx, "v1/zwatch/event-records", params, &resp)
}

func (cli *ZSClient) GetEventRecord(ctx context.Context, uuid string) (*view.EventRecordsInventoryView, error) {
	var resp view.EventRecordsInventoryView
	if err := cli.Get(ctx, "v1/zwatch/event-records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEventRecord Pagination
func (cli *ZSClient) PageEventRecord(ctx context.Context, params *param.QueryParam) ([]view.EventRecordsInventoryView, int, error) {
	var eventRecords []view.EventRecordsInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/event-records", params, &eventRecords)
	return eventRecords, total, err
}

// AttachBareMetal2ProvisionNetworkToCluster operates on BareMetal2ProvisionNetworkToCluster
func (cli *ZSClient) AttachBareMetal2ProvisionNetworkToCluster(ctx context.Context, params param.AttachBareMetal2ProvisionNetworkToClusterParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	resp := view.BareMetal2ProvisionNetworkInventoryView{}
	if err := cli.Post(ctx, "v1/baremetal2/clusters/{clusterUuid}/provision-networks/{networkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIAM2LdapBinding queries IAM2LdapBinding list
func (cli *ZSClient) QueryIAM2LdapBinding(ctx context.Context, params *param.QueryParam) ([]view.LdapResourceRefInventoryView, error) {
	var resp []view.LdapResourceRefInventoryView
	return resp, cli.List(ctx, "v1/iam2/ldap/bindings", params, &resp)
}

func (cli *ZSClient) GetIAM2LdapBinding(ctx context.Context, uuid string) (*view.LdapResourceRefInventoryView, error) {
	var resp view.LdapResourceRefInventoryView
	if err := cli.Get(ctx, "v1/iam2/ldap/bindings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2LdapBinding Pagination
func (cli *ZSClient) PageIAM2LdapBinding(ctx context.Context, params *param.QueryParam) ([]view.LdapResourceRefInventoryView, int, error) {
	var iAM2LdapBindings []view.LdapResourceRefInventoryView
	total, err := cli.Page(ctx, "v1/iam2/ldap/bindings", params, &iAM2LdapBindings)
	return iAM2LdapBindings, total, err
}

// ProvisionSlbInstance operates on ProvisionSlbInstance
func (cli *ZSClient) ProvisionSlbInstance(ctx context.Context, uuid string, params param.ProvisionSlbInstanceParam) (*view.SlbGroupInventoryView, error) {
	resp := view.SlbGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/load-balancers/slb/instances", uuid, "", map[string]interface{}{
		"provisionSlbInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmUserDefinedXmlHookScript operates on VmUserDefinedXmlHookScript
func (cli *ZSClient) SetVmUserDefinedXmlHookScript(ctx context.Context, vmInstanceUuid string, params param.SetVmUserDefinedXmlHookScriptParam) (*view.SetVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.SetVmUserDefinedXmlHookScriptEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"setVmUserDefinedXmlHookScript": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachProvisionNicFromBonding operates on ProvisionNicFromBonding
func (cli *ZSClient) DetachProvisionNicFromBonding(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/baremetal2/bm-instances/bm2-bondings", uuid, string(deleteMode))
}

// GetHostAllocatorStrategies gets HostAllocatorStrategies by uuid
func (cli *ZSClient) GetHostAllocatorStrategies(ctx context.Context) (*view.GetHostAllocatorStrategiesView, error) {
	var resp view.GetHostAllocatorStrategiesView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/allocators/strategies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInterfaceServiceTypeStatistic gets InterfaceServiceTypeStatistic by uuid
func (cli *ZSClient) GetInterfaceServiceTypeStatistic(ctx context.Context) (*view.GetInterfaceServiceTypeStatisticView, error) {
	var resp view.GetInterfaceServiceTypeStatisticView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/hosts-network-interfaces/service-type-statistic", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartConnectionBetweenAliyunRouterInterface starts ConnectionBetweenAliyunRouterInterface
func (cli *ZSClient) StartConnectionBetweenAliyunRouterInterface(ctx context.Context, vbrInterfaceUuid string, params param.StartConnectionBetweenAliyunRouterInterfaceParam) (*view.AliyunRouterInterfaceInventoryView, error) {
	resp := view.AliyunRouterInterfaceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/router-interface", vbrInterfaceUuid, "", map[string]interface{}{
		"startConnectionBetweenAliyunRouterInterface": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteModels deletes Models
func (cli *ZSClient) DeleteModels(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ai/models", uuid, string(deleteMode))
}

// ListVmsFromSchedulingState operates on ListVmsFromSchedulingState
func (cli *ZSClient) ListVmsFromSchedulingState(ctx context.Context, params param.ListVmsFromSchedulingStateParam) (*view.ListVmsFromSchedulingStateView, error) {
	resp := view.ListVmsFromSchedulingStateView{}
	if err := cli.Post(ctx, "v1/list/vms/from/executeState", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromVolumeSnapshot creates RootVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeSnapshot(ctx context.Context, params param.CreateRootVolumeTemplateFromVolumeSnapshotParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(ctx, "v1/images/root-volume-templates/from/volume-snapshots/{snapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AllocateHostResource operates on HostResource
func (cli *ZSClient) AllocateHostResource(ctx context.Context, params param.AllocateHostResourceParam) (*view.AllocateHostResourceEventView, error) {
	resp := view.AllocateHostResourceEventView{}
	if err := cli.Post(ctx, "v1/hosts/{uuid}/allocate-resource", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateLdapEntryForBinding gets CandidateLdapEntryForBinding by uuid
func (cli *ZSClient) GetCandidateLdapEntryForBinding(ctx context.Context) (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.GetWithRespKey(ctx, "v1/ldap/entries/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckElaborationContent operates on ElaborationContent
func (cli *ZSClient) CheckElaborationContent(ctx context.Context, params param.CheckElaborationContentParam) (*view.CheckElaborationContentView, error) {
	resp := view.CheckElaborationContentView{}
	if err := cli.Post(ctx, "v1/errorcode/elaborations/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmConsolePassword deletes VmConsolePassword
func (cli *ZSClient) DeleteVmConsolePassword(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// CreateVmBackup creates VmBackup
func (cli *ZSClient) CreateVmBackup(ctx context.Context, params param.CreateVmBackupParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.Post(ctx, "v1/volumes/{rootVolumeUuid}/vm-backups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmBackupAsync Async
func (cli *ZSClient) CreateVmBackupAsync(ctx context.Context, params param.CreateVmBackupParam) (string, error) {

	resource := "v1/volumes/{rootVolumeUuid}/vm-backups"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetPrimaryStorageLicenseInfo gets PrimaryStorageLicenseInfo by uuid
func (cli *ZSClient) GetPrimaryStorageLicenseInfo(ctx context.Context, uuid string) (*view.GetPrimaryStorageLicenseInfoView, error) {
	var resp view.GetPrimaryStorageLicenseInfoView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEncryptedField gets EncryptedField by uuid
func (cli *ZSClient) GetEncryptedField(ctx context.Context) (*view.GetEncryptedFieldView, error) {
	var resp view.GetEncryptedFieldView
	if err := cli.GetWithRespKey(ctx, "v1/encrypted/fields", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanInvalidLdapBinding operates on InvalidLdapBinding
func (cli *ZSClient) CleanInvalidLdapBinding(ctx context.Context) (*view.AccountInventoryView, error) {
	resp := view.AccountInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/ldap/bindings/actions", "", "", map[string]interface{}{
		"cleanInvalidLdapBinding": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachBaremetalPxeServerToCluster operates on BaremetalPxeServerToCluster
func (cli *ZSClient) AttachBaremetalPxeServerToCluster(ctx context.Context, params param.AttachBaremetalPxeServerToClusterParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.Post(ctx, "v1/clusters/{clusterUuid}/pxeservers/{pxeServerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmStartingCandidateClustersHosts gets VmStartingCandidateClustersHosts by uuid
func (cli *ZSClient) GetVmStartingCandidateClustersHosts(ctx context.Context, uuid string) (*view.GetVmStartingCandidateClustersHostsView, error) {
	var resp view.GetVmStartingCandidateClustersHostsView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoverVmBackupFromImageStoreBackupStorage operates on VmBackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverVmBackupFromImageStoreBackupStorage(ctx context.Context, groupUuid string, params param.RecoverVmBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-backups", groupUuid, "", map[string]interface{}{
		"recoverVmBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachIAM2ProjectFromIAM2Organization operates on IAM2ProjectFromIAM2Organization
func (cli *ZSClient) DetachIAM2ProjectFromIAM2Organization(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects", uuid, string(deleteMode))
}

// DiscoverExternalPrimaryStorage operates on DiscoverExternalPrimaryStorage
func (cli *ZSClient) DiscoverExternalPrimaryStorage(ctx context.Context, params param.DiscoverExternalPrimaryStorageParam) (*view.ExternalPrimaryStorageInventoryView, error) {
	resp := view.ExternalPrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/addon/discover", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeIoThreadPin gets VolumeIoThreadPin by uuid
func (cli *ZSClient) GetVolumeIoThreadPin(ctx context.Context, uuid string) (*view.GetVolumeIoThreadPinView, error) {
	var resp view.GetVolumeIoThreadPinView
	if err := cli.GetWithRespKey(ctx, "v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConnectionAccessPointFromRemote gets ConnectionAccessPointFromRemote by uuid
func (cli *ZSClient) GetConnectionAccessPointFromRemote(ctx context.Context, uuid string) (*view.GetConnectionAccessPointFromRemoteView, error) {
	var resp view.GetConnectionAccessPointFromRemoteView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/aliyun/access-point", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedOspf gets VpcAttachedOspf by uuid
func (cli *ZSClient) GetVpcAttachedOspf(ctx context.Context, params param.GetVpcAttachedOspfParam) (*view.NetworkRouterAreaRefInventoryView, error) {
	resp := view.NetworkRouterAreaRefInventoryView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/attached-ospf", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOffHost operates on PowerOffHost
func (cli *ZSClient) PowerOffHost(ctx context.Context, params param.PowerOffHostParam) (*view.PowerOffHostEventView, error) {
	resp := view.PowerOffHostEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/power-off/actions", "", "", map[string]interface{}{
		"powerOffHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveIAM2VirtualIDGroupFromProjects removes IAM2VirtualIDGroupFromProjects
func (cli *ZSClient) RemoveIAM2VirtualIDGroupFromProjects(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects/groups", uuid, string(deleteMode))
}

// UpdateVmUserDefinedXmlHookScript updates VmUserDefinedXmlHookScript
func (cli *ZSClient) UpdateVmUserDefinedXmlHookScript(ctx context.Context, params param.UpdateVmUserDefinedXmlHookScriptParam) (*view.XmlHookInventoryView, error) {
	resp := view.XmlHookInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/xml-hook-script", "", "", map[string]interface{}{
		"updateVmUserDefinedXmlHookScript": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2ProjectContainerImageTags gets IAM2ProjectContainerImageTags by uuid
func (cli *ZSClient) GetIAM2ProjectContainerImageTags(ctx context.Context, projectId string, repositoryId string, imageName string) (*view.GetIAM2ProjectContainerImageTagsView, error) {
	var resp view.GetIAM2ProjectContainerImageTagsView
	err := cli.GetWithSpec(ctx, "v1/iam2/project", projectId, fmt.Sprintf("repository/%s/image/%s/tag", repositoryId, imageName), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunDiskFromRemote deletes AliyunDiskFromRemote
func (cli *ZSClient) DeleteAliyunDiskFromRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/disk", uuid, string(deleteMode))
}

// GetVersion gets Version by uuid
func (cli *ZSClient) GetVersion(ctx context.Context) (*view.GetVersionView, error) {
	resp := view.GetVersionView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/actions", "", "", map[string]interface{}{
		"getVersion": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateBackupStorageForCreatingImage gets CandidateBackupStorageForCreatingImage by uuid
func (cli *ZSClient) GetCandidateBackupStorageForCreatingImage(ctx context.Context) (*view.GetCandidateBackupStorageForCreatingImageView, error) {
	var resp view.GetCandidateBackupStorageForCreatingImageView
	if err := cli.GetWithRespKey(ctx, "v1/images/candidate-backup-storage", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachAutoScalingTemplateToGroup operates on AutoScalingTemplateToGroup
func (cli *ZSClient) AttachAutoScalingTemplateToGroup(ctx context.Context, params param.AttachAutoScalingTemplateToGroupParam) (*view.AutoScalingGroupInventoryView, error) {
	resp := view.AutoScalingGroupInventoryView{}
	if err := cli.Post(ctx, "v1/autoscaling/template/{uuid}/groups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCpuMemoryCapacity gets CpuMemoryCapacity by uuid
func (cli *ZSClient) GetCpuMemoryCapacity(ctx context.Context) (*view.GetCpuMemoryCapacityView, error) {
	var resp view.GetCpuMemoryCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/capacities/cpu-memory", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIntegrityResource adds IntegrityResource
func (cli *ZSClient) AddIntegrityResource(ctx context.Context, params param.AddIntegrityResourceParam) (*view.AddIntegrityResourceEventView, error) {
	resp := view.AddIntegrityResourceEventView{}
	if err := cli.Post(ctx, "v1/integrity/resource/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckVipPortAvailability operates on VipPortAvailability
func (cli *ZSClient) CheckVipPortAvailability(ctx context.Context, uuid string) (*view.CheckVipPortAvailabilityView, error) {
	var resp view.CheckVipPortAvailabilityView
	if err := cli.GetWithRespKey(ctx, "v1/vips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateClustersForAttachingL2Network gets CandidateClustersForAttachingL2Network by uuid
func (cli *ZSClient) GetCandidateClustersForAttachingL2Network(ctx context.Context, uuid string) (*view.GetCandidateClustersForAttachingL2NetworkView, error) {
	var resp view.GetCandidateClustersForAttachingL2NetworkView
	if err := cli.GetWithRespKey(ctx, "v1/l2-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckScsiLunClusterStatus operates on ScsiLunClusterStatus
func (cli *ZSClient) CheckScsiLunClusterStatus(ctx context.Context, uuid string, clusterUuid string, params param.CheckScsiLunClusterStatusParam) (*view.ScsiLunClusterStatusInventoryView, error) {
	resp := view.ScsiLunClusterStatusInventoryView{}
	err := cli.PutWithSpec(ctx, "v1/storage-devices/scsi-lun", uuid, fmt.Sprintf("cluster/%s", clusterUuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckBatchDataIntegrity operates on BatchDataIntegrity
func (cli *ZSClient) CheckBatchDataIntegrity(ctx context.Context) (*view.CheckBatchDataIntegrityView, error) {
	var resp view.CheckBatchDataIntegrityView
	if err := cli.GetWithRespKey(ctx, "v1/check/batch/data/integrity/", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAutoScalingGroupRemovalInstanceRule updates AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupRemovalInstanceRule(ctx context.Context, uuid string, params param.UpdateAutoScalingGroupRemovalInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/autoscaling/rules/removal-instance", uuid, "", map[string]interface{}{
		"updateAutoScalingGroupRemovalInstanceRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryHybridKeySecret queries HybridKeySecret list
func (cli *ZSClient) QueryHybridKeySecret(ctx context.Context, params *param.QueryParam) ([]view.HybridAccountInventoryView, error) {
	var resp []view.HybridAccountInventoryView
	return resp, cli.List(ctx, "v1/hybrid/hybrid/key", params, &resp)
}

func (cli *ZSClient) GetHybridKeySecret(ctx context.Context, uuid string) (*view.HybridAccountInventoryView, error) {
	var resp view.HybridAccountInventoryView
	if err := cli.Get(ctx, "v1/hybrid/hybrid/key", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHybridKeySecret Pagination
func (cli *ZSClient) PageHybridKeySecret(ctx context.Context, params *param.QueryParam) ([]view.HybridAccountInventoryView, int, error) {
	var hybridKeySecrets []view.HybridAccountInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/hybrid/key", params, &hybridKeySecrets)
	return hybridKeySecrets, total, err
}

// UploadFileToVm operates on UploadFileToVm
func (cli *ZSClient) UploadFileToVm(ctx context.Context, params param.UploadFileToVmParam) (*view.UploadFileToVmEventView, error) {
	resp := view.UploadFileToVmEventView{}
	if err := cli.Post(ctx, "v1/upload-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcVpnGatewayFromLocal queries VpcVpnGatewayFromLocal list
func (cli *ZSClient) QueryVpcVpnGatewayFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVpnGatewayInventoryView, error) {
	var resp []view.VpcVpnGatewayInventoryView
	return resp, cli.List(ctx, "v1/hybrid/vpc-vpn", params, &resp)
}

func (cli *ZSClient) GetVpcVpnGatewayFromLocal(ctx context.Context, uuid string) (*view.VpcVpnGatewayInventoryView, error) {
	var resp view.VpcVpnGatewayInventoryView
	if err := cli.Get(ctx, "v1/hybrid/vpc-vpn", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcVpnGatewayFromLocal Pagination
func (cli *ZSClient) PageVpcVpnGatewayFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVpnGatewayInventoryView, int, error) {
	var vpcVpnGatewayFromLocals []view.VpcVpnGatewayInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/vpc-vpn", params, &vpcVpnGatewayFromLocals)
	return vpcVpnGatewayFromLocals, total, err
}

// ChangeL3NetworkDhcpIpAddress changes L3NetworkDhcpIpAddress
func (cli *ZSClient) ChangeL3NetworkDhcpIpAddress(ctx context.Context, l3NetworkUuid string, params param.ChangeL3NetworkDhcpIpAddressParam) (*view.ChangeL3NetworkDhcpIpAddressEventView, error) {
	resp := view.ChangeL3NetworkDhcpIpAddressEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/l3-networks", l3NetworkUuid, "", map[string]interface{}{
		"changeL3NetworkDhcpIpAddress": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckVolumeSnapshotGroupAvailability operates on VolumeSnapshotGroupAvailability
func (cli *ZSClient) CheckVolumeSnapshotGroupAvailability(ctx context.Context) (*view.CheckVolumeSnapshotGroupAvailabilityView, error) {
	var resp view.CheckVolumeSnapshotGroupAvailabilityView
	if err := cli.GetWithRespKey(ctx, "v1/volume-snapshots/groups/availabilities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SsoClientPushData operates on SsoClientPushData
func (cli *ZSClient) SsoClientPushData(ctx context.Context, params param.SsoClientPushDataParam) (*view.SsoClientPushDataEventView, error) {
	resp := view.SsoClientPushDataEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/sso/resource/data/push", "", "", map[string]interface{}{
		"ssoClientPushData": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddEmailAddressToSNSEmailEndpoint adds EmailAddressToSNSEmailEndpoint
func (cli *ZSClient) AddEmailAddressToSNSEmailEndpoint(ctx context.Context, params param.AddEmailAddressToSNSEmailEndpointParam) (*view.SNSEmailAddressInventoryView, error) {
	resp := view.SNSEmailAddressInventoryView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/emails/email-addresses", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVmNicInSecurityGroup queries VmNicInSecurityGroup list
func (cli *ZSClient) QueryVmNicInSecurityGroup(ctx context.Context, params *param.QueryParam) ([]view.VmNicSecurityGroupRefInventoryView, error) {
	var resp []view.VmNicSecurityGroupRefInventoryView
	return resp, cli.List(ctx, "v1/security-groups/vm-instances/nics", params, &resp)
}

func (cli *ZSClient) GetVmNicInSecurityGroup(ctx context.Context, uuid string) (*view.VmNicSecurityGroupRefInventoryView, error) {
	var resp view.VmNicSecurityGroupRefInventoryView
	if err := cli.Get(ctx, "v1/security-groups/vm-instances/nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmNicInSecurityGroup Pagination
func (cli *ZSClient) PageVmNicInSecurityGroup(ctx context.Context, params *param.QueryParam) ([]view.VmNicSecurityGroupRefInventoryView, int, error) {
	var vmNicInSecurityGroups []view.VmNicSecurityGroupRefInventoryView
	total, err := cli.Page(ctx, "v1/security-groups/vm-instances/nics", params, &vmNicInSecurityGroups)
	return vmNicInSecurityGroups, total, err
}

// BackupDatabaseToPublicCloud operates on DatabaseToPublicCloud
func (cli *ZSClient) BackupDatabaseToPublicCloud(ctx context.Context, params param.BackupDatabaseToPublicCloudParam) (*view.BackupDatabaseToPublicCloudEventView, error) {
	resp := view.BackupDatabaseToPublicCloudEventView{}
	if err := cli.Post(ctx, "v1/hybrid/backup-mysql", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoveryImageFromImageStoreBackupStorage operates on yImageFromImageStoreBackupStorage
func (cli *ZSClient) RecoveryImageFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.RecoveryImageFromImageStoreBackupStorageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage", uuid, "", map[string]interface{}{
		"recoveryImageFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEventFromResourceStack queries EventFromResourceStack list
func (cli *ZSClient) QueryEventFromResourceStack(ctx context.Context, params *param.QueryParam) ([]view.CloudFormationStackEventInventoryView, error) {
	var resp []view.CloudFormationStackEventInventoryView
	return resp, cli.List(ctx, "v1/cloudformation/event", params, &resp)
}

func (cli *ZSClient) GetEventFromResourceStack(ctx context.Context, uuid string) (*view.CloudFormationStackEventInventoryView, error) {
	var resp view.CloudFormationStackEventInventoryView
	if err := cli.Get(ctx, "v1/cloudformation/event", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEventFromResourceStack Pagination
func (cli *ZSClient) PageEventFromResourceStack(ctx context.Context, params *param.QueryParam) ([]view.CloudFormationStackEventInventoryView, int, error) {
	var eventFromResourceStacks []view.CloudFormationStackEventInventoryView
	total, err := cli.Page(ctx, "v1/cloudformation/event", params, &eventFromResourceStacks)
	return eventFromResourceStacks, total, err
}

// RevertVmFromSnapshotGroup operates on VmFromSnapshotGroup
func (cli *ZSClient) RevertVmFromSnapshotGroup(ctx context.Context, uuid string, params param.RevertVmFromSnapshotGroupParam) (*view.RevertVmFromSnapshotGroupEventView, error) {
	resp := view.RevertVmFromSnapshotGroupEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/volume-snapshots/group", uuid, "", map[string]interface{}{
		"revertVmFromSnapshotGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachFirewallRuleSetFromL3 operates on FirewallRuleSetFromL3
func (cli *ZSClient) DetachFirewallRuleSetFromL3(ctx context.Context, params param.DetachFirewallRuleSetFromL3Param) (*view.DetachFirewallRuleSetFromL3EventView, error) {
	resp := view.DetachFirewallRuleSetFromL3EventView{}
	if err := cli.Post(ctx, "v1/vpcfirewalls/l3networks/{l3Uuid}/ruleSets/{ruleSetUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVmSchedulingRulesFromExecuteState operates on ListVmSchedulingRulesFromExecuteState
func (cli *ZSClient) ListVmSchedulingRulesFromExecuteState(ctx context.Context, params param.ListVmSchedulingRulesFromExecuteStateParam) (*view.ListVmSchedulingRulesFromExecuteStateView, error) {
	resp := view.ListVmSchedulingRulesFromExecuteStateView{}
	if err := cli.Post(ctx, "v1/list/vmSchedulingRules/from/conflict/state", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmUserDefinedXml operates on VmUserDefinedXml
func (cli *ZSClient) SetVmUserDefinedXml(ctx context.Context, vmInstanceUuid string, params param.SetVmUserDefinedXmlParam) (*view.SetVmUserDefinedXmlEventView, error) {
	resp := view.SetVmUserDefinedXmlEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"setVmUserDefinedXml": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageQga operates on ImageQga
func (cli *ZSClient) SetImageQga(ctx context.Context, uuid string, params param.SetImageQgaParam) (*view.SetImageQgaEventView, error) {
	resp := view.SetImageQgaEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/images", uuid, "", map[string]interface{}{
		"setImageQga": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVMsFromKVMHost operates on ListVMsFromKVMHost
func (cli *ZSClient) ListVMsFromKVMHost(ctx context.Context, params param.ListVMsFromKVMHostParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/v2v", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// TakeVmConsoleScreenshot operates on TakeVmConsoleScreenshot
func (cli *ZSClient) TakeVmConsoleScreenshot(ctx context.Context, uuid string, params param.TakeVmConsoleScreenshotParam) (*view.TakeVmConsoleScreenshotEventView, error) {
	resp := view.TakeVmConsoleScreenshotEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"takeVmConsoleScreenshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVRouterNetworksFromOspfArea removes VRouterNetworksFromOspfArea
func (cli *ZSClient) RemoveVRouterNetworksFromOspfArea(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/routerArea/networks", uuid, string(deleteMode))
}

// GetAliyunNasMountTargetRemote gets AliyunNasMountTargetRemote by uuid
func (cli *ZSClient) GetAliyunNasMountTargetRemote(ctx context.Context) (*view.GetAliyunNasMountTargetRemoteView, error) {
	var resp view.GetAliyunNasMountTargetRemoteView
	if err := cli.GetWithRespKey(ctx, "v1/nas/aliyun/mount/remote", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateImageGroupFromVmInstance creates ImageGroupFromVmInstance
func (cli *ZSClient) CreateImageGroupFromVmInstance(ctx context.Context, params param.CreateImageGroupFromVmInstanceParam) (*view.ImageGroupInventoryView, error) {
	resp := view.ImageGroupInventoryView{}
	if err := cli.Post(ctx, "v1/images/groups/from/vm-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// TerminateVirtualBorderRouterRemote operates on TerminateVirtualBorderRouterRemote
func (cli *ZSClient) TerminateVirtualBorderRouterRemote(ctx context.Context, uuid string, params param.TerminateVirtualBorderRouterRemoteParam) (*view.TerminateVirtualBorderRouterRemoteEventView, error) {
	resp := view.TerminateVirtualBorderRouterRemoteEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/border-router", uuid, "", map[string]interface{}{
		"terminateVirtualBorderRouterRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmBackup deletes VmBackup
func (cli *ZSClient) DeleteVmBackup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-backups", uuid, string(deleteMode))
}

// SetVmSecurityLevel operates on VmSecurityLevel
func (cli *ZSClient) SetVmSecurityLevel(ctx context.Context, uuid string, params param.SetVmSecurityLevelParam) (*view.SetVmSecurityLevelEventView, error) {
	resp := view.SetVmSecurityLevelEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmSecurityLevel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveMdevDeviceSpecFromVmInstance removes MdevDeviceSpecFromVmInstance
func (cli *ZSClient) RemoveMdevDeviceSpecFromVmInstance(ctx context.Context, mdevSpecUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/mdev-device-specs", mdevSpecUuid, fmt.Sprintf("vm-instances/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// SyncVolumeSize operates on VolumeSize
func (cli *ZSClient) SyncVolumeSize(ctx context.Context, uuid string, params param.SyncVolumeSizeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/volumes", uuid, "", map[string]interface{}{
		"syncVolumeSize": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTrashOnBackupStorage gets TrashOnBackupStorage by uuid
func (cli *ZSClient) GetTrashOnBackupStorage(ctx context.Context) (*view.GetTrashOnBackupStorageView, error) {
	var resp view.GetTrashOnBackupStorageView
	if err := cli.GetWithRespKey(ctx, "v1/backup-storage/trash", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeDiskOfferingState changes DiskOfferingState
func (cli *ZSClient) ChangeDiskOfferingState(ctx context.Context, uuid string, params param.ChangeDiskOfferingStateParam) (*view.DiskOfferingInventoryView, error) {
	resp := view.DiskOfferingInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/disk-offerings", uuid, "", map[string]interface{}{
		"changeDiskOfferingState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RequestConsoleAccess operates on RequestConsoleAccess
func (cli *ZSClient) RequestConsoleAccess(ctx context.Context, params param.RequestConsoleAccessParam) (*view.ConsoleInventoryView, error) {
	resp := view.ConsoleInventoryView{}
	if err := cli.Post(ctx, "v1/consoles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeIAM2VirtualIDGroupState changes IAM2VirtualIDGroupState
func (cli *ZSClient) ChangeIAM2VirtualIDGroupState(ctx context.Context, uuid string, params param.ChangeIAM2VirtualIDGroupStateParam) (*view.IAM2VirtualIDGroupInventoryView, error) {
	resp := view.IAM2VirtualIDGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/projects/groups", uuid, "", map[string]interface{}{
		"changeIAM2VirtualIDGroupState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEventData updates EventData
func (cli *ZSClient) UpdateEventData(ctx context.Context, params param.UpdateEventDataParam) (*view.UpdateEventDataEventView, error) {
	resp := view.UpdateEventDataEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/events/actions", "", "", map[string]interface{}{
		"updateEventData": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncHybridEipFromRemote operates on HybridEipFromRemote
func (cli *ZSClient) SyncHybridEipFromRemote(ctx context.Context, dataCenterUuid string, params param.SyncHybridEipFromRemoteParam) (*view.HybridEipAddressInventoryView, error) {
	resp := view.HybridEipAddressInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/eip", dataCenterUuid, "", map[string]interface{}{
		"syncHybridEipFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunRouteEntryRemote deletes AliyunRouteEntryRemote
func (cli *ZSClient) DeleteAliyunRouteEntryRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/route-entry", uuid, string(deleteMode))
}

// UngenerateSriovPciDevices operates on UngenerateSriovPciDevices
func (cli *ZSClient) UngenerateSriovPciDevices(ctx context.Context, pciDeviceUuid string, params param.UngenerateSriovPciDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/pci-devices", pciDeviceUuid, "", map[string]interface{}{
		"ungenerateSriovPciDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmStaticIp deletes VmStaticIp
func (cli *ZSClient) DeleteVmStaticIp(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// AttachMonitorTriggerActionToTrigger operates on MonitorTriggerActionToTrigger
func (cli *ZSClient) AttachMonitorTriggerActionToTrigger(ctx context.Context, params param.AttachMonitorTriggerActionToTriggerParam) (*view.AttachMonitorTriggerActionToTriggerEventView, error) {
	resp := view.AttachMonitorTriggerActionToTriggerEventView{}
	if err := cli.Post(ctx, "v1/monitoring/triggers/{triggerUuid}/trigger-actions/{actionUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAliyunNasFileSystemRemote gets AliyunNasFileSystemRemote by uuid
func (cli *ZSClient) GetAliyunNasFileSystemRemote(ctx context.Context) (*view.GetAliyunNasFileSystemRemoteView, error) {
	var resp view.GetAliyunNasFileSystemRemoteView
	if err := cli.GetWithRespKey(ctx, "v1/nas/aliyun/remote", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOrganizationQuota updates OrganizationQuota
func (cli *ZSClient) UpdateOrganizationQuota(ctx context.Context, params param.UpdateOrganizationQuotaParam) (*view.QuotaInventoryView, error) {
	resp := view.QuotaInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/Organization/quotas/actions", "", "", map[string]interface{}{
		"updateOrganizationQuota": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePreconfigurationTemplateState changes PreconfigurationTemplateState
func (cli *ZSClient) ChangePreconfigurationTemplateState(ctx context.Context, uuid string, params param.ChangePreconfigurationTemplateStateParam) (*view.PreconfigurationTemplateInventoryView, error) {
	resp := view.PreconfigurationTemplateInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal/preconfigurations", uuid, "", map[string]interface{}{
		"changePreconfigurationTemplateState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetOrganizationSupervisor operates on OrganizationSupervisor
func (cli *ZSClient) SetOrganizationSupervisor(ctx context.Context, uuid string, params param.SetOrganizationSupervisorParam) (*view.SetOrganizationSupervisorEventView, error) {
	resp := view.SetOrganizationSupervisorEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/organizations", uuid, "", map[string]interface{}{
		"setOrganizationSupervisor": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL3NetworksToIPsecConnection operates on L3NetworksToIPsecConnection
func (cli *ZSClient) AttachL3NetworksToIPsecConnection(ctx context.Context, params param.AttachL3NetworksToIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.Post(ctx, "v1/ipsec/{uuid}/l3networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExecuteGuestVmScript operates on ExecuteGuestVmScript
func (cli *ZSClient) ExecuteGuestVmScript(ctx context.Context, uuid string, params param.ExecuteGuestVmScriptParam) (*view.GuestVmScriptExecutedRecordInventoryView, error) {
	resp := view.GuestVmScriptExecutedRecordInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/scripts", uuid, "", map[string]interface{}{
		"executeGuestVmScript": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddNfsPrimaryStorage adds NfsPrimaryStorage
func (cli *ZSClient) AddNfsPrimaryStorage(ctx context.Context) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/nfs", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2ProjectContainerClusterCandidates gets IAM2ProjectContainerClusterCandidates by uuid
func (cli *ZSClient) GetIAM2ProjectContainerClusterCandidates(ctx context.Context) (*view.GetIAM2ProjectContainerClusterCandidatesView, error) {
	var resp view.GetIAM2ProjectContainerClusterCandidatesView
	if err := cli.GetWithRespKey(ctx, "v1/iam2/projects/container/cluster/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachTagToResources operates on TagToResources
func (cli *ZSClient) AttachTagToResources(ctx context.Context, params param.AttachTagToResourcesParam) (*view.AttachTagToResourcesEventView, error) {
	resp := view.AttachTagToResourcesEventView{}
	if err := cli.Post(ctx, "v1/tags/{tagUuid}/resources", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePrimaryStorageState changes PrimaryStorageState
func (cli *ZSClient) ChangePrimaryStorageState(ctx context.Context, uuid string, params param.ChangePrimaryStorageStateParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage", uuid, "", map[string]interface{}{
		"changePrimaryStorageState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedNetflow gets VpcAttachedNetflow by uuid
func (cli *ZSClient) GetVpcAttachedNetflow(ctx context.Context, params param.GetVpcAttachedNetflowParam) (*view.FlowMeterInventoryView, error) {
	resp := view.FlowMeterInventoryView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/attached-netflow", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAuditData gets AuditData by uuid
func (cli *ZSClient) GetAuditData(ctx context.Context) (*view.GetAuditDataView, error) {
	var resp view.GetAuditDataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/audits", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSpiceCertificates gets SpiceCertificates by uuid
func (cli *ZSClient) GetSpiceCertificates(ctx context.Context) (*view.GetSpiceCertificatesView, error) {
	var resp view.GetSpiceCertificatesView
	if err := cli.GetWithRespKey(ctx, "v1/spice/certificates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveUserFromGroup removes UserFromGroup
func (cli *ZSClient) RemoveUserFromGroup(ctx context.Context, groupUuid string, userUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/accounts/groups", groupUuid, fmt.Sprintf("users/%s", userUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DeleteEcsVpcRemote deletes EcsVpcRemote
func (cli *ZSClient) DeleteEcsVpcRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/vpc/remote", uuid, string(deleteMode))
}

// SyncDatabaseBackupFromImageStoreBackupStorage operates on DatabaseBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncDatabaseBackupFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.SyncDatabaseBackupFromImageStoreBackupStorageParam) (*view.DatabaseBackupInventoryView, error) {
	resp := view.DatabaseBackupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/database-backups", uuid, "", map[string]interface{}{
		"syncDatabaseBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFirewallIpSetTemplate deletes FirewallIpSetTemplate
func (cli *ZSClient) DeleteFirewallIpSetTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpcfirewalls/ipset/templates", uuid, string(deleteMode))
}

// SNSDingTalkTestConnection operates on DingTalkTestConnection
func (cli *ZSClient) SNSDingTalkTestConnection(ctx context.Context, params param.SNSDingTalkTestConnectionParam) (*view.SNSDingTalkTestConnectionEventView, error) {
	resp := view.SNSDingTalkTestConnectionEventView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/ding-talk/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportImageFromBackupStorage operates on ImageFromBackupStorage
func (cli *ZSClient) ExportImageFromBackupStorage(ctx context.Context, backupStorageUuid string, params param.ExportImageFromBackupStorageParam) (*view.ExportImageFromBackupStorageEventView, error) {
	resp := view.ExportImageFromBackupStorageEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage", backupStorageUuid, "", map[string]interface{}{
		"exportImageFromBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportImageFromBackupStorageAsync Async
func (cli *ZSClient) ExportImageFromBackupStorageAsync(ctx context.Context, params param.ExportImageFromBackupStorageParam) (string, error) {

	resource := "v1/backup-storage/{backupStorageUuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetModelCenterServices gets ModelCenterServices by uuid
func (cli *ZSClient) GetModelCenterServices(ctx context.Context) (*view.GetModelCenterServicesView, error) {
	var resp view.GetModelCenterServicesView
	if err := cli.GetWithRespKey(ctx, "v1/ai/model-centers/services", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallIpSetTemplate creates FirewallIpSetTemplate
func (cli *ZSClient) CreateFirewallIpSetTemplate(ctx context.Context, params param.CreateFirewallIpSetTemplateParam) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	resp := view.VpcFirewallIpSetTemplateInventoryView{}
	if err := cli.Post(ctx, "v1/vpcfirewalls/ipset/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachMonitorTriggerActionFromTrigger operates on MonitorTriggerActionFromTrigger
func (cli *ZSClient) DetachMonitorTriggerActionFromTrigger(ctx context.Context, triggerUuid string, actionUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/monitoring/triggers", triggerUuid, fmt.Sprintf("trigger-actions/%s", actionUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DetachPolicyRouteRuleSetFromL3 operates on PolicyRouteRuleSetFromL3
func (cli *ZSClient) DetachPolicyRouteRuleSetFromL3(ctx context.Context, ruleSetUuid string, l3Uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/policy-routes/rulesets", ruleSetUuid, fmt.Sprintf("l3networks/%s", l3Uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// CreateL2TfNetwork creates L2TfNetwork
func (cli *ZSClient) CreateL2TfNetwork(ctx context.Context, params param.CreateL2TfNetworkParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.Post(ctx, "v1/l2-networks/tf", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInterdependentL3NetworksImages gets InterdependentL3NetworksImages by uuid
func (cli *ZSClient) GetInterdependentL3NetworksImages(ctx context.Context) (*view.GetInterdependentL3NetworkImageView, error) {
	var resp view.GetInterdependentL3NetworkImageView
	if err := cli.GetWithRespKey(ctx, "v1/images-l3networks/dependencies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateVolumeSnapshotChain operates on VolumeSnapshotChain
func (cli *ZSClient) ValidateVolumeSnapshotChain(ctx context.Context, uuid string, params param.ValidateVolumeSnapshotChainParam) (*view.ValidateVolumeSnapshotChainEventView, error) {
	resp := view.ValidateVolumeSnapshotChainEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/volumes", uuid, "", map[string]interface{}{
		"validateVolumeSnapshotChain": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeHostNetworkInterfaceLldpMode changes HostNetworkInterfaceLldpMode
func (cli *ZSClient) ChangeHostNetworkInterfaceLldpMode(ctx context.Context, params param.ChangeHostNetworkInterfaceLldpModeParam) (*view.HostNetworkInterfaceLldpInventoryView, error) {
	resp := view.HostNetworkInterfaceLldpInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hostNetworkInterface/lldp/actions", "", "", map[string]interface{}{
		"changeHostNetworkInterfaceLldpMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGuestOsMetadata gets GuestOsMetadata by uuid
func (cli *ZSClient) GetGuestOsMetadata(ctx context.Context) (*view.GetGuestOsMetadataView, error) {
	var resp view.GetGuestOsMetadataView
	if err := cli.GetWithRespKey(ctx, "v1/guest-os/metadata", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicsForLoadBalancerServerGroup gets CandidateVmNicsForLoadBalancerServerGroup by uuid
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancerServerGroup(ctx context.Context) (*view.GetCandidateVmNicsForLoadBalancerServerGroupView, error) {
	var resp view.GetCandidateVmNicsForLoadBalancerServerGroupView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers/servergroups/candidate-nics", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachIscsiServerToCluster operates on IscsiServerToCluster
func (cli *ZSClient) AttachIscsiServerToCluster(ctx context.Context, params param.AttachIscsiServerToClusterParam) (*view.IscsiServerInventoryView, error) {
	resp := view.IscsiServerInventoryView{}
	if err := cli.Post(ctx, "v1/clusters/{clusterUuid}/storage-devices/iscsi/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachRoleToAccount operates on RoleToAccount
func (cli *ZSClient) AttachRoleToAccount(ctx context.Context, params param.AttachRoleToAccountParam) (*view.AttachRoleToAccountEventView, error) {
	resp := view.AttachRoleToAccountEventView{}
	if err := cli.Post(ctx, "v1/identities/accounts/{accountUuid}/roles/{roleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryBuildApp queries BuildApp list
func (cli *ZSClient) QueryBuildApp(ctx context.Context, params *param.QueryParam) ([]view.BuildApplicationInventoryView, error) {
	var resp []view.BuildApplicationInventoryView
	return resp, cli.List(ctx, "v1/appcenter/buildapp", params, &resp)
}

func (cli *ZSClient) GetBuildApp(ctx context.Context, uuid string) (*view.BuildApplicationInventoryView, error) {
	var resp view.BuildApplicationInventoryView
	if err := cli.Get(ctx, "v1/appcenter/buildapp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBuildApp Pagination
func (cli *ZSClient) PageBuildApp(ctx context.Context, params *param.QueryParam) ([]view.BuildApplicationInventoryView, int, error) {
	var buildApps []view.BuildApplicationInventoryView
	total, err := cli.Page(ctx, "v1/appcenter/buildapp", params, &buildApps)
	return buildApps, total, err
}

// AttachIsoToVmInstance operates on IsoToVmInstance
func (cli *ZSClient) AttachIsoToVmInstance(ctx context.Context, params param.AttachIsoToVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/{vmInstanceUuid}/iso/{isoUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVRouterRouterId operates on VRouterRouterId
func (cli *ZSClient) SetVRouterRouterId(ctx context.Context, params param.SetVRouterRouterIdParam) (*view.SetVRouterRouterIdEventView, error) {
	resp := view.SetVRouterRouterIdEventView{}
	if err := cli.Post(ctx, "v1/routerArea/{vRouterUuid}/routerid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExpungeVmUserDefinedXmlHookScript operates on VmUserDefinedXmlHookScript
func (cli *ZSClient) ExpungeVmUserDefinedXmlHookScript(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances/xml-hook-script", uuid, string(deleteMode))
}

// DeleteCdpTaskData deletes CdpTaskData
func (cli *ZSClient) DeleteCdpTaskData(uuid string, params param.DeleteCdpTaskDataParam) (*view.DeleteCdpTaskDataEventView, error) {
	resp := view.DeleteCdpTaskDataEventView{}
	if err := cli.PostWithRespKey(fmt.Sprintf("v1/cdp-task/%s/data", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckApiPermission operates on ApiPermission
func (cli *ZSClient) CheckApiPermission(ctx context.Context, params param.CheckApiPermissionParam) (*view.CheckApiPermissionView, error) {
	resp := view.CheckApiPermissionView{}
	if err := cli.PutWithRespKey(ctx, "v1/accounts/permissions/actions", "", "", map[string]interface{}{
		"checkApiPermission": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTextTemplateArg gets TextTemplateArg by uuid
func (cli *ZSClient) GetTextTemplateArg(ctx context.Context) (*view.GetTextTemplateArgView, error) {
	var resp view.GetTextTemplateArgView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/textTemplateArg", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFirewall deletes Firewall
func (cli *ZSClient) DeleteFirewall(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpcfirewalls", uuid, string(deleteMode))
}

// GetVmCapabilities gets VmCapabilities by uuid
func (cli *ZSClient) GetVmCapabilities(ctx context.Context, uuid string) (*view.GetVmCapabilitiesView, error) {
	var resp view.GetVmCapabilitiesView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAccessKeyState changes AccessKeyState
func (cli *ZSClient) ChangeAccessKeyState(ctx context.Context, uuid string, params param.ChangeAccessKeyStateParam) (*view.AccessKeyInventoryView, error) {
	resp := view.AccessKeyInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/accesskeys", uuid, "", map[string]interface{}{
		"changeAccessKeyState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployDistributedModelService operates on DeployDistributedModelService
func (cli *ZSClient) DeployDistributedModelService(ctx context.Context, params param.DeployDistributedModelServiceParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	resp := view.ModelServiceInstanceGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/ai/model-services", "", "", map[string]interface{}{
		"deployDistributedModelService": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2SystemAttributes gets IAM2SystemAttributes by uuid
func (cli *ZSClient) GetIAM2SystemAttributes(ctx context.Context) (*view.GetIAM2SystemAttributesView, error) {
	var resp view.GetIAM2SystemAttributesView
	if err := cli.GetWithRespKey(ctx, "v1/iam2/attributes/system", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeInstanceOfferingState changes InstanceOfferingState
func (cli *ZSClient) ChangeInstanceOfferingState(ctx context.Context, uuid string, params param.ChangeInstanceOfferingStateParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/instance-offerings", uuid, "", map[string]interface{}{
		"changeInstanceOfferingState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVirtualBorderRouterFromLocal queries VirtualBorderRouterFromLocal list
func (cli *ZSClient) QueryVirtualBorderRouterFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VirtualBorderRouterInventoryView, error) {
	var resp []view.VirtualBorderRouterInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/border-router", params, &resp)
}

func (cli *ZSClient) GetVirtualBorderRouterFromLocal(ctx context.Context, uuid string) (*view.VirtualBorderRouterInventoryView, error) {
	var resp view.VirtualBorderRouterInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/border-router", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVirtualBorderRouterFromLocal Pagination
func (cli *ZSClient) PageVirtualBorderRouterFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VirtualBorderRouterInventoryView, int, error) {
	var virtualBorderRouterFromLocals []view.VirtualBorderRouterInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/border-router", params, &virtualBorderRouterFromLocals)
	return virtualBorderRouterFromLocals, total, err
}

// GetBackupStorageCapacity gets BackupStorageCapacity by uuid
func (cli *ZSClient) GetBackupStorageCapacity(ctx context.Context) (*view.GetBackupStorageCapacityView, error) {
	var resp view.GetBackupStorageCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/backup-storage/capacities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateSeMdevDevices operates on SeMdevDevices
func (cli *ZSClient) GenerateSeMdevDevices(ctx context.Context, mttyDeviceUuid string, params param.GenerateSeMdevDevicesParam) (*view.GenerateSeMdevDevicesEventView, error) {
	resp := view.GenerateSeMdevDevicesEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/mtty-devices", mttyDeviceUuid, "", map[string]interface{}{
		"generateSeMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateMiniCluster creates MiniCluster
func (cli *ZSClient) CreateMiniCluster(ctx context.Context, params param.CreateMiniClusterParam) (*view.ClusterInventoryView, error) {
	resp := view.ClusterInventoryView{}
	if err := cli.Post(ctx, "v1/mini-clusters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncImageFromImageStoreBackupStorage operates on ImageFromImageStoreBackupStorage
func (cli *ZSClient) SyncImageFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.SyncImageFromImageStoreBackupStorageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/images", uuid, "", map[string]interface{}{
		"syncImageFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVipState changes VipState
func (cli *ZSClient) ChangeVipState(ctx context.Context, uuid string, params param.ChangeVipStateParam) (*view.VipInventoryView, error) {
	resp := view.VipInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vips", uuid, "", map[string]interface{}{
		"changeVipState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UndoSnapshotCreation operates on UndoSnapshotCreation
func (cli *ZSClient) UndoSnapshotCreation(ctx context.Context, uuid string, params param.UndoSnapshotCreationParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/volumes", uuid, "", map[string]interface{}{
		"undoSnapshotCreation": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddBuildApp adds BuildApp
func (cli *ZSClient) AddBuildApp(ctx context.Context, params param.AddBuildAppParam) (*view.BuildApplicationInventoryView, error) {
	resp := view.BuildApplicationInventoryView{}
	if err := cli.Post(ctx, "v1/appcenter/buildapp/add", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmFromVolumeBackup creates VmFromVolumeBackup
func (cli *ZSClient) CreateVmFromVolumeBackup(ctx context.Context, params param.CreateVmFromVolumeBackupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/from/vm-backup/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIdentityZoneFromRemote gets IdentityZoneFromRemote by uuid
func (cli *ZSClient) GetIdentityZoneFromRemote(ctx context.Context) (*view.GetIdentityZoneFromRemoteView, error) {
	var resp view.GetIdentityZoneFromRemoteView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/identity-zone/remote", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEcsInstanceVncUrl gets EcsInstanceVncUrl by uuid
func (cli *ZSClient) GetEcsInstanceVncUrl(ctx context.Context, uuid string) (*view.GetEcsInstanceVncUrlView, error) {
	var resp view.GetEcsInstanceVncUrlView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/aliyun/ecs-vnc", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddMonToCephPrimaryStorage adds MonToCephPrimaryStorage
func (cli *ZSClient) AddMonToCephPrimaryStorage(ctx context.Context, params param.AddMonToCephPrimaryStorageParam) (*view.CephPrimaryStorageInventoryView, error) {
	resp := view.CephPrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/ceph/{uuid}/mons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryLocalRaidPhysicalDrive queries LocalRaidPhysicalDrive list
func (cli *ZSClient) QueryLocalRaidPhysicalDrive(ctx context.Context, params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, error) {
	var resp []view.RaidPhysicalDriveInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/local-raid/physical-drives", params, &resp)
}

func (cli *ZSClient) GetLocalRaidPhysicalDrive(ctx context.Context, uuid string) (*view.RaidPhysicalDriveInventoryView, error) {
	var resp view.RaidPhysicalDriveInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/local-raid/physical-drives", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLocalRaidPhysicalDrive Pagination
func (cli *ZSClient) PageLocalRaidPhysicalDrive(ctx context.Context, params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, int, error) {
	var localRaidPhysicalDrives []view.RaidPhysicalDriveInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/local-raid/physical-drives", params, &localRaidPhysicalDrives)
	return localRaidPhysicalDrives, total, err
}

// RemoveHostRouteFromL3Network removes HostRouteFromL3Network
func (cli *ZSClient) RemoveHostRouteFromL3Network(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks", uuid, string(deleteMode))
}

// BackupStorageMigrateImage operates on StorageMigrateImage
func (cli *ZSClient) BackupStorageMigrateImage(ctx context.Context, imageUuid string, params param.BackupStorageMigrateImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage/images", imageUuid, "", map[string]interface{}{
		"backupStorageMigrateImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BackupStorageMigrateImageAsync Async
func (cli *ZSClient) BackupStorageMigrateImageAsync(ctx context.Context, params param.BackupStorageMigrateImageParam) (string, error) {

	resource := "v1/backup-storage/images/{imageUuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// QueryArchiveTicketHistory queries ArchiveTicketHistory list
func (cli *ZSClient) QueryArchiveTicketHistory(ctx context.Context, params *param.QueryParam) ([]view.ArchiveTicketStatusHistoryInventoryView, error) {
	var resp []view.ArchiveTicketStatusHistoryInventoryView
	return resp, cli.List(ctx, "v1/tickets/histories/archives", params, &resp)
}

func (cli *ZSClient) GetArchiveTicketHistory(ctx context.Context, uuid string) (*view.ArchiveTicketStatusHistoryInventoryView, error) {
	var resp view.ArchiveTicketStatusHistoryInventoryView
	if err := cli.Get(ctx, "v1/tickets/histories/archives", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageArchiveTicketHistory Pagination
func (cli *ZSClient) PageArchiveTicketHistory(ctx context.Context, params *param.QueryParam) ([]view.ArchiveTicketStatusHistoryInventoryView, int, error) {
	var archiveTicketHistories []view.ArchiveTicketStatusHistoryInventoryView
	total, err := cli.Page(ctx, "v1/tickets/histories/archives", params, &archiveTicketHistories)
	return archiveTicketHistories, total, err
}

// ChangeIAM2VirtualIDType changes IAM2VirtualIDType
func (cli *ZSClient) ChangeIAM2VirtualIDType(ctx context.Context, uuid string, params param.ChangeIAM2VirtualIDTypeParam) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/virtual-ids", uuid, "", map[string]interface{}{
		"changeIAM2VirtualIDType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveBackendServerFromServerGroup removes BackendServerFromServerGroup
func (cli *ZSClient) RemoveBackendServerFromServerGroup(ctx context.Context, serverGroupUuid string, params param.RemoveBackendServerFromServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/load-balancers/servergroups", serverGroupUuid, "", map[string]interface{}{
		"removeBackendServerFromServerGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedVip gets VpcAttachedVip by uuid
func (cli *ZSClient) GetVpcAttachedVip(ctx context.Context, params param.GetVpcAttachedVipParam) (*view.VipInventoryView, error) {
	resp := view.VipInventoryView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/attached-vip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIpv6Range adds Ipv6Range
func (cli *ZSClient) AddIpv6Range(ctx context.Context, params param.AddIpv6RangeParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.Post(ctx, "v1/l3-networks/{l3NetworkUuid}/ipv6-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckBaremetalChassisConfigFile operates on BaremetalChassisConfigFile
func (cli *ZSClient) CheckBaremetalChassisConfigFile(ctx context.Context, params param.CheckBaremetalChassisConfigFileParam) (*view.CheckBaremetalChassisConfigFileView, error) {
	resp := view.CheckBaremetalChassisConfigFileView{}
	if err := cli.Post(ctx, "v1/baremetal/chassis/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOssBucketFileRemote deletes OssBucketFileRemote
func (cli *ZSClient) DeleteOssBucketFileRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/oss-bucket-file/remote", uuid, string(deleteMode))
}

// ChangeMulticastRouterState changes MulticastRouterState
func (cli *ZSClient) ChangeMulticastRouterState(ctx context.Context, uuid string, params param.ChangeMulticastRouterStateParam) (*view.MulticastRouterInventoryView, error) {
	resp := view.MulticastRouterInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/multicast/virtual-routers", uuid, "", map[string]interface{}{
		"changeMulticastRouterState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMaaSUsage gets MaaSUsage by uuid
func (cli *ZSClient) GetMaaSUsage(ctx context.Context) (*view.GetMaaSUsageView, error) {
	var resp view.GetMaaSUsageView
	if err := cli.GetWithRespKey(ctx, "v1/maas/usage", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFreeIp gets FreeIp by uuid
func (cli *ZSClient) GetFreeIp(ctx context.Context) (*view.GetFreeIpView, error) {
	var resp view.GetFreeIpView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks/ip/free", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOssBucketRemote deletes OssBucketRemote
func (cli *ZSClient) DeleteOssBucketRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/oss-bucket/remote", uuid, string(deleteMode))
}

// CreateL2PortGroup creates L2PortGroup
func (cli *ZSClient) CreateL2PortGroup(ctx context.Context, params param.CreateL2PortGroupParam) (*view.CreateL2PortGroupEventView, error) {
	resp := view.CreateL2PortGroupEventView{}
	if err := cli.Post(ctx, "v1/l2-networks/port-group", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateInstanceOfferingUserConfig operates on InstanceOfferingUserConfig
func (cli *ZSClient) ValidateInstanceOfferingUserConfig(ctx context.Context, params param.ValidateInstanceOfferingUserConfigParam) (*view.ValidateInstanceOfferingUserConfigEventView, error) {
	resp := view.ValidateInstanceOfferingUserConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/accounts/actions", "", "", map[string]interface{}{
		"validateInstanceOfferingUserConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryTag queries Tag list
func (cli *ZSClient) QueryTag(ctx context.Context, params *param.QueryParam) ([]view.TagPatternInventoryView, error) {
	var resp []view.TagPatternInventoryView
	return resp, cli.List(ctx, "v1/tags", params, &resp)
}

func (cli *ZSClient) GetTag(ctx context.Context, uuid string) (*view.TagPatternInventoryView, error) {
	var resp view.TagPatternInventoryView
	if err := cli.Get(ctx, "v1/tags", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTag Pagination
func (cli *ZSClient) PageTag(ctx context.Context, params *param.QueryParam) ([]view.TagPatternInventoryView, int, error) {
	var tags []view.TagPatternInventoryView
	total, err := cli.Page(ctx, "v1/tags", params, &tags)
	return tags, total, err
}

// SetVmHostname operates on VmHostname
func (cli *ZSClient) SetVmHostname(ctx context.Context, uuid string, params param.SetVmHostnameParam) (*view.SetVmHostnameEventView, error) {
	resp := view.SetVmHostnameEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmHostname": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// TriggerGCJob operates on TriggerGCJob
func (cli *ZSClient) TriggerGCJob(ctx context.Context, uuid string, params param.TriggerGCJobParam) (*view.TriggerGCJobEventView, error) {
	resp := view.TriggerGCJobEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/gc-jobs", uuid, "", map[string]interface{}{
		"triggerGCJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckBareMetal2IpmiChassisConfigFile operates on BareMetal2IpmiChassisConfigFile
func (cli *ZSClient) CheckBareMetal2IpmiChassisConfigFile(ctx context.Context) (*view.CheckBareMetal2ChassisConfigFileView, error) {
	resp := view.CheckBareMetal2ChassisConfigFileView{}
	if err := cli.Post(ctx, "v1/baremetal2/chassis/ipmi/from-file/check", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVirtualRouterLocal deletes VirtualRouterLocal
func (cli *ZSClient) DeleteVirtualRouterLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/vrouter", uuid, string(deleteMode))
}

// DeleteVpcIkeConfigLocal deletes VpcIkeConfigLocal
func (cli *ZSClient) DeleteVpcIkeConfigLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/vpn-connection/ike", uuid, string(deleteMode))
}

// CreateOssBucketRemote creates OssBucketRemote
func (cli *ZSClient) CreateOssBucketRemote(ctx context.Context, params param.CreateOssBucketRemoteParam) (*view.OssBucketInventoryView, error) {
	resp := view.OssBucketInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/oss-bucket/remote", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSimulatorPrimaryStorage adds SimulatorPrimaryStorage
func (cli *ZSClient) AddSimulatorPrimaryStorage(ctx context.Context, params param.AddSimulatorPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachVRouterRouteTableFromVRouter operates on VRouterRouteTableFromVRouter
func (cli *ZSClient) DetachVRouterRouteTableFromVRouter(ctx context.Context, routeTableUuid string, virtualRouterVmUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/vrouter-route-tables", routeTableUuid, fmt.Sprintf("detach/%s", virtualRouterVmUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetVipUsedPorts gets VipUsedPorts by uuid
func (cli *ZSClient) GetVipUsedPorts(ctx context.Context, uuid string) (*view.GetVipUsedPortsView, error) {
	var resp view.GetVipUsedPortsView
	if err := cli.GetWithRespKey(ctx, "v1/vips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmConsolePassword operates on VmConsolePassword
func (cli *ZSClient) SetVmConsolePassword(ctx context.Context, uuid string, params param.SetVmConsolePasswordParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmConsolePassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachFirewallRuleSetToL3 operates on FirewallRuleSetToL3
func (cli *ZSClient) AttachFirewallRuleSetToL3(ctx context.Context, params param.AttachFirewallRuleSetToL3Param) (*view.VpcFirewallRuleSetL3RefInventoryView, error) {
	resp := view.VpcFirewallRuleSetL3RefInventoryView{}
	if err := cli.Post(ctx, "v1/vpcfirewalls/ruleSets/{ruleSetUuid}/l3networks/{l3Uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpStorageTrashOnPrimaryStorage operates on UpStorageTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpStorageTrashOnPrimaryStorage(ctx context.Context, uuid string, params param.CleanUpStorageTrashOnPrimaryStorageParam) (*view.CleanUpStorageTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpStorageTrashOnPrimaryStorageEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage", uuid, "", map[string]interface{}{
		"cleanUpStorageTrashOnPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVRouterFlowMeterNetwork queries VRouterFlowMeterNetwork list
func (cli *ZSClient) QueryVRouterFlowMeterNetwork(ctx context.Context, params *param.QueryParam) ([]view.NetworkRouterFlowMeterRefInventoryView, error) {
	var resp []view.NetworkRouterFlowMeterRefInventoryView
	return resp, cli.List(ctx, "v1/flowmeters/networks", params, &resp)
}

func (cli *ZSClient) GetVRouterFlowMeterNetwork(ctx context.Context, uuid string) (*view.NetworkRouterFlowMeterRefInventoryView, error) {
	var resp view.NetworkRouterFlowMeterRefInventoryView
	if err := cli.Get(ctx, "v1/flowmeters/networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterFlowMeterNetwork Pagination
func (cli *ZSClient) PageVRouterFlowMeterNetwork(ctx context.Context, params *param.QueryParam) ([]view.NetworkRouterFlowMeterRefInventoryView, int, error) {
	var vRouterFlowMeterNetworks []view.NetworkRouterFlowMeterRefInventoryView
	total, err := cli.Page(ctx, "v1/flowmeters/networks", params, &vRouterFlowMeterNetworks)
	return vRouterFlowMeterNetworks, total, err
}

// GetManagementNodeDirCapacity gets ManagementNodeDirCapacity by uuid
func (cli *ZSClient) GetManagementNodeDirCapacity(ctx context.Context) (*view.GetManagementNodeDirCapacityView, error) {
	var resp view.GetManagementNodeDirCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/mn", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGpuDeviceSpecCandidates gets GpuDeviceSpecCandidates by uuid
func (cli *ZSClient) GetGpuDeviceSpecCandidates(ctx context.Context) (*view.GetGpuDeviceSpecCandidatesView, error) {
	var resp view.GetGpuDeviceSpecCandidatesView
	if err := cli.GetWithRespKey(ctx, "v1/gpu-device-specs/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallRuleSetL3Ref queries FirewallRuleSetL3Ref list
func (cli *ZSClient) QueryFirewallRuleSetL3Ref(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleSetL3RefInventoryView, error) {
	var resp []view.VpcFirewallRuleSetL3RefInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/l3networks/rulesets/refs", params, &resp)
}

func (cli *ZSClient) GetFirewallRuleSetL3Ref(ctx context.Context, uuid string) (*view.VpcFirewallRuleSetL3RefInventoryView, error) {
	var resp view.VpcFirewallRuleSetL3RefInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls/l3networks/rulesets/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRuleSetL3Ref Pagination
func (cli *ZSClient) PageFirewallRuleSetL3Ref(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleSetL3RefInventoryView, int, error) {
	var firewallRuleSetL3Refs []view.VpcFirewallRuleSetL3RefInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/l3networks/rulesets/refs", params, &firewallRuleSetL3Refs)
	return firewallRuleSetL3Refs, total, err
}

// UngroupVolumeSnapshotGroup operates on UngroupVolumeSnapshotGroup
func (cli *ZSClient) UngroupVolumeSnapshotGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volume-snapshots/ungroup", uuid, string(deleteMode))
}

// SubscribeSNSTopic operates on SubscribeSNSTopic
func (cli *ZSClient) SubscribeSNSTopic(ctx context.Context, params param.SubscribeSNSTopicParam) (*view.SubscribeSNSTopicEventView, error) {
	resp := view.SubscribeSNSTopicEventView{}
	if err := cli.Post(ctx, "v1/sns/topics/{topicUuid}/endpoints/{endpointUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicForSecurityGroup gets CandidateVmNicForSecurityGroup by uuid
func (cli *ZSClient) GetCandidateVmNicForSecurityGroup(ctx context.Context, uuid string) (*view.GetCandidateVmNicForSecurityGroupView, error) {
	var resp view.GetCandidateVmNicForSecurityGroupView
	if err := cli.GetWithRespKey(ctx, "v1/security-groups", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmRDP gets VmRDP by uuid
func (cli *ZSClient) GetVmRDP(ctx context.Context, uuid string) (*view.GetVmRDPView, error) {
	var resp view.GetVmRDPView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPciDeviceToVm operates on PciDeviceToVm
func (cli *ZSClient) AttachPciDeviceToVm(ctx context.Context, params param.AttachPciDeviceToVmParam) (*view.PciDeviceInventoryView, error) {
	resp := view.PciDeviceInventoryView{}
	if err := cli.Post(ctx, "v1/pci-device/pci-devices/{pciDeviceUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanupBillingUsage operates on upBillingUsage
func (cli *ZSClient) CleanupBillingUsage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/billings/usage", uuid, string(deleteMode))
}

// GetLdapEntry gets LdapEntry by uuid
func (cli *ZSClient) GetLdapEntry(ctx context.Context) (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.GetWithRespKey(ctx, "v1/ldap/entry", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL2NetworksForAttachingCluster gets CandidateL2NetworksForAttachingCluster by uuid
func (cli *ZSClient) GetCandidateL2NetworksForAttachingCluster(ctx context.Context, uuid string) (*view.GetCandidateL2NetworksForAttachingClusterView, error) {
	var resp view.GetCandidateL2NetworksForAttachingClusterView
	if err := cli.GetWithRespKey(ctx, "v1/cluster", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsVfNicAvailableInL3Network operates on IsVfNicAvailableInL3Network
func (cli *ZSClient) IsVfNicAvailableInL3Network(ctx context.Context, l3NetworkUuid string, hostUuid string) (*view.IsVfNicAvailableInL3NetworkView, error) {
	var resp view.IsVfNicAvailableInL3NetworkView
	err := cli.GetWithSpec(ctx, "v1/l3-networks", l3NetworkUuid, fmt.Sprintf("hosts/%s/vfnicavailable", hostUuid), "vfNicAvailable", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAllMetricMetadata gets AllMetricMetadata by uuid
func (cli *ZSClient) GetAllMetricMetadata(ctx context.Context) (*view.GetAllMetricMetadataView, error) {
	var resp view.GetAllMetricMetadataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/metrics/meta-data", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddOssBucketFromRemote adds OssBucketFromRemote
func (cli *ZSClient) AddOssBucketFromRemote(ctx context.Context, params param.AddOssBucketFromRemoteParam) (*view.OssBucketInventoryView, error) {
	resp := view.OssBucketInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/oss-bucket", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVmBackup operates on VmBackup
func (cli *ZSClient) SyncVmBackup(ctx context.Context, imageStoreUuid string, params param.SyncVmBackupParam) (*view.SyncVmBackupEventView, error) {
	resp := view.SyncVmBackupEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-backups/imageStore", imageStoreUuid, "", map[string]interface{}{
		"syncVmBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshGuestOsMetadata operates on GuestOsMetadata
func (cli *ZSClient) RefreshGuestOsMetadata(ctx context.Context) (*view.RefreshGuestOsMetadataEventView, error) {
	resp := view.RefreshGuestOsMetadataEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/guest-os/metadata/actions", "", "", map[string]interface{}{
		"refreshGuestOsMetadata": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GCAliyunSnapshotRemote operates on GCAliyunSnapshotRemote
func (cli *ZSClient) GCAliyunSnapshotRemote(ctx context.Context, params param.GCAliyunSnapshotRemoteParam) (*view.GCAliyunSnapshotRemoteEventView, error) {
	resp := view.GCAliyunSnapshotRemoteEventView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/snapshot/{dataCenterUuid}/gc", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadBackupFileFromPublicCloud operates on DownloadBackupFileFromPublicCloud
func (cli *ZSClient) DownloadBackupFileFromPublicCloud(ctx context.Context, params param.DownloadBackupFileFromPublicCloudParam) (*view.DownloadBackupFileFromPublicCloudEventView, error) {
	resp := view.DownloadBackupFileFromPublicCloudEventView{}
	if err := cli.Post(ctx, "v1/hybrid/backup-mysql/download", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIAM2VirtualIDsToProjects adds IAM2VirtualIDsToProjects
func (cli *ZSClient) AddIAM2VirtualIDsToProjects(ctx context.Context, params param.AddIAM2VirtualIDsToProjectsParam) (*view.AddIAM2VirtualIDsToProjectsEventView, error) {
	resp := view.AddIAM2VirtualIDsToProjectsEventView{}
	if err := cli.Post(ctx, "v1/iam2/projects/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2ProjectTemplateFromProject creates IAM2ProjectTemplateFromProject
func (cli *ZSClient) CreateIAM2ProjectTemplateFromProject(ctx context.Context, params param.CreateIAM2ProjectTemplateFromProjectParam) (*view.IAM2ProjectTemplateInventoryView, error) {
	resp := view.IAM2ProjectTemplateInventoryView{}
	if err := cli.Post(ctx, "v1/iam2/projects/templates/from/projects/{projectUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTag creates Tag
func (cli *ZSClient) CreateTag(ctx context.Context, params param.CreateTagParam) (*view.TagPatternInventoryView, error) {
	resp := view.TagPatternInventoryView{}
	if err := cli.Post(ctx, "v1/tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromVolumeSnapshotGroup creates VmInstanceFromVolumeSnapshotGroup
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshotGroup(ctx context.Context, params param.CreateVmInstanceFromVolumeSnapshotGroupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/from/volume-snapshots/group/{volumeSnapshotGroupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIAM2ProjectRetirePolicy operates on IAM2ProjectRetirePolicy
func (cli *ZSClient) SetIAM2ProjectRetirePolicy(ctx context.Context, uuid string, params param.SetIAM2ProjectRetirePolicyParam) (*view.SetIAM2ProjectRetirePolicyEventView, error) {
	resp := view.SetIAM2ProjectRetirePolicyEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/projects/retire-policies", uuid, "", map[string]interface{}{
		"setIAM2ProjectRetirePolicy": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunIAM2Script operates on RunIAM2Script
func (cli *ZSClient) RunIAM2Script(ctx context.Context, params param.RunIAM2ScriptParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Post(ctx, "v1/iam2/iam2-script/run", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunIAM2ScriptAsync Async
func (cli *ZSClient) RunIAM2ScriptAsync(ctx context.Context, params param.RunIAM2ScriptParam) (string, error) {

	resource := "v1/iam2/iam2-script/run"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// AttachServiceToObservabilityServer operates on ServiceToObservabilityServer
func (cli *ZSClient) AttachServiceToObservabilityServer(ctx context.Context, params param.AttachServiceToObservabilityServerParam) (*view.ObservabilityServerVmInventoryView, error) {
	resp := view.ObservabilityServerVmInventoryView{}
	if err := cli.Post(ctx, "v1/observability-server/{observabilityServerUuid}/service", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteHostNetworkServiceType deletes HostNetworkServiceType
func (cli *ZSClient) DeleteHostNetworkServiceType(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hosts/service-types", uuid, string(deleteMode))
}

// SubscribeResNotify operates on SubscribeResNotify
func (cli *ZSClient) SubscribeResNotify(ctx context.Context, params param.SubscribeResNotifyParam) (*view.ResNotifySubscriptionInventoryView, error) {
	resp := view.ResNotifySubscriptionInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/resnotify/subscriptions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2ProjectFromTemplate creates IAM2ProjectFromTemplate
func (cli *ZSClient) CreateIAM2ProjectFromTemplate(ctx context.Context, params param.CreateIAM2ProjectFromTemplateParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.Post(ctx, "v1/iam2/projects/from/templates/{templateUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcIpSecConfigFromLocal queries VpcIpSecConfigFromLocal list
func (cli *ZSClient) QueryVpcIpSecConfigFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVpnIpSecConfigInventoryView, error) {
	var resp []view.VpcVpnIpSecConfigInventoryView
	return resp, cli.List(ctx, "v1/hybrid/vpn-connection/ipsec", params, &resp)
}

func (cli *ZSClient) GetVpcIpSecConfigFromLocal(ctx context.Context, uuid string) (*view.VpcVpnIpSecConfigInventoryView, error) {
	var resp view.VpcVpnIpSecConfigInventoryView
	if err := cli.Get(ctx, "v1/hybrid/vpn-connection/ipsec", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcIpSecConfigFromLocal Pagination
func (cli *ZSClient) PageVpcIpSecConfigFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVpnIpSecConfigInventoryView, int, error) {
	var vpcIpSecConfigFromLocals []view.VpcVpnIpSecConfigInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/vpn-connection/ipsec", params, &vpcIpSecConfigFromLocals)
	return vpcIpSecConfigFromLocals, total, err
}

// AddConnectionAccessPointFromRemote adds ConnectionAccessPointFromRemote
func (cli *ZSClient) AddConnectionAccessPointFromRemote(ctx context.Context, params param.AddConnectionAccessPointFromRemoteParam) (*view.ConnectionAccessPointInventoryView, error) {
	resp := view.ConnectionAccessPointInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/access-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachSshKeyPairToVmInstance operates on SshKeyPairToVmInstance
func (cli *ZSClient) AttachSshKeyPairToVmInstance(ctx context.Context, params param.AttachSshKeyPairToVmInstanceParam) (*view.SshKeyPairInventoryView, error) {
	resp := view.SshKeyPairInventoryView{}
	if err := cli.Post(ctx, "v1/ssh-key-pair/{sshKeyPairUuid}/vm-instance/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEmailTriggerAction queries EmailTrigger list
func (cli *ZSClient) QueryEmailTriggerAction(ctx context.Context, params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, error) {
	var resp []view.MonitorTriggerActionInventoryView
	return resp, cli.List(ctx, "v1/monitoring/trigger-actions/emails", params, &resp)
}

func (cli *ZSClient) GetEmailTriggerAction(ctx context.Context, uuid string) (*view.MonitorTriggerActionInventoryView, error) {
	var resp view.MonitorTriggerActionInventoryView
	if err := cli.Get(ctx, "v1/monitoring/trigger-actions/emails", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEmailTriggerAction Pagination
func (cli *ZSClient) PageEmailTriggerAction(ctx context.Context, params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, int, error) {
	var emailTriggers []view.MonitorTriggerActionInventoryView
	total, err := cli.Page(ctx, "v1/monitoring/trigger-actions/emails", params, &emailTriggers)
	return emailTriggers, total, err
}

// DetachBareMetal2GatewayFromCluster operates on BareMetal2GatewayFromCluster
func (cli *ZSClient) DetachBareMetal2GatewayFromCluster(ctx context.Context, clusterUuid string, gatewayUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/baremetal2/clusters", clusterUuid, fmt.Sprintf("gateways/%s", gatewayUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ReloadElaboration operates on ReloadElaboration
func (cli *ZSClient) ReloadElaboration(ctx context.Context) (*view.ReloadElaborationEventView, error) {
	resp := view.ReloadElaborationEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/errorcode/actions", "", "", map[string]interface{}{
		"reloadElaboration": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReconnectVirtualRouter operates on VirtualRouter
func (cli *ZSClient) ReconnectVirtualRouter(ctx context.Context, vmInstanceUuid string, params param.ReconnectVirtualRouterParam) (*view.ApplianceVmInventoryView, error) {
	resp := view.ApplianceVmInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/appliances/virtual-routers", vmInstanceUuid, "", map[string]interface{}{
		"reconnectVirtualRouter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConvertVmFromForeignHypervisor operates on ConvertVmFromForeignHypervisor
func (cli *ZSClient) ConvertVmFromForeignHypervisor(ctx context.Context, params param.ConvertVmFromForeignHypervisorParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Post(ctx, "v1/v2vs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConvertVmFromForeignHypervisorAsync Async
func (cli *ZSClient) ConvertVmFromForeignHypervisorAsync(ctx context.Context, params param.ConvertVmFromForeignHypervisorParam) (string, error) {

	resource := "v1/v2vs"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch deletes ConnectionBetweenL3NetWorkAndAliyunVSwitch
func (cli *ZSClient) DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/connections", uuid, string(deleteMode))
}

// RestartResourceStack operates on RestartResourceStack
func (cli *ZSClient) RestartResourceStack(ctx context.Context, uuid string, params param.RestartResourceStackParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/cloudformation/stack", uuid, "", map[string]interface{}{
		"restartResourceStack": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsImageFromRemote operates on EcsImageFromRemote
func (cli *ZSClient) SyncEcsImageFromRemote(ctx context.Context, params param.SyncEcsImageFromRemoteParam) (*view.EcsImageInventoryView, error) {
	resp := view.EcsImageInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/image/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPoliciesToUser operates on PoliciesToUser
func (cli *ZSClient) AttachPoliciesToUser(ctx context.Context, params param.AttachPoliciesToUserParam) (*view.AttachPoliciesToUserEventView, error) {
	resp := view.AttachPoliciesToUserEventView{}
	if err := cli.Post(ctx, "v1/accounts/users/{userUuid}/policy-collection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachBackupStorageToZone operates on BackupStorageToZone
func (cli *ZSClient) AttachBackupStorageToZone(ctx context.Context, params param.AttachBackupStorageToZoneParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Post(ctx, "v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddPciDeviceSpecToVmInstance adds PciDeviceSpecToVmInstance
func (cli *ZSClient) AddPciDeviceSpecToVmInstance(ctx context.Context, params param.AddPciDeviceSpecToVmInstanceParam) (*view.VmInstancePciDeviceSpecRefInventoryView, error) {
	resp := view.VmInstancePciDeviceSpecRefInventoryView{}
	if err := cli.Post(ctx, "v1/pci-device-specs/{pciSpecUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResizeRootVolume operates on RootVolume
func (cli *ZSClient) ResizeRootVolume(ctx context.Context, uuid string, params param.ResizeRootVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/volumes/resize", uuid, "", map[string]interface{}{
		"resizeRootVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcVpnConfigurationFromRemote gets VpcVpnConfigurationFromRemote by uuid
func (cli *ZSClient) GetVpcVpnConfigurationFromRemote(ctx context.Context, uuid string) (*view.GetVpcVpnConfigurationFromRemoteView, error) {
	var resp view.GetVpcVpnConfigurationFromRemoteView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/vpn-conf", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateImageGroupFromImage creates ImageGroupFromImage
func (cli *ZSClient) CreateImageGroupFromImage(ctx context.Context, params param.CreateImageGroupFromImageParam) (*view.ImageGroupInventoryView, error) {
	resp := view.ImageGroupInventoryView{}
	if err := cli.Post(ctx, "v1/imagegroup/from/image/{rootVolumeTemplateUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// TokenIntrospection operates on TokenIntrospection
func (cli *ZSClient) TokenIntrospection(ctx context.Context, params param.TokenIntrospectionParam) (*view.TokenIntrospectionView, error) {
	resp := view.TokenIntrospectionView{}
	if err := cli.Post(ctx, "v1/token/introspect", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVmBackupFromImageStoreBackupStorage operates on VmBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncVmBackupFromImageStoreBackupStorage(ctx context.Context, groupUuid string, params param.SyncVmBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-backups", groupUuid, "", map[string]interface{}{
		"syncVmBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddCertificateToLoadBalancerListener adds CertificateToLoadBalancerListener
func (cli *ZSClient) AddCertificateToLoadBalancerListener(ctx context.Context, params param.AddCertificateToLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.Post(ctx, "v1/load-balancers/listeners/{listenerUuid}/certificate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRolesToIAM2VirtualID adds RolesToIAM2VirtualID
func (cli *ZSClient) AddRolesToIAM2VirtualID(ctx context.Context, params param.AddRolesToIAM2VirtualIDParam) (*view.AddRolesToIAM2VirtualIDEventView, error) {
	resp := view.AddRolesToIAM2VirtualIDEventView{}
	if err := cli.Post(ctx, "v1/iam2/projects/virtual-ids/{virtualIDUuid}/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryTicketHistory queries TicketHistory list
func (cli *ZSClient) QueryTicketHistory(ctx context.Context, params *param.QueryParam) ([]view.TicketStatusHistoryInventoryView, error) {
	var resp []view.TicketStatusHistoryInventoryView
	return resp, cli.List(ctx, "v1/tickets/histories", params, &resp)
}

func (cli *ZSClient) GetTicketHistory(ctx context.Context, uuid string) (*view.TicketStatusHistoryInventoryView, error) {
	var resp view.TicketStatusHistoryInventoryView
	if err := cli.Get(ctx, "v1/tickets/histories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTicketHistory Pagination
func (cli *ZSClient) PageTicketHistory(ctx context.Context, params *param.QueryParam) ([]view.TicketStatusHistoryInventoryView, int, error) {
	var ticketHistories []view.TicketStatusHistoryInventoryView
	total, err := cli.Page(ctx, "v1/tickets/histories", params, &ticketHistories)
	return ticketHistories, total, err
}

// CreateFaultToleranceVmInstance creates FaultToleranceVmInstance
func (cli *ZSClient) CreateFaultToleranceVmInstance(ctx context.Context, params param.CreateFaultToleranceVmInstanceParam) (*view.CreateFaultToleranceVmInstanceEventView, error) {
	resp := view.CreateFaultToleranceVmInstanceEventView{}
	if err := cli.Post(ctx, "v1/vm-instances/fault-tolerance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteResourceStackVmPortMonitor deletes ResourceStackVmPortMonitor
func (cli *ZSClient) DeleteResourceStackVmPortMonitor(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/cloudformation/stack/monitor/delvm", uuid, string(deleteMode))
}

// DeleteGCJob deletes GCJob
func (cli *ZSClient) DeleteGCJob(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/gc-jobs", uuid, string(deleteMode))
}

// DeleteEmailAddressOfSNSEmailEndpoint deletes EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) DeleteEmailAddressOfSNSEmailEndpoint(ctx context.Context, endpointUuid string, emailAddressUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/sns/application-endpoints/emails", endpointUuid, fmt.Sprintf("email-addresses/%s", emailAddressUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// CleanInvalidLdapIAM2Binding operates on InvalidLdapIAM2Binding
func (cli *ZSClient) CleanInvalidLdapIAM2Binding(ctx context.Context) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/ldap/bindings/actions", "", "", map[string]interface{}{
		"cleanInvalidLdapIAM2Binding": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHybridEip updates HybridEip
func (cli *ZSClient) UpdateHybridEip(ctx context.Context, uuid string, params param.UpdateHybridEipParam) (*view.HybridEipAddressInventoryView, error) {
	resp := view.HybridEipAddressInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/eip", uuid, "", map[string]interface{}{
		"updateHybridEip": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedIpsec gets VpcAttachedIpsec by uuid
func (cli *ZSClient) GetVpcAttachedIpsec(ctx context.Context, params param.GetVpcAttachedIpsecParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/attached-ipsec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImagesFromImageStoreBackupStorage gets ImagesFromImageStoreBackupStorage by uuid
func (cli *ZSClient) GetImagesFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.GetImagesFromImageStoreBackupStorageParam) (*view.GetImagesFromImageStoreBackupStorageView, error) {
	resp := view.GetImagesFromImageStoreBackupStorageView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage", uuid, "", map[string]interface{}{
		"getImagesFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetElaborationCategories gets ElaborationCategories by uuid
func (cli *ZSClient) GetElaborationCategories(ctx context.Context) (*view.GetElaborationCategoriesView, error) {
	var resp view.GetElaborationCategoriesView
	if err := cli.GetWithRespKey(ctx, "v1/errorcode/elaborations/categories", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetScsiLunCandidatesForAttachingVm gets ScsiLunCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetScsiLunCandidatesForAttachingVm(ctx context.Context, uuid string) (*view.GetScsiLunCandidatesForAttachingVmView, error) {
	var resp view.GetScsiLunCandidatesForAttachingVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostMultipathTopology gets HostMultipathTopology by uuid
func (cli *ZSClient) GetHostMultipathTopology(ctx context.Context) (*view.GetHostMultipathTopologyView, error) {
	var resp view.GetHostMultipathTopologyView
	if err := cli.GetWithRespKey(ctx, "v1/storage-devices/multipath/topology", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallIpSetTemplate queries FirewallIpSetTemplate list
func (cli *ZSClient) QueryFirewallIpSetTemplate(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallIpSetTemplateInventoryView, error) {
	var resp []view.VpcFirewallIpSetTemplateInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/ipset/templates", params, &resp)
}

func (cli *ZSClient) GetFirewallIpSetTemplate(ctx context.Context, uuid string) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	var resp view.VpcFirewallIpSetTemplateInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls/ipset/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallIpSetTemplate Pagination
func (cli *ZSClient) PageFirewallIpSetTemplate(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallIpSetTemplateInventoryView, int, error) {
	var firewallIpSetTemplates []view.VpcFirewallIpSetTemplateInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/ipset/templates", params, &firewallIpSetTemplates)
	return firewallIpSetTemplates, total, err
}

// DeleteEcsImageRemote deletes EcsImageRemote
func (cli *ZSClient) DeleteEcsImageRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/image/remote", uuid, string(deleteMode))
}

// GetHostNetworkFacts gets HostNetworkFacts by uuid
func (cli *ZSClient) GetHostNetworkFacts(ctx context.Context, uuid string) (*view.GetHostNetworkFactsView, error) {
	var resp view.GetHostNetworkFactsView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/network-facts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpTrashOnBackupStorage operates on UpTrashOnBackupStorage
func (cli *ZSClient) CleanUpTrashOnBackupStorage(ctx context.Context, uuid string, params param.CleanUpTrashOnBackupStorageParam) (*view.CleanUpTrashOnBackupStorageEventView, error) {
	resp := view.CleanUpTrashOnBackupStorageEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage", uuid, "", map[string]interface{}{
		"cleanUpTrashOnBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateConnectionBetweenL3NetworkAndAliyunVSwitch creates ConnectionBetweenL3NetworkAndAliyunVSwitch
func (cli *ZSClient) CreateConnectionBetweenL3NetworkAndAliyunVSwitch(ctx context.Context, params param.CreateConnectionBetweenL3NetworkAndAliyunVSwitchParam) (*view.ConnectionRelationShipInventoryView, error) {
	resp := view.ConnectionRelationShipInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/connections", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPriceTableFromAccount operates on PriceTableFromAccount
func (cli *ZSClient) DetachPriceTableFromAccount(ctx context.Context, tableUuid string, accountUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/billings/price-tables", tableUuid, fmt.Sprintf("accounts/%s", accountUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddVRouterNetworksToFlowMeter adds VRouterNetworksToFlowMeter
func (cli *ZSClient) AddVRouterNetworksToFlowMeter(ctx context.Context, params param.AddVRouterNetworksToFlowMeterParam) (*view.NetworkRouterFlowMeterRefInventoryView, error) {
	resp := view.NetworkRouterFlowMeterRefInventoryView{}
	if err := cli.Post(ctx, "v1/flowmeters/{flowMeterUuid}/router/{vRouterUuid}/addnetworks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2VirtualIDInGroup gets IAM2VirtualIDInGroup by uuid
func (cli *ZSClient) GetIAM2VirtualIDInGroup(ctx context.Context) (*view.GetIAM2VirtualIDInGroupView, error) {
	var resp view.GetIAM2VirtualIDInGroupView
	if err := cli.GetWithRespKey(ctx, "v1/iam2/IAM2VirtualIDGroup/IAM2VirtualID", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnlockIdentity operates on UnlockIdentity
func (cli *ZSClient) UnlockIdentity(ctx context.Context) (*view.UnlockIdentityView, error) {
	var resp view.UnlockIdentityView
	if err := cli.GetWithRespKey(ctx, "v1/login/control/unlock", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicsForPortMirror gets CandidateVmNicsForPortMirror by uuid
func (cli *ZSClient) GetCandidateVmNicsForPortMirror(ctx context.Context, portMirrorUuid string, typeParam string) (*view.GetCandidateVmNicsForPortMirrorView, error) {
	var resp view.GetCandidateVmNicsForPortMirrorView
	err := cli.GetWithSpec(ctx, "v1/port-mirrors", portMirrorUuid, fmt.Sprintf("vm-instances/candidate-nics/%s", typeParam), "inventories", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmSchedulingRuleState changes VmSchedulingRuleState
func (cli *ZSClient) ChangeVmSchedulingRuleState(ctx context.Context, uuid string, params param.ChangeVmSchedulingRuleStateParam) (*view.VmSchedulingRuleInventoryView, error) {
	resp := view.VmSchedulingRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vmSchedulingRule", uuid, "", map[string]interface{}{
		"changeVmSchedulingRuleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRule creates FirewallRule
func (cli *ZSClient) CreateFirewallRule(ctx context.Context, params param.CreateFirewallRuleParam) (*view.VpcFirewallRuleInventoryView, error) {
	resp := view.VpcFirewallRuleInventoryView{}
	if err := cli.Post(ctx, "v1/vpcfirewalls/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmEmulatorPinning gets VmEmulatorPinning by uuid
func (cli *ZSClient) GetVmEmulatorPinning(ctx context.Context, uuid string) (*view.GetVmEmulatorPinningView, error) {
	var resp view.GetVmEmulatorPinningView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDataVolumeAttachableVm gets DataVolumeAttachableVm by uuid
func (cli *ZSClient) GetDataVolumeAttachableVm(ctx context.Context, uuid string) (*view.GetDataVolumeAttachableVmView, error) {
	var resp view.GetDataVolumeAttachableVmView
	if err := cli.GetWithRespKey(ctx, "v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIpRangeByNetworkCidr adds IpRangeByNetworkCidr
func (cli *ZSClient) AddIpRangeByNetworkCidr(ctx context.Context, params param.AddIpRangeByNetworkCidrParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.Post(ctx, "v1/l3-networks/{l3NetworkUuid}/ip-ranges/by-cidr", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateL2NoVlanNetwork creates L2NoVlanNetwork
func (cli *ZSClient) CreateL2NoVlanNetwork(ctx context.Context) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.Post(ctx, "v1/l2-networks/no-vlan", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddMonToCephBackupStorage adds MonToCephBackupStorage
func (cli *ZSClient) AddMonToCephBackupStorage(ctx context.Context, params param.AddMonToCephBackupStorageParam) (*view.CephBackupStorageInventoryView, error) {
	resp := view.CephBackupStorageInventoryView{}
	if err := cli.Post(ctx, "v1/backup-storage/ceph/{uuid}/mons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachBareMetal2ProvisionNetworkFromCluster operates on BareMetal2ProvisionNetworkFromCluster
func (cli *ZSClient) DetachBareMetal2ProvisionNetworkFromCluster(ctx context.Context, clusterUuid string, networkUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/baremetal2/clusters", clusterUuid, fmt.Sprintf("provision-networks/%s", networkUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DeleteAliyunDiskFromLocal deletes AliyunDiskFromLocal
func (cli *ZSClient) DeleteAliyunDiskFromLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/disk", uuid, string(deleteMode))
}

// GetResourceNames gets ResourceNames by uuid
func (cli *ZSClient) GetResourceNames(ctx context.Context) (*view.GetResourceNamesView, error) {
	var resp view.GetResourceNamesView
	if err := cli.GetWithRespKey(ctx, "v1/resources/names", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2VirtualIDAPIPermission gets IAM2VirtualIDAPIPermission by uuid
func (cli *ZSClient) GetIAM2VirtualIDAPIPermission(ctx context.Context) (*view.GetIAM2VirtualIDAPIPermissionView, error) {
	var resp view.GetIAM2VirtualIDAPIPermissionView
	if err := cli.GetWithRespKey(ctx, "v1/iam2/virtual-ids/api-permissions", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOrganizationQuotaUsage gets OrganizationQuotaUsage by uuid
func (cli *ZSClient) GetOrganizationQuotaUsage(ctx context.Context, uuid string) (*view.GetOrganizationQuotaUsageView, error) {
	var resp view.GetOrganizationQuotaUsageView
	if err := cli.GetWithRespKey(ctx, "v1/iam2/organizations/quota", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceConfigs gets ResourceConfigs by uuid
func (cli *ZSClient) GetResourceConfigs(ctx context.Context, resourceUuid string, category string) (*view.GetResourceConfigsView, error) {
	var resp view.GetResourceConfigsView
	err := cli.GetWithSpec(ctx, "v1/resource-configurations", resourceUuid, fmt.Sprintf("%s", category), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVpcUserVpnGatewayFromRemote operates on VpcUserVpnGatewayFromRemote
func (cli *ZSClient) SyncVpcUserVpnGatewayFromRemote(ctx context.Context, dataCenterUuid string, params param.SyncVpcUserVpnGatewayFromRemoteParam) (*view.VpcUserVpnGatewayInventoryView, error) {
	resp := view.VpcUserVpnGatewayInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/user-vpn", dataCenterUuid, "", map[string]interface{}{
		"syncVpcUserVpnGatewayFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPrimaryStorageFromCluster operates on PrimaryStorageFromCluster
func (cli *ZSClient) DetachPrimaryStorageFromCluster(ctx context.Context, clusterUuid string, primaryStorageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/clusters", clusterUuid, fmt.Sprintf("primary-storage/%s", primaryStorageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// CheckStackTemplateParameters operates on StackTemplateParameters
func (cli *ZSClient) CheckStackTemplateParameters(ctx context.Context, params param.CheckStackTemplateParametersParam) (*view.CheckStackTemplateParametersView, error) {
	resp := view.CheckStackTemplateParametersView{}
	if err := cli.Post(ctx, "v1/cloudformation/stack/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFactoryModeState gets FactoryModeState by uuid
func (cli *ZSClient) GetFactoryModeState(ctx context.Context) (*view.GetFactoryModeStateView, error) {
	var resp view.GetFactoryModeStateView
	if err := cli.GetWithRespKey(ctx, "v1/management-nodes/factory-mode-state", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddServerGroupToLoadBalancerListener adds ServerGroupToLoadBalancerListener
func (cli *ZSClient) AddServerGroupToLoadBalancerListener(ctx context.Context, params param.AddServerGroupToLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.Post(ctx, "v1/load-balancers/listeners/{listenerUuid}/servergroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetActiveAlarmStatus gets ActiveAlarmStatus by uuid
func (cli *ZSClient) GetActiveAlarmStatus(ctx context.Context) (*view.GetActiveAlarmStatusView, error) {
	var resp view.GetActiveAlarmStatusView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/activealarms/status", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployModelEvalService operates on DeployModelEvalService
func (cli *ZSClient) DeployModelEvalService(ctx context.Context, uuid string, params param.DeployModelEvalServiceParam) (*view.ModelEvalServiceInstanceGroupInventoryView, error) {
	resp := view.ModelEvalServiceInstanceGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/ai/model-services/eval", uuid, "", map[string]interface{}{
		"deployModelEvalService": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachNvmeServerToCluster operates on NvmeServerToCluster
func (cli *ZSClient) AttachNvmeServerToCluster(ctx context.Context, params param.AttachNvmeServerToClusterParam) (*view.NvmeServerInventoryView, error) {
	resp := view.NvmeServerInventoryView{}
	if err := cli.Post(ctx, "v1/clusters/{clusterUuid}/storage-devices/nvme/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostResourceAllocation gets HostResourceAllocation by uuid
func (cli *ZSClient) GetHostResourceAllocation(ctx context.Context, params param.GetHostResourceAllocationParam) (*view.GetHostResourceAllocationEventView, error) {
	resp := view.GetHostResourceAllocationEventView{}
	if err := cli.Post(ctx, "v1/hosts/{uuid}/resource-allocation", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachUsbDeviceToVm operates on UsbDeviceToVm
func (cli *ZSClient) AttachUsbDeviceToVm(ctx context.Context, params param.AttachUsbDeviceToVmParam) (*view.UsbDeviceInventoryView, error) {
	resp := view.UsbDeviceInventoryView{}
	if err := cli.Post(ctx, "v1/usb-device/usb-devices/{usbDeviceUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseAddOns gets LicenseAddOns by uuid
func (cli *ZSClient) GetLicenseAddOns(ctx context.Context) (*view.GetLicenseAddOnsView, error) {
	var resp view.GetLicenseAddOnsView
	if err := cli.GetWithRespKey(ctx, "v1/licenses/addons", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAliyunSnapshotRemote operates on AliyunSnapshotRemote
func (cli *ZSClient) SyncAliyunSnapshotRemote(ctx context.Context, params param.SyncAliyunSnapshotRemoteParam) (*view.AliyunSnapshotInventoryView, error) {
	resp := view.AliyunSnapshotInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/snapshot/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunRouterInterfaceFromLocal queries AliyunRouterInterfaceFromLocal list
func (cli *ZSClient) QueryAliyunRouterInterfaceFromLocal(ctx context.Context, params *param.QueryParam) ([]view.AliyunRouterInterfaceInventoryView, error) {
	var resp []view.AliyunRouterInterfaceInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/router-interface", params, &resp)
}

func (cli *ZSClient) GetAliyunRouterInterfaceFromLocal(ctx context.Context, uuid string) (*view.AliyunRouterInterfaceInventoryView, error) {
	var resp view.AliyunRouterInterfaceInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/router-interface", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunRouterInterfaceFromLocal Pagination
func (cli *ZSClient) PageAliyunRouterInterfaceFromLocal(ctx context.Context, params *param.QueryParam) ([]view.AliyunRouterInterfaceInventoryView, int, error) {
	var aliyunRouterInterfaceFromLocals []view.AliyunRouterInterfaceInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/router-interface", params, &aliyunRouterInterfaceFromLocals)
	return aliyunRouterInterfaceFromLocals, total, err
}

// UpdateTicketRequest updates TicketRequest
func (cli *ZSClient) UpdateTicketRequest(ctx context.Context, uuid string, params param.UpdateTicketRequestParam) (*view.TicketInventoryView, error) {
	resp := view.TicketInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/tickets", uuid, "", map[string]interface{}{
		"updateTicketRequest": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcIPsecLog gets VpcIPsecLog by uuid
func (cli *ZSClient) GetVpcIPsecLog(ctx context.Context) (*view.GetVpcIPsecLogView, error) {
	var resp view.GetVpcIPsecLogView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers/ipseclog", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromOvf creates VmInstanceFromOvf
func (cli *ZSClient) CreateVmInstanceFromOvf(ctx context.Context, params param.CreateVmInstanceFromOvfParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/ovf/create-vm-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromOvfAsync Async
func (cli *ZSClient) CreateVmInstanceFromOvfAsync(ctx context.Context, params param.CreateVmInstanceFromOvfParam) (string, error) {

	resource := "v1/ovf/create-vm-instance"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// QueryLdapBinding queries LdapBinding list
func (cli *ZSClient) QueryLdapBinding(ctx context.Context, params *param.QueryParam) ([]view.LdapAccountRefInventoryView, error) {
	var resp []view.LdapAccountRefInventoryView
	return resp, cli.List(ctx, "v1/ldap/bindings", params, &resp)
}

func (cli *ZSClient) GetLdapBinding(ctx context.Context, uuid string) (*view.LdapAccountRefInventoryView, error) {
	var resp view.LdapAccountRefInventoryView
	if err := cli.Get(ctx, "v1/ldap/bindings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLdapBinding Pagination
func (cli *ZSClient) PageLdapBinding(ctx context.Context, params *param.QueryParam) ([]view.LdapAccountRefInventoryView, int, error) {
	var ldapBindings []view.LdapAccountRefInventoryView
	total, err := cli.Page(ctx, "v1/ldap/bindings", params, &ldapBindings)
	return ldapBindings, total, err
}

// ChangeVmImage changes VmImage
func (cli *ZSClient) ChangeVmImage(ctx context.Context, vmInstanceUuid string, params param.ChangeVmImageParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"changeVmImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddResourcesToDirectory adds ResourcesToDirectory
func (cli *ZSClient) AddResourcesToDirectory(ctx context.Context, params param.AddResourcesToDirectoryParam) (*view.AddResourcesToDirectoryEventView, error) {
	resp := view.AddResourcesToDirectoryEventView{}
	if err := cli.Post(ctx, "v1/add/resources/directory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachGuestToolsIsoToVm operates on GuestToolsIsoToVm
func (cli *ZSClient) AttachGuestToolsIsoToVm(ctx context.Context, uuid string, params param.AttachGuestToolsIsoToVmParam) (*view.AttachGuestToolsIsoToVmEventView, error) {
	resp := view.AttachGuestToolsIsoToVmEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"attachGuestToolsIsoToVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachAliyunKey operates on AliyunKey
func (cli *ZSClient) DetachAliyunKey(ctx context.Context, uuid string, params param.DetachAliyunKeyParam) (*view.DetachAliyunKeyEventView, error) {
	resp := view.DetachAliyunKeyEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/key", uuid, "", map[string]interface{}{
		"detachAliyunKey": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBuildApp deletes BuildApp
func (cli *ZSClient) DeleteBuildApp(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/appcenter/buildapp", uuid, string(deleteMode))
}

// ChangeBareMetal2InstancePassword changes BareMetal2InstancePassword
func (cli *ZSClient) ChangeBareMetal2InstancePassword(ctx context.Context, uuid string, params param.ChangeBareMetal2InstancePasswordParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/bm-instances", uuid, "", map[string]interface{}{
		"changeBareMetal2InstancePassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceFromPublishApp gets ResourceFromPublishApp by uuid
func (cli *ZSClient) GetResourceFromPublishApp(ctx context.Context) (*view.GetResourceFromPublishAppView, error) {
	var resp view.GetResourceFromPublishAppView
	if err := cli.GetWithRespKey(ctx, "v1/appcenter/app/resources", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeResourceOwner changes ResourceOwner
func (cli *ZSClient) ChangeResourceOwner(ctx context.Context, params param.ChangeResourceOwnerParam) (*view.AccountResourceRefInventoryView, error) {
	resp := view.AccountResourceRefInventoryView{}
	if err := cli.Post(ctx, "v1/account/{accountUuid}/resources", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostIommuState gets HostIommuState by uuid
func (cli *ZSClient) GetHostIommuState(ctx context.Context, uuid string) (*view.GetHostIommuStateView, error) {
	var resp view.GetHostIommuStateView
	if err := cli.GetWithRespKey(ctx, "v1/pci-device/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVirtualBorderRouterLocal deletes VirtualBorderRouterLocal
func (cli *ZSClient) DeleteVirtualBorderRouterLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/border-router", uuid, string(deleteMode))
}

// GetMetricData gets MetricData by uuid
func (cli *ZSClient) GetMetricData(ctx context.Context) (*view.GetMetricDataView, error) {
	var resp view.GetMetricDataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/metrics", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunSnapshotFromLocal queries AliyunSnapshotFromLocal list
func (cli *ZSClient) QueryAliyunSnapshotFromLocal(ctx context.Context, params *param.QueryParam) ([]view.AliyunSnapshotInventoryView, error) {
	var resp []view.AliyunSnapshotInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/snapshot", params, &resp)
}

func (cli *ZSClient) GetAliyunSnapshotFromLocal(ctx context.Context, uuid string) (*view.AliyunSnapshotInventoryView, error) {
	var resp view.AliyunSnapshotInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/snapshot", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunSnapshotFromLocal Pagination
func (cli *ZSClient) PageAliyunSnapshotFromLocal(ctx context.Context, params *param.QueryParam) ([]view.AliyunSnapshotInventoryView, int, error) {
	var aliyunSnapshotFromLocals []view.AliyunSnapshotInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/snapshot", params, &aliyunSnapshotFromLocals)
	return aliyunSnapshotFromLocals, total, err
}

// EnableCbtTask operates on EnableCbtTask
func (cli *ZSClient) EnableCbtTask(ctx context.Context, params param.EnableCbtTaskParam) (*view.EnableCbtTaskEventView, error) {
	resp := view.EnableCbtTaskEventView{}
	if err := cli.Post(ctx, "v1/cbt-task/enable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAliyunNasAccessGroupRemote gets AliyunNasAccessGroupRemote by uuid
func (cli *ZSClient) GetAliyunNasAccessGroupRemote(ctx context.Context) (*view.GetAliyunNasAccessGroupRemoteView, error) {
	var resp view.GetAliyunNasAccessGroupRemoteView
	if err := cli.GetWithRespKey(ctx, "v1/nas/aliyun/access/remote", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckBuildAppParameters operates on BuildAppParameters
func (cli *ZSClient) CheckBuildAppParameters(ctx context.Context, params param.CheckBuildAppParametersParam) (*view.CheckBuildAppParametersView, error) {
	resp := view.CheckBuildAppParametersView{}
	if err := cli.Post(ctx, "v1/appcenter/buildapp/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddLabelToEventSubscription adds LabelToEventSubscription
func (cli *ZSClient) AddLabelToEventSubscription(ctx context.Context, params param.AddLabelToEventSubscriptionParam) (*view.EventSubscriptionLabelInventoryView, error) {
	resp := view.EventSubscriptionLabelInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/events/subscriptions/{subscriptionUuid}/labels", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcVRouterDistributedRoutingConnections gets VpcVRouterDistributedRoutingConnections by uuid
func (cli *ZSClient) GetVpcVRouterDistributedRoutingConnections(ctx context.Context, uuid string) (*view.GetVpcVRouterDistributedRoutingConnectionsView, error) {
	var resp view.GetVpcVRouterDistributedRoutingConnectionsView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateThirdpartyAlerts updates ThirdpartyAlerts
func (cli *ZSClient) UpdateThirdpartyAlerts(ctx context.Context, params param.UpdateThirdpartyAlertsParam) (*view.UpdateThirdpartyAlertsEventView, error) {
	resp := view.UpdateThirdpartyAlertsEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/third-party/alerts/actions", "", "", map[string]interface{}{
		"updateThirdpartyAlerts": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PullSdnControllerTenant operates on PullSdnControllerTenant
func (cli *ZSClient) PullSdnControllerTenant(ctx context.Context, uuid string, params param.PullSdnControllerTenantParam) (*view.H3cSdnControllerTenantInventoryView, error) {
	resp := view.H3cSdnControllerTenantInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sdn-controllers", uuid, "", map[string]interface{}{
		"pullSdnControllerTenant": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachServiceFromObservabilityServer operates on ServiceFromObservabilityServer
func (cli *ZSClient) DetachServiceFromObservabilityServer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/observability-server", uuid, string(deleteMode))
}

// SetVmUsbRedirect operates on VmUsbRedirect
func (cli *ZSClient) SetVmUsbRedirect(ctx context.Context, uuid string, params param.SetVmUsbRedirectParam) (*view.SetVmUsbRedirectEventView, error) {
	resp := view.SetVmUsbRedirectEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmUsbRedirect": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostCandidatesForVmMigration gets HostCandidatesForVmMigration by uuid
func (cli *ZSClient) GetHostCandidatesForVmMigration(ctx context.Context, uuid string) (*view.GetHostCandidatesForVmMigrationView, error) {
	var resp view.GetHostCandidatesForVmMigrationView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmNicAttachableEips gets VmNicAttachableEips by uuid
func (cli *ZSClient) GetVmNicAttachableEips(ctx context.Context, uuid string) (*view.GetVmNicAttachableEipsView, error) {
	var resp view.GetVmNicAttachableEipsView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/nics", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFactoryModeState updates FactoryModeState
func (cli *ZSClient) UpdateFactoryModeState(ctx context.Context, params param.UpdateFactoryModeStateParam) (*view.UpdateFactoryModeStateEventView, error) {
	resp := view.UpdateFactoryModeStateEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/actions", "", "", map[string]interface{}{
		"updateFactoryModeState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateChronyServers updates ChronyServers
func (cli *ZSClient) UpdateChronyServers(ctx context.Context, params param.UpdateChronyServersParam) (*view.UpdateChronyServersEventView, error) {
	resp := view.UpdateChronyServersEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zops/chrony/actions", "", "", map[string]interface{}{
		"updateChronyServers": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPolicyRouteRuleSetToL3 operates on PolicyRouteRuleSetToL3
func (cli *ZSClient) AttachPolicyRouteRuleSetToL3(ctx context.Context, params param.AttachPolicyRouteRuleSetToL3Param) (*view.AttachPolicyRouteRuleSetToL3EventView, error) {
	resp := view.AttachPolicyRouteRuleSetToL3EventView{}
	if err := cli.Post(ctx, "v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOAuthClient updates OAuthClient
func (cli *ZSClient) UpdateOAuthClient(ctx context.Context, params param.UpdateOAuthClientParam) (*view.OAuth2ClientInventoryView, error) {
	resp := view.OAuth2ClientInventoryView{}
	if err := cli.Post(ctx, "v1/update/oauth2/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetZWatchAlertHistogram gets ZWatchAlertHistogram by uuid
func (cli *ZSClient) GetZWatchAlertHistogram(ctx context.Context) (*view.GetZWatchAlertHistogramView, error) {
	var resp view.GetZWatchAlertHistogramView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/alert-histories/histogram", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunRouterInterfaceRemote deletes AliyunRouterInterfaceRemote
func (cli *ZSClient) DeleteAliyunRouterInterfaceRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/router-interface/remote", uuid, string(deleteMode))
}

// SetImageBootMode operates on ImageBootMode
func (cli *ZSClient) SetImageBootMode(ctx context.Context, uuid string, params param.SetImageBootModeParam) (*view.SetImageBootModeEventView, error) {
	resp := view.SetImageBootModeEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/images", uuid, "", map[string]interface{}{
		"setImageBootMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachAutoScalingTemplateFromGroup operates on AutoScalingTemplateFromGroup
func (cli *ZSClient) DetachAutoScalingTemplateFromGroup(ctx context.Context, templateUuid string, groupUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/autoscaling/template", templateUuid, fmt.Sprintf("groups/%s", groupUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// UpdateVirtualBorderRouterRemote updates VirtualBorderRouterRemote
func (cli *ZSClient) UpdateVirtualBorderRouterRemote(ctx context.Context, uuid string, params param.UpdateVirtualBorderRouterRemoteParam) (*view.VirtualBorderRouterInventoryView, error) {
	resp := view.VirtualBorderRouterInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/border-router", uuid, "", map[string]interface{}{
		"updateVirtualBorderRouterRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmsCapabilities gets VmsCapabilities by uuid
func (cli *ZSClient) GetVmsCapabilities(ctx context.Context) (*view.GetVmsCapabilitiesView, error) {
	var resp view.GetVmsCapabilitiesView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/capabilities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPolicyToUserGroup operates on PolicyToUserGroup
func (cli *ZSClient) AttachPolicyToUserGroup(ctx context.Context, params param.AttachPolicyToUserGroupParam) (*view.AttachPolicyToUserGroupEventView, error) {
	resp := view.AttachPolicyToUserGroupEventView{}
	if err := cli.Post(ctx, "v1/accounts/groups/{groupUuid}/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevokeMonitorTemplateFromMonitorGroup operates on RevokeMonitorTemplateFromMonitorGroup
func (cli *ZSClient) RevokeMonitorTemplateFromMonitorGroup(ctx context.Context, templateUuid string, groupUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zwatch/monitortemplates", templateUuid, fmt.Sprintf("monitorgroups/%s", groupUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DeleteFirewallRule deletes FirewallRule
func (cli *ZSClient) DeleteFirewallRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpcfirewalls/rules", uuid, string(deleteMode))
}

// ShareResource operates on ShareResource
func (cli *ZSClient) ShareResource(ctx context.Context, params param.ShareResourceParam) (*view.ShareResourceEventView, error) {
	resp := view.ShareResourceEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/accounts/resources/actions", "", "", map[string]interface{}{
		"shareResource": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsVpcRemote creates EcsVpcRemote
func (cli *ZSClient) CreateEcsVpcRemote(ctx context.Context, params param.CreateEcsVpcRemoteParam) (*view.EcsVpcInventoryView, error) {
	resp := view.EcsVpcInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/vpc", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAccountQuotaUsage gets AccountQuotaUsage by uuid
func (cli *ZSClient) GetAccountQuotaUsage(ctx context.Context, uuid string) (*view.GetAccountQuotaUsageView, error) {
	var resp view.GetAccountQuotaUsageView
	if err := cli.GetWithRespKey(ctx, "v1/accounts/quota", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryPluginDrivers queries PluginDrivers list
func (cli *ZSClient) QueryPluginDrivers(ctx context.Context, params *param.QueryParam) ([]view.PluginDriverInventoryView, error) {
	var resp []view.PluginDriverInventoryView
	return resp, cli.List(ctx, "v1/external/plugins", params, &resp)
}

func (cli *ZSClient) GetPluginDrivers(ctx context.Context, uuid string) (*view.PluginDriverInventoryView, error) {
	var resp view.PluginDriverInventoryView
	if err := cli.Get(ctx, "v1/external/plugins", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePluginDrivers Pagination
func (cli *ZSClient) PagePluginDrivers(ctx context.Context, params *param.QueryParam) ([]view.PluginDriverInventoryView, int, error) {
	var pluginDrivers []view.PluginDriverInventoryView
	total, err := cli.Page(ctx, "v1/external/plugins", params, &pluginDrivers)
	return pluginDrivers, total, err
}

// RemoveIAM2VirtualIDsFromProjects removes IAM2VirtualIDsFromProjects
func (cli *ZSClient) RemoveIAM2VirtualIDsFromProjects(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects/virtual-ids", uuid, string(deleteMode))
}

// GetCandidateL3NetworksForServerGroup gets CandidateL3NetworksForServerGroup by uuid
func (cli *ZSClient) GetCandidateL3NetworksForServerGroup(ctx context.Context) (*view.GetCandidateL3NetworksForServerGroupView, error) {
	var resp view.GetCandidateL3NetworksForServerGroupView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers/servergroups/candidate-l3network", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmFromCdpBackup creates VmFromCdpBackup
func (cli *ZSClient) CreateVmFromCdpBackup(ctx context.Context, params param.CreateVmFromCdpBackupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/cdp-backups/actions", "", "", map[string]interface{}{
		"createVmFromCdpBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmFromCdpBackupAsync Async
func (cli *ZSClient) CreateVmFromCdpBackupAsync(ctx context.Context, params param.CreateVmFromCdpBackupParam) (string, error) {

	resource := "v1/cdp-backups/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// SyncVpcVpnConnectionFromRemote operates on VpcVpnConnectionFromRemote
func (cli *ZSClient) SyncVpcVpnConnectionFromRemote(ctx context.Context, dataCenterUuid string, params param.SyncVpcVpnConnectionFromRemoteParam) (*view.VpcVpnConnectionInventoryView, error) {
	resp := view.VpcVpnConnectionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/vpn-connection", dataCenterUuid, "", map[string]interface{}{
		"syncVpcVpnConnectionFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpnIkeConfig creates VpnIkeConfig
func (cli *ZSClient) CreateVpnIkeConfig(ctx context.Context, params param.CreateVpnIkeConfigParam) (*view.VpcVpnIkeConfigInventoryView, error) {
	resp := view.VpcVpnIkeConfigInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/vpn-connection/ike", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SubmitLongJob operates on SubmitLongJob
func (cli *ZSClient) SubmitLongJob(ctx context.Context, params param.SubmitLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Post(ctx, "v1/longjobs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolumeBackup creates DataVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeBackup(ctx context.Context, params param.CreateDataVolumeTemplateFromVolumeBackupParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(ctx, "v1/images/data-volume-templates/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DegradeFromLicenseServer operates on DegradeFromLicenseServer
func (cli *ZSClient) DegradeFromLicenseServer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/license-server", uuid, string(deleteMode))
}

// GetDebugSignal gets DebugSignal by uuid
func (cli *ZSClient) GetDebugSignal(ctx context.Context) (*view.GetDebugSignalView, error) {
	var resp view.GetDebugSignalView
	if err := cli.GetWithRespKey(ctx, "v1/debug", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAliyunKeySecret updates AliyunKeySecret
func (cli *ZSClient) UpdateAliyunKeySecret(ctx context.Context, params param.UpdateAliyunKeySecretParam) (*view.HybridAccountInventoryView, error) {
	resp := view.HybridAccountInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/{uuid}/key", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncLicenseCapacity operates on LicenseCapacity
func (cli *ZSClient) SyncLicenseCapacity(ctx context.Context) (*view.SyncLicenseCapacityEventView, error) {
	resp := view.SyncLicenseCapacityEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/license-server/authorized-capacity/sync", "", "", map[string]interface{}{
		"syncLicenseCapacity": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachDataVolumeToHost operates on DataVolumeToHost
func (cli *ZSClient) AttachDataVolumeToHost(ctx context.Context, params param.AttachDataVolumeToHostParam) (*view.AttachDataVolumeToHostEventView, error) {
	resp := view.AttachDataVolumeToHostEventView{}
	if err := cli.Post(ctx, "v1/volumes/{volumeUuid}/hosts/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SecurityMachineEncrypt operates on MachineEncrypt
func (cli *ZSClient) SecurityMachineEncrypt(ctx context.Context, params param.SecurityMachineEncryptParam) (*view.SecurityMachineEncryptEventView, error) {
	resp := view.SecurityMachineEncryptEventView{}
	if err := cli.Post(ctx, "v1/security-machine/encrypt/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAppBuildSystemState changes AppBuildSystemState
func (cli *ZSClient) ChangeAppBuildSystemState(ctx context.Context, uuid string, params param.ChangeAppBuildSystemStateParam) (*view.AppBuildSystemInventoryView, error) {
	resp := view.AppBuildSystemInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/appcenter/buildsystem", uuid, "", map[string]interface{}{
		"changeAppBuildSystemState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMemorySnapshotGroupReference gets MemorySnapshotGroupReference by uuid
func (cli *ZSClient) GetMemorySnapshotGroupReference(ctx context.Context) (*view.GetMemorySnapshotGroupReferenceView, error) {
	var resp view.GetMemorySnapshotGroupReferenceView
	if err := cli.GetWithRespKey(ctx, "v1/memory-snapshots/group/reference", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcVRouterNetworkServiceState gets VpcVRouterNetworkServiceState by uuid
func (cli *ZSClient) GetVpcVRouterNetworkServiceState(ctx context.Context, uuid string) (*view.GetVpcVRouterNetworkServiceStateView, error) {
	var resp view.GetVpcVRouterNetworkServiceStateView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachNetworkServiceFromL3Network operates on NetworkServiceFromL3Network
func (cli *ZSClient) DetachNetworkServiceFromL3Network(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks", uuid, string(deleteMode))
}

// CreateDataVolumeFromVolumeBackup creates DataVolumeFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeFromVolumeBackup(ctx context.Context, params param.CreateDataVolumeFromVolumeBackupParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.Post(ctx, "v1/volumes/data-volume/from/volume-template/{backupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteContainerResourceFromEndpoint deletes ContainerResourceFromEndpoint
func (cli *ZSClient) DeleteContainerResourceFromEndpoint(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/container/management/endpoint", uuid, string(deleteMode))
}

// GetSupportAPIs gets SupportAPIs by uuid
func (cli *ZSClient) GetSupportAPIs(ctx context.Context) (*view.GetSupportAPIsView, error) {
	resp := view.GetSupportAPIsView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/actions", "", "", map[string]interface{}{
		"getSupportAPIs": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSharedMountPointPrimaryStorage adds SharedMountPointPrimaryStorage
func (cli *ZSClient) AddSharedMountPointPrimaryStorage(ctx context.Context) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/smp", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTrashOnPrimaryStorage gets TrashOnPrimaryStorage by uuid
func (cli *ZSClient) GetTrashOnPrimaryStorage(ctx context.Context) (*view.GetTrashOnPrimaryStorageView, error) {
	var resp view.GetTrashOnPrimaryStorageView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/trash", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicsForLoadBalancer gets CandidateVmNicsForLoadBalancer by uuid
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancer(ctx context.Context, uuid string) (*view.GetCandidateVmNicsForLoadBalancerView, error) {
	var resp view.GetCandidateVmNicsForLoadBalancerView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers/listeners", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachBaremetalPxeServerFromCluster operates on BaremetalPxeServerFromCluster
func (cli *ZSClient) DetachBaremetalPxeServerFromCluster(ctx context.Context, clusterUuid string, pxeServerUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/clusters", clusterUuid, fmt.Sprintf("pxeservers/%s", pxeServerUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DeleteVpcUserVpnGatewayRemote deletes VpcUserVpnGatewayRemote
func (cli *ZSClient) DeleteVpcUserVpnGatewayRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/user-gateway", uuid, string(deleteMode))
}

// ChangeAccessControlListRedirectRule changes AccessControlListRedirectRule
func (cli *ZSClient) ChangeAccessControlListRedirectRule(ctx context.Context, uuid string, params param.ChangeAccessControlListRedirectRuleParam) (*view.AccessControlListEntryInventoryView, error) {
	resp := view.AccessControlListEntryInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/access-control-lists/redirectRules", uuid, "", map[string]interface{}{
		"changeAccessControlListRedirectRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddResourceStackVmPortMonitor adds ResourceStackVmPortMonitor
func (cli *ZSClient) AddResourceStackVmPortMonitor(ctx context.Context, params param.AddResourceStackVmPortMonitorParam) (*view.AddResourceStackVmPortMonitorEventView, error) {
	resp := view.AddResourceStackVmPortMonitorEventView{}
	if err := cli.Post(ctx, "v1/cloudformation/stack/monitor/addvm", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSNSApplicationEndpointState changes SNSApplicationEndpointState
func (cli *ZSClient) ChangeSNSApplicationEndpointState(ctx context.Context, uuid string, params param.ChangeSNSApplicationEndpointStateParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-endpoints", uuid, "", map[string]interface{}{
		"changeSNSApplicationEndpointState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedLoadBalancer gets VpcAttachedLoadBalancer by uuid
func (cli *ZSClient) GetVpcAttachedLoadBalancer(ctx context.Context, params param.GetVpcAttachedLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	resp := view.LoadBalancerInventoryView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/attached-lb", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpcVpnConnectionRemote creates VpcVpnConnectionRemote
func (cli *ZSClient) CreateVpcVpnConnectionRemote(ctx context.Context, params param.CreateVpcVpnConnectionRemoteParam) (*view.VpcVpnConnectionInventoryView, error) {
	resp := view.VpcVpnConnectionInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/vpn-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedPortForwardingRules gets VpcAttachedPortForwardingRules by uuid
func (cli *ZSClient) GetVpcAttachedPortForwardingRules(ctx context.Context, params param.GetVpcAttachedPortForwardingRulesParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/attached-pf", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVpcVRouterNetworkServiceState operates on VpcVRouterNetworkServiceState
func (cli *ZSClient) SetVpcVRouterNetworkServiceState(ctx context.Context, params param.SetVpcVRouterNetworkServiceStateParam) (*view.SetVpcVRouterNetworkServiceStateEventView, error) {
	resp := view.SetVpcVRouterNetworkServiceStateEventView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/networkservicestate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDnsToVpcRouter adds DnsToVpcRouter
func (cli *ZSClient) AddDnsToVpcRouter(ctx context.Context, params param.AddDnsToVpcRouterParam) (*view.VpcRouterVmInventoryView, error) {
	resp := view.VpcRouterVmInventoryView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/dns", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAccountBilling queries AccountBilling list
func (cli *ZSClient) QueryAccountBilling(ctx context.Context, params *param.QueryParam) ([]view.BillingInventoryView, error) {
	var resp []view.BillingInventoryView
	return resp, cli.List(ctx, "v1/billing/billings", params, &resp)
}

func (cli *ZSClient) GetAccountBilling(ctx context.Context, uuid string) (*view.BillingInventoryView, error) {
	var resp view.BillingInventoryView
	if err := cli.Get(ctx, "v1/billing/billings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccountBilling Pagination
func (cli *ZSClient) PageAccountBilling(ctx context.Context, params *param.QueryParam) ([]view.BillingInventoryView, int, error) {
	var accountBillings []view.BillingInventoryView
	total, err := cli.Page(ctx, "v1/billing/billings", params, &accountBillings)
	return accountBillings, total, err
}

// GetVmXml gets VmXml by uuid
func (cli *ZSClient) GetVmXml(ctx context.Context, uuid string) (*view.GetVmXmlView, error) {
	var resp view.GetVmXmlView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceFirstBootDevice gets VmInstanceFirstBootDevice by uuid
func (cli *ZSClient) GetVmInstanceFirstBootDevice(ctx context.Context, uuid string) (*view.GetVmInstanceFirstBootDeviceView, error) {
	var resp view.GetVmInstanceFirstBootDeviceView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOvnControllerVm creates OvnControllerVm
func (cli *ZSClient) CreateOvnControllerVm(ctx context.Context) (*view.ApplianceVmInventoryView, error) {
	resp := view.ApplianceVmInventoryView{}
	if err := cli.Post(ctx, "v1/ovn/instances", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIpAddress deletes IpAddress
func (cli *ZSClient) DeleteIpAddress(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks", uuid, string(deleteMode))
}

// DeleteVpcVpnConnectionRemote deletes VpcVpnConnectionRemote
func (cli *ZSClient) DeleteVpcVpnConnectionRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/vpn-connection", uuid, string(deleteMode))
}

// AttachOssBucketToEcsDataCenter operates on OssBucketToEcsDataCenter
func (cli *ZSClient) AttachOssBucketToEcsDataCenter(ctx context.Context, ossBucketUuid string, params param.AttachOssBucketToEcsDataCenterParam) (*view.OssBucketInventoryView, error) {
	resp := view.OssBucketInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/oss-bucket", ossBucketUuid, "", map[string]interface{}{
		"attachOssBucketToEcsDataCenter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIAM2OrganizationAvailability operates on IAM2OrganizationAvailability
func (cli *ZSClient) CheckIAM2OrganizationAvailability(ctx context.Context) (*view.CheckIAM2OrganizationAvailabilityView, error) {
	var resp view.CheckIAM2OrganizationAvailabilityView
	if err := cli.GetWithRespKey(ctx, "v1/iam2/organizations/availabilities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnmountVmInstanceRecoveryPoint operates on UnmountVmInstanceRecoveryPoint
func (cli *ZSClient) UnmountVmInstanceRecoveryPoint(ctx context.Context, params param.UnmountVmInstanceRecoveryPointParam) (*view.UnmountVmInstanceRecoveryPointEventView, error) {
	resp := view.UnmountVmInstanceRecoveryPointEventView{}
	if err := cli.Post(ctx, "v1/cdp-backup-storage/unmount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemovePolicyStatementsFromRole removes PolicyStatementsFromRole
func (cli *ZSClient) RemovePolicyStatementsFromRole(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/identities/roles", uuid, string(deleteMode))
}

// GenerateModelMetadata operates on ModelMetadata
func (cli *ZSClient) GenerateModelMetadata(ctx context.Context, params param.GenerateModelMetadataParam) (*view.GenerateModelMetadataEventView, error) {
	resp := view.GenerateModelMetadataEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/ai/model/metadata/generate", "", "", map[string]interface{}{
		"generateModelMetadata": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsReadyToGo operates on IsReadyToGo
func (cli *ZSClient) IsReadyToGo(ctx context.Context) (*view.IsReadyToGoView, error) {
	var resp view.IsReadyToGoView
	if err := cli.GetWithRespKey(ctx, "v1/management-nodes/ready", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostIommuStatus gets HostIommuStatus by uuid
func (cli *ZSClient) GetHostIommuStatus(ctx context.Context, uuid string) (*view.GetHostIommuStatusView, error) {
	var resp view.GetHostIommuStatusView
	if err := cli.GetWithRespKey(ctx, "v1/pci-device/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DescribeVmInstanceRecoveryPoint operates on DescribeVmInstanceRecoveryPoint
func (cli *ZSClient) DescribeVmInstanceRecoveryPoint(ctx context.Context, uuid string) (*view.DescribeVmInstanceRecoveryPointView, error) {
	var resp view.DescribeVmInstanceRecoveryPointView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceCandidatesForAttachingVm gets PciDeviceCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetPciDeviceCandidatesForAttachingVm(ctx context.Context, uuid string) (*view.GetPciDeviceCandidatesForAttachingVmView, error) {
	var resp view.GetPciDeviceCandidatesForAttachingVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeMonitorTriggerState changes MonitorTriggerState
func (cli *ZSClient) ChangeMonitorTriggerState(ctx context.Context, uuid string, params param.ChangeMonitorTriggerStateParam) (*view.MonitorTriggerInventoryView, error) {
	resp := view.MonitorTriggerInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/monitoring/triggers", uuid, "", map[string]interface{}{
		"changeMonitorTriggerState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBareMetal2ChassisPowerStatus gets BareMetal2ChassisPowerStatus by uuid
func (cli *ZSClient) GetBareMetal2ChassisPowerStatus(ctx context.Context, uuid string) (*view.GetBareMetal2ChassisPowerStatusView, error) {
	var resp view.GetBareMetal2ChassisPowerStatusView
	if err := cli.GetWithRespKey(ctx, "v1/baremetal2/chassis", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTaskProgress gets TaskProgress by uuid
func (cli *ZSClient) GetTaskProgress(ctx context.Context, uuid string) (*view.GetTaskProgressView, error) {
	var resp view.GetTaskProgressView
	if err := cli.GetWithRespKey(ctx, "v1/task-progresses", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartDataProtection starts DataProtection
func (cli *ZSClient) StartDataProtection(ctx context.Context, params param.StartDataProtectionParam) (*view.StartDataProtectionEventView, error) {
	resp := view.StartDataProtectionEventView{}
	if err := cli.Post(ctx, "v1/start/data/protection/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartDataProtectionAsync Async
func (cli *ZSClient) StartDataProtectionAsync(ctx context.Context, params param.StartDataProtectionParam) (string, error) {

	resource := "v1/start/data/protection/"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// ChangeActiveAlarmState changes ActiveAlarmState
func (cli *ZSClient) ChangeActiveAlarmState(ctx context.Context, params param.ChangeActiveAlarmStateParam) (*view.ChangeActiveAlarmStateEventView, error) {
	resp := view.ChangeActiveAlarmStateEventView{}
	if err := cli.Post(ctx, "v1/zwatch/activealarms/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmCleanTraffic operates on VmCleanTraffic
func (cli *ZSClient) SetVmCleanTraffic(ctx context.Context, uuid string, params param.SetVmCleanTrafficParam) (*view.SetVmCleanTrafficEventView, error) {
	resp := view.SetVmCleanTrafficEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmCleanTraffic": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmBootMode operates on VmBootMode
func (cli *ZSClient) SetVmBootMode(ctx context.Context, uuid string, params param.SetVmBootModeParam) (*view.SetVmBootModeEventView, error) {
	resp := view.SetVmBootModeEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmBootMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncImageSize operates on ImageSize
func (cli *ZSClient) SyncImageSize(ctx context.Context, uuid string, params param.SyncImageSizeParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/images", uuid, "", map[string]interface{}{
		"syncImageSize": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNoTriggerSchedulerJobs gets NoTriggerSchedulerJobs by uuid
func (cli *ZSClient) GetNoTriggerSchedulerJobs(ctx context.Context) (*view.GetNoTriggerSchedulerJobsView, error) {
	var resp view.GetNoTriggerSchedulerJobsView
	if err := cli.GetWithRespKey(ctx, "v1/scheduler/jobs/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddProxyToResource adds ProxyToResource
func (cli *ZSClient) AddProxyToResource(ctx context.Context, params param.AddProxyToResourceParam) (*view.UserProxyConfigResourceRefInventoryView, error) {
	resp := view.UserProxyConfigResourceRefInventoryView{}
	if err := cli.Post(ctx, "v1/proxy/{proxyUuid}/resource/{resourceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProtectVmInstanceRecoveryPoint operates on ProtectVmInstanceRecoveryPoint
func (cli *ZSClient) ProtectVmInstanceRecoveryPoint(ctx context.Context, vmInstanceUuid string, params param.ProtectVmInstanceRecoveryPointParam) (*view.ProtectVmInstanceRecoveryPointEventView, error) {
	resp := view.ProtectVmInstanceRecoveryPointEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"protectVmInstanceRecoveryPoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteConnectionAccessPointLocal deletes ConnectionAccessPointLocal
func (cli *ZSClient) DeleteConnectionAccessPointLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/access-point", uuid, string(deleteMode))
}

// QueryPhysicalDriveSelfTestHistory queries PhysicalDriveSelfTestHistory list
func (cli *ZSClient) QueryPhysicalDriveSelfTestHistory(ctx context.Context, params *param.QueryParam) ([]view.PhysicalDriveSmartSelfTestHistoryInventoryView, error) {
	var resp []view.PhysicalDriveSmartSelfTestHistoryInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/local-raid/physical-drives/self-test", params, &resp)
}

func (cli *ZSClient) GetPhysicalDriveSelfTestHistory(ctx context.Context, uuid string) (*view.PhysicalDriveSmartSelfTestHistoryInventoryView, error) {
	var resp view.PhysicalDriveSmartSelfTestHistoryInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/local-raid/physical-drives/self-test", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePhysicalDriveSelfTestHistory Pagination
func (cli *ZSClient) PagePhysicalDriveSelfTestHistory(ctx context.Context, params *param.QueryParam) ([]view.PhysicalDriveSmartSelfTestHistoryInventoryView, int, error) {
	var physicalDriveSelfTestHistories []view.PhysicalDriveSmartSelfTestHistoryInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/local-raid/physical-drives/self-test", params, &physicalDriveSelfTestHistories)
	return physicalDriveSelfTestHistories, total, err
}

// RemoveIAM2VirtualIDsFromProject removes IAM2VirtualIDsFromProject
func (cli *ZSClient) RemoveIAM2VirtualIDsFromProject(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects", uuid, string(deleteMode))
}

// CreateEcsImageFromEcsSnapshot creates EcsImageFromEcsSnapshot
func (cli *ZSClient) CreateEcsImageFromEcsSnapshot(ctx context.Context, params param.CreateEcsImageFromEcsSnapshotParam) (*view.EcsImageInventoryView, error) {
	resp := view.EcsImageInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/image/snapshot", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateResourceStackFromApp creates ResourceStackFromApp
func (cli *ZSClient) CreateResourceStackFromApp(ctx context.Context, params param.CreateResourceStackFromAppParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.Post(ctx, "v1/appcenter/app/stack", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSharedBlockCandidate gets SharedBlockCandidate by uuid
func (cli *ZSClient) GetSharedBlockCandidate(ctx context.Context) (*view.GetSharedBlockCandidateView, error) {
	var resp view.GetSharedBlockCandidateView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/sharedblockgroup/sharedblock-candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsSecurityGroupFromRemote operates on EcsSecurityGroupFromRemote
func (cli *ZSClient) SyncEcsSecurityGroupFromRemote(ctx context.Context, ecsVpcUuid string, params param.SyncEcsSecurityGroupFromRemoteParam) (*view.EcsSecurityGroupInventoryView, error) {
	resp := view.EcsSecurityGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/security-group", ecsVpcUuid, "", map[string]interface{}{
		"syncEcsSecurityGroupFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportBuildApp operates on BuildApp
func (cli *ZSClient) ExportBuildApp(ctx context.Context, uuid string, params param.ExportBuildAppParam) (*view.BuildAppExportHistoryInventoryView, error) {
	resp := view.BuildAppExportHistoryInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/appcenter/buildapp", uuid, "", map[string]interface{}{
		"exportBuildApp": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReclaimSpaceFromImageStore operates on ReclaimSpaceFromImageStore
func (cli *ZSClient) ReclaimSpaceFromImageStore(ctx context.Context, uuid string, params param.ReclaimSpaceFromImageStoreParam) (*view.ReclaimSpaceFromImageStoreEventView, error) {
	resp := view.ReclaimSpaceFromImageStoreEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage/image-store", uuid, "", map[string]interface{}{
		"reclaimSpaceFromImageStore": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReclaimSpaceFromImageStoreAsync Async
func (cli *ZSClient) ReclaimSpaceFromImageStoreAsync(ctx context.Context, params param.ReclaimSpaceFromImageStoreParam) (string, error) {

	resource := "v1/backup-storage/image-store/{uuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetAllEventMetadata gets AllEventMetadata by uuid
func (cli *ZSClient) GetAllEventMetadata(ctx context.Context) (*view.GetAllEventMetadataView, error) {
	var resp view.GetAllEventMetadataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/events/meta-data", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmForAttachingIso gets CandidateVmForAttachingIso by uuid
func (cli *ZSClient) GetCandidateVmForAttachingIso(ctx context.Context, uuid string) (*view.GetCandidateVmForAttachingIsoView, error) {
	var resp view.GetCandidateVmForAttachingIsoView
	if err := cli.GetWithRespKey(ctx, "v1/images/iso", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachDataVolumeToVm operates on DataVolumeToVm
func (cli *ZSClient) AttachDataVolumeToVm(ctx context.Context, params param.AttachDataVolumeToVmParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.Post(ctx, "v1/volumes/{volumeUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAliyunVirtualRouter updates AliyunVirtualRouter
func (cli *ZSClient) UpdateAliyunVirtualRouter(ctx context.Context, uuid string, params param.UpdateAliyunVirtualRouterParam) (*view.VpcVirtualRouterInventoryView, error) {
	resp := view.VpcVirtualRouterInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/vrouter", uuid, "", map[string]interface{}{
		"updateAliyunVirtualRouter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDataVolume deletes DataVolume
func (cli *ZSClient) DeleteDataVolume(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volumes", uuid, string(deleteMode))
}

// GetUploadImageJobDetails gets UploadImageJobDetails by uuid
func (cli *ZSClient) GetUploadImageJobDetails(ctx context.Context, uuid string) (*view.GetUploadImageJobDetailsView, error) {
	var resp view.GetUploadImageJobDetailsView
	if err := cli.GetWithRespKey(ctx, "v1/images/upload-job/details", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachIscsiServerFromCluster operates on IscsiServerFromCluster
func (cli *ZSClient) DetachIscsiServerFromCluster(ctx context.Context, clusterUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/clusters", clusterUuid, fmt.Sprintf("storage-devices/iscsi/servers/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// SetVolumeQos operates on VolumeQos
func (cli *ZSClient) SetVolumeQos(ctx context.Context, uuid string, params param.SetVolumeQosParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/volumes", uuid, "", map[string]interface{}{
		"setVolumeQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachHybridEipFromEcs operates on HybridEipFromEcs
func (cli *ZSClient) DetachHybridEipFromEcs(ctx context.Context, params param.DetachHybridEipFromEcsParam) (*view.DetachHybridEipFromEcsEventView, error) {
	resp := view.DetachHybridEipFromEcsEventView{}
	if err := cli.Post(ctx, "v1/hybrid/eip/{eipUuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeCapabilities gets VolumeCapabilities by uuid
func (cli *ZSClient) GetVolumeCapabilities(ctx context.Context, uuid string) (*view.GetVolumeCapabilitiesView, error) {
	var resp view.GetVolumeCapabilitiesView
	if err := cli.GetWithRespKey(ctx, "v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2GatewayCluster changes BareMetal2GatewayCluster
func (cli *ZSClient) ChangeBareMetal2GatewayCluster(ctx context.Context, gatewayUuid string, params param.ChangeBareMetal2GatewayClusterParam) (*view.BareMetal2GatewayInventoryView, error) {
	resp := view.BareMetal2GatewayInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/gateways", gatewayUuid, "", map[string]interface{}{
		"changeBareMetal2GatewayCluster": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcIkeConfigFromLocal queries VpcIkeConfigFromLocal list
func (cli *ZSClient) QueryVpcIkeConfigFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVpnIkeConfigInventoryView, error) {
	var resp []view.VpcVpnIkeConfigInventoryView
	return resp, cli.List(ctx, "v1/hybrid/vpn-connection/ike", params, &resp)
}

func (cli *ZSClient) GetVpcIkeConfigFromLocal(ctx context.Context, uuid string) (*view.VpcVpnIkeConfigInventoryView, error) {
	var resp view.VpcVpnIkeConfigInventoryView
	if err := cli.Get(ctx, "v1/hybrid/vpn-connection/ike", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcIkeConfigFromLocal Pagination
func (cli *ZSClient) PageVpcIkeConfigFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVpnIkeConfigInventoryView, int, error) {
	var vpcIkeConfigFromLocals []view.VpcVpnIkeConfigInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/vpn-connection/ike", params, &vpcIkeConfigFromLocals)
	return vpcIkeConfigFromLocals, total, err
}

// SetVmInstanceHaLevel operates on VmInstanceHaLevel
func (cli *ZSClient) SetVmInstanceHaLevel(ctx context.Context, params param.SetVmInstanceHaLevelParam) (*view.SetVmInstanceHaLevelEventView, error) {
	resp := view.SetVmInstanceHaLevelEventView{}
	if err := cli.Post(ctx, "v1/vm-instances/{uuid}/ha-levels", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVRouterNetworksFromFlowMeter removes VRouterNetworksFromFlowMeter
func (cli *ZSClient) RemoveVRouterNetworksFromFlowMeter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/flowmeters/networks", uuid, string(deleteMode))
}

// GetCandidateL3NetworksForChangeVmNicNetwork gets CandidateL3NetworksForChangeVmNicNetwork by uuid
func (cli *ZSClient) GetCandidateL3NetworksForChangeVmNicNetwork(ctx context.Context, uuid string) (*view.GetCandidateL3NetworksForChangeVmNicNetworkView, error) {
	var resp view.GetCandidateL3NetworksForChangeVmNicNetworkView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/nics", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteHybridKeySecret deletes HybridKeySecret
func (cli *ZSClient) DeleteHybridKeySecret(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/hybrid/key", uuid, string(deleteMode))
}

// GetCandidatePrimaryStoragesForCreatingVm gets CandidatePrimaryStoragesForCreatingVm by uuid
func (cli *ZSClient) GetCandidatePrimaryStoragesForCreatingVm(ctx context.Context) (*view.GetCandidatePrimaryStoragesForCreatingVmView, error) {
	var resp view.GetCandidatePrimaryStoragesForCreatingVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/candidate-storages", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmConsolePassword gets VmConsolePassword by uuid
func (cli *ZSClient) GetVmConsolePassword(ctx context.Context, uuid string) (*view.GetVmConsolePasswordView, error) {
	var resp view.GetVmConsolePasswordView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceBindableConfig gets ResourceBindableConfig by uuid
func (cli *ZSClient) GetResourceBindableConfig(ctx context.Context) (*view.GetResourceBindableConfigView, error) {
	var resp view.GetResourceBindableConfigView
	if err := cli.GetWithRespKey(ctx, "v1/resource-configurations/bindable", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceHaLevel gets VmInstanceHaLevel by uuid
func (cli *ZSClient) GetVmInstanceHaLevel(ctx context.Context, uuid string) (*view.GetVmInstanceHaLevelView, error) {
	var resp view.GetVmInstanceHaLevelView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateLdapEntryForIAM2Binding gets CandidateLdapEntryForIAM2Binding by uuid
func (cli *ZSClient) GetCandidateLdapEntryForIAM2Binding(ctx context.Context) (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.GetWithRespKey(ctx, "v1/iam2/ldap/entries/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAudit queries Audit list
func (cli *ZSClient) QueryAudit(ctx context.Context, params *param.QueryParam) ([]view.AuditsInventoryView, error) {
	var resp []view.AuditsInventoryView
	return resp, cli.List(ctx, "v1/zwatch/audit-records", params, &resp)
}

func (cli *ZSClient) GetAudit(ctx context.Context, uuid string) (*view.AuditsInventoryView, error) {
	var resp view.AuditsInventoryView
	if err := cli.Get(ctx, "v1/zwatch/audit-records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAudit Pagination
func (cli *ZSClient) PageAudit(ctx context.Context, params *param.QueryParam) ([]view.AuditsInventoryView, int, error) {
	var audits []view.AuditsInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/audit-records", params, &audits)
	return audits, total, err
}

// RemoveResourcesFromDirectory removes ResourcesFromDirectory
func (cli *ZSClient) RemoveResourcesFromDirectory(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/remove/resources/directory", uuid, string(deleteMode))
}

// CreateVmFromVmBackup creates VmFromVmBackup
func (cli *ZSClient) CreateVmFromVmBackup(ctx context.Context, params param.CreateVmFromVmBackupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/from/vm-backups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteExportedDatabaseBackupFromBackupStorage deletes ExportedDatabaseBackupFromBackupStorage
func (cli *ZSClient) DeleteExportedDatabaseBackupFromBackupStorage(ctx context.Context, databaseBackupUuid string, backupStorageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/exported-database-backup", databaseBackupUuid, fmt.Sprintf("backup-storage/%s", backupStorageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AttachNetworkServiceToL3Network operates on NetworkServiceToL3Network
func (cli *ZSClient) AttachNetworkServiceToL3Network(ctx context.Context, params param.AttachNetworkServiceToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.Post(ctx, "v1/l3-networks/{l3NetworkUuid}/network-services", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnexportNbdVolumes operates on UnexportNbdVolumes
func (cli *ZSClient) UnexportNbdVolumes(ctx context.Context, params param.UnexportNbdVolumesParam) (*view.UnexportNbdVolumesEventView, error) {
	resp := view.UnexportNbdVolumesEventView{}
	if err := cli.Post(ctx, "v1/cbt-task/unexportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoveryVirtualBorderRouterRemote operates on yVirtualBorderRouterRemote
func (cli *ZSClient) RecoveryVirtualBorderRouterRemote(ctx context.Context, uuid string, params param.RecoveryVirtualBorderRouterRemoteParam) (*view.RecoveryVirtualBorderRouterRemoteEventView, error) {
	resp := view.RecoveryVirtualBorderRouterRemoteEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/border-router", uuid, "", map[string]interface{}{
		"recoveryVirtualBorderRouterRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcRouter queries VpcRouter list
func (cli *ZSClient) QueryVpcRouter(ctx context.Context, params *param.QueryParam) ([]view.VpcRouterVmInventoryView, error) {
	var resp []view.VpcRouterVmInventoryView
	return resp, cli.List(ctx, "v1/vpc/virtual-routers", params, &resp)
}

func (cli *ZSClient) GetVpcRouter(ctx context.Context, uuid string) (*view.VpcRouterVmInventoryView, error) {
	var resp view.VpcRouterVmInventoryView
	if err := cli.Get(ctx, "v1/vpc/virtual-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcRouter Pagination
func (cli *ZSClient) PageVpcRouter(ctx context.Context, params *param.QueryParam) ([]view.VpcRouterVmInventoryView, int, error) {
	var vpcRouters []view.VpcRouterVmInventoryView
	total, err := cli.Page(ctx, "v1/vpc/virtual-routers", params, &vpcRouters)
	return vpcRouters, total, err
}

// ExecuteAutoScalingRule operates on ExecuteAutoScalingRule
func (cli *ZSClient) ExecuteAutoScalingRule(ctx context.Context, uuid string, params param.ExecuteAutoScalingRuleParam) (*view.ExecuteAutoScalingRuleEventView, error) {
	resp := view.ExecuteAutoScalingRuleEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/autoscaling/rules", uuid, "", map[string]interface{}{
		"executeAutoScalingRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSHttpTestConnection operates on HttpTestConnection
func (cli *ZSClient) SNSHttpTestConnection(ctx context.Context, params param.SNSHttpTestConnectionParam) (*view.SNSHttpTestConnectionEventView, error) {
	resp := view.SNSHttpTestConnectionEventView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/http/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageSecurityLevel operates on ImageSecurityLevel
func (cli *ZSClient) SetImageSecurityLevel(ctx context.Context, uuid string, params param.SetImageSecurityLevelParam) (*view.SetImageSecurityLevelEventView, error) {
	resp := view.SetImageSecurityLevelEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/images", uuid, "", map[string]interface{}{
		"setImageSecurityLevel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2ChassisState changes BareMetal2ChassisState
func (cli *ZSClient) ChangeBareMetal2ChassisState(ctx context.Context, uuid string, params param.ChangeBareMetal2ChassisStateParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"changeBareMetal2ChassisState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryHybridEipFromLocal queries HybridEipFromLocal list
func (cli *ZSClient) QueryHybridEipFromLocal(ctx context.Context, params *param.QueryParam) ([]view.HybridEipAddressInventoryView, error) {
	var resp []view.HybridEipAddressInventoryView
	return resp, cli.List(ctx, "v1/hybrid/eip", params, &resp)
}

func (cli *ZSClient) GetHybridEipFromLocal(ctx context.Context, uuid string) (*view.HybridEipAddressInventoryView, error) {
	var resp view.HybridEipAddressInventoryView
	if err := cli.Get(ctx, "v1/hybrid/eip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHybridEipFromLocal Pagination
func (cli *ZSClient) PageHybridEipFromLocal(ctx context.Context, params *param.QueryParam) ([]view.HybridEipAddressInventoryView, int, error) {
	var hybridEipFromLocals []view.HybridEipAddressInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/eip", params, &hybridEipFromLocals)
	return hybridEipFromLocals, total, err
}

// AddHybridKeySecret adds HybridKeySecret
func (cli *ZSClient) AddHybridKeySecret(ctx context.Context, params param.AddHybridKeySecretParam) (*view.HybridAccountInventoryView, error) {
	resp := view.HybridAccountInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/hybrid/key", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunRouteEntryFromLocal queries AliyunRouteEntryFromLocal list
func (cli *ZSClient) QueryAliyunRouteEntryFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVirtualRouteEntryInventoryView, error) {
	var resp []view.VpcVirtualRouteEntryInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/route-entry", params, &resp)
}

func (cli *ZSClient) GetAliyunRouteEntryFromLocal(ctx context.Context, uuid string) (*view.VpcVirtualRouteEntryInventoryView, error) {
	var resp view.VpcVirtualRouteEntryInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/route-entry", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunRouteEntryFromLocal Pagination
func (cli *ZSClient) PageAliyunRouteEntryFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVirtualRouteEntryInventoryView, int, error) {
	var aliyunRouteEntryFromLocals []view.VpcVirtualRouteEntryInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/route-entry", params, &aliyunRouteEntryFromLocals)
	return aliyunRouteEntryFromLocals, total, err
}

// DetachVmFromVmSchedulingRuleGroup operates on VmFromVmSchedulingRuleGroup
func (cli *ZSClient) DetachVmFromVmSchedulingRuleGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vmSchedulingRuleGroup", uuid, string(deleteMode))
}

// AddVRouterNetworksToOspfArea adds VRouterNetworksToOspfArea
func (cli *ZSClient) AddVRouterNetworksToOspfArea(ctx context.Context, params param.AddVRouterNetworksToOspfAreaParam) (*view.NetworkRouterAreaRefInventoryView, error) {
	resp := view.NetworkRouterAreaRefInventoryView{}
	if err := cli.Post(ctx, "v1/routerArea/{routerAreaUuid}/router/{vRouterUuid}/addnetworks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QuerySNSSmsEndpoint queries SNSSmsEndpoint list
func (cli *ZSClient) QuerySNSSmsEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSAliyunSmsEndpointInventoryView, error) {
	var resp []view.SNSAliyunSmsEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/sms-endpoints", params, &resp)
}

func (cli *ZSClient) GetSNSSmsEndpoint(ctx context.Context, uuid string) (*view.SNSAliyunSmsEndpointInventoryView, error) {
	var resp view.SNSAliyunSmsEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/sms-endpoints", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSSmsEndpoint Pagination
func (cli *ZSClient) PageSNSSmsEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSAliyunSmsEndpointInventoryView, int, error) {
	var sNSSmsEndpoints []view.SNSAliyunSmsEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/sms-endpoints", params, &sNSSmsEndpoints)
	return sNSSmsEndpoints, total, err
}

// AddRolesToIAM2VirtualIDGroup adds RolesToIAM2VirtualIDGroup
func (cli *ZSClient) AddRolesToIAM2VirtualIDGroup(ctx context.Context, params param.AddRolesToIAM2VirtualIDGroupParam) (*view.AddRolesToIAM2VirtualIDGroupEventView, error) {
	resp := view.AddRolesToIAM2VirtualIDGroupEventView{}
	if err := cli.Post(ctx, "v1/iam2/projects/groups/{groupUuid}/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckStaticProvisionIp operates on StaticProvisionIp
func (cli *ZSClient) CheckStaticProvisionIp(ctx context.Context, params param.CheckStaticProvisionIpParam) (*view.CheckStaticProvisionIpView, error) {
	resp := view.CheckStaticProvisionIpView{}
	if err := cli.Post(ctx, "v1/baremetal2/bm-instances/static/provision/ip/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeEventSubscriptionState changes EventSubscriptionState
func (cli *ZSClient) ChangeEventSubscriptionState(ctx context.Context, uuid string, params param.ChangeEventSubscriptionStateParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/change/eventSubscription", uuid, "", map[string]interface{}{
		"changeEventSubscriptionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PushLicenseAddOnsUsage operates on PushLicenseAddOnsUsage
func (cli *ZSClient) PushLicenseAddOnsUsage(ctx context.Context, params param.PushLicenseAddOnsUsageParam) (*view.PushLicenseAddOnsUsageEventView, error) {
	resp := view.PushLicenseAddOnsUsageEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/licenses/addons/usage", "", "", map[string]interface{}{
		"pushLicenseAddOnsUsage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachHybridEipToEcs operates on HybridEipToEcs
func (cli *ZSClient) AttachHybridEipToEcs(ctx context.Context, params param.AttachHybridEipToEcsParam) (*view.HybridEipAddressInventoryView, error) {
	resp := view.HybridEipAddressInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/eip/{eipUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsImageFromLocalImage creates EcsImageFromLocalImage
func (cli *ZSClient) CreateEcsImageFromLocalImage(ctx context.Context, params param.CreateEcsImageFromLocalImageParam) (*view.EcsImageInventoryView, error) {
	resp := view.EcsImageInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/image", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddHostRouteToL3Network adds HostRouteToL3Network
func (cli *ZSClient) AddHostRouteToL3Network(ctx context.Context, params param.AddHostRouteToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.Post(ctx, "v1/l3-networks/{l3NetworkUuid}/hostroute", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddInstanceToMonitorGroup adds InstanceToMonitorGroup
func (cli *ZSClient) AddInstanceToMonitorGroup(ctx context.Context, params param.AddInstanceToMonitorGroupParam) (*view.MonitorGroupInstanceInventoryView, error) {
	resp := view.MonitorGroupInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/monitorgroups/{groupUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBareMetal2ProvisionNetworkIpAddressCapacity gets BareMetal2ProvisionNetworkIpAddressCapacity by uuid
func (cli *ZSClient) GetBareMetal2ProvisionNetworkIpAddressCapacity(ctx context.Context) (*view.GetBareMetal2ProvisionNetworkIpAddressCapacityView, error) {
	var resp view.GetBareMetal2ProvisionNetworkIpAddressCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/baremetal2/provision-networks/ip-capacity", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachMdevDeviceToVm operates on MdevDeviceToVm
func (cli *ZSClient) AttachMdevDeviceToVm(ctx context.Context, params param.AttachMdevDeviceToVmParam) (*view.MdevDeviceInventoryView, error) {
	resp := view.MdevDeviceInventoryView{}
	if err := cli.Post(ctx, "v1/mdev-devices/{mdevDeviceUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DecodeStackTemplate operates on DecodeStackTemplate
func (cli *ZSClient) DecodeStackTemplate(ctx context.Context, params param.DecodeStackTemplateParam) (*view.DecodeStackTemplateView, error) {
	resp := view.DecodeStackTemplateView{}
	if err := cli.Post(ctx, "v1/cloudformation/stack/preview/resource", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVirtualRouter updates VirtualRouter
func (cli *ZSClient) UpdateVirtualRouter(ctx context.Context, vmInstanceUuid string, params param.UpdateVirtualRouterParam) (*view.VirtualRouterVmInventoryView, error) {
	resp := view.VirtualRouterVmInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/appliances/virtual-routers", vmInstanceUuid, "", map[string]interface{}{
		"updateVirtualRouter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVSwitchTypes gets VSwitchTypes by uuid
func (cli *ZSClient) GetVSwitchTypes(ctx context.Context) (*view.GetVSwitchTypesView, error) {
	var resp view.GetVSwitchTypesView
	if err := cli.GetWithRespKey(ctx, "v1/l2-networks/vSwitchTypes", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsInstanceFromLocal queries EcsInstanceFromLocal list
func (cli *ZSClient) QueryEcsInstanceFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsInstanceInventoryView, error) {
	var resp []view.EcsInstanceInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/ecs", params, &resp)
}

func (cli *ZSClient) GetEcsInstanceFromLocal(ctx context.Context, uuid string) (*view.EcsInstanceInventoryView, error) {
	var resp view.EcsInstanceInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/ecs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsInstanceFromLocal Pagination
func (cli *ZSClient) PageEcsInstanceFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsInstanceInventoryView, int, error) {
	var ecsInstanceFromLocals []view.EcsInstanceInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/ecs", params, &ecsInstanceFromLocals)
	return ecsInstanceFromLocals, total, err
}

// CreateL2HardwareVxlanNetworkPool creates L2HardwareVxlanNetworkPool
func (cli *ZSClient) CreateL2HardwareVxlanNetworkPool(ctx context.Context, params param.CreateL2HardwareVxlanNetworkPoolParam) (*view.CreateL2HardwareVxlanNetworkPoolEventView, error) {
	resp := view.CreateL2HardwareVxlanNetworkPoolEventView{}
	if err := cli.Post(ctx, "v1/l2-networks/hardware-vxlan-pool", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLdapServerAvailableAttributes gets LdapServerAvailableAttributes by uuid
func (cli *ZSClient) GetLdapServerAvailableAttributes(ctx context.Context, uuid string) (*view.GetLdapServerAvailableAttributesView, error) {
	var resp view.GetLdapServerAvailableAttributesView
	if err := cli.GetWithRespKey(ctx, "v1/ldap/server/attributes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResizeDataVolume operates on DataVolume
func (cli *ZSClient) ResizeDataVolume(ctx context.Context, uuid string, params param.ResizeDataVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/volumes/data/resize", uuid, "", map[string]interface{}{
		"resizeDataVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEipAttachableVmNics gets EipAttachableVmNics by uuid
func (cli *ZSClient) GetEipAttachableVmNics(ctx context.Context, uuid string) (*view.GetEipAttachableVmNicsView, error) {
	var resp view.GetEipAttachableVmNicsView
	if err := cli.GetWithRespKey(ctx, "v1/eips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIpv6RangeByNetworkCidr adds Ipv6RangeByNetworkCidr
func (cli *ZSClient) AddIpv6RangeByNetworkCidr(ctx context.Context, params param.AddIpv6RangeByNetworkCidrParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.Post(ctx, "v1/l3-networks/{l3NetworkUuid}/ipv6-ranges/by-cidr", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchQuery operates on Query
func (cli *ZSClient) BatchQuery(ctx context.Context) (*view.BatchQueryView, error) {
	var resp view.BatchQueryView
	if err := cli.GetWithRespKey(ctx, "v1/batch-queries", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReloadExternalService operates on ReloadExternalService
func (cli *ZSClient) ReloadExternalService(ctx context.Context, params param.ReloadExternalServiceParam) (*view.ReloadExternalServiceEventView, error) {
	resp := view.ReloadExternalServiceEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/external/services", "", "", map[string]interface{}{
		"reloadExternalService": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIAM2VirtualIDsToGroup adds IAM2VirtualIDsToGroup
func (cli *ZSClient) AddIAM2VirtualIDsToGroup(ctx context.Context, params param.AddIAM2VirtualIDsToGroupParam) (*view.AddIAM2VirtualIDToGroupEventView, error) {
	resp := view.AddIAM2VirtualIDToGroupEventView{}
	if err := cli.Post(ctx, "v1/iam2/projects/groups/{groupUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2VirtualIDLdapBinding creates IAM2VirtualIDLdapBinding
func (cli *ZSClient) CreateIAM2VirtualIDLdapBinding(ctx context.Context, params param.CreateIAM2VirtualIDLdapBindingParam) (*view.LdapResourceRefInventoryView, error) {
	resp := view.LdapResourceRefInventoryView{}
	if err := cli.Post(ctx, "v1/iam2/ldap/bindings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmNicSecurityGroup operates on VmNicSecurityGroup
func (cli *ZSClient) SetVmNicSecurityGroup(ctx context.Context, vmNicUuid string, params param.SetVmNicSecurityGroupParam) (*view.VmNicSecurityGroupRefInventoryView, error) {
	resp := view.VmNicSecurityGroupRefInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-groups/nics", vmNicUuid, "", map[string]interface{}{
		"setVmNicSecurityGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryResourcePrice queries ResourcePrice list
func (cli *ZSClient) QueryResourcePrice(ctx context.Context, params *param.QueryParam) ([]view.PriceInventoryView, error) {
	var resp []view.PriceInventoryView
	return resp, cli.List(ctx, "v1/billings/prices", params, &resp)
}

func (cli *ZSClient) GetResourcePrice(ctx context.Context, uuid string) (*view.PriceInventoryView, error) {
	var resp view.PriceInventoryView
	if err := cli.Get(ctx, "v1/billings/prices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageResourcePrice Pagination
func (cli *ZSClient) PageResourcePrice(ctx context.Context, params *param.QueryParam) ([]view.PriceInventoryView, int, error) {
	var resourcePrices []view.PriceInventoryView
	total, err := cli.Page(ctx, "v1/billings/prices", params, &resourcePrices)
	return resourcePrices, total, err
}

// AddIdentityZoneFromRemote adds IdentityZoneFromRemote
func (cli *ZSClient) AddIdentityZoneFromRemote(ctx context.Context, params param.AddIdentityZoneFromRemoteParam) (*view.IdentityZoneInventoryView, error) {
	resp := view.IdentityZoneInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/identity-zone", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeSnapshotSize gets VolumeSnapshotSize by uuid
func (cli *ZSClient) GetVolumeSnapshotSize(ctx context.Context, uuid string, params param.GetVolumeSnapshotSizeParam) (*view.GetVolumeSnapshotSizeEventView, error) {
	resp := view.GetVolumeSnapshotSizeEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/volume-snapshots", uuid, "", map[string]interface{}{
		"getVolumeSnapshotSize": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchSyncVolumeSize operates on SyncVolumeSize
func (cli *ZSClient) BatchSyncVolumeSize(ctx context.Context, params param.BatchSyncVolumeSizeParam) (*view.BatchSyncVolumeSizeView, error) {
	resp := view.BatchSyncVolumeSizeView{}
	if err := cli.Post(ctx, "v1/volumes/batch-sync-volumes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHypervisorTypes gets HypervisorTypes by uuid
func (cli *ZSClient) GetHypervisorTypes(ctx context.Context) (*view.GetHypervisorTypesView, error) {
	var resp view.GetHypervisorTypesView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/hypervisor-types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmAttachableDataVolume gets VmAttachableDataVolume by uuid
func (cli *ZSClient) GetVmAttachableDataVolume(ctx context.Context, uuid string) (*view.GetVmAttachableDataVolumeView, error) {
	var resp view.GetVmAttachableDataVolumeView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmMonitorNumber gets VmMonitorNumber by uuid
func (cli *ZSClient) GetVmMonitorNumber(ctx context.Context, uuid string) (*view.GetVmMonitorNumberView, error) {
	var resp view.GetVmMonitorNumberView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2VirtualIDFromLdapUid creates IAM2VirtualIDFromLdapUid
func (cli *ZSClient) CreateIAM2VirtualIDFromLdapUid(ctx context.Context, params param.CreateIAM2VirtualIDFromLdapUidParam) (*view.LdapResourceRefInventoryView, error) {
	resp := view.LdapResourceRefInventoryView{}
	if err := cli.Post(ctx, "v1/iam2/virtual-id/ldap/uid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidatePriceUserConfig operates on PriceUserConfig
func (cli *ZSClient) ValidatePriceUserConfig(ctx context.Context, params param.ValidatePriceUserConfigParam) (*view.ValidatePriceUserConfigEventView, error) {
	resp := view.ValidatePriceUserConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/accounts/actions", "", "", map[string]interface{}{
		"validatePriceUserConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2GatewayState changes BareMetal2GatewayState
func (cli *ZSClient) ChangeBareMetal2GatewayState(ctx context.Context, uuid string, params param.ChangeBareMetal2GatewayStateParam) (*view.BareMetal2GatewayInventoryView, error) {
	resp := view.BareMetal2GatewayInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/gateways", uuid, "", map[string]interface{}{
		"changeBareMetal2GatewayState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveActionFromEventSubscription removes ActionFromEventSubscription
func (cli *ZSClient) RemoveActionFromEventSubscription(ctx context.Context, subscriptionUuid string, actionUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zwatch/events/subscriptions", subscriptionUuid, fmt.Sprintf("actions/%s", actionUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// CheckKVMHostConfigFile operates on KVMHostConfigFile
func (cli *ZSClient) CheckKVMHostConfigFile(ctx context.Context) (*view.CheckHostConfigFileView, error) {
	resp := view.CheckHostConfigFileView{}
	if err := cli.Post(ctx, "v1/hosts/kvm/from-file/check", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetContainerUsage gets ContainerUsage by uuid
func (cli *ZSClient) GetContainerUsage(ctx context.Context) (*view.GetContainerUsageView, error) {
	var resp view.GetContainerUsageView
	if err := cli.GetWithRespKey(ctx, "v1/container/usage", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSSnmpTestConnection operates on SnmpTestConnection
func (cli *ZSClient) SNSSnmpTestConnection(ctx context.Context, params param.SNSSnmpTestConnectionParam) (*view.SNSSnmpTestConnectionEventView, error) {
	resp := view.SNSSnmpTestConnectionEventView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/snmp/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDataCenterFromRemote gets DataCenterFromRemote by uuid
func (cli *ZSClient) GetDataCenterFromRemote(ctx context.Context) (*view.GetDataCenterFromRemoteView, error) {
	var resp view.GetDataCenterFromRemoteView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/data-center/remote", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryGCJob queries GCJob list
func (cli *ZSClient) QueryGCJob(ctx context.Context, params *param.QueryParam) ([]view.GarbageCollectorInventoryView, error) {
	var resp []view.GarbageCollectorInventoryView
	return resp, cli.List(ctx, "v1/gc-jobs", params, &resp)
}

func (cli *ZSClient) GetGCJob(ctx context.Context, uuid string) (*view.GarbageCollectorInventoryView, error) {
	var resp view.GarbageCollectorInventoryView
	if err := cli.Get(ctx, "v1/gc-jobs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageGCJob Pagination
func (cli *ZSClient) PageGCJob(ctx context.Context, params *param.QueryParam) ([]view.GarbageCollectorInventoryView, int, error) {
	var gCJobs []view.GarbageCollectorInventoryView
	total, err := cli.Page(ctx, "v1/gc-jobs", params, &gCJobs)
	return gCJobs, total, err
}

// CreateHostNetworkServiceType creates HostNetworkServiceType
func (cli *ZSClient) CreateHostNetworkServiceType(ctx context.Context, params param.CreateHostNetworkServiceTypeParam) (*view.HostNetworkLabelInventoryView, error) {
	resp := view.HostNetworkLabelInventoryView{}
	if err := cli.Post(ctx, "v1/hosts/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsImageLocal deletes EcsImageLocal
func (cli *ZSClient) DeleteEcsImageLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/image", uuid, string(deleteMode))
}

// DetachNvmeServerFromCluster operates on NvmeServerFromCluster
func (cli *ZSClient) DetachNvmeServerFromCluster(ctx context.Context, clusterUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/clusters", clusterUuid, fmt.Sprintf("storage-devices/nvme/servers/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetBackupStorageTypes gets BackupStorageTypes by uuid
func (cli *ZSClient) GetBackupStorageTypes(ctx context.Context) (*view.GetBackupStorageTypesView, error) {
	var resp view.GetBackupStorageTypesView
	if err := cli.GetWithRespKey(ctx, "v1/backup-storage/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeQos gets VolumeQos by uuid
func (cli *ZSClient) GetVolumeQos(ctx context.Context, uuid string) (*view.GetVolumeQosView, error) {
	var resp view.GetVolumeQosView
	if err := cli.GetWithRespKey(ctx, "v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRemoteCidrsToIPsecConnection adds RemoteCidrsToIPsecConnection
func (cli *ZSClient) AddRemoteCidrsToIPsecConnection(ctx context.Context, params param.AddRemoteCidrsToIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.Post(ctx, "v1/ipsec/{uuid}/remote-cidrs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOnBaremetalChassis operates on PowerOnBaremetalChassis
func (cli *ZSClient) PowerOnBaremetalChassis(ctx context.Context, chassisUuid string, params param.PowerOnBaremetalChassisParam) (*view.PowerOnBaremetalChassisEventView, error) {
	resp := view.PowerOnBaremetalChassisEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal/chassis", chassisUuid, "", map[string]interface{}{
		"powerOnBaremetalChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RequestLicenseCapacity operates on RequestLicenseCapacity
func (cli *ZSClient) RequestLicenseCapacity(ctx context.Context, params param.RequestLicenseCapacityParam) (*view.LicenseAuthorizedCapacityInventoryView, error) {
	resp := view.LicenseAuthorizedCapacityInventoryView{}
	if err := cli.Post(ctx, "v1/license-server/capacity-application", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeFromVolumeSnapshot creates DataVolumeFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeFromVolumeSnapshot(ctx context.Context, params param.CreateDataVolumeFromVolumeSnapshotParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.Post(ctx, "v1/volumes/data/from/volume-snapshots/{volumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachIsoFromVmInstance operates on IsoFromVmInstance
func (cli *ZSClient) DetachIsoFromVmInstance(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// DetachSecurityGroupFromL3Network operates on SecurityGroupFromL3Network
func (cli *ZSClient) DetachSecurityGroupFromL3Network(ctx context.Context, securityGroupUuid string, l3NetworkUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/security-groups", securityGroupUuid, fmt.Sprintf("l3-networks/%s", l3NetworkUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetVirtualizerInfo gets VirtualizerInfo by uuid
func (cli *ZSClient) GetVirtualizerInfo(ctx context.Context) (*view.GetVirtualizerInfoView, error) {
	var resp view.GetVirtualizerInfoView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/virtualizer-info", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkIpStatistic gets L3NetworkIpStatistic by uuid
func (cli *ZSClient) GetL3NetworkIpStatistic(ctx context.Context, uuid string) (*view.GetL3NetworkIpStatisticView, error) {
	var resp view.GetL3NetworkIpStatisticView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImageCandidatesForVmToChange gets ImageCandidatesForVmToChange by uuid
func (cli *ZSClient) GetImageCandidatesForVmToChange(ctx context.Context, uuid string) (*view.GetImageCandidatesForVmToChangeView, error) {
	var resp view.GetImageCandidatesForVmToChangeView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeImageState changes ImageState
func (cli *ZSClient) ChangeImageState(ctx context.Context, uuid string, params param.ChangeImageStateParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/images", uuid, "", map[string]interface{}{
		"changeImageState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// KvmRunShell operates on KvmRunShell
func (cli *ZSClient) KvmRunShell(ctx context.Context, params param.KvmRunShellParam) (*view.KvmRunShellEventView, error) {
	resp := view.KvmRunShellEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/kvm/actions", "", "", map[string]interface{}{
		"kvmRunShell": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAliyunNasAccessGroupRule creates AliyunNasAccessGroupRule
func (cli *ZSClient) CreateAliyunNasAccessGroupRule(ctx context.Context, params param.CreateAliyunNasAccessGroupRuleParam) (*view.AliyunNasAccessRuleInventoryView, error) {
	resp := view.AliyunNasAccessRuleInventoryView{}
	if err := cli.Post(ctx, "v1/nas/aliyun/rule", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoverBackupFromImageStoreBackupStorage operates on BackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverBackupFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.RecoverBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/volume-backups", uuid, "", map[string]interface{}{
		"recoverBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeTicketFlowCollectionState changes TicketFlowCollectionState
func (cli *ZSClient) ChangeTicketFlowCollectionState(ctx context.Context, uuid string, params param.ChangeTicketFlowCollectionStateParam) (*view.TicketFlowCollectionInventoryView, error) {
	resp := view.TicketFlowCollectionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/tickets/flow-collections", uuid, "", map[string]interface{}{
		"changeTicketFlowCollectionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExpungeDataVolume operates on DataVolume
func (cli *ZSClient) ExpungeDataVolume(ctx context.Context, uuid string) error {
	params := map[string]interface{}{
		"expungeDataVolume": map[string]interface{}{},
	}
	return cli.Put(ctx, "v1/volumes", uuid, params, nil)
}

// AddActionToEventSubscription adds ActionToEventSubscription
func (cli *ZSClient) AddActionToEventSubscription(ctx context.Context, params param.AddActionToEventSubscriptionParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/events/subscriptions/{subscriptionUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVRouterRouterId gets VRouterRouterId by uuid
func (cli *ZSClient) GetVRouterRouterId(ctx context.Context, uuid string) (*view.GetVRouterRouterIdView, error) {
	var resp view.GetVRouterRouterIdView
	if err := cli.GetWithRespKey(ctx, "v1/routerArea", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetZBoxBackupDetails gets ZBoxBackupDetails by uuid
func (cli *ZSClient) GetZBoxBackupDetails(ctx context.Context, uuid string) (*view.GetZBoxBackupDetailsView, error) {
	var resp view.GetZBoxBackupDetailsView
	if err := cli.GetWithRespKey(ctx, "v1/externalbackup/zbox", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetExternalServices gets ExternalServices by uuid
func (cli *ZSClient) GetExternalServices(ctx context.Context) (*view.GetExternalServicesView, error) {
	var resp view.GetExternalServicesView
	if err := cli.GetWithRespKey(ctx, "v1/external/services", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2ProjectRepository gets IAM2ProjectRepository by uuid
func (cli *ZSClient) GetIAM2ProjectRepository(ctx context.Context) (*view.GetIAM2ProjectRepositoryView, error) {
	var resp view.GetIAM2ProjectRepositoryView
	if err := cli.GetWithRespKey(ctx, "v1/iam2/projects/repositories", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateNetworkInterfaces gets CandidateNetworkInterfaces by uuid
func (cli *ZSClient) GetCandidateNetworkInterfaces(ctx context.Context) (*view.GetCandidateNetworkInterfacesView, error) {
	var resp view.GetCandidateNetworkInterfacesView
	if err := cli.GetWithRespKey(ctx, "v1/cluster/hosts-network-interfaces", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAccessControlListServerGroup changes AccessControlListServerGroup
func (cli *ZSClient) ChangeAccessControlListServerGroup(ctx context.Context, aclUuid string, params param.ChangeAccessControlListServerGroupParam) (*view.LoadBalancerListerAclView, error) {
	resp := view.LoadBalancerListerAclView{}
	if err := cli.PutWithRespKey(ctx, "v1/load-balancers/listener/acl", aclUuid, "", map[string]interface{}{
		"changeAccessControlListServerGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVirtualBorderRouterFromRemote operates on VirtualBorderRouterFromRemote
func (cli *ZSClient) SyncVirtualBorderRouterFromRemote(ctx context.Context, dataCenterUuid string, params param.SyncVirtualBorderRouterFromRemoteParam) (*view.VirtualBorderRouterInventoryView, error) {
	resp := view.VirtualBorderRouterInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/border-router", dataCenterUuid, "", map[string]interface{}{
		"syncVirtualBorderRouterFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAtPersonOfAtFeiShuEndpoint updates AtPersonOfAtFeiShuEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtFeiShuEndpoint(ctx context.Context, uuid string, params param.UpdateAtPersonOfAtFeiShuEndpointParam) (*view.SNSFeiShuAtPersonInventoryView, error) {
	resp := view.SNSFeiShuAtPersonInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-endpoints/feishu/at-persons", uuid, "", map[string]interface{}{
		"updateAtPersonOfAtFeiShuEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateL2HardwareVxlanNetwork creates L2HardwareVxlanNetwork
func (cli *ZSClient) CreateL2HardwareVxlanNetwork(ctx context.Context, params param.CreateL2HardwareVxlanNetworkParam) (*view.CreateL2HardwareVxlanNetworkEventView, error) {
	resp := view.CreateL2HardwareVxlanNetworkEventView{}
	if err := cli.Post(ctx, "v1/l2-networks/hardware-vxlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGlobalConfigOptions gets GlobalConfigOptions by uuid
func (cli *ZSClient) GetGlobalConfigOptions(ctx context.Context, category string, name string) (*view.GetGlobalConfigOptionsView, error) {
	var resp view.GetGlobalConfigOptionsView
	err := cli.GetWithSpec(ctx, "v1/global-configurations", category, fmt.Sprintf("%s", name), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateHybridEip creates HybridEip
func (cli *ZSClient) CreateHybridEip(ctx context.Context, params param.CreateHybridEipParam) (*view.HybridEipAddressInventoryView, error) {
	resp := view.HybridEipAddressInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/eip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ApplyMonitorTemplateToMonitorGroup operates on MonitorTemplateToMonitorGroup
func (cli *ZSClient) ApplyMonitorTemplateToMonitorGroup(ctx context.Context, params param.ApplyMonitorTemplateToMonitorGroupParam) (*view.MonitorGroupTemplateRefInventoryView, error) {
	resp := view.MonitorGroupTemplateRefInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/monitortemplates/{templateUuid}/monitorgroups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PutMetricData operates on PutMetricData
func (cli *ZSClient) PutMetricData(ctx context.Context, params param.PutMetricDataParam) (*view.PutMetricDataEventView, error) {
	resp := view.PutMetricDataEventView{}
	if err := cli.Post(ctx, "v1/zwatch/metrics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAttachablePublicL3ForVRouter gets AttachablePublicL3ForVRouter by uuid
func (cli *ZSClient) GetAttachablePublicL3ForVRouter(ctx context.Context, uuid string) (*view.GetAttachablePublicL3ForVRouterView, error) {
	var resp view.GetAttachablePublicL3ForVRouterView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/appliances/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RerunLongJob operates on RerunLongJob
func (cli *ZSClient) RerunLongJob(ctx context.Context, uuid string, params param.RerunLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/longjobs", uuid, "", map[string]interface{}{
		"rerunLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallRuleSet queries FirewallRuleSet list
func (cli *ZSClient) QueryFirewallRuleSet(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleSetInventoryView, error) {
	var resp []view.VpcFirewallRuleSetInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/ruleSets", params, &resp)
}

func (cli *ZSClient) GetFirewallRuleSet(ctx context.Context, uuid string) (*view.VpcFirewallRuleSetInventoryView, error) {
	var resp view.VpcFirewallRuleSetInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls/ruleSets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRuleSet Pagination
func (cli *ZSClient) PageFirewallRuleSet(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleSetInventoryView, int, error) {
	var firewallRuleSets []view.VpcFirewallRuleSetInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/ruleSets", params, &firewallRuleSets)
	return firewallRuleSets, total, err
}

// DeleteExportedImageFromBackupStorage deletes ExportedImageFromBackupStorage
func (cli *ZSClient) DeleteExportedImageFromBackupStorage(ctx context.Context, backupStorageUuid string, imageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/backup-storage", backupStorageUuid, fmt.Sprintf("exported-images/%s", imageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// UpdateClusterOS updates ClusterOS
func (cli *ZSClient) UpdateClusterOS(ctx context.Context, uuid string, params param.UpdateClusterOSParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/clusters", uuid, "", map[string]interface{}{
		"updateClusterOS": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateClusterOSAsync Async
func (cli *ZSClient) UpdateClusterOSAsync(ctx context.Context, params param.UpdateClusterOSParam) (string, error) {

	resource := "v1/clusters/{uuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetVmUsbRedirect gets VmUsbRedirect by uuid
func (cli *ZSClient) GetVmUsbRedirect(ctx context.Context, uuid string) (*view.GetVmUsbRedirectView, error) {
	var resp view.GetVmUsbRedirectView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateImageGroupFromSnapshot creates ImageGroupFromSnapshot
func (cli *ZSClient) CreateImageGroupFromSnapshot(ctx context.Context, params param.CreateImageGroupFromSnapshotParam) (*view.ImageGroupInventoryView, error) {
	resp := view.ImageGroupInventoryView{}
	if err := cli.Post(ctx, "v1/imagegroup/from/snapshot/{rootVolumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOssBucketFileFromRemote gets OssBucketFileFromRemote by uuid
func (cli *ZSClient) GetOssBucketFileFromRemote(ctx context.Context) (*view.GetOssBucketFileFromRemoteView, error) {
	var resp view.GetOssBucketFileFromRemoteView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/oss/file/remote", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVipToVpcSharedQos operates on VipToVpcSharedQos
func (cli *ZSClient) AttachVipToVpcSharedQos(ctx context.Context, params param.AttachVipToVpcSharedQosParam) (*view.AttachVipToVpcSharedQosEventView, error) {
	resp := view.AttachVipToVpcSharedQosEventView{}
	if err := cli.Post(ctx, "v1/vips/sharedqos/{sharedQosUuid}/vips", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEventData gets EventData by uuid
func (cli *ZSClient) GetEventData(ctx context.Context) (*view.GetEventDataView, error) {
	var resp view.GetEventDataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/events", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIpAvailability operates on IpAvailability
func (cli *ZSClient) CheckIpAvailability(ctx context.Context, l3NetworkUuid string, ip string) (*view.CheckIpAvailabilityView, error) {
	var resp view.CheckIpAvailabilityView
	err := cli.GetWithSpec(ctx, "v1/l3-networks", l3NetworkUuid, fmt.Sprintf("ip/%s/availability", ip), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVmNicFromLoadBalancer removes VmNicFromLoadBalancer
func (cli *ZSClient) RemoveVmNicFromLoadBalancer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/listeners", uuid, string(deleteMode))
}

// RemoveRolesFromIAM2VirtualID removes RolesFromIAM2VirtualID
func (cli *ZSClient) RemoveRolesFromIAM2VirtualID(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects/virtual-ids", uuid, string(deleteMode))
}

// CalculateResourceSpending operates on ResourceSpending
func (cli *ZSClient) CalculateResourceSpending(ctx context.Context, params param.CalculateResourceSpendingParam) (*view.CalculateResourceSpendingView, error) {
	resp := view.CalculateResourceSpendingView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/resources/actions", "", "", map[string]interface{}{
		"calculateResourceSpending": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAlarmRecord queries AlarmRecord list
func (cli *ZSClient) QueryAlarmRecord(ctx context.Context, params *param.QueryParam) ([]view.AlarmRecordsInventoryView, error) {
	var resp []view.AlarmRecordsInventoryView
	return resp, cli.List(ctx, "v1/zwatch/alarm-records", params, &resp)
}

func (cli *ZSClient) GetAlarmRecord(ctx context.Context, uuid string) (*view.AlarmRecordsInventoryView, error) {
	var resp view.AlarmRecordsInventoryView
	if err := cli.Get(ctx, "v1/zwatch/alarm-records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAlarmRecord Pagination
func (cli *ZSClient) PageAlarmRecord(ctx context.Context, params *param.QueryParam) ([]view.AlarmRecordsInventoryView, int, error) {
	var alarmRecords []view.AlarmRecordsInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/alarm-records", params, &alarmRecords)
	return alarmRecords, total, err
}

// DetachBackupStorageFromZone operates on BackupStorageFromZone
func (cli *ZSClient) DetachBackupStorageFromZone(ctx context.Context, zoneUuid string, backupStorageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zones", zoneUuid, fmt.Sprintf("backup-storage/%s", backupStorageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// UpdateCCSCertificateUserState updates CCSCertificateUserState
func (cli *ZSClient) UpdateCCSCertificateUserState(ctx context.Context, params param.UpdateCCSCertificateUserStateParam) (*view.CCSCertificateInventoryView, error) {
	resp := view.CCSCertificateInventoryView{}
	if err := cli.Post(ctx, "v1/crypto/ccs-certificate/update-state/{userUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerResetBaremetalChassis operates on PowerResetBaremetalChassis
func (cli *ZSClient) PowerResetBaremetalChassis(ctx context.Context, chassisUuid string, params param.PowerResetBaremetalChassisParam) (*view.PowerResetBaremetalChassisEventView, error) {
	resp := view.PowerResetBaremetalChassisEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal/chassis", chassisUuid, "", map[string]interface{}{
		"powerResetBaremetalChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpTrashOnPrimaryStorage operates on UpTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpTrashOnPrimaryStorage(ctx context.Context, uuid string, params param.CleanUpTrashOnPrimaryStorageParam) (*view.CleanUpTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpTrashOnPrimaryStorageEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage", uuid, "", map[string]interface{}{
		"cleanUpTrashOnPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDisasterImageStoreBackupStorage adds DisasterImageStoreBackupStorage
func (cli *ZSClient) AddDisasterImageStoreBackupStorage(ctx context.Context, params param.AddDisasterImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.Post(ctx, "v1/backup-storage/image-store/disaster", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVolumesSnapshot creates VolumesSnapshot
func (cli *ZSClient) CreateVolumesSnapshot(ctx context.Context, params param.CreateVolumesSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	resp := view.VolumeSnapshotInventoryView{}
	if err := cli.Post(ctx, "v1/volumes/volume-snapshots", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmSchedulingRulesExecuteState gets VmSchedulingRulesExecuteState by uuid
func (cli *ZSClient) GetVmSchedulingRulesExecuteState(ctx context.Context, params param.GetVmSchedulingRulesExecuteStateParam) (*view.GetVmSchedulingRulesExecuteStateView, error) {
	resp := view.GetVmSchedulingRulesExecuteStateView{}
	if err := cli.Post(ctx, "v1/get/vmSchedulingRules/conflict/state", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIpAddressCapacity gets IpAddressCapacity by uuid
func (cli *ZSClient) GetIpAddressCapacity(ctx context.Context) (*view.GetIpAddressCapacityView, error) {
	var resp view.GetIpAddressCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/ip-capacity", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIAM2ProjectContainerCluster operates on IAM2ProjectContainerCluster
func (cli *ZSClient) SetIAM2ProjectContainerCluster(ctx context.Context, uuid string, params param.SetIAM2ProjectContainerClusterParam) (*view.SetIAM2ProjectContainerClusterEventView, error) {
	resp := view.SetIAM2ProjectContainerClusterEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/projects", uuid, "", map[string]interface{}{
		"setIAM2ProjectContainerCluster": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployAppDevelopmentService operates on DeployAppDevelopmentService
func (cli *ZSClient) DeployAppDevelopmentService(ctx context.Context) (*view.ModelServiceInstanceGroupInventoryView, error) {
	resp := view.ModelServiceInstanceGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/ai/model-services/app/", "", "", map[string]interface{}{
		"deployAppDevelopmentService": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshPluginDrivers operates on PluginDrivers
func (cli *ZSClient) RefreshPluginDrivers(ctx context.Context, params param.RefreshPluginDriversParam) (*view.RefreshPluginDriversEventView, error) {
	resp := view.RefreshPluginDriversEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/external/plugins", "", "", map[string]interface{}{
		"refreshPluginDrivers": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PauseVmInstance operates on PauseVmInstance
func (cli *ZSClient) PauseVmInstance(ctx context.Context, uuid string, params param.PauseVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"pauseVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachUserDefinedXmlHookScriptFromVm operates on UserDefinedXmlHookScriptFromVm
func (cli *ZSClient) DetachUserDefinedXmlHookScriptFromVm(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/xmlhook/vm-instances", uuid, string(deleteMode))
}

// GetSignatureServerEncryptPublicKey gets SignatureServerEncryptPublicKey by uuid
func (cli *ZSClient) GetSignatureServerEncryptPublicKey(ctx context.Context) (*view.GetSignatureServerEncryptPublicKeyView, error) {
	var resp view.GetSignatureServerEncryptPublicKeyView
	if err := cli.GetWithRespKey(ctx, "v1/secret-resource-pool-token/signature-server-encrypt-public-key", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAliyunKeySecret adds AliyunKeySecret
func (cli *ZSClient) AddAliyunKeySecret(ctx context.Context, params param.AddAliyunKeySecretParam) (*view.HybridAccountInventoryView, error) {
	resp := view.HybridAccountInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/key", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddBackupStoragesToReplicationGroup adds BackupStoragesToReplicationGroup
func (cli *ZSClient) AddBackupStoragesToReplicationGroup(ctx context.Context, params param.AddBackupStoragesToReplicationGroupParam) (*view.ImageReplicationGroupBackupStorageRefInventoryView, error) {
	resp := view.ImageReplicationGroupBackupStorageRefInventoryView{}
	if err := cli.Post(ctx, "v1/image-replication-groups/{replicationGroupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDataCenterFromRemote adds DataCenterFromRemote
func (cli *ZSClient) AddDataCenterFromRemote(ctx context.Context, params param.AddDataCenterFromRemoteParam) (*view.DataCenterInventoryView, error) {
	resp := view.DataCenterInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/data-center", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFirewallRuleSet deletes FirewallRuleSet
func (cli *ZSClient) DeleteFirewallRuleSet(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpcfirewalls/ruleSets", uuid, string(deleteMode))
}

// BatchAddBareMetal2IpmiChassis operates on AddBareMetal2IpmiChassis
func (cli *ZSClient) BatchAddBareMetal2IpmiChassis(ctx context.Context) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Post(ctx, "v1/baremetal2/chassis/ipmi/from-file", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LocalStorageMigrateVolume operates on LocalStorageMigrateVolume
func (cli *ZSClient) LocalStorageMigrateVolume(ctx context.Context, volumeUuid string, params param.LocalStorageMigrateVolumeParam) (*view.LocalStorageResourceRefInventoryView, error) {
	resp := view.LocalStorageResourceRefInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage/local-storage/volumes", volumeUuid, "", map[string]interface{}{
		"localStorageMigrateVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachNicToBonding operates on NicToBonding
func (cli *ZSClient) AttachNicToBonding(ctx context.Context, uuid string, params param.AttachNicToBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/bondings", uuid, "", map[string]interface{}{
		"attachNicToBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetOrganizationOperation operates on OrganizationOperation
func (cli *ZSClient) SetOrganizationOperation(ctx context.Context, uuid string, params param.SetOrganizationOperationParam) (*view.SetOrganizationOperationEventView, error) {
	resp := view.SetOrganizationOperationEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/organizations", uuid, "", map[string]interface{}{
		"setOrganizationOperation": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolume creates DataVolumeTemplateFromVolume
func (cli *ZSClient) CreateDataVolumeTemplateFromVolume(ctx context.Context, params param.CreateDataVolumeTemplateFromVolumeParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(ctx, "v1/images/data-volume-templates/from/volumes/{volumeUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolumeAsync Async
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeAsync(ctx context.Context, params param.CreateDataVolumeTemplateFromVolumeParam) (string, error) {

	resource := "v1/images/data-volume-templates/from/volumes/{volumeUuid}"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// RemoveIAM2VirtualIDsFromOrganization removes IAM2VirtualIDsFromOrganization
func (cli *ZSClient) RemoveIAM2VirtualIDsFromOrganization(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/organizations", uuid, string(deleteMode))
}

// ExportDatabaseBackupFromBackupStorage operates on DatabaseBackupFromBackupStorage
func (cli *ZSClient) ExportDatabaseBackupFromBackupStorage(ctx context.Context, databaseBackupUuid string, backupStorageUuid string, params param.ExportDatabaseBackupFromBackupStorageParam) (*view.ExportDatabaseBackupFromBackupStorageEventView, error) {
	resp := view.ExportDatabaseBackupFromBackupStorageEventView{}
	err := cli.PutWithSpec(ctx, "v1/database-backups", databaseBackupUuid, fmt.Sprintf("backup-storage/%s/actions", backupStorageUuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachIAM2ProjectToIAM2Organization operates on IAM2ProjectToIAM2Organization
func (cli *ZSClient) AttachIAM2ProjectToIAM2Organization(ctx context.Context, params param.AttachIAM2ProjectToIAM2OrganizationParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.Post(ctx, "v1/iam2/projects/{projectUuid}/iam2/organizations/{organizationUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEmailMonitorTriggerAction creates EmailMonitorTrigger
func (cli *ZSClient) CreateEmailMonitorTriggerAction(ctx context.Context, params param.CreateEmailMonitorTriggerActionParam) (*view.MonitorTriggerActionInventoryView, error) {
	resp := view.MonitorTriggerActionInventoryView{}
	if err := cli.Post(ctx, "v1/monitoring/trigger-actions/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVpcVRouterDistributedRoutingEnabled operates on VpcVRouterDistributedRoutingEnabled
func (cli *ZSClient) SetVpcVRouterDistributedRoutingEnabled(ctx context.Context, params param.SetVpcVRouterDistributedRoutingEnabledParam) (*view.SetVpcVRouterDistributedRoutingEnabledEventView, error) {
	resp := view.SetVpcVRouterDistributedRoutingEnabledEventView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/distributed-routing", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOnBareMetal2Chassis operates on PowerOnBareMetal2Chassis
func (cli *ZSClient) PowerOnBareMetal2Chassis(ctx context.Context, uuid string, params param.PowerOnBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"powerOnBareMetal2Chassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLocalRaidPhysicalDriveSmart gets LocalRaidPhysicalDriveSmart by uuid
func (cli *ZSClient) GetLocalRaidPhysicalDriveSmart(ctx context.Context, uuid string) (*view.GetLocalRaidPhysicalDriveSmartView, error) {
	var resp view.GetLocalRaidPhysicalDriveSmartView
	if err := cli.GetWithRespKey(ctx, "v1/storage-devices/local-raid/physical-drives", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHybridKeySecret updates HybridKeySecret
func (cli *ZSClient) UpdateHybridKeySecret(ctx context.Context, params param.UpdateHybridKeySecretParam) (*view.HybridAccountInventoryView, error) {
	resp := view.HybridAccountInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/hybrid/{uuid}/key", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PullHuaweiIMasterController operates on PullHuaweiIMasterController
func (cli *ZSClient) PullHuaweiIMasterController(ctx context.Context, uuid string, params param.PullHuaweiIMasterControllerParam) (*view.HuaweiIMasterSdnControllerInventoryView, error) {
	resp := view.HuaweiIMasterSdnControllerInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sdn-controller/huawei-imaster", uuid, "", map[string]interface{}{
		"pullHuaweiIMasterController": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveRolesFromIAM2VirtualIDGroup removes RolesFromIAM2VirtualIDGroup
func (cli *ZSClient) RemoveRolesFromIAM2VirtualIDGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects/groups", uuid, string(deleteMode))
}

// AckAlarmData operates on AlarmData
func (cli *ZSClient) AckAlarmData(ctx context.Context, params param.AckAlarmDataParam) (*view.AlertDataAckInventoryView, error) {
	resp := view.AlertDataAckInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/alarm-histories/acknowledgments", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveDnsFromL3Network removes DnsFromL3Network
func (cli *ZSClient) RemoveDnsFromL3Network(ctx context.Context, l3NetworkUuid string, dns string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/l3-networks", l3NetworkUuid, fmt.Sprintf("dns/%s", dns), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ChangeIAM2OrganizationParent changes IAM2OrganizationParent
func (cli *ZSClient) ChangeIAM2OrganizationParent(ctx context.Context, parentUuid string, params param.ChangeIAM2OrganizationParentParam) (*view.ChangeIAM2OrganizationParentEventView, error) {
	resp := view.ChangeIAM2OrganizationParentEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/organizations", parentUuid, "", map[string]interface{}{
		"changeIAM2OrganizationParent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSWeComTestConnection operates on WeComTestConnection
func (cli *ZSClient) SNSWeComTestConnection(ctx context.Context, params param.SNSWeComTestConnectionParam) (*view.SNSWeComTestConnectionEventView, error) {
	resp := view.SNSWeComTestConnectionEventView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/we-com/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProvisionVirtualRouterConfig operates on ProvisionVirtualRouterConfig
func (cli *ZSClient) ProvisionVirtualRouterConfig(ctx context.Context, vmInstanceUuid string, params param.ProvisionVirtualRouterConfigParam) (*view.ApplianceVmInventoryView, error) {
	resp := view.ApplianceVmInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/appliances/virtual-routers", vmInstanceUuid, "", map[string]interface{}{
		"provisionVirtualRouterConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmQga operates on VmQga
func (cli *ZSClient) SetVmQga(ctx context.Context, uuid string, params param.SetVmQgaParam) (*view.SetVmQgaEventView, error) {
	resp := view.SetVmQgaEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmQga": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidatePassword operates on Password
func (cli *ZSClient) ValidatePassword(ctx context.Context, params param.ValidatePasswordParam) (*view.ValidatePasswordView, error) {
	resp := view.ValidatePasswordView{}
	if err := cli.PutWithRespKey(ctx, "v1/password/verify", "", "", map[string]interface{}{
		"validatePassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetChronyServers gets ChronyServers by uuid
func (cli *ZSClient) GetChronyServers(ctx context.Context) (*view.GetChronyServersView, error) {
	var resp view.GetChronyServersView
	if err := cli.GetWithRespKey(ctx, "v1/zops/chrony/servers", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL3NetworkToVmNic operates on L3NetworkToVmNic
func (cli *ZSClient) AttachL3NetworkToVmNic(ctx context.Context, params param.AttachL3NetworkToVmNicParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.Post(ctx, "v1/nics/{vmNicUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSecurityMachineState changes SecurityMachineState
func (cli *ZSClient) ChangeSecurityMachineState(ctx context.Context, uuid string, params param.ChangeSecurityMachineStateParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-machines", uuid, "", map[string]interface{}{
		"changeSecurityMachineState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmQxlMemory operates on VmQxlMemory
func (cli *ZSClient) SetVmQxlMemory(ctx context.Context, uuid string, params param.SetVmQxlMemoryParam) (*view.SetVmQxlMemoryEventView, error) {
	resp := view.SetVmQxlMemoryEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmQxlMemory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddLocalPrimaryStorage adds LocalPrimaryStorage
func (cli *ZSClient) AddLocalPrimaryStorage(ctx context.Context) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/local-storage", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeFormat gets VolumeFormat by uuid
func (cli *ZSClient) GetVolumeFormat(ctx context.Context) (*view.GetVolumeFormatView, error) {
	var resp view.GetVolumeFormatView
	if err := cli.GetWithRespKey(ctx, "v1/volumes/formats", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAtPersonOfAtDingTalkEndpoint updates AtPersonOfAtDingTalkEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtDingTalkEndpoint(ctx context.Context, uuid string, params param.UpdateAtPersonOfAtDingTalkEndpointParam) (*view.SNSDingTalkAtPersonInventoryView, error) {
	resp := view.SNSDingTalkAtPersonInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-endpoints/ding-talk/at-persons", uuid, "", map[string]interface{}{
		"updateAtPersonOfAtDingTalkEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAliyunMountTarget updates AliyunMountTarget
func (cli *ZSClient) UpdateAliyunMountTarget(ctx context.Context, params param.UpdateAliyunMountTargetParam) (*view.NasMountTargetInventoryView, error) {
	resp := view.NasMountTargetInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/nas/aliyun/mount", "", "", map[string]interface{}{
		"updateAliyunMountTarget": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceAccount gets ResourceAccount by uuid
func (cli *ZSClient) GetResourceAccount(ctx context.Context) (*view.GetResourceAccountView, error) {
	var resp view.GetResourceAccountView
	if err := cli.GetWithRespKey(ctx, "v1/resources/accounts", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSecretResourcePoolState changes SecretResourcePoolState
func (cli *ZSClient) ChangeSecretResourcePoolState(ctx context.Context, uuid string, params param.ChangeSecretResourcePoolStateParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/secret-resource-pools", uuid, "", map[string]interface{}{
		"changeSecretResourcePoolState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSimulatorBackupStorage adds SimulatorBackupStorage
func (cli *ZSClient) AddSimulatorBackupStorage(ctx context.Context, params param.AddSimulatorBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Post(ctx, "v1/backup-storage/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BindModelToService operates on BindModelToService
func (cli *ZSClient) BindModelToService(ctx context.Context, params param.BindModelToServiceParam) (*view.ModelServiceInventoryView, error) {
	resp := view.ModelServiceInventoryView{}
	if err := cli.Post(ctx, "v1/ai/models/{modelUuid}/model-services/{modelServiceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateAffinityGroupForCreatingVm gets CandidateAffinityGroupForCreatingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForCreatingVm(ctx context.Context) (*view.GetCandidateAffinityGroupForCreatingVmView, error) {
	var resp view.GetCandidateAffinityGroupForCreatingVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/candidate-affinityGroup", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckNetworkReachable operates on NetworkReachable
func (cli *ZSClient) CheckNetworkReachable(ctx context.Context) (*view.CheckNetworkReachableView, error) {
	var resp view.CheckNetworkReachableView
	if err := cli.GetWithRespKey(ctx, "v1/zops/check/network", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetFlowMeterRouterId operates on FlowMeterRouterId
func (cli *ZSClient) SetFlowMeterRouterId(ctx context.Context, params param.SetFlowMeterRouterIdParam) (*view.SetFlowMeterRouterIdEventView, error) {
	resp := view.SetFlowMeterRouterIdEventView{}
	if err := cli.Post(ctx, "v1/flowmeters/{vRouterUuid}/routerid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddStorageProtocol adds StorageProtocol
func (cli *ZSClient) AddStorageProtocol(ctx context.Context, params param.AddStorageProtocolParam) (*view.AddStorageProtocolEventView, error) {
	resp := view.AddStorageProtocolEventView{}
	if err := cli.Post(ctx, "v1/primary-storage/protocol", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployModelService operates on DeployModelService
func (cli *ZSClient) DeployModelService(ctx context.Context, uuid string, params param.DeployModelServiceParam) (*view.ModelServiceInstanceGroupInventoryView, error) {
	resp := view.ModelServiceInstanceGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/ai/model-services", uuid, "", map[string]interface{}{
		"deployModelService": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMonitorItem gets MonitorItem by uuid
func (cli *ZSClient) GetMonitorItem(ctx context.Context) (*view.GetMonitorItemView, error) {
	var resp view.GetMonitorItemView
	if err := cli.GetWithRespKey(ctx, "v1/monitoring/items", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseRecords gets LicenseRecords by uuid
func (cli *ZSClient) GetLicenseRecords(ctx context.Context) (*view.GetLicenseRecordsView, error) {
	var resp view.GetLicenseRecordsView
	if err := cli.GetWithRespKey(ctx, "v1/licenses/records", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnregisterLicenseRequestedApplication operates on LicenseRequestedApplication
func (cli *ZSClient) UnregisterLicenseRequestedApplication(ctx context.Context, params param.UnregisterLicenseRequestedApplicationParam) (*view.UnregisterLicenseRequestedApplicationEventView, error) {
	resp := view.UnregisterLicenseRequestedApplicationEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/license/unregister-applications", "", "", map[string]interface{}{
		"unregisterLicenseRequestedApplication": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachSecurityGroupToL3Network operates on SecurityGroupToL3Network
func (cli *ZSClient) AttachSecurityGroupToL3Network(ctx context.Context, params param.AttachSecurityGroupToL3NetworkParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.Post(ctx, "v1/security-groups/{securityGroupUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmNicDriver updates VmNicDriver
func (cli *ZSClient) UpdateVmNicDriver(ctx context.Context, vmInstanceUuid string, params param.UpdateVmNicDriverParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"updateVmNicDriver": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIpOnHostNetworkInterface operates on IpOnHostNetworkInterface
func (cli *ZSClient) SetIpOnHostNetworkInterface(ctx context.Context, params param.SetIpOnHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	resp := view.HostNetworkInterfaceInventoryView{}
	if err := cli.Post(ctx, "v1/hosts/nics/{interfaceUuid}/ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachNicFromBonding operates on NicFromBonding
func (cli *ZSClient) DetachNicFromBonding(ctx context.Context, uuid string, params param.DetachNicFromBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/bondings", uuid, "", map[string]interface{}{
		"detachNicFromBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeMonitorTriggerActionState changes MonitorTriggerActionState
func (cli *ZSClient) ChangeMonitorTriggerActionState(ctx context.Context, uuid string, params param.ChangeMonitorTriggerActionStateParam) (*view.MonitorTriggerActionInventoryView, error) {
	resp := view.MonitorTriggerActionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/monitoring/trigger-actions", uuid, "", map[string]interface{}{
		"changeMonitorTriggerActionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MigrateVm operates on Vm
func (cli *ZSClient) MigrateVm(ctx context.Context, vmInstanceUuid string, params param.MigrateVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"migrateVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MigrateVmAsync Async
func (cli *ZSClient) MigrateVmAsync(ctx context.Context, params param.MigrateVmParam) (string, error) {

	resource := "v1/vm-instances/{vmInstanceUuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// ChangeVmPassword changes VmPassword
func (cli *ZSClient) ChangeVmPassword(ctx context.Context, uuid string, params param.ChangeVmPasswordParam) (*view.ChangeVmPasswordEventView, error) {
	resp := view.ChangeVmPasswordEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"changeVmPassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVmInstance operates on FlattenVmInstance
func (cli *ZSClient) FlattenVmInstance(ctx context.Context, uuid string, params param.FlattenVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"flattenVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVmInstanceAsync Async
func (cli *ZSClient) FlattenVmInstanceAsync(ctx context.Context, params param.FlattenVmInstanceParam) (string, error) {

	resource := "v1/vm-instances/{uuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// DeleteAllEcsInstancesFromDataCenter deletes AllEcsInstancesFromDataCenter
func (cli *ZSClient) DeleteAllEcsInstancesFromDataCenter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/dc-ecs", uuid, string(deleteMode))
}

// GetVpcMulticastRoute gets VpcMulticastRoute by uuid
func (cli *ZSClient) GetVpcMulticastRoute(ctx context.Context, uuid string) (*view.GetVpcMulticastRouteView, error) {
	var resp view.GetVpcMulticastRouteView
	if err := cli.GetWithRespKey(ctx, "v1/multicast/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmUserDefinedXmlHookScript deletes VmUserDefinedXmlHookScript
func (cli *ZSClient) DeleteVmUserDefinedXmlHookScript(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// SyncZBoxCapacity operates on ZBoxCapacity
func (cli *ZSClient) SyncZBoxCapacity(ctx context.Context, uuid string, params param.SyncZBoxCapacityParam) (*view.ZBoxInventoryView, error) {
	resp := view.ZBoxInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zbox", uuid, "", map[string]interface{}{
		"syncZBoxCapacity": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AckEventData operates on EventData
func (cli *ZSClient) AckEventData(ctx context.Context, params param.AckEventDataParam) (*view.AlertDataAckInventoryView, error) {
	resp := view.AlertDataAckInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/event-histories/acknowledgments", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckResourcePermission operates on ResourcePermission
func (cli *ZSClient) CheckResourcePermission(ctx context.Context) (*view.CheckResourcePermissionView, error) {
	var resp view.CheckResourcePermissionView
	if err := cli.GetWithRespKey(ctx, "v1/accounts/resource/api-permissions", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateMiniHosts gets CandidateMiniHosts by uuid
func (cli *ZSClient) GetCandidateMiniHosts(ctx context.Context) (*view.GetCandidateMiniHostsView, error) {
	var resp view.GetCandidateMiniHostsView
	if err := cli.GetWithRespKey(ctx, "v1/mini-clusters/candidate-hosts", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDatasets deletes Datasets
func (cli *ZSClient) DeleteDatasets(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ai/datasets", uuid, string(deleteMode))
}

// RevokeResourceSharing operates on RevokeResourceSharing
func (cli *ZSClient) RevokeResourceSharing(ctx context.Context, params param.RevokeResourceSharingParam) (*view.RevokeResourceSharingEventView, error) {
	resp := view.RevokeResourceSharingEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/accounts/resources/actions", "", "", map[string]interface{}{
		"revokeResourceSharing": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteModelServices deletes ModelServices
func (cli *ZSClient) DeleteModelServices(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ai/model-services/", uuid, string(deleteMode))
}

// ChangeL3NetworkState changes L3NetworkState
func (cli *ZSClient) ChangeL3NetworkState(ctx context.Context, uuid string, params param.ChangeL3NetworkStateParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/l3-networks", uuid, "", map[string]interface{}{
		"changeL3NetworkState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostNUMATopology gets HostNUMATopology by uuid
func (cli *ZSClient) GetHostNUMATopology(ctx context.Context, params param.GetHostNUMATopologyParam) (*view.GetHostNUMATopologyEventView, error) {
	resp := view.GetHostNUMATopologyEventView{}
	if err := cli.Post(ctx, "v1/hosts/{uuid}/numa", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateL2VirtualSwitch creates L2VirtualSwitch
func (cli *ZSClient) CreateL2VirtualSwitch(ctx context.Context, params param.CreateL2VirtualSwitchParam) (*view.CreateL2VirtualSwitchEventView, error) {
	resp := view.CreateL2VirtualSwitchEventView{}
	if err := cli.Post(ctx, "v1/l2-networks/virtual-switch", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmNicToLoadBalancer adds VmNicToLoadBalancer
func (cli *ZSClient) AddVmNicToLoadBalancer(ctx context.Context, params param.AddVmNicToLoadBalancerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.Post(ctx, "v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBuildApp updates BuildApp
func (cli *ZSClient) UpdateBuildApp(ctx context.Context, uuid string, params param.UpdateBuildAppParam) (*view.BuildApplicationInventoryView, error) {
	resp := view.BuildApplicationInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/appcenter/buildapp", uuid, "", map[string]interface{}{
		"updateBuildApp": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetClusterDRSStatus gets ClusterDRSStatus by uuid
func (cli *ZSClient) GetClusterDRSStatus(ctx context.Context) (*view.GetClusterDRSStatusView, error) {
	var resp view.GetClusterDRSStatusView
	if err := cli.GetWithRespKey(ctx, "v1/clusters/drs/status", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAliyunNasPrimaryStorage adds AliyunNasPrimaryStorage
func (cli *ZSClient) AddAliyunNasPrimaryStorage(ctx context.Context, params param.AddAliyunNasPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/aliyun/nas", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmNuma gets VmNuma by uuid
func (cli *ZSClient) GetVmNuma(ctx context.Context, uuid string) (*view.GetVmNumaView, error) {
	var resp view.GetVmNumaView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeZoneState changes ZoneState
func (cli *ZSClient) ChangeZoneState(ctx context.Context, uuid string, params param.ChangeZoneStateParam) (*view.ZoneInventoryView, error) {
	resp := view.ZoneInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zones", uuid, "", map[string]interface{}{
		"changeZoneState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachAppBuildSystemToZone operates on AppBuildSystemToZone
func (cli *ZSClient) AttachAppBuildSystemToZone(ctx context.Context, params param.AttachAppBuildSystemToZoneParam) (*view.AppBuildSystemZoneRefInventoryView, error) {
	resp := view.AppBuildSystemZoneRefInventoryView{}
	if err := cli.Post(ctx, "v1/zones/{zoneUuid}/buildsystem/{buildSystemUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolume creates DataVolume
func (cli *ZSClient) CreateDataVolume(ctx context.Context, params param.CreateDataVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.Post(ctx, "v1/volumes/data", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePluginDrivers deletes PluginDrivers
func (cli *ZSClient) DeletePluginDrivers(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/external/plugins", uuid, string(deleteMode))
}

// BatchCreateBaremetalChassis operates on CreateBaremetalChassis
func (cli *ZSClient) BatchCreateBaremetalChassis(ctx context.Context, params param.BatchCreateBaremetalChassisParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Post(ctx, "v1/baremetal/chassis/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchCreateBaremetalChassisAsync Async
func (cli *ZSClient) BatchCreateBaremetalChassisAsync(ctx context.Context, params param.BatchCreateBaremetalChassisParam) (string, error) {

	resource := "v1/baremetal/chassis/from-file"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// AddSchedulerJobToSchedulerTrigger adds SchedulerJobToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobToSchedulerTrigger(ctx context.Context, params param.AddSchedulerJobToSchedulerTriggerParam) (*view.SchedulerJobSchedulerTriggerInventoryView, error) {
	resp := view.SchedulerJobSchedulerTriggerInventoryView{}
	if err := cli.Post(ctx, "v1/scheduler/jobs/{schedulerJobUuid}/scheduler/triggers/{schedulerTriggerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPolicyFromRole operates on PolicyFromRole
func (cli *ZSClient) DetachPolicyFromRole(ctx context.Context, policyUuid string, roleUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/identities/policies", policyUuid, fmt.Sprintf("roles/%s", roleUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// RestartModelServiceGroups operates on RestartModelServiceGroups
func (cli *ZSClient) RestartModelServiceGroups(ctx context.Context, params param.RestartModelServiceGroupsParam) (*view.RestartModelServiceGroupsEventView, error) {
	resp := view.RestartModelServiceGroupsEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/model-service-instance-groups", "", "", map[string]interface{}{
		"restartModelServiceGroups": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLoadBalancerOwner gets LoadBalancerOwner by uuid
func (cli *ZSClient) GetLoadBalancerOwner(ctx context.Context, uuid string) (*view.GetLoadBalancerOwnerView, error) {
	var resp view.GetLoadBalancerOwnerView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNicQos gets NicQos by uuid
func (cli *ZSClient) GetNicQos(ctx context.Context, uuid string) (*view.GetNicQosView, error) {
	var resp view.GetNicQosView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmNicNetwork changes VmNicNetwork
func (cli *ZSClient) ChangeVmNicNetwork(ctx context.Context, params param.ChangeVmNicNetworkParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/nics/{vmNicUuid}/l3-networks/{destL3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveLabelFromAlarm removes LabelFromAlarm
func (cli *ZSClient) RemoveLabelFromAlarm(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/alarms/labels", uuid, string(deleteMode))
}

// CreateBareMetal2IpmiChassisHardwareInfo creates BareMetal2IpmiChassisHardwareInfo
func (cli *ZSClient) CreateBareMetal2IpmiChassisHardwareInfo(ctx context.Context, params param.CreateBareMetal2IpmiChassisHardwareInfoParam) (*view.CreateBareMetal2ChassisHardwareView, error) {
	resp := view.CreateBareMetal2ChassisHardwareView{}
	if err := cli.Post(ctx, "v1/baremetal2/chassis/ipmi/hardwareinfos", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIAM2VirtualIDLdapBinding deletes IAM2VirtualIDLdapBinding
func (cli *ZSClient) DeleteIAM2VirtualIDLdapBinding(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/ldap/bindings", uuid, string(deleteMode))
}

// UpdateVmPriority updates VmPriority
func (cli *ZSClient) UpdateVmPriority(ctx context.Context, uuid string, params param.UpdateVmPriorityParam) (*view.UpdateVmPriorityEventView, error) {
	resp := view.UpdateVmPriorityEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"updateVmPriority": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachMdevDeviceFromVm operates on MdevDeviceFromVm
func (cli *ZSClient) DetachMdevDeviceFromVm(ctx context.Context, mdevDeviceUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/mdev-devices", mdevDeviceUuid, fmt.Sprintf("vm-instances/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DeleteVmHostname deletes VmHostname
func (cli *ZSClient) DeleteVmHostname(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// GetLicenseCapabilities gets LicenseCapabilities by uuid
func (cli *ZSClient) GetLicenseCapabilities(ctx context.Context) (*view.GetLicenseCapabilitiesView, error) {
	var resp view.GetLicenseCapabilitiesView
	if err := cli.GetWithRespKey(ctx, "v1/licenses/capabilities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRuleTemplate creates FirewallRuleTemplate
func (cli *ZSClient) CreateFirewallRuleTemplate(ctx context.Context, params param.CreateFirewallRuleTemplateParam) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	resp := view.VpcFirewallRuleTemplateInventoryView{}
	if err := cli.Post(ctx, "v1/vpcfirewalls/rules/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeIAM2ProjectState changes IAM2ProjectState
func (cli *ZSClient) ChangeIAM2ProjectState(ctx context.Context, uuid string, params param.ChangeIAM2ProjectStateParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/projects", uuid, "", map[string]interface{}{
		"changeIAM2ProjectState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmSoundType operates on VmSoundType
func (cli *ZSClient) SetVmSoundType(ctx context.Context, uuid string, params param.SetVmSoundTypeParam) (*view.SetVmSoundTypeEventView, error) {
	resp := view.SetVmSoundTypeEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmSoundType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MergeDataOnBackupStorage operates on MergeDataOnBackupStorage
func (cli *ZSClient) MergeDataOnBackupStorage(ctx context.Context, backupStorageUuid string, params param.MergeDataOnBackupStorageParam) (*view.MergeDataOnBackupStorageEventView, error) {
	resp := view.MergeDataOnBackupStorageEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/cdp-task/mergedata", backupStorageUuid, "", map[string]interface{}{
		"mergeDataOnBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCdpBackupStorageRequirement gets CdpBackupStorageRequirement by uuid
func (cli *ZSClient) GetCdpBackupStorageRequirement(ctx context.Context, uuid string) (*view.GetCdpBackupStorageRequirementView, error) {
	var resp view.GetCdpBackupStorageRequirementView
	if err := cli.GetWithRespKey(ctx, "v1/cdp-backup-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAttributesToIAM2VirtualIDGroup adds AttributesToIAM2VirtualIDGroup
func (cli *ZSClient) AddAttributesToIAM2VirtualIDGroup(ctx context.Context, params param.AddAttributesToIAM2VirtualIDGroupParam) (*view.AddAttributesToIAM2VirtualIDGroupEventView, error) {
	resp := view.AddAttributesToIAM2VirtualIDGroupEventView{}
	if err := cli.Post(ctx, "v1/iam2/projects/groups/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAffinityGroupState changes AffinityGroupState
func (cli *ZSClient) ChangeAffinityGroupState(ctx context.Context, uuid string, params param.ChangeAffinityGroupStateParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/affinity-groups", uuid, "", map[string]interface{}{
		"changeAffinityGroupState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSecurityGroupRuleState changes SecurityGroupRuleState
func (cli *ZSClient) ChangeSecurityGroupRuleState(ctx context.Context, securityGroupUuid string, params param.ChangeSecurityGroupRuleStateParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-groups", securityGroupUuid, "", map[string]interface{}{
		"changeSecurityGroupRuleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmNicToSecurityGroup adds VmNicToSecurityGroup
func (cli *ZSClient) AddVmNicToSecurityGroup(ctx context.Context, params param.AddVmNicToSecurityGroupParam) (*view.AddVmNicToSecurityGroupEventView, error) {
	resp := view.AddVmNicToSecurityGroupEventView{}
	if err := cli.Post(ctx, "v1/security-groups/{securityGroupUuid}/vm-instances/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAliyunRouteEntryFromRemote operates on AliyunRouteEntryFromRemote
func (cli *ZSClient) SyncAliyunRouteEntryFromRemote(ctx context.Context, vRouterUuid string, params param.SyncAliyunRouteEntryFromRemoteParam) (*view.VpcVirtualRouteEntryInventoryView, error) {
	resp := view.VpcVirtualRouteEntryInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/route-entry", vRouterUuid, "", map[string]interface{}{
		"syncAliyunRouteEntryFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEmailAddressOfSNSEmailEndpoint updates EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) UpdateEmailAddressOfSNSEmailEndpoint(ctx context.Context, params param.UpdateEmailAddressOfSNSEmailEndpointParam) (*view.SNSEmailAddressInventoryView, error) {
	resp := view.SNSEmailAddressInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-endpoints/emails/email-addresses", "", "", map[string]interface{}{
		"updateEmailAddressOfSNSEmailEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMetricLabelValue gets MetricLabelValue by uuid
func (cli *ZSClient) GetMetricLabelValue(ctx context.Context) (*view.GetMetricLabelValueView, error) {
	var resp view.GetMetricLabelValueView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/metrics/label-values", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateZonesClustersHostsForCreatingVm gets CandidateZonesClustersHostsForCreatingVm by uuid
func (cli *ZSClient) GetCandidateZonesClustersHostsForCreatingVm(ctx context.Context) (*view.GetCandidateZonesClustersHostsForCreatingVmView, error) {
	var resp view.GetCandidateZonesClustersHostsForCreatingVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/candidate-destinations", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateResourcePrice creates ResourcePrice
func (cli *ZSClient) CreateResourcePrice(ctx context.Context, params param.CreateResourcePriceParam) (*view.PriceInventoryView, error) {
	resp := view.PriceInventoryView{}
	if err := cli.Post(ctx, "v1/billings/prices", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveSchedulerJobGroupFromSchedulerTrigger removes SchedulerJobGroupFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobGroupFromSchedulerTrigger(ctx context.Context, schedulerJobGroupUuid string, schedulerTriggerUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/scheduler/jobgroups", schedulerJobGroupUuid, fmt.Sprintf("scheduler/triggers/%s", schedulerTriggerUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ChangeAccountPriceTableBinding changes AccountPriceTableBinding
func (cli *ZSClient) ChangeAccountPriceTableBinding(ctx context.Context, tableUuid string, accountUuid string, params param.ChangeAccountPriceTableBindingParam) (*view.PriceTableInventoryView, error) {
	resp := view.PriceTableInventoryView{}
	err := cli.PutWithSpec(ctx, "v1/billings/price-tables", tableUuid, fmt.Sprintf("accounts/%s", accountUuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVolumeQos deletes VolumeQos
func (cli *ZSClient) DeleteVolumeQos(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volumes", uuid, string(deleteMode))
}

// GetL3NetworkDhcpIpAddress gets L3NetworkDhcpIpAddress by uuid
func (cli *ZSClient) GetL3NetworkDhcpIpAddress(ctx context.Context, uuid string) (*view.GetL3NetworkDhcpIpAddressView, error) {
	var resp view.GetL3NetworkDhcpIpAddressView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRuleSet creates FirewallRuleSet
func (cli *ZSClient) CreateFirewallRuleSet(ctx context.Context, params param.CreateFirewallRuleSetParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.Post(ctx, "v1/vpcfirewalls/ruleSets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateJitSecretResourcePool creates JitSecretResourcePool
func (cli *ZSClient) CreateJitSecretResourcePool(ctx context.Context) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.Post(ctx, "v1/secret-resource-pool/jit", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBaremetalChassisPowerStatus gets BaremetalChassisPowerStatus by uuid
func (cli *ZSClient) GetBaremetalChassisPowerStatus(ctx context.Context, uuid string) (*view.GetBaremetalChassisPowerStatusView, error) {
	var resp view.GetBaremetalChassisPowerStatusView
	if err := cli.GetWithRespKey(ctx, "v1/baremetal/chassis", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVmUserDefinedXmlHookScript queries VmUserDefinedXmlHookScript list
func (cli *ZSClient) QueryVmUserDefinedXmlHookScript(ctx context.Context, params *param.QueryParam) ([]view.XmlHookInventoryView, error) {
	var resp []view.XmlHookInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/xml-hook-script", params, &resp)
}

func (cli *ZSClient) GetVmUserDefinedXmlHookScript(ctx context.Context, uuid string) (*view.XmlHookInventoryView, error) {
	var resp view.XmlHookInventoryView
	if err := cli.Get(ctx, "v1/vm-instances/xml-hook-script", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmUserDefinedXmlHookScript Pagination
func (cli *ZSClient) PageVmUserDefinedXmlHookScript(ctx context.Context, params *param.QueryParam) ([]view.XmlHookInventoryView, int, error) {
	var vmUserDefinedXmlHookScripts []view.XmlHookInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/xml-hook-script", params, &vmUserDefinedXmlHookScripts)
	return vmUserDefinedXmlHookScripts, total, err
}

// RefreshFirewall operates on Firewall
func (cli *ZSClient) RefreshFirewall(ctx context.Context, uuid string, params param.RefreshFirewallParam) (*view.VpcFirewallInventoryView, error) {
	resp := view.VpcFirewallInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vpcfirewalls/refresh", uuid, "", map[string]interface{}{
		"refreshFirewall": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL3NetworksFromIPsecConnection operates on L3NetworksFromIPsecConnection
func (cli *ZSClient) DetachL3NetworksFromIPsecConnection(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ipsec", uuid, string(deleteMode))
}

// UpdateAutoScalingGroupAddingNewInstanceRule updates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupAddingNewInstanceRule(ctx context.Context, uuid string, params param.UpdateAutoScalingGroupAddingNewInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/autoscaling/rules/adding-new-instance", uuid, "", map[string]interface{}{
		"updateAutoScalingGroupAddingNewInstanceRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFaultToleranceVms gets FaultToleranceVms by uuid
func (cli *ZSClient) GetFaultToleranceVms(ctx context.Context) (*view.GetFaultToleranceVmsView, error) {
	var resp view.GetFaultToleranceVmsView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/fault-tolerance/sub-vms", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunKeySecret deletes AliyunKeySecret
func (cli *ZSClient) DeleteAliyunKeySecret(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/key", uuid, string(deleteMode))
}

// CreateVmInstanceFromVolumeSnapshot creates VmInstanceFromVolumeSnapshot
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshot(ctx context.Context, params param.CreateVmInstanceFromVolumeSnapshotParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/from/volume-snapshots/{volumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerResetBareMetal2Chassis operates on PowerResetBareMetal2Chassis
func (cli *ZSClient) PowerResetBareMetal2Chassis(ctx context.Context, uuid string, params param.PowerResetBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"powerResetBareMetal2Chassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryVmMonitoringData operates on PrometheusQueryVmMonitoringData
func (cli *ZSClient) PrometheusQueryVmMonitoringData(ctx context.Context) (*view.PrometheusQueryVmMonitoringDataView, error) {
	var resp view.PrometheusQueryVmMonitoringDataView
	if err := cli.GetWithRespKey(ctx, "v1/prometheus/vm-instances", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateResourceConfigs updates ResourceConfigs
func (cli *ZSClient) UpdateResourceConfigs(ctx context.Context, params param.UpdateResourceConfigsParam) (*view.ResourceConfigStructView, error) {
	resp := view.ResourceConfigStructView{}
	if err := cli.Post(ctx, "v1/resource-configurations/{resourceUuid}/resource-configs/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcUserVpnGatewayFromLocal queries VpcUserVpnGatewayFromLocal list
func (cli *ZSClient) QueryVpcUserVpnGatewayFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcUserVpnGatewayInventoryView, error) {
	var resp []view.VpcUserVpnGatewayInventoryView
	return resp, cli.List(ctx, "v1/hybrid/user-vpn", params, &resp)
}

func (cli *ZSClient) GetVpcUserVpnGatewayFromLocal(ctx context.Context, uuid string) (*view.VpcUserVpnGatewayInventoryView, error) {
	var resp view.VpcUserVpnGatewayInventoryView
	if err := cli.Get(ctx, "v1/hybrid/user-vpn", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcUserVpnGatewayFromLocal Pagination
func (cli *ZSClient) PageVpcUserVpnGatewayFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcUserVpnGatewayInventoryView, int, error) {
	var vpcUserVpnGatewayFromLocals []view.VpcUserVpnGatewayInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/user-vpn", params, &vpcUserVpnGatewayFromLocals)
	return vpcUserVpnGatewayFromLocals, total, err
}

// RevertVolumeFromSnapshot operates on VolumeFromSnapshot
func (cli *ZSClient) RevertVolumeFromSnapshot(ctx context.Context, uuid string, params param.RevertVolumeFromSnapshotParam) (*view.RevertVolumeFromSnapshotEventView, error) {
	resp := view.RevertVolumeFromSnapshotEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/volume-snapshots", uuid, "", map[string]interface{}{
		"revertVolumeFromSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVolumeFromSnapshotAsync Async
func (cli *ZSClient) RevertVolumeFromSnapshotAsync(ctx context.Context, params param.RevertVolumeFromSnapshotParam) (string, error) {

	resource := "v1/volume-snapshots/{uuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetBlockPrimaryStorageMetadata gets BlockPrimaryStorageMetadata by uuid
func (cli *ZSClient) GetBlockPrimaryStorageMetadata(ctx context.Context, params param.GetBlockPrimaryStorageMetadataParam) (*view.BlockPrimaryStorageInventoryView, error) {
	resp := view.BlockPrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/block/metadata", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBonding updates Bonding
func (cli *ZSClient) UpdateBonding(ctx context.Context, uuid string, params param.UpdateBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/bondings", uuid, "", map[string]interface{}{
		"updateBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetManagementNodeArch gets ManagementNodeArch by uuid
func (cli *ZSClient) GetManagementNodeArch(ctx context.Context) (*view.GetManagementNodeArchView, error) {
	resp := view.GetManagementNodeArchView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/actions", "", "", map[string]interface{}{
		"getManagementNodeArch": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachScsiLunFromHost operates on ScsiLunFromHost
func (cli *ZSClient) DetachScsiLunFromHost(ctx context.Context, uuid string, params param.DetachScsiLunFromHostParam) (*view.ScsiLunInventoryView, error) {
	resp := view.ScsiLunInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/storage-devices/scsi-lun", uuid, "", map[string]interface{}{
		"detachScsiLunFromHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableCbtTask operates on DisableCbtTask
func (cli *ZSClient) DisableCbtTask(ctx context.Context, params param.DisableCbtTaskParam) (*view.CbtTaskInventoryView, error) {
	resp := view.CbtTaskInventoryView{}
	if err := cli.Post(ctx, "v1/cbt-task/disable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshLocalRaid operates on LocalRaid
func (cli *ZSClient) RefreshLocalRaid(ctx context.Context, params param.RefreshLocalRaidParam) (*view.RaidControllerInventoryView, error) {
	resp := view.RaidControllerInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/storage-devices/local-raid/actions", "", "", map[string]interface{}{
		"refreshLocalRaid": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSubscribeEvent updates SubscribeEvent
func (cli *ZSClient) UpdateSubscribeEvent(ctx context.Context, uuid string, params param.UpdateSubscribeEventParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/events/subscriptions", uuid, "", map[string]interface{}{
		"updateSubscribeEvent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmSshKey operates on VmSshKey
func (cli *ZSClient) SetVmSshKey(ctx context.Context, uuid string, params param.SetVmSshKeyParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmSshKey": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FailoverFaultToleranceVm operates on FailoverFaultToleranceVm
func (cli *ZSClient) FailoverFaultToleranceVm(ctx context.Context, params param.FailoverFaultToleranceVmParam) (*view.FailoverFaultToleranceVmEventView, error) {
	resp := view.FailoverFaultToleranceVmEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/fault-tolerance", "", "", map[string]interface{}{
		"failoverFaultToleranceVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EjectZBox operates on EjectZBox
func (cli *ZSClient) EjectZBox(ctx context.Context, uuid string, params param.EjectZBoxParam) (*view.ZBoxInventoryView, error) {
	resp := view.ZBoxInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zbox", uuid, "", map[string]interface{}{
		"ejectZBox": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryMetadata operates on PrometheusQueryMetadata
func (cli *ZSClient) PrometheusQueryMetadata(ctx context.Context) (*view.PrometheusQueryMetadataView, error) {
	var resp view.PrometheusQueryMetadataView
	if err := cli.GetWithRespKey(ctx, "v1/prometheus/meta-data", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIpAddress queries IpAddress list
func (cli *ZSClient) QueryIpAddress(ctx context.Context, params *param.QueryParam) ([]view.UsedIpInventoryView, error) {
	var resp []view.UsedIpInventoryView
	return resp, cli.List(ctx, "v1/l3-networks/ip-address", params, &resp)
}

func (cli *ZSClient) GetIpAddress(ctx context.Context, uuid string) (*view.UsedIpInventoryView, error) {
	var resp view.UsedIpInventoryView
	if err := cli.Get(ctx, "v1/l3-networks/ip-address", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIpAddress Pagination
func (cli *ZSClient) PageIpAddress(ctx context.Context, params *param.QueryParam) ([]view.UsedIpInventoryView, int, error) {
	var ipAddress []view.UsedIpInventoryView
	total, err := cli.Page(ctx, "v1/l3-networks/ip-address", params, &ipAddress)
	return ipAddress, total, err
}

// DeleteFirewallRuleTemplate deletes FirewallRuleTemplate
func (cli *ZSClient) DeleteFirewallRuleTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpcfirewalls/rules/templates", uuid, string(deleteMode))
}

// DetachPciDeviceFromVm operates on PciDeviceFromVm
func (cli *ZSClient) DetachPciDeviceFromVm(ctx context.Context, params param.DetachPciDeviceFromVmParam) (*view.PciDeviceInventoryView, error) {
	resp := view.PciDeviceInventoryView{}
	if err := cli.Post(ctx, "v1/pci-device/pci-devices/{pciDeviceUuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExecuteGuestVmCommand operates on ExecuteGuestVmCommand
func (cli *ZSClient) ExecuteGuestVmCommand(ctx context.Context, params param.ExecuteGuestVmCommandParam) (*view.ExecuteGuestVmCommandEventView, error) {
	resp := view.ExecuteGuestVmCommandEventView{}
	if err := cli.Post(ctx, "v1/vm-instances/commands/exec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVpcSharedQosBandwidth changes VpcSharedQosBandwidth
func (cli *ZSClient) ChangeVpcSharedQosBandwidth(ctx context.Context, params param.ChangeVpcSharedQosBandwidthParam) (*view.VpcSharedQosInventoryView, error) {
	resp := view.VpcSharedQosInventoryView{}
	if err := cli.Post(ctx, "v1/vips/sharedqos/{sharedQosUuid}/bandwidth/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveAttributesFromIAM2VirtualIDGroup removes AttributesFromIAM2VirtualIDGroup
func (cli *ZSClient) RemoveAttributesFromIAM2VirtualIDGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects/groups", uuid, string(deleteMode))
}

// AddAttributesToIAM2VirtualID adds AttributesToIAM2VirtualID
func (cli *ZSClient) AddAttributesToIAM2VirtualID(ctx context.Context, params param.AddAttributesToIAM2VirtualIDParam) (*view.AddAttributesToIAM2VirtualIDEventView, error) {
	resp := view.AddAttributesToIAM2VirtualIDEventView{}
	if err := cli.Post(ctx, "v1/iam2/virtual-ids/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVolume operates on FlattenVolume
func (cli *ZSClient) FlattenVolume(ctx context.Context, uuid string, params param.FlattenVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/volumes", uuid, "", map[string]interface{}{
		"flattenVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVolumeAsync Async
func (cli *ZSClient) FlattenVolumeAsync(ctx context.Context, params param.FlattenVolumeParam) (string, error) {

	resource := "v1/volumes/{uuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// CreateAliyunDiskFromRemote creates AliyunDiskFromRemote
func (cli *ZSClient) CreateAliyunDiskFromRemote(ctx context.Context, params param.CreateAliyunDiskFromRemoteParam) (*view.AliyunDiskInventoryView, error) {
	resp := view.AliyunDiskInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/disk", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsSecurityGroupRuleRemote deletes EcsSecurityGroupRuleRemote
func (cli *ZSClient) DeleteEcsSecurityGroupRuleRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/security-group-rule/remote", uuid, string(deleteMode))
}

// GetCandidateAffinityGroupForAttachingVm gets CandidateAffinityGroupForAttachingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForAttachingVm(ctx context.Context) (*view.GetCandidateAffinityGroupForAttachingVmView, error) {
	var resp view.GetCandidateAffinityGroupForAttachingVmView
	if err := cli.GetWithRespKey(ctx, "v1/affinityGroup/attachingVm", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachAliyunDiskFromEcs operates on AliyunDiskFromEcs
func (cli *ZSClient) DetachAliyunDiskFromEcs(ctx context.Context, params param.DetachAliyunDiskFromEcsParam) (*view.DetachAliyunDiskFromEcsEventView, error) {
	resp := view.DetachAliyunDiskFromEcsEventView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/disk/{uuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFirewallIpSetTemplate updates FirewallIpSetTemplate
func (cli *ZSClient) UpdateFirewallIpSetTemplate(ctx context.Context, uuid string, params param.UpdateFirewallIpSetTemplateParam) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	resp := view.VpcFirewallIpSetTemplateInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vpcfirewalls/ipset/templates", uuid, "", map[string]interface{}{
		"updateFirewallIpSetTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAccessControlListRedirectRule adds AccessControlListRedirectRule
func (cli *ZSClient) AddAccessControlListRedirectRule(ctx context.Context, params param.AddAccessControlListRedirectRuleParam) (*view.AccessControlListEntryInventoryView, error) {
	resp := view.AccessControlListEntryInventoryView{}
	if err := cli.Post(ctx, "v1/access-control-lists/{aclUuid}/redirectRules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachHostFromHostSchedulingRuleGroup operates on HostFromHostSchedulingRuleGroup
func (cli *ZSClient) DetachHostFromHostSchedulingRuleGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hostSchedulingRuleGroup", uuid, string(deleteMode))
}

// UpdateAliyunRouteInterfaceRemote updates AliyunRouteInterfaceRemote
func (cli *ZSClient) UpdateAliyunRouteInterfaceRemote(ctx context.Context, uuid string, params param.UpdateAliyunRouteInterfaceRemoteParam) (*view.UpdateAliyunRouteInterfaceRemoteEventView, error) {
	resp := view.UpdateAliyunRouteInterfaceRemoteEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/router-interface", uuid, "", map[string]interface{}{
		"updateAliyunRouteInterfaceRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceSpecCandidates gets PciDeviceSpecCandidates by uuid
func (cli *ZSClient) GetPciDeviceSpecCandidates(ctx context.Context) (*view.GetPciDeviceSpecCandidatesView, error) {
	var resp view.GetPciDeviceSpecCandidatesView
	if err := cli.GetWithRespKey(ctx, "v1/pci-device-specs/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryPassThrough operates on PrometheusQueryPassThrough
func (cli *ZSClient) PrometheusQueryPassThrough(ctx context.Context) (*view.PrometheusQueryPassThroughView, error) {
	var resp view.PrometheusQueryPassThroughView
	if err := cli.GetWithRespKey(ctx, "v1/prometheus/all", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVmNicToVm operates on VmNicToVm
func (cli *ZSClient) AttachVmNicToVm(ctx context.Context, params param.AttachVmNicToVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/{vmInstanceUuid}/nices/{vmNicUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveMonFromCephBackupStorage removes MonFromCephBackupStorage
func (cli *ZSClient) RemoveMonFromCephBackupStorage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/backup-storage/ceph", uuid, string(deleteMode))
}

// GetVmDeviceAddress gets VmDeviceAddress by uuid
func (cli *ZSClient) GetVmDeviceAddress(ctx context.Context) (*view.GetVmDeviceAddressView, error) {
	var resp view.GetVmDeviceAddressView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/devices", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveInstanceFromMonitorGroup removes InstanceFromMonitorGroup
func (cli *ZSClient) RemoveInstanceFromMonitorGroup(ctx context.Context, groupUuid string, instanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zwatch/monitorgroups", groupUuid, fmt.Sprintf("actions/%s", instanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// CleanQueue operates on Queue
func (cli *ZSClient) CleanQueue(ctx context.Context, params param.CleanQueueParam) (*view.CleanQueueEventView, error) {
	resp := view.CleanQueueEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/clean/queue", "", "", map[string]interface{}{
		"cleanQueue": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveAccessControlListFromLoadBalancer removes AccessControlListFromLoadBalancer
func (cli *ZSClient) RemoveAccessControlListFromLoadBalancer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/listeners", uuid, string(deleteMode))
}

// RemoveLabelFromEventSubscription removes LabelFromEventSubscription
func (cli *ZSClient) RemoveLabelFromEventSubscription(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/events/subscriptions/labels", uuid, string(deleteMode))
}

// SdnControllerRemoveHost operates on SdnControllerRemoveHost
func (cli *ZSClient) SdnControllerRemoveHost(ctx context.Context, sdnControllerUuid string, hostUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/sdn-controllers", sdnControllerUuid, fmt.Sprintf("hosts/%s", hostUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DetachCCSCertificateFromUser operates on CCSCertificateFromUser
func (cli *ZSClient) DetachCCSCertificateFromUser(ctx context.Context, params param.DetachCCSCertificateFromUserParam) (*view.DetachCCSCertificateFromUserEventView, error) {
	resp := view.DetachCCSCertificateFromUserEventView{}
	if err := cli.Post(ctx, "v1/crypto/ccs-certificate/detach-user/{userUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetManagementNodeOS gets ManagementNodeOS by uuid
func (cli *ZSClient) GetManagementNodeOS(ctx context.Context) (*view.GetManagementNodeOSView, error) {
	resp := view.GetManagementNodeOSView{}
	if err := cli.PutWithRespKey(ctx, "v1/management/actions", "", "", map[string]interface{}{
		"getManagementNodeOS": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateLdapBinding creates LdapBinding
func (cli *ZSClient) CreateLdapBinding(ctx context.Context, params param.CreateLdapBindingParam) (*view.LdapAccountRefInventoryView, error) {
	resp := view.LdapAccountRefInventoryView{}
	if err := cli.Post(ctx, "v1/ldap/bindings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExecuteDRSScheduling operates on ExecuteDRSScheduling
func (cli *ZSClient) ExecuteDRSScheduling(ctx context.Context, uuid string, params param.ExecuteDRSSchedulingParam) (*view.ExecuteDRSSchedulingEventView, error) {
	resp := view.ExecuteDRSSchedulingEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/clusters/drs", uuid, "", map[string]interface{}{
		"executeDRSScheduling": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsSecurityGroupRuleRemote creates EcsSecurityGroupRuleRemote
func (cli *ZSClient) CreateEcsSecurityGroupRuleRemote(ctx context.Context, params param.CreateEcsSecurityGroupRuleRemoteParam) (*view.EcsSecurityGroupRuleInventoryView, error) {
	resp := view.EcsSecurityGroupRuleInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/security-group-rule", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmQga gets VmQga by uuid
func (cli *ZSClient) GetVmQga(ctx context.Context, uuid string) (*view.GetVmQgaView, error) {
	var resp view.GetVmQgaView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PreviewResourceStack operates on PreviewResourceStack
func (cli *ZSClient) PreviewResourceStack(ctx context.Context, params param.PreviewResourceStackParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post(ctx, "v1/cloudformation/stack/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmvNUMATopology gets VmvNUMATopology by uuid
func (cli *ZSClient) GetVmvNUMATopology(ctx context.Context, uuid string) (*view.GetVmvNUMATopologyView, error) {
	var resp view.GetVmvNUMATopologyView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveSchedulerJobsFromSchedulerJobGroup removes SchedulerJobsFromSchedulerJobGroup
func (cli *ZSClient) RemoveSchedulerJobsFromSchedulerJobGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/scheduler/jobgroups", uuid, string(deleteMode))
}

// ChangeTicketStatus changes TicketStatus
func (cli *ZSClient) ChangeTicketStatus(ctx context.Context, uuid string, params param.ChangeTicketStatusParam) (*view.TicketInventoryView, error) {
	resp := view.TicketInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/tickets", uuid, "", map[string]interface{}{
		"changeTicketStatus": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostPhysicalMemoryFacts gets HostPhysicalMemoryFacts by uuid
func (cli *ZSClient) GetHostPhysicalMemoryFacts(ctx context.Context, uuid string) (*view.GetHostPhysicalMemoryFactsView, error) {
	var resp view.GetHostPhysicalMemoryFactsView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/physical-memory-facts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseInfo gets LicenseInfo by uuid
func (cli *ZSClient) GetLicenseInfo(ctx context.Context) (*view.LicenseInventoryView, error) {
	var resp view.GetLicenseInfoView
	if err := cli.GetWithRespKey(ctx, "v1/licenses", "", "inventory", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeSchedulerState changes SchedulerState
func (cli *ZSClient) ChangeSchedulerState(ctx context.Context, uuid string, params param.ChangeSchedulerStateParam) (*view.SchedulerJobInventoryView, error) {
	resp := view.SchedulerJobInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/schedulers", uuid, "", map[string]interface{}{
		"changeSchedulerState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPriceTableToAccount operates on PriceTableToAccount
func (cli *ZSClient) AttachPriceTableToAccount(ctx context.Context, params param.AttachPriceTableToAccountParam) (*view.PriceTableInventoryView, error) {
	resp := view.PriceTableInventoryView{}
	if err := cli.Post(ctx, "v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateMdevDevices operates on MdevDevices
func (cli *ZSClient) GenerateMdevDevices(ctx context.Context, pciDeviceUuid string, params param.GenerateMdevDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/pci-devices", pciDeviceUuid, "", map[string]interface{}{
		"generateMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PreviewResourceFromApp operates on PreviewResourceFromApp
func (cli *ZSClient) PreviewResourceFromApp(ctx context.Context, params param.PreviewResourceFromAppParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post(ctx, "v1/appcenter/app/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSNSTopicState changes SNSTopicState
func (cli *ZSClient) ChangeSNSTopicState(ctx context.Context, uuid string, params param.ChangeSNSTopicStateParam) (*view.SNSTopicInventoryView, error) {
	resp := view.SNSTopicInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/topics", uuid, "", map[string]interface{}{
		"changeSNSTopicState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachScsiLunToVmInstance operates on ScsiLunToVmInstance
func (cli *ZSClient) AttachScsiLunToVmInstance(ctx context.Context, params param.AttachScsiLunToVmInstanceParam) (*view.ScsiLunInventoryView, error) {
	resp := view.ScsiLunInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/{vmInstanceUuid}/scsi-lun/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveRemoteCidrsFromIPsecConnection removes RemoteCidrsFromIPsecConnection
func (cli *ZSClient) RemoveRemoteCidrsFromIPsecConnection(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ipsec", uuid, string(deleteMode))
}

// GetIAM2ProjectsOfVirtualID gets IAM2ProjectsOfVirtualID by uuid
func (cli *ZSClient) GetIAM2ProjectsOfVirtualID(ctx context.Context) (*view.GetIAM2ProjectsOfVirtualIDView, error) {
	var resp view.GetIAM2ProjectsOfVirtualIDView
	if err := cli.GetWithRespKey(ctx, "v1/iam2/virtual-ids/projects", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmToAffinityGroup adds VmToAffinityGroup
func (cli *ZSClient) AddVmToAffinityGroup(ctx context.Context, params param.AddVmToAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.Post(ctx, "v1/affinity-groups/{affinityGroupUuid}/vm-instances/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageAllocatorStrategies gets PrimaryStorageAllocatorStrategies by uuid
func (cli *ZSClient) GetPrimaryStorageAllocatorStrategies(ctx context.Context) (*view.GetPrimaryStorageAllocatorStrategiesView, error) {
	var resp view.GetPrimaryStorageAllocatorStrategiesView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/allocators/strategies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPlatformTimeZone gets PlatformTimeZone by uuid
func (cli *ZSClient) GetPlatformTimeZone(ctx context.Context) (*view.GetPlatformTimeZoneView, error) {
	var resp view.GetPlatformTimeZoneView
	if err := cli.GetWithRespKey(ctx, "v1/management-nodes/platform-timezone", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPolicyFromUser operates on PolicyFromUser
func (cli *ZSClient) DetachPolicyFromUser(ctx context.Context, userUuid string, policyUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/accounts/users", userUuid, fmt.Sprintf("policies/%s", policyUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// SetVmInstanceDefaultCdRom operates on VmInstanceDefaultCdRom
func (cli *ZSClient) SetVmInstanceDefaultCdRom(ctx context.Context, vmInstanceUuid string, uuid string, params param.SetVmInstanceDefaultCdRomParam) (*view.VmCdRomInventoryView, error) {
	resp := view.VmCdRomInventoryView{}
	err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, fmt.Sprintf("cdroms/%s/actions", uuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshSharedblockDeviceCapacity operates on SharedblockDeviceCapacity
func (cli *ZSClient) RefreshSharedblockDeviceCapacity(ctx context.Context, params param.RefreshSharedblockDeviceCapacityParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	resp := view.SharedBlockGroupPrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/sharedblockgroup/{sharedBlockGroupUuid}/sharedblocks/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FstrimVm operates on FstrimVm
func (cli *ZSClient) FstrimVm(ctx context.Context, params param.FstrimVmParam) (*view.FstrimVmEventView, error) {
	resp := view.FstrimVmEventView{}
	if err := cli.Post(ctx, "v1/vm-instances/{uuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL2NetworkFromCluster operates on L2NetworkFromCluster
func (cli *ZSClient) DetachL2NetworkFromCluster(ctx context.Context, l2NetworkUuid string, clusterUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/l2-networks", l2NetworkUuid, fmt.Sprintf("clusters/%s", clusterUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// SyncAINginxConfiguration operates on AINginxConfiguration
func (cli *ZSClient) SyncAINginxConfiguration(ctx context.Context, params param.SyncAINginxConfigurationParam) (*view.SyncAINginxConfigurationView, error) {
	resp := view.SyncAINginxConfigurationView{}
	if err := cli.Post(ctx, "v1/ai/nginx/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MatchModelServiceTemplateWithModel operates on MatchModelServiceTemplateWithModel
func (cli *ZSClient) MatchModelServiceTemplateWithModel(ctx context.Context, params param.MatchModelServiceTemplateWithModelParam) (*view.MatchModelServiceTemplateWithModelEventView, error) {
	resp := view.MatchModelServiceTemplateWithModelEventView{}
	if err := cli.Post(ctx, "v1/ai/model-services/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmNicState changes VmNicState
func (cli *ZSClient) ChangeVmNicState(ctx context.Context, vmNicUuid string, params param.ChangeVmNicStateParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/nics", vmNicUuid, "", map[string]interface{}{
		"changeVmNicState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddPolicyStatementsToRole adds PolicyStatementsToRole
func (cli *ZSClient) AddPolicyStatementsToRole(ctx context.Context, params param.AddPolicyStatementsToRoleParam) (*view.AddPolicyStatementsToRoleEventView, error) {
	resp := view.AddPolicyStatementsToRoleEventView{}
	if err := cli.Post(ctx, "v1/identities/roles/{uuid}/policy-statements", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnprotectVmInstanceRecoveryPoint operates on UnprotectVmInstanceRecoveryPoint
func (cli *ZSClient) UnprotectVmInstanceRecoveryPoint(ctx context.Context, vmInstanceUuid string, params param.UnprotectVmInstanceRecoveryPointParam) (*view.UnprotectVmInstanceRecoveryPointEventView, error) {
	resp := view.UnprotectVmInstanceRecoveryPointEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"unprotectVmInstanceRecoveryPoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachVipFromVpcSharedQos operates on VipFromVpcSharedQos
func (cli *ZSClient) DetachVipFromVpcSharedQos(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vips/sharedqos", uuid, string(deleteMode))
}

// ApplyRuleSetChanges operates on RuleSetChanges
func (cli *ZSClient) ApplyRuleSetChanges(ctx context.Context, uuid string, params param.ApplyRuleSetChangesParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vpcfirewalls/ruleSets/apply", uuid, "", map[string]interface{}{
		"applyRuleSetChanges": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVm operates on PrimaryStorageMigrateVm
func (cli *ZSClient) PrimaryStorageMigrateVm(ctx context.Context, vmInstanceUuid string, params param.PrimaryStorageMigrateVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"primaryStorageMigrateVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVmAsync Async
func (cli *ZSClient) PrimaryStorageMigrateVmAsync(ctx context.Context, params param.PrimaryStorageMigrateVmParam) (string, error) {

	resource := "v1/vm-instances/{vmInstanceUuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// RecoverDatabaseFromBackup operates on DatabaseFromBackup
func (cli *ZSClient) RecoverDatabaseFromBackup(ctx context.Context, params param.RecoverDatabaseFromBackupParam) (*view.RecoverDatabaseFromBackupEventView, error) {
	resp := view.RecoverDatabaseFromBackupEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/database-backups/actions", "", "", map[string]interface{}{
		"recoverDatabaseFromBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2TickFlowCollection creates IAM2TickFlowCollection
func (cli *ZSClient) CreateIAM2TickFlowCollection(ctx context.Context, params param.CreateIAM2TickFlowCollectionParam) (*view.TicketFlowCollectionInventoryView, error) {
	resp := view.TicketFlowCollectionInventoryView{}
	if err := cli.Post(ctx, "v1/tickets/flow-collections", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UngenerateMdevDevices operates on UngenerateMdevDevices
func (cli *ZSClient) UngenerateMdevDevices(ctx context.Context, pciDeviceUuid string, params param.UngenerateMdevDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/pci-devices", pciDeviceUuid, "", map[string]interface{}{
		"ungenerateMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveDirectory operates on MoveDirectory
func (cli *ZSClient) MoveDirectory(ctx context.Context, params param.MoveDirectoryParam) (*view.MoveDirectoryEventView, error) {
	resp := view.MoveDirectoryEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/move/directory", "", "", map[string]interface{}{
		"moveDirectory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVRouterOspfNeighbor gets VRouterOspfNeighbor by uuid
func (cli *ZSClient) GetVRouterOspfNeighbor(ctx context.Context, uuid string) (*view.GetVRouterOspfNeighborView, error) {
	var resp view.GetVRouterOspfNeighborView
	if err := cli.GetWithRespKey(ctx, "v1/routerArea", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpcVRouter creates VpcVRouter
func (cli *ZSClient) CreateVpcVRouter(ctx context.Context, params param.CreateVpcVRouterParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsInstanceFromRemote operates on EcsInstanceFromRemote
func (cli *ZSClient) SyncEcsInstanceFromRemote(ctx context.Context, params param.SyncEcsInstanceFromRemoteParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/ecs/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMdevDeviceSpecCandidates gets MdevDeviceSpecCandidates by uuid
func (cli *ZSClient) GetMdevDeviceSpecCandidates(ctx context.Context) (*view.GetMdevDeviceSpecCandidatesView, error) {
	var resp view.GetMdevDeviceSpecCandidatesView
	if err := cli.GetWithRespKey(ctx, "v1/mdev-device-specs/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFlowMeterRouterId gets FlowMeterRouterId by uuid
func (cli *ZSClient) GetFlowMeterRouterId(ctx context.Context, uuid string) (*view.GetFlowMeterRouterIdView, error) {
	var resp view.GetFlowMeterRouterIdView
	if err := cli.GetWithRespKey(ctx, "v1/flowmeters", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceCandidatesForNewCreateVm gets PciDeviceCandidatesForNewCreateVm by uuid
func (cli *ZSClient) GetPciDeviceCandidatesForNewCreateVm(ctx context.Context) (*view.GetPciDeviceCandidatesForNewCreateVmView, error) {
	var resp view.GetPciDeviceCandidatesForNewCreateVmView
	if err := cli.GetWithRespKey(ctx, "v1/pci-device/candidate-pci-devices-for-new-create-vm", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryDataCenterFromLocal queries DataCenterFromLocal list
func (cli *ZSClient) QueryDataCenterFromLocal(ctx context.Context, params *param.QueryParam) ([]view.DataCenterInventoryView, error) {
	var resp []view.DataCenterInventoryView
	return resp, cli.List(ctx, "v1/hybrid/data-center", params, &resp)
}

func (cli *ZSClient) GetDataCenterFromLocal(ctx context.Context, uuid string) (*view.DataCenterInventoryView, error) {
	var resp view.DataCenterInventoryView
	if err := cli.Get(ctx, "v1/hybrid/data-center", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDataCenterFromLocal Pagination
func (cli *ZSClient) PageDataCenterFromLocal(ctx context.Context, params *param.QueryParam) ([]view.DataCenterInventoryView, int, error) {
	var dataCenterFromLocals []view.DataCenterInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/data-center", params, &dataCenterFromLocals)
	return dataCenterFromLocals, total, err
}

// GetHostTask gets HostTask by uuid
func (cli *ZSClient) GetHostTask(ctx context.Context) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/task-details", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddResourceToIAM2Project adds ResourceToIAM2Project
func (cli *ZSClient) AddResourceToIAM2Project(ctx context.Context, params param.AddResourceToIAM2ProjectParam) (*view.AddResourceToIAM2ProjectEventView, error) {
	resp := view.AddResourceToIAM2ProjectEventView{}
	if err := cli.Post(ctx, "v1/iam2/projects/add/resource/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAlarmData gets AlarmData by uuid
func (cli *ZSClient) GetAlarmData(ctx context.Context) (*view.GetAlarmDataView, error) {
	var resp view.GetAlarmDataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/alarm-histories", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeV2VConversionHostState changes V2VConversionHostState
func (cli *ZSClient) ChangeV2VConversionHostState(ctx context.Context, uuid string, params param.ChangeV2VConversionHostStateParam) (*view.V2VConversionHostInventoryView, error) {
	resp := view.V2VConversionHostInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/v2v-conversion-hosts", uuid, "", map[string]interface{}{
		"changeV2VConversionHostState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoverResourceSplitBrain operates on ResourceSplitBrain
func (cli *ZSClient) RecoverResourceSplitBrain(ctx context.Context, resourceUuid string, params param.RecoverResourceSplitBrainParam) (*view.RecoverResourceSplitBrainEventView, error) {
	resp := view.RecoverResourceSplitBrainEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage/mini", resourceUuid, "", map[string]interface{}{
		"recoverResourceSplitBrain": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsOpensourceVersion operates on IsOpensourceVersion
func (cli *ZSClient) IsOpensourceVersion(ctx context.Context) (*view.IsOpensourceVersionView, error) {
	var resp view.IsOpensourceVersionView
	if err := cli.GetWithRespKey(ctx, "v1/meta-data/opensource", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsInstanceFromEcsImage creates EcsInstanceFromEcsImage
func (cli *ZSClient) CreateEcsInstanceFromEcsImage(ctx context.Context, params param.CreateEcsInstanceFromEcsImageParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/ecs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceFromResourceStack gets ResourceFromResourceStack by uuid
func (cli *ZSClient) GetResourceFromResourceStack(ctx context.Context) (*view.GetResourceFromResourceStackView, error) {
	var resp view.GetResourceFromResourceStackView
	if err := cli.GetWithRespKey(ctx, "v1/cloudformation/stack/resources", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveResourcesToDirectory operates on MoveResourcesToDirectory
func (cli *ZSClient) MoveResourcesToDirectory(ctx context.Context, params param.MoveResourcesToDirectoryParam) (*view.MoveResourcesToDirectoryEventView, error) {
	resp := view.MoveResourcesToDirectoryEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/move/resources/directory", "", "", map[string]interface{}{
		"moveResourcesToDirectory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSupportedCloudFormationResources gets SupportedCloudFormationResources by uuid
func (cli *ZSClient) GetSupportedCloudFormationResources(ctx context.Context) (*view.GetSupportedCloudFormationResourcesView, error) {
	var resp view.GetSupportedCloudFormationResourcesView
	if err := cli.GetWithRespKey(ctx, "v1/cloudformation/resources", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIdentityZoneInLocal deletes IdentityZoneInLocal
func (cli *ZSClient) DeleteIdentityZoneInLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/identity-zone", uuid, string(deleteMode))
}

// QueryVRouterOspfNetwork queries VRouterOspfNetwork list
func (cli *ZSClient) QueryVRouterOspfNetwork(ctx context.Context, params *param.QueryParam) ([]view.NetworkRouterAreaRefInventoryView, error) {
	var resp []view.NetworkRouterAreaRefInventoryView
	return resp, cli.List(ctx, "v1/routerArea/network", params, &resp)
}

func (cli *ZSClient) GetVRouterOspfNetwork(ctx context.Context, uuid string) (*view.NetworkRouterAreaRefInventoryView, error) {
	var resp view.NetworkRouterAreaRefInventoryView
	if err := cli.Get(ctx, "v1/routerArea/network", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterOspfNetwork Pagination
func (cli *ZSClient) PageVRouterOspfNetwork(ctx context.Context, params *param.QueryParam) ([]view.NetworkRouterAreaRefInventoryView, int, error) {
	var vRouterOspfNetworks []view.NetworkRouterAreaRefInventoryView
	total, err := cli.Page(ctx, "v1/routerArea/network", params, &vRouterOspfNetworks)
	return vRouterOspfNetworks, total, err
}

// ChangeIAM2VirtualIDState changes IAM2VirtualIDState
func (cli *ZSClient) ChangeIAM2VirtualIDState(ctx context.Context, uuid string, params param.ChangeIAM2VirtualIDStateParam) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/virtual-ids", uuid, "", map[string]interface{}{
		"changeIAM2VirtualIDState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnregisterLicenseServer operates on LicenseServer
func (cli *ZSClient) UnregisterLicenseServer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/license-server/client", uuid, string(deleteMode))
}

// CreateVmUserDefinedXmlHookScript creates VmUserDefinedXmlHookScript
func (cli *ZSClient) CreateVmUserDefinedXmlHookScript(ctx context.Context, params param.CreateVmUserDefinedXmlHookScriptParam) (*view.XmlHookInventoryView, error) {
	resp := view.XmlHookInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/xml-hook-script", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradeToLicenseServer operates on UpgradeToLicenseServer
func (cli *ZSClient) UpgradeToLicenseServer(ctx context.Context) (*view.LicenseAuthorizedNodeInventoryView, error) {
	resp := view.LicenseAuthorizedNodeInventoryView{}
	if err := cli.Post(ctx, "v1/license-server", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachAppBuildSystemToZone operates on AppBuildSystemToZone
func (cli *ZSClient) DetachAppBuildSystemToZone(ctx context.Context, zoneUuid string, buildSystemUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zones", zoneUuid, fmt.Sprintf("buildsystem/%s", buildSystemUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetAppBuildSystemCapacity gets AppBuildSystemCapacity by uuid
func (cli *ZSClient) GetAppBuildSystemCapacity(ctx context.Context) (*view.GetAppBuildSystemCapacityView, error) {
	var resp view.GetAppBuildSystemCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/appcenter/buildsystem/capacities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAttachableVpcL3Network gets AttachableVpcL3Network by uuid
func (cli *ZSClient) GetAttachableVpcL3Network(ctx context.Context, params param.GetAttachableVpcL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.Post(ctx, "v1/vpc/virtual-routers/{uuid}/attachable-vpc-l3s", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBaremetalChassisState changes BaremetalChassisState
func (cli *ZSClient) ChangeBaremetalChassisState(ctx context.Context, uuid string, params param.ChangeBaremetalChassisStateParam) (*view.BaremetalChassisInventoryView, error) {
	resp := view.BaremetalChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal/chassis", uuid, "", map[string]interface{}{
		"changeBaremetalChassisState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkMtu gets L3NetworkMtu by uuid
func (cli *ZSClient) GetL3NetworkMtu(ctx context.Context, uuid string) (*view.GetL3NetworkMtuView, error) {
	var resp view.GetL3NetworkMtuView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVipToLoadBalancer operates on VipToLoadBalancer
func (cli *ZSClient) AttachVipToLoadBalancer(ctx context.Context, params param.AttachVipToLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	resp := view.LoadBalancerInventoryView{}
	if err := cli.Post(ctx, "v1/load-balancers/{loadBalancerUuid}/vip/{vipUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSecurityGroupRulePriority updates SecurityGroupRulePriority
func (cli *ZSClient) UpdateSecurityGroupRulePriority(ctx context.Context, securityGroupUuid string, params param.UpdateSecurityGroupRulePriorityParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-groups", securityGroupUuid, "", map[string]interface{}{
		"updateSecurityGroupRulePriority": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryConnectionBetweenL3NetworkAndAliyunVSwitch queries ConnectionBetweenL3NetworkAndAliyunVSwitch list
func (cli *ZSClient) QueryConnectionBetweenL3NetworkAndAliyunVSwitch(ctx context.Context, params *param.QueryParam) ([]view.ConnectionRelationShipInventoryView, error) {
	var resp []view.ConnectionRelationShipInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/relationships", params, &resp)
}

func (cli *ZSClient) GetConnectionBetweenL3NetworkAndAliyunVSwitch(ctx context.Context, uuid string) (*view.ConnectionRelationShipInventoryView, error) {
	var resp view.ConnectionRelationShipInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/relationships", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageConnectionBetweenL3NetworkAndAliyunVSwitch Pagination
func (cli *ZSClient) PageConnectionBetweenL3NetworkAndAliyunVSwitch(ctx context.Context, params *param.QueryParam) ([]view.ConnectionRelationShipInventoryView, int, error) {
	var connectionBetweenL3NetworkAndAliyunVSwitchs []view.ConnectionRelationShipInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/relationships", params, &connectionBetweenL3NetworkAndAliyunVSwitchs)
	return connectionBetweenL3NetworkAndAliyunVSwitchs, total, err
}

// AddDnsToL3Network adds DnsToL3Network
func (cli *ZSClient) AddDnsToL3Network(ctx context.Context, params param.AddDnsToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.Post(ctx, "v1/l3-networks/{l3NetworkUuid}/dns", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryPortMirrorNetworkUsedIp queries PortMirrorNetworkUsedIp list
func (cli *ZSClient) QueryPortMirrorNetworkUsedIp(ctx context.Context, params *param.QueryParam) ([]view.MirrorNetworkUsedIpInventoryView, error) {
	var resp []view.MirrorNetworkUsedIpInventoryView
	return resp, cli.List(ctx, "v1/port-mirrors/networks/usedIps", params, &resp)
}

func (cli *ZSClient) GetPortMirrorNetworkUsedIp(ctx context.Context, uuid string) (*view.MirrorNetworkUsedIpInventoryView, error) {
	var resp view.MirrorNetworkUsedIpInventoryView
	if err := cli.Get(ctx, "v1/port-mirrors/networks/usedIps", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePortMirrorNetworkUsedIp Pagination
func (cli *ZSClient) PagePortMirrorNetworkUsedIp(ctx context.Context, params *param.QueryParam) ([]view.MirrorNetworkUsedIpInventoryView, int, error) {
	var portMirrorNetworkUsedIps []view.MirrorNetworkUsedIpInventoryView
	total, err := cli.Page(ctx, "v1/port-mirrors/networks/usedIps", params, &portMirrorNetworkUsedIps)
	return portMirrorNetworkUsedIps, total, err
}

// SetVmMonitorNumber operates on VmMonitorNumber
func (cli *ZSClient) SetVmMonitorNumber(ctx context.Context, uuid string, params param.SetVmMonitorNumberParam) (*view.SetVmMonitorNumberEventView, error) {
	resp := view.SetVmMonitorNumberEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmMonitorNumber": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeLoadBalancerBackendServer changes LoadBalancerBackendServer
func (cli *ZSClient) ChangeLoadBalancerBackendServer(ctx context.Context, serverGroupUuid string, params param.ChangeLoadBalancerBackendServerParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/load-balancers/servergroups", serverGroupUuid, "", map[string]interface{}{
		"changeLoadBalancerBackendServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsVSwitchRemote creates EcsVSwitchRemote
func (cli *ZSClient) CreateEcsVSwitchRemote(ctx context.Context, params param.CreateEcsVSwitchRemoteParam) (*view.EcsVSwitchInventoryView, error) {
	resp := view.EcsVSwitchInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/vswitch", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmMigrationCandidateHosts gets VmMigrationCandidateHosts by uuid
func (cli *ZSClient) GetVmMigrationCandidateHosts(ctx context.Context, uuid string) (*view.GetVmMigrationCandidateHostsView, error) {
	var resp view.GetVmMigrationCandidateHostsView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForIpSecConnection gets CandidateL3NetworksForIpSecConnection by uuid
func (cli *ZSClient) GetCandidateL3NetworksForIpSecConnection(ctx context.Context) (*view.GetCandidateL3NetworksForIpSecConnectionView, error) {
	var resp view.GetCandidateL3NetworksForIpSecConnectionView
	if err := cli.GetWithRespKey(ctx, "v1/ipsec/candidatesL3Networks", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHostNetworkServiceType updates HostNetworkServiceType
func (cli *ZSClient) UpdateHostNetworkServiceType(ctx context.Context, uuid string, params param.UpdateHostNetworkServiceTypeParam) (*view.HostNetworkLabelInventoryView, error) {
	resp := view.HostNetworkLabelInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/service-types", uuid, "", map[string]interface{}{
		"updateHostNetworkServiceType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSMicrosoftTeamsTestConnection operates on MicrosoftTeamsTestConnection
func (cli *ZSClient) SNSMicrosoftTeamsTestConnection(ctx context.Context, params param.SNSMicrosoftTeamsTestConnectionParam) (*view.SNSMicrosoftTeamsTestConnectionEventView, error) {
	resp := view.SNSMicrosoftTeamsTestConnectionEventView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/microsoft-teams/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLatestGuestToolsForVm gets LatestGuestToolsForVm by uuid
func (cli *ZSClient) GetLatestGuestToolsForVm(ctx context.Context, uuid string) (*view.GuestToolsInventoryView, error) {
	var resp view.GetLatestGuestToolsForVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "inventory", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateVpcUserVpnGatewayRemote creates VpcUserVpnGatewayRemote
func (cli *ZSClient) CreateVpcUserVpnGatewayRemote(ctx context.Context, params param.CreateVpcUserVpnGatewayRemoteParam) (*view.VpcUserVpnGatewayInventoryView, error) {
	resp := view.VpcUserVpnGatewayInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/user-vpn", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOssBackupBucketRemote creates OssBackupBucketRemote
func (cli *ZSClient) CreateOssBackupBucketRemote(ctx context.Context, params param.CreateOssBackupBucketRemoteParam) (*view.CreateOssBackupBucketRemoteEventView, error) {
	resp := view.CreateOssBackupBucketRemoteEventView{}
	if err := cli.Post(ctx, "v1/hybrid/backup-mysql/oss", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOffBaremetalChassis operates on PowerOffBaremetalChassis
func (cli *ZSClient) PowerOffBaremetalChassis(ctx context.Context, chassisUuid string, params param.PowerOffBaremetalChassisParam) (*view.PowerOffBaremetalChassisEventView, error) {
	resp := view.PowerOffBaremetalChassisEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal/chassis", chassisUuid, "", map[string]interface{}{
		"powerOffBaremetalChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateInterfaceVlanIds gets CandidateInterfaceVlanIds by uuid
func (cli *ZSClient) GetCandidateInterfaceVlanIds(ctx context.Context) (*view.GetCandidateInterfaceVlanIdsView, error) {
	var resp view.GetCandidateInterfaceVlanIdsView
	if err := cli.GetWithRespKey(ctx, "v1/host/network-interface-vlan-ids", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNetworkServiceTypes gets NetworkServiceTypes by uuid
func (cli *ZSClient) GetNetworkServiceTypes(ctx context.Context) (*view.GetNetworkServiceTypesView, error) {
	var resp view.GetNetworkServiceTypesView
	if err := cli.GetWithRespKey(ctx, "v1/network-services/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmUserDefinedXml deletes VmUserDefinedXml
func (cli *ZSClient) DeleteVmUserDefinedXml(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// GetAvailableVpcL3Network gets AvailableVpcL3Network by uuid
func (cli *ZSClient) GetAvailableVpcL3Network(ctx context.Context) (*view.GetAvailableVpcL3NetworkView, error) {
	var resp view.GetAvailableVpcL3NetworkView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers/available-vpc-l3s", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCurrentTime gets CurrentTime by uuid
func (cli *ZSClient) GetCurrentTime(ctx context.Context) (*view.GetCurrentTimeView, error) {
	resp := view.GetCurrentTimeView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/actions", "", "", map[string]interface{}{
		"getCurrentTime": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateAccountSpending operates on AccountSpending
func (cli *ZSClient) CalculateAccountSpending(ctx context.Context, accountUuid string, params param.CalculateAccountSpendingParam) (*view.CalculateAccountSpendingView, error) {
	resp := view.CalculateAccountSpendingView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/accounts", accountUuid, "", map[string]interface{}{
		"calculateAccountSpending": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmAttachableL3Network gets VmAttachableL3Network by uuid
func (cli *ZSClient) GetVmAttachableL3Network(ctx context.Context, uuid string) (*view.GetVmAttachableL3NetworkView, error) {
	var resp view.GetVmAttachableL3NetworkView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEcsInstanceVncPassword updates EcsInstanceVncPassword
func (cli *ZSClient) UpdateEcsInstanceVncPassword(ctx context.Context, params param.UpdateEcsInstanceVncPasswordParam) (*view.UpdateEcsInstanceVncPasswordEventView, error) {
	resp := view.UpdateEcsInstanceVncPasswordEventView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/{uuid}/ecs-vnc", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncChronyServers operates on ChronyServers
func (cli *ZSClient) SyncChronyServers(ctx context.Context) (*view.SyncChronyServersEventView, error) {
	resp := view.SyncChronyServersEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zops/chrony/actions", "", "", map[string]interface{}{
		"syncChronyServers": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceProtectedRecoveryPoints gets VmInstanceProtectedRecoveryPoints by uuid
func (cli *ZSClient) GetVmInstanceProtectedRecoveryPoints(ctx context.Context, uuid string) (*view.GetVmInstanceProtectedRecoveryPointsView, error) {
	var resp view.GetVmInstanceProtectedRecoveryPointsView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmToVmSchedulingRuleGroup adds VmToVmSchedulingRuleGroup
func (cli *ZSClient) AddVmToVmSchedulingRuleGroup(ctx context.Context, params param.AddVmToVmSchedulingRuleGroupParam) (*view.AddVmToVmSchedulingRuleGroupEventView, error) {
	resp := view.AddVmToVmSchedulingRuleGroupEventView{}
	if err := cli.Post(ctx, "v1/vmSchedulingRuleGroup/{vmGroupUuid}/vmInstance/{vmUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncBackupFromImageStoreBackupStorage operates on BackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncBackupFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.SyncBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/volume-backups", uuid, "", map[string]interface{}{
		"syncBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostWebSshUrl gets HostWebSshUrl by uuid
func (cli *ZSClient) GetHostWebSshUrl(ctx context.Context, params param.GetHostWebSshUrlParam) (*view.GetHostWebSshUrlEventView, error) {
	resp := view.GetHostWebSshUrlEventView{}
	if err := cli.Post(ctx, "v1/hosts/webssh", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetL3NetworkMtu operates on L3NetworkMtu
func (cli *ZSClient) SetL3NetworkMtu(ctx context.Context, params param.SetL3NetworkMtuParam) (*view.SetL3NetworkMtuEventView, error) {
	resp := view.SetL3NetworkMtuEventView{}
	if err := cli.Post(ctx, "v1/l3-networks/{l3NetworkUuid}/mtu", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkRouterInterfaceIp gets L3NetworkRouterInterfaceIp by uuid
func (cli *ZSClient) GetL3NetworkRouterInterfaceIp(ctx context.Context, uuid string) (*view.GetL3NetworkRouterInterfaceIpView, error) {
	var resp view.GetL3NetworkRouterInterfaceIpView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVmClock operates on VmClock
func (cli *ZSClient) SyncVmClock(ctx context.Context, uuid string, params param.SyncVmClockParam) (*view.SyncVmClockEventView, error) {
	resp := view.SyncVmClockEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"syncVmClock": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcVpnConnectionFromLocal queries VpcVpnConnectionFromLocal list
func (cli *ZSClient) QueryVpcVpnConnectionFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVpnConnectionInventoryView, error) {
	var resp []view.VpcVpnConnectionInventoryView
	return resp, cli.List(ctx, "v1/hybrid/vpn-connection", params, &resp)
}

func (cli *ZSClient) GetVpcVpnConnectionFromLocal(ctx context.Context, uuid string) (*view.VpcVpnConnectionInventoryView, error) {
	var resp view.VpcVpnConnectionInventoryView
	if err := cli.Get(ctx, "v1/hybrid/vpn-connection", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcVpnConnectionFromLocal Pagination
func (cli *ZSClient) PageVpcVpnConnectionFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVpnConnectionInventoryView, int, error) {
	var vpcVpnConnectionFromLocals []view.VpcVpnConnectionInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/vpn-connection", params, &vpcVpnConnectionFromLocals)
	return vpcVpnConnectionFromLocals, total, err
}

// CreateSNSSnmpEndpoint creates SNSSnmpEndpoint
func (cli *ZSClient) CreateSNSSnmpEndpoint(ctx context.Context) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/snmp", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SdnControllerAddHost operates on SdnControllerAddHost
func (cli *ZSClient) SdnControllerAddHost(ctx context.Context, params param.SdnControllerAddHostParam) (*view.SdnControllerInventoryView, error) {
	resp := view.SdnControllerInventoryView{}
	if err := cli.Post(ctx, "v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAliyunSnapshotRemote creates AliyunSnapshotRemote
func (cli *ZSClient) CreateAliyunSnapshotRemote(ctx context.Context, params param.CreateAliyunSnapshotRemoteParam) (*view.AliyunSnapshotInventoryView, error) {
	resp := view.AliyunSnapshotInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/snapshot", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmBootVolume operates on VmBootVolume
func (cli *ZSClient) SetVmBootVolume(ctx context.Context, vmInstanceUuid string, params param.SetVmBootVolumeParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"setVmBootVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVpcHaGroupMonitorIps changes VpcHaGroupMonitorIps
func (cli *ZSClient) ChangeVpcHaGroupMonitorIps(ctx context.Context, uuid string, params param.ChangeVpcHaGroupMonitorIpsParam) (*view.VpcHaGroupInventoryView, error) {
	resp := view.VpcHaGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vpc/hagroups", uuid, "", map[string]interface{}{
		"changeVpcHaGroupMonitorIps": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewSession operates on RenewSession
func (cli *ZSClient) RenewSession(ctx context.Context, sessionUuid string, params param.RenewSessionParam) (*view.SessionInventoryView, error) {
	resp := view.SessionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/accounts/sessions", sessionUuid, "", map[string]interface{}{
		"renewSession": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDataCenterInLocal deletes DataCenterInLocal
func (cli *ZSClient) DeleteDataCenterInLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/data-center", uuid, string(deleteMode))
}

// SetVmConsoleMode operates on VmConsoleMode
func (cli *ZSClient) SetVmConsoleMode(ctx context.Context, uuid string, params param.SetVmConsoleModeParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmConsoleMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPolicyToUser operates on PolicyToUser
func (cli *ZSClient) AttachPolicyToUser(ctx context.Context, params param.AttachPolicyToUserParam) (*view.AttachPolicyToUserEventView, error) {
	resp := view.AttachPolicyToUserEventView{}
	if err := cli.Post(ctx, "v1/accounts/users/{userUuid}/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVfPciDeviceAvailableInL2Network gets VfPciDeviceAvailableInL2Network by uuid
func (cli *ZSClient) GetVfPciDeviceAvailableInL2Network(ctx context.Context) (*view.GetVfPciDeviceAvailableInL2NetworkView, error) {
	var resp view.GetVfPciDeviceAvailableInL2NetworkView
	if err := cli.GetWithRespKey(ctx, "v1/l2-networks/vf-pci-devices-available", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAttributesToIAM2Project adds AttributesToIAM2Project
func (cli *ZSClient) AddAttributesToIAM2Project(ctx context.Context, params param.AddAttributesToIAM2ProjectParam) (*view.AddAttributesToIAM2ProjectEventView, error) {
	resp := view.AddAttributesToIAM2ProjectEventView{}
	if err := cli.Post(ctx, "v1/iam2/projects/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UngenerateSeMdevDevices operates on UngenerateSeMdevDevices
func (cli *ZSClient) UngenerateSeMdevDevices(ctx context.Context, mttyDeviceUuid string, params param.UngenerateSeMdevDevicesParam) (*view.UngenerateSeMdevDevicesEventView, error) {
	resp := view.UngenerateSeMdevDevicesEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/mtty-devices", mttyDeviceUuid, "", map[string]interface{}{
		"ungenerateSeMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmEmulatorPinning operates on VmEmulatorPinning
func (cli *ZSClient) SetVmEmulatorPinning(ctx context.Context, uuid string, params param.SetVmEmulatorPinningParam) (*view.SetVmEmulatorPinningEventView, error) {
	resp := view.SetVmEmulatorPinningEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmEmulatorPinning": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanV2VConversionCache operates on V2VConversionCache
func (cli *ZSClient) CleanV2VConversionCache(ctx context.Context, uuid string, params param.CleanV2VConversionCacheParam) (*view.CleanV2VConversionCacheEventView, error) {
	resp := view.CleanV2VConversionCacheEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/v2v/conversion/host", uuid, "", map[string]interface{}{
		"cleanV2VConversionCache": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnbindModelFromService operates on UnbindModelFromService
func (cli *ZSClient) UnbindModelFromService(ctx context.Context, modelUuid string, modelServiceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/ai/models", modelUuid, fmt.Sprintf("model-services/%s", modelServiceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetEcsInstanceType gets EcsInstanceType by uuid
func (cli *ZSClient) GetEcsInstanceType(ctx context.Context) (*view.GetEcsInstanceTypeView, error) {
	var resp view.GetEcsInstanceTypeView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/ecs/type", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseUKeyStatus gets LicenseUKeyStatus by uuid
func (cli *ZSClient) GetLicenseUKeyStatus(ctx context.Context) (*view.UKeyInventoryView, error) {
	resp := view.UKeyInventoryView{}
	if err := cli.Post(ctx, "v1/licenses/actions", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddTicketTypesToTicketFlowCollection adds TicketTypesToTicketFlowCollection
func (cli *ZSClient) AddTicketTypesToTicketFlowCollection(ctx context.Context, params param.AddTicketTypesToTicketFlowCollectionParam) (*view.AddTicketTypesToTicketFlowCollectionEventView, error) {
	resp := view.AddTicketTypesToTicketFlowCollectionEventView{}
	if err := cli.Post(ctx, "v1/tickets/flow-collections/{ticketFlowCollectionUuid}/ticket-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetL3NetworkRouterInterfaceIp operates on L3NetworkRouterInterfaceIp
func (cli *ZSClient) SetL3NetworkRouterInterfaceIp(ctx context.Context, params param.SetL3NetworkRouterInterfaceIpParam) (*view.SetL3NetworkRouterInterfaceIpEventView, error) {
	resp := view.SetL3NetworkRouterInterfaceIpEventView{}
	if err := cli.Post(ctx, "v1/l3-networks/{l3NetworkUuid}/router-interface-ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEthernetVF queries EthernetVF list
func (cli *ZSClient) QueryEthernetVF(ctx context.Context, params *param.QueryParam) ([]view.EthernetVfPciDeviceInventoryView, error) {
	var resp []view.EthernetVfPciDeviceInventoryView
	return resp, cli.List(ctx, "v1/pci-device/ethernet-vfs", params, &resp)
}

func (cli *ZSClient) GetEthernetVF(ctx context.Context, uuid string) (*view.EthernetVfPciDeviceInventoryView, error) {
	var resp view.EthernetVfPciDeviceInventoryView
	if err := cli.Get(ctx, "v1/pci-device/ethernet-vfs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEthernetVF Pagination
func (cli *ZSClient) PageEthernetVF(ctx context.Context, params *param.QueryParam) ([]view.EthernetVfPciDeviceInventoryView, int, error) {
	var ethernetVFs []view.EthernetVfPciDeviceInventoryView
	total, err := cli.Page(ctx, "v1/pci-device/ethernet-vfs", params, &ethernetVFs)
	return ethernetVFs, total, err
}

// GetBareMetal2GatewayAllocatorStrategies gets BareMetal2GatewayAllocatorStrategies by uuid
func (cli *ZSClient) GetBareMetal2GatewayAllocatorStrategies(ctx context.Context) (*view.GetBareMetal2GatewayAllocatorStrategiesView, error) {
	var resp view.GetBareMetal2GatewayAllocatorStrategiesView
	if err := cli.GetWithRespKey(ctx, "v1/baremetal2/gateways/allocators/strategies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsSecurityGroupFromLocal queries EcsSecurityGroupFromLocal list
func (cli *ZSClient) QueryEcsSecurityGroupFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsSecurityGroupInventoryView, error) {
	var resp []view.EcsSecurityGroupInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/security-group", params, &resp)
}

func (cli *ZSClient) GetEcsSecurityGroupFromLocal(ctx context.Context, uuid string) (*view.EcsSecurityGroupInventoryView, error) {
	var resp view.EcsSecurityGroupInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/security-group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsSecurityGroupFromLocal Pagination
func (cli *ZSClient) PageEcsSecurityGroupFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsSecurityGroupInventoryView, int, error) {
	var ecsSecurityGroupFromLocals []view.EcsSecurityGroupInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/security-group", params, &ecsSecurityGroupFromLocals)
	return ecsSecurityGroupFromLocals, total, err
}

// UpdateFirewallRuleTemplate updates FirewallRuleTemplate
func (cli *ZSClient) UpdateFirewallRuleTemplate(ctx context.Context, uuid string, params param.UpdateFirewallRuleTemplateParam) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	resp := view.VpcFirewallRuleTemplateInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vpcfirewalls/rules/template", uuid, "", map[string]interface{}{
		"updateFirewallRuleTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUsbDeviceCandidatesForAttachingVm gets UsbDeviceCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetUsbDeviceCandidatesForAttachingVm(ctx context.Context, uuid string) (*view.GetUsbDeviceCandidatesForAttachingVmView, error) {
	var resp view.GetUsbDeviceCandidatesForAttachingVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForLoadBalancer gets CandidateL3NetworksForLoadBalancer by uuid
func (cli *ZSClient) GetCandidateL3NetworksForLoadBalancer(ctx context.Context, uuid string) (*view.GetCandidateL3NetworksForLoadBalancerView, error) {
	var resp view.GetCandidateL3NetworksForLoadBalancerView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers/listeners", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// WithdrawLicenseCapacityApplication operates on WithdrawLicenseCapacityApplication
func (cli *ZSClient) WithdrawLicenseCapacityApplication(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/license-server/capacity-application", uuid, string(deleteMode))
}

// PowerResetHost operates on PowerResetHost
func (cli *ZSClient) PowerResetHost(ctx context.Context, uuid string, params param.PowerResetHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/power", uuid, "", map[string]interface{}{
		"powerResetHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallRule queries FirewallRule list
func (cli *ZSClient) QueryFirewallRule(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleInventoryView, error) {
	var resp []view.VpcFirewallRuleInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/rules", params, &resp)
}

func (cli *ZSClient) GetFirewallRule(ctx context.Context, uuid string) (*view.VpcFirewallRuleInventoryView, error) {
	var resp view.VpcFirewallRuleInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls/rules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRule Pagination
func (cli *ZSClient) PageFirewallRule(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleInventoryView, int, error) {
	var firewallRules []view.VpcFirewallRuleInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/rules", params, &firewallRules)
	return firewallRules, total, err
}

// RevertVmFromVmBackup operates on VmFromVmBackup
func (cli *ZSClient) RevertVmFromVmBackup(ctx context.Context, groupUuid string, params param.RevertVmFromVmBackupParam) (*view.RevertVmFromVmBackupEventView, error) {
	resp := view.RevertVmFromVmBackupEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-backups", groupUuid, "", map[string]interface{}{
		"revertVmFromVmBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachCCSCertificateToUser operates on CCSCertificateToUser
func (cli *ZSClient) AttachCCSCertificateToUser(ctx context.Context, params param.AttachCCSCertificateToUserParam) (*view.CCSCertificateInventoryView, error) {
	resp := view.CCSCertificateInventoryView{}
	if err := cli.Post(ctx, "v1/crypto/ccs-certificate/attach-user/{userUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryOssBucketFileName queries OssBucketFileName list
func (cli *ZSClient) QueryOssBucketFileName(ctx context.Context, params *param.QueryParam) ([]view.OssBucketInventoryView, error) {
	var resp []view.OssBucketInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/oss-bucket", params, &resp)
}

func (cli *ZSClient) GetOssBucketFileName(ctx context.Context, uuid string) (*view.OssBucketInventoryView, error) {
	var resp view.OssBucketInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/oss-bucket", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageOssBucketFileName Pagination
func (cli *ZSClient) PageOssBucketFileName(ctx context.Context, params *param.QueryParam) ([]view.OssBucketInventoryView, int, error) {
	var ossBucketFileNames []view.OssBucketInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/oss-bucket", params, &ossBucketFileNames)
	return ossBucketFileNames, total, err
}

// SetVmNuma operates on VmNuma
func (cli *ZSClient) SetVmNuma(ctx context.Context, uuid string, params param.SetVmNumaParam) (*view.SetVmNumaEventView, error) {
	resp := view.SetVmNumaEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmNuma": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVRouterOspfArea queries VRouterOspfArea list
func (cli *ZSClient) QueryVRouterOspfArea(ctx context.Context, params *param.QueryParam) ([]view.RouterAreaInventoryView, error) {
	var resp []view.RouterAreaInventoryView
	return resp, cli.List(ctx, "v1/routerArea", params, &resp)
}

func (cli *ZSClient) GetVRouterOspfArea(ctx context.Context, uuid string) (*view.RouterAreaInventoryView, error) {
	var resp view.RouterAreaInventoryView
	if err := cli.Get(ctx, "v1/routerArea", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterOspfArea Pagination
func (cli *ZSClient) PageVRouterOspfArea(ctx context.Context, params *param.QueryParam) ([]view.RouterAreaInventoryView, int, error) {
	var vRouterOspfAreas []view.RouterAreaInventoryView
	total, err := cli.Page(ctx, "v1/routerArea", params, &vRouterOspfAreas)
	return vRouterOspfAreas, total, err
}

// DeleteAliyunRouterInterfaceLocal deletes AliyunRouterInterfaceLocal
func (cli *ZSClient) DeleteAliyunRouterInterfaceLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/router-interface", uuid, string(deleteMode))
}

// UpdateFirewallRuleSet updates FirewallRuleSet
func (cli *ZSClient) UpdateFirewallRuleSet(ctx context.Context, uuid string, params param.UpdateFirewallRuleSetParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vpcfirewalls/ruleSets", uuid, "", map[string]interface{}{
		"updateFirewallRuleSet": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachAliyunKey operates on AliyunKey
func (cli *ZSClient) AttachAliyunKey(ctx context.Context, uuid string, params param.AttachAliyunKeyParam) (*view.AttachAliyunKeyEventView, error) {
	resp := view.AttachAliyunKeyEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/key", uuid, "", map[string]interface{}{
		"attachAliyunKey": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshSearchIndexes operates on SearchIndexes
func (cli *ZSClient) RefreshSearchIndexes(ctx context.Context) (*view.RefreshSearchIndexesView, error) {
	var resp view.RefreshSearchIndexesView
	if err := cli.GetWithRespKey(ctx, "v1/search/indexes/refresh", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateImageHash operates on ImageHash
func (cli *ZSClient) CalculateImageHash(ctx context.Context, uuid string, params param.CalculateImageHashParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/images", uuid, "", map[string]interface{}{
		"calculateImageHash": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL2NetworkTypes gets L2NetworkTypes by uuid
func (cli *ZSClient) GetL2NetworkTypes(ctx context.Context) (*view.GetL2NetworkTypesView, error) {
	var resp view.GetL2NetworkTypesView
	if err := cli.GetWithRespKey(ctx, "v1/l2-networks/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShutdownHost operates on ShutdownHost
func (cli *ZSClient) ShutdownHost(ctx context.Context, uuid string, params param.ShutdownHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/power", uuid, "", map[string]interface{}{
		"shutdownHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVpcVpnConnectionRemote updates VpcVpnConnectionRemote
func (cli *ZSClient) UpdateVpcVpnConnectionRemote(ctx context.Context, uuid string, params param.UpdateVpcVpnConnectionRemoteParam) (*view.VpcVpnConnectionInventoryView, error) {
	resp := view.VpcVpnConnectionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/vpn-connection", uuid, "", map[string]interface{}{
		"updateVpcVpnConnectionRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmTask gets VmTask by uuid
func (cli *ZSClient) GetVmTask(ctx context.Context) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/task-details", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableCdpTask operates on DisableCdpTask
func (cli *ZSClient) DisableCdpTask(ctx context.Context, params param.DisableCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.Post(ctx, "v1/cdp-task/disable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIpOnHostNetworkBonding operates on IpOnHostNetworkBonding
func (cli *ZSClient) SetIpOnHostNetworkBonding(ctx context.Context, params param.SetIpOnHostNetworkBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.Post(ctx, "v1/hosts/bondings/{bondingUuid}/ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveAttributesFromIAM2VirtualID removes AttributesFromIAM2VirtualID
func (cli *ZSClient) RemoveAttributesFromIAM2VirtualID(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/virtual-ids", uuid, string(deleteMode))
}

// CreateBonding creates Bonding
func (cli *ZSClient) CreateBonding(ctx context.Context, params param.CreateBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.Post(ctx, "v1/hosts/bondings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachUsbDeviceFromVm operates on UsbDeviceFromVm
func (cli *ZSClient) DetachUsbDeviceFromVm(ctx context.Context, params param.DetachUsbDeviceFromVmParam) (*view.UsbDeviceInventoryView, error) {
	resp := view.UsbDeviceInventoryView{}
	if err := cli.Post(ctx, "v1/usb-device/usb-devices/{usbDeviceUuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolumeSnapshot creates DataVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeSnapshot(ctx context.Context, params param.CreateDataVolumeTemplateFromVolumeSnapshotParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(ctx, "v1/images/data-volume-templates/from/volume-snapshots/{snapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachRoleFromAccount operates on RoleFromAccount
func (cli *ZSClient) DetachRoleFromAccount(ctx context.Context, accountUuid string, roleUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/identities/accounts", accountUuid, fmt.Sprintf("roles/%s", roleUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddRendezvousPointToMulticastRouter adds RendezvousPointToMulticastRouter
func (cli *ZSClient) AddRendezvousPointToMulticastRouter(ctx context.Context, params param.AddRendezvousPointToMulticastRouterParam) (*view.MulticastRouterInventoryView, error) {
	resp := view.MulticastRouterInventoryView{}
	if err := cli.Post(ctx, "v1/multicast/virtual-routers/{uuid}/RendezvousPoint", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QuerySNSTopicSubscriber queries SNSTopicSubscriber list
func (cli *ZSClient) QuerySNSTopicSubscriber(ctx context.Context, params *param.QueryParam) ([]view.SNSSubscriberInventoryView, error) {
	var resp []view.SNSSubscriberInventoryView
	return resp, cli.List(ctx, "v1/sns/topics/subscribers", params, &resp)
}

func (cli *ZSClient) GetSNSTopicSubscriber(ctx context.Context, uuid string) (*view.SNSSubscriberInventoryView, error) {
	var resp view.SNSSubscriberInventoryView
	if err := cli.Get(ctx, "v1/sns/topics/subscribers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSTopicSubscriber Pagination
func (cli *ZSClient) PageSNSTopicSubscriber(ctx context.Context, params *param.QueryParam) ([]view.SNSSubscriberInventoryView, int, error) {
	var sNSTopicSubscribers []view.SNSSubscriberInventoryView
	total, err := cli.Page(ctx, "v1/sns/topics/subscribers", params, &sNSTopicSubscribers)
	return sNSTopicSubscribers, total, err
}

// DeleteLdapBinding deletes LdapBinding
func (cli *ZSClient) DeleteLdapBinding(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ldap/bindings", uuid, string(deleteMode))
}

// DebugSignal operates on DebugSignal
func (cli *ZSClient) DebugSignal(ctx context.Context, params param.DebugSignalParam) (*view.DebugSignalEventView, error) {
	resp := view.DebugSignalEventView{}
	if err := cli.Post(ctx, "v1/debug", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromVolume creates VmInstanceFromVolume
func (cli *ZSClient) CreateVmInstanceFromVolume(ctx context.Context, params param.CreateVmInstanceFromVolumeParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/vm-instances/from/volume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcVRouterDistributedRoutingEnabled gets VpcVRouterDistributedRoutingEnabled by uuid
func (cli *ZSClient) GetVpcVRouterDistributedRoutingEnabled(ctx context.Context, uuid string) (*view.GetVpcVRouterDistributedRoutingEnabledView, error) {
	var resp view.GetVpcVRouterDistributedRoutingEnabledView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsSecurityGroupRemote creates EcsSecurityGroupRemote
func (cli *ZSClient) CreateEcsSecurityGroupRemote(ctx context.Context, params param.CreateEcsSecurityGroupRemoteParam) (*view.EcsSecurityGroupInventoryView, error) {
	resp := view.EcsSecurityGroupInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/security-group/remote", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveAttributesFromIAM2Organization removes AttributesFromIAM2Organization
func (cli *ZSClient) RemoveAttributesFromIAM2Organization(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/organizations", uuid, string(deleteMode))
}

// DeleteAliyunSnapshotFromLocal deletes AliyunSnapshotFromLocal
func (cli *ZSClient) DeleteAliyunSnapshotFromLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/snapshot", uuid, string(deleteMode))
}

// GetIAM2ProjectContainerImages gets IAM2ProjectContainerImages by uuid
func (cli *ZSClient) GetIAM2ProjectContainerImages(ctx context.Context, projectId string, repositoryId string) (*view.GetIAM2ProjectContainerImagesView, error) {
	var resp view.GetIAM2ProjectContainerImagesView
	err := cli.GetWithSpec(ctx, "v1/iam2/project", projectId, fmt.Sprintf("repository/%s/image", repositoryId), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachDataVolumeFromVm operates on DataVolumeFromVm
func (cli *ZSClient) DetachDataVolumeFromVm(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volumes", uuid, string(deleteMode))
}

// QueryEcsVSwitchFromLocal queries EcsVSwitchFromLocal list
func (cli *ZSClient) QueryEcsVSwitchFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsVSwitchInventoryView, error) {
	var resp []view.EcsVSwitchInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/vswitch", params, &resp)
}

func (cli *ZSClient) GetEcsVSwitchFromLocal(ctx context.Context, uuid string) (*view.EcsVSwitchInventoryView, error) {
	var resp view.EcsVSwitchInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/vswitch", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsVSwitchFromLocal Pagination
func (cli *ZSClient) PageEcsVSwitchFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsVSwitchInventoryView, int, error) {
	var ecsVSwitchFromLocals []view.EcsVSwitchInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/vswitch", params, &ecsVSwitchFromLocals)
	return ecsVSwitchFromLocals, total, err
}

// CreateRootVolumeTemplateFromRootVolume creates RootVolumeTemplateFromRootVolume
func (cli *ZSClient) CreateRootVolumeTemplateFromRootVolume(ctx context.Context, params param.CreateRootVolumeTemplateFromRootVolumeParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(ctx, "v1/images/root-volume-templates/from/volumes/{rootVolumeUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromRootVolumeAsync Async
func (cli *ZSClient) CreateRootVolumeTemplateFromRootVolumeAsync(ctx context.Context, params param.CreateRootVolumeTemplateFromRootVolumeParam) (string, error) {

	resource := "v1/images/root-volume-templates/from/volumes/{rootVolumeUuid}"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// AttachAliyunDiskToEcs operates on AliyunDiskToEcs
func (cli *ZSClient) AttachAliyunDiskToEcs(ctx context.Context, params param.AttachAliyunDiskToEcsParam) (*view.AliyunDiskInventoryView, error) {
	resp := view.AliyunDiskInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/disk/{diskUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOssBucketNameLocal deletes OssBucketNameLocal
func (cli *ZSClient) DeleteOssBucketNameLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/oss-bucket", uuid, string(deleteMode))
}

// QueryEcsImageFromLocal queries EcsImageFromLocal list
func (cli *ZSClient) QueryEcsImageFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsImageInventoryView, error) {
	var resp []view.EcsImageInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/image", params, &resp)
}

func (cli *ZSClient) GetEcsImageFromLocal(ctx context.Context, uuid string) (*view.EcsImageInventoryView, error) {
	var resp view.EcsImageInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/image", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsImageFromLocal Pagination
func (cli *ZSClient) PageEcsImageFromLocal(ctx context.Context, params *param.QueryParam) ([]view.EcsImageInventoryView, int, error) {
	var ecsImageFromLocals []view.EcsImageInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/image", params, &ecsImageFromLocals)
	return ecsImageFromLocals, total, err
}

// GetObservabilityServerServiceData gets ObservabilityServerServiceData by uuid
func (cli *ZSClient) GetObservabilityServerServiceData(ctx context.Context, params param.GetObservabilityServerServiceDataParam) (*view.ObservabilityServerServiceDataInventoryView, error) {
	resp := view.ObservabilityServerServiceDataInventoryView{}
	if err := cli.Post(ctx, "v1/observability-server/{observabilityServerUuid}/service-data", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunVirtualRouterFromLocal queries AliyunVirtualRouterFromLocal list
func (cli *ZSClient) QueryAliyunVirtualRouterFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVirtualRouterInventoryView, error) {
	var resp []view.VpcVirtualRouterInventoryView
	return resp, cli.List(ctx, "v1/hybrid/aliyun/vrouter", params, &resp)
}

func (cli *ZSClient) GetAliyunVirtualRouterFromLocal(ctx context.Context, uuid string) (*view.VpcVirtualRouterInventoryView, error) {
	var resp view.VpcVirtualRouterInventoryView
	if err := cli.Get(ctx, "v1/hybrid/aliyun/vrouter", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunVirtualRouterFromLocal Pagination
func (cli *ZSClient) PageAliyunVirtualRouterFromLocal(ctx context.Context, params *param.QueryParam) ([]view.VpcVirtualRouterInventoryView, int, error) {
	var aliyunVirtualRouterFromLocals []view.VpcVirtualRouterInventoryView
	total, err := cli.Page(ctx, "v1/hybrid/aliyun/vrouter", params, &aliyunVirtualRouterFromLocals)
	return aliyunVirtualRouterFromLocals, total, err
}

// VerifyLicenseServer operates on VerifyLicenseServer
func (cli *ZSClient) VerifyLicenseServer(ctx context.Context, params param.VerifyLicenseServerParam) (*view.VerifyLicenseServerEventView, error) {
	resp := view.VerifyLicenseServerEventView{}
	if err := cli.Post(ctx, "v1/license-server/register-verify", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachBareMetal2GatewayToCluster operates on BareMetal2GatewayToCluster
func (cli *ZSClient) AttachBareMetal2GatewayToCluster(ctx context.Context, params param.AttachBareMetal2GatewayToClusterParam) (*view.BareMetal2GatewayInventoryView, error) {
	resp := view.BareMetal2GatewayInventoryView{}
	if err := cli.Post(ctx, "v1/baremetal2/clusters/{clusterUuid}/gateways/{gatewayUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAtPersonOfAtWeComEndpoint updates AtPersonOfAtWeComEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtWeComEndpoint(ctx context.Context, uuid string, params param.UpdateAtPersonOfAtWeComEndpointParam) (*view.SNSWeComAtPersonInventoryView, error) {
	resp := view.SNSWeComAtPersonInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-endpoints/we-com/at-persons", uuid, "", map[string]interface{}{
		"updateAtPersonOfAtWeComEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSlbGroupDeployType changes SlbGroupDeployType
func (cli *ZSClient) ChangeSlbGroupDeployType(ctx context.Context, slbGroupUuid string, params param.ChangeSlbGroupDeployTypeParam) (*view.SlbGroupInventoryView, error) {
	resp := view.SlbGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/load-balancers/slb/groups", slbGroupUuid, "", map[string]interface{}{
		"changeSlbGroupDeployType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsSecurityGroupInLocal deletes EcsSecurityGroupInLocal
func (cli *ZSClient) DeleteEcsSecurityGroupInLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/security-group", uuid, string(deleteMode))
}

// DetachDataVolumeFromHost operates on DataVolumeFromHost
func (cli *ZSClient) DetachDataVolumeFromHost(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volumes", uuid, string(deleteMode))
}

// GetVmInstanceRecoveryPoints gets VmInstanceRecoveryPoints by uuid
func (cli *ZSClient) GetVmInstanceRecoveryPoints(ctx context.Context, uuid string) (*view.GetVmInstanceRecoveryPointsView, error) {
	var resp view.GetVmInstanceRecoveryPointsView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSystemTags creates SystemTags
func (cli *ZSClient) CreateSystemTags(ctx context.Context, params param.CreateSystemTagsParam) (*view.SystemTagInventoryView, error) {
	resp := view.SystemTagInventoryView{}
	if err := cli.Post(ctx, "v1/system-tags/{resourceUuid}/tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOssBackupBucketFromRemote gets OssBackupBucketFromRemote by uuid
func (cli *ZSClient) GetOssBackupBucketFromRemote(ctx context.Context) (*view.GetOssBackupBucketFromRemoteView, error) {
	var resp view.GetOssBackupBucketFromRemoteView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/backup-mysql/oss", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkTypes gets L3NetworkTypes by uuid
func (cli *ZSClient) GetL3NetworkTypes(ctx context.Context) (*view.GetL3NetworkTypesView, error) {
	var resp view.GetL3NetworkTypesView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPoliciesFromUser operates on PoliciesFromUser
func (cli *ZSClient) DetachPoliciesFromUser(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/accounts/users", uuid, string(deleteMode))
}

// CleanUpImageCacheOnPrimaryStorage operates on UpImageCacheOnPrimaryStorage
func (cli *ZSClient) CleanUpImageCacheOnPrimaryStorage(ctx context.Context, uuid string, params param.CleanUpImageCacheOnPrimaryStorageParam) (*view.CleanUpImageCacheOnPrimaryStorageEventView, error) {
	resp := view.CleanUpImageCacheOnPrimaryStorageEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage", uuid, "", map[string]interface{}{
		"cleanUpImageCacheOnPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddKVMHostFromConfigFile adds KVMHostFromConfigFile
func (cli *ZSClient) AddKVMHostFromConfigFile(ctx context.Context) (*view.AddHostFromConfigFileEventView, error) {
	resp := view.AddHostFromConfigFileEventView{}
	if err := cli.Post(ctx, "v1/hosts/kvm/from-file", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddKVMHostFromConfigFileAsync Async
func (cli *ZSClient) AddKVMHostFromConfigFileAsync(ctx context.Context, params param.AddKVMHostFromConfigFileParam) (string, error) {

	resource := "v1/hosts/kvm/from-file"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// InspectBareMetal2ChassisByInstance operates on BareMetal2ChassisByInstance
func (cli *ZSClient) InspectBareMetal2ChassisByInstance(ctx context.Context, uuid string, params param.InspectBareMetal2ChassisByInstanceParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/bm-instances", uuid, "", map[string]interface{}{
		"inspectBareMetal2ChassisByInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmBootMode deletes VmBootMode
func (cli *ZSClient) DeleteVmBootMode(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// GetCandidateVMForAttachingAffinityGroup gets CandidateVMForAttachingAffinityGroup by uuid
func (cli *ZSClient) GetCandidateVMForAttachingAffinityGroup(ctx context.Context) (*view.GetCandidateVMForAttachingAffinityGroupView, error) {
	var resp view.GetCandidateVMForAttachingAffinityGroupView
	if err := cli.GetWithRespKey(ctx, "v1/VM/attachingGroup", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcVpnConnectionLocal deletes VpcVpnConnectionLocal
func (cli *ZSClient) DeleteVpcVpnConnectionLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/vpn-connection", uuid, string(deleteMode))
}

// DetachPolicyFromUserGroup operates on PolicyFromUserGroup
func (cli *ZSClient) DetachPolicyFromUserGroup(ctx context.Context, groupUuid string, policyUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/accounts/groups", groupUuid, fmt.Sprintf("policies/%s", policyUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddActionToAlarm adds ActionToAlarm
func (cli *ZSClient) AddActionToAlarm(ctx context.Context, params param.AddActionToAlarmParam) (*view.AlarmInventoryView, error) {
	resp := view.AlarmInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/alarms/{alarmUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFirewallRule updates FirewallRule
func (cli *ZSClient) UpdateFirewallRule(ctx context.Context, uuid string, params param.UpdateFirewallRuleParam) (*view.VpcFirewallRuleInventoryView, error) {
	resp := view.VpcFirewallRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vpcfirewalls/rules", uuid, "", map[string]interface{}{
		"updateFirewallRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZQLQuery operates on ZQLQuery
func (cli *ZSClient) ZQLQuery(ctx context.Context) (*view.ZQLQueryView, error) {
	var resp view.ZQLQueryView
	if err := cli.GetWithRespKey(ctx, "v1/zql", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetElaborations gets Elaborations by uuid
func (cli *ZSClient) GetElaborations(ctx context.Context) (*view.GetElaborationsView, error) {
	var resp view.GetElaborationsView
	if err := cli.GetWithRespKey(ctx, "v1/errorcode/elaborations", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAccessPath gets AccessPath by uuid
func (cli *ZSClient) GetAccessPath(ctx context.Context) (*view.GetAccessPathView, error) {
	var resp view.GetAccessPathView
	if err := cli.GetWithRespKey(ctx, "v1/block-volumes/access/path", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageUsageReport gets PrimaryStorageUsageReport by uuid
func (cli *ZSClient) GetPrimaryStorageUsageReport(ctx context.Context, uuid string) (*view.GetPrimaryStorageUsageReportView, error) {
	var resp view.GetPrimaryStorageUsageReportView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVolumeFromVolumeBackup operates on VolumeFromVolumeBackup
func (cli *ZSClient) RevertVolumeFromVolumeBackup(ctx context.Context, uuid string, params param.RevertVolumeFromVolumeBackupParam) (*view.RevertVolumeFromVolumeBackupEventView, error) {
	resp := view.RevertVolumeFromVolumeBackupEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/volume-backups", uuid, "", map[string]interface{}{
		"revertVolumeFromVolumeBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeFromVolumeTemplate creates DataVolumeFromVolumeTemplate
func (cli *ZSClient) CreateDataVolumeFromVolumeTemplate(ctx context.Context, params param.CreateDataVolumeFromVolumeTemplateParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.Post(ctx, "v1/volumes/data/from/data-volume-templates/{imageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LocalStorageGetVolumeMigratableHosts operates on LocalStorageGetVolumeMigratableHosts
func (cli *ZSClient) LocalStorageGetVolumeMigratableHosts(ctx context.Context, uuid string) (*view.LocalStorageGetVolumeMigratableView, error) {
	var resp view.LocalStorageGetVolumeMigratableView
	if err := cli.GetWithRespKey(ctx, "v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOssBucketNameFromRemote gets OssBucketNameFromRemote by uuid
func (cli *ZSClient) GetOssBucketNameFromRemote(ctx context.Context, uuid string) (*view.GetOssBucketNameFromRemoteView, error) {
	var resp view.GetOssBucketNameFromRemoteView
	if err := cli.GetWithRespKey(ctx, "v1/hybrid/oss", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsVpcFromRemote operates on EcsVpcFromRemote
func (cli *ZSClient) SyncEcsVpcFromRemote(ctx context.Context, params param.SyncEcsVpcFromRemoteParam) (*view.EcsVpcInventoryView, error) {
	resp := view.EcsVpcInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/vpc/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetServiceTypeOnHostNetworkInterface operates on ServiceTypeOnHostNetworkInterface
func (cli *ZSClient) SetServiceTypeOnHostNetworkInterface(ctx context.Context, params param.SetServiceTypeOnHostNetworkInterfaceParam) (*view.HostNetworkInterfaceServiceRefInventoryView, error) {
	resp := view.HostNetworkInterfaceServiceRefInventoryView{}
	if err := cli.Post(ctx, "v1/hosts/nics/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddBackendServerToServerGroup adds BackendServerToServerGroup
func (cli *ZSClient) AddBackendServerToServerGroup(ctx context.Context, params param.AddBackendServerToServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.Post(ctx, "v1/load-balancers/servergroups/{serverGroupUuid}/backendservers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachUserDefinedXmlHookScriptToVm operates on UserDefinedXmlHookScriptToVm
func (cli *ZSClient) AttachUserDefinedXmlHookScriptToVm(ctx context.Context, params param.AttachUserDefinedXmlHookScriptToVmParam) (*view.AttachUserDefinedXmlHookScriptToVmEventView, error) {
	resp := view.AttachUserDefinedXmlHookScriptToVmEventView{}
	if err := cli.Post(ctx, "v1/xmlhook/{xmlHookUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPolicyToRole operates on PolicyToRole
func (cli *ZSClient) AttachPolicyToRole(ctx context.Context, params param.AttachPolicyToRoleParam) (*view.AttachPolicyToRoleEventView, error) {
	resp := view.AttachPolicyToRoleEventView{}
	if err := cli.Post(ctx, "v1/identities/policies/{policyUuid}/roles/{roleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2ProvisionNetworkState changes BareMetal2ProvisionNetworkState
func (cli *ZSClient) ChangeBareMetal2ProvisionNetworkState(ctx context.Context, uuid string, params param.ChangeBareMetal2ProvisionNetworkStateParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	resp := view.BareMetal2ProvisionNetworkInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/provision-networks", uuid, "", map[string]interface{}{
		"changeBareMetal2ProvisionNetworkState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBackupStorageCandidatesForImageMigration gets BackupStorageCandidatesForImageMigration by uuid
func (cli *ZSClient) GetBackupStorageCandidatesForImageMigration(ctx context.Context, uuid string) (*view.GetBackupStorageCandidatesForImageMigrationView, error) {
	var resp view.GetBackupStorageCandidatesForImageMigrationView
	if err := cli.GetWithRespKey(ctx, "v1/backup-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcIpSecConfigLocal deletes VpcIpSecConfigLocal
func (cli *ZSClient) DeleteVpcIpSecConfigLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/vpn-connection/ipsec", uuid, string(deleteMode))
}

// GenerateSriovPciDevices operates on SriovPciDevices
func (cli *ZSClient) GenerateSriovPciDevices(ctx context.Context, pciDeviceUuid string, params param.GenerateSriovPciDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/pci-devices", pciDeviceUuid, "", map[string]interface{}{
		"generateSriovPciDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateAccountBillingSpending operates on AccountBillingSpending
func (cli *ZSClient) CalculateAccountBillingSpending(ctx context.Context, accountUuid string, params param.CalculateAccountBillingSpendingParam) (*view.CalculateAccountBillingSpendingView, error) {
	resp := view.CalculateAccountBillingSpendingView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/accounts", accountUuid, "", map[string]interface{}{
		"calculateAccountBillingSpending": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVRouterOspfArea deletes VRouterOspfArea
func (cli *ZSClient) DeleteVRouterOspfArea(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/routerArea", uuid, string(deleteMode))
}

// GetVipAvailablePort gets VipAvailablePort by uuid
func (cli *ZSClient) GetVipAvailablePort(ctx context.Context, uuid string) (*view.GetVipAvailablePortView, error) {
	var resp view.GetVipAvailablePortView
	if err := cli.GetWithRespKey(ctx, "v1/vips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncDiskFromAliyunFromRemote operates on DiskFromAliyunFromRemote
func (cli *ZSClient) SyncDiskFromAliyunFromRemote(ctx context.Context, params param.SyncDiskFromAliyunFromRemoteParam) (*view.AliyunDiskInventoryView, error) {
	resp := view.AliyunDiskInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/disk/{identityUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVolumeState changes VolumeState
func (cli *ZSClient) ChangeVolumeState(ctx context.Context, uuid string, params param.ChangeVolumeStateParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/volumes", uuid, "", map[string]interface{}{
		"changeVolumeState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MountVmInstanceRecoveryPoint operates on MountVmInstanceRecoveryPoint
func (cli *ZSClient) MountVmInstanceRecoveryPoint(ctx context.Context, params param.MountVmInstanceRecoveryPointParam) (*view.MountVmInstanceRecoveryPointEventView, error) {
	resp := view.MountVmInstanceRecoveryPointEventView{}
	if err := cli.Post(ctx, "v1/cdp-backup-storage/mount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVxlanPoolRemoteVtep creates VxlanPoolRemoteVtep
func (cli *ZSClient) CreateVxlanPoolRemoteVtep(ctx context.Context, params param.CreateVxlanPoolRemoteVtepParam) (*view.RemoteVtepInventoryView, error) {
	resp := view.RemoteVtepInventoryView{}
	if err := cli.Post(ctx, "v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}/remote-vtep-ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceStackFromResource gets ResourceStackFromResource by uuid
func (cli *ZSClient) GetResourceStackFromResource(ctx context.Context) (*view.GetResourceStackFromResourceView, error) {
	var resp view.GetResourceStackFromResourceView
	if err := cli.GetWithRespKey(ctx, "v1/cloudformation/resources/stack", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIAM2VirtualIDConfigFile operates on IAM2VirtualIDConfigFile
func (cli *ZSClient) CheckIAM2VirtualIDConfigFile(ctx context.Context, params param.CheckIAM2VirtualIDConfigFileParam) (*view.CheckIAM2VirtualIDConfigFileView, error) {
	resp := view.CheckIAM2VirtualIDConfigFileView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/virtual-ids/from-file", "", "", map[string]interface{}{
		"checkIAM2VirtualIDConfigFile": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetClusterHostNetworkFacts gets ClusterHostNetworkFacts by uuid
func (cli *ZSClient) GetClusterHostNetworkFacts(ctx context.Context, uuid string) (*view.GetClusterHostNetworkFactsView, error) {
	var resp view.GetClusterHostNetworkFactsView
	if err := cli.GetWithRespKey(ctx, "v1/cluster/hosts-network-facts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachOssBucketFromEcsDataCenter operates on OssBucketFromEcsDataCenter
func (cli *ZSClient) DetachOssBucketFromEcsDataCenter(ctx context.Context, ossBucketUuid string, params param.DetachOssBucketFromEcsDataCenterParam) (*view.OssBucketInventoryView, error) {
	resp := view.OssBucketInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/oss-bucket", ossBucketUuid, "", map[string]interface{}{
		"detachOssBucketFromEcsDataCenter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ParseOvf operates on ParseOvf
func (cli *ZSClient) ParseOvf(ctx context.Context, params param.ParseOvfParam) (*view.ParseOvfView, error) {
	resp := view.ParseOvfView{}
	if err := cli.Post(ctx, "v1/ovf/parse", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFaultToleranceVm queries FaultToleranceVm list
func (cli *ZSClient) QueryFaultToleranceVm(ctx context.Context, params *param.QueryParam) ([]view.FaultToleranceVmGroupInventoryView, error) {
	var resp []view.FaultToleranceVmGroupInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/fault-tolerance", params, &resp)
}

func (cli *ZSClient) GetFaultToleranceVm(ctx context.Context, uuid string) (*view.FaultToleranceVmGroupInventoryView, error) {
	var resp view.FaultToleranceVmGroupInventoryView
	if err := cli.Get(ctx, "v1/vm-instances/fault-tolerance", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFaultToleranceVm Pagination
func (cli *ZSClient) PageFaultToleranceVm(ctx context.Context, params *param.QueryParam) ([]view.FaultToleranceVmGroupInventoryView, int, error) {
	var faultToleranceVms []view.FaultToleranceVmGroupInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/fault-tolerance", params, &faultToleranceVms)
	return faultToleranceVms, total, err
}

// AddSchedulerJobGroupToSchedulerTrigger adds SchedulerJobGroupToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobGroupToSchedulerTrigger(ctx context.Context, params param.AddSchedulerJobGroupToSchedulerTriggerParam) (*view.SchedulerJobGroupSchedulerTriggerRefInventoryView, error) {
	resp := view.SchedulerJobGroupSchedulerTriggerRefInventoryView{}
	if err := cli.Post(ctx, "v1/scheduler/jobgroups/{schedulerJobGroupUuid}/scheduler/triggers/{schedulerTriggerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunNasAccessGroupRule deletes AliyunNasAccessGroupRule
func (cli *ZSClient) DeleteAliyunNasAccessGroupRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/nas/access/rule", uuid, string(deleteMode))
}

// DeleteBonding deletes Bonding
func (cli *ZSClient) DeleteBonding(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hosts/bondings", uuid, string(deleteMode))
}

// DeleteEcsSecurityGroupRemote deletes EcsSecurityGroupRemote
func (cli *ZSClient) DeleteEcsSecurityGroupRemote(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/security-group/remote", uuid, string(deleteMode))
}

// DeleteVmNicFromSecurityGroup deletes VmNicFromSecurityGroup
func (cli *ZSClient) DeleteVmNicFromSecurityGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/security-groups", uuid, string(deleteMode))
}

// UpdateTag updates Tag
func (cli *ZSClient) UpdateTag(ctx context.Context, uuid string, params param.UpdateTagParam) (*view.TagPatternInventoryView, error) {
	resp := view.TagPatternInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/tags", uuid, "", map[string]interface{}{
		"updateTag": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVRouterRouteTableToVRouter operates on VRouterRouteTableToVRouter
func (cli *ZSClient) AttachVRouterRouteTableToVRouter(ctx context.Context, params param.AttachVRouterRouteTableToVRouterParam) (*view.VRouterRouteTableInventoryView, error) {
	resp := view.VRouterRouteTableInventoryView{}
	if err := cli.Post(ctx, "v1/vrouter-route-tables/{routeTableUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVxlanVtep creates VxlanVtep
func (cli *ZSClient) CreateVxlanVtep(ctx context.Context, params param.CreateVxlanVtepParam) (*view.VtepInventoryView, error) {
	resp := view.VtepInventoryView{}
	if err := cli.Post(ctx, "v1/l2-networks/vxlan/vteps", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddMdevDeviceSpecToVmInstance adds MdevDeviceSpecToVmInstance
func (cli *ZSClient) AddMdevDeviceSpecToVmInstance(ctx context.Context, params param.AddMdevDeviceSpecToVmInstanceParam) (*view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	resp := view.VmInstanceMdevDeviceSpecRefInventoryView{}
	if err := cli.Post(ctx, "v1/mdev-device-specs/{mdevSpecUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachScsiLunFromVmInstance operates on ScsiLunFromVmInstance
func (cli *ZSClient) DetachScsiLunFromVmInstance(ctx context.Context, vmInstanceUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, fmt.Sprintf("scsi-lun/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// EnableCdpTask operates on EnableCdpTask
func (cli *ZSClient) EnableCdpTask(ctx context.Context, params param.EnableCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.Post(ctx, "v1/cdp-task/enable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableCdpTaskAsync Async
func (cli *ZSClient) EnableCdpTaskAsync(ctx context.Context, params param.EnableCdpTaskParam) (string, error) {

	resource := "v1/cdp-task/enable/{uuid}"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// SyncConnectionAccessPointFromRemote operates on ConnectionAccessPointFromRemote
func (cli *ZSClient) SyncConnectionAccessPointFromRemote(ctx context.Context, dataCenterUuid string, params param.SyncConnectionAccessPointFromRemoteParam) (*view.ConnectionAccessPointInventoryView, error) {
	resp := view.ConnectionAccessPointInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/access-point", dataCenterUuid, "", map[string]interface{}{
		"syncConnectionAccessPointFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RegisterLicenseRequestedApplication operates on LicenseRequestedApplication
func (cli *ZSClient) RegisterLicenseRequestedApplication(ctx context.Context, params param.RegisterLicenseRequestedApplicationParam) (*view.RegisterLicenseRequestedApplicationEventView, error) {
	resp := view.RegisterLicenseRequestedApplicationEventView{}
	if err := cli.Post(ctx, "v1/licenses/applications", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVpcVpnGatewayFromRemote operates on VpcVpnGatewayFromRemote
func (cli *ZSClient) SyncVpcVpnGatewayFromRemote(ctx context.Context, dataCenterUuid string, params param.SyncVpcVpnGatewayFromRemoteParam) (*view.VpcVpnGatewayInventoryView, error) {
	resp := view.VpcVpnGatewayInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/vpc-vpn", dataCenterUuid, "", map[string]interface{}{
		"syncVpcVpnGatewayFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsVpcInLocal deletes EcsVpcInLocal
func (cli *ZSClient) DeleteEcsVpcInLocal(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/vpc", uuid, string(deleteMode))
}

