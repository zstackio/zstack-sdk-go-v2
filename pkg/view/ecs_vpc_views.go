// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EcsVpcInventoryView EcsVpc
type EcsVpcInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"ecsVpcId,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"deleted,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"cidrBlock,omitempty"`
	rest string `json:"vRouterId,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

