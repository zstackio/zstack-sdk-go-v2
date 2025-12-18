// Copyright (c) ZStack.io, Inc.

package param

// CreateL2TfNetworkDetailParam CreateL2TfNetwork detail param
type CreateL2TfNetworkDetailParam struct {
	IpPrefix string `json:"ipPrefix,omitempty"`
	IpPrefixLength int `json:"ipPrefixLength,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	Isolated bool `json:"isolated,omitempty"`
	Pvlan string `json:"pvlan,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2TfNetworkParam CreateL2TfNetwork request param
type CreateL2TfNetworkParam struct {
	BaseParam
	Params CreateL2TfNetworkDetailParam `json:"params"`
}
