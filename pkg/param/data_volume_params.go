// Copyright (c) ZStack.io, Inc.

package param

// ExpungeDataVolumeDetailParam ExpungeDataVolume详细参数
type ExpungeDataVolumeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ExpungeDataVolumeParam ExpungeDataVolume请求参数
type ExpungeDataVolumeParam struct {
	BaseParam
	Params ExpungeDataVolumeDetailParam `json:"params"` // 详细参数
}

