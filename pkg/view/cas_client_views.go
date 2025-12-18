// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CasClientInventoryView CasClient
type CasClientInventoryView struct {
	rest string `json:"casServerLoginUrl,omitempty"`
	rest string `json:"casServerUrlPrefix,omitempty"`
	rest string `json:"serverName,omitempty"`
	rest string `json:"state,omitempty"`
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

