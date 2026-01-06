// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateSNSTextTemplateParamDetail UpdateSNSTextTemplate detail param
type UpdateSNSTextTemplateParamDetail struct {
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
	Params UpdateSNSTextTemplateParamDetail `json:"params"`
}
// CreateSNSTextTemplateParamDetail CreateSNSTextTemplate detail param
type CreateSNSTextTemplateParamDetail struct {
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
	Params CreateSNSTextTemplateParamDetail `json:"params"`
}
// DeleteSNSTextTemplateParamDetail DeleteSNSTextTemplate detail param
type DeleteSNSTextTemplateParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSNSTextTemplateParam DeleteSNSTextTemplate request param
type DeleteSNSTextTemplateParam struct {
	BaseParam
	Params DeleteSNSTextTemplateParamDetail `json:"params"`
}
