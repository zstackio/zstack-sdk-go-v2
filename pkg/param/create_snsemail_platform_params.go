// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSEmailPlatformDetailParam CreateSNSEmailPlatform detail param
type CreateSNSEmailPlatformDetailParam struct {
	SmtpServer string `json:"smtpServer" validate:"required"`
	SmtpPort int `json:"smtpPort" validate:"required"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	EncryptType string `json:"encryptType,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSEmailPlatformParam CreateSNSEmailPlatform request param
type CreateSNSEmailPlatformParam struct {
	BaseParam
	Params CreateSNSEmailPlatformDetailParam `json:"params"`
}
