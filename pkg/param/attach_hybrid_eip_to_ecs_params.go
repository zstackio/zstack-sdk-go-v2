// Copyright (c) ZStack.io, Inc.

package param

// AttachHybridEipToEcsDetailParam AttachHybridEipToEcs detail param
type AttachHybridEipToEcsDetailParam struct {
	EipUuid string `json:"eipUuid" validate:"required"`
	EcsUuid string `json:"ecsUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
}

// AttachHybridEipToEcsParam AttachHybridEipToEcs request param
type AttachHybridEipToEcsParam struct {
	BaseParam
	Params AttachHybridEipToEcsDetailParam `json:"params"`
}
