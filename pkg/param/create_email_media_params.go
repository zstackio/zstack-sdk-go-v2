// Copyright (c) ZStack.io, Inc.

package param

// CreateEmailMediaDetailParam CreateEmailMedia detail param
type CreateEmailMediaDetailParam struct {
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
	Params CreateEmailMediaDetailParam `json:"params"`
}
