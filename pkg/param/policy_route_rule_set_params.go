// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreatePolicyRouteRuleSetParamDetail CreatePolicyRouteRuleSet detail param
type CreatePolicyRouteRuleSetParamDetail struct {
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
	CreatePolicyRouteRuleSet CreatePolicyRouteRuleSetParamDetail `json:"createPolicyRouteRuleSet"`
}
// UpdatePolicyRouteRuleSetParamDetail UpdatePolicyRouteRuleSet detail param
type UpdatePolicyRouteRuleSetParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdatePolicyRouteRuleSetParam UpdatePolicyRouteRuleSet request param
type UpdatePolicyRouteRuleSetParam struct {
	BaseParam
	UpdatePolicyRouteRuleSet UpdatePolicyRouteRuleSetParamDetail `json:"updatePolicyRouteRuleSet"`
}
// DeletePolicyRouteRuleSetParamDetail DeletePolicyRouteRuleSet detail param
type DeletePolicyRouteRuleSetParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePolicyRouteRuleSetParam DeletePolicyRouteRuleSet request param
type DeletePolicyRouteRuleSetParam struct {
	BaseParam
	DeletePolicyRouteRuleSet DeletePolicyRouteRuleSetParamDetail `json:"deletePolicyRouteRuleSet"`
}
