// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteIAM2OrganizationParamDetail DeleteIAM2Organization detail param
type DeleteIAM2OrganizationParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteIAM2OrganizationParam DeleteIAM2Organization request param
type DeleteIAM2OrganizationParam struct {
	BaseParam
	Params DeleteIAM2OrganizationParamDetail `json:"params"`
}
// UpdateIAM2OrganizationParamDetail UpdateIAM2Organization detail param
type UpdateIAM2OrganizationParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	Type string `json:"type,omitempty"`
}

// UpdateIAM2OrganizationParam UpdateIAM2Organization request param
type UpdateIAM2OrganizationParam struct {
	BaseParam
	Params UpdateIAM2OrganizationParamDetail `json:"params"`
}
// CreateIAM2OrganizationParamDetail CreateIAM2Organization detail param
type CreateIAM2OrganizationParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	ParentUuid string `json:"parentUuid,omitempty"`
	Attributes []AttributeParam `json:"attributes,omitempty"`
	Quota map[string]int64 `json:"quota,omitempty"`
	SrcType string `json:"srcType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2OrganizationParam CreateIAM2Organization request param
type CreateIAM2OrganizationParam struct {
	BaseParam
	Params CreateIAM2OrganizationParamDetail `json:"params"`
}
