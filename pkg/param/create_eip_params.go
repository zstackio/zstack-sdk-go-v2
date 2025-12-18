// Copyright (c) ZStack.io, Inc.

package param

// CreateEipDetailParam CreateEip detail param
type CreateEipDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VipUuid string `json:"vipUuid" validate:"required"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	UsedIpUuid string `json:"usedIpUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEipParam CreateEip request param
type CreateEipParam struct {
	BaseParam
	Params CreateEipDetailParam `json:"params"`
}
