// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// FiberChannelStorageInventoryView FiberChannelStorage
type FiberChannelStorageInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Wwnn string `json:"wwnn,omitempty"`
	State string `json:"state,omitempty"`
	FiberChannelLuns []FiberChannelLunInventoryView `json:"fiberChannelLuns,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

