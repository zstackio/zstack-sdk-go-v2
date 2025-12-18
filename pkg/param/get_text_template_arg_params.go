// Copyright (c) ZStack.io, Inc.

package param

// GetTextTemplateArgDetailParam GetTextTemplateArg detail param
type GetTextTemplateArgDetailParam struct {
}

// GetTextTemplateArgParam GetTextTemplateArg request param
type GetTextTemplateArgParam struct {
	BaseParam
	Params GetTextTemplateArgDetailParam `json:"params"`
}
