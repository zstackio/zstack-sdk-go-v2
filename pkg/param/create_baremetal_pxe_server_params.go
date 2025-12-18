// Copyright (c) ZStack.io, Inc.

package param

// CreateBaremetalPxeServerDetailParam CreateBaremetalPxeServer detail param
type CreateBaremetalPxeServerDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Hostname string `json:"hostname" validate:"required"`
	SshUsername string `json:"sshUsername" validate:"required"`
	SshPassword string `json:"sshPassword" validate:"required"`
	SshPort int `json:"sshPort,omitempty"`
	StoragePath string `json:"storagePath" validate:"required"`
	DhcpInterface string `json:"dhcpInterface" validate:"required"`
	DhcpRangeBegin string `json:"dhcpRangeBegin,omitempty"`
	DhcpRangeEnd string `json:"dhcpRangeEnd,omitempty"`
	DhcpRangeNetmask string `json:"dhcpRangeNetmask,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBaremetalPxeServerParam CreateBaremetalPxeServer request param
type CreateBaremetalPxeServerParam struct {
	BaseParam
	Params CreateBaremetalPxeServerDetailParam `json:"params"`
}
