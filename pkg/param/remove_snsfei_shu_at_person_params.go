// Copyright (c) ZStack.io, Inc.

package param

// RemoveSNSFeiShuAtPersonDetailParam RemoveSNSFeiShuAtPerson detail param
type RemoveSNSFeiShuAtPersonDetailParam struct {
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	UserId string `json:"userId" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveSNSFeiShuAtPersonParam RemoveSNSFeiShuAtPerson request param
type RemoveSNSFeiShuAtPersonParam struct {
	BaseParam
	Params RemoveSNSFeiShuAtPersonDetailParam `json:"params"`
}
