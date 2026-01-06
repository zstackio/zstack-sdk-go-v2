// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SessionInventoryView Session
type SessionInventoryView struct {
	Uuid        string    `json:"uuid,omitempty"`
	AccountUuid string    `json:"accountUuid,omitempty"`
	UserUuid    string    `json:"userUuid,omitempty"`
	UserType    string    `json:"userType,omitempty"`
	ExpiredDate time.Time `json:"expiredDate,omitempty"`
	CreateDate  time.Time `json:"createDate,omitempty"`
}

type WebUISessionView struct {
	SessionId       string `json:"sessionId"`   // Resource UUID
	AccountUuid     string `json:"accountUuid"` // Account UUID
	UserUuid        string `json:"userUuid"`    // User UUID
	UserName        string `json:"username"`    // Username
	LoginType       string `json:"loginType"`
	CurrentIdentity string `json:"currentIdentity"`
	ZSVersion       string `json:"zsVersion"` // ZStack Cloud detailed version
}
