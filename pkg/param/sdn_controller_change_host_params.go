// Copyright (c) ZStack.io, Inc.

package param

// SdnControllerChangeHostDetailParam SdnControllerChangeHost detail param
type SdnControllerChangeHostDetailParam struct {
	SdnControllerUuid string `json:"sdnControllerUuid" validate:"required"`
	HostUuid string `json:"hostUuid" validate:"required"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	NicNames []string `json:"nicNames,omitempty"`
	VtepIp string `json:"vtepIp,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	BondMode string `json:"bondMode,omitempty"`
	LacpMode string `json:"lacpMode,omitempty"`
}

// SdnControllerChangeHostParam SdnControllerChangeHost request param
type SdnControllerChangeHostParam struct {
	BaseParam
	Params SdnControllerChangeHostDetailParam `json:"params"`
}
