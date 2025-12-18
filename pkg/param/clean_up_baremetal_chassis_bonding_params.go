// Copyright (c) ZStack.io, Inc.

package param

// CleanUpBaremetalChassisBondingDetailParam CleanUpBaremetalChassisBonding detail param
type CleanUpBaremetalChassisBondingDetailParam struct {
	ChassisUuid string `json:"chassisUuid" validate:"required"`
}

// CleanUpBaremetalChassisBondingParam CleanUpBaremetalChassisBonding request param
type CleanUpBaremetalChassisBondingParam struct {
	BaseParam
	Params CleanUpBaremetalChassisBondingDetailParam `json:"params"`
}
