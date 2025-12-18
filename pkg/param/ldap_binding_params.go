// Copyright (c) ZStack.io, Inc.

package param

// CreateLdapBindingDetailParam CreateLdapBinding详细参数
type CreateLdapBindingDetailParam struct {
	rest string `json:"ldapUid" validate:"required"` // 必填
	rest string `json:"accountUuid" validate:"required"` // 必填
}

// CreateLdapBindingParam CreateLdapBinding请求参数
type CreateLdapBindingParam struct {
	BaseParam
	Params CreateLdapBindingDetailParam `json:"params"` // 详细参数
}

