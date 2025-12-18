// Copyright (c) ZStack.io, Inc.

package param

// RefreshSharedblockDeviceCapacityDetailParam RefreshSharedblockDeviceCapacity详细参数
type RefreshSharedblockDeviceCapacityDetailParam struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"sharedBlockGroupUuid" validate:"required"` // 必填
}

// RefreshSharedblockDeviceCapacityParam RefreshSharedblockDeviceCapacity请求参数
type RefreshSharedblockDeviceCapacityParam struct {
	BaseParam
	Params RefreshSharedblockDeviceCapacityDetailParam `json:"params"` // 详细参数
}

