// Copyright (c) ZStack.io, Inc.

package param

// GetSignatureServerEncryptPublicKeyDetailParam GetSignatureServerEncryptPublicKey详细参数
type GetSignatureServerEncryptPublicKeyDetailParam struct {
}

// GetSignatureServerEncryptPublicKeyParam GetSignatureServerEncryptPublicKey请求参数
type GetSignatureServerEncryptPublicKeyParam struct {
	BaseParam
	Params GetSignatureServerEncryptPublicKeyDetailParam `json:"params"` // 详细参数
}

