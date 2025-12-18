// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CertificateInventoryView Certificate
type CertificateInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"certificate,omitempty"`
	rest string `json:"description,omitempty"`
	rest []LoadBalancerListenerCertificateRefInventoryView `json:"listeners,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

