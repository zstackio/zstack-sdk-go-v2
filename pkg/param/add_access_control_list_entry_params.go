// Copyright (c) ZStack.io, Inc.

package param

// AddAccessControlListEntryDetailParam AddAccessControlListEntry详细参数
type AddAccessControlListEntryDetailParam struct {
	rest string `json:"aclUuid" validate:"required"` // 必填
	rest string `json:"entries" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAccessControlListEntryParam AddAccessControlListEntry请求参数
type AddAccessControlListEntryParam struct {
	BaseParam
	Params AddAccessControlListEntryDetailParam `json:"params"` // 详细参数
}

