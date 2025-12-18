// Copyright (c) ZStack.io, Inc.

package param

// GetEncryptedFieldDetailParam GetEncryptedField详细参数
type GetEncryptedFieldDetailParam struct {
	rest string `json:"encryptedType,omitempty"`
}

// GetEncryptedFieldParam GetEncryptedField请求参数
type GetEncryptedFieldParam struct {
	BaseParam
	Params GetEncryptedFieldDetailParam `json:"params"` // 详细参数
}

