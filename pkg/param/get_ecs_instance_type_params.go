// Copyright (c) ZStack.io, Inc.

package param

// GetEcsInstanceTypeDetailParam GetEcsInstanceType detail param
type GetEcsInstanceTypeDetailParam struct {
	IdentityZoneUuid string `json:"identityZoneUuid" validate:"required"`
	EcsImageUuid string `json:"ecsImageUuid,omitempty"`
}

// GetEcsInstanceTypeParam GetEcsInstanceType request param
type GetEcsInstanceTypeParam struct {
	BaseParam
	Params GetEcsInstanceTypeDetailParam `json:"params"`
}
