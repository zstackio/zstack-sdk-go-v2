// Copyright (c) ZStack.io, Inc.

package view

// DeployModelEvalServiceEventView DeployModelEvalServiceEvent
type DeployModelEvalServiceEventView struct {
	Inventory ModelEvalServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
	Tasks []ModelEvaluationTaskInventoryView `json:"tasks,omitempty"`
	Success bool `json:"success,omitempty"`
}

