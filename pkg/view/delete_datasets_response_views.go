// Copyright (c) ZStack.io, Inc.

package view

// DeleteDatasetsEventView DeleteDatasetsEvent
type DeleteDatasetsEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

