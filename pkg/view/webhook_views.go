// Copyright (c) ZStack.io, Inc.

package view

import "time"

// WebhookInventoryView Webhook
type WebhookInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"opaque,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

