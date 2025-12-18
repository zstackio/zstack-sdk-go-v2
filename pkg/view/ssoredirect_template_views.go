// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SSORedirectTemplateInventoryView SSORedirectTemplate
type SSORedirectTemplateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"clientUuid,omitempty"`
	rest string `json:"redirectTemplate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

