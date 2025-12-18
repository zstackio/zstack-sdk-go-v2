// Copyright (c) ZStack.io, Inc.

package param

// RemoveSNSDingTalkAtPersonDetailParam RemoveSNSDingTalkAtPerson detail param
type RemoveSNSDingTalkAtPersonDetailParam struct {
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveSNSDingTalkAtPersonParam RemoveSNSDingTalkAtPerson request param
type RemoveSNSDingTalkAtPersonParam struct {
	BaseParam
	Params RemoveSNSDingTalkAtPersonDetailParam `json:"params"`
}
