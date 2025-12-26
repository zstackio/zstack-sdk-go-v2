// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIAM2LdapBinding queries IAM2LdapBinding list
func (cli *ZSClient) QueryIAM2LdapBinding(params *param.QueryParam) ([]view.LdapResourceRefInventoryView, error) {
	var resp []view.LdapResourceRefInventoryView
	return resp, cli.List("v1/iam2/ldap/bindings", params, &resp)
}
