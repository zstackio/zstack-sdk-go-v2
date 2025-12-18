// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VirtualizerInfoInventoryView VirtualizerInfo
type VirtualizerInfoInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	InfoList []interface{} `json:"infoList,omitempty"`
}

