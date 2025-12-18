// Copyright (c) ZStack.io, Inc.

package param

// UpdateIPsecConnectionDetailParam UpdateIPsecConnection详细参数
type UpdateIPsecConnectionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateIPsecConnectionParam UpdateIPsecConnection请求参数
type UpdateIPsecConnectionParam struct {
	BaseParam
	Params UpdateIPsecConnectionDetailParam `json:"params"` // 详细参数
}

