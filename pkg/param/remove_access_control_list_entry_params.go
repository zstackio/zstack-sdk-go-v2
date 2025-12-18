// Copyright (c) ZStack.io, Inc.

package param

// RemoveAccessControlListEntryDetailParam RemoveAccessControlListEntry详细参数
type RemoveAccessControlListEntryDetailParam struct {
	rest string `json:"aclUuid" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveAccessControlListEntryParam RemoveAccessControlListEntry请求参数
type RemoveAccessControlListEntryParam struct {
	BaseParam
	Params RemoveAccessControlListEntryDetailParam `json:"params"` // 详细参数
}

