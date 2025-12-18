// Copyright (c) ZStack.io, Inc.

package param

// CreateVmFromCdpBackupDetailParam CreateVmFromCdpBackup detail param
type CreateVmFromCdpBackupDetailParam struct {
	Name string `json:"name" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
	CdpTaskUuid string `json:"cdpTaskUuid" validate:"required"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid" validate:"required"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume string `json:"primaryStorageUuidForDataVolume,omitempty"`
	RecoverBandwidth int64 `json:"recoverBandwidth,omitempty"`
	Description string `json:"description,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmFromCdpBackupParam CreateVmFromCdpBackup request param
type CreateVmFromCdpBackupParam struct {
	BaseParam
	Params CreateVmFromCdpBackupDetailParam `json:"params"`
}
