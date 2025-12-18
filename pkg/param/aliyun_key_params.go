// Copyright (c) ZStack.io, Inc.

package param

// DetachAliyunKeyDetailParam DetachAliyunKey详细参数
type DetachAliyunKeyDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// DetachAliyunKeyParam DetachAliyunKey请求参数
type DetachAliyunKeyParam struct {
	BaseParam
	Params DetachAliyunKeyDetailParam `json:"params"` // 详细参数
}

