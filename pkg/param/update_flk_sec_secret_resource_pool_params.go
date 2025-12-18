// Copyright (c) ZStack.io, Inc.

package param

// UpdateFlkSecSecretResourcePoolDetailParam UpdateFlkSecSecretResourcePool detail param
type UpdateFlkSecSecretResourcePoolDetailParam struct {
	EncryptResult string `json:"encryptResult,omitempty"`
	ActivatedToken string `json:"activatedToken,omitempty"`
	ProtectToken string `json:"protectToken,omitempty"`
	HmacToken string `json:"hmacToken,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Model string `json:"model,omitempty"`
	HeartbeatInterval int `json:"heartbeatInterval,omitempty"`
}

// UpdateFlkSecSecretResourcePoolParam UpdateFlkSecSecretResourcePool request param
type UpdateFlkSecSecretResourcePoolParam struct {
	BaseParam
	Params UpdateFlkSecSecretResourcePoolDetailParam `json:"params"`
}
