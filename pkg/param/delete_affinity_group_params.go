// Copyright (c) ZStack.io, Inc.

package param

// DeleteAffinityGroupDetailParam DeleteAffinityGroup detail param
type DeleteAffinityGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAffinityGroupParam DeleteAffinityGroup request param
type DeleteAffinityGroupParam struct {
	BaseParam
	Params DeleteAffinityGroupDetailParam `json:"params"`
}
