// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryLdapServer queries LdapServer list
func (cli *ZSClient) QueryLdapServer(params *param.QueryParam) ([]view.LdapServerInventoryView, error) {
	var resp []view.LdapServerInventoryView
	return resp, cli.List("v1/ldap/servers", params, &resp)
}
