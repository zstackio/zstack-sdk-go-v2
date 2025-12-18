// Copyright (c) ZStack.io, Inc.

package param

// ReconnectSdnControllerDetailParam ReconnectSdnController详细参数
type ReconnectSdnControllerDetailParam struct {
	rest string `json:"sdnControllerUuid" validate:"required"` // 必填
}

// ReconnectSdnControllerParam ReconnectSdnController请求参数
type ReconnectSdnControllerParam struct {
	BaseParam
	Params ReconnectSdnControllerDetailParam `json:"params"` // 详细参数
}

