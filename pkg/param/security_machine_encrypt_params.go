// Copyright (c) ZStack.io, Inc.

package param

// SecurityMachineEncryptDetailParam SecurityMachineEncrypt detail param
type SecurityMachineEncryptDetailParam struct {
	Text string `json:"text" validate:"required"`
	AlgType string `json:"algType" validate:"required"`
}

// SecurityMachineEncryptParam SecurityMachineEncrypt request param
type SecurityMachineEncryptParam struct {
	BaseParam
	Params SecurityMachineEncryptDetailParam `json:"params"`
}
