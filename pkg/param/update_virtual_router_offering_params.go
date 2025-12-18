// Copyright (c) ZStack.io, Inc.

package param

// UpdateVirtualRouterOfferingDetailParam UpdateVirtualRouterOffering detail param
type UpdateVirtualRouterOfferingDetailParam struct {
	IsDefault bool `json:"isDefault,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
}

// UpdateVirtualRouterOfferingParam UpdateVirtualRouterOffering request param
type UpdateVirtualRouterOfferingParam struct {
	BaseParam
	Params UpdateVirtualRouterOfferingDetailParam `json:"params"`
}
