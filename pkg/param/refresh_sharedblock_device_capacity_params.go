// Copyright (c) ZStack.io, Inc.

package param

// RefreshSharedblockDeviceCapacityDetailParam RefreshSharedblockDeviceCapacity detail param
type RefreshSharedblockDeviceCapacityDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
	SharedBlockGroupUuid string `json:"sharedBlockGroupUuid" validate:"required"`
}

// RefreshSharedblockDeviceCapacityParam RefreshSharedblockDeviceCapacity request param
type RefreshSharedblockDeviceCapacityParam struct {
	BaseParam
	Params RefreshSharedblockDeviceCapacityDetailParam `json:"params"`
}
