// Copyright (c) ZStack.io, Inc.

package param

// CleanUpBareMetal2BondingDetailParam CleanUpBareMetal2Bonding detail param
type CleanUpBareMetal2BondingDetailParam struct {
	ChassisUuid string `json:"chassisUuid" validate:"required"`
}

// CleanUpBareMetal2BondingParam CleanUpBareMetal2Bonding request param
type CleanUpBareMetal2BondingParam struct {
	BaseParam
	Params CleanUpBareMetal2BondingDetailParam `json:"params"`
}
