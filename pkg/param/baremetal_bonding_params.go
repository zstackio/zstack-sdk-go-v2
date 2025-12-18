// Copyright (c) ZStack.io, Inc.

package param

// CreateBaremetalBondingDetailParam CreateBaremetalBonding详细参数
type CreateBaremetalBondingDetailParam struct {
	rest string `json:"chassisUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest int `json:"mode" validate:"required"` // 必填
	rest string `json:"slaves" validate:"required"` // 必填
	rest string `json:"opts,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateBaremetalBondingParam CreateBaremetalBonding请求参数
type CreateBaremetalBondingParam struct {
	BaseParam
	Params CreateBaremetalBondingDetailParam `json:"params"` // 详细参数
}

