// Copyright (c) ZStack.io, Inc.

package param

// AddSNSDingTalkAtPersonDetailParam AddSNSDingTalkAtPerson detail param
type AddSNSDingTalkAtPersonDetailParam struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	Remark string `json:"remark,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSNSDingTalkAtPersonParam AddSNSDingTalkAtPerson request param
type AddSNSDingTalkAtPersonParam struct {
	BaseParam
	Params AddSNSDingTalkAtPersonDetailParam `json:"params"`
}
