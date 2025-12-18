// Copyright (c) ZStack.io, Inc.

package param

// RemoveSdnControllerDetailParam RemoveSdnController详细参数
type RemoveSdnControllerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveSdnControllerParam RemoveSdnController请求参数
type RemoveSdnControllerParam struct {
	BaseParam
	Params RemoveSdnControllerDetailParam `json:"params"` // 详细参数
}

