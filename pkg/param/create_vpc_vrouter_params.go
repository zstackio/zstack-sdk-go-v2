// Copyright (c) ZStack.io, Inc.

package param

// CreateVpcVRouterDetailParam CreateVpcVRouter detail param
type CreateVpcVRouterDetailParam struct {
	Name string `json:"name" validate:"required"`
	VirtualRouterOfferingUuid string `json:"virtualRouterOfferingUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcVRouterParam CreateVpcVRouter request param
type CreateVpcVRouterParam struct {
	BaseParam
	Params CreateVpcVRouterDetailParam `json:"params"`
}
