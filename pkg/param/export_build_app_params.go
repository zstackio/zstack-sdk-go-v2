// Copyright (c) ZStack.io, Inc.

package param

// ExportBuildAppDetailParam ExportBuildApp detail param
type ExportBuildAppDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExportBuildAppParam ExportBuildApp request param
type ExportBuildAppParam struct {
	BaseParam
	Params ExportBuildAppDetailParam `json:"params"`
}
