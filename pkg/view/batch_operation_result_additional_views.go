// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BatchOperationResultView BatchOperationResult
type BatchOperationResultView struct {
	Inventory interface{} `json:"inventory,omitempty"`
	Uuid      string      `json:"uuid,omitempty"`
	Success   bool        `json:"success,omitempty"`
}
