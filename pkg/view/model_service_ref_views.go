// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelServiceRefInventoryView ModelServiceRef
type ModelServiceRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ModelUuid *string `json:"modelUuid,omitempty"`
	ModelServiceUuid *string `json:"modelServiceUuid,omitempty"`
}

