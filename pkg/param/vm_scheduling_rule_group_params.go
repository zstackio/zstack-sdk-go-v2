// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteVmSchedulingRuleGroupParamDetail DeleteVmSchedulingRuleGroup detail param
type DeleteVmSchedulingRuleGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteVmSchedulingRuleGroupParam DeleteVmSchedulingRuleGroup request param
type DeleteVmSchedulingRuleGroupParam struct {
	BaseParam
	DeleteVmSchedulingRuleGroup DeleteVmSchedulingRuleGroupParamDetail `json:"deleteVmSchedulingRuleGroup"`
}
// UpdateVmSchedulingRuleGroupParamDetail UpdateVmSchedulingRuleGroup detail param
type UpdateVmSchedulingRuleGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVmSchedulingRuleGroupParam UpdateVmSchedulingRuleGroup request param
type UpdateVmSchedulingRuleGroupParam struct {
	BaseParam
	UpdateVmSchedulingRuleGroup UpdateVmSchedulingRuleGroupParamDetail `json:"updateVmSchedulingRuleGroup"`
}
// CreateVmSchedulingRuleGroupParamDetail CreateVmSchedulingRuleGroup detail param
type CreateVmSchedulingRuleGroupParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmSchedulingRuleGroupParam CreateVmSchedulingRuleGroup request param
type CreateVmSchedulingRuleGroupParam struct {
	BaseParam
	CreateVmSchedulingRuleGroup CreateVmSchedulingRuleGroupParamDetail `json:"createVmSchedulingRuleGroup"`
}
