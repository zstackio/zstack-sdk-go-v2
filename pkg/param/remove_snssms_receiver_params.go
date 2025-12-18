// Copyright (c) ZStack.io, Inc.

package param

// RemoveSNSSmsReceiverDetailParam RemoveSNSSmsReceiver detail param
type RemoveSNSSmsReceiverDetailParam struct {
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	PhoneNumberList []string `json:"phoneNumberList,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveSNSSmsReceiverParam RemoveSNSSmsReceiver request param
type RemoveSNSSmsReceiverParam struct {
	BaseParam
	Params RemoveSNSSmsReceiverDetailParam `json:"params"`
}
