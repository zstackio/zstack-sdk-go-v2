// Copyright (c) ZStack.io, Inc.

package param

// ReconnectSdnControllerDetailParam ReconnectSdnController detail param
type ReconnectSdnControllerDetailParam struct {
	SdnControllerUuid string `json:"sdnControllerUuid" validate:"required"`
}

// ReconnectSdnControllerParam ReconnectSdnController request param
type ReconnectSdnControllerParam struct {
	BaseParam
	Params ReconnectSdnControllerDetailParam `json:"params"`
}
