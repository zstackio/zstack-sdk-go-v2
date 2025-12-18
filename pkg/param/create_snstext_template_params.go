// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSTextTemplateDetailParam CreateSNSTextTemplate detail param
type CreateSNSTextTemplateDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ApplicationPlatformType string `json:"applicationPlatformType" validate:"required"`
	Subject string `json:"subject,omitempty"`
	RecoverySubject string `json:"recoverySubject,omitempty"`
	Template string `json:"template" validate:"required"`
	RecoveryTemplate string `json:"recoveryTemplate,omitempty"`
	DefaultTemplate bool `json:"defaultTemplate,omitempty"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSTextTemplateParam CreateSNSTextTemplate request param
type CreateSNSTextTemplateParam struct {
	BaseParam
	Params CreateSNSTextTemplateDetailParam `json:"params"`
}
