// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmSchedulingRuleStateDetailParam ChangeVmSchedulingRuleState detail param
type ChangeVmSchedulingRuleStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	State string `json:"state" validate:"required"`
}

// ChangeVmSchedulingRuleStateParam ChangeVmSchedulingRuleState request param
type ChangeVmSchedulingRuleStateParam struct {
	BaseParam
	Params ChangeVmSchedulingRuleStateDetailParam `json:"params"`
}
