// Copyright (c) ZStack.io, Inc.

package param

// GetSupportedCloudFormationResourcesDetailParam GetSupportedCloudFormationResources详细参数
type GetSupportedCloudFormationResourcesDetailParam struct {
	rest string `json:"version,omitempty"`
	rest string `json:"type,omitempty"`
}

// GetSupportedCloudFormationResourcesParam GetSupportedCloudFormationResources请求参数
type GetSupportedCloudFormationResourcesParam struct {
	BaseParam
	Params GetSupportedCloudFormationResourcesDetailParam `json:"params"` // 详细参数
}

