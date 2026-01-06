// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LoadBalancerListenerCertificateRefInventoryView LoadBalancerListenerCertificateRef
type LoadBalancerListenerCertificateRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ListenerUuid string `json:"listenerUuid,omitempty"`
	CertificateUuid string `json:"certificateUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

