// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2OrganizationStateParamDetail ChangeIAM2OrganizationState detail param
type ChangeIAM2OrganizationStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeIAM2OrganizationStateParam ChangeIAM2OrganizationState request param
type ChangeIAM2OrganizationStateParam struct {
	BaseParam
	Params ChangeIAM2OrganizationStateParamDetail `json:"changeIAM2OrganizationState"`
}
// CreateAutoScalingGroupAddingNewInstanceRuleParamDetail CreateAutoScalingGroupAddingNewInstanceRule detail param
type CreateAutoScalingGroupAddingNewInstanceRuleParamDetail struct {
	AdjustmentType string `json:"adjustmentType" validate:"required"`
	AdjustmentValue int `json:"adjustmentValue" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	AutoScalingGroupUuid string `json:"autoScalingGroupUuid" validate:"required"`
	Type *string `json:"type,omitempty"`
	Cooldown *int64 `json:"cooldown,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingGroupAddingNewInstanceRuleParam CreateAutoScalingGroupAddingNewInstanceRule request param
type CreateAutoScalingGroupAddingNewInstanceRuleParam struct {
	BaseParam
	Params CreateAutoScalingGroupAddingNewInstanceRuleParamDetail `json:"params"`
}
// SetServiceTypeOnHostNetworkBondingParamDetail SetServiceTypeOnHostNetworkBonding detail param
type SetServiceTypeOnHostNetworkBondingParamDetail struct {
	BondingUuids []string `json:"bondingUuids" validate:"required"`
	VlanIds []int `json:"vlanIds,omitempty"`
	ServiceTypes []string `json:"serviceTypes,omitempty"`
}

// SetServiceTypeOnHostNetworkBondingParam SetServiceTypeOnHostNetworkBonding request param
type SetServiceTypeOnHostNetworkBondingParam struct {
	BaseParam
	Params SetServiceTypeOnHostNetworkBondingParamDetail `json:"params"`
}
// GetCreateEcsImageProgressParamDetail GetCreateEcsImageProgress detail param
type GetCreateEcsImageProgressParamDetail struct {
}

// GetCreateEcsImageProgressParam GetCreateEcsImageProgress request param
type GetCreateEcsImageProgressParam struct {
	BaseParam
	Params GetCreateEcsImageProgressParamDetail `json:"getCreateEcsImageProgress"`
}
// AddAttributesToIAM2OrganizationParamDetail AddAttributesToIAM2Organization detail param
type AddAttributesToIAM2OrganizationParamDetail struct {
	Attributes []AttributeParam `json:"attributes" validate:"required"`
}

// AddAttributesToIAM2OrganizationParam AddAttributesToIAM2Organization request param
type AddAttributesToIAM2OrganizationParam struct {
	BaseParam
	Params AddAttributesToIAM2OrganizationParamDetail `json:"params"`
}
// AddAccessControlListToLoadBalancerParamDetail AddAccessControlListToLoadBalancer detail param
type AddAccessControlListToLoadBalancerParamDetail struct {
	AclUuids []string `json:"aclUuids" validate:"required"`
	AclType string `json:"aclType" validate:"required"`
	ServerGroupUuids []string `json:"serverGroupUuids,omitempty"`
}

// AddAccessControlListToLoadBalancerParam AddAccessControlListToLoadBalancer request param
type AddAccessControlListToLoadBalancerParam struct {
	BaseParam
	Params AddAccessControlListToLoadBalancerParamDetail `json:"params"`
}
// LogOutParamDetail LogOut detail param
type LogOutParamDetail struct {
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LogOutParam LogOut request param
type LogOutParam struct {
	BaseParam
	Params LogOutParamDetail `json:"logOut"`
}
// GetVmXmlHookScriptParamDetail GetVmXmlHookScript detail param
type GetVmXmlHookScriptParamDetail struct {
}

// GetVmXmlHookScriptParam GetVmXmlHookScript request param
type GetVmXmlHookScriptParam struct {
	BaseParam
	Params GetVmXmlHookScriptParamDetail `json:"getVmXmlHookScript"`
}
// AttachHybridKeyParamDetail AttachHybridKey detail param
type AttachHybridKeyParamDetail struct {
}

// AttachHybridKeyParam AttachHybridKey request param
type AttachHybridKeyParam struct {
	BaseParam
	Params AttachHybridKeyParamDetail `json:"attachHybridKey"`
}
// GetImageQgaParamDetail GetImageQga detail param
type GetImageQgaParamDetail struct {
}

// GetImageQgaParam GetImageQga request param
type GetImageQgaParam struct {
	BaseParam
	Params GetImageQgaParamDetail `json:"getImageQga"`
}
// GetInterdependentL3NetworksBackupStoragesParamDetail GetInterdependentL3NetworksBackupStorages detail param
type GetInterdependentL3NetworksBackupStoragesParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	BackupStorageUuid *string `json:"backupStorageUuid,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
}

// GetInterdependentL3NetworksBackupStoragesParam GetInterdependentL3NetworksBackupStorages request param
type GetInterdependentL3NetworksBackupStoragesParam struct {
	BaseParam
	Params GetInterdependentL3NetworksBackupStoragesParamDetail `json:"getInterdependentL3NetworksBackupStorages"`
}
// DeleteBackupFileInPublicParamDetail DeleteBackupFileInPublic detail param
type DeleteBackupFileInPublicParamDetail struct {
	Type string `json:"type" validate:"required"`
	RegionId string `json:"regionId" validate:"required"`
	File string `json:"file" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteBackupFileInPublicParam DeleteBackupFileInPublic request param
type DeleteBackupFileInPublicParam struct {
	BaseParam
	Params DeleteBackupFileInPublicParamDetail `json:"deleteBackupFileInPublic"`
}
// BatchCreateIAM2VirtualIDFromConfigFileParamDetail BatchCreateIAM2VirtualIDFromConfigFile detail param
type BatchCreateIAM2VirtualIDFromConfigFileParamDetail struct {
	VirtualIDInfos string `json:"virtualIDInfos" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// BatchCreateIAM2VirtualIDFromConfigFileParam BatchCreateIAM2VirtualIDFromConfigFile request param
type BatchCreateIAM2VirtualIDFromConfigFileParam struct {
	BaseParam
	Params BatchCreateIAM2VirtualIDFromConfigFileParamDetail `json:"params"`
}
// SetVmClockTrackParamDetail SetVmClockTrack detail param
type SetVmClockTrackParamDetail struct {
	Track string `json:"track" validate:"required"`
	SyncAfterVMResume *bool `json:"syncAfterVMResume,omitempty"`
	IntervalInSeconds *int `json:"intervalInSeconds,omitempty"`
}

// SetVmClockTrackParam SetVmClockTrack request param
type SetVmClockTrackParam struct {
	BaseParam
	Params SetVmClockTrackParamDetail `json:"setVmClockTrack"`
}
// UpdateEmailMonitorTriggerActionParamDetail UpdateEmailMonitorTriggerAction detail param
type UpdateEmailMonitorTriggerActionParamDetail struct {
	Name string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	MediaUuid *string `json:"mediaUuid,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateEmailMonitorTriggerActionParam UpdateEmailMonitorTriggerAction request param
type UpdateEmailMonitorTriggerActionParam struct {
	BaseParam
	Params UpdateEmailMonitorTriggerActionParamDetail `json:"updateEmailMonitorTriggerAction"`
}
// SyncDataCenterFromRemoteParamDetail SyncDataCenterFromRemote detail param
type SyncDataCenterFromRemoteParamDetail struct {
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncDataCenterFromRemoteParam SyncDataCenterFromRemote request param
type SyncDataCenterFromRemoteParam struct {
	BaseParam
	Params SyncDataCenterFromRemoteParamDetail `json:"syncDataCenterFromRemote"`
}
// ChangeBackupStorageStateParamDetail ChangeBackupStorageState detail param
type ChangeBackupStorageStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBackupStorageStateParam ChangeBackupStorageState request param
type ChangeBackupStorageStateParam struct {
	BaseParam
	Params ChangeBackupStorageStateParamDetail `json:"changeBackupStorageState"`
}
// SetVmInstanceHygonMdevParamDetail SetVmInstanceHygonMdev detail param
type SetVmInstanceHygonMdevParamDetail struct {
	HygonSecurityElementEnable string `json:"hygonSecurityElementEnable" validate:"required"`
}

// SetVmInstanceHygonMdevParam SetVmInstanceHygonMdev request param
type SetVmInstanceHygonMdevParam struct {
	BaseParam
	Params SetVmInstanceHygonMdevParamDetail `json:"params"`
}
// GetCandidateIsoForAttachingVmParamDetail GetCandidateIsoForAttachingVm detail param
type GetCandidateIsoForAttachingVmParamDetail struct {
}

// GetCandidateIsoForAttachingVmParam GetCandidateIsoForAttachingVm request param
type GetCandidateIsoForAttachingVmParam struct {
	BaseParam
	Params GetCandidateIsoForAttachingVmParamDetail `json:"getCandidateIsoForAttachingVm"`
}
// SecurityMachineDetectSyncParamDetail SecurityMachineDetectSync detail param
type SecurityMachineDetectSyncParamDetail struct {
}

// SecurityMachineDetectSyncParam SecurityMachineDetectSync request param
type SecurityMachineDetectSyncParam struct {
	BaseParam
	Params SecurityMachineDetectSyncParamDetail `json:"params"`
}
// ChangeSecurityGroupStateParamDetail ChangeSecurityGroupState detail param
type ChangeSecurityGroupStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSecurityGroupStateParam ChangeSecurityGroupState request param
type ChangeSecurityGroupStateParam struct {
	BaseParam
	Params ChangeSecurityGroupStateParamDetail `json:"changeSecurityGroupState"`
}
// ChangeBareMetal2ChassisOfferingStateParamDetail ChangeBareMetal2ChassisOfferingState detail param
type ChangeBareMetal2ChassisOfferingStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBareMetal2ChassisOfferingStateParam ChangeBareMetal2ChassisOfferingState request param
type ChangeBareMetal2ChassisOfferingStateParam struct {
	BaseParam
	Params ChangeBareMetal2ChassisOfferingStateParamDetail `json:"changeBareMetal2ChassisOfferingState"`
}
// GetPrometheusMetricLabelValueParamDetail GetPrometheusMetricLabelValue detail param
type GetPrometheusMetricLabelValueParamDetail struct {
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	StartTime *int64 `json:"startTime,omitempty"`
	EndTime *int64 `json:"endTime,omitempty"`
	LabelNames []string `json:"labelNames,omitempty"`
	FilterLabels []string `json:"filterLabels,omitempty"`
}

// GetPrometheusMetricLabelValueParam GetPrometheusMetricLabelValue request param
type GetPrometheusMetricLabelValueParam struct {
	BaseParam
	Params GetPrometheusMetricLabelValueParamDetail `json:"getPrometheusMetricLabelValue"`
}
// UpdateAlarmDataParamDetail UpdateAlarmData detail param
type UpdateAlarmDataParamDetail struct {
	DataUuid *string `json:"dataUuid,omitempty"`
	DataStartTime *int64 `json:"dataStartTime,omitempty"`
	DataEndTime *int64 `json:"dataEndTime,omitempty"`
	UpdateMode string `json:"updateMode" validate:"required"`
	ReadStatus *string `json:"readStatus,omitempty"`
}

// UpdateAlarmDataParam UpdateAlarmData request param
type UpdateAlarmDataParam struct {
	BaseParam
	Params UpdateAlarmDataParamDetail `json:"updateAlarmData"`
}
// CreateVpnIpsecConfigParamDetail CreateVpnIpsecConfig detail param
type CreateVpnIpsecConfigParamDetail struct {
	Name string `json:"name" validate:"required"`
	Pfs *string `json:"pfs,omitempty"`
	EncAlg *string `json:"encAlg,omitempty"`
	AuthAlg *string `json:"authAlg,omitempty"`
	Lifetime *int `json:"lifetime,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpnIpsecConfigParam CreateVpnIpsecConfig request param
type CreateVpnIpsecConfigParam struct {
	BaseParam
	Params CreateVpnIpsecConfigParamDetail `json:"params"`
}
// SNSEmailTestConnectionParamDetail SNSEmailTestConnection detail param
type SNSEmailTestConnectionParamDetail struct {
	Emails []string `json:"emails,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	EndpointUuid *string `json:"endpointUuid,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Text *string `json:"text,omitempty"`
}

// SNSEmailTestConnectionParam SNSEmailTestConnection request param
type SNSEmailTestConnectionParam struct {
	BaseParam
	Params SNSEmailTestConnectionParamDetail `json:"params"`
}
// RegisterLicenseServerParamDetail RegisterLicenseServer detail param
type RegisterLicenseServerParamDetail struct {
	Ip string `json:"ip" validate:"required"`
	LoginParams map[string]interface{} `json:"loginParams" validate:"required"`
}

// RegisterLicenseServerParam RegisterLicenseServer request param
type RegisterLicenseServerParam struct {
	BaseParam
	Params RegisterLicenseServerParamDetail `json:"params"`
}
// ChangeAutoScalingGroupStateParamDetail ChangeAutoScalingGroupState detail param
type ChangeAutoScalingGroupStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeAutoScalingGroupStateParam ChangeAutoScalingGroupState request param
type ChangeAutoScalingGroupStateParam struct {
	BaseParam
	Params ChangeAutoScalingGroupStateParamDetail `json:"changeAutoScalingGroupState"`
}
// CreateAutoScalingGroupRemovalInstanceRuleParamDetail CreateAutoScalingGroupRemovalInstanceRule detail param
type CreateAutoScalingGroupRemovalInstanceRuleParamDetail struct {
	AdjustmentType string `json:"adjustmentType" validate:"required"`
	AdjustmentValue int `json:"adjustmentValue" validate:"required"`
	RemovalPolicy string `json:"removalPolicy" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	AutoScalingGroupUuid string `json:"autoScalingGroupUuid" validate:"required"`
	Type *string `json:"type,omitempty"`
	Cooldown *int64 `json:"cooldown,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingGroupRemovalInstanceRuleParam CreateAutoScalingGroupRemovalInstanceRule request param
type CreateAutoScalingGroupRemovalInstanceRuleParam struct {
	BaseParam
	Params CreateAutoScalingGroupRemovalInstanceRuleParamDetail `json:"params"`
}
// ChangeSlbGroupMonitorIpsParamDetail ChangeSlbGroupMonitorIps detail param
type ChangeSlbGroupMonitorIpsParamDetail struct {
	MonitorIps []string `json:"monitorIps" validate:"required"`
}

// ChangeSlbGroupMonitorIpsParam ChangeSlbGroupMonitorIps request param
type ChangeSlbGroupMonitorIpsParam struct {
	BaseParam
	Params ChangeSlbGroupMonitorIpsParamDetail `json:"changeSlbGroupMonitorIps"`
}
// DeleteModelEvaluationTasksParamDetail DeleteModelEvaluationTasks detail param
type DeleteModelEvaluationTasksParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// DeleteModelEvaluationTasksParam DeleteModelEvaluationTasks request param
type DeleteModelEvaluationTasksParam struct {
	BaseParam
	Params DeleteModelEvaluationTasksParamDetail `json:"deleteModelEvaluationTasks"`
}
// AttachL3NetworkToVmParamDetail AttachL3NetworkToVm detail param
type AttachL3NetworkToVmParamDetail struct {
	StaticIp *string `json:"staticIp,omitempty"`
	DriverType *string `json:"driverType,omitempty"`
	CustomMac *string `json:"customMac,omitempty"`
	VmNicParams *string `json:"vmNicParams,omitempty"`
}

// AttachL3NetworkToVmParam AttachL3NetworkToVm request param
type AttachL3NetworkToVmParam struct {
	BaseParam
	Params AttachL3NetworkToVmParamDetail `json:"params"`
}
// AttachPrimaryStorageToClusterParamDetail AttachPrimaryStorageToCluster detail param
type AttachPrimaryStorageToClusterParamDetail struct {
}

// AttachPrimaryStorageToClusterParam AttachPrimaryStorageToCluster request param
type AttachPrimaryStorageToClusterParam struct {
	BaseParam
	Params AttachPrimaryStorageToClusterParamDetail `json:"params"`
}
// AttachL2NetworkToClusterParamDetail AttachL2NetworkToCluster detail param
type AttachL2NetworkToClusterParamDetail struct {
	L2ProviderType *string `json:"l2ProviderType,omitempty"`
}

// AttachL2NetworkToClusterParam AttachL2NetworkToCluster request param
type AttachL2NetworkToClusterParam struct {
	BaseParam
	Params AttachL2NetworkToClusterParamDetail `json:"attachL2NetworkToCluster"`
}
// ChangeVmNicTypeParamDetail ChangeVmNicType detail param
type ChangeVmNicTypeParamDetail struct {
	VmNicType string `json:"vmNicType" validate:"required"`
}

// ChangeVmNicTypeParam ChangeVmNicType request param
type ChangeVmNicTypeParam struct {
	BaseParam
	Params ChangeVmNicTypeParamDetail `json:"changeVmNicType"`
}
// ChangeFirewallRuleStateParamDetail ChangeFirewallRuleState detail param
type ChangeFirewallRuleStateParamDetail struct {
	State string `json:"state" validate:"required"`
}

// ChangeFirewallRuleStateParam ChangeFirewallRuleState request param
type ChangeFirewallRuleStateParam struct {
	BaseParam
	Params ChangeFirewallRuleStateParamDetail `json:"changeFirewallRuleState"`
}
// GetMdevDeviceCandidatesParamDetail GetMdevDeviceCandidates detail param
type GetMdevDeviceCandidatesParamDetail struct {
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	VmInstanceUuid *string `json:"vmInstanceUuid,omitempty"`
	Types []string `json:"types,omitempty"`
	MinAvailableDevicesPerHost *int `json:"minAvailableDevicesPerHost,omitempty"`
}

// GetMdevDeviceCandidatesParam GetMdevDeviceCandidates request param
type GetMdevDeviceCandidatesParam struct {
	BaseParam
	Params GetMdevDeviceCandidatesParamDetail `json:"getMdevDeviceCandidates"`
}
// GetTwoFactorAuthenticationStateParamDetail GetTwoFactorAuthenticationState detail param
type GetTwoFactorAuthenticationStateParamDetail struct {
}

// GetTwoFactorAuthenticationStateParam GetTwoFactorAuthenticationState request param
type GetTwoFactorAuthenticationStateParam struct {
	BaseParam
	Params GetTwoFactorAuthenticationStateParamDetail `json:"getTwoFactorAuthenticationState"`
}
// BootstrapMiniHostParamDetail BootstrapMiniHost detail param
type BootstrapMiniHostParamDetail struct {
	Local MiniHostInfoParam `json:"local" validate:"required"`
	Peer MiniHostInfoParam `json:"peer" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// BootstrapMiniHostParam BootstrapMiniHost request param
type BootstrapMiniHostParam struct {
	BaseParam
	Params BootstrapMiniHostParamDetail `json:"params"`
}
// RemoveActionFromAlarmParamDetail RemoveActionFromAlarm detail param
type RemoveActionFromAlarmParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveActionFromAlarmParam RemoveActionFromAlarm request param
type RemoveActionFromAlarmParam struct {
	BaseParam
	Params RemoveActionFromAlarmParamDetail `json:"removeActionFromAlarm"`
}
// ChangeEipStateParamDetail ChangeEipState detail param
type ChangeEipStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeEipStateParam ChangeEipState request param
type ChangeEipStateParam struct {
	BaseParam
	Params ChangeEipStateParamDetail `json:"changeEipState"`
}
// DetachSshKeyPairFromVmInstanceParamDetail DetachSshKeyPairFromVmInstance detail param
type DetachSshKeyPairFromVmInstanceParamDetail struct {
}

// DetachSshKeyPairFromVmInstanceParam DetachSshKeyPairFromVmInstance request param
type DetachSshKeyPairFromVmInstanceParam struct {
	BaseParam
	Params DetachSshKeyPairFromVmInstanceParamDetail `json:"detachSshKeyPairFromVmInstance"`
}
// GetPrimaryStorageCandidatesForVmMigrationParamDetail GetPrimaryStorageCandidatesForVmMigration detail param
type GetPrimaryStorageCandidatesForVmMigrationParamDetail struct {
	WithDataVolumes *bool `json:"withDataVolumes,omitempty"`
	MigrateStorageOnly *bool `json:"migrateStorageOnly,omitempty"`
}

// GetPrimaryStorageCandidatesForVmMigrationParam GetPrimaryStorageCandidatesForVmMigration request param
type GetPrimaryStorageCandidatesForVmMigrationParam struct {
	BaseParam
	Params GetPrimaryStorageCandidatesForVmMigrationParamDetail `json:"getPrimaryStorageCandidatesForVmMigration"`
}
// PrimaryStorageMigrateVolumeParamDetail PrimaryStorageMigrateVolume detail param
type PrimaryStorageMigrateVolumeParamDetail struct {
	DstPrimaryStorageUuid string `json:"dstPrimaryStorageUuid" validate:"required"`
	VolumeProvisioningStrategy *string `json:"volumeProvisioningStrategy,omitempty"`
}

// PrimaryStorageMigrateVolumeParam PrimaryStorageMigrateVolume request param
type PrimaryStorageMigrateVolumeParam struct {
	BaseParam
	Params PrimaryStorageMigrateVolumeParamDetail `json:"primaryStorageMigrateVolume"`
}
// DeleteHybridEipRemoteParamDetail DeleteHybridEipRemote detail param
type DeleteHybridEipRemoteParamDetail struct {
	Type string `json:"type" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteHybridEipRemoteParam DeleteHybridEipRemote request param
type DeleteHybridEipRemoteParam struct {
	BaseParam
	Params DeleteHybridEipRemoteParamDetail `json:"deleteHybridEipRemote"`
}
// DeleteModelServiceInstanceGroupsParamDetail DeleteModelServiceInstanceGroups detail param
type DeleteModelServiceInstanceGroupsParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteModelServiceInstanceGroupsParam DeleteModelServiceInstanceGroups request param
type DeleteModelServiceInstanceGroupsParam struct {
	BaseParam
	Params DeleteModelServiceInstanceGroupsParamDetail `json:"deleteModelServiceInstanceGroups"`
}
// GetVmBootOrderParamDetail GetVmBootOrder detail param
type GetVmBootOrderParamDetail struct {
}

// GetVmBootOrderParam GetVmBootOrder request param
type GetVmBootOrderParam struct {
	BaseParam
	Params GetVmBootOrderParamDetail `json:"getVmBootOrder"`
}
// SetVmBootOrderParamDetail SetVmBootOrder detail param
type SetVmBootOrderParamDetail struct {
	BootOrder []string `json:"bootOrder,omitempty"`
}

// SetVmBootOrderParam SetVmBootOrder request param
type SetVmBootOrderParam struct {
	BaseParam
	Params SetVmBootOrderParamDetail `json:"setVmBootOrder"`
}
// GetDatabaseBackupFromImageStoreParamDetail GetDatabaseBackupFromImageStore detail param
type GetDatabaseBackupFromImageStoreParamDetail struct {
	Url string `json:"url" validate:"required"`
	RegistryPort *int `json:"registryPort,omitempty"`
}

// GetDatabaseBackupFromImageStoreParam GetDatabaseBackupFromImageStore request param
type GetDatabaseBackupFromImageStoreParam struct {
	BaseParam
	Params GetDatabaseBackupFromImageStoreParamDetail `json:"getDatabaseBackupFromImageStore"`
}
// SyncEcsVSwitchFromRemoteParamDetail SyncEcsVSwitchFromRemote detail param
type SyncEcsVSwitchFromRemoteParamDetail struct {
	VSwitchId *string `json:"vSwitchId,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsVSwitchFromRemoteParam SyncEcsVSwitchFromRemote request param
type SyncEcsVSwitchFromRemoteParam struct {
	BaseParam
	Params SyncEcsVSwitchFromRemoteParamDetail `json:"params"`
}
// LocateLocalRaidPhysicalDriveParamDetail LocateLocalRaidPhysicalDrive detail param
type LocateLocalRaidPhysicalDriveParamDetail struct {
	Locate *bool `json:"locate,omitempty"`
}

// LocateLocalRaidPhysicalDriveParam LocateLocalRaidPhysicalDrive request param
type LocateLocalRaidPhysicalDriveParam struct {
	BaseParam
	Params LocateLocalRaidPhysicalDriveParamDetail `json:"locateLocalRaidPhysicalDrive"`
}
// CleanUpBaremetalChassisBondingParamDetail CleanUpBaremetalChassisBonding detail param
type CleanUpBaremetalChassisBondingParamDetail struct {
}

// CleanUpBaremetalChassisBondingParam CleanUpBaremetalChassisBonding request param
type CleanUpBaremetalChassisBondingParam struct {
	BaseParam
	Params CleanUpBaremetalChassisBondingParamDetail `json:"cleanUpBaremetalChassisBonding"`
}
// RemovePciDeviceSpecFromVmInstanceParamDetail RemovePciDeviceSpecFromVmInstance detail param
type RemovePciDeviceSpecFromVmInstanceParamDetail struct {
}

// RemovePciDeviceSpecFromVmInstanceParam RemovePciDeviceSpecFromVmInstance request param
type RemovePciDeviceSpecFromVmInstanceParam struct {
	BaseParam
	Params RemovePciDeviceSpecFromVmInstanceParamDetail `json:"removePciDeviceSpecFromVmInstance"`
}
// AddIAM2VirtualIDGroupToProjectsParamDetail AddIAM2VirtualIDGroupToProjects detail param
type AddIAM2VirtualIDGroupToProjectsParamDetail struct {
	Structs []AddIAM2VirtualIDGroupToProjects_IAM2ProjectRoleRefStructParam `json:"structs,omitempty"`
}

// AddIAM2VirtualIDGroupToProjectsParam AddIAM2VirtualIDGroupToProjects request param
type AddIAM2VirtualIDGroupToProjectsParam struct {
	BaseParam
	Params AddIAM2VirtualIDGroupToProjectsParamDetail `json:"params"`
}
// RemoveServerGroupFromLoadBalancerListenerParamDetail RemoveServerGroupFromLoadBalancerListener detail param
type RemoveServerGroupFromLoadBalancerListenerParamDetail struct {
	ServerGroupUuid string `json:"serverGroupUuid" validate:"required"`
}

// RemoveServerGroupFromLoadBalancerListenerParam RemoveServerGroupFromLoadBalancerListener request param
type RemoveServerGroupFromLoadBalancerListenerParam struct {
	BaseParam
	Params RemoveServerGroupFromLoadBalancerListenerParamDetail `json:"removeServerGroupFromLoadBalancerListener"`
}
// AddSharedBlockToSharedBlockGroupParamDetail AddSharedBlockToSharedBlockGroup detail param
type AddSharedBlockToSharedBlockGroupParamDetail struct {
	DiskUuid string `json:"diskUuid" validate:"required"`
}

// AddSharedBlockToSharedBlockGroupParam AddSharedBlockToSharedBlockGroup request param
type AddSharedBlockToSharedBlockGroupParam struct {
	BaseParam
	Params AddSharedBlockToSharedBlockGroupParamDetail `json:"params"`
}
// RefreshCaptchaParamDetail RefreshCaptcha detail param
type RefreshCaptchaParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RefreshCaptchaParam RefreshCaptcha request param
type RefreshCaptchaParam struct {
	BaseParam
	Params RefreshCaptchaParamDetail `json:"refreshCaptcha"`
}
// DeleteEcsVSwitchInLocalParamDetail DeleteEcsVSwitchInLocal detail param
type DeleteEcsVSwitchInLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsVSwitchInLocalParam DeleteEcsVSwitchInLocal request param
type DeleteEcsVSwitchInLocalParam struct {
	BaseParam
	Params DeleteEcsVSwitchInLocalParamDetail `json:"deleteEcsVSwitchInLocal"`
}
// DeleteTagParamDetail DeleteTag detail param
type DeleteTagParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteTagParam DeleteTag request param
type DeleteTagParam struct {
	BaseParam
	Params DeleteTagParamDetail `json:"deleteTag"`
}
// AddIAM2VirtualIDsToOrganizationParamDetail AddIAM2VirtualIDsToOrganization detail param
type AddIAM2VirtualIDsToOrganizationParamDetail struct {
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
}

// AddIAM2VirtualIDsToOrganizationParam AddIAM2VirtualIDsToOrganization request param
type AddIAM2VirtualIDsToOrganizationParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToOrganizationParamDetail `json:"params"`
}
// AttachProvisionNicToBondingParamDetail AttachProvisionNicToBonding detail param
type AttachProvisionNicToBondingParamDetail struct {
	ProvisionNicUuid string `json:"provisionNicUuid" validate:"required"`
	ProvisionIp *string `json:"provisionIp,omitempty"`
	CustomMac *string `json:"customMac,omitempty"`
}

// AttachProvisionNicToBondingParam AttachProvisionNicToBonding request param
type AttachProvisionNicToBondingParam struct {
	BaseParam
	Params AttachProvisionNicToBondingParamDetail `json:"params"`
}
// ExportNbdVolumesParamDetail ExportNbdVolumes detail param
type ExportNbdVolumesParamDetail struct {
	VolumeUuids []string `json:"volumeUuids" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Force *bool `json:"force,omitempty"`
}

// ExportNbdVolumesParam ExportNbdVolumes request param
type ExportNbdVolumesParam struct {
	BaseParam
	Params ExportNbdVolumesParamDetail `json:"params"`
}
// SelfTestLocalRaidParamDetail SelfTestLocalRaid detail param
type SelfTestLocalRaidParamDetail struct {
}

// SelfTestLocalRaidParam SelfTestLocalRaid request param
type SelfTestLocalRaidParam struct {
	BaseParam
	Params SelfTestLocalRaidParamDetail `json:"selfTestLocalRaid"`
}
// ChangeSNSApplicationPlatformStateParamDetail ChangeSNSApplicationPlatformState detail param
type ChangeSNSApplicationPlatformStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSNSApplicationPlatformStateParam ChangeSNSApplicationPlatformState request param
type ChangeSNSApplicationPlatformStateParam struct {
	BaseParam
	Params ChangeSNSApplicationPlatformStateParamDetail `json:"changeSNSApplicationPlatformState"`
}
// PowerOffBareMetal2ChassisParamDetail PowerOffBareMetal2Chassis detail param
type PowerOffBareMetal2ChassisParamDetail struct {
}

// PowerOffBareMetal2ChassisParam PowerOffBareMetal2Chassis request param
type PowerOffBareMetal2ChassisParam struct {
	BaseParam
	Params PowerOffBareMetal2ChassisParamDetail `json:"powerOffBareMetal2Chassis"`
}
// SdnControllerChangeHostParamDetail SdnControllerChangeHost detail param
type SdnControllerChangeHostParamDetail struct {
	VSwitchType *string `json:"vSwitchType,omitempty"`
	NicNames []string `json:"nicNames,omitempty"`
	VtepIp *string `json:"vtepIp,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	BondMode *string `json:"bondMode,omitempty"`
	LacpMode *string `json:"lacpMode,omitempty"`
}

// SdnControllerChangeHostParam SdnControllerChangeHost request param
type SdnControllerChangeHostParam struct {
	BaseParam
	Params SdnControllerChangeHostParamDetail `json:"sdnControllerChangeHost"`
}
// UpdateResourcePriceParamDetail UpdateResourcePrice detail param
type UpdateResourcePriceParamDetail struct {
	EndDateInLong *int64 `json:"endDateInLong,omitempty"`
	SetEndDateInLongBaseOnCurrentTime *bool `json:"setEndDateInLongBaseOnCurrentTime,omitempty"`
}

// UpdateResourcePriceParam UpdateResourcePrice request param
type UpdateResourcePriceParam struct {
	BaseParam
	Params UpdateResourcePriceParamDetail `json:"updateResourcePrice"`
}
// DetachTagFromResourcesParamDetail DetachTagFromResources detail param
type DetachTagFromResourcesParamDetail struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
}

// DetachTagFromResourcesParam DetachTagFromResources request param
type DetachTagFromResourcesParam struct {
	BaseParam
	Params DetachTagFromResourcesParamDetail `json:"detachTagFromResources"`
}
// ChangeHostStateParamDetail ChangeHostState detail param
type ChangeHostStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeHostStateParam ChangeHostState request param
type ChangeHostStateParam struct {
	BaseParam
	Params ChangeHostStateParamDetail `json:"changeHostState"`
}
// UpdateVmNicMacParamDetail UpdateVmNicMac detail param
type UpdateVmNicMacParamDetail struct {
	Mac string `json:"mac" validate:"required"`
}

// UpdateVmNicMacParam UpdateVmNicMac request param
type UpdateVmNicMacParam struct {
	BaseParam
	Params UpdateVmNicMacParamDetail `json:"updateVmNicMac"`
}
// DeleteVmInstanceHaLevelParamDetail DeleteVmInstanceHaLevel detail param
type DeleteVmInstanceHaLevelParamDetail struct {
}

// DeleteVmInstanceHaLevelParam DeleteVmInstanceHaLevel request param
type DeleteVmInstanceHaLevelParam struct {
	BaseParam
	Params DeleteVmInstanceHaLevelParamDetail `json:"deleteVmInstanceHaLevel"`
}
// DeleteResourcePriceParamDetail DeleteResourcePrice detail param
type DeleteResourcePriceParamDetail struct {
	CutoffPrice *bool `json:"cutoffPrice,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteResourcePriceParam DeleteResourcePrice request param
type DeleteResourcePriceParam struct {
	BaseParam
	Params DeleteResourcePriceParamDetail `json:"deleteResourcePrice"`
}
// CleanUpBareMetal2BondingParamDetail CleanUpBareMetal2Bonding detail param
type CleanUpBareMetal2BondingParamDetail struct {
}

// CleanUpBareMetal2BondingParam CleanUpBareMetal2Bonding request param
type CleanUpBareMetal2BondingParam struct {
	BaseParam
	Params CleanUpBareMetal2BondingParamDetail `json:"cleanUpBareMetal2Bonding"`
}
// DeleteMetricDataParamDetail DeleteMetricData detail param
type DeleteMetricDataParamDetail struct {
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	Labels []string `json:"labels,omitempty"`
}

// DeleteMetricDataParam DeleteMetricData request param
type DeleteMetricDataParam struct {
	BaseParam
	Params DeleteMetricDataParamDetail `json:"deleteMetricData"`
}
// AddLabelToAlarmParamDetail AddLabelToAlarm detail param
type AddLabelToAlarmParamDetail struct {
	Key string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
	Operator string `json:"operator" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddLabelToAlarmParam AddLabelToAlarm request param
type AddLabelToAlarmParam struct {
	BaseParam
	Params AddLabelToAlarmParamDetail `json:"params"`
}
// SyncAliyunRouterInterfaceFromRemoteParamDetail SyncAliyunRouterInterfaceFromRemote detail param
type SyncAliyunRouterInterfaceFromRemoteParamDetail struct {
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncAliyunRouterInterfaceFromRemoteParam SyncAliyunRouterInterfaceFromRemote request param
type SyncAliyunRouterInterfaceFromRemoteParam struct {
	BaseParam
	Params SyncAliyunRouterInterfaceFromRemoteParamDetail `json:"syncAliyunRouterInterfaceFromRemote"`
}
// ExportVmOvaPackageParamDetail ExportVmOvaPackage detail param
type ExportVmOvaPackageParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	VmUuid string `json:"vmUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// ExportVmOvaPackageParam ExportVmOvaPackage request param
type ExportVmOvaPackageParam struct {
	BaseParam
	Params ExportVmOvaPackageParamDetail `json:"params"`
}
// RevertVmFromCdpBackupParamDetail RevertVmFromCdpBackup detail param
type RevertVmFromCdpBackupParamDetail struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume *string `json:"primaryStorageUuidForDataVolume,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	UseExistingVolume *bool `json:"useExistingVolume,omitempty"`
	RecoverBandwidth *int64 `json:"recoverBandwidth,omitempty"`
}

// RevertVmFromCdpBackupParam RevertVmFromCdpBackup request param
type RevertVmFromCdpBackupParam struct {
	BaseParam
	Params RevertVmFromCdpBackupParamDetail `json:"revertVmFromCdpBackup"`
}
// SNSFeiShuTestConnectionParamDetail SNSFeiShuTestConnection detail param
type SNSFeiShuTestConnectionParamDetail struct {
	Url *string `json:"url,omitempty"`
	AtAll *bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	Secret *string `json:"secret,omitempty"`
	TestMsg string `json:"testMsg" validate:"required"`
	EndpointUuid *string `json:"endpointUuid,omitempty"`
}

// SNSFeiShuTestConnectionParam SNSFeiShuTestConnection request param
type SNSFeiShuTestConnectionParam struct {
	BaseParam
	Params SNSFeiShuTestConnectionParamDetail `json:"params"`
}
// GetSchedulerExecutionReportParamDetail GetSchedulerExecutionReport detail param
type GetSchedulerExecutionReportParamDetail struct {
	StartTime int64 `json:"startTime" validate:"required"`
	IntervalTimeUnit string `json:"intervalTimeUnit" validate:"required"`
	Range int `json:"range" validate:"required"`
	SchedulerJobTypes []string `json:"schedulerJobTypes" validate:"required"`
}

// GetSchedulerExecutionReportParam GetSchedulerExecutionReport request param
type GetSchedulerExecutionReportParam struct {
	BaseParam
	Params GetSchedulerExecutionReportParamDetail `json:"getSchedulerExecutionReport"`
}
// CreateFirewallRuleFromConfigFileParamDetail CreateFirewallRuleFromConfigFile detail param
type CreateFirewallRuleFromConfigFileParamDetail struct {
	RuleInfo string `json:"ruleInfo" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFirewallRuleFromConfigFileParam CreateFirewallRuleFromConfigFile request param
type CreateFirewallRuleFromConfigFileParam struct {
	BaseParam
	Params CreateFirewallRuleFromConfigFileParamDetail `json:"params"`
}
// GetSupportedIdentityModelsParamDetail GetSupportedIdentityModels detail param
type GetSupportedIdentityModelsParamDetail struct {
}

// GetSupportedIdentityModelsParam GetSupportedIdentityModels request param
type GetSupportedIdentityModelsParam struct {
	BaseParam
	Params GetSupportedIdentityModelsParamDetail `json:"getSupportedIdentityModels"`
}
// SuspendLongJobParamDetail SuspendLongJob detail param
type SuspendLongJobParamDetail struct {
}

// SuspendLongJobParam SuspendLongJob request param
type SuspendLongJobParam struct {
	BaseParam
	Params SuspendLongJobParamDetail `json:"suspendLongJob"`
}
// AddUserToGroupParamDetail AddUserToGroup detail param
type AddUserToGroupParamDetail struct {
	UserUuid string `json:"userUuid" validate:"required"`
}

// AddUserToGroupParam AddUserToGroup request param
type AddUserToGroupParam struct {
	BaseParam
	Params AddUserToGroupParamDetail `json:"params"`
}
// UpdateVRouterOspfAreaParamDetail UpdateVRouterOspfArea detail param
type UpdateVRouterOspfAreaParamDetail struct {
	AreaAuth *string `json:"areaAuth,omitempty"`
	AreaType *string `json:"areaType,omitempty"`
	Password *string `json:"password,omitempty"`
	KeyId *int `json:"keyId,omitempty"`
}

// UpdateVRouterOspfAreaParam UpdateVRouterOspfArea request param
type UpdateVRouterOspfAreaParam struct {
	BaseParam
	Params UpdateVRouterOspfAreaParamDetail `json:"updateVRouterOspfArea"`
}
// GetPrimaryStorageTypesParamDetail GetPrimaryStorageTypes detail param
type GetPrimaryStorageTypesParamDetail struct {
}

// GetPrimaryStorageTypesParam GetPrimaryStorageTypes request param
type GetPrimaryStorageTypesParam struct {
	BaseParam
	Params GetPrimaryStorageTypesParamDetail `json:"getPrimaryStorageTypes"`
}
// BatchDeleteVolumeSnapshotParamDetail BatchDeleteVolumeSnapshot detail param
type BatchDeleteVolumeSnapshotParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// BatchDeleteVolumeSnapshotParam BatchDeleteVolumeSnapshot request param
type BatchDeleteVolumeSnapshotParam struct {
	BaseParam
	Params BatchDeleteVolumeSnapshotParamDetail `json:"batchDeleteVolumeSnapshot"`
}
// ReloadLicenseParamDetail ReloadLicense detail param
type ReloadLicenseParamDetail struct {
	ManagementNodeUuids []string `json:"managementNodeUuids,omitempty"`
	AdditionSession *string `json:"additionSession,omitempty"`
}

// ReloadLicenseParam ReloadLicense request param
type ReloadLicenseParam struct {
	BaseParam
	Params ReloadLicenseParamDetail `json:"reloadLicense"`
}
// DeleteNicQosParamDetail DeleteNicQos detail param
type DeleteNicQosParamDetail struct {
	Direction string `json:"direction" validate:"required"`
}

// DeleteNicQosParam DeleteNicQos request param
type DeleteNicQosParam struct {
	BaseParam
	Params DeleteNicQosParamDetail `json:"deleteNicQos"`
}
// ChangeL2NetworkVlanIdParamDetail ChangeL2NetworkVlanId detail param
type ChangeL2NetworkVlanIdParamDetail struct {
	Vlan *int `json:"vlan,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ChangeL2NetworkVlanIdParam ChangeL2NetworkVlanId request param
type ChangeL2NetworkVlanIdParam struct {
	BaseParam
	Params ChangeL2NetworkVlanIdParamDetail `json:"changeL2NetworkVlanId"`
}
// GetResourceStackVmStatusParamDetail GetResourceStackVmStatus detail param
type GetResourceStackVmStatusParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetResourceStackVmStatusParam GetResourceStackVmStatus request param
type GetResourceStackVmStatusParam struct {
	BaseParam
	Params GetResourceStackVmStatusParamDetail `json:"getResourceStackVmStatus"`
}
// DetachHybridKeyParamDetail DetachHybridKey detail param
type DetachHybridKeyParamDetail struct {
}

// DetachHybridKeyParam DetachHybridKey request param
type DetachHybridKeyParam struct {
	BaseParam
	Params DetachHybridKeyParamDetail `json:"detachHybridKey"`
}
// RemoveDnsFromVpcRouterParamDetail RemoveDnsFromVpcRouter detail param
type RemoveDnsFromVpcRouterParamDetail struct {
	Dns string `json:"dns" validate:"required"`
}

// RemoveDnsFromVpcRouterParam RemoveDnsFromVpcRouter request param
type RemoveDnsFromVpcRouterParam struct {
	BaseParam
	Params RemoveDnsFromVpcRouterParamDetail `json:"removeDnsFromVpcRouter"`
}
// DeleteHybridEipFromLocalParamDetail DeleteHybridEipFromLocal detail param
type DeleteHybridEipFromLocalParamDetail struct {
	Type string `json:"type" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteHybridEipFromLocalParam DeleteHybridEipFromLocal request param
type DeleteHybridEipFromLocalParam struct {
	BaseParam
	Params DeleteHybridEipFromLocalParamDetail `json:"deleteHybridEipFromLocal"`
}
// GetAvailableTriggersParamDetail GetAvailableTriggers detail param
type GetAvailableTriggersParamDetail struct {
}

// GetAvailableTriggersParam GetAvailableTriggers request param
type GetAvailableTriggersParam struct {
	BaseParam
	Params GetAvailableTriggersParamDetail `json:"getAvailableTriggers"`
}
// ReimageVmInstanceParamDetail ReimageVmInstance detail param
type ReimageVmInstanceParamDetail struct {
}

// ReimageVmInstanceParam ReimageVmInstance request param
type ReimageVmInstanceParam struct {
	BaseParam
	Params ReimageVmInstanceParamDetail `json:"reimageVmInstance"`
}
// UpdateDatasetsParamDetail UpdateDatasets detail param
type UpdateDatasetsParamDetail struct {
	UpdateDatasetStructs []UpdateDatasets_UpdateDatasetStructParam `json:"updateDatasetStructs" validate:"required"`
}

// UpdateDatasetsParam UpdateDatasets request param
type UpdateDatasetsParam struct {
	BaseParam
	Params UpdateDatasetsParamDetail `json:"updateDatasets"`
}
// SyncEcsSecurityGroupRuleFromRemoteParamDetail SyncEcsSecurityGroupRuleFromRemote detail param
type SyncEcsSecurityGroupRuleFromRemoteParamDetail struct {
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsSecurityGroupRuleFromRemoteParam SyncEcsSecurityGroupRuleFromRemote request param
type SyncEcsSecurityGroupRuleFromRemoteParam struct {
	BaseParam
	Params SyncEcsSecurityGroupRuleFromRemoteParamDetail `json:"syncEcsSecurityGroupRuleFromRemote"`
}
// SyncIdentityFromRemoteParamDetail SyncIdentityFromRemote detail param
type SyncIdentityFromRemoteParamDetail struct {
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncIdentityFromRemoteParam SyncIdentityFromRemote request param
type SyncIdentityFromRemoteParam struct {
	BaseParam
	Params SyncIdentityFromRemoteParamDetail `json:"syncIdentityFromRemote"`
}
// SetImageStoreBackupStorageQuotaParamDetail SetImageStoreBackupStorageQuota detail param
type SetImageStoreBackupStorageQuotaParamDetail struct {
	Uuids []string `json:"uuids,omitempty"`
	MaxCapacity int64 `json:"maxCapacity" validate:"required"`
}

// SetImageStoreBackupStorageQuotaParam SetImageStoreBackupStorageQuota request param
type SetImageStoreBackupStorageQuotaParam struct {
	BaseParam
	Params SetImageStoreBackupStorageQuotaParamDetail `json:"setImageStoreBackupStorageQuota"`
}
// ChangeClusterStateParamDetail ChangeClusterState detail param
type ChangeClusterStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeClusterStateParam ChangeClusterState request param
type ChangeClusterStateParam struct {
	BaseParam
	Params ChangeClusterStateParamDetail `json:"changeClusterState"`
}
// ChangeVfNicHaStateParamDetail ChangeVfNicHaState detail param
type ChangeVfNicHaStateParamDetail struct {
	HaState string `json:"haState" validate:"required"`
}

// ChangeVfNicHaStateParam ChangeVfNicHaState request param
type ChangeVfNicHaStateParam struct {
	BaseParam
	Params ChangeVfNicHaStateParamDetail `json:"changeVfNicHaState"`
}
// CreateOvnControllerOfferingParamDetail CreateOvnControllerOffering detail param
type CreateOvnControllerOfferingParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ManagementNetworkUuid string `json:"managementNetworkUuid" validate:"required"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	CpuNum int `json:"cpuNum" validate:"required"`
	MemorySize int64 `json:"memorySize" validate:"required"`
	ReservedMemorySize *int64 `json:"reservedMemorySize,omitempty"`
	AllocatorStrategy *string `json:"allocatorStrategy,omitempty"`
	SortKey int `json:"sortKey,omitempty"`
	Type *string `json:"type,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateOvnControllerOfferingParam CreateOvnControllerOffering request param
type CreateOvnControllerOfferingParam struct {
	BaseParam
	Params CreateOvnControllerOfferingParamDetail `json:"params"`
}
// GetIAM2OrganizationVirtualIDNumberParamDetail GetIAM2OrganizationVirtualIDNumber detail param
type GetIAM2OrganizationVirtualIDNumberParamDetail struct {
}

// GetIAM2OrganizationVirtualIDNumberParam GetIAM2OrganizationVirtualIDNumber request param
type GetIAM2OrganizationVirtualIDNumberParam struct {
	BaseParam
	Params GetIAM2OrganizationVirtualIDNumberParamDetail `json:"getIAM2OrganizationVirtualIDNumber"`
}
// DeleteEcsInstanceLocalParamDetail DeleteEcsInstanceLocal detail param
type DeleteEcsInstanceLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsInstanceLocalParam DeleteEcsInstanceLocal request param
type DeleteEcsInstanceLocalParam struct {
	BaseParam
	Params DeleteEcsInstanceLocalParamDetail `json:"deleteEcsInstanceLocal"`
}
// ChangePortMirrorStateParamDetail ChangePortMirrorState detail param
type ChangePortMirrorStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangePortMirrorStateParam ChangePortMirrorState request param
type ChangePortMirrorStateParam struct {
	BaseParam
	Params ChangePortMirrorStateParamDetail `json:"changePortMirrorState"`
}
// UnsubscribeSNSTopicParamDetail UnsubscribeSNSTopic detail param
type UnsubscribeSNSTopicParamDetail struct {
}

// UnsubscribeSNSTopicParam UnsubscribeSNSTopic request param
type UnsubscribeSNSTopicParam struct {
	BaseParam
	Params UnsubscribeSNSTopicParamDetail `json:"unsubscribeSNSTopic"`
}
// SetNicQosParamDetail SetNicQos detail param
type SetNicQosParamDetail struct {
	OutboundBandwidth *int64 `json:"outboundBandwidth,omitempty"`
	InboundBandwidth *int64 `json:"inboundBandwidth,omitempty"`
}

// SetNicQosParam SetNicQos request param
type SetNicQosParam struct {
	BaseParam
	Params SetNicQosParamDetail `json:"setNicQos"`
}
// CancelLongJobParamDetail CancelLongJob detail param
type CancelLongJobParamDetail struct {
}

// CancelLongJobParam CancelLongJob request param
type CancelLongJobParam struct {
	BaseParam
	Params CancelLongJobParamDetail `json:"cancelLongJob"`
}
// GetRouteTableVpcVRouterCandidateParamDetail GetRouteTableVpcVRouterCandidate detail param
type GetRouteTableVpcVRouterCandidateParamDetail struct {
	TableUuid *string `json:"tableUuid,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetRouteTableVpcVRouterCandidateParam GetRouteTableVpcVRouterCandidate request param
type GetRouteTableVpcVRouterCandidateParam struct {
	BaseParam
	Params GetRouteTableVpcVRouterCandidateParamDetail `json:"getRouteTableVpcVRouterCandidate"`
}
// GenerateAccountBillingParamDetail GenerateAccountBilling detail param
type GenerateAccountBillingParamDetail struct {
}

// GenerateAccountBillingParam GenerateAccountBilling request param
type GenerateAccountBillingParam struct {
	BaseParam
	Params GenerateAccountBillingParamDetail `json:"generateAccountBilling"`
}
// SyncAliyunVirtualRouterFromRemoteParamDetail SyncAliyunVirtualRouterFromRemote detail param
type SyncAliyunVirtualRouterFromRemoteParamDetail struct {
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncAliyunVirtualRouterFromRemoteParam SyncAliyunVirtualRouterFromRemote request param
type SyncAliyunVirtualRouterFromRemoteParam struct {
	BaseParam
	Params SyncAliyunVirtualRouterFromRemoteParamDetail `json:"syncAliyunVirtualRouterFromRemote"`
}
// GetInvocationRecordsParamDetail GetInvocationRecords detail param
type GetInvocationRecordsParamDetail struct {
	RecordUuid *string `json:"recordUuid,omitempty"`
	VmInstanceUuid *string `json:"vmInstanceUuid,omitempty"`
	IncludeOutput *bool `json:"includeOutput,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetInvocationRecordsParam GetInvocationRecords request param
type GetInvocationRecordsParam struct {
	BaseParam
	Params GetInvocationRecordsParamDetail `json:"getInvocationRecords"`
}
// ChangeRoleStateParamDetail ChangeRoleState detail param
type ChangeRoleStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeRoleStateParam ChangeRoleState request param
type ChangeRoleStateParam struct {
	BaseParam
	Params ChangeRoleStateParamDetail `json:"changeRoleState"`
}
// GetVRouterFlowCounterParamDetail GetVRouterFlowCounter detail param
type GetVRouterFlowCounterParamDetail struct {
}

// GetVRouterFlowCounterParam GetVRouterFlowCounter request param
type GetVRouterFlowCounterParam struct {
	BaseParam
	Params GetVRouterFlowCounterParamDetail `json:"getVRouterFlowCounter"`
}
// CreateAliyunRouterInterfaceRemoteParamDetail CreateAliyunRouterInterfaceRemote detail param
type CreateAliyunRouterInterfaceRemoteParamDetail struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	AccessPointUuid *string `json:"accessPointUuid,omitempty"`
	Spec *string `json:"spec,omitempty"`
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	RouterType string `json:"routerType" validate:"required"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunRouterInterfaceRemoteParam CreateAliyunRouterInterfaceRemote request param
type CreateAliyunRouterInterfaceRemoteParam struct {
	BaseParam
	Params CreateAliyunRouterInterfaceRemoteParamDetail `json:"params"`
}
// GetBareMetal2SupportedBootModeParamDetail GetBareMetal2SupportedBootMode detail param
type GetBareMetal2SupportedBootModeParamDetail struct {
}

// GetBareMetal2SupportedBootModeParam GetBareMetal2SupportedBootMode request param
type GetBareMetal2SupportedBootModeParam struct {
	BaseParam
	Params GetBareMetal2SupportedBootModeParamDetail `json:"getBareMetal2SupportedBootMode"`
}
// GetHostPowerStatusParamDetail GetHostPowerStatus detail param
type GetHostPowerStatusParamDetail struct {
	Method *string `json:"method,omitempty"`
}

// GetHostPowerStatusParam GetHostPowerStatus request param
type GetHostPowerStatusParam struct {
	BaseParam
	Params GetHostPowerStatusParamDetail `json:"getHostPowerStatus"`
}
// GetChainTaskParamDetail GetChainTask detail param
type GetChainTaskParamDetail struct {
	SyncSignatures []string `json:"syncSignatures,omitempty"`
}

// GetChainTaskParam GetChainTask request param
type GetChainTaskParam struct {
	BaseParam
	Params GetChainTaskParamDetail `json:"getChainTask"`
}
// UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParamDetail UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch detail param
type UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParam UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch request param
type UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParam struct {
	BaseParam
	Params UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParamDetail `json:"updateConnectionBetweenL3NetWorkAndAliyunVSwitch"`
}
// ChangeHostPasswordParamDetail ChangeHostPassword detail param
type ChangeHostPasswordParamDetail struct {
	Password string `json:"password" validate:"required"`
}

// ChangeHostPasswordParam ChangeHostPassword request param
type ChangeHostPasswordParam struct {
	BaseParam
	Params ChangeHostPasswordParamDetail `json:"changeHostPassword"`
}
// CreateSlbInstanceParamDetail CreateSlbInstance detail param
type CreateSlbInstanceParamDetail struct {
	Name string `json:"name" validate:"required"`
	SlbGroupUuid string `json:"slbGroupUuid" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSlbInstanceParam CreateSlbInstance request param
type CreateSlbInstanceParam struct {
	BaseParam
	Params CreateSlbInstanceParamDetail `json:"params"`
}
// ChangePortForwardingRuleStateParamDetail ChangePortForwardingRuleState detail param
type ChangePortForwardingRuleStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangePortForwardingRuleStateParam ChangePortForwardingRuleState request param
type ChangePortForwardingRuleStateParam struct {
	BaseParam
	Params ChangePortForwardingRuleStateParamDetail `json:"changePortForwardingRuleState"`
}
// IsLicenseServerParamDetail IsLicenseServer detail param
type IsLicenseServerParamDetail struct {
}

// IsLicenseServerParam IsLicenseServer request param
type IsLicenseServerParam struct {
	BaseParam
	Params IsLicenseServerParamDetail `json:"isLicenseServer"`
}
// PrometheusQueryLabelValuesParamDetail PrometheusQueryLabelValues detail param
type PrometheusQueryLabelValuesParamDetail struct {
	Labels []string `json:"labels" validate:"required"`
}

// PrometheusQueryLabelValuesParam PrometheusQueryLabelValues request param
type PrometheusQueryLabelValuesParam struct {
	BaseParam
	Params PrometheusQueryLabelValuesParamDetail `json:"prometheusQueryLabelValues"`
}
// ValidateClusterSupportDRSParamDetail ValidateClusterSupportDRS detail param
type ValidateClusterSupportDRSParamDetail struct {
}

// ValidateClusterSupportDRSParam ValidateClusterSupportDRS request param
type ValidateClusterSupportDRSParam struct {
	BaseParam
	Params ValidateClusterSupportDRSParamDetail `json:"validateClusterSupportDRS"`
}
// ShrinkVolumeSnapshotParamDetail ShrinkVolumeSnapshot detail param
type ShrinkVolumeSnapshotParamDetail struct {
}

// ShrinkVolumeSnapshotParam ShrinkVolumeSnapshot request param
type ShrinkVolumeSnapshotParam struct {
	BaseParam
	Params ShrinkVolumeSnapshotParamDetail `json:"shrinkVolumeSnapshot"`
}
// AddHostToHostSchedulingRuleGroupParamDetail AddHostToHostSchedulingRuleGroup detail param
type AddHostToHostSchedulingRuleGroupParamDetail struct {
}

// AddHostToHostSchedulingRuleGroupParam AddHostToHostSchedulingRuleGroup request param
type AddHostToHostSchedulingRuleGroupParam struct {
	BaseParam
	Params AddHostToHostSchedulingRuleGroupParamDetail `json:"addHostToHostSchedulingRuleGroup"`
}
// CreateBuildAppParamDetail CreateBuildApp detail param
type CreateBuildAppParamDetail struct {
	BuildSystemUuid string `json:"buildSystemUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	DataPath string `json:"dataPath" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBuildAppParam CreateBuildApp request param
type CreateBuildAppParam struct {
	BaseParam
	Params CreateBuildAppParamDetail `json:"params"`
}
// GetVmNicAttachedNetworkServiceParamDetail GetVmNicAttachedNetworkService detail param
type GetVmNicAttachedNetworkServiceParamDetail struct {
}

// GetVmNicAttachedNetworkServiceParam GetVmNicAttachedNetworkService request param
type GetVmNicAttachedNetworkServiceParam struct {
	BaseParam
	Params GetVmNicAttachedNetworkServiceParamDetail `json:"getVmNicAttachedNetworkService"`
}
// GetVmHostnameParamDetail GetVmHostname detail param
type GetVmHostnameParamDetail struct {
}

// GetVmHostnameParam GetVmHostname request param
type GetVmHostnameParam struct {
	BaseParam
	Params GetVmHostnameParamDetail `json:"getVmHostname"`
}
// AddSchedulerJobsToSchedulerJobGroupParamDetail AddSchedulerJobsToSchedulerJobGroup detail param
type AddSchedulerJobsToSchedulerJobGroupParamDetail struct {
	SchedulerJobUuids []string `json:"schedulerJobUuids" validate:"required"`
}

// AddSchedulerJobsToSchedulerJobGroupParam AddSchedulerJobsToSchedulerJobGroup request param
type AddSchedulerJobsToSchedulerJobGroupParam struct {
	BaseParam
	Params AddSchedulerJobsToSchedulerJobGroupParamDetail `json:"params"`
}
// DetachL3NetworkFromVmParamDetail DetachL3NetworkFromVm detail param
type DetachL3NetworkFromVmParamDetail struct {
}

// DetachL3NetworkFromVmParam DetachL3NetworkFromVm request param
type DetachL3NetworkFromVmParam struct {
	BaseParam
	Params DetachL3NetworkFromVmParamDetail `json:"detachL3NetworkFromVm"`
}
// DeleteVpcUserVpnGatewayLocalParamDetail DeleteVpcUserVpnGatewayLocal detail param
type DeleteVpcUserVpnGatewayLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVpcUserVpnGatewayLocalParam DeleteVpcUserVpnGatewayLocal request param
type DeleteVpcUserVpnGatewayLocalParam struct {
	BaseParam
	Params DeleteVpcUserVpnGatewayLocalParamDetail `json:"deleteVpcUserVpnGatewayLocal"`
}
// CreateVRouterOspfAreaParamDetail CreateVRouterOspfArea detail param
type CreateVRouterOspfAreaParamDetail struct {
	AreaId string `json:"areaId" validate:"required"`
	AreaAuth *string `json:"areaAuth,omitempty"`
	AreaType *string `json:"areaType,omitempty"`
	Password *string `json:"password,omitempty"`
	KeyId *int `json:"keyId,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVRouterOspfAreaParam CreateVRouterOspfArea request param
type CreateVRouterOspfAreaParam struct {
	BaseParam
	Params CreateVRouterOspfAreaParamDetail `json:"params"`
}
// SetSecurityMachineKeyParamDetail SetSecurityMachineKey detail param
type SetSecurityMachineKeyParamDetail struct {
	Type string `json:"type" validate:"required"`
	TokenName string `json:"tokenName" validate:"required"`
	DryRun *bool `json:"dryRun,omitempty"`
}

// SetSecurityMachineKeyParam SetSecurityMachineKey request param
type SetSecurityMachineKeyParam struct {
	BaseParam
	Params SetSecurityMachineKeyParamDetail `json:"params"`
}
// CreateOAuthClientParamDetail CreateOAuthClient detail param
type CreateOAuthClientParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ClientId string `json:"clientId" validate:"required"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	AuthorizationUrl *string `json:"authorizationUrl,omitempty"`
	TokenUrl string `json:"tokenUrl" validate:"required"`
	UserinfoUrl *string `json:"userinfoUrl,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	LoginType string `json:"loginType" validate:"required"`
	GrantType string `json:"grantType" validate:"required"`
	IdentityProvider *string `json:"identityProvider,omitempty"`
	PluginUuid *string `json:"pluginUuid,omitempty"`
	UrlTemplate string `json:"urlTemplate" validate:"required"`
	ClientType string `json:"clientType" validate:"required"`
	ScopeList []string `json:"scopeList,omitempty"`
	Attributes []ExtendedAttributeParam `json:"attributes,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateOAuthClientParam CreateOAuthClient request param
type CreateOAuthClientParam struct {
	BaseParam
	Params CreateOAuthClientParamDetail `json:"params"`
}
// GetVpcAttachedEipParamDetail GetVpcAttachedEip detail param
type GetVpcAttachedEipParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVpcAttachedEipParam GetVpcAttachedEip request param
type GetVpcAttachedEipParam struct {
	BaseParam
	Params GetVpcAttachedEipParamDetail `json:"params"`
}
// RemoveSchedulerJobFromSchedulerTriggerParamDetail RemoveSchedulerJobFromSchedulerTrigger detail param
type RemoveSchedulerJobFromSchedulerTriggerParamDetail struct {
}

// RemoveSchedulerJobFromSchedulerTriggerParam RemoveSchedulerJobFromSchedulerTrigger request param
type RemoveSchedulerJobFromSchedulerTriggerParam struct {
	BaseParam
	Params RemoveSchedulerJobFromSchedulerTriggerParamDetail `json:"removeSchedulerJobFromSchedulerTrigger"`
}
// ChangeMediaStateParamDetail ChangeMediaState detail param
type ChangeMediaStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeMediaStateParam ChangeMediaState request param
type ChangeMediaStateParam struct {
	BaseParam
	Params ChangeMediaStateParamDetail `json:"changeMediaState"`
}
// ChangeIPSecConnectionStateParamDetail ChangeIPSecConnectionState detail param
type ChangeIPSecConnectionStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeIPSecConnectionStateParam ChangeIPSecConnectionState request param
type ChangeIPSecConnectionStateParam struct {
	BaseParam
	Params ChangeIPSecConnectionStateParamDetail `json:"changeIPSecConnectionState"`
}
// StopAllResourcesInIAM2ProjectParamDetail StopAllResourcesInIAM2Project detail param
type StopAllResourcesInIAM2ProjectParamDetail struct {
}

// StopAllResourcesInIAM2ProjectParam StopAllResourcesInIAM2Project request param
type StopAllResourcesInIAM2ProjectParam struct {
	BaseParam
	Params StopAllResourcesInIAM2ProjectParamDetail `json:"stopAllResourcesInIAM2Project"`
}
// UpdateVmNetworkConfigParamDetail UpdateVmNetworkConfig detail param
type UpdateVmNetworkConfigParamDetail struct {
	VmNicUuids []string `json:"vmNicUuids" validate:"required"`
}

// UpdateVmNetworkConfigParam UpdateVmNetworkConfig request param
type UpdateVmNetworkConfigParam struct {
	BaseParam
	Params UpdateVmNetworkConfigParamDetail `json:"updateVmNetworkConfig"`
}
// RemoveTicketTypesFromTicketFlowCollectionParamDetail RemoveTicketTypesFromTicketFlowCollection detail param
type RemoveTicketTypesFromTicketFlowCollectionParamDetail struct {
	TicketTypeUuids []string `json:"ticketTypeUuids" validate:"required"`
}

// RemoveTicketTypesFromTicketFlowCollectionParam RemoveTicketTypesFromTicketFlowCollection request param
type RemoveTicketTypesFromTicketFlowCollectionParam struct {
	BaseParam
	Params RemoveTicketTypesFromTicketFlowCollectionParamDetail `json:"removeTicketTypesFromTicketFlowCollection"`
}
// DeleteEcsVSwitchRemoteParamDetail DeleteEcsVSwitchRemote detail param
type DeleteEcsVSwitchRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsVSwitchRemoteParam DeleteEcsVSwitchRemote request param
type DeleteEcsVSwitchRemoteParam struct {
	BaseParam
	Params DeleteEcsVSwitchRemoteParamDetail `json:"deleteEcsVSwitchRemote"`
}
// SetVmStaticIpParamDetail SetVmStaticIp detail param
type SetVmStaticIpParamDetail struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Ip *string `json:"ip,omitempty"`
	Ip6 *string `json:"ip6,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	Ipv6Gateway *string `json:"ipv6Gateway,omitempty"`
	Ipv6Prefix *string `json:"ipv6Prefix,omitempty"`
}

// SetVmStaticIpParam SetVmStaticIp request param
type SetVmStaticIpParam struct {
	BaseParam
	Params SetVmStaticIpParamDetail `json:"setVmStaticIp"`
}
// GetVmSshKeyParamDetail GetVmSshKey detail param
type GetVmSshKeyParamDetail struct {
}

// GetVmSshKeyParam GetVmSshKey request param
type GetVmSshKeyParam struct {
	BaseParam
	Params GetVmSshKeyParamDetail `json:"getVmSshKey"`
}
// GetVmGuestToolsInfoParamDetail GetVmGuestToolsInfo detail param
type GetVmGuestToolsInfoParamDetail struct {
	Debug []string `json:"debug,omitempty"`
}

// GetVmGuestToolsInfoParam GetVmGuestToolsInfo request param
type GetVmGuestToolsInfoParam struct {
	BaseParam
	Params GetVmGuestToolsInfoParamDetail `json:"getVmGuestToolsInfo"`
}
// ValidateDiskOfferingUserConfigParamDetail ValidateDiskOfferingUserConfig detail param
type ValidateDiskOfferingUserConfigParamDetail struct {
	Config string `json:"config" validate:"required"`
}

// ValidateDiskOfferingUserConfigParam ValidateDiskOfferingUserConfig request param
type ValidateDiskOfferingUserConfigParam struct {
	BaseParam
	Params ValidateDiskOfferingUserConfigParamDetail `json:"validateDiskOfferingUserConfig"`
}
// DeleteVpcVpnGatewayLocalParamDetail DeleteVpcVpnGatewayLocal detail param
type DeleteVpcVpnGatewayLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVpcVpnGatewayLocalParam DeleteVpcVpnGatewayLocal request param
type DeleteVpcVpnGatewayLocalParam struct {
	BaseParam
	Params DeleteVpcVpnGatewayLocalParamDetail `json:"deleteVpcVpnGatewayLocal"`
}
// SetVmRDPParamDetail SetVmRDP detail param
type SetVmRDPParamDetail struct {
	Enable bool `json:"enable" validate:"required"`
}

// SetVmRDPParam SetVmRDP request param
type SetVmRDPParam struct {
	BaseParam
	Params SetVmRDPParamDetail `json:"setVmRDP"`
}
// RunSchedulerTriggerParamDetail RunSchedulerTrigger detail param
type RunSchedulerTriggerParamDetail struct {
	JobUuids []string `json:"jobUuids,omitempty"`
}

// RunSchedulerTriggerParam RunSchedulerTrigger request param
type RunSchedulerTriggerParam struct {
	BaseParam
	Params RunSchedulerTriggerParamDetail `json:"runSchedulerTrigger"`
}
// CreateAliyunVpcVirtualRouterEntryRemoteParamDetail CreateAliyunVpcVirtualRouterEntryRemote detail param
type CreateAliyunVpcVirtualRouterEntryRemoteParamDetail struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	DstCidrBlock string `json:"dstCidrBlock" validate:"required"`
	NextHopUuid string `json:"nextHopUuid" validate:"required"`
	NextHopType string `json:"nextHopType" validate:"required"`
	VRouterType string `json:"vRouterType" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunVpcVirtualRouterEntryRemoteParam CreateAliyunVpcVirtualRouterEntryRemote request param
type CreateAliyunVpcVirtualRouterEntryRemoteParam struct {
	BaseParam
	Params CreateAliyunVpcVirtualRouterEntryRemoteParamDetail `json:"params"`
}
// PowerOnHostParamDetail PowerOnHost detail param
type PowerOnHostParamDetail struct {
	ReturnEarly *bool `json:"returnEarly,omitempty"`
}

// PowerOnHostParam PowerOnHost request param
type PowerOnHostParam struct {
	BaseParam
	Params PowerOnHostParamDetail `json:"powerOnHost"`
}
// DeleteAliyunSnapshotFromRemoteParamDetail DeleteAliyunSnapshotFromRemote detail param
type DeleteAliyunSnapshotFromRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAliyunSnapshotFromRemoteParam DeleteAliyunSnapshotFromRemote request param
type DeleteAliyunSnapshotFromRemoteParam struct {
	BaseParam
	Params DeleteAliyunSnapshotFromRemoteParamDetail `json:"deleteAliyunSnapshotFromRemote"`
}
// RemoveCertificateFromLoadBalancerListenerParamDetail RemoveCertificateFromLoadBalancerListener detail param
type RemoveCertificateFromLoadBalancerListenerParamDetail struct {
	CertificateUuid string `json:"certificateUuid" validate:"required"`
}

// RemoveCertificateFromLoadBalancerListenerParam RemoveCertificateFromLoadBalancerListener request param
type RemoveCertificateFromLoadBalancerListenerParam struct {
	BaseParam
	Params RemoveCertificateFromLoadBalancerListenerParamDetail `json:"removeCertificateFromLoadBalancerListener"`
}
// GetPortForwardingAttachableVmNicsParamDetail GetPortForwardingAttachableVmNics detail param
type GetPortForwardingAttachableVmNicsParamDetail struct {
}

// GetPortForwardingAttachableVmNicsParam GetPortForwardingAttachableVmNics request param
type GetPortForwardingAttachableVmNicsParam struct {
	BaseParam
	Params GetPortForwardingAttachableVmNicsParamDetail `json:"getPortForwardingAttachableVmNics"`
}
// RemoveRendezvousPointFromMulticastRouterParamDetail RemoveRendezvousPointFromMulticastRouter detail param
type RemoveRendezvousPointFromMulticastRouterParamDetail struct {
	RpAddress string `json:"rpAddress" validate:"required"`
	GroupAddress string `json:"groupAddress" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveRendezvousPointFromMulticastRouterParam RemoveRendezvousPointFromMulticastRouter request param
type RemoveRendezvousPointFromMulticastRouterParam struct {
	BaseParam
	Params RemoveRendezvousPointFromMulticastRouterParamDetail `json:"removeRendezvousPointFromMulticastRouter"`
}
// AddIAM2VirtualIDsToProjectParamDetail AddIAM2VirtualIDsToProject detail param
type AddIAM2VirtualIDsToProjectParamDetail struct {
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
	RoleUuids []string `json:"roleUuids,omitempty"`
}

// AddIAM2VirtualIDsToProjectParam AddIAM2VirtualIDsToProject request param
type AddIAM2VirtualIDsToProjectParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToProjectParamDetail `json:"params"`
}
// SubscribeEventParamDetail SubscribeEvent detail param
type SubscribeEventParamDetail struct {
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace" validate:"required"`
	EventName string `json:"eventName" validate:"required"`
	Actions []CreateAlarm_ActionParamParam `json:"actions,omitempty"`
	Labels []LabelParam `json:"labels,omitempty"`
	EmergencyLevel *string `json:"emergencyLevel,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SubscribeEventParam SubscribeEvent request param
type SubscribeEventParam struct {
	BaseParam
	Params SubscribeEventParamDetail `json:"params"`
}
// GetPrimaryStorageCandidatesForVolumeMigrationParamDetail GetPrimaryStorageCandidatesForVolumeMigration detail param
type GetPrimaryStorageCandidatesForVolumeMigrationParamDetail struct {
}

// GetPrimaryStorageCandidatesForVolumeMigrationParam GetPrimaryStorageCandidatesForVolumeMigration request param
type GetPrimaryStorageCandidatesForVolumeMigrationParam struct {
	BaseParam
	Params GetPrimaryStorageCandidatesForVolumeMigrationParamDetail `json:"getPrimaryStorageCandidatesForVolumeMigration"`
}
// UpgradeBackupStorageCdpTasksParamDetail UpgradeBackupStorageCdpTasks detail param
type UpgradeBackupStorageCdpTasksParamDetail struct {
}

// UpgradeBackupStorageCdpTasksParam UpgradeBackupStorageCdpTasks request param
type UpgradeBackupStorageCdpTasksParam struct {
	BaseParam
	Params UpgradeBackupStorageCdpTasksParamDetail `json:"upgradeBackupStorageCdpTasks"`
}
// DeleteVxlanL2NetworkParamDetail DeleteVxlanL2Network detail param
type DeleteVxlanL2NetworkParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVxlanL2NetworkParam DeleteVxlanL2Network request param
type DeleteVxlanL2NetworkParam struct {
	BaseParam
	Params DeleteVxlanL2NetworkParamDetail `json:"deleteVxlanL2Network"`
}
// RemoveVmFromAffinityGroupParamDetail RemoveVmFromAffinityGroup detail param
type RemoveVmFromAffinityGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RemoveVmFromAffinityGroupParam RemoveVmFromAffinityGroup request param
type RemoveVmFromAffinityGroupParam struct {
	BaseParam
	Params RemoveVmFromAffinityGroupParamDetail `json:"removeVmFromAffinityGroup"`
}
// SetVolumeIoThreadPinParamDetail SetVolumeIoThreadPin detail param
type SetVolumeIoThreadPinParamDetail struct {
	VmUuid string `json:"vmUuid" validate:"required"`
	Pin string `json:"pin" validate:"required"`
	IoThreadId int `json:"ioThreadId" validate:"required"`
}

// SetVolumeIoThreadPinParam SetVolumeIoThreadPin request param
type SetVolumeIoThreadPinParam struct {
	BaseParam
	Params SetVolumeIoThreadPinParamDetail `json:"setVolumeIoThreadPin"`
}
// UpdatePriorityConfigParamDetail UpdatePriorityConfig detail param
type UpdatePriorityConfigParamDetail struct {
	CpuShares *int `json:"cpuShares,omitempty"`
	OomScoreAdj *int `json:"oomScoreAdj,omitempty"`
}

// UpdatePriorityConfigParam UpdatePriorityConfig request param
type UpdatePriorityConfigParam struct {
	BaseParam
	Params UpdatePriorityConfigParamDetail `json:"updatePriorityConfig"`
}
// IdentifyHostParamDetail IdentifyHost detail param
type IdentifyHostParamDetail struct {
	Interval *int64 `json:"interval,omitempty"`
}

// IdentifyHostParam IdentifyHost request param
type IdentifyHostParam struct {
	BaseParam
	Params IdentifyHostParamDetail `json:"identifyHost"`
}
// CreateRootVolumeTemplateFromVolumeBackupParamDetail CreateRootVolumeTemplateFromVolumeBackup detail param
type CreateRootVolumeTemplateFromVolumeBackupParamDetail struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	System bool `json:"system,omitempty"`
	Virtio *bool `json:"virtio,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateRootVolumeTemplateFromVolumeBackupParam CreateRootVolumeTemplateFromVolumeBackup request param
type CreateRootVolumeTemplateFromVolumeBackupParam struct {
	BaseParam
	Params CreateRootVolumeTemplateFromVolumeBackupParamDetail `json:"params"`
}
// CheckFirewallRuleConfigFileParamDetail CheckFirewallRuleConfigFile detail param
type CheckFirewallRuleConfigFileParamDetail struct {
	RuleInfo string `json:"ruleInfo" validate:"required"`
}

// CheckFirewallRuleConfigFileParam CheckFirewallRuleConfigFile request param
type CheckFirewallRuleConfigFileParam struct {
	BaseParam
	Params CheckFirewallRuleConfigFileParamDetail `json:"params"`
}
// GetVmConsoleAddressParamDetail GetVmConsoleAddress detail param
type GetVmConsoleAddressParamDetail struct {
}

// GetVmConsoleAddressParam GetVmConsoleAddress request param
type GetVmConsoleAddressParam struct {
	BaseParam
	Params GetVmConsoleAddressParamDetail `json:"getVmConsoleAddress"`
}
// GetLoadBalancerListenerACLEntriesParamDetail GetLoadBalancerListenerACLEntries detail param
type GetLoadBalancerListenerACLEntriesParamDetail struct {
	ListenerUuids []string `json:"listenerUuids,omitempty"`
	Type *string `json:"type,omitempty"`
}

// GetLoadBalancerListenerACLEntriesParam GetLoadBalancerListenerACLEntries request param
type GetLoadBalancerListenerACLEntriesParam struct {
	BaseParam
	Params GetLoadBalancerListenerACLEntriesParamDetail `json:"getLoadBalancerListenerACLEntries"`
}
// UpdateHostIommuStateParamDetail UpdateHostIommuState detail param
type UpdateHostIommuStateParamDetail struct {
	State string `json:"state" validate:"required"`
}

// UpdateHostIommuStateParam UpdateHostIommuState request param
type UpdateHostIommuStateParam struct {
	BaseParam
	Params UpdateHostIommuStateParamDetail `json:"updateHostIommuState"`
}
// UnsubscribeEventParamDetail UnsubscribeEvent detail param
type UnsubscribeEventParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// UnsubscribeEventParam UnsubscribeEvent request param
type UnsubscribeEventParam struct {
	BaseParam
	Params UnsubscribeEventParamDetail `json:"unsubscribeEvent"`
}
// CreateObservabilityServerParamDetail CreateObservabilityServer detail param
type CreateObservabilityServerParamDetail struct {
	Name string `json:"name" validate:"required"`
	ObservabilityServerOfferingUuid string `json:"observabilityServerOfferingUuid" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateObservabilityServerParam CreateObservabilityServer request param
type CreateObservabilityServerParam struct {
	BaseParam
	Params CreateObservabilityServerParamDetail `json:"params"`
}
// RemoveMonFromCephPrimaryStorageParamDetail RemoveMonFromCephPrimaryStorage detail param
type RemoveMonFromCephPrimaryStorageParamDetail struct {
	MonHostnames []string `json:"monHostnames" validate:"required"`
}

// RemoveMonFromCephPrimaryStorageParam RemoveMonFromCephPrimaryStorage request param
type RemoveMonFromCephPrimaryStorageParam struct {
	BaseParam
	Params RemoveMonFromCephPrimaryStorageParamDetail `json:"removeMonFromCephPrimaryStorage"`
}
// GetVmsSchedulingStateFromSchedulingRuleParamDetail GetVmsSchedulingStateFromSchedulingRule detail param
type GetVmsSchedulingStateFromSchedulingRuleParamDetail struct {
	RuleUuid string `json:"ruleUuid" validate:"required"`
	VmUuids []string `json:"vmUuids" validate:"required"`
}

// GetVmsSchedulingStateFromSchedulingRuleParam GetVmsSchedulingStateFromSchedulingRule request param
type GetVmsSchedulingStateFromSchedulingRuleParam struct {
	BaseParam
	Params GetVmsSchedulingStateFromSchedulingRuleParamDetail `json:"params"`
}
// ChangeAlarmStateParamDetail ChangeAlarmState detail param
type ChangeAlarmStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeAlarmStateParam ChangeAlarmState request param
type ChangeAlarmStateParam struct {
	BaseParam
	Params ChangeAlarmStateParamDetail `json:"changeAlarmState"`
}
// GetLocalStorageHostDiskCapacityParamDetail GetLocalStorageHostDiskCapacity detail param
type GetLocalStorageHostDiskCapacityParamDetail struct {
	HostUuid *string `json:"hostUuid,omitempty"`
}

// GetLocalStorageHostDiskCapacityParam GetLocalStorageHostDiskCapacity request param
type GetLocalStorageHostDiskCapacityParam struct {
	BaseParam
	Params GetLocalStorageHostDiskCapacityParamDetail `json:"getLocalStorageHostDiskCapacity"`
}
// DeleteVmSshKeyParamDetail DeleteVmSshKey detail param
type DeleteVmSshKeyParamDetail struct {
}

// DeleteVmSshKeyParam DeleteVmSshKey request param
type DeleteVmSshKeyParam struct {
	BaseParam
	Params DeleteVmSshKeyParamDetail `json:"deleteVmSshKey"`
}
// GetPolicyRouteRuleSetFromVirtualRouterParamDetail GetPolicyRouteRuleSetFromVirtualRouter detail param
type GetPolicyRouteRuleSetFromVirtualRouterParamDetail struct {
}

// GetPolicyRouteRuleSetFromVirtualRouterParam GetPolicyRouteRuleSetFromVirtualRouter request param
type GetPolicyRouteRuleSetFromVirtualRouterParam struct {
	BaseParam
	Params GetPolicyRouteRuleSetFromVirtualRouterParamDetail `json:"getPolicyRouteRuleSetFromVirtualRouter"`
}
// DeleteVxlanPoolRemoteVtepParamDetail DeleteVxlanPoolRemoteVtep detail param
type DeleteVxlanPoolRemoteVtepParamDetail struct {
	RemoteVtepIp string `json:"remoteVtepIp" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVxlanPoolRemoteVtepParam DeleteVxlanPoolRemoteVtep request param
type DeleteVxlanPoolRemoteVtepParam struct {
	BaseParam
	Params DeleteVxlanPoolRemoteVtepParamDetail `json:"deleteVxlanPoolRemoteVtep"`
}
// RemoveAttributesFromIAM2ProjectParamDetail RemoveAttributesFromIAM2Project detail param
type RemoveAttributesFromIAM2ProjectParamDetail struct {
	AttributeUuids []string `json:"attributeUuids" validate:"required"`
}

// RemoveAttributesFromIAM2ProjectParam RemoveAttributesFromIAM2Project request param
type RemoveAttributesFromIAM2ProjectParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2ProjectParamDetail `json:"removeAttributesFromIAM2Project"`
}
// RecoverDataVolumeParamDetail RecoverDataVolume detail param
type RecoverDataVolumeParamDetail struct {
}

// RecoverDataVolumeParam RecoverDataVolume request param
type RecoverDataVolumeParam struct {
	BaseParam
	Params RecoverDataVolumeParamDetail `json:"recoverDataVolume"`
}
// RemoveIAM2VirtualIDsFromGroupParamDetail RemoveIAM2VirtualIDsFromGroup detail param
type RemoveIAM2VirtualIDsFromGroupParamDetail struct {
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
}

// RemoveIAM2VirtualIDsFromGroupParam RemoveIAM2VirtualIDsFromGroup request param
type RemoveIAM2VirtualIDsFromGroupParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromGroupParamDetail `json:"removeIAM2VirtualIDsFromGroup"`
}
// AttachBareMetal2ProvisionNetworkToClusterParamDetail AttachBareMetal2ProvisionNetworkToCluster detail param
type AttachBareMetal2ProvisionNetworkToClusterParamDetail struct {
}

// AttachBareMetal2ProvisionNetworkToClusterParam AttachBareMetal2ProvisionNetworkToCluster request param
type AttachBareMetal2ProvisionNetworkToClusterParam struct {
	BaseParam
	Params AttachBareMetal2ProvisionNetworkToClusterParamDetail `json:"params"`
}
// ProvisionSlbInstanceParamDetail ProvisionSlbInstance detail param
type ProvisionSlbInstanceParamDetail struct {
}

// ProvisionSlbInstanceParam ProvisionSlbInstance request param
type ProvisionSlbInstanceParam struct {
	BaseParam
	Params ProvisionSlbInstanceParamDetail `json:"provisionSlbInstance"`
}
// SetVmUserDefinedXmlHookScriptParamDetail SetVmUserDefinedXmlHookScript detail param
type SetVmUserDefinedXmlHookScriptParamDetail struct {
	XmlHookScriptBase64 string `json:"xmlHookScriptBase64" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SetVmUserDefinedXmlHookScriptParam SetVmUserDefinedXmlHookScript request param
type SetVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params SetVmUserDefinedXmlHookScriptParamDetail `json:"setVmUserDefinedXmlHookScript"`
}
// DetachProvisionNicFromBondingParamDetail DetachProvisionNicFromBonding detail param
type DetachProvisionNicFromBondingParamDetail struct {
	ProvisionNicUuid string `json:"provisionNicUuid" validate:"required"`
}

// DetachProvisionNicFromBondingParam DetachProvisionNicFromBonding request param
type DetachProvisionNicFromBondingParam struct {
	BaseParam
	Params DetachProvisionNicFromBondingParamDetail `json:"detachProvisionNicFromBonding"`
}
// GetHostAllocatorStrategiesParamDetail GetHostAllocatorStrategies detail param
type GetHostAllocatorStrategiesParamDetail struct {
}

// GetHostAllocatorStrategiesParam GetHostAllocatorStrategies request param
type GetHostAllocatorStrategiesParam struct {
	BaseParam
	Params GetHostAllocatorStrategiesParamDetail `json:"getHostAllocatorStrategies"`
}
// GetInterfaceServiceTypeStatisticParamDetail GetInterfaceServiceTypeStatistic detail param
type GetInterfaceServiceTypeStatisticParamDetail struct {
	InterfaceUuid *string `json:"interfaceUuid,omitempty"`
	VlanId *int `json:"vlanId,omitempty"`
	InterfaceType *string `json:"interfaceType,omitempty"`
	ServiceType []string `json:"serviceType,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	SortBy *string `json:"sortBy,omitempty"`
	SortDirection *string `json:"sortDirection,omitempty"`
	Start *int `json:"start,omitempty"`
	Limit *int `json:"limit,omitempty"`
	ReplyWithCount *bool `json:"replyWithCount,omitempty"`
}

// GetInterfaceServiceTypeStatisticParam GetInterfaceServiceTypeStatistic request param
type GetInterfaceServiceTypeStatisticParam struct {
	BaseParam
	Params GetInterfaceServiceTypeStatisticParamDetail `json:"getInterfaceServiceTypeStatistic"`
}
// StartConnectionBetweenAliyunRouterInterfaceParamDetail StartConnectionBetweenAliyunRouterInterface detail param
type StartConnectionBetweenAliyunRouterInterfaceParamDetail struct {
	VrouterInterfaceUuid string `json:"vrouterInterfaceUuid" validate:"required"`
}

// StartConnectionBetweenAliyunRouterInterfaceParam StartConnectionBetweenAliyunRouterInterface request param
type StartConnectionBetweenAliyunRouterInterfaceParam struct {
	BaseParam
	Params StartConnectionBetweenAliyunRouterInterfaceParamDetail `json:"startConnectionBetweenAliyunRouterInterface"`
}
// DeleteModelsParamDetail DeleteModels detail param
type DeleteModelsParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteModelsParam DeleteModels request param
type DeleteModelsParam struct {
	BaseParam
	Params DeleteModelsParamDetail `json:"deleteModels"`
}
// ListVmsFromSchedulingStateParamDetail ListVmsFromSchedulingState detail param
type ListVmsFromSchedulingStateParamDetail struct {
	RuleUuid string `json:"ruleUuid" validate:"required"`
	ExecuteStates []string `json:"executeStates" validate:"required"`
}

// ListVmsFromSchedulingStateParam ListVmsFromSchedulingState request param
type ListVmsFromSchedulingStateParam struct {
	BaseParam
	Params ListVmsFromSchedulingStateParamDetail `json:"params"`
}
// CreateRootVolumeTemplateFromVolumeSnapshotParamDetail CreateRootVolumeTemplateFromVolumeSnapshot detail param
type CreateRootVolumeTemplateFromVolumeSnapshotParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids" validate:"required"`
	Platform *string `json:"platform,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	System bool `json:"system,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateRootVolumeTemplateFromVolumeSnapshotParam CreateRootVolumeTemplateFromVolumeSnapshot request param
type CreateRootVolumeTemplateFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateRootVolumeTemplateFromVolumeSnapshotParamDetail `json:"params"`
}
// AllocateHostResourceParamDetail AllocateHostResource detail param
type AllocateHostResourceParamDetail struct {
	Strategy string `json:"strategy" validate:"required"`
	Scene string `json:"scene" validate:"required"`
	Vcpu int `json:"vcpu" validate:"required"`
	MemSize *int64 `json:"memSize,omitempty"`
}

// AllocateHostResourceParam AllocateHostResource request param
type AllocateHostResourceParam struct {
	BaseParam
	Params AllocateHostResourceParamDetail `json:"params"`
}
// GetCandidateLdapEntryForBindingParamDetail GetCandidateLdapEntryForBinding detail param
type GetCandidateLdapEntryForBindingParamDetail struct {
	LdapFilter string `json:"ldapFilter" validate:"required"`
	Limit *int `json:"limit,omitempty"`
}

// GetCandidateLdapEntryForBindingParam GetCandidateLdapEntryForBinding request param
type GetCandidateLdapEntryForBindingParam struct {
	BaseParam
	Params GetCandidateLdapEntryForBindingParamDetail `json:"getCandidateLdapEntryForBinding"`
}
// CheckElaborationContentParamDetail CheckElaborationContent detail param
type CheckElaborationContentParamDetail struct {
	ElaborateFile *string `json:"elaborateFile,omitempty"`
	ElaborateContent *string `json:"elaborateContent,omitempty"`
}

// CheckElaborationContentParam CheckElaborationContent request param
type CheckElaborationContentParam struct {
	BaseParam
	Params CheckElaborationContentParamDetail `json:"params"`
}
// DeleteVmConsolePasswordParamDetail DeleteVmConsolePassword detail param
type DeleteVmConsolePasswordParamDetail struct {
}

// DeleteVmConsolePasswordParam DeleteVmConsolePassword request param
type DeleteVmConsolePasswordParam struct {
	BaseParam
	Params DeleteVmConsolePasswordParamDetail `json:"deleteVmConsolePassword"`
}
// CreateVmBackupParamDetail CreateVmBackup detail param
type CreateVmBackupParamDetail struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Mode *string `json:"mode,omitempty"`
	VolumeReadBandwidth *int64 `json:"volumeReadBandwidth,omitempty"`
	VolumeWriteBandwidth *int64 `json:"volumeWriteBandwidth,omitempty"`
	NetworkReadBandwidth *int64 `json:"networkReadBandwidth,omitempty"`
	NetworkWriteBandwidth *int64 `json:"networkWriteBandwidth,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmBackupParam CreateVmBackup request param
type CreateVmBackupParam struct {
	BaseParam
	Params CreateVmBackupParamDetail `json:"params"`
}
// GetPrimaryStorageLicenseInfoParamDetail GetPrimaryStorageLicenseInfo detail param
type GetPrimaryStorageLicenseInfoParamDetail struct {
}

// GetPrimaryStorageLicenseInfoParam GetPrimaryStorageLicenseInfo request param
type GetPrimaryStorageLicenseInfoParam struct {
	BaseParam
	Params GetPrimaryStorageLicenseInfoParamDetail `json:"getPrimaryStorageLicenseInfo"`
}
// GetEncryptedFieldParamDetail GetEncryptedField detail param
type GetEncryptedFieldParamDetail struct {
	EncryptedType *string `json:"encryptedType,omitempty"`
}

// GetEncryptedFieldParam GetEncryptedField request param
type GetEncryptedFieldParam struct {
	BaseParam
	Params GetEncryptedFieldParamDetail `json:"getEncryptedField"`
}
// CleanInvalidLdapBindingParamDetail CleanInvalidLdapBinding detail param
type CleanInvalidLdapBindingParamDetail struct {
}

// CleanInvalidLdapBindingParam CleanInvalidLdapBinding request param
type CleanInvalidLdapBindingParam struct {
	BaseParam
	Params CleanInvalidLdapBindingParamDetail `json:"cleanInvalidLdapBinding"`
}
// AttachBaremetalPxeServerToClusterParamDetail AttachBaremetalPxeServerToCluster detail param
type AttachBaremetalPxeServerToClusterParamDetail struct {
}

// AttachBaremetalPxeServerToClusterParam AttachBaremetalPxeServerToCluster request param
type AttachBaremetalPxeServerToClusterParam struct {
	BaseParam
	Params AttachBaremetalPxeServerToClusterParamDetail `json:"params"`
}
// GetVmStartingCandidateClustersHostsParamDetail GetVmStartingCandidateClustersHosts detail param
type GetVmStartingCandidateClustersHostsParamDetail struct {
}

// GetVmStartingCandidateClustersHostsParam GetVmStartingCandidateClustersHosts request param
type GetVmStartingCandidateClustersHostsParam struct {
	BaseParam
	Params GetVmStartingCandidateClustersHostsParamDetail `json:"getVmStartingCandidateClustersHosts"`
}
// RecoverVmBackupFromImageStoreBackupStorageParamDetail RecoverVmBackupFromImageStoreBackupStorage detail param
type RecoverVmBackupFromImageStoreBackupStorageParamDetail struct {
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// RecoverVmBackupFromImageStoreBackupStorageParam RecoverVmBackupFromImageStoreBackupStorage request param
type RecoverVmBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params RecoverVmBackupFromImageStoreBackupStorageParamDetail `json:"recoverVmBackupFromImageStoreBackupStorage"`
}
// DetachIAM2ProjectFromIAM2OrganizationParamDetail DetachIAM2ProjectFromIAM2Organization detail param
type DetachIAM2ProjectFromIAM2OrganizationParamDetail struct {
}

// DetachIAM2ProjectFromIAM2OrganizationParam DetachIAM2ProjectFromIAM2Organization request param
type DetachIAM2ProjectFromIAM2OrganizationParam struct {
	BaseParam
	Params DetachIAM2ProjectFromIAM2OrganizationParamDetail `json:"detachIAM2ProjectFromIAM2Organization"`
}
// DiscoverExternalPrimaryStorageParamDetail DiscoverExternalPrimaryStorage detail param
type DiscoverExternalPrimaryStorageParamDetail struct {
	Url string `json:"url" validate:"required"`
	Identity *string `json:"identity,omitempty"`
	Config *string `json:"config,omitempty"`
}

// DiscoverExternalPrimaryStorageParam DiscoverExternalPrimaryStorage request param
type DiscoverExternalPrimaryStorageParam struct {
	BaseParam
	Params DiscoverExternalPrimaryStorageParamDetail `json:"params"`
}
// GetVolumeIoThreadPinParamDetail GetVolumeIoThreadPin detail param
type GetVolumeIoThreadPinParamDetail struct {
}

// GetVolumeIoThreadPinParam GetVolumeIoThreadPin request param
type GetVolumeIoThreadPinParam struct {
	BaseParam
	Params GetVolumeIoThreadPinParamDetail `json:"getVolumeIoThreadPin"`
}
// GetConnectionAccessPointFromRemoteParamDetail GetConnectionAccessPointFromRemote detail param
type GetConnectionAccessPointFromRemoteParamDetail struct {
}

// GetConnectionAccessPointFromRemoteParam GetConnectionAccessPointFromRemote request param
type GetConnectionAccessPointFromRemoteParam struct {
	BaseParam
	Params GetConnectionAccessPointFromRemoteParamDetail `json:"getConnectionAccessPointFromRemote"`
}
// GetVpcAttachedOspfParamDetail GetVpcAttachedOspf detail param
type GetVpcAttachedOspfParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVpcAttachedOspfParam GetVpcAttachedOspf request param
type GetVpcAttachedOspfParam struct {
	BaseParam
	Params GetVpcAttachedOspfParamDetail `json:"params"`
}
// PowerOffHostParamDetail PowerOffHost detail param
type PowerOffHostParamDetail struct {
	AdminPassword string `json:"adminPassword" validate:"required"`
	HostUuids []string `json:"hostUuids" validate:"required"`
	WaitTaskCompleted *bool `json:"waitTaskCompleted,omitempty"`
	MaxWaitTime *int64 `json:"maxWaitTime,omitempty"`
}

// PowerOffHostParam PowerOffHost request param
type PowerOffHostParam struct {
	BaseParam
	Params PowerOffHostParamDetail `json:"powerOffHost"`
}
// RemoveIAM2VirtualIDGroupFromProjectsParamDetail RemoveIAM2VirtualIDGroupFromProjects detail param
type RemoveIAM2VirtualIDGroupFromProjectsParamDetail struct {
	ProjectUuids []string `json:"projectUuids,omitempty"`
	GroupUuids []string `json:"groupUuids,omitempty"`
}

// RemoveIAM2VirtualIDGroupFromProjectsParam RemoveIAM2VirtualIDGroupFromProjects request param
type RemoveIAM2VirtualIDGroupFromProjectsParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDGroupFromProjectsParamDetail `json:"removeIAM2VirtualIDGroupFromProjects"`
}
// UpdateVmUserDefinedXmlHookScriptParamDetail UpdateVmUserDefinedXmlHookScript detail param
type UpdateVmUserDefinedXmlHookScriptParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	HookScript *string `json:"hookScript,omitempty"`
	StartupStrategy *string `json:"startupStrategy,omitempty"`
}

// UpdateVmUserDefinedXmlHookScriptParam UpdateVmUserDefinedXmlHookScript request param
type UpdateVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params UpdateVmUserDefinedXmlHookScriptParamDetail `json:"updateVmUserDefinedXmlHookScript"`
}
// GetIAM2ProjectContainerImageTagsParamDetail GetIAM2ProjectContainerImageTags detail param
type GetIAM2ProjectContainerImageTagsParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetIAM2ProjectContainerImageTagsParam GetIAM2ProjectContainerImageTags request param
type GetIAM2ProjectContainerImageTagsParam struct {
	BaseParam
	Params GetIAM2ProjectContainerImageTagsParamDetail `json:"getIAM2ProjectContainerImageTags"`
}
// DeleteAliyunDiskFromRemoteParamDetail DeleteAliyunDiskFromRemote detail param
type DeleteAliyunDiskFromRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAliyunDiskFromRemoteParam DeleteAliyunDiskFromRemote request param
type DeleteAliyunDiskFromRemoteParam struct {
	BaseParam
	Params DeleteAliyunDiskFromRemoteParamDetail `json:"deleteAliyunDiskFromRemote"`
}
// GetVersionParamDetail GetVersion detail param
type GetVersionParamDetail struct {
}

// GetVersionParam GetVersion request param
type GetVersionParam struct {
	BaseParam
	Params GetVersionParamDetail `json:"getVersion"`
}
// GetCandidateBackupStorageForCreatingImageParamDetail GetCandidateBackupStorageForCreatingImage detail param
type GetCandidateBackupStorageForCreatingImageParamDetail struct {
	VolumeUuid *string `json:"volumeUuid,omitempty"`
	VolumeSnapshotUuid *string `json:"volumeSnapshotUuid,omitempty"`
}

// GetCandidateBackupStorageForCreatingImageParam GetCandidateBackupStorageForCreatingImage request param
type GetCandidateBackupStorageForCreatingImageParam struct {
	BaseParam
	Params GetCandidateBackupStorageForCreatingImageParamDetail `json:"getCandidateBackupStorageForCreatingImage"`
}
// AttachAutoScalingTemplateToGroupParamDetail AttachAutoScalingTemplateToGroup detail param
type AttachAutoScalingTemplateToGroupParamDetail struct {
}

// AttachAutoScalingTemplateToGroupParam AttachAutoScalingTemplateToGroup request param
type AttachAutoScalingTemplateToGroupParam struct {
	BaseParam
	Params AttachAutoScalingTemplateToGroupParamDetail `json:"params"`
}
// GetCpuMemoryCapacityParamDetail GetCpuMemoryCapacity detail param
type GetCpuMemoryCapacityParamDetail struct {
	ZoneUuids []string `json:"zoneUuids,omitempty"`
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	HostUuids []string `json:"hostUuids,omitempty"`
	HypervisorType *string `json:"hypervisorType,omitempty"`
	All bool `json:"all,omitempty"`
}

// GetCpuMemoryCapacityParam GetCpuMemoryCapacity request param
type GetCpuMemoryCapacityParam struct {
	BaseParam
	Params GetCpuMemoryCapacityParamDetail `json:"getCpuMemoryCapacity"`
}
// AddIntegrityResourceParamDetail AddIntegrityResource detail param
type AddIntegrityResourceParamDetail struct {
	ResourceType string `json:"resourceType" validate:"required"`
	IntegrityResourceDataRangeInDays *int `json:"integrityResourceDataRangeInDays,omitempty"`
}

// AddIntegrityResourceParam AddIntegrityResource request param
type AddIntegrityResourceParam struct {
	BaseParam
	Params AddIntegrityResourceParamDetail `json:"params"`
}
// CheckVipPortAvailabilityParamDetail CheckVipPortAvailability detail param
type CheckVipPortAvailabilityParamDetail struct {
	Port int `json:"port" validate:"required"`
	ProtocolType string `json:"protocolType" validate:"required"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// CheckVipPortAvailabilityParam CheckVipPortAvailability request param
type CheckVipPortAvailabilityParam struct {
	BaseParam
	Params CheckVipPortAvailabilityParamDetail `json:"checkVipPortAvailability"`
}
// GetCandidateClustersForAttachingL2NetworkParamDetail GetCandidateClustersForAttachingL2Network detail param
type GetCandidateClustersForAttachingL2NetworkParamDetail struct {
	ClusterTypes []string `json:"clusterTypes,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetCandidateClustersForAttachingL2NetworkParam GetCandidateClustersForAttachingL2Network request param
type GetCandidateClustersForAttachingL2NetworkParam struct {
	BaseParam
	Params GetCandidateClustersForAttachingL2NetworkParamDetail `json:"getCandidateClustersForAttachingL2Network"`
}
// CheckScsiLunClusterStatusParamDetail CheckScsiLunClusterStatus detail param
type CheckScsiLunClusterStatusParamDetail struct {
}

// CheckScsiLunClusterStatusParam CheckScsiLunClusterStatus request param
type CheckScsiLunClusterStatusParam struct {
	BaseParam
	Params CheckScsiLunClusterStatusParamDetail `json:"checkScsiLunClusterStatus"`
}
// CheckBatchDataIntegrityParamDetail CheckBatchDataIntegrity detail param
type CheckBatchDataIntegrityParamDetail struct {
	ResourceUuids []string `json:"resourceUuids,omitempty"`
	ResourceType string `json:"resourceType" validate:"required"`
}

// CheckBatchDataIntegrityParam CheckBatchDataIntegrity request param
type CheckBatchDataIntegrityParam struct {
	BaseParam
	Params CheckBatchDataIntegrityParamDetail `json:"checkBatchDataIntegrity"`
}
// UpdateAutoScalingGroupRemovalInstanceRuleParamDetail UpdateAutoScalingGroupRemovalInstanceRule detail param
type UpdateAutoScalingGroupRemovalInstanceRuleParamDetail struct {
	AdjustmentType *string `json:"adjustmentType,omitempty"`
	AdjustmentValue *int `json:"adjustmentValue,omitempty"`
	RemovalPolicy *string `json:"removalPolicy,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Cooldown *int64 `json:"cooldown,omitempty"`
}

// UpdateAutoScalingGroupRemovalInstanceRuleParam UpdateAutoScalingGroupRemovalInstanceRule request param
type UpdateAutoScalingGroupRemovalInstanceRuleParam struct {
	BaseParam
	Params UpdateAutoScalingGroupRemovalInstanceRuleParamDetail `json:"updateAutoScalingGroupRemovalInstanceRule"`
}
// UploadFileToVmParamDetail UploadFileToVm detail param
type UploadFileToVmParamDetail struct {
	VmInstanceUuids []string `json:"vmInstanceUuids" validate:"required"`
	FileContent string `json:"fileContent" validate:"required"`
	RemotePath string `json:"remotePath" validate:"required"`
}

// UploadFileToVmParam UploadFileToVm request param
type UploadFileToVmParam struct {
	BaseParam
	Params UploadFileToVmParamDetail `json:"params"`
}
// ChangeL3NetworkDhcpIpAddressParamDetail ChangeL3NetworkDhcpIpAddress detail param
type ChangeL3NetworkDhcpIpAddressParamDetail struct {
	DhcpServerIp *string `json:"dhcpServerIp,omitempty"`
	Dhcpv6ServerIp *string `json:"dhcpv6ServerIp,omitempty"`
}

// ChangeL3NetworkDhcpIpAddressParam ChangeL3NetworkDhcpIpAddress request param
type ChangeL3NetworkDhcpIpAddressParam struct {
	BaseParam
	Params ChangeL3NetworkDhcpIpAddressParamDetail `json:"changeL3NetworkDhcpIpAddress"`
}
// CheckVolumeSnapshotGroupAvailabilityParamDetail CheckVolumeSnapshotGroupAvailability detail param
type CheckVolumeSnapshotGroupAvailabilityParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// CheckVolumeSnapshotGroupAvailabilityParam CheckVolumeSnapshotGroupAvailability request param
type CheckVolumeSnapshotGroupAvailabilityParam struct {
	BaseParam
	Params CheckVolumeSnapshotGroupAvailabilityParamDetail `json:"checkVolumeSnapshotGroupAvailability"`
}
// SsoClientPushDataParamDetail SsoClientPushData detail param
type SsoClientPushDataParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DataType *string `json:"dataType,omitempty"`
	ServerUrl *string `json:"serverUrl,omitempty"`
}

// SsoClientPushDataParam SsoClientPushData request param
type SsoClientPushDataParam struct {
	BaseParam
	Params SsoClientPushDataParamDetail `json:"ssoClientPushData"`
}
// AddEmailAddressToSNSEmailEndpointParamDetail AddEmailAddressToSNSEmailEndpoint detail param
type AddEmailAddressToSNSEmailEndpointParamDetail struct {
	EmailAddress string `json:"emailAddress" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddEmailAddressToSNSEmailEndpointParam AddEmailAddressToSNSEmailEndpoint request param
type AddEmailAddressToSNSEmailEndpointParam struct {
	BaseParam
	Params AddEmailAddressToSNSEmailEndpointParamDetail `json:"params"`
}
// BackupDatabaseToPublicCloudParamDetail BackupDatabaseToPublicCloud detail param
type BackupDatabaseToPublicCloudParamDetail struct {
	Type string `json:"type" validate:"required"`
	RegionId string `json:"regionId" validate:"required"`
	Local *bool `json:"local,omitempty"`
}

// BackupDatabaseToPublicCloudParam BackupDatabaseToPublicCloud request param
type BackupDatabaseToPublicCloudParam struct {
	BaseParam
	Params BackupDatabaseToPublicCloudParamDetail `json:"params"`
}
// RecoveryImageFromImageStoreBackupStorageParamDetail RecoveryImageFromImageStoreBackupStorage detail param
type RecoveryImageFromImageStoreBackupStorageParamDetail struct {
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
}

// RecoveryImageFromImageStoreBackupStorageParam RecoveryImageFromImageStoreBackupStorage request param
type RecoveryImageFromImageStoreBackupStorageParam struct {
	BaseParam
	Params RecoveryImageFromImageStoreBackupStorageParamDetail `json:"recoveryImageFromImageStoreBackupStorage"`
}
// RevertVmFromSnapshotGroupParamDetail RevertVmFromSnapshotGroup detail param
type RevertVmFromSnapshotGroupParamDetail struct {
}

// RevertVmFromSnapshotGroupParam RevertVmFromSnapshotGroup request param
type RevertVmFromSnapshotGroupParam struct {
	BaseParam
	Params RevertVmFromSnapshotGroupParamDetail `json:"revertVmFromSnapshotGroup"`
}
// DetachFirewallRuleSetFromL3ParamDetail DetachFirewallRuleSetFromL3 detail param
type DetachFirewallRuleSetFromL3ParamDetail struct {
	VpcFirewallUuid string `json:"vpcFirewallUuid" validate:"required"`
	Forward string `json:"forward" validate:"required"`
}

// DetachFirewallRuleSetFromL3Param DetachFirewallRuleSetFromL3 request param
type DetachFirewallRuleSetFromL3Param struct {
	BaseParam
	Params DetachFirewallRuleSetFromL3ParamDetail `json:"params"`
}
// ListVmSchedulingRulesFromExecuteStateParamDetail ListVmSchedulingRulesFromExecuteState detail param
type ListVmSchedulingRulesFromExecuteStateParamDetail struct {
	ExecuteStates []string `json:"executeStates" validate:"required"`
}

// ListVmSchedulingRulesFromExecuteStateParam ListVmSchedulingRulesFromExecuteState request param
type ListVmSchedulingRulesFromExecuteStateParam struct {
	BaseParam
	Params ListVmSchedulingRulesFromExecuteStateParamDetail `json:"params"`
}
// SetVmUserDefinedXmlParamDetail SetVmUserDefinedXml detail param
type SetVmUserDefinedXmlParamDetail struct {
	XmlBase64 string `json:"xmlBase64" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SetVmUserDefinedXmlParam SetVmUserDefinedXml request param
type SetVmUserDefinedXmlParam struct {
	BaseParam
	Params SetVmUserDefinedXmlParamDetail `json:"setVmUserDefinedXml"`
}
// SetImageQgaParamDetail SetImageQga detail param
type SetImageQgaParamDetail struct {
	Enable bool `json:"enable" validate:"required"`
}

// SetImageQgaParam SetImageQga request param
type SetImageQgaParam struct {
	BaseParam
	Params SetImageQgaParamDetail `json:"setImageQga"`
}
// ListVMsFromKVMHostParamDetail ListVMsFromKVMHost detail param
type ListVMsFromKVMHostParamDetail struct {
	LibvirtURI string `json:"libvirtURI" validate:"required"`
	ConversionHostUuid string `json:"conversionHostUuid" validate:"required"`
	SshPrivKey *string `json:"sshPrivKey,omitempty"`
	V2vType *string `json:"v2vType,omitempty"`
}

// ListVMsFromKVMHostParam ListVMsFromKVMHost request param
type ListVMsFromKVMHostParam struct {
	BaseParam
	Params ListVMsFromKVMHostParamDetail `json:"params"`
}
// TakeVmConsoleScreenshotParamDetail TakeVmConsoleScreenshot detail param
type TakeVmConsoleScreenshotParamDetail struct {
}

// TakeVmConsoleScreenshotParam TakeVmConsoleScreenshot request param
type TakeVmConsoleScreenshotParam struct {
	BaseParam
	Params TakeVmConsoleScreenshotParamDetail `json:"takeVmConsoleScreenshot"`
}
// RemoveVRouterNetworksFromOspfAreaParamDetail RemoveVRouterNetworksFromOspfArea detail param
type RemoveVRouterNetworksFromOspfAreaParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveVRouterNetworksFromOspfAreaParam RemoveVRouterNetworksFromOspfArea request param
type RemoveVRouterNetworksFromOspfAreaParam struct {
	BaseParam
	Params RemoveVRouterNetworksFromOspfAreaParamDetail `json:"removeVRouterNetworksFromOspfArea"`
}
// GetAliyunNasMountTargetRemoteParamDetail GetAliyunNasMountTargetRemote detail param
type GetAliyunNasMountTargetRemoteParamDetail struct {
	NasFSUuid string `json:"nasFSUuid" validate:"required"`
	MountDomain *string `json:"mountDomain,omitempty"`
}

// GetAliyunNasMountTargetRemoteParam GetAliyunNasMountTargetRemote request param
type GetAliyunNasMountTargetRemoteParam struct {
	BaseParam
	Params GetAliyunNasMountTargetRemoteParamDetail `json:"getAliyunNasMountTargetRemote"`
}
// CreateImageGroupFromVmInstanceParamDetail CreateImageGroupFromVmInstance detail param
type CreateImageGroupFromVmInstanceParamDetail struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateImageGroupFromVmInstanceParam CreateImageGroupFromVmInstance request param
type CreateImageGroupFromVmInstanceParam struct {
	BaseParam
	Params CreateImageGroupFromVmInstanceParamDetail `json:"params"`
}
// TerminateVirtualBorderRouterRemoteParamDetail TerminateVirtualBorderRouterRemote detail param
type TerminateVirtualBorderRouterRemoteParamDetail struct {
}

// TerminateVirtualBorderRouterRemoteParam TerminateVirtualBorderRouterRemote request param
type TerminateVirtualBorderRouterRemoteParam struct {
	BaseParam
	Params TerminateVirtualBorderRouterRemoteParamDetail `json:"terminateVirtualBorderRouterRemote"`
}
// DeleteVmBackupParamDetail DeleteVmBackup detail param
type DeleteVmBackupParamDetail struct {
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	HandleDependency *bool `json:"handleDependency,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVmBackupParam DeleteVmBackup request param
type DeleteVmBackupParam struct {
	BaseParam
	Params DeleteVmBackupParamDetail `json:"deleteVmBackup"`
}
// SetVmSecurityLevelParamDetail SetVmSecurityLevel detail param
type SetVmSecurityLevelParamDetail struct {
	SecurityLevel *string `json:"securityLevel,omitempty"`
}

// SetVmSecurityLevelParam SetVmSecurityLevel request param
type SetVmSecurityLevelParam struct {
	BaseParam
	Params SetVmSecurityLevelParamDetail `json:"setVmSecurityLevel"`
}
// RemoveMdevDeviceSpecFromVmInstanceParamDetail RemoveMdevDeviceSpecFromVmInstance detail param
type RemoveMdevDeviceSpecFromVmInstanceParamDetail struct {
}

// RemoveMdevDeviceSpecFromVmInstanceParam RemoveMdevDeviceSpecFromVmInstance request param
type RemoveMdevDeviceSpecFromVmInstanceParam struct {
	BaseParam
	Params RemoveMdevDeviceSpecFromVmInstanceParamDetail `json:"removeMdevDeviceSpecFromVmInstance"`
}
// SyncVolumeSizeParamDetail SyncVolumeSize detail param
type SyncVolumeSizeParamDetail struct {
}

// SyncVolumeSizeParam SyncVolumeSize request param
type SyncVolumeSizeParam struct {
	BaseParam
	Params SyncVolumeSizeParamDetail `json:"syncVolumeSize"`
}
// GetTrashOnBackupStorageParamDetail GetTrashOnBackupStorage detail param
type GetTrashOnBackupStorageParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	TrashType *string `json:"trashType,omitempty"`
}

// GetTrashOnBackupStorageParam GetTrashOnBackupStorage request param
type GetTrashOnBackupStorageParam struct {
	BaseParam
	Params GetTrashOnBackupStorageParamDetail `json:"getTrashOnBackupStorage"`
}
// ChangeDiskOfferingStateParamDetail ChangeDiskOfferingState detail param
type ChangeDiskOfferingStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeDiskOfferingStateParam ChangeDiskOfferingState request param
type ChangeDiskOfferingStateParam struct {
	BaseParam
	Params ChangeDiskOfferingStateParamDetail `json:"changeDiskOfferingState"`
}
// RequestConsoleAccessParamDetail RequestConsoleAccess detail param
type RequestConsoleAccessParamDetail struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// RequestConsoleAccessParam RequestConsoleAccess request param
type RequestConsoleAccessParam struct {
	BaseParam
	Params RequestConsoleAccessParamDetail `json:"params"`
}
// ChangeIAM2VirtualIDGroupStateParamDetail ChangeIAM2VirtualIDGroupState detail param
type ChangeIAM2VirtualIDGroupStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeIAM2VirtualIDGroupStateParam ChangeIAM2VirtualIDGroupState request param
type ChangeIAM2VirtualIDGroupStateParam struct {
	BaseParam
	Params ChangeIAM2VirtualIDGroupStateParamDetail `json:"changeIAM2VirtualIDGroupState"`
}
// UpdateEventDataParamDetail UpdateEventData detail param
type UpdateEventDataParamDetail struct {
	DataUuid *string `json:"dataUuid,omitempty"`
	DataStartTime *int64 `json:"dataStartTime,omitempty"`
	DataEndTime *int64 `json:"dataEndTime,omitempty"`
	UpdateMode string `json:"updateMode" validate:"required"`
	ReadStatus *string `json:"readStatus,omitempty"`
}

// UpdateEventDataParam UpdateEventData request param
type UpdateEventDataParam struct {
	BaseParam
	Params UpdateEventDataParamDetail `json:"updateEventData"`
}
// SyncHybridEipFromRemoteParamDetail SyncHybridEipFromRemote detail param
type SyncHybridEipFromRemoteParamDetail struct {
	Type string `json:"type" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncHybridEipFromRemoteParam SyncHybridEipFromRemote request param
type SyncHybridEipFromRemoteParam struct {
	BaseParam
	Params SyncHybridEipFromRemoteParamDetail `json:"syncHybridEipFromRemote"`
}
// DeleteAliyunRouteEntryRemoteParamDetail DeleteAliyunRouteEntryRemote detail param
type DeleteAliyunRouteEntryRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAliyunRouteEntryRemoteParam DeleteAliyunRouteEntryRemote request param
type DeleteAliyunRouteEntryRemoteParam struct {
	BaseParam
	Params DeleteAliyunRouteEntryRemoteParamDetail `json:"deleteAliyunRouteEntryRemote"`
}
// UngenerateSriovPciDevicesParamDetail UngenerateSriovPciDevices detail param
type UngenerateSriovPciDevicesParamDetail struct {
}

// UngenerateSriovPciDevicesParam UngenerateSriovPciDevices request param
type UngenerateSriovPciDevicesParam struct {
	BaseParam
	Params UngenerateSriovPciDevicesParamDetail `json:"ungenerateSriovPciDevices"`
}
// DeleteVmStaticIpParamDetail DeleteVmStaticIp detail param
type DeleteVmStaticIpParamDetail struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	StaticIp *string `json:"staticIp,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVmStaticIpParam DeleteVmStaticIp request param
type DeleteVmStaticIpParam struct {
	BaseParam
	Params DeleteVmStaticIpParamDetail `json:"deleteVmStaticIp"`
}
// AttachMonitorTriggerActionToTriggerParamDetail AttachMonitorTriggerActionToTrigger detail param
type AttachMonitorTriggerActionToTriggerParamDetail struct {
}

// AttachMonitorTriggerActionToTriggerParam AttachMonitorTriggerActionToTrigger request param
type AttachMonitorTriggerActionToTriggerParam struct {
	BaseParam
	Params AttachMonitorTriggerActionToTriggerParamDetail `json:"params"`
}
// GetAliyunNasFileSystemRemoteParamDetail GetAliyunNasFileSystemRemote detail param
type GetAliyunNasFileSystemRemoteParamDetail struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	FileSystemId *string `json:"fileSystemId,omitempty"`
}

// GetAliyunNasFileSystemRemoteParam GetAliyunNasFileSystemRemote request param
type GetAliyunNasFileSystemRemoteParam struct {
	BaseParam
	Params GetAliyunNasFileSystemRemoteParamDetail `json:"getAliyunNasFileSystemRemote"`
}
// UpdateOrganizationQuotaParamDetail UpdateOrganizationQuota detail param
type UpdateOrganizationQuotaParamDetail struct {
	IdentityUuid string `json:"identityUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Value int64 `json:"value" validate:"required"`
}

// UpdateOrganizationQuotaParam UpdateOrganizationQuota request param
type UpdateOrganizationQuotaParam struct {
	BaseParam
	Params UpdateOrganizationQuotaParamDetail `json:"updateOrganizationQuota"`
}
// ChangePreconfigurationTemplateStateParamDetail ChangePreconfigurationTemplateState detail param
type ChangePreconfigurationTemplateStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangePreconfigurationTemplateStateParam ChangePreconfigurationTemplateState request param
type ChangePreconfigurationTemplateStateParam struct {
	BaseParam
	Params ChangePreconfigurationTemplateStateParamDetail `json:"changePreconfigurationTemplateState"`
}
// SetOrganizationSupervisorParamDetail SetOrganizationSupervisor detail param
type SetOrganizationSupervisorParamDetail struct {
	VirtualIDUuid string `json:"virtualIDUuid" validate:"required"`
}

// SetOrganizationSupervisorParam SetOrganizationSupervisor request param
type SetOrganizationSupervisorParam struct {
	BaseParam
	Params SetOrganizationSupervisorParamDetail `json:"setOrganizationSupervisor"`
}
// AttachL3NetworksToIPsecConnectionParamDetail AttachL3NetworksToIPsecConnection detail param
type AttachL3NetworksToIPsecConnectionParamDetail struct {
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AttachL3NetworksToIPsecConnectionParam AttachL3NetworksToIPsecConnection request param
type AttachL3NetworksToIPsecConnectionParam struct {
	BaseParam
	Params AttachL3NetworksToIPsecConnectionParamDetail `json:"params"`
}
// ExecuteGuestVmScriptParamDetail ExecuteGuestVmScript detail param
type ExecuteGuestVmScriptParamDetail struct {
	VmInstanceUuids []string `json:"vmInstanceUuids" validate:"required"`
	ScriptTimeout *int `json:"scriptTimeout,omitempty"`
	LogPath *string `json:"logPath,omitempty"`
	RecordUuid *string `json:"recordUuid,omitempty"`
	RuntimeParams *string `json:"runtimeParams,omitempty"`
}

// ExecuteGuestVmScriptParam ExecuteGuestVmScript request param
type ExecuteGuestVmScriptParam struct {
	BaseParam
	Params ExecuteGuestVmScriptParamDetail `json:"executeGuestVmScript"`
}
// AddNfsPrimaryStorageParamDetail AddNfsPrimaryStorage detail param
type AddNfsPrimaryStorageParamDetail struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddNfsPrimaryStorageParam AddNfsPrimaryStorage request param
type AddNfsPrimaryStorageParam struct {
	BaseParam
	Params AddNfsPrimaryStorageParamDetail `json:"params"`
}
// GetIAM2ProjectContainerClusterCandidatesParamDetail GetIAM2ProjectContainerClusterCandidates detail param
type GetIAM2ProjectContainerClusterCandidatesParamDetail struct {
}

// GetIAM2ProjectContainerClusterCandidatesParam GetIAM2ProjectContainerClusterCandidates request param
type GetIAM2ProjectContainerClusterCandidatesParam struct {
	BaseParam
	Params GetIAM2ProjectContainerClusterCandidatesParamDetail `json:"getIAM2ProjectContainerClusterCandidates"`
}
// AttachTagToResourcesParamDetail AttachTagToResources detail param
type AttachTagToResourcesParamDetail struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	Tokens map[string]string `json:"tokens,omitempty"`
}

// AttachTagToResourcesParam AttachTagToResources request param
type AttachTagToResourcesParam struct {
	BaseParam
	Params AttachTagToResourcesParamDetail `json:"params"`
}
// ChangePrimaryStorageStateParamDetail ChangePrimaryStorageState detail param
type ChangePrimaryStorageStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangePrimaryStorageStateParam ChangePrimaryStorageState request param
type ChangePrimaryStorageStateParam struct {
	BaseParam
	Params ChangePrimaryStorageStateParamDetail `json:"changePrimaryStorageState"`
}
// GetVpcAttachedNetflowParamDetail GetVpcAttachedNetflow detail param
type GetVpcAttachedNetflowParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVpcAttachedNetflowParam GetVpcAttachedNetflow request param
type GetVpcAttachedNetflowParam struct {
	BaseParam
	Params GetVpcAttachedNetflowParamDetail `json:"params"`
}
// GetAuditDataParamDetail GetAuditData detail param
type GetAuditDataParamDetail struct {
	StartTime *int64 `json:"startTime,omitempty"`
	EndTime *int64 `json:"endTime,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Conditions []string `json:"conditions,omitempty"`
	AuditType *string `json:"auditType,omitempty"`
}

// GetAuditDataParam GetAuditData request param
type GetAuditDataParam struct {
	BaseParam
	Params GetAuditDataParamDetail `json:"getAuditData"`
}
// GetSpiceCertificatesParamDetail GetSpiceCertificates detail param
type GetSpiceCertificatesParamDetail struct {
}

// GetSpiceCertificatesParam GetSpiceCertificates request param
type GetSpiceCertificatesParam struct {
	BaseParam
	Params GetSpiceCertificatesParamDetail `json:"getSpiceCertificates"`
}
// RemoveUserFromGroupParamDetail RemoveUserFromGroup detail param
type RemoveUserFromGroupParamDetail struct {
}

// RemoveUserFromGroupParam RemoveUserFromGroup request param
type RemoveUserFromGroupParam struct {
	BaseParam
	Params RemoveUserFromGroupParamDetail `json:"removeUserFromGroup"`
}
// DeleteEcsVpcRemoteParamDetail DeleteEcsVpcRemote detail param
type DeleteEcsVpcRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsVpcRemoteParam DeleteEcsVpcRemote request param
type DeleteEcsVpcRemoteParam struct {
	BaseParam
	Params DeleteEcsVpcRemoteParamDetail `json:"deleteEcsVpcRemote"`
}
// SyncDatabaseBackupFromImageStoreBackupStorageParamDetail SyncDatabaseBackupFromImageStoreBackupStorage detail param
type SyncDatabaseBackupFromImageStoreBackupStorageParamDetail struct {
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// SyncDatabaseBackupFromImageStoreBackupStorageParam SyncDatabaseBackupFromImageStoreBackupStorage request param
type SyncDatabaseBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncDatabaseBackupFromImageStoreBackupStorageParamDetail `json:"syncDatabaseBackupFromImageStoreBackupStorage"`
}
// DeleteFirewallIpSetTemplateParamDetail DeleteFirewallIpSetTemplate detail param
type DeleteFirewallIpSetTemplateParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteFirewallIpSetTemplateParam DeleteFirewallIpSetTemplate request param
type DeleteFirewallIpSetTemplateParam struct {
	BaseParam
	Params DeleteFirewallIpSetTemplateParamDetail `json:"deleteFirewallIpSetTemplate"`
}
// SNSDingTalkTestConnectionParamDetail SNSDingTalkTestConnection detail param
type SNSDingTalkTestConnectionParamDetail struct {
	Url *string `json:"url,omitempty"`
	AtAll *bool `json:"atAll,omitempty"`
	AtPersonPhoneNumbers []string `json:"atPersonPhoneNumbers,omitempty"`
	Secret *string `json:"secret,omitempty"`
	TestMsg string `json:"testMsg" validate:"required"`
	EndpointUuid *string `json:"endpointUuid,omitempty"`
}

// SNSDingTalkTestConnectionParam SNSDingTalkTestConnection request param
type SNSDingTalkTestConnectionParam struct {
	BaseParam
	Params SNSDingTalkTestConnectionParamDetail `json:"params"`
}
// ExportImageFromBackupStorageParamDetail ExportImageFromBackupStorage detail param
type ExportImageFromBackupStorageParamDetail struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	ExportFormat *string `json:"exportFormat,omitempty"`
}

// ExportImageFromBackupStorageParam ExportImageFromBackupStorage request param
type ExportImageFromBackupStorageParam struct {
	BaseParam
	Params ExportImageFromBackupStorageParamDetail `json:"exportImageFromBackupStorage"`
}
// GetModelCenterServicesParamDetail GetModelCenterServices detail param
type GetModelCenterServicesParamDetail struct {
	ModelCenterUuids []string `json:"modelCenterUuids,omitempty"`
}

// GetModelCenterServicesParam GetModelCenterServices request param
type GetModelCenterServicesParam struct {
	BaseParam
	Params GetModelCenterServicesParamDetail `json:"getModelCenterServices"`
}
// CreateFirewallIpSetTemplateParamDetail CreateFirewallIpSetTemplate detail param
type CreateFirewallIpSetTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	SourceValue *string `json:"sourceValue,omitempty"`
	DestValue *string `json:"destValue,omitempty"`
	Type string `json:"type" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFirewallIpSetTemplateParam CreateFirewallIpSetTemplate request param
type CreateFirewallIpSetTemplateParam struct {
	BaseParam
	Params CreateFirewallIpSetTemplateParamDetail `json:"params"`
}
// DetachMonitorTriggerActionFromTriggerParamDetail DetachMonitorTriggerActionFromTrigger detail param
type DetachMonitorTriggerActionFromTriggerParamDetail struct {
}

// DetachMonitorTriggerActionFromTriggerParam DetachMonitorTriggerActionFromTrigger request param
type DetachMonitorTriggerActionFromTriggerParam struct {
	BaseParam
	Params DetachMonitorTriggerActionFromTriggerParamDetail `json:"detachMonitorTriggerActionFromTrigger"`
}
// DetachPolicyRouteRuleSetFromL3ParamDetail DetachPolicyRouteRuleSetFromL3 detail param
type DetachPolicyRouteRuleSetFromL3ParamDetail struct {
}

// DetachPolicyRouteRuleSetFromL3Param DetachPolicyRouteRuleSetFromL3 request param
type DetachPolicyRouteRuleSetFromL3Param struct {
	BaseParam
	Params DetachPolicyRouteRuleSetFromL3ParamDetail `json:"detachPolicyRouteRuleSetFromL3"`
}
// CreateL2TfNetworkParamDetail CreateL2TfNetwork detail param
type CreateL2TfNetworkParamDetail struct {
	IpPrefix *string `json:"ipPrefix,omitempty"`
	IpPrefixLength *int `json:"ipPrefixLength,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	PhysicalInterface *string `json:"physicalInterface,omitempty"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Pvlan *string `json:"pvlan,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2TfNetworkParam CreateL2TfNetwork request param
type CreateL2TfNetworkParam struct {
	BaseParam
	Params CreateL2TfNetworkParamDetail `json:"params"`
}
// GetInterdependentL3NetworksImagesParamDetail GetInterdependentL3NetworksImages detail param
type GetInterdependentL3NetworksImagesParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	ImageUuid *string `json:"imageUuid,omitempty"`
	RaiseException *bool `json:"raiseException,omitempty"`
}

// GetInterdependentL3NetworksImagesParam GetInterdependentL3NetworksImages request param
type GetInterdependentL3NetworksImagesParam struct {
	BaseParam
	Params GetInterdependentL3NetworksImagesParamDetail `json:"getInterdependentL3NetworksImages"`
}
// ValidateVolumeSnapshotChainParamDetail ValidateVolumeSnapshotChain detail param
type ValidateVolumeSnapshotChainParamDetail struct {
}

// ValidateVolumeSnapshotChainParam ValidateVolumeSnapshotChain request param
type ValidateVolumeSnapshotChainParam struct {
	BaseParam
	Params ValidateVolumeSnapshotChainParamDetail `json:"validateVolumeSnapshotChain"`
}
// ChangeHostNetworkInterfaceLldpModeParamDetail ChangeHostNetworkInterfaceLldpMode detail param
type ChangeHostNetworkInterfaceLldpModeParamDetail struct {
	InterfaceUuids []string `json:"interfaceUuids" validate:"required"`
	Mode *string `json:"mode,omitempty"`
}

// ChangeHostNetworkInterfaceLldpModeParam ChangeHostNetworkInterfaceLldpMode request param
type ChangeHostNetworkInterfaceLldpModeParam struct {
	BaseParam
	Params ChangeHostNetworkInterfaceLldpModeParamDetail `json:"changeHostNetworkInterfaceLldpMode"`
}
// GetGuestOsMetadataParamDetail GetGuestOsMetadata detail param
type GetGuestOsMetadataParamDetail struct {
}

// GetGuestOsMetadataParam GetGuestOsMetadata request param
type GetGuestOsMetadataParam struct {
	BaseParam
	Params GetGuestOsMetadataParamDetail `json:"getGuestOsMetadata"`
}
// GetCandidateVmNicsForLoadBalancerServerGroupParamDetail GetCandidateVmNicsForLoadBalancerServerGroup detail param
type GetCandidateVmNicsForLoadBalancerServerGroupParamDetail struct {
	ServergroupUuid *string `json:"servergroupUuid,omitempty"`
	LoadBalancerUuid *string `json:"loadBalancerUuid,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
}

// GetCandidateVmNicsForLoadBalancerServerGroupParam GetCandidateVmNicsForLoadBalancerServerGroup request param
type GetCandidateVmNicsForLoadBalancerServerGroupParam struct {
	BaseParam
	Params GetCandidateVmNicsForLoadBalancerServerGroupParamDetail `json:"getCandidateVmNicsForLoadBalancerServerGroup"`
}
// AttachIscsiServerToClusterParamDetail AttachIscsiServerToCluster detail param
type AttachIscsiServerToClusterParamDetail struct {
}

// AttachIscsiServerToClusterParam AttachIscsiServerToCluster request param
type AttachIscsiServerToClusterParam struct {
	BaseParam
	Params AttachIscsiServerToClusterParamDetail `json:"params"`
}
// AttachRoleToAccountParamDetail AttachRoleToAccount detail param
type AttachRoleToAccountParamDetail struct {
}

// AttachRoleToAccountParam AttachRoleToAccount request param
type AttachRoleToAccountParam struct {
	BaseParam
	Params AttachRoleToAccountParamDetail `json:"params"`
}
// AttachIsoToVmInstanceParamDetail AttachIsoToVmInstance detail param
type AttachIsoToVmInstanceParamDetail struct {
}

// AttachIsoToVmInstanceParam AttachIsoToVmInstance request param
type AttachIsoToVmInstanceParam struct {
	BaseParam
	Params AttachIsoToVmInstanceParamDetail `json:"attachIsoToVmInstance"`
}
// SetVRouterRouterIdParamDetail SetVRouterRouterId detail param
type SetVRouterRouterIdParamDetail struct {
	RouterId string `json:"routerId" validate:"required"`
}

// SetVRouterRouterIdParam SetVRouterRouterId request param
type SetVRouterRouterIdParam struct {
	BaseParam
	Params SetVRouterRouterIdParamDetail `json:"params"`
}
// ExpungeVmUserDefinedXmlHookScriptParamDetail ExpungeVmUserDefinedXmlHookScript detail param
type ExpungeVmUserDefinedXmlHookScriptParamDetail struct {
}

// ExpungeVmUserDefinedXmlHookScriptParam ExpungeVmUserDefinedXmlHookScript request param
type ExpungeVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params ExpungeVmUserDefinedXmlHookScriptParamDetail `json:"expungeVmUserDefinedXmlHookScript"`
}
// DeleteCdpTaskDataParamDetail DeleteCdpTaskData detail param
type DeleteCdpTaskDataParamDetail struct {
}

// DeleteCdpTaskDataParam DeleteCdpTaskData request param
type DeleteCdpTaskDataParam struct {
	BaseParam
	Params DeleteCdpTaskDataParamDetail `json:"params"`
}
// CheckApiPermissionParamDetail CheckApiPermission detail param
type CheckApiPermissionParamDetail struct {
	UserUuid *string `json:"userUuid,omitempty"`
	ApiNames []string `json:"apiNames" validate:"required"`
}

// CheckApiPermissionParam CheckApiPermission request param
type CheckApiPermissionParam struct {
	BaseParam
	Params CheckApiPermissionParamDetail `json:"checkApiPermission"`
}
// GetTextTemplateArgParamDetail GetTextTemplateArg detail param
type GetTextTemplateArgParamDetail struct {
}

// GetTextTemplateArgParam GetTextTemplateArg request param
type GetTextTemplateArgParam struct {
	BaseParam
	Params GetTextTemplateArgParamDetail `json:"getTextTemplateArg"`
}
// DeleteFirewallParamDetail DeleteFirewall detail param
type DeleteFirewallParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteFirewallParam DeleteFirewall request param
type DeleteFirewallParam struct {
	BaseParam
	Params DeleteFirewallParamDetail `json:"deleteFirewall"`
}
// GetVmCapabilitiesParamDetail GetVmCapabilities detail param
type GetVmCapabilitiesParamDetail struct {
}

// GetVmCapabilitiesParam GetVmCapabilities request param
type GetVmCapabilitiesParam struct {
	BaseParam
	Params GetVmCapabilitiesParamDetail `json:"getVmCapabilities"`
}
// ChangeAccessKeyStateParamDetail ChangeAccessKeyState detail param
type ChangeAccessKeyStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeAccessKeyStateParam ChangeAccessKeyState request param
type ChangeAccessKeyStateParam struct {
	BaseParam
	Params ChangeAccessKeyStateParamDetail `json:"changeAccessKeyState"`
}
// DeployDistributedModelServiceParamDetail DeployDistributedModelService detail param
type DeployDistributedModelServiceParamDetail struct {
	ModelServices []ModelServiceParam `json:"modelServices" validate:"required"`
	ServiceCreationStrategy string `json:"serviceCreationStrategy" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// DeployDistributedModelServiceParam DeployDistributedModelService request param
type DeployDistributedModelServiceParam struct {
	BaseParam
	Params DeployDistributedModelServiceParamDetail `json:"deployDistributedModelService"`
}
// GetIAM2SystemAttributesParamDetail GetIAM2SystemAttributes detail param
type GetIAM2SystemAttributesParamDetail struct {
}

// GetIAM2SystemAttributesParam GetIAM2SystemAttributes request param
type GetIAM2SystemAttributesParam struct {
	BaseParam
	Params GetIAM2SystemAttributesParamDetail `json:"getIAM2SystemAttributes"`
}
// ChangeInstanceOfferingStateParamDetail ChangeInstanceOfferingState detail param
type ChangeInstanceOfferingStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeInstanceOfferingStateParam ChangeInstanceOfferingState request param
type ChangeInstanceOfferingStateParam struct {
	BaseParam
	Params ChangeInstanceOfferingStateParamDetail `json:"changeInstanceOfferingState"`
}
// GetBackupStorageCapacityParamDetail GetBackupStorageCapacity detail param
type GetBackupStorageCapacityParamDetail struct {
	ZoneUuids []string `json:"zoneUuids,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	All bool `json:"all,omitempty"`
}

// GetBackupStorageCapacityParam GetBackupStorageCapacity request param
type GetBackupStorageCapacityParam struct {
	BaseParam
	Params GetBackupStorageCapacityParamDetail `json:"getBackupStorageCapacity"`
}
// GenerateSeMdevDevicesParamDetail GenerateSeMdevDevices detail param
type GenerateSeMdevDevicesParamDetail struct {
	VirtPartNum int `json:"virtPartNum" validate:"required"`
}

// GenerateSeMdevDevicesParam GenerateSeMdevDevices request param
type GenerateSeMdevDevicesParam struct {
	BaseParam
	Params GenerateSeMdevDevicesParamDetail `json:"generateSeMdevDevices"`
}
// CreateMiniClusterParamDetail CreateMiniCluster detail param
type CreateMiniClusterParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	HostManagementIps []string `json:"hostManagementIps" validate:"required"`
	Username *string `json:"username,omitempty"`
	Password string `json:"password" validate:"required"`
	SshPort *int `json:"sshPort,omitempty"`
	Description *string `json:"description,omitempty"`
	HypervisorType string `json:"hypervisorType" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMiniClusterParam CreateMiniCluster request param
type CreateMiniClusterParam struct {
	BaseParam
	Params CreateMiniClusterParamDetail `json:"params"`
}
// SyncImageFromImageStoreBackupStorageParamDetail SyncImageFromImageStoreBackupStorage detail param
type SyncImageFromImageStoreBackupStorageParamDetail struct {
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
}

// SyncImageFromImageStoreBackupStorageParam SyncImageFromImageStoreBackupStorage request param
type SyncImageFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncImageFromImageStoreBackupStorageParamDetail `json:"syncImageFromImageStoreBackupStorage"`
}
// ChangeVipStateParamDetail ChangeVipState detail param
type ChangeVipStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeVipStateParam ChangeVipState request param
type ChangeVipStateParam struct {
	BaseParam
	Params ChangeVipStateParamDetail `json:"changeVipState"`
}
// UndoSnapshotCreationParamDetail UndoSnapshotCreation detail param
type UndoSnapshotCreationParamDetail struct {
	SnapShotUuid string `json:"snapShotUuid" validate:"required"`
}

// UndoSnapshotCreationParam UndoSnapshotCreation request param
type UndoSnapshotCreationParam struct {
	BaseParam
	Params UndoSnapshotCreationParamDetail `json:"undoSnapshotCreation"`
}
// AddBuildAppParamDetail AddBuildApp detail param
type AddBuildAppParamDetail struct {
	Url string `json:"url" validate:"required"`
	Type *string `json:"type,omitempty"`
	BackupStorageUuid *string `json:"backupStorageUuid,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBuildAppParam AddBuildApp request param
type AddBuildAppParam struct {
	BaseParam
	Params AddBuildAppParamDetail `json:"params"`
}
// CreateVmFromVolumeBackupParamDetail CreateVmFromVolumeBackup detail param
type CreateVmFromVolumeBackupParamDetail struct {
	Name string `json:"name" validate:"required"`
	BackupStorageUuid *string `json:"backupStorageUuid,omitempty"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid" validate:"required"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	Type *string `json:"type,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	Description *string `json:"description,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmFromVolumeBackupParam CreateVmFromVolumeBackup request param
type CreateVmFromVolumeBackupParam struct {
	BaseParam
	Params CreateVmFromVolumeBackupParamDetail `json:"params"`
}
// GetIdentityZoneFromRemoteParamDetail GetIdentityZoneFromRemote detail param
type GetIdentityZoneFromRemoteParamDetail struct {
	Type *string `json:"type,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	RegionId *string `json:"regionId,omitempty"`
}

// GetIdentityZoneFromRemoteParam GetIdentityZoneFromRemote request param
type GetIdentityZoneFromRemoteParam struct {
	BaseParam
	Params GetIdentityZoneFromRemoteParamDetail `json:"getIdentityZoneFromRemote"`
}
// GetEcsInstanceVncUrlParamDetail GetEcsInstanceVncUrl detail param
type GetEcsInstanceVncUrlParamDetail struct {
}

// GetEcsInstanceVncUrlParam GetEcsInstanceVncUrl request param
type GetEcsInstanceVncUrlParam struct {
	BaseParam
	Params GetEcsInstanceVncUrlParamDetail `json:"getEcsInstanceVncUrl"`
}
// AddMonToCephPrimaryStorageParamDetail AddMonToCephPrimaryStorage detail param
type AddMonToCephPrimaryStorageParamDetail struct {
	MonUrls []string `json:"monUrls" validate:"required"`
}

// AddMonToCephPrimaryStorageParam AddMonToCephPrimaryStorage request param
type AddMonToCephPrimaryStorageParam struct {
	BaseParam
	Params AddMonToCephPrimaryStorageParamDetail `json:"params"`
}
// RemoveHostRouteFromL3NetworkParamDetail RemoveHostRouteFromL3Network detail param
type RemoveHostRouteFromL3NetworkParamDetail struct {
	Prefix string `json:"prefix" validate:"required"`
}

// RemoveHostRouteFromL3NetworkParam RemoveHostRouteFromL3Network request param
type RemoveHostRouteFromL3NetworkParam struct {
	BaseParam
	Params RemoveHostRouteFromL3NetworkParamDetail `json:"removeHostRouteFromL3Network"`
}
// BackupStorageMigrateImageParamDetail BackupStorageMigrateImage detail param
type BackupStorageMigrateImageParamDetail struct {
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// BackupStorageMigrateImageParam BackupStorageMigrateImage request param
type BackupStorageMigrateImageParam struct {
	BaseParam
	Params BackupStorageMigrateImageParamDetail `json:"backupStorageMigrateImage"`
}
// ChangeIAM2VirtualIDTypeParamDetail ChangeIAM2VirtualIDType detail param
type ChangeIAM2VirtualIDTypeParamDetail struct {
	Type string `json:"type" validate:"required"`
}

// ChangeIAM2VirtualIDTypeParam ChangeIAM2VirtualIDType request param
type ChangeIAM2VirtualIDTypeParam struct {
	BaseParam
	Params ChangeIAM2VirtualIDTypeParamDetail `json:"changeIAM2VirtualIDType"`
}
// RemoveBackendServerFromServerGroupParamDetail RemoveBackendServerFromServerGroup detail param
type RemoveBackendServerFromServerGroupParamDetail struct {
	VmNicUuids []string `json:"vmNicUuids,omitempty"`
	ServerIps []string `json:"serverIps,omitempty"`
}

// RemoveBackendServerFromServerGroupParam RemoveBackendServerFromServerGroup request param
type RemoveBackendServerFromServerGroupParam struct {
	BaseParam
	Params RemoveBackendServerFromServerGroupParamDetail `json:"removeBackendServerFromServerGroup"`
}
// GetVpcAttachedVipParamDetail GetVpcAttachedVip detail param
type GetVpcAttachedVipParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVpcAttachedVipParam GetVpcAttachedVip request param
type GetVpcAttachedVipParam struct {
	BaseParam
	Params GetVpcAttachedVipParamDetail `json:"params"`
}
// AddIpv6RangeParamDetail AddIpv6Range detail param
type AddIpv6RangeParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	StartIp string `json:"startIp" validate:"required"`
	EndIp string `json:"endIp" validate:"required"`
	Gateway string `json:"gateway" validate:"required"`
	PrefixLen int `json:"prefixLen" validate:"required"`
	AddressMode string `json:"addressMode" validate:"required"`
	IpRangeType *string `json:"ipRangeType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIpv6RangeParam AddIpv6Range request param
type AddIpv6RangeParam struct {
	BaseParam
	Params AddIpv6RangeParamDetail `json:"params"`
}
// CheckBaremetalChassisConfigFileParamDetail CheckBaremetalChassisConfigFile detail param
type CheckBaremetalChassisConfigFileParamDetail struct {
	BaremetalChassisInfo string `json:"baremetalChassisInfo" validate:"required"`
}

// CheckBaremetalChassisConfigFileParam CheckBaremetalChassisConfigFile request param
type CheckBaremetalChassisConfigFileParam struct {
	BaseParam
	Params CheckBaremetalChassisConfigFileParamDetail `json:"params"`
}
// DeleteOssBucketFileRemoteParamDetail DeleteOssBucketFileRemote detail param
type DeleteOssBucketFileRemoteParamDetail struct {
	FileName string `json:"fileName" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteOssBucketFileRemoteParam DeleteOssBucketFileRemote request param
type DeleteOssBucketFileRemoteParam struct {
	BaseParam
	Params DeleteOssBucketFileRemoteParamDetail `json:"deleteOssBucketFileRemote"`
}
// ChangeMulticastRouterStateParamDetail ChangeMulticastRouterState detail param
type ChangeMulticastRouterStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeMulticastRouterStateParam ChangeMulticastRouterState request param
type ChangeMulticastRouterStateParam struct {
	BaseParam
	Params ChangeMulticastRouterStateParamDetail `json:"changeMulticastRouterState"`
}
// GetMaaSUsageParamDetail GetMaaSUsage detail param
type GetMaaSUsageParamDetail struct {
}

// GetMaaSUsageParam GetMaaSUsage request param
type GetMaaSUsageParam struct {
	BaseParam
	Params GetMaaSUsageParamDetail `json:"getMaaSUsage"`
}
// GetFreeIpParamDetail GetFreeIp detail param
type GetFreeIpParamDetail struct {
	L3NetworkUuid *string `json:"l3NetworkUuid,omitempty"`
	IpRangeUuid *string `json:"ipRangeUuid,omitempty"`
	Start *string `json:"start,omitempty"`
	IpRangeType *string `json:"ipRangeType,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
	Limit int `json:"limit,omitempty"`
}

// GetFreeIpParam GetFreeIp request param
type GetFreeIpParam struct {
	BaseParam
	Params GetFreeIpParamDetail `json:"getFreeIp"`
}
// DeleteOssBucketRemoteParamDetail DeleteOssBucketRemote detail param
type DeleteOssBucketRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteOssBucketRemoteParam DeleteOssBucketRemote request param
type DeleteOssBucketRemoteParam struct {
	BaseParam
	Params DeleteOssBucketRemoteParamDetail `json:"deleteOssBucketRemote"`
}
// CreateL2PortGroupParamDetail CreateL2PortGroup detail param
type CreateL2PortGroupParamDetail struct {
	VSwitchUuid string `json:"vSwitchUuid" validate:"required"`
	VlanMode *string `json:"vlanMode,omitempty"`
	Vlan int `json:"vlan" validate:"required"`
	VlanRanges *string `json:"vlanRanges,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	PhysicalInterface *string `json:"physicalInterface,omitempty"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Pvlan *string `json:"pvlan,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2PortGroupParam CreateL2PortGroup request param
type CreateL2PortGroupParam struct {
	BaseParam
	Params CreateL2PortGroupParamDetail `json:"params"`
}
// ValidateInstanceOfferingUserConfigParamDetail ValidateInstanceOfferingUserConfig detail param
type ValidateInstanceOfferingUserConfigParamDetail struct {
	Config string `json:"config" validate:"required"`
}

// ValidateInstanceOfferingUserConfigParam ValidateInstanceOfferingUserConfig request param
type ValidateInstanceOfferingUserConfigParam struct {
	BaseParam
	Params ValidateInstanceOfferingUserConfigParamDetail `json:"validateInstanceOfferingUserConfig"`
}
// SetVmHostnameParamDetail SetVmHostname detail param
type SetVmHostnameParamDetail struct {
	Hostname string `json:"hostname" validate:"required"`
}

// SetVmHostnameParam SetVmHostname request param
type SetVmHostnameParam struct {
	BaseParam
	Params SetVmHostnameParamDetail `json:"setVmHostname"`
}
// TriggerGCJobParamDetail TriggerGCJob detail param
type TriggerGCJobParamDetail struct {
}

// TriggerGCJobParam TriggerGCJob request param
type TriggerGCJobParam struct {
	BaseParam
	Params TriggerGCJobParamDetail `json:"triggerGCJob"`
}
// CheckBareMetal2IpmiChassisConfigFileParamDetail CheckBareMetal2IpmiChassisConfigFile detail param
type CheckBareMetal2IpmiChassisConfigFileParamDetail struct {
	ChassisInfo string `json:"chassisInfo" validate:"required"`
}

// CheckBareMetal2IpmiChassisConfigFileParam CheckBareMetal2IpmiChassisConfigFile request param
type CheckBareMetal2IpmiChassisConfigFileParam struct {
	BaseParam
	Params CheckBareMetal2IpmiChassisConfigFileParamDetail `json:"params"`
}
// DeleteVirtualRouterLocalParamDetail DeleteVirtualRouterLocal detail param
type DeleteVirtualRouterLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVirtualRouterLocalParam DeleteVirtualRouterLocal request param
type DeleteVirtualRouterLocalParam struct {
	BaseParam
	Params DeleteVirtualRouterLocalParamDetail `json:"deleteVirtualRouterLocal"`
}
// DeleteVpcIkeConfigLocalParamDetail DeleteVpcIkeConfigLocal detail param
type DeleteVpcIkeConfigLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVpcIkeConfigLocalParam DeleteVpcIkeConfigLocal request param
type DeleteVpcIkeConfigLocalParam struct {
	BaseParam
	Params DeleteVpcIkeConfigLocalParamDetail `json:"deleteVpcIkeConfigLocal"`
}
// CreateOssBucketRemoteParamDetail CreateOssBucketRemote detail param
type CreateOssBucketRemoteParamDetail struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	BucketName string `json:"bucketName" validate:"required"`
	Description *string `json:"description,omitempty"`
	OssDomain *string `json:"ossDomain,omitempty"`
	OssKey *string `json:"ossKey,omitempty"`
	OssSecret *string `json:"ossSecret,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateOssBucketRemoteParam CreateOssBucketRemote request param
type CreateOssBucketRemoteParam struct {
	BaseParam
	Params CreateOssBucketRemoteParamDetail `json:"params"`
}
// AddSimulatorPrimaryStorageParamDetail AddSimulatorPrimaryStorage detail param
type AddSimulatorPrimaryStorageParamDetail struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSimulatorPrimaryStorageParam AddSimulatorPrimaryStorage request param
type AddSimulatorPrimaryStorageParam struct {
	BaseParam
	Params AddSimulatorPrimaryStorageParamDetail `json:"params"`
}
// DetachVRouterRouteTableFromVRouterParamDetail DetachVRouterRouteTableFromVRouter detail param
type DetachVRouterRouteTableFromVRouterParamDetail struct {
}

// DetachVRouterRouteTableFromVRouterParam DetachVRouterRouteTableFromVRouter request param
type DetachVRouterRouteTableFromVRouterParam struct {
	BaseParam
	Params DetachVRouterRouteTableFromVRouterParamDetail `json:"detachVRouterRouteTableFromVRouter"`
}
// GetVipUsedPortsParamDetail GetVipUsedPorts detail param
type GetVipUsedPortsParamDetail struct {
	Protocol string `json:"protocol" validate:"required"`
}

// GetVipUsedPortsParam GetVipUsedPorts request param
type GetVipUsedPortsParam struct {
	BaseParam
	Params GetVipUsedPortsParamDetail `json:"getVipUsedPorts"`
}
// SetVmConsolePasswordParamDetail SetVmConsolePassword detail param
type SetVmConsolePasswordParamDetail struct {
	ConsolePassword string `json:"consolePassword" validate:"required"`
}

// SetVmConsolePasswordParam SetVmConsolePassword request param
type SetVmConsolePasswordParam struct {
	BaseParam
	Params SetVmConsolePasswordParamDetail `json:"setVmConsolePassword"`
}
// AttachFirewallRuleSetToL3ParamDetail AttachFirewallRuleSetToL3 detail param
type AttachFirewallRuleSetToL3ParamDetail struct {
	VpcFirewallUuid string `json:"vpcFirewallUuid" validate:"required"`
	Forward string `json:"forward" validate:"required"`
}

// AttachFirewallRuleSetToL3Param AttachFirewallRuleSetToL3 request param
type AttachFirewallRuleSetToL3Param struct {
	BaseParam
	Params AttachFirewallRuleSetToL3ParamDetail `json:"params"`
}
// CleanUpStorageTrashOnPrimaryStorageParamDetail CleanUpStorageTrashOnPrimaryStorage detail param
type CleanUpStorageTrashOnPrimaryStorageParamDetail struct {
	Force *bool `json:"force,omitempty"`
}

// CleanUpStorageTrashOnPrimaryStorageParam CleanUpStorageTrashOnPrimaryStorage request param
type CleanUpStorageTrashOnPrimaryStorageParam struct {
	BaseParam
	Params CleanUpStorageTrashOnPrimaryStorageParamDetail `json:"cleanUpStorageTrashOnPrimaryStorage"`
}
// GetManagementNodeDirCapacityParamDetail GetManagementNodeDirCapacity detail param
type GetManagementNodeDirCapacityParamDetail struct {
	ManagementNodeUuids []string `json:"managementNodeUuids,omitempty"`
}

// GetManagementNodeDirCapacityParam GetManagementNodeDirCapacity request param
type GetManagementNodeDirCapacityParam struct {
	BaseParam
	Params GetManagementNodeDirCapacityParamDetail `json:"getManagementNodeDirCapacity"`
}
// GetGpuDeviceSpecCandidatesParamDetail GetGpuDeviceSpecCandidates detail param
type GetGpuDeviceSpecCandidatesParamDetail struct {
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	VmInstanceUuid *string `json:"vmInstanceUuid,omitempty"`
	VmInstanceUuids []string `json:"vmInstanceUuids,omitempty"`
}

// GetGpuDeviceSpecCandidatesParam GetGpuDeviceSpecCandidates request param
type GetGpuDeviceSpecCandidatesParam struct {
	BaseParam
	Params GetGpuDeviceSpecCandidatesParamDetail `json:"getGpuDeviceSpecCandidates"`
}
// UngroupVolumeSnapshotGroupParamDetail UngroupVolumeSnapshotGroup detail param
type UngroupVolumeSnapshotGroupParamDetail struct {
}

// UngroupVolumeSnapshotGroupParam UngroupVolumeSnapshotGroup request param
type UngroupVolumeSnapshotGroupParam struct {
	BaseParam
	Params UngroupVolumeSnapshotGroupParamDetail `json:"ungroupVolumeSnapshotGroup"`
}
// SubscribeSNSTopicParamDetail SubscribeSNSTopic detail param
type SubscribeSNSTopicParamDetail struct {
}

// SubscribeSNSTopicParam SubscribeSNSTopic request param
type SubscribeSNSTopicParam struct {
	BaseParam
	Params SubscribeSNSTopicParamDetail `json:"params"`
}
// GetCandidateVmNicForSecurityGroupParamDetail GetCandidateVmNicForSecurityGroup detail param
type GetCandidateVmNicForSecurityGroupParamDetail struct {
}

// GetCandidateVmNicForSecurityGroupParam GetCandidateVmNicForSecurityGroup request param
type GetCandidateVmNicForSecurityGroupParam struct {
	BaseParam
	Params GetCandidateVmNicForSecurityGroupParamDetail `json:"getCandidateVmNicForSecurityGroup"`
}
// GetVmRDPParamDetail GetVmRDP detail param
type GetVmRDPParamDetail struct {
}

// GetVmRDPParam GetVmRDP request param
type GetVmRDPParam struct {
	BaseParam
	Params GetVmRDPParamDetail `json:"getVmRDP"`
}
// AttachPciDeviceToVmParamDetail AttachPciDeviceToVm detail param
type AttachPciDeviceToVmParamDetail struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// AttachPciDeviceToVmParam AttachPciDeviceToVm request param
type AttachPciDeviceToVmParam struct {
	BaseParam
	Params AttachPciDeviceToVmParamDetail `json:"params"`
}
// CleanupBillingUsageParamDetail CleanupBillingUsage detail param
type CleanupBillingUsageParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// CleanupBillingUsageParam CleanupBillingUsage request param
type CleanupBillingUsageParam struct {
	BaseParam
	Params CleanupBillingUsageParamDetail `json:"cleanupBillingUsage"`
}
// GetLdapEntryParamDetail GetLdapEntry detail param
type GetLdapEntryParamDetail struct {
	LdapFilter string `json:"ldapFilter" validate:"required"`
	Limit *int `json:"limit,omitempty"`
	LdapServerUuid *string `json:"ldapServerUuid,omitempty"`
}

// GetLdapEntryParam GetLdapEntry request param
type GetLdapEntryParam struct {
	BaseParam
	Params GetLdapEntryParamDetail `json:"getLdapEntry"`
}
// GetCandidateL2NetworksForAttachingClusterParamDetail GetCandidateL2NetworksForAttachingCluster detail param
type GetCandidateL2NetworksForAttachingClusterParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetCandidateL2NetworksForAttachingClusterParam GetCandidateL2NetworksForAttachingCluster request param
type GetCandidateL2NetworksForAttachingClusterParam struct {
	BaseParam
	Params GetCandidateL2NetworksForAttachingClusterParamDetail `json:"getCandidateL2NetworksForAttachingCluster"`
}
// IsVfNicAvailableInL3NetworkParamDetail IsVfNicAvailableInL3Network detail param
type IsVfNicAvailableInL3NetworkParamDetail struct {
}

// IsVfNicAvailableInL3NetworkParam IsVfNicAvailableInL3Network request param
type IsVfNicAvailableInL3NetworkParam struct {
	BaseParam
	Params IsVfNicAvailableInL3NetworkParamDetail `json:"isVfNicAvailableInL3Network"`
}
// GetAllMetricMetadataParamDetail GetAllMetricMetadata detail param
type GetAllMetricMetadataParamDetail struct {
	Name string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// GetAllMetricMetadataParam GetAllMetricMetadata request param
type GetAllMetricMetadataParam struct {
	BaseParam
	Params GetAllMetricMetadataParamDetail `json:"getAllMetricMetadata"`
}
// AddOssBucketFromRemoteParamDetail AddOssBucketFromRemote detail param
type AddOssBucketFromRemoteParamDetail struct {
	BucketName string `json:"bucketName" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	Description *string `json:"description,omitempty"`
	OssDomain *string `json:"ossDomain,omitempty"`
	OssKey *string `json:"ossKey,omitempty"`
	OssSecret *string `json:"ossSecret,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddOssBucketFromRemoteParam AddOssBucketFromRemote request param
type AddOssBucketFromRemoteParam struct {
	BaseParam
	Params AddOssBucketFromRemoteParamDetail `json:"params"`
}
// SyncVmBackupParamDetail SyncVmBackup detail param
type SyncVmBackupParamDetail struct {
}

// SyncVmBackupParam SyncVmBackup request param
type SyncVmBackupParam struct {
	BaseParam
	Params SyncVmBackupParamDetail `json:"syncVmBackup"`
}
// RefreshGuestOsMetadataParamDetail RefreshGuestOsMetadata detail param
type RefreshGuestOsMetadataParamDetail struct {
}

// RefreshGuestOsMetadataParam RefreshGuestOsMetadata request param
type RefreshGuestOsMetadataParam struct {
	BaseParam
	Params RefreshGuestOsMetadataParamDetail `json:"refreshGuestOsMetadata"`
}
// GCAliyunSnapshotRemoteParamDetail GCAliyunSnapshotRemote detail param
type GCAliyunSnapshotRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// GCAliyunSnapshotRemoteParam GCAliyunSnapshotRemote request param
type GCAliyunSnapshotRemoteParam struct {
	BaseParam
	Params GCAliyunSnapshotRemoteParamDetail `json:"params"`
}
// DownloadBackupFileFromPublicCloudParamDetail DownloadBackupFileFromPublicCloud detail param
type DownloadBackupFileFromPublicCloudParamDetail struct {
	RegionId string `json:"regionId" validate:"required"`
	File string `json:"file" validate:"required"`
	Type string `json:"type" validate:"required"`
}

// DownloadBackupFileFromPublicCloudParam DownloadBackupFileFromPublicCloud request param
type DownloadBackupFileFromPublicCloudParam struct {
	BaseParam
	Params DownloadBackupFileFromPublicCloudParamDetail `json:"params"`
}
// AddIAM2VirtualIDsToProjectsParamDetail AddIAM2VirtualIDsToProjects detail param
type AddIAM2VirtualIDsToProjectsParamDetail struct {
	ProjectUuids []string `json:"projectUuids" validate:"required"`
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
	RoleUuids []string `json:"roleUuids,omitempty"`
}

// AddIAM2VirtualIDsToProjectsParam AddIAM2VirtualIDsToProjects request param
type AddIAM2VirtualIDsToProjectsParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToProjectsParamDetail `json:"params"`
}
// CreateIAM2ProjectTemplateFromProjectParamDetail CreateIAM2ProjectTemplateFromProject detail param
type CreateIAM2ProjectTemplateFromProjectParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectTemplateFromProjectParam CreateIAM2ProjectTemplateFromProject request param
type CreateIAM2ProjectTemplateFromProjectParam struct {
	BaseParam
	Params CreateIAM2ProjectTemplateFromProjectParamDetail `json:"params"`
}
// CreateTagParamDetail CreateTag detail param
type CreateTagParamDetail struct {
	Name string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
	Description *string `json:"description,omitempty"`
	Color *string `json:"color,omitempty"`
	Type *string `json:"type,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateTagParam CreateTag request param
type CreateTagParam struct {
	BaseParam
	Params CreateTagParamDetail `json:"params"`
}
// UpdateConsolePasswordParamDetail UpdateConsolePassword detail param
type UpdateConsolePasswordParamDetail struct {
	Password string `json:"password" validate:"required"`
}

// UpdateConsolePasswordParam UpdateConsolePassword request param
type UpdateConsolePasswordParam struct {
	BaseParam
	Params UpdateConsolePasswordParamDetail `json:"updateConsolePassword"`
}
// CreateVmInstanceFromVolumeSnapshotGroupParamDetail CreateVmInstanceFromVolumeSnapshotGroup detail param
type CreateVmInstanceFromVolumeSnapshotGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	InstanceOfferingUuid *string `json:"instanceOfferingUuid,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	ReservedMemorySize *int64 `json:"reservedMemorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags map[string]interface{} `json:"dataVolumeSystemTags,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceFromVolumeSnapshotGroupParam CreateVmInstanceFromVolumeSnapshotGroup request param
type CreateVmInstanceFromVolumeSnapshotGroupParam struct {
	BaseParam
	Params CreateVmInstanceFromVolumeSnapshotGroupParamDetail `json:"params"`
}
// SetIAM2ProjectRetirePolicyParamDetail SetIAM2ProjectRetirePolicy detail param
type SetIAM2ProjectRetirePolicyParamDetail struct {
	Policy string `json:"policy" validate:"required"`
}

// SetIAM2ProjectRetirePolicyParam SetIAM2ProjectRetirePolicy request param
type SetIAM2ProjectRetirePolicyParam struct {
	BaseParam
	Params SetIAM2ProjectRetirePolicyParamDetail `json:"setIAM2ProjectRetirePolicy"`
}
// RunIAM2ScriptParamDetail RunIAM2Script detail param
type RunIAM2ScriptParamDetail struct {
	ScriptContent string `json:"scriptContent" validate:"required"`
	ScriptExecutor *string `json:"scriptExecutor,omitempty"`
	ScriptParams []string `json:"scriptParams,omitempty"`
}

// RunIAM2ScriptParam RunIAM2Script request param
type RunIAM2ScriptParam struct {
	BaseParam
	Params RunIAM2ScriptParamDetail `json:"params"`
}
// AttachServiceToObservabilityServerParamDetail AttachServiceToObservabilityServer detail param
type AttachServiceToObservabilityServerParamDetail struct {
	ServiceType string `json:"serviceType" validate:"required"`
	ServiceUuid string `json:"serviceUuid" validate:"required"`
}

// AttachServiceToObservabilityServerParam AttachServiceToObservabilityServer request param
type AttachServiceToObservabilityServerParam struct {
	BaseParam
	Params AttachServiceToObservabilityServerParamDetail `json:"params"`
}
// DeleteHostNetworkServiceTypeParamDetail DeleteHostNetworkServiceType detail param
type DeleteHostNetworkServiceTypeParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteHostNetworkServiceTypeParam DeleteHostNetworkServiceType request param
type DeleteHostNetworkServiceTypeParam struct {
	BaseParam
	Params DeleteHostNetworkServiceTypeParamDetail `json:"deleteHostNetworkServiceType"`
}
// SubscribeResNotifyParamDetail SubscribeResNotify detail param
type SubscribeResNotifyParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceTypes []string `json:"resourceTypes,omitempty"`
	EventTypes []string `json:"eventTypes,omitempty"`
	Type *string `json:"type,omitempty"`
	WebhookUrl string `json:"webhookUrl" validate:"required"`
	Secret *string `json:"secret,omitempty"`
	CustomHeaders *string `json:"customHeaders,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SubscribeResNotifyParam SubscribeResNotify request param
type SubscribeResNotifyParam struct {
	BaseParam
	Params SubscribeResNotifyParamDetail `json:"params"`
}
// CreateIAM2ProjectFromTemplateParamDetail CreateIAM2ProjectFromTemplate detail param
type CreateIAM2ProjectFromTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	RoleUuids []string `json:"roleUuids,omitempty"`
	OrganizationUuid *string `json:"organizationUuid,omitempty"`
	ResourceTemplates []string `json:"resourceTemplates,omitempty"`
	LinkAccountUuid *string `json:"linkAccountUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectFromTemplateParam CreateIAM2ProjectFromTemplate request param
type CreateIAM2ProjectFromTemplateParam struct {
	BaseParam
	Params CreateIAM2ProjectFromTemplateParamDetail `json:"params"`
}
// AddConnectionAccessPointFromRemoteParamDetail AddConnectionAccessPointFromRemote detail param
type AddConnectionAccessPointFromRemoteParamDetail struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	AccessPointId string `json:"accessPointId" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddConnectionAccessPointFromRemoteParam AddConnectionAccessPointFromRemote request param
type AddConnectionAccessPointFromRemoteParam struct {
	BaseParam
	Params AddConnectionAccessPointFromRemoteParamDetail `json:"params"`
}
// AttachSshKeyPairToVmInstanceParamDetail AttachSshKeyPairToVmInstance detail param
type AttachSshKeyPairToVmInstanceParamDetail struct {
}

// AttachSshKeyPairToVmInstanceParam AttachSshKeyPairToVmInstance request param
type AttachSshKeyPairToVmInstanceParam struct {
	BaseParam
	Params AttachSshKeyPairToVmInstanceParamDetail `json:"params"`
}
// DetachBareMetal2GatewayFromClusterParamDetail DetachBareMetal2GatewayFromCluster detail param
type DetachBareMetal2GatewayFromClusterParamDetail struct {
}

// DetachBareMetal2GatewayFromClusterParam DetachBareMetal2GatewayFromCluster request param
type DetachBareMetal2GatewayFromClusterParam struct {
	BaseParam
	Params DetachBareMetal2GatewayFromClusterParamDetail `json:"detachBareMetal2GatewayFromCluster"`
}
// ReloadElaborationParamDetail ReloadElaboration detail param
type ReloadElaborationParamDetail struct {
}

// ReloadElaborationParam ReloadElaboration request param
type ReloadElaborationParam struct {
	BaseParam
	Params ReloadElaborationParamDetail `json:"reloadElaboration"`
}
// ReconnectVirtualRouterParamDetail ReconnectVirtualRouter detail param
type ReconnectVirtualRouterParamDetail struct {
}

// ReconnectVirtualRouterParam ReconnectVirtualRouter request param
type ReconnectVirtualRouterParam struct {
	BaseParam
	Params ReconnectVirtualRouterParamDetail `json:"reconnectVirtualRouter"`
}
// ConvertVmFromForeignHypervisorParamDetail ConvertVmFromForeignHypervisor detail param
type ConvertVmFromForeignHypervisorParamDetail struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ConversionHostUuid *string `json:"conversionHostUuid,omitempty"`
	SshPrivKey *string `json:"sshPrivKey,omitempty"`
	CpuNum int `json:"cpuNum" validate:"required"`
	MemorySize int64 `json:"memorySize" validate:"required"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Type *string `json:"type,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	ConvertStrategy *string `json:"convertStrategy,omitempty"`
	PauseVm *bool `json:"pauseVm,omitempty"`
	VolumeFilters []VolumeFilterInfoParam `json:"volumeFilters,omitempty"`
	RootFileSystem *string `json:"rootFileSystem,omitempty"`
	LongJobName *string `json:"longJobName,omitempty"`
	LongJobDescription *string `json:"longJobDescription,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// ConvertVmFromForeignHypervisorParam ConvertVmFromForeignHypervisor request param
type ConvertVmFromForeignHypervisorParam struct {
	BaseParam
	Params ConvertVmFromForeignHypervisorParamDetail `json:"params"`
}
// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchParamDetail DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch detail param
type DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchParam DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch request param
type DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchParam struct {
	BaseParam
	Params DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchParamDetail `json:"deleteConnectionBetweenL3NetWorkAndAliyunVSwitch"`
}
// RestartResourceStackParamDetail RestartResourceStack detail param
type RestartResourceStackParamDetail struct {
}

// RestartResourceStackParam RestartResourceStack request param
type RestartResourceStackParam struct {
	BaseParam
	Params RestartResourceStackParamDetail `json:"restartResourceStack"`
}
// SyncEcsImageFromRemoteParamDetail SyncEcsImageFromRemote detail param
type SyncEcsImageFromRemoteParamDetail struct {
	Type *string `json:"type,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsImageFromRemoteParam SyncEcsImageFromRemote request param
type SyncEcsImageFromRemoteParam struct {
	BaseParam
	Params SyncEcsImageFromRemoteParamDetail `json:"params"`
}
// AttachPoliciesToUserParamDetail AttachPoliciesToUser detail param
type AttachPoliciesToUserParamDetail struct {
	PolicyUuids []string `json:"policyUuids" validate:"required"`
}

// AttachPoliciesToUserParam AttachPoliciesToUser request param
type AttachPoliciesToUserParam struct {
	BaseParam
	Params AttachPoliciesToUserParamDetail `json:"params"`
}
// AttachBackupStorageToZoneParamDetail AttachBackupStorageToZone detail param
type AttachBackupStorageToZoneParamDetail struct {
}

// AttachBackupStorageToZoneParam AttachBackupStorageToZone request param
type AttachBackupStorageToZoneParam struct {
	BaseParam
	Params AttachBackupStorageToZoneParamDetail `json:"params"`
}
// AddPciDeviceSpecToVmInstanceParamDetail AddPciDeviceSpecToVmInstance detail param
type AddPciDeviceSpecToVmInstanceParamDetail struct {
	PciDeviceNumber *int `json:"pciDeviceNumber,omitempty"`
}

// AddPciDeviceSpecToVmInstanceParam AddPciDeviceSpecToVmInstance request param
type AddPciDeviceSpecToVmInstanceParam struct {
	BaseParam
	Params AddPciDeviceSpecToVmInstanceParamDetail `json:"params"`
}
// ResizeRootVolumeParamDetail ResizeRootVolume detail param
type ResizeRootVolumeParamDetail struct {
	VmInstanceUuid *string `json:"vmInstanceUuid,omitempty"`
	Size int64 `json:"size" validate:"required"`
}

// ResizeRootVolumeParam ResizeRootVolume request param
type ResizeRootVolumeParam struct {
	BaseParam
	Params ResizeRootVolumeParamDetail `json:"resizeRootVolume"`
}
// GetVpcVpnConfigurationFromRemoteParamDetail GetVpcVpnConfigurationFromRemote detail param
type GetVpcVpnConfigurationFromRemoteParamDetail struct {
}

// GetVpcVpnConfigurationFromRemoteParam GetVpcVpnConfigurationFromRemote request param
type GetVpcVpnConfigurationFromRemoteParam struct {
	BaseParam
	Params GetVpcVpnConfigurationFromRemoteParamDetail `json:"getVpcVpnConfigurationFromRemote"`
}
// CreateImageGroupFromImageParamDetail CreateImageGroupFromImage detail param
type CreateImageGroupFromImageParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	DataVolumeTemplateUuids []string `json:"dataVolumeTemplateUuids,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateImageGroupFromImageParam CreateImageGroupFromImage request param
type CreateImageGroupFromImageParam struct {
	BaseParam
	Params CreateImageGroupFromImageParamDetail `json:"params"`
}
// TokenIntrospectionParamDetail TokenIntrospection detail param
type TokenIntrospectionParamDetail struct {
	Token string `json:"token" validate:"required"`
	TokenType string `json:"tokenType" validate:"required"`
}

// TokenIntrospectionParam TokenIntrospection request param
type TokenIntrospectionParam struct {
	BaseParam
	Params TokenIntrospectionParamDetail `json:"params"`
}
// SyncVmBackupFromImageStoreBackupStorageParamDetail SyncVmBackupFromImageStoreBackupStorage detail param
type SyncVmBackupFromImageStoreBackupStorageParamDetail struct {
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// SyncVmBackupFromImageStoreBackupStorageParam SyncVmBackupFromImageStoreBackupStorage request param
type SyncVmBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncVmBackupFromImageStoreBackupStorageParamDetail `json:"syncVmBackupFromImageStoreBackupStorage"`
}
// AddCertificateToLoadBalancerListenerParamDetail AddCertificateToLoadBalancerListener detail param
type AddCertificateToLoadBalancerListenerParamDetail struct {
	CertificateUuid string `json:"certificateUuid" validate:"required"`
}

// AddCertificateToLoadBalancerListenerParam AddCertificateToLoadBalancerListener request param
type AddCertificateToLoadBalancerListenerParam struct {
	BaseParam
	Params AddCertificateToLoadBalancerListenerParamDetail `json:"params"`
}
// AddRolesToIAM2VirtualIDParamDetail AddRolesToIAM2VirtualID detail param
type AddRolesToIAM2VirtualIDParamDetail struct {
	RoleUuids []string `json:"roleUuids" validate:"required"`
	ProjectUuid *string `json:"projectUuid,omitempty"`
}

// AddRolesToIAM2VirtualIDParam AddRolesToIAM2VirtualID request param
type AddRolesToIAM2VirtualIDParam struct {
	BaseParam
	Params AddRolesToIAM2VirtualIDParamDetail `json:"params"`
}
// CreateFaultToleranceVmInstanceParamDetail CreateFaultToleranceVmInstance detail param
type CreateFaultToleranceVmInstanceParamDetail struct {
	Name string `json:"name" validate:"required"`
	InstanceOfferingUuid *string `json:"instanceOfferingUuid,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	ReservedMemorySize *int64 `json:"reservedMemorySize,omitempty"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	Description *string `json:"description,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	Type *string `json:"type,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFaultToleranceVmInstanceParam CreateFaultToleranceVmInstance request param
type CreateFaultToleranceVmInstanceParam struct {
	BaseParam
	Params CreateFaultToleranceVmInstanceParamDetail `json:"params"`
}
// DeleteResourceStackVmPortMonitorParamDetail DeleteResourceStackVmPortMonitor detail param
type DeleteResourceStackVmPortMonitorParamDetail struct {
	StackUuid *string `json:"stackUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Port *int `json:"port,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteResourceStackVmPortMonitorParam DeleteResourceStackVmPortMonitor request param
type DeleteResourceStackVmPortMonitorParam struct {
	BaseParam
	Params DeleteResourceStackVmPortMonitorParamDetail `json:"deleteResourceStackVmPortMonitor"`
}
// DeleteGCJobParamDetail DeleteGCJob detail param
type DeleteGCJobParamDetail struct {
}

// DeleteGCJobParam DeleteGCJob request param
type DeleteGCJobParam struct {
	BaseParam
	Params DeleteGCJobParamDetail `json:"deleteGCJob"`
}
// DeleteEmailAddressOfSNSEmailEndpointParamDetail DeleteEmailAddressOfSNSEmailEndpoint detail param
type DeleteEmailAddressOfSNSEmailEndpointParamDetail struct {
}

// DeleteEmailAddressOfSNSEmailEndpointParam DeleteEmailAddressOfSNSEmailEndpoint request param
type DeleteEmailAddressOfSNSEmailEndpointParam struct {
	BaseParam
	Params DeleteEmailAddressOfSNSEmailEndpointParamDetail `json:"deleteEmailAddressOfSNSEmailEndpoint"`
}
// CleanInvalidLdapIAM2BindingParamDetail CleanInvalidLdapIAM2Binding detail param
type CleanInvalidLdapIAM2BindingParamDetail struct {
}

// CleanInvalidLdapIAM2BindingParam CleanInvalidLdapIAM2Binding request param
type CleanInvalidLdapIAM2BindingParam struct {
	BaseParam
	Params CleanInvalidLdapIAM2BindingParamDetail `json:"cleanInvalidLdapIAM2Binding"`
}
// UpdateHybridEipParamDetail UpdateHybridEip detail param
type UpdateHybridEipParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
}

// UpdateHybridEipParam UpdateHybridEip request param
type UpdateHybridEipParam struct {
	BaseParam
	Params UpdateHybridEipParamDetail `json:"updateHybridEip"`
}
// GetVpcAttachedIpsecParamDetail GetVpcAttachedIpsec detail param
type GetVpcAttachedIpsecParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVpcAttachedIpsecParam GetVpcAttachedIpsec request param
type GetVpcAttachedIpsecParam struct {
	BaseParam
	Params GetVpcAttachedIpsecParamDetail `json:"params"`
}
// GetImagesFromImageStoreBackupStorageParamDetail GetImagesFromImageStoreBackupStorage detail param
type GetImagesFromImageStoreBackupStorageParamDetail struct {
}

// GetImagesFromImageStoreBackupStorageParam GetImagesFromImageStoreBackupStorage request param
type GetImagesFromImageStoreBackupStorageParam struct {
	BaseParam
	Params GetImagesFromImageStoreBackupStorageParamDetail `json:"getImagesFromImageStoreBackupStorage"`
}
// GetElaborationCategoriesParamDetail GetElaborationCategories detail param
type GetElaborationCategoriesParamDetail struct {
}

// GetElaborationCategoriesParam GetElaborationCategories request param
type GetElaborationCategoriesParam struct {
	BaseParam
	Params GetElaborationCategoriesParamDetail `json:"getElaborationCategories"`
}
// GetScsiLunCandidatesForAttachingVmParamDetail GetScsiLunCandidatesForAttachingVm detail param
type GetScsiLunCandidatesForAttachingVmParamDetail struct {
}

// GetScsiLunCandidatesForAttachingVmParam GetScsiLunCandidatesForAttachingVm request param
type GetScsiLunCandidatesForAttachingVmParam struct {
	BaseParam
	Params GetScsiLunCandidatesForAttachingVmParamDetail `json:"getScsiLunCandidatesForAttachingVm"`
}
// GetHostMultipathTopologyParamDetail GetHostMultipathTopology detail param
type GetHostMultipathTopologyParamDetail struct {
	HostUuid string `json:"hostUuid" validate:"required"`
	DiskUuids []string `json:"diskUuids" validate:"required"`
}

// GetHostMultipathTopologyParam GetHostMultipathTopology request param
type GetHostMultipathTopologyParam struct {
	BaseParam
	Params GetHostMultipathTopologyParamDetail `json:"getHostMultipathTopology"`
}
// DeleteEcsImageRemoteParamDetail DeleteEcsImageRemote detail param
type DeleteEcsImageRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsImageRemoteParam DeleteEcsImageRemote request param
type DeleteEcsImageRemoteParam struct {
	BaseParam
	Params DeleteEcsImageRemoteParamDetail `json:"deleteEcsImageRemote"`
}
// GetHostNetworkFactsParamDetail GetHostNetworkFacts detail param
type GetHostNetworkFactsParamDetail struct {
}

// GetHostNetworkFactsParam GetHostNetworkFacts request param
type GetHostNetworkFactsParam struct {
	BaseParam
	Params GetHostNetworkFactsParamDetail `json:"getHostNetworkFacts"`
}
// CleanUpTrashOnBackupStorageParamDetail CleanUpTrashOnBackupStorage detail param
type CleanUpTrashOnBackupStorageParamDetail struct {
	TrashId *int64 `json:"trashId,omitempty"`
}

// CleanUpTrashOnBackupStorageParam CleanUpTrashOnBackupStorage request param
type CleanUpTrashOnBackupStorageParam struct {
	BaseParam
	Params CleanUpTrashOnBackupStorageParamDetail `json:"cleanUpTrashOnBackupStorage"`
}
// CreateConnectionBetweenL3NetworkAndAliyunVSwitchParamDetail CreateConnectionBetweenL3NetworkAndAliyunVSwitch detail param
type CreateConnectionBetweenL3NetworkAndAliyunVSwitchParamDetail struct {
	L3networkUuid string `json:"l3networkUuid" validate:"required"`
	VpcUuid string `json:"vpcUuid" validate:"required"`
	VbrUuid string `json:"vbrUuid" validate:"required"`
	CpeIp string `json:"cpeIp" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Direction string `json:"direction" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateConnectionBetweenL3NetworkAndAliyunVSwitchParam CreateConnectionBetweenL3NetworkAndAliyunVSwitch request param
type CreateConnectionBetweenL3NetworkAndAliyunVSwitchParam struct {
	BaseParam
	Params CreateConnectionBetweenL3NetworkAndAliyunVSwitchParamDetail `json:"params"`
}
// DetachPriceTableFromAccountParamDetail DetachPriceTableFromAccount detail param
type DetachPriceTableFromAccountParamDetail struct {
}

// DetachPriceTableFromAccountParam DetachPriceTableFromAccount request param
type DetachPriceTableFromAccountParam struct {
	BaseParam
	Params DetachPriceTableFromAccountParamDetail `json:"detachPriceTableFromAccount"`
}
// AddVRouterNetworksToFlowMeterParamDetail AddVRouterNetworksToFlowMeter detail param
type AddVRouterNetworksToFlowMeterParamDetail struct {
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddVRouterNetworksToFlowMeterParam AddVRouterNetworksToFlowMeter request param
type AddVRouterNetworksToFlowMeterParam struct {
	BaseParam
	Params AddVRouterNetworksToFlowMeterParamDetail `json:"params"`
}
// GetIAM2VirtualIDInGroupParamDetail GetIAM2VirtualIDInGroup detail param
type GetIAM2VirtualIDInGroupParamDetail struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
	Count *bool `json:"count,omitempty"`
	SortDirection *string `json:"sortDirection,omitempty"`
	SortBy *string `json:"sortBy,omitempty"`
	ReplyWithCount *bool `json:"replyWithCount,omitempty"`
}

// GetIAM2VirtualIDInGroupParam GetIAM2VirtualIDInGroup request param
type GetIAM2VirtualIDInGroupParam struct {
	BaseParam
	Params GetIAM2VirtualIDInGroupParamDetail `json:"getIAM2VirtualIDInGroup"`
}
// UnlockIdentityParamDetail UnlockIdentity detail param
type UnlockIdentityParamDetail struct {
	ResourceName string `json:"resourceName" validate:"required"`
	LoginType string `json:"loginType" validate:"required"`
}

// UnlockIdentityParam UnlockIdentity request param
type UnlockIdentityParam struct {
	BaseParam
	Params UnlockIdentityParamDetail `json:"unlockIdentity"`
}
// ChangeVmSchedulingRuleStateParamDetail ChangeVmSchedulingRuleState detail param
type ChangeVmSchedulingRuleStateParamDetail struct {
	State string `json:"state" validate:"required"`
}

// ChangeVmSchedulingRuleStateParam ChangeVmSchedulingRuleState request param
type ChangeVmSchedulingRuleStateParam struct {
	BaseParam
	Params ChangeVmSchedulingRuleStateParamDetail `json:"changeVmSchedulingRuleState"`
}
// GetCandidateVmNicsForPortMirrorParamDetail GetCandidateVmNicsForPortMirror detail param
type GetCandidateVmNicsForPortMirrorParamDetail struct {
}

// GetCandidateVmNicsForPortMirrorParam GetCandidateVmNicsForPortMirror request param
type GetCandidateVmNicsForPortMirrorParam struct {
	BaseParam
	Params GetCandidateVmNicsForPortMirrorParamDetail `json:"getCandidateVmNicsForPortMirror"`
}
// CreateFirewallRuleParamDetail CreateFirewallRule detail param
type CreateFirewallRuleParamDetail struct {
	RuleSetUuid string `json:"ruleSetUuid" validate:"required"`
	Action string `json:"action" validate:"required"`
	Protocol *string `json:"protocol,omitempty"`
	DestPort *string `json:"destPort,omitempty"`
	SourcePort *string `json:"sourcePort,omitempty"`
	SourceIp *string `json:"sourceIp,omitempty"`
	DestIp *string `json:"destIp,omitempty"`
	AllowStates *string `json:"allowStates,omitempty"`
	TcpFlag *string `json:"tcpFlag,omitempty"`
	IcmpTypeName *string `json:"icmpTypeName,omitempty"`
	RuleNumber int `json:"ruleNumber" validate:"required"`
	EnableLog *bool `json:"enableLog,omitempty"`
	State string `json:"state" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFirewallRuleParam CreateFirewallRule request param
type CreateFirewallRuleParam struct {
	BaseParam
	Params CreateFirewallRuleParamDetail `json:"params"`
}
// GetVmEmulatorPinningParamDetail GetVmEmulatorPinning detail param
type GetVmEmulatorPinningParamDetail struct {
}

// GetVmEmulatorPinningParam GetVmEmulatorPinning request param
type GetVmEmulatorPinningParam struct {
	BaseParam
	Params GetVmEmulatorPinningParamDetail `json:"getVmEmulatorPinning"`
}
// GetDataVolumeAttachableVmParamDetail GetDataVolumeAttachableVm detail param
type GetDataVolumeAttachableVmParamDetail struct {
}

// GetDataVolumeAttachableVmParam GetDataVolumeAttachableVm request param
type GetDataVolumeAttachableVmParam struct {
	BaseParam
	Params GetDataVolumeAttachableVmParamDetail `json:"getDataVolumeAttachableVm"`
}
// AddIpRangeByNetworkCidrParamDetail AddIpRangeByNetworkCidr detail param
type AddIpRangeByNetworkCidrParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	NetworkCidr string `json:"networkCidr" validate:"required"`
	Gateway *string `json:"gateway,omitempty"`
	IpRangeType *string `json:"ipRangeType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIpRangeByNetworkCidrParam AddIpRangeByNetworkCidr request param
type AddIpRangeByNetworkCidrParam struct {
	BaseParam
	Params AddIpRangeByNetworkCidrParamDetail `json:"params"`
}
// CreateL2NoVlanNetworkParamDetail CreateL2NoVlanNetwork detail param
type CreateL2NoVlanNetworkParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	PhysicalInterface *string `json:"physicalInterface,omitempty"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Pvlan *string `json:"pvlan,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2NoVlanNetworkParam CreateL2NoVlanNetwork request param
type CreateL2NoVlanNetworkParam struct {
	BaseParam
	Params CreateL2NoVlanNetworkParamDetail `json:"params"`
}
// AddMonToCephBackupStorageParamDetail AddMonToCephBackupStorage detail param
type AddMonToCephBackupStorageParamDetail struct {
	MonUrls []string `json:"monUrls" validate:"required"`
}

// AddMonToCephBackupStorageParam AddMonToCephBackupStorage request param
type AddMonToCephBackupStorageParam struct {
	BaseParam
	Params AddMonToCephBackupStorageParamDetail `json:"params"`
}
// DetachBareMetal2ProvisionNetworkFromClusterParamDetail DetachBareMetal2ProvisionNetworkFromCluster detail param
type DetachBareMetal2ProvisionNetworkFromClusterParamDetail struct {
}

// DetachBareMetal2ProvisionNetworkFromClusterParam DetachBareMetal2ProvisionNetworkFromCluster request param
type DetachBareMetal2ProvisionNetworkFromClusterParam struct {
	BaseParam
	Params DetachBareMetal2ProvisionNetworkFromClusterParamDetail `json:"detachBareMetal2ProvisionNetworkFromCluster"`
}
// DeleteAliyunDiskFromLocalParamDetail DeleteAliyunDiskFromLocal detail param
type DeleteAliyunDiskFromLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAliyunDiskFromLocalParam DeleteAliyunDiskFromLocal request param
type DeleteAliyunDiskFromLocalParam struct {
	BaseParam
	Params DeleteAliyunDiskFromLocalParamDetail `json:"deleteAliyunDiskFromLocal"`
}
// GetResourceNamesParamDetail GetResourceNames detail param
type GetResourceNamesParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// GetResourceNamesParam GetResourceNames request param
type GetResourceNamesParam struct {
	BaseParam
	Params GetResourceNamesParamDetail `json:"getResourceNames"`
}
// GetIAM2VirtualIDAPIPermissionParamDetail GetIAM2VirtualIDPermission detail param
type GetIAM2VirtualIDAPIPermissionParamDetail struct {
	ApisToCheck []string `json:"apisToCheck,omitempty"`
	OnlyCheckParams bool `json:"onlyCheckParams,omitempty"`
}

// GetIAM2VirtualIDAPIPermissionParam GetIAM2VirtualIDPermission request param
type GetIAM2VirtualIDAPIPermissionParam struct {
	BaseParam
	Params GetIAM2VirtualIDAPIPermissionParamDetail `json:"getIAM2VirtualIDAPIPermission"`
}
// GetOrganizationQuotaUsageParamDetail GetOrganizationQuotaUsage detail param
type GetOrganizationQuotaUsageParamDetail struct {
}

// GetOrganizationQuotaUsageParam GetOrganizationQuotaUsage request param
type GetOrganizationQuotaUsageParam struct {
	BaseParam
	Params GetOrganizationQuotaUsageParamDetail `json:"getOrganizationQuotaUsage"`
}
// GetResourceConfigsParamDetail GetResourceConfigs detail param
type GetResourceConfigsParamDetail struct {
	Names []string `json:"names" validate:"required"`
}

// GetResourceConfigsParam GetResourceConfigs request param
type GetResourceConfigsParam struct {
	BaseParam
	Params GetResourceConfigsParamDetail `json:"getResourceConfigs"`
}
// SyncVpcUserVpnGatewayFromRemoteParamDetail SyncVpcUserVpnGatewayFromRemote detail param
type SyncVpcUserVpnGatewayFromRemoteParamDetail struct {
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncVpcUserVpnGatewayFromRemoteParam SyncVpcUserVpnGatewayFromRemote request param
type SyncVpcUserVpnGatewayFromRemoteParam struct {
	BaseParam
	Params SyncVpcUserVpnGatewayFromRemoteParamDetail `json:"syncVpcUserVpnGatewayFromRemote"`
}
// DetachPrimaryStorageFromClusterParamDetail DetachPrimaryStorageFromCluster detail param
type DetachPrimaryStorageFromClusterParamDetail struct {
}

// DetachPrimaryStorageFromClusterParam DetachPrimaryStorageFromCluster request param
type DetachPrimaryStorageFromClusterParam struct {
	BaseParam
	Params DetachPrimaryStorageFromClusterParamDetail `json:"detachPrimaryStorageFromCluster"`
}
// CheckStackTemplateParametersParamDetail CheckStackTemplateParameters detail param
type CheckStackTemplateParametersParamDetail struct {
	Type *string `json:"type,omitempty"`
	TemplateContent *string `json:"templateContent,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

// CheckStackTemplateParametersParam CheckStackTemplateParameters request param
type CheckStackTemplateParametersParam struct {
	BaseParam
	Params CheckStackTemplateParametersParamDetail `json:"params"`
}
// GetFactoryModeStateParamDetail GetFactoryModeState detail param
type GetFactoryModeStateParamDetail struct {
}

// GetFactoryModeStateParam GetFactoryModeState request param
type GetFactoryModeStateParam struct {
	BaseParam
	Params GetFactoryModeStateParamDetail `json:"getFactoryModeState"`
}
// AddServerGroupToLoadBalancerListenerParamDetail AddServerGroupToLoadBalancerListener detail param
type AddServerGroupToLoadBalancerListenerParamDetail struct {
	ServerGroupUuid string `json:"serverGroupUuid" validate:"required"`
}

// AddServerGroupToLoadBalancerListenerParam AddServerGroupToLoadBalancerListener request param
type AddServerGroupToLoadBalancerListenerParam struct {
	BaseParam
	Params AddServerGroupToLoadBalancerListenerParamDetail `json:"params"`
}
// GetActiveAlarmStatusParamDetail GetActiveAlarmStatus detail param
type GetActiveAlarmStatusParamDetail struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
}

// GetActiveAlarmStatusParam GetActiveAlarmStatus request param
type GetActiveAlarmStatusParam struct {
	BaseParam
	Params GetActiveAlarmStatusParamDetail `json:"getActiveAlarmStatus"`
}
// DeployModelEvalServiceParamDetail DeployModelEvalService detail param
type DeployModelEvalServiceParamDetail struct {
	TaskType string `json:"taskType" validate:"required"`
	Limits int `json:"limits" validate:"required"`
	Temperature *float32 `json:"temperature,omitempty"`
	TopK *int `json:"topK,omitempty"`
	TopP *float32 `json:"topP,omitempty"`
	MaxNewTokens *int `json:"maxNewTokens,omitempty"`
	RepetitionPenalty *float32 `json:"repetitionPenalty,omitempty"`
	MaxLength *int `json:"maxLength,omitempty"`
	Prompt *string `json:"prompt,omitempty"`
	Model *string `json:"model,omitempty"`
	Url *string `json:"url,omitempty"`
	Parallel *int `json:"parallel,omitempty"`
	LogEveryQuery *int `json:"logEveryQuery,omitempty"`
	Api *string `json:"api,omitempty"`
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`
	ConnectTimeout *int `json:"connectTimeout,omitempty"`
	ReadTimeout *int `json:"readTimeout,omitempty"`
	Description *string `json:"description,omitempty"`
	ModelUuid *string `json:"modelUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	VmImageUuid *string `json:"vmImageUuid,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	DatasetUuids []string `json:"datasetUuids,omitempty"`
	ModelServiceGroupUuids []string `json:"modelServiceGroupUuids,omitempty"`
	DockerImage *string `json:"dockerImage,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	Name string `json:"name" validate:"required"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	StartupParameters map[string]string `json:"startupParameters,omitempty"`
	Type string `json:"type" validate:"required"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	ServiceBootUptime *int `json:"serviceBootUptime,omitempty"`
	ServiceLivez *string `json:"serviceLivez,omitempty"`
	ServiceReadyz *string `json:"serviceReadyz,omitempty"`
	RootDiskOfferingUuid *string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize *int64 `json:"rootDiskSize,omitempty"`
	ProjectUuid *string `json:"projectUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// DeployModelEvalServiceParam DeployModelEvalService request param
type DeployModelEvalServiceParam struct {
	BaseParam
	Params DeployModelEvalServiceParamDetail `json:"deployModelEvalService"`
}
// AttachNvmeServerToClusterParamDetail AttachNvmeServerToCluster detail param
type AttachNvmeServerToClusterParamDetail struct {
}

// AttachNvmeServerToClusterParam AttachNvmeServerToCluster request param
type AttachNvmeServerToClusterParam struct {
	BaseParam
	Params AttachNvmeServerToClusterParamDetail `json:"params"`
}
// GetHostResourceAllocationParamDetail GetHostResourceAllocation detail param
type GetHostResourceAllocationParamDetail struct {
	Strategy string `json:"strategy" validate:"required"`
	Scene string `json:"scene" validate:"required"`
	Vcpu int `json:"vcpu" validate:"required"`
	MemSize *int64 `json:"memSize,omitempty"`
}

// GetHostResourceAllocationParam GetHostResourceAllocation request param
type GetHostResourceAllocationParam struct {
	BaseParam
	Params GetHostResourceAllocationParamDetail `json:"params"`
}
// AttachUsbDeviceToVmParamDetail AttachUsbDeviceToVm detail param
type AttachUsbDeviceToVmParamDetail struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	AttachType *string `json:"attachType,omitempty"`
}

// AttachUsbDeviceToVmParam AttachUsbDeviceToVm request param
type AttachUsbDeviceToVmParam struct {
	BaseParam
	Params AttachUsbDeviceToVmParamDetail `json:"params"`
}
// GetLicenseAddOnsParamDetail GetLicenseAddOns detail param
type GetLicenseAddOnsParamDetail struct {
}

// GetLicenseAddOnsParam GetLicenseAddOns request param
type GetLicenseAddOnsParam struct {
	BaseParam
	Params GetLicenseAddOnsParamDetail `json:"getLicenseAddOns"`
}
// SyncAliyunSnapshotRemoteParamDetail SyncAliyunSnapshotRemote detail param
type SyncAliyunSnapshotRemoteParamDetail struct {
	SnapshotId *string `json:"snapshotId,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncAliyunSnapshotRemoteParam SyncAliyunSnapshotRemote request param
type SyncAliyunSnapshotRemoteParam struct {
	BaseParam
	Params SyncAliyunSnapshotRemoteParamDetail `json:"params"`
}
// UpdateTicketRequestParamDetail UpdateTicketRequest detail param
type UpdateTicketRequestParamDetail struct {
	Requests []TicketRequestParam `json:"requests" validate:"required"`
	Description *string `json:"description,omitempty"`
}

// UpdateTicketRequestParam UpdateTicketRequest request param
type UpdateTicketRequestParam struct {
	BaseParam
	Params UpdateTicketRequestParamDetail `json:"updateTicketRequest"`
}
// GetVpcIPsecLogParamDetail GetVpcIPsecLog detail param
type GetVpcIPsecLogParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Lines *int `json:"lines,omitempty"`
}

// GetVpcIPsecLogParam GetVpcIPsecLog request param
type GetVpcIPsecLogParam struct {
	BaseParam
	Params GetVpcIPsecLogParamDetail `json:"getVpcIPsecLog"`
}
// CreateVmInstanceFromOvfParamDetail CreateVmInstanceFromOvf detail param
type CreateVmInstanceFromOvfParamDetail struct {
	XmlBase64 string `json:"xmlBase64" validate:"required"`
	JsonImageInfos string `json:"jsonImageInfos" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	JsonCreateVmParam string `json:"jsonCreateVmParam" validate:"required"`
	DeleteImageAfterSuccess *bool `json:"deleteImageAfterSuccess,omitempty"`
	DeleteImageOnFail *bool `json:"deleteImageOnFail,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceFromOvfParam CreateVmInstanceFromOvf request param
type CreateVmInstanceFromOvfParam struct {
	BaseParam
	Params CreateVmInstanceFromOvfParamDetail `json:"params"`
}
// ChangeVmImageParamDetail ChangeVmImage detail param
type ChangeVmImageParamDetail struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// ChangeVmImageParam ChangeVmImage request param
type ChangeVmImageParam struct {
	BaseParam
	Params ChangeVmImageParamDetail `json:"changeVmImage"`
}
// AddResourcesToDirectoryParamDetail AddResourcesToDirectory detail param
type AddResourcesToDirectoryParamDetail struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	DirectoryUuid string `json:"directoryUuid" validate:"required"`
}

// AddResourcesToDirectoryParam AddResourcesToDirectory request param
type AddResourcesToDirectoryParam struct {
	BaseParam
	Params AddResourcesToDirectoryParamDetail `json:"params"`
}
// AttachGuestToolsIsoToVmParamDetail AttachGuestToolsIsoToVm detail param
type AttachGuestToolsIsoToVmParamDetail struct {
}

// AttachGuestToolsIsoToVmParam AttachGuestToolsIsoToVm request param
type AttachGuestToolsIsoToVmParam struct {
	BaseParam
	Params AttachGuestToolsIsoToVmParamDetail `json:"attachGuestToolsIsoToVm"`
}
// DetachAliyunKeyParamDetail DetachAliyunKey detail param
type DetachAliyunKeyParamDetail struct {
}

// DetachAliyunKeyParam DetachAliyunKey request param
type DetachAliyunKeyParam struct {
	BaseParam
	Params DetachAliyunKeyParamDetail `json:"detachAliyunKey"`
}
// DeleteBuildAppParamDetail DeleteBuildApp detail param
type DeleteBuildAppParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteBuildAppParam DeleteBuildApp request param
type DeleteBuildAppParam struct {
	BaseParam
	Params DeleteBuildAppParamDetail `json:"deleteBuildApp"`
}
// ChangeBareMetal2InstancePasswordParamDetail ChangeBareMetal2InstancePassword detail param
type ChangeBareMetal2InstancePasswordParamDetail struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ChangeBareMetal2InstancePasswordParam ChangeBareMetal2InstancePassword request param
type ChangeBareMetal2InstancePasswordParam struct {
	BaseParam
	Params ChangeBareMetal2InstancePasswordParamDetail `json:"changeBareMetal2InstancePassword"`
}
// GetResourceFromPublishAppParamDetail GetResourceFromPublishApp detail param
type GetResourceFromPublishAppParamDetail struct {
	Uuid string `json:"uuid,omitempty"`
}

// GetResourceFromPublishAppParam GetResourceFromPublishApp request param
type GetResourceFromPublishAppParam struct {
	BaseParam
	Params GetResourceFromPublishAppParamDetail `json:"getResourceFromPublishApp"`
}
// ChangeResourceOwnerParamDetail ChangeResourceOwner detail param
type ChangeResourceOwnerParamDetail struct {
	ResourceUuid string `json:"resourceUuid" validate:"required"`
}

// ChangeResourceOwnerParam ChangeResourceOwner request param
type ChangeResourceOwnerParam struct {
	BaseParam
	Params ChangeResourceOwnerParamDetail `json:"params"`
}
// GetHostIommuStateParamDetail GetHostIommuState detail param
type GetHostIommuStateParamDetail struct {
}

// GetHostIommuStateParam GetHostIommuState request param
type GetHostIommuStateParam struct {
	BaseParam
	Params GetHostIommuStateParamDetail `json:"getHostIommuState"`
}
// DeleteVirtualBorderRouterLocalParamDetail DeleteVirtualBorderRouterLocal detail param
type DeleteVirtualBorderRouterLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVirtualBorderRouterLocalParam DeleteVirtualBorderRouterLocal request param
type DeleteVirtualBorderRouterLocalParam struct {
	BaseParam
	Params DeleteVirtualBorderRouterLocalParamDetail `json:"deleteVirtualBorderRouterLocal"`
}
// GetMetricDataParamDetail GetMetricData detail param
type GetMetricDataParamDetail struct {
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	StartTime *int64 `json:"startTime,omitempty"`
	EndTime *int64 `json:"endTime,omitempty"`
	OffsetAheadOfCurrentTime *int64 `json:"offsetAheadOfCurrentTime,omitempty"`
	Period *int `json:"period,omitempty"`
	Labels []string `json:"labels,omitempty"`
	ValueConditions []string `json:"valueConditions,omitempty"`
	Functions []string `json:"functions,omitempty"`
}

// GetMetricDataParam GetMetricData request param
type GetMetricDataParam struct {
	BaseParam
	Params GetMetricDataParamDetail `json:"getMetricData"`
}
// EnableCbtTaskParamDetail EnableCbtTask detail param
type EnableCbtTaskParamDetail struct {
	BitmapName *string `json:"bitmapName,omitempty"`
}

// EnableCbtTaskParam EnableCbtTask request param
type EnableCbtTaskParam struct {
	BaseParam
	Params EnableCbtTaskParamDetail `json:"params"`
}
// GetAliyunNasAccessGroupRemoteParamDetail GetAliyunNasAccessGroupRemote detail param
type GetAliyunNasAccessGroupRemoteParamDetail struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	GroupName *string `json:"groupName,omitempty"`
}

// GetAliyunNasAccessGroupRemoteParam GetAliyunNasAccessGroupRemote request param
type GetAliyunNasAccessGroupRemoteParam struct {
	BaseParam
	Params GetAliyunNasAccessGroupRemoteParamDetail `json:"getAliyunNasAccessGroupRemote"`
}
// CheckBuildAppParametersParamDetail CheckBuildAppParameters detail param
type CheckBuildAppParametersParamDetail struct {
	Type *string `json:"type,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
}

// CheckBuildAppParametersParam CheckBuildAppParameters request param
type CheckBuildAppParametersParam struct {
	BaseParam
	Params CheckBuildAppParametersParamDetail `json:"params"`
}
// AddLabelToEventSubscriptionParamDetail AddLabelToEventSubscription detail param
type AddLabelToEventSubscriptionParamDetail struct {
	Key string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
	Operator string `json:"operator" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddLabelToEventSubscriptionParam AddLabelToEventSubscription request param
type AddLabelToEventSubscriptionParam struct {
	BaseParam
	Params AddLabelToEventSubscriptionParamDetail `json:"params"`
}
// GetVpcVRouterDistributedRoutingConnectionsParamDetail GetVpcVRouterDistributedRoutingConnections detail param
type GetVpcVRouterDistributedRoutingConnectionsParamDetail struct {
}

// GetVpcVRouterDistributedRoutingConnectionsParam GetVpcVRouterDistributedRoutingConnections request param
type GetVpcVRouterDistributedRoutingConnectionsParam struct {
	BaseParam
	Params GetVpcVRouterDistributedRoutingConnectionsParamDetail `json:"getVpcVRouterDistributedRoutingConnections"`
}
// UpdateThirdpartyAlertsParamDetail UpdateThirdpartyAlerts detail param
type UpdateThirdpartyAlertsParamDetail struct {
	Uuid string `json:"uuid,omitempty"`
	StartTimeMillis *int64 `json:"startTimeMillis,omitempty"`
	EndTimeMillis *int64 `json:"endTimeMillis,omitempty"`
	UpdateReadStatus *string `json:"updateReadStatus,omitempty"`
}

// UpdateThirdpartyAlertsParam UpdateThirdpartyAlerts request param
type UpdateThirdpartyAlertsParam struct {
	BaseParam
	Params UpdateThirdpartyAlertsParamDetail `json:"updateThirdpartyAlerts"`
}
// PullSdnControllerTenantParamDetail PullSdnControllerTenant detail param
type PullSdnControllerTenantParamDetail struct {
}

// PullSdnControllerTenantParam PullSdnControllerTenant request param
type PullSdnControllerTenantParam struct {
	BaseParam
	Params PullSdnControllerTenantParamDetail `json:"pullSdnControllerTenant"`
}
// DetachServiceFromObservabilityServerParamDetail DetachServiceFromObservabilityServer detail param
type DetachServiceFromObservabilityServerParamDetail struct {
	ServiceType string `json:"serviceType" validate:"required"`
	ServiceUuid string `json:"serviceUuid" validate:"required"`
}

// DetachServiceFromObservabilityServerParam DetachServiceFromObservabilityServer request param
type DetachServiceFromObservabilityServerParam struct {
	BaseParam
	Params DetachServiceFromObservabilityServerParamDetail `json:"detachServiceFromObservabilityServer"`
}
// GenerateHygonMdevDevicesParamDetail GenerateHygonMdevDevices detail param
type GenerateHygonMdevDevicesParamDetail struct {
	MaxQemuNum int `json:"maxQemuNum" validate:"required"`
}

// GenerateHygonMdevDevicesParam GenerateHygonMdevDevices request param
type GenerateHygonMdevDevicesParam struct {
	BaseParam
	Params GenerateHygonMdevDevicesParamDetail `json:"generateHygonMdevDevices"`
}
// SetVmUsbRedirectParamDetail SetVmUsbRedirect detail param
type SetVmUsbRedirectParamDetail struct {
	Enable bool `json:"enable" validate:"required"`
}

// SetVmUsbRedirectParam SetVmUsbRedirect request param
type SetVmUsbRedirectParam struct {
	BaseParam
	Params SetVmUsbRedirectParamDetail `json:"setVmUsbRedirect"`
}
// GetHostCandidatesForVmMigrationParamDetail GetHostCandidatesForVmMigration detail param
type GetHostCandidatesForVmMigrationParamDetail struct {
	DstPrimaryStorageUuid string `json:"dstPrimaryStorageUuid" validate:"required"`
	Limit *int `json:"limit,omitempty"`
}

// GetHostCandidatesForVmMigrationParam GetHostCandidatesForVmMigration request param
type GetHostCandidatesForVmMigrationParam struct {
	BaseParam
	Params GetHostCandidatesForVmMigrationParamDetail `json:"getHostCandidatesForVmMigration"`
}
// GetVmNicAttachableEipsParamDetail GetVmNicAttachableEips detail param
type GetVmNicAttachableEipsParamDetail struct {
	IpVersion *int `json:"ipVersion,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVmNicAttachableEipsParam GetVmNicAttachableEips request param
type GetVmNicAttachableEipsParam struct {
	BaseParam
	Params GetVmNicAttachableEipsParamDetail `json:"getVmNicAttachableEips"`
}
// UpdateFactoryModeStateParamDetail UpdateFactoryModeState detail param
type UpdateFactoryModeStateParamDetail struct {
	FactoryModeState bool `json:"factoryModeState" validate:"required"`
}

// UpdateFactoryModeStateParam UpdateFactoryModeState request param
type UpdateFactoryModeStateParam struct {
	BaseParam
	Params UpdateFactoryModeStateParamDetail `json:"updateFactoryModeState"`
}
// UpdateChronyServersParamDetail UpdateChronyServers detail param
type UpdateChronyServersParamDetail struct {
	InternalHostnames []string `json:"internalHostnames,omitempty"`
	ExternalHostnames []string `json:"externalHostnames,omitempty"`
}

// UpdateChronyServersParam UpdateChronyServers request param
type UpdateChronyServersParam struct {
	BaseParam
	Params UpdateChronyServersParamDetail `json:"updateChronyServers"`
}
// AttachPolicyRouteRuleSetToL3ParamDetail AttachPolicyRouteRuleSetToL3 detail param
type AttachPolicyRouteRuleSetToL3ParamDetail struct {
}

// AttachPolicyRouteRuleSetToL3Param AttachPolicyRouteRuleSetToL3 request param
type AttachPolicyRouteRuleSetToL3Param struct {
	BaseParam
	Params AttachPolicyRouteRuleSetToL3ParamDetail `json:"params"`
}
// UpdateOAuthClientParamDetail UpdateOAuthClient detail param
type UpdateOAuthClientParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	AuthorizationUrl *string `json:"authorizationUrl,omitempty"`
	TokenUrl *string `json:"tokenUrl,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	UserinfoUrl *string `json:"userinfoUrl,omitempty"`
	IdentityProvider *string `json:"identityProvider,omitempty"`
	PluginUuid *string `json:"pluginUuid,omitempty"`
	LoginType *string `json:"loginType,omitempty"`
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	ScopeList []string `json:"scopeList,omitempty"`
}

// UpdateOAuthClientParam UpdateOAuthClient request param
type UpdateOAuthClientParam struct {
	BaseParam
	Params UpdateOAuthClientParamDetail `json:"params"`
}
// GetZWatchAlertHistogramParamDetail GetZWatchAlertHistogram detail param
type GetZWatchAlertHistogramParamDetail struct {
	TableName string `json:"tableName" validate:"required"`
	StartTime int64 `json:"startTime" validate:"required"`
	EndTime int64 `json:"endTime" validate:"required"`
	IntervalHours int `json:"intervalHours" validate:"required"`
	GroupColumns []string `json:"groupColumns,omitempty"`
}

// GetZWatchAlertHistogramParam GetZWatchAlertHistogram request param
type GetZWatchAlertHistogramParam struct {
	BaseParam
	Params GetZWatchAlertHistogramParamDetail `json:"getZWatchAlertHistogram"`
}
// DeleteAliyunRouterInterfaceRemoteParamDetail DeleteAliyunRouterInterfaceRemote detail param
type DeleteAliyunRouterInterfaceRemoteParamDetail struct {
	VRouterType string `json:"vRouterType" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAliyunRouterInterfaceRemoteParam DeleteAliyunRouterInterfaceRemote request param
type DeleteAliyunRouterInterfaceRemoteParam struct {
	BaseParam
	Params DeleteAliyunRouterInterfaceRemoteParamDetail `json:"deleteAliyunRouterInterfaceRemote"`
}
// SetImageBootModeParamDetail SetImageBootMode detail param
type SetImageBootModeParamDetail struct {
	BootMode string `json:"bootMode" validate:"required"`
}

// SetImageBootModeParam SetImageBootMode request param
type SetImageBootModeParam struct {
	BaseParam
	Params SetImageBootModeParamDetail `json:"setImageBootMode"`
}
// DetachAutoScalingTemplateFromGroupParamDetail DetachAutoScalingTemplateFromGroup detail param
type DetachAutoScalingTemplateFromGroupParamDetail struct {
}

// DetachAutoScalingTemplateFromGroupParam DetachAutoScalingTemplateFromGroup request param
type DetachAutoScalingTemplateFromGroupParam struct {
	BaseParam
	Params DetachAutoScalingTemplateFromGroupParamDetail `json:"detachAutoScalingTemplateFromGroup"`
}
// UpdateVirtualBorderRouterRemoteParamDetail UpdateVirtualBorderRouterRemote detail param
type UpdateVirtualBorderRouterRemoteParamDetail struct {
	LocalGatewayIp *string `json:"localGatewayIp,omitempty"`
	PeerGatewayIp *string `json:"peerGatewayIp,omitempty"`
	PeeringSubnetMask *string `json:"peeringSubnetMask,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	VlanId *string `json:"vlanId,omitempty"`
	CircuitCode *string `json:"circuitCode,omitempty"`
}

// UpdateVirtualBorderRouterRemoteParam UpdateVirtualBorderRouterRemote request param
type UpdateVirtualBorderRouterRemoteParam struct {
	BaseParam
	Params UpdateVirtualBorderRouterRemoteParamDetail `json:"updateVirtualBorderRouterRemote"`
}
// GetVmsCapabilitiesParamDetail GetVmsCapabilities detail param
type GetVmsCapabilitiesParamDetail struct {
	VmUuids []string `json:"vmUuids" validate:"required"`
}

// GetVmsCapabilitiesParam GetVmsCapabilities request param
type GetVmsCapabilitiesParam struct {
	BaseParam
	Params GetVmsCapabilitiesParamDetail `json:"getVmsCapabilities"`
}
// AttachPolicyToUserGroupParamDetail AttachPolicyToUserGroup detail param
type AttachPolicyToUserGroupParamDetail struct {
	PolicyUuid string `json:"policyUuid" validate:"required"`
}

// AttachPolicyToUserGroupParam AttachPolicyToUserGroup request param
type AttachPolicyToUserGroupParam struct {
	BaseParam
	Params AttachPolicyToUserGroupParamDetail `json:"params"`
}
// RevokeMonitorTemplateFromMonitorGroupParamDetail RevokeMonitorTemplateFromMonitorGroup detail param
type RevokeMonitorTemplateFromMonitorGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RevokeMonitorTemplateFromMonitorGroupParam RevokeMonitorTemplateFromMonitorGroup request param
type RevokeMonitorTemplateFromMonitorGroupParam struct {
	BaseParam
	Params RevokeMonitorTemplateFromMonitorGroupParamDetail `json:"revokeMonitorTemplateFromMonitorGroup"`
}
// DeleteFirewallRuleParamDetail DeleteFirewallRule detail param
type DeleteFirewallRuleParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteFirewallRuleParam DeleteFirewallRule request param
type DeleteFirewallRuleParam struct {
	BaseParam
	Params DeleteFirewallRuleParamDetail `json:"deleteFirewallRule"`
}
// ShareResourceParamDetail ShareResource detail param
type ShareResourceParamDetail struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	AccountUuids []string `json:"accountUuids,omitempty"`
	ToPublic bool `json:"toPublic,omitempty"`
}

// ShareResourceParam ShareResource request param
type ShareResourceParam struct {
	BaseParam
	Params ShareResourceParamDetail `json:"shareResource"`
}
// CreateEcsVpcRemoteParamDetail CreateEcsVpcRemote detail param
type CreateEcsVpcRemoteParamDetail struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	CidrBlock string `json:"cidrBlock" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	VRouterName string `json:"vRouterName" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsVpcRemoteParam CreateEcsVpcRemote request param
type CreateEcsVpcRemoteParam struct {
	BaseParam
	Params CreateEcsVpcRemoteParamDetail `json:"params"`
}
// GetAccountQuotaUsageParamDetail GetAccountQuotaUsage detail param
type GetAccountQuotaUsageParamDetail struct {
}

// GetAccountQuotaUsageParam GetAccountQuotaUsage request param
type GetAccountQuotaUsageParam struct {
	BaseParam
	Params GetAccountQuotaUsageParamDetail `json:"getAccountQuotaUsage"`
}
// RemoveIAM2VirtualIDsFromProjectsParamDetail RemoveIAM2VirtualIDsFromProjects detail param
type RemoveIAM2VirtualIDsFromProjectsParamDetail struct {
	ProjectUuids []string `json:"projectUuids" validate:"required"`
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
}

// RemoveIAM2VirtualIDsFromProjectsParam RemoveIAM2VirtualIDsFromProjects request param
type RemoveIAM2VirtualIDsFromProjectsParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromProjectsParamDetail `json:"removeIAM2VirtualIDsFromProjects"`
}
// GetCandidateL3NetworksForServerGroupParamDetail GetCandidateL3NetworksForServerGroup detail param
type GetCandidateL3NetworksForServerGroupParamDetail struct {
	ServerGroupUuid *string `json:"serverGroupUuid,omitempty"`
	LoadBalancerUuid *string `json:"loadBalancerUuid,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetCandidateL3NetworksForServerGroupParam GetCandidateL3NetworksForServerGroup request param
type GetCandidateL3NetworksForServerGroupParam struct {
	BaseParam
	Params GetCandidateL3NetworksForServerGroupParamDetail `json:"getCandidateL3NetworksForServerGroup"`
}
// CreateVmFromCdpBackupParamDetail CreateVmFromCdpBackup detail param
type CreateVmFromCdpBackupParamDetail struct {
	Name string `json:"name" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
	CdpTaskUuid string `json:"cdpTaskUuid" validate:"required"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid" validate:"required"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume *string `json:"primaryStorageUuidForDataVolume,omitempty"`
	RecoverBandwidth *int64 `json:"recoverBandwidth,omitempty"`
	Description *string `json:"description,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmFromCdpBackupParam CreateVmFromCdpBackup request param
type CreateVmFromCdpBackupParam struct {
	BaseParam
	Params CreateVmFromCdpBackupParamDetail `json:"createVmFromCdpBackup"`
}
// SyncVpcVpnConnectionFromRemoteParamDetail SyncVpcVpnConnectionFromRemote detail param
type SyncVpcVpnConnectionFromRemoteParamDetail struct {
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncVpcVpnConnectionFromRemoteParam SyncVpcVpnConnectionFromRemote request param
type SyncVpcVpnConnectionFromRemoteParam struct {
	BaseParam
	Params SyncVpcVpnConnectionFromRemoteParamDetail `json:"syncVpcVpnConnectionFromRemote"`
}
// CreateVpnIkeConfigParamDetail CreateVpnIkeConfig detail param
type CreateVpnIkeConfigParamDetail struct {
	Name string `json:"name" validate:"required"`
	Psk string `json:"psk" validate:"required"`
	Pfs *string `json:"pfs,omitempty"`
	Version *string `json:"version,omitempty"`
	Mode *string `json:"mode,omitempty"`
	EncAlg *string `json:"encAlg,omitempty"`
	AuthAlg *string `json:"authAlg,omitempty"`
	Lifetime *int `json:"lifetime,omitempty"`
	LocalIp string `json:"localIp" validate:"required"`
	RemoteIp string `json:"remoteIp" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpnIkeConfigParam CreateVpnIkeConfig request param
type CreateVpnIkeConfigParam struct {
	BaseParam
	Params CreateVpnIkeConfigParamDetail `json:"params"`
}
// SubmitLongJobParamDetail SubmitLongJob detail param
type SubmitLongJobParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	JobName string `json:"jobName" validate:"required"`
	JobData string `json:"jobData" validate:"required"`
	TargetResourceUuid *string `json:"targetResourceUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SubmitLongJobParam SubmitLongJob request param
type SubmitLongJobParam struct {
	BaseParam
	Params SubmitLongJobParamDetail `json:"params"`
}
// CreateDataVolumeTemplateFromVolumeBackupParamDetail CreateDataVolumeTemplateFromVolumeBackup detail param
type CreateDataVolumeTemplateFromVolumeBackupParamDetail struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	System bool `json:"system,omitempty"`
	Virtio *bool `json:"virtio,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeTemplateFromVolumeBackupParam CreateDataVolumeTemplateFromVolumeBackup request param
type CreateDataVolumeTemplateFromVolumeBackupParam struct {
	BaseParam
	Params CreateDataVolumeTemplateFromVolumeBackupParamDetail `json:"params"`
}
// DegradeFromLicenseServerParamDetail DegradeFromLicenseServer detail param
type DegradeFromLicenseServerParamDetail struct {
}

// DegradeFromLicenseServerParam DegradeFromLicenseServer request param
type DegradeFromLicenseServerParam struct {
	BaseParam
	Params DegradeFromLicenseServerParamDetail `json:"degradeFromLicenseServer"`
}
// UpdateNfvInstProvisionConfigParamDetail UpdateNfvInstProvisionConfig detail param
type UpdateNfvInstProvisionConfigParamDetail struct {
}

// UpdateNfvInstProvisionConfigParam UpdateNfvInstProvisionConfig request param
type UpdateNfvInstProvisionConfigParam struct {
	BaseParam
	Params UpdateNfvInstProvisionConfigParamDetail `json:"updateNfvInstProvisionConfig"`
}
// GetDebugSignalParamDetail GetDebugSignal detail param
type GetDebugSignalParamDetail struct {
}

// GetDebugSignalParam GetDebugSignal request param
type GetDebugSignalParam struct {
	BaseParam
	Params GetDebugSignalParamDetail `json:"getDebugSignal"`
}
// UpdateAliyunKeySecretParamDetail UpdateAliyunKeySecret detail param
type UpdateAliyunKeySecretParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateAliyunKeySecretParam UpdateAliyunKeySecret request param
type UpdateAliyunKeySecretParam struct {
	BaseParam
	Params UpdateAliyunKeySecretParamDetail `json:"params"`
}
// SyncLicenseCapacityParamDetail SyncLicenseCapacity detail param
type SyncLicenseCapacityParamDetail struct {
}

// SyncLicenseCapacityParam SyncLicenseCapacity request param
type SyncLicenseCapacityParam struct {
	BaseParam
	Params SyncLicenseCapacityParamDetail `json:"syncLicenseCapacity"`
}
// AttachDataVolumeToHostParamDetail AttachDataVolumeToHost detail param
type AttachDataVolumeToHostParamDetail struct {
	MountPath string `json:"mountPath" validate:"required"`
}

// AttachDataVolumeToHostParam AttachDataVolumeToHost request param
type AttachDataVolumeToHostParam struct {
	BaseParam
	Params AttachDataVolumeToHostParamDetail `json:"params"`
}
// SecurityMachineEncryptParamDetail SecurityMachineEncrypt detail param
type SecurityMachineEncryptParamDetail struct {
	Text string `json:"text" validate:"required"`
	AlgType string `json:"algType" validate:"required"`
}

// SecurityMachineEncryptParam SecurityMachineEncrypt request param
type SecurityMachineEncryptParam struct {
	BaseParam
	Params SecurityMachineEncryptParamDetail `json:"params"`
}
// ChangeAppBuildSystemStateParamDetail ChangeAppBuildSystemState detail param
type ChangeAppBuildSystemStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeAppBuildSystemStateParam ChangeAppBuildSystemState request param
type ChangeAppBuildSystemStateParam struct {
	BaseParam
	Params ChangeAppBuildSystemStateParamDetail `json:"changeAppBuildSystemState"`
}
// GetMemorySnapshotGroupReferenceParamDetail GetMemorySnapshotGroupReference detail param
type GetMemorySnapshotGroupReferenceParamDetail struct {
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	ResourceType string `json:"resourceType" validate:"required"`
}

// GetMemorySnapshotGroupReferenceParam GetMemorySnapshotGroupReference request param
type GetMemorySnapshotGroupReferenceParam struct {
	BaseParam
	Params GetMemorySnapshotGroupReferenceParamDetail `json:"getMemorySnapshotGroupReference"`
}
// GetVpcVRouterNetworkServiceStateParamDetail GetVpcVRouterNetworkServiceState detail param
type GetVpcVRouterNetworkServiceStateParamDetail struct {
	NetworkService string `json:"networkService" validate:"required"`
	L3NetworkUuid *string `json:"l3NetworkUuid,omitempty"`
}

// GetVpcVRouterNetworkServiceStateParam GetVpcVRouterNetworkServiceState request param
type GetVpcVRouterNetworkServiceStateParam struct {
	BaseParam
	Params GetVpcVRouterNetworkServiceStateParamDetail `json:"getVpcVRouterNetworkServiceState"`
}
// DetachNetworkServiceFromL3NetworkParamDetail DetachNetworkServiceFromL3Network detail param
type DetachNetworkServiceFromL3NetworkParamDetail struct {
	NetworkServices map[string]interface{} `json:"networkServices,omitempty"`
	Service *string `json:"service,omitempty"`
}

// DetachNetworkServiceFromL3NetworkParam DetachNetworkServiceFromL3Network request param
type DetachNetworkServiceFromL3NetworkParam struct {
	BaseParam
	Params DetachNetworkServiceFromL3NetworkParamDetail `json:"detachNetworkServiceFromL3Network"`
}
// CreateDataVolumeFromVolumeBackupParamDetail CreateDataVolumeFromVolumeBackup detail param
type CreateDataVolumeFromVolumeBackupParamDetail struct {
	Name string `json:"name" validate:"required"`
	VmInstanceUuid *string `json:"vmInstanceUuid,omitempty"`
	BackupStorageUuid *string `json:"backupStorageUuid,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeFromVolumeBackupParam CreateDataVolumeFromVolumeBackup request param
type CreateDataVolumeFromVolumeBackupParam struct {
	BaseParam
	Params CreateDataVolumeFromVolumeBackupParamDetail `json:"params"`
}
// DeleteContainerResourceFromEndpointParamDetail DeleteContainerResourceFromEndpoint detail param
type DeleteContainerResourceFromEndpointParamDetail struct {
}

// DeleteContainerResourceFromEndpointParam DeleteContainerResourceFromEndpoint request param
type DeleteContainerResourceFromEndpointParam struct {
	BaseParam
	Params DeleteContainerResourceFromEndpointParamDetail `json:"deleteContainerResourceFromEndpoint"`
}
// GetSupportAPIsParamDetail GetSupports detail param
type GetSupportAPIsParamDetail struct {
}

// GetSupportAPIsParam GetSupports request param
type GetSupportAPIsParam struct {
	BaseParam
	Params GetSupportAPIsParamDetail `json:"getSupportAPIs"`
}
// AddSharedMountPointPrimaryStorageParamDetail AddSharedMountPointPrimaryStorage detail param
type AddSharedMountPointPrimaryStorageParamDetail struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSharedMountPointPrimaryStorageParam AddSharedMountPointPrimaryStorage request param
type AddSharedMountPointPrimaryStorageParam struct {
	BaseParam
	Params AddSharedMountPointPrimaryStorageParamDetail `json:"params"`
}
// GetTrashOnPrimaryStorageParamDetail GetTrashOnPrimaryStorage detail param
type GetTrashOnPrimaryStorageParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	TrashType *string `json:"trashType,omitempty"`
}

// GetTrashOnPrimaryStorageParam GetTrashOnPrimaryStorage request param
type GetTrashOnPrimaryStorageParam struct {
	BaseParam
	Params GetTrashOnPrimaryStorageParamDetail `json:"getTrashOnPrimaryStorage"`
}
// GetCandidateVmNicsForLoadBalancerParamDetail GetCandidateVmNicsForLoadBalancer detail param
type GetCandidateVmNicsForLoadBalancerParamDetail struct {
}

// GetCandidateVmNicsForLoadBalancerParam GetCandidateVmNicsForLoadBalancer request param
type GetCandidateVmNicsForLoadBalancerParam struct {
	BaseParam
	Params GetCandidateVmNicsForLoadBalancerParamDetail `json:"getCandidateVmNicsForLoadBalancer"`
}
// DetachBaremetalPxeServerFromClusterParamDetail DetachBaremetalPxeServerFromCluster detail param
type DetachBaremetalPxeServerFromClusterParamDetail struct {
}

// DetachBaremetalPxeServerFromClusterParam DetachBaremetalPxeServerFromCluster request param
type DetachBaremetalPxeServerFromClusterParam struct {
	BaseParam
	Params DetachBaremetalPxeServerFromClusterParamDetail `json:"detachBaremetalPxeServerFromCluster"`
}
// DeleteVpcUserVpnGatewayRemoteParamDetail DeleteVpcUserVpnGatewayRemote detail param
type DeleteVpcUserVpnGatewayRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVpcUserVpnGatewayRemoteParam DeleteVpcUserVpnGatewayRemote request param
type DeleteVpcUserVpnGatewayRemoteParam struct {
	BaseParam
	Params DeleteVpcUserVpnGatewayRemoteParamDetail `json:"deleteVpcUserVpnGatewayRemote"`
}
// ChangeAccessControlListRedirectRuleParamDetail ChangeAccessControlListRedirectRule detail param
type ChangeAccessControlListRedirectRuleParamDetail struct {
	Name string `json:"name,omitempty"`
}

// ChangeAccessControlListRedirectRuleParam ChangeAccessControlListRedirectRule request param
type ChangeAccessControlListRedirectRuleParam struct {
	BaseParam
	Params ChangeAccessControlListRedirectRuleParamDetail `json:"changeAccessControlListRedirectRule"`
}
// AddResourceStackVmPortMonitorParamDetail AddResourceStackVmPortMonitor detail param
type AddResourceStackVmPortMonitorParamDetail struct {
	StackUuid *string `json:"stackUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Port int `json:"port" validate:"required"`
}

// AddResourceStackVmPortMonitorParam AddResourceStackVmPortMonitor request param
type AddResourceStackVmPortMonitorParam struct {
	BaseParam
	Params AddResourceStackVmPortMonitorParamDetail `json:"params"`
}
// ChangeSNSApplicationEndpointStateParamDetail ChangeSNSApplicationEndpointState detail param
type ChangeSNSApplicationEndpointStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSNSApplicationEndpointStateParam ChangeSNSApplicationEndpointState request param
type ChangeSNSApplicationEndpointStateParam struct {
	BaseParam
	Params ChangeSNSApplicationEndpointStateParamDetail `json:"changeSNSApplicationEndpointState"`
}
// GetVpcAttachedLoadBalancerParamDetail GetVpcAttachedLoadBalancer detail param
type GetVpcAttachedLoadBalancerParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVpcAttachedLoadBalancerParam GetVpcAttachedLoadBalancer request param
type GetVpcAttachedLoadBalancerParam struct {
	BaseParam
	Params GetVpcAttachedLoadBalancerParamDetail `json:"params"`
}
// CreateVpcVpnConnectionRemoteParamDetail CreateVpcVpnConnectionRemote detail param
type CreateVpcVpnConnectionRemoteParamDetail struct {
	UserGatewayUuid string `json:"userGatewayUuid" validate:"required"`
	VpnGatewayUuid string `json:"vpnGatewayUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	LocalCidr string `json:"localCidr" validate:"required"`
	RemoteCidr string `json:"remoteCidr" validate:"required"`
	Active bool `json:"active" validate:"required"`
	IkeConfUuid string `json:"ikeConfUuid" validate:"required"`
	IpsecConfUuid string `json:"ipsecConfUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcVpnConnectionRemoteParam CreateVpcVpnConnectionRemote request param
type CreateVpcVpnConnectionRemoteParam struct {
	BaseParam
	Params CreateVpcVpnConnectionRemoteParamDetail `json:"params"`
}
// GetVpcAttachedPortForwardingRulesParamDetail GetVpcAttachedPortForwardingRules detail param
type GetVpcAttachedPortForwardingRulesParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVpcAttachedPortForwardingRulesParam GetVpcAttachedPortForwardingRules request param
type GetVpcAttachedPortForwardingRulesParam struct {
	BaseParam
	Params GetVpcAttachedPortForwardingRulesParamDetail `json:"params"`
}
// SetVpcVRouterNetworkServiceStateParamDetail SetVpcVRouterNetworkServiceState detail param
type SetVpcVRouterNetworkServiceStateParamDetail struct {
	NetworkService string `json:"networkService" validate:"required"`
	State string `json:"state" validate:"required"`
	L3NetworkUuid *string `json:"l3NetworkUuid,omitempty"`
}

// SetVpcVRouterNetworkServiceStateParam SetVpcVRouterNetworkServiceState request param
type SetVpcVRouterNetworkServiceStateParam struct {
	BaseParam
	Params SetVpcVRouterNetworkServiceStateParamDetail `json:"params"`
}
// DetachNfvInstFromGroupParamDetail DetachNfvInstFromGroup detail param
type DetachNfvInstFromGroupParamDetail struct {
}

// DetachNfvInstFromGroupParam DetachNfvInstFromGroup request param
type DetachNfvInstFromGroupParam struct {
	BaseParam
	Params DetachNfvInstFromGroupParamDetail `json:"detachNfvInstFromGroup"`
}
// AddDnsToVpcRouterParamDetail AddDnsToVpcRouter detail param
type AddDnsToVpcRouterParamDetail struct {
	Dns string `json:"dns" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddDnsToVpcRouterParam AddDnsToVpcRouter request param
type AddDnsToVpcRouterParam struct {
	BaseParam
	Params AddDnsToVpcRouterParamDetail `json:"params"`
}
// GetVmXmlParamDetail GetVmXml detail param
type GetVmXmlParamDetail struct {
}

// GetVmXmlParam GetVmXml request param
type GetVmXmlParam struct {
	BaseParam
	Params GetVmXmlParamDetail `json:"getVmXml"`
}
// GetVmInstanceFirstBootDeviceParamDetail GetVmInstanceFirstBootDevice detail param
type GetVmInstanceFirstBootDeviceParamDetail struct {
}

// GetVmInstanceFirstBootDeviceParam GetVmInstanceFirstBootDevice request param
type GetVmInstanceFirstBootDeviceParam struct {
	BaseParam
	Params GetVmInstanceFirstBootDeviceParamDetail `json:"getVmInstanceFirstBootDevice"`
}
// CreateOvnControllerVmParamDetail CreateOvnControllerVm detail param
type CreateOvnControllerVmParamDetail struct {
	Name string `json:"name" validate:"required"`
	InstanceOfferingUuid *string `json:"instanceOfferingUuid,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	ReservedMemorySize *int64 `json:"reservedMemorySize,omitempty"`
	ImageUuid *string `json:"imageUuid,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	VmNicParams *string `json:"vmNicParams,omitempty"`
	Type *string `json:"type,omitempty"`
	RootDiskOfferingUuid *string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize *int64 `json:"rootDiskSize,omitempty"`
	DataDiskSizes []int64 `json:"dataDiskSizes,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	Description *string `json:"description,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	DataVolumeSystemTagsOnIndex map[string]interface{} `json:"dataVolumeSystemTagsOnIndex,omitempty"`
	SshKeyPairUuids []string `json:"sshKeyPairUuids,omitempty"`
	Platform *string `json:"platform,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	Virtio *bool `json:"virtio,omitempty"`
	DiskAOs []CreateVmInstance_DiskAOParam `json:"diskAOs,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateOvnControllerVmParam CreateOvnControllerVm request param
type CreateOvnControllerVmParam struct {
	BaseParam
	Params CreateOvnControllerVmParamDetail `json:"params"`
}
// DeleteIpAddressParamDetail DeleteIpAddress detail param
type DeleteIpAddressParamDetail struct {
	UsedIpUuids []string `json:"usedIpUuids" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteIpAddressParam DeleteIpAddress request param
type DeleteIpAddressParam struct {
	BaseParam
	Params DeleteIpAddressParamDetail `json:"deleteIpAddress"`
}
// DeleteVpcVpnConnectionRemoteParamDetail DeleteVpcVpnConnectionRemote detail param
type DeleteVpcVpnConnectionRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVpcVpnConnectionRemoteParam DeleteVpcVpnConnectionRemote request param
type DeleteVpcVpnConnectionRemoteParam struct {
	BaseParam
	Params DeleteVpcVpnConnectionRemoteParamDetail `json:"deleteVpcVpnConnectionRemote"`
}
// AttachOssBucketToEcsDataCenterParamDetail AttachOssBucketToEcsDataCenter detail param
type AttachOssBucketToEcsDataCenterParamDetail struct {
}

// AttachOssBucketToEcsDataCenterParam AttachOssBucketToEcsDataCenter request param
type AttachOssBucketToEcsDataCenterParam struct {
	BaseParam
	Params AttachOssBucketToEcsDataCenterParamDetail `json:"attachOssBucketToEcsDataCenter"`
}
// CheckIAM2OrganizationAvailabilityParamDetail CheckIAM2OrganizationAvailability detail param
type CheckIAM2OrganizationAvailabilityParamDetail struct {
}

// CheckIAM2OrganizationAvailabilityParam CheckIAM2OrganizationAvailability request param
type CheckIAM2OrganizationAvailabilityParam struct {
	BaseParam
	Params CheckIAM2OrganizationAvailabilityParamDetail `json:"checkIAM2OrganizationAvailability"`
}
// UnmountVmInstanceRecoveryPointParamDetail UnmountVmInstanceRecoveryPoint detail param
type UnmountVmInstanceRecoveryPointParamDetail struct {
	VmUuid string `json:"vmUuid" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
}

// UnmountVmInstanceRecoveryPointParam UnmountVmInstanceRecoveryPoint request param
type UnmountVmInstanceRecoveryPointParam struct {
	BaseParam
	Params UnmountVmInstanceRecoveryPointParamDetail `json:"params"`
}
// RemovePolicyStatementsFromRoleParamDetail RemovePolicyStatementsFromRole detail param
type RemovePolicyStatementsFromRoleParamDetail struct {
	PolicyStatementUuids []string `json:"policyStatementUuids" validate:"required"`
}

// RemovePolicyStatementsFromRoleParam RemovePolicyStatementsFromRole request param
type RemovePolicyStatementsFromRoleParam struct {
	BaseParam
	Params RemovePolicyStatementsFromRoleParamDetail `json:"removePolicyStatementsFromRole"`
}
// GenerateModelMetadataParamDetail GenerateModelMetadata detail param
type GenerateModelMetadataParamDetail struct {
	ModelCenterUuid string `json:"modelCenterUuid" validate:"required"`
	ModelUuids []string `json:"modelUuids,omitempty"`
}

// GenerateModelMetadataParam GenerateModelMetadata request param
type GenerateModelMetadataParam struct {
	BaseParam
	Params GenerateModelMetadataParamDetail `json:"generateModelMetadata"`
}
// IsReadyToGoParamDetail IsReadyToGo detail param
type IsReadyToGoParamDetail struct {
	ManagementNodeId *string `json:"managementNodeId,omitempty"`
}

// IsReadyToGoParam IsReadyToGo request param
type IsReadyToGoParam struct {
	BaseParam
	Params IsReadyToGoParamDetail `json:"isReadyToGo"`
}
// GetHostIommuStatusParamDetail GetHostIommuStatus detail param
type GetHostIommuStatusParamDetail struct {
}

// GetHostIommuStatusParam GetHostIommuStatus request param
type GetHostIommuStatusParam struct {
	BaseParam
	Params GetHostIommuStatusParamDetail `json:"getHostIommuStatus"`
}
// DescribeVmInstanceRecoveryPointParamDetail DescribeVmInstanceRecoveryPoint detail param
type DescribeVmInstanceRecoveryPointParamDetail struct {
	GroupId int64 `json:"groupId" validate:"required"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// DescribeVmInstanceRecoveryPointParam DescribeVmInstanceRecoveryPoint request param
type DescribeVmInstanceRecoveryPointParam struct {
	BaseParam
	Params DescribeVmInstanceRecoveryPointParamDetail `json:"describeVmInstanceRecoveryPoint"`
}
// GetPciDeviceCandidatesForAttachingVmParamDetail GetPciDeviceCandidatesForAttachingVm detail param
type GetPciDeviceCandidatesForAttachingVmParamDetail struct {
	Types []string `json:"types,omitempty"`
	PciSpecUuids []string `json:"pciSpecUuids,omitempty"`
}

// GetPciDeviceCandidatesForAttachingVmParam GetPciDeviceCandidatesForAttachingVm request param
type GetPciDeviceCandidatesForAttachingVmParam struct {
	BaseParam
	Params GetPciDeviceCandidatesForAttachingVmParamDetail `json:"getPciDeviceCandidatesForAttachingVm"`
}
// ChangeMonitorTriggerStateParamDetail ChangeMonitorTriggerState detail param
type ChangeMonitorTriggerStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeMonitorTriggerStateParam ChangeMonitorTriggerState request param
type ChangeMonitorTriggerStateParam struct {
	BaseParam
	Params ChangeMonitorTriggerStateParamDetail `json:"changeMonitorTriggerState"`
}
// GetBareMetal2ChassisPowerStatusParamDetail GetBareMetal2ChassisPowerStatus detail param
type GetBareMetal2ChassisPowerStatusParamDetail struct {
}

// GetBareMetal2ChassisPowerStatusParam GetBareMetal2ChassisPowerStatus request param
type GetBareMetal2ChassisPowerStatusParam struct {
	BaseParam
	Params GetBareMetal2ChassisPowerStatusParamDetail `json:"getBareMetal2ChassisPowerStatus"`
}
// GetTaskProgressParamDetail GetTaskProgress detail param
type GetTaskProgressParamDetail struct {
	All bool `json:"all,omitempty"`
}

// GetTaskProgressParam GetTaskProgress request param
type GetTaskProgressParam struct {
	BaseParam
	Params GetTaskProgressParamDetail `json:"getTaskProgress"`
}
// StartDataProtectionParamDetail StartDataProtection detail param
type StartDataProtectionParamDetail struct {
	EncryptType string `json:"encryptType" validate:"required"`
	AuditsIntegrityDate *int `json:"auditsIntegrityDate,omitempty"`
}

// StartDataProtectionParam StartDataProtection request param
type StartDataProtectionParam struct {
	BaseParam
	Params StartDataProtectionParamDetail `json:"params"`
}
// ChangeActiveAlarmStateParamDetail ChangeActiveAlarmState detail param
type ChangeActiveAlarmStateParamDetail struct {
	Namespace string `json:"namespace" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeActiveAlarmStateParam ChangeActiveAlarmState request param
type ChangeActiveAlarmStateParam struct {
	BaseParam
	Params ChangeActiveAlarmStateParamDetail `json:"params"`
}
// SetVmCleanTrafficParamDetail SetVmCleanTraffic detail param
type SetVmCleanTrafficParamDetail struct {
	Enable bool `json:"enable" validate:"required"`
}

// SetVmCleanTrafficParam SetVmCleanTraffic request param
type SetVmCleanTrafficParam struct {
	BaseParam
	Params SetVmCleanTrafficParamDetail `json:"setVmCleanTraffic"`
}
// SetVmBootModeParamDetail SetVmBootMode detail param
type SetVmBootModeParamDetail struct {
	BootMode string `json:"bootMode" validate:"required"`
}

// SetVmBootModeParam SetVmBootMode request param
type SetVmBootModeParam struct {
	BaseParam
	Params SetVmBootModeParamDetail `json:"setVmBootMode"`
}
// SyncImageSizeParamDetail SyncImageSize detail param
type SyncImageSizeParamDetail struct {
}

// SyncImageSizeParam SyncImageSize request param
type SyncImageSizeParam struct {
	BaseParam
	Params SyncImageSizeParamDetail `json:"syncImageSize"`
}
// GetNoTriggerSchedulerJobsParamDetail GetNoTriggerSchedulerJobs detail param
type GetNoTriggerSchedulerJobsParamDetail struct {
}

// GetNoTriggerSchedulerJobsParam GetNoTriggerSchedulerJobs request param
type GetNoTriggerSchedulerJobsParam struct {
	BaseParam
	Params GetNoTriggerSchedulerJobsParamDetail `json:"getNoTriggerSchedulerJobs"`
}
// AddProxyToResourceParamDetail AddProxyToResource detail param
type AddProxyToResourceParamDetail struct {
}

// AddProxyToResourceParam AddProxyToResource request param
type AddProxyToResourceParam struct {
	BaseParam
	Params AddProxyToResourceParamDetail `json:"params"`
}
// ProtectVmInstanceRecoveryPointParamDetail ProtectVmInstanceRecoveryPoint detail param
type ProtectVmInstanceRecoveryPointParamDetail struct {
	GroupId int64 `json:"groupId" validate:"required"`
	Description *string `json:"description,omitempty"`
}

// ProtectVmInstanceRecoveryPointParam ProtectVmInstanceRecoveryPoint request param
type ProtectVmInstanceRecoveryPointParam struct {
	BaseParam
	Params ProtectVmInstanceRecoveryPointParamDetail `json:"protectVmInstanceRecoveryPoint"`
}
// DeleteConnectionAccessPointLocalParamDetail DeleteConnectionAccessPointLocal detail param
type DeleteConnectionAccessPointLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteConnectionAccessPointLocalParam DeleteConnectionAccessPointLocal request param
type DeleteConnectionAccessPointLocalParam struct {
	BaseParam
	Params DeleteConnectionAccessPointLocalParamDetail `json:"deleteConnectionAccessPointLocal"`
}
// RemoveIAM2VirtualIDsFromProjectParamDetail RemoveIAM2VirtualIDsFromProject detail param
type RemoveIAM2VirtualIDsFromProjectParamDetail struct {
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
}

// RemoveIAM2VirtualIDsFromProjectParam RemoveIAM2VirtualIDsFromProject request param
type RemoveIAM2VirtualIDsFromProjectParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromProjectParamDetail `json:"removeIAM2VirtualIDsFromProject"`
}
// CreateEcsImageFromEcsSnapshotParamDetail CreateEcsImageFromEcsSnapshot detail param
type CreateEcsImageFromEcsSnapshotParamDetail struct {
	SnapshotUuid string `json:"snapshotUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsImageFromEcsSnapshotParam CreateEcsImageFromEcsSnapshot request param
type CreateEcsImageFromEcsSnapshotParam struct {
	BaseParam
	Params CreateEcsImageFromEcsSnapshotParamDetail `json:"params"`
}
// CreateResourceStackFromAppParamDetail CreateResourceStackFromApp detail param
type CreateResourceStackFromAppParamDetail struct {
	AppUuid string `json:"appUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Rollback *bool `json:"rollback,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
	WithoutAppInfo *bool `json:"withoutAppInfo,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateResourceStackFromAppParam CreateResourceStackFromApp request param
type CreateResourceStackFromAppParam struct {
	BaseParam
	Params CreateResourceStackFromAppParamDetail `json:"params"`
}
// GetSharedBlockCandidateParamDetail GetSharedBlockCandidate detail param
type GetSharedBlockCandidateParamDetail struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// GetSharedBlockCandidateParam GetSharedBlockCandidate request param
type GetSharedBlockCandidateParam struct {
	BaseParam
	Params GetSharedBlockCandidateParamDetail `json:"getSharedBlockCandidate"`
}
// SyncEcsSecurityGroupFromRemoteParamDetail SyncEcsSecurityGroupFromRemote detail param
type SyncEcsSecurityGroupFromRemoteParamDetail struct {
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsSecurityGroupFromRemoteParam SyncEcsSecurityGroupFromRemote request param
type SyncEcsSecurityGroupFromRemoteParam struct {
	BaseParam
	Params SyncEcsSecurityGroupFromRemoteParamDetail `json:"syncEcsSecurityGroupFromRemote"`
}
// ExportBuildAppParamDetail ExportBuildApp detail param
type ExportBuildAppParamDetail struct {
}

// ExportBuildAppParam ExportBuildApp request param
type ExportBuildAppParam struct {
	BaseParam
	Params ExportBuildAppParamDetail `json:"exportBuildApp"`
}
// ReclaimSpaceFromImageStoreParamDetail ReclaimSpaceFromImageStore detail param
type ReclaimSpaceFromImageStoreParamDetail struct {
}

// ReclaimSpaceFromImageStoreParam ReclaimSpaceFromImageStore request param
type ReclaimSpaceFromImageStoreParam struct {
	BaseParam
	Params ReclaimSpaceFromImageStoreParamDetail `json:"reclaimSpaceFromImageStore"`
}
// GetAllEventMetadataParamDetail GetAllEventMetadata detail param
type GetAllEventMetadataParamDetail struct {
	Name string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// GetAllEventMetadataParam GetAllEventMetadata request param
type GetAllEventMetadataParam struct {
	BaseParam
	Params GetAllEventMetadataParamDetail `json:"getAllEventMetadata"`
}
// GetCandidateVmForAttachingIsoParamDetail GetCandidateVmForAttachingIso detail param
type GetCandidateVmForAttachingIsoParamDetail struct {
}

// GetCandidateVmForAttachingIsoParam GetCandidateVmForAttachingIso request param
type GetCandidateVmForAttachingIsoParam struct {
	BaseParam
	Params GetCandidateVmForAttachingIsoParamDetail `json:"getCandidateVmForAttachingIso"`
}
// AttachDataVolumeToVmParamDetail AttachDataVolumeToVm detail param
type AttachDataVolumeToVmParamDetail struct {
}

// AttachDataVolumeToVmParam AttachDataVolumeToVm request param
type AttachDataVolumeToVmParam struct {
	BaseParam
	Params AttachDataVolumeToVmParamDetail `json:"params"`
}
// UpdateAliyunVirtualRouterParamDetail UpdateAliyunVirtualRouter detail param
type UpdateAliyunVirtualRouterParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateAliyunVirtualRouterParam UpdateAliyunVirtualRouter request param
type UpdateAliyunVirtualRouterParam struct {
	BaseParam
	Params UpdateAliyunVirtualRouterParamDetail `json:"updateAliyunVirtualRouter"`
}
// DeleteDataVolumeParamDetail DeleteDataVolume detail param
type DeleteDataVolumeParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteDataVolumeParam DeleteDataVolume request param
type DeleteDataVolumeParam struct {
	BaseParam
	Params DeleteDataVolumeParamDetail `json:"deleteDataVolume"`
}
// GetUploadImageJobDetailsParamDetail GetUploadImageJobDetails detail param
type GetUploadImageJobDetailsParamDetail struct {
}

// GetUploadImageJobDetailsParam GetUploadImageJobDetails request param
type GetUploadImageJobDetailsParam struct {
	BaseParam
	Params GetUploadImageJobDetailsParamDetail `json:"getUploadImageJobDetails"`
}
// DetachIscsiServerFromClusterParamDetail DetachIscsiServerFromCluster detail param
type DetachIscsiServerFromClusterParamDetail struct {
}

// DetachIscsiServerFromClusterParam DetachIscsiServerFromCluster request param
type DetachIscsiServerFromClusterParam struct {
	BaseParam
	Params DetachIscsiServerFromClusterParamDetail `json:"detachIscsiServerFromCluster"`
}
// SetVolumeQosParamDetail SetVolumeQos detail param
type SetVolumeQosParamDetail struct {
	Mode *string `json:"mode,omitempty"`
	VolumeBandwidth *int64 `json:"volumeBandwidth,omitempty"`
	ReadBandwidth *int64 `json:"readBandwidth,omitempty"`
	WriteBandwidth *int64 `json:"writeBandwidth,omitempty"`
	TotalBandwidth *int64 `json:"totalBandwidth,omitempty"`
	ReadIOPS *int64 `json:"readIOPS,omitempty"`
	WriteIOPS *int64 `json:"writeIOPS,omitempty"`
	TotalIOPS *int64 `json:"totalIOPS,omitempty"`
}

// SetVolumeQosParam SetVolumeQos request param
type SetVolumeQosParam struct {
	BaseParam
	Params SetVolumeQosParamDetail `json:"setVolumeQos"`
}
// DetachHybridEipFromEcsParamDetail DetachHybridEipFromEcs detail param
type DetachHybridEipFromEcsParamDetail struct {
	Type string `json:"type" validate:"required"`
}

// DetachHybridEipFromEcsParam DetachHybridEipFromEcs request param
type DetachHybridEipFromEcsParam struct {
	BaseParam
	Params DetachHybridEipFromEcsParamDetail `json:"params"`
}
// GetVolumeCapabilitiesParamDetail GetVolumeCapabilities detail param
type GetVolumeCapabilitiesParamDetail struct {
}

// GetVolumeCapabilitiesParam GetVolumeCapabilities request param
type GetVolumeCapabilitiesParam struct {
	BaseParam
	Params GetVolumeCapabilitiesParamDetail `json:"getVolumeCapabilities"`
}
// ChangeBareMetal2GatewayClusterParamDetail ChangeBareMetal2GatewayCluster detail param
type ChangeBareMetal2GatewayClusterParamDetail struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// ChangeBareMetal2GatewayClusterParam ChangeBareMetal2GatewayCluster request param
type ChangeBareMetal2GatewayClusterParam struct {
	BaseParam
	Params ChangeBareMetal2GatewayClusterParamDetail `json:"changeBareMetal2GatewayCluster"`
}
// SetVmInstanceHaLevelParamDetail SetVmInstanceHaLevel detail param
type SetVmInstanceHaLevelParamDetail struct {
	Level string `json:"level" validate:"required"`
}

// SetVmInstanceHaLevelParam SetVmInstanceHaLevel request param
type SetVmInstanceHaLevelParam struct {
	BaseParam
	Params SetVmInstanceHaLevelParamDetail `json:"params"`
}
// RemoveVRouterNetworksFromFlowMeterParamDetail RemoveVRouterNetworksFromFlowMeter detail param
type RemoveVRouterNetworksFromFlowMeterParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveVRouterNetworksFromFlowMeterParam RemoveVRouterNetworksFromFlowMeter request param
type RemoveVRouterNetworksFromFlowMeterParam struct {
	BaseParam
	Params RemoveVRouterNetworksFromFlowMeterParamDetail `json:"removeVRouterNetworksFromFlowMeter"`
}
// GetCandidateL3NetworksForChangeVmNicNetworkParamDetail GetCandidateL3NetworksForChangeVmNicNetwork detail param
type GetCandidateL3NetworksForChangeVmNicNetworkParamDetail struct {
}

// GetCandidateL3NetworksForChangeVmNicNetworkParam GetCandidateL3NetworksForChangeVmNicNetwork request param
type GetCandidateL3NetworksForChangeVmNicNetworkParam struct {
	BaseParam
	Params GetCandidateL3NetworksForChangeVmNicNetworkParamDetail `json:"getCandidateL3NetworksForChangeVmNicNetwork"`
}
// DeleteHybridKeySecretParamDetail DeleteHybridKeySecret detail param
type DeleteHybridKeySecretParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteHybridKeySecretParam DeleteHybridKeySecret request param
type DeleteHybridKeySecretParam struct {
	BaseParam
	Params DeleteHybridKeySecretParamDetail `json:"deleteHybridKeySecret"`
}
// GetCandidatePrimaryStoragesForCreatingVmParamDetail GetCandidatePrimaryStoragesForCreatingVm detail param
type GetCandidatePrimaryStoragesForCreatingVmParamDetail struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	RootDiskOfferingUuid *string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize *int64 `json:"rootDiskSize,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	DataDiskSizes []int64 `json:"dataDiskSizes,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	InstanceOfferingUuid *string `json:"instanceOfferingUuid,omitempty"`
}

// GetCandidatePrimaryStoragesForCreatingVmParam GetCandidatePrimaryStoragesForCreatingVm request param
type GetCandidatePrimaryStoragesForCreatingVmParam struct {
	BaseParam
	Params GetCandidatePrimaryStoragesForCreatingVmParamDetail `json:"getCandidatePrimaryStoragesForCreatingVm"`
}
// GetVmConsolePasswordParamDetail GetVmConsolePassword detail param
type GetVmConsolePasswordParamDetail struct {
}

// GetVmConsolePasswordParam GetVmConsolePassword request param
type GetVmConsolePasswordParam struct {
	BaseParam
	Params GetVmConsolePasswordParamDetail `json:"getVmConsolePassword"`
}
// GetResourceBindableConfigParamDetail GetResourceBindableConfig detail param
type GetResourceBindableConfigParamDetail struct {
	Category *string `json:"category,omitempty"`
}

// GetResourceBindableConfigParam GetResourceBindableConfig request param
type GetResourceBindableConfigParam struct {
	BaseParam
	Params GetResourceBindableConfigParamDetail `json:"getResourceBindableConfig"`
}
// GetVmInstanceHaLevelParamDetail GetVmInstanceHaLevel detail param
type GetVmInstanceHaLevelParamDetail struct {
}

// GetVmInstanceHaLevelParam GetVmInstanceHaLevel request param
type GetVmInstanceHaLevelParam struct {
	BaseParam
	Params GetVmInstanceHaLevelParamDetail `json:"getVmInstanceHaLevel"`
}
// GetCandidateLdapEntryForIAM2BindingParamDetail GetCandidateLdapEntryForIAM2Binding detail param
type GetCandidateLdapEntryForIAM2BindingParamDetail struct {
	LdapFilter string `json:"ldapFilter" validate:"required"`
	Limit *int `json:"limit,omitempty"`
}

// GetCandidateLdapEntryForIAM2BindingParam GetCandidateLdapEntryForIAM2Binding request param
type GetCandidateLdapEntryForIAM2BindingParam struct {
	BaseParam
	Params GetCandidateLdapEntryForIAM2BindingParamDetail `json:"getCandidateLdapEntryForIAM2Binding"`
}
// RemoveResourcesFromDirectoryParamDetail RemoveResourcesFromDirectory detail param
type RemoveResourcesFromDirectoryParamDetail struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	DirectoryUuid string `json:"directoryUuid" validate:"required"`
}

// RemoveResourcesFromDirectoryParam RemoveResourcesFromDirectory request param
type RemoveResourcesFromDirectoryParam struct {
	BaseParam
	Params RemoveResourcesFromDirectoryParamDetail `json:"removeResourcesFromDirectory"`
}
// CreateVmFromVmBackupParamDetail CreateVmFromVmBackup detail param
type CreateVmFromVmBackupParamDetail struct {
	Name string `json:"name" validate:"required"`
	BackupStorageUuid *string `json:"backupStorageUuid,omitempty"`
	InstanceOfferingUuid *string `json:"instanceOfferingUuid,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	ReservedMemorySize *int64 `json:"reservedMemorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	VmNicParams *string `json:"vmNicParams,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume *string `json:"primaryStorageUuidForDataVolume,omitempty"`
	Description *string `json:"description,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmFromVmBackupParam CreateVmFromVmBackup request param
type CreateVmFromVmBackupParam struct {
	BaseParam
	Params CreateVmFromVmBackupParamDetail `json:"params"`
}
// DeleteExportedDatabaseBackupFromBackupStorageParamDetail DeleteExportedDatabaseBackupFromBackupStorage detail param
type DeleteExportedDatabaseBackupFromBackupStorageParamDetail struct {
}

// DeleteExportedDatabaseBackupFromBackupStorageParam DeleteExportedDatabaseBackupFromBackupStorage request param
type DeleteExportedDatabaseBackupFromBackupStorageParam struct {
	BaseParam
	Params DeleteExportedDatabaseBackupFromBackupStorageParamDetail `json:"deleteExportedDatabaseBackupFromBackupStorage"`
}
// AttachNetworkServiceToL3NetworkParamDetail AttachNetworkServiceToL3Network detail param
type AttachNetworkServiceToL3NetworkParamDetail struct {
	NetworkServices map[string]interface{} `json:"networkServices" validate:"required"`
}

// AttachNetworkServiceToL3NetworkParam AttachNetworkServiceToL3Network request param
type AttachNetworkServiceToL3NetworkParam struct {
	BaseParam
	Params AttachNetworkServiceToL3NetworkParamDetail `json:"params"`
}
// UnexportNbdVolumesParamDetail UnexportNbdVolumes detail param
type UnexportNbdVolumesParamDetail struct {
	VolumeUuids []string `json:"volumeUuids" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// UnexportNbdVolumesParam UnexportNbdVolumes request param
type UnexportNbdVolumesParam struct {
	BaseParam
	Params UnexportNbdVolumesParamDetail `json:"params"`
}
// RecoveryVirtualBorderRouterRemoteParamDetail RecoveryVirtualBorderRouterRemote detail param
type RecoveryVirtualBorderRouterRemoteParamDetail struct {
}

// RecoveryVirtualBorderRouterRemoteParam RecoveryVirtualBorderRouterRemote request param
type RecoveryVirtualBorderRouterRemoteParam struct {
	BaseParam
	Params RecoveryVirtualBorderRouterRemoteParamDetail `json:"recoveryVirtualBorderRouterRemote"`
}
// ExecuteAutoScalingRuleParamDetail ExecuteAutoScalingRule detail param
type ExecuteAutoScalingRuleParamDetail struct {
}

// ExecuteAutoScalingRuleParam ExecuteAutoScalingRule request param
type ExecuteAutoScalingRuleParam struct {
	BaseParam
	Params ExecuteAutoScalingRuleParamDetail `json:"executeAutoScalingRule"`
}
// SNSHttpTestConnectionParamDetail SNSHttpTestConnection detail param
type SNSHttpTestConnectionParamDetail struct {
	Url *string `json:"url,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	EndpointUuid *string `json:"endpointUuid,omitempty"`
}

// SNSHttpTestConnectionParam SNSHttpTestConnection request param
type SNSHttpTestConnectionParam struct {
	BaseParam
	Params SNSHttpTestConnectionParamDetail `json:"params"`
}
// SetImageSecurityLevelParamDetail SetImageSecurityLevel detail param
type SetImageSecurityLevelParamDetail struct {
	SecurityLevel *string `json:"securityLevel,omitempty"`
}

// SetImageSecurityLevelParam SetImageSecurityLevel request param
type SetImageSecurityLevelParam struct {
	BaseParam
	Params SetImageSecurityLevelParamDetail `json:"setImageSecurityLevel"`
}
// ChangeBareMetal2ChassisStateParamDetail ChangeBareMetal2ChassisState detail param
type ChangeBareMetal2ChassisStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBareMetal2ChassisStateParam ChangeBareMetal2ChassisState request param
type ChangeBareMetal2ChassisStateParam struct {
	BaseParam
	Params ChangeBareMetal2ChassisStateParamDetail `json:"changeBareMetal2ChassisState"`
}
// AddHybridKeySecretParamDetail AddHybridKeySecret detail param
type AddHybridKeySecretParamDetail struct {
	Name string `json:"name" validate:"required"`
	Key string `json:"key" validate:"required"`
	Secret string `json:"secret" validate:"required"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	Sync *bool `json:"sync,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddHybridKeySecretParam AddHybridKeySecret request param
type AddHybridKeySecretParam struct {
	BaseParam
	Params AddHybridKeySecretParamDetail `json:"params"`
}
// DetachVmFromVmSchedulingRuleGroupParamDetail DetachVmFromVmSchedulingRuleGroup detail param
type DetachVmFromVmSchedulingRuleGroupParamDetail struct {
	VmUuid string `json:"vmUuid" validate:"required"`
}

// DetachVmFromVmSchedulingRuleGroupParam DetachVmFromVmSchedulingRuleGroup request param
type DetachVmFromVmSchedulingRuleGroupParam struct {
	BaseParam
	Params DetachVmFromVmSchedulingRuleGroupParamDetail `json:"detachVmFromVmSchedulingRuleGroup"`
}
// AddVRouterNetworksToOspfAreaParamDetail AddVRouterNetworksToOspfArea detail param
type AddVRouterNetworksToOspfAreaParamDetail struct {
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddVRouterNetworksToOspfAreaParam AddVRouterNetworksToOspfArea request param
type AddVRouterNetworksToOspfAreaParam struct {
	BaseParam
	Params AddVRouterNetworksToOspfAreaParamDetail `json:"params"`
}
// AddRolesToIAM2VirtualIDGroupParamDetail AddRolesToIAM2VirtualIDGroup detail param
type AddRolesToIAM2VirtualIDGroupParamDetail struct {
	RoleUuids []string `json:"roleUuids" validate:"required"`
	ProjectUuid *string `json:"projectUuid,omitempty"`
}

// AddRolesToIAM2VirtualIDGroupParam AddRolesToIAM2VirtualIDGroup request param
type AddRolesToIAM2VirtualIDGroupParam struct {
	BaseParam
	Params AddRolesToIAM2VirtualIDGroupParamDetail `json:"params"`
}
// CheckStaticProvisionIpParamDetail CheckStaticProvisionIp detail param
type CheckStaticProvisionIpParamDetail struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ProvisionIp string `json:"provisionIp" validate:"required"`
}

// CheckStaticProvisionIpParam CheckStaticProvisionIp request param
type CheckStaticProvisionIpParam struct {
	BaseParam
	Params CheckStaticProvisionIpParamDetail `json:"params"`
}
// ChangeEventSubscriptionStateParamDetail ChangeEventSubscriptionState detail param
type ChangeEventSubscriptionStateParamDetail struct {
	State string `json:"state" validate:"required"`
}

// ChangeEventSubscriptionStateParam ChangeEventSubscriptionState request param
type ChangeEventSubscriptionStateParam struct {
	BaseParam
	Params ChangeEventSubscriptionStateParamDetail `json:"changeEventSubscriptionState"`
}
// PushLicenseAddOnsUsageParamDetail PushLicenseAddOnsUsage detail param
type PushLicenseAddOnsUsageParamDetail struct {
	AddOnsUsage string `json:"addOnsUsage" validate:"required"`
}

// PushLicenseAddOnsUsageParam PushLicenseAddOnsUsage request param
type PushLicenseAddOnsUsageParam struct {
	BaseParam
	Params PushLicenseAddOnsUsageParamDetail `json:"pushLicenseAddOnsUsage"`
}
// AttachHybridEipToEcsParamDetail AttachHybridEipToEcs detail param
type AttachHybridEipToEcsParamDetail struct {
	EcsUuid string `json:"ecsUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
}

// AttachHybridEipToEcsParam AttachHybridEipToEcs request param
type AttachHybridEipToEcsParam struct {
	BaseParam
	Params AttachHybridEipToEcsParamDetail `json:"params"`
}
// CreateEcsImageFromLocalImageParamDetail CreateEcsImageFromLocalImage detail param
type CreateEcsImageFromLocalImageParamDetail struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	BackupStorageUuid *string `json:"backupStorageUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsImageFromLocalImageParam CreateEcsImageFromLocalImage request param
type CreateEcsImageFromLocalImageParam struct {
	BaseParam
	Params CreateEcsImageFromLocalImageParamDetail `json:"params"`
}
// AddHostRouteToL3NetworkParamDetail AddHostRouteToL3Network detail param
type AddHostRouteToL3NetworkParamDetail struct {
	Prefix string `json:"prefix" validate:"required"`
	Nexthop string `json:"nexthop" validate:"required"`
}

// AddHostRouteToL3NetworkParam AddHostRouteToL3Network request param
type AddHostRouteToL3NetworkParam struct {
	BaseParam
	Params AddHostRouteToL3NetworkParamDetail `json:"params"`
}
// AddInstanceToMonitorGroupParamDetail AddInstanceToMonitorGroup detail param
type AddInstanceToMonitorGroupParamDetail struct {
	InstanceUuid string `json:"instanceUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddInstanceToMonitorGroupParam AddInstanceToMonitorGroup request param
type AddInstanceToMonitorGroupParam struct {
	BaseParam
	Params AddInstanceToMonitorGroupParamDetail `json:"params"`
}
// GetBareMetal2ProvisionNetworkIpAddressCapacityParamDetail GetBareMetal2ProvisionNetworkIpAddressCapacity detail param
type GetBareMetal2ProvisionNetworkIpAddressCapacityParamDetail struct {
	NetworkUuids []string `json:"networkUuids" validate:"required"`
}

// GetBareMetal2ProvisionNetworkIpAddressCapacityParam GetBareMetal2ProvisionNetworkIpAddressCapacity request param
type GetBareMetal2ProvisionNetworkIpAddressCapacityParam struct {
	BaseParam
	Params GetBareMetal2ProvisionNetworkIpAddressCapacityParamDetail `json:"getBareMetal2ProvisionNetworkIpAddressCapacity"`
}
// AttachMdevDeviceToVmParamDetail AttachMdevDeviceToVm detail param
type AttachMdevDeviceToVmParamDetail struct {
}

// AttachMdevDeviceToVmParam AttachMdevDeviceToVm request param
type AttachMdevDeviceToVmParam struct {
	BaseParam
	Params AttachMdevDeviceToVmParamDetail `json:"params"`
}
// DecodeStackTemplateParamDetail DecodeStackTemplate detail param
type DecodeStackTemplateParamDetail struct {
	Type *string `json:"type,omitempty"`
	TemplateContent *string `json:"templateContent,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
	Preparameters *string `json:"preparameters,omitempty"`
}

// DecodeStackTemplateParam DecodeStackTemplate request param
type DecodeStackTemplateParam struct {
	BaseParam
	Params DecodeStackTemplateParamDetail `json:"params"`
}
// UpdateVirtualRouterParamDetail UpdateVirtualRouter detail param
type UpdateVirtualRouterParamDetail struct {
	DefaultRouteL3NetworkUuid *string `json:"defaultRouteL3NetworkUuid,omitempty"`
}

// UpdateVirtualRouterParam UpdateVirtualRouter request param
type UpdateVirtualRouterParam struct {
	BaseParam
	Params UpdateVirtualRouterParamDetail `json:"updateVirtualRouter"`
}
// GetVSwitchTypesParamDetail GetVSwitchTypes detail param
type GetVSwitchTypesParamDetail struct {
}

// GetVSwitchTypesParam GetVSwitchTypes request param
type GetVSwitchTypesParam struct {
	BaseParam
	Params GetVSwitchTypesParamDetail `json:"getVSwitchTypes"`
}
// CreateL2HardwareVxlanNetworkPoolParamDetail CreateL2HardwareVxlanNetworkPool detail param
type CreateL2HardwareVxlanNetworkPoolParamDetail struct {
	SdnControllerUuid string `json:"sdnControllerUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	PhysicalInterface string `json:"physicalInterface" validate:"required"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Pvlan *string `json:"pvlan,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2HardwareVxlanNetworkPoolParam CreateL2HardwareVxlanNetworkPool request param
type CreateL2HardwareVxlanNetworkPoolParam struct {
	BaseParam
	Params CreateL2HardwareVxlanNetworkPoolParamDetail `json:"params"`
}
// GetLdapServerAvailableAttributesParamDetail GetLdapServerAvailableAttributes detail param
type GetLdapServerAvailableAttributesParamDetail struct {
}

// GetLdapServerAvailableAttributesParam GetLdapServerAvailableAttributes request param
type GetLdapServerAvailableAttributesParam struct {
	BaseParam
	Params GetLdapServerAvailableAttributesParamDetail `json:"getLdapServerAvailableAttributes"`
}
// ResizeDataVolumeParamDetail ResizeDataVolume detail param
type ResizeDataVolumeParamDetail struct {
	Size int64 `json:"size" validate:"required"`
}

// ResizeDataVolumeParam ResizeDataVolume request param
type ResizeDataVolumeParam struct {
	BaseParam
	Params ResizeDataVolumeParamDetail `json:"resizeDataVolume"`
}
// GetEipAttachableVmNicsParamDetail GetEipAttachableVmNics detail param
type GetEipAttachableVmNicsParamDetail struct {
	VipUuid *string `json:"vipUuid,omitempty"`
	VmUuid *string `json:"vmUuid,omitempty"`
	VmName *string `json:"vmName,omitempty"`
	NetworkServiceProvider *string `json:"networkServiceProvider,omitempty"`
	AttachedToVm *bool `json:"attachedToVm,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetEipAttachableVmNicsParam GetEipAttachableVmNics request param
type GetEipAttachableVmNicsParam struct {
	BaseParam
	Params GetEipAttachableVmNicsParamDetail `json:"getEipAttachableVmNics"`
}
// AddIpv6RangeByNetworkCidrParamDetail AddIpv6RangeByNetworkCidr detail param
type AddIpv6RangeByNetworkCidrParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	NetworkCidr string `json:"networkCidr" validate:"required"`
	AddressMode string `json:"addressMode" validate:"required"`
	IpRangeType *string `json:"ipRangeType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIpv6RangeByNetworkCidrParam AddIpv6RangeByNetworkCidr request param
type AddIpv6RangeByNetworkCidrParam struct {
	BaseParam
	Params AddIpv6RangeByNetworkCidrParamDetail `json:"params"`
}
// BatchQueryParamDetail BatchQuery detail param
type BatchQueryParamDetail struct {
	Script *string `json:"script,omitempty"`
}

// BatchQueryParam BatchQuery request param
type BatchQueryParam struct {
	BaseParam
	Params BatchQueryParamDetail `json:"batchQuery"`
}
// ReloadExternalServiceParamDetail ReloadExternalService detail param
type ReloadExternalServiceParamDetail struct {
	Name string `json:"name" validate:"required"`
}

// ReloadExternalServiceParam ReloadExternalService request param
type ReloadExternalServiceParam struct {
	BaseParam
	Params ReloadExternalServiceParamDetail `json:"reloadExternalService"`
}
// AddIAM2VirtualIDsToGroupParamDetail AddIAM2VirtualIDsToGroup detail param
type AddIAM2VirtualIDsToGroupParamDetail struct {
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
}

// AddIAM2VirtualIDsToGroupParam AddIAM2VirtualIDsToGroup request param
type AddIAM2VirtualIDsToGroupParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToGroupParamDetail `json:"params"`
}
// CreateIAM2VirtualIDLdapBindingParamDetail CreateIAM2VirtualIDLdapBinding detail param
type CreateIAM2VirtualIDLdapBindingParamDetail struct {
	VirtualIDUuid string `json:"virtualIDUuid" validate:"required"`
	LdapUid string `json:"ldapUid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDLdapBindingParam CreateIAM2VirtualIDLdapBinding request param
type CreateIAM2VirtualIDLdapBindingParam struct {
	BaseParam
	Params CreateIAM2VirtualIDLdapBindingParamDetail `json:"params"`
}
// SetVmNicSecurityGroupParamDetail SetVmNicSecurityGroup detail param
type SetVmNicSecurityGroupParamDetail struct {
	Refs []SetVmNicSecurityGroup_VmNicSecurityGroupRefAOParam `json:"refs" validate:"required"`
}

// SetVmNicSecurityGroupParam SetVmNicSecurityGroup request param
type SetVmNicSecurityGroupParam struct {
	BaseParam
	Params SetVmNicSecurityGroupParamDetail `json:"setVmNicSecurityGroup"`
}
// AddIdentityZoneFromRemoteParamDetail AddIdentityZoneFromRemote detail param
type AddIdentityZoneFromRemoteParamDetail struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ZoneId *string `json:"zoneId,omitempty"`
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIdentityZoneFromRemoteParam AddIdentityZoneFromRemote request param
type AddIdentityZoneFromRemoteParam struct {
	BaseParam
	Params AddIdentityZoneFromRemoteParamDetail `json:"params"`
}
// GetVolumeSnapshotSizeParamDetail GetVolumeSnapshotSize detail param
type GetVolumeSnapshotSizeParamDetail struct {
}

// GetVolumeSnapshotSizeParam GetVolumeSnapshotSize request param
type GetVolumeSnapshotSizeParam struct {
	BaseParam
	Params GetVolumeSnapshotSizeParamDetail `json:"getVolumeSnapshotSize"`
}
// BatchSyncVolumeSizeParamDetail BatchSyncVolumeSize detail param
type BatchSyncVolumeSizeParamDetail struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// BatchSyncVolumeSizeParam BatchSyncVolumeSize request param
type BatchSyncVolumeSizeParam struct {
	BaseParam
	Params BatchSyncVolumeSizeParamDetail `json:"batchSyncVolumeSize"`
}
// GetHypervisorTypesParamDetail GetHypervisorTypes detail param
type GetHypervisorTypesParamDetail struct {
}

// GetHypervisorTypesParam GetHypervisorTypes request param
type GetHypervisorTypesParam struct {
	BaseParam
	Params GetHypervisorTypesParamDetail `json:"getHypervisorTypes"`
}
// GetVmAttachableDataVolumeParamDetail GetVmAttachableDataVolume detail param
type GetVmAttachableDataVolumeParamDetail struct {
}

// GetVmAttachableDataVolumeParam GetVmAttachableDataVolume request param
type GetVmAttachableDataVolumeParam struct {
	BaseParam
	Params GetVmAttachableDataVolumeParamDetail `json:"getVmAttachableDataVolume"`
}
// GetVmMonitorNumberParamDetail GetVmMonitorNumber detail param
type GetVmMonitorNumberParamDetail struct {
}

// GetVmMonitorNumberParam GetVmMonitorNumber request param
type GetVmMonitorNumberParam struct {
	BaseParam
	Params GetVmMonitorNumberParamDetail `json:"getVmMonitorNumber"`
}
// CreateIAM2VirtualIDFromLdapUidParamDetail CreateIAM2VirtualIDFromLdapUid detail param
type CreateIAM2VirtualIDFromLdapUidParamDetail struct {
	LdapUid string `json:"ldapUid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDFromLdapUidParam CreateIAM2VirtualIDFromLdapUid request param
type CreateIAM2VirtualIDFromLdapUidParam struct {
	BaseParam
	Params CreateIAM2VirtualIDFromLdapUidParamDetail `json:"params"`
}
// ValidatePriceUserConfigParamDetail ValidatePriceUserConfig detail param
type ValidatePriceUserConfigParamDetail struct {
	Config string `json:"config" validate:"required"`
}

// ValidatePriceUserConfigParam ValidatePriceUserConfig request param
type ValidatePriceUserConfigParam struct {
	BaseParam
	Params ValidatePriceUserConfigParamDetail `json:"validatePriceUserConfig"`
}
// ChangeBareMetal2GatewayStateParamDetail ChangeBareMetal2GatewayState detail param
type ChangeBareMetal2GatewayStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBareMetal2GatewayStateParam ChangeBareMetal2GatewayState request param
type ChangeBareMetal2GatewayStateParam struct {
	BaseParam
	Params ChangeBareMetal2GatewayStateParamDetail `json:"changeBareMetal2GatewayState"`
}
// RemoveActionFromEventSubscriptionParamDetail RemoveActionFromEventSubscription detail param
type RemoveActionFromEventSubscriptionParamDetail struct {
}

// RemoveActionFromEventSubscriptionParam RemoveActionFromEventSubscription request param
type RemoveActionFromEventSubscriptionParam struct {
	BaseParam
	Params RemoveActionFromEventSubscriptionParamDetail `json:"removeActionFromEventSubscription"`
}
// CheckKVMHostConfigFileParamDetail CheckKVMHostConfigFile detail param
type CheckKVMHostConfigFileParamDetail struct {
	HostInfo string `json:"hostInfo" validate:"required"`
}

// CheckKVMHostConfigFileParam CheckKVMHostConfigFile request param
type CheckKVMHostConfigFileParam struct {
	BaseParam
	Params CheckKVMHostConfigFileParamDetail `json:"params"`
}
// GetContainerUsageParamDetail GetContainerUsage detail param
type GetContainerUsageParamDetail struct {
}

// GetContainerUsageParam GetContainerUsage request param
type GetContainerUsageParam struct {
	BaseParam
	Params GetContainerUsageParamDetail `json:"getContainerUsage"`
}
// SNSSnmpTestConnectionParamDetail SNSSnmpTestConnection detail param
type SNSSnmpTestConnectionParamDetail struct {
	PlatformUuid *string `json:"platformUuid,omitempty"`
	EndpointUuid *string `json:"endpointUuid,omitempty"`
}

// SNSSnmpTestConnectionParam SNSSnmpTestConnection request param
type SNSSnmpTestConnectionParam struct {
	BaseParam
	Params SNSSnmpTestConnectionParamDetail `json:"params"`
}
// GetDataCenterFromRemoteParamDetail GetDataCenterFromRemote detail param
type GetDataCenterFromRemoteParamDetail struct {
	Type string `json:"type" validate:"required"`
	Endpoint *string `json:"endpoint,omitempty"`
}

// GetDataCenterFromRemoteParam GetDataCenterFromRemote request param
type GetDataCenterFromRemoteParam struct {
	BaseParam
	Params GetDataCenterFromRemoteParamDetail `json:"getDataCenterFromRemote"`
}
// CreateHostNetworkServiceTypeParamDetail CreateHostNetworkServiceType detail param
type CreateHostNetworkServiceTypeParamDetail struct {
	ServiceType string `json:"serviceType" validate:"required"`
	System *bool `json:"system,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateHostNetworkServiceTypeParam CreateHostNetworkServiceType request param
type CreateHostNetworkServiceTypeParam struct {
	BaseParam
	Params CreateHostNetworkServiceTypeParamDetail `json:"params"`
}
// DeleteEcsImageLocalParamDetail DeleteEcsImageLocal detail param
type DeleteEcsImageLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsImageLocalParam DeleteEcsImageLocal request param
type DeleteEcsImageLocalParam struct {
	BaseParam
	Params DeleteEcsImageLocalParamDetail `json:"deleteEcsImageLocal"`
}
// DetachNvmeServerFromClusterParamDetail DetachNvmeServerFromCluster detail param
type DetachNvmeServerFromClusterParamDetail struct {
}

// DetachNvmeServerFromClusterParam DetachNvmeServerFromCluster request param
type DetachNvmeServerFromClusterParam struct {
	BaseParam
	Params DetachNvmeServerFromClusterParamDetail `json:"detachNvmeServerFromCluster"`
}
// GetBackupStorageTypesParamDetail GetBackupStorageTypes detail param
type GetBackupStorageTypesParamDetail struct {
}

// GetBackupStorageTypesParam GetBackupStorageTypes request param
type GetBackupStorageTypesParam struct {
	BaseParam
	Params GetBackupStorageTypesParamDetail `json:"getBackupStorageTypes"`
}
// GetVolumeQosParamDetail GetVolumeQos detail param
type GetVolumeQosParamDetail struct {
	ForceSync *bool `json:"forceSync,omitempty"`
}

// GetVolumeQosParam GetVolumeQos request param
type GetVolumeQosParam struct {
	BaseParam
	Params GetVolumeQosParamDetail `json:"getVolumeQos"`
}
// AddRemoteCidrsToIPsecConnectionParamDetail AddRemoteCidrsToIPsecConnection detail param
type AddRemoteCidrsToIPsecConnectionParamDetail struct {
	PeerCidrs []string `json:"peerCidrs" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddRemoteCidrsToIPsecConnectionParam AddRemoteCidrsToIPsecConnection request param
type AddRemoteCidrsToIPsecConnectionParam struct {
	BaseParam
	Params AddRemoteCidrsToIPsecConnectionParamDetail `json:"params"`
}
// PowerOnBaremetalChassisParamDetail PowerOnBaremetalChassis detail param
type PowerOnBaremetalChassisParamDetail struct {
}

// PowerOnBaremetalChassisParam PowerOnBaremetalChassis request param
type PowerOnBaremetalChassisParam struct {
	BaseParam
	Params PowerOnBaremetalChassisParamDetail `json:"powerOnBaremetalChassis"`
}
// RequestLicenseCapacityParamDetail RequestLicenseCapacity detail param
type RequestLicenseCapacityParamDetail struct {
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	QuotaType string `json:"quotaType" validate:"required"`
	Quota int64 `json:"quota" validate:"required"`
	ClientAuthorizedNodeUuid string `json:"clientAuthorizedNodeUuid" validate:"required"`
	LicenseType string `json:"licenseType" validate:"required"`
	ResourceInfo *string `json:"resourceInfo,omitempty"`
}

// RequestLicenseCapacityParam RequestLicenseCapacity request param
type RequestLicenseCapacityParam struct {
	BaseParam
	Params RequestLicenseCapacityParamDetail `json:"params"`
}
// CreateDataVolumeFromVolumeSnapshotParamDetail CreateDataVolumeFromVolumeSnapshot detail param
type CreateDataVolumeFromVolumeSnapshotParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeFromVolumeSnapshotParam CreateDataVolumeFromVolumeSnapshot request param
type CreateDataVolumeFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateDataVolumeFromVolumeSnapshotParamDetail `json:"params"`
}
// DetachIsoFromVmInstanceParamDetail DetachIsoFromVmInstance detail param
type DetachIsoFromVmInstanceParamDetail struct {
	IsoUuid *string `json:"isoUuid,omitempty"`
}

// DetachIsoFromVmInstanceParam DetachIsoFromVmInstance request param
type DetachIsoFromVmInstanceParam struct {
	BaseParam
	Params DetachIsoFromVmInstanceParamDetail `json:"detachIsoFromVmInstance"`
}
// DetachSecurityGroupFromL3NetworkParamDetail DetachSecurityGroupFromL3Network detail param
type DetachSecurityGroupFromL3NetworkParamDetail struct {
}

// DetachSecurityGroupFromL3NetworkParam DetachSecurityGroupFromL3Network request param
type DetachSecurityGroupFromL3NetworkParam struct {
	BaseParam
	Params DetachSecurityGroupFromL3NetworkParamDetail `json:"detachSecurityGroupFromL3Network"`
}
// GetVirtualizerInfoParamDetail GetVirtualizerInfo detail param
type GetVirtualizerInfoParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// GetVirtualizerInfoParam GetVirtualizerInfo request param
type GetVirtualizerInfoParam struct {
	BaseParam
	Params GetVirtualizerInfoParamDetail `json:"getVirtualizerInfo"`
}
// GetL3NetworkIpStatisticParamDetail GetL3NetworkIpStatistic detail param
type GetL3NetworkIpStatisticParamDetail struct {
	ResourceType *string `json:"resourceType,omitempty"`
	Ip *string `json:"ip,omitempty"`
	SortBy *string `json:"sortBy,omitempty"`
	SortDirection *string `json:"sortDirection,omitempty"`
	Start *int `json:"start,omitempty"`
	Limit *int `json:"limit,omitempty"`
	ReplyWithCount *bool `json:"replyWithCount,omitempty"`
}

// GetL3NetworkIpStatisticParam GetL3NetworkIpStatistic request param
type GetL3NetworkIpStatisticParam struct {
	BaseParam
	Params GetL3NetworkIpStatisticParamDetail `json:"getL3NetworkIpStatistic"`
}
// GetImageCandidatesForVmToChangeParamDetail GetImageCandidatesForVmToChange detail param
type GetImageCandidatesForVmToChangeParamDetail struct {
}

// GetImageCandidatesForVmToChangeParam GetImageCandidatesForVmToChange request param
type GetImageCandidatesForVmToChangeParam struct {
	BaseParam
	Params GetImageCandidatesForVmToChangeParamDetail `json:"getImageCandidatesForVmToChange"`
}
// ChangeImageStateParamDetail ChangeImageState detail param
type ChangeImageStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeImageStateParam ChangeImageState request param
type ChangeImageStateParam struct {
	BaseParam
	Params ChangeImageStateParamDetail `json:"changeImageState"`
}
// KvmRunShellParamDetail KvmRunShell detail param
type KvmRunShellParamDetail struct {
	HostUuids []string `json:"hostUuids" validate:"required"`
	Script string `json:"script" validate:"required"`
}

// KvmRunShellParam KvmRunShell request param
type KvmRunShellParam struct {
	BaseParam
	Params KvmRunShellParamDetail `json:"kvmRunShell"`
}
// CreateAliyunNasAccessGroupRuleParamDetail CreateAliyunNasAccessGroupRule detail param
type CreateAliyunNasAccessGroupRuleParamDetail struct {
	AccessGroupUuid string `json:"accessGroupUuid" validate:"required"`
	SourceCidrIp string `json:"sourceCidrIp" validate:"required"`
	RwAccessType *string `json:"rwAccessType,omitempty"`
	Priority *int `json:"priority,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunNasAccessGroupRuleParam CreateAliyunNasAccessGroupRule request param
type CreateAliyunNasAccessGroupRuleParam struct {
	BaseParam
	Params CreateAliyunNasAccessGroupRuleParamDetail `json:"params"`
}
// RecoverBackupFromImageStoreBackupStorageParamDetail RecoverBackupFromImageStoreBackupStorage detail param
type RecoverBackupFromImageStoreBackupStorageParamDetail struct {
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// RecoverBackupFromImageStoreBackupStorageParam RecoverBackupFromImageStoreBackupStorage request param
type RecoverBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params RecoverBackupFromImageStoreBackupStorageParamDetail `json:"recoverBackupFromImageStoreBackupStorage"`
}
// ChangeTicketFlowCollectionStateParamDetail ChangeTicketFlowCollectionState detail param
type ChangeTicketFlowCollectionStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeTicketFlowCollectionStateParam ChangeTicketFlowCollectionState request param
type ChangeTicketFlowCollectionStateParam struct {
	BaseParam
	Params ChangeTicketFlowCollectionStateParamDetail `json:"changeTicketFlowCollectionState"`
}
// ExpungeDataVolumeParamDetail ExpungeDataVolume detail param
type ExpungeDataVolumeParamDetail struct {
}

// ExpungeDataVolumeParam ExpungeDataVolume request param
type ExpungeDataVolumeParam struct {
	BaseParam
	Params ExpungeDataVolumeParamDetail `json:"expungeDataVolume"`
}
// AddActionToEventSubscriptionParamDetail AddActionToEventSubscription detail param
type AddActionToEventSubscriptionParamDetail struct {
	ActionUuid string `json:"actionUuid" validate:"required"`
	ActionType string `json:"actionType" validate:"required"`
}

// AddActionToEventSubscriptionParam AddActionToEventSubscription request param
type AddActionToEventSubscriptionParam struct {
	BaseParam
	Params AddActionToEventSubscriptionParamDetail `json:"params"`
}
// GetVRouterRouterIdParamDetail GetVRouterRouterId detail param
type GetVRouterRouterIdParamDetail struct {
}

// GetVRouterRouterIdParam GetVRouterRouterId request param
type GetVRouterRouterIdParam struct {
	BaseParam
	Params GetVRouterRouterIdParamDetail `json:"getVRouterRouterId"`
}
// GetZBoxBackupDetailsParamDetail GetZBoxBackupDetails detail param
type GetZBoxBackupDetailsParamDetail struct {
}

// GetZBoxBackupDetailsParam GetZBoxBackupDetails request param
type GetZBoxBackupDetailsParam struct {
	BaseParam
	Params GetZBoxBackupDetailsParamDetail `json:"getZBoxBackupDetails"`
}
// GetExternalServicesParamDetail GetExternalServices detail param
type GetExternalServicesParamDetail struct {
}

// GetExternalServicesParam GetExternalServices request param
type GetExternalServicesParam struct {
	BaseParam
	Params GetExternalServicesParamDetail `json:"getExternalServices"`
}
// GetIAM2ProjectRepositoryParamDetail GetIAM2ProjectRepository detail param
type GetIAM2ProjectRepositoryParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetIAM2ProjectRepositoryParam GetIAM2ProjectRepository request param
type GetIAM2ProjectRepositoryParam struct {
	BaseParam
	Params GetIAM2ProjectRepositoryParamDetail `json:"getIAM2ProjectRepository"`
}
// GetCandidateNetworkInterfacesParamDetail GetCandidateNetworkInterfaces detail param
type GetCandidateNetworkInterfacesParamDetail struct {
	HostUuids []string `json:"hostUuids" validate:"required"`
	InterfaceType *string `json:"interfaceType,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetCandidateNetworkInterfacesParam GetCandidateNetworkInterfaces request param
type GetCandidateNetworkInterfacesParam struct {
	BaseParam
	Params GetCandidateNetworkInterfacesParamDetail `json:"getCandidateNetworkInterfaces"`
}
// ChangeAccessControlListServerGroupParamDetail ChangeAccessControlListServerGroup detail param
type ChangeAccessControlListServerGroupParamDetail struct {
	ServerGroupUuids []string `json:"serverGroupUuids" validate:"required"`
	ListenerUuid string `json:"listenerUuid" validate:"required"`
}

// ChangeAccessControlListServerGroupParam ChangeAccessControlListServerGroup request param
type ChangeAccessControlListServerGroupParam struct {
	BaseParam
	Params ChangeAccessControlListServerGroupParamDetail `json:"changeAccessControlListServerGroup"`
}
// SyncVirtualBorderRouterFromRemoteParamDetail SyncVirtualBorderRouterFromRemote detail param
type SyncVirtualBorderRouterFromRemoteParamDetail struct {
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncVirtualBorderRouterFromRemoteParam SyncVirtualBorderRouterFromRemote request param
type SyncVirtualBorderRouterFromRemoteParam struct {
	BaseParam
	Params SyncVirtualBorderRouterFromRemoteParamDetail `json:"syncVirtualBorderRouterFromRemote"`
}
// UpdateAtPersonOfAtFeiShuEndpointParamDetail UpdateAtPersonOfAtFeiShuEndpoint detail param
type UpdateAtPersonOfAtFeiShuEndpointParamDetail struct {
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	UserId *string `json:"userId,omitempty"`
	Remark *string `json:"remark,omitempty"`
}

// UpdateAtPersonOfAtFeiShuEndpointParam UpdateAtPersonOfAtFeiShuEndpoint request param
type UpdateAtPersonOfAtFeiShuEndpointParam struct {
	BaseParam
	Params UpdateAtPersonOfAtFeiShuEndpointParamDetail `json:"updateAtPersonOfAtFeiShuEndpoint"`
}
// CreateL2HardwareVxlanNetworkParamDetail CreateL2HardwareVxlanNetwork detail param
type CreateL2HardwareVxlanNetworkParamDetail struct {
	Vni *int `json:"vni,omitempty"`
	PoolUuid string `json:"poolUuid" validate:"required"`
	H3cTenantUuid *string `json:"h3cTenantUuid,omitempty"`
	Vlan *int `json:"vlan,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	PhysicalInterface *string `json:"physicalInterface,omitempty"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Pvlan *string `json:"pvlan,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2HardwareVxlanNetworkParam CreateL2HardwareVxlanNetwork request param
type CreateL2HardwareVxlanNetworkParam struct {
	BaseParam
	Params CreateL2HardwareVxlanNetworkParamDetail `json:"params"`
}
// GetGlobalConfigOptionsParamDetail GetGlobalConfigOptions detail param
type GetGlobalConfigOptionsParamDetail struct {
}

// GetGlobalConfigOptionsParam GetGlobalConfigOptions request param
type GetGlobalConfigOptionsParam struct {
	BaseParam
	Params GetGlobalConfigOptionsParamDetail `json:"getGlobalConfigOptions"`
}
// CreateHybridEipParamDetail CreateHybridEip detail param
type CreateHybridEipParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	BandWidthMb int64 `json:"bandWidthMb" validate:"required"`
	Type string `json:"type" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ChargeType string `json:"chargeType" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateHybridEipParam CreateHybridEip request param
type CreateHybridEipParam struct {
	BaseParam
	Params CreateHybridEipParamDetail `json:"params"`
}
// ApplyMonitorTemplateToMonitorGroupParamDetail ApplyMonitorTemplateToMonitorGroup detail param
type ApplyMonitorTemplateToMonitorGroupParamDetail struct {
}

// ApplyMonitorTemplateToMonitorGroupParam ApplyMonitorTemplateToMonitorGroup request param
type ApplyMonitorTemplateToMonitorGroupParam struct {
	BaseParam
	Params ApplyMonitorTemplateToMonitorGroupParamDetail `json:"params"`
}
// PutMetricDataParamDetail PutMetricData detail param
type PutMetricDataParamDetail struct {
	Namespace string `json:"namespace" validate:"required"`
	Data []MetricDatumParam `json:"data" validate:"required"`
}

// PutMetricDataParam PutMetricData request param
type PutMetricDataParam struct {
	BaseParam
	Params PutMetricDataParamDetail `json:"params"`
}
// GetAttachablePublicL3ForVRouterParamDetail GetAttachablePublicL3ForVRouter detail param
type GetAttachablePublicL3ForVRouterParamDetail struct {
}

// GetAttachablePublicL3ForVRouterParam GetAttachablePublicL3ForVRouter request param
type GetAttachablePublicL3ForVRouterParam struct {
	BaseParam
	Params GetAttachablePublicL3ForVRouterParamDetail `json:"getAttachablePublicL3ForVRouter"`
}
// RerunLongJobParamDetail RerunLongJob detail param
type RerunLongJobParamDetail struct {
}

// RerunLongJobParam RerunLongJob request param
type RerunLongJobParam struct {
	BaseParam
	Params RerunLongJobParamDetail `json:"rerunLongJob"`
}
// DeleteExportedImageFromBackupStorageParamDetail DeleteExportedImageFromBackupStorage detail param
type DeleteExportedImageFromBackupStorageParamDetail struct {
}

// DeleteExportedImageFromBackupStorageParam DeleteExportedImageFromBackupStorage request param
type DeleteExportedImageFromBackupStorageParam struct {
	BaseParam
	Params DeleteExportedImageFromBackupStorageParamDetail `json:"deleteExportedImageFromBackupStorage"`
}
// UpdateClusterOSParamDetail UpdateClusterOS detail param
type UpdateClusterOSParamDetail struct {
	HostUuid *string `json:"hostUuid,omitempty"`
	ExcludePackages []string `json:"excludePackages,omitempty"`
	UpdatePackages []string `json:"updatePackages,omitempty"`
	ReleaseVersion *string `json:"releaseVersion,omitempty"`
	Force *bool `json:"force,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateClusterOSParam UpdateClusterOS request param
type UpdateClusterOSParam struct {
	BaseParam
	Params UpdateClusterOSParamDetail `json:"updateClusterOS"`
}
// GetVmUsbRedirectParamDetail GetVmUsbRedirect detail param
type GetVmUsbRedirectParamDetail struct {
}

// GetVmUsbRedirectParam GetVmUsbRedirect request param
type GetVmUsbRedirectParam struct {
	BaseParam
	Params GetVmUsbRedirectParamDetail `json:"getVmUsbRedirect"`
}
// CreateImageGroupFromSnapshotParamDetail CreateImageGroupFromSnapshot detail param
type CreateImageGroupFromSnapshotParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	DataVolumeSnapshotUuids []string `json:"dataVolumeSnapshotUuids,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateImageGroupFromSnapshotParam CreateImageGroupFromSnapshot request param
type CreateImageGroupFromSnapshotParam struct {
	BaseParam
	Params CreateImageGroupFromSnapshotParamDetail `json:"params"`
}
// GetOssBucketFileFromRemoteParamDetail GetOssBucketFileFromRemote detail param
type GetOssBucketFileFromRemoteParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	OssDomain *string `json:"ossDomain,omitempty"`
	OssKey *string `json:"ossKey,omitempty"`
	OssSecret *string `json:"ossSecret,omitempty"`
}

// GetOssBucketFileFromRemoteParam GetOssBucketFileFromRemote request param
type GetOssBucketFileFromRemoteParam struct {
	BaseParam
	Params GetOssBucketFileFromRemoteParamDetail `json:"getOssBucketFileFromRemote"`
}
// AttachVipToVpcSharedQosParamDetail AttachVipToVpcSharedQos detail param
type AttachVipToVpcSharedQosParamDetail struct {
	VipLists []string `json:"vipLists,omitempty"`
	VipUuids []string `json:"vipUuids,omitempty"`
}

// AttachVipToVpcSharedQosParam AttachVipToVpcSharedQos request param
type AttachVipToVpcSharedQosParam struct {
	BaseParam
	Params AttachVipToVpcSharedQosParamDetail `json:"params"`
}
// GetEventDataParamDetail GetEventData detail param
type GetEventDataParamDetail struct {
	StartTime *int64 `json:"startTime,omitempty"`
	EndTime *int64 `json:"endTime,omitempty"`
	OffsetAheadOfCurrentTime *int64 `json:"offsetAheadOfCurrentTime,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Conditions []string `json:"conditions,omitempty"`
	Count *bool `json:"count,omitempty"`
	Start *int `json:"start,omitempty"`
	ConditionExpression *string `json:"conditionExpression,omitempty"`
	EndpointUuid *string `json:"endpointUuid,omitempty"`
}

// GetEventDataParam GetEventData request param
type GetEventDataParam struct {
	BaseParam
	Params GetEventDataParamDetail `json:"getEventData"`
}
// CheckIpAvailabilityParamDetail CheckIpAvailability detail param
type CheckIpAvailabilityParamDetail struct {
	ArpCheck *bool `json:"arpCheck,omitempty"`
	IpRangeCheck *bool `json:"ipRangeCheck,omitempty"`
}

// CheckIpAvailabilityParam CheckIpAvailability request param
type CheckIpAvailabilityParam struct {
	BaseParam
	Params CheckIpAvailabilityParamDetail `json:"checkIpAvailability"`
}
// RemoveVmNicFromLoadBalancerParamDetail RemoveVmNicFromLoadBalancer detail param
type RemoveVmNicFromLoadBalancerParamDetail struct {
	VmNicUuids []string `json:"vmNicUuids" validate:"required"`
}

// RemoveVmNicFromLoadBalancerParam RemoveVmNicFromLoadBalancer request param
type RemoveVmNicFromLoadBalancerParam struct {
	BaseParam
	Params RemoveVmNicFromLoadBalancerParamDetail `json:"removeVmNicFromLoadBalancer"`
}
// RemoveRolesFromIAM2VirtualIDParamDetail RemoveRolesFromIAM2VirtualID detail param
type RemoveRolesFromIAM2VirtualIDParamDetail struct {
	RoleUuids []string `json:"roleUuids" validate:"required"`
	ProjectUuid *string `json:"projectUuid,omitempty"`
}

// RemoveRolesFromIAM2VirtualIDParam RemoveRolesFromIAM2VirtualID request param
type RemoveRolesFromIAM2VirtualIDParam struct {
	BaseParam
	Params RemoveRolesFromIAM2VirtualIDParamDetail `json:"removeRolesFromIAM2VirtualID"`
}
// CalculateResourceSpendingParamDetail CalculateResourceSpending detail param
type CalculateResourceSpendingParamDetail struct {
	ResourceType *string `json:"resourceType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	DateStart *string `json:"dateStart,omitempty"`
	DateEnd *string `json:"dateEnd,omitempty"`
	Start *int `json:"start,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// CalculateResourceSpendingParam CalculateResourceSpending request param
type CalculateResourceSpendingParam struct {
	BaseParam
	Params CalculateResourceSpendingParamDetail `json:"calculateResourceSpending"`
}
// DetachBackupStorageFromZoneParamDetail DetachBackupStorageFromZone detail param
type DetachBackupStorageFromZoneParamDetail struct {
}

// DetachBackupStorageFromZoneParam DetachBackupStorageFromZone request param
type DetachBackupStorageFromZoneParam struct {
	BaseParam
	Params DetachBackupStorageFromZoneParamDetail `json:"detachBackupStorageFromZone"`
}
// UpdateCCSCertificateUserStateParamDetail UpdateCCSCertificateUserState detail param
type UpdateCCSCertificateUserStateParamDetail struct {
	State string `json:"state" validate:"required"`
}

// UpdateCCSCertificateUserStateParam UpdateCCSCertificateUserState request param
type UpdateCCSCertificateUserStateParam struct {
	BaseParam
	Params UpdateCCSCertificateUserStateParamDetail `json:"params"`
}
// PowerResetBaremetalChassisParamDetail PowerResetBaremetalChassis detail param
type PowerResetBaremetalChassisParamDetail struct {
}

// PowerResetBaremetalChassisParam PowerResetBaremetalChassis request param
type PowerResetBaremetalChassisParam struct {
	BaseParam
	Params PowerResetBaremetalChassisParamDetail `json:"powerResetBaremetalChassis"`
}
// CleanUpTrashOnPrimaryStorageParamDetail CleanUpTrashOnPrimaryStorage detail param
type CleanUpTrashOnPrimaryStorageParamDetail struct {
	TrashId *int64 `json:"trashId,omitempty"`
}

// CleanUpTrashOnPrimaryStorageParam CleanUpTrashOnPrimaryStorage request param
type CleanUpTrashOnPrimaryStorageParam struct {
	BaseParam
	Params CleanUpTrashOnPrimaryStorageParamDetail `json:"cleanUpTrashOnPrimaryStorage"`
}
// AddDisasterImageStoreBackupStorageParamDetail AddDisasterImageStoreBackupStorage detail param
type AddDisasterImageStoreBackupStorageParamDetail struct {
	AttachPoint *string `json:"attachPoint,omitempty"`
	EndPoint *string `json:"endPoint,omitempty"`
	Hostname string `json:"hostname" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	SshPort *int `json:"sshPort,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ImportImages *bool `json:"importImages,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddDisasterImageStoreBackupStorageParam AddDisasterImageStoreBackupStorage request param
type AddDisasterImageStoreBackupStorageParam struct {
	BaseParam
	Params AddDisasterImageStoreBackupStorageParamDetail `json:"params"`
}
// GetVmSchedulingRulesExecuteStateParamDetail GetVmSchedulingRulesExecuteState detail param
type GetVmSchedulingRulesExecuteStateParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// GetVmSchedulingRulesExecuteStateParam GetVmSchedulingRulesExecuteState request param
type GetVmSchedulingRulesExecuteStateParam struct {
	BaseParam
	Params GetVmSchedulingRulesExecuteStateParamDetail `json:"params"`
}
// CreateVolumesSnapshotParamDetail CreateVolumesSnapshot detail param
type CreateVolumesSnapshotParamDetail struct {
	VolumeUuids []string `json:"volumeUuids" validate:"required"`
}

// CreateVolumesSnapshotParam CreateVolumesSnapshot request param
type CreateVolumesSnapshotParam struct {
	BaseParam
	Params CreateVolumesSnapshotParamDetail `json:"params"`
}
// GetIpAddressCapacityParamDetail GetIpAddressCapacity detail param
type GetIpAddressCapacityParamDetail struct {
	ZoneUuids []string `json:"zoneUuids,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	IpRangeUuids []string `json:"ipRangeUuids,omitempty"`
	All bool `json:"all,omitempty"`
}

// GetIpAddressCapacityParam GetIpAddressCapacity request param
type GetIpAddressCapacityParam struct {
	BaseParam
	Params GetIpAddressCapacityParamDetail `json:"getIpAddressCapacity"`
}
// SetIAM2ProjectContainerClusterParamDetail SetIAM2ProjectContainerCluster detail param
type SetIAM2ProjectContainerClusterParamDetail struct {
	ContainerUuid string `json:"containerUuid" validate:"required"`
	ClusterId int64 `json:"clusterId" validate:"required"`
}

// SetIAM2ProjectContainerClusterParam SetIAM2ProjectContainerCluster request param
type SetIAM2ProjectContainerClusterParam struct {
	BaseParam
	Params SetIAM2ProjectContainerClusterParamDetail `json:"setIAM2ProjectContainerCluster"`
}
// DeployAppDevelopmentServiceParamDetail DeployAppDevelopmentService detail param
type DeployAppDevelopmentServiceParamDetail struct {
	Uuid string `json:"uuid,omitempty"`
	Description *string `json:"description,omitempty"`
	ModelUuid *string `json:"modelUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	VmImageUuid *string `json:"vmImageUuid,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	DatasetUuids []string `json:"datasetUuids,omitempty"`
	ModelServiceGroupUuids []string `json:"modelServiceGroupUuids,omitempty"`
	DockerImage *string `json:"dockerImage,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	Name string `json:"name" validate:"required"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	StartupParameters map[string]string `json:"startupParameters,omitempty"`
	Type string `json:"type" validate:"required"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	ServiceBootUptime *int `json:"serviceBootUptime,omitempty"`
	ServiceLivez *string `json:"serviceLivez,omitempty"`
	ServiceReadyz *string `json:"serviceReadyz,omitempty"`
	RootDiskOfferingUuid *string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize *int64 `json:"rootDiskSize,omitempty"`
	ProjectUuid *string `json:"projectUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// DeployAppDevelopmentServiceParam DeployAppDevelopmentService request param
type DeployAppDevelopmentServiceParam struct {
	BaseParam
	Params DeployAppDevelopmentServiceParamDetail `json:"deployAppDevelopmentService"`
}
// RefreshPluginDriversParamDetail RefreshPluginDrivers detail param
type RefreshPluginDriversParamDetail struct {
	Name string `json:"name,omitempty"`
}

// RefreshPluginDriversParam RefreshPluginDrivers request param
type RefreshPluginDriversParam struct {
	BaseParam
	Params RefreshPluginDriversParamDetail `json:"refreshPluginDrivers"`
}
// PauseVmInstanceParamDetail PauseVmInstance detail param
type PauseVmInstanceParamDetail struct {
}

// PauseVmInstanceParam PauseVmInstance request param
type PauseVmInstanceParam struct {
	BaseParam
	Params PauseVmInstanceParamDetail `json:"pauseVmInstance"`
}
// DetachUserDefinedXmlHookScriptFromVmParamDetail DetachUserDefinedXmlHookScriptFromVm detail param
type DetachUserDefinedXmlHookScriptFromVmParamDetail struct {
	StartupStrategy *string `json:"startupStrategy,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DetachUserDefinedXmlHookScriptFromVmParam DetachUserDefinedXmlHookScriptFromVm request param
type DetachUserDefinedXmlHookScriptFromVmParam struct {
	BaseParam
	Params DetachUserDefinedXmlHookScriptFromVmParamDetail `json:"detachUserDefinedXmlHookScriptFromVm"`
}
// GetSignatureServerEncryptPublicKeyParamDetail GetSignatureServerEncryptPublicKey detail param
type GetSignatureServerEncryptPublicKeyParamDetail struct {
}

// GetSignatureServerEncryptPublicKeyParam GetSignatureServerEncryptPublicKey request param
type GetSignatureServerEncryptPublicKeyParam struct {
	BaseParam
	Params GetSignatureServerEncryptPublicKeyParamDetail `json:"getSignatureServerEncryptPublicKey"`
}
// AddAliyunKeySecretParamDetail AddAliyunKeySecret detail param
type AddAliyunKeySecretParamDetail struct {
	Name string `json:"name" validate:"required"`
	Key string `json:"key" validate:"required"`
	Secret string `json:"secret" validate:"required"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	Sync *bool `json:"sync,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunKeySecretParam AddAliyunKeySecret request param
type AddAliyunKeySecretParam struct {
	BaseParam
	Params AddAliyunKeySecretParamDetail `json:"params"`
}
// AddBackupStoragesToReplicationGroupParamDetail AddBackupStoragesToReplicationGroup detail param
type AddBackupStoragesToReplicationGroupParamDetail struct {
	BackupStorageUuids []string `json:"backupStorageUuids" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBackupStoragesToReplicationGroupParam AddBackupStoragesToReplicationGroup request param
type AddBackupStoragesToReplicationGroupParam struct {
	BaseParam
	Params AddBackupStoragesToReplicationGroupParamDetail `json:"params"`
}
// AddDataCenterFromRemoteParamDetail AddDataCenterFromRemote detail param
type AddDataCenterFromRemoteParamDetail struct {
	RegionId string `json:"regionId" validate:"required"`
	Type string `json:"type" validate:"required"`
	SyncZones *bool `json:"syncZones,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddDataCenterFromRemoteParam AddDataCenterFromRemote request param
type AddDataCenterFromRemoteParam struct {
	BaseParam
	Params AddDataCenterFromRemoteParamDetail `json:"params"`
}
// DeleteFirewallRuleSetParamDetail DeleteFirewallRuleSet detail param
type DeleteFirewallRuleSetParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteFirewallRuleSetParam DeleteFirewallRuleSet request param
type DeleteFirewallRuleSetParam struct {
	BaseParam
	Params DeleteFirewallRuleSetParamDetail `json:"deleteFirewallRuleSet"`
}
// BatchAddBareMetal2IpmiChassisParamDetail BatchAddBareMetal2IpmiChassis detail param
type BatchAddBareMetal2IpmiChassisParamDetail struct {
	ChassisInfo string `json:"chassisInfo" validate:"required"`
	LongJobName *string `json:"longJobName,omitempty"`
	LongJobDescription *string `json:"longJobDescription,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// BatchAddBareMetal2IpmiChassisParam BatchAddBareMetal2IpmiChassis request param
type BatchAddBareMetal2IpmiChassisParam struct {
	BaseParam
	Params BatchAddBareMetal2IpmiChassisParamDetail `json:"params"`
}
// LocalStorageMigrateVolumeParamDetail LocalStorageMigrateVolume detail param
type LocalStorageMigrateVolumeParamDetail struct {
	DestHostUuid string `json:"destHostUuid" validate:"required"`
}

// LocalStorageMigrateVolumeParam LocalStorageMigrateVolume request param
type LocalStorageMigrateVolumeParam struct {
	BaseParam
	Params LocalStorageMigrateVolumeParamDetail `json:"localStorageMigrateVolume"`
}
// AttachNicToBondingParamDetail AttachNicToBonding detail param
type AttachNicToBondingParamDetail struct {
	SlaveUuids []string `json:"slaveUuids" validate:"required"`
	Type *string `json:"type,omitempty"`
}

// AttachNicToBondingParam AttachNicToBonding request param
type AttachNicToBondingParam struct {
	BaseParam
	Params AttachNicToBondingParamDetail `json:"attachNicToBonding"`
}
// SetOrganizationOperationParamDetail SetOrganizationOperation detail param
type SetOrganizationOperationParamDetail struct {
	VirtualIDUuid string `json:"virtualIDUuid" validate:"required"`
}

// SetOrganizationOperationParam SetOrganizationOperation request param
type SetOrganizationOperationParam struct {
	BaseParam
	Params SetOrganizationOperationParamDetail `json:"setOrganizationOperation"`
}
// CreateDataVolumeTemplateFromVolumeParamDetail CreateDataVolumeTemplateFromVolume detail param
type CreateDataVolumeTemplateFromVolumeParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeTemplateFromVolumeParam CreateDataVolumeTemplateFromVolume request param
type CreateDataVolumeTemplateFromVolumeParam struct {
	BaseParam
	Params CreateDataVolumeTemplateFromVolumeParamDetail `json:"params"`
}
// RemoveIAM2VirtualIDsFromOrganizationParamDetail RemoveIAM2VirtualIDsFromOrganization detail param
type RemoveIAM2VirtualIDsFromOrganizationParamDetail struct {
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
}

// RemoveIAM2VirtualIDsFromOrganizationParam RemoveIAM2VirtualIDsFromOrganization request param
type RemoveIAM2VirtualIDsFromOrganizationParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromOrganizationParamDetail `json:"removeIAM2VirtualIDsFromOrganization"`
}
// ExportDatabaseBackupFromBackupStorageParamDetail ExportDatabaseBackupFromBackupStorage detail param
type ExportDatabaseBackupFromBackupStorageParamDetail struct {
}

// ExportDatabaseBackupFromBackupStorageParam ExportDatabaseBackupFromBackupStorage request param
type ExportDatabaseBackupFromBackupStorageParam struct {
	BaseParam
	Params ExportDatabaseBackupFromBackupStorageParamDetail `json:"exportDatabaseBackupFromBackupStorage"`
}
// AttachIAM2ProjectToIAM2OrganizationParamDetail AttachIAM2ProjectToIAM2Organization detail param
type AttachIAM2ProjectToIAM2OrganizationParamDetail struct {
}

// AttachIAM2ProjectToIAM2OrganizationParam AttachIAM2ProjectToIAM2Organization request param
type AttachIAM2ProjectToIAM2OrganizationParam struct {
	BaseParam
	Params AttachIAM2ProjectToIAM2OrganizationParamDetail `json:"params"`
}
// CreateEmailMonitorTriggerActionParamDetail CreateEmailMonitorTriggerAction detail param
type CreateEmailMonitorTriggerActionParamDetail struct {
	Email string `json:"email" validate:"required"`
	MediaUuid string `json:"mediaUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	TriggerUuids []string `json:"triggerUuids,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEmailMonitorTriggerActionParam CreateEmailMonitorTriggerAction request param
type CreateEmailMonitorTriggerActionParam struct {
	BaseParam
	Params CreateEmailMonitorTriggerActionParamDetail `json:"params"`
}
// SetVpcVRouterDistributedRoutingEnabledParamDetail SetVpcVRouterDistributedRoutingEnabled detail param
type SetVpcVRouterDistributedRoutingEnabledParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// SetVpcVRouterDistributedRoutingEnabledParam SetVpcVRouterDistributedRoutingEnabled request param
type SetVpcVRouterDistributedRoutingEnabledParam struct {
	BaseParam
	Params SetVpcVRouterDistributedRoutingEnabledParamDetail `json:"params"`
}
// PowerOnBareMetal2ChassisParamDetail PowerOnBareMetal2Chassis detail param
type PowerOnBareMetal2ChassisParamDetail struct {
	BootDev *string `json:"bootDev,omitempty"`
}

// PowerOnBareMetal2ChassisParam PowerOnBareMetal2Chassis request param
type PowerOnBareMetal2ChassisParam struct {
	BaseParam
	Params PowerOnBareMetal2ChassisParamDetail `json:"powerOnBareMetal2Chassis"`
}
// GetLocalRaidPhysicalDriveSmartParamDetail GetLocalRaidPhysicalDriveSmart detail param
type GetLocalRaidPhysicalDriveSmartParamDetail struct {
}

// GetLocalRaidPhysicalDriveSmartParam GetLocalRaidPhysicalDriveSmart request param
type GetLocalRaidPhysicalDriveSmartParam struct {
	BaseParam
	Params GetLocalRaidPhysicalDriveSmartParamDetail `json:"getLocalRaidPhysicalDriveSmart"`
}
// UpdateHybridKeySecretParamDetail UpdateHybridKeySecret detail param
type UpdateHybridKeySecretParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateHybridKeySecretParam UpdateHybridKeySecret request param
type UpdateHybridKeySecretParam struct {
	BaseParam
	Params UpdateHybridKeySecretParamDetail `json:"params"`
}
// PullHuaweiIMasterControllerParamDetail PullHuaweiIMasterController detail param
type PullHuaweiIMasterControllerParamDetail struct {
	PullSwitch *bool `json:"pullSwitch,omitempty"`
}

// PullHuaweiIMasterControllerParam PullHuaweiIMasterController request param
type PullHuaweiIMasterControllerParam struct {
	BaseParam
	Params PullHuaweiIMasterControllerParamDetail `json:"pullHuaweiIMasterController"`
}
// RemoveRolesFromIAM2VirtualIDGroupParamDetail RemoveRolesFromIAM2VirtualIDGroup detail param
type RemoveRolesFromIAM2VirtualIDGroupParamDetail struct {
	RoleUuids []string `json:"roleUuids" validate:"required"`
	ProjectUuid *string `json:"projectUuid,omitempty"`
}

// RemoveRolesFromIAM2VirtualIDGroupParam RemoveRolesFromIAM2VirtualIDGroup request param
type RemoveRolesFromIAM2VirtualIDGroupParam struct {
	BaseParam
	Params RemoveRolesFromIAM2VirtualIDGroupParamDetail `json:"removeRolesFromIAM2VirtualIDGroup"`
}
// AckAlarmDataParamDetail AckAlarmData detail param
type AckAlarmDataParamDetail struct {
	AlarmUuid string `json:"alarmUuid" validate:"required"`
	AlertDataUuid string `json:"alertDataUuid" validate:"required"`
	DataType string `json:"dataType" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	AckPeriodSec int `json:"ackPeriodSec" validate:"required"`
}

// AckAlarmDataParam AckAlarmData request param
type AckAlarmDataParam struct {
	BaseParam
	Params AckAlarmDataParamDetail `json:"params"`
}
// RemoveDnsFromL3NetworkParamDetail RemoveDnsFromL3Network detail param
type RemoveDnsFromL3NetworkParamDetail struct {
}

// RemoveDnsFromL3NetworkParam RemoveDnsFromL3Network request param
type RemoveDnsFromL3NetworkParam struct {
	BaseParam
	Params RemoveDnsFromL3NetworkParamDetail `json:"removeDnsFromL3Network"`
}
// ChangeIAM2OrganizationParentParamDetail ChangeIAM2OrganizationParent detail param
type ChangeIAM2OrganizationParentParamDetail struct {
	ChildrenUuids []string `json:"childrenUuids" validate:"required"`
}

// ChangeIAM2OrganizationParentParam ChangeIAM2OrganizationParent request param
type ChangeIAM2OrganizationParentParam struct {
	BaseParam
	Params ChangeIAM2OrganizationParentParamDetail `json:"changeIAM2OrganizationParent"`
}
// SNSWeComTestConnectionParamDetail SNSWeComTestConnection detail param
type SNSWeComTestConnectionParamDetail struct {
	Url *string `json:"url,omitempty"`
	AtAll *bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	TestMsg string `json:"testMsg" validate:"required"`
	EndpointUuid *string `json:"endpointUuid,omitempty"`
}

// SNSWeComTestConnectionParam SNSWeComTestConnection request param
type SNSWeComTestConnectionParam struct {
	BaseParam
	Params SNSWeComTestConnectionParamDetail `json:"params"`
}
// ProvisionVirtualRouterConfigParamDetail ProvisionVirtualRouterConfig detail param
type ProvisionVirtualRouterConfigParamDetail struct {
}

// ProvisionVirtualRouterConfigParam ProvisionVirtualRouterConfig request param
type ProvisionVirtualRouterConfigParam struct {
	BaseParam
	Params ProvisionVirtualRouterConfigParamDetail `json:"provisionVirtualRouterConfig"`
}
// SetVmQgaParamDetail SetVmQga detail param
type SetVmQgaParamDetail struct {
	Enable bool `json:"enable" validate:"required"`
}

// SetVmQgaParam SetVmQga request param
type SetVmQgaParam struct {
	BaseParam
	Params SetVmQgaParamDetail `json:"setVmQga"`
}
// ValidatePasswordParamDetail ValidatePassword detail param
type ValidatePasswordParamDetail struct {
	LoginName string `json:"loginName" validate:"required"`
	Password string `json:"password" validate:"required"`
	LoginType string `json:"loginType" validate:"required"`
}

// ValidatePasswordParam ValidatePassword request param
type ValidatePasswordParam struct {
	BaseParam
	Params ValidatePasswordParamDetail `json:"validatePassword"`
}
// GetChronyServersParamDetail GetChronyServers detail param
type GetChronyServersParamDetail struct {
}

// GetChronyServersParam GetChronyServers request param
type GetChronyServersParam struct {
	BaseParam
	Params GetChronyServersParamDetail `json:"getChronyServers"`
}
// AttachL3NetworkToVmNicParamDetail AttachL3NetworkToVmNic detail param
type AttachL3NetworkToVmNicParamDetail struct {
	StaticIp *string `json:"staticIp,omitempty"`
}

// AttachL3NetworkToVmNicParam AttachL3NetworkToVmNic request param
type AttachL3NetworkToVmNicParam struct {
	BaseParam
	Params AttachL3NetworkToVmNicParamDetail `json:"params"`
}
// ChangeSecurityMachineStateParamDetail ChangeSecurityMachineState detail param
type ChangeSecurityMachineStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSecurityMachineStateParam ChangeSecurityMachineState request param
type ChangeSecurityMachineStateParam struct {
	BaseParam
	Params ChangeSecurityMachineStateParamDetail `json:"changeSecurityMachineState"`
}
// SetVmQxlMemoryParamDetail SetVmQxlMemory detail param
type SetVmQxlMemoryParamDetail struct {
	Ram *int `json:"ram,omitempty"`
	Vram *int `json:"vram,omitempty"`
	Vgamem *int `json:"vgamem,omitempty"`
}

// SetVmQxlMemoryParam SetVmQxlMemory request param
type SetVmQxlMemoryParam struct {
	BaseParam
	Params SetVmQxlMemoryParamDetail `json:"setVmQxlMemory"`
}
// GetVolumeFormatParamDetail GetVolumeFormat detail param
type GetVolumeFormatParamDetail struct {
}

// GetVolumeFormatParam GetVolumeFormat request param
type GetVolumeFormatParam struct {
	BaseParam
	Params GetVolumeFormatParamDetail `json:"getVolumeFormat"`
}
// AddLocalPrimaryStorageParamDetail AddLocalPrimaryStorage detail param
type AddLocalPrimaryStorageParamDetail struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddLocalPrimaryStorageParam AddLocalPrimaryStorage request param
type AddLocalPrimaryStorageParam struct {
	BaseParam
	Params AddLocalPrimaryStorageParamDetail `json:"params"`
}
// UpdateAtPersonOfAtDingTalkEndpointParamDetail UpdateAtPersonOfAtDingTalkEndpoint detail param
type UpdateAtPersonOfAtDingTalkEndpointParamDetail struct {
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Remark *string `json:"remark,omitempty"`
}

// UpdateAtPersonOfAtDingTalkEndpointParam UpdateAtPersonOfAtDingTalkEndpoint request param
type UpdateAtPersonOfAtDingTalkEndpointParam struct {
	BaseParam
	Params UpdateAtPersonOfAtDingTalkEndpointParamDetail `json:"updateAtPersonOfAtDingTalkEndpoint"`
}
// UpdateAliyunMountTargetParamDetail UpdateAliyunMountTarget detail param
type UpdateAliyunMountTargetParamDetail struct {
	AccessGroupUuid string `json:"accessGroupUuid" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateAliyunMountTargetParam UpdateAliyunMountTarget request param
type UpdateAliyunMountTargetParam struct {
	BaseParam
	Params UpdateAliyunMountTargetParamDetail `json:"updateAliyunMountTarget"`
}
// GetResourceAccountParamDetail GetResourceAccount detail param
type GetResourceAccountParamDetail struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
}

// GetResourceAccountParam GetResourceAccount request param
type GetResourceAccountParam struct {
	BaseParam
	Params GetResourceAccountParamDetail `json:"getResourceAccount"`
}
// ChangeSecretResourcePoolStateParamDetail ChangeSecretResourcePoolState detail param
type ChangeSecretResourcePoolStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSecretResourcePoolStateParam ChangeSecretResourcePoolState request param
type ChangeSecretResourcePoolStateParam struct {
	BaseParam
	Params ChangeSecretResourcePoolStateParamDetail `json:"changeSecretResourcePoolState"`
}
// AddSimulatorBackupStorageParamDetail AddSimulatorBackupStorage detail param
type AddSimulatorBackupStorageParamDetail struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ImportImages *bool `json:"importImages,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSimulatorBackupStorageParam AddSimulatorBackupStorage request param
type AddSimulatorBackupStorageParam struct {
	BaseParam
	Params AddSimulatorBackupStorageParamDetail `json:"params"`
}
// BindModelToServiceParamDetail BindModelToService detail param
type BindModelToServiceParamDetail struct {
}

// BindModelToServiceParam BindModelToService request param
type BindModelToServiceParam struct {
	BaseParam
	Params BindModelToServiceParamDetail `json:"param"`
}
// GetCandidateAffinityGroupForCreatingVmParamDetail GetCandidateAffinityGroupForCreatingVm detail param
type GetCandidateAffinityGroupForCreatingVmParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
}

// GetCandidateAffinityGroupForCreatingVmParam GetCandidateAffinityGroupForCreatingVm request param
type GetCandidateAffinityGroupForCreatingVmParam struct {
	BaseParam
	Params GetCandidateAffinityGroupForCreatingVmParamDetail `json:"getCandidateAffinityGroupForCreatingVm"`
}
// CheckNetworkReachableParamDetail CheckNetworkReachable detail param
type CheckNetworkReachableParamDetail struct {
	SourceHostnames []string `json:"sourceHostnames,omitempty"`
	TargetHostnames []string `json:"targetHostnames" validate:"required"`
}

// CheckNetworkReachableParam CheckNetworkReachable request param
type CheckNetworkReachableParam struct {
	BaseParam
	Params CheckNetworkReachableParamDetail `json:"checkNetworkReachable"`
}
// SetFlowMeterRouterIdParamDetail SetFlowMeterRouterId detail param
type SetFlowMeterRouterIdParamDetail struct {
	RouterId int64 `json:"routerId" validate:"required"`
}

// SetFlowMeterRouterIdParam SetFlowMeterRouterId request param
type SetFlowMeterRouterIdParam struct {
	BaseParam
	Params SetFlowMeterRouterIdParamDetail `json:"params"`
}
// AddStorageProtocolParamDetail AddStorageProtocol detail param
type AddStorageProtocolParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	OutputProtocol string `json:"outputProtocol" validate:"required"`
}

// AddStorageProtocolParam AddStorageProtocol request param
type AddStorageProtocolParam struct {
	BaseParam
	Params AddStorageProtocolParamDetail `json:"params"`
}
// DeployModelServiceParamDetail DeployModelService detail param
type DeployModelServiceParamDetail struct {
	Description *string `json:"description,omitempty"`
	ModelUuid *string `json:"modelUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	VmImageUuid *string `json:"vmImageUuid,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	DatasetUuids []string `json:"datasetUuids,omitempty"`
	ModelServiceGroupUuids []string `json:"modelServiceGroupUuids,omitempty"`
	DockerImage *string `json:"dockerImage,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	Name string `json:"name" validate:"required"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	StartupParameters map[string]string `json:"startupParameters,omitempty"`
	Type string `json:"type" validate:"required"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	ServiceBootUptime *int `json:"serviceBootUptime,omitempty"`
	ServiceLivez *string `json:"serviceLivez,omitempty"`
	ServiceReadyz *string `json:"serviceReadyz,omitempty"`
	RootDiskOfferingUuid *string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize *int64 `json:"rootDiskSize,omitempty"`
	ProjectUuid *string `json:"projectUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// DeployModelServiceParam DeployModelService request param
type DeployModelServiceParam struct {
	BaseParam
	Params DeployModelServiceParamDetail `json:"deployModelService"`
}
// GetMonitorItemParamDetail GetMonitorItem detail param
type GetMonitorItemParamDetail struct {
	ResourceType string `json:"resourceType" validate:"required"`
}

// GetMonitorItemParam GetMonitorItem request param
type GetMonitorItemParam struct {
	BaseParam
	Params GetMonitorItemParamDetail `json:"getMonitorItem"`
}
// GetLicenseRecordsParamDetail GetLicenseRecords detail param
type GetLicenseRecordsParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
	ReplyWithCount *bool `json:"replyWithCount,omitempty"`
	Count *bool `json:"count,omitempty"`
	SortBy *string `json:"sortBy,omitempty"`
	SortDirection *string `json:"sortDirection,omitempty"`
}

// GetLicenseRecordsParam GetLicenseRecords request param
type GetLicenseRecordsParam struct {
	BaseParam
	Params GetLicenseRecordsParamDetail `json:"getLicenseRecords"`
}
// UnregisterLicenseRequestedApplicationParamDetail UnregisterLicenseRequestedApplication detail param
type UnregisterLicenseRequestedApplicationParamDetail struct {
	AppId string `json:"appId" validate:"required"`
}

// UnregisterLicenseRequestedApplicationParam UnregisterLicenseRequestedApplication request param
type UnregisterLicenseRequestedApplicationParam struct {
	BaseParam
	Params UnregisterLicenseRequestedApplicationParamDetail `json:"unregisterLicenseRequestedApplication"`
}
// AttachSecurityGroupToL3NetworkParamDetail AttachSecurityGroupToL3Network detail param
type AttachSecurityGroupToL3NetworkParamDetail struct {
}

// AttachSecurityGroupToL3NetworkParam AttachSecurityGroupToL3Network request param
type AttachSecurityGroupToL3NetworkParam struct {
	BaseParam
	Params AttachSecurityGroupToL3NetworkParamDetail `json:"params"`
}
// UpdateVmNicDriverParamDetail UpdateVmNicDriver detail param
type UpdateVmNicDriverParamDetail struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	DriverType string `json:"driverType" validate:"required"`
}

// UpdateVmNicDriverParam UpdateVmNicDriver request param
type UpdateVmNicDriverParam struct {
	BaseParam
	Params UpdateVmNicDriverParamDetail `json:"updateVmNicDriver"`
}
// SetIpOnHostNetworkInterfaceParamDetail SetIpOnHostNetworkInterface detail param
type SetIpOnHostNetworkInterfaceParamDetail struct {
	IpAddress *string `json:"ipAddress,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
}

// SetIpOnHostNetworkInterfaceParam SetIpOnHostNetworkInterface request param
type SetIpOnHostNetworkInterfaceParam struct {
	BaseParam
	Params SetIpOnHostNetworkInterfaceParamDetail `json:"params"`
}
// ProvisionNfvInstGroupParamDetail ProvisionNfvInstGroup detail param
type ProvisionNfvInstGroupParamDetail struct {
}

// ProvisionNfvInstGroupParam ProvisionNfvInstGroup request param
type ProvisionNfvInstGroupParam struct {
	BaseParam
	Params ProvisionNfvInstGroupParamDetail `json:"provisionNfvInstGroup"`
}
// DetachNicFromBondingParamDetail DetachNicFromBonding detail param
type DetachNicFromBondingParamDetail struct {
	SlaveUuids []string `json:"slaveUuids" validate:"required"`
	Type *string `json:"type,omitempty"`
}

// DetachNicFromBondingParam DetachNicFromBonding request param
type DetachNicFromBondingParam struct {
	BaseParam
	Params DetachNicFromBondingParamDetail `json:"detachNicFromBonding"`
}
// ChangeMonitorTriggerActionStateParamDetail ChangeMonitorTriggerActionState detail param
type ChangeMonitorTriggerActionStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeMonitorTriggerActionStateParam ChangeMonitorTriggerActionState request param
type ChangeMonitorTriggerActionStateParam struct {
	BaseParam
	Params ChangeMonitorTriggerActionStateParamDetail `json:"changeMonitorTriggerActionState"`
}
// MigrateVmParamDetail MigrateVm detail param
type MigrateVmParamDetail struct {
	HostUuid *string `json:"hostUuid,omitempty"`
	MigrateFromDestination *bool `json:"migrateFromDestination,omitempty"`
	AllowUnknown *bool `json:"allowUnknown,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	DownTime *int `json:"downTime,omitempty"`
}

// MigrateVmParam MigrateVm request param
type MigrateVmParam struct {
	BaseParam
	Params MigrateVmParamDetail `json:"migrateVm"`
}
// ChangeVmPasswordParamDetail ChangeVmPassword detail param
type ChangeVmPasswordParamDetail struct {
	Password string `json:"password" validate:"required"`
	Account string `json:"account" validate:"required"`
}

// ChangeVmPasswordParam ChangeVmPassword request param
type ChangeVmPasswordParam struct {
	BaseParam
	Params ChangeVmPasswordParamDetail `json:"changeVmPassword"`
}
// FlattenVmInstanceParamDetail FlattenVmInstance detail param
type FlattenVmInstanceParamDetail struct {
	Full *bool `json:"full,omitempty"`
	DryRun *bool `json:"dryRun,omitempty"`
}

// FlattenVmInstanceParam FlattenVmInstance request param
type FlattenVmInstanceParam struct {
	BaseParam
	Params FlattenVmInstanceParamDetail `json:"flattenVmInstance"`
}
// DeleteAllEcsInstancesFromDataCenterParamDetail DeleteAllEcsInstancesFromDataCenter detail param
type DeleteAllEcsInstancesFromDataCenterParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAllEcsInstancesFromDataCenterParam DeleteAllEcsInstancesFromDataCenter request param
type DeleteAllEcsInstancesFromDataCenterParam struct {
	BaseParam
	Params DeleteAllEcsInstancesFromDataCenterParamDetail `json:"deleteAllEcsInstancesFromDataCenter"`
}
// GetVpcMulticastRouteParamDetail GetVpcMulticastRoute detail param
type GetVpcMulticastRouteParamDetail struct {
}

// GetVpcMulticastRouteParam GetVpcMulticastRoute request param
type GetVpcMulticastRouteParam struct {
	BaseParam
	Params GetVpcMulticastRouteParamDetail `json:"getVpcMulticastRoute"`
}
// DeleteVmUserDefinedXmlHookScriptParamDetail DeleteVmUserDefinedXmlHookScript detail param
type DeleteVmUserDefinedXmlHookScriptParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVmUserDefinedXmlHookScriptParam DeleteVmUserDefinedXmlHookScript request param
type DeleteVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params DeleteVmUserDefinedXmlHookScriptParamDetail `json:"deleteVmUserDefinedXmlHookScript"`
}
// AddL3NetworkToGroupParamDetail AddL3NetworkToGroup detail param
type AddL3NetworkToGroupParamDetail struct {
	FrontEndL3NetworkUuid string `json:"frontEndL3NetworkUuid" validate:"required"`
	BackendL3NetworkUuids []string `json:"backendL3NetworkUuids" validate:"required"`
}

// AddL3NetworkToGroupParam AddL3NetworkToGroup request param
type AddL3NetworkToGroupParam struct {
	BaseParam
	Params AddL3NetworkToGroupParamDetail `json:"params"`
}
// SyncZBoxCapacityParamDetail SyncZBoxCapacity detail param
type SyncZBoxCapacityParamDetail struct {
}

// SyncZBoxCapacityParam SyncZBoxCapacity request param
type SyncZBoxCapacityParam struct {
	BaseParam
	Params SyncZBoxCapacityParamDetail `json:"syncZBoxCapacity"`
}
// AckEventDataParamDetail AckEventData detail param
type AckEventDataParamDetail struct {
	EventSubscriptionUuid string `json:"eventSubscriptionUuid" validate:"required"`
	AlertDataUuid string `json:"alertDataUuid" validate:"required"`
	DataType string `json:"dataType" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	AckPeriodSec int `json:"ackPeriodSec" validate:"required"`
}

// AckEventDataParam AckEventData request param
type AckEventDataParam struct {
	BaseParam
	Params AckEventDataParamDetail `json:"params"`
}
// CheckResourcePermissionParamDetail CheckResourcePermission detail param
type CheckResourcePermissionParamDetail struct {
	ResourceType string `json:"resourceType" validate:"required"`
}

// CheckResourcePermissionParam CheckResourcePermission request param
type CheckResourcePermissionParam struct {
	BaseParam
	Params CheckResourcePermissionParamDetail `json:"checkResourcePermission"`
}
// ProvisionNfvInstConfigParamDetail ProvisionNfvInstConfig detail param
type ProvisionNfvInstConfigParamDetail struct {
}

// ProvisionNfvInstConfigParam ProvisionNfvInstConfig request param
type ProvisionNfvInstConfigParam struct {
	BaseParam
	Params ProvisionNfvInstConfigParamDetail `json:"provisionNfvInstConfig"`
}
// GetCandidateMiniHostsParamDetail GetCandidateMiniHosts detail param
type GetCandidateMiniHostsParamDetail struct {
	Local *bool `json:"local,omitempty"`
	Configure *bool `json:"configure,omitempty"`
}

// GetCandidateMiniHostsParam GetCandidateMiniHosts request param
type GetCandidateMiniHostsParam struct {
	BaseParam
	Params GetCandidateMiniHostsParamDetail `json:"getCandidateMiniHosts"`
}
// DeleteDatasetsParamDetail DeleteDatasets detail param
type DeleteDatasetsParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteDatasetsParam DeleteDatasets request param
type DeleteDatasetsParam struct {
	BaseParam
	Params DeleteDatasetsParamDetail `json:"deleteDatasets"`
}
// RevokeResourceSharingParamDetail RevokeResourceSharing detail param
type RevokeResourceSharingParamDetail struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	ToPublic bool `json:"toPublic,omitempty"`
	AccountUuids []string `json:"accountUuids,omitempty"`
	All bool `json:"all,omitempty"`
}

// RevokeResourceSharingParam RevokeResourceSharing request param
type RevokeResourceSharingParam struct {
	BaseParam
	Params RevokeResourceSharingParamDetail `json:"revokeResourceSharing"`
}
// DeleteModelServicesParamDetail DeleteModelServices detail param
type DeleteModelServicesParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteModelServicesParam DeleteModelServices request param
type DeleteModelServicesParam struct {
	BaseParam
	Params DeleteModelServicesParamDetail `json:"deleteModelServices"`
}
// ChangeL3NetworkStateParamDetail ChangeL3NetworkState detail param
type ChangeL3NetworkStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeL3NetworkStateParam ChangeL3NetworkState request param
type ChangeL3NetworkStateParam struct {
	BaseParam
	Params ChangeL3NetworkStateParamDetail `json:"changeL3NetworkState"`
}
// GetHostNUMATopologyParamDetail GetHostNUMATopology detail param
type GetHostNUMATopologyParamDetail struct {
}

// GetHostNUMATopologyParam GetHostNUMATopology request param
type GetHostNUMATopologyParam struct {
	BaseParam
	Params GetHostNUMATopologyParamDetail `json:"params"`
}
// CreateL2VirtualSwitchParamDetail CreateL2VirtualSwitch detail param
type CreateL2VirtualSwitchParamDetail struct {
	IsDistributed *bool `json:"isDistributed,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	PhysicalInterface string `json:"physicalInterface" validate:"required"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Pvlan *string `json:"pvlan,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2VirtualSwitchParam CreateL2VirtualSwitch request param
type CreateL2VirtualSwitchParam struct {
	BaseParam
	Params CreateL2VirtualSwitchParamDetail `json:"params"`
}
// AddVmNicToLoadBalancerParamDetail AddVmNicToLoadBalancer detail param
type AddVmNicToLoadBalancerParamDetail struct {
	VmNicUuids []string `json:"vmNicUuids" validate:"required"`
}

// AddVmNicToLoadBalancerParam AddVmNicToLoadBalancer request param
type AddVmNicToLoadBalancerParam struct {
	BaseParam
	Params AddVmNicToLoadBalancerParamDetail `json:"params"`
}
// UpdateBuildAppParamDetail UpdateBuildApp detail param
type UpdateBuildAppParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Version *string `json:"version,omitempty"`
}

// UpdateBuildAppParam UpdateBuildApp request param
type UpdateBuildAppParam struct {
	BaseParam
	Params UpdateBuildAppParamDetail `json:"updateBuildApp"`
}
// GetClusterDRSStatusParamDetail GetClusterDRSStatus detail param
type GetClusterDRSStatusParamDetail struct {
	DrsUuid string `json:"drsUuid" validate:"required"`
}

// GetClusterDRSStatusParam GetClusterDRSStatus request param
type GetClusterDRSStatusParam struct {
	BaseParam
	Params GetClusterDRSStatusParamDetail `json:"getClusterDRSStatus"`
}
// AddAliyunNasPrimaryStorageParamDetail AddAliyunNasPrimaryStorage detail param
type AddAliyunNasPrimaryStorageParamDetail struct {
	NasUuid string `json:"nasUuid" validate:"required"`
	AccessGroupUuid string `json:"accessGroupUuid" validate:"required"`
	VSwitchUuid *string `json:"vSwitchUuid,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasPrimaryStorageParam AddAliyunNasPrimaryStorage request param
type AddAliyunNasPrimaryStorageParam struct {
	BaseParam
	Params AddAliyunNasPrimaryStorageParamDetail `json:"params"`
}
// GetVmNumaParamDetail GetVmNuma detail param
type GetVmNumaParamDetail struct {
}

// GetVmNumaParam GetVmNuma request param
type GetVmNumaParam struct {
	BaseParam
	Params GetVmNumaParamDetail `json:"getVmNuma"`
}
// ChangeZoneStateParamDetail ChangeZoneState detail param
type ChangeZoneStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeZoneStateParam ChangeZoneState request param
type ChangeZoneStateParam struct {
	BaseParam
	Params ChangeZoneStateParamDetail `json:"changeZoneState"`
}
// AttachAppBuildSystemToZoneParamDetail AttachAppBuildSystemToZone detail param
type AttachAppBuildSystemToZoneParamDetail struct {
}

// AttachAppBuildSystemToZoneParam AttachAppBuildSystemToZone request param
type AttachAppBuildSystemToZoneParam struct {
	BaseParam
	Params AttachAppBuildSystemToZoneParamDetail `json:"params"`
}
// CreateDataVolumeParamDetail CreateDataVolume detail param
type CreateDataVolumeParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	DiskOfferingUuid *string `json:"diskOfferingUuid,omitempty"`
	DiskSize *int64 `json:"diskSize,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeParam CreateDataVolume request param
type CreateDataVolumeParam struct {
	BaseParam
	Params CreateDataVolumeParamDetail `json:"params"`
}
// UngenerateHygonMdevDevicesParamDetail UngenerateHygonMdevDevices detail param
type UngenerateHygonMdevDevicesParamDetail struct {
}

// UngenerateHygonMdevDevicesParam UngenerateHygonMdevDevices request param
type UngenerateHygonMdevDevicesParam struct {
	BaseParam
	Params UngenerateHygonMdevDevicesParamDetail `json:"ungenerateHygonMdevDevices"`
}
// DeletePluginDriversParamDetail DeletePluginDrivers detail param
type DeletePluginDriversParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePluginDriversParam DeletePluginDrivers request param
type DeletePluginDriversParam struct {
	BaseParam
	Params DeletePluginDriversParamDetail `json:"deletePluginDrivers"`
}
// BatchCreateBaremetalChassisParamDetail BatchCreateBaremetalChassis detail param
type BatchCreateBaremetalChassisParamDetail struct {
	BaremetalChassisInfo string `json:"baremetalChassisInfo" validate:"required"`
	LongJobName *string `json:"longJobName,omitempty"`
	LongJobDescription *string `json:"longJobDescription,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// BatchCreateBaremetalChassisParam BatchCreateBaremetalChassis request param
type BatchCreateBaremetalChassisParam struct {
	BaseParam
	Params BatchCreateBaremetalChassisParamDetail `json:"params"`
}
// AddSchedulerJobToSchedulerTriggerParamDetail AddSchedulerJobToSchedulerTrigger detail param
type AddSchedulerJobToSchedulerTriggerParamDetail struct {
	TriggerNow *bool `json:"triggerNow,omitempty"`
}

// AddSchedulerJobToSchedulerTriggerParam AddSchedulerJobToSchedulerTrigger request param
type AddSchedulerJobToSchedulerTriggerParam struct {
	BaseParam
	Params AddSchedulerJobToSchedulerTriggerParamDetail `json:"params"`
}
// DetachPolicyFromRoleParamDetail DetachPolicyFromRole detail param
type DetachPolicyFromRoleParamDetail struct {
}

// DetachPolicyFromRoleParam DetachPolicyFromRole request param
type DetachPolicyFromRoleParam struct {
	BaseParam
	Params DetachPolicyFromRoleParamDetail `json:"detachPolicyFromRole"`
}
// RestartModelServiceGroupsParamDetail RestartModelServiceGroups detail param
type RestartModelServiceGroupsParamDetail struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// RestartModelServiceGroupsParam RestartModelServiceGroups request param
type RestartModelServiceGroupsParam struct {
	BaseParam
	Params RestartModelServiceGroupsParamDetail `json:"restartModelServiceGroups"`
}
// GetLoadBalancerOwnerParamDetail GetLoadBalancerOwner detail param
type GetLoadBalancerOwnerParamDetail struct {
}

// GetLoadBalancerOwnerParam GetLoadBalancerOwner request param
type GetLoadBalancerOwnerParam struct {
	BaseParam
	Params GetLoadBalancerOwnerParamDetail `json:"getLoadBalancerOwner"`
}
// GetNicQosParamDetail GetNicQos detail param
type GetNicQosParamDetail struct {
	ForceSync *bool `json:"forceSync,omitempty"`
}

// GetNicQosParam GetNicQos request param
type GetNicQosParam struct {
	BaseParam
	Params GetNicQosParamDetail `json:"getNicQos"`
}
// ChangeVmNicNetworkParamDetail ChangeVmNicNetwork detail param
type ChangeVmNicNetworkParamDetail struct {
	StaticIp *string `json:"staticIp,omitempty"`
}

// ChangeVmNicNetworkParam ChangeVmNicNetwork request param
type ChangeVmNicNetworkParam struct {
	BaseParam
	Params ChangeVmNicNetworkParamDetail `json:"params"`
}
// CreateBareMetal2IpmiChassisHardwareInfoParamDetail CreateBareMetal2IpmiChassisHardwareInfo detail param
type CreateBareMetal2IpmiChassisHardwareInfoParamDetail struct {
	IpmiAddress string `json:"ipmiAddress" validate:"required"`
	IpmiPort int `json:"ipmiPort" validate:"required"`
	HardwareInfo string `json:"hardwareInfo" validate:"required"`
	ConvertInfo *string `json:"convertInfo,omitempty"`
}

// CreateBareMetal2IpmiChassisHardwareInfoParam CreateBareMetal2IpmiChassisHardwareInfo request param
type CreateBareMetal2IpmiChassisHardwareInfoParam struct {
	BaseParam
	Params CreateBareMetal2IpmiChassisHardwareInfoParamDetail `json:"params"`
}
// RemoveLabelFromAlarmParamDetail RemoveLabelFromAlarm detail param
type RemoveLabelFromAlarmParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveLabelFromAlarmParam RemoveLabelFromAlarm request param
type RemoveLabelFromAlarmParam struct {
	BaseParam
	Params RemoveLabelFromAlarmParamDetail `json:"removeLabelFromAlarm"`
}
// DeleteIAM2VirtualIDLdapBindingParamDetail DeleteIAM2VirtualIDLdapBinding detail param
type DeleteIAM2VirtualIDLdapBindingParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteIAM2VirtualIDLdapBindingParam DeleteIAM2VirtualIDLdapBinding request param
type DeleteIAM2VirtualIDLdapBindingParam struct {
	BaseParam
	Params DeleteIAM2VirtualIDLdapBindingParamDetail `json:"deleteIAM2VirtualIDLdapBinding"`
}
// UpdateVmPriorityParamDetail UpdateVmPriority detail param
type UpdateVmPriorityParamDetail struct {
	Priority string `json:"priority" validate:"required"`
}

// UpdateVmPriorityParam UpdateVmPriority request param
type UpdateVmPriorityParam struct {
	BaseParam
	Params UpdateVmPriorityParamDetail `json:"updateVmPriority"`
}
// DetachMdevDeviceFromVmParamDetail DetachMdevDeviceFromVm detail param
type DetachMdevDeviceFromVmParamDetail struct {
}

// DetachMdevDeviceFromVmParam DetachMdevDeviceFromVm request param
type DetachMdevDeviceFromVmParam struct {
	BaseParam
	Params DetachMdevDeviceFromVmParamDetail `json:"detachMdevDeviceFromVm"`
}
// DeleteVmHostnameParamDetail DeleteVmHostname detail param
type DeleteVmHostnameParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVmHostnameParam DeleteVmHostname request param
type DeleteVmHostnameParam struct {
	BaseParam
	Params DeleteVmHostnameParamDetail `json:"deleteVmHostname"`
}
// GetLicenseCapabilitiesParamDetail GetLicenseCapabilities detail param
type GetLicenseCapabilitiesParamDetail struct {
}

// GetLicenseCapabilitiesParam GetLicenseCapabilities request param
type GetLicenseCapabilitiesParam struct {
	BaseParam
	Params GetLicenseCapabilitiesParamDetail `json:"getLicenseCapabilities"`
}
// CreateFirewallRuleTemplateParamDetail CreateFirewallRuleTemplate detail param
type CreateFirewallRuleTemplateParamDetail struct {
	Action string `json:"action" validate:"required"`
	Protocol *string `json:"protocol,omitempty"`
	Name string `json:"name" validate:"required"`
	DestPort *string `json:"destPort,omitempty"`
	SourcePort *string `json:"sourcePort,omitempty"`
	SourceIp *string `json:"sourceIp,omitempty"`
	DestIp *string `json:"destIp,omitempty"`
	AllowStates *string `json:"allowStates,omitempty"`
	TcpFlag *string `json:"tcpFlag,omitempty"`
	IcmpTypeName *string `json:"icmpTypeName,omitempty"`
	RuleNumber int `json:"ruleNumber" validate:"required"`
	EnableLog *bool `json:"enableLog,omitempty"`
	State *string `json:"state,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFirewallRuleTemplateParam CreateFirewallRuleTemplate request param
type CreateFirewallRuleTemplateParam struct {
	BaseParam
	Params CreateFirewallRuleTemplateParamDetail `json:"params"`
}
// ChangeIAM2ProjectStateParamDetail ChangeIAM2ProjectState detail param
type ChangeIAM2ProjectStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeIAM2ProjectStateParam ChangeIAM2ProjectState request param
type ChangeIAM2ProjectStateParam struct {
	BaseParam
	Params ChangeIAM2ProjectStateParamDetail `json:"changeIAM2ProjectState"`
}
// SetVmSoundTypeParamDetail SetVmSoundType detail param
type SetVmSoundTypeParamDetail struct {
	SoundType string `json:"soundType" validate:"required"`
}

// SetVmSoundTypeParam SetVmSoundType request param
type SetVmSoundTypeParam struct {
	BaseParam
	Params SetVmSoundTypeParamDetail `json:"setVmSoundType"`
}
// MergeDataOnBackupStorageParamDetail MergeDataOnBackupStorage detail param
type MergeDataOnBackupStorageParamDetail struct {
}

// MergeDataOnBackupStorageParam MergeDataOnBackupStorage request param
type MergeDataOnBackupStorageParam struct {
	BaseParam
	Params MergeDataOnBackupStorageParamDetail `json:"mergeDataOnBackupStorage"`
}
// GetCdpBackupStorageRequirementParamDetail GetCdpBackupStorageRequirement detail param
type GetCdpBackupStorageRequirementParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetCdpBackupStorageRequirementParam GetCdpBackupStorageRequirement request param
type GetCdpBackupStorageRequirementParam struct {
	BaseParam
	Params GetCdpBackupStorageRequirementParamDetail `json:"getCdpBackupStorageRequirement"`
}
// AddAttributesToIAM2VirtualIDGroupParamDetail AddAttributesToIAM2VirtualIDGroup detail param
type AddAttributesToIAM2VirtualIDGroupParamDetail struct {
	Attributes []AttributeParam `json:"attributes" validate:"required"`
}

// AddAttributesToIAM2VirtualIDGroupParam AddAttributesToIAM2VirtualIDGroup request param
type AddAttributesToIAM2VirtualIDGroupParam struct {
	BaseParam
	Params AddAttributesToIAM2VirtualIDGroupParamDetail `json:"params"`
}
// ChangeAffinityGroupStateParamDetail ChangeAffinityGroupState detail param
type ChangeAffinityGroupStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeAffinityGroupStateParam ChangeAffinityGroupState request param
type ChangeAffinityGroupStateParam struct {
	BaseParam
	Params ChangeAffinityGroupStateParamDetail `json:"changeAffinityGroupState"`
}
// ChangeSecurityGroupRuleStateParamDetail ChangeSecurityGroupRuleState detail param
type ChangeSecurityGroupRuleStateParamDetail struct {
	RuleUuids []string `json:"ruleUuids" validate:"required"`
	State string `json:"state" validate:"required"`
}

// ChangeSecurityGroupRuleStateParam ChangeSecurityGroupRuleState request param
type ChangeSecurityGroupRuleStateParam struct {
	BaseParam
	Params ChangeSecurityGroupRuleStateParamDetail `json:"changeSecurityGroupRuleState"`
}
// AddVmNicToSecurityGroupParamDetail AddVmNicToSecurityGroup detail param
type AddVmNicToSecurityGroupParamDetail struct {
	VmNicUuids []string `json:"vmNicUuids" validate:"required"`
}

// AddVmNicToSecurityGroupParam AddVmNicToSecurityGroup request param
type AddVmNicToSecurityGroupParam struct {
	BaseParam
	Params AddVmNicToSecurityGroupParamDetail `json:"params"`
}
// SyncAliyunRouteEntryFromRemoteParamDetail SyncAliyunRouteEntryFromRemote detail param
type SyncAliyunRouteEntryFromRemoteParamDetail struct {
	VRouterType string `json:"vRouterType" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncAliyunRouteEntryFromRemoteParam SyncAliyunRouteEntryFromRemote request param
type SyncAliyunRouteEntryFromRemoteParam struct {
	BaseParam
	Params SyncAliyunRouteEntryFromRemoteParamDetail `json:"syncAliyunRouteEntryFromRemote"`
}
// UpdateEmailAddressOfSNSEmailEndpointParamDetail UpdateEmailAddressOfSNSEmailEndpoint detail param
type UpdateEmailAddressOfSNSEmailEndpointParamDetail struct {
	EmailAddressUuid string `json:"emailAddressUuid" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	EmailAddress string `json:"emailAddress" validate:"required"`
}

// UpdateEmailAddressOfSNSEmailEndpointParam UpdateEmailAddressOfSNSEmailEndpoint request param
type UpdateEmailAddressOfSNSEmailEndpointParam struct {
	BaseParam
	Params UpdateEmailAddressOfSNSEmailEndpointParamDetail `json:"updateEmailAddressOfSNSEmailEndpoint"`
}
// GetMetricLabelValueParamDetail GetMetricLabelValue detail param
type GetMetricLabelValueParamDetail struct {
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	StartTime *int64 `json:"startTime,omitempty"`
	EndTime *int64 `json:"endTime,omitempty"`
	LabelNames []string `json:"labelNames" validate:"required"`
	FilterLabels []string `json:"filterLabels,omitempty"`
}

// GetMetricLabelValueParam GetMetricLabelValue request param
type GetMetricLabelValueParam struct {
	BaseParam
	Params GetMetricLabelValueParamDetail `json:"getMetricLabelValue"`
}
// GetCandidateZonesClustersHostsForCreatingVmParamDetail GetCandidateZonesClustersHostsForCreatingVm detail param
type GetCandidateZonesClustersHostsForCreatingVmParamDetail struct {
	InstanceOfferingUuid *string `json:"instanceOfferingUuid,omitempty"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	RootDiskOfferingUuid *string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize *int64 `json:"rootDiskSize,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
}

// GetCandidateZonesClustersHostsForCreatingVmParam GetCandidateZonesClustersHostsForCreatingVm request param
type GetCandidateZonesClustersHostsForCreatingVmParam struct {
	BaseParam
	Params GetCandidateZonesClustersHostsForCreatingVmParamDetail `json:"getCandidateZonesClustersHostsForCreatingVm"`
}
// CreateResourcePriceParamDetail CreateResourcePrice detail param
type CreateResourcePriceParamDetail struct {
	ResourceName string `json:"resourceName" validate:"required"`
	ResourceUnit *string `json:"resourceUnit,omitempty"`
	TimeUnit string `json:"timeUnit" validate:"required"`
	Price float64 `json:"price" validate:"required"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	DateInLong *int64 `json:"dateInLong,omitempty"`
	TableUuid *string `json:"tableUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateResourcePriceParam CreateResourcePrice request param
type CreateResourcePriceParam struct {
	BaseParam
	Params CreateResourcePriceParamDetail `json:"params"`
}
// RemoveSchedulerJobGroupFromSchedulerTriggerParamDetail RemoveSchedulerJobGroupFromSchedulerTrigger detail param
type RemoveSchedulerJobGroupFromSchedulerTriggerParamDetail struct {
}

// RemoveSchedulerJobGroupFromSchedulerTriggerParam RemoveSchedulerJobGroupFromSchedulerTrigger request param
type RemoveSchedulerJobGroupFromSchedulerTriggerParam struct {
	BaseParam
	Params RemoveSchedulerJobGroupFromSchedulerTriggerParamDetail `json:"removeSchedulerJobGroupFromSchedulerTrigger"`
}
// ChangeAccountPriceTableBindingParamDetail ChangeAccountPriceTableBinding detail param
type ChangeAccountPriceTableBindingParamDetail struct {
}

// ChangeAccountPriceTableBindingParam ChangeAccountPriceTableBinding request param
type ChangeAccountPriceTableBindingParam struct {
	BaseParam
	Params ChangeAccountPriceTableBindingParamDetail `json:"changeAccountPriceTableBinding"`
}
// DeleteVolumeQosParamDetail DeleteVolumeQos detail param
type DeleteVolumeQosParamDetail struct {
	Mode *string `json:"mode,omitempty"`
}

// DeleteVolumeQosParam DeleteVolumeQos request param
type DeleteVolumeQosParam struct {
	BaseParam
	Params DeleteVolumeQosParamDetail `json:"deleteVolumeQos"`
}
// GetL3NetworkDhcpIpAddressParamDetail GetL3NetworkDhcpIpAddress detail param
type GetL3NetworkDhcpIpAddressParamDetail struct {
}

// GetL3NetworkDhcpIpAddressParam GetL3NetworkDhcpIpAddress request param
type GetL3NetworkDhcpIpAddressParam struct {
	BaseParam
	Params GetL3NetworkDhcpIpAddressParamDetail `json:"getL3NetworkDhcpIpAddress"`
}
// CreateFirewallRuleSetParamDetail CreateFirewallRuleSet detail param
type CreateFirewallRuleSetParamDetail struct {
	Name string `json:"name" validate:"required"`
	ActionType *string `json:"actionType,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFirewallRuleSetParam CreateFirewallRuleSet request param
type CreateFirewallRuleSetParam struct {
	BaseParam
	Params CreateFirewallRuleSetParamDetail `json:"params"`
}
// CreateJitSecretResourcePoolParamDetail CreateJitSecretResourcePool detail param
type CreateJitSecretResourcePoolParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	Ability *string `json:"ability,omitempty"`
	Type string `json:"type" validate:"required"`
	HeartbeatInterval int `json:"heartbeatInterval" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateJitSecretResourcePoolParam CreateJitSecretResourcePool request param
type CreateJitSecretResourcePoolParam struct {
	BaseParam
	Params CreateJitSecretResourcePoolParamDetail `json:"params"`
}
// GetBaremetalChassisPowerStatusParamDetail GetBaremetalChassisPowerStatus detail param
type GetBaremetalChassisPowerStatusParamDetail struct {
}

// GetBaremetalChassisPowerStatusParam GetBaremetalChassisPowerStatus request param
type GetBaremetalChassisPowerStatusParam struct {
	BaseParam
	Params GetBaremetalChassisPowerStatusParamDetail `json:"getBaremetalChassisPowerStatus"`
}
// RefreshFirewallParamDetail RefreshFirewall detail param
type RefreshFirewallParamDetail struct {
}

// RefreshFirewallParam RefreshFirewall request param
type RefreshFirewallParam struct {
	BaseParam
	Params RefreshFirewallParamDetail `json:"refreshFirewall"`
}
// DetachL3NetworksFromIPsecConnectionParamDetail DetachL3NetworksFromIPsecConnection detail param
type DetachL3NetworksFromIPsecConnectionParamDetail struct {
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
}

// DetachL3NetworksFromIPsecConnectionParam DetachL3NetworksFromIPsecConnection request param
type DetachL3NetworksFromIPsecConnectionParam struct {
	BaseParam
	Params DetachL3NetworksFromIPsecConnectionParamDetail `json:"detachL3NetworksFromIPsecConnection"`
}
// UpdateAutoScalingGroupAddingNewInstanceRuleParamDetail UpdateAutoScalingGroupAddingNewInstanceRule detail param
type UpdateAutoScalingGroupAddingNewInstanceRuleParamDetail struct {
	AdjustmentType *string `json:"adjustmentType,omitempty"`
	AdjustmentValue *int `json:"adjustmentValue,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Cooldown *int64 `json:"cooldown,omitempty"`
}

// UpdateAutoScalingGroupAddingNewInstanceRuleParam UpdateAutoScalingGroupAddingNewInstanceRule request param
type UpdateAutoScalingGroupAddingNewInstanceRuleParam struct {
	BaseParam
	Params UpdateAutoScalingGroupAddingNewInstanceRuleParamDetail `json:"updateAutoScalingGroupAddingNewInstanceRule"`
}
// GetFaultToleranceVmsParamDetail GetFaultToleranceVms detail param
type GetFaultToleranceVmsParamDetail struct {
	FaultToleranceVmUuid string `json:"faultToleranceVmUuid" validate:"required"`
}

// GetFaultToleranceVmsParam GetFaultToleranceVms request param
type GetFaultToleranceVmsParam struct {
	BaseParam
	Params GetFaultToleranceVmsParamDetail `json:"getFaultToleranceVms"`
}
// DeleteAliyunKeySecretParamDetail DeleteAliyunKeySecret detail param
type DeleteAliyunKeySecretParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAliyunKeySecretParam DeleteAliyunKeySecret request param
type DeleteAliyunKeySecretParam struct {
	BaseParam
	Params DeleteAliyunKeySecretParamDetail `json:"deleteAliyunKeySecret"`
}
// CreateVmInstanceFromVolumeSnapshotParamDetail CreateVmInstanceFromVolumeSnapshot detail param
type CreateVmInstanceFromVolumeSnapshotParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	InstanceOfferingUuid *string `json:"instanceOfferingUuid,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	ReservedMemorySize *int64 `json:"reservedMemorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	VmNicParams *string `json:"vmNicParams,omitempty"`
	Type *string `json:"type,omitempty"`
	Platform *string `json:"platform,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceFromVolumeSnapshotParam CreateVmInstanceFromVolumeSnapshot request param
type CreateVmInstanceFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateVmInstanceFromVolumeSnapshotParamDetail `json:"params"`
}
// PowerResetBareMetal2ChassisParamDetail PowerResetBareMetal2Chassis detail param
type PowerResetBareMetal2ChassisParamDetail struct {
	BootDev *string `json:"bootDev,omitempty"`
}

// PowerResetBareMetal2ChassisParam PowerResetBareMetal2Chassis request param
type PowerResetBareMetal2ChassisParam struct {
	BaseParam
	Params PowerResetBareMetal2ChassisParamDetail `json:"powerResetBareMetal2Chassis"`
}
// PrometheusQueryVmMonitoringDataParamDetail PrometheusQueryVmMonitoringData detail param
type PrometheusQueryVmMonitoringDataParamDetail struct {
	VmUuids []string `json:"vmUuids" validate:"required"`
	Instant bool `json:"instant,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
	EndTime *int64 `json:"endTime,omitempty"`
	Step *string `json:"step,omitempty"`
	Expression string `json:"expression" validate:"required"`
	RelativeTime *string `json:"relativeTime,omitempty"`
}

// PrometheusQueryVmMonitoringDataParam PrometheusQueryVmMonitoringData request param
type PrometheusQueryVmMonitoringDataParam struct {
	BaseParam
	Params PrometheusQueryVmMonitoringDataParamDetail `json:"prometheusQueryVmMonitoringData"`
}
// UpdateResourceConfigsParamDetail UpdateResourceConfigs detail param
type UpdateResourceConfigsParamDetail struct {
	ResourceConfigs []UpdateResourceConfigs_ResourceConfigAOParam `json:"resourceConfigs" validate:"required"`
}

// UpdateResourceConfigsParam UpdateResourceConfigs request param
type UpdateResourceConfigsParam struct {
	BaseParam
	Params UpdateResourceConfigsParamDetail `json:"params"`
}
// RevertVolumeFromSnapshotParamDetail RevertVolumeFromSnapshot detail param
type RevertVolumeFromSnapshotParamDetail struct {
}

// RevertVolumeFromSnapshotParam RevertVolumeFromSnapshot request param
type RevertVolumeFromSnapshotParam struct {
	BaseParam
	Params RevertVolumeFromSnapshotParamDetail `json:"revertVolumeFromSnapshot"`
}
// GetBlockPrimaryStorageMetadataParamDetail GetBlockPrimaryStorageMetadata detail param
type GetBlockPrimaryStorageMetadataParamDetail struct {
	VendorName string `json:"vendorName" validate:"required"`
	Metadata string `json:"metadata" validate:"required"`
}

// GetBlockPrimaryStorageMetadataParam GetBlockPrimaryStorageMetadata request param
type GetBlockPrimaryStorageMetadataParam struct {
	BaseParam
	Params GetBlockPrimaryStorageMetadataParamDetail `json:"param"`
}
// UpdateBondingParamDetail UpdateBonding detail param
type UpdateBondingParamDetail struct {
	SlaveUuids []string `json:"slaveUuids,omitempty"`
	SlaveNames []string `json:"slaveNames,omitempty"`
	Type *string `json:"type,omitempty"`
	Mode *string `json:"mode,omitempty"`
	XmitHashPolicy *string `json:"xmitHashPolicy,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateBondingParam UpdateBonding request param
type UpdateBondingParam struct {
	BaseParam
	Params UpdateBondingParamDetail `json:"updateBonding"`
}
// GetManagementNodeArchParamDetail GetManagementNodeArch detail param
type GetManagementNodeArchParamDetail struct {
}

// GetManagementNodeArchParam GetManagementNodeArch request param
type GetManagementNodeArchParam struct {
	BaseParam
	Params GetManagementNodeArchParamDetail `json:"getManagementNodeArch"`
}
// DetachScsiLunFromHostParamDetail DetachScsiLunFromHost detail param
type DetachScsiLunFromHostParamDetail struct {
	HostUuid *string `json:"hostUuid,omitempty"`
}

// DetachScsiLunFromHostParam DetachScsiLunFromHost request param
type DetachScsiLunFromHostParam struct {
	BaseParam
	Params DetachScsiLunFromHostParamDetail `json:"detachScsiLunFromHost"`
}
// DisableCbtTaskParamDetail DisableCbtTask detail param
type DisableCbtTaskParamDetail struct {
	Force *bool `json:"force,omitempty"`
}

// DisableCbtTaskParam DisableCbtTask request param
type DisableCbtTaskParam struct {
	BaseParam
	Params DisableCbtTaskParamDetail `json:"params"`
}
// RefreshLocalRaidParamDetail RefreshLocalRaid detail param
type RefreshLocalRaidParamDetail struct {
	HostUuid string `json:"hostUuid" validate:"required"`
}

// RefreshLocalRaidParam RefreshLocalRaid request param
type RefreshLocalRaidParam struct {
	BaseParam
	Params RefreshLocalRaidParamDetail `json:"refreshLocalRaid"`
}
// UpdateSubscribeEventParamDetail UpdateSubscribeEvent detail param
type UpdateSubscribeEventParamDetail struct {
	EmergencyLevel *string `json:"emergencyLevel,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateSubscribeEventParam UpdateSubscribeEvent request param
type UpdateSubscribeEventParam struct {
	BaseParam
	Params UpdateSubscribeEventParamDetail `json:"updateSubscribeEvent"`
}
// SetVmSshKeyParamDetail SetVmSshKey detail param
type SetVmSshKeyParamDetail struct {
	SshKey string `json:"SshKey" validate:"required"`
}

// SetVmSshKeyParam SetVmSshKey request param
type SetVmSshKeyParam struct {
	BaseParam
	Params SetVmSshKeyParamDetail `json:"setVmSshKey"`
}
// FailoverFaultToleranceVmParamDetail FailoverFaultToleranceVm detail param
type FailoverFaultToleranceVmParamDetail struct {
	FaultToleranceVmUuid string `json:"faultToleranceVmUuid" validate:"required"`
}

// FailoverFaultToleranceVmParam FailoverFaultToleranceVm request param
type FailoverFaultToleranceVmParam struct {
	BaseParam
	Params FailoverFaultToleranceVmParamDetail `json:"failoverFaultToleranceVm"`
}
// EjectZBoxParamDetail EjectZBox detail param
type EjectZBoxParamDetail struct {
}

// EjectZBoxParam EjectZBox request param
type EjectZBoxParam struct {
	BaseParam
	Params EjectZBoxParamDetail `json:"ejectZBox"`
}
// PrometheusQueryMetadataParamDetail PrometheusQueryMetadata detail param
type PrometheusQueryMetadataParamDetail struct {
	Matches []string `json:"matches" validate:"required"`
}

// PrometheusQueryMetadataParam PrometheusQueryMetadata request param
type PrometheusQueryMetadataParam struct {
	BaseParam
	Params PrometheusQueryMetadataParamDetail `json:"prometheusQueryMetadata"`
}
// DeleteFirewallRuleTemplateParamDetail DeleteFirewallRuleTemplate detail param
type DeleteFirewallRuleTemplateParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteFirewallRuleTemplateParam DeleteFirewallRuleTemplate request param
type DeleteFirewallRuleTemplateParam struct {
	BaseParam
	Params DeleteFirewallRuleTemplateParamDetail `json:"deleteFirewallRuleTemplate"`
}
// DetachPciDeviceFromVmParamDetail DetachPciDeviceFromVm detail param
type DetachPciDeviceFromVmParamDetail struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// DetachPciDeviceFromVmParam DetachPciDeviceFromVm request param
type DetachPciDeviceFromVmParam struct {
	BaseParam
	Params DetachPciDeviceFromVmParamDetail `json:"params"`
}
// ExecuteGuestVmCommandParamDetail ExecuteGuestVmCommand detail param
type ExecuteGuestVmCommandParamDetail struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Platform string `json:"platform" validate:"required"`
	Command string `json:"command" validate:"required"`
	CommandTimeout *int `json:"commandTimeout,omitempty"`
}

// ExecuteGuestVmCommandParam ExecuteGuestVmCommand request param
type ExecuteGuestVmCommandParam struct {
	BaseParam
	Params ExecuteGuestVmCommandParamDetail `json:"params"`
}
// ChangeVpcSharedQosBandwidthParamDetail ChangeVpcSharedQosBandwidth detail param
type ChangeVpcSharedQosBandwidthParamDetail struct {
	Bandwidth int64 `json:"bandwidth" validate:"required"`
}

// ChangeVpcSharedQosBandwidthParam ChangeVpcSharedQosBandwidth request param
type ChangeVpcSharedQosBandwidthParam struct {
	BaseParam
	Params ChangeVpcSharedQosBandwidthParamDetail `json:"params"`
}
// RemoveAttributesFromIAM2VirtualIDGroupParamDetail RemoveAttributesFromIAM2VirtualIDGroup detail param
type RemoveAttributesFromIAM2VirtualIDGroupParamDetail struct {
	AttributeUuids []string `json:"attributeUuids" validate:"required"`
}

// RemoveAttributesFromIAM2VirtualIDGroupParam RemoveAttributesFromIAM2VirtualIDGroup request param
type RemoveAttributesFromIAM2VirtualIDGroupParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2VirtualIDGroupParamDetail `json:"removeAttributesFromIAM2VirtualIDGroup"`
}
// AddAttributesToIAM2VirtualIDParamDetail AddAttributesToIAM2VirtualID detail param
type AddAttributesToIAM2VirtualIDParamDetail struct {
	Attributes []AttributeParam `json:"attributes" validate:"required"`
}

// AddAttributesToIAM2VirtualIDParam AddAttributesToIAM2VirtualID request param
type AddAttributesToIAM2VirtualIDParam struct {
	BaseParam
	Params AddAttributesToIAM2VirtualIDParamDetail `json:"params"`
}
// FlattenVolumeParamDetail FlattenVolume detail param
type FlattenVolumeParamDetail struct {
	DryRun *bool `json:"dryRun,omitempty"`
}

// FlattenVolumeParam FlattenVolume request param
type FlattenVolumeParam struct {
	BaseParam
	Params FlattenVolumeParamDetail `json:"flattenVolume"`
}
// CreateAliyunDiskFromRemoteParamDetail CreateAliyunDiskFromRemote detail param
type CreateAliyunDiskFromRemoteParamDetail struct {
	IdentityUuid string `json:"identityUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	SizeWithGB *int `json:"sizeWithGB,omitempty"`
	Description *string `json:"description,omitempty"`
	DiskCategory *string `json:"diskCategory,omitempty"`
	SnapshotUuid *string `json:"snapshotUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunDiskFromRemoteParam CreateAliyunDiskFromRemote request param
type CreateAliyunDiskFromRemoteParam struct {
	BaseParam
	Params CreateAliyunDiskFromRemoteParamDetail `json:"params"`
}
// DeleteEcsSecurityGroupRuleRemoteParamDetail DeleteEcsSecurityGroupRuleRemote detail param
type DeleteEcsSecurityGroupRuleRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsSecurityGroupRuleRemoteParam DeleteEcsSecurityGroupRuleRemote request param
type DeleteEcsSecurityGroupRuleRemoteParam struct {
	BaseParam
	Params DeleteEcsSecurityGroupRuleRemoteParamDetail `json:"deleteEcsSecurityGroupRuleRemote"`
}
// DetachAliyunDiskFromEcsParamDetail DetachAliyunDiskFromEcs detail param
type DetachAliyunDiskFromEcsParamDetail struct {
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// DetachAliyunDiskFromEcsParam DetachAliyunDiskFromEcs request param
type DetachAliyunDiskFromEcsParam struct {
	BaseParam
	Params DetachAliyunDiskFromEcsParamDetail `json:"params"`
}
// GetCandidateAffinityGroupForAttachingVmParamDetail GetCandidateAffinityGroupForAttachingVm detail param
type GetCandidateAffinityGroupForAttachingVmParamDetail struct {
	VmUuid string `json:"vmUuid" validate:"required"`
}

// GetCandidateAffinityGroupForAttachingVmParam GetCandidateAffinityGroupForAttachingVm request param
type GetCandidateAffinityGroupForAttachingVmParam struct {
	BaseParam
	Params GetCandidateAffinityGroupForAttachingVmParamDetail `json:"getCandidateAffinityGroupForAttachingVm"`
}
// UpdateFirewallIpSetTemplateParamDetail UpdateFirewallIpSetTemplate detail param
type UpdateFirewallIpSetTemplateParamDetail struct {
	Name string `json:"name,omitempty"`
	SourceValue *string `json:"sourceValue,omitempty"`
	DestValue *string `json:"destValue,omitempty"`
	Type *string `json:"type,omitempty"`
}

// UpdateFirewallIpSetTemplateParam UpdateFirewallIpSetTemplate request param
type UpdateFirewallIpSetTemplateParam struct {
	BaseParam
	Params UpdateFirewallIpSetTemplateParamDetail `json:"updateFirewallIpSetTemplate"`
}
// AddAccessControlListRedirectRuleParamDetail AddAccessControlListRedirectRule detail param
type AddAccessControlListRedirectRuleParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Url *string `json:"url,omitempty"`
	RedirectPort *int `json:"redirectPort,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAccessControlListRedirectRuleParam AddAccessControlListRedirectRule request param
type AddAccessControlListRedirectRuleParam struct {
	BaseParam
	Params AddAccessControlListRedirectRuleParamDetail `json:"params"`
}
// DetachHostFromHostSchedulingRuleGroupParamDetail DetachHostFromHostSchedulingRuleGroup detail param
type DetachHostFromHostSchedulingRuleGroupParamDetail struct {
	HostUuid string `json:"hostUuid" validate:"required"`
}

// DetachHostFromHostSchedulingRuleGroupParam DetachHostFromHostSchedulingRuleGroup request param
type DetachHostFromHostSchedulingRuleGroupParam struct {
	BaseParam
	Params DetachHostFromHostSchedulingRuleGroupParamDetail `json:"detachHostFromHostSchedulingRuleGroup"`
}
// UpdateAliyunRouteInterfaceRemoteParamDetail UpdateAliyunRouteInterfaceRemote detail param
type UpdateAliyunRouteInterfaceRemoteParamDetail struct {
	Op string `json:"op" validate:"required"`
	VRouterType string `json:"vRouterType" validate:"required"`
}

// UpdateAliyunRouteInterfaceRemoteParam UpdateAliyunRouteInterfaceRemote request param
type UpdateAliyunRouteInterfaceRemoteParam struct {
	BaseParam
	Params UpdateAliyunRouteInterfaceRemoteParamDetail `json:"updateAliyunRouteInterfaceRemote"`
}
// GetPciDeviceSpecCandidatesParamDetail GetPciDeviceSpecCandidates detail param
type GetPciDeviceSpecCandidatesParamDetail struct {
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	VmInstanceUuid *string `json:"vmInstanceUuid,omitempty"`
	VmInstanceUuids []string `json:"vmInstanceUuids,omitempty"`
	Types []string `json:"types,omitempty"`
}

// GetPciDeviceSpecCandidatesParam GetPciDeviceSpecCandidates request param
type GetPciDeviceSpecCandidatesParam struct {
	BaseParam
	Params GetPciDeviceSpecCandidatesParamDetail `json:"getPciDeviceSpecCandidates"`
}
// PrometheusQueryPassThroughParamDetail PrometheusQueryPassThrough detail param
type PrometheusQueryPassThroughParamDetail struct {
	Instant bool `json:"instant,omitempty"`
	StartTime *int64 `json:"startTime,omitempty"`
	EndTime *int64 `json:"endTime,omitempty"`
	Step *string `json:"step,omitempty"`
	Expression string `json:"expression" validate:"required"`
	RelativeTime *string `json:"relativeTime,omitempty"`
}

// PrometheusQueryPassThroughParam PrometheusQueryPassThrough request param
type PrometheusQueryPassThroughParam struct {
	BaseParam
	Params PrometheusQueryPassThroughParamDetail `json:"prometheusQueryPassThrough"`
}
// AttachVmNicToVmParamDetail AttachVmNicToVm detail param
type AttachVmNicToVmParamDetail struct {
}

// AttachVmNicToVmParam AttachVmNicToVm request param
type AttachVmNicToVmParam struct {
	BaseParam
	Params AttachVmNicToVmParamDetail `json:"params"`
}
// RemoveMonFromCephBackupStorageParamDetail RemoveMonFromCephBackupStorage detail param
type RemoveMonFromCephBackupStorageParamDetail struct {
	MonHostnames []string `json:"monHostnames" validate:"required"`
}

// RemoveMonFromCephBackupStorageParam RemoveMonFromCephBackupStorage request param
type RemoveMonFromCephBackupStorageParam struct {
	BaseParam
	Params RemoveMonFromCephBackupStorageParamDetail `json:"removeMonFromCephBackupStorage"`
}
// GetVmDeviceAddressParamDetail GetVmDeviceAddress detail param
type GetVmDeviceAddressParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceTypes []string `json:"resourceTypes" validate:"required"`
}

// GetVmDeviceAddressParam GetVmDeviceAddress request param
type GetVmDeviceAddressParam struct {
	BaseParam
	Params GetVmDeviceAddressParamDetail `json:"getVmDeviceAddress"`
}
// RemoveInstanceFromMonitorGroupParamDetail RemoveInstanceFromMonitorGroup detail param
type RemoveInstanceFromMonitorGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveInstanceFromMonitorGroupParam RemoveInstanceFromMonitorGroup request param
type RemoveInstanceFromMonitorGroupParam struct {
	BaseParam
	Params RemoveInstanceFromMonitorGroupParamDetail `json:"removeInstanceFromMonitorGroup"`
}
// CleanQueueParamDetail CleanQueue detail param
type CleanQueueParamDetail struct {
	SignatureName string `json:"signatureName" validate:"required"`
	TaskIndex *int `json:"taskIndex,omitempty"`
	IsCleanUp *bool `json:"isCleanUp,omitempty"`
	IsRunningTask *bool `json:"isRunningTask,omitempty"`
	ManagementiUuid *string `json:"managementiUuid,omitempty"`
}

// CleanQueueParam CleanQueue request param
type CleanQueueParam struct {
	BaseParam
	Params CleanQueueParamDetail `json:"cleanQueue"`
}
// RemoveAccessControlListFromLoadBalancerParamDetail RemoveAccessControlListFromLoadBalancer detail param
type RemoveAccessControlListFromLoadBalancerParamDetail struct {
	AclUuids []string `json:"aclUuids" validate:"required"`
	ServerGroupUuids []string `json:"serverGroupUuids,omitempty"`
}

// RemoveAccessControlListFromLoadBalancerParam RemoveAccessControlListFromLoadBalancer request param
type RemoveAccessControlListFromLoadBalancerParam struct {
	BaseParam
	Params RemoveAccessControlListFromLoadBalancerParamDetail `json:"removeAccessControlListFromLoadBalancer"`
}
// RemoveLabelFromEventSubscriptionParamDetail RemoveLabelFromEventSubscription detail param
type RemoveLabelFromEventSubscriptionParamDetail struct {
}

// RemoveLabelFromEventSubscriptionParam RemoveLabelFromEventSubscription request param
type RemoveLabelFromEventSubscriptionParam struct {
	BaseParam
	Params RemoveLabelFromEventSubscriptionParamDetail `json:"removeLabelFromEventSubscription"`
}
// SdnControllerRemoveHostParamDetail SdnControllerRemoveHost detail param
type SdnControllerRemoveHostParamDetail struct {
	VSwitchType *string `json:"vSwitchType,omitempty"`
}

// SdnControllerRemoveHostParam SdnControllerRemoveHost request param
type SdnControllerRemoveHostParam struct {
	BaseParam
	Params SdnControllerRemoveHostParamDetail `json:"sdnControllerRemoveHost"`
}
// DetachCCSCertificateFromUserParamDetail DetachCCSCertificateFromUser detail param
type DetachCCSCertificateFromUserParamDetail struct {
}

// DetachCCSCertificateFromUserParam DetachCCSCertificateFromUser request param
type DetachCCSCertificateFromUserParam struct {
	BaseParam
	Params DetachCCSCertificateFromUserParamDetail `json:"params"`
}
// GetManagementNodeOSParamDetail GetManagementNodeOS detail param
type GetManagementNodeOSParamDetail struct {
}

// GetManagementNodeOSParam GetManagementNodeOS request param
type GetManagementNodeOSParam struct {
	BaseParam
	Params GetManagementNodeOSParamDetail `json:"getManagementNodeOS"`
}
// CreateLdapBindingParamDetail CreateLdapBinding detail param
type CreateLdapBindingParamDetail struct {
	LdapUid string `json:"ldapUid" validate:"required"`
	AccountUuid string `json:"accountUuid" validate:"required"`
}

// CreateLdapBindingParam CreateLdapBinding request param
type CreateLdapBindingParam struct {
	BaseParam
	Params CreateLdapBindingParamDetail `json:"params"`
}
// ExecuteDRSSchedulingParamDetail ExecuteDRSScheduling detail param
type ExecuteDRSSchedulingParamDetail struct {
}

// ExecuteDRSSchedulingParam ExecuteDRSScheduling request param
type ExecuteDRSSchedulingParam struct {
	BaseParam
	Params ExecuteDRSSchedulingParamDetail `json:"executeDRSScheduling"`
}
// CreateEcsSecurityGroupRuleRemoteParamDetail CreateEcsSecurityGroupRuleRemote detail param
type CreateEcsSecurityGroupRuleRemoteParamDetail struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	Direction string `json:"direction" validate:"required"`
	Protocol string `json:"protocol" validate:"required"`
	PortRange string `json:"portRange" validate:"required"`
	Cidr string `json:"cidr" validate:"required"`
	Policy *string `json:"policy,omitempty"`
	Nictype *string `json:"nictype,omitempty"`
	Priority *int `json:"priority,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsSecurityGroupRuleRemoteParam CreateEcsSecurityGroupRuleRemote request param
type CreateEcsSecurityGroupRuleRemoteParam struct {
	BaseParam
	Params CreateEcsSecurityGroupRuleRemoteParamDetail `json:"params"`
}
// GetVmQgaParamDetail GetVmQga detail param
type GetVmQgaParamDetail struct {
}

// GetVmQgaParam GetVmQga request param
type GetVmQgaParam struct {
	BaseParam
	Params GetVmQgaParamDetail `json:"getVmQga"`
}
// PreviewResourceStackParamDetail PreviewResourceStack detail param
type PreviewResourceStackParamDetail struct {
	Type *string `json:"type,omitempty"`
	TemplateContent *string `json:"templateContent,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
	PreParameters *string `json:"preParameters,omitempty"`
}

// PreviewResourceStackParam PreviewResourceStack request param
type PreviewResourceStackParam struct {
	BaseParam
	Params PreviewResourceStackParamDetail `json:"params"`
}
// GetVmvNUMATopologyParamDetail GetVmvNUMATopology detail param
type GetVmvNUMATopologyParamDetail struct {
}

// GetVmvNUMATopologyParam GetVmvNUMATopology request param
type GetVmvNUMATopologyParam struct {
	BaseParam
	Params GetVmvNUMATopologyParamDetail `json:"getVmvNUMATopology"`
}
// RemoveSchedulerJobsFromSchedulerJobGroupParamDetail RemoveSchedulerJobsFromSchedulerJobGroup detail param
type RemoveSchedulerJobsFromSchedulerJobGroupParamDetail struct {
	SchedulerJobUuids []string `json:"schedulerJobUuids" validate:"required"`
}

// RemoveSchedulerJobsFromSchedulerJobGroupParam RemoveSchedulerJobsFromSchedulerJobGroup request param
type RemoveSchedulerJobsFromSchedulerJobGroupParam struct {
	BaseParam
	Params RemoveSchedulerJobsFromSchedulerJobGroupParamDetail `json:"removeSchedulerJobsFromSchedulerJobGroup"`
}
// ChangeTicketStatusParamDetail ChangeTicketStatus detail param
type ChangeTicketStatusParamDetail struct {
	StatusEvent string `json:"statusEvent" validate:"required"`
	Comment *string `json:"comment,omitempty"`
}

// ChangeTicketStatusParam ChangeTicketStatus request param
type ChangeTicketStatusParam struct {
	BaseParam
	Params ChangeTicketStatusParamDetail `json:"changeTicketStatus"`
}
// GetHostPhysicalMemoryFactsParamDetail GetHostPhysicalMemoryFacts detail param
type GetHostPhysicalMemoryFactsParamDetail struct {
}

// GetHostPhysicalMemoryFactsParam GetHostPhysicalMemoryFacts request param
type GetHostPhysicalMemoryFactsParam struct {
	BaseParam
	Params GetHostPhysicalMemoryFactsParamDetail `json:"getHostPhysicalMemoryFacts"`
}
// GetLicenseInfoParamDetail GetLicenseInfo detail param
type GetLicenseInfoParamDetail struct {
	AdditionSession *string `json:"additionSession,omitempty"`
}

// GetLicenseInfoParam GetLicenseInfo request param
type GetLicenseInfoParam struct {
	BaseParam
	Params GetLicenseInfoParamDetail `json:"getLicenseInfo"`
}
// ChangeSchedulerStateParamDetail ChangeSchedulerState detail param
type ChangeSchedulerStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSchedulerStateParam ChangeSchedulerState request param
type ChangeSchedulerStateParam struct {
	BaseParam
	Params ChangeSchedulerStateParamDetail `json:"changeSchedulerState"`
}
// AttachPriceTableToAccountParamDetail AttachPriceTableToAccount detail param
type AttachPriceTableToAccountParamDetail struct {
}

// AttachPriceTableToAccountParam AttachPriceTableToAccount request param
type AttachPriceTableToAccountParam struct {
	BaseParam
	Params AttachPriceTableToAccountParamDetail `json:"params"`
}
// GenerateMdevDevicesParamDetail GenerateMdevDevices detail param
type GenerateMdevDevicesParamDetail struct {
	MdevSpecUuid string `json:"mdevSpecUuid" validate:"required"`
}

// GenerateMdevDevicesParam GenerateMdevDevices request param
type GenerateMdevDevicesParam struct {
	BaseParam
	Params GenerateMdevDevicesParamDetail `json:"generateMdevDevices"`
}
// PreviewResourceFromAppParamDetail PreviewResourceFromApp detail param
type PreviewResourceFromAppParamDetail struct {
	AppUuid string `json:"appUuid" validate:"required"`
	Parameters *string `json:"parameters,omitempty"`
}

// PreviewResourceFromAppParam PreviewResourceFromApp request param
type PreviewResourceFromAppParam struct {
	BaseParam
	Params PreviewResourceFromAppParamDetail `json:"params"`
}
// ChangeSNSTopicStateParamDetail ChangeSNSTopicState detail param
type ChangeSNSTopicStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSNSTopicStateParam ChangeSNSTopicState request param
type ChangeSNSTopicStateParam struct {
	BaseParam
	Params ChangeSNSTopicStateParamDetail `json:"changeSNSTopicState"`
}
// AttachScsiLunToVmInstanceParamDetail AttachScsiLunToVmInstance detail param
type AttachScsiLunToVmInstanceParamDetail struct {
	DisableMultiPathAttach *bool `json:"disableMultiPathAttach,omitempty"`
}

// AttachScsiLunToVmInstanceParam AttachScsiLunToVmInstance request param
type AttachScsiLunToVmInstanceParam struct {
	BaseParam
	Params AttachScsiLunToVmInstanceParamDetail `json:"params"`
}
// RemoveRemoteCidrsFromIPsecConnectionParamDetail RemoveRemoteCidrsFromIPsecConnection detail param
type RemoveRemoteCidrsFromIPsecConnectionParamDetail struct {
	PeerCidrs []string `json:"peerCidrs" validate:"required"`
}

// RemoveRemoteCidrsFromIPsecConnectionParam RemoveRemoteCidrsFromIPsecConnection request param
type RemoveRemoteCidrsFromIPsecConnectionParam struct {
	BaseParam
	Params RemoveRemoteCidrsFromIPsecConnectionParamDetail `json:"removeRemoteCidrsFromIPsecConnection"`
}
// GetIAM2ProjectsOfVirtualIDParamDetail GetIAM2ProjectsOfVirtualID detail param
type GetIAM2ProjectsOfVirtualIDParamDetail struct {
}

// GetIAM2ProjectsOfVirtualIDParam GetIAM2ProjectsOfVirtualID request param
type GetIAM2ProjectsOfVirtualIDParam struct {
	BaseParam
	Params GetIAM2ProjectsOfVirtualIDParamDetail `json:"getIAM2ProjectsOfVirtualID"`
}
// AddVmToAffinityGroupParamDetail AddVmToAffinityGroup detail param
type AddVmToAffinityGroupParamDetail struct {
}

// AddVmToAffinityGroupParam AddVmToAffinityGroup request param
type AddVmToAffinityGroupParam struct {
	BaseParam
	Params AddVmToAffinityGroupParamDetail `json:"params"`
}
// GetPrimaryStorageAllocatorStrategiesParamDetail GetPrimaryStorageAllocatorStrategies detail param
type GetPrimaryStorageAllocatorStrategiesParamDetail struct {
}

// GetPrimaryStorageAllocatorStrategiesParam GetPrimaryStorageAllocatorStrategies request param
type GetPrimaryStorageAllocatorStrategiesParam struct {
	BaseParam
	Params GetPrimaryStorageAllocatorStrategiesParamDetail `json:"getPrimaryStorageAllocatorStrategies"`
}
// GetPlatformTimeZoneParamDetail GetPlatformTimeZone detail param
type GetPlatformTimeZoneParamDetail struct {
}

// GetPlatformTimeZoneParam GetPlatformTimeZone request param
type GetPlatformTimeZoneParam struct {
	BaseParam
	Params GetPlatformTimeZoneParamDetail `json:"getPlatformTimeZone"`
}
// DetachPolicyFromUserParamDetail DetachPolicyFromUser detail param
type DetachPolicyFromUserParamDetail struct {
}

// DetachPolicyFromUserParam DetachPolicyFromUser request param
type DetachPolicyFromUserParam struct {
	BaseParam
	Params DetachPolicyFromUserParamDetail `json:"detachPolicyFromUser"`
}
// SetVmInstanceDefaultCdRomParamDetail SetVmInstanceDefaultCdRom detail param
type SetVmInstanceDefaultCdRomParamDetail struct {
}

// SetVmInstanceDefaultCdRomParam SetVmInstanceDefaultCdRom request param
type SetVmInstanceDefaultCdRomParam struct {
	BaseParam
	Params SetVmInstanceDefaultCdRomParamDetail `json:"setVmInstanceDefaultCdRom"`
}
// RefreshSharedblockDeviceCapacityParamDetail RefreshSharedblockDeviceCapacity detail param
type RefreshSharedblockDeviceCapacityParamDetail struct {
}

// RefreshSharedblockDeviceCapacityParam RefreshSharedblockDeviceCapacity request param
type RefreshSharedblockDeviceCapacityParam struct {
	BaseParam
	Params RefreshSharedblockDeviceCapacityParamDetail `json:"params"`
}
// FstrimVmParamDetail FstrimVm detail param
type FstrimVmParamDetail struct {
}

// FstrimVmParam FstrimVm request param
type FstrimVmParam struct {
	BaseParam
	Params FstrimVmParamDetail `json:"params"`
}
// DetachL2NetworkFromClusterParamDetail DetachL2NetworkFromCluster detail param
type DetachL2NetworkFromClusterParamDetail struct {
}

// DetachL2NetworkFromClusterParam DetachL2NetworkFromCluster request param
type DetachL2NetworkFromClusterParam struct {
	BaseParam
	Params DetachL2NetworkFromClusterParamDetail `json:"detachL2NetworkFromCluster"`
}
// SyncAINginxConfigurationParamDetail SyncAINginxConfiguration detail param
type SyncAINginxConfigurationParamDetail struct {
	GroupUuids []string `json:"groupUuids,omitempty"`
	DryRun *bool `json:"dryRun,omitempty"`
	SyncAll *bool `json:"syncAll,omitempty"`
}

// SyncAINginxConfigurationParam SyncAINginxConfiguration request param
type SyncAINginxConfigurationParam struct {
	BaseParam
	Params SyncAINginxConfigurationParamDetail `json:"param"`
}
// MatchModelServiceTemplateWithModelParamDetail MatchModelServiceTemplateWithModel detail param
type MatchModelServiceTemplateWithModelParamDetail struct {
	ModelUuids []string `json:"modelUuids" validate:"required"`
	ModelServiceUuids []string `json:"modelServiceUuids" validate:"required"`
	InstanceNumber int `json:"instanceNumber" validate:"required"`
	Uuid string `json:"uuid,omitempty"`
	Description *string `json:"description,omitempty"`
	ModelUuid *string `json:"modelUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	VmImageUuid *string `json:"vmImageUuid,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	DatasetUuids []string `json:"datasetUuids,omitempty"`
	ModelServiceGroupUuids []string `json:"modelServiceGroupUuids,omitempty"`
	DockerImage *string `json:"dockerImage,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	Name string `json:"name" validate:"required"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	StartupParameters map[string]string `json:"startupParameters,omitempty"`
	Type string `json:"type" validate:"required"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	ServiceBootUptime *int `json:"serviceBootUptime,omitempty"`
	ServiceLivez *string `json:"serviceLivez,omitempty"`
	ServiceReadyz *string `json:"serviceReadyz,omitempty"`
	RootDiskOfferingUuid *string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize *int64 `json:"rootDiskSize,omitempty"`
	ProjectUuid *string `json:"projectUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// MatchModelServiceTemplateWithModelParam MatchModelServiceTemplateWithModel request param
type MatchModelServiceTemplateWithModelParam struct {
	BaseParam
	Params MatchModelServiceTemplateWithModelParamDetail `json:"param"`
}
// ChangeVmNicStateParamDetail ChangeVmNicState detail param
type ChangeVmNicStateParamDetail struct {
	State string `json:"state" validate:"required"`
}

// ChangeVmNicStateParam ChangeVmNicState request param
type ChangeVmNicStateParam struct {
	BaseParam
	Params ChangeVmNicStateParamDetail `json:"changeVmNicState"`
}
// AddPolicyStatementsToRoleParamDetail AddPolicyStatementsToRole detail param
type AddPolicyStatementsToRoleParamDetail struct {
	Statements []PolicyStatementParam `json:"statements" validate:"required"`
}

// AddPolicyStatementsToRoleParam AddPolicyStatementsToRole request param
type AddPolicyStatementsToRoleParam struct {
	BaseParam
	Params AddPolicyStatementsToRoleParamDetail `json:"params"`
}
// UnprotectVmInstanceRecoveryPointParamDetail UnprotectVmInstanceRecoveryPoint detail param
type UnprotectVmInstanceRecoveryPointParamDetail struct {
	GroupId int64 `json:"groupId" validate:"required"`
}

// UnprotectVmInstanceRecoveryPointParam UnprotectVmInstanceRecoveryPoint request param
type UnprotectVmInstanceRecoveryPointParam struct {
	BaseParam
	Params UnprotectVmInstanceRecoveryPointParamDetail `json:"unprotectVmInstanceRecoveryPoint"`
}
// DetachVipFromVpcSharedQosParamDetail DetachVipFromVpcSharedQos detail param
type DetachVipFromVpcSharedQosParamDetail struct {
	VipLists []string `json:"vipLists,omitempty"`
	VipUuids []string `json:"vipUuids,omitempty"`
}

// DetachVipFromVpcSharedQosParam DetachVipFromVpcSharedQos request param
type DetachVipFromVpcSharedQosParam struct {
	BaseParam
	Params DetachVipFromVpcSharedQosParamDetail `json:"detachVipFromVpcSharedQos"`
}
// ApplyRuleSetChangesParamDetail ApplyRuleSetChanges detail param
type ApplyRuleSetChangesParamDetail struct {
}

// ApplyRuleSetChangesParam ApplyRuleSetChanges request param
type ApplyRuleSetChangesParam struct {
	BaseParam
	Params ApplyRuleSetChangesParamDetail `json:"applyRuleSetChanges"`
}
// PrimaryStorageMigrateVmParamDetail PrimaryStorageMigrateVm detail param
type PrimaryStorageMigrateVmParamDetail struct {
	DstPrimaryStorageUuid string `json:"dstPrimaryStorageUuid" validate:"required"`
	DstHostUuid *string `json:"dstHostUuid,omitempty"`
	WithDataVolumes *bool `json:"withDataVolumes,omitempty"`
	DataVolumeUuids []string `json:"dataVolumeUuids,omitempty"`
	WithSnapshots *bool `json:"withSnapshots,omitempty"`
	DownTime *int `json:"downTime,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	Bandwidth *int64 `json:"bandwidth,omitempty"`
	VolumeProvisioningStrategy *string `json:"volumeProvisioningStrategy,omitempty"`
}

// PrimaryStorageMigrateVmParam PrimaryStorageMigrateVm request param
type PrimaryStorageMigrateVmParam struct {
	BaseParam
	Params PrimaryStorageMigrateVmParamDetail `json:"primaryStorageMigrateVm"`
}
// RecoverDatabaseFromBackupParamDetail RecoverDatabaseFromBackup detail param
type RecoverDatabaseFromBackupParamDetail struct {
	Uuid string `json:"uuid,omitempty"`
	BackupStorageUrl *string `json:"backupStorageUrl,omitempty"`
	BackupInstallPath *string `json:"backupInstallPath,omitempty"`
	MysqlRootPassword string `json:"mysqlRootPassword" validate:"required"`
}

// RecoverDatabaseFromBackupParam RecoverDatabaseFromBackup request param
type RecoverDatabaseFromBackupParam struct {
	BaseParam
	Params RecoverDatabaseFromBackupParamDetail `json:"recoverDatabaseFromBackup"`
}
// CreateIAM2TickFlowCollectionParamDetail CreateIAM2TickFlowCollection detail param
type CreateIAM2TickFlowCollectionParamDetail struct {
	Flows []CreateIAM2TickFlowCollection_IAM2FlowStructParam `json:"flows,omitempty"`
	ProjectUuid string `json:"projectUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	TicketTypeUuids []string `json:"ticketTypeUuids,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2TickFlowCollectionParam CreateIAM2TickFlowCollection request param
type CreateIAM2TickFlowCollectionParam struct {
	BaseParam
	Params CreateIAM2TickFlowCollectionParamDetail `json:"params"`
}
// UngenerateMdevDevicesParamDetail UngenerateMdevDevices detail param
type UngenerateMdevDevicesParamDetail struct {
}

// UngenerateMdevDevicesParam UngenerateMdevDevices request param
type UngenerateMdevDevicesParam struct {
	BaseParam
	Params UngenerateMdevDevicesParamDetail `json:"ungenerateMdevDevices"`
}
// MoveDirectoryParamDetail MoveDirectory detail param
type MoveDirectoryParamDetail struct {
	TargetParentUuid string `json:"targetParentUuid" validate:"required"`
	DirectoryUuid string `json:"directoryUuid" validate:"required"`
}

// MoveDirectoryParam MoveDirectory request param
type MoveDirectoryParam struct {
	BaseParam
	Params MoveDirectoryParamDetail `json:"moveDirectory"`
}
// GetVRouterOspfNeighborParamDetail GetVRouterOspfNeighbor detail param
type GetVRouterOspfNeighborParamDetail struct {
}

// GetVRouterOspfNeighborParam GetVRouterOspfNeighbor request param
type GetVRouterOspfNeighborParam struct {
	BaseParam
	Params GetVRouterOspfNeighborParamDetail `json:"getVRouterOspfNeighbor"`
}
// CreateVpcVRouterParamDetail CreateVpcVRouter detail param
type CreateVpcVRouterParamDetail struct {
	Name string `json:"name" validate:"required"`
	VirtualRouterOfferingUuid string `json:"virtualRouterOfferingUuid" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcVRouterParam CreateVpcVRouter request param
type CreateVpcVRouterParam struct {
	BaseParam
	Params CreateVpcVRouterParamDetail `json:"params"`
}
// SyncEcsInstanceFromRemoteParamDetail SyncEcsInstanceFromRemote detail param
type SyncEcsInstanceFromRemoteParamDetail struct {
	OnlyZstack *bool `json:"onlyZstack,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsInstanceFromRemoteParam SyncEcsInstanceFromRemote request param
type SyncEcsInstanceFromRemoteParam struct {
	BaseParam
	Params SyncEcsInstanceFromRemoteParamDetail `json:"params"`
}
// GetMdevDeviceSpecCandidatesParamDetail GetMdevDeviceSpecCandidates detail param
type GetMdevDeviceSpecCandidatesParamDetail struct {
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	VmInstanceUuid *string `json:"vmInstanceUuid,omitempty"`
	VmInstanceUuids []string `json:"vmInstanceUuids,omitempty"`
	Types []string `json:"types,omitempty"`
}

// GetMdevDeviceSpecCandidatesParam GetMdevDeviceSpecCandidates request param
type GetMdevDeviceSpecCandidatesParam struct {
	BaseParam
	Params GetMdevDeviceSpecCandidatesParamDetail `json:"getMdevDeviceSpecCandidates"`
}
// GetFlowMeterRouterIdParamDetail GetFlowMeterRouterId detail param
type GetFlowMeterRouterIdParamDetail struct {
}

// GetFlowMeterRouterIdParam GetFlowMeterRouterId request param
type GetFlowMeterRouterIdParam struct {
	BaseParam
	Params GetFlowMeterRouterIdParamDetail `json:"getFlowMeterRouterId"`
}
// GetPciDeviceCandidatesForNewCreateVmParamDetail GetPciDeviceCandidatesForNewCreateVm detail param
type GetPciDeviceCandidatesForNewCreateVmParamDetail struct {
	HostUuid *string `json:"hostUuid,omitempty"`
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	Types []string `json:"types,omitempty"`
}

// GetPciDeviceCandidatesForNewCreateVmParam GetPciDeviceCandidatesForNewCreateVm request param
type GetPciDeviceCandidatesForNewCreateVmParam struct {
	BaseParam
	Params GetPciDeviceCandidatesForNewCreateVmParamDetail `json:"getPciDeviceCandidatesForNewCreateVm"`
}
// GetHostTaskParamDetail GetHostTask detail param
type GetHostTaskParamDetail struct {
	HostUuids []string `json:"hostUuids" validate:"required"`
	SyncSignatures []string `json:"syncSignatures,omitempty"`
}

// GetHostTaskParam GetHostTask request param
type GetHostTaskParam struct {
	BaseParam
	Params GetHostTaskParamDetail `json:"getHostTask"`
}
// AddResourceToIAM2ProjectParamDetail AddResourceToIAM2Project detail param
type AddResourceToIAM2ProjectParamDetail struct {
	ProjectUuid string `json:"projectUuid" validate:"required"`
	ResourceTemplates []string `json:"resourceTemplates" validate:"required"`
}

// AddResourceToIAM2ProjectParam AddResourceToIAM2Project request param
type AddResourceToIAM2ProjectParam struct {
	BaseParam
	Params AddResourceToIAM2ProjectParamDetail `json:"params"`
}
// GetAlarmDataParamDetail GetAlarmData detail param
type GetAlarmDataParamDetail struct {
	StartTime *int64 `json:"startTime,omitempty"`
	EndTime *int64 `json:"endTime,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Conditions []string `json:"conditions,omitempty"`
	Count *bool `json:"count,omitempty"`
	ExcludeOtherAccount *bool `json:"excludeOtherAccount,omitempty"`
	Start *int `json:"start,omitempty"`
	EndpointUuid *string `json:"endpointUuid,omitempty"`
}

// GetAlarmDataParam GetAlarmData request param
type GetAlarmDataParam struct {
	BaseParam
	Params GetAlarmDataParamDetail `json:"getAlarmData"`
}
// ChangeV2VConversionHostStateParamDetail ChangeV2VConversionHostState detail param
type ChangeV2VConversionHostStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeV2VConversionHostStateParam ChangeV2VConversionHostState request param
type ChangeV2VConversionHostStateParam struct {
	BaseParam
	Params ChangeV2VConversionHostStateParamDetail `json:"changeV2VConversionHostState"`
}
// RecoverResourceSplitBrainParamDetail RecoverResourceSplitBrain detail param
type RecoverResourceSplitBrainParamDetail struct {
	HostUuid string `json:"hostUuid" validate:"required"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	ForceRecover *bool `json:"forceRecover,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

// RecoverResourceSplitBrainParam RecoverResourceSplitBrain request param
type RecoverResourceSplitBrainParam struct {
	BaseParam
	Params RecoverResourceSplitBrainParamDetail `json:"recoverResourceSplitBrain"`
}
// IsOpensourceVersionParamDetail IsOpensourceVersion detail param
type IsOpensourceVersionParamDetail struct {
}

// IsOpensourceVersionParam IsOpensourceVersion request param
type IsOpensourceVersionParam struct {
	BaseParam
	Params IsOpensourceVersionParamDetail `json:"isOpensourceVersion"`
}
// CreateBareMetal2InstanceFromVolumeBackupParamDetail CreateBareMetal2InstanceFromVolumeBackup detail param
type CreateBareMetal2InstanceFromVolumeBackupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	GatewayUuid *string `json:"gatewayUuid,omitempty"`
	ChassisUuid *string `json:"chassisUuid,omitempty"`
	ChassisOfferingUuid *string `json:"chassisOfferingUuid,omitempty"`
	ChassisDiskUuid *string `json:"chassisDiskUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	GatewayAllocatorStrategy *string `json:"gatewayAllocatorStrategy,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBareMetal2InstanceFromVolumeBackupParam CreateBareMetal2InstanceFromVolumeBackup request param
type CreateBareMetal2InstanceFromVolumeBackupParam struct {
	BaseParam
	Params CreateBareMetal2InstanceFromVolumeBackupParamDetail `json:"params"`
}
// CreateEcsInstanceFromEcsImageParamDetail CreateEcsInstanceFromEcsImage detail param
type CreateEcsInstanceFromEcsImageParamDetail struct {
	EcsRootVolumeType *string `json:"ecsRootVolumeType,omitempty"`
	Description *string `json:"description,omitempty"`
	EcsRootVolumeGBSize *int64 `json:"ecsRootVolumeGBSize,omitempty"`
	CreateMode *string `json:"createMode,omitempty"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	AllocatePublicIp *string `json:"allocatePublicIp,omitempty"`
	EcsConsolePassword *string `json:"ecsConsolePassword,omitempty"`
	Name string `json:"name" validate:"required"`
	EcsImageUuid string `json:"ecsImageUuid" validate:"required"`
	InstanceOfferingUuid *string `json:"instanceOfferingUuid,omitempty"`
	InstanceType *string `json:"instanceType,omitempty"`
	EcsVSwitchUuid string `json:"ecsVSwitchUuid" validate:"required"`
	EcsSecurityGroupUuid string `json:"ecsSecurityGroupUuid" validate:"required"`
	EcsRootPassword string `json:"ecsRootPassword" validate:"required"`
	EcsBandWidth *int64 `json:"ecsBandWidth,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsInstanceFromEcsImageParam CreateEcsInstanceFromEcsImage request param
type CreateEcsInstanceFromEcsImageParam struct {
	BaseParam
	Params CreateEcsInstanceFromEcsImageParamDetail `json:"params"`
}
// GetResourceFromResourceStackParamDetail GetResourceFromResourceStack detail param
type GetResourceFromResourceStackParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetResourceFromResourceStackParam GetResourceFromResourceStack request param
type GetResourceFromResourceStackParam struct {
	BaseParam
	Params GetResourceFromResourceStackParamDetail `json:"getResourceFromResourceStack"`
}
// MoveResourcesToDirectoryParamDetail MoveResourcesToDirectory detail param
type MoveResourcesToDirectoryParamDetail struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	DirectoryUuid string `json:"directoryUuid" validate:"required"`
}

// MoveResourcesToDirectoryParam MoveResourcesToDirectory request param
type MoveResourcesToDirectoryParam struct {
	BaseParam
	Params MoveResourcesToDirectoryParamDetail `json:"moveResourcesToDirectory"`
}
// GetSupportedCloudFormationResourcesParamDetail GetSupportedCloudFormationResources detail param
type GetSupportedCloudFormationResourcesParamDetail struct {
	Version *string `json:"version,omitempty"`
	Type *string `json:"type,omitempty"`
}

// GetSupportedCloudFormationResourcesParam GetSupportedCloudFormationResources request param
type GetSupportedCloudFormationResourcesParam struct {
	BaseParam
	Params GetSupportedCloudFormationResourcesParamDetail `json:"getSupportedCloudFormationResources"`
}
// DeleteIdentityZoneInLocalParamDetail DeleteIdentityZoneInLocal detail param
type DeleteIdentityZoneInLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteIdentityZoneInLocalParam DeleteIdentityZoneInLocal request param
type DeleteIdentityZoneInLocalParam struct {
	BaseParam
	Params DeleteIdentityZoneInLocalParamDetail `json:"deleteIdentityZoneInLocal"`
}
// ChangeIAM2VirtualIDStateParamDetail ChangeIAM2VirtualIDState detail param
type ChangeIAM2VirtualIDStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeIAM2VirtualIDStateParam ChangeIAM2VirtualIDState request param
type ChangeIAM2VirtualIDStateParam struct {
	BaseParam
	Params ChangeIAM2VirtualIDStateParamDetail `json:"changeIAM2VirtualIDState"`
}
// UnregisterLicenseServerParamDetail UnregisterLicenseServer detail param
type UnregisterLicenseServerParamDetail struct {
	ClientAuthorizedNodeUuid *string `json:"clientAuthorizedNodeUuid,omitempty"`
}

// UnregisterLicenseServerParam UnregisterLicenseServer request param
type UnregisterLicenseServerParam struct {
	BaseParam
	Params UnregisterLicenseServerParamDetail `json:"unregisterLicenseServer"`
}
// CreateVmUserDefinedXmlHookScriptParamDetail CreateVmUserDefinedXmlHookScript detail param
type CreateVmUserDefinedXmlHookScriptParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	HookScript string `json:"hookScript" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmUserDefinedXmlHookScriptParam CreateVmUserDefinedXmlHookScript request param
type CreateVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params CreateVmUserDefinedXmlHookScriptParamDetail `json:"params"`
}
// UpgradeToLicenseServerParamDetail UpgradeToLicenseServer detail param
type UpgradeToLicenseServerParamDetail struct {
}

// UpgradeToLicenseServerParam UpgradeToLicenseServer request param
type UpgradeToLicenseServerParam struct {
	BaseParam
	Params UpgradeToLicenseServerParamDetail `json:"params"`
}
// DetachAppBuildSystemToZoneParamDetail DetachAppBuildSystemToZone detail param
type DetachAppBuildSystemToZoneParamDetail struct {
}

// DetachAppBuildSystemToZoneParam DetachAppBuildSystemToZone request param
type DetachAppBuildSystemToZoneParam struct {
	BaseParam
	Params DetachAppBuildSystemToZoneParamDetail `json:"detachAppBuildSystemToZone"`
}
// GetAppBuildSystemCapacityParamDetail GetAppBuildSystemCapacity detail param
type GetAppBuildSystemCapacityParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetAppBuildSystemCapacityParam GetAppBuildSystemCapacity request param
type GetAppBuildSystemCapacityParam struct {
	BaseParam
	Params GetAppBuildSystemCapacityParamDetail `json:"getAppBuildSystemCapacity"`
}
// ChangeBaremetalChassisStateParamDetail ChangeBaremetalChassisState detail param
type ChangeBaremetalChassisStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBaremetalChassisStateParam ChangeBaremetalChassisState request param
type ChangeBaremetalChassisStateParam struct {
	BaseParam
	Params ChangeBaremetalChassisStateParamDetail `json:"changeBaremetalChassisState"`
}
// GetAttachableVpcL3NetworkParamDetail GetAttachableVpcL3Network detail param
type GetAttachableVpcL3NetworkParamDetail struct {
}

// GetAttachableVpcL3NetworkParam GetAttachableVpcL3Network request param
type GetAttachableVpcL3NetworkParam struct {
	BaseParam
	Params GetAttachableVpcL3NetworkParamDetail `json:"params"`
}
// ChangeNfvInstGroupOperationModeParamDetail ChangeNfvInstGroupOperationMode detail param
type ChangeNfvInstGroupOperationModeParamDetail struct {
	OperationMode string `json:"operationMode" validate:"required"`
}

// ChangeNfvInstGroupOperationModeParam ChangeNfvInstGroupOperationMode request param
type ChangeNfvInstGroupOperationModeParam struct {
	BaseParam
	Params ChangeNfvInstGroupOperationModeParamDetail `json:"changeNfvInstGroupOperationMode"`
}
// GetL3NetworkMtuParamDetail GetL3NetworkMtu detail param
type GetL3NetworkMtuParamDetail struct {
}

// GetL3NetworkMtuParam GetL3NetworkMtu request param
type GetL3NetworkMtuParam struct {
	BaseParam
	Params GetL3NetworkMtuParamDetail `json:"getL3NetworkMtu"`
}
// AttachVipToLoadBalancerParamDetail AttachVipToLoadBalancer detail param
type AttachVipToLoadBalancerParamDetail struct {
}

// AttachVipToLoadBalancerParam AttachVipToLoadBalancer request param
type AttachVipToLoadBalancerParam struct {
	BaseParam
	Params AttachVipToLoadBalancerParamDetail `json:"params"`
}
// UpdateSecurityGroupRulePriorityParamDetail UpdateSecurityGroupRulePriority detail param
type UpdateSecurityGroupRulePriorityParamDetail struct {
	Type string `json:"type" validate:"required"`
	Rules []UpdateSecurityGroupRulePriority_SecurityGroupRulePriorityAOParam `json:"rules" validate:"required"`
}

// UpdateSecurityGroupRulePriorityParam UpdateSecurityGroupRulePriority request param
type UpdateSecurityGroupRulePriorityParam struct {
	BaseParam
	Params UpdateSecurityGroupRulePriorityParamDetail `json:"updateSecurityGroupRulePriority"`
}
// AddDnsToL3NetworkParamDetail AddDnsToL3Network detail param
type AddDnsToL3NetworkParamDetail struct {
	Dns string `json:"dns" validate:"required"`
}

// AddDnsToL3NetworkParam AddDnsToL3Network request param
type AddDnsToL3NetworkParam struct {
	BaseParam
	Params AddDnsToL3NetworkParamDetail `json:"params"`
}
// SetVmMonitorNumberParamDetail SetVmMonitorNumber detail param
type SetVmMonitorNumberParamDetail struct {
	MonitorNumber int `json:"monitorNumber" validate:"required"`
}

// SetVmMonitorNumberParam SetVmMonitorNumber request param
type SetVmMonitorNumberParam struct {
	BaseParam
	Params SetVmMonitorNumberParamDetail `json:"setVmMonitorNumber"`
}
// ChangeLoadBalancerBackendServerParamDetail ChangeLoadBalancerBackendServer detail param
type ChangeLoadBalancerBackendServerParamDetail struct {
	VmNics []interface{} `json:"vmNics,omitempty"`
	Servers []interface{} `json:"servers,omitempty"`
}

// ChangeLoadBalancerBackendServerParam ChangeLoadBalancerBackendServer request param
type ChangeLoadBalancerBackendServerParam struct {
	BaseParam
	Params ChangeLoadBalancerBackendServerParamDetail `json:"changeLoadBalancerBackendServer"`
}
// CreateEcsVSwitchRemoteParamDetail CreateEcsVSwitchRemote detail param
type CreateEcsVSwitchRemoteParamDetail struct {
	VpcUuid string `json:"vpcUuid" validate:"required"`
	IdentityZoneUuid string `json:"identityZoneUuid" validate:"required"`
	CidrBlock string `json:"cidrBlock" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsVSwitchRemoteParam CreateEcsVSwitchRemote request param
type CreateEcsVSwitchRemoteParam struct {
	BaseParam
	Params CreateEcsVSwitchRemoteParamDetail `json:"params"`
}
// GetVmMigrationCandidateHostsParamDetail GetVmMigrationCandidateHosts detail param
type GetVmMigrationCandidateHostsParamDetail struct {
}

// GetVmMigrationCandidateHostsParam GetVmMigrationCandidateHosts request param
type GetVmMigrationCandidateHostsParam struct {
	BaseParam
	Params GetVmMigrationCandidateHostsParamDetail `json:"getVmMigrationCandidateHosts"`
}
// GetCandidateL3NetworksForIpSecConnectionParamDetail GetCandidateL3NetworksForIpSecConnection detail param
type GetCandidateL3NetworksForIpSecConnectionParamDetail struct {
	Uuid string `json:"uuid,omitempty"`
	PublicL3Uuid *string `json:"publicL3Uuid,omitempty"`
	VipUuid *string `json:"vipUuid,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetCandidateL3NetworksForIpSecConnectionParam GetCandidateL3NetworksForIpSecConnection request param
type GetCandidateL3NetworksForIpSecConnectionParam struct {
	BaseParam
	Params GetCandidateL3NetworksForIpSecConnectionParamDetail `json:"getCandidateL3NetworksForIpSecConnection"`
}
// UpdateHostNetworkServiceTypeParamDetail UpdateHostNetworkServiceType detail param
type UpdateHostNetworkServiceTypeParamDetail struct {
	ServiceType string `json:"serviceType" validate:"required"`
	System *bool `json:"system,omitempty"`
}

// UpdateHostNetworkServiceTypeParam UpdateHostNetworkServiceType request param
type UpdateHostNetworkServiceTypeParam struct {
	BaseParam
	Params UpdateHostNetworkServiceTypeParamDetail `json:"updateHostNetworkServiceType"`
}
// SNSMicrosoftTeamsTestConnectionParamDetail SNSMicrosoftTeamsTestConnection detail param
type SNSMicrosoftTeamsTestConnectionParamDetail struct {
	Url *string `json:"url,omitempty"`
	TestMsg string `json:"testMsg" validate:"required"`
	EndpointUuid *string `json:"endpointUuid,omitempty"`
}

// SNSMicrosoftTeamsTestConnectionParam SNSMicrosoftTeamsTestConnection request param
type SNSMicrosoftTeamsTestConnectionParam struct {
	BaseParam
	Params SNSMicrosoftTeamsTestConnectionParamDetail `json:"params"`
}
// GetLatestGuestToolsForVmParamDetail GetLatestGuestToolsForVm detail param
type GetLatestGuestToolsForVmParamDetail struct {
}

// GetLatestGuestToolsForVmParam GetLatestGuestToolsForVm request param
type GetLatestGuestToolsForVmParam struct {
	BaseParam
	Params GetLatestGuestToolsForVmParamDetail `json:"getLatestGuestToolsForVm"`
}
// CreateVpcUserVpnGatewayRemoteParamDetail CreateVpcUserVpnGatewayRemote detail param
type CreateVpcUserVpnGatewayRemoteParamDetail struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	Ip string `json:"ip" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcUserVpnGatewayRemoteParam CreateVpcUserVpnGatewayRemote request param
type CreateVpcUserVpnGatewayRemoteParam struct {
	BaseParam
	Params CreateVpcUserVpnGatewayRemoteParamDetail `json:"params"`
}
// CreateOssBackupBucketRemoteParamDetail CreateOssBackupBucketRemote detail param
type CreateOssBackupBucketRemoteParamDetail struct {
	RegionId string `json:"regionId" validate:"required"`
	OssDomain *string `json:"ossDomain,omitempty"`
	OssKey *string `json:"ossKey,omitempty"`
	OssSecret *string `json:"ossSecret,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateOssBackupBucketRemoteParam CreateOssBackupBucketRemote request param
type CreateOssBackupBucketRemoteParam struct {
	BaseParam
	Params CreateOssBackupBucketRemoteParamDetail `json:"params"`
}
// PowerOffBaremetalChassisParamDetail PowerOffBaremetalChassis detail param
type PowerOffBaremetalChassisParamDetail struct {
}

// PowerOffBaremetalChassisParam PowerOffBaremetalChassis request param
type PowerOffBaremetalChassisParam struct {
	BaseParam
	Params PowerOffBaremetalChassisParamDetail `json:"powerOffBaremetalChassis"`
}
// GetCandidateInterfaceVlanIdsParamDetail GetCandidateInterfaceVlanIds detail param
type GetCandidateInterfaceVlanIdsParamDetail struct {
	InterfaceUuids []string `json:"interfaceUuids" validate:"required"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetCandidateInterfaceVlanIdsParam GetCandidateInterfaceVlanIds request param
type GetCandidateInterfaceVlanIdsParam struct {
	BaseParam
	Params GetCandidateInterfaceVlanIdsParamDetail `json:"getCandidateInterfaceVlanIds"`
}
// GetNetworkServiceTypesParamDetail GetNetworkServiceTypes detail param
type GetNetworkServiceTypesParamDetail struct {
}

// GetNetworkServiceTypesParam GetNetworkServiceTypes request param
type GetNetworkServiceTypesParam struct {
	BaseParam
	Params GetNetworkServiceTypesParamDetail `json:"getNetworkServiceTypes"`
}
// DeleteVmUserDefinedXmlParamDetail DeleteVmUserDefinedXml detail param
type DeleteVmUserDefinedXmlParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVmUserDefinedXmlParam DeleteVmUserDefinedXml request param
type DeleteVmUserDefinedXmlParam struct {
	BaseParam
	Params DeleteVmUserDefinedXmlParamDetail `json:"deleteVmUserDefinedXml"`
}
// GetAvailableVpcL3NetworkParamDetail GetAvailableVpcL3Network detail param
type GetAvailableVpcL3NetworkParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
}

// GetAvailableVpcL3NetworkParam GetAvailableVpcL3Network request param
type GetAvailableVpcL3NetworkParam struct {
	BaseParam
	Params GetAvailableVpcL3NetworkParamDetail `json:"getAvailableVpcL3Network"`
}
// GetCurrentTimeParamDetail GetCurrentTime detail param
type GetCurrentTimeParamDetail struct {
}

// GetCurrentTimeParam GetCurrentTime request param
type GetCurrentTimeParam struct {
	BaseParam
	Params GetCurrentTimeParamDetail `json:"getCurrentTime"`
}
// CalculateAccountSpendingParamDetail CalculateAccountSpending detail param
type CalculateAccountSpendingParamDetail struct {
	HypervisorType *string `json:"hypervisorType,omitempty"`
	DateStart *int64 `json:"dateStart,omitempty"`
	DateEnd *int64 `json:"dateEnd,omitempty"`
	Simple *bool `json:"simple,omitempty"`
}

// CalculateAccountSpendingParam CalculateAccountSpending request param
type CalculateAccountSpendingParam struct {
	BaseParam
	Params CalculateAccountSpendingParamDetail `json:"calculateAccountSpending"`
}
// GetVmAttachableL3NetworkParamDetail GetVmAttachableL3Network detail param
type GetVmAttachableL3NetworkParamDetail struct {
}

// GetVmAttachableL3NetworkParam GetVmAttachableL3Network request param
type GetVmAttachableL3NetworkParam struct {
	BaseParam
	Params GetVmAttachableL3NetworkParamDetail `json:"getVmAttachableL3Network"`
}
// UpdateEcsInstanceVncPasswordParamDetail UpdateEcsInstanceVncPassword detail param
type UpdateEcsInstanceVncPasswordParamDetail struct {
	Password string `json:"password" validate:"required"`
}

// UpdateEcsInstanceVncPasswordParam UpdateEcsInstanceVncPassword request param
type UpdateEcsInstanceVncPasswordParam struct {
	BaseParam
	Params UpdateEcsInstanceVncPasswordParamDetail `json:"params"`
}
// SyncChronyServersParamDetail SyncChronyServers detail param
type SyncChronyServersParamDetail struct {
}

// SyncChronyServersParam SyncChronyServers request param
type SyncChronyServersParam struct {
	BaseParam
	Params SyncChronyServersParamDetail `json:"syncChronyServers"`
}
// GetVmInstanceProtectedRecoveryPointsParamDetail GetVmInstanceProtectedRecoveryPoints detail param
type GetVmInstanceProtectedRecoveryPointsParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVmInstanceProtectedRecoveryPointsParam GetVmInstanceProtectedRecoveryPoints request param
type GetVmInstanceProtectedRecoveryPointsParam struct {
	BaseParam
	Params GetVmInstanceProtectedRecoveryPointsParamDetail `json:"getVmInstanceProtectedRecoveryPoints"`
}
// AddVmToVmSchedulingRuleGroupParamDetail AddVmToVmSchedulingRuleGroup detail param
type AddVmToVmSchedulingRuleGroupParamDetail struct {
}

// AddVmToVmSchedulingRuleGroupParam AddVmToVmSchedulingRuleGroup request param
type AddVmToVmSchedulingRuleGroupParam struct {
	BaseParam
	Params AddVmToVmSchedulingRuleGroupParamDetail `json:"addVmToVmSchedulingRuleGroup"`
}
// SyncBackupFromImageStoreBackupStorageParamDetail SyncBackupFromImageStoreBackupStorage detail param
type SyncBackupFromImageStoreBackupStorageParamDetail struct {
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// SyncBackupFromImageStoreBackupStorageParam SyncBackupFromImageStoreBackupStorage request param
type SyncBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncBackupFromImageStoreBackupStorageParamDetail `json:"syncBackupFromImageStoreBackupStorage"`
}
// GetHostWebSshUrlParamDetail GetHostWebSshUrl detail param
type GetHostWebSshUrlParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Https *bool `json:"https,omitempty"`
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// GetHostWebSshUrlParam GetHostWebSshUrl request param
type GetHostWebSshUrlParam struct {
	BaseParam
	Params GetHostWebSshUrlParamDetail `json:"params"`
}
// SetL3NetworkMtuParamDetail SetL3NetworkMtu detail param
type SetL3NetworkMtuParamDetail struct {
	Mtu int `json:"mtu" validate:"required"`
}

// SetL3NetworkMtuParam SetL3NetworkMtu request param
type SetL3NetworkMtuParam struct {
	BaseParam
	Params SetL3NetworkMtuParamDetail `json:"params"`
}
// GetL3NetworkRouterInterfaceIpParamDetail GetL3NetworkRouterInterfaceIp detail param
type GetL3NetworkRouterInterfaceIpParamDetail struct {
}

// GetL3NetworkRouterInterfaceIpParam GetL3NetworkRouterInterfaceIp request param
type GetL3NetworkRouterInterfaceIpParam struct {
	BaseParam
	Params GetL3NetworkRouterInterfaceIpParamDetail `json:"getL3NetworkRouterInterfaceIp"`
}
// SyncVmClockParamDetail SyncVmClock detail param
type SyncVmClockParamDetail struct {
}

// SyncVmClockParam SyncVmClock request param
type SyncVmClockParam struct {
	BaseParam
	Params SyncVmClockParamDetail `json:"syncVmClock"`
}
// CreateSNSSnmpEndpointParamDetail CreateSNSSnmpEndpoint detail param
type CreateSNSSnmpEndpointParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PlatformUuid *string `json:"platformUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSSnmpEndpointParam CreateSNSSnmpEndpoint request param
type CreateSNSSnmpEndpointParam struct {
	BaseParam
	Params CreateSNSSnmpEndpointParamDetail `json:"params"`
}
// SdnControllerAddHostParamDetail SdnControllerAddHost detail param
type SdnControllerAddHostParamDetail struct {
	VSwitchType *string `json:"vSwitchType,omitempty"`
	NicNames []string `json:"nicNames" validate:"required"`
	VtepIp *string `json:"vtepIp,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	BondMode *string `json:"bondMode,omitempty"`
	LacpMode *string `json:"lacpMode,omitempty"`
}

// SdnControllerAddHostParam SdnControllerAddHost request param
type SdnControllerAddHostParam struct {
	BaseParam
	Params SdnControllerAddHostParamDetail `json:"sdnControllerAddHost"`
}
// GetLicenseNodeUsageDetailsParamDetail GetLicenseNodeUsageDetails detail param
type GetLicenseNodeUsageDetailsParamDetail struct {
	NodeUuid *string `json:"nodeUuid,omitempty"`
}

// GetLicenseNodeUsageDetailsParam GetLicenseNodeUsageDetails request param
type GetLicenseNodeUsageDetailsParam struct {
	BaseParam
	Params GetLicenseNodeUsageDetailsParamDetail `json:"getLicenseNodeUsageDetails"`
}
// CreateAliyunSnapshotRemoteParamDetail CreateAliyunSnapshotRemote detail param
type CreateAliyunSnapshotRemoteParamDetail struct {
	DiskUuid string `json:"diskUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunSnapshotRemoteParam CreateAliyunSnapshotRemote request param
type CreateAliyunSnapshotRemoteParam struct {
	BaseParam
	Params CreateAliyunSnapshotRemoteParamDetail `json:"params"`
}
// SetVmBootVolumeParamDetail SetVmBootVolume detail param
type SetVmBootVolumeParamDetail struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
}

// SetVmBootVolumeParam SetVmBootVolume request param
type SetVmBootVolumeParam struct {
	BaseParam
	Params SetVmBootVolumeParamDetail `json:"setVmBootVolume"`
}
// ChangeVpcHaGroupMonitorIpsParamDetail ChangeVpcHaGroupMonitorIps detail param
type ChangeVpcHaGroupMonitorIpsParamDetail struct {
	MonitorIps []string `json:"monitorIps,omitempty"`
}

// ChangeVpcHaGroupMonitorIpsParam ChangeVpcHaGroupMonitorIps request param
type ChangeVpcHaGroupMonitorIpsParam struct {
	BaseParam
	Params ChangeVpcHaGroupMonitorIpsParamDetail `json:"changeVpcHaGroupMonitorIps"`
}
// RenewSessionParamDetail RenewSession detail param
type RenewSessionParamDetail struct {
	Duration *int64 `json:"duration,omitempty"`
}

// RenewSessionParam RenewSession request param
type RenewSessionParam struct {
	BaseParam
	Params RenewSessionParamDetail `json:"renewSession"`
}
// DeleteDataCenterInLocalParamDetail DeleteDataCenterInLocal detail param
type DeleteDataCenterInLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteDataCenterInLocalParam DeleteDataCenterInLocal request param
type DeleteDataCenterInLocalParam struct {
	BaseParam
	Params DeleteDataCenterInLocalParamDetail `json:"deleteDataCenterInLocal"`
}
// SetVmConsoleModeParamDetail SetVmConsoleMode detail param
type SetVmConsoleModeParamDetail struct {
	Mode string `json:"mode" validate:"required"`
}

// SetVmConsoleModeParam SetVmConsoleMode request param
type SetVmConsoleModeParam struct {
	BaseParam
	Params SetVmConsoleModeParamDetail `json:"setVmConsoleMode"`
}
// AttachPolicyToUserParamDetail AttachPolicyToUser detail param
type AttachPolicyToUserParamDetail struct {
	PolicyUuid string `json:"policyUuid" validate:"required"`
}

// AttachPolicyToUserParam AttachPolicyToUser request param
type AttachPolicyToUserParam struct {
	BaseParam
	Params AttachPolicyToUserParamDetail `json:"params"`
}
// GetVfPciDeviceAvailableInL2NetworkParamDetail GetVfPciDeviceAvailableInL2Network detail param
type GetVfPciDeviceAvailableInL2NetworkParamDetail struct {
	L2NetworkUuids []string `json:"l2NetworkUuids" validate:"required"`
}

// GetVfPciDeviceAvailableInL2NetworkParam GetVfPciDeviceAvailableInL2Network request param
type GetVfPciDeviceAvailableInL2NetworkParam struct {
	BaseParam
	Params GetVfPciDeviceAvailableInL2NetworkParamDetail `json:"getVfPciDeviceAvailableInL2Network"`
}
// AddAttributesToIAM2ProjectParamDetail AddAttributesToIAM2Project detail param
type AddAttributesToIAM2ProjectParamDetail struct {
	Attributes []AttributeParam `json:"attributes" validate:"required"`
}

// AddAttributesToIAM2ProjectParam AddAttributesToIAM2Project request param
type AddAttributesToIAM2ProjectParam struct {
	BaseParam
	Params AddAttributesToIAM2ProjectParamDetail `json:"params"`
}
// UngenerateSeMdevDevicesParamDetail UngenerateSeMdevDevices detail param
type UngenerateSeMdevDevicesParamDetail struct {
}

// UngenerateSeMdevDevicesParam UngenerateSeMdevDevices request param
type UngenerateSeMdevDevicesParam struct {
	BaseParam
	Params UngenerateSeMdevDevicesParamDetail `json:"ungenerateSeMdevDevices"`
}
// SetVmEmulatorPinningParamDetail SetVmEmulatorPinning detail param
type SetVmEmulatorPinningParamDetail struct {
	EmulatorPinning string `json:"emulatorPinning" validate:"required"`
}

// SetVmEmulatorPinningParam SetVmEmulatorPinning request param
type SetVmEmulatorPinningParam struct {
	BaseParam
	Params SetVmEmulatorPinningParamDetail `json:"setVmEmulatorPinning"`
}
// CleanV2VConversionCacheParamDetail CleanV2VConversionCache detail param
type CleanV2VConversionCacheParamDetail struct {
}

// CleanV2VConversionCacheParam CleanV2VConversionCache request param
type CleanV2VConversionCacheParam struct {
	BaseParam
	Params CleanV2VConversionCacheParamDetail `json:"cleanV2VConversionCache"`
}
// UnbindModelFromServiceParamDetail UnbindModelFromService detail param
type UnbindModelFromServiceParamDetail struct {
}

// UnbindModelFromServiceParam UnbindModelFromService request param
type UnbindModelFromServiceParam struct {
	BaseParam
	Params UnbindModelFromServiceParamDetail `json:"unbindModelFromService"`
}
// GetEcsInstanceTypeParamDetail GetEcsInstanceType detail param
type GetEcsInstanceTypeParamDetail struct {
	IdentityZoneUuid string `json:"identityZoneUuid" validate:"required"`
	EcsImageUuid *string `json:"ecsImageUuid,omitempty"`
}

// GetEcsInstanceTypeParam GetEcsInstanceType request param
type GetEcsInstanceTypeParam struct {
	BaseParam
	Params GetEcsInstanceTypeParamDetail `json:"getEcsInstanceType"`
}
// GetLicenseUKeyStatusParamDetail GetLicenseUKeyStatus detail param
type GetLicenseUKeyStatusParamDetail struct {
}

// GetLicenseUKeyStatusParam GetLicenseUKeyStatus request param
type GetLicenseUKeyStatusParam struct {
	BaseParam
	Params GetLicenseUKeyStatusParamDetail `json:"params"`
}
// AddTicketTypesToTicketFlowCollectionParamDetail AddTicketTypesToTicketFlowCollection detail param
type AddTicketTypesToTicketFlowCollectionParamDetail struct {
	TicketTypeUuids []string `json:"ticketTypeUuids" validate:"required"`
}

// AddTicketTypesToTicketFlowCollectionParam AddTicketTypesToTicketFlowCollection request param
type AddTicketTypesToTicketFlowCollectionParam struct {
	BaseParam
	Params AddTicketTypesToTicketFlowCollectionParamDetail `json:"params"`
}
// SetL3NetworkRouterInterfaceIpParamDetail SetL3NetworkRouterInterfaceIp detail param
type SetL3NetworkRouterInterfaceIpParamDetail struct {
	RouterInterfaceIp string `json:"routerInterfaceIp" validate:"required"`
}

// SetL3NetworkRouterInterfaceIpParam SetL3NetworkRouterInterfaceIp request param
type SetL3NetworkRouterInterfaceIpParam struct {
	BaseParam
	Params SetL3NetworkRouterInterfaceIpParamDetail `json:"params"`
}
// GetConnectionBetweenL3NetworkAndAliyunVSwitchParamDetail GetConnectionBetweenL3NetworkAndAliyunVSwitch detail param
type GetConnectionBetweenL3NetworkAndAliyunVSwitchParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceType string `json:"resourceType" validate:"required"`
}

// GetConnectionBetweenL3NetworkAndAliyunVSwitchParam GetConnectionBetweenL3NetworkAndAliyunVSwitch request param
type GetConnectionBetweenL3NetworkAndAliyunVSwitchParam struct {
	BaseParam
	Params GetConnectionBetweenL3NetworkAndAliyunVSwitchParamDetail `json:"params"`
}
// GetBareMetal2GatewayAllocatorStrategiesParamDetail GetBareMetal2GatewayAllocatorStrategies detail param
type GetBareMetal2GatewayAllocatorStrategiesParamDetail struct {
}

// GetBareMetal2GatewayAllocatorStrategiesParam GetBareMetal2GatewayAllocatorStrategies request param
type GetBareMetal2GatewayAllocatorStrategiesParam struct {
	BaseParam
	Params GetBareMetal2GatewayAllocatorStrategiesParamDetail `json:"getBareMetal2GatewayAllocatorStrategies"`
}
// UpdateFirewallRuleTemplateParamDetail UpdateFirewallRuleTemplate detail param
type UpdateFirewallRuleTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	Action string `json:"action" validate:"required"`
	Protocol *string `json:"protocol,omitempty"`
	DestPort *string `json:"destPort,omitempty"`
	SourcePort *string `json:"sourcePort,omitempty"`
	SourceIp *string `json:"sourceIp,omitempty"`
	DestIp *string `json:"destIp,omitempty"`
	AllowStates *string `json:"allowStates,omitempty"`
	TcpFlag *string `json:"tcpFlag,omitempty"`
	IcmpTypeName *string `json:"icmpTypeName,omitempty"`
	RuleNumber int `json:"ruleNumber" validate:"required"`
	EnableLog *bool `json:"enableLog,omitempty"`
	State *string `json:"state,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateFirewallRuleTemplateParam UpdateFirewallRuleTemplate request param
type UpdateFirewallRuleTemplateParam struct {
	BaseParam
	Params UpdateFirewallRuleTemplateParamDetail `json:"updateFirewallRuleTemplate"`
}
// GetUsbDeviceCandidatesForAttachingVmParamDetail GetUsbDeviceCandidatesForAttachingVm detail param
type GetUsbDeviceCandidatesForAttachingVmParamDetail struct {
	AttachType *string `json:"attachType,omitempty"`
}

// GetUsbDeviceCandidatesForAttachingVmParam GetUsbDeviceCandidatesForAttachingVm request param
type GetUsbDeviceCandidatesForAttachingVmParam struct {
	BaseParam
	Params GetUsbDeviceCandidatesForAttachingVmParamDetail `json:"getUsbDeviceCandidatesForAttachingVm"`
}
// GetCandidateL3NetworksForLoadBalancerParamDetail GetCandidateL3NetworksForLoadBalancer detail param
type GetCandidateL3NetworksForLoadBalancerParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetCandidateL3NetworksForLoadBalancerParam GetCandidateL3NetworksForLoadBalancer request param
type GetCandidateL3NetworksForLoadBalancerParam struct {
	BaseParam
	Params GetCandidateL3NetworksForLoadBalancerParamDetail `json:"getCandidateL3NetworksForLoadBalancer"`
}
// WithdrawLicenseCapacityApplicationParamDetail WithdrawLicenseCapacityApplication detail param
type WithdrawLicenseCapacityApplicationParamDetail struct {
	ResourceUuidList []string `json:"resourceUuidList" validate:"required"`
	ClientAuthorizedNodeUuid string `json:"clientAuthorizedNodeUuid" validate:"required"`
	LicenseType string `json:"licenseType" validate:"required"`
}

// WithdrawLicenseCapacityApplicationParam WithdrawLicenseCapacityApplication request param
type WithdrawLicenseCapacityApplicationParam struct {
	BaseParam
	Params WithdrawLicenseCapacityApplicationParamDetail `json:"withdrawLicenseCapacityApplication"`
}
// PowerResetHostParamDetail PowerResetHost detail param
type PowerResetHostParamDetail struct {
	ReturnEarly *bool `json:"returnEarly,omitempty"`
	Method *string `json:"method,omitempty"`
}

// PowerResetHostParam PowerResetHost request param
type PowerResetHostParam struct {
	BaseParam
	Params PowerResetHostParamDetail `json:"powerResetHost"`
}
// RevertVmFromVmBackupParamDetail RevertVmFromVmBackup detail param
type RevertVmFromVmBackupParamDetail struct {
	BackupStorageUuid *string `json:"backupStorageUuid,omitempty"`
}

// RevertVmFromVmBackupParam RevertVmFromVmBackup request param
type RevertVmFromVmBackupParam struct {
	BaseParam
	Params RevertVmFromVmBackupParamDetail `json:"revertVmFromVmBackup"`
}
// AttachCCSCertificateToUserParamDetail AttachCCSCertificateToUser detail param
type AttachCCSCertificateToUserParamDetail struct {
	CertificateUuid *string `json:"certificateUuid,omitempty"`
	State *string `json:"state,omitempty"`
}

// AttachCCSCertificateToUserParam AttachCCSCertificateToUser request param
type AttachCCSCertificateToUserParam struct {
	BaseParam
	Params AttachCCSCertificateToUserParamDetail `json:"params"`
}
// SetVmNumaParamDetail SetVmNuma detail param
type SetVmNumaParamDetail struct {
	Enable bool `json:"enable" validate:"required"`
}

// SetVmNumaParam SetVmNuma request param
type SetVmNumaParam struct {
	BaseParam
	Params SetVmNumaParamDetail `json:"setVmNuma"`
}
// DeleteAliyunRouterInterfaceLocalParamDetail DeleteAliyunRouterInterfaceLocal detail param
type DeleteAliyunRouterInterfaceLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAliyunRouterInterfaceLocalParam DeleteAliyunRouterInterfaceLocal request param
type DeleteAliyunRouterInterfaceLocalParam struct {
	BaseParam
	Params DeleteAliyunRouterInterfaceLocalParamDetail `json:"deleteAliyunRouterInterfaceLocal"`
}
// UpdateFirewallRuleSetParamDetail UpdateFirewallRuleSet detail param
type UpdateFirewallRuleSetParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ActionType *string `json:"actionType,omitempty"`
}

// UpdateFirewallRuleSetParam UpdateFirewallRuleSet request param
type UpdateFirewallRuleSetParam struct {
	BaseParam
	Params UpdateFirewallRuleSetParamDetail `json:"updateFirewallRuleSet"`
}
// AttachAliyunKeyParamDetail AttachAliyunKey detail param
type AttachAliyunKeyParamDetail struct {
}

// AttachAliyunKeyParam AttachAliyunKey request param
type AttachAliyunKeyParam struct {
	BaseParam
	Params AttachAliyunKeyParamDetail `json:"attachAliyunKey"`
}
// RefreshSearchIndexesParamDetail RefreshSearchIndexes detail param
type RefreshSearchIndexesParamDetail struct {
}

// RefreshSearchIndexesParam RefreshSearchIndexes request param
type RefreshSearchIndexesParam struct {
	BaseParam
	Params RefreshSearchIndexesParamDetail `json:"refreshSearchIndexes"`
}
// CalculateImageHashParamDetail CalculateImageHash detail param
type CalculateImageHashParamDetail struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	Algorithm *string `json:"algorithm,omitempty"`
}

// CalculateImageHashParam CalculateImageHash request param
type CalculateImageHashParam struct {
	BaseParam
	Params CalculateImageHashParamDetail `json:"calculateImageHash"`
}
// GetL2NetworkTypesParamDetail GetL2NetworkTypes detail param
type GetL2NetworkTypesParamDetail struct {
}

// GetL2NetworkTypesParam GetL2NetworkTypes request param
type GetL2NetworkTypesParam struct {
	BaseParam
	Params GetL2NetworkTypesParamDetail `json:"getL2NetworkTypes"`
}
// ShutdownHostParamDetail ShutdownHost detail param
type ShutdownHostParamDetail struct {
	ReturnEarly *bool `json:"returnEarly,omitempty"`
	Force *bool `json:"force,omitempty"`
	Method *string `json:"method,omitempty"`
}

// ShutdownHostParam ShutdownHost request param
type ShutdownHostParam struct {
	BaseParam
	Params ShutdownHostParamDetail `json:"shutdownHost"`
}
// UpdateVpcVpnConnectionRemoteParamDetail UpdateVpcVpnConnectionRemote detail param
type UpdateVpcVpnConnectionRemoteParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	LocalCidr *string `json:"localCidr,omitempty"`
	RemoteCidr *string `json:"remoteCidr,omitempty"`
	Active *bool `json:"active,omitempty"`
	IkeConfUuid *string `json:"ikeConfUuid,omitempty"`
	IpsecConfUuid *string `json:"ipsecConfUuid,omitempty"`
}

// UpdateVpcVpnConnectionRemoteParam UpdateVpcVpnConnectionRemote request param
type UpdateVpcVpnConnectionRemoteParam struct {
	BaseParam
	Params UpdateVpcVpnConnectionRemoteParamDetail `json:"updateVpcVpnConnectionRemote"`
}
// GetVmTaskParamDetail GetVmTask detail param
type GetVmTaskParamDetail struct {
	VmInstanceUuids []string `json:"vmInstanceUuids" validate:"required"`
	SyncSignatures []string `json:"syncSignatures,omitempty"`
}

// GetVmTaskParam GetVmTask request param
type GetVmTaskParam struct {
	BaseParam
	Params GetVmTaskParamDetail `json:"getVmTask"`
}
// DisableCdpTaskParamDetail DisableCdpTask detail param
type DisableCdpTaskParamDetail struct {
	Force *bool `json:"force,omitempty"`
}

// DisableCdpTaskParam DisableCdpTask request param
type DisableCdpTaskParam struct {
	BaseParam
	Params DisableCdpTaskParamDetail `json:"params"`
}
// SetIpOnHostNetworkBondingParamDetail SetIpOnHostNetworkBonding detail param
type SetIpOnHostNetworkBondingParamDetail struct {
	IpAddress *string `json:"ipAddress,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
}

// SetIpOnHostNetworkBondingParam SetIpOnHostNetworkBonding request param
type SetIpOnHostNetworkBondingParam struct {
	BaseParam
	Params SetIpOnHostNetworkBondingParamDetail `json:"params"`
}
// RemoveAttributesFromIAM2VirtualIDParamDetail RemoveAttributesFromIAM2VirtualID detail param
type RemoveAttributesFromIAM2VirtualIDParamDetail struct {
	AttributeUuids []string `json:"attributeUuids" validate:"required"`
}

// RemoveAttributesFromIAM2VirtualIDParam RemoveAttributesFromIAM2VirtualID request param
type RemoveAttributesFromIAM2VirtualIDParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2VirtualIDParamDetail `json:"removeAttributesFromIAM2VirtualID"`
}
// CreateBondingParamDetail CreateBonding detail param
type CreateBondingParamDetail struct {
	HostUuids []string `json:"hostUuids" validate:"required"`
	BondingName string `json:"bondingName" validate:"required"`
	SlaveUuids []string `json:"slaveUuids,omitempty"`
	SlaveNames []string `json:"slaveNames,omitempty"`
	Type string `json:"type" validate:"required"`
	Mode string `json:"mode" validate:"required"`
	XmitHashPolicy *string `json:"xmitHashPolicy,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBondingParam CreateBonding request param
type CreateBondingParam struct {
	BaseParam
	Params CreateBondingParamDetail `json:"params"`
}
// DetachUsbDeviceFromVmParamDetail DetachUsbDeviceFromVm detail param
type DetachUsbDeviceFromVmParamDetail struct {
}

// DetachUsbDeviceFromVmParam DetachUsbDeviceFromVm request param
type DetachUsbDeviceFromVmParam struct {
	BaseParam
	Params DetachUsbDeviceFromVmParamDetail `json:"params"`
}
// CreateDataVolumeTemplateFromVolumeSnapshotParamDetail CreateDataVolumeTemplateFromVolumeSnapshot detail param
type CreateDataVolumeTemplateFromVolumeSnapshotParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeTemplateFromVolumeSnapshotParam CreateDataVolumeTemplateFromVolumeSnapshot request param
type CreateDataVolumeTemplateFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateDataVolumeTemplateFromVolumeSnapshotParamDetail `json:"params"`
}
// DetachRoleFromAccountParamDetail DetachRoleFromAccount detail param
type DetachRoleFromAccountParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DetachRoleFromAccountParam DetachRoleFromAccount request param
type DetachRoleFromAccountParam struct {
	BaseParam
	Params DetachRoleFromAccountParamDetail `json:"detachRoleFromAccount"`
}
// AddRendezvousPointToMulticastRouterParamDetail AddRendezvousPointToMulticastRouter detail param
type AddRendezvousPointToMulticastRouterParamDetail struct {
	RpAddress string `json:"rpAddress" validate:"required"`
	GroupAddress string `json:"groupAddress" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddRendezvousPointToMulticastRouterParam AddRendezvousPointToMulticastRouter request param
type AddRendezvousPointToMulticastRouterParam struct {
	BaseParam
	Params AddRendezvousPointToMulticastRouterParamDetail `json:"params"`
}
// DeleteLdapBindingParamDetail DeleteLdapBinding detail param
type DeleteLdapBindingParamDetail struct {
}

// DeleteLdapBindingParam DeleteLdapBinding request param
type DeleteLdapBindingParam struct {
	BaseParam
	Params DeleteLdapBindingParamDetail `json:"deleteLdapBinding"`
}
// AttachNfvInstToGroupParamDetail AttachNfvInstToGroup detail param
type AttachNfvInstToGroupParamDetail struct {
}

// AttachNfvInstToGroupParam AttachNfvInstToGroup request param
type AttachNfvInstToGroupParam struct {
	BaseParam
	Params AttachNfvInstToGroupParamDetail `json:"attachNfvInstToGroup"`
}
// DebugSignalParamDetail DebugSignal detail param
type DebugSignalParamDetail struct {
	Signals []string `json:"signals" validate:"required"`
}

// DebugSignalParam DebugSignal request param
type DebugSignalParam struct {
	BaseParam
	Params DebugSignalParamDetail `json:"params"`
}
// CreateVmInstanceFromVolumeParamDetail CreateVmInstanceFromVolume detail param
type CreateVmInstanceFromVolumeParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	InstanceOfferingUuid *string `json:"instanceOfferingUuid,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	ReservedMemorySize *int64 `json:"reservedMemorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	VmNicParams *string `json:"vmNicParams,omitempty"`
	Type *string `json:"type,omitempty"`
	VolumeUuid string `json:"volumeUuid" validate:"required"`
	Platform *string `json:"platform,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceFromVolumeParam CreateVmInstanceFromVolume request param
type CreateVmInstanceFromVolumeParam struct {
	BaseParam
	Params CreateVmInstanceFromVolumeParamDetail `json:"params"`
}
// GetVpcVRouterDistributedRoutingEnabledParamDetail GetVpcVRouterDistributedRoutingEnabled detail param
type GetVpcVRouterDistributedRoutingEnabledParamDetail struct {
}

// GetVpcVRouterDistributedRoutingEnabledParam GetVpcVRouterDistributedRoutingEnabled request param
type GetVpcVRouterDistributedRoutingEnabledParam struct {
	BaseParam
	Params GetVpcVRouterDistributedRoutingEnabledParamDetail `json:"getVpcVRouterDistributedRoutingEnabled"`
}
// CreateEcsSecurityGroupRemoteParamDetail CreateEcsSecurityGroupRemote detail param
type CreateEcsSecurityGroupRemoteParamDetail struct {
	VpcUuid string `json:"vpcUuid" validate:"required"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name" validate:"required"`
	Strategy *string `json:"strategy,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsSecurityGroupRemoteParam CreateEcsSecurityGroupRemote request param
type CreateEcsSecurityGroupRemoteParam struct {
	BaseParam
	Params CreateEcsSecurityGroupRemoteParamDetail `json:"params"`
}
// RemoveAttributesFromIAM2OrganizationParamDetail RemoveAttributesFromIAM2Organization detail param
type RemoveAttributesFromIAM2OrganizationParamDetail struct {
	AttributeUuids []string `json:"attributeUuids" validate:"required"`
}

// RemoveAttributesFromIAM2OrganizationParam RemoveAttributesFromIAM2Organization request param
type RemoveAttributesFromIAM2OrganizationParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2OrganizationParamDetail `json:"removeAttributesFromIAM2Organization"`
}
// DeleteAliyunSnapshotFromLocalParamDetail DeleteAliyunSnapshotFromLocal detail param
type DeleteAliyunSnapshotFromLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAliyunSnapshotFromLocalParam DeleteAliyunSnapshotFromLocal request param
type DeleteAliyunSnapshotFromLocalParam struct {
	BaseParam
	Params DeleteAliyunSnapshotFromLocalParamDetail `json:"deleteAliyunSnapshotFromLocal"`
}
// GetIAM2ProjectContainerImagesParamDetail GetIAM2ProjectContainerImages detail param
type GetIAM2ProjectContainerImagesParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetIAM2ProjectContainerImagesParam GetIAM2ProjectContainerImages request param
type GetIAM2ProjectContainerImagesParam struct {
	BaseParam
	Params GetIAM2ProjectContainerImagesParamDetail `json:"getIAM2ProjectContainerImages"`
}
// DetachDataVolumeFromVmParamDetail DetachDataVolumeFromVm detail param
type DetachDataVolumeFromVmParamDetail struct {
	VmUuid *string `json:"vmUuid,omitempty"`
}

// DetachDataVolumeFromVmParam DetachDataVolumeFromVm request param
type DetachDataVolumeFromVmParam struct {
	BaseParam
	Params DetachDataVolumeFromVmParamDetail `json:"detachDataVolumeFromVm"`
}
// CreateRootVolumeTemplateFromRootVolumeParamDetail CreateRootVolumeTemplateFromRootVolume detail param
type CreateRootVolumeTemplateFromRootVolumeParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	Platform *string `json:"platform,omitempty"`
	System bool `json:"system,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateRootVolumeTemplateFromRootVolumeParam CreateRootVolumeTemplateFromRootVolume request param
type CreateRootVolumeTemplateFromRootVolumeParam struct {
	BaseParam
	Params CreateRootVolumeTemplateFromRootVolumeParamDetail `json:"params"`
}
// AttachAliyunDiskToEcsParamDetail AttachAliyunDiskToEcs detail param
type AttachAliyunDiskToEcsParamDetail struct {
	EcsUuid string `json:"ecsUuid" validate:"required"`
	DeleteWithInstance *bool `json:"deleteWithInstance,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AttachAliyunDiskToEcsParam AttachAliyunDiskToEcs request param
type AttachAliyunDiskToEcsParam struct {
	BaseParam
	Params AttachAliyunDiskToEcsParamDetail `json:"params"`
}
// DeleteOssBucketNameLocalParamDetail DeleteOssBucketNameLocal detail param
type DeleteOssBucketNameLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteOssBucketNameLocalParam DeleteOssBucketNameLocal request param
type DeleteOssBucketNameLocalParam struct {
	BaseParam
	Params DeleteOssBucketNameLocalParamDetail `json:"deleteOssBucketNameLocal"`
}
// GetObservabilityServerServiceDataParamDetail GetObservabilityServerServiceData detail param
type GetObservabilityServerServiceDataParamDetail struct {
	ServiceType string `json:"serviceType" validate:"required"`
	ServiceUuid string `json:"serviceUuid" validate:"required"`
	Limit *int `json:"limit,omitempty"`
	StartTime *string `json:"startTime,omitempty"`
	EndTime *string `json:"endTime,omitempty"`
	SortDirection *string `json:"sortDirection,omitempty"`
	LabelFilters map[string]string `json:"labelFilters,omitempty"`
}

// GetObservabilityServerServiceDataParam GetObservabilityServerServiceData request param
type GetObservabilityServerServiceDataParam struct {
	BaseParam
	Params GetObservabilityServerServiceDataParamDetail `json:"params"`
}
// VerifyLicenseServerParamDetail VerifyLicenseServer detail param
type VerifyLicenseServerParamDetail struct {
	AppId string `json:"appId" validate:"required"`
	ClientAccessKeyId string `json:"clientAccessKeyId" validate:"required"`
	ClientAccessKeySecret string `json:"clientAccessKeySecret" validate:"required"`
}

// VerifyLicenseServerParam VerifyLicenseServer request param
type VerifyLicenseServerParam struct {
	BaseParam
	Params VerifyLicenseServerParamDetail `json:"params"`
}
// AttachBareMetal2GatewayToClusterParamDetail AttachBareMetal2GatewayToCluster detail param
type AttachBareMetal2GatewayToClusterParamDetail struct {
}

// AttachBareMetal2GatewayToClusterParam AttachBareMetal2GatewayToCluster request param
type AttachBareMetal2GatewayToClusterParam struct {
	BaseParam
	Params AttachBareMetal2GatewayToClusterParamDetail `json:"params"`
}
// UpdateAtPersonOfAtWeComEndpointParamDetail UpdateAtPersonOfAtWeComEndpoint detail param
type UpdateAtPersonOfAtWeComEndpointParamDetail struct {
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	UserId *string `json:"userId,omitempty"`
	Remark *string `json:"remark,omitempty"`
}

// UpdateAtPersonOfAtWeComEndpointParam UpdateAtPersonOfAtWeComEndpoint request param
type UpdateAtPersonOfAtWeComEndpointParam struct {
	BaseParam
	Params UpdateAtPersonOfAtWeComEndpointParamDetail `json:"updateAtPersonOfAtWeComEndpoint"`
}
// ChangeSlbGroupDeployTypeParamDetail ChangeSlbGroupDeployType detail param
type ChangeSlbGroupDeployTypeParamDetail struct {
	DeployType string `json:"deployType" validate:"required"`
}

// ChangeSlbGroupDeployTypeParam ChangeSlbGroupDeployType request param
type ChangeSlbGroupDeployTypeParam struct {
	BaseParam
	Params ChangeSlbGroupDeployTypeParamDetail `json:"changeSlbGroupDeployType"`
}
// DeleteEcsSecurityGroupInLocalParamDetail DeleteEcsSecurityGroupInLocal detail param
type DeleteEcsSecurityGroupInLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsSecurityGroupInLocalParam DeleteEcsSecurityGroupInLocal request param
type DeleteEcsSecurityGroupInLocalParam struct {
	BaseParam
	Params DeleteEcsSecurityGroupInLocalParamDetail `json:"deleteEcsSecurityGroupInLocal"`
}
// DetachDataVolumeFromHostParamDetail DetachDataVolumeFromHost detail param
type DetachDataVolumeFromHostParamDetail struct {
	HostUuid *string `json:"hostUuid,omitempty"`
}

// DetachDataVolumeFromHostParam DetachDataVolumeFromHost request param
type DetachDataVolumeFromHostParam struct {
	BaseParam
	Params DetachDataVolumeFromHostParamDetail `json:"detachDataVolumeFromHost"`
}
// GetVmInstanceRecoveryPointsParamDetail GetVmInstanceRecoveryPoints detail param
type GetVmInstanceRecoveryPointsParamDetail struct {
	StartTime *string `json:"startTime,omitempty"`
	EndTime *string `json:"endTime,omitempty"`
	Scale *string `json:"scale,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVmInstanceRecoveryPointsParam GetVmInstanceRecoveryPoints request param
type GetVmInstanceRecoveryPointsParam struct {
	BaseParam
	Params GetVmInstanceRecoveryPointsParamDetail `json:"getVmInstanceRecoveryPoints"`
}
// CreateSystemTagsParamDetail CreateSystemTags detail param
type CreateSystemTagsParamDetail struct {
	ResourceType string `json:"resourceType" validate:"required"`
	Tags []string `json:"tags" validate:"required"`
}

// CreateSystemTagsParam CreateSystemTags request param
type CreateSystemTagsParam struct {
	BaseParam
	Params CreateSystemTagsParamDetail `json:"params"`
}
// GetOssBackupBucketFromRemoteParamDetail GetOssBackupBucketFromRemote detail param
type GetOssBackupBucketFromRemoteParamDetail struct {
}

// GetOssBackupBucketFromRemoteParam GetOssBackupBucketFromRemote request param
type GetOssBackupBucketFromRemoteParam struct {
	BaseParam
	Params GetOssBackupBucketFromRemoteParamDetail `json:"getOssBackupBucketFromRemote"`
}
// GetL3NetworkTypesParamDetail GetL3NetworkTypes detail param
type GetL3NetworkTypesParamDetail struct {
}

// GetL3NetworkTypesParam GetL3NetworkTypes request param
type GetL3NetworkTypesParam struct {
	BaseParam
	Params GetL3NetworkTypesParamDetail `json:"getL3NetworkTypes"`
}
// DetachPoliciesFromUserParamDetail DetachPoliciesFromUser detail param
type DetachPoliciesFromUserParamDetail struct {
	PolicyUuids []string `json:"policyUuids" validate:"required"`
}

// DetachPoliciesFromUserParam DetachPoliciesFromUser request param
type DetachPoliciesFromUserParam struct {
	BaseParam
	Params DetachPoliciesFromUserParamDetail `json:"detachPoliciesFromUser"`
}
// CleanUpImageCacheOnPrimaryStorageParamDetail CleanUpImageCacheOnPrimaryStorage detail param
type CleanUpImageCacheOnPrimaryStorageParamDetail struct {
	Force *bool `json:"force,omitempty"`
}

// CleanUpImageCacheOnPrimaryStorageParam CleanUpImageCacheOnPrimaryStorage request param
type CleanUpImageCacheOnPrimaryStorageParam struct {
	BaseParam
	Params CleanUpImageCacheOnPrimaryStorageParamDetail `json:"cleanUpImageCacheOnPrimaryStorage"`
}
// AddKVMHostFromConfigFileParamDetail AddKVMHostFromConfigFile detail param
type AddKVMHostFromConfigFileParamDetail struct {
	HostInfo string `json:"hostInfo" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddKVMHostFromConfigFileParam AddKVMHostFromConfigFile request param
type AddKVMHostFromConfigFileParam struct {
	BaseParam
	Params AddKVMHostFromConfigFileParamDetail `json:"params"`
}
// InspectBareMetal2ChassisByInstanceParamDetail InspectBareMetal2ChassisByInstance detail param
type InspectBareMetal2ChassisByInstanceParamDetail struct {
}

// InspectBareMetal2ChassisByInstanceParam InspectBareMetal2ChassisByInstance request param
type InspectBareMetal2ChassisByInstanceParam struct {
	BaseParam
	Params InspectBareMetal2ChassisByInstanceParamDetail `json:"inspectBareMetal2ChassisByInstance"`
}
// DeleteVmBootModeParamDetail DeleteVmBootMode detail param
type DeleteVmBootModeParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVmBootModeParam DeleteVmBootMode request param
type DeleteVmBootModeParam struct {
	BaseParam
	Params DeleteVmBootModeParamDetail `json:"deleteVmBootMode"`
}
// GetCandidateVMForAttachingAffinityGroupParamDetail GetCandidateVMForAttachingAffinityGroup detail param
type GetCandidateVMForAttachingAffinityGroupParamDetail struct {
	AffinityGroupUuid string `json:"affinityGroupUuid" validate:"required"`
}

// GetCandidateVMForAttachingAffinityGroupParam GetCandidateVMForAttachingAffinityGroup request param
type GetCandidateVMForAttachingAffinityGroupParam struct {
	BaseParam
	Params GetCandidateVMForAttachingAffinityGroupParamDetail `json:"getCandidateVMForAttachingAffinityGroup"`
}
// DeleteVpcVpnConnectionLocalParamDetail DeleteVpcVpnConnectionLocal detail param
type DeleteVpcVpnConnectionLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVpcVpnConnectionLocalParam DeleteVpcVpnConnectionLocal request param
type DeleteVpcVpnConnectionLocalParam struct {
	BaseParam
	Params DeleteVpcVpnConnectionLocalParamDetail `json:"deleteVpcVpnConnectionLocal"`
}
// DetachPolicyFromUserGroupParamDetail DetachPolicyFromUserGroup detail param
type DetachPolicyFromUserGroupParamDetail struct {
}

// DetachPolicyFromUserGroupParam DetachPolicyFromUserGroup request param
type DetachPolicyFromUserGroupParam struct {
	BaseParam
	Params DetachPolicyFromUserGroupParamDetail `json:"detachPolicyFromUserGroup"`
}
// AddActionToAlarmParamDetail AddActionToAlarm detail param
type AddActionToAlarmParamDetail struct {
	ActionUuid string `json:"actionUuid" validate:"required"`
	ActionType string `json:"actionType" validate:"required"`
}

// AddActionToAlarmParam AddActionToAlarm request param
type AddActionToAlarmParam struct {
	BaseParam
	Params AddActionToAlarmParamDetail `json:"params"`
}
// UpdateFirewallRuleParamDetail UpdateFirewallRule detail param
type UpdateFirewallRuleParamDetail struct {
	RuleSetUuid string `json:"ruleSetUuid" validate:"required"`
	Action string `json:"action" validate:"required"`
	Protocol *string `json:"protocol,omitempty"`
	DestPort *string `json:"destPort,omitempty"`
	SourcePort *string `json:"sourcePort,omitempty"`
	SourceIp *string `json:"sourceIp,omitempty"`
	DestIp *string `json:"destIp,omitempty"`
	AllowStates *string `json:"allowStates,omitempty"`
	TcpFlag *string `json:"tcpFlag,omitempty"`
	IcmpTypeName *string `json:"icmpTypeName,omitempty"`
	RuleNumber int `json:"ruleNumber" validate:"required"`
	EnableLog *bool `json:"enableLog,omitempty"`
	State string `json:"state" validate:"required"`
	Description *string `json:"description,omitempty"`
}

// UpdateFirewallRuleParam UpdateFirewallRule request param
type UpdateFirewallRuleParam struct {
	BaseParam
	Params UpdateFirewallRuleParamDetail `json:"updateFirewallRule"`
}
// ZQLQueryParamDetail ZQLQuery detail param
type ZQLQueryParamDetail struct {
	Zql *string `json:"zql,omitempty"`
}

// ZQLQueryParam ZQLQuery request param
type ZQLQueryParam struct {
	BaseParam
	Params ZQLQueryParamDetail `json:"zQLQuery"`
}
// CreateBareMetal2InstanceFromVmBackupParamDetail CreateBareMetal2InstanceFromVmBackup detail param
type CreateBareMetal2InstanceFromVmBackupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	GatewayUuid *string `json:"gatewayUuid,omitempty"`
	ChassisUuid *string `json:"chassisUuid,omitempty"`
	ChassisOfferingUuid *string `json:"chassisOfferingUuid,omitempty"`
	ChassisDiskUuid *string `json:"chassisDiskUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume *string `json:"primaryStorageUuidForDataVolume,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	GatewayAllocatorStrategy *string `json:"gatewayAllocatorStrategy,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBareMetal2InstanceFromVmBackupParam CreateBareMetal2InstanceFromVmBackup request param
type CreateBareMetal2InstanceFromVmBackupParam struct {
	BaseParam
	Params CreateBareMetal2InstanceFromVmBackupParamDetail `json:"params"`
}
// GetElaborationsParamDetail GetElaborations detail param
type GetElaborationsParamDetail struct {
	Category *string `json:"category,omitempty"`
	Code *string `json:"code,omitempty"`
	Regex *string `json:"regex,omitempty"`
}

// GetElaborationsParam GetElaborations request param
type GetElaborationsParam struct {
	BaseParam
	Params GetElaborationsParamDetail `json:"getElaborations"`
}
// GetAccessPathParamDetail GetAccessPath detail param
type GetAccessPathParamDetail struct {
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
}

// GetAccessPathParam GetAccessPath request param
type GetAccessPathParam struct {
	BaseParam
	Params GetAccessPathParamDetail `json:"getAccessPath"`
}
// GetPrimaryStorageUsageReportParamDetail GetPrimaryStorageUsageReport detail param
type GetPrimaryStorageUsageReportParamDetail struct {
	Uris []string `json:"uris,omitempty"`
}

// GetPrimaryStorageUsageReportParam GetPrimaryStorageUsageReport request param
type GetPrimaryStorageUsageReportParam struct {
	BaseParam
	Params GetPrimaryStorageUsageReportParamDetail `json:"getPrimaryStorageUsageReport"`
}
// RevertVolumeFromVolumeBackupParamDetail RevertVolumeFromVolumeBackup detail param
type RevertVolumeFromVolumeBackupParamDetail struct {
	BackupStorageUuid *string `json:"backupStorageUuid,omitempty"`
}

// RevertVolumeFromVolumeBackupParam RevertVolumeFromVolumeBackup request param
type RevertVolumeFromVolumeBackupParam struct {
	BaseParam
	Params RevertVolumeFromVolumeBackupParamDetail `json:"revertVolumeFromVolumeBackup"`
}
// CreateDataVolumeFromVolumeTemplateParamDetail CreateDataVolumeFromVolumeTemplate detail param
type CreateDataVolumeFromVolumeTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
	HostUuid *string `json:"hostUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeFromVolumeTemplateParam CreateDataVolumeFromVolumeTemplate request param
type CreateDataVolumeFromVolumeTemplateParam struct {
	BaseParam
	Params CreateDataVolumeFromVolumeTemplateParamDetail `json:"params"`
}
// LocalStorageGetVolumeMigratableHostsParamDetail LocalStorageGetVolumeMigratableHosts detail param
type LocalStorageGetVolumeMigratableHostsParamDetail struct {
}

// LocalStorageGetVolumeMigratableHostsParam LocalStorageGetVolumeMigratableHosts request param
type LocalStorageGetVolumeMigratableHostsParam struct {
	BaseParam
	Params LocalStorageGetVolumeMigratableHostsParamDetail `json:"localStorageGetVolumeMigratableHosts"`
}
// GetOssBucketNameFromRemoteParamDetail GetOssBucketNameFromRemote detail param
type GetOssBucketNameFromRemoteParamDetail struct {
	OssDomain *string `json:"ossDomain,omitempty"`
	OssKey *string `json:"ossKey,omitempty"`
	OssSecret *string `json:"ossSecret,omitempty"`
}

// GetOssBucketNameFromRemoteParam GetOssBucketNameFromRemote request param
type GetOssBucketNameFromRemoteParam struct {
	BaseParam
	Params GetOssBucketNameFromRemoteParamDetail `json:"getOssBucketNameFromRemote"`
}
// SyncEcsVpcFromRemoteParamDetail SyncEcsVpcFromRemote detail param
type SyncEcsVpcFromRemoteParamDetail struct {
	EcsVpcId *string `json:"ecsVpcId,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsVpcFromRemoteParam SyncEcsVpcFromRemote request param
type SyncEcsVpcFromRemoteParam struct {
	BaseParam
	Params SyncEcsVpcFromRemoteParamDetail `json:"params"`
}
// SetServiceTypeOnHostNetworkInterfaceParamDetail SetServiceTypeOnHostNetworkInterface detail param
type SetServiceTypeOnHostNetworkInterfaceParamDetail struct {
	InterfaceUuids []string `json:"interfaceUuids" validate:"required"`
	VlanIds []int `json:"vlanIds,omitempty"`
	ServiceTypes []string `json:"serviceTypes,omitempty"`
}

// SetServiceTypeOnHostNetworkInterfaceParam SetServiceTypeOnHostNetworkInterface request param
type SetServiceTypeOnHostNetworkInterfaceParam struct {
	BaseParam
	Params SetServiceTypeOnHostNetworkInterfaceParamDetail `json:"params"`
}
// AddBackendServerToServerGroupParamDetail AddBackendServerToServerGroup detail param
type AddBackendServerToServerGroupParamDetail struct {
	VmNics []interface{} `json:"vmNics,omitempty"`
	Servers []interface{} `json:"servers,omitempty"`
}

// AddBackendServerToServerGroupParam AddBackendServerToServerGroup request param
type AddBackendServerToServerGroupParam struct {
	BaseParam
	Params AddBackendServerToServerGroupParamDetail `json:"params"`
}
// AttachUserDefinedXmlHookScriptToVmParamDetail AttachUserDefinedXmlHookScriptToVm detail param
type AttachUserDefinedXmlHookScriptToVmParamDetail struct {
	StartupStrategy *string `json:"startupStrategy,omitempty"`
}

// AttachUserDefinedXmlHookScriptToVmParam AttachUserDefinedXmlHookScriptToVm request param
type AttachUserDefinedXmlHookScriptToVmParam struct {
	BaseParam
	Params AttachUserDefinedXmlHookScriptToVmParamDetail `json:"params"`
}
// AttachPolicyToRoleParamDetail AttachPolicyToRole detail param
type AttachPolicyToRoleParamDetail struct {
}

// AttachPolicyToRoleParam AttachPolicyToRole request param
type AttachPolicyToRoleParam struct {
	BaseParam
	Params AttachPolicyToRoleParamDetail `json:"params"`
}
// ChangeBareMetal2ProvisionNetworkStateParamDetail ChangeBareMetal2ProvisionNetworkState detail param
type ChangeBareMetal2ProvisionNetworkStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBareMetal2ProvisionNetworkStateParam ChangeBareMetal2ProvisionNetworkState request param
type ChangeBareMetal2ProvisionNetworkStateParam struct {
	BaseParam
	Params ChangeBareMetal2ProvisionNetworkStateParamDetail `json:"changeBareMetal2ProvisionNetworkState"`
}
// GetBackupStorageCandidatesForImageMigrationParamDetail GetBackupStorageCandidatesForImageMigration detail param
type GetBackupStorageCandidatesForImageMigrationParamDetail struct {
}

// GetBackupStorageCandidatesForImageMigrationParam GetBackupStorageCandidatesForImageMigration request param
type GetBackupStorageCandidatesForImageMigrationParam struct {
	BaseParam
	Params GetBackupStorageCandidatesForImageMigrationParamDetail `json:"getBackupStorageCandidatesForImageMigration"`
}
// DeleteVpcIpSecConfigLocalParamDetail DeleteVpcIpSecConfigLocal detail param
type DeleteVpcIpSecConfigLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVpcIpSecConfigLocalParam DeleteVpcIpSecConfigLocal request param
type DeleteVpcIpSecConfigLocalParam struct {
	BaseParam
	Params DeleteVpcIpSecConfigLocalParamDetail `json:"deleteVpcIpSecConfigLocal"`
}
// GenerateSriovPciDevicesParamDetail GenerateSriovPciDevices detail param
type GenerateSriovPciDevicesParamDetail struct {
	VirtPartNum int `json:"virtPartNum" validate:"required"`
}

// GenerateSriovPciDevicesParam GenerateSriovPciDevices request param
type GenerateSriovPciDevicesParam struct {
	BaseParam
	Params GenerateSriovPciDevicesParamDetail `json:"generateSriovPciDevices"`
}
// CalculateAccountBillingSpendingParamDetail CalculateAccountBillingSpending detail param
type CalculateAccountBillingSpendingParamDetail struct {
	DateStart *int64 `json:"dateStart,omitempty"`
	DateEnd *int64 `json:"dateEnd,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	Simple *bool `json:"simple,omitempty"`
}

// CalculateAccountBillingSpendingParam CalculateAccountBillingSpending request param
type CalculateAccountBillingSpendingParam struct {
	BaseParam
	Params CalculateAccountBillingSpendingParamDetail `json:"calculateAccountBillingSpending"`
}
// DeleteVRouterOspfAreaParamDetail DeleteVRouterOspfArea detail param
type DeleteVRouterOspfAreaParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVRouterOspfAreaParam DeleteVRouterOspfArea request param
type DeleteVRouterOspfAreaParam struct {
	BaseParam
	Params DeleteVRouterOspfAreaParamDetail `json:"deleteVRouterOspfArea"`
}
// GetVipAvailablePortParamDetail GetVipAvailablePort detail param
type GetVipAvailablePortParamDetail struct {
	ProtocolType string `json:"protocolType" validate:"required"`
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetVipAvailablePortParam GetVipAvailablePort request param
type GetVipAvailablePortParam struct {
	BaseParam
	Params GetVipAvailablePortParamDetail `json:"getVipAvailablePort"`
}
// SyncDiskFromAliyunFromRemoteParamDetail SyncDiskFromAliyunFromRemote detail param
type SyncDiskFromAliyunFromRemoteParamDetail struct {
	DiskId *string `json:"diskId,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncDiskFromAliyunFromRemoteParam SyncDiskFromAliyunFromRemote request param
type SyncDiskFromAliyunFromRemoteParam struct {
	BaseParam
	Params SyncDiskFromAliyunFromRemoteParamDetail `json:"params"`
}
// ChangeVolumeStateParamDetail ChangeVolumeState detail param
type ChangeVolumeStateParamDetail struct {
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeVolumeStateParam ChangeVolumeState request param
type ChangeVolumeStateParam struct {
	BaseParam
	Params ChangeVolumeStateParamDetail `json:"changeVolumeState"`
}
// MountVmInstanceRecoveryPointParamDetail MountVmInstanceRecoveryPoint detail param
type MountVmInstanceRecoveryPointParamDetail struct {
	VmUuid string `json:"vmUuid" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
	Https *bool `json:"https,omitempty"`
}

// MountVmInstanceRecoveryPointParam MountVmInstanceRecoveryPoint request param
type MountVmInstanceRecoveryPointParam struct {
	BaseParam
	Params MountVmInstanceRecoveryPointParamDetail `json:"params"`
}
// CreateVxlanPoolRemoteVtepParamDetail CreateVxlanPoolRemoteVtep detail param
type CreateVxlanPoolRemoteVtepParamDetail struct {
	RemoteVtepIp string `json:"remoteVtepIp" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVxlanPoolRemoteVtepParam CreateVxlanPoolRemoteVtep request param
type CreateVxlanPoolRemoteVtepParam struct {
	BaseParam
	Params CreateVxlanPoolRemoteVtepParamDetail `json:"params"`
}
// GetResourceStackFromResourceParamDetail GetResourceStackFromResource detail param
type GetResourceStackFromResourceParamDetail struct {
	ResourceUuid string `json:"resourceUuid" validate:"required"`
}

// GetResourceStackFromResourceParam GetResourceStackFromResource request param
type GetResourceStackFromResourceParam struct {
	BaseParam
	Params GetResourceStackFromResourceParamDetail `json:"getResourceStackFromResource"`
}
// CheckIAM2VirtualIDConfigFileParamDetail CheckIAM2VirtualIDConfigFile detail param
type CheckIAM2VirtualIDConfigFileParamDetail struct {
	VirtualIDInfos string `json:"virtualIDInfos" validate:"required"`
}

// CheckIAM2VirtualIDConfigFileParam CheckIAM2VirtualIDConfigFile request param
type CheckIAM2VirtualIDConfigFileParam struct {
	BaseParam
	Params CheckIAM2VirtualIDConfigFileParamDetail `json:"checkIAM2VirtualIDConfigFile"`
}
// GetClusterHostNetworkFactsParamDetail GetClusterHostNetworkFacts detail param
type GetClusterHostNetworkFactsParamDetail struct {
	Limit *int `json:"limit,omitempty"`
	Start *int `json:"start,omitempty"`
}

// GetClusterHostNetworkFactsParam GetClusterHostNetworkFacts request param
type GetClusterHostNetworkFactsParam struct {
	BaseParam
	Params GetClusterHostNetworkFactsParamDetail `json:"getClusterHostNetworkFacts"`
}
// DetachOssBucketFromEcsDataCenterParamDetail DetachOssBucketFromEcsDataCenter detail param
type DetachOssBucketFromEcsDataCenterParamDetail struct {
}

// DetachOssBucketFromEcsDataCenterParam DetachOssBucketFromEcsDataCenter request param
type DetachOssBucketFromEcsDataCenterParam struct {
	BaseParam
	Params DetachOssBucketFromEcsDataCenterParamDetail `json:"detachOssBucketFromEcsDataCenter"`
}
// ParseOvfParamDetail ParseOvf detail param
type ParseOvfParamDetail struct {
	XmlBase64 string `json:"xmlBase64" validate:"required"`
}

// ParseOvfParam ParseOvf request param
type ParseOvfParam struct {
	BaseParam
	Params ParseOvfParamDetail `json:"params"`
}
// AddSchedulerJobGroupToSchedulerTriggerParamDetail AddSchedulerJobGroupToSchedulerTrigger detail param
type AddSchedulerJobGroupToSchedulerTriggerParamDetail struct {
	TriggerNow *bool `json:"triggerNow,omitempty"`
}

// AddSchedulerJobGroupToSchedulerTriggerParam AddSchedulerJobGroupToSchedulerTrigger request param
type AddSchedulerJobGroupToSchedulerTriggerParam struct {
	BaseParam
	Params AddSchedulerJobGroupToSchedulerTriggerParamDetail `json:"params"`
}
// DeleteAliyunNasAccessGroupRuleParamDetail DeleteAliyunNasAccessGroupRule detail param
type DeleteAliyunNasAccessGroupRuleParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAliyunNasAccessGroupRuleParam DeleteAliyunNasAccessGroupRule request param
type DeleteAliyunNasAccessGroupRuleParam struct {
	BaseParam
	Params DeleteAliyunNasAccessGroupRuleParamDetail `json:"deleteAliyunNasAccessGroupRule"`
}
// DeleteBondingParamDetail DeleteBonding detail param
type DeleteBondingParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteBondingParam DeleteBonding request param
type DeleteBondingParam struct {
	BaseParam
	Params DeleteBondingParamDetail `json:"deleteBonding"`
}
// DeleteEcsSecurityGroupRemoteParamDetail DeleteEcsSecurityGroupRemote detail param
type DeleteEcsSecurityGroupRemoteParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsSecurityGroupRemoteParam DeleteEcsSecurityGroupRemote request param
type DeleteEcsSecurityGroupRemoteParam struct {
	BaseParam
	Params DeleteEcsSecurityGroupRemoteParamDetail `json:"deleteEcsSecurityGroupRemote"`
}
// DeleteVmNicFromSecurityGroupParamDetail DeleteVmNicFromSecurityGroup detail param
type DeleteVmNicFromSecurityGroupParamDetail struct {
	VmNicUuids []string `json:"vmNicUuids" validate:"required"`
}

// DeleteVmNicFromSecurityGroupParam DeleteVmNicFromSecurityGroup request param
type DeleteVmNicFromSecurityGroupParam struct {
	BaseParam
	Params DeleteVmNicFromSecurityGroupParamDetail `json:"deleteVmNicFromSecurityGroup"`
}
// UpdateTagParamDetail UpdateTag detail param
type UpdateTagParamDetail struct {
	Name string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
	Description *string `json:"description,omitempty"`
	Color *string `json:"color,omitempty"`
}

// UpdateTagParam UpdateTag request param
type UpdateTagParam struct {
	BaseParam
	Params UpdateTagParamDetail `json:"updateTag"`
}
// AttachVRouterRouteTableToVRouterParamDetail AttachVRouterRouteTableToVRouter detail param
type AttachVRouterRouteTableToVRouterParamDetail struct {
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid" validate:"required"`
}

// AttachVRouterRouteTableToVRouterParam AttachVRouterRouteTableToVRouter request param
type AttachVRouterRouteTableToVRouterParam struct {
	BaseParam
	Params AttachVRouterRouteTableToVRouterParamDetail `json:"params"`
}
// CreateVxlanVtepParamDetail CreateVxlanVtep detail param
type CreateVxlanVtepParamDetail struct {
	HostUuid string `json:"hostUuid" validate:"required"`
	PoolUuid string `json:"poolUuid" validate:"required"`
	VtepIp *string `json:"vtepIp,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVxlanVtepParam CreateVxlanVtep request param
type CreateVxlanVtepParam struct {
	BaseParam
	Params CreateVxlanVtepParamDetail `json:"params"`
}
// AddMdevDeviceSpecToVmInstanceParamDetail AddMdevDeviceSpecToVmInstance detail param
type AddMdevDeviceSpecToVmInstanceParamDetail struct {
	MdevDeviceNumber *int `json:"mdevDeviceNumber,omitempty"`
}

// AddMdevDeviceSpecToVmInstanceParam AddMdevDeviceSpecToVmInstance request param
type AddMdevDeviceSpecToVmInstanceParam struct {
	BaseParam
	Params AddMdevDeviceSpecToVmInstanceParamDetail `json:"params"`
}
// DetachScsiLunFromVmInstanceParamDetail DetachScsiLunFromVmInstance detail param
type DetachScsiLunFromVmInstanceParamDetail struct {
}

// DetachScsiLunFromVmInstanceParam DetachScsiLunFromVmInstance request param
type DetachScsiLunFromVmInstanceParam struct {
	BaseParam
	Params DetachScsiLunFromVmInstanceParamDetail `json:"detachScsiLunFromVmInstance"`
}
// EnableCdpTaskParamDetail EnableCdpTask detail param
type EnableCdpTaskParamDetail struct {
}

// EnableCdpTaskParam EnableCdpTask request param
type EnableCdpTaskParam struct {
	BaseParam
	Params EnableCdpTaskParamDetail `json:"params"`
}
// SyncConnectionAccessPointFromRemoteParamDetail SyncConnectionAccessPointFromRemote detail param
type SyncConnectionAccessPointFromRemoteParamDetail struct {
	AccessPointId *string `json:"accessPointId,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncConnectionAccessPointFromRemoteParam SyncConnectionAccessPointFromRemote request param
type SyncConnectionAccessPointFromRemoteParam struct {
	BaseParam
	Params SyncConnectionAccessPointFromRemoteParamDetail `json:"syncConnectionAccessPointFromRemote"`
}
// RegisterLicenseRequestedApplicationParamDetail RegisterLicenseRequestedApplication detail param
type RegisterLicenseRequestedApplicationParamDetail struct {
	LicenseRequestCode string `json:"licenseRequestCode" validate:"required"`
	ClientPubKey *string `json:"clientPubKey,omitempty"`
	CurrentTimeMillis *int64 `json:"currentTimeMillis,omitempty"`
}

// RegisterLicenseRequestedApplicationParam RegisterLicenseRequestedApplication request param
type RegisterLicenseRequestedApplicationParam struct {
	BaseParam
	Params RegisterLicenseRequestedApplicationParamDetail `json:"params"`
}
// SyncVpcVpnGatewayFromRemoteParamDetail SyncVpcVpnGatewayFromRemote detail param
type SyncVpcVpnGatewayFromRemoteParamDetail struct {
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncVpcVpnGatewayFromRemoteParam SyncVpcVpnGatewayFromRemote request param
type SyncVpcVpnGatewayFromRemoteParam struct {
	BaseParam
	Params SyncVpcVpnGatewayFromRemoteParamDetail `json:"syncVpcVpnGatewayFromRemote"`
}
// DeleteEcsVpcInLocalParamDetail DeleteEcsVpcInLocal detail param
type DeleteEcsVpcInLocalParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteEcsVpcInLocalParam DeleteEcsVpcInLocal request param
type DeleteEcsVpcInLocalParam struct {
	BaseParam
	Params DeleteEcsVpcInLocalParamDetail `json:"deleteEcsVpcInLocal"`
}
