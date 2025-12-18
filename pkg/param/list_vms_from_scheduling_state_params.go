// Copyright (c) ZStack.io, Inc.

package param

// ListVmsFromSchedulingStateDetailParam ListVmsFromSchedulingState detail param
type ListVmsFromSchedulingStateDetailParam struct {
	RuleUuid string `json:"ruleUuid" validate:"required"`
	ExecuteStates []string `json:"executeStates" validate:"required"`
}

// ListVmsFromSchedulingStateParam ListVmsFromSchedulingState request param
type ListVmsFromSchedulingStateParam struct {
	BaseParam
	Params ListVmsFromSchedulingStateDetailParam `json:"params"`
}
