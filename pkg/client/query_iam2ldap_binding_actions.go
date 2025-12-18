// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2LdapBinding queries IAM2LdapBinding list
func (cli *ZSClient) QueryIAM2LdapBinding(params param.QueryParam) ([]view.LdapResourceRefInventoryView, error) {
	var resp []view.LdapResourceRefInventoryView
	return resp, cli.List("v1/iam2/ldap/bindings", &params, &resp)
}
