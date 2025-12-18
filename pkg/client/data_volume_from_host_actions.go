// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachDataVolumeFromHost 操作DataVolumeFromHost
func (cli *ZSClient) DetachDataVolumeFromHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{volumeUuid}/hosts", uuid, string(deleteMode))
}

