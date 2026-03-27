// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// ReconnectBareMetal2GatewayParamDetail ReconnectBareMetal2Gateway detail param
type ReconnectBareMetal2GatewayParamDetail struct {
}

// ReconnectBareMetal2GatewayParam ReconnectBareMetal2Gateway request param
type ReconnectBareMetal2GatewayParam struct {
	BaseParam
	Params ReconnectBareMetal2GatewayParamDetail `json:"reconnectBareMetal2Gateway"`
}
// UpdateBareMetal2GatewayParamDetail UpdateBareMetal2Gateway detail param
type UpdateBareMetal2GatewayParamDetail struct {
	SshPort *int `json:"sshPort,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
}

// UpdateBareMetal2GatewayParam UpdateBareMetal2Gateway request param
type UpdateBareMetal2GatewayParam struct {
	BaseParam
	Params UpdateBareMetal2GatewayParamDetail `json:"updateBareMetal2Gateway"`
}
// DeleteBareMetal2GatewayParamDetail DeleteBareMetal2Gateway detail param
type DeleteBareMetal2GatewayParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteBareMetal2GatewayParam DeleteBareMetal2Gateway request param
type DeleteBareMetal2GatewayParam struct {
	BaseParam
	Params DeleteBareMetal2GatewayParamDetail `json:"deleteBareMetal2Gateway"`
}
// AddBareMetal2GatewayParamDetail AddBareMetal2Gateway detail param
type AddBareMetal2GatewayParamDetail struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	SshPort *int `json:"sshPort,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBareMetal2GatewayParam AddBareMetal2Gateway request param
type AddBareMetal2GatewayParam struct {
	BaseParam
	Params AddBareMetal2GatewayParamDetail `json:"params"`
}
