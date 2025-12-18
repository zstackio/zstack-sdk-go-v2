// Copyright (c) ZStack.io, Inc.

package param

// ChangeSecretResourcePoolStateDetailParam ChangeSecretResourcePoolState详细参数
type ChangeSecretResourcePoolStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeSecretResourcePoolStateParam ChangeSecretResourcePoolState请求参数
type ChangeSecretResourcePoolStateParam struct {
	BaseParam
	Params ChangeSecretResourcePoolStateDetailParam `json:"params"` // 详细参数
}

