// Copyright (c) ZStack.io, Inc.

package view

import "time"

// OAuth2TokenInventoryView OAuth2Token
type OAuth2TokenInventoryView struct {
	rest string `json:"accessToken,omitempty"`
	rest string `json:"idToken,omitempty"`
	rest string `json:"refreshToken,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"clientUuid,omitempty"`
	rest string `json:"userUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

