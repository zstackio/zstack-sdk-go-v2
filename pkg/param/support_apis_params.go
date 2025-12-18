// Copyright (c) ZStack.io, Inc.

package param

// GetSupportAPIsDetailParam GetSupports详细参数
type GetSupportAPIsDetailParam struct {
}

// GetSupportAPIsParam GetSupports请求参数
type GetSupportAPIsParam struct {
	BaseParam
	Params GetSupportAPIsDetailParam `json:"params"` // 详细参数
}

