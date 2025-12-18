// Copyright (c) ZStack.io, Inc.

package param

// CheckElaborationContentDetailParam CheckElaborationContent详细参数
type CheckElaborationContentDetailParam struct {
	rest string `json:"elaborateFile,omitempty"`
	rest string `json:"elaborateContent,omitempty"`
}

// CheckElaborationContentParam CheckElaborationContent请求参数
type CheckElaborationContentParam struct {
	BaseParam
	Params CheckElaborationContentDetailParam `json:"params"` // 详细参数
}

