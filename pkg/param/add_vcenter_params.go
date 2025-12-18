// Copyright (c) ZStack.io, Inc.

package param

// AddVCenterDetailParam AddVCenter detail param
type AddVCenterDetailParam struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Https bool `json:"https,omitempty"`
	Port int `json:"port,omitempty"`
	DomainName string `json:"domainName" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddVCenterParam AddVCenter request param
type AddVCenterParam struct {
	BaseParam
	Params AddVCenterDetailParam `json:"params"`
}
