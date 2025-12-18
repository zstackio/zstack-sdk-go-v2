// Copyright (c) ZStack.io, Inc.

package param

// CreateSnmpAgentDetailParam CreateSnmpAgent detail param
type CreateSnmpAgentDetailParam struct {
	Version string `json:"version" validate:"required"`
	ReadCommunity string `json:"readCommunity,omitempty"`
	UserName string `json:"userName,omitempty"`
	AuthAlgorithm string `json:"authAlgorithm,omitempty"`
	AuthPassword string `json:"authPassword,omitempty"`
	PrivacyAlgorithm string `json:"privacyAlgorithm,omitempty"`
	PrivacyPassword string `json:"privacyPassword,omitempty"`
	Port int `json:"port" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSnmpAgentParam CreateSnmpAgent request param
type CreateSnmpAgentParam struct {
	BaseParam
	Params CreateSnmpAgentDetailParam `json:"params"`
}
