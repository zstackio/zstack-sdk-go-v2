// Copyright (c) ZStack.io, Inc.

package param

// PreviewResourceFromAppDetailParam PreviewResourceFromApp detail param
type PreviewResourceFromAppDetailParam struct {
	AppUuid string `json:"appUuid" validate:"required"`
	Parameters string `json:"parameters,omitempty"`
}

// PreviewResourceFromAppParam PreviewResourceFromApp request param
type PreviewResourceFromAppParam struct {
	BaseParam
	Params PreviewResourceFromAppDetailParam `json:"params"`
}
