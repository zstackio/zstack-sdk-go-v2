// Copyright (c) ZStack.io, Inc.

package view

import "time"

// NetworkServiceProviderInventoryView NetworkServiceProvider
type NetworkServiceProviderInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"networkServiceTypes,omitempty"`
	rest []string `json:"attachedL2NetworkUuids,omitempty"`
}

