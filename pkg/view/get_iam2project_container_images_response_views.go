// Copyright (c) ZStack.io, Inc.

package view

// GetIAM2ProjectContainerImagesView GetIAM2ProjectContainerImages
type GetIAM2ProjectContainerImagesView struct {
	Inventories []ZakuImageInventoryView `json:"inventories,omitempty"`
	Total int `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

