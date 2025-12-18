// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmNicSecurityPolicyInventoryView VmNicSecurityPolicy
type VmNicSecurityPolicyInventoryView struct {
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	IngressPolicy string `json:"ingressPolicy,omitempty"`
	EgressPolicy string `json:"egressPolicy,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

