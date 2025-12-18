// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2VirtualIDGroupDetailParam CreateIAM2VirtualIDGroup detail param
type CreateIAM2VirtualIDGroupDetailParam struct {
	ProjectUuid string `json:"projectUuid,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Attributes []interface{} `json:"attributes,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDGroupParam CreateIAM2VirtualIDGroup request param
type CreateIAM2VirtualIDGroupParam struct {
	BaseParam
	Params CreateIAM2VirtualIDGroupDetailParam `json:"params"`
}
