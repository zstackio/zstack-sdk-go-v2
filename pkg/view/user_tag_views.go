// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserTagInventoryView UserTag
type UserTagInventoryView struct {
	TagPatternUuid string `json:"tagPatternUuid,omitempty"`
	TagPattern TagPatternInventoryView `json:"tagPattern,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Tag string `json:"tag,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// CreateUserTagEventView CreateUserTagEvent
type CreateUserTagEventView struct {
	Inventory UserTagInventoryView `json:"inventory,omitempty"`
}

// QueryUserTagView QueryUserTag
type QueryUserTagView struct {
	Inventories []UserTagInventoryView `json:"inventories,omitempty"`
}

