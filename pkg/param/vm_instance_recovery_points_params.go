// Copyright (c) ZStack.io, Inc.

package param

// GetVmInstanceRecoveryPointsDetailParam GetVmInstanceRecoveryPoints详细参数
type GetVmInstanceRecoveryPointsDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"startTime,omitempty"`
	rest string `json:"endTime,omitempty"`
	rest string `json:"scale,omitempty"`
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVmInstanceRecoveryPointsParam GetVmInstanceRecoveryPoints请求参数
type GetVmInstanceRecoveryPointsParam struct {
	BaseParam
	Params GetVmInstanceRecoveryPointsDetailParam `json:"params"` // 详细参数
}

