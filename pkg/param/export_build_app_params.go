// Copyright (c) ZStack.io, Inc.

package param

// ExportBuildAppDetailParam ExportBuildApp详细参数
type ExportBuildAppDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ExportBuildAppParam ExportBuildApp请求参数
type ExportBuildAppParam struct {
	BaseParam
	Params ExportBuildAppDetailParam `json:"params"` // 详细参数
}

