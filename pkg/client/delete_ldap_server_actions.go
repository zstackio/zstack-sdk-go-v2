// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteLdapServer deletes LdapServer
func (cli *ZSClient) DeleteLdapServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ldap/servers/{uuid}", uuid, string(deleteMode))
}
