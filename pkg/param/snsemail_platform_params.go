// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSEmailPlatformDetailParam CreateSNSEmailPlatform详细参数
type CreateSNSEmailPlatformDetailParam struct {
	rest string `json:"smtpServer" validate:"required"` // 必填
	rest int `json:"smtpPort" validate:"required"` // 必填
	rest string `json:"username,omitempty"`
	rest string `json:"password,omitempty"`
	rest string `json:"encryptType,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateSNSEmailPlatformParam CreateSNSEmailPlatform请求参数
type CreateSNSEmailPlatformParam struct {
	BaseParam
	Params CreateSNSEmailPlatformDetailParam `json:"params"` // 详细参数
}

