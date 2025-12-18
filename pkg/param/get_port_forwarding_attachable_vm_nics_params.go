// Copyright (c) ZStack.io, Inc.

package param

// GetPortForwardingAttachableVmNicsDetailParam GetPortForwardingAttachableVmNics detail param
type GetPortForwardingAttachableVmNicsDetailParam struct {
	RuleUuid string `json:"ruleUuid" validate:"required"`
}

// GetPortForwardingAttachableVmNicsParam GetPortForwardingAttachableVmNics request param
type GetPortForwardingAttachableVmNicsParam struct {
	BaseParam
	Params GetPortForwardingAttachableVmNicsDetailParam `json:"params"`
}
