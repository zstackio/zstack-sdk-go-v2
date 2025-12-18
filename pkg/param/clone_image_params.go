// Copyright (c) ZStack.io, Inc.

package param

// CloneImageDetailParam CloneImage detail param
type CloneImageDetailParam struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	Strategy string `json:"strategy,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CloneImageParam CloneImage request param
type CloneImageParam struct {
	BaseParam
	Params CloneImageDetailParam `json:"params"`
}
