// Copyright (c) ZStack.io, Inc.

package param

// SNSSnmpTestConnectionDetailParam SNSSnmpTestConnection详细参数
type SNSSnmpTestConnectionDetailParam struct {
	rest string `json:"platformUuid,omitempty"`
	rest string `json:"endpointUuid,omitempty"`
}

// SNSSnmpTestConnectionParam SNSSnmpTestConnection请求参数
type SNSSnmpTestConnectionParam struct {
	BaseParam
	Params SNSSnmpTestConnectionDetailParam `json:"params"` // 详细参数
}

