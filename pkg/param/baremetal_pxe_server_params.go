// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteBaremetalPxeServerParamDetail DeleteBaremetalPxeServer detail param
type DeleteBaremetalPxeServerParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteBaremetalPxeServerParam DeleteBaremetalPxeServer request param
type DeleteBaremetalPxeServerParam struct {
	BaseParam
	Params DeleteBaremetalPxeServerParamDetail `json:"deleteBaremetalPxeServer"`
}
// UpdateBaremetalPxeServerParamDetail UpdateBaremetalPxeServer detail param
type UpdateBaremetalPxeServerParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	DhcpRangeBegin *string `json:"dhcpRangeBegin,omitempty"`
	DhcpRangeEnd *string `json:"dhcpRangeEnd,omitempty"`
	DhcpRangeNetmask *string `json:"dhcpRangeNetmask,omitempty"`
}

// UpdateBaremetalPxeServerParam UpdateBaremetalPxeServer request param
type UpdateBaremetalPxeServerParam struct {
	BaseParam
	Params UpdateBaremetalPxeServerParamDetail `json:"updateBaremetalPxeServer"`
}
// StartBaremetalPxeServerParamDetail StartBaremetalPxeServer detail param
type StartBaremetalPxeServerParamDetail struct {
}

// StartBaremetalPxeServerParam StartBaremetalPxeServer request param
type StartBaremetalPxeServerParam struct {
	BaseParam
	Params StartBaremetalPxeServerParamDetail `json:"startBaremetalPxeServer"`
}
// ReconnectBaremetalPxeServerParamDetail ReconnectBaremetalPxeServer detail param
type ReconnectBaremetalPxeServerParamDetail struct {
}

// ReconnectBaremetalPxeServerParam ReconnectBaremetalPxeServer request param
type ReconnectBaremetalPxeServerParam struct {
	BaseParam
	Params ReconnectBaremetalPxeServerParamDetail `json:"reconnectBaremetalPxeServer"`
}
// StopBaremetalPxeServerParamDetail StopBaremetalPxeServer detail param
type StopBaremetalPxeServerParamDetail struct {
}

// StopBaremetalPxeServerParam StopBaremetalPxeServer request param
type StopBaremetalPxeServerParam struct {
	BaseParam
	Params StopBaremetalPxeServerParamDetail `json:"stopBaremetalPxeServer"`
}
// CreateBaremetalPxeServerParamDetail CreateBaremetalPxeServer detail param
type CreateBaremetalPxeServerParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Hostname string `json:"hostname" validate:"required"`
	SshUsername string `json:"sshUsername" validate:"required"`
	SshPassword string `json:"sshPassword" validate:"required"`
	SshPort *int `json:"sshPort,omitempty"`
	StoragePath string `json:"storagePath" validate:"required"`
	DhcpInterface string `json:"dhcpInterface" validate:"required"`
	DhcpRangeBegin *string `json:"dhcpRangeBegin,omitempty"`
	DhcpRangeEnd *string `json:"dhcpRangeEnd,omitempty"`
	DhcpRangeNetmask *string `json:"dhcpRangeNetmask,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBaremetalPxeServerParam CreateBaremetalPxeServer request param
type CreateBaremetalPxeServerParam struct {
	BaseParam
	Params CreateBaremetalPxeServerParamDetail `json:"params"`
}
