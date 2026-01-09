// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ObservabilityServerServiceDataInventoryView ObservabilityServerServiceData
type ObservabilityServerServiceDataInventoryView struct {
	Job *string `json:"job,omitempty"`
	FileName *string `json:"fileName,omitempty"`
}

