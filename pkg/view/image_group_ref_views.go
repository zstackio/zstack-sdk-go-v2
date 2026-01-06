// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ImageGroupRefInventoryView ImageGroupRef
type ImageGroupRefInventoryView struct {
	ImageUuid string `json:"imageUuid,omitempty"`
	ImageGroupUuid string `json:"imageGroupUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryImageGroupRefView QueryImageGroupRef
type QueryImageGroupRefView struct {
	Inventories []ImageGroupRefInventoryView `json:"inventories,omitempty"`
}

