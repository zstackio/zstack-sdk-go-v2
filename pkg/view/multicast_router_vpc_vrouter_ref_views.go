// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MulticastRouterVpcVRouterRefInventoryView MulticastRouterVpcVRouterRef
type MulticastRouterVpcVRouterRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VpcRouterUuid string `json:"vpcRouterUuid,omitempty"`
}

