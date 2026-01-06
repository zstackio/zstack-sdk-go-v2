// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcHaGroupNetworkServiceRefInventoryView VpcHaGroupNetworkServiceRef
type VpcHaGroupNetworkServiceRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	VpcHaRouterUuid string `json:"vpcHaRouterUuid,omitempty"`
	NetworkServiceName string `json:"networkServiceName,omitempty"`
	NetworkServiceUuid string `json:"networkServiceUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryVpcHaGroupNetworkServiceRefView QueryVpcHaGroupNetworkServiceRef
type QueryVpcHaGroupNetworkServiceRefView struct {
	Inventories []VpcHaGroupNetworkServiceRefInventoryView `json:"inventories,omitempty"`
}

