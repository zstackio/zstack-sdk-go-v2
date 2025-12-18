// Copyright (c) ZStack.io, Inc.

package param

// GetVmSchedulingRulesExecuteStateDetailParam GetVmSchedulingRulesExecuteState详细参数
type GetVmSchedulingRulesExecuteStateDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
}

// GetVmSchedulingRulesExecuteStateParam GetVmSchedulingRulesExecuteState请求参数
type GetVmSchedulingRulesExecuteStateParam struct {
	BaseParam
	Params GetVmSchedulingRulesExecuteStateDetailParam `json:"params"` // 详细参数
}

