// Copyright (c) ZStack.io, Inc.

package param

// ReloadElaborationDetailParam ReloadElaboration详细参数
type ReloadElaborationDetailParam struct {
}

// ReloadElaborationParam ReloadElaboration请求参数
type ReloadElaborationParam struct {
	BaseParam
	Params ReloadElaborationDetailParam `json:"params"` // 详细参数
}

