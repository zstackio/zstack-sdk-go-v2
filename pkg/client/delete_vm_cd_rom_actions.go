// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVmCdRom deletes VmCdRom
func (cli *ZSClient) DeleteVmCdRom(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/cdroms/{uuid}", uuid, string(deleteMode))
}
