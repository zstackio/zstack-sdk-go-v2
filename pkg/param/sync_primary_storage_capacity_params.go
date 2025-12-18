// Copyright (c) ZStack.io, Inc.

package param

// SyncPrimaryStorageCapacityDetailParam SyncPrimaryStorageCapacity detail param
type SyncPrimaryStorageCapacityDetailParam struct {
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
}

// SyncPrimaryStorageCapacityParam SyncPrimaryStorageCapacity request param
type SyncPrimaryStorageCapacityParam struct {
	BaseParam
	Params SyncPrimaryStorageCapacityDetailParam `json:"params"`
}
