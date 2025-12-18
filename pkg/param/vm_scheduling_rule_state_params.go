// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmSchedulingRuleStateDetailParam ChangeVmSchedulingRuleState详细参数
type ChangeVmSchedulingRuleStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"state" validate:"required"` // 必填
}

// ChangeVmSchedulingRuleStateParam ChangeVmSchedulingRuleState请求参数
type ChangeVmSchedulingRuleStateParam struct {
	BaseParam
	Params ChangeVmSchedulingRuleStateDetailParam `json:"params"` // 详细参数
}

