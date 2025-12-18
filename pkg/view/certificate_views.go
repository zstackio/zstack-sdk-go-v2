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
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

