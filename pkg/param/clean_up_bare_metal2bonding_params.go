// Copyright (c) ZStack.io, Inc.

package param

// CleanUpBareMetal2BondingDetailParam CleanUpBareMetal2Bonding详细参数
type CleanUpBareMetal2BondingDetailParam struct {
	rest string `json:"chassisUuid" validate:"required"` // 必填
}

// CleanUpBareMetal2BondingParam CleanUpBareMetal2Bonding请求参数
type CleanUpBareMetal2BondingParam struct {
	BaseParam
	Params CleanUpBareMetal2BondingDetailParam `json:"params"` // 详细参数
}

