// Copyright (c) ZStack.io, Inc.

package param

// CreateVmNicDetailParam CreateVmNic detail param
type CreateVmNicDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Ip string `json:"ip,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmNicParam CreateVmNic request param
type CreateVmNicParam struct {
	BaseParam
	Params CreateVmNicDetailParam `json:"params"`
}
