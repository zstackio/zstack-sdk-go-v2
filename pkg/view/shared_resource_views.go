// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SharedResourceInventoryView SharedResource
type SharedResourceInventoryView struct {
	BaseInfoView
	BaseTimeView
	OwnerAccountUuid string `json:"ownerAccountUuid,omitempty"`
	ReceiverAccountUuid string `json:"receiverAccountUuid,omitempty"`
	ToPublic bool `json:"toPublic,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
}

// QuerySharedResourceView QuerySharedResource
type QuerySharedResourceView struct {
	Inventories []SharedResourceInventoryView `json:"inventories,omitempty"`
}

