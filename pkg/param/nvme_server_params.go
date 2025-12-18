// Copyright (c) ZStack.io, Inc.

package param

// DeleteNvmeServerDetailParam DeleteNvmeServer详细参数
type DeleteNvmeServerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteNvmeServerParam DeleteNvmeServer请求参数
type DeleteNvmeServerParam struct {
	BaseParam
	Params DeleteNvmeServerDetailParam `json:"params"` // 详细参数
}

