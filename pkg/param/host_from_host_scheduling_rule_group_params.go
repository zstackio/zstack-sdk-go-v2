// Copyright (c) ZStack.io, Inc.

package param

// DetachHostFromHostSchedulingRuleGroupDetailParam DetachHostFromHostSchedulingRuleGroup详细参数
type DetachHostFromHostSchedulingRuleGroupDetailParam struct {
	rest string `json:"hostGroupUuid" validate:"required"` // 必填
	rest string `json:"hostUuid" validate:"required"` // 必填
}

// DetachHostFromHostSchedulingRuleGroupParam DetachHostFromHostSchedulingRuleGroup请求参数
type DetachHostFromHostSchedulingRuleGroupParam struct {
	BaseParam
	Params DetachHostFromHostSchedulingRuleGroupDetailParam `json:"params"` // 详细参数
}

