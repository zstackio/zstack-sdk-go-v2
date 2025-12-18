// Copyright (c) ZStack.io, Inc.

package param

// GetManagementNodeDirCapacityDetailParam GetManagementNodeDirCapacity detail param
type GetManagementNodeDirCapacityDetailParam struct {
	ManagementNodeUuids []string `json:"managementNodeUuids,omitempty"`
}

// GetManagementNodeDirCapacityParam GetManagementNodeDirCapacity request param
type GetManagementNodeDirCapacityParam struct {
	BaseParam
	Params GetManagementNodeDirCapacityDetailParam `json:"params"`
}
