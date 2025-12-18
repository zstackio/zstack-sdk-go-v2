// Copyright (c) ZStack.io, Inc.

package param

// CloneVmInstanceDetailParam CloneVmInstance detail param
type CloneVmInstanceDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Strategy string `json:"strategy,omitempty"`
	VmNicParams []interface{} `json:"vmNicParams,omitempty"`
	Names []string `json:"names" validate:"required"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume string `json:"primaryStorageUuidForDataVolume,omitempty"`
	Full bool `json:"full,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
}

// CloneVmInstanceParam CloneVmInstance request param
type CloneVmInstanceParam struct {
	BaseParam
	Params CloneVmInstanceDetailParam `json:"params"`
}
