// Copyright (c) ZStack.io, Inc.

package param

// AttachTagToResourcesDetailParam AttachTagToResources detail param
type AttachTagToResourcesDetailParam struct {
	TagUuid string `json:"tagUuid" validate:"required"`
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	Tokens map[string]string `json:"tokens,omitempty"`
}

// AttachTagToResourcesParam AttachTagToResources request param
type AttachTagToResourcesParam struct {
	BaseParam
	Params AttachTagToResourcesDetailParam `json:"params"`
}
