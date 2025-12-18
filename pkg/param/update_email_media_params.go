// Copyright (c) ZStack.io, Inc.

package param

// UpdateEmailMediaDetailParam UpdateEmailMedia detail param
type UpdateEmailMediaDetailParam struct {
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
	Params UpdateEmailMediaDetailParam `json:"params"`
}
