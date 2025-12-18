// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunKeySecretDetailParam AddAliyunKeySecret详细参数
type AddAliyunKeySecretDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"key" validate:"required"` // 必填
	rest string `json:"secret" validate:"required"` // 必填
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest bool `json:"sync,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAliyunKeySecretParam AddAliyunKeySecret请求参数
type AddAliyunKeySecretParam struct {
	BaseParam
	Params AddAliyunKeySecretDetailParam `json:"params"` // 详细参数
}

