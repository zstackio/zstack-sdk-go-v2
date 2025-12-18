// Copyright (c) ZStack.io, Inc.

package param

// DeleteVpcHaGroupDetailParam DeleteVpcHaGroup detail param
type DeleteVpcHaGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVpcHaGroupParam DeleteVpcHaGroup request param
type DeleteVpcHaGroupParam struct {
	BaseParam
	Params DeleteVpcHaGroupDetailParam `json:"params"`
}
