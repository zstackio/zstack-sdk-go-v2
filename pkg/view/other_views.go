// Copyright (c) ZStack.io, Inc.

package view

// AckAlertDataEventView AckAlertDataEvent
type AckAlertDataEventView struct {
	Inventory AlertDataAckInventoryView `json:"inventory,omitempty"`
}

// AllocateHostResourceEventView AllocateHostResourceEvent
type AllocateHostResourceEventView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	VCPUPin []interface{} `json:"vCPUPin,omitempty"`
}

// ApplyDRSAdviceEventView ApplyDRSAdviceEvent
type ApplyDRSAdviceEventView struct {
	VmMigrationActivityUuid string `json:"vmMigrationActivityUuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ApplyMonitorTemplateToMonitorGroupEventView ApplyMonitorTemplateToMonitorGroupEvent
type ApplyMonitorTemplateToMonitorGroupEventView struct {
	Inventory MonitorGroupTemplateRefInventoryView `json:"inventory,omitempty"`
}

// ApplyRuleSetChangesEventView ApplyRuleSetChangesEvent
type ApplyRuleSetChangesEventView struct {
	Inventory VpcFirewallRuleSetInventoryView `json:"inventory,omitempty"`
}

// ApplyTemplateConfigEventView ApplyTemplateConfigEvent
type ApplyTemplateConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// BackupDatabaseToPublicCloudEventView BackupDatabaseToPublicCloudEvent
type BackupDatabaseToPublicCloudEventView struct {
	Local string `json:"local,omitempty"`
	Remote string `json:"remote,omitempty"`
	RegionId string `json:"regionId,omitempty"`
}

// BackupStorageMigrateImageEventView BackupStorageMigrateImageEvent
type BackupStorageMigrateImageEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// BatchAddBareMetal2ChassisEventView BatchAddBareMetal2ChassisEvent
type BatchAddBareMetal2ChassisEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// BatchCreateBaremetalChassisEventView BatchCreateBaremetalChassisEvent
type BatchCreateBaremetalChassisEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// BatchCreateIAM2VirtualIDFromConfigFileEventView BatchCreateIAM2VirtualIDFromConfigFileEvent
type BatchCreateIAM2VirtualIDFromConfigFileEventView struct {
	NumberOfImportedUser int `json:"numberOfImportedUser,omitempty"`
}

// BatchDeleteVolumeSnapshotEventView BatchDeleteVolumeSnapshotEvent
type BatchDeleteVolumeSnapshotEventView struct {
	Results []BatchDeleteVolumeSnapshotStructView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// BatchQueryView BatchQuery
type BatchQueryView struct {
	Result map[string]interface{} `json:"result,omitempty"`
}

// BatchSyncVolumeSizeView BatchSyncVolumeSize
type BatchSyncVolumeSizeView struct {
	SuccessCount int `json:"successCount,omitempty"`
	FailCount int `json:"failCount,omitempty"`
	Success bool `json:"success,omitempty"`
}

// BindModelToServiceEventView BindModelToServiceEvent
type BindModelToServiceEventView struct {
	Inventory ModelServiceInventoryView `json:"inventory,omitempty"`
}

// BootstrapMiniHostEventView BootstrapMiniHostEvent
type BootstrapMiniHostEventView struct {
	Stage string `json:"stage,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CalculateAccountBillingSpendingView CalculateAccountBillingSpending
type CalculateAccountBillingSpendingView struct {
	Total float64 `json:"total,omitempty"`
	Spending []SpendingView `json:"spending,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CalculateAccountSpendingView CalculateAccountSpending
type CalculateAccountSpendingView struct {
	Total float64 `json:"total,omitempty"`
	Spending []SpendingView `json:"spending,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CalculateImageHashEventView CalculateImageHashEvent
type CalculateImageHashEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// CalculateResourceSpendingView CalculateResourceSpending
type CalculateResourceSpendingView struct {
	Spending []ResourceSpendingView `json:"spending,omitempty"`
	Pagination PaginationView `json:"pagination,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CancelLongJobEventView CancelLongJobEvent
type CancelLongJobEventView struct {
	Success bool `json:"success,omitempty"`
}

// CheckApiPermissionView CheckApiPermission
type CheckApiPermissionView struct {
	Inventory map[string]string `json:"inventory,omitempty"`
}

// CheckBareMetal2ChassisConfigFileView CheckBareMetal2ChassisConfigFile
type CheckBareMetal2ChassisConfigFileView struct {
	Success bool `json:"success,omitempty"`
}

// CheckBaremetalChassisConfigFileView CheckBaremetalChassisConfigFile
type CheckBaremetalChassisConfigFileView struct {
	Success bool `json:"success,omitempty"`
}

// CheckBatchDataIntegrityView CheckBatchDataIntegrity
type CheckBatchDataIntegrityView struct {
	ResourceMap map[string]bool `json:"resourceMap,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CheckBuildAppParametersView CheckBuildAppParameters
type CheckBuildAppParametersView struct {
	Parameters []StackParametersView `json:"parameters,omitempty"`
}

// CheckElaborationContentView CheckElaborationContent
type CheckElaborationContentView struct {
	Results []ElaborationCheckResultView `json:"results,omitempty"`
}

// CheckFirewallRuleConfigFileView CheckFirewallRuleConfigFile
type CheckFirewallRuleConfigFileView struct {
	Success bool `json:"success,omitempty"`
}

// CheckIAM2OrganizationAvailabilityView CheckIAM2OrganizationAvailability
type CheckIAM2OrganizationAvailabilityView struct {
	Exists bool `json:"exists,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CheckIAM2VirtualIDConfigFileView CheckIAM2VirtualIDConfigFile
type CheckIAM2VirtualIDConfigFileView struct {
	Results []ErrorResultView `json:"results,omitempty"`
}

// CheckIpAvailabilityView CheckIpAvailability
type CheckIpAvailabilityView struct {
	Available bool `json:"available,omitempty"`
	Reason string `json:"reason,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CheckHostConfigFileView CheckHostConfigFile
type CheckHostConfigFileView struct {
	Success bool `json:"success,omitempty"`
}

// CheckNetworkReachableView CheckNetworkReachable
type CheckNetworkReachableView struct {
	Results []NetworkReachablePairView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CheckResourcePermissionView CheckResourcePermission
type CheckResourcePermissionView struct {
	Apis []string `json:"apis,omitempty"`
}

// CheckScsiLunClusterStatusView CheckScsiLunClusterStatus
type CheckScsiLunClusterStatusView struct {
	Inventory ScsiLunClusterStatusInventoryView `json:"inventory,omitempty"`
}

// CheckStackTemplateParametersView CheckStackTemplateParameters
type CheckStackTemplateParametersView struct {
	Parameters []StackParametersView `json:"parameters,omitempty"`
	Preparameters []StackParametersView `json:"preparameters,omitempty"`
}

// CheckStaticProvisionIpView CheckStaticProvisionIp
type CheckStaticProvisionIpView struct {
	Success bool `json:"success,omitempty"`
}

// CheckVipPortAvailabilityView CheckVipPortAvailability
type CheckVipPortAvailabilityView struct {
	Available bool `json:"available,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CheckVolumeSnapshotGroupAvailabilityView CheckVolumeSnapshotGroupAvailability
type CheckVolumeSnapshotGroupAvailabilityView struct {
	Results []VolumeSnapshotGroupAvailabilityView `json:"results,omitempty"`
}

// CleanInvalidLdapBindingEventView CleanInvalidLdapBindingEvent
type CleanInvalidLdapBindingEventView struct {
	Inventories []AccountInventoryView `json:"inventories,omitempty"`
}

// CleanInvalidLdapIAM2BindingEventView CleanInvalidLdapIAM2BindingEvent
type CleanInvalidLdapIAM2BindingEventView struct {
	Inventories []IAM2VirtualIDInventoryView `json:"inventories,omitempty"`
}

// CleanLongJobEventView CleanLongJobEvent
type CleanLongJobEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanQueueEventView CleanQueueEvent
type CleanQueueEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanUpBaremetal2BondingEventView CleanUpBaremetal2BondingEvent
type CleanUpBaremetal2BondingEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanUpBaremetalChassisBondingEventView CleanUpBaremetalChassisBondingEvent
type CleanUpBaremetalChassisBondingEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanUpImageCacheOnPrimaryStorageEventView CleanUpImageCacheOnPrimaryStorageEvent
type CleanUpImageCacheOnPrimaryStorageEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanUpStorageTrashOnPrimaryStorageEventView CleanUpStorageTrashOnPrimaryStorageEvent
type CleanUpStorageTrashOnPrimaryStorageEventView struct {
	Result map[string]interface{} `json:"result,omitempty"`
	Total int `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CleanUpTrashOnBackupStorageEventView CleanUpTrashOnBackupStorageEvent
type CleanUpTrashOnBackupStorageEventView struct {
	Result CleanTrashResultView `json:"result,omitempty"`
	Results []TrashCleanupResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CleanUpTrashOnPrimaryStorageEventView CleanUpTrashOnPrimaryStorageEvent
type CleanUpTrashOnPrimaryStorageEventView struct {
	Result CleanTrashResultView `json:"result,omitempty"`
	Results []TrashCleanupResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CleanV2VConversionCacheEventView CleanV2VConversionCacheEvent
type CleanV2VConversionCacheEventView struct {
	Success bool `json:"success,omitempty"`
}

// CleanupBillingUsageEventView CleanupBillingUsageEvent
type CleanupBillingUsageEventView struct {
	Success bool `json:"success,omitempty"`
}

// ConvertVmFromForeignHypervisorEventView ConvertVmFromForeignHypervisorEvent
type ConvertVmFromForeignHypervisorEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// DebugSignalEventView DebugSignalEvent
type DebugSignalEventView struct {
	Success bool `json:"success,omitempty"`
}

// DecodeStackTemplateView DecodeStackTemplate
type DecodeStackTemplateView struct {
	Resources []ResourceStructView `json:"resources,omitempty"`
}

// DegradeFromLicenseServerEventView DegradeFromLicenseServerEvent
type DegradeFromLicenseServerEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeployAppDevelopmentServiceEventView DeployAppDevelopmentServiceEvent
type DeployAppDevelopmentServiceEventView struct {
	Inventory ModelServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
	App ApplicationDevelopmentServiceInventoryView `json:"app,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeployDistributedModelServiceEventView DeployDistributedModelServiceEvent
type DeployDistributedModelServiceEventView struct {
	Inventory ModelServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
}

// DeployModelEvalServiceEventView DeployModelEvalServiceEvent
type DeployModelEvalServiceEventView struct {
	Inventory ModelEvalServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
	Tasks []ModelEvaluationTaskInventoryView `json:"tasks,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DeployModelServiceEventView DeployModelServiceEvent
type DeployModelServiceEventView struct {
	Inventory ModelServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
}

// DescribeVmInstanceRecoveryPointView DescribeVmInstanceRecoveryPoint
type DescribeVmInstanceRecoveryPointView struct {
	RealSizes map[string]int64 `json:"realSizes,omitempty"`
	VirtualSizes map[string]int64 `json:"virtualSizes,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DisableCbtTaskEventView DisableCbtTaskEvent
type DisableCbtTaskEventView struct {
	Inventory CbtTaskInventoryView `json:"inventory,omitempty"`
}

// DisableCdpTaskEventView DisableCdpTaskEvent
type DisableCdpTaskEventView struct {
	Inventory CdpTaskInventoryView `json:"inventory,omitempty"`
}

// DiscoverExternalPrimaryStorageEventView DiscoverExternalPrimaryStorageEvent
type DiscoverExternalPrimaryStorageEventView struct {
	Inventory ExternalPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// DownloadBackupFileFromPublicCloudEventView DownloadBackupFileFromPublicCloudEvent
type DownloadBackupFileFromPublicCloudEventView struct {
	Local string `json:"local,omitempty"`
}

// EjectZBoxEventView EjectZBoxEvent
type EjectZBoxEventView struct {
	Inventory ZBoxInventoryView `json:"inventory,omitempty"`
}

// EnableCbtTaskEventView EnableCbtTaskEvent
type EnableCbtTaskEventView struct {
	VolumeCbtBackupInfos []VolumeCbtBackupInfoView `json:"volumeCbtBackupInfos,omitempty"`
	Success bool `json:"success,omitempty"`
}

// EnableCdpTaskEventView EnableCdpTaskEvent
type EnableCdpTaskEventView struct {
	Inventory CdpTaskInventoryView `json:"inventory,omitempty"`
}

// ExecuteAutoScalingRuleEventView ExecuteAutoScalingRuleEvent
type ExecuteAutoScalingRuleEventView struct {
	ScalingActivityUuid string `json:"scalingActivityUuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ExecuteDRSSchedulingEventView ExecuteDRSSchedulingEvent
type ExecuteDRSSchedulingEventView struct {
	Success bool `json:"success,omitempty"`
}

// ExecuteGuestVmCommandEventView ExecuteGuestVmCommandEvent
type ExecuteGuestVmCommandEventView struct {
	Stream string `json:"stream,omitempty"`
	VmInstance VmInstanceInventoryView `json:"vmInstance,omitempty"`
}

// ExecuteGuestVmScriptEventView ExecuteGuestVmScriptEvent
type ExecuteGuestVmScriptEventView struct {
	Inventory GuestVmScriptExecutedRecordInventoryView `json:"inventory,omitempty"`
}

// ExportBuildAppEventView ExportBuildAppEvent
type ExportBuildAppEventView struct {
	Inventory BuildAppExportHistoryInventoryView `json:"inventory,omitempty"`
}

// ExportDatabaseBackupFromBackupStorageEventView ExportDatabaseBackupFromBackupStorageEvent
type ExportDatabaseBackupFromBackupStorageEventView struct {
	DatabaseBackupUrl string `json:"databaseBackupUrl,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ExportImageFromBackupStorageEventView ExportImageFromBackupStorageEvent
type ExportImageFromBackupStorageEventView struct {
	ImageUrl string `json:"imageUrl,omitempty"`
	ExportMd5Sum string `json:"exportMd5Sum,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ExportNbdVolumesEventView ExportNbdVolumesEvent
type ExportNbdVolumesEventView struct {
	VolumeInfos []VolumeCbtBackupInfoView `json:"volumeInfos,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ExportVmOvaPackageEventView ExportVmOvaPackageEvent
type ExportVmOvaPackageEventView struct {
	Inventory ImagePackageInventoryView `json:"inventory,omitempty"`
}

// FailoverFaultToleranceVmEventView FailoverFaultToleranceVmEvent
type FailoverFaultToleranceVmEventView struct {
	PrimaryVmInventory VmInstanceInventoryView `json:"primaryVmInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// FlattenVmInstanceEventView FlattenVmInstanceEvent
type FlattenVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// FlattenVolumeEventView FlattenVolumeEvent
type FlattenVolumeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// FstrimVmEventView FstrimVmEvent
type FstrimVmEventView struct {
	Success bool `json:"success,omitempty"`
}

// GCAliyunSnapshotRemoteEventView GCAliyunSnapshotRemoteEvent
type GCAliyunSnapshotRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// GenerateAccountBillingEventView GenerateAccountBillingEvent
type GenerateAccountBillingEventView struct {
	Success bool `json:"success,omitempty"`
}

// GenerateHygonMdevDevicesEventView GenerateHygonMdevDevicesEvent
type GenerateHygonMdevDevicesEventView struct {
	Success bool `json:"success,omitempty"`
}

// GenerateVirtualPciDevicesEventView GenerateVirtualPciDevicesEvent
type GenerateVirtualPciDevicesEventView struct {
	Success bool `json:"success,omitempty"`
}

// GenerateModelMetadataEventView GenerateModelMetadataEvent
type GenerateModelMetadataEventView struct {
	Success bool `json:"success,omitempty"`
}

// GenerateSeMdevDevicesEventView GenerateSeMdevDevicesEvent
type GenerateSeMdevDevicesEventView struct {
	Success bool `json:"success,omitempty"`
}

// GenerateSshKeyPairView GenerateSshKeyPair
type GenerateSshKeyPairView struct {
	Inventory SshPrivateKeyPairInventoryView `json:"inventory,omitempty"`
}

// IdentifyHostEventView IdentifyHostEvent
type IdentifyHostEventView struct {
	Success bool `json:"success,omitempty"`
}

// InspectBareMetal2ChassisByInstanceEventView InspectBareMetal2ChassisByInstanceEvent
type InspectBareMetal2ChassisByInstanceEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// InspectBareMetal2ChassisEventView InspectBareMetal2ChassisEvent
type InspectBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// InspectBaremetalChassisEventView InspectBaremetalChassisEvent
type InspectBaremetalChassisEventView struct {
	Inventory BaremetalChassisInventoryView `json:"inventory,omitempty"`
}

// IsLicenseServerView IsLicenseServer
type IsLicenseServerView struct {
	LicenseServer bool `json:"licenseServer,omitempty"`
	Success bool `json:"success,omitempty"`
}

// IsOpensourceVersionView IsOpensourceVersion
type IsOpensourceVersionView struct {
	Opensource bool `json:"opensource,omitempty"`
	Success bool `json:"success,omitempty"`
}

// IsReadyToGoView IsReadyToGo
type IsReadyToGoView struct {
	ManagementNodeId string `json:"managementNodeId,omitempty"`
	Success bool `json:"success,omitempty"`
}

// IsVfNicAvailableInL3NetworkView IsVfNicAvailableInL3Network
type IsVfNicAvailableInL3NetworkView struct {
	VfNicAvailable bool `json:"vfNicAvailable,omitempty"`
}

// KvmRunShellEventView KvmRunShellEvent
type KvmRunShellEventView struct {
	Inventory map[string]ShellResultView `json:"inventory,omitempty"`
}

// ListVMsFromKVMHostEventView ListVMsFromKVMHostEvent
type ListVMsFromKVMHostEventView struct {
	Inventories []VmInstanceInventoryView `json:"inventories,omitempty"`
	LibvirtVersion string `json:"libvirtVersion,omitempty"`
	QemuVersion string `json:"qemuVersion,omitempty"`
	V2vCaps map[string]bool `json:"v2vCaps,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ListVmSchedulingRulesFromExecuteStateView ListVmSchedulingRulesFromExecuteState
type ListVmSchedulingRulesFromExecuteStateView struct {
	Uuids []string `json:"uuids,omitempty"`
}

// ListVmsFromSchedulingStateView ListVmsFromSchedulingState
type ListVmsFromSchedulingStateView struct {
	Uuids []string `json:"uuids,omitempty"`
}

// LocalStorageGetVolumeMigratableView LocalStorageGetVolumeMigratable
type LocalStorageGetVolumeMigratableView struct {
	Inventories []HostInventoryView `json:"inventories,omitempty"`
}

// LocalStorageMigrateVolumeEventView LocalStorageMigrateVolumeEvent
type LocalStorageMigrateVolumeEventView struct {
	Inventory LocalStorageResourceRefInventoryView `json:"inventory,omitempty"`
}

// LocateHostNetworkInterfaceEventView LocateHostNetworkInterfaceEvent
type LocateHostNetworkInterfaceEventView struct {
	Success bool `json:"success,omitempty"`
}

// LocateLocalRaidPhysicalDriveEventView LocateLocalRaidPhysicalDriveEvent
type LocateLocalRaidPhysicalDriveEventView struct {
	Inventory RaidPhysicalDriveInventoryView `json:"inventory,omitempty"`
}

// LogInView LogIn
type LogInView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// LogInByLdapView LogInByLdap
type LogInByLdapView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
	AccountInventory AccountInventoryView `json:"accountInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// LogOutView LogOut
type LogOutView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// LoginByCasView LoginByCas
type LoginByCasView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// LoginIAM2PlatformView LoginIAM2Platform
type LoginIAM2PlatformView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// LoginIAM2ProjectView LoginIAM2Project
type LoginIAM2ProjectView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// LoginIAM2VirtualIDView LoginIAM2VirtualID
type LoginIAM2VirtualIDView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// LoginIAM2VirtualIDWithLdapView LoginIAM2VirtualIDWithLdap
type LoginIAM2VirtualIDWithLdapView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// MatchModelServiceTemplateWithModelEventView MatchModelServiceTemplateWithModelEvent
type MatchModelServiceTemplateWithModelEventView struct {
	Result map[string]interface{} `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}

// MergeDataOnBackupStorageEventView MergeDataOnBackupStorageEvent
type MergeDataOnBackupStorageEventView struct {
	Success bool `json:"success,omitempty"`
}

// MountVmInstanceRecoveryPointEventView MountVmInstanceRecoveryPointEvent
type MountVmInstanceRecoveryPointEventView struct {
	ResourcePath string `json:"resourcePath,omitempty"`
	FailedVolumes map[string]string `json:"failedVolumes,omitempty"`
	Success bool `json:"success,omitempty"`
}

// MoveDirectoryEventView MoveDirectoryEvent
type MoveDirectoryEventView struct {
	Success bool `json:"success,omitempty"`
}

// MoveResourcesToDirectoryEventView MoveResourcesToDirectoryEvent
type MoveResourcesToDirectoryEventView struct {
	Success bool `json:"success,omitempty"`
}

// ParseOvfView ParseOvf
type ParseOvfView struct {
	OvfInfo OvfInfoView `json:"ovfInfo,omitempty"`
}

// PauseVmInstanceEventView PauseVmInstanceEvent
type PauseVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// PowerOffBareMetal2ChassisEventView PowerOffBareMetal2ChassisEvent
type PowerOffBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// PowerOffBaremetalChassisEventView PowerOffBaremetalChassisEvent
type PowerOffBaremetalChassisEventView struct {
	Success bool `json:"success,omitempty"`
}

// PowerOffHostEventView PowerOffHostEvent
type PowerOffHostEventView struct {
	Results []PowerOffHardwareResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

// PowerOnBareMetal2ChassisEventView PowerOnBareMetal2ChassisEvent
type PowerOnBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// PowerOnBaremetalChassisEventView PowerOnBaremetalChassisEvent
type PowerOnBaremetalChassisEventView struct {
	Success bool `json:"success,omitempty"`
}

// PowerOnHostEventView PowerOnHostEvent
type PowerOnHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// PowerResetBareMetal2ChassisEventView PowerResetBareMetal2ChassisEvent
type PowerResetBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// PowerResetBaremetalChassisEventView PowerResetBaremetalChassisEvent
type PowerResetBaremetalChassisEventView struct {
	Success bool `json:"success,omitempty"`
}

// PowerResetHostEventView PowerResetHostEvent
type PowerResetHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// PreviewResourceStackView PreviewResourceStack
type PreviewResourceStackView struct {
	Preview PreviewResourceStructView `json:"preview,omitempty"`
}

// PrimaryStorageMigrateVmEventView PrimaryStorageMigrateVmEvent
type PrimaryStorageMigrateVmEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// PrimaryStorageMigrateVolumeEventView PrimaryStorageMigrateVolumeEvent
type PrimaryStorageMigrateVolumeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// PrometheusQueryLabelValuesView PrometheusQueryLabelValues
type PrometheusQueryLabelValuesView struct {
	Inventories interface{} `json:"inventories,omitempty"`
}

// PrometheusQueryMetadataView PrometheusQueryMetadata
type PrometheusQueryMetadataView struct {
	Inventories interface{} `json:"inventories,omitempty"`
}

// PrometheusQueryPassThroughView PrometheusQueryPassThrough
type PrometheusQueryPassThroughView struct {
	Inventories interface{} `json:"inventories,omitempty"`
}

// PrometheusQueryVmMonitoringDataView PrometheusQueryVmMonitoringData
type PrometheusQueryVmMonitoringDataView struct {
	Inventories interface{} `json:"inventories,omitempty"`
}

// ProtectVmInstanceRecoveryPointEventView ProtectVmInstanceRecoveryPointEvent
type ProtectVmInstanceRecoveryPointEventView struct {
	Success bool `json:"success,omitempty"`
}

// ProvisionNfvInstConfigEventView ProvisionNfvInstConfigEvent
type ProvisionNfvInstConfigEventView struct {
	Inventory ApplianceVmInventoryView `json:"inventory,omitempty"`
}

// ProvisionNfvInstGroupEventView ProvisionNfvInstGroupEvent
type ProvisionNfvInstGroupEventView struct {
	Inventory NfvInstGroupInventoryView `json:"inventory,omitempty"`
}

// ProvisionSlbGroupInstanceEventView ProvisionSlbGroupInstanceEvent
type ProvisionSlbGroupInstanceEventView struct {
	Inventory SlbGroupInventoryView `json:"inventory,omitempty"`
}

// ProvisionVirtualRouterConfigEventView ProvisionVirtualRouterConfigEvent
type ProvisionVirtualRouterConfigEventView struct {
	Inventory ApplianceVmInventoryView `json:"inventory,omitempty"`
}

// PublishAppEventView PublishAppEvent
type PublishAppEventView struct {
	Inventory PublishAppInventoryView `json:"inventory,omitempty"`
}

// PullHuaweiIMasterControllerEventView PullHuaweiIMasterControllerEvent
type PullHuaweiIMasterControllerEventView struct {
	Inventories []HuaweiIMasterSdnControllerInventoryView `json:"inventories,omitempty"`
}

// PullSdnControllerTenantEventView PullSdnControllerTenantEvent
type PullSdnControllerTenantEventView struct {
	Inventories []H3cSdnControllerTenantInventoryView `json:"inventories,omitempty"`
}

// PushLicenseAddOnsUsageEventView PushLicenseAddOnsUsageEvent
type PushLicenseAddOnsUsageEventView struct {
	Success bool `json:"success,omitempty"`
}

// PutMetricDataEventView PutMetricDataEvent
type PutMetricDataEventView struct {
	Success bool `json:"success,omitempty"`
}

// ReclaimSpaceFromImageStoreEventView ReclaimSpaceFromImageStoreEvent
type ReclaimSpaceFromImageStoreEventView struct {
	GcResult ImageStoreGcResultView `json:"gcResult,omitempty"`
}

// ReconnectAppBuildSystemEventView ReconnectAppBuildSystemEvent
type ReconnectAppBuildSystemEventView struct {
	Inventory AppBuildSystemInventoryView `json:"inventory,omitempty"`
}

// ReconnectBackupStorageEventView ReconnectBackupStorageEvent
type ReconnectBackupStorageEventView struct {
	Inventory BackupStorageInventoryView `json:"inventory,omitempty"`
}

// ReconnectBareMetal2GatewayEventView ReconnectBareMetal2GatewayEvent
type ReconnectBareMetal2GatewayEventView struct {
	Inventory BareMetal2GatewayInventoryView `json:"inventory,omitempty"`
}

// ReconnectBareMetal2InstanceEventView ReconnectBareMetal2InstanceEvent
type ReconnectBareMetal2InstanceEventView struct {
	Inventory BareMetal2InstanceInventoryView `json:"inventory,omitempty"`
}

// ReconnectBaremetalPxeServerEventView ReconnectBaremetalPxeServerEvent
type ReconnectBaremetalPxeServerEventView struct {
	Inventory BaremetalPxeServerInventoryView `json:"inventory,omitempty"`
}

// ReconnectConsoleProxyAgentEventView ReconnectConsoleProxyAgentEvent
type ReconnectConsoleProxyAgentEventView struct {
	Inventory map[string]interface{} `json:"inventory,omitempty"`
}

// ReconnectHostEventView ReconnectHostEvent
type ReconnectHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
}

// ReconnectIPsecConnectionEventView ReconnectIPsecConnectionEvent
type ReconnectIPsecConnectionEventView struct {
	Inventory IPsecConnectionInventoryView `json:"inventory,omitempty"`
}

// ReconnectImageStoreBackupStorageEventView ReconnectImageStoreBackupStorageEvent
type ReconnectImageStoreBackupStorageEventView struct {
	Inventory ImageStoreBackupStorageInventoryView `json:"inventory,omitempty"`
}

// ReconnectNfvInstEventView ReconnectNfvInstEvent
type ReconnectNfvInstEventView struct {
	Inventory ApplianceVmInventoryView `json:"inventory,omitempty"`
}

// ReconnectPrimaryStorageEventView ReconnectPrimaryStorageEvent
type ReconnectPrimaryStorageEventView struct {
	Inventory PrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// ReconnectSdnControllerEventView ReconnectSdnControllerEvent
type ReconnectSdnControllerEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// ReconnectSftpBackupStorageEventView ReconnectSftpBackupStorageEvent
type ReconnectSftpBackupStorageEventView struct {
	Inventory SftpBackupStorageInventoryView `json:"inventory,omitempty"`
}

// ReconnectVirtualRouterEventView ReconnectVirtualRouterEvent
type ReconnectVirtualRouterEventView struct {
	Inventory ApplianceVmInventoryView `json:"inventory,omitempty"`
}

// ReconnectZdfsEventView ReconnectZdfsEvent
type ReconnectZdfsEventView struct {
	Inventory ZdfsInventoryView `json:"inventory,omitempty"`
}

// RefreshCaptchaView RefreshCaptcha
type RefreshCaptchaView struct {
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	Captcha string `json:"captcha,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RefreshFiberChannelStorageEventView RefreshFiberChannelStorageEvent
type RefreshFiberChannelStorageEventView struct {
	Inventories []FiberChannelStorageInventoryView `json:"inventories,omitempty"`
}

// RefreshFirewallEventView RefreshFirewallEvent
type RefreshFirewallEventView struct {
	Inventory VpcFirewallInventoryView `json:"inventory,omitempty"`
}

// RefreshGuestOsMetadataEventView RefreshGuestOsMetadataEvent
type RefreshGuestOsMetadataEventView struct {
	Success bool `json:"success,omitempty"`
}

// RefreshIscsiServerEventView RefreshIscsiServerEvent
type RefreshIscsiServerEventView struct {
	Inventory IscsiServerInventoryView `json:"inventory,omitempty"`
}

// RefreshLoadBalancerEventView RefreshLoadBalancerEvent
type RefreshLoadBalancerEventView struct {
	Inventory LoadBalancerInventoryView `json:"inventory,omitempty"`
}

// RefreshLocalRaidEventView RefreshLocalRaidEvent
type RefreshLocalRaidEventView struct {
	Inventories []RaidControllerInventoryView `json:"inventories,omitempty"`
}

// RefreshNvmeTargetEventView RefreshNvmeTargetEvent
type RefreshNvmeTargetEventView struct {
	Inventories []NvmeTargetInventoryView `json:"inventories,omitempty"`
}

// RefreshPluginDriversEventView RefreshPluginDriversEvent
type RefreshPluginDriversEventView struct {
	Success bool `json:"success,omitempty"`
}

// RefreshSSOServerTokenEventView RefreshSSOServerTokenEvent
type RefreshSSOServerTokenEventView struct {
	Inventory SSOServerTokenInventoryView `json:"inventory,omitempty"`
}

// RefreshSearchIndexesView RefreshSearchIndexes
type RefreshSearchIndexesView struct {
	Success bool `json:"success,omitempty"`
}

// RefreshSharedBlockDeviceCapacityEventView RefreshSharedBlockDeviceCapacityEvent
type RefreshSharedBlockDeviceCapacityEventView struct {
	Inventory SharedBlockGroupPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// RegisterLicenseRequestedApplicationEventView RegisterLicenseRequestedApplicationEvent
type RegisterLicenseRequestedApplicationEventView struct {
	AppId string `json:"appId,omitempty"`
	ServicePubKey string `json:"servicePubKey,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RegisterLicenseServerEventView RegisterLicenseServerEvent
type RegisterLicenseServerEventView struct {
	LicenseClient LicenseAuthorizedNodeInventoryView `json:"licenseClient,omitempty"`
	LicenseServer LicenseAuthorizedNodeInventoryView `json:"licenseServer,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ReimageVmInstanceEventView ReimageVmInstanceEvent
type ReimageVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// ReloadElaborationEventView ReloadElaborationEvent
type ReloadElaborationEventView struct {
	Success bool `json:"success,omitempty"`
}

// ReloadExternalServiceEventView ReloadExternalServiceEvent
type ReloadExternalServiceEventView struct {
	Success bool `json:"success,omitempty"`
}

// ReloadLicenseView ReloadLicense
type ReloadLicenseView struct {
	Inventory LicenseInventoryView `json:"inventory,omitempty"`
}

// RenewSessionEventView RenewSessionEvent
type RenewSessionEventView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
}

// RequestConsoleAccessEventView RequestConsoleAccessEvent
type RequestConsoleAccessEventView struct {
	Inventory ConsoleInventoryView `json:"inventory,omitempty"`
}

// RequestLicenseCapacityEventView RequestLicenseCapacityEvent
type RequestLicenseCapacityEventView struct {
	Inventory LicenseAuthorizedCapacityInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// RerunLongJobEventView RerunLongJobEvent
type RerunLongJobEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// ResetGlobalConfigEventView ResetGlobalConfigEvent
type ResetGlobalConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// ResetTemplateConfigEventView ResetTemplateConfigEvent
type ResetTemplateConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// ResetTwoFactorAuthenticationSecretEventView ResetTwoFactorAuthenticationSecretEvent
type ResetTwoFactorAuthenticationSecretEventView struct {
	Inventory TwoFactorAuthenticationSecretInventoryView `json:"inventory,omitempty"`
}

// RestartModelServiceGroupsEventView RestartModelServiceGroupsEvent
type RestartModelServiceGroupsEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
}

// RestartResourceStackEventView RestartResourceStackEvent
type RestartResourceStackEventView struct {
	Inventory ResourceStackInventoryView `json:"inventory,omitempty"`
}

// ResumeLongJobEventView ResumeLongJobEvent
type ResumeLongJobEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// ResumeVmInstanceEventView ResumeVmInstanceEvent
type ResumeVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// RevertTemplateConfigEventView RevertTemplateConfigEvent
type RevertTemplateConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// RevertVmFromCdpBackupEventView RevertVmFromCdpBackupEvent
type RevertVmFromCdpBackupEventView struct {
	Success bool `json:"success,omitempty"`
}

// RevertVmFromSnapshotGroupEventView RevertVmFromSnapshotGroupEvent
type RevertVmFromSnapshotGroupEventView struct {
	Results []RevertSnapshotGroupResultView `json:"results,omitempty"`
}

// RevertVmFromVmBackupEventView RevertVmFromVmBackupEvent
type RevertVmFromVmBackupEventView struct {
	Success bool `json:"success,omitempty"`
}

// RevertVolumeFromSnapshotEventView RevertVolumeFromSnapshotEvent
type RevertVolumeFromSnapshotEventView struct {
	Success bool `json:"success,omitempty"`
}

// RevertVolumeFromVolumeBackupEventView RevertVolumeFromVolumeBackupEvent
type RevertVolumeFromVolumeBackupEventView struct {
	Success bool `json:"success,omitempty"`
}

// RevokeMonitorTemplateFromMonitorGroupEventView RevokeMonitorTemplateFromMonitorGroupEvent
type RevokeMonitorTemplateFromMonitorGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// RevokeResourceSharingEventView RevokeResourceSharingEvent
type RevokeResourceSharingEventView struct {
	Success bool `json:"success,omitempty"`
}

// RunIAM2ScriptEventView RunIAM2ScriptEvent
type RunIAM2ScriptEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// RunSchedulerTriggerEventView RunSchedulerTriggerEvent
type RunSchedulerTriggerEventView struct {
	Success bool `json:"success,omitempty"`
}

// SNSDingTalkTestConnectionEventView SNSDingTalkTestConnectionEvent
type SNSDingTalkTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// SNSEmailTestConnectionEventView SNSEmailTestConnectionEvent
type SNSEmailTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// SNSFeiShuTestConnectionEventView SNSFeiShuTestConnectionEvent
type SNSFeiShuTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// SNSHttpTestConnectionEventView SNSHttpTestConnectionEvent
type SNSHttpTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp string `json:"webhookResp,omitempty"`
}

// SNSMicrosoftTeamsTestConnectionEventView SNSMicrosoftTeamsTestConnectionEvent
type SNSMicrosoftTeamsTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// SNSSnmpTestConnectionEventView SNSSnmpTestConnectionEvent
type SNSSnmpTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// SNSWeComTestConnectionEventView SNSWeComTestConnectionEvent
type SNSWeComTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

// SdnControllerAddHostEventView SdnControllerAddHostEvent
type SdnControllerAddHostEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// SdnControllerChangeHostEventView SdnControllerChangeHostEvent
type SdnControllerChangeHostEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// SdnControllerRemoveHostEventView SdnControllerRemoveHostEvent
type SdnControllerRemoveHostEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// SecurityMachineDetectSyncEventView SecurityMachineDetectSyncEvent
type SecurityMachineDetectSyncEventView struct {
	Success bool `json:"success,omitempty"`
}

// SecurityMachineEncryptEventView SecurityMachineEncryptEvent
type SecurityMachineEncryptEventView struct {
	Text string `json:"text,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SelfTestLocalRaidEventView SelfTestLocalRaidEvent
type SelfTestLocalRaidEventView struct {
	Result string `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetFlowMeterRouterIdEventView SetFlowMeterRouterIdEvent
type SetFlowMeterRouterIdEventView struct {
	RouterId int64 `json:"routerId,omitempty"`
}

// SetIAM2ProjectContainerClusterEventView SetIAM2ProjectContainerClusterEvent
type SetIAM2ProjectContainerClusterEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetIAM2ProjectLoginExpiredEventView SetIAM2ProjectLoginExpiredEvent
type SetIAM2ProjectLoginExpiredEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetIAM2ProjectRetirePolicyEventView SetIAM2ProjectRetirePolicyEvent
type SetIAM2ProjectRetirePolicyEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetImageBootModeEventView SetImageBootModeEvent
type SetImageBootModeEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetImageQgaEventView SetImageQgaEvent
type SetImageQgaEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetImageSecurityLevelEventView SetImageSecurityLevelEvent
type SetImageSecurityLevelEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetImageStoreBackupStorageQuotaEventView SetImageStoreBackupStorageQuotaEvent
type SetImageStoreBackupStorageQuotaEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetIpOnHostNetworkBondingEventView SetIpOnHostNetworkBondingEvent
type SetIpOnHostNetworkBondingEventView struct {
	Inventory HostNetworkBondingInventoryView `json:"inventory,omitempty"`
}

// SetIpOnHostNetworkInterfaceEventView SetIpOnHostNetworkInterfaceEvent
type SetIpOnHostNetworkInterfaceEventView struct {
	Inventory HostNetworkInterfaceInventoryView `json:"inventory,omitempty"`
}

// SetL3NetworkMtuEventView SetL3NetworkMtuEvent
type SetL3NetworkMtuEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetL3NetworkRouterInterfaceIpEventView SetL3NetworkRouterInterfaceIpEvent
type SetL3NetworkRouterInterfaceIpEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetNicQosEventView SetNicQosEvent
type SetNicQosEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetOrganizationOperationEventView SetOrganizationOperationEvent
type SetOrganizationOperationEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetOrganizationSupervisorEventView SetOrganizationSupervisorEvent
type SetOrganizationSupervisorEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetSecurityMachineKeyEventView SetSecurityMachineKeyEvent
type SetSecurityMachineKeyEventView struct {
	Inventories []SecurityMachineInventoryView `json:"inventories,omitempty"`
}

// SetServiceTypeOnHostNetworkBondingEventView SetServiceTypeOnHostNetworkBondingEvent
type SetServiceTypeOnHostNetworkBondingEventView struct {
	Inventory []HostNetworkBondingServiceRefInventoryView `json:"inventory,omitempty"`
}

// SetServiceTypeOnHostNetworkInterfaceEventView SetServiceTypeOnHostNetworkInterfaceEvent
type SetServiceTypeOnHostNetworkInterfaceEventView struct {
	Inventory []HostNetworkInterfaceServiceRefInventoryView `json:"inventory,omitempty"`
}

// SetVRouterRouterIdEventView SetVRouterRouterIdEvent
type SetVRouterRouterIdEventView struct {
	RouterId string `json:"routerId,omitempty"`
}

// SetVipQosEventView SetVipQosEvent
type SetVipQosEventView struct {
	Inventory VipQosInventoryView `json:"inventory,omitempty"`
}

// SetVmBootModeEventView SetVmBootModeEvent
type SetVmBootModeEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmBootOrderEventView SetVmBootOrderEvent
type SetVmBootOrderEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmBootVolumeEventView SetVmBootVolumeEvent
type SetVmBootVolumeEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmCleanTrafficEventView SetVmCleanTrafficEvent
type SetVmCleanTrafficEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmClockTrackEventView SetVmClockTrackEvent
type SetVmClockTrackEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmConsoleModeEventView SetVmConsoleModeEvent
type SetVmConsoleModeEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmConsolePasswordEventView SetVmConsolePasswordEvent
type SetVmConsolePasswordEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmEmulatorPinningEventView SetVmEmulatorPinningEvent
type SetVmEmulatorPinningEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmHostnameEventView SetVmHostnameEvent
type SetVmHostnameEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmInstanceDefaultCdRomEventView SetVmInstanceDefaultCdRomEvent
type SetVmInstanceDefaultCdRomEventView struct {
	Inventory VmCdRomInventoryView `json:"inventory,omitempty"`
}

// SetVmInstanceHaLevelEventView SetVmInstanceHaLevelEvent
type SetVmInstanceHaLevelEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmInstanceHygonMdevEventView SetVmInstanceHygonMdevEvent
type SetVmInstanceHygonMdevEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmMonitorNumberEventView SetVmMonitorNumberEvent
type SetVmMonitorNumberEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmNicSecurityGroupEventView SetVmNicSecurityGroupEvent
type SetVmNicSecurityGroupEventView struct {
	Inventory []VmNicSecurityGroupRefInventoryView `json:"inventory,omitempty"`
}

// SetVmNumaEventView SetVmNumaEvent
type SetVmNumaEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmQgaEventView SetVmQgaEvent
type SetVmQgaEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmQxlMemoryEventView SetVmQxlMemoryEvent
type SetVmQxlMemoryEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmRDPEventView SetVmRDPEvent
type SetVmRDPEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmSecurityLevelEventView SetVmSecurityLevelEvent
type SetVmSecurityLevelEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmSoundTypeEventView SetVmSoundTypeEvent
type SetVmSoundTypeEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmSshKeyEventView SetVmSshKeyEvent
type SetVmSshKeyEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmStaticIpEventView SetVmStaticIpEvent
type SetVmStaticIpEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmUsbRedirectEventView SetVmUsbRedirectEvent
type SetVmUsbRedirectEventView struct {
	Success bool `json:"success,omitempty"`
}

// SetVmUserDefinedXmlHookScriptEventView SetVmUserDefinedXmlHookScriptEvent
type SetVmUserDefinedXmlHookScriptEventView struct {
	VmUserDefinedXmlHookScript string `json:"vmUserDefinedXmlHookScript,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetVmUserDefinedXmlEventView SetVmUserDefinedXmlEvent
type SetVmUserDefinedXmlEventView struct {
	VmUserDefinedXml string `json:"vmUserDefinedXml,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetVolumeIoThreadPinEventView SetVolumeIoThreadPinEvent
type SetVolumeIoThreadPinEventView struct {
	VolumeUuid string `json:"volumeUuid,omitempty"`
	IoThreadId int `json:"ioThreadId,omitempty"`
	Pin string `json:"pin,omitempty"`
}

// SetVolumeQosEventView SetVolumeQosEvent
type SetVolumeQosEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// SetVpcVRouterDistributedRoutingEnabledEventView SetVpcVRouterDistributedRoutingEnabledEvent
type SetVpcVRouterDistributedRoutingEnabledEventView struct {
	Enabled bool `json:"enabled,omitempty"`
}

// SetVpcVRouterNetworkServiceStateEventView SetVpcVRouterNetworkServiceStateEvent
type SetVpcVRouterNetworkServiceStateEventView struct {
	State string `json:"state,omitempty"`
}

// ShareResourceEventView ShareResourceEvent
type ShareResourceEventView struct {
	Success bool `json:"success,omitempty"`
}

// ShrinkVolumeSnapshotEventView ShrinkVolumeSnapshotEvent
type ShrinkVolumeSnapshotEventView struct {
	ShrinkResult ShrinkResultView `json:"shrinkResult,omitempty"`
}

// ShutdownHostEventView ShutdownHostEvent
type ShutdownHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SsoClientPushDataEventView SsoClientPushDataEvent
type SsoClientPushDataEventView struct {
	Success bool `json:"success,omitempty"`
}

// SubmitLongJobEventView SubmitLongJobEvent
type SubmitLongJobEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// SubscribeEventEventView SubscribeEventEvent
type SubscribeEventEventView struct {
	Inventory EventSubscriptionInventoryView `json:"inventory,omitempty"`
}

// SubscribeSNSTopicEventView SubscribeSNSTopicEvent
type SubscribeSNSTopicEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncAINginxConfigurationView SyncAINginxConfiguration
type SyncAINginxConfigurationView struct {
	UnSyncedRules []NginxRedirectRuleView `json:"unSyncedRules,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SyncAliyunRouteEntryFromRemoteEventView SyncAliyunRouteEntryFromRemoteEvent
type SyncAliyunRouteEntryFromRemoteEventView struct {
	Inventories []VpcVirtualRouteEntryInventoryView `json:"inventories,omitempty"`
}

// SyncAliyunRouterInterfaceFromRemoteEventView SyncAliyunRouterInterfaceFromRemoteEvent
type SyncAliyunRouterInterfaceFromRemoteEventView struct {
	Inventories []AliyunRouterInterfaceInventoryView `json:"inventories,omitempty"`
}

// SyncAliyunSnapshotRemoteEventView SyncAliyunSnapshotRemoteEvent
type SyncAliyunSnapshotRemoteEventView struct {
	Inventories []AliyunSnapshotInventoryView `json:"inventories,omitempty"`
}

// SyncAliyunVirtualRouterFromRemoteEventView SyncAliyunVirtualRouterFromRemoteEvent
type SyncAliyunVirtualRouterFromRemoteEventView struct {
	Inventories []VpcVirtualRouterInventoryView `json:"inventories,omitempty"`
}

// SyncBackupFromImageStoreBackupStorageEventView SyncBackupFromImageStoreBackupStorageEvent
type SyncBackupFromImageStoreBackupStorageEventView struct {
	Inventory VolumeBackupInventoryView `json:"inventory,omitempty"`
}

// SyncChronyServersEventView SyncChronyServersEvent
type SyncChronyServersEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncConnectionAccessPointFromRemoteEventView SyncConnectionAccessPointFromRemoteEvent
type SyncConnectionAccessPointFromRemoteEventView struct {
	Inventories []ConnectionAccessPointInventoryView `json:"inventories,omitempty"`
}

// SyncContainerManagementEndpointEventView SyncContainerManagementEndpointEvent
type SyncContainerManagementEndpointEventView struct {
	Inventory ContainerManagementEndpointInventoryView `json:"inventory,omitempty"`
}

// SyncDataCenterFromRemoteEventView SyncDataCenterFromRemoteEvent
type SyncDataCenterFromRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncDatabaseBackupFromImageStoreBackupStorageEventView SyncDatabaseBackupFromImageStoreBackupStorageEvent
type SyncDatabaseBackupFromImageStoreBackupStorageEventView struct {
	Inventory DatabaseBackupInventoryView `json:"inventory,omitempty"`
}

// SyncDatabaseBackupEventView SyncDatabaseBackupEvent
type SyncDatabaseBackupEventView struct {
	Result SyncBackupResultView `json:"result,omitempty"`
}

// SyncDiskFromAliyunFromRemoteEventView SyncDiskFromAliyunFromRemoteEvent
type SyncDiskFromAliyunFromRemoteEventView struct {
	Inventories []AliyunDiskInventoryView `json:"inventories,omitempty"`
}

// SyncEcsImageFromRemoteEventView SyncEcsImageFromRemoteEvent
type SyncEcsImageFromRemoteEventView struct {
	Inventories []EcsImageInventoryView `json:"inventories,omitempty"`
}

// SyncEcsInstanceFromRemoteEventView SyncEcsInstanceFromRemoteEvent
type SyncEcsInstanceFromRemoteEventView struct {
	Inventories []EcsInstanceInventoryView `json:"inventories,omitempty"`
}

// SyncEcsSecurityGroupFromRemoteEventView SyncEcsSecurityGroupFromRemoteEvent
type SyncEcsSecurityGroupFromRemoteEventView struct {
	Inventories []EcsSecurityGroupInventoryView `json:"inventories,omitempty"`
}

// SyncEcsSecurityGroupRuleFromRemoteEventView SyncEcsSecurityGroupRuleFromRemoteEvent
type SyncEcsSecurityGroupRuleFromRemoteEventView struct {
	Inventories []EcsSecurityGroupRuleInventoryView `json:"inventories,omitempty"`
}

// SyncEcsVSwitchFromRemoteEventView SyncEcsVSwitchFromRemoteEvent
type SyncEcsVSwitchFromRemoteEventView struct {
	Inventories []EcsVSwitchInventoryView `json:"inventories,omitempty"`
}

// SyncEcsVpcFromRemoteEventView SyncEcsVpcFromRemoteEvent
type SyncEcsVpcFromRemoteEventView struct {
	Inventories []EcsVpcInventoryView `json:"inventories,omitempty"`
}

// SyncHybridEipFromRemoteEventView SyncHybridEipFromRemoteEvent
type SyncHybridEipFromRemoteEventView struct {
	Inventories []HybridEipAddressInventoryView `json:"inventories,omitempty"`
}

// SyncIdentityFromRemoteEventView SyncIdentityFromRemoteEvent
type SyncIdentityFromRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncImageFromImageStoreBackupStorageEventView SyncImageFromImageStoreBackupStorageEvent
type SyncImageFromImageStoreBackupStorageEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// SyncImageEventView SyncImageEvent
type SyncImageEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncImageSizeEventView SyncImageSizeEvent
type SyncImageSizeEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
}

// SyncLdapServerEventView SyncLdapServerEvent
type SyncLdapServerEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// SyncLicenseCapacityEventView SyncLicenseCapacityEvent
type SyncLicenseCapacityEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncNfvInstGroupEventView SyncNfvInstGroupEvent
type SyncNfvInstGroupEventView struct {
	Inventory NfvInstGroupInventoryView `json:"inventory,omitempty"`
}

// SyncPrimaryStorageCapacityEventView SyncPrimaryStorageCapacityEvent
type SyncPrimaryStorageCapacityEventView struct {
	Inventory PrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// SyncVCenterEventView SyncVCenterEvent
type SyncVCenterEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncVirtualBorderRouterFromRemoteEventView SyncVirtualBorderRouterFromRemoteEvent
type SyncVirtualBorderRouterFromRemoteEventView struct {
	Inventories []VirtualBorderRouterInventoryView `json:"inventories,omitempty"`
}

// SyncVmBackupFromImageStoreBackupStorageEventView SyncVmBackupFromImageStoreBackupStorageEvent
type SyncVmBackupFromImageStoreBackupStorageEventView struct {
	Inventories []VolumeBackupInventoryView `json:"inventories,omitempty"`
}

// SyncVmBackupEventView SyncVmBackupEvent
type SyncVmBackupEventView struct {
	Result SyncBackupResultView `json:"result,omitempty"`
}

// SyncVmClockEventView SyncVmClockEvent
type SyncVmClockEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncVolumeBackupEventView SyncVolumeBackupEvent
type SyncVolumeBackupEventView struct {
	Result SyncBackupResultView `json:"result,omitempty"`
}

// SyncVolumeSizeEventView SyncVolumeSizeEvent
type SyncVolumeSizeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// SyncVpcUserVpnGatewayFromRemoteEventView SyncVpcUserVpnGatewayFromRemoteEvent
type SyncVpcUserVpnGatewayFromRemoteEventView struct {
	Inventories []VpcUserVpnGatewayInventoryView `json:"inventories,omitempty"`
}

// SyncVpcVpnConnectionFromRemoteEventView SyncVpcVpnConnectionFromRemoteEvent
type SyncVpcVpnConnectionFromRemoteEventView struct {
	Inventories []VpcVpnConnectionInventoryView `json:"inventories,omitempty"`
}

// SyncVpcVpnGatewayFromRemoteEventView SyncVpcVpnGatewayFromRemoteEvent
type SyncVpcVpnGatewayFromRemoteEventView struct {
	Inventories []VpcVpnGatewayInventoryView `json:"inventories,omitempty"`
}

// SyncZBoxCapacityEventView SyncZBoxCapacityEvent
type SyncZBoxCapacityEventView struct {
	Inventory ZBoxInventoryView `json:"inventory,omitempty"`
}

// TakeVmConsoleScreenshotEventView TakeVmConsoleScreenshotEvent
type TakeVmConsoleScreenshotEventView struct {
	ImageData string `json:"imageData,omitempty"`
}

// TerminateVirtualBorderRouterRemoteEventView TerminateVirtualBorderRouterRemoteEvent
type TerminateVirtualBorderRouterRemoteEventView struct {
	Success bool `json:"success,omitempty"`
}

// TokenIntrospectionView TokenIntrospection
type TokenIntrospectionView struct {
	Active bool `json:"active,omitempty"`
	Success bool `json:"success,omitempty"`
}

// TriggerGCJobEventView TriggerGCJobEvent
type TriggerGCJobEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnbindModelFromServiceEventView UnbindModelFromServiceEvent
type UnbindModelFromServiceEventView struct {
	Inventory ModelServiceInventoryView `json:"inventory,omitempty"`
}

// UndoSnapshotCreationEventView UndoSnapshotCreationEvent
type UndoSnapshotCreationEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// UnexportNbdVolumesEventView UnexportNbdVolumesEvent
type UnexportNbdVolumesEventView struct {
	Success bool `json:"success,omitempty"`
}

// UngenerateHygonMdevDevicesEventView UngenerateHygonMdevDevicesEvent
type UngenerateHygonMdevDevicesEventView struct {
	Success bool `json:"success,omitempty"`
}

// UngenerateVirtualPciDevicesEventView UngenerateVirtualPciDevicesEvent
type UngenerateVirtualPciDevicesEventView struct {
	Success bool `json:"success,omitempty"`
}

// UngenerateSeMdevDevicesEventView UngenerateSeMdevDevicesEvent
type UngenerateSeMdevDevicesEventView struct {
	Success bool `json:"success,omitempty"`
}

// UngroupVolumeSnapshotGroupEventView UngroupVolumeSnapshotGroupEvent
type UngroupVolumeSnapshotGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnlockIdentityView UnlockIdentity
type UnlockIdentityView struct {
	Success bool `json:"success,omitempty"`
}

// UnmountVmInstanceRecoveryPointEventView UnmountVmInstanceRecoveryPointEvent
type UnmountVmInstanceRecoveryPointEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnprotectVmInstanceRecoveryPointEventView UnprotectVmInstanceRecoveryPointEvent
type UnprotectVmInstanceRecoveryPointEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnregisterLicenseRequestedApplicationEventView UnregisterLicenseRequestedApplicationEvent
type UnregisterLicenseRequestedApplicationEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnregisterLicenseServerEventView UnregisterLicenseServerEvent
type UnregisterLicenseServerEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnsubscribeEventEventView UnsubscribeEventEvent
type UnsubscribeEventEventView struct {
	Success bool `json:"success,omitempty"`
}

// UnsubscribeSNSTopicEventView UnsubscribeSNSTopicEvent
type UnsubscribeSNSTopicEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpgradeBackupStorageCdpTasksEventView UpgradeBackupStorageCdpTasksEvent
type UpgradeBackupStorageCdpTasksEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpgradeToLicenseServerEventView UpgradeToLicenseServerEvent
type UpgradeToLicenseServerEventView struct {
	Inventory LicenseAuthorizedNodeInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UploadFileToVmEventView UploadFileToVmEvent
type UploadFileToVmEventView struct {
	Success bool `json:"success,omitempty"`
}

// ValidateClusterSupportDRSView ValidateClusterSupportDRS
type ValidateClusterSupportDRSView struct {
	Supported bool `json:"supported,omitempty"`
	Reason ErrorCodeView `json:"reason,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ValidateDiskOfferingUserConfigEventView ValidateDiskOfferingUserConfigEvent
type ValidateDiskOfferingUserConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// ValidateInstanceOfferingUserConfigEventView ValidateInstanceOfferingUserConfigEvent
type ValidateInstanceOfferingUserConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// ValidatePasswordView ValidatePassword
type ValidatePasswordView struct {
	Available bool `json:"available,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ValidatePriceUserConfigEventView ValidatePriceUserConfigEvent
type ValidatePriceUserConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// ValidateSNSAliyunSmsEndpointEventView ValidateSNSAliyunSmsEndpointEvent
type ValidateSNSAliyunSmsEndpointEventView struct {
	Success bool `json:"success,omitempty"`
}

// ValidateSNSEmailPlatformEventView ValidateSNSEmailPlatformEvent
type ValidateSNSEmailPlatformEventView struct {
	Success bool `json:"success,omitempty"`
}

// ValidateSNSApplicationEndpointEventView ValidateSNSApplicationEndpointEvent
type ValidateSNSApplicationEndpointEventView struct {
	Success bool `json:"success,omitempty"`
}

// ValidateSecurityGroupRuleView ValidateSecurityGroupRule
type ValidateSecurityGroupRuleView struct {
	Available bool `json:"available,omitempty"`
	Code string `json:"code,omitempty"`
	Reason string `json:"reason,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ValidateSessionView ValidateSession
type ValidateSessionView struct {
	Valid bool `json:"valid,omitempty"`
}

// ValidateVmSchedulingRuleView ValidateVmSchedulingRule
type ValidateVmSchedulingRuleView struct {
	Success bool `json:"success,omitempty"`
}

// ValidateVolumeSnapshotChainEventView ValidateVolumeSnapshotChainEvent
type ValidateVolumeSnapshotChainEventView struct {
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

// WithdrawLicenseCapacityApplicationEventView WithdrawLicenseCapacityApplicationEvent
type WithdrawLicenseCapacityApplicationEventView struct {
	Success bool `json:"success,omitempty"`
}

// ZQLQueryView ZQLQuery
type ZQLQueryView struct {
	Results []ZQLQueryReturnView `json:"results,omitempty"`
}

