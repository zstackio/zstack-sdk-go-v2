// Copyright (c) ZStack.io, Inc.

package param

// GetVmInstanceProtectedRecoveryPointsDetailParam GetVmInstanceProtectedRecoveryPoints detail param
type GetVmInstanceProtectedRecoveryPointsDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVmInstanceProtectedRecoveryPointsParam GetVmInstanceProtectedRecoveryPoints request param
type GetVmInstanceProtectedRecoveryPointsParam struct {
	BaseParam
	Params GetVmInstanceProtectedRecoveryPointsDetailParam `json:"params"`
}
