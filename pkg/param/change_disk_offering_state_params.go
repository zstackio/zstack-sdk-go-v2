// Copyright (c) ZStack.io, Inc.

package param

// ChangeDiskOfferingStateDetailParam ChangeDiskOfferingState detail param
type ChangeDiskOfferingStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeDiskOfferingStateParam ChangeDiskOfferingState request param
type ChangeDiskOfferingStateParam struct {
	BaseParam
	Params ChangeDiskOfferingStateDetailParam `json:"params"`
}
