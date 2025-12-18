// Copyright (c) ZStack.io, Inc.

package param

// CreateUserDetailParam CreateUser detail param
type CreateUserDetailParam struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateUserParam CreateUser request param
type CreateUserParam struct {
	BaseParam
	Params CreateUserDetailParam `json:"params"`
}
