// Copyright (c) ZStack.io, Inc.

package param

// UpdateSdnControllerDetailParam UpdateSdnController详细参数
type UpdateSdnControllerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateSdnControllerParam UpdateSdnController请求参数
type UpdateSdnControllerParam struct {
	BaseParam
	Params UpdateSdnControllerDetailParam `json:"params"` // 详细参数
}

