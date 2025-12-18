// Copyright (c) ZStack.io, Inc.

package param

// UpdateVirtualRouterSoftwareVersionDetailParam UpdateVirtualRouterSoftwareVersion detail param
type UpdateVirtualRouterSoftwareVersionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SoftwareName string `json:"softwareName" validate:"required"`
	TargetVersion string `json:"targetVersion" validate:"required"`
}

// UpdateVirtualRouterSoftwareVersionParam UpdateVirtualRouterSoftwareVersion request param
type UpdateVirtualRouterSoftwareVersionParam struct {
	BaseParam
	Params UpdateVirtualRouterSoftwareVersionDetailParam `json:"params"`
}
