// Copyright (c) ZStack.io, Inc.

package param

// GetLocalStorageHostDiskCapacityDetailParam GetLocalStorageHostDiskCapacity详细参数
type GetLocalStorageHostDiskCapacityDetailParam struct {
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuid" validate:"required"` // 必填
}

// GetLocalStorageHostDiskCapacityParam GetLocalStorageHostDiskCapacity请求参数
type GetLocalStorageHostDiskCapacityParam struct {
	BaseParam
	Params GetLocalStorageHostDiskCapacityDetailParam `json:"params"` // 详细参数
}

