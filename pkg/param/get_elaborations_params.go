// Copyright (c) ZStack.io, Inc.

package param

// GetElaborationsDetailParam GetElaborations detail param
type GetElaborationsDetailParam struct {
	Category string `json:"category,omitempty"`
	Code string `json:"code,omitempty"`
	Regex string `json:"regex,omitempty"`
}

// GetElaborationsParam GetElaborations request param
type GetElaborationsParam struct {
	BaseParam
	Params GetElaborationsDetailParam `json:"params"`
}
