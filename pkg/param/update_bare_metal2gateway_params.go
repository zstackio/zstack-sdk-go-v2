// Copyright (c) ZStack.io, Inc.

package param

// UpdateBareMetal2GatewayDetailParam UpdateBareMetal2Gateway detail param
type UpdateBareMetal2GatewayDetailParam struct {
	SshPort int `json:"sshPort,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
}

// UpdateBareMetal2GatewayParam UpdateBareMetal2Gateway request param
type UpdateBareMetal2GatewayParam struct {
	BaseParam
	Params UpdateBareMetal2GatewayDetailParam `json:"params"`
}
