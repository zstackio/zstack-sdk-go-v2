// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZnsTenantRouterInventoryView ZnsTenantRouter
type ZnsTenantRouterInventoryView struct {
	BaseInfoView
	BaseTimeView
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	TenantUuid string `json:"tenantUuid,omitempty"`
	ZnsResourceUuid string `json:"znsResourceUuid,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
}

// QueryZnsTenantRouterView QueryZnsTenantRouter
type QueryZnsTenantRouterView struct {
	Inventories []ZnsTenantRouterInventoryView `json:"inventories,omitempty"`
}

