// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// RebootBaremetalInstanceParamDetail RebootBaremetalInstance detail param
type RebootBaremetalInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	PxeBoot bool `json:"pxeBoot,omitempty"`
}

// RebootBaremetalInstanceParam RebootBaremetalInstance request param
type RebootBaremetalInstanceParam struct {
	BaseParam
	Params RebootBaremetalInstanceParamDetail `json:"params"`
}
// StartBaremetalInstanceParamDetail StartBaremetalInstance detail param
type StartBaremetalInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	PxeBoot bool `json:"pxeBoot,omitempty"`
}

// StartBaremetalInstanceParam StartBaremetalInstance request param
type StartBaremetalInstanceParam struct {
	BaseParam
	Params StartBaremetalInstanceParamDetail `json:"params"`
}
// CreateBaremetalInstanceParamDetail CreateBaremetalInstance detail param
type CreateBaremetalInstanceParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ChassisUuid string `json:"chassisUuid" validate:"required"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	TemplateUuid string `json:"templateUuid,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password" validate:"required"`
	NicCfgs map[string]string `json:"nicCfgs,omitempty"`
	BondingCfgs map[string]string `json:"bondingCfgs,omitempty"`
	CustomConfigurations map[string]string `json:"customConfigurations,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBaremetalInstanceParam CreateBaremetalInstance request param
type CreateBaremetalInstanceParam struct {
	BaseParam
	Params CreateBaremetalInstanceParamDetail `json:"params"`
}
// DestroyBaremetalInstanceParamDetail DestroyBaremetalInstance detail param
type DestroyBaremetalInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DestroyBaremetalInstanceParam DestroyBaremetalInstance request param
type DestroyBaremetalInstanceParam struct {
	BaseParam
	Params DestroyBaremetalInstanceParamDetail `json:"params"`
}
// ExpungeBaremetalInstanceParamDetail ExpungeBaremetalInstance detail param
type ExpungeBaremetalInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExpungeBaremetalInstanceParam ExpungeBaremetalInstance request param
type ExpungeBaremetalInstanceParam struct {
	BaseParam
	Params ExpungeBaremetalInstanceParamDetail `json:"params"`
}
// UpdateBaremetalInstanceParamDetail UpdateBaremetalInstance detail param
type UpdateBaremetalInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Password string `json:"password,omitempty"`
	Platform string `json:"platform,omitempty"`
}

// UpdateBaremetalInstanceParam UpdateBaremetalInstance request param
type UpdateBaremetalInstanceParam struct {
	BaseParam
	Params UpdateBaremetalInstanceParamDetail `json:"params"`
}
// StopBaremetalInstanceParamDetail StopBaremetalInstance detail param
type StopBaremetalInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Type string `json:"type,omitempty"`
}

// StopBaremetalInstanceParam StopBaremetalInstance request param
type StopBaremetalInstanceParam struct {
	BaseParam
	Params StopBaremetalInstanceParamDetail `json:"params"`
}
// RecoverBaremetalInstanceParamDetail RecoverBaremetalInstance detail param
type RecoverBaremetalInstanceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RecoverBaremetalInstanceParam RecoverBaremetalInstance request param
type RecoverBaremetalInstanceParam struct {
	BaseParam
	Params RecoverBaremetalInstanceParamDetail `json:"params"`
}
