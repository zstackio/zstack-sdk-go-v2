// Copyright (c) ZStack.io, Inc.

package param

// CreateCasClientDetailParam CreateCasClient detail param
type CreateCasClientDetailParam struct {
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
	Params CreateCasClientDetailParam `json:"params"`
}
