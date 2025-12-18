// Copyright (c) ZStack.io, Inc.

package param

// GetManagementNodeArchDetailParam GetManagementNodeArch详细参数
type GetManagementNodeArchDetailParam struct {
}

// GetManagementNodeArchParam GetManagementNodeArch请求参数
type GetManagementNodeArchParam struct {
	BaseParam
	Params GetManagementNodeArchDetailParam `json:"params"` // 详细参数
}

