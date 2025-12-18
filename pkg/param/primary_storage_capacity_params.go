// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageCapacityDetailParam GetPrimaryStorageCapacity详细参数
type GetPrimaryStorageCapacityDetailParam struct {
	rest []string `json:"zoneUuids,omitempty"`
	rest []string `json:"clusterUuids,omitempty"`
	rest []string `json:"primaryStorageUuids,omitempty"`
	rest bool `json:"all,omitempty"`
}

// GetPrimaryStorageCapacityParam GetPrimaryStorageCapacity请求参数
type GetPrimaryStorageCapacityParam struct {
	BaseParam
	Params GetPrimaryStorageCapacityDetailParam `json:"params"` // 详细参数
}

