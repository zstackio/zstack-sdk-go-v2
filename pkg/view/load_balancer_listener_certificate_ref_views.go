// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LoadBalancerListenerCertificateRefInventoryView LoadBalancerListenerCertificateRef
type LoadBalancerListenerCertificateRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"listenerUuid,omitempty"`
	rest string `json:"certificateUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

