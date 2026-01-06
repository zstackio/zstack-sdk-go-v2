// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TwoFactorAuthenticationInventoryView TwoFactorAuthentication
type TwoFactorAuthenticationInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Secret string `json:"secret,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
	UserType string `json:"userType,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryTwoFactorAuthenticationView QueryTwoFactorAuthentication
type QueryTwoFactorAuthenticationView struct {
	Inventories []TwoFactorAuthenticationInventoryView `json:"inventories,omitempty"`
}

