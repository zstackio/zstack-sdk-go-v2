// Copyright (c) ZStack.io, Inc.

package param

// GetVmSchedulingRulesExecuteStateDetailParam GetVmSchedulingRulesExecuteState detail param
type GetVmSchedulingRulesExecuteStateDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// GetVmSchedulingRulesExecuteStateParam GetVmSchedulingRulesExecuteState request param
type GetVmSchedulingRulesExecuteStateParam struct {
	BaseParam
	Params GetVmSchedulingRulesExecuteStateDetailParam `json:"params"`
}
