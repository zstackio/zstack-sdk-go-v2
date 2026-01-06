// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcSnatStateInventoryView VpcSnatState
type VpcSnatStateInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VpcUuid string `json:"vpcUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryVpcSnatStateView QueryVpcSnatState
type QueryVpcSnatStateView struct {
	Inventories []VpcSnatStateInventoryView `json:"inventories,omitempty"`
}

