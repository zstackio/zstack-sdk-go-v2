// Copyright (c) ZStack.io, Inc.

package view

// GetIAM2ProjectContainerImageTagsView GetIAM2ProjectContainerImageTags
type GetIAM2ProjectContainerImageTagsView struct {
	Inventories []ContainerImageTagInventoryView `json:"inventories,omitempty"`
	Total int `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

