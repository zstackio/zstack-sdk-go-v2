// Copyright (c) ZStack.io, Inc.

package param

// AddSNSWeComAtPersonDetailParam AddSNSWeComAtPerson detail param
type AddSNSWeComAtPersonDetailParam struct {
	UserId string `json:"userId" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	Remark string `json:"remark,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSNSWeComAtPersonParam AddSNSWeComAtPerson request param
type AddSNSWeComAtPersonParam struct {
	BaseParam
	Params AddSNSWeComAtPersonDetailParam `json:"params"`
}
