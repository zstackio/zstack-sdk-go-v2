// Copyright (c) ZStack.io, Inc.

package param

// AttachEipDetailParam AttachEip detail param
type AttachEipDetailParam struct {
	EipUuid string `json:"eipUuid" validate:"required"`
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	UsedIpUuid string `json:"usedIpUuid,omitempty"`
}

// AttachEipParam AttachEip request param
type AttachEipParam struct {
	BaseParam
	Params AttachEipDetailParam `json:"params"`
}
