// Copyright (c) ZStack.io, Inc.

package param

// StopSnmpAgentDetailParam StopSnmpAgent详细参数
type StopSnmpAgentDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// StopSnmpAgentParam StopSnmpAgent请求参数
type StopSnmpAgentParam struct {
	BaseParam
	Params StopSnmpAgentDetailParam `json:"params"` // 详细参数
}

