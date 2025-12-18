// Copyright (c) ZStack.io, Inc.

package view

// GetVmXmlView GetVmXml
type GetVmXmlView struct {
	Match bool `json:"match,omitempty"`
	RunningXml string `json:"runningXml,omitempty"`
	UserDefinedXml string `json:"userDefinedXml,omitempty"`
	Success bool `json:"success,omitempty"`
}

