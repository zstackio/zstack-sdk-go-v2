// Copyright (c) ZStack.io, Inc.

package param

// SyncLdapServerDetailParam SyncLdapServer detail param
type SyncLdapServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SyncLdapServerParam SyncLdapServer request param
type SyncLdapServerParam struct {
	BaseParam
	Params SyncLdapServerDetailParam `json:"params"`
}
