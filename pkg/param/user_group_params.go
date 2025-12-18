// Copyright (c) ZStack.io, Inc.

package param

// CreateUserGroupDetailParam CreateUserGroup详细参数
type CreateUserGroupDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateUserGroupParam CreateUserGroup请求参数
type CreateUserGroupParam struct {
	BaseParam
	Params CreateUserGroupDetailParam `json:"params"` // 详细参数
}

