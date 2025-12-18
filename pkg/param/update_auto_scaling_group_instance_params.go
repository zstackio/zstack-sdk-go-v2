// Copyright (c) ZStack.io, Inc.

package param

// UpdateAutoScalingGroupInstanceDetailParam UpdateAutoScalingGroupInstance detail param
type UpdateAutoScalingGroupInstanceDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	InstanceUuid string `json:"instanceUuid" validate:"required"`
	ProtectionStrategy string `json:"protectionStrategy,omitempty"`
}

// UpdateAutoScalingGroupInstanceParam UpdateAutoScalingGroupInstance request param
type UpdateAutoScalingGroupInstanceParam struct {
	BaseParam
	Params UpdateAutoScalingGroupInstanceDetailParam `json:"params"`
}
