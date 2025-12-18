// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLdapServer 查询LdapServer列表
func (cli *ZSClient) QueryLdapServer(params param.QueryParam) ([]view.QueryLdapServerView, error) {
	var resp []view.QueryLdapServerView
	return resp, cli.List("v1/ldap/servers", &params, &resp)
}

