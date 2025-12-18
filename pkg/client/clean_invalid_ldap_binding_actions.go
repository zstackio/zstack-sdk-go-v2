// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CleanInvalidLdapBinding operates on CleanInvalidLdapBinding
func (cli *ZSClient) CleanInvalidLdapBinding(uuid string, params param.CleanInvalidLdapBindingParam) (*view.CleanInvalidLdapBindingEventView, error) {
	resp := view.CleanInvalidLdapBindingEventView{}
	if err := cli.Put("v1/ldap/bindings/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
