// Copyright (c) ZStack.io, Inc.

package param

// RefreshSearchIndexesDetailParam RefreshSearchIndexes详细参数
type RefreshSearchIndexesDetailParam struct {
}

// RefreshSearchIndexesParam RefreshSearchIndexes请求参数
type RefreshSearchIndexesParam struct {
	BaseParam
	Params RefreshSearchIndexesDetailParam `json:"params"` // 详细参数
}

