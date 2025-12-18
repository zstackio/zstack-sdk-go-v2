// Copyright (c) ZStack.io, Inc.

package param

// UpdateIscsiServerDetailParam UpdateIscsiServer详细参数
type UpdateIscsiServerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"chapUserName,omitempty"`
	rest string `json:"chapUserPassword,omitempty"`
	rest string `json:"state,omitempty"`
}

// UpdateIscsiServerParam UpdateIscsiServer请求参数
type UpdateIscsiServerParam struct {
	BaseParam
	Params UpdateIscsiServerDetailParam `json:"params"` // 详细参数
}

