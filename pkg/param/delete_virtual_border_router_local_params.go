// Copyright (c) ZStack.io, Inc.

package param

// DeleteVirtualBorderRouterLocalDetailParam DeleteVirtualBorderRouterLocal detail param
type DeleteVirtualBorderRouterLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVirtualBorderRouterLocalParam DeleteVirtualBorderRouterLocal request param
type DeleteVirtualBorderRouterLocalParam struct {
	BaseParam
	Params DeleteVirtualBorderRouterLocalDetailParam `json:"params"`
}
