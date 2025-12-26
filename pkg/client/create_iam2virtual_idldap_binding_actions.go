// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateIAM2VirtualIDLdapBinding creates IAM2VirtualIDLdapBinding
func (cli *ZSClient) CreateIAM2VirtualIDLdapBinding(params param.CreateIAM2VirtualIDLdapBindingParam) (*view.CreateIAM2VirtualIDLdapBindingEventView, error) {
	resp := view.CreateIAM2VirtualIDLdapBindingEventView{}
	if err := cli.Post("v1/iam2/ldap/bindings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
