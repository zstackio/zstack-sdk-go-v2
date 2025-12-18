// Copyright (c) ZStack.io, Inc.

package param

// AllocateHostResourceDetailParam AllocateHostResource详细参数
type AllocateHostResourceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"strategy" validate:"required"` // 必填
	rest string `json:"scene" validate:"required"` // 必填
	rest int `json:"vcpu" validate:"required"` // 必填
	rest int64 `json:"memSize,omitempty"`
}

// AllocateHostResourceParam AllocateHostResource请求参数
type AllocateHostResourceParam struct {
	BaseParam
	Params AllocateHostResourceDetailParam `json:"params"` // 详细参数
}

