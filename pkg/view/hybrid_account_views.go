// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HybridAccountInventoryView HybridAccount
type HybridAccountInventoryView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Akey string `json:"akey,omitempty"`
	HybridAccountId string `json:"hybridAccountId,omitempty"`
	HybridUserId string `json:"hybridUserId,omitempty"`
	HybridUserName string `json:"hybridUserName,omitempty"`
	Current string `json:"current,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

