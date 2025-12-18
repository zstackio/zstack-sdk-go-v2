// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcSnatStateInventoryView VpcSnatState
type VpcSnatStateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vpcUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

