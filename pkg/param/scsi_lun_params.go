// Copyright (c) ZStack.io, Inc.

package param

// UpdateScsiLunDetailParam UpdateScsiLun详细参数
type UpdateScsiLunDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"state,omitempty"`
}

// UpdateScsiLunParam UpdateScsiLun请求参数
type UpdateScsiLunParam struct {
	BaseParam
	Params UpdateScsiLunDetailParam `json:"params"` // 详细参数
}

