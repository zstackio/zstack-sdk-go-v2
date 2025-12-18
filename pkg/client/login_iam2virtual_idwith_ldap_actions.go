// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LoginIAM2VirtualIDWithLdap 操作LoginIAM2VirtualIDWithLdap
func (cli *ZSClient) LoginIAM2VirtualIDWithLdap(uuid string, params param.LoginIAM2VirtualIDWithLdapParam) (*view.LoginIAM2VirtualIDWithLdapView, error) {
	resp := view.LoginIAM2VirtualIDWithLdapView{}
	if err := cli.Put("v1/iam2/login/virtual-ids/ldap", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

