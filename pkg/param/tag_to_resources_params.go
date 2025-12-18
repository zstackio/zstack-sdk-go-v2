// Copyright (c) ZStack.io, Inc.

package param

// AttachTagToResourcesDetailParam AttachTagToResources详细参数
type AttachTagToResourcesDetailParam struct {
	rest string `json:"tagUuid" validate:"required"` // 必填
	rest []string `json:"resourceUuids" validate:"required"` // 必填
	rest map[string]string `json:"tokens,omitempty"`
}

// AttachTagToResourcesParam AttachTagToResources请求参数
type AttachTagToResourcesParam struct {
	BaseParam
	Params AttachTagToResourcesDetailParam `json:"params"` // 详细参数
}

