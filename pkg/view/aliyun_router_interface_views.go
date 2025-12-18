// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AliyunRouterInterfaceInventoryView AliyunRouterInterface
type AliyunRouterInterfaceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest string `json:"routerInterfaceId,omitempty"`
	rest string `json:"virtualRouterUuid,omitempty"`
	rest string `json:"accessPointUuid,omitempty"`
	rest string `json:"role,omitempty"`
	rest string `json:"vRouterType,omitempty"`
	rest string `json:"spec,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"oppositeInterfaceUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

