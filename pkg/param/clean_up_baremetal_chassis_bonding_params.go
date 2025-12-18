// Copyright (c) ZStack.io, Inc.

package param

// CleanUpBaremetalChassisBondingDetailParam CleanUpBaremetalChassisBonding详细参数
type CleanUpBaremetalChassisBondingDetailParam struct {
	rest string `json:"chassisUuid" validate:"required"` // 必填
}

// CleanUpBaremetalChassisBondingParam CleanUpBaremetalChassisBonding请求参数
type CleanUpBaremetalChassisBondingParam struct {
	BaseParam
	Params CleanUpBaremetalChassisBondingDetailParam `json:"params"` // 详细参数
}

