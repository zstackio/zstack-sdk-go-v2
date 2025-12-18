// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmNicSecurityPolicyInventoryView VmNicSecurityPolicy
type VmNicSecurityPolicyInventoryView struct {
	rest string `json:"vmNicUuid,omitempty"`
	rest string `json:"ingressPolicy,omitempty"`
	rest string `json:"egressPolicy,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

