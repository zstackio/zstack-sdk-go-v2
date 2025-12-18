// Copyright (c) ZStack.io, Inc.

package param

// DeleteVirtualRouterLocalDetailParam DeleteVirtualRouterLocal detail param
type DeleteVirtualRouterLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVirtualRouterLocalParam DeleteVirtualRouterLocal request param
type DeleteVirtualRouterLocalParam struct {
	BaseParam
	Params DeleteVirtualRouterLocalDetailParam `json:"params"`
}
