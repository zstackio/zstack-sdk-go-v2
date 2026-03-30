// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ExternalTenantResourceRefInventoryView ExternalTenantResourceRef
type ExternalTenantResourceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	Source string `json:"source,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
	UserId string `json:"userId,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

