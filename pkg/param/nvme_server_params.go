// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteNvmeServerParamDetail DeleteNvmeServer detail param
type DeleteNvmeServerParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteNvmeServerParam DeleteNvmeServer request param
type DeleteNvmeServerParam struct {
	BaseParam
	Params DeleteNvmeServerParamDetail `json:"params"`
}
// AddNvmeServerParamDetail AddNvmeServer detail param
type AddNvmeServerParamDetail struct {
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
	Params AddNvmeServerParamDetail `json:"params"`
}
