// Copyright (c) ZStack.io, Inc.

package param

// StopBaremetalPxeServerDetailParam StopBaremetalPxeServer详细参数
type StopBaremetalPxeServerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// StopBaremetalPxeServerParam StopBaremetalPxeServer请求参数
type StopBaremetalPxeServerParam struct {
	BaseParam
	Params StopBaremetalPxeServerDetailParam `json:"params"` // 详细参数
}

