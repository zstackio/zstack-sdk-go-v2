// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelServicesDetailParam DeleteModelServices detail param
type DeleteModelServicesDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteModelServicesParam DeleteModelServices request param
type DeleteModelServicesParam struct {
	BaseParam
	Params DeleteModelServicesDetailParam `json:"params"`
}
