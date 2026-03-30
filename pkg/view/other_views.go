// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AddAttributesToIAM2OrganizationEventView AddAttributesToIAM2OrganizationEvent
type AddAttributesToIAM2OrganizationEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetCreateEcsImageProgressView GetCreateEcsImageProgress
type GetCreateEcsImageProgressView struct {
	Progress ProgressPropertyView `json:"progress,omitempty"`
}

// LogOutView LogOut
type LogOutView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetVmXmlHookScriptView GetVmXmlHookScript
type GetVmXmlHookScriptView struct {
	UserDefinedXmlHookScript string `json:"userDefinedXmlHookScript,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AttachHybridKeyEventView AttachHybridKeyEvent
type AttachHybridKeyEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetImageQgaView GetImageQga
type GetImageQgaView struct {
	Enable bool `json:"enable,omitempty"`
}

// GetInterdependentL3NetworksBackupStoragesView GetInterdependentL3NetworksBackupStorages
type GetInterdependentL3NetworksBackupStoragesView struct {
	Inventories ListView `json:"inventories,omitempty"`
}

// DeleteBackupDatabaseInPublicEventView DeleteBackupDatabaseInPublicEvent
type DeleteBackupDatabaseInPublicEventView struct {
	Success bool `json:"success,omitempty"`
}

// BatchCreateIAM2VirtualIDFromConfigFileEventView BatchCreateIAM2VirtualIDFromConfigFileEvent
type BatchCreateIAM2VirtualIDFromConfigFileEventView struct {
	NumberOfImportedUser int `json:"numberOfImportedUser,omitempty"`
}

// SyncDataCenterFromRemoteEventView SyncDataCenterFromRemoteEvent
type SyncDataCenterFromRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// SecurityMachineDetectSyncEventView SecurityMachineDetectSyncEvent
type SecurityMachineDetectSyncEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetPrometheusMetricLabelValueView GetPrometheusMetricLabelValue
type GetPrometheusMetricLabelValueView struct {
	LabelValues map[string]interface{} `json:"labelValues,omitempty"`
}

// UpdateAlarmDataEventView UpdateAlarmDataEvent
type UpdateAlarmDataEventView struct {
	Success bool `json:"success,omitempty"`
}

// LoginIAM2VirtualIDWithLdapView LoginIAM2VirtualIDWithLdap
type LoginIAM2VirtualIDWithLdapView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// SNSEmailTestConnectionEventView SNSEmailTestConnectionEvent
type SNSEmailTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// RegisterLicenseServerEventView RegisterLicenseServerEvent
type RegisterLicenseServerEventView struct {
	LicenseClient LicenseAuthorizedNodeInventoryView `json:"licenseClient,omitempty"`
	LicenseServer LicenseAuthorizedNodeInventoryView `json:"licenseServer,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteModelEvaluationTasksEventView DeleteModelEvaluationTasksEvent
type DeleteModelEvaluationTasksEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetTwoFactorAuthenticationStateView GetTwoFactorAuthenticationState
type GetTwoFactorAuthenticationStateView struct {
	State string `json:"state,omitempty"`
}

// BootstrapMiniHostEventView BootstrapMiniHostEvent
type BootstrapMiniHostEventView struct {
	Stage string `json:"stage,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DetachSshKeyPairFromVmInstanceEventView DetachSshKeyPairFromVmInstanceEvent
type DetachSshKeyPairFromVmInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// PrimaryStorageMigrateVolumeEventView PrimaryStorageMigrateVolumeEvent
type PrimaryStorageMigrateVolumeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// DeleteHybridEipRemoteEventView DeleteHybridEipRemoteEvent
type DeleteHybridEipRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteModelServiceInstanceGroupsEventView DeleteModelServiceInstanceGroupsEvent
type DeleteModelServiceInstanceGroupsEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetVmBootOrderView GetVmBootOrder
type GetVmBootOrderView struct {
	Orders []string `json:"orders,omitempty"`
}

// GetDatabaseBackupFromImageStoreView GetDatabaseBackupFromImageStore
type GetDatabaseBackupFromImageStoreView struct {
	Backups []DatabaseBackupStructView `json:"backups,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CleanUpBaremetalChassisBondingEventView CleanUpBaremetalChassisBondingEvent
type CleanUpBaremetalChassisBondingEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemovePciDeviceSpecFromVmInstanceEventView RemovePciDeviceSpecFromVmInstanceEvent
type RemovePciDeviceSpecFromVmInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddIAM2VirtualIDGroupToProjectsEventView AddIAM2VirtualIDGroupToProjectsEvent
type AddIAM2VirtualIDGroupToProjectsEventView struct {
	Success bool `json:"success,omitempty"`
}

// RefreshCaptchaView RefreshCaptcha
type RefreshCaptchaView struct {
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	Captcha string `json:"captcha,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteEcsVSwitchInLocalEventView DeleteEcsVSwitchInLocalEvent
type DeleteEcsVSwitchInLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteTagEventView DeleteTagEvent
type DeleteTagEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddIAM2VirtualIDsToOrganizationEventView AddIAM2VirtualIDsToOrganizationEvent
type AddIAM2VirtualIDsToOrganizationEventView struct {
	Success bool `json:"success,omitempty"`
}

// ExportNbdVolumesEventView ExportNbdVolumesEvent
type ExportNbdVolumesEventView struct {
	VolumeInfos []VolumeCbtBackupInfoView `json:"volumeInfos,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SelfTestLocalRaidEventView SelfTestLocalRaidEvent
type SelfTestLocalRaidEventView struct {
	Result string `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}

// PowerOffBareMetal2ChassisEventView PowerOffBareMetal2ChassisEvent
type PowerOffBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// SdnControllerChangeHostEventView SdnControllerChangeHostEvent
type SdnControllerChangeHostEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// DetachTagFromResourcesEventView DetachTagFromResourcesEvent
type DetachTagFromResourcesEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteVmInstanceHaLevelEventView DeleteVmInstanceHaLevelEvent
type DeleteVmInstanceHaLevelEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteResourcePriceEventView DeleteResourcePriceEvent
type DeleteResourcePriceEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanUpBaremetal2BondingEventView CleanUpBaremetal2BondingEvent
type CleanUpBaremetal2BondingEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteMetricDataEventView DeleteMetricDataEvent
type DeleteMetricDataEventView struct {
	Success bool `json:"success,omitempty"`
}

// RevertVmFromCdpBackupEventView RevertVmFromCdpBackupEvent
type RevertVmFromCdpBackupEventView struct {
	Success bool `json:"success,omitempty"`
}

// SNSFeiShuTestConnectionEventView SNSFeiShuTestConnectionEvent
type SNSFeiShuTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// GetSchedulerExecutionReportView GetSchedulerExecutionReport
type GetSchedulerExecutionReportView struct {
	SuccessRecords []int `json:"successRecords,omitempty"`
	FailureRecords []int `json:"failureRecords,omitempty"`
	PartialSuccessRecords []int `json:"partialSuccessRecords,omitempty"`
	WaitingRecords []int `json:"waitingRecords,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetSupportedIdentityModelsView GetSupportedIdentityModels
type GetSupportedIdentityModelsView struct {
	Configs []string `json:"configs,omitempty"`
}

// AddUserToGroupEventView AddUserToGroupEvent
type AddUserToGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetPrimaryStorageTypesView GetPrimaryStorageTypes
type GetPrimaryStorageTypesView struct {
	Types []string `json:"types,omitempty"`
}

// BatchDeleteVolumeSnapshotEventView BatchDeleteVolumeSnapshotEvent
type BatchDeleteVolumeSnapshotEventView struct {
	Results []BatchDeleteVolumeSnapshotStructView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ReloadLicenseView ReloadLicense
type ReloadLicenseView struct {
	Inventory LicenseInventoryView `json:"inventory,omitempty"`
}

// DeleteNicQosEventView DeleteNicQosEvent
type DeleteNicQosEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetResourceStackVmStatusView GetResourceStackVmStatus
type GetResourceStackVmStatusView struct {
	PortStatus map[string]interface{} `json:"portStatus,omitempty"`
}

// DetachHybridKeyEventView DetachHybridKeyEvent
type DetachHybridKeyEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteHybridEipFromLocalEventView DeleteHybridEipFromLocalEvent
type DeleteHybridEipFromLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// ReimageVmInstanceEventView ReimageVmInstanceEvent
type ReimageVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// UpdateDatasetsEventView UpdateDatasetsEvent
type UpdateDatasetsEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
}

// SyncIdentityFromRemoteEventView SyncIdentityFromRemoteEvent
type SyncIdentityFromRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetImageStoreBackupStorageQuotaEventView SetImageStoreBackupStorageQuotaEvent
type SetImageStoreBackupStorageQuotaEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetIAM2OrganizationVirtualIDNumberView GetIAM2OrganizationVirtualIDNumber
type GetIAM2OrganizationVirtualIDNumberView struct {
	VirtualTotalNumber int `json:"virtualTotalNumber,omitempty"`
	VirtualDirectNumber int `json:"virtualDirectNumber,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteEcsInstanceLocalEventView DeleteEcsInstanceLocalEvent
type DeleteEcsInstanceLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnsubscribeSNSTopicEventView UnsubscribeSNSTopicEvent
type UnsubscribeSNSTopicEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetNicQosEventView SetNicQosEvent
type SetNicQosEventView struct {
	Success bool `json:"success,omitempty"`
}

// CancelLongJobEventView CancelLongJobEvent
type CancelLongJobEventView struct {
	Success bool `json:"success,omitempty"`
}

// GenerateAccountBillingEventView GenerateAccountBillingEvent
type GenerateAccountBillingEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetInvocationRecordsView GetInvocationRecords
type GetInvocationRecordsView struct {
	Inventories []InvocationRecordView `json:"inventories,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetVRouterFlowCounterView GetVRouterFlowCounter
type GetVRouterFlowCounterView struct {
	Counters []FlowCounterView `json:"counters,omitempty"`
}

// GetBareMetal2SupportedBootModeView GetBareMetal2SupportedBootMode
type GetBareMetal2SupportedBootModeView struct {
	SupportedBootMode string `json:"supportedBootMode,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetChainTaskView GetChainTask
type GetChainTaskView struct {
	Results map[string]ChainInfoView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ChangeHostPasswordEventView ChangeHostPasswordEvent
type ChangeHostPasswordEventView struct {
	Success bool `json:"success,omitempty"`
}

// IsLicenseServerView IsLicenseServer
type IsLicenseServerView struct {
	LicenseServer bool `json:"licenseServer,omitempty"`
	Success bool `json:"success,omitempty"`
}

// PrometheusQueryLabelValuesView PrometheusQueryLabelValues
type PrometheusQueryLabelValuesView struct {
	Inventories MapView `json:"inventories,omitempty"`
}

// ValidateClusterSupportDRSView ValidateClusterSupportDRS
type ValidateClusterSupportDRSView struct {
	Supported bool `json:"supported,omitempty"`
	Reason ErrorCodeView `json:"reason,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ShrinkVolumeSnapshotEventView ShrinkVolumeSnapshotEvent
type ShrinkVolumeSnapshotEventView struct {
	ShrinkResult ShrinkResultView `json:"shrinkResult,omitempty"`
}

// AddHostToHostSchedulingRuleGroupEventView AddHostToHostSchedulingRuleGroupEvent
type AddHostToHostSchedulingRuleGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmNicAttachedNetworkServiceView GetVmNicAttachedNetworkService
type GetVmNicAttachedNetworkServiceView struct {
	NetworkServices []string `json:"networkServices,omitempty"`
}

// GetVmHostnameView GetVmHostname
type GetVmHostnameView struct {
	Hostname string `json:"hostname,omitempty"`
}

// DeleteVpcUserVpnGatewayLocalEventView DeleteVpcUserVpnGatewayLocalEvent
type DeleteVpcUserVpnGatewayLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveSchedulerJobFromSchedulerTriggerEventView RemoveSchedulerJobFromSchedulerTriggerEvent
type RemoveSchedulerJobFromSchedulerTriggerEventView struct {
	Success bool `json:"success,omitempty"`
}

// StopAllResourcesInIAM2ProjectEventView StopAllResourcesInIAM2ProjectEvent
type StopAllResourcesInIAM2ProjectEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateVmNetworkConfigEventView UpdateVmNetworkConfigEvent
type UpdateVmNetworkConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveTicketTypesFromTicketFlowCollectionEventView RemoveTicketTypesFromTicketFlowCollectionEvent
type RemoveTicketTypesFromTicketFlowCollectionEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteEcsVSwitchRemoteEventView DeleteEcsVSwitchRemoteEvent
type DeleteEcsVSwitchRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmStaticIpEventView SetVmStaticIpEvent
type SetVmStaticIpEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmSshKeyView GetVmSshKey
type GetVmSshKeyView struct {
	SshKey string `json:"sshKey,omitempty"`
}

// GetVmGuestToolsInfoView GetVmGuestToolsInfo
type GetVmGuestToolsInfoView struct {
	Version string `json:"version,omitempty"`
	Status string `json:"status,omitempty"`
	Features map[string]string `json:"features,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ValidateDiskOfferingUserConfigEventView ValidateDiskOfferingUserConfigEvent
type ValidateDiskOfferingUserConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteVpcVpnGatewayLocalEventView DeleteVpcVpnGatewayLocalEvent
type DeleteVpcVpnGatewayLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmRDPEventView SetVmRDPEvent
type SetVmRDPEventView struct {
	Success bool `json:"success,omitempty"`
}

// RunSchedulerTriggerEventView RunSchedulerTriggerEvent
type RunSchedulerTriggerEventView struct {
	Success bool `json:"success,omitempty"`
}

// PowerOnHostEventView PowerOnHostEvent
type PowerOnHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteAliyunSnapshotFromRemoteEventView DeleteAliyunSnapshotFromRemoteEvent
type DeleteAliyunSnapshotFromRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddIAM2VirtualIDsToProjectEventView AddIAM2VirtualIDsToProjectEvent
type AddIAM2VirtualIDsToProjectEventView struct {
	Success bool `json:"success,omitempty"`
}

// SubscribeEventEventView SubscribeEventEvent
type SubscribeEventEventView struct {
	Inventory EventSubscriptionInventoryView `json:"inventory,omitempty"`
}

// UpgradeBackupStorageCdpTasksEventView UpgradeBackupStorageCdpTasksEvent
type UpgradeBackupStorageCdpTasksEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVolumeIoThreadPinEventView SetVolumeIoThreadPinEvent
type SetVolumeIoThreadPinEventView struct {
	VolumeUuid string `json:"volumeUuid,omitempty"`
	IoThreadId int `json:"ioThreadId,omitempty"`
	Pin string `json:"pin,omitempty"`
}

// UpdatePriorityConfigEventView UpdatePriorityConfigEvent
type UpdatePriorityConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// IdentifyHostEventView IdentifyHostEvent
type IdentifyHostEventView struct {
	Success bool `json:"success,omitempty"`
}

// CheckFirewallRuleConfigFileView CheckFirewallRuleConfigFile
type CheckFirewallRuleConfigFileView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmConsoleAddressView GetVmConsoleAddress
type GetVmConsoleAddressView struct {
	HostIp string `json:"hostIp,omitempty"`
	Port int `json:"port,omitempty"`
	Path string `json:"path,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	VdiPortInfo VdiPortInfoView `json:"vdiPortInfo,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetLoadBalancerListenerACLEntriesView GetLoadBalancerListenerACLEntries
type GetLoadBalancerListenerACLEntriesView struct {
	Inventories map[string]interface{} `json:"inventories,omitempty"`
}

// UpdateHostIommuStateEventView UpdateHostIommuStateEvent
type UpdateHostIommuStateEventView struct {
	State string `json:"state,omitempty"`
}

// UnsubscribeEventEventView UnsubscribeEventEvent
type UnsubscribeEventEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmsSchedulingStateFromSchedulingRuleView GetVmsSchedulingStateFromSchedulingRule
type GetVmsSchedulingStateFromSchedulingRuleView struct {
	RuleMapState map[string]string `json:"ruleMapState,omitempty"`
}

// GetLocalStorageHostDiskCapacityView GetLocalStorageHostDiskCapacity
type GetLocalStorageHostDiskCapacityView struct {
	Inventories []HostDiskCapacityView `json:"inventories,omitempty"`
}

// DeleteVxlanPoolRemoteVtepEventView DeleteVxlanPoolRemoteVtepEvent
type DeleteVxlanPoolRemoteVtepEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveAttributesFromIAM2ProjectEventView RemoveAttributesFromIAM2ProjectEvent
type RemoveAttributesFromIAM2ProjectEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveIAM2VirtualIDsFromGroupEventView RemoveIAM2VirtualIDsFromGroupEvent
type RemoveIAM2VirtualIDsFromGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// ProvisionSlbGroupInstanceEventView ProvisionSlbGroupInstanceEvent
type ProvisionSlbGroupInstanceEventView struct {
	Inventory SlbGroupInventoryView `json:"inventory,omitempty"`
}

// SetVmUserDefinedXmlHookScriptEventView SetVmUserDefinedXmlHookScriptEvent
type SetVmUserDefinedXmlHookScriptEventView struct {
	VmUserDefinedXmlHookScript string `json:"vmUserDefinedXmlHookScript,omitempty"`
	Success bool `json:"success,omitempty"`
}

// LoginIAM2PlatformView LoginIAM2Platform
type LoginIAM2PlatformView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// GetHostAllocatorStrategiesView GetHostAllocatorStrategies
type GetHostAllocatorStrategiesView struct {
	Strategies []string `json:"strategies,omitempty"`
}

// GetInterfaceServiceTypeStatisticView GetInterfaceServiceTypeStatistic
type GetInterfaceServiceTypeStatisticView struct {
	ServiceTypeStatistics []ServiceTypeStatisticDataView `json:"serviceTypeStatistics,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteModelsEventView DeleteModelsEvent
type DeleteModelsEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ListVmsFromSchedulingStateView ListVmsFromSchedulingState
type ListVmsFromSchedulingStateView struct {
	Uuids []string `json:"uuids,omitempty"`
}

// AllocateHostResourceEventView AllocateHostResourceEvent
type AllocateHostResourceEventView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	VCPUPin []interface{} `json:"vCPUPin,omitempty"`
}

// GetLdapEntryView GetLdapEntry
type GetLdapEntryView struct {
	Inventories ListView `json:"inventories,omitempty"`
}

// CheckElaborationContentView CheckElaborationContent
type CheckElaborationContentView struct {
	Results []ElaborationCheckResultView `json:"results,omitempty"`
}

// GetPrimaryStorageLicenseInfoView GetPrimaryStorageLicenseInfo
type GetPrimaryStorageLicenseInfoView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	ExpireTime string `json:"expireTime,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetEncryptedFieldView GetEncryptedField
type GetEncryptedFieldView struct {
	EncryptedFields []string `json:"encryptedFields,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RemoveIAM2ProjectLoginExpiredEventView RemoveIAM2ProjectLoginExpiredEvent
type RemoveIAM2ProjectLoginExpiredEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmStartingCandidateClustersHostsView GetVmStartingCandidateClustersHosts
type GetVmStartingCandidateClustersHostsView struct {
	Hosts []HostInventoryView `json:"hosts,omitempty"`
	Clusters []ClusterInventoryView `json:"clusters,omitempty"`
}

// DiscoverExternalPrimaryStorageEventView DiscoverExternalPrimaryStorageEvent
type DiscoverExternalPrimaryStorageEventView struct {
	Inventory ExternalPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// GetVolumeIoThreadPinView GetVolumeIoThreadPin
type GetVolumeIoThreadPinView struct {
	VolumeUuid string `json:"volumeUuid,omitempty"`
	IoThreadId string `json:"ioThreadId,omitempty"`
	Pin string `json:"pin,omitempty"`
}

// PowerOffHostEventView PowerOffHostEvent
type PowerOffHostEventView struct {
	Results []PowerOffHardwareResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RemoveIAM2VirtualIDGroupFromProjectsEventView RemoveIAM2VirtualIDGroupFromProjectsEvent
type RemoveIAM2VirtualIDGroupFromProjectsEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetIAM2ProjectContainerImageTagsView GetIAM2ProjectContainerImageTags
type GetIAM2ProjectContainerImageTagsView struct {
	Inventories []ContainerImageTagInventoryView `json:"inventories,omitempty"`
	Total int `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteAliyunDiskFromRemoteEventView DeleteAliyunDiskFromRemoteEvent
type DeleteAliyunDiskFromRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVersionView GetVersion
type GetVersionView struct {
	Version string `json:"version,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetCpuMemoryCapacityView GetCpuMemoryCapacity
type GetCpuMemoryCapacityView struct {
	TotalCpu int64 `json:"totalCpu,omitempty"`
	AvailableCpu int64 `json:"availableCpu,omitempty"`
	TotalMemory int64 `json:"totalMemory,omitempty"`
	AvailableMemory int64 `json:"availableMemory,omitempty"`
	ManagedCpuNum int64 `json:"managedCpuNum,omitempty"`
	CapacityData []CpuMemoryCapacityDataView `json:"capacityData,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AddIntegrityResourceEventView AddIntegrityResourceEvent
type AddIntegrityResourceEventView struct {
	Success bool `json:"success,omitempty"`
}

// CheckVipPortAvailabilityView CheckVipPortAvailability
type CheckVipPortAvailabilityView struct {
	Available bool `json:"available,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CheckScsiLunClusterStatusView CheckScsiLunClusterStatus
type CheckScsiLunClusterStatusView struct {
	Inventory ScsiLunClusterStatusInventoryView `json:"inventory,omitempty"`
}

// CheckBatchDataIntegrityView CheckBatchDataIntegrity
type CheckBatchDataIntegrityView struct {
	ResourceMap map[string]bool `json:"resourceMap,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UploadFileToVmEventView UploadFileToVmEvent
type UploadFileToVmEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeL3NetworkDhcpIpAddressEventView ChangeL3NetworkDhcpIpAddressEvent
type ChangeL3NetworkDhcpIpAddressEventView struct {
	DhcpServerIp string `json:"dhcpServerIp,omitempty"`
	Dhcpv6ServerIp string `json:"dhcpv6ServerIp,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CheckVolumeSnapshotGroupAvailabilityView CheckVolumeSnapshotGroupAvailability
type CheckVolumeSnapshotGroupAvailabilityView struct {
	Results []VolumeSnapshotGroupAvailabilityView `json:"results,omitempty"`
}

// SsoClientPushDataEventView SsoClientPushDataEvent
type SsoClientPushDataEventView struct {
	Success bool `json:"success,omitempty"`
}

// BackupDatabaseToPublicCloudEventView BackupDatabaseToPublicCloudEvent
type BackupDatabaseToPublicCloudEventView struct {
	Local string `json:"local,omitempty"`
	Remote string `json:"remote,omitempty"`
	RegionId string `json:"regionId,omitempty"`
}

// LogInView LogIn
type LogInView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// RevertVmFromSnapshotGroupEventView RevertVmFromSnapshotGroupEvent
type RevertVmFromSnapshotGroupEventView struct {
	Results []RevertSnapshotGroupResultView `json:"results,omitempty"`
}

// DetachFirewallRuleSetFromL3EventView DetachFirewallRuleSetFromL3Event
type DetachFirewallRuleSetFromL3EventView struct {
	Success bool `json:"success,omitempty"`
}

// GetLoginCaptchaView GetLoginCaptcha
type GetLoginCaptchaView struct {
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	Captcha string `json:"captcha,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ListVmSchedulingRulesFromExecuteStateView ListVmSchedulingRulesFromExecuteState
type ListVmSchedulingRulesFromExecuteStateView struct {
	Uuids []string `json:"uuids,omitempty"`
}

// SetVmUserDefinedXmlEventView SetVmUserDefinedXmlEvent
type SetVmUserDefinedXmlEventView struct {
	VmUserDefinedXml string `json:"vmUserDefinedXml,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetImageQgaEventView SetImageQgaEvent
type SetImageQgaEventView struct {
	Success bool `json:"success,omitempty"`
}

// ListVMsFromKVMHostEventView ListVMsFromKVMHostEvent
type ListVMsFromKVMHostEventView struct {
	Inventories []VmInstanceInventoryView `json:"inventories,omitempty"`
	LibvirtVersion string `json:"libvirtVersion,omitempty"`
	QemuVersion string `json:"qemuVersion,omitempty"`
	V2vCaps map[string]bool `json:"v2vCaps,omitempty"`
	Success bool `json:"success,omitempty"`
}

// TakeVmConsoleScreenshotEventView TakeVmConsoleScreenshotEvent
type TakeVmConsoleScreenshotEventView struct {
	ImageData string `json:"imageData,omitempty"`
}

// RemoveVRouterNetworksFromOspfAreaEventView RemoveVRouterNetworksFromOspfAreaEvent
type RemoveVRouterNetworksFromOspfAreaEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetAliyunNasMountTargetRemoteView GetAliyunNasMountTargetRemote
type GetAliyunNasMountTargetRemoteView struct {
	Inventories []AliyunNasMountTargetPropertyView `json:"inventories,omitempty"`
}

// TerminateVirtualBorderRouterRemoteEventView TerminateVirtualBorderRouterRemoteEvent
type TerminateVirtualBorderRouterRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteVmBackupEventView DeleteVmBackupEvent
type DeleteVmBackupEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmSecurityLevelEventView SetVmSecurityLevelEvent
type SetVmSecurityLevelEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveMdevDeviceSpecFromVmInstanceEventView RemoveMdevDeviceSpecFromVmInstanceEvent
type RemoveMdevDeviceSpecFromVmInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// RequestConsoleAccessEventView RequestConsoleAccessEvent
type RequestConsoleAccessEventView struct {
	Inventory ConsoleInventoryView `json:"inventory,omitempty"`
}

// UpdateEventDataEventView UpdateEventDataEvent
type UpdateEventDataEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteAliyunRouteEntryRemoteEventView DeleteAliyunRouteEntryRemoteEvent
type DeleteAliyunRouteEntryRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// UngenerateVirtualPciDevicesEventView UngenerateVirtualPciDevicesEvent
type UngenerateVirtualPciDevicesEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachMonitorTriggerActionToTriggerEventView AttachMonitorTriggerActionToTriggerEvent
type AttachMonitorTriggerActionToTriggerEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetAliyunNasFileSystemRemoteView GetAliyunNasFileSystemRemote
type GetAliyunNasFileSystemRemoteView struct {
	Inventories []AliyunNasFileSystemPropertyView `json:"inventories,omitempty"`
}

// SetOrganizationSupervisorEventView SetOrganizationSupervisorEvent
type SetOrganizationSupervisorEventView struct {
	Success bool `json:"success,omitempty"`
}

// ExecuteGuestVmScriptEventView ExecuteGuestVmScriptEvent
type ExecuteGuestVmScriptEventView struct {
	Inventory GuestVmScriptExecutedRecordInventoryView `json:"inventory,omitempty"`
}

// GetIAM2ProjectContainerClusterCandidatesView GetIAM2ProjectContainerClusterCandidates
type GetIAM2ProjectContainerClusterCandidatesView struct {
	Inventories []ContainerClusterInventoryView `json:"inventories,omitempty"`
}

// AttachTagToResourcesEventView AttachTagToResourcesEvent
type AttachTagToResourcesEventView struct {
	Results []AttachTagResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetAuditDataView GetAuditData
type GetAuditDataView struct {
	Audits []AuditDataView `json:"audits,omitempty"`
}

// GetSpiceCertificatesView GetSpiceCertificates
type GetSpiceCertificatesView struct {
	CertificateStr string `json:"certificateStr,omitempty"`
}

// RemoveUserFromGroupEventView RemoveUserFromGroupEvent
type RemoveUserFromGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteEcsVpcRemoteEventView DeleteEcsVpcRemoteEvent
type DeleteEcsVpcRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteFirewallIpSetTemplateEventView DeleteFirewallIpSetTemplateEvent
type DeleteFirewallIpSetTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

// SNSDingTalkTestConnectionEventView SNSDingTalkTestConnectionEvent
type SNSDingTalkTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// ExportImageFromBackupStorageEventView ExportImageFromBackupStorageEvent
type ExportImageFromBackupStorageEventView struct {
	ImageUrl string `json:"imageUrl,omitempty"`
	ExportMd5Sum string `json:"exportMd5Sum,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetModelCenterServicesView GetModelCenterServices
type GetModelCenterServicesView struct {
	Services []ModelCenterServiceInventoryView `json:"services,omitempty"`
}

// DetachMonitorTriggerActionFromTriggerEventView DetachMonitorTriggerActionFromTriggerEvent
type DetachMonitorTriggerActionFromTriggerEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachPolicyRouteRuleSetFromL3EventView DetachPolicyRouteRuleSetFromL3Event
type DetachPolicyRouteRuleSetFromL3EventView struct {
	Success bool `json:"success,omitempty"`
}

// GetInterdependentL3NetworkImageView GetInterdependentL3NetworkImage
type GetInterdependentL3NetworkImageView struct {
	Inventories ListView `json:"inventories,omitempty"`
}

// ValidateVolumeSnapshotChainEventView ValidateVolumeSnapshotChainEvent
type ValidateVolumeSnapshotChainEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetGuestOsMetadataView GetGuestOsMetadata
type GetGuestOsMetadataView struct {
	Inventories []GuestOsCharacterInventoryView `json:"inventories,omitempty"`
}

// AttachRoleToAccountEventView AttachRoleToAccountEvent
type AttachRoleToAccountEventView struct {
	Success bool `json:"success,omitempty"`
}

// LoginByCasView LoginByCas
type LoginByCasView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// SetVRouterRouterIdEventView SetVRouterRouterIdEvent
type SetVRouterRouterIdEventView struct {
	RouterId string `json:"routerId,omitempty"`
}

// ExpungeVmUserDefinedXmlHookScriptEventView ExpungeVmUserDefinedXmlHookScriptEvent
type ExpungeVmUserDefinedXmlHookScriptEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteCdpTaskDataEventView DeleteCdpTaskDataEvent
type DeleteCdpTaskDataEventView struct {
	Success bool `json:"success,omitempty"`
}

// CheckApiPermissionView CheckApiPermission
type CheckApiPermissionView struct {
	Inventory map[string]string `json:"inventory,omitempty"`
}

// GetTextTemplateArgView GetTextTemplateArg
type GetTextTemplateArgView struct {
	DefaultSupportedParams map[string]interface{} `json:"defaultSupportedParams,omitempty"`
}

// DeleteFirewallEventView DeleteFirewallEvent
type DeleteFirewallEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmCapabilitiesView GetVmCapabilities
type GetVmCapabilitiesView struct {
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
}

// DeployDistributedModelServiceEventView DeployDistributedModelServiceEvent
type DeployDistributedModelServiceEventView struct {
	Inventory ModelServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
}

// GetIAM2SystemAttributesView GetIAM2SystemAttributes
type GetIAM2SystemAttributesView struct {
	Inventories []IAM2AttributeInventoryView `json:"inventories,omitempty"`
}

// GetBackupStorageCapacityView GetBackupStorageCapacity
type GetBackupStorageCapacityView struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GenerateSeMdevDevicesEventView GenerateSeMdevDevicesEvent
type GenerateSeMdevDevicesEventView struct {
	Success bool `json:"success,omitempty"`
}

// UndoSnapshotCreationEventView UndoSnapshotCreationEvent
type UndoSnapshotCreationEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// GetIdentityZoneFromRemoteView GetIdentityZoneFromRemote
type GetIdentityZoneFromRemoteView struct {
	Inventories []IdentityZonePropertyView `json:"inventories,omitempty"`
}

// GetEcsInstanceVncUrlView GetEcsInstanceVncUrl
type GetEcsInstanceVncUrlView struct {
	EcsId string `json:"ecsId,omitempty"`
	VncUrl string `json:"vncUrl,omitempty"`
}

// CheckBaremetalChassisConfigFileView CheckBaremetalChassisConfigFile
type CheckBaremetalChassisConfigFileView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteOssBucketFileRemoteEventView DeleteOssBucketFileRemoteEvent
type DeleteOssBucketFileRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetMaaSUsageView GetMaaSUsage
type GetMaaSUsageView struct {
	Usages []MaaSUsageView `json:"usages,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetFreeIpView GetFreeIp
type GetFreeIpView struct {
	Inventories []FreeIpInventoryView `json:"inventories,omitempty"`
}

// DeleteOssBucketRemoteEventView DeleteOssBucketRemoteEvent
type DeleteOssBucketRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// LogInByLdapView LogInByLdap
type LogInByLdapView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
	AccountInventory AccountInventoryView `json:"accountInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CreateL2PortGroupEventView CreateL2PortGroupEvent
type CreateL2PortGroupEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// ValidateInstanceOfferingUserConfigEventView ValidateInstanceOfferingUserConfigEvent
type ValidateInstanceOfferingUserConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmHostnameEventView SetVmHostnameEvent
type SetVmHostnameEventView struct {
	Success bool `json:"success,omitempty"`
}

// TriggerGCJobEventView TriggerGCJobEvent
type TriggerGCJobEventView struct {
	Success bool `json:"success,omitempty"`
}

// CheckBareMetal2ChassisConfigFileView CheckBareMetal2ChassisConfigFile
type CheckBareMetal2ChassisConfigFileView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteVirtualRouterLocalEventView DeleteVirtualRouterLocalEvent
type DeleteVirtualRouterLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteVpcIkeConfigLocalEventView DeleteVpcIkeConfigLocalEvent
type DeleteVpcIkeConfigLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVipUsedPortsView GetVipUsedPorts
type GetVipUsedPortsView struct {
	Inventories []VipPortRangeInventoryView `json:"inventories,omitempty"`
}

// CleanUpStorageTrashOnPrimaryStorageEventView CleanUpStorageTrashOnPrimaryStorageEvent
type CleanUpStorageTrashOnPrimaryStorageEventView struct {
	Result map[string]interface{} `json:"result,omitempty"`
	Total int `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetManagementNodeDirCapacityView GetManagementNodeDirCapacity
type GetManagementNodeDirCapacityView struct {
	Result map[string]interface{} `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UngroupVolumeSnapshotGroupEventView UngroupVolumeSnapshotGroupEvent
type UngroupVolumeSnapshotGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// SubscribeSNSTopicEventView SubscribeSNSTopicEvent
type SubscribeSNSTopicEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmRDPView GetVmRDP
type GetVmRDPView struct {
	Enable bool `json:"enable,omitempty"`
}

// CleanupBillingUsageEventView CleanupBillingUsageEvent
type CleanupBillingUsageEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetCandidateL2NetworksForAttachingClusterView GetCandidateL2NetworksForAttachingCluster
type GetCandidateL2NetworksForAttachingClusterView struct {
	Inventories []L2NetworkDataView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

// IsVfNicAvailableInL3NetworkView IsVfNicAvailableInL3Network
type IsVfNicAvailableInL3NetworkView struct {
	VfNicAvailable bool `json:"vfNicAvailable,omitempty"`
}

// GetAllMetricMetadataView GetAllMetricMetadata
type GetAllMetricMetadataView struct {
	Metrics []MetricStructView `json:"metrics,omitempty"`
}

// SyncVmBackupEventView SyncVmBackupEvent
type SyncVmBackupEventView struct {
	Result SyncBackupResultView `json:"result,omitempty"`
}

// RefreshGuestOsMetadataEventView RefreshGuestOsMetadataEvent
type RefreshGuestOsMetadataEventView struct {
	Success bool `json:"success,omitempty"`
}

// GCAliyunSnapshotRemoteEventView GCAliyunSnapshotRemoteEvent
type GCAliyunSnapshotRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// DownloadBackupFileFromPublicCloudEventView DownloadBackupFileFromPublicCloudEvent
type DownloadBackupFileFromPublicCloudEventView struct {
	Local string `json:"local,omitempty"`
}

// AddIAM2VirtualIDsToProjectsEventView AddIAM2VirtualIDsToProjectsEvent
type AddIAM2VirtualIDsToProjectsEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetIAM2ProjectRetirePolicyEventView SetIAM2ProjectRetirePolicyEvent
type SetIAM2ProjectRetirePolicyEventView struct {
	Success bool `json:"success,omitempty"`
}

// RunIAM2ScriptEventView RunIAM2ScriptEvent
type RunIAM2ScriptEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// DeleteHostNetworkServiceTypeEventView DeleteHostNetworkServiceTypeEvent
type DeleteHostNetworkServiceTypeEventView struct {
	Success bool `json:"success,omitempty"`
}

// SubscribeResNotifyEventView SubscribeResNotifyEvent
type SubscribeResNotifyEventView struct {
	Inventory ResNotifySubscriptionInventoryView `json:"inventory,omitempty"`
}

// ReloadElaborationEventView ReloadElaborationEvent
type ReloadElaborationEventView struct {
	Success bool `json:"success,omitempty"`
}

// ConvertVmFromForeignHypervisorEventView ConvertVmFromForeignHypervisorEvent
type ConvertVmFromForeignHypervisorEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchEventView DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchEvent
type DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchEventView struct {
	Success bool `json:"success,omitempty"`
}

// RestartResourceStackEventView RestartResourceStackEvent
type RestartResourceStackEventView struct {
	Inventory ResourceStackInventoryView `json:"inventory,omitempty"`
}

// AttachPoliciesToUserEventView AttachPoliciesToUserEvent
type AttachPoliciesToUserEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVpcVpnConfigurationFromRemoteView GetVpcVpnConfigurationFromRemote
type GetVpcVpnConfigurationFromRemoteView struct {
	IkeConf VpcVpnIkeConfigStructView `json:"ikeConf,omitempty"`
	IpSecConf VpcVpnIpSecConfigStructView `json:"ipSecConf,omitempty"`
}

// TokenIntrospectionView TokenIntrospection
type TokenIntrospectionView struct {
	Active bool `json:"active,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AddRolesToIAM2VirtualIDEventView AddRolesToIAM2VirtualIDEvent
type AddRolesToIAM2VirtualIDEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateFaultToleranceVmInstanceEventView CreateFaultToleranceVmInstanceEvent
type CreateFaultToleranceVmInstanceEventView struct {
	PrimaryVmInventory VmInstanceInventoryView `json:"primaryVmInventory,omitempty"`
	SecondaryVmInventory VmInstanceInventoryView `json:"secondaryVmInventory,omitempty"`
	FaultToleranceVmGroupInventory VmInstanceInventoryView `json:"faultToleranceVmGroupInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteResourceStackVmPortMonitorEventView DeleteResourceStackVmPortMonitorEvent
type DeleteResourceStackVmPortMonitorEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteGCJobEventView DeleteGCJobEvent
type DeleteGCJobEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteEmailAddressOfSNSEmailEndpointEventView DeleteEmailAddressOfSNSEmailEndpointEvent
type DeleteEmailAddressOfSNSEmailEndpointEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetImagesFromImageStoreBackupStorageView GetImagesFromImageStoreBackupStorage
type GetImagesFromImageStoreBackupStorageView struct {
	Infos []ImageStoreImageStructView `json:"infos,omitempty"`
}

// GetElaborationCategoriesView GetElaborationCategories
type GetElaborationCategoriesView struct {
	Categories []ElaborationCategoryView `json:"categories,omitempty"`
}

// GetHostMultipathTopologyView GetHostMultipathTopology
type GetHostMultipathTopologyView struct {
	Results []MultipathTopologyStructView `json:"results,omitempty"`
}

// DeleteEcsImageRemoteEventView DeleteEcsImageRemoteEvent
type DeleteEcsImageRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetHostNetworkFactsView GetHostNetworkFacts
type GetHostNetworkFactsView struct {
	Bondings []HostNetworkBondingInventoryView `json:"bondings,omitempty"`
	Nics []HostNetworkInterfaceInventoryView `json:"nics,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CleanUpTrashOnBackupStorageEventView CleanUpTrashOnBackupStorageEvent
type CleanUpTrashOnBackupStorageEventView struct {
	Result CleanTrashResultView `json:"result,omitempty"`
	Results []TrashCleanupResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UnlockIdentityView UnlockIdentity
type UnlockIdentityView struct {
	Success bool `json:"success,omitempty"`
}

// SetIAM2ProjectLoginExpiredEventView SetIAM2ProjectLoginExpiredEvent
type SetIAM2ProjectLoginExpiredEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmEmulatorPinningView GetVmEmulatorPinning
type GetVmEmulatorPinningView struct {
	EmulatorPinning string `json:"emulatorPinning,omitempty"`
}

// DeleteAliyunDiskFromLocalEventView DeleteAliyunDiskFromLocalEvent
type DeleteAliyunDiskFromLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetIAM2VirtualIDAPIPermissionView GetIAM2VirtualIDPermission
type GetIAM2VirtualIDAPIPermissionView struct {
	Permissions map[string]PermissionView `json:"permissions,omitempty"`
	NoPermission bool `json:"noPermission,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetOrganizationQuotaUsageView GetOrganizationQuotaUsage
type GetOrganizationQuotaUsageView struct {
	Usages []QuotaUsageView `json:"usages,omitempty"`
}

// GetResourceConfigsView GetResourceConfigs
type GetResourceConfigsView struct {
	Configs []ResourceConfigStructView `json:"configs,omitempty"`
}

// CheckStackTemplateParametersView CheckStackTemplateParameters
type CheckStackTemplateParametersView struct {
	Parameters []StackParametersView `json:"parameters,omitempty"`
	Preparameters []StackParametersView `json:"preparameters,omitempty"`
}

// GetFactoryModeStateView GetFactoryModeState
type GetFactoryModeStateView struct {
	FactoryModeState bool `json:"factoryModeState,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetActiveAlarmStatusView GetActiveAlarmStatus
type GetActiveAlarmStatusView struct {
	Statuses []ActiveAlarmStatusView `json:"statuses,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeployModelEvalServiceEventView DeployModelEvalServiceEvent
type DeployModelEvalServiceEventView struct {
	Inventory ModelEvalServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
	Tasks []ModelEvaluationTaskInventoryView `json:"tasks,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetHostResourceAllocationEventView GetHostResourceAllocationEvent
type GetHostResourceAllocationEventView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	VCPUPin []interface{} `json:"vCPUPin,omitempty"`
}

// GetLicenseAddOnsView GetLicenseAddOns
type GetLicenseAddOnsView struct {
	Addons []LicenseAddOnInventoryView `json:"addons,omitempty"`
}

// GetVpcIPsecLogView GetVpcIPsecLog
type GetVpcIPsecLogView struct {
	IpsecLog string `json:"ipsecLog,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AddResourcesToDirectoryEventView AddResourcesToDirectoryEvent
type AddResourcesToDirectoryEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachGuestToolsIsoToVmEventView AttachGuestToolsIsoToVmEvent
type AttachGuestToolsIsoToVmEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachAliyunKeyEventView DetachAliyunKeyEvent
type DetachAliyunKeyEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteBuildAppEventView DeleteBuildAppEvent
type DeleteBuildAppEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetResourceFromPublishAppView GetResourceFromPublishApp
type GetResourceFromPublishAppView struct {
	Resources []PublishAppResourceStructView `json:"resources,omitempty"`
}

// GetHostIommuStateView GetHostIommuState
type GetHostIommuStateView struct {
	State string `json:"state,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteVirtualBorderRouterLocalEventView DeleteVirtualBorderRouterLocalEvent
type DeleteVirtualBorderRouterLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetMetricDataView GetMetricData
type GetMetricDataView struct {
	Data []DatapointView `json:"data,omitempty"`
}

// EnableCbtTaskEventView EnableCbtTaskEvent
type EnableCbtTaskEventView struct {
	VolumeCbtBackupInfos []VolumeCbtBackupInfoView `json:"volumeCbtBackupInfos,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetAliyunNasAccessGroupRemoteView GetAliyunNasAccessGroupRemote
type GetAliyunNasAccessGroupRemoteView struct {
	Inventories []AliyunNasAccessGroupPropertyView `json:"inventories,omitempty"`
}

// CheckBuildAppParametersView CheckBuildAppParameters
type CheckBuildAppParametersView struct {
	Parameters []StackParametersView `json:"parameters,omitempty"`
}

// GetVpcVRouterDistributedRoutingConnectionsView GetVpcVRouterDistributedRoutingConnections
type GetVpcVRouterDistributedRoutingConnectionsView struct {
	Inventories map[string]interface{} `json:"inventories,omitempty"`
}

// UpdateThirdpartyAlertsEventView UpdateThirdpartyAlertsEvent
type UpdateThirdpartyAlertsEventView struct {
	Success bool `json:"success,omitempty"`
}

// PullSdnControllerTenantEventView PullSdnControllerTenantEvent
type PullSdnControllerTenantEventView struct {
	Inventories []H3cSdnControllerTenantInventoryView `json:"inventories,omitempty"`
}

// SetVmUsbRedirectEventView SetVmUsbRedirectEvent
type SetVmUsbRedirectEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateFactoryModeStateEventView UpdateFactoryModeStateEvent
type UpdateFactoryModeStateEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateChronyServersEventView UpdateChronyServersEvent
type UpdateChronyServersEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachPolicyRouteRuleSetToL3EventView AttachPolicyRouteRuleSetToL3Event
type AttachPolicyRouteRuleSetToL3EventView struct {
	Success bool `json:"success,omitempty"`
}

// GetZWatchAlertHistogramView GetZWatchAlertHistogram
type GetZWatchAlertHistogramView struct {
	Histograms []HistogramView `json:"histograms,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteAliyunRouterInterfaceRemoteEventView DeleteAliyunRouterInterfaceRemoteEvent
type DeleteAliyunRouterInterfaceRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetImageBootModeEventView SetImageBootModeEvent
type SetImageBootModeEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmsCapabilitiesView GetVmsCapabilities
type GetVmsCapabilitiesView struct {
	VmsCaps map[string]VmCapabilitiesView `json:"vmsCaps,omitempty"`
}

// AttachPolicyToUserGroupEventView AttachPolicyToUserGroupEvent
type AttachPolicyToUserGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// RevokeMonitorTemplateFromMonitorGroupEventView RevokeMonitorTemplateFromMonitorGroupEvent
type RevokeMonitorTemplateFromMonitorGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteFirewallRuleEventView DeleteFirewallRuleEvent
type DeleteFirewallRuleEventView struct {
	Success bool `json:"success,omitempty"`
}

// ShareResourceEventView ShareResourceEvent
type ShareResourceEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetAccountQuotaUsageView GetAccountQuotaUsage
type GetAccountQuotaUsageView struct {
	Usages []QuotaUsageView `json:"usages,omitempty"`
}

// RemoveIAM2VirtualIDsFromProjectsEventView RemoveIAM2VirtualIDsFromProjectsEvent
type RemoveIAM2VirtualIDsFromProjectsEventView struct {
	Success bool `json:"success,omitempty"`
}

// SubmitLongJobEventView SubmitLongJobEvent
type SubmitLongJobEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// DegradeFromLicenseServerEventView DegradeFromLicenseServerEvent
type DegradeFromLicenseServerEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetDebugSignalView GetDebugSignal
type GetDebugSignalView struct {
	Signals []string `json:"signals,omitempty"`
}

// SyncLicenseCapacityEventView SyncLicenseCapacityEvent
type SyncLicenseCapacityEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachDataVolumeToHostEventView AttachDataVolumeToHostEvent
type AttachDataVolumeToHostEventView struct {
	Success bool `json:"success,omitempty"`
}

// SecurityMachineEncryptEventView SecurityMachineEncryptEvent
type SecurityMachineEncryptEventView struct {
	Text string `json:"text,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetVpcVRouterNetworkServiceStateView GetVpcVRouterNetworkServiceState
type GetVpcVRouterNetworkServiceStateView struct {
	State string `json:"state,omitempty"`
}

// DeleteContainerResourceFromEndpointEventView DeleteContainerResourceFromEndpointEvent
type DeleteContainerResourceFromEndpointEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetSupportAPIsView GetSupports
type GetSupportAPIsView struct {
	SupportApis []string `json:"supportApis,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteVpcUserVpnGatewayRemoteEventView DeleteVpcUserVpnGatewayRemoteEvent
type DeleteVpcUserVpnGatewayRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddResourceStackVmPortMonitorEventView AddResourceStackVmPortMonitorEvent
type AddResourceStackVmPortMonitorEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVpcVRouterNetworkServiceStateEventView SetVpcVRouterNetworkServiceStateEvent
type SetVpcVRouterNetworkServiceStateEventView struct {
	State string `json:"state,omitempty"`
}

// GetVmXmlView GetVmXml
type GetVmXmlView struct {
	Match bool `json:"match,omitempty"`
	RunningXml string `json:"runningXml,omitempty"`
	UserDefinedXml string `json:"userDefinedXml,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetVmInstanceFirstBootDeviceView GetVmInstanceFirstBootDevice
type GetVmInstanceFirstBootDeviceView struct {
	FirstBootDevice string `json:"firstBootDevice,omitempty"`
}

// DeleteIpAddressEventView DeleteIpAddressEvent
type DeleteIpAddressEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteVpcVpnConnectionRemoteEventView DeleteVpcVpnConnectionRemoteEvent
type DeleteVpcVpnConnectionRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// CheckIAM2OrganizationAvailabilityView CheckIAM2OrganizationAvailability
type CheckIAM2OrganizationAvailabilityView struct {
	Exists bool `json:"exists,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UnmountVmInstanceRecoveryPointEventView UnmountVmInstanceRecoveryPointEvent
type UnmountVmInstanceRecoveryPointEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemovePolicyStatementsFromRoleEventView RemovePolicyStatementsFromRoleEvent
type RemovePolicyStatementsFromRoleEventView struct {
	Success bool `json:"success,omitempty"`
}

// GenerateModelMetadataEventView GenerateModelMetadataEvent
type GenerateModelMetadataEventView struct {
	Success bool `json:"success,omitempty"`
}

// IsReadyToGoView IsReadyToGo
type IsReadyToGoView struct {
	ManagementNodeId string `json:"managementNodeId,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetHostIommuStatusView GetHostIommuStatus
type GetHostIommuStatusView struct {
	Status string `json:"status,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DescribeVmInstanceRecoveryPointView DescribeVmInstanceRecoveryPoint
type DescribeVmInstanceRecoveryPointView struct {
	RealSizes map[string]int64 `json:"realSizes,omitempty"`
	VirtualSizes map[string]int64 `json:"virtualSizes,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetBareMetal2ChassisPowerStatusView GetBareMetal2ChassisPowerStatus
type GetBareMetal2ChassisPowerStatusView struct {
	Status string `json:"status,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetTaskProgressView GetTaskProgress
type GetTaskProgressView struct {
	Inventories []TaskProgressInventoryView `json:"inventories,omitempty"`
}

// StartDataProtectionEventView StartDataProtectionEvent
type StartDataProtectionEventView struct {
	Success bool `json:"success,omitempty"`
}

// ValidateSessionView ValidateSession
type ValidateSessionView struct {
	Valid bool `json:"valid,omitempty"`
}

// ChangeActiveAlarmStateEventView ChangeActiveAlarmStateEvent
type ChangeActiveAlarmStateEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmCleanTrafficEventView SetVmCleanTrafficEvent
type SetVmCleanTrafficEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmBootModeEventView SetVmBootModeEvent
type SetVmBootModeEventView struct {
	Success bool `json:"success,omitempty"`
}

// PublishAppEventView PublishAppEvent
type PublishAppEventView struct {
	Inventory PublishAppInventoryView `json:"inventory,omitempty"`
}

// ProtectVmInstanceRecoveryPointEventView ProtectVmInstanceRecoveryPointEvent
type ProtectVmInstanceRecoveryPointEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteConnectionAccessPointLocalEventView DeleteConnectionAccessPointLocalEvent
type DeleteConnectionAccessPointLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveIAM2VirtualIDsFromProjectEventView RemoveIAM2VirtualIDsFromProjectEvent
type RemoveIAM2VirtualIDsFromProjectEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetSharedBlockCandidateView GetSharedBlockCandidate
type GetSharedBlockCandidateView struct {
	Results []SharedBlockCandidateStructView `json:"results,omitempty"`
}

// ReclaimSpaceFromImageStoreEventView ReclaimSpaceFromImageStoreEvent
type ReclaimSpaceFromImageStoreEventView struct {
	GcResult ImageStoreGcResultView `json:"gcResult,omitempty"`
}

// GetAllEventMetadataView GetAllEventMetadata
type GetAllEventMetadataView struct {
	Events []EventStructView `json:"events,omitempty"`
}

// DeleteDataVolumeEventView DeleteDataVolumeEvent
type DeleteDataVolumeEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetUploadImageJobDetailsView GetUploadImageJobDetails
type GetUploadImageJobDetailsView struct {
	ExistingJobDetails []JobDetailsView `json:"existingJobDetails,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DetachHybridEipFromEcsEventView DetachHybridEipFromEcsEvent
type DetachHybridEipFromEcsEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVolumeCapabilitiesView GetVolumeCapabilities
type GetVolumeCapabilitiesView struct {
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
}

// SetVmInstanceHaLevelEventView SetVmInstanceHaLevelEvent
type SetVmInstanceHaLevelEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveVRouterNetworksFromFlowMeterEventView RemoveVRouterNetworksFromFlowMeterEvent
type RemoveVRouterNetworksFromFlowMeterEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteHybridKeySecretEventView DeleteHybridKeySecretEvent
type DeleteHybridKeySecretEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetCandidatePrimaryStoragesForCreatingVmView GetCandidatePrimaryStoragesForCreatingVm
type GetCandidatePrimaryStoragesForCreatingVmView struct {
	RootVolumePrimaryStorages []PrimaryStorageInventoryView `json:"rootVolumePrimaryStorages,omitempty"`
	DataVolumePrimaryStorages map[string]interface{} `json:"dataVolumePrimaryStorages,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetVmConsolePasswordView GetVmConsolePassword
type GetVmConsolePasswordView struct {
	Password string `json:"password,omitempty"`
}

// GetResourceBindableConfigView GetResourceBindableConfig
type GetResourceBindableConfigView struct {
	BindableConfigs []ResourceBindableConfigStructView `json:"bindableConfigs,omitempty"`
}

// GetVmInstanceHaLevelView GetVmInstanceHaLevel
type GetVmInstanceHaLevelView struct {
	Level string `json:"level,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RemoveResourcesFromDirectoryEventView RemoveResourcesFromDirectoryEvent
type RemoveResourcesFromDirectoryEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteExportedDatabaseBackupFromBackupStorageEventView DeleteExportedDatabaseBackupFromBackupStorageEvent
type DeleteExportedDatabaseBackupFromBackupStorageEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnexportNbdVolumesEventView UnexportNbdVolumesEvent
type UnexportNbdVolumesEventView struct {
	Success bool `json:"success,omitempty"`
}

// RecoveryVirtualBorderRouterRemoteEventView RecoveryVirtualBorderRouterRemoteEvent
type RecoveryVirtualBorderRouterRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// ExecuteAutoScalingRuleEventView ExecuteAutoScalingRuleEvent
type ExecuteAutoScalingRuleEventView struct {
	ScalingActivityUuid string `json:"scalingActivityUuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SNSHttpTestConnectionEventView SNSHttpTestConnectionEvent
type SNSHttpTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp string `json:"webhookResp,omitempty"`
}

// SetImageSecurityLevelEventView SetImageSecurityLevelEvent
type SetImageSecurityLevelEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachVmFromVmSchedulingRuleGroupEventView DetachVmFromVmSchedulingRuleGroupEvent
type DetachVmFromVmSchedulingRuleGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddRolesToIAM2VirtualIDGroupEventView AddRolesToIAM2VirtualIDGroupEvent
type AddRolesToIAM2VirtualIDGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// CheckStaticProvisionIpView CheckStaticProvisionIp
type CheckStaticProvisionIpView struct {
	Success bool `json:"success,omitempty"`
}

// PushLicenseAddOnsUsageEventView PushLicenseAddOnsUsageEvent
type PushLicenseAddOnsUsageEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetBareMetal2ProvisionNetworkIpAddressCapacityView GetBareMetal2ProvisionNetworkIpAddressCapacity
type GetBareMetal2ProvisionNetworkIpAddressCapacityView struct {
	CapacityData []BareMetal2ProvisionNetworkIpCapacityView `json:"capacityData,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DecodeStackTemplateView DecodeStackTemplate
type DecodeStackTemplateView struct {
	Resources []ResourceStructView `json:"resources,omitempty"`
}

// GetVSwitchTypesView GetVSwitchTypes
type GetVSwitchTypesView struct {
	Types []string `json:"types,omitempty"`
}

// CreateL2HardwareVxlanNetworkPoolEventView CreateL2HardwareVxlanNetworkPoolEvent
type CreateL2HardwareVxlanNetworkPoolEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// GetLdapServerAvailableAttributesView GetLdapServerAvailableAttributes
type GetLdapServerAvailableAttributesView struct {
	Inventories ListView `json:"inventories,omitempty"`
}

// BatchQueryView BatchQuery
type BatchQueryView struct {
	Result map[string]interface{} `json:"result,omitempty"`
}

// ReloadExternalServiceEventView ReloadExternalServiceEvent
type ReloadExternalServiceEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddIAM2VirtualIDToGroupEventView AddIAM2VirtualIDToGroupEvent
type AddIAM2VirtualIDToGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVolumeSnapshotSizeEventView GetVolumeSnapshotSizeEvent
type GetVolumeSnapshotSizeEventView struct {
	Size int64 `json:"size,omitempty"`
	ActualSize int64 `json:"actualSize,omitempty"`
	Success bool `json:"success,omitempty"`
}

// BatchSyncVolumeSizeView BatchSyncVolumeSize
type BatchSyncVolumeSizeView struct {
	SuccessCount int `json:"successCount,omitempty"`
	FailCount int `json:"failCount,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetHypervisorTypesView GetHypervisorTypes
type GetHypervisorTypesView struct {
	HypervisorTypes []string `json:"hypervisorTypes,omitempty"`
}

// GetVmMonitorNumberView GetVmMonitorNumber
type GetVmMonitorNumberView struct {
	MonitorNumber int `json:"monitorNumber,omitempty"`
}

// ValidatePriceUserConfigEventView ValidatePriceUserConfigEvent
type ValidatePriceUserConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveActionFromEventSubscriptionEventView RemoveActionFromEventSubscriptionEvent
type RemoveActionFromEventSubscriptionEventView struct {
	Success bool `json:"success,omitempty"`
}

// CheckHostConfigFileView CheckHostConfigFile
type CheckHostConfigFileView struct {
	Success bool `json:"success,omitempty"`
}

// GetContainerUsageView GetContainerUsage
type GetContainerUsageView struct {
	Usages []ContainerUsageView `json:"usages,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SNSSnmpTestConnectionEventView SNSSnmpTestConnectionEvent
type SNSSnmpTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// GetDataCenterFromRemoteView GetDataCenterFromRemote
type GetDataCenterFromRemoteView struct {
	Inventories []DataCenterPropertyView `json:"inventories,omitempty"`
}

// DeleteEcsImageLocalEventView DeleteEcsImageLocalEvent
type DeleteEcsImageLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetBackupStorageTypesView GetBackupStorageTypes
type GetBackupStorageTypesView struct {
	Types []string `json:"types,omitempty"`
}

// GetVolumeQosView GetVolumeQos
type GetVolumeQosView struct {
	VolumeUuid string `json:"volumeUuid,omitempty"`
	VolumeBandwidth int64 `json:"volumeBandwidth,omitempty"`
	VolumeBandwidthRead int64 `json:"volumeBandwidthRead,omitempty"`
	VolumeBandwidthWrite int64 `json:"volumeBandwidthWrite,omitempty"`
	IopsTotal int64 `json:"iopsTotal,omitempty"`
	IopsRead int64 `json:"iopsRead,omitempty"`
	IopsWrite int64 `json:"iopsWrite,omitempty"`
	VolumeBandwidthUpthreshold int64 `json:"volumeBandwidthUpthreshold,omitempty"`
	VolumeBandwidthReadUpthreshold int64 `json:"volumeBandwidthReadUpthreshold,omitempty"`
	VolumeBandwidthWriteUpthreshold int64 `json:"volumeBandwidthWriteUpthreshold,omitempty"`
	IopsTotalUpthreshold int64 `json:"iopsTotalUpthreshold,omitempty"`
	IopsReadUpthreshold int64 `json:"iopsReadUpthreshold,omitempty"`
	IopsWriteUpthreshold int64 `json:"iopsWriteUpthreshold,omitempty"`
	Success bool `json:"success,omitempty"`
}

// PowerOnBaremetalChassisEventView PowerOnBaremetalChassisEvent
type PowerOnBaremetalChassisEventView struct {
	Success bool `json:"success,omitempty"`
}

// RequestLicenseCapacityEventView RequestLicenseCapacityEvent
type RequestLicenseCapacityEventView struct {
	Inventory LicenseAuthorizedCapacityInventoryView `json:"inventory,omitempty"`
}

// GetVirtualizerInfoView GetVirtualizerInfo
type GetVirtualizerInfoView struct {
	Inventories []VirtualizerInfoInventoryView `json:"inventories,omitempty"`
}

// GetL3NetworkIpStatisticView GetL3NetworkIpStatistic
type GetL3NetworkIpStatisticView struct {
	IpStatistics []IpStatisticDataView `json:"ipStatistics,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

// KvmRunShellEventView KvmRunShellEvent
type KvmRunShellEventView struct {
	Inventory map[string]ShellResultView `json:"inventory,omitempty"`
}

// ExpungeDataVolumeEventView ExpungeDataVolumeEvent
type ExpungeDataVolumeEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVRouterRouterIdView GetVRouterRouterId
type GetVRouterRouterIdView struct {
	RouterId string `json:"routerId,omitempty"`
}

// GetZBoxBackupDetailsView GetZBoxBackupDetails
type GetZBoxBackupDetailsView struct {
	VmBackupInfos []VmExternalBackupInfoView `json:"vmBackupInfos,omitempty"`
	VolumeBackupInfos []VolumeExternalBackupInfoView `json:"volumeBackupInfos,omitempty"`
	BackupStorageBackupInfos []BackupStorageExternalBackupInfoView `json:"backupStorageBackupInfos,omitempty"`
	Version string `json:"version,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetExternalServicesView GetExternalServices
type GetExternalServicesView struct {
	Inventories []ExternalServiceInventoryView `json:"inventories,omitempty"`
}

// GetIAM2ProjectRepositoryView GetIAM2ProjectRepository
type GetIAM2ProjectRepositoryView struct {
	Inventories []ProjectRepositoryInventoryView `json:"inventories,omitempty"`
}

// GetCandidateNetworkInterfacesView GetCandidateNetworkInterfaces
type GetCandidateNetworkInterfacesView struct {
	SlaveNames []string `json:"slaveNames,omitempty"`
	CandidateNics []HostNetworkInterfaceInventoryView `json:"candidateNics,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ChangeAccessControlListServerGroupEventView ChangeAccessControlListServerGroupEvent
type ChangeAccessControlListServerGroupEventView struct {
	Inventory LoadBalancerListerAclView `json:"inventory,omitempty"`
}

// CreateL2HardwareVxlanNetworkEventView CreateL2HardwareVxlanNetworkEvent
type CreateL2HardwareVxlanNetworkEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// GetGlobalConfigOptionsView GetGlobalConfigOptions
type GetGlobalConfigOptionsView struct {
	Options GlobalConfigOptionsView `json:"options,omitempty"`
}

// PutMetricDataEventView PutMetricDataEvent
type PutMetricDataEventView struct {
	Success bool `json:"success,omitempty"`
}

// RerunLongJobEventView RerunLongJobEvent
type RerunLongJobEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// DeleteExportedImageFromBackupStorageEventView DeleteExportedImageFromBackupStorageEvent
type DeleteExportedImageFromBackupStorageEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmUsbRedirectView GetVmUsbRedirect
type GetVmUsbRedirectView struct {
	Enable bool `json:"enable,omitempty"`
}

// GetOssBucketFileFromRemoteView GetOssBucketFileFromRemote
type GetOssBucketFileFromRemoteView struct {
	Files []string `json:"files,omitempty"`
}

// AttachVipToVpcSharedQosEventView AttachVipToVpcSharedQosEvent
type AttachVipToVpcSharedQosEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetEventDataView GetEventData
type GetEventDataView struct {
	Events []EventDataView `json:"events,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CheckIpAvailabilityView CheckIpAvailability
type CheckIpAvailabilityView struct {
	Available bool `json:"available,omitempty"`
	Reason string `json:"reason,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RemoveRolesFromIAM2VirtualIDEventView RemoveRolesFromIAM2VirtualIDEvent
type RemoveRolesFromIAM2VirtualIDEventView struct {
	Success bool `json:"success,omitempty"`
}

// CalculateResourceSpendingView CalculateResourceSpending
type CalculateResourceSpendingView struct {
	Spending []ResourceSpendingView `json:"spending,omitempty"`
	Pagination PaginationView `json:"pagination,omitempty"`
	Success bool `json:"success,omitempty"`
}

// PowerResetBaremetalChassisEventView PowerResetBaremetalChassisEvent
type PowerResetBaremetalChassisEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanUpTrashOnPrimaryStorageEventView CleanUpTrashOnPrimaryStorageEvent
type CleanUpTrashOnPrimaryStorageEventView struct {
	Result CleanTrashResultView `json:"result,omitempty"`
	Results []TrashCleanupResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetVmSchedulingRulesExecuteStateView GetVmSchedulingRulesExecuteState
type GetVmSchedulingRulesExecuteStateView struct {
	RuleMapState map[string]string `json:"ruleMapState,omitempty"`
}

// GetIpAddressCapacityView GetIpAddressCapacity
type GetIpAddressCapacityView struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	UsedIpAddressNumber int64 `json:"usedIpAddressNumber,omitempty"`
	Ipv4TotalCapacity int64 `json:"ipv4TotalCapacity,omitempty"`
	Ipv4AvailableCapacity int64 `json:"ipv4AvailableCapacity,omitempty"`
	Ipv4UsedIpAddressNumber int64 `json:"ipv4UsedIpAddressNumber,omitempty"`
	Ipv6TotalCapacity int64 `json:"ipv6TotalCapacity,omitempty"`
	Ipv6AvailableCapacity int64 `json:"ipv6AvailableCapacity,omitempty"`
	Ipv6UsedIpAddressNumber int64 `json:"ipv6UsedIpAddressNumber,omitempty"`
	CapacityData []IpCapacityDataView `json:"capacityData,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetIAM2ProjectContainerClusterEventView SetIAM2ProjectContainerClusterEvent
type SetIAM2ProjectContainerClusterEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeployAppDevelopmentServiceEventView DeployAppDevelopmentServiceEvent
type DeployAppDevelopmentServiceEventView struct {
	Inventory ModelServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
	App ApplicationDevelopmentServiceInventoryView `json:"app,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RefreshPluginDriversEventView RefreshPluginDriversEvent
type RefreshPluginDriversEventView struct {
	Success bool `json:"success,omitempty"`
}

// PauseVmInstanceEventView PauseVmInstanceEvent
type PauseVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// DetachUserDefinedXmlHookScriptFromVmEventView DetachUserDefinedXmlHookScriptFromVmEvent
type DetachUserDefinedXmlHookScriptFromVmEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetSignatureServerEncryptPublicKeyView GetSignatureServerEncryptPublicKey
type GetSignatureServerEncryptPublicKeyView struct {
	PublicKey string `json:"publicKey,omitempty"`
}

// DeleteFirewallRuleSetEventView DeleteFirewallRuleSetEvent
type DeleteFirewallRuleSetEventView struct {
	Success bool `json:"success,omitempty"`
}

// LocalStorageMigrateVolumeEventView LocalStorageMigrateVolumeEvent
type LocalStorageMigrateVolumeEventView struct {
	Inventory LocalStorageResourceRefInventoryView `json:"inventory,omitempty"`
}

// SetOrganizationOperationEventView SetOrganizationOperationEvent
type SetOrganizationOperationEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveIAM2VirtualIDsFromOrganizationEventView RemoveIAM2VirtualIDsFromOrganizationEvent
type RemoveIAM2VirtualIDsFromOrganizationEventView struct {
	Success bool `json:"success,omitempty"`
}

// ExportDatabaseBackupFromBackupStorageEventView ExportDatabaseBackupFromBackupStorageEvent
type ExportDatabaseBackupFromBackupStorageEventView struct {
	DatabaseBackupUrl string `json:"databaseBackupUrl,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetVpcVRouterDistributedRoutingEnabledEventView SetVpcVRouterDistributedRoutingEnabledEvent
type SetVpcVRouterDistributedRoutingEnabledEventView struct {
	Enabled bool `json:"enabled,omitempty"`
}

// PowerOnBareMetal2ChassisEventView PowerOnBareMetal2ChassisEvent
type PowerOnBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// GetLocalRaidPhysicalDriveSmartView GetLocalRaidPhysicalDriveSmart
type GetLocalRaidPhysicalDriveSmartView struct {
	Result []SmartDataStructView `json:"result,omitempty"`
}

// PullHuaweiIMasterControllerEventView PullHuaweiIMasterControllerEvent
type PullHuaweiIMasterControllerEventView struct {
	Inventories []HuaweiIMasterSdnControllerInventoryView `json:"inventories,omitempty"`
}

// RemoveRolesFromIAM2VirtualIDGroupEventView RemoveRolesFromIAM2VirtualIDGroupEvent
type RemoveRolesFromIAM2VirtualIDGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeIAM2OrganizationParentEventView ChangeIAM2OrganizationParentEvent
type ChangeIAM2OrganizationParentEventView struct {
	Success bool `json:"success,omitempty"`
}

// SNSWeComTestConnectionEventView SNSWeComTestConnectionEvent
type SNSWeComTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// ProvisionVirtualRouterConfigEventView ProvisionVirtualRouterConfigEvent
type ProvisionVirtualRouterConfigEventView struct {
	Inventory ApplianceVmInventoryView `json:"inventory,omitempty"`
}

// SetVmQgaEventView SetVmQgaEvent
type SetVmQgaEventView struct {
	Success bool `json:"success,omitempty"`
}

// ValidatePasswordView ValidatePassword
type ValidatePasswordView struct {
	Available bool `json:"available,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetChronyServersView GetChronyServers
type GetChronyServersView struct {
	Servers []ChronyServerInfoPairView `json:"servers,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetVmQxlMemoryEventView SetVmQxlMemoryEvent
type SetVmQxlMemoryEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVolumeFormatView GetVolumeFormat
type GetVolumeFormatView struct {
	Formats []VolumeFormatReplyStructView `json:"formats,omitempty"`
}

// GetResourceAccountView GetResourceAccount
type GetResourceAccountView struct {
	Inventories map[string]AccountInventoryView `json:"inventories,omitempty"`
}

// BindModelToServiceEventView BindModelToServiceEvent
type BindModelToServiceEventView struct {
	Inventory ModelServiceInventoryView `json:"inventory,omitempty"`
}

// CheckNetworkReachableView CheckNetworkReachable
type CheckNetworkReachableView struct {
	Results []NetworkReachablePairView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetFlowMeterRouterIdEventView SetFlowMeterRouterIdEvent
type SetFlowMeterRouterIdEventView struct {
	RouterId int64 `json:"routerId,omitempty"`
}

// AddStorageProtocolEventView AddStorageProtocolEvent
type AddStorageProtocolEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeployModelServiceEventView DeployModelServiceEvent
type DeployModelServiceEventView struct {
	Inventory ModelServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
}

// GetMonitorItemView GetMonitorItem
type GetMonitorItemView struct {
	Inventories []ItemInventoryView `json:"inventories,omitempty"`
}

// UnregisterLicenseRequestedApplicationEventView UnregisterLicenseRequestedApplicationEvent
type UnregisterLicenseRequestedApplicationEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeVmPasswordEventView ChangeVmPasswordEvent
type ChangeVmPasswordEventView struct {
	Success bool `json:"success,omitempty"`
}

// FlattenVmInstanceEventView FlattenVmInstanceEvent
type FlattenVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// DeleteAllEcsInstancesFromDataCenterEventView DeleteAllEcsInstancesFromDataCenterEvent
type DeleteAllEcsInstancesFromDataCenterEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVpcMulticastRouteView GetVpcMulticastRoute
type GetVpcMulticastRouteView struct {
	Inventories []MulticastRouteInventoryView `json:"inventories,omitempty"`
}

// DeleteVmUserDefinedXmlHookScriptEventView DeleteVmUserDefinedXmlHookScriptEvent
type DeleteVmUserDefinedXmlHookScriptEventView struct {
	Success bool `json:"success,omitempty"`
}

// CheckResourcePermissionView CheckResourcePermission
type CheckResourcePermissionView struct {
	Apis []string `json:"apis,omitempty"`
}

// GetCandidateMiniHostsView GetCandidateMiniHosts
type GetCandidateMiniHostsView struct {
	Hosts []MiniCandidateHostStructView `json:"hosts,omitempty"`
}

// DeleteDatasetsEventView DeleteDatasetsEvent
type DeleteDatasetsEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RevokeResourceSharingEventView RevokeResourceSharingEvent
type RevokeResourceSharingEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteModelServicesEventView DeleteModelServicesEvent
type DeleteModelServicesEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
}

// GetHostNUMATopologyEventView GetHostNUMATopologyEvent
type GetHostNUMATopologyEventView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Topology map[string]HostNUMANodeView `json:"topology,omitempty"`
}

// CreateL2VirtualSwitchEventView CreateL2VirtualSwitchEvent
type CreateL2VirtualSwitchEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// GetClusterDRSStatusView GetClusterDRSStatus
type GetClusterDRSStatusView struct {
	HostLoadOverThreshold []HostLoadView `json:"hostLoadOverThreshold,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetVmNumaView GetVmNuma
type GetVmNumaView struct {
	Enable bool `json:"enable,omitempty"`
}

// DeletePluginDriversEventView DeletePluginDriversEvent
type DeletePluginDriversEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachPolicyFromRoleEventView DetachPolicyFromRoleEvent
type DetachPolicyFromRoleEventView struct {
	Success bool `json:"success,omitempty"`
}

// RestartModelServiceGroupsEventView RestartModelServiceGroupsEvent
type RestartModelServiceGroupsEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
}

// GetLoadBalancerOwnerView GetLoadBalancerOwner
type GetLoadBalancerOwnerView struct {
	Type string `json:"type,omitempty"`
	Vpc VpcRouterVmInventoryView `json:"vpc,omitempty"`
	VpcHa VpcHaGroupInventoryView `json:"vpcHa,omitempty"`
	Slb SlbGroupInventoryView `json:"slb,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetNicQosView GetNicQos
type GetNicQosView struct {
	OutboundBandwidth int64 `json:"outboundBandwidth,omitempty"`
	InboundBandwidth int64 `json:"inboundBandwidth,omitempty"`
	OutboundBandwidthUpthreshold int64 `json:"outboundBandwidthUpthreshold,omitempty"`
	InboundBandwidthUpthreshold int64 `json:"inboundBandwidthUpthreshold,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CreateBareMetal2ChassisHardwareView CreateBareMetal2ChassisHardware
type CreateBareMetal2ChassisHardwareView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteIAM2VirtualIDLdapBindingEventView DeleteIAM2VirtualIDLdapBindingEvent
type DeleteIAM2VirtualIDLdapBindingEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateVmPriorityEventView UpdateVmPriorityEvent
type UpdateVmPriorityEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteVmHostnameEventView DeleteVmHostnameEvent
type DeleteVmHostnameEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetLicenseCapabilitiesView GetLicenseCapabilities
type GetLicenseCapabilitiesView struct {
	Capabilities map[string]string `json:"capabilities,omitempty"`
}

// SetVmSoundTypeEventView SetVmSoundTypeEvent
type SetVmSoundTypeEventView struct {
	Success bool `json:"success,omitempty"`
}

// MergeDataOnBackupStorageEventView MergeDataOnBackupStorageEvent
type MergeDataOnBackupStorageEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetCdpBackupStorageRequirementView GetCdpBackupStorageRequirement
type GetCdpBackupStorageRequirementView struct {
	NextStep string `json:"nextStep,omitempty"`
	Required map[string]string `json:"required,omitempty"`
	Current map[string]string `json:"current,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AddAttributesToIAM2VirtualIDGroupEventView AddAttributesToIAM2VirtualIDGroupEvent
type AddAttributesToIAM2VirtualIDGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddVmNicToSecurityGroupEventView AddVmNicToSecurityGroupEvent
type AddVmNicToSecurityGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetMetricLabelValueView GetMetricLabelValue
type GetMetricLabelValueView struct {
	Labels []interface{} `json:"labels,omitempty"`
}

// GetCandidateZonesClustersHostsForCreatingVmView GetCandidateZonesClustersHostsForCreatingVm
type GetCandidateZonesClustersHostsForCreatingVmView struct {
	Zones []ZoneInventoryView `json:"zones,omitempty"`
	Clusters []ClusterInventoryView `json:"clusters,omitempty"`
	Hosts []HostInventoryView `json:"hosts,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RemoveSchedulerJobGroupFromSchedulerTriggerEventView RemoveSchedulerJobGroupFromSchedulerTriggerEvent
type RemoveSchedulerJobGroupFromSchedulerTriggerEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetL3NetworkDhcpIpAddressView GetL3NetworkDhcpIpAddress
type GetL3NetworkDhcpIpAddressView struct {
	Ip string `json:"ip,omitempty"`
	Ip6 string `json:"ip6,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetBaremetalChassisPowerStatusView GetBaremetalChassisPowerStatus
type GetBaremetalChassisPowerStatusView struct {
	Status string `json:"status,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetFaultToleranceVmsView GetFaultToleranceVms
type GetFaultToleranceVmsView struct {
	PrimaryVmInventory VmInstanceInventoryView `json:"primaryVmInventory,omitempty"`
	SecondaryVmInventory VmInstanceInventoryView `json:"secondaryVmInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteAliyunKeySecretEventView DeleteAliyunKeySecretEvent
type DeleteAliyunKeySecretEventView struct {
	Success bool `json:"success,omitempty"`
}

// PowerResetBareMetal2ChassisEventView PowerResetBareMetal2ChassisEvent
type PowerResetBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// PrometheusQueryVmMonitoringDataView PrometheusQueryVmMonitoringData
type PrometheusQueryVmMonitoringDataView struct {
	Inventories MapView `json:"inventories,omitempty"`
}

// UpdateResourceConfigsEventView UpdateResourceConfigsEvent
type UpdateResourceConfigsEventView struct {
	Inventories []ResourceConfigStructView `json:"inventories,omitempty"`
}

// RevertVolumeFromSnapshotEventView RevertVolumeFromSnapshotEvent
type RevertVolumeFromSnapshotEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetManagementNodeArchView GetManagementNodeArch
type GetManagementNodeArchView struct {
	Architecture string `json:"architecture,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DisableCbtTaskEventView DisableCbtTaskEvent
type DisableCbtTaskEventView struct {
	Inventory CbtTaskInventoryView `json:"inventory,omitempty"`
}

// FailoverFaultToleranceVmEventView FailoverFaultToleranceVmEvent
type FailoverFaultToleranceVmEventView struct {
	PrimaryVmInventory VmInstanceInventoryView `json:"primaryVmInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// EjectZBoxEventView EjectZBoxEvent
type EjectZBoxEventView struct {
	Inventory ZBoxInventoryView `json:"inventory,omitempty"`
}

// PrometheusQueryMetadataView PrometheusQueryMetadata
type PrometheusQueryMetadataView struct {
	Inventories MapView `json:"inventories,omitempty"`
}

// DeleteFirewallRuleTemplateEventView DeleteFirewallRuleTemplateEvent
type DeleteFirewallRuleTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

// ExecuteGuestVmCommandEventView ExecuteGuestVmCommandEvent
type ExecuteGuestVmCommandEventView struct {
	Stream string `json:"stream,omitempty"`
	VmInstance VmInstanceInventoryView `json:"vmInstance,omitempty"`
}

// RemoveAttributesFromIAM2VirtualIDGroupEventView RemoveAttributesFromIAM2VirtualIDGroupEvent
type RemoveAttributesFromIAM2VirtualIDGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddAttributesToIAM2VirtualIDEventView AddAttributesToIAM2VirtualIDEvent
type AddAttributesToIAM2VirtualIDEventView struct {
	Success bool `json:"success,omitempty"`
}

// FlattenVolumeEventView FlattenVolumeEvent
type FlattenVolumeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// DeleteEcsSecurityGroupRuleRemoteEventView DeleteEcsSecurityGroupRuleRemoteEvent
type DeleteEcsSecurityGroupRuleRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachAliyunDiskFromEcsEventView DetachAliyunDiskFromEcsEvent
type DetachAliyunDiskFromEcsEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachHostFromHostSchedulingRuleGroupEventView DetachHostFromHostSchedulingRuleGroupEvent
type DetachHostFromHostSchedulingRuleGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateAliyunRouteInterfaceRemoteEventView UpdateAliyunRouteInterfaceRemoteEvent
type UpdateAliyunRouteInterfaceRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// PrometheusQueryPassThroughView PrometheusQueryPassThrough
type PrometheusQueryPassThroughView struct {
	Inventories MapView `json:"inventories,omitempty"`
}

// GetVmDeviceAddressView GetVmDeviceAddress
type GetVmDeviceAddressView struct {
	Addresses map[string]interface{} `json:"addresses,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RemoveInstanceFromMonitorGroupEventView RemoveInstanceFromMonitorGroupEvent
type RemoveInstanceFromMonitorGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanQueueEventView CleanQueueEvent
type CleanQueueEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveLabelFromEventSubscriptionEventView RemoveLabelFromEventSubscriptionEvent
type RemoveLabelFromEventSubscriptionEventView struct {
	Success bool `json:"success,omitempty"`
}

// SdnControllerRemoveHostEventView SdnControllerRemoveHostEvent
type SdnControllerRemoveHostEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// DetachCCSCertificateFromUserEventView DetachCCSCertificateFromUserEvent
type DetachCCSCertificateFromUserEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetManagementNodeOSView GetManagementNodeOS
type GetManagementNodeOSView struct {
	Name string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ExecuteDRSSchedulingEventView ExecuteDRSSchedulingEvent
type ExecuteDRSSchedulingEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmQgaView GetVmQga
type GetVmQgaView struct {
	Enable bool `json:"enable,omitempty"`
}

// PreviewResourceStackView PreviewResourceStack
type PreviewResourceStackView struct {
	Preview PreviewResourceStructView `json:"preview,omitempty"`
}

// GetVmvNUMATopologyView GetVmvNUMATopology
type GetVmvNUMATopologyView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	Topology []interface{} `json:"topology,omitempty"`
}

// RemoveSchedulerJobsFromSchedulerJobGroupEventView RemoveSchedulerJobsFromSchedulerJobGroupEvent
type RemoveSchedulerJobsFromSchedulerJobGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// GenerateVirtualPciDevicesEventView GenerateVirtualPciDevicesEvent
type GenerateVirtualPciDevicesEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetPrimaryStorageAllocatorStrategiesView GetPrimaryStorageAllocatorStrategies
type GetPrimaryStorageAllocatorStrategiesView struct {
	Strategies []string `json:"strategies,omitempty"`
}

// GetPlatformTimeZoneView GetPlatformTimeZone
type GetPlatformTimeZoneView struct {
	Timezone string `json:"timezone,omitempty"`
	Offset string `json:"offset,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DetachPolicyFromUserEventView DetachPolicyFromUserEvent
type DetachPolicyFromUserEventView struct {
	Success bool `json:"success,omitempty"`
}

// FstrimVmEventView FstrimVmEvent
type FstrimVmEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncAINginxConfigurationView SyncAINginxConfiguration
type SyncAINginxConfigurationView struct {
	UnSyncedRules []NginxRedirectRuleView `json:"unSyncedRules,omitempty"`
	Success bool `json:"success,omitempty"`
}

// MatchModelServiceTemplateWithModelEventView MatchModelServiceTemplateWithModelEvent
type MatchModelServiceTemplateWithModelEventView struct {
	Result map[string]interface{} `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AddPolicyStatementsToRoleEventView AddPolicyStatementsToRoleEvent
type AddPolicyStatementsToRoleEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnprotectVmInstanceRecoveryPointEventView UnprotectVmInstanceRecoveryPointEvent
type UnprotectVmInstanceRecoveryPointEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachVipFromVpcSharedQosEventView DetachVipFromVpcSharedQosEvent
type DetachVipFromVpcSharedQosEventView struct {
	Success bool `json:"success,omitempty"`
}

// PrimaryStorageMigrateVmEventView PrimaryStorageMigrateVmEvent
type PrimaryStorageMigrateVmEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// RecoverDatabaseFromBackupEventView RecoverDatabaseFromBackupEvent
type RecoverDatabaseFromBackupEventView struct {
	LogListenPort int `json:"logListenPort,omitempty"`
	Success bool `json:"success,omitempty"`
}

// MoveDirectoryEventView MoveDirectoryEvent
type MoveDirectoryEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVRouterOspfNeighborView GetVRouterOspfNeighbor
type GetVRouterOspfNeighborView struct {
	Neighbors []NeighborView `json:"neighbors,omitempty"`
}

// GetFlowMeterRouterIdView GetFlowMeterRouterId
type GetFlowMeterRouterIdView struct {
	RouterId int64 `json:"routerId,omitempty"`
}

// AddResourceToIAM2ProjectEventView AddResourceToIAM2ProjectEvent
type AddResourceToIAM2ProjectEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetAlarmDataView GetAlarmData
type GetAlarmDataView struct {
	Histories []AlarmDataView `json:"histories,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RecoverResourceSplitBrainEventView RecoverResourceSplitBrainEvent
type RecoverResourceSplitBrainEventView struct {
	Success bool `json:"success,omitempty"`
}

// IsOpensourceVersionView IsOpensourceVersion
type IsOpensourceVersionView struct {
	Opensource bool `json:"opensource,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetResourceFromResourceStackView GetResourceFromResourceStack
type GetResourceFromResourceStackView struct {
	Resources []interface{} `json:"resources,omitempty"`
}

// MoveResourcesToDirectoryEventView MoveResourcesToDirectoryEvent
type MoveResourcesToDirectoryEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetSupportedCloudFormationResourcesView GetSupportedCloudFormationResources
type GetSupportedCloudFormationResourcesView struct {
	Resources []SupportedResourceStructView `json:"resources,omitempty"`
}

// DeleteIdentityZoneInLocalEventView DeleteIdentityZoneInLocalEvent
type DeleteIdentityZoneInLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnregisterLicenseServerEventView UnregisterLicenseServerEvent
type UnregisterLicenseServerEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpgradeToLicenseServerEventView UpgradeToLicenseServerEvent
type UpgradeToLicenseServerEventView struct {
	Inventory LicenseAuthorizedNodeInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DetachAppBuildSystemToZoneEventView DetachAppBuildSystemToZoneEvent
type DetachAppBuildSystemToZoneEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetAppBuildSystemCapacityView GetAppBuildSystemCapacity
type GetAppBuildSystemCapacityView struct {
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetL3NetworkMtuView GetL3NetworkMtu
type GetL3NetworkMtuView struct {
	Mtu int `json:"mtu,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetVmMonitorNumberEventView SetVmMonitorNumberEvent
type SetVmMonitorNumberEventView struct {
	Success bool `json:"success,omitempty"`
}

// SNSMicrosoftTeamsTestConnectionEventView SNSMicrosoftTeamsTestConnectionEvent
type SNSMicrosoftTeamsTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// CreateOssBackupBucketRemoteEventView CreateOssBackupBucketRemoteEvent
type CreateOssBackupBucketRemoteEventView struct {
	BucketName string `json:"bucketName,omitempty"`
	Success bool `json:"success,omitempty"`
}

// PowerOffBaremetalChassisEventView PowerOffBaremetalChassisEvent
type PowerOffBaremetalChassisEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetCandidateInterfaceVlanIdsView GetCandidateInterfaceVlanIds
type GetCandidateInterfaceVlanIdsView struct {
	VlanIds []int `json:"vlanIds,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetNetworkServiceTypesView GetNetworkServiceTypes
type GetNetworkServiceTypesView struct {
	Types map[string]interface{} `json:"types,omitempty"`
}

// DeleteVmUserDefinedXmlEventView DeleteVmUserDefinedXmlEvent
type DeleteVmUserDefinedXmlEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetAvailableVpcL3NetworkView GetAvailableVpcL3Network
type GetAvailableVpcL3NetworkView struct {
	Inventories ListView `json:"inventories,omitempty"`
}

// GetCurrentTimeView GetCurrentTime
type GetCurrentTimeView struct {
	CurrentTime map[string]int64 `json:"currentTime,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CalculateAccountSpendingView CalculateAccountSpending
type CalculateAccountSpendingView struct {
	Total float64 `json:"total,omitempty"`
	Spending []SpendingView `json:"spending,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UpdateEcsInstanceVncPasswordEventView UpdateEcsInstanceVncPasswordEvent
type UpdateEcsInstanceVncPasswordEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncChronyServersEventView SyncChronyServersEvent
type SyncChronyServersEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmInstanceProtectedRecoveryPointsView GetVmInstanceProtectedRecoveryPoints
type GetVmInstanceProtectedRecoveryPointsView struct {
	RecoveryPoints map[string]interface{} `json:"recoveryPoints,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AddVmToVmSchedulingRuleGroupEventView AddVmToVmSchedulingRuleGroupEvent
type AddVmToVmSchedulingRuleGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetHostWebSshUrlEventView GetHostWebSshUrlEvent
type GetHostWebSshUrlEventView struct {
	Url string `json:"url,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetL3NetworkMtuEventView SetL3NetworkMtuEvent
type SetL3NetworkMtuEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetL3NetworkRouterInterfaceIpView GetL3NetworkRouterInterfaceIp
type GetL3NetworkRouterInterfaceIpView struct {
	RouterInterfaceIp string `json:"routerInterfaceIp,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SyncVmClockEventView SyncVmClockEvent
type SyncVmClockEventView struct {
	Success bool `json:"success,omitempty"`
}

// SdnControllerAddHostEventView SdnControllerAddHostEvent
type SdnControllerAddHostEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// RenewSessionEventView RenewSessionEvent
type RenewSessionEventView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// DeleteDataCenterInLocalEventView DeleteDataCenterInLocalEvent
type DeleteDataCenterInLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachPolicyToUserEventView AttachPolicyToUserEvent
type AttachPolicyToUserEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVfPciDeviceAvailableInL2NetworkView GetVfPciDeviceAvailableInL2Network
type GetVfPciDeviceAvailableInL2NetworkView struct {
	L2VfAvailableClusters map[string]interface{} `json:"l2VfAvailableClusters,omitempty"`
}

// AddAttributesToIAM2ProjectEventView AddAttributesToIAM2ProjectEvent
type AddAttributesToIAM2ProjectEventView struct {
	Success bool `json:"success,omitempty"`
}

// UngenerateSeMdevDevicesEventView UngenerateSeMdevDevicesEvent
type UngenerateSeMdevDevicesEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmEmulatorPinningEventView SetVmEmulatorPinningEvent
type SetVmEmulatorPinningEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanV2VConversionCacheEventView CleanV2VConversionCacheEvent
type CleanV2VConversionCacheEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnbindModelFromServiceEventView UnbindModelFromServiceEvent
type UnbindModelFromServiceEventView struct {
	Inventory ModelServiceInventoryView `json:"inventory,omitempty"`
}

// GetEcsInstanceTypeView GetEcsInstanceType
type GetEcsInstanceTypeView struct {
	Types []EcsInstanceTypeView `json:"types,omitempty"`
}

// GetLicenseUKeyStatusEventView GetLicenseUKeyStatusEvent
type GetLicenseUKeyStatusEventView struct {
	Inventories []UKeyInventoryView `json:"inventories,omitempty"`
}

// SetL3NetworkRouterInterfaceIpEventView SetL3NetworkRouterInterfaceIpEvent
type SetL3NetworkRouterInterfaceIpEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddTicketTypesToTicketFlowCollectionEventView AddTicketTypesToTicketFlowCollectionEvent
type AddTicketTypesToTicketFlowCollectionEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetConnectionBetweenL3NetworkAndAliyunVSwitchView GetConnectionBetweenL3NetworkAndAliyunVSwitch
type GetConnectionBetweenL3NetworkAndAliyunVSwitchView struct {
	Inventories []ConnectionRelationShipPropertyView `json:"inventories,omitempty"`
}

// GetBareMetal2GatewayAllocatorStrategiesView GetBareMetal2GatewayAllocatorStrategies
type GetBareMetal2GatewayAllocatorStrategiesView struct {
	Strategies []string `json:"strategies,omitempty"`
}

// WithdrawLicenseCapacityApplicationEventView WithdrawLicenseCapacityApplicationEvent
type WithdrawLicenseCapacityApplicationEventView struct {
	Success bool `json:"success,omitempty"`
}

// PowerResetHostEventView PowerResetHostEvent
type PowerResetHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RevertVmFromVmBackupEventView RevertVmFromVmBackupEvent
type RevertVmFromVmBackupEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmNumaEventView SetVmNumaEvent
type SetVmNumaEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteAliyunRouterInterfaceLocalEventView DeleteAliyunRouterInterfaceLocalEvent
type DeleteAliyunRouterInterfaceLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachAliyunKeyEventView AttachAliyunKeyEvent
type AttachAliyunKeyEventView struct {
	Success bool `json:"success,omitempty"`
}

// RefreshSearchIndexesView RefreshSearchIndexes
type RefreshSearchIndexesView struct {
	Success bool `json:"success,omitempty"`
}

// GetL2NetworkTypesView GetL2NetworkTypes
type GetL2NetworkTypesView struct {
	Types []string `json:"types,omitempty"`
}

// ShutdownHostEventView ShutdownHostEvent
type ShutdownHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DisableCdpTaskEventView DisableCdpTaskEvent
type DisableCdpTaskEventView struct {
	Inventory CdpTaskInventoryView `json:"inventory,omitempty"`
}

// RemoveAttributesFromIAM2VirtualIDEventView RemoveAttributesFromIAM2VirtualIDEvent
type RemoveAttributesFromIAM2VirtualIDEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachRoleFromAccountEventView DetachRoleFromAccountEvent
type DetachRoleFromAccountEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteLdapBindingEventView DeleteLdapBindingEvent
type DeleteLdapBindingEventView struct {
	Success bool `json:"success,omitempty"`
}

// DebugSignalEventView DebugSignalEvent
type DebugSignalEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVpcVRouterDistributedRoutingEnabledView GetVpcVRouterDistributedRoutingEnabled
type GetVpcVRouterDistributedRoutingEnabledView struct {
	Enabled bool `json:"enabled,omitempty"`
}

// RemoveAttributesFromIAM2OrganizationEventView RemoveAttributesFromIAM2OrganizationEvent
type RemoveAttributesFromIAM2OrganizationEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteAliyunSnapshotFromLocalEventView DeleteAliyunSnapshotFromLocalEvent
type DeleteAliyunSnapshotFromLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetIAM2ProjectContainerImagesView GetIAM2ProjectContainerImages
type GetIAM2ProjectContainerImagesView struct {
	Inventories []ZakuImageInventoryView `json:"inventories,omitempty"`
	Total int `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteOssBucketNameLocalEventView DeleteOssBucketNameLocalEvent
type DeleteOssBucketNameLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetObservabilityServerServiceDataView GetObservabilityServerServiceData
type GetObservabilityServerServiceDataView struct {
	Inventories []ObservabilityServerServiceDataInventoryView `json:"inventories,omitempty"`
	Total int `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

// VerifyLicenseServerEventView VerifyLicenseServerEvent
type VerifyLicenseServerEventView struct {
	AccessKeyId string `json:"accessKeyId,omitempty"`
	AccessKeySecret string `json:"accessKeySecret,omitempty"`
	LicenseClient LicenseAuthorizedNodeInventoryView `json:"licenseClient,omitempty"`
	LicenseServer LicenseAuthorizedNodeInventoryView `json:"licenseServer,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteEcsSecurityGroupInLocalEventView DeleteEcsSecurityGroupInLocalEvent
type DeleteEcsSecurityGroupInLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachDataVolumeFromHostEventView DetachDataVolumeFromHostEvent
type DetachDataVolumeFromHostEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVmInstanceRecoveryPointsView GetVmInstanceRecoveryPoints
type GetVmInstanceRecoveryPointsView struct {
	RecoveryPoints map[string]interface{} `json:"recoveryPoints,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetOssBackupBucketFromRemoteView GetOssBackupBucketFromRemote
type GetOssBackupBucketFromRemoteView struct {
	Buckets []OssBucketFilesPropertyView `json:"buckets,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetL3NetworkTypesView GetL3NetworkTypes
type GetL3NetworkTypesView struct {
	Types []string `json:"types,omitempty"`
}

// DetachPoliciesFromUserEventView DetachPoliciesFromUserEvent
type DetachPoliciesFromUserEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanUpImageCacheOnPrimaryStorageEventView CleanUpImageCacheOnPrimaryStorageEvent
type CleanUpImageCacheOnPrimaryStorageEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddHostFromConfigFileEventView AddHostFromConfigFileEvent
type AddHostFromConfigFileEventView struct {
	Results []AddHostFromFileResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteVmBootModeEventView DeleteVmBootModeEvent
type DeleteVmBootModeEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteVpcVpnConnectionLocalEventView DeleteVpcVpnConnectionLocalEvent
type DeleteVpcVpnConnectionLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachPolicyFromUserGroupEventView DetachPolicyFromUserGroupEvent
type DetachPolicyFromUserGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// ZQLQueryView ZQLQuery
type ZQLQueryView struct {
	Results []ZQLQueryReturnView `json:"results,omitempty"`
}

// GetElaborationsView GetElaborations
type GetElaborationsView struct {
	Contents []ElaborationContentView `json:"contents,omitempty"`
}

// GetAccessPathView GetAccessPath
type GetAccessPathView struct {
	PathInfos []AccessPathInfoView `json:"pathInfos,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetPrimaryStorageUsageReportView GetPrimaryStorageUsageReport
type GetPrimaryStorageUsageReportView struct {
	UriUsageForecast map[string]UsageReportView `json:"uriUsageForecast,omitempty"`
	UsageReport UsageReportView `json:"usageReport,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RevertVolumeFromVolumeBackupEventView RevertVolumeFromVolumeBackupEvent
type RevertVolumeFromVolumeBackupEventView struct {
	Success bool `json:"success,omitempty"`
}

// LocalStorageGetVolumeMigratableView LocalStorageGetVolumeMigratable
type LocalStorageGetVolumeMigratableView struct {
	Inventories []HostInventoryView `json:"inventories,omitempty"`
}

// GetOssBucketNameFromRemoteView GetOssBucketNameFromRemote
type GetOssBucketNameFromRemoteView struct {
	Inventories []OssBucketPropertyView `json:"inventories,omitempty"`
}

// AttachUserDefinedXmlHookScriptToVmEventView AttachUserDefinedXmlHookScriptToVmEvent
type AttachUserDefinedXmlHookScriptToVmEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachPolicyToRoleEventView AttachPolicyToRoleEvent
type AttachPolicyToRoleEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteVpcIpSecConfigLocalEventView DeleteVpcIpSecConfigLocalEvent
type DeleteVpcIpSecConfigLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

// CalculateAccountBillingSpendingView CalculateAccountBillingSpending
type CalculateAccountBillingSpendingView struct {
	Total float64 `json:"total,omitempty"`
	Spending []SpendingView `json:"spending,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteVRouterOspfAreaEventView DeleteVRouterOspfAreaEvent
type DeleteVRouterOspfAreaEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVipAvailablePortView GetVipAvailablePort
type GetVipAvailablePortView struct {
	AvailablePort []int `json:"availablePort,omitempty"`
	Success bool `json:"success,omitempty"`
}

// MountVmInstanceRecoveryPointEventView MountVmInstanceRecoveryPointEvent
type MountVmInstanceRecoveryPointEventView struct {
	ResourcePath string `json:"resourcePath,omitempty"`
	FailedVolumes map[string]string `json:"failedVolumes,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetResourceStackFromResourceView GetResourceStackFromResource
type GetResourceStackFromResourceView struct {
	Stack map[string]string `json:"stack,omitempty"`
}

// CheckIAM2VirtualIDConfigFileView CheckIAM2VirtualIDConfigFile
type CheckIAM2VirtualIDConfigFileView struct {
	Results []ErrorResultView `json:"results,omitempty"`
}

// GetClusterHostNetworkFactsView GetClusterHostNetworkFacts
type GetClusterHostNetworkFactsView struct {
	Bondings []HostNetworkBondingInventoryView `json:"bondings,omitempty"`
	Nics []HostNetworkInterfaceInventoryView `json:"nics,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ParseOvfView ParseOvf
type ParseOvfView struct {
	OvfInfo OvfInfoView `json:"ovfInfo,omitempty"`
}

// DeleteAliyunNasAccessGroupRuleEventView DeleteAliyunNasAccessGroupRuleEvent
type DeleteAliyunNasAccessGroupRuleEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetLoginProceduresView GetLoginProcedures
type GetLoginProceduresView struct {
	Procedures []LoginAuthenticationProcedureDescView `json:"procedures,omitempty"`
}

// DeleteBondingEventView DeleteBondingEvent
type DeleteBondingEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteEcsSecurityGroupRemoteEventView DeleteEcsSecurityGroupRemoteEvent
type DeleteEcsSecurityGroupRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteVmNicFromSecurityGroupEventView DeleteVmNicFromSecurityGroupEvent
type DeleteVmNicFromSecurityGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// EnableCdpTaskEventView EnableCdpTaskEvent
type EnableCdpTaskEventView struct {
	Inventory CdpTaskInventoryView `json:"inventory,omitempty"`
}

// RegisterLicenseRequestedApplicationEventView RegisterLicenseRequestedApplicationEvent
type RegisterLicenseRequestedApplicationEventView struct {
	AppId string `json:"appId,omitempty"`
	ServicePubKey string `json:"servicePubKey,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeleteEcsVpcInLocalEventView DeleteEcsVpcInLocalEvent
type DeleteEcsVpcInLocalEventView struct {
	Success bool `json:"success,omitempty"`
}

