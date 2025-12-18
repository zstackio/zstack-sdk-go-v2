// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CleanInvalidLdapIAM2Binding 操作CleanInvalidLdapIAM2Binding
func (cli *ZSClient) CleanInvalidLdapIAM2Binding(uuid string, params param.CleanInvalidLdapIAM2BindingParam) (*view.CleanInvalidLdapIAM2BindingEventView, error) {
	resp := view.CleanInvalidLdapIAM2BindingEventView{}
	if err := cli.Put("v1/iam2/ldap/bindings/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

