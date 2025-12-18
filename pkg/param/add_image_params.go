// Copyright (c) ZStack.io, Inc.

package param

// AddImageDetailParam AddImage详细参数
type AddImageDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"mediaType,omitempty"`
	rest string `json:"guestOsType,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest bool `json:"system,omitempty"`
	rest string `json:"format,omitempty"`
	rest string `json:"platform,omitempty"`
	rest []string `json:"backupStorageUuids" validate:"required"` // 必填
	rest string `json:"type,omitempty"`
	rest bool `json:"virtio,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddImageParam AddImage请求参数
type AddImageParam struct {
	BaseParam
	Params AddImageDetailParam `json:"params"` // 详细参数
}

