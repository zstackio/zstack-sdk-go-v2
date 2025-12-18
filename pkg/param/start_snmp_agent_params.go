// Copyright (c) ZStack.io, Inc.

package param

// StartSnmpAgentDetailParam StartSnmpAgent detail param
type StartSnmpAgentDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StartSnmpAgentParam StartSnmpAgent request param
type StartSnmpAgentParam struct {
	BaseParam
	Params StartSnmpAgentDetailParam `json:"params"`
}
