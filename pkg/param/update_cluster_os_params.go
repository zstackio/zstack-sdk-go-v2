// Copyright (c) ZStack.io, Inc.

package param

// UpdateClusterOSDetailParam UpdateClusterOS detail param
type UpdateClusterOSDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	HostUuid string `json:"hostUuid,omitempty"`
	ExcludePackages []string `json:"excludePackages,omitempty"`
	UpdatePackages []string `json:"updatePackages,omitempty"`
	ReleaseVersion string `json:"releaseVersion,omitempty"`
	Force bool `json:"force,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateClusterOSParam UpdateClusterOS request param
type UpdateClusterOSParam struct {
	BaseParam
	Params UpdateClusterOSDetailParam `json:"params"`
}
