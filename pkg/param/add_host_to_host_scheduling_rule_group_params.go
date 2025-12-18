// Copyright (c) ZStack.io, Inc.

package param

// AddHostToHostSchedulingRuleGroupDetailParam AddHostToHostSchedulingRuleGroup详细参数
type AddHostToHostSchedulingRuleGroupDetailParam struct {
	rest string `json:"hostGroupUuid" validate:"required"` // 必填
	rest string `json:"hostUuid" validate:"required"` // 必填
}

// AddHostToHostSchedulingRuleGroupParam AddHostToHostSchedulingRuleGroup请求参数
type AddHostToHostSchedulingRuleGroupParam struct {
	BaseParam
	Params AddHostToHostSchedulingRuleGroupDetailParam `json:"params"` // 详细参数
}

