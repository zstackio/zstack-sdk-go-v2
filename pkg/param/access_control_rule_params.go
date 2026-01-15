// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddAccessControlRuleParamDetail AddAccessControlRule detail param
type AddAccessControlRuleParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Rule string `json:"rule" validate:"required"`
	ControlStrategy string `json:"controlStrategy" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAccessControlRuleParam AddAccessControlRule request param
type AddAccessControlRuleParam struct {
	BaseParam
	Params AddAccessControlRuleParamDetail `json:"addAccessControlRule"`
}
// UpdateAccessControlRuleParamDetail UpdateAccessControlRule detail param
type UpdateAccessControlRuleParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Rule string `json:"rule,omitempty"`
}

// UpdateAccessControlRuleParam UpdateAccessControlRule request param
type UpdateAccessControlRuleParam struct {
	BaseParam
	Params UpdateAccessControlRuleParamDetail `json:"updateAccessControlRule"`
}
// DeleteAccessControlRuleParamDetail DeleteAccessControlRule detail param
type DeleteAccessControlRuleParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAccessControlRuleParam DeleteAccessControlRule request param
type DeleteAccessControlRuleParam struct {
	BaseParam
	Params DeleteAccessControlRuleParamDetail `json:"deleteAccessControlRule"`
}
