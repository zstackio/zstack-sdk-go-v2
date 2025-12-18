// Copyright (c) ZStack.io, Inc.

package param

// CreateVmFromVmBackupDetailParam CreateVmFromVmBackup detail param
type CreateVmFromVmBackupDetailParam struct {
	Name string `json:"name" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	VmNicParams string `json:"vmNicParams,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume string `json:"primaryStorageUuidForDataVolume,omitempty"`
	Description string `json:"description,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmFromVmBackupParam CreateVmFromVmBackup request param
type CreateVmFromVmBackupParam struct {
	BaseParam
	Params CreateVmFromVmBackupDetailParam `json:"params"`
}
