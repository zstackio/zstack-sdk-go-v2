// Copyright (c) ZStack.io, Inc.

package param

// GetLocalStorageHostDiskCapacityDetailParam GetLocalStorageHostDiskCapacity detail param
type GetLocalStorageHostDiskCapacityDetailParam struct {
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
}

// GetLocalStorageHostDiskCapacityParam GetLocalStorageHostDiskCapacity request param
type GetLocalStorageHostDiskCapacityParam struct {
	BaseParam
	Params GetLocalStorageHostDiskCapacityDetailParam `json:"params"`
}
