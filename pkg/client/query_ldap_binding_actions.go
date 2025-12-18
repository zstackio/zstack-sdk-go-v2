// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLdapBinding queries LdapBinding list
func (cli *ZSClient) QueryLdapBinding(params param.QueryParam) ([]view.LdapAccountRefInventoryView, error) {
	var resp []view.LdapAccountRefInventoryView
	return resp, cli.List("v1/ldap/bindings", &params, &resp)
}
