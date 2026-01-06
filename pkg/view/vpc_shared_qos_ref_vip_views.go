// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcSharedQosRefVipInventoryView VpcSharedQosRefVip
type VpcSharedQosRefVipInventoryView struct {
	Id int64 `json:"id,omitempty"`
	SharedQosUuid string `json:"sharedQosUuid,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

