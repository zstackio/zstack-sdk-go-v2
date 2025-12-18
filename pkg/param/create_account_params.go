// Copyright (c) ZStack.io, Inc.

package param

// CreateAccountDetailParam CreateAccount detail param
type CreateAccountDetailParam struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAccountParam CreateAccount request param
type CreateAccountParam struct {
	BaseParam
	Params CreateAccountDetailParam `json:"params"`
}
