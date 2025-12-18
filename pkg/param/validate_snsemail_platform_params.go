// Copyright (c) ZStack.io, Inc.

package param

// ValidateSNSEmailPlatformDetailParam ValidateSNSEmailPlatform详细参数
type ValidateSNSEmailPlatformDetailParam struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"smtpServer,omitempty"`
	rest int `json:"smtpPort,omitempty"`
	rest string `json:"username,omitempty"`
	rest string `json:"password,omitempty"`
	rest string `json:"encryptType,omitempty"`
}

// ValidateSNSEmailPlatformParam ValidateSNSEmailPlatform请求参数
type ValidateSNSEmailPlatformParam struct {
	BaseParam
	Params ValidateSNSEmailPlatformDetailParam `json:"params"` // 详细参数
}

