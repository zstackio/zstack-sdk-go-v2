// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SSOServerTokenInventoryView SSOServerToken
type SSOServerTokenInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"accessToken,omitempty"`
	rest string `json:"idToken,omitempty"`
	rest string `json:"refreshToken,omitempty"`
	rest string `json:"userUuid,omitempty"`
	rest string `json:"sessionUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

