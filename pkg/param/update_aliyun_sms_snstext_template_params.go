// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunSmsSNSTextTemplateDetailParam UpdateAliyunSmsSNSTextTemplate detail param
type UpdateAliyunSmsSNSTextTemplateDetailParam struct {
	AlarmTemplateCode string `json:"alarmTemplateCode,omitempty"`
	Sign string `json:"sign,omitempty"`
	EventTemplateCode string `json:"eventTemplateCode,omitempty"`
	EventTemplate string `json:"eventTemplate,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Subject string `json:"subject,omitempty"`
	RecoverySubject string `json:"recoverySubject,omitempty"`
	Template string `json:"template,omitempty"`
	RecoveryTemplate string `json:"recoveryTemplate,omitempty"`
	DefaultTemplate bool `json:"defaultTemplate,omitempty"`
}

// UpdateAliyunSmsSNSTextTemplateParam UpdateAliyunSmsSNSTextTemplate request param
type UpdateAliyunSmsSNSTextTemplateParam struct {
	BaseParam
	Params UpdateAliyunSmsSNSTextTemplateDetailParam `json:"params"`
}
