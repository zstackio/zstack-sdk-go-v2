// Copyright (c) ZStack.io, Inc.

package param

// AddSdnControllerDetailParam AddSdnController detail param
type AddSdnControllerDetailParam struct {
	VendorType string `json:"vendorType" validate:"required"`
	VendorVersion string `json:"vendorVersion,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Ip string `json:"ip" validate:"required"`
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSdnControllerParam AddSdnController request param
type AddSdnControllerParam struct {
	BaseParam
	Params AddSdnControllerDetailParam `json:"params"`
}
