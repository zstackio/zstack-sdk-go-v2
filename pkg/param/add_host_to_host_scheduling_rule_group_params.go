// Copyright (c) ZStack.io, Inc.

package param

// AddHostToHostSchedulingRuleGroupDetailParam AddHostToHostSchedulingRuleGroup detail param
type AddHostToHostSchedulingRuleGroupDetailParam struct {
	HostGroupUuid string `json:"hostGroupUuid" validate:"required"`
	HostUuid string `json:"hostUuid" validate:"required"`
}

// AddHostToHostSchedulingRuleGroupParam AddHostToHostSchedulingRuleGroup request param
type AddHostToHostSchedulingRuleGroupParam struct {
	BaseParam
	Params AddHostToHostSchedulingRuleGroupDetailParam `json:"params"`
}
