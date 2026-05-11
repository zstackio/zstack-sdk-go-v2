// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZnsTenantInventoryView ZnsTenant
type ZnsTenantInventoryView struct {
	BaseInfoView
	BaseTimeView
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	ZnsResourceUuid string `json:"znsResourceUuid,omitempty"`
	Description string `json:"description,omitempty"`
}

// QueryZnsTenantView QueryZnsTenant
type QueryZnsTenantView struct {
	Inventories []ZnsTenantInventoryView `json:"inventories,omitempty"`
}

