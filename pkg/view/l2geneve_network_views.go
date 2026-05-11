// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L2GeneveNetworkInventoryView L2GeneveNetwork
type L2GeneveNetworkInventoryView struct {
	BaseInfoView
	BaseTimeView
	GeneveId int `json:"geneveId,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	VirtualNetworkId int `json:"virtualNetworkId,omitempty"`
	Isolated bool `json:"isolated,omitempty"`
	Pvlan string `json:"pvlan,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

// CreateL2NetworkEventView CreateL2NetworkEvent
type CreateL2NetworkEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

