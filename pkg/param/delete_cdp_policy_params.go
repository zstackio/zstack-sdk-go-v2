// Copyright (c) ZStack.io, Inc.

package param

// DeleteCdpPolicyDetailParam DeleteCdpPolicy detail param
type DeleteCdpPolicyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteCdpPolicyParam DeleteCdpPolicy request param
type DeleteCdpPolicyParam struct {
	BaseParam
	Params DeleteCdpPolicyDetailParam `json:"params"`
}
