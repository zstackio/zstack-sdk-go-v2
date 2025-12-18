// Copyright (c) ZStack.io, Inc.

package param

// UpdateClusterOSDetailParam UpdateClusterOS详细参数
type UpdateClusterOSDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"hostUuid,omitempty"`
	rest []string `json:"excludePackages,omitempty"`
	rest []string `json:"updatePackages,omitempty"`
	rest string `json:"releaseVersion,omitempty"`
	rest bool `json:"force,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// UpdateClusterOSParam UpdateClusterOS请求参数
type UpdateClusterOSParam struct {
	BaseParam
	Params UpdateClusterOSDetailParam `json:"params"` // 详细参数
}

