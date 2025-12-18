// Copyright (c) ZStack.io, Inc.

package param

// GetEncryptedFieldDetailParam GetEncryptedField detail param
type GetEncryptedFieldDetailParam struct {
	EncryptedType string `json:"encryptedType,omitempty"`
}

// GetEncryptedFieldParam GetEncryptedField request param
type GetEncryptedFieldParam struct {
	BaseParam
	Params GetEncryptedFieldDetailParam `json:"params"`
}
