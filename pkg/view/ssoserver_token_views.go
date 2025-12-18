// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SSOServerTokenInventoryView SSOServerToken
type SSOServerTokenInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
	IdToken string `json:"idToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
	SessionUuid string `json:"sessionUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

