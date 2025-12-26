// Copyright (c) ZStack.io, Inc.

package view

// GetMaaSUsageView GetMaaSUsage
type GetMaaSUsageView struct {
	Usages []MaaSUsageView `json:"usages,omitempty"`
	Success bool `json:"success,omitempty"`
}

