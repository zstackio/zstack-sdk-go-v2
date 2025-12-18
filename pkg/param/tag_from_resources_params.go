// Copyright (c) ZStack.io, Inc.

package param

// DetachTagFromResourcesDetailParam DetachTagFromResources详细参数
type DetachTagFromResourcesDetailParam struct {
	rest string `json:"tagUuid" validate:"required"` // 必填
	rest []string `json:"resourceUuids" validate:"required"` // 必填
}

// DetachTagFromResourcesParam DetachTagFromResources请求参数
type DetachTagFromResourcesParam struct {
	BaseParam
	Params DetachTagFromResourcesDetailParam `json:"params"` // 详细参数
}

