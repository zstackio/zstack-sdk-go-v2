// Copyright (c) ZStack.io, Inc.

package param

// CreateClusterDetailParam CreateCluster detail param
type CreateClusterDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	HypervisorType string `json:"hypervisorType" validate:"required"`
	Type string `json:"type,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateClusterParam CreateCluster request param
type CreateClusterParam struct {
	BaseParam
	Params CreateClusterDetailParam `json:"params"`
}
