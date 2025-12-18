// Copyright (c) ZStack.io, Inc.

package param

// CleanInvalidLdapIAM2BindingDetailParam CleanInvalidLdapIAM2Binding详细参数
type CleanInvalidLdapIAM2BindingDetailParam struct {
}

// CleanInvalidLdapIAM2BindingParam CleanInvalidLdapIAM2Binding请求参数
type CleanInvalidLdapIAM2BindingParam struct {
	BaseParam
	Params CleanInvalidLdapIAM2BindingDetailParam `json:"params"` // 详细参数
}

