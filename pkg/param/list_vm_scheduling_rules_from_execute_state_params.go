// Copyright (c) ZStack.io, Inc.

package param

// ListVmSchedulingRulesFromExecuteStateDetailParam ListVmSchedulingRulesFromExecuteState detail param
type ListVmSchedulingRulesFromExecuteStateDetailParam struct {
	ExecuteStates []string `json:"executeStates" validate:"required"`
}

// ListVmSchedulingRulesFromExecuteStateParam ListVmSchedulingRulesFromExecuteState request param
type ListVmSchedulingRulesFromExecuteStateParam struct {
	BaseParam
	Params ListVmSchedulingRulesFromExecuteStateDetailParam `json:"params"`
}
