// Copyright (c) ZStack.io, Inc.

package view

import "time"

// FiberChannelStorageInventoryView FiberChannelStorage
type FiberChannelStorageInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"wwnn,omitempty"`
	rest string `json:"state,omitempty"`
	rest []FiberChannelLunInventoryView `json:"fiberChannelLuns,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

