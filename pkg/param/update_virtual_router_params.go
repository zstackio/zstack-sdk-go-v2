// Copyright (c) ZStack.io, Inc.

package param

// UpdateVirtualRouterDetailParam UpdateVirtualRouter detail param
type UpdateVirtualRouterDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	DefaultRouteL3NetworkUuid string `json:"defaultRouteL3NetworkUuid,omitempty"`
}

// UpdateVirtualRouterParam UpdateVirtualRouter request param
type UpdateVirtualRouterParam struct {
	BaseParam
	Params UpdateVirtualRouterDetailParam `json:"params"`
}
