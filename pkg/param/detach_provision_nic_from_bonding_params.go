// Copyright (c) ZStack.io, Inc.

package param

// DetachProvisionNicFromBondingDetailParam DetachProvisionNicFromBonding detail param
type DetachProvisionNicFromBondingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ProvisionNicUuid string `json:"provisionNicUuid" validate:"required"`
}

// DetachProvisionNicFromBondingParam DetachProvisionNicFromBonding request param
type DetachProvisionNicFromBondingParam struct {
	BaseParam
	Params DetachProvisionNicFromBondingDetailParam `json:"params"`
}
