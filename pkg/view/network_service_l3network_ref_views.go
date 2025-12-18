// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NetworkServiceL3NetworkRefInventoryView NetworkServiceL3NetworkRef
type NetworkServiceL3NetworkRefInventoryView struct {
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	NetworkServiceProviderUuid string `json:"networkServiceProviderUuid,omitempty"`
	NetworkServiceType string `json:"networkServiceType,omitempty"`
}

