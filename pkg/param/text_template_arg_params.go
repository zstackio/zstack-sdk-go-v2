// Copyright (c) ZStack.io, Inc.

package param

// GetTextTemplateArgDetailParam GetTextTemplateArg详细参数
type GetTextTemplateArgDetailParam struct {
}

// GetTextTemplateArgParam GetTextTemplateArg请求参数
type GetTextTemplateArgParam struct {
	BaseParam
	Params GetTextTemplateArgDetailParam `json:"params"` // 详细参数
}

