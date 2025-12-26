// Copyright (c) ZStack.io, Inc.

package view

// GetChainTaskView GetChainTask
type GetChainTaskView struct {
	Results map[string]ChainInfoView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

