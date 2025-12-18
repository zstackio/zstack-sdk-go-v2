// Copyright (c) ZStack.io, Inc.

package param

// AddThirdpartyPlatformDetailParam AddThirdpartyPlatform detail param
type AddThirdpartyPlatformDetailParam struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
	Url string `json:"url" validate:"required"`
	Template string `json:"template" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddThirdpartyPlatformParam AddThirdpartyPlatform request param
type AddThirdpartyPlatformParam struct {
	BaseParam
	Params AddThirdpartyPlatformDetailParam `json:"params"`
}
