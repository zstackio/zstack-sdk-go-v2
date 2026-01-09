// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ConnectionRelationShipPropertyView ConnectionRelationShipProperty
type ConnectionRelationShipPropertyView struct {
	Uuid string `json:"uuid,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	ConnectionType string `json:"connectionType,omitempty"`
	Direction *string `json:"direction,omitempty"`
	RelationShips *string `json:"relationShips,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
}

