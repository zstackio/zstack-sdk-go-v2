// Copyright (c) ZStack.io, Inc.

package param

// GetVirtualRouterSoftwareVersionDetailParam GetVirtualRouterSoftwareVersion detail param
type GetVirtualRouterSoftwareVersionDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
	SoftwareName string `json:"softwareName" validate:"required"`
	NeedUpdate bool `json:"needUpdate,omitempty"`
}

// GetVirtualRouterSoftwareVersionParam GetVirtualRouterSoftwareVersion request param
type GetVirtualRouterSoftwareVersionParam struct {
	BaseParam
	Params GetVirtualRouterSoftwareVersionDetailParam `json:"params"`
}
