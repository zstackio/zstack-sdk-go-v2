// Copyright (c) ZStack.io, Inc.

package param

// DetachTagFromResourcesDetailParam DetachTagFromResources detail param
type DetachTagFromResourcesDetailParam struct {
	TagUuid string `json:"tagUuid" validate:"required"`
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
}

// DetachTagFromResourcesParam DetachTagFromResources request param
type DetachTagFromResourcesParam struct {
	BaseParam
	Params DetachTagFromResourcesDetailParam `json:"params"`
}
