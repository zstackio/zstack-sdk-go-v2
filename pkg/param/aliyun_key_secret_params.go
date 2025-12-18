// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunKeySecretDetailParam UpdateAliyunKeySecret详细参数
type UpdateAliyunKeySecretDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateAliyunKeySecretParam UpdateAliyunKeySecret请求参数
type UpdateAliyunKeySecretParam struct {
	BaseParam
	Params UpdateAliyunKeySecretDetailParam `json:"params"` // 详细参数
}

