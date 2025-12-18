// Copyright (c) ZStack.io, Inc.

package param

// CreateSlbInstanceDetailParam CreateSlbInstance detail param
type CreateSlbInstanceDetailParam struct {
	Name string `json:"name" validate:"required"`
	SlbGroupUuid string `json:"slbGroupUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSlbInstanceParam CreateSlbInstance request param
type CreateSlbInstanceParam struct {
	BaseParam
	Params CreateSlbInstanceDetailParam `json:"params"`
}
