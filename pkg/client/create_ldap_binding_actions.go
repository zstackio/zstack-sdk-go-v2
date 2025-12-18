// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateLdapBinding creates LdapBinding
func (cli *ZSClient) CreateLdapBinding(params param.CreateLdapBindingParam) (*view.CreateLdapBindingEventView, error) {
	resp := view.CreateLdapBindingEventView{}
	if err := cli.Post("v1/ldap/bindings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
