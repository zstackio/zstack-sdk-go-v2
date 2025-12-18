// Copyright (c) ZStack.io, Inc.

package param

// CreateVmSchedulingRuleDetailParam CreateVmSchedulingRule detail param
type CreateVmSchedulingRuleDetailParam struct {
	Rule string `json:"rule" validate:"required"`
	Mode string `json:"mode" validate:"required"`
	VmGroupUuid string `json:"vmGroupUuid" validate:"required"`
	HostGroupUuid string `json:"hostGroupUuid,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Policy string `json:"policy,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	SubType string `json:"subType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmSchedulingRuleParam CreateVmSchedulingRule request param
type CreateVmSchedulingRuleParam struct {
	BaseParam
	Params CreateVmSchedulingRuleDetailParam `json:"params"`
}
