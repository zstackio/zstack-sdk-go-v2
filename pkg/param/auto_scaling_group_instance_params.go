// Copyright (c) ZStack.io, Inc.

package param

// DeleteAutoScalingGroupInstanceDetailParam DeleteAutoScalingGroupInstance详细参数
type DeleteAutoScalingGroupInstanceDetailParam struct {
	rest string `json:"instanceUuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingGroupInstanceParam DeleteAutoScalingGroupInstance请求参数
type DeleteAutoScalingGroupInstanceParam struct {
	BaseParam
	Params DeleteAutoScalingGroupInstanceDetailParam `json:"params"` // 详细参数
}

