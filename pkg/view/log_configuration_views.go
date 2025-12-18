// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LogConfigurationInventoryView LogConfiguration
type LogConfigurationInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Type string `json:"type,omitempty"`
	ManagementNodeUuids []string `json:"managementNodeUuids,omitempty"`
}

