// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLdapServerAvailableAttributes gets LdapServerAvailableAttributes by uuid
func (cli *ZSClient) GetLdapServerAvailableAttributes(uuid string) (*view.GetLdapServerAvailableAttributesView, error) {
	var resp view.GetLdapServerAvailableAttributesView
	if err := cli.Get("v1/ldap/server/attributes/{uuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
