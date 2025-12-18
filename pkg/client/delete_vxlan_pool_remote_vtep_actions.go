// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVxlanPoolRemoteVtep deletes VxlanPoolRemoteVtep
func (cli *ZSClient) DeleteVxlanPoolRemoteVtep(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}/delete/remote-vtep-ip", uuid, string(deleteMode))
}
