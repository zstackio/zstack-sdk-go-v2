// Copyright (c) ZStack.io, Inc.

package param

// ListVmsFromSchedulingStateDetailParam ListVmsFromSchedulingState详细参数
type ListVmsFromSchedulingStateDetailParam struct {
	rest string `json:"ruleUuid" validate:"required"` // 必填
	rest []string `json:"executeStates" validate:"required"` // 必填
}

// ListVmsFromSchedulingStateParam ListVmsFromSchedulingState请求参数
type ListVmsFromSchedulingStateParam struct {
	BaseParam
	Params ListVmsFromSchedulingStateDetailParam `json:"params"` // 详细参数
}

