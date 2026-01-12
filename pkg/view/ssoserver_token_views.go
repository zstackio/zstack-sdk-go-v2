// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SSOServerTokenInventoryView SSOServerToken
type SSOServerTokenInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccessToken *string `json:"accessToken,omitempty"`
	IdToken *string `json:"idToken,omitempty"`
	RefreshToken *string `json:"refreshToken,omitempty"`
	UserUuid *string `json:"userUuid,omitempty"`
	SessionUuid *string `json:"sessionUuid,omitempty"`
}

// RefreshSSOServerTokenEventView RefreshSSOServerTokenEvent
type RefreshSSOServerTokenEventView struct {
	Inventory SSOServerTokenInventoryView `json:"inventory,omitempty"`
}

