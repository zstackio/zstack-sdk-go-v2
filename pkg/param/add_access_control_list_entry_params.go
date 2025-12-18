// Copyright (c) ZStack.io, Inc.

package param

// AddAccessControlListEntryDetailParam AddAccessControlListEntry detail param
type AddAccessControlListEntryDetailParam struct {
	AclUuid string `json:"aclUuid" validate:"required"`
	Entries string `json:"entries" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAccessControlListEntryParam AddAccessControlListEntry request param
type AddAccessControlListEntryParam struct {
	BaseParam
	Params AddAccessControlListEntryDetailParam `json:"params"`
}
