// Copyright (c) ZStack.io, Inc.

package param

// PublishAppDetailParam PublishApp detail param
type PublishAppDetailParam struct {
	BuildAppUuid string `json:"buildAppUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// PublishAppParam PublishApp request param
type PublishAppParam struct {
	BaseParam
	Params PublishAppDetailParam `json:"params"`
}
