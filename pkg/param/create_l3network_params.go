// Copyright (c) ZStack.io, Inc.

package param

// CreateL3NetworkDetailParam CreateL3Network detail param
type CreateL3NetworkDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	L2NetworkUuid string `json:"l2NetworkUuid" validate:"required"`
	Category string `json:"category,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	System bool `json:"system,omitempty"`
	DnsDomain string `json:"dnsDomain,omitempty"`
	EnableIPAM bool `json:"enableIPAM,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL3NetworkParam CreateL3Network request param
type CreateL3NetworkParam struct {
	BaseParam
	Params CreateL3NetworkDetailParam `json:"params"`
}
