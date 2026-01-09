// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcHaGroupApplianceVmRefInventoryView VpcHaGroupApplianceVmRef
type VpcHaGroupApplianceVmRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VpcHaRouterUuid *string `json:"vpcHaRouterUuid,omitempty"`
}

