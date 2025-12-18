// Copyright (c) ZStack.io, Inc.

package param

// GetL3NetworkTypesDetailParam GetL3NetworkTypes详细参数
type GetL3NetworkTypesDetailParam struct {
}

// GetL3NetworkTypesParam GetL3NetworkTypes请求参数
type GetL3NetworkTypesParam struct {
	BaseParam
	Params GetL3NetworkTypesDetailParam `json:"params"` // 详细参数
}

