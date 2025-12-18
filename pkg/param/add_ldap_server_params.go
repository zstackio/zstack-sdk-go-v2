// Copyright (c) ZStack.io, Inc.

package param

// AddLdapServerDetailParam AddLdapServer详细参数
type AddLdapServerDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"base" validate:"required"` // 必填
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"encryption" validate:"required"` // 必填
	rest string `json:"scope" validate:"required"` // 必填
}

// AddLdapServerParam AddLdapServer请求参数
type AddLdapServerParam struct {
	BaseParam
	Params AddLdapServerDetailParam `json:"params"` // 详细参数
}

