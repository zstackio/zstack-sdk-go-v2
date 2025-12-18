// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSTextTemplateDetailParam UpdateSNSTextTemplate detail param
type UpdateSNSTextTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Subject string `json:"subject,omitempty"`
	RecoverySubject string `json:"recoverySubject,omitempty"`
	Template string `json:"template,omitempty"`
	RecoveryTemplate string `json:"recoveryTemplate,omitempty"`
	DefaultTemplate bool `json:"defaultTemplate,omitempty"`
}

// UpdateSNSTextTemplateParam UpdateSNSTextTemplate request param
type UpdateSNSTextTemplateParam struct {
	BaseParam
	Params UpdateSNSTextTemplateDetailParam `json:"params"`
}
