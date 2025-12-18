// Copyright (c) ZStack.io, Inc.

package view

import "time"

// TwoFactorAuthenticationSecretInventoryView TwoFactorAuthenticationSecret
type TwoFactorAuthenticationSecretInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"secret,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"userUuid,omitempty"`
	rest string `json:"userType,omitempty"`
}

