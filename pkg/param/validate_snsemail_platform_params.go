// Copyright (c) ZStack.io, Inc.

package param

// ValidateSNSEmailPlatformDetailParam ValidateSNSEmailPlatform detail param
type ValidateSNSEmailPlatformDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
	SmtpServer string `json:"smtpServer,omitempty"`
	SmtpPort int `json:"smtpPort,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	EncryptType string `json:"encryptType,omitempty"`
}

// ValidateSNSEmailPlatformParam ValidateSNSEmailPlatform request param
type ValidateSNSEmailPlatformParam struct {
	BaseParam
	Params ValidateSNSEmailPlatformDetailParam `json:"params"`
}
