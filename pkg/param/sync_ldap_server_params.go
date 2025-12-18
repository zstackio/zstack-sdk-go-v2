// Copyright (c) ZStack.io, Inc.

package param

// SyncLdapServerDetailParam SyncLdapServer详细参数
type SyncLdapServerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// SyncLdapServerParam SyncLdapServer请求参数
type SyncLdapServerParam struct {
	BaseParam
	Params SyncLdapServerDetailParam `json:"params"` // 详细参数
}

