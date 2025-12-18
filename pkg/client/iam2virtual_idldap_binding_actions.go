// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateIAM2VirtualIDLdapBinding 创建IAM2VirtualIDLdapBinding
func (cli *ZSClient) CreateIAM2VirtualIDLdapBinding(params param.CreateIAM2VirtualIDLdapBindingParam) (*view.CreateIAM2VirtualIDLdapBindingEventView, error) {
	resp := view.CreateIAM2VirtualIDLdapBindingEventView{}
	if err := cli.Post("v1/iam2/ldap/bindings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

