// Copyright (c) ZStack.io, Inc.

package param

// DeleteVRouterOspfAreaDetailParam DeleteVRouterOspfArea detail param
type DeleteVRouterOspfAreaDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVRouterOspfAreaParam DeleteVRouterOspfArea request param
type DeleteVRouterOspfAreaParam struct {
	BaseParam
	Params DeleteVRouterOspfAreaDetailParam `json:"params"`
}
