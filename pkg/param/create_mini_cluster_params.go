// Copyright (c) ZStack.io, Inc.

package param

// CreateMiniClusterDetailParam CreateMiniCluster detail param
type CreateMiniClusterDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	HostManagementIps []string `json:"hostManagementIps" validate:"required"`
	Username string `json:"username,omitempty"`
	Password string `json:"password" validate:"required"`
	SshPort int `json:"sshPort,omitempty"`
	Description string `json:"description,omitempty"`
	HypervisorType string `json:"hypervisorType" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMiniClusterParam CreateMiniCluster request param
type CreateMiniClusterParam struct {
	BaseParam
	Params CreateMiniClusterDetailParam `json:"params"`
}
