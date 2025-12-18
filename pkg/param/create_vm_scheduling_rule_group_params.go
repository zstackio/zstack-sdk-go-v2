// Copyright (c) ZStack.io, Inc.

package param

// CreateVmSchedulingRuleGroupDetailParam CreateVmSchedulingRuleGroup detail param
type CreateVmSchedulingRuleGroupDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmSchedulingRuleGroupParam CreateVmSchedulingRuleGroup request param
type CreateVmSchedulingRuleGroupParam struct {
	BaseParam
	Params CreateVmSchedulingRuleGroupDetailParam `json:"params"`
}
