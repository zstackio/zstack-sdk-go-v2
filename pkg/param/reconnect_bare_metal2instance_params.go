// Copyright (c) ZStack.io, Inc.

package param

// ReconnectBareMetal2InstanceDetailParam ReconnectBareMetal2Instance详细参数
type ReconnectBareMetal2InstanceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectBareMetal2InstanceParam ReconnectBareMetal2Instance请求参数
type ReconnectBareMetal2InstanceParam struct {
	BaseParam
	Params ReconnectBareMetal2InstanceDetailParam `json:"params"` // 详细参数
}

