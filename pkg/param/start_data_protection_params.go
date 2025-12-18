// Copyright (c) ZStack.io, Inc.

package param

// StartDataProtectionDetailParam StartDataProtection detail param
type StartDataProtectionDetailParam struct {
	EncryptType string `json:"encryptType" validate:"required"`
	AuditsIntegrityDate int `json:"auditsIntegrityDate,omitempty"`
}

// StartDataProtectionParam StartDataProtection request param
type StartDataProtectionParam struct {
	BaseParam
	Params StartDataProtectionDetailParam `json:"params"`
}
