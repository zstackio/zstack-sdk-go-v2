// Copyright (c) ZStack.io, Inc.

package param

// GetChainTaskDetailParam GetChainTask detail param
type GetChainTaskDetailParam struct {
	SyncSignatures []string `json:"syncSignatures,omitempty"`
}

// GetChainTaskParam GetChainTask request param
type GetChainTaskParam struct {
	BaseParam
	Params GetChainTaskDetailParam `json:"params"`
}
