// Copyright (c) ZStack.io, Inc.

package param

// ReconnectVirtualRouterDetailParam ReconnectVirtualRouter detail param
type ReconnectVirtualRouterDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// ReconnectVirtualRouterParam ReconnectVirtualRouter request param
type ReconnectVirtualRouterParam struct {
	BaseParam
	Params ReconnectVirtualRouterDetailParam `json:"params"`
}
