// Copyright (c) ZStack.io, Inc.

package param

// DeletePolicyRouteTableDetailParam DeletePolicyRouteTable detail param
type DeletePolicyRouteTableDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePolicyRouteTableParam DeletePolicyRouteTable request param
type DeletePolicyRouteTableParam struct {
	BaseParam
	Params DeletePolicyRouteTableDetailParam `json:"params"`
}
