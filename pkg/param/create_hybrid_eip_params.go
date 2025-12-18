// Copyright (c) ZStack.io, Inc.

package param

// CreateHybridEipDetailParam CreateHybridEip detail param
type CreateHybridEipDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	BandWidthMb int64 `json:"bandWidthMb" validate:"required"`
	Type string `json:"type" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ChargeType string `json:"chargeType" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateHybridEipParam CreateHybridEip request param
type CreateHybridEipParam struct {
	BaseParam
	Params CreateHybridEipDetailParam `json:"params"`
}
