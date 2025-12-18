// Copyright (c) ZStack.io, Inc.

package param

// GetSignatureServerEncryptPublicKeyDetailParam GetSignatureServerEncryptPublicKey detail param
type GetSignatureServerEncryptPublicKeyDetailParam struct {
}

// GetSignatureServerEncryptPublicKeyParam GetSignatureServerEncryptPublicKey request param
type GetSignatureServerEncryptPublicKeyParam struct {
	BaseParam
	Params GetSignatureServerEncryptPublicKeyDetailParam `json:"params"`
}
