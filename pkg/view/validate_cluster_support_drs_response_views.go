// Copyright (c) ZStack.io, Inc.

package view

// ValidateClusterSupportDRSView ValidateClusterSupportDRS
type ValidateClusterSupportDRSView struct {
	Supported bool `json:"supported,omitempty"`
	Reason ErrorCodeView `json:"reason,omitempty"`
	Success bool `json:"success,omitempty"`
}

