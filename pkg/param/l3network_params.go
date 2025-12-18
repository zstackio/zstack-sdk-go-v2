// Copyright (c) ZStack.io, Inc.

package param

// UpdateL3NetworkDetailParam UpdateL3Network详细参数
type UpdateL3NetworkDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"dnsDomain,omitempty"`
	rest string `json:"category,omitempty"`
	rest bool `json:"system,omitempty"`
}

// UpdateL3NetworkParam UpdateL3Network请求参数
type UpdateL3NetworkParam struct {
	BaseParam
	Params UpdateL3NetworkDetailParam `json:"params"` // 详细参数
}

