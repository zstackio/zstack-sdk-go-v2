// Copyright (c) ZStack.io, Inc.

package param

// StopSnmpAgentDetailParam StopSnmpAgent detail param
type StopSnmpAgentDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StopSnmpAgentParam StopSnmpAgent request param
type StopSnmpAgentParam struct {
	BaseParam
	Params StopSnmpAgentDetailParam `json:"params"`
}
