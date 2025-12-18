// Copyright (c) ZStack.io, Inc.

package param

// DeleteVCenterDetailParam DeleteVCenter详细参数
type DeleteVCenterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVCenterParam DeleteVCenter请求参数
type DeleteVCenterParam struct {
	BaseParam
	Params DeleteVCenterDetailParam `json:"params"` // 详细参数
}

