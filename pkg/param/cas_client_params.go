// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateCasClientParamDetail CreateCasClient detail param
type CreateCasClientParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	CasServerLoginUrl string `json:"casServerLoginUrl" validate:"required"`
	CasServerUrlPrefix string `json:"casServerUrlPrefix" validate:"required"`
	ServerName string `json:"serverName" validate:"required"`
	LoginType string `json:"loginType" validate:"required"`
	UrlTemplate string `json:"urlTemplate" validate:"required"`
	Attributes []ExtendedAttributeParam `json:"attributes,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateCasClientParam CreateCasClient request param
type CreateCasClientParam struct {
	BaseParam
	CreateCasClient CreateCasClientParamDetail `json:"createCasClient"`
}
// UpdateCasClientParamDetail UpdateCasClient detail param
type UpdateCasClientParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	CasServerLoginUrl string `json:"casServerLoginUrl,omitempty"`
	CasServerUrlPrefix string `json:"casServerUrlPrefix,omitempty"`
	ServerName string `json:"serverName,omitempty"`
	LoginType string `json:"loginType,omitempty"`
}

// UpdateCasClientParam UpdateCasClient request param
type UpdateCasClientParam struct {
	BaseParam
	UpdateCasClient UpdateCasClientParamDetail `json:"updateCasClient"`
}
