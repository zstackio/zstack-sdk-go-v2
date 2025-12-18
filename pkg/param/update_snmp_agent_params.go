// Copyright (c) ZStack.io, Inc.

package param

// UpdateSnmpAgentDetailParam UpdateSnmpAgent detail param
type UpdateSnmpAgentDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Version string `json:"version" validate:"required"`
	ReadCommunity string `json:"readCommunity,omitempty"`
	UserName string `json:"userName,omitempty"`
	AuthAlgorithm string `json:"authAlgorithm,omitempty"`
	AuthPassword string `json:"authPassword,omitempty"`
	PrivacyAlgorithm string `json:"privacyAlgorithm,omitempty"`
	PrivacyPassword string `json:"privacyPassword,omitempty"`
	Port int `json:"port" validate:"required"`
}

// UpdateSnmpAgentParam UpdateSnmpAgent request param
type UpdateSnmpAgentParam struct {
	BaseParam
	Params UpdateSnmpAgentDetailParam `json:"params"`
}
