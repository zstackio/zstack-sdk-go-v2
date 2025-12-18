// Copyright (c) ZStack.io, Inc.

package param

// CreateAutoScalingGroupAddingNewInstanceRuleDetailParam CreateAutoScalingGroupAddingNewInstanceRule detail param
type CreateAutoScalingGroupAddingNewInstanceRuleDetailParam struct {
	AdjustmentType string `json:"adjustmentType" validate:"required"`
	AdjustmentValue int `json:"adjustmentValue" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	AutoScalingGroupUuid string `json:"autoScalingGroupUuid" validate:"required"`
	Type string `json:"type,omitempty"`
	Cooldown int64 `json:"cooldown,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingGroupAddingNewInstanceRuleParam CreateAutoScalingGroupAddingNewInstanceRule request param
type CreateAutoScalingGroupAddingNewInstanceRuleParam struct {
	BaseParam
	Params CreateAutoScalingGroupAddingNewInstanceRuleDetailParam `json:"params"`
}
