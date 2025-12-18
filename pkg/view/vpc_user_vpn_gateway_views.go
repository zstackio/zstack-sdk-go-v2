// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcUserVpnGatewayInventoryView VpcUserVpnGateway
type VpcUserVpnGatewayInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AccountName string `json:"accountName,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	Type string `json:"type,omitempty"`
	GatewayId string `json:"gatewayId,omitempty"`
	Ip string `json:"ip,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

