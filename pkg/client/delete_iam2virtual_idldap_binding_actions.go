// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteIAM2VirtualIDLdapBinding deletes IAM2VirtualIDLdapBinding
func (cli *ZSClient) DeleteIAM2VirtualIDLdapBinding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/ldap/bindings/{uuid}", uuid, string(deleteMode))
}
