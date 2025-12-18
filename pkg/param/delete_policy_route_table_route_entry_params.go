// Copyright (c) ZStack.io, Inc.

package param

// DeletePolicyRouteTableRouteEntryDetailParam DeletePolicyRouteTableRouteEntry detail param
type DeletePolicyRouteTableRouteEntryDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePolicyRouteTableRouteEntryParam DeletePolicyRouteTableRouteEntry request param
type DeletePolicyRouteTableRouteEntryParam struct {
	BaseParam
	Params DeletePolicyRouteTableRouteEntryDetailParam `json:"params"`
}
