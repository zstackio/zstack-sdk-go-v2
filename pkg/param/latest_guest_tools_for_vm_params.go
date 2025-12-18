// Copyright (c) ZStack.io, Inc.

package param

// GetLatestGuestToolsForVmDetailParam GetLatestGuestToolsForVm详细参数
type GetLatestGuestToolsForVmDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetLatestGuestToolsForVmParam GetLatestGuestToolsForVm请求参数
type GetLatestGuestToolsForVmParam struct {
	BaseParam
	Params GetLatestGuestToolsForVmDetailParam `json:"params"` // 详细参数
}

