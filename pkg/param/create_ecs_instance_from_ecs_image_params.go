// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsInstanceFromEcsImageDetailParam CreateEcsInstanceFromEcsImage detail param
type CreateEcsInstanceFromEcsImageDetailParam struct {
	EcsRootVolumeType string `json:"ecsRootVolumeType,omitempty"`
	Description string `json:"description,omitempty"`
	EcsRootVolumeGBSize int64 `json:"ecsRootVolumeGBSize,omitempty"`
	CreateMode string `json:"createMode,omitempty"`
	PrivateIpAddress string `json:"privateIpAddress,omitempty"`
	AllocatePublicIp string `json:"allocatePublicIp,omitempty"`
	EcsConsolePassword string `json:"ecsConsolePassword,omitempty"`
	Name string `json:"name" validate:"required"`
	EcsImageUuid string `json:"ecsImageUuid" validate:"required"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
	InstanceType string `json:"instanceType,omitempty"`
	EcsVSwitchUuid string `json:"ecsVSwitchUuid" validate:"required"`
	EcsSecurityGroupUuid string `json:"ecsSecurityGroupUuid" validate:"required"`
	EcsRootPassword string `json:"ecsRootPassword" validate:"required"`
	EcsBandWidth int64 `json:"ecsBandWidth,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsInstanceFromEcsImageParam CreateEcsInstanceFromEcsImage request param
type CreateEcsInstanceFromEcsImageParam struct {
	BaseParam
	Params CreateEcsInstanceFromEcsImageDetailParam `json:"params"`
}
