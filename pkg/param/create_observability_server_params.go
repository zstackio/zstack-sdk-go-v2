// Copyright (c) ZStack.io, Inc.

package param

// CreateObservabilityServerDetailParam CreateObservabilityServer detail param
type CreateObservabilityServerDetailParam struct {
	Name string `json:"name" validate:"required"`
	ObservabilityServerOfferingUuid string `json:"observabilityServerOfferingUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateObservabilityServerParam CreateObservabilityServer request param
type CreateObservabilityServerParam struct {
	BaseParam
	Params CreateObservabilityServerDetailParam `json:"params"`
}
