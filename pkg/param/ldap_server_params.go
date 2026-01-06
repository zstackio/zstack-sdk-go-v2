// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddLdapServerParamDetail AddLdapServer detail param
type AddLdapServerParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Url string `json:"url" validate:"required"`
	Base string `json:"base" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Encryption string `json:"encryption" validate:"required"`
	Scope string `json:"scope" validate:"required"`
}

// AddLdapServerParam AddLdapServer request param
type AddLdapServerParam struct {
	BaseParam
	Params AddLdapServerParamDetail `json:"params"`
}
// SyncLdapServerParamDetail SyncLdapServer detail param
type SyncLdapServerParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SyncLdapServerParam SyncLdapServer request param
type SyncLdapServerParam struct {
	BaseParam
	Params SyncLdapServerParamDetail `json:"params"`
}
// DeleteLdapServerParamDetail DeleteLdapServer detail param
type DeleteLdapServerParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteLdapServerParam DeleteLdapServer request param
type DeleteLdapServerParam struct {
	BaseParam
	Params DeleteLdapServerParamDetail `json:"params"`
}
// UpdateLdapServerParamDetail UpdateLdapServer detail param
type UpdateLdapServerParamDetail struct {
	LdapServerUuid string `json:"ldapServerUuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
	Base string `json:"base,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Encryption string `json:"encryption,omitempty"`
}

// UpdateLdapServerParam UpdateLdapServer request param
type UpdateLdapServerParam struct {
	BaseParam
	Params UpdateLdapServerParamDetail `json:"params"`
}
