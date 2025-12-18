// Copyright (c) ZStack.io, Inc.

package param

// CreateFirewallRuleSetDetailParam CreateFirewallRuleSet detail param
type CreateFirewallRuleSetDetailParam struct {
	Name string `json:"name" validate:"required"`
	ActionType string `json:"actionType,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFirewallRuleSetParam CreateFirewallRuleSet request param
type CreateFirewallRuleSetParam struct {
	BaseParam
	Params CreateFirewallRuleSetDetailParam `json:"params"`
}
