// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateVpcSharedQosParamDetail CreateVpcSharedQos detail param
type CreateVpcSharedQosParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VpcUuid string `json:"vpcUuid" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Bandwidth int64 `json:"bandwidth,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcSharedQosParam CreateVpcSharedQos request param
type CreateVpcSharedQosParam struct {
	BaseParam
	CreateVpcSharedQos CreateVpcSharedQosParamDetail `json:"createVpcSharedQos"`
}
// UpdateVpcSharedQosParamDetail UpdateVpcSharedQos detail param
type UpdateVpcSharedQosParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

// UpdateVpcSharedQosParam UpdateVpcSharedQos request param
type UpdateVpcSharedQosParam struct {
	BaseParam
	UpdateVpcSharedQos UpdateVpcSharedQosParamDetail `json:"updateVpcSharedQos"`
}
// DeleteVpcSharedQosParamDetail DeleteVpcSharedQos detail param
type DeleteVpcSharedQosParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVpcSharedQosParam DeleteVpcSharedQos request param
type DeleteVpcSharedQosParam struct {
	BaseParam
	DeleteVpcSharedQos DeleteVpcSharedQosParamDetail `json:"deleteVpcSharedQos"`
}
