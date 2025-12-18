// Copyright (c) ZStack.io, Inc.

package param

// DeletePolicyRouteRuleSetDetailParam DeletePolicyRouteRuleSet detail param
type DeletePolicyRouteRuleSetDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePolicyRouteRuleSetParam DeletePolicyRouteRuleSet request param
type DeletePolicyRouteRuleSetParam struct {
	BaseParam
	Params DeletePolicyRouteRuleSetDetailParam `json:"params"`
}
