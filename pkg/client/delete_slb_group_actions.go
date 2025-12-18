// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteSlbGroup deletes SlbGroup
func (cli *ZSClient) DeleteSlbGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/slb/group/{uuid}", uuid, string(deleteMode))
}
