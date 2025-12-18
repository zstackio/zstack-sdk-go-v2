// Copyright (c) ZStack.io, Inc.

package param

// DeleteAutoScalingGroupDetailParam DeleteAutoScalingGroup detail param
type DeleteAutoScalingGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingGroupParam DeleteAutoScalingGroup request param
type DeleteAutoScalingGroupParam struct {
	BaseParam
	Params DeleteAutoScalingGroupDetailParam `json:"params"`
}
