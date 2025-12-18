// Copyright (c) ZStack.io, Inc.

package param

// SdnControllerRemoveHostDetailParam SdnControllerRemoveHost detail param
type SdnControllerRemoveHostDetailParam struct {
	SdnControllerUuid string `json:"sdnControllerUuid" validate:"required"`
	HostUuid string `json:"hostUuid" validate:"required"`
	VSwitchType string `json:"vSwitchType,omitempty"`
}

// SdnControllerRemoveHostParam SdnControllerRemoveHost request param
type SdnControllerRemoveHostParam struct {
	BaseParam
	Params SdnControllerRemoveHostDetailParam `json:"params"`
}
