// Copyright (c) ZStack.io, Inc.

package param

// DeleteInstanceOfferingDetailParam DeleteInstanceOffering detail param
type DeleteInstanceOfferingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteInstanceOfferingParam DeleteInstanceOffering request param
type DeleteInstanceOfferingParam struct {
	BaseParam
	Params DeleteInstanceOfferingDetailParam `json:"params"`
}
