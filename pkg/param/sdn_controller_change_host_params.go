// Copyright (c) ZStack.io, Inc.

package param

// SdnControllerChangeHostDetailParam SdnControllerChangeHost详细参数
type SdnControllerChangeHostDetailParam struct {
	rest string `json:"sdnControllerUuid" validate:"required"` // 必填
	rest string `json:"hostUuid" validate:"required"` // 必填
	rest string `json:"vSwitchType,omitempty"`
	rest []string `json:"nicNames,omitempty"`
	rest string `json:"vtepIp,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest string `json:"bondMode,omitempty"`
	rest string `json:"lacpMode,omitempty"`
}

// SdnControllerChangeHostParam SdnControllerChangeHost请求参数
type SdnControllerChangeHostParam struct {
	BaseParam
	Params SdnControllerChangeHostDetailParam `json:"params"` // 详细参数
}

