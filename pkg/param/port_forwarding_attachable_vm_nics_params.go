// Copyright (c) ZStack.io, Inc.

package param

// GetPortForwardingAttachableVmNicsDetailParam GetPortForwardingAttachableVmNics详细参数
type GetPortForwardingAttachableVmNicsDetailParam struct {
	rest string `json:"ruleUuid" validate:"required"` // 必填
}

// GetPortForwardingAttachableVmNicsParam GetPortForwardingAttachableVmNics请求参数
type GetPortForwardingAttachableVmNicsParam struct {
	BaseParam
	Params GetPortForwardingAttachableVmNicsDetailParam `json:"params"` // 详细参数
}

