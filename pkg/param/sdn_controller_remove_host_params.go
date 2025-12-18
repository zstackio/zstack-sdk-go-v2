// Copyright (c) ZStack.io, Inc.

package param

// SdnControllerRemoveHostDetailParam SdnControllerRemoveHost详细参数
type SdnControllerRemoveHostDetailParam struct {
	rest string `json:"sdnControllerUuid" validate:"required"` // 必填
	rest string `json:"hostUuid" validate:"required"` // 必填
	rest string `json:"vSwitchType,omitempty"`
}

// SdnControllerRemoveHostParam SdnControllerRemoveHost请求参数
type SdnControllerRemoveHostParam struct {
	BaseParam
	Params SdnControllerRemoveHostDetailParam `json:"params"` // 详细参数
}

