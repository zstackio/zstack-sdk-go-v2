// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachDataVolumeFromVm operates on DataVolumeFromVm
func (cli *ZSClient) DetachDataVolumeFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{uuid}/vm-instances", uuid, string(deleteMode))
}
