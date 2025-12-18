// Copyright (c) ZStack.io, Inc.

package param

// DeletePolicyRouteRuleDetailParam DeletePolicyRouteRule detail param
type DeletePolicyRouteRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePolicyRouteRuleParam DeletePolicyRouteRule request param
type DeletePolicyRouteRuleParam struct {
	BaseParam
	Params DeletePolicyRouteRuleDetailParam `json:"params"`
}
