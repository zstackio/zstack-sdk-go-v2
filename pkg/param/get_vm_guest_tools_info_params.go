// Copyright (c) ZStack.io, Inc.

package param

// GetVmGuestToolsInfoDetailParam GetVmGuestToolsInfo detail param
type GetVmGuestToolsInfoDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Debug []string `json:"debug,omitempty"`
}

// GetVmGuestToolsInfoParam GetVmGuestToolsInfo request param
type GetVmGuestToolsInfoParam struct {
	BaseParam
	Params GetVmGuestToolsInfoDetailParam `json:"params"`
}
