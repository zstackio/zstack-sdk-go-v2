// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIAM2VirtualIDLdapBinding deletes IAM2VirtualIDLdapBinding
func (cli *ZSClient) DeleteIAM2VirtualIDLdapBinding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/ldap/bindings/{uuid}", uuid, string(deleteMode))
}
