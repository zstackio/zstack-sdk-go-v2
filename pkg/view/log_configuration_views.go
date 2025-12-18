// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LogConfigurationInventoryView LogConfiguration
type LogConfigurationInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest []string `json:"managementNodeUuids,omitempty"`
}

