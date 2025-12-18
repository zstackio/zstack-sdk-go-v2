// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateLdapServer updates LdapServer
func (cli *ZSClient) UpdateLdapServer(uuid string, params param.UpdateLdapServerParam) (*view.UpdateLdapServerEventView, error) {
	resp := view.UpdateLdapServerEventView{}
	if err := cli.Put("v1/ldap/servers/{ldapServerUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
