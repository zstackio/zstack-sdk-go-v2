// Copyright (c) ZStack.io, Inc.

package param

// SdnControllerAddHostDetailParam SdnControllerAddHost详细参数
type SdnControllerAddHostDetailParam struct {
	rest string `json:"sdnControllerUuid" validate:"required"` // 必填
	rest string `json:"hostUuid" validate:"required"` // 必填
	rest string `json:"vSwitchType,omitempty"`
	rest []string `json:"nicNames" validate:"required"` // 必填
	rest string `json:"vtepIp,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest string `json:"bondMode,omitempty"`
	rest string `json:"lacpMode,omitempty"`
}

// SdnControllerAddHostParam SdnControllerAddHost请求参数
type SdnControllerAddHostParam struct {
	BaseParam
	Params SdnControllerAddHostDetailParam `json:"params"` // 详细参数
}

