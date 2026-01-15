// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateSecurityGroupParamDetail CreateSecurityGroup detail param
type CreateSecurityGroupParamDetail struct {
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
	CreateSecurityGroup CreateSecurityGroupParamDetail `json:"createSecurityGroup"`
}
// DeleteSecurityGroupParamDetail DeleteSecurityGroup detail param
type DeleteSecurityGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSecurityGroupParam DeleteSecurityGroup request param
type DeleteSecurityGroupParam struct {
	BaseParam
	DeleteSecurityGroup DeleteSecurityGroupParamDetail `json:"deleteSecurityGroup"`
}
// UpdateSecurityGroupParamDetail UpdateSecurityGroup detail param
type UpdateSecurityGroupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateSecurityGroupParam UpdateSecurityGroup request param
type UpdateSecurityGroupParam struct {
	BaseParam
	UpdateSecurityGroup UpdateSecurityGroupParamDetail `json:"updateSecurityGroup"`
}
