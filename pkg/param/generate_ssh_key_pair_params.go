// Copyright (c) ZStack.io, Inc.

package param

// GenerateSshKeyPairDetailParam GenerateSshKeyPair详细参数
type GenerateSshKeyPairDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
}

// GenerateSshKeyPairParam GenerateSshKeyPair请求参数
type GenerateSshKeyPairParam struct {
	BaseParam
	Params GenerateSshKeyPairDetailParam `json:"params"` // 详细参数
}

