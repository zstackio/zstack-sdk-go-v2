// Copyright (c) ZStack.io, Inc.

package param

// CancelLongJobDetailParam CancelLongJob详细参数
type CancelLongJobDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// CancelLongJobParam CancelLongJob请求参数
type CancelLongJobParam struct {
	BaseParam
	Params CancelLongJobDetailParam `json:"params"` // 详细参数
}

