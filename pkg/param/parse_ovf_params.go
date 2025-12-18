// Copyright (c) ZStack.io, Inc.

package param

// ParseOvfDetailParam ParseOvf detail param
type ParseOvfDetailParam struct {
	XmlBase64 string `json:"xmlBase64" validate:"required"`
}

// ParseOvfParam ParseOvf request param
type ParseOvfParam struct {
	BaseParam
	Params ParseOvfDetailParam `json:"params"`
}
