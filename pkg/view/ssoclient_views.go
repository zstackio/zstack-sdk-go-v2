// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SSOClientInventoryView SSOClient
type SSOClientInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description *string `json:"description,omitempty"`
	ClientType *string `json:"clientType,omitempty"`
	LoginType *string `json:"loginType,omitempty"`
	LoginMNUrl *string `json:"loginMNUrl,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	Attributes []SSOClientAttributeInventoryView `json:"attributes,omitempty"`
}

// DeleteSSOClientEventView DeleteSSOClientEvent
type DeleteSSOClientEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetSSOClientView GetSSOClient
type GetSSOClientView struct {
	Inventories []SSOClientInventoryView `json:"inventories,omitempty"`
}

