// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SSOClientInventoryView SSOClient
type SSOClientInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"clientType,omitempty"`
	rest string `json:"loginType,omitempty"`
	rest string `json:"loginMNUrl,omitempty"`
	rest string `json:"redirectUrl,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest []SSOClientAttributeInventoryView `json:"attributes,omitempty"`
}

