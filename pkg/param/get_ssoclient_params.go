// Copyright (c) ZStack.io, Inc.

package param

// GetSSOClientDetailParam GetSSOClient detail param
type GetSSOClientDetailParam struct {
}

// GetSSOClientParam GetSSOClient request param
type GetSSOClientParam struct {
	BaseParam
	Params GetSSOClientDetailParam `json:"params"`
}
