// Copyright (c) ZStack.io, Inc.

package param

// GetSupportedCloudFormationResourcesDetailParam GetSupportedCloudFormationResources detail param
type GetSupportedCloudFormationResourcesDetailParam struct {
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
}

// GetSupportedCloudFormationResourcesParam GetSupportedCloudFormationResources request param
type GetSupportedCloudFormationResourcesParam struct {
	BaseParam
	Params GetSupportedCloudFormationResourcesDetailParam `json:"params"`
}
