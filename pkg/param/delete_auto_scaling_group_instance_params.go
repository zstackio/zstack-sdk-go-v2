// Copyright (c) ZStack.io, Inc.

package param

// DeleteAutoScalingGroupInstanceDetailParam DeleteAutoScalingGroupInstance detail param
type DeleteAutoScalingGroupInstanceDetailParam struct {
	InstanceUuid string `json:"instanceUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingGroupInstanceParam DeleteAutoScalingGroupInstance request param
type DeleteAutoScalingGroupInstanceParam struct {
	BaseParam
	Params DeleteAutoScalingGroupInstanceDetailParam `json:"params"`
}
