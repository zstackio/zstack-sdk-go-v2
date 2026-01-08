// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ImageGroupRefInventoryView ImageGroupRef
type ImageGroupRefInventoryView struct {
	ImageUuid      string    `json:"imageUuid,omitempty"`
	ImageGroupUuid string    `json:"imageGroupUuid,omitempty"`
	CreateDate     time.Time `json:"createDate,omitempty"`
	LastOpDate     time.Time `json:"lastOpDate,omitempty"`
}

// QueryImageGroupRefView QueryImageGroupRef
type QueryImageGroupRefView struct {
	Inventories []ImageGroupRefInventoryView `json:"inventories,omitempty"`
}
