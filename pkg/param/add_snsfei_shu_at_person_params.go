// Copyright (c) ZStack.io, Inc.

package param

// AddSNSFeiShuAtPersonDetailParam AddSNSFeiShuAtPerson detail param
type AddSNSFeiShuAtPersonDetailParam struct {
	UserId string `json:"userId" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	Remark string `json:"remark,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSNSFeiShuAtPersonParam AddSNSFeiShuAtPerson request param
type AddSNSFeiShuAtPersonParam struct {
	BaseParam
	Params AddSNSFeiShuAtPersonDetailParam `json:"params"`
}
