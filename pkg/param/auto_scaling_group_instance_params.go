// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateAutoScalingGroupInstanceParamDetail UpdateAutoScalingGroupInstance detail param
type UpdateAutoScalingGroupInstanceParamDetail struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	InstanceUuid string `json:"instanceUuid" validate:"required"`
	ProtectionStrategy string `json:"protectionStrategy,omitempty"`
}

// UpdateAutoScalingGroupInstanceParam UpdateAutoScalingGroupInstance request param
type UpdateAutoScalingGroupInstanceParam struct {
	BaseParam
	Params UpdateAutoScalingGroupInstanceParamDetail `json:"params"`
}
// DeleteAutoScalingGroupInstanceParamDetail DeleteAutoScalingGroupInstance detail param
type DeleteAutoScalingGroupInstanceParamDetail struct {
	InstanceUuid string `json:"instanceUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingGroupInstanceParam DeleteAutoScalingGroupInstance request param
type DeleteAutoScalingGroupInstanceParam struct {
	BaseParam
	Params DeleteAutoScalingGroupInstanceParamDetail `json:"params"`
}
