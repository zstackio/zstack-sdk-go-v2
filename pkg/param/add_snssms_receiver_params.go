// Copyright (c) ZStack.io, Inc.

package param

// AddSNSSmsReceiverDetailParam AddSNSSmsReceiver detail param
type AddSNSSmsReceiverDetailParam struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	PhoneNumberList []string `json:"phoneNumberList,omitempty"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSNSSmsReceiverParam AddSNSSmsReceiver request param
type AddSNSSmsReceiverParam struct {
	BaseParam
	Params AddSNSSmsReceiverDetailParam `json:"params"`
}
