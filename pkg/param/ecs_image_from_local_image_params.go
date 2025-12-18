// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsImageFromLocalImageDetailParam CreateEcsImageFromLocalImage详细参数
type CreateEcsImageFromLocalImageDetailParam struct {
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateEcsImageFromLocalImageParam CreateEcsImageFromLocalImage请求参数
type CreateEcsImageFromLocalImageParam struct {
	BaseParam
	Params CreateEcsImageFromLocalImageDetailParam `json:"params"` // 详细参数
}

