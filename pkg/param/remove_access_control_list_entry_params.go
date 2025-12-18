// Copyright (c) ZStack.io, Inc.

package param

// RemoveAccessControlListEntryDetailParam RemoveAccessControlListEntry detail param
type RemoveAccessControlListEntryDetailParam struct {
	AclUuid string `json:"aclUuid" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveAccessControlListEntryParam RemoveAccessControlListEntry request param
type RemoveAccessControlListEntryParam struct {
	BaseParam
	Params RemoveAccessControlListEntryDetailParam `json:"params"`
}
