// Copyright (c) ZStack.io, Inc.

package param

// FlattenVolumeDetailParam FlattenVolume详细参数
type FlattenVolumeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"dryRun,omitempty"`
}

// FlattenVolumeParam FlattenVolume请求参数
type FlattenVolumeParam struct {
	BaseParam
	Params FlattenVolumeDetailParam `json:"params"` // 详细参数
}

