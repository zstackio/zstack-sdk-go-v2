// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ResourceInventoryView Resource
type ResourceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ConcreteResourceType *string `json:"concreteResourceType,omitempty"`
}

// GetResourceNamesView GetResourceNames
type GetResourceNamesView struct {
	Inventories []ResourceInventoryView `json:"inventories,omitempty"`
}

