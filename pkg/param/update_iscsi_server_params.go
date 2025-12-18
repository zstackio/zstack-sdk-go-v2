// Copyright (c) ZStack.io, Inc.

package param

// UpdateIscsiServerDetailParam UpdateIscsiServer detail param
type UpdateIscsiServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	ChapUserName string `json:"chapUserName,omitempty"`
	ChapUserPassword string `json:"chapUserPassword,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateIscsiServerParam UpdateIscsiServer request param
type UpdateIscsiServerParam struct {
	BaseParam
	Params UpdateIscsiServerDetailParam `json:"params"`
}
