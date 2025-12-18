// Copyright (c) ZStack.io, Inc.

package param

// CreateSshKeyPairDetailParam CreateSshKeyPair detail param
type CreateSshKeyPairDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PublicKey string `json:"publicKey" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSshKeyPairParam CreateSshKeyPair request param
type CreateSshKeyPairParam struct {
	BaseParam
	Params CreateSshKeyPairDetailParam `json:"params"`
}
