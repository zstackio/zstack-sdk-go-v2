// Copyright (c) ZStack.io, Inc.

package param

// CreateVniRangeDetailParam CreateVniRange detail param
type CreateVniRangeDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	StartVni int `json:"startVni" validate:"required"`
	EndVni int `json:"endVni" validate:"required"`
	L2NetworkUuid string `json:"l2NetworkUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVniRangeParam CreateVniRange request param
type CreateVniRangeParam struct {
	BaseParam
	Params CreateVniRangeDetailParam `json:"params"`
}
