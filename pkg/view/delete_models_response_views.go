// Copyright (c) ZStack.io, Inc.

package view

// DeleteModelsEventView DeleteModelsEvent
type DeleteModelsEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

