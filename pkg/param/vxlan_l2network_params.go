// Copyright (c) ZStack.io, Inc.

package param

// DeleteVxlanL2NetworkDetailParam DeleteVxlanL2Network详细参数
type DeleteVxlanL2NetworkDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVxlanL2NetworkParam DeleteVxlanL2Network请求参数
type DeleteVxlanL2NetworkParam struct {
	BaseParam
	Params DeleteVxlanL2NetworkDetailParam `json:"params"` // 详细参数
}

