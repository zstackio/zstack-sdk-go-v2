// Copyright (c) ZStack.io, Inc.

package param

// RemoveSNSWeComAtPersonDetailParam RemoveSNSWeComAtPerson detail param
type RemoveSNSWeComAtPersonDetailParam struct {
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	UserId string `json:"userId" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveSNSWeComAtPersonParam RemoveSNSWeComAtPerson request param
type RemoveSNSWeComAtPersonParam struct {
	BaseParam
	Params RemoveSNSWeComAtPersonDetailParam `json:"params"`
}
