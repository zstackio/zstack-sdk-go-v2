// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunSmsSNSTextTemplateDetailParam CreateAliyunSmsSNSTextTemplate detail param
type CreateAliyunSmsSNSTextTemplateDetailParam struct {
	Sign string `json:"sign" validate:"required"`
	AlarmTemplateCode string `json:"alarmTemplateCode" validate:"required"`
	EventTemplateCode string `json:"eventTemplateCode" validate:"required"`
	EventTemplate string `json:"eventTemplate,omitempty"`
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

// CreateAliyunSmsSNSTextTemplateParam CreateAliyunSmsSNSTextTemplate request param
type CreateAliyunSmsSNSTextTemplateParam struct {
	BaseParam
	Params CreateAliyunSmsSNSTextTemplateDetailParam `json:"params"`
}
