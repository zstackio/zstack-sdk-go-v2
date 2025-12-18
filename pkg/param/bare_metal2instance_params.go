// Copyright (c) ZStack.io, Inc.

package param

// StartBareMetal2InstanceDetailParam StartBareMetal2Instance详细参数
type StartBareMetal2InstanceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"gatewayUuid,omitempty"`
	rest string `json:"chassisUuid,omitempty"`
	rest string `json:"chassisOfferingUuid,omitempty"`
}

// StartBareMetal2InstanceParam StartBareMetal2Instance请求参数
type StartBareMetal2InstanceParam struct {
	BaseParam
	Params StartBareMetal2InstanceDetailParam `json:"params"` // 详细参数
}

