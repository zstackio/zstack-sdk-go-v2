// Copyright (c) ZStack.io, Inc.

package view

// DeleteModelEvaluationTasksEventView DeleteModelEvaluationTasksEvent
type DeleteModelEvaluationTasksEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

