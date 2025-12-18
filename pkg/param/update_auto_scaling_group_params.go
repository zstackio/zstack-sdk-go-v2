// Copyright (c) ZStack.io, Inc.

package param

// UpdateAutoScalingGroupDetailParam UpdateAutoScalingGroup detail param
type UpdateAutoScalingGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	MinResourceSize int `json:"minResourceSize,omitempty"`
	MaxResourceSize int `json:"maxResourceSize,omitempty"`
	RemovalPolicy string `json:"removalPolicy,omitempty"`
}

// UpdateAutoScalingGroupParam UpdateAutoScalingGroup request param
type UpdateAutoScalingGroupParam struct {
	BaseParam
	Params UpdateAutoScalingGroupDetailParam `json:"params"`
}
