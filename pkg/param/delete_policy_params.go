// Copyright (c) ZStack.io, Inc.

package param

// DeletePolicyDetailParam DeletePolicy detail param
type DeletePolicyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePolicyParam DeletePolicy request param
type DeletePolicyParam struct {
	BaseParam
	Params DeletePolicyDetailParam `json:"params"`
}
