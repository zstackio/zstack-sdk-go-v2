// Copyright (c) ZStack.io, Inc.

package param

// GetVolumeFormatDetailParam GetVolumeFormat详细参数
type GetVolumeFormatDetailParam struct {
}

// GetVolumeFormatParam GetVolumeFormat请求参数
type GetVolumeFormatParam struct {
	BaseParam
	Params GetVolumeFormatDetailParam `json:"params"` // 详细参数
}

