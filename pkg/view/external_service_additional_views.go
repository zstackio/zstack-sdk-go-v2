// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ExternalServiceInventoryView ExternalService
type ExternalServiceInventoryView struct {
	Name         string                          `json:"name,omitempty"`
	Status       string                          `json:"status,omitempty"`
	Capabilities ExternalServiceCapabilitiesView `json:"capabilities,omitempty"`
}
