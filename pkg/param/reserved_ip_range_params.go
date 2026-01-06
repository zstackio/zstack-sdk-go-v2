// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteReservedIpRangeParamDetail DeleteReservedIpRange detail param
type DeleteReservedIpRangeParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteReservedIpRangeParam DeleteReservedIpRange request param
type DeleteReservedIpRangeParam struct {
	BaseParam
	Params DeleteReservedIpRangeParamDetail `json:"params"`
}
// AddReservedIpRangeParamDetail AddReservedIpRange detail param
type AddReservedIpRangeParamDetail struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	StartIp string `json:"startIp" validate:"required"`
	EndIp string `json:"endIp" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddReservedIpRangeParam AddReservedIpRange request param
type AddReservedIpRangeParam struct {
	BaseParam
	Params AddReservedIpRangeParamDetail `json:"params"`
}
