// Copyright (c) ZStack.io, Inc.

package param

// UpdatePolicyRouteRuleSetDetailParam UpdatePolicyRouteRuleSet detail param
type UpdatePolicyRouteRuleSetDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdatePolicyRouteRuleSetParam UpdatePolicyRouteRuleSet request param
type UpdatePolicyRouteRuleSetParam struct {
	BaseParam
	Params UpdatePolicyRouteRuleSetDetailParam `json:"params"`
}
