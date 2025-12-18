// Copyright (c) ZStack.io, Inc.

package param

// GetLdapServerAvailableAttributesDetailParam GetLdapServerAvailableAttributes detail param
type GetLdapServerAvailableAttributesDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetLdapServerAvailableAttributesParam GetLdapServerAvailableAttributes request param
type GetLdapServerAvailableAttributesParam struct {
	BaseParam
	Params GetLdapServerAvailableAttributesDetailParam `json:"params"`
}
