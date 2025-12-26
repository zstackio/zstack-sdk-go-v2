// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateLdapEntryForBinding gets CandidateLdapEntryForBinding by uuid
func (cli *ZSClient) GetCandidateLdapEntryForBinding(uuid string) (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.Get("v1/ldap/entries/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
