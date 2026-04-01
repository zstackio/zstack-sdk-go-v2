// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteUserGroupParamDetail DeleteUserGroup detail param
type DeleteUserGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteUserGroupParam DeleteUserGroup request param
type DeleteUserGroupParam struct {
	BaseParam
	Params DeleteUserGroupParamDetail `json:"deleteUserGroup"`
}
// CreateUserGroupParamDetail CreateUserGroup detail param
type CreateUserGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateUserGroupParam CreateUserGroup request param
type CreateUserGroupParam struct {
	BaseParam
	Params CreateUserGroupParamDetail `json:"params"`
}
// UpdateUserGroupParamDetail UpdateUserGroup detail param
type UpdateUserGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateUserGroupParam UpdateUserGroup request param
type UpdateUserGroupParam struct {
	BaseParam
	Params UpdateUserGroupParamDetail `json:"updateUserGroup"`
}
