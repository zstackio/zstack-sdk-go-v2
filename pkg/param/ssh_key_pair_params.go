// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateSshKeyPairParamDetail CreateSshKeyPair detail param
type CreateSshKeyPairParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PublicKey string `json:"publicKey" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSshKeyPairParam CreateSshKeyPair request param
type CreateSshKeyPairParam struct {
	BaseParam
	CreateSshKeyPair CreateSshKeyPairParamDetail `json:"createSshKeyPair"`
}
// GenerateSshKeyPairParamDetail GenerateSshKeyPair detail param
type GenerateSshKeyPairParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

// GenerateSshKeyPairParam GenerateSshKeyPair request param
type GenerateSshKeyPairParam struct {
	BaseParam
	GenerateSshKeyPair GenerateSshKeyPairParamDetail `json:"generateSshKeyPair"`
}
// UpdateSshKeyPairParamDetail UpdateSshKeyPair detail param
type UpdateSshKeyPairParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateSshKeyPairParam UpdateSshKeyPair request param
type UpdateSshKeyPairParam struct {
	BaseParam
	UpdateSshKeyPair UpdateSshKeyPairParamDetail `json:"updateSshKeyPair"`
}
// DeleteSshKeyPairParamDetail DeleteSshKeyPair detail param
type DeleteSshKeyPairParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteSshKeyPairParam DeleteSshKeyPair request param
type DeleteSshKeyPairParam struct {
	BaseParam
	DeleteSshKeyPair DeleteSshKeyPairParamDetail `json:"deleteSshKeyPair"`
}
