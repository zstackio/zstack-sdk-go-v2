// Copyright (c) ZStack.io, Inc.

package param

// GenerateAccountBillingDetailParam GenerateAccountBilling详细参数
type GenerateAccountBillingDetailParam struct {
	rest string `json:"accountUuid" validate:"required"` // 必填
}

// GenerateAccountBillingParam GenerateAccountBilling请求参数
type GenerateAccountBillingParam struct {
	BaseParam
	Params GenerateAccountBillingDetailParam `json:"params"` // 详细参数
}

