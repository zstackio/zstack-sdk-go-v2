// Copyright (c) ZStack.io, Inc.

package param

// DeleteMulticastRouterDetailParam DeleteMulticastRouter detail param
type DeleteMulticastRouterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMulticastRouterParam DeleteMulticastRouter request param
type DeleteMulticastRouterParam struct {
	BaseParam
	Params DeleteMulticastRouterDetailParam `json:"params"`
}
