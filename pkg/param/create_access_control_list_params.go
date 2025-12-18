// Copyright (c) ZStack.io, Inc.

package param

// CreateAccessControlListDetailParam CreateAccessControlList detail param
type CreateAccessControlListDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAccessControlListParam CreateAccessControlList request param
type CreateAccessControlListParam struct {
	BaseParam
	Params CreateAccessControlListDetailParam `json:"params"`
}
