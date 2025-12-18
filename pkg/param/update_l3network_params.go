// Copyright (c) ZStack.io, Inc.

package param

// UpdateL3NetworkDetailParam UpdateL3Network detail param
type UpdateL3NetworkDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DnsDomain string `json:"dnsDomain,omitempty"`
	Category string `json:"category,omitempty"`
	System bool `json:"system,omitempty"`
}

// UpdateL3NetworkParam UpdateL3Network request param
type UpdateL3NetworkParam struct {
	BaseParam
	Params UpdateL3NetworkDetailParam `json:"params"`
}
