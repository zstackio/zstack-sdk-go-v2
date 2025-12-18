// Copyright (c) ZStack.io, Inc.

package param

// CreateVipDetailParam CreateVip detail param
type CreateVipDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
	IpRangeUuid string `json:"ipRangeUuid,omitempty"`
	RequiredIp string `json:"requiredIp,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVipParam CreateVip request param
type CreateVipParam struct {
	BaseParam
	Params CreateVipDetailParam `json:"params"`
}
