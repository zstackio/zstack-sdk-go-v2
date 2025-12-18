// Copyright (c) ZStack.io, Inc.

package param

// GetSupportAPIsDetailParam GetSupports detail param
type GetSupportAPIsDetailParam struct {
}

// GetSupportAPIsParam GetSupports request param
type GetSupportAPIsParam struct {
	BaseParam
	Params GetSupportAPIsDetailParam `json:"params"`
}
