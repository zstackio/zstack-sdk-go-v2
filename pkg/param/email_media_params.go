// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateEmailMediaParamDetail CreateEmailMedia detail param
type CreateEmailMediaParamDetail struct {
	SmtpServer string `json:"smtpServer" validate:"required"`
	SmtpPort int `json:"smtpPort" validate:"required"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEmailMediaParam CreateEmailMedia request param
type CreateEmailMediaParam struct {
	BaseParam
	CreateEmailMedia CreateEmailMediaParamDetail `json:"createEmailMedia"`
}
// UpdateEmailMediaParamDetail UpdateEmailMedia detail param
type UpdateEmailMediaParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	SmtpServer string `json:"smtpServer,omitempty"`
	SmtpPort int `json:"smtpPort,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// UpdateEmailMediaParam UpdateEmailMedia request param
type UpdateEmailMediaParam struct {
	BaseParam
	UpdateEmailMedia UpdateEmailMediaParamDetail `json:"updateEmailMedia"`
}
