// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SSOClientAttributeInventoryView SSOClientAttribute
type SSOClientAttributeInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"value,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"purpose,omitempty"`
	rest string `json:"ssoClientUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

