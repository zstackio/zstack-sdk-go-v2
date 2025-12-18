// Copyright (c) ZStack.io, Inc.

package param

// SyncPrimaryStorageCapacityDetailParam SyncPrimaryStorageCapacity详细参数
type SyncPrimaryStorageCapacityDetailParam struct {
	rest string `json:"primaryStorageUuid" validate:"required"` // 必填
}

// SyncPrimaryStorageCapacityParam SyncPrimaryStorageCapacity请求参数
type SyncPrimaryStorageCapacityParam struct {
	BaseParam
	Params SyncPrimaryStorageCapacityDetailParam `json:"params"` // 详细参数
}

