// Copyright (c) ZStack.io, Inc.

package param

// RefreshSearchIndexesDetailParam RefreshSearchIndexes detail param
type RefreshSearchIndexesDetailParam struct {
}

// RefreshSearchIndexesParam RefreshSearchIndexes request param
type RefreshSearchIndexesParam struct {
	BaseParam
	Params RefreshSearchIndexesDetailParam `json:"params"`
}
