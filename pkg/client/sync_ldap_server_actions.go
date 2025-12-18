// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncLdapServer 操作SyncLdapServer
func (cli *ZSClient) SyncLdapServer(uuid string, params param.SyncLdapServerParam) (*view.SyncLdapServerEventView, error) {
	resp := view.SyncLdapServerEventView{}
	if err := cli.Put("v1/ldap/servers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

