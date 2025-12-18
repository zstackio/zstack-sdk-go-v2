// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddLdapServer adds LdapServer
func (cli *ZSClient) AddLdapServer(params param.AddLdapServerParam) (*view.AddLdapServerEventView, error) {
	resp := view.AddLdapServerEventView{}
	if err := cli.Post("v1/ldap/servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
