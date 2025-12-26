// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// SdnControllerRemoveHost operates on SdnControllerRemoveHost
func (cli *ZSClient) SdnControllerRemoveHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}", uuid, string(deleteMode))
}
