// Copyright (c) ZStack.io, Inc.

package param

// GetVersionDetailParam GetVersion详细参数
type GetVersionDetailParam struct {
}

// GetVersionParam GetVersion请求参数
type GetVersionParam struct {
	BaseParam
	Params GetVersionDetailParam `json:"params"` // 详细参数
}

