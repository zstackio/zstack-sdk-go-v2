// Copyright (c) ZStack.io, Inc.

package param

// SecurityMachineEncryptDetailParam SecurityMachineEncrypt详细参数
type SecurityMachineEncryptDetailParam struct {
	rest string `json:"text" validate:"required"` // 必填
	rest string `json:"algType" validate:"required"` // 必填
}

// SecurityMachineEncryptParam SecurityMachineEncrypt请求参数
type SecurityMachineEncryptParam struct {
	BaseParam
	Params SecurityMachineEncryptDetailParam `json:"params"` // 详细参数
}

