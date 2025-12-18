// Copyright (c) ZStack.io, Inc.

package param

// GetBareMetal2SupportedBootModeDetailParam GetBareMetal2SupportedBootMode detail param
type GetBareMetal2SupportedBootModeDetailParam struct {
}

// GetBareMetal2SupportedBootModeParam GetBareMetal2SupportedBootMode request param
type GetBareMetal2SupportedBootModeParam struct {
	BaseParam
	Params GetBareMetal2SupportedBootModeDetailParam `json:"params"`
}
