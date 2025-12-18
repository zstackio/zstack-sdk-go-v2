// Copyright (c) ZStack.io, Inc.

package param

// RemoveVmSchedulingRuleDetailParam RemoveVmSchedulingRule detail param
type RemoveVmSchedulingRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveVmSchedulingRuleParam RemoveVmSchedulingRule request param
type RemoveVmSchedulingRuleParam struct {
	BaseParam
	Params RemoveVmSchedulingRuleDetailParam `json:"params"`
}
