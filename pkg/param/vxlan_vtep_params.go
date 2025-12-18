// Copyright (c) ZStack.io, Inc.

package param

// CreateVxlanVtepDetailParam CreateVxlanVtep详细参数
type CreateVxlanVtepDetailParam struct {
	rest string `json:"hostUuid" validate:"required"` // 必填
	rest string `json:"poolUuid" validate:"required"` // 必填
	rest string `json:"vtepIp,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVxlanVtepParam CreateVxlanVtep请求参数
type CreateVxlanVtepParam struct {
	BaseParam
	Params CreateVxlanVtepDetailParam `json:"params"` // 详细参数
}

