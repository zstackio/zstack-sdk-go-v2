// Copyright (c) ZStack.io, Inc.

package param

// CreatePolicyRouteTableDetailParam CreatePolicyRouteTable detail param
type CreatePolicyRouteTableDetailParam struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	Number int `json:"number" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePolicyRouteTableParam CreatePolicyRouteTable request param
type CreatePolicyRouteTableParam struct {
	BaseParam
	Params CreatePolicyRouteTableDetailParam `json:"params"`
}
