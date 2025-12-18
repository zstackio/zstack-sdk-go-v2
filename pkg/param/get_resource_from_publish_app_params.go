// Copyright (c) ZStack.io, Inc.

package param

// GetResourceFromPublishAppDetailParam GetResourceFromPublishApp detail param
type GetResourceFromPublishAppDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
}

// GetResourceFromPublishAppParam GetResourceFromPublishApp request param
type GetResourceFromPublishAppParam struct {
	BaseParam
	Params GetResourceFromPublishAppDetailParam `json:"params"`
}
