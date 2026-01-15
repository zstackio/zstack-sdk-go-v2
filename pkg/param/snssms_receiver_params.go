// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddSNSSmsReceiverParamDetail AddSNSSmsReceiver detail param
type AddSNSSmsReceiverParamDetail struct {
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
	Params AddSNSSmsReceiverParamDetail `json:"addSNSSmsReceiver"`
}
// RemoveSNSSmsReceiverParamDetail RemoveSNSSmsReceiver detail param
type RemoveSNSSmsReceiverParamDetail struct {
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	PhoneNumberList []string `json:"phoneNumberList,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveSNSSmsReceiverParam RemoveSNSSmsReceiver request param
type RemoveSNSSmsReceiverParam struct {
	BaseParam
	Params RemoveSNSSmsReceiverParamDetail `json:"removeSNSSmsReceiver"`
}
