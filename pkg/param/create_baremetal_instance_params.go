// Copyright (c) ZStack.io, Inc.

package param

// CreateBaremetalInstanceDetailParam CreateBaremetalInstance detail param
type CreateBaremetalInstanceDetailParam struct {
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
	Params CreateBaremetalInstanceDetailParam `json:"params"`
}
