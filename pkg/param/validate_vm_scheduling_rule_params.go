// Copyright (c) ZStack.io, Inc.

package param

// ValidateVmSchedulingRuleDetailParam ValidateVmSchedulingRule detail param
type ValidateVmSchedulingRuleDetailParam struct {
	VmGroupUuid string `json:"vmGroupUuid" validate:"required"`
	HostGroupUuid string `json:"hostGroupUuid,omitempty"`
	Rule string `json:"rule" validate:"required"`
	Mode string `json:"mode" validate:"required"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
}

// ValidateVmSchedulingRuleParam ValidateVmSchedulingRule request param
type ValidateVmSchedulingRuleParam struct {
	BaseParam
	Params ValidateVmSchedulingRuleDetailParam `json:"params"`
}
