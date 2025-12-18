// Copyright (c) ZStack.io, Inc.

package param

// DetachNicFromBondingDetailParam DetachNicFromBonding详细参数
type DetachNicFromBondingDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"slaveUuids" validate:"required"` // 必填
	rest string `json:"type,omitempty"`
}

// DetachNicFromBondingParam DetachNicFromBonding请求参数
type DetachNicFromBondingParam struct {
	BaseParam
	Params DetachNicFromBondingDetailParam `json:"params"` // 详细参数
}

