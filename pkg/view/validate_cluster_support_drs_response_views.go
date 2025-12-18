// Copyright (c) ZStack.io, Inc.

package view

// ValidateClusterSupportDRSView ValidateClusterSupportDRS
type ValidateClusterSupportDRSView struct {
	Supported bool `json:"supported,omitempty"`
	Reason interface{} `json:"reason,omitempty"`
	Success bool `json:"success,omitempty"`
}

