// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateIAM2ProjectParamDetail CreateIAM2Project detail param
type CreateIAM2ProjectParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Attributes []AttributeParam `json:"attributes,omitempty"`
	Quota map[string]int64 `json:"quota,omitempty"`
	RoleUuids []string `json:"roleUuids,omitempty"`
	ResourceTemplates []string `json:"resourceTemplates,omitempty"`
	OrganizationUuid *string `json:"organizationUuid,omitempty"`
	LinkAccountUuid *string `json:"linkAccountUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectParam CreateIAM2Project request param
type CreateIAM2ProjectParam struct {
	BaseParam
	Params CreateIAM2ProjectParamDetail `json:"params"`
}
// DeleteIAM2ProjectParamDetail DeleteIAM2Project detail param
type DeleteIAM2ProjectParamDetail struct {
}

// DeleteIAM2ProjectParam DeleteIAM2Project request param
type DeleteIAM2ProjectParam struct {
	BaseParam
	Params DeleteIAM2ProjectParamDetail `json:"deleteIAM2Project"`
}
// RecoverIAM2ProjectParamDetail RecoverIAM2Project detail param
type RecoverIAM2ProjectParamDetail struct {
}

// RecoverIAM2ProjectParam RecoverIAM2Project request param
type RecoverIAM2ProjectParam struct {
	BaseParam
	Params RecoverIAM2ProjectParamDetail `json:"recoverIAM2Project"`
}
// ExpungeIAM2ProjectParamDetail ExpungeIAM2Project detail param
type ExpungeIAM2ProjectParamDetail struct {
}

// ExpungeIAM2ProjectParam ExpungeIAM2Project request param
type ExpungeIAM2ProjectParam struct {
	BaseParam
	Params ExpungeIAM2ProjectParamDetail `json:"expungeIAM2Project"`
}
// LoginIAM2ProjectParamDetail LoginIAM2Project detail param
type LoginIAM2ProjectParamDetail struct {
	ProjectName string `json:"projectName" validate:"required"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LoginIAM2ProjectParam LoginIAM2Project request param
type LoginIAM2ProjectParam struct {
	BaseParam
	Params LoginIAM2ProjectParamDetail `json:"loginIAM2Project"`
}
// UpdateIAM2ProjectParamDetail UpdateIAM2Project detail param
type UpdateIAM2ProjectParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateIAM2ProjectParam UpdateIAM2Project request param
type UpdateIAM2ProjectParam struct {
	BaseParam
	Params UpdateIAM2ProjectParamDetail `json:"updateIAM2Project"`
}
