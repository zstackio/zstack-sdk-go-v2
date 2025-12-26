// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLdapEntry gets LdapEntry by uuid
func (cli *ZSClient) GetLdapEntry(uuid string) (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.Get("v1/ldap/entry", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
