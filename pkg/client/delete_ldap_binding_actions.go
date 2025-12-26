// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteLdapBinding deletes LdapBinding
func (cli *ZSClient) DeleteLdapBinding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ldap/bindings/{uuid}", uuid, string(deleteMode))
}
