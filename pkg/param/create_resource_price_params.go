// Copyright (c) ZStack.io, Inc.

package param

// CreateResourcePriceDetailParam CreateResourcePrice detail param
type CreateResourcePriceDetailParam struct {
	ResourceName string `json:"resourceName" validate:"required"`
	ResourceUnit string `json:"resourceUnit,omitempty"`
	TimeUnit string `json:"timeUnit" validate:"required"`
	Price float64 `json:"price" validate:"required"`
	AccountUuid string `json:"accountUuid,omitempty"`
	DateInLong int64 `json:"dateInLong,omitempty"`
	TableUuid string `json:"tableUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateResourcePriceParam CreateResourcePrice request param
type CreateResourcePriceParam struct {
	BaseParam
	Params CreateResourcePriceDetailParam `json:"params"`
}
