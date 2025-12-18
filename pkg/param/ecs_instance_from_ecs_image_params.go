// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsInstanceFromEcsImageDetailParam CreateEcsInstanceFromEcsImage详细参数
type CreateEcsInstanceFromEcsImageDetailParam struct {
	rest string `json:"ecsRootVolumeType,omitempty"`
	rest string `json:"description,omitempty"`
	rest int64 `json:"ecsRootVolumeGBSize,omitempty"`
	rest string `json:"createMode,omitempty"`
	rest string `json:"privateIpAddress,omitempty"`
	rest string `json:"allocatePublicIp,omitempty"`
	rest string `json:"ecsConsolePassword,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"ecsImageUuid" validate:"required"` // 必填
	rest string `json:"instanceOfferingUuid,omitempty"`
	rest string `json:"instanceType,omitempty"`
	rest string `json:"ecsVSwitchUuid" validate:"required"` // 必填
	rest string `json:"ecsSecurityGroupUuid" validate:"required"` // 必填
	rest string `json:"ecsRootPassword" validate:"required"` // 必填
	rest int64 `json:"ecsBandWidth,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateEcsInstanceFromEcsImageParam CreateEcsInstanceFromEcsImage请求参数
type CreateEcsInstanceFromEcsImageParam struct {
	BaseParam
	Params CreateEcsInstanceFromEcsImageDetailParam `json:"params"` // 详细参数
}

