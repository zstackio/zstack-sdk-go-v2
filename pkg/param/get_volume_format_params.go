// Copyright (c) ZStack.io, Inc.

package param

// GetVolumeFormatDetailParam GetVolumeFormat detail param
type GetVolumeFormatDetailParam struct {
}

// GetVolumeFormatParam GetVolumeFormat request param
type GetVolumeFormatParam struct {
	BaseParam
	Params GetVolumeFormatDetailParam `json:"params"`
}
