// Copyright (c) ZStack.io, Inc.

package param

// DetachHostFromHostSchedulingRuleGroupDetailParam DetachHostFromHostSchedulingRuleGroup detail param
type DetachHostFromHostSchedulingRuleGroupDetailParam struct {
	HostGroupUuid string `json:"hostGroupUuid" validate:"required"`
	HostUuid string `json:"hostUuid" validate:"required"`
}

// DetachHostFromHostSchedulingRuleGroupParam DetachHostFromHostSchedulingRuleGroup request param
type DetachHostFromHostSchedulingRuleGroupParam struct {
	BaseParam
	Params DetachHostFromHostSchedulingRuleGroupDetailParam `json:"params"`
}
