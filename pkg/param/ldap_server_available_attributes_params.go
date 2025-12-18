// Copyright (c) ZStack.io, Inc.

package param

// GetLdapServerAvailableAttributesDetailParam GetLdapServerAvailableAttributes详细参数
type GetLdapServerAvailableAttributesDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetLdapServerAvailableAttributesParam GetLdapServerAvailableAttributes请求参数
type GetLdapServerAvailableAttributesParam struct {
	BaseParam
	Params GetLdapServerAvailableAttributesDetailParam `json:"params"` // 详细参数
}

