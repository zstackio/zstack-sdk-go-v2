// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LoadBalancerServerGroupServerIpInventoryView LoadBalancerServerGroupServerIp
type LoadBalancerServerGroupServerIpInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
	IpAddress string `json:"ipAddress,omitempty"`
	Weight int64 `json:"weight,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

