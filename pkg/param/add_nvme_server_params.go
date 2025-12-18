// Copyright (c) ZStack.io, Inc.

package param

// AddNvmeServerDetailParam AddNvmeServer detail param
type AddNvmeServerDetailParam struct {
	Name string `json:"name,omitempty"`
	Ip string `json:"ip" validate:"required"`
	Port int `json:"port,omitempty"`
	Transport string `json:"transport" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddNvmeServerParam AddNvmeServer request param
type AddNvmeServerParam struct {
	BaseParam
	Params AddNvmeServerDetailParam `json:"params"`
}
