// Copyright (c) ZStack.io, Inc.

package param

// GenerateSshKeyPairDetailParam GenerateSshKeyPair detail param
type GenerateSshKeyPairDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

// GenerateSshKeyPairParam GenerateSshKeyPair request param
type GenerateSshKeyPairParam struct {
	BaseParam
	Params GenerateSshKeyPairDetailParam `json:"params"`
}
