// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VirtualRouterMetadataInventoryView VirtualRouterMetadata
type VirtualRouterMetadataInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"zvrVersion,omitempty"`
	rest string `json:"vyosVersion,omitempty"`
	rest string `json:"kernelVersion,omitempty"`
}

