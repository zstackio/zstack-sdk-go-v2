// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddIpRangeParamDetail AddIpRange detail param
type AddIpRangeParamDetail struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	StartIp string `json:"startIp" validate:"required"`
	EndIp string `json:"endIp" validate:"required"`
	Netmask string `json:"netmask" validate:"required"`
	Gateway string `json:"gateway,omitempty"`
	IpRangeType string `json:"ipRangeType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIpRangeParam AddIpRange request param
type AddIpRangeParam struct {
	BaseParam
	Params AddIpRangeParamDetail `json:"params"`
}
// UpdateIpRangeParamDetail UpdateIpRange detail param
type UpdateIpRangeParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateIpRangeParam UpdateIpRange request param
type UpdateIpRangeParam struct {
	BaseParam
	Params UpdateIpRangeParamDetail `json:"params"`
}
// DeleteIpRangeParamDetail DeleteIpRange detail param
type DeleteIpRangeParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteIpRangeParam DeleteIpRange request param
type DeleteIpRangeParam struct {
	BaseParam
	Params DeleteIpRangeParamDetail `json:"params"`
}
