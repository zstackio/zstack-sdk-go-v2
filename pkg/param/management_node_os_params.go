// Copyright (c) ZStack.io, Inc.

package param

// GetManagementNodeOSDetailParam GetManagementNodeOS详细参数
type GetManagementNodeOSDetailParam struct {
}

// GetManagementNodeOSParam GetManagementNodeOS请求参数
type GetManagementNodeOSParam struct {
	BaseParam
	Params GetManagementNodeOSDetailParam `json:"params"` // 详细参数
}

