// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ServiceStatusView ServiceStatus
type ServiceStatusView struct {
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
}
