// Copyright (c) ZStack.io, Inc.

package param

// DetachHybridEipFromEcsDetailParam DetachHybridEipFromEcs detail param
type DetachHybridEipFromEcsDetailParam struct {
	EipUuid string `json:"eipUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
}

// DetachHybridEipFromEcsParam DetachHybridEipFromEcs request param
type DetachHybridEipFromEcsParam struct {
	BaseParam
	Params DetachHybridEipFromEcsDetailParam `json:"params"`
}
