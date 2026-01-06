// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TwoFactorAuthenticationSecretInventoryView TwoFactorAuthenticationSecret
type TwoFactorAuthenticationSecretInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Secret string `json:"secret,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
	UserType string `json:"userType,omitempty"`
}

// GetTwoFactorAuthenticationSecretView GetTwoFactorAuthenticationSecret
type GetTwoFactorAuthenticationSecretView struct {
	Inventory TwoFactorAuthenticationSecretInventoryView `json:"inventory,omitempty"`
}

// ResetTwoFactorAuthenticationSecretEventView ResetTwoFactorAuthenticationSecretEvent
type ResetTwoFactorAuthenticationSecretEventView struct {
	Inventory TwoFactorAuthenticationSecretInventoryView `json:"inventory,omitempty"`
}

