// Copyright (c) ZStack.io, Inc.

package param

// GetVmInstanceRecoveryPointsDetailParam GetVmInstanceRecoveryPoints detail param
type GetVmInstanceRecoveryPointsDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StartTime string `json:"startTime,omitempty"`
	EndTime string `json:"endTime,omitempty"`
	Scale string `json:"scale,omitempty"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVmInstanceRecoveryPointsParam GetVmInstanceRecoveryPoints request param
type GetVmInstanceRecoveryPointsParam struct {
	BaseParam
	Params GetVmInstanceRecoveryPointsDetailParam `json:"params"`
}
