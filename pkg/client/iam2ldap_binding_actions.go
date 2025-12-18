// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2LdapBinding 查询IAM2LdapBinding列表
func (cli *ZSClient) QueryIAM2LdapBinding(params param.QueryParam) ([]view.QueryIAM2LdapBindingView, error) {
	var resp []view.QueryIAM2LdapBindingView
	return resp, cli.List("v1/iam2/ldap/bindings", &params, &resp)
}

