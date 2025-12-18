// Copyright (c) ZStack.io, Inc.

package param

// ListVmSchedulingRulesFromExecuteStateDetailParam ListVmSchedulingRulesFromExecuteState详细参数
type ListVmSchedulingRulesFromExecuteStateDetailParam struct {
	rest []string `json:"executeStates" validate:"required"` // 必填
}

// ListVmSchedulingRulesFromExecuteStateParam ListVmSchedulingRulesFromExecuteState请求参数
type ListVmSchedulingRulesFromExecuteStateParam struct {
	BaseParam
	Params ListVmSchedulingRulesFromExecuteStateDetailParam `json:"params"` // 详细参数
}

