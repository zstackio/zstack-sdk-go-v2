// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IdentityZonePropertyView IdentityZoneProperty
type IdentityZonePropertyView struct {
	ZoneId                    string   `json:"zoneId,omitempty"`
	LocalName                 string   `json:"localName,omitempty"`
	AvailableInstanceTypes    []string `json:"availableInstanceTypes,omitempty"`
	AvailableResourceCreation []string `json:"availableResourceCreation,omitempty"`
	AvailableDiskCategories   []string `json:"availableDiskCategories,omitempty"`
}
