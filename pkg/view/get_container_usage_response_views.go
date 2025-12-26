// Copyright (c) ZStack.io, Inc.

package view

// GetContainerUsageView GetContainerUsage
type GetContainerUsageView struct {
	Usages []ContainerUsageView `json:"usages,omitempty"`
	Success bool `json:"success,omitempty"`
}

