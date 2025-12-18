// Copyright (c) ZStack.io, Inc.

package param

// CreateAutoScalingGroupRemovalInstanceRuleDetailParam CreateAutoScalingGroupRemovalInstanceRule detail param
type CreateAutoScalingGroupRemovalInstanceRuleDetailParam struct {
	AdjustmentType string `json:"adjustmentType" validate:"required"`
	AdjustmentValue int `json:"adjustmentValue" validate:"required"`
	RemovalPolicy string `json:"removalPolicy" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	AutoScalingGroupUuid string `json:"autoScalingGroupUuid" validate:"required"`
	Type string `json:"type,omitempty"`
	Cooldown int64 `json:"cooldown,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingGroupRemovalInstanceRuleParam CreateAutoScalingGroupRemovalInstanceRule request param
type CreateAutoScalingGroupRemovalInstanceRuleParam struct {
	BaseParam
	Params CreateAutoScalingGroupRemovalInstanceRuleDetailParam `json:"params"`
}
