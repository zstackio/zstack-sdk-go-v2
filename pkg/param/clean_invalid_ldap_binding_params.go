// Copyright (c) ZStack.io, Inc.

package param

// CleanInvalidLdapBindingDetailParam CleanInvalidLdapBinding详细参数
type CleanInvalidLdapBindingDetailParam struct {
}

// CleanInvalidLdapBindingParam CleanInvalidLdapBinding请求参数
type CleanInvalidLdapBindingParam struct {
	BaseParam
	Params CleanInvalidLdapBindingDetailParam `json:"params"` // 详细参数
}

