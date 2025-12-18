// Copyright (c) ZStack.io, Inc.

package param

// CreatePolicyRouteRuleSetDetailParam CreatePolicyRouteRuleSet detail param
type CreatePolicyRouteRuleSetDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePolicyRouteRuleSetParam CreatePolicyRouteRuleSet request param
type CreatePolicyRouteRuleSetParam struct {
	BaseParam
	Params CreatePolicyRouteRuleSetDetailParam `json:"params"`
}
