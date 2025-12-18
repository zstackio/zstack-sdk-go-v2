// Copyright (c) ZStack.io, Inc.

package param

// CheckElaborationContentDetailParam CheckElaborationContent detail param
type CheckElaborationContentDetailParam struct {
	ElaborateFile string `json:"elaborateFile,omitempty"`
	ElaborateContent string `json:"elaborateContent,omitempty"`
}

// CheckElaborationContentParam CheckElaborationContent request param
type CheckElaborationContentParam struct {
	BaseParam
	Params CheckElaborationContentDetailParam `json:"params"`
}
