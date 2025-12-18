// Copyright (c) ZStack.io, Inc.

package param

// UpdateSlbGroupDetailParam UpdateSlbGroup detail param
type UpdateSlbGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	SlbOfferingUuid string `json:"slbOfferingUuid,omitempty"`
}

// UpdateSlbGroupParam UpdateSlbGroup request param
type UpdateSlbGroupParam struct {
	BaseParam
	Params UpdateSlbGroupDetailParam `json:"params"`
}
