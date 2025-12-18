// Copyright (c) ZStack.io, Inc.

package param

// CreateSecurityGroupDetailParam CreateSecurityGroup detail param
type CreateSecurityGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSecurityGroupParam CreateSecurityGroup request param
type CreateSecurityGroupParam struct {
	BaseParam
	Params CreateSecurityGroupDetailParam `json:"params"`
}
