// Copyright (c) ZStack.io, Inc.

package param

// CreateUserGroupDetailParam CreateUserGroup detail param
type CreateUserGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateUserGroupParam CreateUserGroup request param
type CreateUserGroupParam struct {
	BaseParam
	Params CreateUserGroupDetailParam `json:"params"`
}
