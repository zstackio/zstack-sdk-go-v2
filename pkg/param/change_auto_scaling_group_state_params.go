// Copyright (c) ZStack.io, Inc.

package param

// ChangeAutoScalingGroupStateDetailParam ChangeAutoScalingGroupState detail param
type ChangeAutoScalingGroupStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeAutoScalingGroupStateParam ChangeAutoScalingGroupState request param
type ChangeAutoScalingGroupStateParam struct {
	BaseParam
	Params ChangeAutoScalingGroupStateDetailParam `json:"params"`
}
