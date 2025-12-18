// Copyright (c) ZStack.io, Inc.

package param

// GetManagementNodeDirCapacityDetailParam GetManagementNodeDirCapacity详细参数
type GetManagementNodeDirCapacityDetailParam struct {
	rest []string `json:"managementNodeUuids,omitempty"`
}

// GetManagementNodeDirCapacityParam GetManagementNodeDirCapacity请求参数
type GetManagementNodeDirCapacityParam struct {
	BaseParam
	Params GetManagementNodeDirCapacityDetailParam `json:"params"` // 详细参数
}

