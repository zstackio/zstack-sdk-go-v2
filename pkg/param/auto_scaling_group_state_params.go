// Copyright (c) ZStack.io, Inc.

package param

// ChangeAutoScalingGroupStateDetailParam ChangeAutoScalingGroupState详细参数
type ChangeAutoScalingGroupStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeAutoScalingGroupStateParam ChangeAutoScalingGroupState请求参数
type ChangeAutoScalingGroupStateParam struct {
	BaseParam
	Params ChangeAutoScalingGroupStateDetailParam `json:"params"` // 详细参数
}

