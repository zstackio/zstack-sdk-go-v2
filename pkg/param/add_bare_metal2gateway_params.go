// Copyright (c) ZStack.io, Inc.

package param

// AddBareMetal2GatewayDetailParam AddBareMetal2Gateway detail param
type AddBareMetal2GatewayDetailParam struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	SshPort int `json:"sshPort,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBareMetal2GatewayParam AddBareMetal2Gateway request param
type AddBareMetal2GatewayParam struct {
	BaseParam
	Params AddBareMetal2GatewayDetailParam `json:"params"`
}
