// Copyright (c) ZStack.io, Inc.

package view

// DeleteModelServiceInstanceGroupsEventView DeleteModelServiceInstanceGroupsEvent
type DeleteModelServiceInstanceGroupsEventView struct {
	Results []BatchOperationResultView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

