// Copyright (c) ZStack.io, Inc.

package param

// UpdateSshKeyPairDetailParam UpdateSshKeyPair detail param
type UpdateSshKeyPairDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateSshKeyPairParam UpdateSshKeyPair request param
type UpdateSshKeyPairParam struct {
	BaseParam
	Params UpdateSshKeyPairDetailParam `json:"params"`
}
