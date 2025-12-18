// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OAuth2TokenInventoryView OAuth2Token
type OAuth2TokenInventoryView struct {
	AccessToken string `json:"accessToken,omitempty"`
	IdToken string `json:"idToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	ClientUuid string `json:"clientUuid,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

