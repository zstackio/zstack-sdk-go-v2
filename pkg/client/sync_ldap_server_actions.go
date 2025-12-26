// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncLdapServer operates on SyncLdapServer
func (cli *ZSClient) SyncLdapServer(uuid string, params param.SyncLdapServerParam) (*view.SyncLdapServerEventView, error) {
	resp := view.SyncLdapServerEventView{}
	if err := cli.Put("v1/ldap/servers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
