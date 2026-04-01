// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ContainerUsageView ContainerUsage
type ContainerUsageView struct {
	Name string `json:"name,omitempty"`
	Value int64 `json:"value,omitempty"`
}

