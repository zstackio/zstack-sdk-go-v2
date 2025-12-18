// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteLdapBinding deletes LdapBinding
func (cli *ZSClient) DeleteLdapBinding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ldap/bindings/{uuid}", uuid, string(deleteMode))
}
