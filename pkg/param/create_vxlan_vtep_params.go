// Copyright (c) ZStack.io, Inc.

package param

// CreateVxlanVtepDetailParam CreateVxlanVtep detail param
type CreateVxlanVtepDetailParam struct {
	HostUuid string `json:"hostUuid" validate:"required"`
	PoolUuid string `json:"poolUuid" validate:"required"`
	VtepIp string `json:"vtepIp,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVxlanVtepParam CreateVxlanVtep request param
type CreateVxlanVtepParam struct {
	BaseParam
	Params CreateVxlanVtepDetailParam `json:"params"`
}
