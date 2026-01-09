// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PriceInventoryView Price
type PriceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
	ResourceUnit *string `json:"resourceUnit,omitempty"`
	TimeUnit *string `json:"timeUnit,omitempty"`
	Price *float64 `json:"price,omitempty"`
	DateInLong *int64 `json:"dateInLong,omitempty"`
	EndDateInLong *int64 `json:"endDateInLong,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	TableUuid *string `json:"tableUuid,omitempty"`
	PciDeviceOfferings []PricePciDeviceOfferingRefInventoryView `json:"pciDeviceOfferings,omitempty"`
	BareMetal2VmOfferings []PriceBareMetal2ChassisOfferingRefInventoryView `json:"bareMetal2VmOfferings,omitempty"`
}

// UpdateResourcePriceEventView UpdateResourcePriceEvent
type UpdateResourcePriceEventView struct {
	Inventory PriceInventoryView `json:"inventory,omitempty"`
}

// QueryResourcePriceView QueryResourcePrice
type QueryResourcePriceView struct {
	Inventories []PriceInventoryView `json:"inventories,omitempty"`
}

// CreateResourcePriceEventView CreateResourcePriceEvent
type CreateResourcePriceEventView struct {
	Inventory PriceInventoryView `json:"inventory,omitempty"`
}

