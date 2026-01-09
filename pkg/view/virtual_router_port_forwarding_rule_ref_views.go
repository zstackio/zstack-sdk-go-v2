// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VirtualRouterPortForwardingRuleRefInventoryView VirtualRouterPortForwardingRuleRef
type VirtualRouterPortForwardingRuleRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VipUuid *string `json:"vipUuid,omitempty"`
	VirtualRouterVmUuid *string `json:"virtualRouterVmUuid,omitempty"`
}

