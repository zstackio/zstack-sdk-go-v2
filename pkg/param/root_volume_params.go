// Copyright (c) ZStack.io, Inc.

package param

// ResizeRootVolumeDetailParam ResizeRootVolume详细参数
type ResizeRootVolumeDetailParam struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest int64 `json:"size" validate:"required"` // 必填
}

// ResizeRootVolumeParam ResizeRootVolume请求参数
type ResizeRootVolumeParam struct {
	BaseParam
	Params ResizeRootVolumeDetailParam `json:"params"` // 详细参数
}

