// Copyright (c) ZStack.io, Inc.

package param

// GetVmInstanceProtectedRecoveryPointsDetailParam GetVmInstanceProtectedRecoveryPoints详细参数
type GetVmInstanceProtectedRecoveryPointsDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVmInstanceProtectedRecoveryPointsParam GetVmInstanceProtectedRecoveryPoints请求参数
type GetVmInstanceProtectedRecoveryPointsParam struct {
	BaseParam
	Params GetVmInstanceProtectedRecoveryPointsDetailParam `json:"params"` // 详细参数
}

