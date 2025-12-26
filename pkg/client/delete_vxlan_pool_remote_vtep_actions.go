// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVxlanPoolRemoteVtep deletes VxlanPoolRemoteVtep
func (cli *ZSClient) DeleteVxlanPoolRemoteVtep(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}/delete/remote-vtep-ip", uuid, string(deleteMode))
}
