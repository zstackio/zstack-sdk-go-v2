// Copyright (c) ZStack.io, Inc.

package param

// UpdateCasClientDetailParam UpdateCasClient detail param
type UpdateCasClientDetailParam struct {
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
	Params UpdateCasClientDetailParam `json:"params"`
}
