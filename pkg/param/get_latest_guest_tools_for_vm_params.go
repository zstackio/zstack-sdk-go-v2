// Copyright (c) ZStack.io, Inc.

package param

// GetLatestGuestToolsForVmDetailParam GetLatestGuestToolsForVm detail param
type GetLatestGuestToolsForVmDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetLatestGuestToolsForVmParam GetLatestGuestToolsForVm request param
type GetLatestGuestToolsForVmParam struct {
	BaseParam
	Params GetLatestGuestToolsForVmDetailParam `json:"params"`
}
