// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// QuotaInventoryView Quota
type QuotaInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	IdentityUuid *string `json:"identityUuid,omitempty"`
	IdentityType *string `json:"identityType,omitempty"`
	Value *int64 `json:"value,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
}

// QueryQuotaView QueryQuota
type QueryQuotaView struct {
	Inventories []QuotaInventoryView `json:"inventories,omitempty"`
}

// UpdateOrganizationQuotaEventView UpdateOrganizationQuotaEvent
type UpdateOrganizationQuotaEventView struct {
	Inventory QuotaInventoryView `json:"inventory,omitempty"`
}

// UpdateQuotaEventView UpdateQuotaEvent
type UpdateQuotaEventView struct {
	Inventory QuotaInventoryView `json:"inventory,omitempty"`
}

