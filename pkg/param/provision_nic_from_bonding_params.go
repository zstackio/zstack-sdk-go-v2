// Copyright (c) ZStack.io, Inc.

package param

// DetachProvisionNicFromBondingDetailParam DetachProvisionNicFromBonding详细参数
type DetachProvisionNicFromBondingDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"provisionNicUuid" validate:"required"` // 必填
}

// DetachProvisionNicFromBondingParam DetachProvisionNicFromBonding请求参数
type DetachProvisionNicFromBondingParam struct {
	BaseParam
	Params DetachProvisionNicFromBondingDetailParam `json:"params"` // 详细参数
}

