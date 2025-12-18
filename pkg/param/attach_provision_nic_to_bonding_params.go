// Copyright (c) ZStack.io, Inc.

package param

// AttachProvisionNicToBondingDetailParam AttachProvisionNicToBonding detail param
type AttachProvisionNicToBondingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ProvisionNicUuid string `json:"provisionNicUuid" validate:"required"`
	BondingUuid string `json:"bondingUuid" validate:"required"`
	ProvisionIp string `json:"provisionIp,omitempty"`
	CustomMac string `json:"customMac,omitempty"`
}

// AttachProvisionNicToBondingParam AttachProvisionNicToBonding request param
type AttachProvisionNicToBondingParam struct {
	BaseParam
	Params AttachProvisionNicToBondingDetailParam `json:"params"`
}
