// Copyright (c) ZStack.io, Inc.

package param

// ReconnectBaremetalPxeServerDetailParam ReconnectBaremetalPxeServer详细参数
type ReconnectBaremetalPxeServerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectBaremetalPxeServerParam ReconnectBaremetalPxeServer请求参数
type ReconnectBaremetalPxeServerParam struct {
	BaseParam
	Params ReconnectBaremetalPxeServerDetailParam `json:"params"` // 详细参数
}

