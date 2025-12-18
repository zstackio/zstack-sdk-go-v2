// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsInstanceDetailParam UpdateEcsInstance detail param
type UpdateEcsInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Password string `json:"password,omitempty"`
}

// UpdateEcsInstanceParam UpdateEcsInstance request param
type UpdateEcsInstanceParam struct {
	BaseParam
	Params UpdateEcsInstanceDetailParam `json:"params"`
}
