// Copyright (c) ZStack.io, Inc.

package param

// GetAttachableVpcL3NetworkDetailParam GetAttachableVpcL3Network详细参数
type GetAttachableVpcL3NetworkDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetAttachableVpcL3NetworkParam GetAttachableVpcL3Network请求参数
type GetAttachableVpcL3NetworkParam struct {
	BaseParam
	Params GetAttachableVpcL3NetworkDetailParam `json:"params"` // 详细参数
}

