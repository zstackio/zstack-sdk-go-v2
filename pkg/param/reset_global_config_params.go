// Copyright (c) ZStack.io, Inc.

package param

// ResetGlobalConfigDetailParam ResetGlobalConfig detail param
type ResetGlobalConfigDetailParam struct {
}

// ResetGlobalConfigParam ResetGlobalConfig request param
type ResetGlobalConfigParam struct {
	BaseParam
	Params ResetGlobalConfigDetailParam `json:"params"`
}
