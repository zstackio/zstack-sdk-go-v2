// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateIAM2VirtualIDGroupParamDetail CreateIAM2VirtualIDGroup detail param
type CreateIAM2VirtualIDGroupParamDetail struct {
	ProjectUuid *string `json:"projectUuid,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Attributes []AttributeParam `json:"attributes,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDGroupParam CreateIAM2VirtualIDGroup request param
type CreateIAM2VirtualIDGroupParam struct {
	BaseParam
	Params CreateIAM2VirtualIDGroupParamDetail `json:"params"`
}
// DeleteIAM2VirtualIDGroupParamDetail DeleteIAM2VirtualIDGroup detail param
type DeleteIAM2VirtualIDGroupParamDetail struct {
}

// DeleteIAM2VirtualIDGroupParam DeleteIAM2VirtualIDGroup request param
type DeleteIAM2VirtualIDGroupParam struct {
	BaseParam
	Params DeleteIAM2VirtualIDGroupParamDetail `json:"deleteIAM2VirtualIDGroup"`
}
// UpdateIAM2VirtualIDGroupParamDetail UpdateIAM2VirtualIDGroup detail param
type UpdateIAM2VirtualIDGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateIAM2VirtualIDGroupParam UpdateIAM2VirtualIDGroup request param
type UpdateIAM2VirtualIDGroupParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDGroupParamDetail `json:"updateIAM2VirtualIDGroup"`
}
