// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SSOTokenInventoryView SSOToken
type SSOTokenInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"clientUuid,omitempty"`
	rest string `json:"userUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

