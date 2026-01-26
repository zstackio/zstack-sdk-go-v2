// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MaaSUsageView MaaSUsage
type MaaSUsageView struct {
	Name string `json:"name,omitempty"`
	Value int64 `json:"value,omitempty"`
}

