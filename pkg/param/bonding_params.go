// Copyright (c) ZStack.io, Inc.

package param

// CreateBondingDetailParam CreateBonding详细参数
type CreateBondingDetailParam struct {
	rest []string `json:"hostUuids" validate:"required"` // 必填
	rest string `json:"bondingName" validate:"required"` // 必填
	rest []string `json:"slaveUuids,omitempty"`
	rest []string `json:"slaveNames,omitempty"`
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"mode" validate:"required"` // 必填
	rest string `json:"xmitHashPolicy,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateBondingParam CreateBonding请求参数
type CreateBondingParam struct {
	BaseParam
	Params CreateBondingDetailParam `json:"params"` // 详细参数
}

