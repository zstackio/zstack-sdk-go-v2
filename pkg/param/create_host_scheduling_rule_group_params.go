// Copyright (c) ZStack.io, Inc.

package param

// CreateHostSchedulingRuleGroupDetailParam CreateHostSchedulingRuleGroup detail param
type CreateHostSchedulingRuleGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateHostSchedulingRuleGroupParam CreateHostSchedulingRuleGroup request param
type CreateHostSchedulingRuleGroupParam struct {
	BaseParam
	Params CreateHostSchedulingRuleGroupDetailParam `json:"params"`
}
