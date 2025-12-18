// Copyright (c) ZStack.io, Inc.

package param

// RefreshIscsiServerDetailParam RefreshIscsiServer详细参数
type RefreshIscsiServerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// RefreshIscsiServerParam RefreshIscsiServer请求参数
type RefreshIscsiServerParam struct {
	BaseParam
	Params RefreshIscsiServerDetailParam `json:"params"` // 详细参数
}

