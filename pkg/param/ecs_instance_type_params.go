// Copyright (c) ZStack.io, Inc.

package param

// GetEcsInstanceTypeDetailParam GetEcsInstanceType详细参数
type GetEcsInstanceTypeDetailParam struct {
	rest string `json:"identityZoneUuid" validate:"required"` // 必填
	rest string `json:"ecsImageUuid,omitempty"`
}

// GetEcsInstanceTypeParam GetEcsInstanceType请求参数
type GetEcsInstanceTypeParam struct {
	BaseParam
	Params GetEcsInstanceTypeDetailParam `json:"params"` // 详细参数
}

