// Copyright (c) ZStack.io, Inc.

package param

// SdnControllerAddHostDetailParam SdnControllerAddHost detail param
type SdnControllerAddHostDetailParam struct {
	SdnControllerUuid string `json:"sdnControllerUuid" validate:"required"`
	HostUuid string `json:"hostUuid" validate:"required"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	NicNames []string `json:"nicNames" validate:"required"`
	VtepIp string `json:"vtepIp,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	BondMode string `json:"bondMode,omitempty"`
	LacpMode string `json:"lacpMode,omitempty"`
}

// SdnControllerAddHostParam SdnControllerAddHost request param
type SdnControllerAddHostParam struct {
	BaseParam
	Params SdnControllerAddHostDetailParam `json:"params"`
}
