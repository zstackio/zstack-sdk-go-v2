// Copyright (c) ZStack.io, Inc.

package param

// AddIscsiServerDetailParam AddIscsiServer detail param
type AddIscsiServerDetailParam struct {
	Name string `json:"name,omitempty"`
	Ip string `json:"ip" validate:"required"`
	Port int `json:"port,omitempty"`
	ChapUserName string `json:"chapUserName,omitempty"`
	ChapUserPassword string `json:"chapUserPassword,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIscsiServerParam AddIscsiServer request param
type AddIscsiServerParam struct {
	BaseParam
	Params AddIscsiServerDetailParam `json:"params"`
}
