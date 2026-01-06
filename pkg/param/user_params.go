// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteUserParamDetail DeleteUser detail param
type DeleteUserParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteUserParam DeleteUser request param
type DeleteUserParam struct {
	BaseParam
	Params DeleteUserParamDetail `json:"params"`
}
// UpdateUserParamDetail UpdateUser detail param
type UpdateUserParamDetail struct {
	Uuid string `json:"uuid,omitempty"`
	Password string `json:"password,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	OldPassword string `json:"oldPassword,omitempty"`
}

// UpdateUserParam UpdateUser request param
type UpdateUserParam struct {
	BaseParam
	Params UpdateUserParamDetail `json:"params"`
}
// CreateUserParamDetail CreateUser detail param
type CreateUserParamDetail struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateUserParam CreateUser request param
type CreateUserParam struct {
	BaseParam
	Params CreateUserParamDetail `json:"params"`
}
