// Copyright (c) ZStack.io, Inc.

package param

// GetVirtualRouterSoftwareVersionDetailParam GetVirtualRouterSoftwareVersion详细参数
type GetVirtualRouterSoftwareVersionDetailParam struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"softwareName" validate:"required"` // 必填
	rest bool `json:"needUpdate,omitempty"`
}

// GetVirtualRouterSoftwareVersionParam GetVirtualRouterSoftwareVersion请求参数
type GetVirtualRouterSoftwareVersionParam struct {
	BaseParam
	Params GetVirtualRouterSoftwareVersionDetailParam `json:"params"` // 详细参数
}

