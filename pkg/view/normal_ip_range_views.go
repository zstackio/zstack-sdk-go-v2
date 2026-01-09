// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NormalIpRangeInventoryView NormalIpRange
type NormalIpRangeInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	L3NetworkUuid *string `json:"l3NetworkUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	StartIp *string `json:"startIp,omitempty"`
	EndIp *string `json:"endIp,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	NetworkCidr *string `json:"networkCidr,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
	AddressMode *string `json:"addressMode,omitempty"`
	PrefixLen *int `json:"prefixLen,omitempty"`
	IpRangeType string `json:"ipRangeType,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

