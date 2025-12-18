// Copyright (c) ZStack.io, Inc.

package param

// CreateVRouterOspfAreaDetailParam CreateVRouterOspfArea detail param
type CreateVRouterOspfAreaDetailParam struct {
	AreaId string `json:"areaId" validate:"required"`
	AreaAuth string `json:"areaAuth,omitempty"`
	AreaType string `json:"areaType,omitempty"`
	Password string `json:"password,omitempty"`
	KeyId int `json:"keyId,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVRouterOspfAreaParam CreateVRouterOspfArea request param
type CreateVRouterOspfAreaParam struct {
	BaseParam
	Params CreateVRouterOspfAreaDetailParam `json:"params"`
}
