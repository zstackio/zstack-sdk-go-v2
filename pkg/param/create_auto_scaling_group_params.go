// Copyright (c) ZStack.io, Inc.

package param

// CreateAutoScalingGroupDetailParam CreateAutoScalingGroup detail param
type CreateAutoScalingGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ScalingResourceType string `json:"scalingResourceType" validate:"required"`
	MinResourceSize int `json:"minResourceSize" validate:"required"`
	MaxResourceSize int `json:"maxResourceSize" validate:"required"`
	DefaultCooldown int64 `json:"defaultCooldown" validate:"required"`
	RemovalPolicy string `json:"removalPolicy" validate:"required"`
	DefaultEnable bool `json:"defaultEnable,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingGroupParam CreateAutoScalingGroup request param
type CreateAutoScalingGroupParam struct {
	BaseParam
	Params CreateAutoScalingGroupDetailParam `json:"params"`
}
