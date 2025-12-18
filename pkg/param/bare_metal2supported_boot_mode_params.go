// Copyright (c) ZStack.io, Inc.

package param

// GetBareMetal2SupportedBootModeDetailParam GetBareMetal2SupportedBootMode详细参数
type GetBareMetal2SupportedBootModeDetailParam struct {
}

// GetBareMetal2SupportedBootModeParam GetBareMetal2SupportedBootMode请求参数
type GetBareMetal2SupportedBootModeParam struct {
	BaseParam
	Params GetBareMetal2SupportedBootModeDetailParam `json:"params"` // 详细参数
}

