// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateIAM2VirtualIDParamDetail CreateIAM2VirtualID detail param
type CreateIAM2VirtualIDParamDetail struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Description *string `json:"description,omitempty"`
	Attributes []AttributeParam `json:"attributes,omitempty"`
	ProjectUuid *string `json:"projectUuid,omitempty"`
	OrganizationUuid *string `json:"organizationUuid,omitempty"`
	WithoutDefaultRole *bool `json:"withoutDefaultRole,omitempty"`
	Type *string `json:"type,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDParam CreateIAM2VirtualID request param
type CreateIAM2VirtualIDParam struct {
	BaseParam
	Params CreateIAM2VirtualIDParamDetail `json:"params"`
}
// DeleteIAM2VirtualIDParamDetail DeleteIAM2VirtualID detail param
type DeleteIAM2VirtualIDParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteIAM2VirtualIDParam DeleteIAM2VirtualID request param
type DeleteIAM2VirtualIDParam struct {
	BaseParam
	Params DeleteIAM2VirtualIDParamDetail `json:"deleteIAM2VirtualID"`
}
// LoginIAM2VirtualIDParamDetail LoginIAM2VirtualID detail param
type LoginIAM2VirtualIDParamDetail struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	CaptchaUuid *string `json:"captchaUuid,omitempty"`
	VerifyCode *string `json:"verifyCode,omitempty"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LoginIAM2VirtualIDParam LoginIAM2VirtualID request param
type LoginIAM2VirtualIDParam struct {
	BaseParam
	Params LoginIAM2VirtualIDParamDetail `json:"loginIAM2VirtualID"`
}
// UpdateIAM2VirtualIDParamDetail UpdateIAM2VirtualID detail param
type UpdateIAM2VirtualIDParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Password *string `json:"password,omitempty"`
	OldPassword *string `json:"oldPassword,omitempty"`
}

// UpdateIAM2VirtualIDParam UpdateIAM2VirtualID request param
type UpdateIAM2VirtualIDParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDParamDetail `json:"updateIAM2VirtualID"`
}
