// Copyright (c) ZStack.io, Inc.

package param

// DeleteBondingDetailParam DeleteBonding detail param
type DeleteBondingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBondingParam DeleteBonding request param
type DeleteBondingParam struct {
	BaseParam
	Params DeleteBondingDetailParam `json:"params"`
}
