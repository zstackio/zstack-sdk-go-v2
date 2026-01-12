// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// CasClientInventoryView CasClient
type CasClientInventoryView struct {
	BaseInfoView
	BaseTimeView
	CasServerLoginUrl *string `json:"casServerLoginUrl,omitempty"`
	CasServerUrlPrefix *string `json:"casServerUrlPrefix,omitempty"`
	ServerName *string `json:"serverName,omitempty"`
	State string `json:"state,omitempty"`
	Description *string `json:"description,omitempty"`
	ClientType *string `json:"clientType,omitempty"`
	LoginType *string `json:"loginType,omitempty"`
	LoginMNUrl *string `json:"loginMNUrl,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	Attributes []SSOClientAttributeInventoryView `json:"attributes,omitempty"`
}

// CreateCasClientEventView CreateCasClientEvent
type CreateCasClientEventView struct {
	Inventory CasClientInventoryView `json:"inventory,omitempty"`
}

// UpdateCasClientEventView UpdateCasClientEvent
type UpdateCasClientEventView struct {
	Inventory CasClientInventoryView `json:"inventory,omitempty"`
}

