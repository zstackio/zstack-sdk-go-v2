// Copyright (c) ZStack.io, Inc.

package param

// GetTaskProgressDetailParam GetTaskProgress detail param
type GetTaskProgressDetailParam struct {
	ApiId string `json:"apiId,omitempty"`
	All bool `json:"all,omitempty"`
}

// GetTaskProgressParam GetTaskProgress request param
type GetTaskProgressParam struct {
	BaseParam
	Params GetTaskProgressDetailParam `json:"params"`
}
