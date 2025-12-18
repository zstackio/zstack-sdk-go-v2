// Copyright (c) ZStack.io, Inc.

package param

// DeleteSshKeyPairDetailParam DeleteSshKeyPair detail param
type DeleteSshKeyPairDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteSshKeyPairParam DeleteSshKeyPair request param
type DeleteSshKeyPairParam struct {
	BaseParam
	Params DeleteSshKeyPairDetailParam `json:"params"`
}
