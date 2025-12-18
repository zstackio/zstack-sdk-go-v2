// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SdnControllerRemoveHost operates on SdnControllerRemoveHost
func (cli *ZSClient) SdnControllerRemoveHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}", uuid, string(deleteMode))
}
