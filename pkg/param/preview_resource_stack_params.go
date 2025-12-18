// Copyright (c) ZStack.io, Inc.

package param

// PreviewResourceStackDetailParam PreviewResourceStack detail param
type PreviewResourceStackDetailParam struct {
	Type string `json:"type,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	PreParameters string `json:"preParameters,omitempty"`
}

// PreviewResourceStackParam PreviewResourceStack request param
type PreviewResourceStackParam struct {
	BaseParam
	Params PreviewResourceStackDetailParam `json:"params"`
}
