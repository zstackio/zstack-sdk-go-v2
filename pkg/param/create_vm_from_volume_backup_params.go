// Copyright (c) ZStack.io, Inc.

package param

// CreateVmFromVolumeBackupDetailParam CreateVmFromVolumeBackup detail param
type CreateVmFromVolumeBackupDetailParam struct {
	Name string `json:"name" validate:"required"`
	BackupUuid string `json:"backupUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid" validate:"required"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	Description string `json:"description,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmFromVolumeBackupParam CreateVmFromVolumeBackup request param
type CreateVmFromVolumeBackupParam struct {
	BaseParam
	Params CreateVmFromVolumeBackupDetailParam `json:"params"`
}
