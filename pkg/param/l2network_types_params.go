// Copyright (c) ZStack.io, Inc.

package param

// GetL2NetworkTypesDetailParam GetL2NetworkTypes详细参数
type GetL2NetworkTypesDetailParam struct {
}

// GetL2NetworkTypesParam GetL2NetworkTypes请求参数
type GetL2NetworkTypesParam struct {
	BaseParam
	Params GetL2NetworkTypesDetailParam `json:"params"` // 详细参数
}

