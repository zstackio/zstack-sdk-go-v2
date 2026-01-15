// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelServiceRefInventoryView ModelServiceRef
type ModelServiceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ModelUuid string `json:"modelUuid,omitempty"`
	ModelServiceUuid string `json:"modelServiceUuid,omitempty"`
}

