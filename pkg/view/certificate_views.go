// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// CertificateInventoryView Certificate
type CertificateInventoryView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	Description string `json:"description,omitempty"`
	Listeners []LoadBalancerListenerCertificateRefInventoryView `json:"listeners,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// UpdateCertificateEventView UpdateCertificateEvent
type UpdateCertificateEventView struct {
	Inventory CertificateInventoryView `json:"inventory,omitempty"`
}

// CreateCertificateEventView CreateCertificateEvent
type CreateCertificateEventView struct {
	Inventory CertificateInventoryView `json:"inventory,omitempty"`
}

// DeleteCertificateEventView DeleteCertificateEvent
type DeleteCertificateEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryCertificateView QueryCertificate
type QueryCertificateView struct {
	Inventories []CertificateInventoryView `json:"inventories,omitempty"`
}

