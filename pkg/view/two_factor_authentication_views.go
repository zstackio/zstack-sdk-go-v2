// Copyright (c) ZStack.io, Inc.

package view

import "time"

// TwoFactorAuthenticationInventoryView TwoFactorAuthentication
type TwoFactorAuthenticationInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"secret,omitempty"`
	rest string `json:"userUuid,omitempty"`
	rest string `json:"userType,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

