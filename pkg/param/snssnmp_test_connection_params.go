// Copyright (c) ZStack.io, Inc.

package param

// SNSSnmpTestConnectionDetailParam SNSSnmpTestConnection detail param
type SNSSnmpTestConnectionDetailParam struct {
	PlatformUuid string `json:"platformUuid,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
}

// SNSSnmpTestConnectionParam SNSSnmpTestConnection request param
type SNSSnmpTestConnectionParam struct {
	BaseParam
	Params SNSSnmpTestConnectionDetailParam `json:"params"`
}
