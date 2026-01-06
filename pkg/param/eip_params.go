// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateEipParamDetail CreateEip detail param
type CreateEipParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VipUuid string `json:"vipUuid" validate:"required"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	UsedIpUuid string `json:"usedIpUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEipParam CreateEip request param
type CreateEipParam struct {
	BaseParam
	Params CreateEipParamDetail `json:"params"`
}
// AttachEipParamDetail AttachEip detail param
type AttachEipParamDetail struct {
	EipUuid string `json:"eipUuid" validate:"required"`
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	UsedIpUuid string `json:"usedIpUuid,omitempty"`
}

// AttachEipParam AttachEip request param
type AttachEipParam struct {
	BaseParam
	Params AttachEipParamDetail `json:"params"`
}
// UpdateEipParamDetail UpdateEip detail param
type UpdateEipParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateEipParam UpdateEip request param
type UpdateEipParam struct {
	BaseParam
	Params UpdateEipParamDetail `json:"params"`
}
// DeleteEipParamDetail DeleteEip detail param
type DeleteEipParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEipParam DeleteEip request param
type DeleteEipParam struct {
	BaseParam
	Params DeleteEipParamDetail `json:"params"`
}
// DetachEipParamDetail DetachEip detail param
type DetachEipParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DetachEipParam DetachEip request param
type DetachEipParam struct {
	BaseParam
	Params DetachEipParamDetail `json:"params"`
}
