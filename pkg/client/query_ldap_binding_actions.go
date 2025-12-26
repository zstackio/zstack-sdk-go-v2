// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryLdapBinding queries LdapBinding list
func (cli *ZSClient) QueryLdapBinding(params *param.QueryParam) ([]view.LdapAccountRefInventoryView, error) {
	var resp []view.LdapAccountRefInventoryView
	return resp, cli.List("v1/ldap/bindings", params, &resp)
}
