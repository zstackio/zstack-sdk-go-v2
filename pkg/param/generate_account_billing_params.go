// Copyright (c) ZStack.io, Inc.

package param

// GenerateAccountBillingDetailParam GenerateAccountBilling detail param
type GenerateAccountBillingDetailParam struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
}

// GenerateAccountBillingParam GenerateAccountBilling request param
type GenerateAccountBillingParam struct {
	BaseParam
	Params GenerateAccountBillingDetailParam `json:"params"`
}
