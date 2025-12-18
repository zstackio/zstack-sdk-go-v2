// Copyright (c) ZStack.io, Inc.

package param

// ReloadElaborationDetailParam ReloadElaboration detail param
type ReloadElaborationDetailParam struct {
}

// ReloadElaborationParam ReloadElaboration request param
type ReloadElaborationParam struct {
	BaseParam
	Params ReloadElaborationDetailParam `json:"params"`
}
