// Copyright (c) ZStack.io, Inc.

package param

// GetVmGuestToolsInfoDetailParam GetVmGuestToolsInfo详细参数
type GetVmGuestToolsInfoDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"debug,omitempty"`
}

// GetVmGuestToolsInfoParam GetVmGuestToolsInfo请求参数
type GetVmGuestToolsInfoParam struct {
	BaseParam
	Params GetVmGuestToolsInfoDetailParam `json:"params"` // 详细参数
}

