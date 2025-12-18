// Copyright (c) ZStack.io, Inc.

package param

// AddBareMetal2GatewayDetailParam AddBareMetal2Gateway详细参数
type AddBareMetal2GatewayDetailParam struct {
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest int `json:"sshPort,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddBareMetal2GatewayParam AddBareMetal2Gateway请求参数
type AddBareMetal2GatewayParam struct {
	BaseParam
	Params AddBareMetal2GatewayDetailParam `json:"params"` // 详细参数
}

