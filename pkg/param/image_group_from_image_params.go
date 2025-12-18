// Copyright (c) ZStack.io, Inc.

package param

// CreateImageGroupFromImageDetailParam CreateImageGroupFromImage详细参数
type CreateImageGroupFromImageDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"rootVolumeTemplateUuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest []string `json:"dataVolumeTemplateUuids,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateImageGroupFromImageParam CreateImageGroupFromImage请求参数
type CreateImageGroupFromImageParam struct {
	BaseParam
	Params CreateImageGroupFromImageDetailParam `json:"params"` // 详细参数
}

