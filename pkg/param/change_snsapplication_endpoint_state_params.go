// Copyright (c) ZStack.io, Inc.

package param

// ChangeSNSApplicationEndpointStateDetailParam ChangeSNSApplicationEndpointState detail param
type ChangeSNSApplicationEndpointStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSNSApplicationEndpointStateParam ChangeSNSApplicationEndpointState request param
type ChangeSNSApplicationEndpointStateParam struct {
	BaseParam
	Params ChangeSNSApplicationEndpointStateDetailParam `json:"params"`
}
